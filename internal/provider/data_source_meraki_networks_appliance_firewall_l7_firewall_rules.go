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
	_ datasource.DataSource              = &NetworksApplianceFirewallL7FirewallRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceFirewallL7FirewallRulesDataSource{}
)

func NewNetworksApplianceFirewallL7FirewallRulesDataSource() datasource.DataSource {
	return &NetworksApplianceFirewallL7FirewallRulesDataSource{}
}

type NetworksApplianceFirewallL7FirewallRulesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceFirewallL7FirewallRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceFirewallL7FirewallRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_firewall_l7_firewall_rules"
}

func (d *NetworksApplianceFirewallL7FirewallRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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

								"policy": schema.StringAttribute{
									Computed: true,
								},
								"type": schema.StringAttribute{
									Computed: true,
								},
								"value": schema.StringAttribute{
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

func (d *NetworksApplianceFirewallL7FirewallRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceFirewallL7FirewallRules NetworksApplianceFirewallL7FirewallRules
	diags := req.Config.Get(ctx, &networksApplianceFirewallL7FirewallRules)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceFirewallL7FirewallRules")
		vvNetworkID := networksApplianceFirewallL7FirewallRules.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceFirewallL7FirewallRules(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceFirewallL7FirewallRules",
				err.Error(),
			)
			return
		}

		networksApplianceFirewallL7FirewallRules = ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesItemToBody(networksApplianceFirewallL7FirewallRules, response1)
		diags = resp.State.Set(ctx, &networksApplianceFirewallL7FirewallRules)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceFirewallL7FirewallRules struct {
	NetworkID types.String                                                 `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceFirewallL7FirewallRules `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceFirewallL7FirewallRules struct {
	Rules *[]ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRules `tfsdk:"rules"`
}

type ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRules struct {
	Policy types.String `tfsdk:"policy"`
	Type   types.String `tfsdk:"type"`
	Value  types.String `tfsdk:"value"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesItemToBody(state NetworksApplianceFirewallL7FirewallRules, response *merakigosdk.ResponseApplianceGetNetworkApplianceFirewallL7FirewallRules) NetworksApplianceFirewallL7FirewallRules {
	itemState := ResponseApplianceGetNetworkApplianceFirewallL7FirewallRules{
		Rules: func() *[]ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRules {
			if response.Rules != nil {
				result := make([]ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRules{
						Policy: func() types.String {
							if rules.Policy != "" {
								return types.StringValue(rules.Policy)
							}
							return types.String{}
						}(),
						Type: func() types.String {
							if rules.Type != "" {
								return types.StringValue(rules.Type)
							}
							return types.String{}
						}(),
						Value: func() types.String {
							if rules.Value == nil {
								return types.StringNull()
							}
							return types.StringValue(*rules.Value)
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
