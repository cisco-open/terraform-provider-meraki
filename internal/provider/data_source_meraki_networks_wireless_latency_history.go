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
	_ datasource.DataSource              = &NetworksWirelessLatencyHistoryDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessLatencyHistoryDataSource{}
)

func NewNetworksWirelessLatencyHistoryDataSource() datasource.DataSource {
	return &NetworksWirelessLatencyHistoryDataSource{}
}

type NetworksWirelessLatencyHistoryDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessLatencyHistoryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessLatencyHistoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_latency_history"
}

func (d *NetworksWirelessLatencyHistoryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"access_category": schema.StringAttribute{
				MarkdownDescription: `accessCategory query parameter. Filter by access category.`,
				Optional:            true,
			},
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
				MarkdownDescription: `clientId query parameter. Filter results by network client.`,
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
				MarkdownDescription: `Array of ResponseWirelessGetNetworkWirelessLatencyHistory`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"avg_latency_ms": schema.Int64Attribute{
							MarkdownDescription: `Average latency in milliseconds`,
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

func (d *NetworksWirelessLatencyHistoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessLatencyHistory NetworksWirelessLatencyHistory
	diags := req.Config.Get(ctx, &networksWirelessLatencyHistory)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessLatencyHistory")
		vvNetworkID := networksWirelessLatencyHistory.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkWirelessLatencyHistoryQueryParams{}

		queryParams1.T0 = networksWirelessLatencyHistory.T0.ValueString()
		queryParams1.T1 = networksWirelessLatencyHistory.T1.ValueString()
		queryParams1.Timespan = networksWirelessLatencyHistory.Timespan.ValueFloat64()
		queryParams1.Resolution = int(networksWirelessLatencyHistory.Resolution.ValueInt64())
		queryParams1.AutoResolution = networksWirelessLatencyHistory.AutoResolution.ValueBool()
		queryParams1.ClientID = networksWirelessLatencyHistory.ClientID.ValueString()
		queryParams1.DeviceSerial = networksWirelessLatencyHistory.DeviceSerial.ValueString()
		queryParams1.ApTag = networksWirelessLatencyHistory.ApTag.ValueString()
		queryParams1.Band = networksWirelessLatencyHistory.Band.ValueString()
		queryParams1.SSID = int(networksWirelessLatencyHistory.SSID.ValueInt64())
		queryParams1.AccessCategory = networksWirelessLatencyHistory.AccessCategory.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessLatencyHistory(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessLatencyHistory",
				err.Error(),
			)
			return
		}

		networksWirelessLatencyHistory = ResponseWirelessGetNetworkWirelessLatencyHistoryItemsToBody(networksWirelessLatencyHistory, response1)
		diags = resp.State.Set(ctx, &networksWirelessLatencyHistory)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessLatencyHistory struct {
	NetworkID      types.String                                            `tfsdk:"network_id"`
	T0             types.String                                            `tfsdk:"t0"`
	T1             types.String                                            `tfsdk:"t1"`
	Timespan       types.Float64                                           `tfsdk:"timespan"`
	Resolution     types.Int64                                             `tfsdk:"resolution"`
	AutoResolution types.Bool                                              `tfsdk:"auto_resolution"`
	ClientID       types.String                                            `tfsdk:"client_id"`
	DeviceSerial   types.String                                            `tfsdk:"device_serial"`
	ApTag          types.String                                            `tfsdk:"ap_tag"`
	Band           types.String                                            `tfsdk:"band"`
	SSID           types.Int64                                             `tfsdk:"ssid"`
	AccessCategory types.String                                            `tfsdk:"access_category"`
	Items          *[]ResponseItemWirelessGetNetworkWirelessLatencyHistory `tfsdk:"items"`
}

type ResponseItemWirelessGetNetworkWirelessLatencyHistory struct {
	AvgLatencyMs types.Int64  `tfsdk:"avg_latency_ms"`
	EndTs        types.String `tfsdk:"end_ts"`
	StartTs      types.String `tfsdk:"start_ts"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessLatencyHistoryItemsToBody(state NetworksWirelessLatencyHistory, response *merakigosdk.ResponseWirelessGetNetworkWirelessLatencyHistory) NetworksWirelessLatencyHistory {
	var items []ResponseItemWirelessGetNetworkWirelessLatencyHistory
	for _, item := range *response {
		itemState := ResponseItemWirelessGetNetworkWirelessLatencyHistory{
			AvgLatencyMs: func() types.Int64 {
				if item.AvgLatencyMs != nil {
					return types.Int64Value(int64(*item.AvgLatencyMs))
				}
				return types.Int64{}
			}(),
			EndTs:   types.StringValue(item.EndTs),
			StartTs: types.StringValue(item.StartTs),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
