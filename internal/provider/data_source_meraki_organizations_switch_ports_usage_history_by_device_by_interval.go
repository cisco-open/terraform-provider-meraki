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
	"time"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsSwitchPortsUsageHistoryByDeviceByIntervalDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSwitchPortsUsageHistoryByDeviceByIntervalDataSource{}
)

func NewOrganizationsSwitchPortsUsageHistoryByDeviceByIntervalDataSource() datasource.DataSource {
	return &OrganizationsSwitchPortsUsageHistoryByDeviceByIntervalDataSource{}
}

type OrganizationsSwitchPortsUsageHistoryByDeviceByIntervalDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSwitchPortsUsageHistoryByDeviceByIntervalDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSwitchPortsUsageHistoryByDeviceByIntervalDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_switch_ports_usage_history_by_device_by_interval"
}

func (d *OrganizationsSwitchPortsUsageHistoryByDeviceByIntervalDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"configuration_updated_after": schema.StringAttribute{
				MarkdownDescription: `configurationUpdatedAfter query parameter. Optional parameter to filter items to switches where the configuration has been updated after the given timestamp.`,
				Optional:            true,
			},
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"interval": schema.Int64Attribute{
				MarkdownDescription: `interval query parameter. The time interval in seconds for returned data. The valid intervals are: 300, 1200, 14400, 86400. The default is 1200. Interval is calculated if time params are provided.`,
				Optional:            true,
			},
			"mac": schema.StringAttribute{
				MarkdownDescription: `mac query parameter. Optional parameter to filter items to switches with MAC addresses that contain the search term or are an exact match.`,
				Optional:            true,
			},
			"macs": schema.ListAttribute{
				MarkdownDescription: `macs query parameter. Optional parameter to filter items to switches that have one of the provided MAC addresses.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `name query parameter. Optional parameter to filter items to switches with names that contain the search term or are an exact match.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter items to switches in one of the provided networks.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 50. Default is 10.`,
				Optional:            true,
			},
			"port_profile_ids": schema.ListAttribute{
				MarkdownDescription: `portProfileIds query parameter. Optional parameter to filter items to switches that contain switchports belonging to one of the specified port profiles.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial query parameter. Optional parameter to filter items to switches with serial number that contains the search term or are an exact match.`,
				Optional:            true,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Optional parameter to filter items to switches that have one of the provided serials.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
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
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. If interval is provided, the timespan will be autocalculated.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"items": schema.ListNestedAttribute{
						MarkdownDescription: `Switches`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"mac": schema.StringAttribute{
									MarkdownDescription: `The MAC address of the switch.`,
									Computed:            true,
								},
								"model": schema.StringAttribute{
									MarkdownDescription: `The model of the switch.`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `The name of the switch.`,
									Computed:            true,
								},
								"network": schema.SingleNestedAttribute{
									MarkdownDescription: `Identifying information of the switch's network.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `The ID of the network.`,
											Computed:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: `The name of the network.`,
											Computed:            true,
										},
									},
								},
								"ports": schema.SetNestedAttribute{
									MarkdownDescription: `The number of ports on the switch with usage data.`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"intervals": schema.SetNestedAttribute{
												MarkdownDescription: `An array of intervals for a port with bandwidth, traffic, and power usage data.`,
												Computed:            true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{

														"bandwidth": schema.SingleNestedAttribute{
															MarkdownDescription: `A breakdown of the average speed of data that has passed through this port during the interval.`,
															Computed:            true,
															Attributes: map[string]schema.Attribute{

																"usage": schema.SingleNestedAttribute{
																	MarkdownDescription: `Bandwidth usage data for the given interval.`,
																	Computed:            true,
																	Attributes: map[string]schema.Attribute{

																		"downstream": schema.Float64Attribute{
																			MarkdownDescription: `The average speed of the data received (in kilobits-per-second).`,
																			Computed:            true,
																		},
																		"total": schema.Float64Attribute{
																			MarkdownDescription: `The average speed of the data sent and received (in kilobits-per-second).`,
																			Computed:            true,
																		},
																		"upstream": schema.Float64Attribute{
																			MarkdownDescription: `The average speed of the data sent (in kilobits-per-second).`,
																			Computed:            true,
																		},
																	},
																},
															},
														},
														"data": schema.SingleNestedAttribute{
															MarkdownDescription: `A breakdown of how many kilobytes have passed through this port during the interval timespan.`,
															Computed:            true,
															Attributes: map[string]schema.Attribute{

																"usage": schema.SingleNestedAttribute{
																	MarkdownDescription: `Usage data for the given interval.`,
																	Computed:            true,
																	Attributes: map[string]schema.Attribute{

																		"downstream": schema.Int64Attribute{
																			MarkdownDescription: `The amount of data received (in kilobytes).`,
																			Computed:            true,
																		},
																		"total": schema.Int64Attribute{
																			MarkdownDescription: `The total amount of data sent and received (in kilobytes).`,
																			Computed:            true,
																		},
																		"upstream": schema.Int64Attribute{
																			MarkdownDescription: `The amount of data sent (in kilobytes).`,
																			Computed:            true,
																		},
																	},
																},
															},
														},
														"end_ts": schema.StringAttribute{
															MarkdownDescription: `The end timestamp of the given interval.`,
															Computed:            true,
														},
														"energy": schema.SingleNestedAttribute{
															MarkdownDescription: `How much energy (in watt-hours) has been delivered by this port during the interval.`,
															Computed:            true,
															Attributes: map[string]schema.Attribute{

																"usage": schema.SingleNestedAttribute{
																	MarkdownDescription: `Energy data for the given interval.`,
																	Computed:            true,
																	Attributes: map[string]schema.Attribute{

																		"total": schema.Float64Attribute{
																			MarkdownDescription: `The total energy in watt-hours delivered by this port during the interval`,
																			Computed:            true,
																		},
																	},
																},
															},
														},
														"start_ts": schema.StringAttribute{
															MarkdownDescription: `The starting timestamp of the given interval.`,
															Computed:            true,
														},
													},
												},
											},
											"port_id": schema.StringAttribute{
												MarkdownDescription: `The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module.`,
												Computed:            true,
											},
										},
									},
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `The serial number of the switch.`,
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

func (d *OrganizationsSwitchPortsUsageHistoryByDeviceByIntervalDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSwitchPortsUsageHistoryByDeviceByInterval OrganizationsSwitchPortsUsageHistoryByDeviceByInterval
	diags := req.Config.Get(ctx, &organizationsSwitchPortsUsageHistoryByDeviceByInterval)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSwitchPortsUsageHistoryByDeviceByInterval")
		vvOrganizationID := organizationsSwitchPortsUsageHistoryByDeviceByInterval.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalQueryParams{}

		queryParams1.T0 = organizationsSwitchPortsUsageHistoryByDeviceByInterval.T0.ValueString()
		queryParams1.T1 = organizationsSwitchPortsUsageHistoryByDeviceByInterval.T1.ValueString()
		queryParams1.Timespan = organizationsSwitchPortsUsageHistoryByDeviceByInterval.Timespan.ValueFloat64()
		queryParams1.Interval = int(organizationsSwitchPortsUsageHistoryByDeviceByInterval.Interval.ValueInt64())
		queryParams1.PerPage = int(organizationsSwitchPortsUsageHistoryByDeviceByInterval.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsSwitchPortsUsageHistoryByDeviceByInterval.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsSwitchPortsUsageHistoryByDeviceByInterval.EndingBefore.ValueString()
		queryParams1.ConfigurationUpdateAfter = organizationsSwitchPortsUsageHistoryByDeviceByInterval.ConfigurationUpdatedAfter.ValueString()
		queryParams1.Mac = organizationsSwitchPortsUsageHistoryByDeviceByInterval.Mac.ValueString()
		queryParams1.Macs = elementsToStrings(ctx, organizationsSwitchPortsUsageHistoryByDeviceByInterval.Macs)
		queryParams1.Name = organizationsSwitchPortsUsageHistoryByDeviceByInterval.Name.ValueString()
		queryParams1.NetworkIds = elementsToStrings(ctx, organizationsSwitchPortsUsageHistoryByDeviceByInterval.NetworkIDs)
		queryParams1.PortProfileIds = elementsToStrings(ctx, organizationsSwitchPortsUsageHistoryByDeviceByInterval.PortProfileIDs)
		queryParams1.Serial = organizationsSwitchPortsUsageHistoryByDeviceByInterval.Serial.ValueString()
		queryParams1.Serials = elementsToStrings(ctx, organizationsSwitchPortsUsageHistoryByDeviceByInterval.Serials)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetOrganizationSwitchPortsUsageHistoryByDeviceByInterval(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSwitchPortsUsageHistoryByDeviceByInterval",
				err.Error(),
			)
			return
		}

		organizationsSwitchPortsUsageHistoryByDeviceByInterval = ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemToBody(organizationsSwitchPortsUsageHistoryByDeviceByInterval, response1)
		diags = resp.State.Set(ctx, &organizationsSwitchPortsUsageHistoryByDeviceByInterval)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSwitchPortsUsageHistoryByDeviceByInterval struct {
	OrganizationID            types.String                                                            `tfsdk:"organization_id"`
	T0                        types.String                                                            `tfsdk:"t0"`
	T1                        types.String                                                            `tfsdk:"t1"`
	Timespan                  types.Float64                                                           `tfsdk:"timespan"`
	Interval                  types.Int64                                                             `tfsdk:"interval"`
	PerPage                   types.Int64                                                             `tfsdk:"per_page"`
	StartingAfter             types.String                                                            `tfsdk:"starting_after"`
	EndingBefore              types.String                                                            `tfsdk:"ending_before"`
	ConfigurationUpdatedAfter types.String                                                            `tfsdk:"configuration_updated_after"`
	Mac                       types.String                                                            `tfsdk:"mac"`
	Macs                      types.List                                                              `tfsdk:"macs"`
	Name                      types.String                                                            `tfsdk:"name"`
	NetworkIDs                types.List                                                              `tfsdk:"network_ids"`
	PortProfileIDs            types.List                                                              `tfsdk:"port_profile_ids"`
	Serial                    types.String                                                            `tfsdk:"serial"`
	Serials                   types.List                                                              `tfsdk:"serials"`
	Item                      *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByInterval `tfsdk:"item"`
}

type ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByInterval struct {
	Items *[]ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItems `tfsdk:"items"`
	Meta  *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalMeta    `tfsdk:"meta"`
}

type ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItems struct {
	Mac     types.String                                                                        `tfsdk:"mac"`
	Model   types.String                                                                        `tfsdk:"model"`
	Name    types.String                                                                        `tfsdk:"name"`
	Network *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsNetwork `tfsdk:"network"`
	Ports   *[]ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPorts `tfsdk:"ports"`
	Serial  types.String                                                                        `tfsdk:"serial"`
}

type ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPorts struct {
	Intervals *[]ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervals `tfsdk:"intervals"`
	PortID    types.String                                                                                 `tfsdk:"port_id"`
}

type ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervals struct {
	Bandwidth *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsBandwidth `tfsdk:"bandwidth"`
	Data      *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsData      `tfsdk:"data"`
	EndTs     types.String                                                                                        `tfsdk:"end_ts"`
	Energy    *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsEnergy    `tfsdk:"energy"`
	StartTs   types.String                                                                                        `tfsdk:"start_ts"`
}

type ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsBandwidth struct {
	Usage *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsBandwidthUsage `tfsdk:"usage"`
}

type ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsBandwidthUsage struct {
	Downstream types.Float64 `tfsdk:"downstream"`
	Total      types.Float64 `tfsdk:"total"`
	Upstream   types.Float64 `tfsdk:"upstream"`
}

type ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsData struct {
	Usage *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsDataUsage `tfsdk:"usage"`
}

type ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsDataUsage struct {
	Downstream types.Int64 `tfsdk:"downstream"`
	Total      types.Int64 `tfsdk:"total"`
	Upstream   types.Int64 `tfsdk:"upstream"`
}

type ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsEnergy struct {
	Usage *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsEnergyUsage `tfsdk:"usage"`
}

type ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsEnergyUsage struct {
	Total types.Float64 `tfsdk:"total"`
}

type ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalMeta struct {
	Counts *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalMetaCounts `tfsdk:"counts"`
}

type ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalMetaCounts struct {
	Items *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalMetaCountsItems `tfsdk:"items"`
}

type ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemToBody(state OrganizationsSwitchPortsUsageHistoryByDeviceByInterval, response *merakigosdk.ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByInterval) OrganizationsSwitchPortsUsageHistoryByDeviceByInterval {
	itemState := ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByInterval{
		Items: func() *[]ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItems {
			if response.Items != nil {
				result := make([]ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItems{
						Mac: func() types.String {
							if items.Mac != "" {
								return types.StringValue(items.Mac)
							}
							return types.String{}
						}(),
						Model: func() types.String {
							if items.Model != "" {
								return types.StringValue(items.Model)
							}
							return types.String{}
						}(),
						Name: func() types.String {
							if items.Name != "" {
								return types.StringValue(items.Name)
							}
							return types.String{}
						}(),
						Network: func() *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsNetwork {
							if items.Network != nil {
								return &ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsNetwork{
									ID: func() types.String {
										if items.Network.ID != "" {
											return types.StringValue(items.Network.ID)
										}
										return types.String{}
									}(),
									Name: func() types.String {
										if items.Network.Name != "" {
											return types.StringValue(items.Network.Name)
										}
										return types.String{}
									}(),
								}
							}
							return nil
						}(),
						Ports: func() *[]ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPorts {
							if items.Ports != nil {
								result := make([]ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPorts, len(*items.Ports))
								for i, ports := range *items.Ports {
									result[i] = ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPorts{
										Intervals: func() *[]ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervals {
											if ports.Intervals != nil {
												result := make([]ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervals, len(*ports.Intervals))
												for i, intervals := range *ports.Intervals {
													result[i] = ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervals{
														Bandwidth: func() *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsBandwidth {
															if intervals.Bandwidth.Usage != nil {
																return &ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsBandwidth{
																	Usage: func() *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsBandwidthUsage {
																		if intervals.Bandwidth.Usage != nil {
																			return &ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsBandwidthUsage{
																				Downstream: func() types.Float64 {
																					if intervals.Bandwidth.Usage.Downstream != nil {
																						return types.Float64Value(float64(*intervals.Bandwidth.Usage.Downstream))
																					}
																					return types.Float64{}
																				}(),
																				Total: func() types.Float64 {
																					if intervals.Bandwidth.Usage.Total != nil {
																						return types.Float64Value(float64(*intervals.Bandwidth.Usage.Total))
																					}
																					return types.Float64{}
																				}(),
																				Upstream: func() types.Float64 {
																					if intervals.Bandwidth.Usage.Upstream != nil {
																						return types.Float64Value(float64(*intervals.Bandwidth.Usage.Upstream))
																					}
																					return types.Float64{}
																				}(),
																			}
																		}
																		return nil
																	}(),
																}
															}
															return nil
														}(),
														Data: func() *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsData {
															if intervals.Data.Usage != nil {
																return &ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsData{
																	Usage: func() *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsDataUsage {
																		if intervals.Data.Usage != nil {
																			return &ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsDataUsage{
																				Downstream: func() types.Int64 {
																					if intervals.Data.Usage.Downstream != nil {
																						return types.Int64Value(int64(*intervals.Data.Usage.Downstream))
																					}
																					return types.Int64{}
																				}(),
																				Total: func() types.Int64 {
																					if intervals.Data.Usage.Total != nil {
																						return types.Int64Value(int64(*intervals.Data.Usage.Total))
																					}
																					return types.Int64{}
																				}(),
																				Upstream: func() types.Int64 {
																					if intervals.Data.Usage.Upstream != nil {
																						return types.Int64Value(int64(*intervals.Data.Usage.Upstream))
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
														EndTs: func() types.String {
															if intervals.EndTs.Format(time.RFC3339) != "" {
																return types.StringValue(intervals.EndTs.Format(time.RFC3339))
															}
															return types.String{}
														}(),
														Energy: func() *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsEnergy {
															if intervals.Energy.Usage != nil {
																return &ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsEnergy{
																	Usage: func() *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsEnergyUsage {
																		if intervals.Energy.Usage != nil {
																			return &ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalItemsPortsIntervalsEnergyUsage{
																				Total: func() types.Float64 {
																					if intervals.Energy.Usage.Total != nil {
																						return types.Float64Value(float64(*intervals.Energy.Usage.Total))
																					}
																					return types.Float64{}
																				}(),
																			}
																		}
																		return nil
																	}(),
																}
															}
															return nil
														}(),
														StartTs: func() types.String {
															if intervals.StartTs.Format(time.RFC3339) != "" {
																return types.StringValue(intervals.StartTs.Format(time.RFC3339))
															}
															return types.String{}
														}(),
													}
												}
												return &result
											}
											return nil
										}(),
										PortID: func() types.String {
											if ports.PortID != "" {
												return types.StringValue(ports.PortID)
											}
											return types.String{}
										}(),
									}
								}
								return &result
							}
							return nil
						}(),
						Serial: func() types.String {
							if items.Serial != "" {
								return types.StringValue(items.Serial)
							}
							return types.String{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		Meta: func() *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalMeta {
			if response.Meta != nil {
				return &ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalMeta{
					Counts: func() *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalMetaCounts{
								Items: func() *ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseSwitchGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalMetaCountsItems{
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
