// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0
package provider

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"limit_down": schema.Int64Attribute{
								MarkdownDescription: `The maximum download limit (integer, in Kbps). null indicates no limit`,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
							},
							"limit_up": schema.Int64Attribute{
								MarkdownDescription: `The maximum upload limit (integer, in Kbps). null indicates no limit`,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"settings": schema.StringAttribute{
						MarkdownDescription: `How bandwidth limits are enforced. Can be 'network default', 'ignore' or 'custom'.
                                        Allowed values: [custom,ignore,network default]`,
						Optional: true,
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
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"rules": schema.ListNestedAttribute{
						MarkdownDescription: `A list of the Bonjour forwarding rules for your group policy. If 'settings' is set to 'custom', at least one rule must be specified.`,
						Optional:            true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"description": schema.StringAttribute{
									MarkdownDescription: `A description for your Bonjour forwarding rule. Optional.`,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"services": schema.ListAttribute{
									MarkdownDescription: `A list of Bonjour services. At least one service must be specified. Available services are 'All Services', 'AFP', 'AirPlay', 'Apple screen share', 'BitTorrent', 'Chromecast', 'FTP', 'iChat', 'iTunes', 'Printers', 'Samba', 'Scanners', 'Spotify' and 'SSH'`,
									Optional:            true,
									PlanModifiers: []planmodifier.List{
										listplanmodifier.UseStateForUnknown(),
									},

									ElementType: types.StringType,
								},
								"vlan_id": schema.StringAttribute{
									MarkdownDescription: `The ID of the service VLAN. Required.`,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
					},
					"settings": schema.StringAttribute{
						MarkdownDescription: `How Bonjour rules are applied. Can be 'network default', 'ignore' or 'custom'.
                                        Allowed values: [custom,ignore,network default]`,
						Optional: true,
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
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"allowed_url_patterns": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for allowed URL patterns`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"patterns": schema.ListAttribute{
								MarkdownDescription: `A list of URL patterns that are allowed`,
								Optional:            true,
								PlanModifiers: []planmodifier.List{
									listplanmodifier.UseStateForUnknown(),
								},

								ElementType: types.StringType,
							},
							"settings": schema.StringAttribute{
								MarkdownDescription: `How URL patterns are applied. Can be 'network default', 'append' or 'override'.
                                              Allowed values: [append,network default,override]`,
								Optional: true,
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
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"categories": schema.ListAttribute{
								MarkdownDescription: `A list of URL categories to block`,
								Optional:            true,
								PlanModifiers: []planmodifier.List{
									listplanmodifier.UseStateForUnknown(),
								},

								ElementType: types.StringType,
							},
							"settings": schema.StringAttribute{
								MarkdownDescription: `How URL categories are applied. Can be 'network default', 'append' or 'override'.
                                              Allowed values: [append,network default,override]`,
								Optional: true,
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
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"patterns": schema.ListAttribute{
								MarkdownDescription: `A list of URL patterns that are blocked`,
								Optional:            true,
								PlanModifiers: []planmodifier.List{
									listplanmodifier.UseStateForUnknown(),
								},

								ElementType: types.StringType,
							},
							"settings": schema.StringAttribute{
								MarkdownDescription: `How URL patterns are applied. Can be 'network default', 'append' or 'override'.
                                              Allowed values: [append,network default,override]`,
								Optional: true,
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

					"l3_firewall_rules": schema.ListNestedAttribute{
						MarkdownDescription: `An ordered array of the L3 firewall rules`,
						Optional:            true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"comment": schema.StringAttribute{
									MarkdownDescription: `Description of the rule (optional)`,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"dest_cidr": schema.StringAttribute{
									MarkdownDescription: `Destination IP address (in IP or CIDR notation), a fully-qualified domain name (FQDN, if your network supports it) or 'any'.`,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"dest_port": schema.StringAttribute{
									MarkdownDescription: `Destination port (integer in the range 1-65535), a port range (e.g. 8080-9090), or 'any'`,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"policy": schema.StringAttribute{
									MarkdownDescription: `'allow' or 'deny' traffic specified by this rule`,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"protocol": schema.StringAttribute{
									MarkdownDescription: `The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')`,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
					},
					"l7_firewall_rules": schema.ListNestedAttribute{
						MarkdownDescription: `An ordered array of L7 firewall rules`,
						Optional:            true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"policy": schema.StringAttribute{
									MarkdownDescription: `The policy applied to matching traffic. Must be 'deny'.
                                              Allowed values: [deny]`,
									Optional: true,
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
									MarkdownDescription: `Type of the L7 Rule. Must be 'application', 'applicationCategory', 'host', 'port' or 'ipRange'
                                              Allowed values: [application,applicationCategory,host,ipRange,port]`,
									Optional: true,
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
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
					},
					"settings": schema.StringAttribute{
						MarkdownDescription: `How firewall and traffic shaping rules are enforced. Can be 'network default', 'ignore' or 'custom'.
                                        Allowed values: [custom,ignore,network default]`,
						Optional: true,
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
					"traffic_shaping_rules": schema.ListNestedAttribute{
						MarkdownDescription: `    An array of traffic shaping rules. Rules are applied in the order that
    they are specified in. An empty list (or null) means no rules. Note that
    you are allowed a maximum of 8 rules.
`,
						Optional: true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"definitions": schema.ListNestedAttribute{
									MarkdownDescription: `    A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.
`,
									Optional: true,
									PlanModifiers: []planmodifier.List{
										listplanmodifier.UseStateForUnknown(),
									},
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"type": schema.StringAttribute{
												MarkdownDescription: `The type of definition. Can be one of 'application', 'applicationCategory', 'host', 'port', 'ipRange' or 'localNet'.
                                                    Allowed values: [application,applicationCategory,host,ipRange,localNet,port]`,
												Optional: true,
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
									Optional: true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"pcp_tag_value": schema.Int64Attribute{
									MarkdownDescription: `    The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority).
    null means 'Do not set PCP tag'.
`,
									Optional: true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"per_client_bandwidth_limits": schema.SingleNestedAttribute{
									MarkdownDescription: `    An object describing the bandwidth settings for your rule.
`,
									Optional: true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"bandwidth_limits": schema.SingleNestedAttribute{
											MarkdownDescription: `The bandwidth limits object, specifying the upload ('limitUp') and download ('limitDown') speed in Kbps. These are only enforced if 'settings' is set to 'custom'.`,
											Optional:            true,
											PlanModifiers: []planmodifier.Object{
												objectplanmodifier.UseStateForUnknown(),
											},
											Attributes: map[string]schema.Attribute{

												"limit_down": schema.Int64Attribute{
													MarkdownDescription: `The maximum download limit (integer, in Kbps).`,
													Optional:            true,
													PlanModifiers: []planmodifier.Int64{
														int64planmodifier.UseStateForUnknown(),
													},
												},
												"limit_up": schema.Int64Attribute{
													MarkdownDescription: `The maximum upload limit (integer, in Kbps).`,
													Optional:            true,
													PlanModifiers: []planmodifier.Int64{
														int64planmodifier.UseStateForUnknown(),
													},
												},
											},
										},
										"settings": schema.StringAttribute{
											MarkdownDescription: `How bandwidth limits are applied by your rule. Can be one of 'network default', 'ignore' or 'custom'.`,
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
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name for your group policy. Required.`,
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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether scheduling is enabled (true) or disabled (false). Defaults to false. If true, the schedule objects for each day of the week (monday - sunday) are parsed.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"friday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Friday.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
									&TimeFormatPlanModifier{},
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
									&TimeFormatPlanModifier{},
								},
							},
						},
					},
					"monday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Monday.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
									&TimeFormatPlanModifier{},
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
									&TimeFormatPlanModifier{},
								},
							},
						},
					},
					"saturday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Saturday.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
									&TimeFormatPlanModifier{},
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
									&TimeFormatPlanModifier{},
								},
							},
						},
					},
					"sunday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Sunday.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
									&TimeFormatPlanModifier{},
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
									&TimeFormatPlanModifier{},
								},
							},
						},
					},
					"thursday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Thursday.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
									&TimeFormatPlanModifier{},
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
									&TimeFormatPlanModifier{},
								},
							},
						},
					},
					"tuesday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Tuesday.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
									&TimeFormatPlanModifier{},
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
									&TimeFormatPlanModifier{},
								},
							},
						},
					},
					"wednesday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Wednesday.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
									&TimeFormatPlanModifier{},
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
									&TimeFormatPlanModifier{},
								},
							},
						},
					},
				},
			},
			"splash_auth_settings": schema.StringAttribute{
				MarkdownDescription: `Whether clients bound to your policy will bypass splash authorization or behave according to the network's rules. Can be one of 'network default' or 'bypass'. Only available if your network has a wireless configuration.
                                  Allowed values: [bypass,network default]`,
				Optional: true,
				Computed: true,
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
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"settings": schema.StringAttribute{
						MarkdownDescription: `How VLAN tagging is applied. Can be 'network default', 'ignore' or 'custom'.
                                        Allowed values: [custom,ignore,network default]`,
						Optional: true,
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
	// Has Paths
	vvNetworkID := data.NetworkID.ValueString()
	//Has Item and has items and post

	vvName := data.Name.ValueString()

	responseVerifyItem, restyResp1, err := r.client.Networks.GetNetworkGroupPolicies(vvNetworkID)
	//Has Post
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkGroupPolicies",
					restyResp1.String(),
				)
				return
			}
		}
	}

	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
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
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkGroupPolicy",
			err.Error(),
		)
		return
	}

	responseGet, restyResp1, err := r.client.Networks.GetNetworkGroupPolicies(vvNetworkID)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkGroupPolicies",
				restyResp1.String(),
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
				"Fail Parsing GroupPolicyID",
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
					restyRespGet.String(),
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

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	vvGroupPolicyID := data.GroupPolicyID.ValueString()
	responseGet, restyRespGet, err := r.client.Networks.GetNetworkGroupPolicy(vvNetworkID, vvGroupPolicyID)
	if err != nil || restyRespGet == nil {
		if restyRespGet != nil {
			if restyRespGet.StatusCode() == 404 {
				resp.Diagnostics.AddWarning(
					"Resource not found",
					"Deleting resource",
				)
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkGroupPolicy",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkGroupPolicy",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseNetworksGetNetworkGroupPolicyItemToBodyRs(data, responseGet, true)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *NetworksGroupPoliciesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: networkId,groupPolicyId. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("group_policy_id"), idParts[1])...)
}

func (r *NetworksGroupPoliciesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworksGroupPoliciesRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvNetworkID := plan.NetworkID.ValueString()
	vvGroupPolicyID := plan.GroupPolicyID.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkGroupPolicy(vvNetworkID, vvGroupPolicyID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkGroupPolicy",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkGroupPolicy",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
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
	_, err := r.client.Networks.DeleteNetworkGroupPolicy(vvNetworkID, vvGroupPolicyID, nil)
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
	Scheduling                *ResponseNetworksGetNetworkGroupPolicySchedulingRs                `tfsdk:"scheduling"`
	SplashAuthSettings        types.String                                                      `tfsdk:"splash_auth_settings"`
	VLANTagging               *ResponseNetworksGetNetworkGroupPolicyVlanTaggingRs               `tfsdk:"vlan_tagging"`
	Name                      types.String                                                      `tfsdk:"name"`
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
	Services    types.List   `tfsdk:"services"`
	VLANID      types.String `tfsdk:"vlan_id"`
}

type ResponseNetworksGetNetworkGroupPolicyContentFilteringRs struct {
	AllowedURLPatterns   *ResponseNetworksGetNetworkGroupPolicyContentFilteringAllowedUrlPatternsRs   `tfsdk:"allowed_url_patterns"`
	BlockedURLCategories *ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlCategoriesRs `tfsdk:"blocked_url_categories"`
	BlockedURLPatterns   *ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlPatternsRs   `tfsdk:"blocked_url_patterns"`
}

type ResponseNetworksGetNetworkGroupPolicyContentFilteringAllowedUrlPatternsRs struct {
	Patterns types.List   `tfsdk:"patterns"`
	Settings types.String `tfsdk:"settings"`
}

type ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlCategoriesRs struct {
	Categories types.List   `tfsdk:"categories"`
	Settings   types.String `tfsdk:"settings"`
}

type ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlPatternsRs struct {
	Patterns types.List   `tfsdk:"patterns"`
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
			//[debug] Is Array: False
		}
		settings := r.Bandwidth.Settings.ValueString()
		requestNetworksCreateNetworkGroupPolicyBandwidth = &merakigosdk.RequestNetworksCreateNetworkGroupPolicyBandwidth{
			BandwidthLimits: requestNetworksCreateNetworkGroupPolicyBandwidthBandwidthLimits,
			Settings:        settings,
		}
		//[debug] Is Array: False
	}
	var requestNetworksCreateNetworkGroupPolicyBonjourForwarding *merakigosdk.RequestNetworksCreateNetworkGroupPolicyBonjourForwarding

	if r.BonjourForwarding != nil {

		log.Printf("[DEBUG] #TODO []RequestNetworksCreateNetworkGroupPolicyBonjourForwardingRules")
		var requestNetworksCreateNetworkGroupPolicyBonjourForwardingRules []merakigosdk.RequestNetworksCreateNetworkGroupPolicyBonjourForwardingRules

		if r.BonjourForwarding.Rules != nil {
			for _, rItem1 := range *r.BonjourForwarding.Rules {
				description := rItem1.Description.ValueString()

				var services []string = nil
				rItem1.Services.ElementsAs(ctx, &services, false)
				vlanID := rItem1.VLANID.ValueString()
				requestNetworksCreateNetworkGroupPolicyBonjourForwardingRules = append(requestNetworksCreateNetworkGroupPolicyBonjourForwardingRules, merakigosdk.RequestNetworksCreateNetworkGroupPolicyBonjourForwardingRules{
					Description: description,
					Services:    services,
					VLANID:      vlanID,
				})
				//[debug] Is Array: True
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
		//[debug] Is Array: False
	}
	var requestNetworksCreateNetworkGroupPolicyContentFiltering *merakigosdk.RequestNetworksCreateNetworkGroupPolicyContentFiltering

	if r.ContentFiltering != nil {
		var requestNetworksCreateNetworkGroupPolicyContentFilteringAllowedURLPatterns *merakigosdk.RequestNetworksCreateNetworkGroupPolicyContentFilteringAllowedURLPatterns

		if r.ContentFiltering.AllowedURLPatterns != nil {

			var patterns []string = nil
			r.ContentFiltering.AllowedURLPatterns.Patterns.ElementsAs(ctx, &patterns, false)
			settings := r.ContentFiltering.AllowedURLPatterns.Settings.ValueString()
			requestNetworksCreateNetworkGroupPolicyContentFilteringAllowedURLPatterns = &merakigosdk.RequestNetworksCreateNetworkGroupPolicyContentFilteringAllowedURLPatterns{
				Patterns: patterns,
				Settings: settings,
			}
			//[debug] Is Array: False
		}
		var requestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLCategories *merakigosdk.RequestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLCategories

		if r.ContentFiltering.BlockedURLCategories != nil {

			var categories []string = nil
			r.ContentFiltering.BlockedURLCategories.Categories.ElementsAs(ctx, &categories, false)
			settings := r.ContentFiltering.BlockedURLCategories.Settings.ValueString()
			requestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLCategories = &merakigosdk.RequestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLCategories{
				Categories: categories,
				Settings:   settings,
			}
			//[debug] Is Array: False
		}
		var requestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLPatterns *merakigosdk.RequestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLPatterns

		if r.ContentFiltering.BlockedURLPatterns != nil {

			var patterns []string = nil
			r.ContentFiltering.BlockedURLPatterns.Patterns.ElementsAs(ctx, &patterns, false)
			settings := r.ContentFiltering.BlockedURLPatterns.Settings.ValueString()
			requestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLPatterns = &merakigosdk.RequestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLPatterns{
				Patterns: patterns,
				Settings: settings,
			}
			//[debug] Is Array: False
		}
		requestNetworksCreateNetworkGroupPolicyContentFiltering = &merakigosdk.RequestNetworksCreateNetworkGroupPolicyContentFiltering{
			AllowedURLPatterns:   requestNetworksCreateNetworkGroupPolicyContentFilteringAllowedURLPatterns,
			BlockedURLCategories: requestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLCategories,
			BlockedURLPatterns:   requestNetworksCreateNetworkGroupPolicyContentFilteringBlockedURLPatterns,
		}
		//[debug] Is Array: False
	}
	var requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShaping *merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShaping

	if r.FirewallAndTrafficShaping != nil {

		log.Printf("[DEBUG] #TODO []RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules")
		var requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules []merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules

		if r.FirewallAndTrafficShaping.L3FirewallRules != nil {
			for _, rItem1 := range *r.FirewallAndTrafficShaping.L3FirewallRules {
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
				//[debug] Is Array: True
			}
		}

		log.Printf("[DEBUG] #TODO []RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules")
		var requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules []merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules

		if r.FirewallAndTrafficShaping.L7FirewallRules != nil {
			for _, rItem1 := range *r.FirewallAndTrafficShaping.L7FirewallRules {
				policy := rItem1.Policy.ValueString()
				typeR := rItem1.Type.ValueString()
				value := rItem1.Value.ValueString()
				requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules = append(requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules, merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules{
					Policy: policy,
					Type:   typeR,
					Value:  value,
				})
				//[debug] Is Array: True
			}
		}
		settings := r.FirewallAndTrafficShaping.Settings.ValueString()

		log.Printf("[DEBUG] #TODO []RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules")
		var requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules []merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules

		if r.FirewallAndTrafficShaping.TrafficShapingRules != nil {
			for _, rItem1 := range *r.FirewallAndTrafficShaping.TrafficShapingRules {

				log.Printf("[DEBUG] #TODO []RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions")
				var requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions []merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions

				if rItem1.Definitions != nil {
					for _, rItem2 := range *rItem1.Definitions {
						typeR := rItem2.Type.ValueString()
						value := rItem2.Value.ValueString()
						requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions = append(requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions, merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions{
							Type:  typeR,
							Value: value,
						})
						//[debug] Is Array: True
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

				if rItem1.PerClientBandwidthLimits != nil {
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
						//[debug] Is Array: False
					}
					settings := rItem1.PerClientBandwidthLimits.Settings.ValueString()
					requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits = &merakigosdk.RequestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits{
						BandwidthLimits: requestNetworksCreateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits,
						Settings:        settings,
					}
					//[debug] Is Array: False
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
				//[debug] Is Array: True
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
		//[debug] Is Array: False
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
			//[debug] Is Array: False
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
			//[debug] Is Array: False
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
			//[debug] Is Array: False
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
			//[debug] Is Array: False
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
			//[debug] Is Array: False
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
			//[debug] Is Array: False
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
			//[debug] Is Array: False
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
		//[debug] Is Array: False
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
		vlanID := r.VLANTagging.VLANID.ValueString()
		requestNetworksCreateNetworkGroupPolicyVLANTagging = &merakigosdk.RequestNetworksCreateNetworkGroupPolicyVLANTagging{
			Settings: settings,
			VLANID:   vlanID,
		}
		//[debug] Is Array: False
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
			//[debug] Is Array: False
		}
		settings := r.Bandwidth.Settings.ValueString()
		requestNetworksUpdateNetworkGroupPolicyBandwidth = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicyBandwidth{
			BandwidthLimits: requestNetworksUpdateNetworkGroupPolicyBandwidthBandwidthLimits,
			Settings:        settings,
		}
		//[debug] Is Array: False
	}
	var requestNetworksUpdateNetworkGroupPolicyBonjourForwarding *merakigosdk.RequestNetworksUpdateNetworkGroupPolicyBonjourForwarding

	if r.BonjourForwarding != nil {

		log.Printf("[DEBUG] #TODO []RequestNetworksUpdateNetworkGroupPolicyBonjourForwardingRules")
		var requestNetworksUpdateNetworkGroupPolicyBonjourForwardingRules []merakigosdk.RequestNetworksUpdateNetworkGroupPolicyBonjourForwardingRules

		if r.BonjourForwarding.Rules != nil {
			for _, rItem1 := range *r.BonjourForwarding.Rules {
				description := rItem1.Description.ValueString()

				var services []string = nil
				rItem1.Services.ElementsAs(ctx, &services, false)
				vlanID := rItem1.VLANID.ValueString()
				requestNetworksUpdateNetworkGroupPolicyBonjourForwardingRules = append(requestNetworksUpdateNetworkGroupPolicyBonjourForwardingRules, merakigosdk.RequestNetworksUpdateNetworkGroupPolicyBonjourForwardingRules{
					Description: description,
					Services:    services,
					VLANID:      vlanID,
				})
				//[debug] Is Array: True
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
		//[debug] Is Array: False
	}
	var requestNetworksUpdateNetworkGroupPolicyContentFiltering *merakigosdk.RequestNetworksUpdateNetworkGroupPolicyContentFiltering

	if r.ContentFiltering != nil {
		var requestNetworksUpdateNetworkGroupPolicyContentFilteringAllowedURLPatterns *merakigosdk.RequestNetworksUpdateNetworkGroupPolicyContentFilteringAllowedURLPatterns

		if r.ContentFiltering.AllowedURLPatterns != nil {

			var patterns []string = nil
			r.ContentFiltering.AllowedURLPatterns.Patterns.ElementsAs(ctx, &patterns, false)
			settings := r.ContentFiltering.AllowedURLPatterns.Settings.ValueString()
			requestNetworksUpdateNetworkGroupPolicyContentFilteringAllowedURLPatterns = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicyContentFilteringAllowedURLPatterns{
				Patterns: patterns,
				Settings: settings,
			}
			//[debug] Is Array: False
		}
		var requestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLCategories *merakigosdk.RequestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLCategories

		if r.ContentFiltering.BlockedURLCategories != nil {

			var categories []string = nil
			r.ContentFiltering.BlockedURLCategories.Categories.ElementsAs(ctx, &categories, false)
			settings := r.ContentFiltering.BlockedURLCategories.Settings.ValueString()
			requestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLCategories = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLCategories{
				Categories: categories,
				Settings:   settings,
			}
			//[debug] Is Array: False
		}
		var requestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLPatterns *merakigosdk.RequestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLPatterns

		if r.ContentFiltering.BlockedURLPatterns != nil {

			var patterns []string = nil
			r.ContentFiltering.BlockedURLPatterns.Patterns.ElementsAs(ctx, &patterns, false)
			settings := r.ContentFiltering.BlockedURLPatterns.Settings.ValueString()
			requestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLPatterns = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLPatterns{
				Patterns: patterns,
				Settings: settings,
			}
			//[debug] Is Array: False
		}
		requestNetworksUpdateNetworkGroupPolicyContentFiltering = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicyContentFiltering{
			AllowedURLPatterns:   requestNetworksUpdateNetworkGroupPolicyContentFilteringAllowedURLPatterns,
			BlockedURLCategories: requestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLCategories,
			BlockedURLPatterns:   requestNetworksUpdateNetworkGroupPolicyContentFilteringBlockedURLPatterns,
		}
		//[debug] Is Array: False
	}
	var requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShaping *merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShaping

	if r.FirewallAndTrafficShaping != nil {

		log.Printf("[DEBUG] #TODO []RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules")
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
				//[debug] Is Array: True
			}
		}

		log.Printf("[DEBUG] #TODO []RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules")
		var requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules []merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules

		if r.FirewallAndTrafficShaping.L7FirewallRules != nil {
			for _, rItem1 := range *r.FirewallAndTrafficShaping.L7FirewallRules {
				policy := rItem1.Policy.ValueString()
				typeR := rItem1.Type.ValueString()
				value := rItem1.Value.ValueString()
				requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules = append(requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules, merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules{
					Policy: policy,
					Type:   typeR,
					Value:  value,
				})
				//[debug] Is Array: True
			}
		}
		settings := r.FirewallAndTrafficShaping.Settings.ValueString()

		log.Printf("[DEBUG] #TODO []RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules")
		var requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules []merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules

		if r.FirewallAndTrafficShaping.TrafficShapingRules != nil {
			for _, rItem1 := range *r.FirewallAndTrafficShaping.TrafficShapingRules {

				log.Printf("[DEBUG] #TODO []RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions")
				var requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions []merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions

				if rItem1.Definitions != nil {
					for _, rItem2 := range *rItem1.Definitions {
						typeR := rItem2.Type.ValueString()
						value := rItem2.Value.ValueString()
						requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions = append(requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions, merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions{
							Type:  typeR,
							Value: value,
						})
						//[debug] Is Array: True
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

				if rItem1.PerClientBandwidthLimits != nil {
					var requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits *merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits

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
						requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits{
							LimitDown: int64ToIntPointer(limitDown),
							LimitUp:   int64ToIntPointer(limitUp),
						}
						//[debug] Is Array: False
					}
					settings := rItem1.PerClientBandwidthLimits.Settings.ValueString()
					requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits{
						BandwidthLimits: requestNetworksUpdateNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits,
						Settings:        settings,
					}
					//[debug] Is Array: False
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
				//[debug] Is Array: True
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
		//[debug] Is Array: False
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
			//[debug] Is Array: False
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
			//[debug] Is Array: False
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
			//[debug] Is Array: False
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
			//[debug] Is Array: False
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
			//[debug] Is Array: False
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
			//[debug] Is Array: False
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
			//[debug] Is Array: False
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
		//[debug] Is Array: False
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
		vlanID := r.VLANTagging.VLANID.ValueString()
		requestNetworksUpdateNetworkGroupPolicyVLANTagging = &merakigosdk.RequestNetworksUpdateNetworkGroupPolicyVLANTagging{
			Settings: settings,
			VLANID:   vlanID,
		}
		//[debug] Is Array: False
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
						return nil
					}(),
					Settings: func() types.String {
						if response.Bandwidth.Settings != "" {
							return types.StringValue(response.Bandwidth.Settings)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
		BonjourForwarding: func() *ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRs {
			if response.BonjourForwarding != nil {
				return &ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRs{
					Rules: func() *[]ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRulesRs {
						if response.BonjourForwarding.Rules != nil {
							result := make([]ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRulesRs, len(*response.BonjourForwarding.Rules))
							for i, rules := range *response.BonjourForwarding.Rules {
								result[i] = ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRulesRs{
									Description: func() types.String {
										if rules.Description != "" {
											return types.StringValue(rules.Description)
										}
										return types.String{}
									}(),
									Services: StringSliceToList(rules.Services),
									VLANID: func() types.String {
										if rules.VLANID != "" {
											return types.StringValue(rules.VLANID)
										}
										return types.String{}
									}(),
								}
							}
							return &result
						}
						return nil
					}(),
					Settings: func() types.String {
						if response.BonjourForwarding.Settings != "" {
							return types.StringValue(response.BonjourForwarding.Settings)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
		ContentFiltering: func() *ResponseNetworksGetNetworkGroupPolicyContentFilteringRs {
			if response.ContentFiltering != nil {
				return &ResponseNetworksGetNetworkGroupPolicyContentFilteringRs{
					AllowedURLPatterns: func() *ResponseNetworksGetNetworkGroupPolicyContentFilteringAllowedUrlPatternsRs {
						if response.ContentFiltering.AllowedURLPatterns != nil {
							return &ResponseNetworksGetNetworkGroupPolicyContentFilteringAllowedUrlPatternsRs{
								Patterns: StringSliceToList(response.ContentFiltering.AllowedURLPatterns.Patterns),
								Settings: func() types.String {
									if response.ContentFiltering.AllowedURLPatterns.Settings != "" {
										return types.StringValue(response.ContentFiltering.AllowedURLPatterns.Settings)
									}
									return types.String{}
								}(),
							}
						}
						return nil
					}(),
					BlockedURLCategories: func() *ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlCategoriesRs {
						if response.ContentFiltering.BlockedURLCategories != nil {
							return &ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlCategoriesRs{
								Categories: StringSliceToList(response.ContentFiltering.BlockedURLCategories.Categories),
								Settings: func() types.String {
									if response.ContentFiltering.BlockedURLCategories.Settings != "" {
										return types.StringValue(response.ContentFiltering.BlockedURLCategories.Settings)
									}
									return types.String{}
								}(),
							}
						}
						return nil
					}(),
					BlockedURLPatterns: func() *ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlPatternsRs {
						if response.ContentFiltering.BlockedURLPatterns != nil {
							return &ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlPatternsRs{
								Patterns: StringSliceToList(response.ContentFiltering.BlockedURLPatterns.Patterns),
								Settings: func() types.String {
									if response.ContentFiltering.BlockedURLPatterns.Settings != "" {
										return types.StringValue(response.ContentFiltering.BlockedURLPatterns.Settings)
									}
									return types.String{}
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		FirewallAndTrafficShaping: func() *ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingRs {
			if response.FirewallAndTrafficShaping != nil {
				return &ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingRs{
					L3FirewallRules: func() *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRulesRs {
						if response.FirewallAndTrafficShaping.L3FirewallRules != nil {
							result := make([]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRulesRs, len(*response.FirewallAndTrafficShaping.L3FirewallRules))
							for i, l3FirewallRules := range *response.FirewallAndTrafficShaping.L3FirewallRules {
								result[i] = ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRulesRs{
									Comment: func() types.String {
										if l3FirewallRules.Comment != "" {
											return types.StringValue(l3FirewallRules.Comment)
										}
										return types.String{}
									}(),
									DestCidr: func() types.String {
										if l3FirewallRules.DestCidr != "" {
											return types.StringValue(l3FirewallRules.DestCidr)
										}
										return types.String{}
									}(),
									DestPort: func() types.String {
										if l3FirewallRules.DestPort != "" {
											return types.StringValue(l3FirewallRules.DestPort)
										}
										return types.String{}
									}(),
									Policy: func() types.String {
										if l3FirewallRules.Policy != "" {
											return types.StringValue(l3FirewallRules.Policy)
										}
										return types.String{}
									}(),
									Protocol: func() types.String {
										if l3FirewallRules.Protocol != "" {
											return types.StringValue(l3FirewallRules.Protocol)
										}
										return types.String{}
									}(),
								}
							}
							return &result
						}
						return nil
					}(),
					L7FirewallRules: func() *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRulesRs {
						if response.FirewallAndTrafficShaping.L7FirewallRules != nil {
							result := make([]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRulesRs, len(*response.FirewallAndTrafficShaping.L7FirewallRules))
							for i, l7FirewallRules := range *response.FirewallAndTrafficShaping.L7FirewallRules {
								result[i] = ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRulesRs{
									Policy: func() types.String {
										if l7FirewallRules.Policy != "" {
											return types.StringValue(l7FirewallRules.Policy)
										}
										return types.String{}
									}(),
									Type: func() types.String {
										if l7FirewallRules.Type != "" {
											return types.StringValue(l7FirewallRules.Type)
										}
										return types.String{}
									}(),
									Value: func() types.String {
										if l7FirewallRules.Value != "" {
											return types.StringValue(l7FirewallRules.Value)
										}
										return types.String{}
									}(),
								}
							}
							return &result
						}
						return nil
					}(),
					Settings: func() types.String {
						if response.FirewallAndTrafficShaping.Settings != "" {
							return types.StringValue(response.FirewallAndTrafficShaping.Settings)
						}
						return types.String{}
					}(),
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
													Type: func() types.String {
														if definitions.Type != "" {
															return types.StringValue(definitions.Type)
														}
														return types.String{}
													}(),
													Value: func() types.String {
														if definitions.Value != "" {
															return types.StringValue(definitions.Value)
														}
														return types.String{}
													}(),
												}
											}
											return &result
										}
										return nil
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
													return nil
												}(),
												Settings: func() types.String {
													if trafficShapingRules.PerClientBandwidthLimits.Settings != "" {
														return types.StringValue(trafficShapingRules.PerClientBandwidthLimits.Settings)
													}
													return types.String{}
												}(),
											}
										}
										return nil
									}(),
									Priority: func() types.String {
										firewallAndTrafficShaping := state.FirewallAndTrafficShaping
										if firewallAndTrafficShaping != nil {
											if firewallAndTrafficShaping.TrafficShapingRules != nil {
												if len(*firewallAndTrafficShaping.TrafficShapingRules) > i {
													return types.StringValue((*firewallAndTrafficShaping.TrafficShapingRules)[i].Priority.ValueString())
												}
											}
										}
										return types.String{}
									}(),
								}
							}
							return &result
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		GroupPolicyID: func() types.String {
			if response.GroupPolicyID != "" {
				return types.StringValue(response.GroupPolicyID)
			}
			return types.String{}
		}(),
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
								From: func() types.String {
									if response.Scheduling.Friday.From != "" {
										return types.StringValue(response.Scheduling.Friday.From)
									}
									return types.String{}
								}(),
								To: func() types.String {
									if response.Scheduling.Friday.To != "" {
										return types.StringValue(response.Scheduling.Friday.To)
									}
									return types.String{}
								}(),
							}
						}
						return nil
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
								From: func() types.String {
									if response.Scheduling.Monday.From != "" {
										return types.StringValue(response.Scheduling.Monday.From)
									}
									return types.String{}
								}(),
								To: func() types.String {
									if response.Scheduling.Monday.To != "" {
										return types.StringValue(response.Scheduling.Monday.To)
									}
									return types.String{}
								}(),
							}
						}
						return nil
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
								From: func() types.String {
									if response.Scheduling.Saturday.From != "" {
										return types.StringValue(response.Scheduling.Saturday.From)
									}
									return types.String{}
								}(),
								To: func() types.String {
									if response.Scheduling.Saturday.To != "" {
										return types.StringValue(response.Scheduling.Saturday.To)
									}
									return types.String{}
								}(),
							}
						}
						return nil
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
								From: func() types.String {
									if response.Scheduling.Sunday.From != "" {
										return types.StringValue(response.Scheduling.Sunday.From)
									}
									return types.String{}
								}(),
								To: func() types.String {
									if response.Scheduling.Sunday.To != "" {
										return types.StringValue(response.Scheduling.Sunday.To)
									}
									return types.String{}
								}(),
							}
						}
						return nil
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
								From: func() types.String {
									if response.Scheduling.Thursday.From != "" {
										return types.StringValue(response.Scheduling.Thursday.From)
									}
									return types.String{}
								}(),
								To: func() types.String {
									if response.Scheduling.Thursday.To != "" {
										return types.StringValue(response.Scheduling.Thursday.To)
									}
									return types.String{}
								}(),
							}
						}
						return nil
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
								From: func() types.String {
									if response.Scheduling.Tuesday.From != "" {
										return types.StringValue(response.Scheduling.Tuesday.From)
									}
									return types.String{}
								}(),
								To: func() types.String {
									if response.Scheduling.Tuesday.To != "" {
										return types.StringValue(response.Scheduling.Tuesday.To)
									}
									return types.String{}
								}(),
							}
						}
						return nil
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
								From: func() types.String {
									if response.Scheduling.Wednesday.From != "" {
										return types.StringValue(response.Scheduling.Wednesday.From)
									}
									return types.String{}
								}(),
								To: func() types.String {
									if response.Scheduling.Wednesday.To != "" {
										return types.StringValue(response.Scheduling.Wednesday.To)
									}
									return types.String{}
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		SplashAuthSettings: func() types.String {
			if response.SplashAuthSettings != "" {
				return types.StringValue(response.SplashAuthSettings)
			}
			return types.String{}
		}(),
		VLANTagging: func() *ResponseNetworksGetNetworkGroupPolicyVlanTaggingRs {
			if response.VLANTagging != nil {
				return &ResponseNetworksGetNetworkGroupPolicyVlanTaggingRs{
					Settings: func() types.String {
						if response.VLANTagging.Settings != "" {
							return types.StringValue(response.VLANTagging.Settings)
						}
						return types.String{}
					}(),
					VLANID: func() types.String {
						if response.VLANTagging.VLANID != "" {
							return types.StringValue(response.VLANTagging.VLANID)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
	}
	// Set Priority

	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksGroupPoliciesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksGroupPoliciesRs)
}

// TimeFormatPlanModifier is a plan modifier that normalizes time formats
// to handle variations like "09:00" vs "9:00" as equivalent
type TimeFormatPlanModifier struct{}

func (m *TimeFormatPlanModifier) Description(ctx context.Context) string {
	return "Normalizes time formats to handle variations like '09:00' vs '9:00' as equivalent"
}

func (m *TimeFormatPlanModifier) MarkdownDescription(ctx context.Context) string {
	return "Normalizes time formats to handle variations like '09:00' vs '9:00' as equivalent"
}

func (m *TimeFormatPlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// If the plan value is unknown, don't modify it
	if req.PlanValue.IsUnknown() {
		return
	}

	// If the state value is unknown, don't modify it
	if req.StateValue.IsUnknown() {
		return
	}

	planValue := req.PlanValue.ValueString()
	stateValue := req.StateValue.ValueString()

	// Normalize both values to a standard format
	normalizedPlan := normalizeTimeFormat(planValue)
	normalizedState := normalizeTimeFormat(stateValue)

	// If both normalized values are the same, use the state value to avoid unnecessary changes
	if normalizedPlan == normalizedState {
		resp.PlanValue = req.StateValue
		return
	}

	// Otherwise, keep the plan value as is
	resp.PlanValue = req.PlanValue
}

// normalizeTimeFormat converts time strings to a standard format
// "09:00" -> "9:00", "9:00" -> "9:00"
func normalizeTimeFormat(timeStr string) string {
	if timeStr == "" {
		return timeStr
	}

	// Split by colon to handle HH:MM format
	parts := strings.Split(timeStr, ":")
	if len(parts) != 2 {
		return timeStr // Return as-is if not in HH:MM format
	}

	// Remove leading zeros from hours
	hours := strings.TrimLeft(parts[0], "0")
	if hours == "" {
		hours = "0"
	}

	// Keep minutes as-is (with leading zeros if present)
	minutes := parts[1]

	return hours + ":" + minutes
}
