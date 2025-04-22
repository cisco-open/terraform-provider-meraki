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
	_ datasource.DataSource              = &DevicesWirelessConnectionStatsDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesWirelessConnectionStatsDataSource{}
)

func NewDevicesWirelessConnectionStatsDataSource() datasource.DataSource {
	return &DevicesWirelessConnectionStatsDataSource{}
}

type DevicesWirelessConnectionStatsDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesWirelessConnectionStatsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesWirelessConnectionStatsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_wireless_connection_stats"
}

func (d *DevicesWirelessConnectionStatsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
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
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"connection_stats": schema.SingleNestedAttribute{
						MarkdownDescription: `The connection stats of the device`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"assoc": schema.Int64Attribute{
								MarkdownDescription: `The number of failed association attempts`,
								Computed:            true,
							},
							"auth": schema.Int64Attribute{
								MarkdownDescription: `The number of failed authentication attempts`,
								Computed:            true,
							},
							"dhcp": schema.Int64Attribute{
								MarkdownDescription: `The number of failed DHCP attempts`,
								Computed:            true,
							},
							"dns": schema.Int64Attribute{
								MarkdownDescription: `The number of failed DNS attempts`,
								Computed:            true,
							},
							"success": schema.Int64Attribute{
								MarkdownDescription: `The number of successful connection attempts`,
								Computed:            true,
							},
						},
					},
					"serial": schema.StringAttribute{
						MarkdownDescription: `The serial number for the device`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *DevicesWirelessConnectionStatsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesWirelessConnectionStats DevicesWirelessConnectionStats
	diags := req.Config.Get(ctx, &devicesWirelessConnectionStats)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceWirelessConnectionStats")
		vvSerial := devicesWirelessConnectionStats.Serial.ValueString()
		queryParams1 := merakigosdk.GetDeviceWirelessConnectionStatsQueryParams{}

		queryParams1.T0 = devicesWirelessConnectionStats.T0.ValueString()
		queryParams1.T1 = devicesWirelessConnectionStats.T1.ValueString()
		queryParams1.Timespan = devicesWirelessConnectionStats.Timespan.ValueFloat64()
		queryParams1.Band = devicesWirelessConnectionStats.Band.ValueString()
		queryParams1.SSID = int(devicesWirelessConnectionStats.SSID.ValueInt64())
		queryParams1.VLAN = int(devicesWirelessConnectionStats.VLAN.ValueInt64())
		queryParams1.ApTag = devicesWirelessConnectionStats.ApTag.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetDeviceWirelessConnectionStats(vvSerial, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceWirelessConnectionStats",
				err.Error(),
			)
			return
		}

		devicesWirelessConnectionStats = ResponseWirelessGetDeviceWirelessConnectionStatsItemToBody(devicesWirelessConnectionStats, response1)
		diags = resp.State.Set(ctx, &devicesWirelessConnectionStats)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesWirelessConnectionStats struct {
	Serial   types.String                                      `tfsdk:"serial"`
	T0       types.String                                      `tfsdk:"t0"`
	T1       types.String                                      `tfsdk:"t1"`
	Timespan types.Float64                                     `tfsdk:"timespan"`
	Band     types.String                                      `tfsdk:"band"`
	SSID     types.Int64                                       `tfsdk:"ssid"`
	VLAN     types.Int64                                       `tfsdk:"vlan"`
	ApTag    types.String                                      `tfsdk:"ap_tag"`
	Item     *ResponseWirelessGetDeviceWirelessConnectionStats `tfsdk:"item"`
}

type ResponseWirelessGetDeviceWirelessConnectionStats struct {
	ConnectionStats *ResponseWirelessGetDeviceWirelessConnectionStatsConnectionStats `tfsdk:"connection_stats"`
	Serial          types.String                                                     `tfsdk:"serial"`
}

type ResponseWirelessGetDeviceWirelessConnectionStatsConnectionStats struct {
	Assoc   types.Int64 `tfsdk:"assoc"`
	Auth    types.Int64 `tfsdk:"auth"`
	Dhcp    types.Int64 `tfsdk:"dhcp"`
	DNS     types.Int64 `tfsdk:"dns"`
	Success types.Int64 `tfsdk:"success"`
}

// ToBody
func ResponseWirelessGetDeviceWirelessConnectionStatsItemToBody(state DevicesWirelessConnectionStats, response *merakigosdk.ResponseWirelessGetDeviceWirelessConnectionStats) DevicesWirelessConnectionStats {
	itemState := ResponseWirelessGetDeviceWirelessConnectionStats{
		ConnectionStats: func() *ResponseWirelessGetDeviceWirelessConnectionStatsConnectionStats {
			if response.ConnectionStats != nil {
				return &ResponseWirelessGetDeviceWirelessConnectionStatsConnectionStats{
					Assoc: func() types.Int64 {
						if response.ConnectionStats.Assoc != nil {
							return types.Int64Value(int64(*response.ConnectionStats.Assoc))
						}
						return types.Int64{}
					}(),
					Auth: func() types.Int64 {
						if response.ConnectionStats.Auth != nil {
							return types.Int64Value(int64(*response.ConnectionStats.Auth))
						}
						return types.Int64{}
					}(),
					Dhcp: func() types.Int64 {
						if response.ConnectionStats.Dhcp != nil {
							return types.Int64Value(int64(*response.ConnectionStats.Dhcp))
						}
						return types.Int64{}
					}(),
					DNS: func() types.Int64 {
						if response.ConnectionStats.DNS != nil {
							return types.Int64Value(int64(*response.ConnectionStats.DNS))
						}
						return types.Int64{}
					}(),
					Success: func() types.Int64 {
						if response.ConnectionStats.Success != nil {
							return types.Int64Value(int64(*response.ConnectionStats.Success))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
		Serial: types.StringValue(response.Serial),
	}
	state.Item = &itemState
	return state
}
