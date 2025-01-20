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
	_ datasource.DataSource              = &OrganizationsWirelessControllerDevicesInterfacesPacketsOverviewByDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessControllerDevicesInterfacesPacketsOverviewByDeviceDataSource{}
)

func NewOrganizationsWirelessControllerDevicesInterfacesPacketsOverviewByDeviceDataSource() datasource.DataSource {
	return &OrganizationsWirelessControllerDevicesInterfacesPacketsOverviewByDeviceDataSource{}
}

type OrganizationsWirelessControllerDevicesInterfacesPacketsOverviewByDeviceDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessControllerDevicesInterfacesPacketsOverviewByDeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessControllerDevicesInterfacesPacketsOverviewByDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_controller_devices_interfaces_packets_overview_by_device"
}

func (d *OrganizationsWirelessControllerDevicesInterfacesPacketsOverviewByDeviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"names": schema.ListAttribute{
				MarkdownDescription: `names query parameter. Optional parameter to filter wireless LAN controller by its interface name. This filter uses multiple exact matches.`,
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
				MarkdownDescription: `serials query parameter. Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 1 day from today.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 1 day after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 1 day. The default is 1 hour.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"items": schema.ListNestedAttribute{
						MarkdownDescription: `Wireless LAN controller interfaces packets statuses`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"interfaces": schema.SetNestedAttribute{
									MarkdownDescription: `Interfaces belongs to the wireless LAN controller`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"name": schema.StringAttribute{
												MarkdownDescription: `The name of the wireless LAN controller interface`,
												Computed:            true,
											},
											"readings": schema.SetNestedAttribute{
												MarkdownDescription: `The status of packets counter on the interfaces of the wireless LAN controller`,
												Computed:            true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{

														"name": schema.StringAttribute{
															MarkdownDescription: `The type of packets being counted`,
															Computed:            true,
														},
														"rate": schema.SingleNestedAttribute{
															MarkdownDescription: `The interface packet rates measured in packets per second`,
															Computed:            true,
															Attributes: map[string]schema.Attribute{

																"recv": schema.Int64Attribute{
																	MarkdownDescription: `The rate of packets received during the timespan`,
																	Computed:            true,
																},
																"send": schema.Int64Attribute{
																	MarkdownDescription: `The rate of packets sent during the timespan`,
																	Computed:            true,
																},
																"total": schema.Int64Attribute{
																	MarkdownDescription: `The rate of all packets sent and received during the timespan`,
																	Computed:            true,
																},
															},
														},
														"recv": schema.Int64Attribute{
															MarkdownDescription: `The total count of packets received by the interface during the timespan`,
															Computed:            true,
														},
														"send": schema.Int64Attribute{
															MarkdownDescription: `The total count of packets sent by the interface during the timespan`,
															Computed:            true,
														},
														"total": schema.Int64Attribute{
															MarkdownDescription: `The total count of sent and received packets during the timespan`,
															Computed:            true,
														},
													},
												},
											},
										},
									},
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `The cloud ID of the wireless LAN controller`,
									Computed:            true,
								},
							},
						},
					},
					"meta": schema.SingleNestedAttribute{
						MarkdownDescription: `Metadata relevant to the paginated dataset`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"counts": schema.SingleNestedAttribute{
								MarkdownDescription: `Counts relating to the paginated dataset`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"items": schema.SingleNestedAttribute{
										MarkdownDescription: `Counts relating to the paginated items`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"remaining": schema.Int64Attribute{
												MarkdownDescription: `The number of items in the dataset that are available on subsequent pages`,
												Computed:            true,
											},
											"total": schema.Int64Attribute{
												MarkdownDescription: `The total number of items in the dataset`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsWirelessControllerDevicesInterfacesPacketsOverviewByDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessControllerDevicesInterfacesPacketsOverviewByDevice OrganizationsWirelessControllerDevicesInterfacesPacketsOverviewByDevice
	diags := req.Config.Get(ctx, &organizationsWirelessControllerDevicesInterfacesPacketsOverviewByDevice)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice")
		vvOrganizationID := organizationsWirelessControllerDevicesInterfacesPacketsOverviewByDevice.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceQueryParams{}

		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessControllerDevicesInterfacesPacketsOverviewByDevice.Serials)
		queryParams1.Names = elementsToStrings(ctx, organizationsWirelessControllerDevicesInterfacesPacketsOverviewByDevice.Names)
		queryParams1.T0 = organizationsWirelessControllerDevicesInterfacesPacketsOverviewByDevice.T0.ValueString()
		queryParams1.T1 = organizationsWirelessControllerDevicesInterfacesPacketsOverviewByDevice.T1.ValueString()
		queryParams1.Timespan = organizationsWirelessControllerDevicesInterfacesPacketsOverviewByDevice.Timespan.ValueFloat64()
		queryParams1.PerPage = int(organizationsWirelessControllerDevicesInterfacesPacketsOverviewByDevice.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessControllerDevicesInterfacesPacketsOverviewByDevice.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessControllerDevicesInterfacesPacketsOverviewByDevice.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.WirelessController.GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice",
				err.Error(),
			)
			return
		}

		organizationsWirelessControllerDevicesInterfacesPacketsOverviewByDevice = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemToBody(organizationsWirelessControllerDevicesInterfacesPacketsOverviewByDevice, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessControllerDevicesInterfacesPacketsOverviewByDevice)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessControllerDevicesInterfacesPacketsOverviewByDevice struct {
	OrganizationID types.String                                                                                         `tfsdk:"organization_id"`
	Serials        types.List                                                                                           `tfsdk:"serials"`
	Names          types.List                                                                                           `tfsdk:"names"`
	T0             types.String                                                                                         `tfsdk:"t0"`
	T1             types.String                                                                                         `tfsdk:"t1"`
	Timespan       types.Float64                                                                                        `tfsdk:"timespan"`
	PerPage        types.Int64                                                                                          `tfsdk:"per_page"`
	StartingAfter  types.String                                                                                         `tfsdk:"starting_after"`
	EndingBefore   types.String                                                                                         `tfsdk:"ending_before"`
	Item           *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice `tfsdk:"item"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItems `tfsdk:"items"`
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMeta    `tfsdk:"meta"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItems struct {
	Interfaces *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfaces `tfsdk:"interfaces"`
	Serial     types.String                                                                                                          `tfsdk:"serial"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfaces struct {
	Name     types.String                                                                                                                  `tfsdk:"name"`
	Readings *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfacesReadings `tfsdk:"readings"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfacesReadings struct {
	Name  types.String                                                                                                                    `tfsdk:"name"`
	Rate  *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfacesReadingsRate `tfsdk:"rate"`
	Recv  types.Int64                                                                                                                     `tfsdk:"recv"`
	Send  types.Int64                                                                                                                     `tfsdk:"send"`
	Total types.Int64                                                                                                                     `tfsdk:"total"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfacesReadingsRate struct {
	Recv  types.Int64 `tfsdk:"recv"`
	Send  types.Int64 `tfsdk:"send"`
	Total types.Int64 `tfsdk:"total"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMetaCounts `tfsdk:"counts"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMetaCountsItems `tfsdk:"items"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemToBody(state OrganizationsWirelessControllerDevicesInterfacesPacketsOverviewByDevice, response *merakigosdk.ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice) OrganizationsWirelessControllerDevicesInterfacesPacketsOverviewByDevice {
	itemState := ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice{
		Items: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItems {
			if response.Items != nil {
				result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItems{
						Interfaces: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfaces {
							if items.Interfaces != nil {
								result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfaces, len(*items.Interfaces))
								for i, interfaces := range *items.Interfaces {
									result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfaces{
										Name: types.StringValue(interfaces.Name),
										Readings: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfacesReadings {
											if interfaces.Readings != nil {
												result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfacesReadings, len(*interfaces.Readings))
												for i, readings := range *interfaces.Readings {
													result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfacesReadings{
														Name: types.StringValue(readings.Name),
														Rate: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfacesReadingsRate {
															if readings.Rate != nil {
																return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceItemsInterfacesReadingsRate{
																	Recv: func() types.Int64 {
																		if readings.Rate.Recv != nil {
																			return types.Int64Value(int64(*readings.Rate.Recv))
																		}
																		return types.Int64{}
																	}(),
																	Send: func() types.Int64 {
																		if readings.Rate.Send != nil {
																			return types.Int64Value(int64(*readings.Rate.Send))
																		}
																		return types.Int64{}
																	}(),
																	Total: func() types.Int64 {
																		if readings.Rate.Total != nil {
																			return types.Int64Value(int64(*readings.Rate.Total))
																		}
																		return types.Int64{}
																	}(),
																}
															}
															return nil
														}(),
														Recv: func() types.Int64 {
															if readings.Recv != nil {
																return types.Int64Value(int64(*readings.Recv))
															}
															return types.Int64{}
														}(),
														Send: func() types.Int64 {
															if readings.Send != nil {
																return types.Int64Value(int64(*readings.Send))
															}
															return types.Int64{}
														}(),
														Total: func() types.Int64 {
															if readings.Total != nil {
																return types.Int64Value(int64(*readings.Total))
															}
															return types.Int64{}
														}(),
													}
												}
												return &result
											}
											return nil
										}(),
									}
								}
								return &result
							}
							return nil
						}(),
						Serial: types.StringValue(items.Serial),
					}
				}
				return &result
			}
			return nil
		}(),
		Meta: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMeta {
			if response.Meta != nil {
				return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMeta{
					Counts: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMetaCounts{
								Items: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceMetaCountsItems{
											Remaining: func() types.Int64 {
												if response.Meta.Counts.Items.Remaining != nil {
													return types.Int64Value(int64(*response.Meta.Counts.Items.Remaining))
												}
												return types.Int64{}
											}(),
											Total: func() types.Int64 {
												if response.Meta.Counts.Items.Total != nil {
													return types.Int64Value(int64(*response.Meta.Counts.Items.Total))
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
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
