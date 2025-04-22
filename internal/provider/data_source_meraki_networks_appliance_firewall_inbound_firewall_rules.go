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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksApplianceFirewallInboundFirewallRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceFirewallInboundFirewallRulesDataSource{}
)

func NewNetworksApplianceFirewallInboundFirewallRulesDataSource() datasource.DataSource {
	return &NetworksApplianceFirewallInboundFirewallRulesDataSource{}
}

type NetworksApplianceFirewallInboundFirewallRulesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceFirewallInboundFirewallRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceFirewallInboundFirewallRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_firewall_inbound_firewall_rules"
}

func (d *NetworksApplianceFirewallInboundFirewallRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"rules": schema.SetNestedAttribute{
						MarkdownDescription: `An ordered array of the firewall rules (not including the default rule)`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"comment": schema.StringAttribute{
									MarkdownDescription: `Description of the rule (optional)`,
									Computed:            true,
								},
								"dest_cidr": schema.StringAttribute{
									MarkdownDescription: `Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'`,
									Computed:            true,
								},
								"dest_port": schema.StringAttribute{
									MarkdownDescription: `Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'`,
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
								"src_cidr": schema.StringAttribute{
									MarkdownDescription: `Comma-separated list of source IP address(es) (in IP or CIDR notation), or 'any' (note: FQDN not supported for source addresses)`,
									Computed:            true,
								},
								"src_port": schema.StringAttribute{
									MarkdownDescription: `Comma-separated list of source port(s) (integer in the range 1-65535), or 'any'`,
									Computed:            true,
								},
								"syslog_enabled": schema.BoolAttribute{
									MarkdownDescription: `Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)`,
									Computed:            true,
								},
							},
						},
					},
					"syslog_default_rule": schema.BoolAttribute{
						MarkdownDescription: `Log the special default rule (boolean value - enable only if you've configured a syslog server) (optional)`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *NetworksApplianceFirewallInboundFirewallRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceFirewallInboundFirewallRules NetworksApplianceFirewallInboundFirewallRules
	diags := req.Config.Get(ctx, &networksApplianceFirewallInboundFirewallRules)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceFirewallInboundFirewallRules")
		vvNetworkID := networksApplianceFirewallInboundFirewallRules.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceFirewallInboundFirewallRules(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceFirewallInboundFirewallRules",
				err.Error(),
			)
			return
		}

		networksApplianceFirewallInboundFirewallRules = ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRulesItemToBody(networksApplianceFirewallInboundFirewallRules, response1)
		diags = resp.State.Set(ctx, &networksApplianceFirewallInboundFirewallRules)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceFirewallInboundFirewallRules struct {
	NetworkID types.String                                                      `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRules `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRules struct {
	Rules             *[]ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRulesRules `tfsdk:"rules"`
	SyslogDefaultRule types.Bool                                                               `tfsdk:"syslog_default_rule"`
}

type ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRulesRules struct {
	Comment       types.String `tfsdk:"comment"`
	DestCidr      types.String `tfsdk:"dest_cidr"`
	DestPort      types.String `tfsdk:"dest_port"`
	Policy        types.String `tfsdk:"policy"`
	Protocol      types.String `tfsdk:"protocol"`
	SrcCidr       types.String `tfsdk:"src_cidr"`
	SrcPort       types.String `tfsdk:"src_port"`
	SyslogEnabled types.Bool   `tfsdk:"syslog_enabled"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRulesItemToBody(state NetworksApplianceFirewallInboundFirewallRules, response *merakigosdk.ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRules) NetworksApplianceFirewallInboundFirewallRules {
	itemState := ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRules{
		Rules: func() *[]ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRulesRules {
			if response.Rules != nil {
				result := make([]ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRulesRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseApplianceGetNetworkApplianceFirewallInboundFirewallRulesRules{
						Comment:  types.StringValue(rules.Comment),
						DestCidr: types.StringValue(rules.DestCidr),
						DestPort: types.StringValue(rules.DestPort),
						Policy:   types.StringValue(rules.Policy),
						Protocol: types.StringValue(rules.Protocol),
						SrcCidr:  types.StringValue(rules.SrcCidr),
						SrcPort:  types.StringValue(rules.SrcPort),
						SyslogEnabled: func() types.Bool {
							if rules.SyslogEnabled != nil {
								return types.BoolValue(*rules.SyslogEnabled)
							}
							return types.Bool{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		SyslogDefaultRule: func() types.Bool {
			if response.SyslogDefaultRule != nil {
				return types.BoolValue(*response.SyslogDefaultRule)
			}
			return types.Bool{}
		}(),
	}
	state.Item = &itemState
	return state
}
