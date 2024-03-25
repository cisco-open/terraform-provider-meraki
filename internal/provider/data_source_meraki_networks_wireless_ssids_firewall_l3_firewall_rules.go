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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksWirelessSSIDsFirewallL3FirewallRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessSSIDsFirewallL3FirewallRulesDataSource{}
)

func NewNetworksWirelessSSIDsFirewallL3FirewallRulesDataSource() datasource.DataSource {
	return &NetworksWirelessSSIDsFirewallL3FirewallRulesDataSource{}
}

type NetworksWirelessSSIDsFirewallL3FirewallRulesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessSSIDsFirewallL3FirewallRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessSSIDsFirewallL3FirewallRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_firewall_l3_firewall_rules"
}

func (d *NetworksWirelessSSIDsFirewallL3FirewallRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: `number path parameter.`,
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
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessSSIDsFirewallL3FirewallRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessSSIDsFirewallL3FirewallRules NetworksWirelessSSIDsFirewallL3FirewallRules
	diags := req.Config.Get(ctx, &networksWirelessSSIDsFirewallL3FirewallRules)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessSSIDFirewallL3FirewallRules")
		vvNetworkID := networksWirelessSSIDsFirewallL3FirewallRules.NetworkID.ValueString()
		vvNumber := networksWirelessSSIDsFirewallL3FirewallRules.Number.ValueString()

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessSSIDFirewallL3FirewallRules(vvNetworkID, vvNumber)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDFirewallL3FirewallRules",
				err.Error(),
			)
			return
		}

		networksWirelessSSIDsFirewallL3FirewallRules = ResponseWirelessGetNetworkWirelessSSIDFirewallL3FirewallRulesItemToBody(networksWirelessSSIDsFirewallL3FirewallRules, response1)
		diags = resp.State.Set(ctx, &networksWirelessSSIDsFirewallL3FirewallRules)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessSSIDsFirewallL3FirewallRules struct {
	NetworkID types.String                                                   `tfsdk:"network_id"`
	Number    types.String                                                   `tfsdk:"number"`
	Item      *ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRules `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRules struct {
	Rules *[]ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRulesRules `tfsdk:"rules"`
}

type ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRulesRules struct {
	Comment  types.String `tfsdk:"comment"`
	DestCidr types.String `tfsdk:"dest_cidr"`
	DestPort types.String `tfsdk:"dest_port"`
	Policy   types.String `tfsdk:"policy"`
	Protocol types.String `tfsdk:"protocol"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessSSIDFirewallL3FirewallRulesItemToBody(state NetworksWirelessSSIDsFirewallL3FirewallRules, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDFirewallL3FirewallRules) NetworksWirelessSSIDsFirewallL3FirewallRules {
	itemState := ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRules{
		Rules: func() *[]ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRulesRules {
			if response.Rules != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRulesRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRulesRules{
						Comment:  types.StringValue(rules.Comment),
						DestCidr: types.StringValue(rules.DestCidr),
						DestPort: types.StringValue(rules.DestPort),
						Policy:   types.StringValue(rules.Policy),
						Protocol: types.StringValue(rules.Protocol),
					}
				}
				return &result
			}
			return &[]ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRulesRules{}
		}(),
	}
	state.Item = &itemState
	return state
}
