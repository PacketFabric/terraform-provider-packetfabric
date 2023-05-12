package packetfabric

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const HostURL string = "http://localhost:9090"

const postMethod = "POST"
const getMethod = "GET"
const deleteMethod = "DELETE"
const patchMethod = "PATCH"
const putMethod = "PUT"

type PFClient struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Ctx        context.Context
	Username   string
	Password   string
}

type AuthResponse struct {
	Token       *string `json:"token"`
	TimeExpires string  `json:"time_expires"`
}

type PFAuth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func NewPFClient(host, token *string) (*PFClient, error) {
	c := _createBasicClient(host)

	if token != nil {
		c.Token = *token
	}
	return &c, nil
}

func NewPFClientByUserAndPass(ctx context.Context, host *string, username, password string) (*PFClient, error) {
	c := _createBasicClient(host)
	auth := PFAuth{Login: username, Password: password}
	loginResp := AuthResponse{}
	_, err := c.sendRequest("/v2/auth/login", postMethod, auth, &loginResp)
	if err != nil {
		return nil, err
	}
	c.Token = *loginResp.Token
	return &c, nil
}

func _createBasicClient(host *string) PFClient {
	c := PFClient{
		HTTPClient: &http.Client{Timeout: 60 * time.Second},
		HostURL:    HostURL,
	}
	if host != nil {
		c.HostURL = *host
	}
	return c
}

func (c *PFClient) _doRequest(req *http.Request, authToken *string, customHeaders ...map[string]string) (*http.Response, []byte, error) {
	token := c.Token
	if authToken != nil {
		token = *authToken
	}
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("pf-request-source", "terraform")

	if len(customHeaders) > 0 {
		headers := customHeaders[0]
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return res, nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res, nil, err
	}
	if res.StatusCode == http.StatusBadRequest ||
		res.StatusCode == http.StatusUnauthorized ||
		res.StatusCode == http.StatusNotFound {
		return res, nil, fmt.Errorf("Status: %d, body: %s", res.StatusCode, body)
	}
	return res, body, err
}

func (c *PFClient) sendRequest(uri, method string, payload interface{}, resp interface{}) (interface{}, error) {
	var req *http.Request
	var err error
	c.Ctx = context.Background()
	formatedURL := fmt.Sprintf("%s%s", c.HostURL, uri)
	switch method {
	case getMethod:
		req, _ = http.NewRequestWithContext(c.Ctx, method, formatedURL, nil)
	case postMethod, patchMethod, putMethod:
		rb, mErr := json.Marshal(payload)
		if mErr != nil {
			return nil, mErr
		}
		req, _ = http.NewRequestWithContext(c.Ctx, method, formatedURL, strings.NewReader(string(rb)))
	case deleteMethod:
		if payload != nil {
			rb, pErr := json.Marshal(payload)
			if pErr != nil {
				return nil, pErr
			}
			req, _ = http.NewRequestWithContext(c.Ctx, method, formatedURL, strings.NewReader(string(rb)))
		} else {
			req, _ = http.NewRequestWithContext(c.Ctx, method, formatedURL, nil)
		}
	}
	res, body, err := c._doRequest(req, &c.Token)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		err = json.Unmarshal(body, &resp)
		if err != nil {
			return nil, err
		}
	}
	c._logDebug(formatedURL, method, payload, resp, body)
	return res, nil
}

func (c *PFClient) sendMultipartRequest(uri, method, fileField, filePath string, payload map[string]string, resp interface{}) (interface{}, error) {
	var req *http.Request
	var err error
	c.Ctx = context.Background()
	formatedURL := fmt.Sprintf("%s%s", c.HostURL, uri)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)
	// Get the MIME type of the file based on the file extension
	mimeType := mime.TypeByExtension(filepath.Ext(filePath))

	// Create a custom part header with the correct MIME type
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, fileField, filepath.Base(filePath)))
	h.Set("Content-Type", mimeType)

	// Create the form file with the custom part header
	part, err := writer.CreatePart(h)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	for key, val := range payload {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, _ = http.NewRequestWithContext(c.Ctx, method, formatedURL, buffer)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, body, err := c._doRequest(req, &c.Token, map[string]string{})
	if err != nil {
		return nil, err
	}
	if resp != nil {
		err = json.Unmarshal(body, &resp)
		if err != nil {
			return nil, err
		}
	}
	c._logDebug(formatedURL, method, payload, resp, body)
	return res, nil
}

// For debug use only.
func (c *PFClient) _logDebug(url, method string, payload, resp interface{}, body []byte) {
	debug := make(map[string]interface{})
	debug["url"] = url
	if payload != nil {
		debug["payload"] = payload
	}
	if resp != nil {
		debug["resp"] = resp
	}
	if body != nil {
		debug["body"] = string(body)
	}
	tflog.Debug(c.Ctx, fmt.Sprintf("\n##[CLIENT | SEND_REQUEST]## SENDING %s REQUEST", method), debug)
}
