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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersDataSource{}
)

func NewNetworksSwitchDhcpServerPolicyArpInspectionTrustedServersDataSource() datasource.DataSource {
	return &NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersDataSource{}
}

type NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_dhcp_server_policy_arp_inspection_trusted_servers"
}

func (d *NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"ipv4": schema.SingleNestedAttribute{
							MarkdownDescription: `IPv4 attributes of the trusted server.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"address": schema.StringAttribute{
									MarkdownDescription: `IPv4 address of the trusted server.`,
									Computed:            true,
								},
							},
						},
						"mac": schema.StringAttribute{
							MarkdownDescription: `Mac address of the trusted server.`,
							Computed:            true,
						},
						"trusted_server_id": schema.StringAttribute{
							MarkdownDescription: `ID of the trusted server.`,
							Computed:            true,
						},
						"vlan": schema.Int64Attribute{
							MarkdownDescription: `Vlan ID of the trusted server.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchDhcpServerPolicyArpInspectionTrustedServers NetworksSwitchDhcpServerPolicyArpInspectionTrustedServers
	diags := req.Config.Get(ctx, &networksSwitchDhcpServerPolicyArpInspectionTrustedServers)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers")
		vvNetworkID := networksSwitchDhcpServerPolicyArpInspectionTrustedServers.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersQueryParams{}

		queryParams1.PerPage = int(networksSwitchDhcpServerPolicyArpInspectionTrustedServers.PerPage.ValueInt64())
		queryParams1.StartingAfter = networksSwitchDhcpServerPolicyArpInspectionTrustedServers.StartingAfter.ValueString()
		queryParams1.EndingBefore = networksSwitchDhcpServerPolicyArpInspectionTrustedServers.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers",
				err.Error(),
			)
			return
		}

		networksSwitchDhcpServerPolicyArpInspectionTrustedServers = ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersItemsToBody(networksSwitchDhcpServerPolicyArpInspectionTrustedServers, response1)
		diags = resp.State.Set(ctx, &networksSwitchDhcpServerPolicyArpInspectionTrustedServers)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchDhcpServerPolicyArpInspectionTrustedServers struct {
	NetworkID     types.String                                                                     `tfsdk:"network_id"`
	PerPage       types.Int64                                                                      `tfsdk:"per_page"`
	StartingAfter types.String                                                                     `tfsdk:"starting_after"`
	EndingBefore  types.String                                                                     `tfsdk:"ending_before"`
	Items         *[]ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers `tfsdk:"items"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers struct {
	IPv4            *ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4 `tfsdk:"ipv4"`
	Mac             types.String                                                                       `tfsdk:"mac"`
	TrustedServerID types.String                                                                       `tfsdk:"trusted_server_id"`
	VLAN            types.Int64                                                                        `tfsdk:"vlan"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4 struct {
	Address types.String `tfsdk:"address"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersItemsToBody(state NetworksSwitchDhcpServerPolicyArpInspectionTrustedServers, response *merakigosdk.ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers) NetworksSwitchDhcpServerPolicyArpInspectionTrustedServers {
	var items []ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers
	for _, item := range *response {
		itemState := ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers{
			IPv4: func() *ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4 {
				if item.IPv4 != nil {
					return &ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4{
						Address: types.StringValue(item.IPv4.Address),
					}
				}
				return nil
			}(),
			Mac:             types.StringValue(item.Mac),
			TrustedServerID: types.StringValue(item.TrustedServerID),
			VLAN: func() types.Int64 {
				if item.VLAN != nil {
					return types.Int64Value(int64(*item.VLAN))
				}
				return types.Int64{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
