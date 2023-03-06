package provider

import (
	"context"
	"regexp"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"login": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 255),
				Description:  "User login.",
			},
			"first_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 255),
				Description:  "User first name.",
			},
			"last_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 255),
				Description:  "User last name.",
			},
			"email": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "User e-mail.",
			},
			"phone": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`^[0-9 ()+.-]+(\s?(x|ex|ext|ete|extn)?(\.|\.\s|\s)?[\d]{1,9})?$`), "Phone number must match the pattern ^[0-9 ()+.-]+(\\s?(x|ex|ext|ete|extn)?(\\.|\\.\\s|\\s)?[\\d]{1,9})?$"),
				Description:  "User phone number.",
			},
			"password": {
				Type:         schema.TypeString,
				Required:     true,
				Sensitive:    true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(8, 64),
				Description:  "User password. Keep it in secret.",
			},
			"timezone": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User time-zone. You can find the list of available timezones [here](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).",
			},
			"group": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"admin", "regular", "read-only", "support", "sales"}, false),
				Description:  "User group. Available options are admin, regular, read-only, support, and sales.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	user := extractUser(d)
	resp, err := c.CreateUsers(user)
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		d.SetId(resp.Login)
	}
	return diags
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	userID := d.Get("login").(string)
	resp, err := c.ReadUsers(userID)
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = d.Set("first_name", resp.FirstName)
		_ = d.Set("last_name", resp.LastName)
		_ = d.Set("email", resp.Email)
		_ = d.Set("phone", resp.Phone)
		_ = d.Set("timezone", resp.Timezone)
		_ = d.Set("group", resp.Group)
	}
	return diags
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	userID := d.Get("login").(string)
	userUpdate := packetfabric.UserUpdate{
		FirstName: d.Get("first_name").(string),
		LastName:  d.Get("last_name").(string),
		Phone:     d.Get("phone").(string),
		Login:     d.Get("login").(string),
		Timezone:  d.Get("timezone").(string),
		Group:     d.Get("group").(string),
	}
	_, err := c.UpdateUser(userUpdate, userID)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	userID := d.Get("login").(string)
	_, err := c.DeleteUsers(userID)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}

func extractUser(d *schema.ResourceData) packetfabric.User {
	user := packetfabric.User{}
	if firstName, ok := d.GetOk("first_name"); ok {
		user.FirstName = firstName.(string)
	}
	if lastName, ok := d.GetOk("last_name"); ok {
		user.LastName = lastName.(string)
	}
	if email, ok := d.GetOk("email"); ok {
		user.Email = email.(string)
	}
	if phone, ok := d.GetOk("phone"); ok {
		user.Phone = phone.(string)
	}
	if login, ok := d.GetOk("login"); ok {
		user.Login = login.(string)
	}
	if password, ok := d.GetOk("password"); ok {
		user.Password = password.(string)
	}
	if timezone, ok := d.GetOk("timezone"); ok {
		user.Timezone = timezone.(string)
	}
	if group, ok := d.GetOk("group"); ok {
		user.Group = group.(string)
	}
	return user
}
