package packetfabric

import (
	"errors"
	"fmt"
)

const CloudProviderCredentialURI = "/v2/services/cloud/credentials"

// This struct represents a Cloud Provider Credential Create
// https://docs.packetfabric.com/api/v2/swagger/#/Cloud%20Provider%20Credentials/cloud_provider_credential_post
type CloudProviderCredentialCreate struct {
	CloudProvider    string           `json:"cloud_provider"`
	Description      string           `json:"description"`
	CloudCredentials CloudCredentials `json:"cloud_credentials"`
}

// This struct represents a Cloud Provider Credential Update
// https://docs.packetfabric.com/api/v2/swagger/#/Cloud%20Provider%20Credentials/cloud_provider_credential_update
type CloudProviderCredentialUpdate struct {
	Description      string           `json:"description"`
	CloudCredentials CloudCredentials `json:"cloud_credentials"`
}

type CloudCredentials struct {
	AWSAccessKey         string `json:"aws_access_key,omitempty"`
	AWSSecretKey         string `json:"aws_secret_key,omitempty"`
	GoogleServiceAccount string `json:"google_service_account,omitempty"`
}

// This struct represents a Cloud Provider Credential create response
type CloudProviderCredentialResponse struct {
	CloudProviderCredentialUUID string `json:"cloud_provider_credential_uuid"`
	Description                 string `json:"description"`
	CloudProvider               string `json:"cloud_provider"`
	IsUnused                    bool   `json:"is_unused"`
	TimeCreated                 string `json:"time_created"`
	TimeUpdated                 string `json:"time_updated"`
}

// This struct represents a Cloud Provider Credential delete response
type CloudProviderCredentialDelResp struct {
	Message string `json:"message"`
}

// This function represents the Action to create a new Cloud Provider Credential
func (c *PFClient) CreateCloudProviderCredential(creds CloudProviderCredentialCreate) (*CloudProviderCredentialResponse, error) {
	resp := &CloudProviderCredentialResponse{}
	_, err := c.sendRequest(CloudProviderCredentialURI, postMethod, creds, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action tp update an existing  Cloud Provider Credential
func (c *PFClient) UpdateCloudProviderCredential(creds CloudProviderCredentialUpdate, cpcID string) (*CloudProviderCredentialResponse, error) {
	formatedURI := fmt.Sprintf("%s/%s", CloudProviderCredentialURI, cpcID)
	resp := &CloudProviderCredentialResponse{}
	_, err := c.sendRequest(formatedURI, patchMethod, creds, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to Read an existing Cloud Provider Credential
func (c *PFClient) ReadCloudProviderCredential(cpcID string) (*CloudProviderCredentialResponse, error) {
	if cpcID == "" {
		return nil, errors.New(errorMsg)
	}
	resp := []*CloudProviderCredentialResponse{}
	_, err := c.sendRequest(CloudProviderCredentialURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	for _, credential := range resp {
		if credential.CloudProviderCredentialUUID == cpcID {
			return credential, nil
		}
	}
	return nil, fmt.Errorf("cloud provider credential %s not found", cpcID)
}

func (c *PFClient) ListCloudProviderCredentials() ([]CloudProviderCredentialResponse, error) {
	expectedResp := make([]CloudProviderCredentialResponse, 0)
	_, err := c.sendRequest(CloudProviderCredentialURI, getMethod, nil, &expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

// This function represents the Action to Delete an existing Cloud Provider Credential
func (c *PFClient) DeleteCloudProviderCredential(cpcID string) (*CloudProviderCredentialDelResp, error) {
	if cpcID == "" {
		return nil, errors.New(errorMsg)
	}
	formatedURI := fmt.Sprintf("%s/%s", CloudProviderCredentialURI, cpcID)
	resp := &CloudProviderCredentialDelResp{}
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
