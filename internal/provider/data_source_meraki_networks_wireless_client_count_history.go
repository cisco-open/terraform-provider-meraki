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
	_ datasource.DataSource              = &NetworksWirelessClientCountHistoryDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessClientCountHistoryDataSource{}
)

func NewNetworksWirelessClientCountHistoryDataSource() datasource.DataSource {
	return &NetworksWirelessClientCountHistoryDataSource{}
}

type NetworksWirelessClientCountHistoryDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessClientCountHistoryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessClientCountHistoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_client_count_history"
}

func (d *NetworksWirelessClientCountHistoryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ap_tag": schema.StringAttribute{
				MarkdownDescription: `apTag query parameter. Filter results by AP tag.`,
				Optional:            true,
			},
			"auto_resolution": schema.BoolAttribute{
				MarkdownDescription: `autoResolution query parameter. Automatically select a data resolution based on the given timespan; this overrides the value specified by the 'resolution' parameter. The default setting is false.`,
				Optional:            true,
			},
			"band": schema.StringAttribute{
				MarkdownDescription: `band query parameter. Filter results by band (either '2.4', '5' or '6').`,
				Optional:            true,
			},
			"client_id": schema.StringAttribute{
				MarkdownDescription: `clientId query parameter. Filter results by network client to return per-device client counts over time inner joined by the queried client's connection history.`,
				Optional:            true,
			},
			"device_serial": schema.StringAttribute{
				MarkdownDescription: `deviceSerial query parameter. Filter results by device.`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"resolution": schema.Int64Attribute{
				MarkdownDescription: `resolution query parameter. The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400.`,
				Optional:            true,
			},
			"ssid": schema.Int64Attribute{
				MarkdownDescription: `ssid query parameter. Filter results by SSID number.`,
				Optional:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 31 days from today.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 31 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseWirelessGetNetworkWirelessClientCountHistory`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"client_count": schema.Int64Attribute{
							MarkdownDescription: `Number of connected clients`,
							Computed:            true,
						},
						"end_ts": schema.StringAttribute{
							MarkdownDescription: `The end time of the query range`,
							Computed:            true,
						},
						"start_ts": schema.StringAttribute{
							MarkdownDescription: `The start time of the query range`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessClientCountHistoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessClientCountHistory NetworksWirelessClientCountHistory
	diags := req.Config.Get(ctx, &networksWirelessClientCountHistory)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessClientCountHistory")
		vvNetworkID := networksWirelessClientCountHistory.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkWirelessClientCountHistoryQueryParams{}

		queryParams1.T0 = networksWirelessClientCountHistory.T0.ValueString()
		queryParams1.T1 = networksWirelessClientCountHistory.T1.ValueString()
		queryParams1.Timespan = networksWirelessClientCountHistory.Timespan.ValueFloat64()
		queryParams1.Resolution = int(networksWirelessClientCountHistory.Resolution.ValueInt64())
		queryParams1.AutoResolution = networksWirelessClientCountHistory.AutoResolution.ValueBool()
		queryParams1.ClientID = networksWirelessClientCountHistory.ClientID.ValueString()
		queryParams1.DeviceSerial = networksWirelessClientCountHistory.DeviceSerial.ValueString()
		queryParams1.ApTag = networksWirelessClientCountHistory.ApTag.ValueString()
		queryParams1.Band = networksWirelessClientCountHistory.Band.ValueString()
		queryParams1.SSID = int(networksWirelessClientCountHistory.SSID.ValueInt64())

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessClientCountHistory(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessClientCountHistory",
				err.Error(),
			)
			return
		}

		networksWirelessClientCountHistory = ResponseWirelessGetNetworkWirelessClientCountHistoryItemsToBody(networksWirelessClientCountHistory, response1)
		diags = resp.State.Set(ctx, &networksWirelessClientCountHistory)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessClientCountHistory struct {
	NetworkID      types.String                                                `tfsdk:"network_id"`
	T0             types.String                                                `tfsdk:"t0"`
	T1             types.String                                                `tfsdk:"t1"`
	Timespan       types.Float64                                               `tfsdk:"timespan"`
	Resolution     types.Int64                                                 `tfsdk:"resolution"`
	AutoResolution types.Bool                                                  `tfsdk:"auto_resolution"`
	ClientID       types.String                                                `tfsdk:"client_id"`
	DeviceSerial   types.String                                                `tfsdk:"device_serial"`
	ApTag          types.String                                                `tfsdk:"ap_tag"`
	Band           types.String                                                `tfsdk:"band"`
	SSID           types.Int64                                                 `tfsdk:"ssid"`
	Items          *[]ResponseItemWirelessGetNetworkWirelessClientCountHistory `tfsdk:"items"`
}

type ResponseItemWirelessGetNetworkWirelessClientCountHistory struct {
	ClientCount types.Int64  `tfsdk:"client_count"`
	EndTs       types.String `tfsdk:"end_ts"`
	StartTs     types.String `tfsdk:"start_ts"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessClientCountHistoryItemsToBody(state NetworksWirelessClientCountHistory, response *merakigosdk.ResponseWirelessGetNetworkWirelessClientCountHistory) NetworksWirelessClientCountHistory {
	var items []ResponseItemWirelessGetNetworkWirelessClientCountHistory
	for _, item := range *response {
		itemState := ResponseItemWirelessGetNetworkWirelessClientCountHistory{
			ClientCount: func() types.Int64 {
				if item.ClientCount != nil {
					return types.Int64Value(int64(*item.ClientCount))
				}
				return types.Int64{}
			}(),
			EndTs: func() types.String {
				if item.EndTs != "" {
					return types.StringValue(item.EndTs)
				}
				return types.String{}
			}(),
			StartTs: func() types.String {
				if item.StartTs != "" {
					return types.StringValue(item.StartTs)
				}
				return types.String{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
