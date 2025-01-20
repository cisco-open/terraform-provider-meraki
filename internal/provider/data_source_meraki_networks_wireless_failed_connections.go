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
	_ datasource.DataSource              = &NetworksWirelessFailedConnectionsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessFailedConnectionsDataSource{}
)

func NewNetworksWirelessFailedConnectionsDataSource() datasource.DataSource {
	return &NetworksWirelessFailedConnectionsDataSource{}
}

type NetworksWirelessFailedConnectionsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessFailedConnectionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessFailedConnectionsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_failed_connections"
}

func (d *NetworksWirelessFailedConnectionsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ap_tag": schema.StringAttribute{
				MarkdownDescription: `apTag query parameter. Filter results by AP Tag`,
				Optional:            true,
			},
			"band": schema.StringAttribute{
				MarkdownDescription: `band query parameter. Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information.`,
				Optional:            true,
			},
			"client_id": schema.StringAttribute{
				MarkdownDescription: `clientId query parameter. Filter by client MAC`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial query parameter. Filter by AP`,
				Optional:            true,
			},
			"ssid": schema.Int64Attribute{
				MarkdownDescription: `ssid query parameter. Filter results by SSID`,
				Optional:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 180 days from today.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 7 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.`,
				Optional:            true,
			},
			"vlan": schema.Int64Attribute{
				MarkdownDescription: `vlan query parameter. Filter results by VLAN`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseWirelessGetNetworkWirelessFailedConnections`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"client_mac": schema.StringAttribute{
							MarkdownDescription: `Client Mac`,
							Computed:            true,
						},
						"failure_step": schema.StringAttribute{
							MarkdownDescription: `The failed onboarding step. One of: assoc, auth, dhcp, dns.`,
							Computed:            true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `Serial Number`,
							Computed:            true,
						},
						"ssid_number": schema.Int64Attribute{
							MarkdownDescription: `SSID Number`,
							Computed:            true,
						},
						"ts": schema.StringAttribute{
							MarkdownDescription: `The timestamp when the client mac failed`,
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `The failure type in the onboarding step`,
							Computed:            true,
						},
						"vlan": schema.Int64Attribute{
							MarkdownDescription: `LAN`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessFailedConnectionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessFailedConnections NetworksWirelessFailedConnections
	diags := req.Config.Get(ctx, &networksWirelessFailedConnections)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessFailedConnections")
		vvNetworkID := networksWirelessFailedConnections.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkWirelessFailedConnectionsQueryParams{}

		queryParams1.T0 = networksWirelessFailedConnections.T0.ValueString()
		queryParams1.T1 = networksWirelessFailedConnections.T1.ValueString()
		queryParams1.Timespan = networksWirelessFailedConnections.Timespan.ValueFloat64()
		queryParams1.Band = networksWirelessFailedConnections.Band.ValueString()
		queryParams1.SSID = int(networksWirelessFailedConnections.SSID.ValueInt64())
		queryParams1.VLAN = int(networksWirelessFailedConnections.VLAN.ValueInt64())
		queryParams1.ApTag = networksWirelessFailedConnections.ApTag.ValueString()
		queryParams1.Serial = networksWirelessFailedConnections.Serial.ValueString()
		queryParams1.ClientID = networksWirelessFailedConnections.ClientID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessFailedConnections(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessFailedConnections",
				err.Error(),
			)
			return
		}

		networksWirelessFailedConnections = ResponseWirelessGetNetworkWirelessFailedConnectionsItemsToBody(networksWirelessFailedConnections, response1)
		diags = resp.State.Set(ctx, &networksWirelessFailedConnections)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessFailedConnections struct {
	NetworkID types.String                                               `tfsdk:"network_id"`
	T0        types.String                                               `tfsdk:"t0"`
	T1        types.String                                               `tfsdk:"t1"`
	Timespan  types.Float64                                              `tfsdk:"timespan"`
	Band      types.String                                               `tfsdk:"band"`
	SSID      types.Int64                                                `tfsdk:"ssid"`
	VLAN      types.Int64                                                `tfsdk:"vlan"`
	ApTag     types.String                                               `tfsdk:"ap_tag"`
	Serial    types.String                                               `tfsdk:"serial"`
	ClientID  types.String                                               `tfsdk:"client_id"`
	Items     *[]ResponseItemWirelessGetNetworkWirelessFailedConnections `tfsdk:"items"`
}

type ResponseItemWirelessGetNetworkWirelessFailedConnections struct {
	ClientMac   types.String `tfsdk:"client_mac"`
	FailureStep types.String `tfsdk:"failure_step"`
	Serial      types.String `tfsdk:"serial"`
	SSIDNumber  types.Int64  `tfsdk:"ssid_number"`
	Ts          types.String `tfsdk:"ts"`
	Type        types.String `tfsdk:"type"`
	VLAN        types.Int64  `tfsdk:"vlan"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessFailedConnectionsItemsToBody(state NetworksWirelessFailedConnections, response *merakigosdk.ResponseWirelessGetNetworkWirelessFailedConnections) NetworksWirelessFailedConnections {
	var items []ResponseItemWirelessGetNetworkWirelessFailedConnections
	for _, item := range *response {
		itemState := ResponseItemWirelessGetNetworkWirelessFailedConnections{
			ClientMac:   types.StringValue(item.ClientMac),
			FailureStep: types.StringValue(item.FailureStep),
			Serial:      types.StringValue(item.Serial),
			SSIDNumber: func() types.Int64 {
				if item.SSIDNumber != nil {
					return types.Int64Value(int64(*item.SSIDNumber))
				}
				return types.Int64{}
			}(),
			Ts:   types.StringValue(item.Ts),
			Type: types.StringValue(item.Type),
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
