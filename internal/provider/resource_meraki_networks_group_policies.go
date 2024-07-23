package provider

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"log"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksGroupPoliciesResource{}
	_ resource.ResourceWithConfigure = &NetworksGroupPoliciesResource{}
)

func NewNetworksGroupPoliciesResource() resource.Resource {
	return &NetworksGroupPoliciesResource{}
}

type NetworksGroupPoliciesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksGroupPoliciesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksGroupPoliciesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_group_policies"
}

func (r *NetworksGroupPoliciesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"bandwidth": schema.SingleNestedAttribute{
				MarkdownDescription: `    The bandwidth settings for clients bound to your group policy.
`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"bandwidth_limits": schema.SingleNestedAttribute{
						MarkdownDescription: `The bandwidth limits object, specifying upload and download speed for clients bound to the group policy. These are only enforced if 'settings' is set to 'custom'.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"limit_down": schema.Int64Attribute{
								MarkdownDescription: `The maximum download limit (integer, in Kbps). null indicates no limit`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
							},
							"limit_up": schema.Int64Attribute{
								MarkdownDescription: `The maximum upload limit (integer, in Kbps). null indicates no limit`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"settings": schema.StringAttribute{
						MarkdownDescription: `How bandwidth limits are enforced. Can be 'network default', 'ignore' or 'custom'.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"custom",
								"ignore",
								"network default",
							),
						},
					},
				},
			},
			"bonjour_forwarding": schema.SingleNestedAttribute{
				MarkdownDescription: `The Bonjour settings for your group policy. Only valid if your network has a wireless configuration.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"rules": schema.SetNestedAttribute{
						MarkdownDescription: `A list of the Bonjour forwarding rules for your group policy. If 'settings' is set to 'custom', at least one rule must be specified.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"description": schema.StringAttribute{
									MarkdownDescription: `A description for your Bonjour forwarding rule. Optional.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"services": schema.SetAttribute{
									MarkdownDescription: `A list of Bonjour services. At least one service must be specified. Available services are 'All Services', 'AirPlay', 'AFP', 'BitTorrent', 'FTP', 'iChat', 'iTunes', 'Printers', 'Samba', 'Scanners' and 'SSH'`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Set{
										setplanmodifier.UseStateForUnknown(),
									},

									ElementType: types.StringType,
								},
								"vlan_id": schema.StringAttribute{
									MarkdownDescription: `The ID of the service VLAN. Required.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
					},
					"settings": schema.StringAttribute{
						MarkdownDescription: `How Bonjour rules are applied. Can be 'network default', 'ignore' or 'custom'.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"custom",
								"ignore",
								"network default",
							),
						},
					},
				},
			},
			"content_filtering": schema.SingleNestedAttribute{
				MarkdownDescription: `The content filtering settings for your group policy`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"allowed_url_patterns": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for allowed URL patterns`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"patterns": schema.SetAttribute{
								MarkdownDescription: `A list of URL patterns that are allowed`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Set{
									setplanmodifier.UseStateForUnknown(),
								},

								ElementType: types.StringType,
							},
							"settings": schema.StringAttribute{
								MarkdownDescription: `How URL patterns are applied. Can be 'network default', 'append' or 'override'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"append",
										"network default",
										"override",
									),
								},
							},
						},
					},
					"blocked_url_categories": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for blocked URL categories`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"categories": schema.SetAttribute{
								MarkdownDescription: `A list of URL categories to block`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Set{
									setplanmodifier.UseStateForUnknown(),
								},

								ElementType: types.StringType,
							},
							"settings": schema.StringAttribute{
								MarkdownDescription: `How URL categories are applied. Can be 'network default', 'append' or 'override'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"append",
										"network default",
										"override",
									),
								},
							},
						},
					},
					"blocked_url_patterns": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for blocked URL patterns`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"patterns": schema.SetAttribute{
								MarkdownDescription: `A list of URL patterns that are blocked`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Set{
									setplanmodifier.UseStateForUnknown(),
								},

								ElementType: types.StringType,
							},
							"settings": schema.StringAttribute{
								MarkdownDescription: `How URL patterns are applied. Can be 'network default', 'append' or 'override'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"append",
										"network default",
										"override",
									),
								},
							},
						},
					},
				},
			},
			"firewall_and_traffic_shaping": schema.SingleNestedAttribute{
				MarkdownDescription: `    The firewall and traffic shaping rules and settings for your policy.
`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"l3_firewall_rules": schema.SetNestedAttribute{
						MarkdownDescription: `An ordered array of the L3 firewall rules`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"comment": schema.StringAttribute{
									MarkdownDescription: `Description of the rule (optional)`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"dest_cidr": schema.StringAttribute{
									MarkdownDescription: `Destination IP address (in IP or CIDR notation), a fully-qualified domain name (FQDN, if your network supports it) or 'any'.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"dest_port": schema.StringAttribute{
									MarkdownDescription: `Destination port (integer in the range 1-65535), a port range (e.g. 8080-9090), or 'any'`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"policy": schema.StringAttribute{
									MarkdownDescription: `'allow' or 'deny' traffic specified by this rule`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"protocol": schema.StringAttribute{
									MarkdownDescription: `The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
					},
					"l7_firewall_rules": schema.SetNestedAttribute{
						MarkdownDescription: `An ordered array of L7 firewall rules`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"policy": schema.StringAttribute{
									MarkdownDescription: `The policy applied to matching traffic. Must be 'deny'.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
									Validators: []validator.String{
										stringvalidator.OneOf(
											"deny",
										),
									},
								},
								"type": schema.StringAttribute{
									MarkdownDescription: `Type of the L7 Rule. Must be 'application', 'applicationCategory', 'host', 'port' or 'ipRange'`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
									Validators: []validator.String{
										stringvalidator.OneOf(
											"application",
											"applicationCategory",
											"host",
											"ipRange",
											"port",
										),
									},
								},
								"value": schema.StringAttribute{
									MarkdownDescription: `The 'value' of what you want to block. If 'type' is 'host', 'port' or 'ipRange', 'value' must be a string matching either a hostname (e.g. somewhere.com), a port (e.g. 8080), or an IP range (e.g. 192.1.0.0/16). If 'type' is 'application' or 'applicationCategory', then 'value' must be an object with an ID for the application.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
					},
					"settings": schema.StringAttribute{
						MarkdownDescription: `How firewall and traffic shaping rules are enforced. Can be 'network default', 'ignore' or 'custom'.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"custom",
								"ignore",
								"network default",
							),
						},
					},
					"traffic_shaping_rules": schema.SetNestedAttribute{
						MarkdownDescription: `    An array of traffic shaping rules. Rules are applied in the order that
    they are specified in. An empty list (or null) means no rules. Note that
    you are allowed a maximum of 8 rules.
`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"definitions": schema.SetNestedAttribute{
									MarkdownDescription: `    A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.
`,
									Computed: true,
									Optional: true,
									PlanModifiers: []planmodifier.Set{
										setplanmodifier.UseStateForUnknown(),
									},
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"type": schema.StringAttribute{
												MarkdownDescription: `The type of definition. Can be one of 'application', 'applicationCategory', 'host', 'port', 'ipRange' or 'localNet'.`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
												Validators: []validator.String{
													stringvalidator.OneOf(
														"application",
														"applicationCategory",
														"host",
														"ipRange",
														"localNet",
														"port",
													),
												},
											},
											"value": schema.StringAttribute{
												MarkdownDescription: `    If "type" is 'host', 'port', 'ipRange' or 'localNet', then "value" must be a string, matching either
    a hostname (e.g. "somesite.com"), a port (e.g. 8080), or an IP range ("192.1.0.0",
    "192.1.0.0/16", or "10.1.0.0/16:80"). 'localNet' also supports CIDR notation, excluding
    custom ports.
     If "type" is 'application' or 'applicationCategory', then "value" must be an object
    with the structure { "id": "meraki:layer7/..." }, where "id" is the application category or
    application ID (for a list of IDs for your network, use the trafficShaping/applicationCategories
    endpoint).
`,
												Computed: true,
												Optional: true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
										},
									},
								},
								"dscp_tag_value": schema.Int64Attribute{
									MarkdownDescription: `    The DSCP tag applied by your rule. null means 'Do not change DSCP tag'.
    For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.
`,
									Computed: true,
									Optional: true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"pcp_tag_value": schema.Int64Attribute{
									MarkdownDescription: `    The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority).
    null means 'Do not set PCP tag'.
`,
									Computed: true,
									Optional: true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"per_client_bandwidth_limits": schema.SingleNestedAttribute{
									MarkdownDescription: `    An object describing the bandwidth settings for your rule.
`,
									Computed: true,
									Optional: true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"bandwidth_limits": schema.SingleNestedAttribute{
											MarkdownDescription: `The bandwidth limits object, specifying the upload ('limitUp') and download ('limitDown') speed in Kbps. These are only enforced if 'settings' is set to 'custom'.`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.Object{
												objectplanmodifier.UseStateForUnknown(),
											},
											Attributes: map[string]schema.Attribute{

												"limit_down": schema.Int64Attribute{
													MarkdownDescription: `The maximum download limit (integer, in Kbps).`,
													Computed:            true,
													Optional:            true,
													PlanModifiers: []planmodifier.Int64{
														int64planmodifier.UseStateForUnknown(),
													},
												},
												"limit_up": schema.Int64Attribute{
													MarkdownDescription: `The maximum upload limit (integer, in Kbps).`,
													Computed:            true,
													Optional:            true,
													PlanModifiers: []planmodifier.Int64{
														int64planmodifier.UseStateForUnknown(),
													},
												},
											},
										},
										"settings": schema.StringAttribute{
											MarkdownDescription: `How bandwidth limits are applied by your rule. Can be one of 'network default', 'ignore' or 'custom'.`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"priority": schema.StringAttribute{
									MarkdownDescription: `    A string, indicating the priority level for packets bound to your rule.
    Can be 'low', 'normal' or 'high'.
`,
									Computed: true,
									Optional: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
					},
				},
			},
			"group_policy_id": schema.StringAttribute{
				MarkdownDescription: `The ID of the group policy`,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name for your group policy. Required.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"scheduling": schema.SingleNestedAttribute{
				MarkdownDescription: `    The schedule for the group policy. Schedules are applied to days of the week.
`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether scheduling is enabled (true) or disabled (false). Defaults to false. If true, the schedule objects for each day of the week (monday - sunday) are parsed.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"friday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Friday.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"monday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Monday.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"saturday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Saturday.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"sunday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Sunday.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"thursday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Thursday.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"tuesday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Tuesday.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"wednesday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Wednesday.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
				},
			},
			"splash_auth_settings": schema.StringAttribute{
				MarkdownDescription: `Whether clients bound to your policy will bypass splash authorization or behave according to the network's rules. Can be one of 'network default' or 'bypass'. Only available if your network has a wireless configuration.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"bypass",
						"network default",
					),
				},
			},
			"vlan_tagging": schema.SingleNestedAttribute{
				MarkdownDescription: `The VLAN tagging settings for your group policy. Only available if your network has a wireless configuration.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"settings": schema.StringAttribute{
						MarkdownDescription: `How VLAN tagging is applied. Can be 'network default', 'ignore' or 'custom'.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"custom",
								"ignore",
								"network default",
							),
						},
					},
					"vlan_id": schema.StringAttribute{
						MarkdownDescription: `The ID of the vlan you want to tag. This only applies if 'settings' is set to 'custom'.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
		},
	}
}

//path params to set ['groupPolicyId']

func (r *NetworksGroupPoliciesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksGroupPoliciesRs

	var item types.Object
	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvName := data.Name.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Networks.GetNetworkGroupPolicies(vvNetworkID)
	// Have Create
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 && restyResp1.StatusCode() != 400 {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkGroupPolicies",
					err.Error(),
				)
				return
			}
		}
	}
	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		log.Printf("[DEBUG] resp: %v", responseStruct)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvGroupPolicyID, ok := result2["GroupPolicyID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter GroupPolicyID",
					"Fail Parsing GroupPolicyID",
				)
				return
			}
			r.client.Networks.UpdateNetworkGroupPolicy(vvNetworkID, vvGroupPolicyID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Networks.GetNetworkGroupPolicy(vvNetworkID, vvGroupPolicyID)
			if responseVerifyItem2 != nil {
				data = ResponseNetworksGetNetworkGroupPolicyItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Networks.CreateNetworkGroupPolicy(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkGroupPolicy",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkGroupPolicy",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Networks.GetNetworkGroupPolicies(vvNetworkID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkGroupPolicies",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkGroupPolicies",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvGroupPolicyID, ok := result2["GroupPolicyID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter GroupPolicyID",
				err.Error(),
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Networks.GetNetworkGroupPolicy(vvNetworkID, vvGroupPolicyID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseNetworksGetNetworkGroupPolicyItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkGroupPolicy",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkGroupPolicy",
				err.Error(),
			)
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"Error in result.",
			"Error in result.",
		)
		return
	}
}

func (r *NetworksGroupPoliciesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksGroupPoliciesRs

	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvGroupPolicyID := data.GroupPolicyID.ValueString()
	// group_policy_id
	responseGet, restyRespGet, err := r.client.Networks.GetNetworkGroupPolicy(vvNetworkID, vvGroupPolicyID)
	if err != nil || restyRespGet == nil {
		if restyRespGet != nil {
			if restyRespGet.StatusCode() == 404 && restyRespGet.StatusCode() != 400 {
				resp.Diagnostics.AddWarning(
					"Resource not found",
					"Deleting resource",
				)
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkGroupPolicy",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkGroupPolicy",
			err.Error(),
		)
		return
	}

	data = ResponseNetworksGetNetworkGroupPolicyItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksGroupPoliciesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("group_policy_id"), idParts[1])...)
}

func (r *NetworksGroupPoliciesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksGroupPoliciesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvGroupPolicyID := data.GroupPolicyID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkGroupPolicy(vvNetworkID, vvGroupPolicyID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkGroupPolicy",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkGroupPolicy",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksGroupPoliciesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksGroupPoliciesRs
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &state, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}

	vvNetworkID := state.NetworkID.ValueString()
	vvGroupPolicyID := state.GroupPolicyID.ValueString()
	_, err := r.client.Networks.DeleteNetworkGroupPolicy(vvNetworkID, vvGroupPolicyID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkGroupPolicy", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksGroupPoliciesRs struct {
	NetworkID                 types.String                                                      `tfsdk:"network_id"`
	GroupPolicyID             types.String                                                      `tfsdk:"group_policy_id"`
	Bandwidth                 *ResponseNetworksGetNetworkGroupPolicyBandwidthRs                 `tfsdk:"bandwidth"`
	BonjourForwarding         *ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRs         `tfsdk:"bonjour_forwarding"`
	ContentFiltering          *ResponseNetworksGetNetworkGroupPolicyContentFilteringRs          `tfsdk:"content_filtering"`
	FirewallAndTrafficShaping *ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingRs `tfsdk:"firewall_and_traffic_shaping"`
	Name                      types.String                                                      `tfsdk:"name"`
	Scheduling                *ResponseNetworksGetNetworkGroupPolicySchedulingRs                `tfsdk:"scheduling"`
	SplashAuthSettings        types.String                                                      `tfsdk:"splash_auth_settings"`
	VLANTagging               *ResponseNetworksGetNetworkGroupPolicyVlanTaggingRs               `tfsdk:"vlan_tagging"`
}

type ResponseNetworksGetNetworkGroupPolicyBandwidthRs struct {
	BandwidthLimits *ResponseNetworksGetNetworkGroupPolicyBandwidthBandwidthLimitsRs `tfsdk:"bandwidth_limits"`
	Settings        types.String                                                     `tfsdk:"settings"`
}

type ResponseNetworksGetNetworkGroupPolicyBandwidthBandwidthLimitsRs struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

type ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRs struct {
	Rules    *[]ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRulesRs `tfsdk:"rules"`
	Settings types.String                                                     `tfsdk:"settings"`
}

type ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRulesRs struct {
	Description types.String `tfsdk:"description"`
	Services    types.Set    `tfsdk:"services"`
	VLANID      types.String `tfsdk:"vlan_id"`
}

type ResponseNetworksGetNetworkGroupPolicyContentFilteringRs struct {
	AllowedURLPatterns   *ResponseNetworksGetNetworkGroupPolicyContentFilteringAllowedUrlPatternsRs   `tfsdk:"allowed_url_patterns"`
	BlockedURLCategories *ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlCategoriesRs `tfsdk:"blocked_url_categories"`
	BlockedURLPatterns   *ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlPatternsRs   `tfsdk:"blocked_url_patterns"`
}

type ResponseNetworksGetNetworkGroupPolicyContentFilteringAllowedUrlPatternsRs struct {
	Patterns types.Set    `tfsdk:"patterns"`
	Settings types.String `tfsdk:"settings"`
}

type ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlCategoriesRs struct {
	Categories types.Set    `tfsdk:"categories"`
	Settings   types.String `tfsdk:"settings"`
}

type ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlPatternsRs struct {
	Patterns types.Set    `tfsdk:"patterns"`
	Settings types.String `tfsdk:"settings"`
}

type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingRs struct {
	L3FirewallRules     *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRulesRs     `tfsdk:"l3_firewall_rules"`
	L7FirewallRules     *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRulesRs     `tfsdk:"l7_firewall_rules"`
	Settings            types.String                                                                           `tfsdk:"settings"`
	TrafficShapingRules *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesRs `tfsdk:"traffic_shaping_rules"`
}

type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRulesRs struct {
	Comment  types.String `tfsdk:"comment"`
	DestCidr types.String `tfsdk:"dest_cidr"`
	DestPort types.String `tfsdk:"dest_port"`
	Policy   types.String `tfsdk:"policy"`
	Protocol types.String `tfsdk:"protocol"`
}

type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRulesRs struct {
	Policy types.String `tfsdk:"policy"`
	Type   types.String `tfsdk:"type"`
	Value  types.String `tfsdk:"value"`
}

type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesRs struct {
	Definitions              *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitionsRs            `tfsdk:"definitions"`
	DscpTagValue             types.Int64                                                                                                  `tfsdk:"dscp_tag_value"`
	PcpTagValue              types.Int64                                                                                                  `tfsdk:"pcp_tag_value"`
	PerClientBandwidthLimits *ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsRs `tfsdk:"per_client_bandwidth_limits"`
	Priority                 types.String                                                                                                 `tfsdk:"priority"`
}

type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitionsRs struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsRs struct {
	BandwidthLimits *ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimitsRs `tfsdk:"bandwidth_limits"`
	Settings        types.String                                                                                                                `tfsdk:"settings"`
}

type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimitsRs struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

type ResponseNetworksGetNetworkGroupPolicySchedulingRs struct {
	Enabled   types.Bool                                                  `tfsdk:"enabled"`
	Friday    *ResponseNetworksGetNetworkGroupPolicySchedulingFridayRs    `tfsdk:"friday"`
	Monday    *ResponseNetworksGetNetworkGroupPolicySchedulingMondayRs    `tfsdk:"monday"`
	Saturday  *ResponseNetworksGetNetworkGroupPolicySchedulingSaturdayRs  `tfsdk:"saturday"`
	Sunday    *ResponseNetworksGetNetworkGroupPolicySchedulingSundayRs    `tfsdk:"sunday"`
	Thursday  *ResponseNetworksGetNetworkGroupPolicySchedulingThursdayRs  `tfsdk:"thursday"`
	Tuesday   *ResponseNetworksGetNetworkGroupPolicySchedulingTuesdayRs   `tfsdk:"tuesday"`
	Wednesday *ResponseNetworksGetNetworkGroupPolicySchedulingWednesdayRs `tfsdk:"wednesday"`
}

type ResponseNetworksGetNetworkGroupPolicySchedulingFridayRs struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseNetworksGetNetworkGroupPolicySchedulingMondayRs struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseNetworksGetNetworkGroupPolicySchedulingSaturdayRs struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseNetworksGetNetworkGroupPolicySchedulingSundayRs struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseNetworksGetNetworkGroupPolicySchedulingThursdayRs struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseNetworksGetNetworkGroupPolicySchedulingTuesdayRs struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseNetworksGetNetworkGroupPolicySchedulingWednesdayRs struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseNetworksGetNetworkGroupPolicyVlanTaggingRs struct {
	Settings types.String `tfsdk:"settings"`
	VLANID   types.String `tfsdk:"vlan_id"`
}

// FromBody
func (r *NetworksGroupPoliciesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksCreateNetworkGroupPolicy {
	emptyString := ""
	var requestNetworksCreateNetworkGroupPolicyBandwidth *merakigosdk.RequestNetworksCreateNetworkGroupPolicyBandwidth
	if r.Bandwidth != nil {
		var requestNetworksCreateNetworkGroupPolicyBandwidthBandwidthLimits *merakigosdk.RequestNetworksCreateNetworkGroupPolicyBandwidthBandwidthLimits
		if r.Bandwidth.BandwidthLimits != nil {
			limitDown := func() *int64 {
				if !r.Bandwidth.BandwidthLimits.LimitDown.IsUnknown() && !r.Bandwidth.BandwidthLimits.LimitDown.IsNull() {
					return r.Bandwidth.BandwidthLimits.LimitDown.ValueInt64Pointer()
				}
				return nil
			}()
			limitUp := func() *int64 {
				if !r.Bandwidth.BandwidthLimits.LimitUp.IsUnknown() && !r.Bandwidth.BandwidthLimits.LimitUp.IsNull() {
					return r.Bandwidth.BandwidthLimits.LimitUp.ValueInt64Pointer()
				}
				return nil
			}()
			requestNetworksCreateNetworkGroupPolicyBandwidthBandwidthLimits = &merakigosdk.RequestNetworksCreateNetworkGroupPolicyBandwidthBandwidthLimits{
				LimitDown: int64ToIntPointer(limitDown),
				LimitUp:   int64ToIntPointer(limitUp),
			}
		}
		settings := r.Bandwidth.Settings.ValueString()
		requestNetworksCreateNetworkGroupPolicyBandwidth = &merakigosdk.RequestNetworksCreateNetworkGroupPolicyBandwidth{
			BandwidthLimits: requestNetworksCreateNetworkGroupPolicyBandwidthBandwidthLimits,
			Settings:        settings,
		}
	}
	var requestNetworksCreateNetworkGroupPolicyBonjourForwarding *merakigosdk.RequestNetworksCreateNetworkGroupPolicyBonjourForwarding
	if r.BonjourForwarding != nil {
		var requestNetworksCreateNetworkGroupPolicyBonjourForwardingRules []merakigosdk.RequestNetworksCreateNetworkGroupPolicyBonjourForwardingRules
		if r.BonjourForwarding.Rules != nil {
			for _, rItem1 := range *r.BonjourForwarding.Rules { //BonjourForwarding.Rules// name: rules
				description := rItem1.Description.ValueString()
				var services []string
				rItem1.Services.ElementsAs(ctx, &services, false)
				vLANID := rItem1.VLANID.ValueString()
				requestNetworksCreateNetworkGroupPolicyBonjourForwardingRules = append(requestNetworksCreateNetworkGroupPolicyBonjourForwardingRules, merakigosdk.RequestNetworksCreateNetworkGroupPolicyBonjourForwardingRules{
					Description: description,
					Services:    services,
					VLANID:      vLANID,
				})
			}
		}
		settings := r.BonjourForwarding.Settings.ValueString()
		requestNetworksCreateNetworkGroupPolicyBonjourForwarding = &merakigosdk.RequestNetworksCreateNetworkGroupPolicyBonjourForwarding{
			Rules: func() *[]merakigosdk.RequestNetworksCreateNetworkGroupPolicyBonjourForwardingRules {
				if len(requestNetworksCreateNetworkGroupPolicyBonjourForwardingRules) > 0 {
					return &requestNetworksCreateNetworkGroupPolicyBonjourForwardingRules
				}
				return nil
			}(),
			Settings: settings,
		}
	}
	var requestNetworksCreateNetworkGroupPolicyContentFiltering *merakigosdk.RequestNetworksCreateNetworkGroupPolicyContentFiltering
	if r.ContentFiltering != nil {
		var requestNetworksCreateNetworkGroupPolicyContentFilteringAllowedURLPatterns *merakigosdk.RequestNetworksCreateNetworkGroupPolicyContentFilteringAllowedURLPatterns
		if r.ContentFiltering.AllowedURLPatterns != nil {
			var patterns []string
			r.ContentFiltering.AllowedURLPatterns.Patterns.ElementsAs(ctx, &patterns, false)
			settings := r.ContentFiltering.AllowedURLPatterns.Settings.ValueString()
			requestNetworksCreateNetworkGroupPolicyContentFilteringAllowedURLPatterns = &merakigosdk.RequestNetworksCreateNetworkGroupPolicyContentFilteringAllowedURLPatterns{
				Patterns: patterns,
				Settings: settings,
			}
		}
		var requestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLCategories *merakigosdk.RequestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLCategories
		if r.ContentFiltering.BlockedURLCategories != nil {
			var categories []string
			r.ContentFiltering.BlockedURLCategories.Categories.ElementsAs(ctx, &categories, false)
			settings := r.ContentFiltering.BlockedURLCategories.Settings.ValueString()
			requestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLCategories = &merakigosdk.RequestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLCategories{
				Categories: categories,
				Settings:   settings,
			}
		}
		var requestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLPatterns *merakigosdk.RequestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLPatterns
		if r.ContentFiltering.BlockedURLPatterns != nil {
			var patterns []string
			r.ContentFiltering.BlockedURLPatterns.Patterns.ElementsAs(ctx, &patterns, false)
			settings := r.ContentFiltering.BlockedURLPatterns.Settings.ValueString()
			requestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLPatterns = &merakigosdk.RequestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLPatterns{
				Patterns: patterns,
				Settings: settings,
			}
		}
		requestNetworksCreateNetworkGroupPolicyContentFiltering = &merakigosdk.RequestNetworksCreateNetworkGroupPolicyContentFiltering{
			AllowedURLPatterns:   requestNetworksCreateNetworkGroupPolicyContentFilteringAllowedURLPatterns,
			BlockedURLCategories: requestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLCategories,
			BlockedURLPatterns:   requestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLPatterns,
		}
	}
	var requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShaping *merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShaping
	if r.FirewallAndTrafficShaping != nil {
		var requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules []merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules
		if r.FirewallAndTrafficShaping.L3FirewallRules != nil {
			for _, rItem1 := range *r.FirewallAndTrafficShaping.L3FirewallRules { //FirewallAndTrafficShaping.L3FirewallRules// name: l3FirewallRules
				comment := rItem1.Comment.ValueString()
				destCidr := rItem1.DestCidr.ValueString()
				destPort := rItem1.DestPort.ValueString()
				policy := rItem1.Policy.ValueString()
				protocol := rItem1.Protocol.ValueString()
				requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules = append(requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules, merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules{
					Comment:  comment,
					DestCidr: destCidr,
					DestPort: destPort,
					Policy:   policy,
					Protocol: protocol,
				})
			}
		}
		var requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules []merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules
		if r.FirewallAndTrafficShaping.L7FirewallRules != nil {
			for _, rItem1 := range *r.FirewallAndTrafficShaping.L7FirewallRules { //FirewallAndTrafficShaping.L7FirewallRules// name: l7FirewallRules
				policy := rItem1.Policy.ValueString()
				typeR := rItem1.Type.ValueString()
				value := rItem1.Value.ValueString()
				requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules = append(requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules, merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules{
					Policy: policy,
					Type:   typeR,
					Value:  value,
				})
			}
		}
		settings := r.FirewallAndTrafficShaping.Settings.ValueString()
		var requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules []merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules
		if r.FirewallAndTrafficShaping.TrafficShapingRules != nil {
			for _, rItem1 := range *r.FirewallAndTrafficShaping.TrafficShapingRules { //FirewallAndTrafficShaping.TrafficShapingRules// name: trafficShapingRules
				var requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions []merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions
				if rItem1.Definitions != nil {
					for _, rItem2 := range *rItem1.Definitions { //FirewallAndTrafficShaping.TrafficShapingRules.Definitions// name: definitions
						typeR := rItem2.Type.ValueString()
						value := rItem2.Value.ValueString()
						requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions = append(requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions, merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions{
							Type:  typeR,
							Value: value,
						})
					}
				}
				dscpTagValue := func() *int64 {
					if !rItem1.DscpTagValue.IsUnknown() && !rItem1.DscpTagValue.IsNull() {
						return rItem1.DscpTagValue.ValueInt64Pointer()
					}
					return nil
				}()
				pcpTagValue := func() *int64 {
					if !rItem1.PcpTagValue.IsUnknown() && !rItem1.PcpTagValue.IsNull() {
						return rItem1.PcpTagValue.ValueInt64Pointer()
					}
					return nil
				}()
				var requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits *merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits
				var requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits *merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits
				if rItem1.PerClientBandwidthLimits.BandwidthLimits != nil {
					limitDown := func() *int64 {
						if !rItem1.PerClientBandwidthLimits.BandwidthLimits.LimitDown.IsUnknown() && !rItem1.PerClientBandwidthLimits.BandwidthLimits.LimitDown.IsNull() {
							return rItem1.PerClientBandwidthLimits.BandwidthLimits.LimitDown.ValueInt64Pointer()
						}
						return nil
					}()
					limitUp := func() *int64 {
						if !rItem1.PerClientBandwidthLimits.BandwidthLimits.LimitUp.IsUnknown() && !rItem1.PerClientBandwidthLimits.BandwidthLimits.LimitUp.IsNull() {
							return rItem1.PerClientBandwidthLimits.BandwidthLimits.LimitUp.ValueInt64Pointer()
						}
						return nil
					}()
					requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits = &merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits{
						LimitDown: int64ToIntPointer(limitDown),
						LimitUp:   int64ToIntPointer(limitUp),
					}
				}
				settings := rItem1.PerClientBandwidthLimits.Settings.ValueString()
				requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits = &merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits{
					BandwidthLimits: requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits,
					Settings:        settings,
				}
				priority := rItem1.Priority.ValueString()
				requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules = append(requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules, merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules{
					Definitions: func() *[]merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions {
						if len(requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions) > 0 {
							return &requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions
						}
						return nil
					}(),
					DscpTagValue:             int64ToIntPointer(dscpTagValue),
					PcpTagValue:              int64ToIntPointer(pcpTagValue),
					PerClientBandwidthLimits: requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits,
					Priority:                 priority,
				})
			}
		}
		requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShaping = &merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShaping{
			L3FirewallRules: func() *[]merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules {
				if len(requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules) > 0 {
					return &requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules
				}
				return nil
			}(),
			L7FirewallRules: func() *[]merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules {
				if len(requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules) > 0 {
					return &requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules
				}
				return nil
			}(),
			Settings: settings,
			TrafficShapingRules: func() *[]merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules {
				if len(requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules) > 0 {
					return &requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules
				}
				return nil
			}(),
		}
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestNetworksCreateNetworkGroupPolicyScheduling *merakigosdk.RequestNetworksCreateNetworkGroupPolicyScheduling
	if r.Scheduling != nil {
		enabled := func() *bool {
			if !r.Scheduling.Enabled.IsUnknown() && !r.Scheduling.Enabled.IsNull() {
				return r.Scheduling.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		var requestNetworksCreateNetworkGroupPolicySchedulingFriday *merakigosdk.RequestNetworksCreateNetworkGroupPolicySchedulingFriday
		if r.Scheduling.Friday != nil {
			active := func() *bool {
				if !r.Scheduling.Friday.Active.IsUnknown() && !r.Scheduling.Friday.Active.IsNull() {
					return r.Scheduling.Friday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.Scheduling.Friday.From.ValueString()
			to := r.Scheduling.Friday.To.ValueString()
			requestNetworksCreateNetworkGroupPolicySchedulingFriday = &merakigosdk.RequestNetworksCreateNetworkGroupPolicySchedulingFriday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestNetworksCreateNetworkGroupPolicySchedulingMonday *merakigosdk.RequestNetworksCreateNetworkGroupPolicySchedulingMonday
		if r.Scheduling.Monday != nil {
			active := func() *bool {
				if !r.Scheduling.Monday.Active.IsUnknown() && !r.Scheduling.Monday.Active.IsNull() {
					return r.Scheduling.Monday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.Scheduling.Monday.From.ValueString()
			to := r.Scheduling.Monday.To.ValueString()
			requestNetworksCreateNetworkGroupPolicySchedulingMonday = &merakigosdk.RequestNetworksCreateNetworkGroupPolicySchedulingMonday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestNetworksCreateNetworkGroupPolicySchedulingSaturday *merakigosdk.RequestNetworksCreateNetworkGroupPolicySchedulingSaturday
		if r.Scheduling.Saturday != nil {
			active := func() *bool {
				if !r.Scheduling.Saturday.Active.IsUnknown() && !r.Scheduling.Saturday.Active.IsNull() {
					return r.Scheduling.Saturday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.Scheduling.Saturday.From.ValueString()
			to := r.Scheduling.Saturday.To.ValueString()
			requestNetworksCreateNetworkGroupPolicySchedulingSaturday = &merakigosdk.RequestNetworksCreateNetworkGroupPolicySchedulingSaturday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestNetworksCreateNetworkGroupPolicySchedulingSunday *merakigosdk.RequestNetworksCreateNetworkGroupPolicySchedulingSunday
		if r.Scheduling.Sunday != nil {
			active := func() *bool {
				if !r.Scheduling.Sunday.Active.IsUnknown() && !r.Scheduling.Sunday.Active.IsNull() {
					return r.Scheduling.Sunday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.Scheduling.Sunday.From.ValueString()
			to := r.Scheduling.Sunday.To.ValueString()
			requestNetworksCreateNetworkGroupPolicySchedulingSunday = &merakigosdk.RequestNetworksCreateNetworkGroupPolicySchedulingSunday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestNetworksCreateNetworkGroupPolicySchedulingThursday *merakigosdk.RequestNetworksCreateNetworkGroupPolicySchedulingThursday
		if r.Scheduling.Thursday != nil {
			active := func() *bool {
				if !r.Scheduling.Thursday.Active.IsUnknown() && !r.Scheduling.Thursday.Active.IsNull() {
					return r.Scheduling.Thursday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.Scheduling.Thursday.From.ValueString()
			to := r.Scheduling.Thursday.To.ValueString()
			requestNetworksCreateNetworkGroupPolicySchedulingThursday = &merakigosdk.RequestNetworksCreateNetworkGroupPolicySchedulingThursday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestNetworksCreateNetworkGroupPolicySchedulingTuesday *merakigosdk.RequestNetworksCreateNetworkGroupPolicySchedulingTuesday
		if r.Scheduling.Tuesday != nil {
			active := func() *bool {
				if !r.Scheduling.Tuesday.Active.IsUnknown() && !r.Scheduling.Tuesday.Active.IsNull() {
					return r.Scheduling.Tuesday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.Scheduling.Tuesday.From.ValueString()
			to := r.Scheduling.Tuesday.To.ValueString()
			requestNetworksCreateNetworkGroupPolicySchedulingTuesday = &merakigosdk.RequestNetworksCreateNetworkGroupPolicySchedulingTuesday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestNetworksCreateNetworkGroupPolicySchedulingWednesday *merakigosdk.RequestNetworksCreateNetworkGroupPolicySchedulingWednesday
		if r.Scheduling.Wednesday != nil {
			active := func() *bool {
				if !r.Scheduling.Wednesday.Active.IsUnknown() && !r.Scheduling.Wednesday.Active.IsNull() {
					return r.Scheduling.Wednesday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.Scheduling.Wednesday.From.ValueString()
			to := r.Scheduling.Wednesday.To.ValueString()
			requestNetworksCreateNetworkGroupPolicySchedulingWednesday = &merakigosdk.RequestNetworksCreateNetworkGroupPolicySchedulingWednesday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		requestNetworksCreateNetworkGroupPolicyScheduling = &merakigosdk.RequestNetworksCreateNetworkGroupPolicyScheduling{
			Enabled:   enabled,
			Friday:    requestNetworksCreateNetworkGroupPolicySchedulingFriday,
			Monday:    requestNetworksCreateNetworkGroupPolicySchedulingMonday,
			Saturday:  requestNetworksCreateNetworkGroupPolicySchedulingSaturday,
			Sunday:    requestNetworksCreateNetworkGroupPolicySchedulingSunday,
			Thursday:  requestNetworksCreateNetworkGroupPolicySchedulingThursday,
			Tuesday:   requestNetworksCreateNetworkGroupPolicySchedulingTuesday,
			Wednesday: requestNetworksCreateNetworkGroupPolicySchedulingWednesday,
		}
	}
	splashAuthSettings := new(string)
	if !r.SplashAuthSettings.IsUnknown() && !r.SplashAuthSettings.IsNull() {
		*splashAuthSettings = r.SplashAuthSettings.ValueString()
	} else {
		splashAuthSettings = &emptyString
	}
	var requestNetworksCreateNetworkGroupPolicyVLANTagging *merakigosdk.RequestNetworksCreateNetworkGroupPolicyVLANTagging
	if r.VLANTagging != nil {
		settings := r.VLANTagging.Settings.ValueString()
		vLANID := r.VLANTagging.VLANID.ValueString()
		requestNetworksCreateNetworkGroupPolicyVLANTagging = &merakigosdk.RequestNetworksCreateNetworkGroupPolicyVLANTagging{
			Settings: settings,
			VLANID:   vLANID,
		}
	}
	out := merakigosdk.RequestNetworksCreateNetworkGroupPolicy{
		Bandwidth:                 requestNetworksCreateNetworkGroupPolicyBandwidth,
		BonjourForwarding:         requestNetworksCreateNetworkGroupPolicyBonjourForwarding,
		ContentFiltering:          requestNetworksCreateNetworkGroupPolicyContentFiltering,
		FirewallAndTrafficShaping: requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShaping,
		Name:                      *name,
		Scheduling:                requestNetworksCreateNetworkGroupPolicyScheduling,
		SplashAuthSettings:        *splashAuthSettings,
		VLANTagging:               requestNetworksCreateNetworkGroupPolicyVLANTagging,
	}
	return &out
}

func (r *NetworksGroupPoliciesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestNetworksUpdateNetworkGroupPolicy {
	emptyString := ""
	var requestNetworksUpdateNetworkGroupPolicyBandwidth *merakigosdk.RequestNetworksUpdateNetworkGroupPolicyBandwidth
	if r.Bandwidth != nil {
		var requestNetworksUpdateNetworkGroupPolicyBandwidthBandwidthLimits *merakigosdk.RequestNetworksUpdateNetworkGroupPolicyBandwidthBandwidthLimits
		if r.Bandwidth.BandwidthLimits != nil {
			limitDown := func() *int64 {
				if !r.Bandwidth.BandwidthLimits.LimitDown.IsUnknown() && !r.Bandwidth.BandwidthLimits.LimitDown.IsNull() {
					return r.Bandwidth.BandwidthLimits.LimitDown.ValueInt64Pointer()
				}
				return nil
			}()
			limitUp := func() *int64 {
				if !r.Bandwidth.BandwidthLimits.LimitUp.IsUnknown() && !r.Bandwidth.BandwidthLimits.LimitUp.IsNull() {
					return r.Bandwidth.BandwidthLimits.LimitUp.ValueInt64Pointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkGroupPolicyBandwidthBandwidthLimits = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicyBandwidthBandwidthLimits{
				LimitDown: int64ToIntPointer(limitDown),
				LimitUp:   int64ToIntPointer(limitUp),
			}
		}
		settings := r.Bandwidth.Settings.ValueString()
		requestNetworksUpdateNetworkGroupPolicyBandwidth = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicyBandwidth{
			BandwidthLimits: requestNetworksUpdateNetworkGroupPolicyBandwidthBandwidthLimits,
			Settings:        settings,
		}
	}
	var requestNetworksUpdateNetworkGroupPolicyBonjourForwarding *merakigosdk.RequestNetworksUpdateNetworkGroupPolicyBonjourForwarding
	if r.BonjourForwarding != nil {
		var requestNetworksUpdateNetworkGroupPolicyBonjourForwardingRules []merakigosdk.RequestNetworksUpdateNetworkGroupPolicyBonjourForwardingRules
		if r.BonjourForwarding.Rules != nil {
			for _, rItem1 := range *r.BonjourForwarding.Rules {
				description := rItem1.Description.ValueString()
				var services []string
				rItem1.Services.ElementsAs(ctx, &services, false)
				vLANID := rItem1.VLANID.ValueString()
				requestNetworksUpdateNetworkGroupPolicyBonjourForwardingRules = append(requestNetworksUpdateNetworkGroupPolicyBonjourForwardingRules, merakigosdk.RequestNetworksUpdateNetworkGroupPolicyBonjourForwardingRules{
					Description: description,
					Services:    services,
					VLANID:      vLANID,
				})
			}
		}
		settings := r.BonjourForwarding.Settings.ValueString()
		requestNetworksUpdateNetworkGroupPolicyBonjourForwarding = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicyBonjourForwarding{
			Rules: func() *[]merakigosdk.RequestNetworksUpdateNetworkGroupPolicyBonjourForwardingRules {
				if len(requestNetworksUpdateNetworkGroupPolicyBonjourForwardingRules) > 0 {
					return &requestNetworksUpdateNetworkGroupPolicyBonjourForwardingRules
				}
				return nil
			}(),
			Settings: settings,
		}
	}
	var requestNetworksUpdateNetworkGroupPolicyContentFiltering *merakigosdk.RequestNetworksUpdateNetworkGroupPolicyContentFiltering
	if r.ContentFiltering != nil {
		var requestNetworksUpdateNetworkGroupPolicyContentFilteringAllowedURLPatterns *merakigosdk.RequestNetworksUpdateNetworkGroupPolicyContentFilteringAllowedURLPatterns
		if r.ContentFiltering.AllowedURLPatterns != nil {
			var patterns []string
			r.ContentFiltering.AllowedURLPatterns.Patterns.ElementsAs(ctx, &patterns, false)
			settings := r.ContentFiltering.AllowedURLPatterns.Settings.ValueString()
			requestNetworksUpdateNetworkGroupPolicyContentFilteringAllowedURLPatterns = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicyContentFilteringAllowedURLPatterns{
				Patterns: patterns,
				Settings: settings,
			}
		}
		var requestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLCategories *merakigosdk.RequestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLCategories
		if r.ContentFiltering.BlockedURLCategories != nil {
			var categories []string
			r.ContentFiltering.BlockedURLCategories.Categories.ElementsAs(ctx, &categories, false)
			settings := r.ContentFiltering.BlockedURLCategories.Settings.ValueString()
			requestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLCategories = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLCategories{
				Categories: categories,
				Settings:   settings,
			}
		}
		var requestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLPatterns *merakigosdk.RequestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLPatterns
		if r.ContentFiltering.BlockedURLPatterns != nil {
			var patterns []string
			r.ContentFiltering.BlockedURLPatterns.Patterns.ElementsAs(ctx, &patterns, false)
			settings := r.ContentFiltering.BlockedURLPatterns.Settings.ValueString()
			requestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLPatterns = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLPatterns{
				Patterns: patterns,
				Settings: settings,
			}
		}
		requestNetworksUpdateNetworkGroupPolicyContentFiltering = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicyContentFiltering{
			AllowedURLPatterns:   requestNetworksUpdateNetworkGroupPolicyContentFilteringAllowedURLPatterns,
			BlockedURLCategories: requestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLCategories,
			BlockedURLPatterns:   requestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLPatterns,
		}
	}
	var requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShaping *merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShaping
	if r.FirewallAndTrafficShaping != nil {
		var requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules []merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules
		if r.FirewallAndTrafficShaping.L3FirewallRules != nil {
			for _, rItem1 := range *r.FirewallAndTrafficShaping.L3FirewallRules {
				comment := rItem1.Comment.ValueString()
				destCidr := rItem1.DestCidr.ValueString()
				destPort := rItem1.DestPort.ValueString()
				policy := rItem1.Policy.ValueString()
				protocol := rItem1.Protocol.ValueString()
				requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules = append(requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules, merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules{
					Comment:  comment,
					DestCidr: destCidr,
					DestPort: destPort,
					Policy:   policy,
					Protocol: protocol,
				})
			}
		}
		var requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules []merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules

		for _, rItem1 := range *r.FirewallAndTrafficShaping.L7FirewallRules {
			policy := rItem1.Policy.ValueString()
			typeR := rItem1.Type.ValueString()
			value := rItem1.Value.ValueString()
			requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules = append(requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules, merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules{
				Policy: policy,
				Type:   typeR,
				Value:  value,
			})
		}

		settings := r.FirewallAndTrafficShaping.Settings.ValueString()
		var requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules []merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules
		if r.FirewallAndTrafficShaping.TrafficShapingRules != nil {
			for _, rItem1 := range *r.FirewallAndTrafficShaping.TrafficShapingRules {
				var requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions []merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions
				if rItem1.Definitions != nil {
					for _, rItem2 := range *rItem1.Definitions {
						typeR := rItem2.Type.ValueString()
						value := rItem2.Value.ValueString()
						requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions = append(requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions, merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions{
							Type:  typeR,
							Value: value,
						})
					}
				}
				dscpTagValue := func() *int64 {
					if !rItem1.DscpTagValue.IsUnknown() && !rItem1.DscpTagValue.IsNull() {
						return rItem1.DscpTagValue.ValueInt64Pointer()
					}
					return nil
				}()
				pcpTagValue := func() *int64 {
					if !rItem1.PcpTagValue.IsUnknown() && !rItem1.PcpTagValue.IsNull() {
						return rItem1.PcpTagValue.ValueInt64Pointer()
					}
					return nil
				}()
				var requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits *merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits
				if rItem1.PerClientBandwidthLimits.BandwidthLimits != nil {
					limitDown := func() *int64 {
						if !rItem1.PerClientBandwidthLimits.BandwidthLimits.LimitDown.IsUnknown() && !rItem1.PerClientBandwidthLimits.BandwidthLimits.LimitDown.IsNull() {
							return rItem1.PerClientBandwidthLimits.BandwidthLimits.LimitDown.ValueInt64Pointer()
						}
						return nil
					}()
					limitUp := func() *int64 {
						if !rItem1.PerClientBandwidthLimits.BandwidthLimits.LimitUp.IsUnknown() && !rItem1.PerClientBandwidthLimits.BandwidthLimits.LimitUp.IsNull() {
							return rItem1.PerClientBandwidthLimits.BandwidthLimits.LimitUp.ValueInt64Pointer()
						}
						return nil
					}()
					requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits{
						BandwidthLimits: &merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits{
							LimitDown: int64ToIntPointer(limitDown),
							LimitUp:   int64ToIntPointer(limitUp),
						},
						Settings: rItem1.PerClientBandwidthLimits.Settings.ValueString(),
					}
				}
				priority := rItem1.Priority.ValueString()
				requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules = append(requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules, merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules{
					Definitions: func() *[]merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions {
						if len(requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions) > 0 {
							return &requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions
						}
						return nil
					}(),
					DscpTagValue:             int64ToIntPointer(dscpTagValue),
					PcpTagValue:              int64ToIntPointer(pcpTagValue),
					PerClientBandwidthLimits: requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits,
					Priority:                 priority,
				})
			}
		}
		requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShaping = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShaping{
			L3FirewallRules: func() *[]merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules {
				if len(requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules) > 0 {
					return &requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules
				}
				return nil
			}(),
			L7FirewallRules: func() *[]merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules {
				if len(requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules) > 0 {
					return &requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules
				}
				return nil
			}(),
			Settings: settings,
			TrafficShapingRules: func() *[]merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules {
				if len(requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules) > 0 {
					return &requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules
				}
				return nil
			}(),
		}
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestNetworksUpdateNetworkGroupPolicyScheduling *merakigosdk.RequestNetworksUpdateNetworkGroupPolicyScheduling
	if r.Scheduling != nil {
		enabled := func() *bool {
			if !r.Scheduling.Enabled.IsUnknown() && !r.Scheduling.Enabled.IsNull() {
				return r.Scheduling.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		var requestNetworksUpdateNetworkGroupPolicySchedulingFriday *merakigosdk.RequestNetworksUpdateNetworkGroupPolicySchedulingFriday
		if r.Scheduling.Friday != nil {
			active := func() *bool {
				if !r.Scheduling.Friday.Active.IsUnknown() && !r.Scheduling.Friday.Active.IsNull() {
					return r.Scheduling.Friday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.Scheduling.Friday.From.ValueString()
			to := r.Scheduling.Friday.To.ValueString()
			requestNetworksUpdateNetworkGroupPolicySchedulingFriday = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicySchedulingFriday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestNetworksUpdateNetworkGroupPolicySchedulingMonday *merakigosdk.RequestNetworksUpdateNetworkGroupPolicySchedulingMonday
		if r.Scheduling.Monday != nil {
			active := func() *bool {
				if !r.Scheduling.Monday.Active.IsUnknown() && !r.Scheduling.Monday.Active.IsNull() {
					return r.Scheduling.Monday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.Scheduling.Monday.From.ValueString()
			to := r.Scheduling.Monday.To.ValueString()
			requestNetworksUpdateNetworkGroupPolicySchedulingMonday = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicySchedulingMonday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestNetworksUpdateNetworkGroupPolicySchedulingSaturday *merakigosdk.RequestNetworksUpdateNetworkGroupPolicySchedulingSaturday
		if r.Scheduling.Saturday != nil {
			active := func() *bool {
				if !r.Scheduling.Saturday.Active.IsUnknown() && !r.Scheduling.Saturday.Active.IsNull() {
					return r.Scheduling.Saturday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.Scheduling.Saturday.From.ValueString()
			to := r.Scheduling.Saturday.To.ValueString()
			requestNetworksUpdateNetworkGroupPolicySchedulingSaturday = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicySchedulingSaturday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestNetworksUpdateNetworkGroupPolicySchedulingSunday *merakigosdk.RequestNetworksUpdateNetworkGroupPolicySchedulingSunday
		if r.Scheduling.Sunday != nil {
			active := func() *bool {
				if !r.Scheduling.Sunday.Active.IsUnknown() && !r.Scheduling.Sunday.Active.IsNull() {
					return r.Scheduling.Sunday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.Scheduling.Sunday.From.ValueString()
			to := r.Scheduling.Sunday.To.ValueString()
			requestNetworksUpdateNetworkGroupPolicySchedulingSunday = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicySchedulingSunday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestNetworksUpdateNetworkGroupPolicySchedulingThursday *merakigosdk.RequestNetworksUpdateNetworkGroupPolicySchedulingThursday
		if r.Scheduling.Thursday != nil {
			active := func() *bool {
				if !r.Scheduling.Thursday.Active.IsUnknown() && !r.Scheduling.Thursday.Active.IsNull() {
					return r.Scheduling.Thursday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.Scheduling.Thursday.From.ValueString()
			to := r.Scheduling.Thursday.To.ValueString()
			requestNetworksUpdateNetworkGroupPolicySchedulingThursday = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicySchedulingThursday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestNetworksUpdateNetworkGroupPolicySchedulingTuesday *merakigosdk.RequestNetworksUpdateNetworkGroupPolicySchedulingTuesday
		if r.Scheduling.Tuesday != nil {
			active := func() *bool {
				if !r.Scheduling.Tuesday.Active.IsUnknown() && !r.Scheduling.Tuesday.Active.IsNull() {
					return r.Scheduling.Tuesday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.Scheduling.Tuesday.From.ValueString()
			to := r.Scheduling.Tuesday.To.ValueString()
			requestNetworksUpdateNetworkGroupPolicySchedulingTuesday = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicySchedulingTuesday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestNetworksUpdateNetworkGroupPolicySchedulingWednesday *merakigosdk.RequestNetworksUpdateNetworkGroupPolicySchedulingWednesday
		if r.Scheduling.Wednesday != nil {
			active := func() *bool {
				if !r.Scheduling.Wednesday.Active.IsUnknown() && !r.Scheduling.Wednesday.Active.IsNull() {
					return r.Scheduling.Wednesday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.Scheduling.Wednesday.From.ValueString()
			to := r.Scheduling.Wednesday.To.ValueString()
			requestNetworksUpdateNetworkGroupPolicySchedulingWednesday = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicySchedulingWednesday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		requestNetworksUpdateNetworkGroupPolicyScheduling = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicyScheduling{
			Enabled:   enabled,
			Friday:    requestNetworksUpdateNetworkGroupPolicySchedulingFriday,
			Monday:    requestNetworksUpdateNetworkGroupPolicySchedulingMonday,
			Saturday:  requestNetworksUpdateNetworkGroupPolicySchedulingSaturday,
			Sunday:    requestNetworksUpdateNetworkGroupPolicySchedulingSunday,
			Thursday:  requestNetworksUpdateNetworkGroupPolicySchedulingThursday,
			Tuesday:   requestNetworksUpdateNetworkGroupPolicySchedulingTuesday,
			Wednesday: requestNetworksUpdateNetworkGroupPolicySchedulingWednesday,
		}
	}
	splashAuthSettings := new(string)
	if !r.SplashAuthSettings.IsUnknown() && !r.SplashAuthSettings.IsNull() {
		*splashAuthSettings = r.SplashAuthSettings.ValueString()
	} else {
		splashAuthSettings = &emptyString
	}
	var requestNetworksUpdateNetworkGroupPolicyVLANTagging *merakigosdk.RequestNetworksUpdateNetworkGroupPolicyVLANTagging
	if r.VLANTagging != nil {
		settings := r.VLANTagging.Settings.ValueString()
		vLANID := r.VLANTagging.VLANID.ValueString()
		requestNetworksUpdateNetworkGroupPolicyVLANTagging = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicyVLANTagging{
			Settings: settings,
			VLANID:   vLANID,
		}
	}
	out := merakigosdk.RequestNetworksUpdateNetworkGroupPolicy{
		Bandwidth:                 requestNetworksUpdateNetworkGroupPolicyBandwidth,
		BonjourForwarding:         requestNetworksUpdateNetworkGroupPolicyBonjourForwarding,
		ContentFiltering:          requestNetworksUpdateNetworkGroupPolicyContentFiltering,
		FirewallAndTrafficShaping: requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShaping,
		Name:                      *name,
		Scheduling:                requestNetworksUpdateNetworkGroupPolicyScheduling,
		SplashAuthSettings:        *splashAuthSettings,
		VLANTagging:               requestNetworksUpdateNetworkGroupPolicyVLANTagging,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseNetworksGetNetworkGroupPolicyItemToBodyRs(state NetworksGroupPoliciesRs, response *merakigosdk.ResponseNetworksGetNetworkGroupPolicy, is_read bool) NetworksGroupPoliciesRs {
	itemState := NetworksGroupPoliciesRs{
		Bandwidth: func() *ResponseNetworksGetNetworkGroupPolicyBandwidthRs {
			if response.Bandwidth != nil {
				return &ResponseNetworksGetNetworkGroupPolicyBandwidthRs{
					BandwidthLimits: func() *ResponseNetworksGetNetworkGroupPolicyBandwidthBandwidthLimitsRs {
						if response.Bandwidth.BandwidthLimits != nil {
							return &ResponseNetworksGetNetworkGroupPolicyBandwidthBandwidthLimitsRs{
								LimitDown: func() types.Int64 {
									if response.Bandwidth.BandwidthLimits.LimitDown != nil {
										return types.Int64Value(int64(*response.Bandwidth.BandwidthLimits.LimitDown))
									}
									return types.Int64{}
								}(),
								LimitUp: func() types.Int64 {
									if response.Bandwidth.BandwidthLimits.LimitUp != nil {
										return types.Int64Value(int64(*response.Bandwidth.BandwidthLimits.LimitUp))
									}
									return types.Int64{}
								}(),
							}
						}
						return &ResponseNetworksGetNetworkGroupPolicyBandwidthBandwidthLimitsRs{}
					}(),
					Settings: types.StringValue(response.Bandwidth.Settings),
				}
			}
			return &ResponseNetworksGetNetworkGroupPolicyBandwidthRs{}
		}(),
		BonjourForwarding: func() *ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRs {
			if response.BonjourForwarding != nil {
				return &ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRs{
					Rules: func() *[]ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRulesRs {
						if response.BonjourForwarding.Rules != nil {
							result := make([]ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRulesRs, len(*response.BonjourForwarding.Rules))
							for i, rules := range *response.BonjourForwarding.Rules {
								result[i] = ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRulesRs{
									Description: types.StringValue(rules.Description),
									Services:    StringSliceToSet(rules.Services),
									VLANID:      types.StringValue(rules.VLANID),
								}
							}
							return &result
						}
						return &[]ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRulesRs{}
					}(),
					Settings: types.StringValue(response.BonjourForwarding.Settings),
				}
			}
			return &ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRs{}
		}(),
		ContentFiltering: func() *ResponseNetworksGetNetworkGroupPolicyContentFilteringRs {
			if response.ContentFiltering != nil {
				return &ResponseNetworksGetNetworkGroupPolicyContentFilteringRs{
					AllowedURLPatterns: func() *ResponseNetworksGetNetworkGroupPolicyContentFilteringAllowedUrlPatternsRs {
						if response.ContentFiltering.AllowedURLPatterns != nil {
							return &ResponseNetworksGetNetworkGroupPolicyContentFilteringAllowedUrlPatternsRs{
								Patterns: StringSliceToSet(response.ContentFiltering.AllowedURLPatterns.Patterns),
								Settings: types.StringValue(response.ContentFiltering.AllowedURLPatterns.Settings),
							}
						}
						return &ResponseNetworksGetNetworkGroupPolicyContentFilteringAllowedUrlPatternsRs{}
					}(),
					BlockedURLCategories: func() *ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlCategoriesRs {
						if response.ContentFiltering.BlockedURLCategories != nil {
							return &ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlCategoriesRs{
								Categories: StringSliceToSet(response.ContentFiltering.BlockedURLCategories.Categories),
								Settings:   types.StringValue(response.ContentFiltering.BlockedURLCategories.Settings),
							}
						}
						return &ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlCategoriesRs{}
					}(),
					BlockedURLPatterns: func() *ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlPatternsRs {
						if response.ContentFiltering.BlockedURLPatterns != nil {
							return &ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlPatternsRs{
								Patterns: StringSliceToSet(response.ContentFiltering.BlockedURLPatterns.Patterns),
								Settings: types.StringValue(response.ContentFiltering.BlockedURLPatterns.Settings),
							}
						}
						return &ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlPatternsRs{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkGroupPolicyContentFilteringRs{}
		}(),
		FirewallAndTrafficShaping: func() *ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingRs {
			if response.FirewallAndTrafficShaping != nil {
				return &ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingRs{
					L3FirewallRules: func() *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRulesRs {
						if response.FirewallAndTrafficShaping.L3FirewallRules != nil {
							result := make([]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRulesRs, len(*response.FirewallAndTrafficShaping.L3FirewallRules))
							for i, l3FirewallRules := range *response.FirewallAndTrafficShaping.L3FirewallRules {
								result[i] = ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRulesRs{
									Comment:  types.StringValue(l3FirewallRules.Comment),
									DestCidr: types.StringValue(l3FirewallRules.DestCidr),
									DestPort: types.StringValue(l3FirewallRules.DestPort),
									Policy:   types.StringValue(l3FirewallRules.Policy),
									Protocol: types.StringValue(l3FirewallRules.Protocol),
								}
							}
							return &result
						}
						return &[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRulesRs{}
					}(),
					L7FirewallRules: func() *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRulesRs {
						if response.FirewallAndTrafficShaping.L7FirewallRules != nil {
							result := make([]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRulesRs, len(*response.FirewallAndTrafficShaping.L7FirewallRules))
							for i, l7FirewallRules := range *response.FirewallAndTrafficShaping.L7FirewallRules {
								result[i] = ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRulesRs{
									Policy: types.StringValue(l7FirewallRules.Policy),
									Type:   types.StringValue(l7FirewallRules.Type),
									Value:  types.StringValue(l7FirewallRules.Value),
								}
							}
							return &result
						}
						return &[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRulesRs{}
					}(),
					Settings: types.StringValue(response.FirewallAndTrafficShaping.Settings),
					TrafficShapingRules: func() *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesRs {
						if response.FirewallAndTrafficShaping.TrafficShapingRules != nil {
							result := make([]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesRs, len(*response.FirewallAndTrafficShaping.TrafficShapingRules))
							for i, trafficShapingRules := range *response.FirewallAndTrafficShaping.TrafficShapingRules {
								result[i] = ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesRs{
									Definitions: func() *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitionsRs {
										if trafficShapingRules.Definitions != nil {
											result := make([]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitionsRs, len(*trafficShapingRules.Definitions))
											for i, definitions := range *trafficShapingRules.Definitions {
												result[i] = ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitionsRs{
													Type:  types.StringValue(definitions.Type),
													Value: types.StringValue(definitions.Value),
												}
											}
											return &result
										}
										return &[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitionsRs{}
									}(),
									DscpTagValue: func() types.Int64 {
										if trafficShapingRules.DscpTagValue != nil {
											return types.Int64Value(int64(*trafficShapingRules.DscpTagValue))
										}
										return types.Int64{}
									}(),
									PcpTagValue: func() types.Int64 {
										if trafficShapingRules.PcpTagValue != nil {
											return types.Int64Value(int64(*trafficShapingRules.PcpTagValue))
										}
										return types.Int64{}
									}(),
									PerClientBandwidthLimits: func() *ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsRs {
										if trafficShapingRules.PerClientBandwidthLimits != nil {
											return &ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsRs{
												BandwidthLimits: func() *ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimitsRs {
													if trafficShapingRules.PerClientBandwidthLimits.BandwidthLimits != nil {
														return &ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimitsRs{
															LimitDown: func() types.Int64 {
																if trafficShapingRules.PerClientBandwidthLimits.BandwidthLimits.LimitDown != nil {
																	return types.Int64Value(int64(*trafficShapingRules.PerClientBandwidthLimits.BandwidthLimits.LimitDown))
																}
																return types.Int64{}
															}(),
															LimitUp: func() types.Int64 {
																if trafficShapingRules.PerClientBandwidthLimits.BandwidthLimits.LimitUp != nil {
																	return types.Int64Value(int64(*trafficShapingRules.PerClientBandwidthLimits.BandwidthLimits.LimitUp))
																}
																return types.Int64{}
															}(),
														}
													}
													return &ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimitsRs{}
												}(),
												Settings: types.StringValue(trafficShapingRules.PerClientBandwidthLimits.Settings),
											}
										}
										return &ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsRs{}
									}(),
								}
							}
							return &result
						}
						return &[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesRs{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingRs{}
		}(),
		GroupPolicyID: types.StringValue(response.GroupPolicyID),
		Name:          types.StringValue(response.Name),
		Scheduling: func() *ResponseNetworksGetNetworkGroupPolicySchedulingRs {
			if response.Scheduling != nil {
				return &ResponseNetworksGetNetworkGroupPolicySchedulingRs{
					Enabled: func() types.Bool {
						if response.Scheduling.Enabled != nil {
							return types.BoolValue(*response.Scheduling.Enabled)
						}
						return types.Bool{}
					}(),
					Friday: func() *ResponseNetworksGetNetworkGroupPolicySchedulingFridayRs {
						if response.Scheduling.Friday != nil {
							return &ResponseNetworksGetNetworkGroupPolicySchedulingFridayRs{
								Active: func() types.Bool {
									if response.Scheduling.Friday.Active != nil {
										return types.BoolValue(*response.Scheduling.Friday.Active)
									}
									return types.Bool{}
								}(),
								From: types.StringValue(response.Scheduling.Friday.From),
								To:   types.StringValue(response.Scheduling.Friday.To),
							}
						}
						return &ResponseNetworksGetNetworkGroupPolicySchedulingFridayRs{}
					}(),
					Monday: func() *ResponseNetworksGetNetworkGroupPolicySchedulingMondayRs {
						if response.Scheduling.Monday != nil {
							return &ResponseNetworksGetNetworkGroupPolicySchedulingMondayRs{
								Active: func() types.Bool {
									if response.Scheduling.Monday.Active != nil {
										return types.BoolValue(*response.Scheduling.Monday.Active)
									}
									return types.Bool{}
								}(),
								From: types.StringValue(response.Scheduling.Monday.From),
								To:   types.StringValue(response.Scheduling.Monday.To),
							}
						}
						return &ResponseNetworksGetNetworkGroupPolicySchedulingMondayRs{}
					}(),
					Saturday: func() *ResponseNetworksGetNetworkGroupPolicySchedulingSaturdayRs {
						if response.Scheduling.Saturday != nil {
							return &ResponseNetworksGetNetworkGroupPolicySchedulingSaturdayRs{
								Active: func() types.Bool {
									if response.Scheduling.Saturday.Active != nil {
										return types.BoolValue(*response.Scheduling.Saturday.Active)
									}
									return types.Bool{}
								}(),
								From: types.StringValue(response.Scheduling.Saturday.From),
								To:   types.StringValue(response.Scheduling.Saturday.To),
							}
						}
						return &ResponseNetworksGetNetworkGroupPolicySchedulingSaturdayRs{}
					}(),
					Sunday: func() *ResponseNetworksGetNetworkGroupPolicySchedulingSundayRs {
						if response.Scheduling.Sunday != nil {
							return &ResponseNetworksGetNetworkGroupPolicySchedulingSundayRs{
								Active: func() types.Bool {
									if response.Scheduling.Sunday.Active != nil {
										return types.BoolValue(*response.Scheduling.Sunday.Active)
									}
									return types.Bool{}
								}(),
								From: types.StringValue(response.Scheduling.Sunday.From),
								To:   types.StringValue(response.Scheduling.Sunday.To),
							}
						}
						return &ResponseNetworksGetNetworkGroupPolicySchedulingSundayRs{}
					}(),
					Thursday: func() *ResponseNetworksGetNetworkGroupPolicySchedulingThursdayRs {
						if response.Scheduling.Thursday != nil {
							return &ResponseNetworksGetNetworkGroupPolicySchedulingThursdayRs{
								Active: func() types.Bool {
									if response.Scheduling.Thursday.Active != nil {
										return types.BoolValue(*response.Scheduling.Thursday.Active)
									}
									return types.Bool{}
								}(),
								From: types.StringValue(response.Scheduling.Thursday.From),
								To:   types.StringValue(response.Scheduling.Thursday.To),
							}
						}
						return &ResponseNetworksGetNetworkGroupPolicySchedulingThursdayRs{}
					}(),
					Tuesday: func() *ResponseNetworksGetNetworkGroupPolicySchedulingTuesdayRs {
						if response.Scheduling.Tuesday != nil {
							return &ResponseNetworksGetNetworkGroupPolicySchedulingTuesdayRs{
								Active: func() types.Bool {
									if response.Scheduling.Tuesday.Active != nil {
										return types.BoolValue(*response.Scheduling.Tuesday.Active)
									}
									return types.Bool{}
								}(),
								From: types.StringValue(response.Scheduling.Tuesday.From),
								To:   types.StringValue(response.Scheduling.Tuesday.To),
							}
						}
						return &ResponseNetworksGetNetworkGroupPolicySchedulingTuesdayRs{}
					}(),
					Wednesday: func() *ResponseNetworksGetNetworkGroupPolicySchedulingWednesdayRs {
						if response.Scheduling.Wednesday != nil {
							return &ResponseNetworksGetNetworkGroupPolicySchedulingWednesdayRs{
								Active: func() types.Bool {
									if response.Scheduling.Wednesday.Active != nil {
										return types.BoolValue(*response.Scheduling.Wednesday.Active)
									}
									return types.Bool{}
								}(),
								From: types.StringValue(response.Scheduling.Wednesday.From),
								To:   types.StringValue(response.Scheduling.Wednesday.To),
							}
						}
						return &ResponseNetworksGetNetworkGroupPolicySchedulingWednesdayRs{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkGroupPolicySchedulingRs{}
		}(),
		SplashAuthSettings: types.StringValue(response.SplashAuthSettings),
		VLANTagging: func() *ResponseNetworksGetNetworkGroupPolicyVlanTaggingRs {
			if response.VLANTagging != nil {
				return &ResponseNetworksGetNetworkGroupPolicyVlanTaggingRs{
					Settings: types.StringValue(response.VLANTagging.Settings),
					VLANID:   types.StringValue(response.VLANTagging.VLANID),
				}
			}
			return &ResponseNetworksGetNetworkGroupPolicyVlanTaggingRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksGroupPoliciesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksGroupPoliciesRs)
}
