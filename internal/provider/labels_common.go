package provider

import (
	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func createLabels(c *packetfabric.PFClient, circuitId string, labels interface{}) (diag.Diagnostics, bool) {
	var labelsData []string
	for _, label := range labels.([]interface{}) {
		labelsData = append(labelsData, label.(string))
	}
	labelPayload := packetfabric.LabelsPayload{Labels: labelsData}
	_, err := c.CreateLabel(circuitId, labelPayload)
	if err != nil {
		return diag.FromErr(err), false
	}
	return diag.Diagnostics{}, true
}

func updateLabels(c *packetfabric.PFClient, circuitId string, labels interface{}) (diag.Diagnostics, bool) {
	var labelsData []string
	for _, label := range labels.([]interface{}) {
		labelsData = append(labelsData, label.(string))
	}
	labelPayload := packetfabric.LabelsPayload{Labels: labelsData}
	_, err := c.UpdateLabel(circuitId, labelPayload)
	if err != nil {
		return diag.FromErr(err), false
	}
	return diag.Diagnostics{}, true
}

func getLabels(c *packetfabric.PFClient, circuitId string) (*packetfabric.LabelsResponse, error) {
	resp, err := c.GetLabels(circuitId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
