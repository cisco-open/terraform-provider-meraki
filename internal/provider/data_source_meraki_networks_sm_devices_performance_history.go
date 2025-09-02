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
	_ datasource.DataSource              = &NetworksSmDevicesPerformanceHistoryDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSmDevicesPerformanceHistoryDataSource{}
)

func NewNetworksSmDevicesPerformanceHistoryDataSource() datasource.DataSource {
	return &NetworksSmDevicesPerformanceHistoryDataSource{}
}

type NetworksSmDevicesPerformanceHistoryDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSmDevicesPerformanceHistoryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSmDevicesPerformanceHistoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_devices_performance_history"
}

func (d *NetworksSmDevicesPerformanceHistoryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"device_id": schema.StringAttribute{
				MarkdownDescription: `deviceId path parameter. Device ID`,
				Required:            true,
			},
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
				MarkdownDescription: `Array of ResponseSmGetNetworkSmDevicePerformanceHistory`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"cpu_percent_used": schema.Float64Attribute{
							MarkdownDescription: `The percentage of CPU used as a decimal format.`,
							Computed:            true,
						},
						"disk_usage": schema.SingleNestedAttribute{
							MarkdownDescription: `An object containing disk usage details.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"c": schema.SingleNestedAttribute{
									MarkdownDescription: `An object containing current disk usage details.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"space": schema.Int64Attribute{
											MarkdownDescription: `The available disk space.`,
											Computed:            true,
										},
										"used": schema.Int64Attribute{
											MarkdownDescription: `The used disk space.`,
											Computed:            true,
										},
									},
								},
							},
						},
						"mem_active": schema.Int64Attribute{
							MarkdownDescription: `The active RAM on the device.`,
							Computed:            true,
						},
						"mem_free": schema.Int64Attribute{
							MarkdownDescription: `Memory that is not yet in use by the system.`,
							Computed:            true,
						},
						"mem_inactive": schema.Int64Attribute{
							MarkdownDescription: `The inactive RAM on the device.`,
							Computed:            true,
						},
						"mem_wired": schema.Int64Attribute{
							MarkdownDescription: `Memory used for core OS functions on the device.`,
							Computed:            true,
						},
						"network_received": schema.Int64Attribute{
							MarkdownDescription: `Network bandwith received.`,
							Computed:            true,
						},
						"network_sent": schema.Int64Attribute{
							MarkdownDescription: `Network bandwith transmitted.`,
							Computed:            true,
						},
						"swap_used": schema.Int64Attribute{
							MarkdownDescription: `The amount of space being used on the startup disk to swap unused files to and from RAM.`,
							Computed:            true,
						},
						"ts": schema.StringAttribute{
							MarkdownDescription: `The time at which the performance was measured.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSmDevicesPerformanceHistoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSmDevicesPerformanceHistory NetworksSmDevicesPerformanceHistory
	diags := req.Config.Get(ctx, &networksSmDevicesPerformanceHistory)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSmDevicePerformanceHistory")
		vvNetworkID := networksSmDevicesPerformanceHistory.NetworkID.ValueString()
		vvDeviceID := networksSmDevicesPerformanceHistory.DeviceID.ValueString()
		queryParams1 := merakigosdk.GetNetworkSmDevicePerformanceHistoryQueryParams{}

		queryParams1.PerPage = int(networksSmDevicesPerformanceHistory.PerPage.ValueInt64())
		queryParams1.StartingAfter = networksSmDevicesPerformanceHistory.StartingAfter.ValueString()
		queryParams1.EndingBefore = networksSmDevicesPerformanceHistory.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Sm.GetNetworkSmDevicePerformanceHistory(vvNetworkID, vvDeviceID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmDevicePerformanceHistory",
				err.Error(),
			)
			return
		}

		networksSmDevicesPerformanceHistory = ResponseSmGetNetworkSmDevicePerformanceHistoryItemsToBody(networksSmDevicesPerformanceHistory, response1)
		diags = resp.State.Set(ctx, &networksSmDevicesPerformanceHistory)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSmDevicesPerformanceHistory struct {
	NetworkID     types.String                                          `tfsdk:"network_id"`
	DeviceID      types.String                                          `tfsdk:"device_id"`
	PerPage       types.Int64                                           `tfsdk:"per_page"`
	StartingAfter types.String                                          `tfsdk:"starting_after"`
	EndingBefore  types.String                                          `tfsdk:"ending_before"`
	Items         *[]ResponseItemSmGetNetworkSmDevicePerformanceHistory `tfsdk:"items"`
}

type ResponseItemSmGetNetworkSmDevicePerformanceHistory struct {
	CPUPercentUsed  types.Float64                                                `tfsdk:"cpu_percent_used"`
	DiskUsage       *ResponseItemSmGetNetworkSmDevicePerformanceHistoryDiskUsage `tfsdk:"disk_usage"`
	MemActive       types.Int64                                                  `tfsdk:"mem_active"`
	MemFree         types.Int64                                                  `tfsdk:"mem_free"`
	MemInactive     types.Int64                                                  `tfsdk:"mem_inactive"`
	MemWired        types.Int64                                                  `tfsdk:"mem_wired"`
	NetworkReceived types.Int64                                                  `tfsdk:"network_received"`
	NetworkSent     types.Int64                                                  `tfsdk:"network_sent"`
	SwapUsed        types.Int64                                                  `tfsdk:"swap_used"`
	Ts              types.String                                                 `tfsdk:"ts"`
}

type ResponseItemSmGetNetworkSmDevicePerformanceHistoryDiskUsage struct {
	C *ResponseItemSmGetNetworkSmDevicePerformanceHistoryDiskUsageC `tfsdk:"c"`
}

type ResponseItemSmGetNetworkSmDevicePerformanceHistoryDiskUsageC struct {
	Space types.Int64 `tfsdk:"space"`
	Used  types.Int64 `tfsdk:"used"`
}

// ToBody
func ResponseSmGetNetworkSmDevicePerformanceHistoryItemsToBody(state NetworksSmDevicesPerformanceHistory, response *merakigosdk.ResponseSmGetNetworkSmDevicePerformanceHistory) NetworksSmDevicesPerformanceHistory {
	var items []ResponseItemSmGetNetworkSmDevicePerformanceHistory
	for _, item := range *response {
		itemState := ResponseItemSmGetNetworkSmDevicePerformanceHistory{
			CPUPercentUsed: func() types.Float64 {
				if item.CPUPercentUsed != nil {
					return types.Float64Value(float64(*item.CPUPercentUsed))
				}
				return types.Float64{}
			}(),
			DiskUsage: func() *ResponseItemSmGetNetworkSmDevicePerformanceHistoryDiskUsage {
				if item.DiskUsage != nil {
					return &ResponseItemSmGetNetworkSmDevicePerformanceHistoryDiskUsage{
						C: func() *ResponseItemSmGetNetworkSmDevicePerformanceHistoryDiskUsageC {
							if item.DiskUsage.C != nil {
								return &ResponseItemSmGetNetworkSmDevicePerformanceHistoryDiskUsageC{
									Space: func() types.Int64 {
										if item.DiskUsage.C.Space != nil {
											return types.Int64Value(int64(*item.DiskUsage.C.Space))
										}
										return types.Int64{}
									}(),
									Used: func() types.Int64 {
										if item.DiskUsage.C.Used != nil {
											return types.Int64Value(int64(*item.DiskUsage.C.Used))
										}
										return types.Int64{}
									}(),
								}
							}
							return nil
						}(),
					}
				}
				return nil
			}(),
			MemActive: func() types.Int64 {
				if item.MemActive != nil {
					return types.Int64Value(int64(*item.MemActive))
				}
				return types.Int64{}
			}(),
			MemFree: func() types.Int64 {
				if item.MemFree != nil {
					return types.Int64Value(int64(*item.MemFree))
				}
				return types.Int64{}
			}(),
			MemInactive: func() types.Int64 {
				if item.MemInactive != nil {
					return types.Int64Value(int64(*item.MemInactive))
				}
				return types.Int64{}
			}(),
			MemWired: func() types.Int64 {
				if item.MemWired != nil {
					return types.Int64Value(int64(*item.MemWired))
				}
				return types.Int64{}
			}(),
			NetworkReceived: func() types.Int64 {
				if item.NetworkReceived != nil {
					return types.Int64Value(int64(*item.NetworkReceived))
				}
				return types.Int64{}
			}(),
			NetworkSent: func() types.Int64 {
				if item.NetworkSent != nil {
					return types.Int64Value(int64(*item.NetworkSent))
				}
				return types.Int64{}
			}(),
			SwapUsed: func() types.Int64 {
				if item.SwapUsed != nil {
					return types.Int64Value(int64(*item.SwapUsed))
				}
				return types.Int64{}
			}(),
			Ts: func() types.String {
				if item.Ts != "" {
					return types.StringValue(item.Ts)
				}
				return types.String{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
