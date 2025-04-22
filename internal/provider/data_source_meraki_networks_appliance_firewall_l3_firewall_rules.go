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
	_ datasource.DataSource              = &NetworksApplianceFirewallL3FirewallRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceFirewallL3FirewallRulesDataSource{}
)

func NewNetworksApplianceFirewallL3FirewallRulesDataSource() datasource.DataSource {
	return &NetworksApplianceFirewallL3FirewallRulesDataSource{}
}

type NetworksApplianceFirewallL3FirewallRulesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceFirewallL3FirewallRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceFirewallL3FirewallRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_firewall_l3_firewall_rules"
}

func (d *NetworksApplianceFirewallL3FirewallRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"comment": schema.StringAttribute{
									Computed: true,
								},
								"dest_cidr": schema.StringAttribute{
									Computed: true,
								},
								"dest_port": schema.StringAttribute{
									Computed: true,
								},
								"policy": schema.StringAttribute{
									Computed: true,
								},
								"protocol": schema.StringAttribute{
									Computed: true,
								},
								"src_cidr": schema.StringAttribute{
									Computed: true,
								},
								"src_port": schema.StringAttribute{
									Computed: true,
								},
								"syslog_enabled": schema.BoolAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksApplianceFirewallL3FirewallRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceFirewallL3FirewallRules NetworksApplianceFirewallL3FirewallRules
	diags := req.Config.Get(ctx, &networksApplianceFirewallL3FirewallRules)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceFirewallL3FirewallRules")
		vvNetworkID := networksApplianceFirewallL3FirewallRules.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceFirewallL3FirewallRules(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceFirewallL3FirewallRules",
				err.Error(),
			)
			return
		}

		networksApplianceFirewallL3FirewallRules = ResponseApplianceGetNetworkApplianceFirewallL3FirewallRulesItemToBody(networksApplianceFirewallL3FirewallRules, response1)
		diags = resp.State.Set(ctx, &networksApplianceFirewallL3FirewallRules)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceFirewallL3FirewallRules struct {
	NetworkID types.String                                                 `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceFirewallL3FirewallRules `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceFirewallL3FirewallRules struct {
	Rules *[]ResponseApplianceGetNetworkApplianceFirewallL3FirewallRulesRules `tfsdk:"rules"`
}

type ResponseApplianceGetNetworkApplianceFirewallL3FirewallRulesRules struct {
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
func ResponseApplianceGetNetworkApplianceFirewallL3FirewallRulesItemToBody(state NetworksApplianceFirewallL3FirewallRules, response *merakigosdk.ResponseApplianceGetNetworkApplianceFirewallL3FirewallRules) NetworksApplianceFirewallL3FirewallRules {
	itemState := ResponseApplianceGetNetworkApplianceFirewallL3FirewallRules{
		Rules: func() *[]ResponseApplianceGetNetworkApplianceFirewallL3FirewallRulesRules {
			if response.Rules != nil {
				result := make([]ResponseApplianceGetNetworkApplianceFirewallL3FirewallRulesRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseApplianceGetNetworkApplianceFirewallL3FirewallRulesRules{
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
	}
	state.Item = &itemState
	return state
}
