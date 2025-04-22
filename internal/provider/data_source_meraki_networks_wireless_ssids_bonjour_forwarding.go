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
	_ datasource.DataSource              = &NetworksWirelessSSIDsBonjourForwardingDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessSSIDsBonjourForwardingDataSource{}
)

func NewNetworksWirelessSSIDsBonjourForwardingDataSource() datasource.DataSource {
	return &NetworksWirelessSSIDsBonjourForwardingDataSource{}
}

type NetworksWirelessSSIDsBonjourForwardingDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessSSIDsBonjourForwardingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessSSIDsBonjourForwardingDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_bonjour_forwarding"
}

func (d *NetworksWirelessSSIDsBonjourForwardingDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `If true, Bonjour forwarding is enabled on the SSID.`,
						Computed:            true,
					},
					"exception": schema.SingleNestedAttribute{
						MarkdownDescription: `Bonjour forwarding exception`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `If true, Bonjour forwarding exception is enabled on this SSID. Exception is required to enable L2 isolation and Bonjour forwarding to work together.`,
								Computed:            true,
							},
						},
					},
					"rules": schema.SetNestedAttribute{
						MarkdownDescription: `Bonjour forwarding rules`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"description": schema.StringAttribute{
									MarkdownDescription: `Desctiption of the bonjour forwarding rule`,
									Computed:            true,
								},
								"services": schema.ListAttribute{
									MarkdownDescription: `A list of Bonjour services. At least one service must be specified. Available services are 'All Services', 'AFP', 'AirPlay', 'Apple screen share', 'BitTorrent', 'Chromecast', 'FTP', 'iChat', 'iTunes', 'Printers', 'Samba', 'Scanners', 'Spotify' and 'SSH'`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"vlan_id": schema.StringAttribute{
									MarkdownDescription: `The ID of the service VLAN. Required`,
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

func (d *NetworksWirelessSSIDsBonjourForwardingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessSSIDsBonjourForwarding NetworksWirelessSSIDsBonjourForwarding
	diags := req.Config.Get(ctx, &networksWirelessSSIDsBonjourForwarding)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessSSIDBonjourForwarding")
		vvNetworkID := networksWirelessSSIDsBonjourForwarding.NetworkID.ValueString()
		vvNumber := networksWirelessSSIDsBonjourForwarding.Number.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessSSIDBonjourForwarding(vvNetworkID, vvNumber)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDBonjourForwarding",
				err.Error(),
			)
			return
		}

		networksWirelessSSIDsBonjourForwarding = ResponseWirelessGetNetworkWirelessSSIDBonjourForwardingItemToBody(networksWirelessSSIDsBonjourForwarding, response1)
		diags = resp.State.Set(ctx, &networksWirelessSSIDsBonjourForwarding)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessSSIDsBonjourForwarding struct {
	NetworkID types.String                                             `tfsdk:"network_id"`
	Number    types.String                                             `tfsdk:"number"`
	Item      *ResponseWirelessGetNetworkWirelessSsidBonjourForwarding `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessSsidBonjourForwarding struct {
	Enabled   types.Bool                                                        `tfsdk:"enabled"`
	Exception *ResponseWirelessGetNetworkWirelessSsidBonjourForwardingException `tfsdk:"exception"`
	Rules     *[]ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRules   `tfsdk:"rules"`
}

type ResponseWirelessGetNetworkWirelessSsidBonjourForwardingException struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRules struct {
	Description types.String `tfsdk:"description"`
	Services    types.List   `tfsdk:"services"`
	VLANID      types.String `tfsdk:"vlan_id"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessSSIDBonjourForwardingItemToBody(state NetworksWirelessSSIDsBonjourForwarding, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDBonjourForwarding) NetworksWirelessSSIDsBonjourForwarding {
	itemState := ResponseWirelessGetNetworkWirelessSsidBonjourForwarding{
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		Exception: func() *ResponseWirelessGetNetworkWirelessSsidBonjourForwardingException {
			if response.Exception != nil {
				return &ResponseWirelessGetNetworkWirelessSsidBonjourForwardingException{
					Enabled: func() types.Bool {
						if response.Exception.Enabled != nil {
							return types.BoolValue(*response.Exception.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
		}(),
		Rules: func() *[]ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRules {
			if response.Rules != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRules{
						Description: types.StringValue(rules.Description),
						Services:    StringSliceToList(rules.Services),
						VLANID:      types.StringValue(rules.VLANID),
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
