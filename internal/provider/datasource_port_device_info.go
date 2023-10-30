package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePortDeviceInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePortDeviceInfoRead,
		Schema: map[string]*schema.Schema{
			PfPortCircuitId:                    schemaStringRequiredNotEmpty(PfPortCircuitIdDescription2),
			PfAdjacentRouter:                   schemaStringComputed(PfAdjacentRouterDescription),
			PfDeviceName:                       schemaStringComputed(PfDeviceNameDescription),
			PfDeviceMake:                       schemaStringComputed(PfDeviceMakeDescription),
			PfAdminStatus:                      schemaStringComputed(PfAdminStatusDescription3),
			PfOperStatus:                       schemaStringComputed(PfOperStatusDescription),
			PfAutoNegotiation:                  schemaBoolComputed(PfAutoNegotiationDescription),
			PfIfaceName:                        schemaStringComputed(PfIfaceNameDescription),
			PfSpeed:                            schemaStringComputed(PfSpeedDescriptionC),
			PfOpticsDiagnosticsLaneTxPowerDbm:  schemaFloatComputed(PfOpticsDiagnosticsLaneTxPowerDbmDescription),
			PfOpticsDiagnosticsLaneTxPower:     schemaFloatComputed(PfOpticsDiagnosticsLaneTxPowerDescription),
			PfOpticsDiagnosticsLaneIndex:       schemaStringComputed(PfOpticsDiagnosticsLaneIndexDescription),
			PfOpticsDiagnosticsLaneRxPowerDbm:  schemaFloatComputed(PfOpticsDiagnosticsLaneRxPowerDbmDescription),
			PfOpticsDiagnosticsLaneRxPower:     schemaFloatComputed(PfOpticsDiagnosticsLaneRxPowerDescription),
			PfOpticsDiagnosticsLaneBiasCurrent: schemaFloatComputed(PfOpticsDiagnosticsLaneBiasCurrentDescription),
			PfOpticsDiagnosticsLaneTxStatus:    schemaStringComputed(PfOpticsDiagnosticsLaneTxStatusDescription),
			PfOpticsDiagnosticsLaneRxStatus:    schemaStringComputed(PfOpticsDiagnosticsLaneRxStatusDescription),
			PfPolltime:                         schemaIntComputed(PfPolltimeDescription),
			PfTimeFlapped:                      schemaStringComputed(PfTimeFlappedDescription),
			PfTrafficRxBps:                     schemaIntComputed(PfTrafficRxBpsDescription),
			PfTrafficRxBytes:                   schemaIntComputed(PfTrafficRxBytesDescription),
			PfTrafficRxIpv6Bytes:               schemaIntComputed(PfTrafficRxIpv6BytesDescription),
			PfTrafficRxIpv6Packets:             schemaIntComputed(PfTrafficRxIpv6PacketsDescription),
			PfTrafficRxPackets:                 schemaIntComputed(PfTrafficRxPacketsDescription),
			PfTrafficRxPps:                     schemaIntComputed(PfTrafficRxPpsDescription),
			PfTrafficTxBps:                     schemaIntComputed(PfTrafficTxBpsDescription),
			PfTrafficTxBytes:                   schemaIntComputed(PfTrafficTxBytesDescription),
			PfTrafficTxIpv6Bytes:               schemaIntComputed(PfTrafficTxIpv6BytesDescription),
			PfTrafficTxIpv6Packets:             schemaIntComputed(PfTrafficTxIpv6PacketsDescription),
			PfTrafficTxPackets:                 schemaIntComputed(PfTrafficTxBytesDescription),
			PfTrafficTxPps:                     schemaIntComputed(PfTrafficTxPpsDescription),
			PfWiringMedia:                      schemaStringComputed(PfWiringMediaDescription),
			PfWiringModule:                     schemaStringComputed(PfWiringModuleDescription),
			PfWiringPanel:                      schemaStringComputed(PfWiringPanelDescription),
			PfWiringPosition:                   schemaStringComputed(PfWiringPositionDescription),
			PfWiringReach:                      schemaStringComputed(PfWiringReachDescription),
			PfWiringType:                       schemaStringComputed(PfWiringTypeDescription),
			PfLagSpeed:                         schemaIntComputed(PfLagSpeedDescription),
			PfDeviceCanLag:                     schemaBoolComputed(PfDeviceCanLagDescription),
		},
	}
}

func dataSourcePortDeviceInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	portCID, ok := d.GetOk("port_circuit_id")
	if !ok {
		return diag.Errorf("please provide a valid port circuti ID")
	}
	portInfo, err := c.GetPortDeviceInfo(portCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	if portInfo != nil {
		_ = setResourceDataKeys(d, portInfo, PfDeviceName, PfDeviceMake, PfAdminStatus, PfOperStatus, PfAutoNegotiation, PfIfaceName, PfSpeed, PfTimeFlapped, PfTrafficRxBps, PfTrafficRxBytes, PfTrafficRxIpv6Bytes, PfTrafficRxIpv6Packets, PfTrafficRxPackets, PfTrafficRxPps, PfTrafficTxBps, PfTrafficTxBytes, PfTrafficTxIpv6Bytes, PfTrafficTxIpv6Packets, PfTrafficTxPackets, PfTrafficTxPps, PfWiringMedia, PfWiringModule, PfWiringPanel, PfWiringPosition, PfWiringReach, PfWiringType, PfLagSpeed, PfDeviceCanLag)
		if portInfo.Polltime != nil {
			_ = d.Set(PfPolltime, portInfo.Polltime) // constant: "polltime" -> PfPolltime
		}
		if portInfo.AdjacentRouter != nil {
			_ = d.Set(PfAdjacentRouter, portInfo.AdjacentRouter) // constant: "adjacent_router" -> PfAdjacentRouter
		}
		for _, optics := range portInfo.OpticsDiagnosticsLaneValues {
			_ = d.Set(PfOpticsDiagnosticsLaneTxPowerDbm, optics.TxPowerDbm) // constant: "optics_diagnostics_lane_tx_power_dbm" -> PfOpticsDiagnosticsLaneTxPowerDbm
			_ = d.Set(PfOpticsDiagnosticsLaneTxPower, optics.TxPower) // constant: "optics_diagnostics_lane_tx_power" -> PfOpticsDiagnosticsLaneTxPower
			_ = d.Set(PfOpticsDiagnosticsLaneIndex, optics.LaneIndex) // constant: "optics_diagnostics_lane_index" -> PfOpticsDiagnosticsLaneIndex
			_ = d.Set(PfOpticsDiagnosticsLaneRxPowerDbm, optics.RxPowerDbm) // constant: "optics_diagnostics_lane_rx_power_dbm" -> PfOpticsDiagnosticsLaneRxPowerDbm
			_ = d.Set(PfOpticsDiagnosticsLaneRxPower, optics.RxPower) // constant: "optics_diagnostics_lane_rx_power" -> PfOpticsDiagnosticsLaneRxPower
			_ = d.Set(PfOpticsDiagnosticsLaneBiasCurrent, optics.BiasCurrent) // constant: "optics_diagnostics_lane_bias_current" -> PfOpticsDiagnosticsLaneBiasCurrent
			_ = d.Set(PfOpticsDiagnosticsLaneTxStatus, optics.TxStatus) // constant: "optics_diagnostics_lane_tx_status" -> PfOpticsDiagnosticsLaneTxStatus
			_ = d.Set(PfOpticsDiagnosticsLaneRxStatus, optics.RxStatus) // constant: "optics_diagnostics_lane_rx_status" -> PfOpticsDiagnosticsLaneRxStatus
		}
	}
	d.SetId(portCID.(string))
	return diags
}
