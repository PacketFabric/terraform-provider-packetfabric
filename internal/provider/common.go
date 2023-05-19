package provider

import (
	"fmt"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Labels
func createLabels(c *packetfabric.PFClient, circuitId string, labels interface{}) (diag.Diagnostics, bool) {
	var labelsData []string
	// Convert labels (TypeSet) to a list
	labelsList := labels.(*schema.Set).List()
	for _, label := range labelsList {
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
	// Convert labels (TypeSet) to a list
	labelsList := labels.(*schema.Set).List()
	for _, label := range labelsList {
		labelsData = append(labelsData, label.(string))
	}
	labelPayload := packetfabric.LabelsPayload{Labels: labelsData}
	_, err := c.UpdateLabel(circuitId, labelPayload)
	if err != nil {
		return diag.FromErr(err), false
	}
	return diag.Diagnostics{}, true
}

func getLabels(c *packetfabric.PFClient, circuitId string) ([]string, error) {
	resp, err := c.GetLabels(circuitId)
	if err != nil {
		return nil, err
	}
	return resp.Labels, nil
}

// ETA (Early Termination Liability)
func addETLWarning(c *packetfabric.PFClient, cID string) ([]diag.Diagnostic, error) {
	var diags []diag.Diagnostic
	etlCost, err := c.GetEarlyTerminationLiability(cID)
	if err != nil {
		return nil, err
	}
	if etlCost > 0 {
		etlWarning := fmt.Sprintf("Resource ID: %s - Early Termination Liability (ETL) cost: $%.2f", cID, etlCost)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  etlWarning,
		})
	}
	return diags, nil
}
