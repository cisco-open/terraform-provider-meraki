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
	_ datasource.DataSource              = &OrganizationsWirelessDevicesChannelUtilizationByDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessDevicesChannelUtilizationByDeviceDataSource{}
)

func NewOrganizationsWirelessDevicesChannelUtilizationByDeviceDataSource() datasource.DataSource {
	return &OrganizationsWirelessDevicesChannelUtilizationByDeviceDataSource{}
}

type OrganizationsWirelessDevicesChannelUtilizationByDeviceDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessDevicesChannelUtilizationByDeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessDevicesChannelUtilizationByDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_devices_channel_utilization_by_device"
}

func (d *OrganizationsWirelessDevicesChannelUtilizationByDeviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"interval": schema.Int64Attribute{
				MarkdownDescription: `interval query parameter. The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Filter results by network.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.`,
				Optional:            true,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Filter results by device.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 90 days from today.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 90 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 90 days. The default is 7 days.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationByDevice`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"by_band": schema.SetNestedAttribute{
							MarkdownDescription: `Channel utilization broken down by band.`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"band": schema.StringAttribute{
										MarkdownDescription: `The band for the given metrics.`,
										Computed:            true,
									},
									"non_wifi": schema.SingleNestedAttribute{
										MarkdownDescription: `An object containing non-wifi utilization.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"percentage": schema.Float64Attribute{
												MarkdownDescription: `Percentage of non-wifi channel utiliation for the given band.`,
												Computed:            true,
											},
										},
									},
									"total": schema.SingleNestedAttribute{
										MarkdownDescription: `An object containing total channel utilization.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"percentage": schema.Float64Attribute{
												MarkdownDescription: `Percentage of total channel utiliation for the given band.`,
												Computed:            true,
											},
										},
									},
									"wifi": schema.SingleNestedAttribute{
										MarkdownDescription: `An object containing wifi utilization.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"percentage": schema.Float64Attribute{
												MarkdownDescription: `Percentage of wifi channel utiliation for the given band.`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
						"mac": schema.StringAttribute{
							MarkdownDescription: `The MAC address of the device.`,
							Computed:            true,
						},
						"network": schema.SingleNestedAttribute{
							MarkdownDescription: `Network for the given utilization metrics.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `Network ID of the given utilization metrics.`,
									Computed:            true,
								},
							},
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `The serial number for the device.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsWirelessDevicesChannelUtilizationByDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessDevicesChannelUtilizationByDevice OrganizationsWirelessDevicesChannelUtilizationByDevice
	diags := req.Config.Get(ctx, &organizationsWirelessDevicesChannelUtilizationByDevice)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessDevicesChannelUtilizationByDevice")
		vvOrganizationID := organizationsWirelessDevicesChannelUtilizationByDevice.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessDevicesChannelUtilizationByDeviceQueryParams{}

		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsWirelessDevicesChannelUtilizationByDevice.NetworkIDs)
		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessDevicesChannelUtilizationByDevice.Serials)
		queryParams1.PerPage = int(organizationsWirelessDevicesChannelUtilizationByDevice.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessDevicesChannelUtilizationByDevice.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessDevicesChannelUtilizationByDevice.EndingBefore.ValueString()
		queryParams1.T0 = organizationsWirelessDevicesChannelUtilizationByDevice.T0.ValueString()
		queryParams1.T1 = organizationsWirelessDevicesChannelUtilizationByDevice.T1.ValueString()
		queryParams1.Timespan = organizationsWirelessDevicesChannelUtilizationByDevice.Timespan.ValueFloat64()
		queryParams1.Interval = int(organizationsWirelessDevicesChannelUtilizationByDevice.Interval.ValueInt64())

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetOrganizationWirelessDevicesChannelUtilizationByDevice(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessDevicesChannelUtilizationByDevice",
				err.Error(),
			)
			return
		}

		organizationsWirelessDevicesChannelUtilizationByDevice = ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceItemsToBody(organizationsWirelessDevicesChannelUtilizationByDevice, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessDevicesChannelUtilizationByDevice)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessDevicesChannelUtilizationByDevice struct {
	OrganizationID types.String                                                                    `tfsdk:"organization_id"`
	NetworkIDs     types.List                                                                      `tfsdk:"network_ids"`
	Serials        types.List                                                                      `tfsdk:"serials"`
	PerPage        types.Int64                                                                     `tfsdk:"per_page"`
	StartingAfter  types.String                                                                    `tfsdk:"starting_after"`
	EndingBefore   types.String                                                                    `tfsdk:"ending_before"`
	T0             types.String                                                                    `tfsdk:"t0"`
	T1             types.String                                                                    `tfsdk:"t1"`
	Timespan       types.Float64                                                                   `tfsdk:"timespan"`
	Interval       types.Int64                                                                     `tfsdk:"interval"`
	Items          *[]ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDevice `tfsdk:"items"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDevice struct {
	ByBand  *[]ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBand `tfsdk:"by_band"`
	Mac     types.String                                                                          `tfsdk:"mac"`
	Network *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceNetwork  `tfsdk:"network"`
	Serial  types.String                                                                          `tfsdk:"serial"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBand struct {
	Band    types.String                                                                               `tfsdk:"band"`
	NonWifi *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandNonWifi `tfsdk:"non_wifi"`
	Total   *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandTotal   `tfsdk:"total"`
	Wifi    *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandWifi    `tfsdk:"wifi"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandNonWifi struct {
	Percentage types.Float64 `tfsdk:"percentage"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandTotal struct {
	Percentage types.Float64 `tfsdk:"percentage"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandWifi struct {
	Percentage types.Float64 `tfsdk:"percentage"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceNetwork struct {
	ID types.String `tfsdk:"id"`
}

// ToBody
func ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceItemsToBody(state OrganizationsWirelessDevicesChannelUtilizationByDevice, response *merakigosdk.ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationByDevice) OrganizationsWirelessDevicesChannelUtilizationByDevice {
	var items []ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDevice
	for _, item := range *response {
		itemState := ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDevice{
			ByBand: func() *[]ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBand {
				if item.ByBand != nil {
					result := make([]ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBand, len(*item.ByBand))
					for i, byBand := range *item.ByBand {
						result[i] = ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBand{
							Band: types.StringValue(byBand.Band),
							NonWifi: func() *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandNonWifi {
								if byBand.NonWifi != nil {
									return &ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandNonWifi{
										Percentage: func() types.Float64 {
											if byBand.NonWifi.Percentage != nil {
												return types.Float64Value(float64(*byBand.NonWifi.Percentage))
											}
											return types.Float64{}
										}(),
									}
								}
								return nil
							}(),
							Total: func() *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandTotal {
								if byBand.Total != nil {
									return &ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandTotal{
										Percentage: func() types.Float64 {
											if byBand.Total.Percentage != nil {
												return types.Float64Value(float64(*byBand.Total.Percentage))
											}
											return types.Float64{}
										}(),
									}
								}
								return nil
							}(),
							Wifi: func() *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandWifi {
								if byBand.Wifi != nil {
									return &ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceByBandWifi{
										Percentage: func() types.Float64 {
											if byBand.Wifi.Percentage != nil {
												return types.Float64Value(float64(*byBand.Wifi.Percentage))
											}
											return types.Float64{}
										}(),
									}
								}
								return nil
							}(),
						}
					}
					return &result
				}
				return nil
			}(),
			Mac: types.StringValue(item.Mac),
			Network: func() *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceNetwork {
				if item.Network != nil {
					return &ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByDeviceNetwork{
						ID: types.StringValue(item.Network.ID),
					}
				}
				return nil
			}(),
			Serial: types.StringValue(item.Serial),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
