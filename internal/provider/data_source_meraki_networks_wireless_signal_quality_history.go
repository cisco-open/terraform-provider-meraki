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
	_ datasource.DataSource              = &NetworksWirelessSignalQualityHistoryDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessSignalQualityHistoryDataSource{}
)

func NewNetworksWirelessSignalQualityHistoryDataSource() datasource.DataSource {
	return &NetworksWirelessSignalQualityHistoryDataSource{}
}

type NetworksWirelessSignalQualityHistoryDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessSignalQualityHistoryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessSignalQualityHistoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_signal_quality_history"
}

func (d *NetworksWirelessSignalQualityHistoryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ap_tag": schema.StringAttribute{
				MarkdownDescription: `apTag query parameter. Filter results by AP tag; either :clientId or :deviceSerial must be jointly specified.`,
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
				MarkdownDescription: `Array of ResponseWirelessGetNetworkWirelessSignalQualityHistory`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"end_ts": schema.StringAttribute{
							MarkdownDescription: `The end time of the query range`,
							Computed:            true,
						},
						"rssi": schema.Int64Attribute{
							MarkdownDescription: `Received signal strength indicator`,
							Computed:            true,
						},
						"snr": schema.Int64Attribute{
							MarkdownDescription: `Signal to noise ratio`,
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

func (d *NetworksWirelessSignalQualityHistoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessSignalQualityHistory NetworksWirelessSignalQualityHistory
	diags := req.Config.Get(ctx, &networksWirelessSignalQualityHistory)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessSignalQualityHistory")
		vvNetworkID := networksWirelessSignalQualityHistory.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkWirelessSignalQualityHistoryQueryParams{}

		queryParams1.T0 = networksWirelessSignalQualityHistory.T0.ValueString()
		queryParams1.T1 = networksWirelessSignalQualityHistory.T1.ValueString()
		queryParams1.Timespan = networksWirelessSignalQualityHistory.Timespan.ValueFloat64()
		queryParams1.Resolution = int(networksWirelessSignalQualityHistory.Resolution.ValueInt64())
		queryParams1.AutoResolution = networksWirelessSignalQualityHistory.AutoResolution.ValueBool()
		queryParams1.ClientID = networksWirelessSignalQualityHistory.ClientID.ValueString()
		queryParams1.DeviceSerial = networksWirelessSignalQualityHistory.DeviceSerial.ValueString()
		queryParams1.ApTag = networksWirelessSignalQualityHistory.ApTag.ValueString()
		queryParams1.Band = networksWirelessSignalQualityHistory.Band.ValueString()
		queryParams1.SSID = int(networksWirelessSignalQualityHistory.SSID.ValueInt64())

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessSignalQualityHistory(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSignalQualityHistory",
				err.Error(),
			)
			return
		}

		networksWirelessSignalQualityHistory = ResponseWirelessGetNetworkWirelessSignalQualityHistoryItemsToBody(networksWirelessSignalQualityHistory, response1)
		diags = resp.State.Set(ctx, &networksWirelessSignalQualityHistory)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessSignalQualityHistory struct {
	NetworkID      types.String                                                  `tfsdk:"network_id"`
	T0             types.String                                                  `tfsdk:"t0"`
	T1             types.String                                                  `tfsdk:"t1"`
	Timespan       types.Float64                                                 `tfsdk:"timespan"`
	Resolution     types.Int64                                                   `tfsdk:"resolution"`
	AutoResolution types.Bool                                                    `tfsdk:"auto_resolution"`
	ClientID       types.String                                                  `tfsdk:"client_id"`
	DeviceSerial   types.String                                                  `tfsdk:"device_serial"`
	ApTag          types.String                                                  `tfsdk:"ap_tag"`
	Band           types.String                                                  `tfsdk:"band"`
	SSID           types.Int64                                                   `tfsdk:"ssid"`
	Items          *[]ResponseItemWirelessGetNetworkWirelessSignalQualityHistory `tfsdk:"items"`
}

type ResponseItemWirelessGetNetworkWirelessSignalQualityHistory struct {
	EndTs   types.String `tfsdk:"end_ts"`
	Rssi    types.Int64  `tfsdk:"rssi"`
	Snr     types.Int64  `tfsdk:"snr"`
	StartTs types.String `tfsdk:"start_ts"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessSignalQualityHistoryItemsToBody(state NetworksWirelessSignalQualityHistory, response *merakigosdk.ResponseWirelessGetNetworkWirelessSignalQualityHistory) NetworksWirelessSignalQualityHistory {
	var items []ResponseItemWirelessGetNetworkWirelessSignalQualityHistory
	for _, item := range *response {
		itemState := ResponseItemWirelessGetNetworkWirelessSignalQualityHistory{
			EndTs: types.StringValue(item.EndTs),
			Rssi: func() types.Int64 {
				if item.Rssi != nil {
					return types.Int64Value(int64(*item.Rssi))
				}
				return types.Int64{}
			}(),
			Snr: func() types.Int64 {
				if item.Snr != nil {
					return types.Int64Value(int64(*item.Snr))
				}
				return types.Int64{}
			}(),
			StartTs: types.StringValue(item.StartTs),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
