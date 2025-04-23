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

// DATA SOURCE NORMAL
import (
	"context"
	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksGroupPoliciesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksGroupPoliciesDataSource{}
)

func NewNetworksGroupPoliciesDataSource() datasource.DataSource {
	return &NetworksGroupPoliciesDataSource{}
}

type NetworksGroupPoliciesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksGroupPoliciesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksGroupPoliciesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_group_policies"
}

func (d *NetworksGroupPoliciesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"group_policy_id": schema.StringAttribute{
				MarkdownDescription: `groupPolicyId path parameter. Group policy ID`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"bandwidth": schema.SingleNestedAttribute{
						MarkdownDescription: `    The bandwidth settings for clients bound to your group policy.
`,
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"bandwidth_limits": schema.SingleNestedAttribute{
								MarkdownDescription: `The bandwidth limits object, specifying upload and download speed for clients bound to the group policy. These are only enforced if 'settings' is set to 'custom'.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"limit_down": schema.Int64Attribute{
										MarkdownDescription: `The maximum download limit (integer, in Kbps). null indicates no limit`,
										Computed:            true,
									},
									"limit_up": schema.Int64Attribute{
										MarkdownDescription: `The maximum upload limit (integer, in Kbps). null indicates no limit`,
										Computed:            true,
									},
								},
							},
							"settings": schema.StringAttribute{
								MarkdownDescription: `How bandwidth limits are enforced. Can be 'network default', 'ignore' or 'custom'.`,
								Computed:            true,
							},
						},
					},
					"bonjour_forwarding": schema.SingleNestedAttribute{
						MarkdownDescription: `The Bonjour settings for your group policy. Only valid if your network has a wireless configuration.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"rules": schema.SetNestedAttribute{
								MarkdownDescription: `A list of the Bonjour forwarding rules for your group policy. If 'settings' is set to 'custom', at least one rule must be specified.`,
								Computed:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{

										"description": schema.StringAttribute{
											MarkdownDescription: `A description for your Bonjour forwarding rule. Optional.`,
											Computed:            true,
										},
										"services": schema.ListAttribute{
											MarkdownDescription: `A list of Bonjour services. At least one service must be specified. Available services are 'All Services', 'AFP', 'AirPlay', 'Apple screen share', 'BitTorrent', 'Chromecast', 'FTP', 'iChat', 'iTunes', 'Printers', 'Samba', 'Scanners', 'Spotify' and 'SSH'`,
											Computed:            true,
											ElementType:         types.StringType,
										},
										"vlan_id": schema.StringAttribute{
											MarkdownDescription: `The ID of the service VLAN. Required.`,
											Computed:            true,
										},
									},
								},
							},
							"settings": schema.StringAttribute{
								MarkdownDescription: `How Bonjour rules are applied. Can be 'network default', 'ignore' or 'custom'.`,
								Computed:            true,
							},
						},
					},
					"content_filtering": schema.SingleNestedAttribute{
						MarkdownDescription: `The content filtering settings for your group policy`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"allowed_url_patterns": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for allowed URL patterns`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"patterns": schema.ListAttribute{
										MarkdownDescription: `A list of URL patterns that are allowed`,
										Computed:            true,
										ElementType:         types.StringType,
									},
									"settings": schema.StringAttribute{
										MarkdownDescription: `How URL patterns are applied. Can be 'network default', 'append' or 'override'.`,
										Computed:            true,
									},
								},
							},
							"blocked_url_categories": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for blocked URL categories`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"categories": schema.ListAttribute{
										MarkdownDescription: `A list of URL categories to block`,
										Computed:            true,
										ElementType:         types.StringType,
									},
									"settings": schema.StringAttribute{
										MarkdownDescription: `How URL categories are applied. Can be 'network default', 'append' or 'override'.`,
										Computed:            true,
									},
								},
							},
							"blocked_url_patterns": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for blocked URL patterns`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"patterns": schema.ListAttribute{
										MarkdownDescription: `A list of URL patterns that are blocked`,
										Computed:            true,
										ElementType:         types.StringType,
									},
									"settings": schema.StringAttribute{
										MarkdownDescription: `How URL patterns are applied. Can be 'network default', 'append' or 'override'.`,
										Computed:            true,
									},
								},
							},
						},
					},
					"firewall_and_traffic_shaping": schema.SingleNestedAttribute{
						MarkdownDescription: `    The firewall and traffic shaping rules and settings for your policy.
`,
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"l3_firewall_rules": schema.SetNestedAttribute{
								MarkdownDescription: `An ordered array of the L3 firewall rules`,
								Computed:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{

										"comment": schema.StringAttribute{
											MarkdownDescription: `Description of the rule (optional)`,
											Computed:            true,
										},
										"dest_cidr": schema.StringAttribute{
											MarkdownDescription: `Destination IP address (in IP or CIDR notation), a fully-qualified domain name (FQDN, if your network supports it) or 'any'.`,
											Computed:            true,
										},
										"dest_port": schema.StringAttribute{
											MarkdownDescription: `Destination port (integer in the range 1-65535), a port range (e.g. 8080-9090), or 'any'`,
											Computed:            true,
										},
										"policy": schema.StringAttribute{
											MarkdownDescription: `'allow' or 'deny' traffic specified by this rule`,
											Computed:            true,
										},
										"protocol": schema.StringAttribute{
											MarkdownDescription: `The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')`,
											Computed:            true,
										},
									},
								},
							},
							"l7_firewall_rules": schema.SetNestedAttribute{
								MarkdownDescription: `An ordered array of L7 firewall rules`,
								Computed:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{

										"policy": schema.StringAttribute{
											MarkdownDescription: `The policy applied to matching traffic. Must be 'deny'.`,
											Computed:            true,
										},
										"type": schema.StringAttribute{
											MarkdownDescription: `Type of the L7 Rule. Must be 'application', 'applicationCategory', 'host', 'port' or 'ipRange'`,
											Computed:            true,
										},
										"value": schema.StringAttribute{
											MarkdownDescription: `The 'value' of what you want to block. If 'type' is 'host', 'port' or 'ipRange', 'value' must be a string matching either a hostname (e.g. somewhere.com), a port (e.g. 8080), or an IP range (e.g. 192.1.0.0/16). If 'type' is 'application' or 'applicationCategory', then 'value' must be an object with an ID for the application.`,
											Computed:            true,
										},
									},
								},
							},
							"settings": schema.StringAttribute{
								MarkdownDescription: `How firewall and traffic shaping rules are enforced. Can be 'network default', 'ignore' or 'custom'.`,
								Computed:            true,
							},
							"traffic_shaping_rules": schema.SetNestedAttribute{
								MarkdownDescription: `    An array of traffic shaping rules. Rules are applied in the order that
    they are specified in. An empty list (or null) means no rules. Note that
    you are allowed a maximum of 8 rules.
`,
								Computed: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{

										"definitions": schema.SetNestedAttribute{
											MarkdownDescription: `    A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.
`,
											Computed: true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{

													"type": schema.StringAttribute{
														MarkdownDescription: `The type of definition. Can be one of 'application', 'applicationCategory', 'host', 'port', 'ipRange' or 'localNet'.`,
														Computed:            true,
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
													},
												},
											},
										},
										"dscp_tag_value": schema.Int64Attribute{
											MarkdownDescription: `    The DSCP tag applied by your rule. null means 'Do not change DSCP tag'.
    For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.
`,
											Computed: true,
										},
										"pcp_tag_value": schema.Int64Attribute{
											MarkdownDescription: `    The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority).
    null means 'Do not set PCP tag'.
`,
											Computed: true,
										},
										"per_client_bandwidth_limits": schema.SingleNestedAttribute{
											MarkdownDescription: `    An object describing the bandwidth settings for your rule.
`,
											Computed: true,
											Attributes: map[string]schema.Attribute{

												"bandwidth_limits": schema.SingleNestedAttribute{
													MarkdownDescription: `The bandwidth limits object, specifying the upload ('limitUp') and download ('limitDown') speed in Kbps. These are only enforced if 'settings' is set to 'custom'.`,
													Computed:            true,
													Attributes: map[string]schema.Attribute{

														"limit_down": schema.Int64Attribute{
															MarkdownDescription: `The maximum download limit (integer, in Kbps).`,
															Computed:            true,
														},
														"limit_up": schema.Int64Attribute{
															MarkdownDescription: `The maximum upload limit (integer, in Kbps).`,
															Computed:            true,
														},
													},
												},
												"settings": schema.StringAttribute{
													MarkdownDescription: `How bandwidth limits are applied by your rule. Can be one of 'network default', 'ignore' or 'custom'.`,
													Computed:            true,
												},
											},
										},
										"priority": schema.StringAttribute{
											MarkdownDescription: `    A string, indicating the priority level for packets bound to your rule.
    Can be 'low', 'normal' or 'high'.
`,
											Computed: true,
										},
									},
								},
							},
						},
					},
					"group_policy_id": schema.StringAttribute{
						MarkdownDescription: `The ID of the group policy`,
						Computed:            true,
					},
					"scheduling": schema.SingleNestedAttribute{
						MarkdownDescription: `    The schedule for the group policy. Schedules are applied to days of the week.
`,
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Whether scheduling is enabled (true) or disabled (false). Defaults to false. If true, the schedule objects for each day of the week (monday - sunday) are parsed.`,
								Computed:            true,
							},
							"friday": schema.SingleNestedAttribute{
								MarkdownDescription: `The schedule object for Friday.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"active": schema.BoolAttribute{
										MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
										Computed:            true,
									},
									"from": schema.StringAttribute{
										MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
										Computed:            true,
									},
									"to": schema.StringAttribute{
										MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
										Computed:            true,
									},
								},
							},
							"monday": schema.SingleNestedAttribute{
								MarkdownDescription: `The schedule object for Monday.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"active": schema.BoolAttribute{
										MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
										Computed:            true,
									},
									"from": schema.StringAttribute{
										MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
										Computed:            true,
									},
									"to": schema.StringAttribute{
										MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
										Computed:            true,
									},
								},
							},
							"saturday": schema.SingleNestedAttribute{
								MarkdownDescription: `The schedule object for Saturday.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"active": schema.BoolAttribute{
										MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
										Computed:            true,
									},
									"from": schema.StringAttribute{
										MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
										Computed:            true,
									},
									"to": schema.StringAttribute{
										MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
										Computed:            true,
									},
								},
							},
							"sunday": schema.SingleNestedAttribute{
								MarkdownDescription: `The schedule object for Sunday.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"active": schema.BoolAttribute{
										MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
										Computed:            true,
									},
									"from": schema.StringAttribute{
										MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
										Computed:            true,
									},
									"to": schema.StringAttribute{
										MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
										Computed:            true,
									},
								},
							},
							"thursday": schema.SingleNestedAttribute{
								MarkdownDescription: `The schedule object for Thursday.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"active": schema.BoolAttribute{
										MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
										Computed:            true,
									},
									"from": schema.StringAttribute{
										MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
										Computed:            true,
									},
									"to": schema.StringAttribute{
										MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
										Computed:            true,
									},
								},
							},
							"tuesday": schema.SingleNestedAttribute{
								MarkdownDescription: `The schedule object for Tuesday.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"active": schema.BoolAttribute{
										MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
										Computed:            true,
									},
									"from": schema.StringAttribute{
										MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
										Computed:            true,
									},
									"to": schema.StringAttribute{
										MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
										Computed:            true,
									},
								},
							},
							"wednesday": schema.SingleNestedAttribute{
								MarkdownDescription: `The schedule object for Wednesday.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"active": schema.BoolAttribute{
										MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
										Computed:            true,
									},
									"from": schema.StringAttribute{
										MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
										Computed:            true,
									},
									"to": schema.StringAttribute{
										MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
										Computed:            true,
									},
								},
							},
						},
					},
					"splash_auth_settings": schema.StringAttribute{
						MarkdownDescription: `Whether clients bound to your policy will bypass splash authorization or behave according to the network's rules. Can be one of 'network default' or 'bypass'. Only available if your network has a wireless configuration.`,
						Computed:            true,
					},
					"vlan_tagging": schema.SingleNestedAttribute{
						MarkdownDescription: `The VLAN tagging settings for your group policy. Only available if your network has a wireless configuration.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"settings": schema.StringAttribute{
								MarkdownDescription: `How VLAN tagging is applied. Can be 'network default', 'ignore' or 'custom'.`,
								Computed:            true,
							},
							"vlan_id": schema.StringAttribute{
								MarkdownDescription: `The ID of the vlan you want to tag. This only applies if 'settings' is set to 'custom'.`,
								Computed:            true,
							},
						},
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseNetworksGetNetworkGroupPolicies`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"bandwidth": schema.SingleNestedAttribute{
							MarkdownDescription: `    The bandwidth settings for clients bound to your group policy.
`,
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"bandwidth_limits": schema.SingleNestedAttribute{
									MarkdownDescription: `The bandwidth limits object, specifying upload and download speed for clients bound to the group policy. These are only enforced if 'settings' is set to 'custom'.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"limit_down": schema.Int64Attribute{
											MarkdownDescription: `The maximum download limit (integer, in Kbps). null indicates no limit`,
											Computed:            true,
										},
										"limit_up": schema.Int64Attribute{
											MarkdownDescription: `The maximum upload limit (integer, in Kbps). null indicates no limit`,
											Computed:            true,
										},
									},
								},
								"settings": schema.StringAttribute{
									MarkdownDescription: `How bandwidth limits are enforced. Can be 'network default', 'ignore' or 'custom'.`,
									Computed:            true,
								},
							},
						},
						"bonjour_forwarding": schema.SingleNestedAttribute{
							MarkdownDescription: `The Bonjour settings for your group policy. Only valid if your network has a wireless configuration.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"rules": schema.SetNestedAttribute{
									MarkdownDescription: `A list of the Bonjour forwarding rules for your group policy. If 'settings' is set to 'custom', at least one rule must be specified.`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"description": schema.StringAttribute{
												MarkdownDescription: `A description for your Bonjour forwarding rule. Optional.`,
												Computed:            true,
											},
											"services": schema.ListAttribute{
												MarkdownDescription: `A list of Bonjour services. At least one service must be specified. Available services are 'All Services', 'AFP', 'AirPlay', 'Apple screen share', 'BitTorrent', 'Chromecast', 'FTP', 'iChat', 'iTunes', 'Printers', 'Samba', 'Scanners', 'Spotify' and 'SSH'`,
												Computed:            true,
												ElementType:         types.StringType,
											},
											"vlan_id": schema.StringAttribute{
												MarkdownDescription: `The ID of the service VLAN. Required.`,
												Computed:            true,
											},
										},
									},
								},
								"settings": schema.StringAttribute{
									MarkdownDescription: `How Bonjour rules are applied. Can be 'network default', 'ignore' or 'custom'.`,
									Computed:            true,
								},
							},
						},
						"content_filtering": schema.SingleNestedAttribute{
							MarkdownDescription: `The content filtering settings for your group policy`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"allowed_url_patterns": schema.SingleNestedAttribute{
									MarkdownDescription: `Settings for allowed URL patterns`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"patterns": schema.ListAttribute{
											MarkdownDescription: `A list of URL patterns that are allowed`,
											Computed:            true,
											ElementType:         types.StringType,
										},
										"settings": schema.StringAttribute{
											MarkdownDescription: `How URL patterns are applied. Can be 'network default', 'append' or 'override'.`,
											Computed:            true,
										},
									},
								},
								"blocked_url_categories": schema.SingleNestedAttribute{
									MarkdownDescription: `Settings for blocked URL categories`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"categories": schema.ListAttribute{
											MarkdownDescription: `A list of URL categories to block`,
											Computed:            true,
											ElementType:         types.StringType,
										},
										"settings": schema.StringAttribute{
											MarkdownDescription: `How URL categories are applied. Can be 'network default', 'append' or 'override'.`,
											Computed:            true,
										},
									},
								},
								"blocked_url_patterns": schema.SingleNestedAttribute{
									MarkdownDescription: `Settings for blocked URL patterns`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"patterns": schema.ListAttribute{
											MarkdownDescription: `A list of URL patterns that are blocked`,
											Computed:            true,
											ElementType:         types.StringType,
										},
										"settings": schema.StringAttribute{
											MarkdownDescription: `How URL patterns are applied. Can be 'network default', 'append' or 'override'.`,
											Computed:            true,
										},
									},
								},
							},
						},
						"firewall_and_traffic_shaping": schema.SingleNestedAttribute{
							MarkdownDescription: `    The firewall and traffic shaping rules and settings for your policy.
`,
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"l3_firewall_rules": schema.SetNestedAttribute{
									MarkdownDescription: `An ordered array of the L3 firewall rules`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"comment": schema.StringAttribute{
												MarkdownDescription: `Description of the rule (optional)`,
												Computed:            true,
											},
											"dest_cidr": schema.StringAttribute{
												MarkdownDescription: `Destination IP address (in IP or CIDR notation), a fully-qualified domain name (FQDN, if your network supports it) or 'any'.`,
												Computed:            true,
											},
											"dest_port": schema.StringAttribute{
												MarkdownDescription: `Destination port (integer in the range 1-65535), a port range (e.g. 8080-9090), or 'any'`,
												Computed:            true,
											},
											"policy": schema.StringAttribute{
												MarkdownDescription: `'allow' or 'deny' traffic specified by this rule`,
												Computed:            true,
											},
											"protocol": schema.StringAttribute{
												MarkdownDescription: `The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')`,
												Computed:            true,
											},
										},
									},
								},
								"l7_firewall_rules": schema.SetNestedAttribute{
									MarkdownDescription: `An ordered array of L7 firewall rules`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"policy": schema.StringAttribute{
												MarkdownDescription: `The policy applied to matching traffic. Must be 'deny'.`,
												Computed:            true,
											},
											"type": schema.StringAttribute{
												MarkdownDescription: `Type of the L7 Rule. Must be 'application', 'applicationCategory', 'host', 'port' or 'ipRange'`,
												Computed:            true,
											},
											"value": schema.StringAttribute{
												MarkdownDescription: `The 'value' of what you want to block. If 'type' is 'host', 'port' or 'ipRange', 'value' must be a string matching either a hostname (e.g. somewhere.com), a port (e.g. 8080), or an IP range (e.g. 192.1.0.0/16). If 'type' is 'application' or 'applicationCategory', then 'value' must be an object with an ID for the application.`,
												Computed:            true,
											},
										},
									},
								},
								"settings": schema.StringAttribute{
									MarkdownDescription: `How firewall and traffic shaping rules are enforced. Can be 'network default', 'ignore' or 'custom'.`,
									Computed:            true,
								},
								"traffic_shaping_rules": schema.SetNestedAttribute{
									MarkdownDescription: `    An array of traffic shaping rules. Rules are applied in the order that
    they are specified in. An empty list (or null) means no rules. Note that
    you are allowed a maximum of 8 rules.
`,
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"definitions": schema.SetNestedAttribute{
												MarkdownDescription: `    A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.
`,
												Computed: true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{

														"type": schema.StringAttribute{
															MarkdownDescription: `The type of definition. Can be one of 'application', 'applicationCategory', 'host', 'port', 'ipRange' or 'localNet'.`,
															Computed:            true,
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
														},
													},
												},
											},
											"dscp_tag_value": schema.Int64Attribute{
												MarkdownDescription: `    The DSCP tag applied by your rule. null means 'Do not change DSCP tag'.
    For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.
`,
												Computed: true,
											},
											"pcp_tag_value": schema.Int64Attribute{
												MarkdownDescription: `    The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority).
    null means 'Do not set PCP tag'.
`,
												Computed: true,
											},
											"per_client_bandwidth_limits": schema.SingleNestedAttribute{
												MarkdownDescription: `    An object describing the bandwidth settings for your rule.
`,
												Computed: true,
												Attributes: map[string]schema.Attribute{

													"bandwidth_limits": schema.SingleNestedAttribute{
														MarkdownDescription: `The bandwidth limits object, specifying the upload ('limitUp') and download ('limitDown') speed in Kbps. These are only enforced if 'settings' is set to 'custom'.`,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"limit_down": schema.Int64Attribute{
																MarkdownDescription: `The maximum download limit (integer, in Kbps).`,
																Computed:            true,
															},
															"limit_up": schema.Int64Attribute{
																MarkdownDescription: `The maximum upload limit (integer, in Kbps).`,
																Computed:            true,
															},
														},
													},
													"settings": schema.StringAttribute{
														MarkdownDescription: `How bandwidth limits are applied by your rule. Can be one of 'network default', 'ignore' or 'custom'.`,
														Computed:            true,
													},
												},
											},
											"priority": schema.StringAttribute{
												MarkdownDescription: `    A string, indicating the priority level for packets bound to your rule.
    Can be 'low', 'normal' or 'high'.
`,
												Computed: true,
											},
										},
									},
								},
							},
						},
						"group_policy_id": schema.StringAttribute{
							MarkdownDescription: `The ID of the group policy`,
							Computed:            true,
						},
						"scheduling": schema.SingleNestedAttribute{
							MarkdownDescription: `    The schedule for the group policy. Schedules are applied to days of the week.
`,
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"enabled": schema.BoolAttribute{
									MarkdownDescription: `Whether scheduling is enabled (true) or disabled (false). Defaults to false. If true, the schedule objects for each day of the week (monday - sunday) are parsed.`,
									Computed:            true,
								},
								"friday": schema.SingleNestedAttribute{
									MarkdownDescription: `The schedule object for Friday.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"active": schema.BoolAttribute{
											MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
											Computed:            true,
										},
										"from": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
										"to": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
									},
								},
								"monday": schema.SingleNestedAttribute{
									MarkdownDescription: `The schedule object for Monday.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"active": schema.BoolAttribute{
											MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
											Computed:            true,
										},
										"from": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
										"to": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
									},
								},
								"saturday": schema.SingleNestedAttribute{
									MarkdownDescription: `The schedule object for Saturday.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"active": schema.BoolAttribute{
											MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
											Computed:            true,
										},
										"from": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
										"to": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
									},
								},
								"sunday": schema.SingleNestedAttribute{
									MarkdownDescription: `The schedule object for Sunday.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"active": schema.BoolAttribute{
											MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
											Computed:            true,
										},
										"from": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
										"to": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
									},
								},
								"thursday": schema.SingleNestedAttribute{
									MarkdownDescription: `The schedule object for Thursday.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"active": schema.BoolAttribute{
											MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
											Computed:            true,
										},
										"from": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
										"to": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
									},
								},
								"tuesday": schema.SingleNestedAttribute{
									MarkdownDescription: `The schedule object for Tuesday.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"active": schema.BoolAttribute{
											MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
											Computed:            true,
										},
										"from": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
										"to": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
									},
								},
								"wednesday": schema.SingleNestedAttribute{
									MarkdownDescription: `The schedule object for Wednesday.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"active": schema.BoolAttribute{
											MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
											Computed:            true,
										},
										"from": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
										"to": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
									},
								},
							},
						},
						"splash_auth_settings": schema.StringAttribute{
							MarkdownDescription: `Whether clients bound to your policy will bypass splash authorization or behave according to the network's rules. Can be one of 'network default' or 'bypass'. Only available if your network has a wireless configuration.`,
							Computed:            true,
						},
						"vlan_tagging": schema.SingleNestedAttribute{
							MarkdownDescription: `The VLAN tagging settings for your group policy. Only available if your network has a wireless configuration.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"settings": schema.StringAttribute{
									MarkdownDescription: `How VLAN tagging is applied. Can be 'network default', 'ignore' or 'custom'.`,
									Computed:            true,
								},
								"vlan_id": schema.StringAttribute{
									MarkdownDescription: `The ID of the vlan you want to tag. This only applies if 'settings' is set to 'custom'.`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksGroupPoliciesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksGroupPolicies NetworksGroupPolicies
	diags := req.Config.Get(ctx, &networksGroupPolicies)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksGroupPolicies.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksGroupPolicies.NetworkID.IsNull(), !networksGroupPolicies.GroupPolicyID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkGroupPolicies")
		vvNetworkID := networksGroupPolicies.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkGroupPolicies(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkGroupPolicies",
				err.Error(),
			)
			return
		}

		networksGroupPolicies = ResponseNetworksGetNetworkGroupPoliciesItemsToBody(networksGroupPolicies, response1)
		diags = resp.State.Set(ctx, &networksGroupPolicies)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkGroupPolicy")
		vvNetworkID := networksGroupPolicies.NetworkID.ValueString()
		vvGroupPolicyID := networksGroupPolicies.GroupPolicyID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Networks.GetNetworkGroupPolicy(vvNetworkID, vvGroupPolicyID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkGroupPolicy",
				err.Error(),
			)
			return
		}

		networksGroupPolicies = ResponseNetworksGetNetworkGroupPolicyItemToBody(networksGroupPolicies, response2)
		diags = resp.State.Set(ctx, &networksGroupPolicies)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksGroupPolicies struct {
	NetworkID     types.String                                   `tfsdk:"network_id"`
	GroupPolicyID types.String                                   `tfsdk:"group_policy_id"`
	Items         *[]ResponseItemNetworksGetNetworkGroupPolicies `tfsdk:"items"`
	Item          *ResponseNetworksGetNetworkGroupPolicy         `tfsdk:"item"`
}

type ResponseItemNetworksGetNetworkGroupPolicies struct {
	Bandwidth                 *ResponseItemNetworksGetNetworkGroupPoliciesBandwidth                 `tfsdk:"bandwidth"`
	BonjourForwarding         *ResponseItemNetworksGetNetworkGroupPoliciesBonjourForwarding         `tfsdk:"bonjour_forwarding"`
	ContentFiltering          *ResponseItemNetworksGetNetworkGroupPoliciesContentFiltering          `tfsdk:"content_filtering"`
	FirewallAndTrafficShaping *ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShaping `tfsdk:"firewall_and_traffic_shaping"`
	GroupPolicyID             types.String                                                          `tfsdk:"group_policy_id"`
	Scheduling                *ResponseItemNetworksGetNetworkGroupPoliciesScheduling                `tfsdk:"scheduling"`
	SplashAuthSettings        types.String                                                          `tfsdk:"splash_auth_settings"`
	VLANTagging               *ResponseItemNetworksGetNetworkGroupPoliciesVlanTagging               `tfsdk:"vlan_tagging"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesBandwidth struct {
	BandwidthLimits *ResponseItemNetworksGetNetworkGroupPoliciesBandwidthBandwidthLimits `tfsdk:"bandwidth_limits"`
	Settings        types.String                                                         `tfsdk:"settings"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesBandwidthBandwidthLimits struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesBonjourForwarding struct {
	Rules    *[]ResponseItemNetworksGetNetworkGroupPoliciesBonjourForwardingRules `tfsdk:"rules"`
	Settings types.String                                                         `tfsdk:"settings"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesBonjourForwardingRules struct {
	Description types.String `tfsdk:"description"`
	Services    types.List   `tfsdk:"services"`
	VLANID      types.String `tfsdk:"vlan_id"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesContentFiltering struct {
	AllowedURLPatterns   *ResponseItemNetworksGetNetworkGroupPoliciesContentFilteringAllowedUrlPatterns   `tfsdk:"allowed_url_patterns"`
	BlockedURLCategories *ResponseItemNetworksGetNetworkGroupPoliciesContentFilteringBlockedUrlCategories `tfsdk:"blocked_url_categories"`
	BlockedURLPatterns   *ResponseItemNetworksGetNetworkGroupPoliciesContentFilteringBlockedUrlPatterns   `tfsdk:"blocked_url_patterns"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesContentFilteringAllowedUrlPatterns struct {
	Patterns types.List   `tfsdk:"patterns"`
	Settings types.String `tfsdk:"settings"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesContentFilteringBlockedUrlCategories struct {
	Categories types.List   `tfsdk:"categories"`
	Settings   types.String `tfsdk:"settings"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesContentFilteringBlockedUrlPatterns struct {
	Patterns types.List   `tfsdk:"patterns"`
	Settings types.String `tfsdk:"settings"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShaping struct {
	L3FirewallRules     *[]ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingL3FirewallRules     `tfsdk:"l3_firewall_rules"`
	L7FirewallRules     *[]ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingL7FirewallRules     `tfsdk:"l7_firewall_rules"`
	Settings            types.String                                                                               `tfsdk:"settings"`
	TrafficShapingRules *[]ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules `tfsdk:"traffic_shaping_rules"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingL3FirewallRules struct {
	Comment  types.String `tfsdk:"comment"`
	DestCidr types.String `tfsdk:"dest_cidr"`
	DestPort types.String `tfsdk:"dest_port"`
	Policy   types.String `tfsdk:"policy"`
	Protocol types.String `tfsdk:"protocol"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingL7FirewallRules struct {
	Policy types.String `tfsdk:"policy"`
	Type   types.String `tfsdk:"type"`
	Value  types.String `tfsdk:"value"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules struct {
	Definitions              *[]ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesDefinitions            `tfsdk:"definitions"`
	DscpTagValue             types.Int64                                                                                                      `tfsdk:"dscp_tag_value"`
	PcpTagValue              types.Int64                                                                                                      `tfsdk:"pcp_tag_value"`
	PerClientBandwidthLimits *ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits `tfsdk:"per_client_bandwidth_limits"`
	Priority                 types.String                                                                                                     `tfsdk:"priority"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesDefinitions struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits struct {
	BandwidthLimits *ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits `tfsdk:"bandwidth_limits"`
	Settings        types.String                                                                                                                    `tfsdk:"settings"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesScheduling struct {
	Enabled   types.Bool                                                      `tfsdk:"enabled"`
	Friday    *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingFriday    `tfsdk:"friday"`
	Monday    *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingMonday    `tfsdk:"monday"`
	Saturday  *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingSaturday  `tfsdk:"saturday"`
	Sunday    *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingSunday    `tfsdk:"sunday"`
	Thursday  *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingThursday  `tfsdk:"thursday"`
	Tuesday   *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingTuesday   `tfsdk:"tuesday"`
	Wednesday *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingWednesday `tfsdk:"wednesday"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesSchedulingFriday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesSchedulingMonday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesSchedulingSaturday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesSchedulingSunday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesSchedulingThursday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesSchedulingTuesday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesSchedulingWednesday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseItemNetworksGetNetworkGroupPoliciesVlanTagging struct {
	Settings types.String `tfsdk:"settings"`
	VLANID   types.String `tfsdk:"vlan_id"`
}

type ResponseNetworksGetNetworkGroupPolicy struct {
	Bandwidth                 *ResponseNetworksGetNetworkGroupPolicyBandwidth                 `tfsdk:"bandwidth"`
	BonjourForwarding         *ResponseNetworksGetNetworkGroupPolicyBonjourForwarding         `tfsdk:"bonjour_forwarding"`
	ContentFiltering          *ResponseNetworksGetNetworkGroupPolicyContentFiltering          `tfsdk:"content_filtering"`
	FirewallAndTrafficShaping *ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShaping `tfsdk:"firewall_and_traffic_shaping"`
	GroupPolicyID             types.String                                                    `tfsdk:"group_policy_id"`
	Scheduling                *ResponseNetworksGetNetworkGroupPolicyScheduling                `tfsdk:"scheduling"`
	SplashAuthSettings        types.String                                                    `tfsdk:"splash_auth_settings"`
	VLANTagging               *ResponseNetworksGetNetworkGroupPolicyVlanTagging               `tfsdk:"vlan_tagging"`
}

type ResponseNetworksGetNetworkGroupPolicyBandwidth struct {
	BandwidthLimits *ResponseNetworksGetNetworkGroupPolicyBandwidthBandwidthLimits `tfsdk:"bandwidth_limits"`
	Settings        types.String                                                   `tfsdk:"settings"`
}

type ResponseNetworksGetNetworkGroupPolicyBandwidthBandwidthLimits struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

type ResponseNetworksGetNetworkGroupPolicyBonjourForwarding struct {
	Rules    *[]ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRules `tfsdk:"rules"`
	Settings types.String                                                   `tfsdk:"settings"`
}

type ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRules struct {
	Description types.String `tfsdk:"description"`
	Services    types.List   `tfsdk:"services"`
	VLANID      types.String `tfsdk:"vlan_id"`
}

type ResponseNetworksGetNetworkGroupPolicyContentFiltering struct {
	AllowedURLPatterns   *ResponseNetworksGetNetworkGroupPolicyContentFilteringAllowedUrlPatterns   `tfsdk:"allowed_url_patterns"`
	BlockedURLCategories *ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlCategories `tfsdk:"blocked_url_categories"`
	BlockedURLPatterns   *ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlPatterns   `tfsdk:"blocked_url_patterns"`
}

type ResponseNetworksGetNetworkGroupPolicyContentFilteringAllowedUrlPatterns struct {
	Patterns types.List   `tfsdk:"patterns"`
	Settings types.String `tfsdk:"settings"`
}

type ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlCategories struct {
	Categories types.List   `tfsdk:"categories"`
	Settings   types.String `tfsdk:"settings"`
}

type ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlPatterns struct {
	Patterns types.List   `tfsdk:"patterns"`
	Settings types.String `tfsdk:"settings"`
}

type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShaping struct {
	L3FirewallRules     *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules     `tfsdk:"l3_firewall_rules"`
	L7FirewallRules     *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules     `tfsdk:"l7_firewall_rules"`
	Settings            types.String                                                                         `tfsdk:"settings"`
	TrafficShapingRules *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules `tfsdk:"traffic_shaping_rules"`
}

type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules struct {
	Comment  types.String `tfsdk:"comment"`
	DestCidr types.String `tfsdk:"dest_cidr"`
	DestPort types.String `tfsdk:"dest_port"`
	Policy   types.String `tfsdk:"policy"`
	Protocol types.String `tfsdk:"protocol"`
}

type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules struct {
	Policy types.String `tfsdk:"policy"`
	Type   types.String `tfsdk:"type"`
	Value  types.String `tfsdk:"value"`
}

type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules struct {
	Definitions              *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions            `tfsdk:"definitions"`
	DscpTagValue             types.Int64                                                                                                `tfsdk:"dscp_tag_value"`
	PcpTagValue              types.Int64                                                                                                `tfsdk:"pcp_tag_value"`
	PerClientBandwidthLimits *ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits `tfsdk:"per_client_bandwidth_limits"`
	Priority                 types.String                                                                                               `tfsdk:"priority"`
}

type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits struct {
	BandwidthLimits *ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits `tfsdk:"bandwidth_limits"`
	Settings        types.String                                                                                                              `tfsdk:"settings"`
}

type ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

type ResponseNetworksGetNetworkGroupPolicyScheduling struct {
	Enabled   types.Bool                                                `tfsdk:"enabled"`
	Friday    *ResponseNetworksGetNetworkGroupPolicySchedulingFriday    `tfsdk:"friday"`
	Monday    *ResponseNetworksGetNetworkGroupPolicySchedulingMonday    `tfsdk:"monday"`
	Saturday  *ResponseNetworksGetNetworkGroupPolicySchedulingSaturday  `tfsdk:"saturday"`
	Sunday    *ResponseNetworksGetNetworkGroupPolicySchedulingSunday    `tfsdk:"sunday"`
	Thursday  *ResponseNetworksGetNetworkGroupPolicySchedulingThursday  `tfsdk:"thursday"`
	Tuesday   *ResponseNetworksGetNetworkGroupPolicySchedulingTuesday   `tfsdk:"tuesday"`
	Wednesday *ResponseNetworksGetNetworkGroupPolicySchedulingWednesday `tfsdk:"wednesday"`
}

type ResponseNetworksGetNetworkGroupPolicySchedulingFriday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseNetworksGetNetworkGroupPolicySchedulingMonday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseNetworksGetNetworkGroupPolicySchedulingSaturday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseNetworksGetNetworkGroupPolicySchedulingSunday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseNetworksGetNetworkGroupPolicySchedulingThursday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseNetworksGetNetworkGroupPolicySchedulingTuesday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseNetworksGetNetworkGroupPolicySchedulingWednesday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseNetworksGetNetworkGroupPolicyVlanTagging struct {
	Settings types.String `tfsdk:"settings"`
	VLANID   types.String `tfsdk:"vlan_id"`
}

// ToBody
func ResponseNetworksGetNetworkGroupPoliciesItemsToBody(state NetworksGroupPolicies, response *merakigosdk.ResponseNetworksGetNetworkGroupPolicies) NetworksGroupPolicies {
	var items []ResponseItemNetworksGetNetworkGroupPolicies
	for _, item := range *response {
		itemState := ResponseItemNetworksGetNetworkGroupPolicies{
			Bandwidth: func() *ResponseItemNetworksGetNetworkGroupPoliciesBandwidth {
				if item.Bandwidth != nil {
					return &ResponseItemNetworksGetNetworkGroupPoliciesBandwidth{
						BandwidthLimits: func() *ResponseItemNetworksGetNetworkGroupPoliciesBandwidthBandwidthLimits {
							if item.Bandwidth.BandwidthLimits != nil {
								return &ResponseItemNetworksGetNetworkGroupPoliciesBandwidthBandwidthLimits{
									LimitDown: func() types.Int64 {
										if item.Bandwidth.BandwidthLimits.LimitDown != nil {
											return types.Int64Value(int64(*item.Bandwidth.BandwidthLimits.LimitDown))
										}
										return types.Int64{}
									}(),
									LimitUp: func() types.Int64 {
										if item.Bandwidth.BandwidthLimits.LimitUp != nil {
											return types.Int64Value(int64(*item.Bandwidth.BandwidthLimits.LimitUp))
										}
										return types.Int64{}
									}(),
								}
							}
							return nil
						}(),
						Settings: types.StringValue(item.Bandwidth.Settings),
					}
				}
				return nil
			}(),
			BonjourForwarding: func() *ResponseItemNetworksGetNetworkGroupPoliciesBonjourForwarding {
				if item.BonjourForwarding != nil {
					return &ResponseItemNetworksGetNetworkGroupPoliciesBonjourForwarding{
						Rules: func() *[]ResponseItemNetworksGetNetworkGroupPoliciesBonjourForwardingRules {
							if item.BonjourForwarding.Rules != nil {
								result := make([]ResponseItemNetworksGetNetworkGroupPoliciesBonjourForwardingRules, len(*item.BonjourForwarding.Rules))
								for i, rules := range *item.BonjourForwarding.Rules {
									result[i] = ResponseItemNetworksGetNetworkGroupPoliciesBonjourForwardingRules{
										Description: types.StringValue(rules.Description),
										Services:    StringSliceToList(rules.Services),
										VLANID:      types.StringValue(rules.VLANID),
									}
								}
								return &result
							}
							return nil
						}(),
						Settings: types.StringValue(item.BonjourForwarding.Settings),
					}
				}
				return nil
			}(),
			ContentFiltering: func() *ResponseItemNetworksGetNetworkGroupPoliciesContentFiltering {
				if item.ContentFiltering != nil {
					return &ResponseItemNetworksGetNetworkGroupPoliciesContentFiltering{
						AllowedURLPatterns: func() *ResponseItemNetworksGetNetworkGroupPoliciesContentFilteringAllowedUrlPatterns {
							if item.ContentFiltering.AllowedURLPatterns != nil {
								return &ResponseItemNetworksGetNetworkGroupPoliciesContentFilteringAllowedUrlPatterns{
									Patterns: StringSliceToList(item.ContentFiltering.AllowedURLPatterns.Patterns),
									Settings: types.StringValue(item.ContentFiltering.AllowedURLPatterns.Settings),
								}
							}
							return nil
						}(),
						BlockedURLCategories: func() *ResponseItemNetworksGetNetworkGroupPoliciesContentFilteringBlockedUrlCategories {
							if item.ContentFiltering.BlockedURLCategories != nil {
								return &ResponseItemNetworksGetNetworkGroupPoliciesContentFilteringBlockedUrlCategories{
									Categories: StringSliceToList(item.ContentFiltering.BlockedURLCategories.Categories),
									Settings:   types.StringValue(item.ContentFiltering.BlockedURLCategories.Settings),
								}
							}
							return nil
						}(),
						BlockedURLPatterns: func() *ResponseItemNetworksGetNetworkGroupPoliciesContentFilteringBlockedUrlPatterns {
							if item.ContentFiltering.BlockedURLPatterns != nil {
								return &ResponseItemNetworksGetNetworkGroupPoliciesContentFilteringBlockedUrlPatterns{
									Patterns: StringSliceToList(item.ContentFiltering.BlockedURLPatterns.Patterns),
									Settings: types.StringValue(item.ContentFiltering.BlockedURLPatterns.Settings),
								}
							}
							return nil
						}(),
					}
				}
				return nil
			}(),
			FirewallAndTrafficShaping: func() *ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShaping {
				if item.FirewallAndTrafficShaping != nil {
					return &ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShaping{
						L3FirewallRules: func() *[]ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingL3FirewallRules {
							if item.FirewallAndTrafficShaping.L3FirewallRules != nil {
								result := make([]ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingL3FirewallRules, len(*item.FirewallAndTrafficShaping.L3FirewallRules))
								for i, l3FirewallRules := range *item.FirewallAndTrafficShaping.L3FirewallRules {
									result[i] = ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingL3FirewallRules{
										Comment:  types.StringValue(l3FirewallRules.Comment),
										DestCidr: types.StringValue(l3FirewallRules.DestCidr),
										DestPort: types.StringValue(l3FirewallRules.DestPort),
										Policy:   types.StringValue(l3FirewallRules.Policy),
										Protocol: types.StringValue(l3FirewallRules.Protocol),
									}
								}
								return &result
							}
							return nil
						}(),
						L7FirewallRules: func() *[]ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingL7FirewallRules {
							if item.FirewallAndTrafficShaping.L7FirewallRules != nil {
								result := make([]ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingL7FirewallRules, len(*item.FirewallAndTrafficShaping.L7FirewallRules))
								for i, l7FirewallRules := range *item.FirewallAndTrafficShaping.L7FirewallRules {
									result[i] = ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingL7FirewallRules{
										Policy: types.StringValue(l7FirewallRules.Policy),
										Type:   types.StringValue(l7FirewallRules.Type),
										Value:  types.StringValue(l7FirewallRules.Value),
									}
								}
								return &result
							}
							return nil
						}(),
						Settings: types.StringValue(item.FirewallAndTrafficShaping.Settings),
						TrafficShapingRules: func() *[]ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules {
							if item.FirewallAndTrafficShaping.TrafficShapingRules != nil {
								result := make([]ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules, len(*item.FirewallAndTrafficShaping.TrafficShapingRules))
								for i, trafficShapingRules := range *item.FirewallAndTrafficShaping.TrafficShapingRules {
									result[i] = ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules{
										Definitions: func() *[]ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesDefinitions {
											if trafficShapingRules.Definitions != nil {
												result := make([]ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesDefinitions, len(*trafficShapingRules.Definitions))
												for i, definitions := range *trafficShapingRules.Definitions {
													result[i] = ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesDefinitions{
														Type:  types.StringValue(definitions.Type),
														Value: types.StringValue(definitions.Value),
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
										PerClientBandwidthLimits: func() *ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits {
											if trafficShapingRules.PerClientBandwidthLimits != nil {
												return &ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits{
													BandwidthLimits: func() *ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits {
														if trafficShapingRules.PerClientBandwidthLimits.BandwidthLimits != nil {
															return &ResponseItemNetworksGetNetworkGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits{
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
													Settings: types.StringValue(trafficShapingRules.PerClientBandwidthLimits.Settings),
												}
											}
											return nil
										}(),
										Priority: types.StringValue(trafficShapingRules.Priority),
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
			GroupPolicyID: types.StringValue(item.GroupPolicyID),
			Scheduling: func() *ResponseItemNetworksGetNetworkGroupPoliciesScheduling {
				if item.Scheduling != nil {
					return &ResponseItemNetworksGetNetworkGroupPoliciesScheduling{
						Enabled: func() types.Bool {
							if item.Scheduling.Enabled != nil {
								return types.BoolValue(*item.Scheduling.Enabled)
							}
							return types.Bool{}
						}(),
						Friday: func() *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingFriday {
							if item.Scheduling.Friday != nil {
								return &ResponseItemNetworksGetNetworkGroupPoliciesSchedulingFriday{
									Active: func() types.Bool {
										if item.Scheduling.Friday.Active != nil {
											return types.BoolValue(*item.Scheduling.Friday.Active)
										}
										return types.Bool{}
									}(),
									From: types.StringValue(item.Scheduling.Friday.From),
									To:   types.StringValue(item.Scheduling.Friday.To),
								}
							}
							return nil
						}(),
						Monday: func() *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingMonday {
							if item.Scheduling.Monday != nil {
								return &ResponseItemNetworksGetNetworkGroupPoliciesSchedulingMonday{
									Active: func() types.Bool {
										if item.Scheduling.Monday.Active != nil {
											return types.BoolValue(*item.Scheduling.Monday.Active)
										}
										return types.Bool{}
									}(),
									From: types.StringValue(item.Scheduling.Monday.From),
									To:   types.StringValue(item.Scheduling.Monday.To),
								}
							}
							return nil
						}(),
						Saturday: func() *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingSaturday {
							if item.Scheduling.Saturday != nil {
								return &ResponseItemNetworksGetNetworkGroupPoliciesSchedulingSaturday{
									Active: func() types.Bool {
										if item.Scheduling.Saturday.Active != nil {
											return types.BoolValue(*item.Scheduling.Saturday.Active)
										}
										return types.Bool{}
									}(),
									From: types.StringValue(item.Scheduling.Saturday.From),
									To:   types.StringValue(item.Scheduling.Saturday.To),
								}
							}
							return nil
						}(),
						Sunday: func() *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingSunday {
							if item.Scheduling.Sunday != nil {
								return &ResponseItemNetworksGetNetworkGroupPoliciesSchedulingSunday{
									Active: func() types.Bool {
										if item.Scheduling.Sunday.Active != nil {
											return types.BoolValue(*item.Scheduling.Sunday.Active)
										}
										return types.Bool{}
									}(),
									From: types.StringValue(item.Scheduling.Sunday.From),
									To:   types.StringValue(item.Scheduling.Sunday.To),
								}
							}
							return nil
						}(),
						Thursday: func() *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingThursday {
							if item.Scheduling.Thursday != nil {
								return &ResponseItemNetworksGetNetworkGroupPoliciesSchedulingThursday{
									Active: func() types.Bool {
										if item.Scheduling.Thursday.Active != nil {
											return types.BoolValue(*item.Scheduling.Thursday.Active)
										}
										return types.Bool{}
									}(),
									From: types.StringValue(item.Scheduling.Thursday.From),
									To:   types.StringValue(item.Scheduling.Thursday.To),
								}
							}
							return nil
						}(),
						Tuesday: func() *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingTuesday {
							if item.Scheduling.Tuesday != nil {
								return &ResponseItemNetworksGetNetworkGroupPoliciesSchedulingTuesday{
									Active: func() types.Bool {
										if item.Scheduling.Tuesday.Active != nil {
											return types.BoolValue(*item.Scheduling.Tuesday.Active)
										}
										return types.Bool{}
									}(),
									From: types.StringValue(item.Scheduling.Tuesday.From),
									To:   types.StringValue(item.Scheduling.Tuesday.To),
								}
							}
							return nil
						}(),
						Wednesday: func() *ResponseItemNetworksGetNetworkGroupPoliciesSchedulingWednesday {
							if item.Scheduling.Wednesday != nil {
								return &ResponseItemNetworksGetNetworkGroupPoliciesSchedulingWednesday{
									Active: func() types.Bool {
										if item.Scheduling.Wednesday.Active != nil {
											return types.BoolValue(*item.Scheduling.Wednesday.Active)
										}
										return types.Bool{}
									}(),
									From: types.StringValue(item.Scheduling.Wednesday.From),
									To:   types.StringValue(item.Scheduling.Wednesday.To),
								}
							}
							return nil
						}(),
					}
				}
				return nil
			}(),
			SplashAuthSettings: types.StringValue(item.SplashAuthSettings),
			VLANTagging: func() *ResponseItemNetworksGetNetworkGroupPoliciesVlanTagging {
				if item.VLANTagging != nil {
					return &ResponseItemNetworksGetNetworkGroupPoliciesVlanTagging{
						Settings: types.StringValue(item.VLANTagging.Settings),
						VLANID:   types.StringValue(item.VLANTagging.VLANID),
					}
				}
				return nil
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseNetworksGetNetworkGroupPolicyItemToBody(state NetworksGroupPolicies, response *merakigosdk.ResponseNetworksGetNetworkGroupPolicy) NetworksGroupPolicies {
	itemState := ResponseNetworksGetNetworkGroupPolicy{
		Bandwidth: func() *ResponseNetworksGetNetworkGroupPolicyBandwidth {
			if response.Bandwidth != nil {
				return &ResponseNetworksGetNetworkGroupPolicyBandwidth{
					BandwidthLimits: func() *ResponseNetworksGetNetworkGroupPolicyBandwidthBandwidthLimits {
						if response.Bandwidth.BandwidthLimits != nil {
							return &ResponseNetworksGetNetworkGroupPolicyBandwidthBandwidthLimits{
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
					Settings: types.StringValue(response.Bandwidth.Settings),
				}
			}
			return nil
		}(),
		BonjourForwarding: func() *ResponseNetworksGetNetworkGroupPolicyBonjourForwarding {
			if response.BonjourForwarding != nil {
				return &ResponseNetworksGetNetworkGroupPolicyBonjourForwarding{
					Rules: func() *[]ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRules {
						if response.BonjourForwarding.Rules != nil {
							result := make([]ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRules, len(*response.BonjourForwarding.Rules))
							for i, rules := range *response.BonjourForwarding.Rules {
								result[i] = ResponseNetworksGetNetworkGroupPolicyBonjourForwardingRules{
									Description: types.StringValue(rules.Description),
									Services:    StringSliceToList(rules.Services),
									VLANID:      types.StringValue(rules.VLANID),
								}
							}
							return &result
						}
						return nil
					}(),
					Settings: types.StringValue(response.BonjourForwarding.Settings),
				}
			}
			return nil
		}(),
		ContentFiltering: func() *ResponseNetworksGetNetworkGroupPolicyContentFiltering {
			if response.ContentFiltering != nil {
				return &ResponseNetworksGetNetworkGroupPolicyContentFiltering{
					AllowedURLPatterns: func() *ResponseNetworksGetNetworkGroupPolicyContentFilteringAllowedUrlPatterns {
						if response.ContentFiltering.AllowedURLPatterns != nil {
							return &ResponseNetworksGetNetworkGroupPolicyContentFilteringAllowedUrlPatterns{
								Patterns: StringSliceToList(response.ContentFiltering.AllowedURLPatterns.Patterns),
								Settings: types.StringValue(response.ContentFiltering.AllowedURLPatterns.Settings),
							}
						}
						return nil
					}(),
					BlockedURLCategories: func() *ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlCategories {
						if response.ContentFiltering.BlockedURLCategories != nil {
							return &ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlCategories{
								Categories: StringSliceToList(response.ContentFiltering.BlockedURLCategories.Categories),
								Settings:   types.StringValue(response.ContentFiltering.BlockedURLCategories.Settings),
							}
						}
						return nil
					}(),
					BlockedURLPatterns: func() *ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlPatterns {
						if response.ContentFiltering.BlockedURLPatterns != nil {
							return &ResponseNetworksGetNetworkGroupPolicyContentFilteringBlockedUrlPatterns{
								Patterns: StringSliceToList(response.ContentFiltering.BlockedURLPatterns.Patterns),
								Settings: types.StringValue(response.ContentFiltering.BlockedURLPatterns.Settings),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		FirewallAndTrafficShaping: func() *ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShaping {
			if response.FirewallAndTrafficShaping != nil {
				return &ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShaping{
					L3FirewallRules: func() *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules {
						if response.FirewallAndTrafficShaping.L3FirewallRules != nil {
							result := make([]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules, len(*response.FirewallAndTrafficShaping.L3FirewallRules))
							for i, l3FirewallRules := range *response.FirewallAndTrafficShaping.L3FirewallRules {
								result[i] = ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules{
									Comment:  types.StringValue(l3FirewallRules.Comment),
									DestCidr: types.StringValue(l3FirewallRules.DestCidr),
									DestPort: types.StringValue(l3FirewallRules.DestPort),
									Policy:   types.StringValue(l3FirewallRules.Policy),
									Protocol: types.StringValue(l3FirewallRules.Protocol),
								}
							}
							return &result
						}
						return nil
					}(),
					L7FirewallRules: func() *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules {
						if response.FirewallAndTrafficShaping.L7FirewallRules != nil {
							result := make([]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules, len(*response.FirewallAndTrafficShaping.L7FirewallRules))
							for i, l7FirewallRules := range *response.FirewallAndTrafficShaping.L7FirewallRules {
								result[i] = ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules{
									Policy: types.StringValue(l7FirewallRules.Policy),
									Type:   types.StringValue(l7FirewallRules.Type),
									Value:  types.StringValue(l7FirewallRules.Value),
								}
							}
							return &result
						}
						return nil
					}(),
					Settings: types.StringValue(response.FirewallAndTrafficShaping.Settings),
					TrafficShapingRules: func() *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules {
						if response.FirewallAndTrafficShaping.TrafficShapingRules != nil {
							result := make([]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules, len(*response.FirewallAndTrafficShaping.TrafficShapingRules))
							for i, trafficShapingRules := range *response.FirewallAndTrafficShaping.TrafficShapingRules {
								result[i] = ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules{
									Definitions: func() *[]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions {
										if trafficShapingRules.Definitions != nil {
											result := make([]ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions, len(*trafficShapingRules.Definitions))
											for i, definitions := range *trafficShapingRules.Definitions {
												result[i] = ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions{
													Type:  types.StringValue(definitions.Type),
													Value: types.StringValue(definitions.Value),
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
									PerClientBandwidthLimits: func() *ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits {
										if trafficShapingRules.PerClientBandwidthLimits != nil {
											return &ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimits{
												BandwidthLimits: func() *ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits {
													if trafficShapingRules.PerClientBandwidthLimits.BandwidthLimits != nil {
														return &ResponseNetworksGetNetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits{
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
												Settings: types.StringValue(trafficShapingRules.PerClientBandwidthLimits.Settings),
											}
										}
										return nil
									}(),
									Priority: types.StringValue(trafficShapingRules.Priority),
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
		GroupPolicyID: types.StringValue(response.GroupPolicyID),
		Scheduling: func() *ResponseNetworksGetNetworkGroupPolicyScheduling {
			if response.Scheduling != nil {
				return &ResponseNetworksGetNetworkGroupPolicyScheduling{
					Enabled: func() types.Bool {
						if response.Scheduling.Enabled != nil {
							return types.BoolValue(*response.Scheduling.Enabled)
						}
						return types.Bool{}
					}(),
					Friday: func() *ResponseNetworksGetNetworkGroupPolicySchedulingFriday {
						if response.Scheduling.Friday != nil {
							return &ResponseNetworksGetNetworkGroupPolicySchedulingFriday{
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
						return nil
					}(),
					Monday: func() *ResponseNetworksGetNetworkGroupPolicySchedulingMonday {
						if response.Scheduling.Monday != nil {
							return &ResponseNetworksGetNetworkGroupPolicySchedulingMonday{
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
						return nil
					}(),
					Saturday: func() *ResponseNetworksGetNetworkGroupPolicySchedulingSaturday {
						if response.Scheduling.Saturday != nil {
							return &ResponseNetworksGetNetworkGroupPolicySchedulingSaturday{
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
						return nil
					}(),
					Sunday: func() *ResponseNetworksGetNetworkGroupPolicySchedulingSunday {
						if response.Scheduling.Sunday != nil {
							return &ResponseNetworksGetNetworkGroupPolicySchedulingSunday{
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
						return nil
					}(),
					Thursday: func() *ResponseNetworksGetNetworkGroupPolicySchedulingThursday {
						if response.Scheduling.Thursday != nil {
							return &ResponseNetworksGetNetworkGroupPolicySchedulingThursday{
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
						return nil
					}(),
					Tuesday: func() *ResponseNetworksGetNetworkGroupPolicySchedulingTuesday {
						if response.Scheduling.Tuesday != nil {
							return &ResponseNetworksGetNetworkGroupPolicySchedulingTuesday{
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
						return nil
					}(),
					Wednesday: func() *ResponseNetworksGetNetworkGroupPolicySchedulingWednesday {
						if response.Scheduling.Wednesday != nil {
							return &ResponseNetworksGetNetworkGroupPolicySchedulingWednesday{
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
						return nil
					}(),
				}
			}
			return nil
		}(),
		SplashAuthSettings: types.StringValue(response.SplashAuthSettings),
		VLANTagging: func() *ResponseNetworksGetNetworkGroupPolicyVlanTagging {
			if response.VLANTagging != nil {
				return &ResponseNetworksGetNetworkGroupPolicyVlanTagging{
					Settings: types.StringValue(response.VLANTagging.Settings),
					VLANID:   types.StringValue(response.VLANTagging.VLANID),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
