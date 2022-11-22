package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourcePortDeviceInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePortDeviceInfoRead,
		Schema: map[string]*schema.Schema{
			"port_circuit_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The port circuit ID.",
			},
			"adjacent_router": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The adjcent router.",
			},
			"device_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The device ID.",
			},
			"device_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The device name.",
			},
			"device_make": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The device make name.",
			},
			"admin_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The current admin status.",
			},
			"oper_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The current oerational status.",
			},
			"auto_negotiation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True if auto negotiation is on.",
			},
			"iface_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The interface name.",
			},
			"speed": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The port speed.",
			},
			"site_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The site ID.",
			},
			"optics_diagnostics_lane_tx_power_dbm": {
				Type:        schema.TypeFloat,
				Optional:    true,
				Description: "The optics diagnostics lane TX Power dbm.",
			},
			"optics_diagnostics_lane_tx_power": {
				Type:        schema.TypeFloat,
				Optional:    true,
				Description: "The optics diagnostics lane TX Power.",
			},
			"optics_diagnostics_lane_index": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The optics diagnostics lane Index.",
			},
			"optics_diagnostics_lane_rx_power_dbm": {
				Type:        schema.TypeFloat,
				Optional:    true,
				Description: "The optics diagnostics lane RX Power dbm.",
			},
			"optics_diagnostics_lane_rx_power": {
				Type:        schema.TypeFloat,
				Optional:    true,
				Description: "The optics diagnostics lane RX Power.",
			},
			"optics_diagnostics_lane_bias_current": {
				Type:        schema.TypeFloat,
				Optional:    true,
				Description: "The optics diagnostics lane bias current.",
			},
			"polltime": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The port pool time.",
			},
			"time_flapped": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The port time flapped.",
			},
			"traffic_rx_bps": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The port traffic RX bps.",
			},
			"traffic_rx_bytes": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The port traffic RX bytes.",
			},
			"traffic_rx_ipv6_bytes": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The port traffic RX IPv6 bytes.",
			},
			"traffic_rx_ipv6_packets": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The port traffic RX IPv6 packets.",
			},
			"traffic_rx_packets": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The port traffic RX packets.",
			},
			"traffic_rx_pps": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The port traffic RX pps.",
			},
			"traffic_tx_bps": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The port traffic TX bps.",
			},
			"traffic_tx_bytes": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The port traffic TX bytes.",
			},
			"traffic_tx_ipv6_bytes": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The port traffic TX IPv6 bytes.",
			},
			"traffic_tx_ipv6_packets": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The port traffic TX IPv6 packets.",
			},
			"traffic_tx_packets": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The port traffic TX bytes.",
			},
			"traffic_tx_pps": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The port traffic TX pps.",
			},
			"wiring_media": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The port wiring media.",
			},
			"wiring_module": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The port wiring module.",
			},
			"wiring_panel": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The port wiring panel.",
			},
			"wiring_position": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The port wiring position.",
			},
			"wiring_reach": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The port wiring reach.",
			},
			"wiring_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The port wiring type.",
			},
			"lag_speed": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The port lag speed.",
			},
			"device_can_lag": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True if device can lag.",
			},
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
		if portInfo.AdjacentRouter != nil {
			_ = d.Set("adjacent_router", portInfo.AdjacentRouter)
		}
		_ = d.Set("device_id", portInfo.DeviceID)
		_ = d.Set("device_name", portInfo.DeviceName)
		_ = d.Set("device_make", portInfo.DeviceMake)
		_ = d.Set("admin_status", portInfo.AdminStatus)
		_ = d.Set("oper_status", portInfo.OperStatus)
		_ = d.Set("auto_negotiation", portInfo.AutoNegotiation)
		_ = d.Set("iface_name", portInfo.IfaceName)
		_ = d.Set("speed", portInfo.Speed)
		_ = d.Set("site_id", portInfo.SiteID)
		for _, optics := range portInfo.OpticsDiagnosticsLaneValues {
			_ = d.Set("optics_diagnostics_lane_tx_power_dbm", optics.TxPowerDbm)
			_ = d.Set("optics_diagnostics_lane_tx_power", optics.TxPower)
			_ = d.Set("optics_diagnostics_lane_index", optics.LaneIndex)
			_ = d.Set("optics_diagnostics_lane_rx_power_dbm", optics.RxPowerDbm)
			_ = d.Set("optics_diagnostics_lane_rx_power", optics.RxPower)
			_ = d.Set("optics_diagnostics_lane_bias_current", optics.BiasCurrent)
		}
		if portInfo.Polltime != nil {
			_ = d.Set("polltime", portInfo.Polltime)
		}
		_ = d.Set("time_flapped", portInfo.TimeFlapped)
		_ = d.Set("traffic_rx_bps", portInfo.TrafficRxBps)
		_ = d.Set("traffic_rx_bytes", portInfo.TrafficRxBytes)
		_ = d.Set("traffic_rx_ipv6_bytes", portInfo.TrafficRxIpv6Bytes)
		_ = d.Set("traffic_rx_ipv6_packets", portInfo.TrafficRxIpv6Packets)
		_ = d.Set("traffic_rx_packets", portInfo.TrafficRxPackets)
		_ = d.Set("traffic_rx_pps", portInfo.TrafficRxPps)
		_ = d.Set("traffic_tx_bps", portInfo.TrafficTxBps)
		_ = d.Set("traffic_tx_bytes", portInfo.TrafficTxBytes)
		_ = d.Set("traffic_tx_ipv6_bytes", portInfo.TrafficTxIpv6Bytes)
		_ = d.Set("traffic_tx_ipv6_packets", portInfo.TrafficTxIpv6Packets)
		_ = d.Set("traffic_tx_packets", portInfo.TrafficTxPackets)
		_ = d.Set("traffic_tx_pps", portInfo.TrafficTxPps)
		_ = d.Set("wiring_media", portInfo.WiringMedia)
		_ = d.Set("wiring_module", portInfo.WiringModule)
		_ = d.Set("wiring_panel", portInfo.WiringPanel)
		_ = d.Set("wiring_position", portInfo.WiringPosition)
		_ = d.Set("wiring_reach", portInfo.WiringReach)
		_ = d.Set("wiring_type", portInfo.WiringType)
		_ = d.Set("lag_speed", portInfo.LagSpeed)
		_ = d.Set("device_can_lag", portInfo.DeviceCanLag)
	}
	d.SetId(portCID.(string))
	return diags
}
