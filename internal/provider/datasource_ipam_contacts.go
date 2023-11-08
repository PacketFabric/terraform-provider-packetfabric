package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIpamContacts() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceIpamContactsRead,
		Schema: map[string]*schema.Schema{
			"ipam_contacts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"contact_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"org_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phone": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"arin_org_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"apnic_org_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ripe_org_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datasourceIpamContactsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	ipamContacts, err := c.ReadIpamContacts()
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("ipam_contacts", flattenIpamContacts(&ipamContacts))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenIpamContacts(ipamContacts *[]packetfabric.IpamContact) []interface{} {
	if ipamContacts != nil {
		flattens := make([]interface{}, len(*ipamContacts))
		for i, ipamContact := range *ipamContacts {
			flatten := make(map[string]interface{})
			flatten["uuid"] = ipamContact.UUID
			flatten["contact_name"] = ipamContact.ContactName
			flatten["org_name"] = ipamContact.OrgName
			flatten["address"] = ipamContact.Address
			flatten["phone"] = ipamContact.Phone
			flatten["email"] = ipamContact.Email
			flatten["arin_org_id"] = ipamContact.ArinOrgId
			flatten["apnic_org_id"] = ipamContact.ApnicOrgId
			flatten["ripe_org_id"] = ipamContact.RipeOrgId
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}
