package packetfabric

const hostedIBMConnURI = "/v2/services/cloud/hosted/ibm"

type HostedIBMConn struct {
	IbmAccountID           string `json:"ibm_account_id,omitempty"`
	IbmBgpAsn              int    `json:"ibm_bgp_asn,omitempty"`
	IbmBgpCerCidr          string `json:"ibm_bgp_cer_cidr,omitempty"`
	IbmBgpIbmCidr          string `json:"ibm_bgp_ibm_cidr,omitempty"`
	Description            string `json:"description,omitempty"`
	AccountUUID            string `json:"account_uuid,omitempty"`
	Pop                    string `json:"pop,omitempty"`
	Port                   string `json:"port,omitempty"`
	Vlan                   int    `json:"vlan,omitempty"`
	SrcSvlan               int    `json:"src_svlan,omitempty"`
	Zone                   string `json:"zone,omitempty"`
	Speed                  string `json:"speed,omitempty"`
	PublishedQuoteLineUUID string `json:"published_quote_line_uuid,omitempty"`
	PONumber               string `json:"po_number,omitempty"`
}

func (c *PFClient) CreateHostedIBMConn(conn HostedIBMConn) (*CloudServiceConnCreateResp, error) {
	expectedResp := &CloudServiceConnCreateResp{}
	_, err := c.sendRequest(hostedIBMConnURI, postMethod, conn, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}
