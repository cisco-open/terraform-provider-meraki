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
	_ datasource.DataSource              = &OrganizationsDevicesSystemMemoryUsageHistoryByIntervalDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsDevicesSystemMemoryUsageHistoryByIntervalDataSource{}
)

func NewOrganizationsDevicesSystemMemoryUsageHistoryByIntervalDataSource() datasource.DataSource {
	return &OrganizationsDevicesSystemMemoryUsageHistoryByIntervalDataSource{}
}

type OrganizationsDevicesSystemMemoryUsageHistoryByIntervalDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsDevicesSystemMemoryUsageHistoryByIntervalDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsDevicesSystemMemoryUsageHistoryByIntervalDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_devices_system_memory_usage_history_by_interval"
}

func (d *OrganizationsDevicesSystemMemoryUsageHistoryByIntervalDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"interval": schema.Int64Attribute{
				MarkdownDescription: `interval query parameter. The time interval in seconds for returned data. The valid intervals are: 300, 1200, 3600, 14400. The default is 300. Interval is calculated if time params are provided.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter the result set by the included set of network IDs`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 20. Default is 10.`,
				Optional:            true,
			},
			"product_types": schema.ListAttribute{
				MarkdownDescription: `productTypes query parameter. Optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, and secureConnect.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Optional parameter to filter device availabilities history by device serial numbers`,
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
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 2 hours. If interval is provided, the timespan will be autocalculated.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"items": schema.ListNestedAttribute{
						MarkdownDescription: `The top-level property containing all memory utilization data.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"free": schema.SingleNestedAttribute{
									MarkdownDescription: `Information regarding memory availability on the device over the entire timespan`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"median": schema.Int64Attribute{
											MarkdownDescription: `Median memory in kB free on the device over the entire timespan rounded up to nearest integer`,
											Computed:            true,
										},
									},
								},
								"intervals": schema.SetNestedAttribute{
									MarkdownDescription: `Time interval snapshots of system memory utilization on the device with the most recent snapshot first`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"end_ts": schema.StringAttribute{
												MarkdownDescription: `Timestamp for the end of the historical snapshot, inclusive.`,
												Computed:            true,
											},
											"memory": schema.SingleNestedAttribute{
												MarkdownDescription: `Information regarding memory usage and availability on the device`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"free": schema.SingleNestedAttribute{
														MarkdownDescription: `Information regarding memory availability on the device over the interval`,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"maximum": schema.Int64Attribute{
																MarkdownDescription: `Maximum memory in kB free on the device over the interval`,
																Computed:            true,
															},
															"median": schema.Int64Attribute{
																MarkdownDescription: `Median memory in kB free on the device over the interval rounded up to nearest integer`,
																Computed:            true,
															},
															"minimum": schema.Int64Attribute{
																MarkdownDescription: `Minimum memory in kB free on the device over the interval`,
																Computed:            true,
															},
														},
													},
													"used": schema.SingleNestedAttribute{
														MarkdownDescription: `Information regarding memory usage on the device over the interval`,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"maximum": schema.Int64Attribute{
																MarkdownDescription: `Maximum memory in kB used on the device over the interval`,
																Computed:            true,
															},
															"median": schema.Int64Attribute{
																MarkdownDescription: `Median memory in kB used on the device over the interval rounded up to nearest integer`,
																Computed:            true,
															},
															"minimum": schema.Int64Attribute{
																MarkdownDescription: `Minimum memory in kB used on the device over the interval`,
																Computed:            true,
															},
															"percentages": schema.SingleNestedAttribute{
																MarkdownDescription: `Memory utilization percentages on the device over the interval`,
																Computed:            true,
																Attributes: map[string]schema.Attribute{

																	"maximum": schema.Int64Attribute{
																		MarkdownDescription: `Maximum memory utilization percentage on the device over the interval`,
																		Computed:            true,
																	},
																},
															},
														},
													},
												},
											},
											"start_ts": schema.StringAttribute{
												MarkdownDescription: `Timestamp for the beginning of the historical snapshot, exclusive.`,
												Computed:            true,
											},
										},
									},
								},
								"mac": schema.StringAttribute{
									MarkdownDescription: `MAC address of the device`,
									Computed:            true,
								},
								"model": schema.StringAttribute{
									MarkdownDescription: `Model of the device`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Name of the device`,
									Computed:            true,
								},
								"network": schema.SingleNestedAttribute{
									MarkdownDescription: `Information regarding the network the device belongs to`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `The network ID`,
											Computed:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: `The name of the network`,
											Computed:            true,
										},
										"tags": schema.ListAttribute{
											MarkdownDescription: `List of custom tags for the network`,
											Computed:            true,
											ElementType:         types.StringType,
										},
									},
								},
								"provisioned": schema.Int64Attribute{
									MarkdownDescription: `The total RAM size provisioned on the device, in kB`,
									Computed:            true,
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `Unique serial number for the device`,
									Computed:            true,
								},
								"tags": schema.ListAttribute{
									MarkdownDescription: `List of custom tags for the device`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"used": schema.SingleNestedAttribute{
									MarkdownDescription: `Information regarding memory usage on the device over the entire timespan`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"median": schema.Int64Attribute{
											MarkdownDescription: `Median memory in kB used on the device over the entire timespan rounded up to nearest integer`,
											Computed:            true,
										},
									},
								},
							},
						},
					},
					"meta": schema.SingleNestedAttribute{
						MarkdownDescription: `Other metadata related to this result set.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"counts": schema.SingleNestedAttribute{
								MarkdownDescription: `Count metadata related to this result set.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"items": schema.SingleNestedAttribute{
										MarkdownDescription: `The count metadata.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"remaining": schema.Int64Attribute{
												MarkdownDescription: `The number of serials remaining based on current pagination location within the dataset.`,
												Computed:            true,
											},
											"total": schema.Int64Attribute{
												MarkdownDescription: `The total number of serials.`,
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

func (d *OrganizationsDevicesSystemMemoryUsageHistoryByIntervalDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsDevicesSystemMemoryUsageHistoryByInterval OrganizationsDevicesSystemMemoryUsageHistoryByInterval
	diags := req.Config.Get(ctx, &organizationsDevicesSystemMemoryUsageHistoryByInterval)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationDevicesSystemMemoryUsageHistoryByInterval")
		vvOrganizationID := organizationsDevicesSystemMemoryUsageHistoryByInterval.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationDevicesSystemMemoryUsageHistoryByIntervalQueryParams{}

		queryParams1.PerPage = int(organizationsDevicesSystemMemoryUsageHistoryByInterval.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsDevicesSystemMemoryUsageHistoryByInterval.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsDevicesSystemMemoryUsageHistoryByInterval.EndingBefore.ValueString()
		queryParams1.T0 = organizationsDevicesSystemMemoryUsageHistoryByInterval.T0.ValueString()
		queryParams1.T1 = organizationsDevicesSystemMemoryUsageHistoryByInterval.T1.ValueString()
		queryParams1.Timespan = organizationsDevicesSystemMemoryUsageHistoryByInterval.Timespan.ValueFloat64()
		queryParams1.Interval = int(organizationsDevicesSystemMemoryUsageHistoryByInterval.Interval.ValueInt64())
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsDevicesSystemMemoryUsageHistoryByInterval.NetworkIDs)
		queryParams1.Serials = elementsToStrings(ctx, organizationsDevicesSystemMemoryUsageHistoryByInterval.Serials)
		queryParams1.ProductTypes = elementsToStrings(ctx, organizationsDevicesSystemMemoryUsageHistoryByInterval.ProductTypes)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationDevicesSystemMemoryUsageHistoryByInterval(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationDevicesSystemMemoryUsageHistoryByInterval",
				err.Error(),
			)
			return
		}

		organizationsDevicesSystemMemoryUsageHistoryByInterval = ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemToBody(organizationsDevicesSystemMemoryUsageHistoryByInterval, response1)
		diags = resp.State.Set(ctx, &organizationsDevicesSystemMemoryUsageHistoryByInterval)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsDevicesSystemMemoryUsageHistoryByInterval struct {
	OrganizationID types.String                                                                   `tfsdk:"organization_id"`
	PerPage        types.Int64                                                                    `tfsdk:"per_page"`
	StartingAfter  types.String                                                                   `tfsdk:"starting_after"`
	EndingBefore   types.String                                                                   `tfsdk:"ending_before"`
	T0             types.String                                                                   `tfsdk:"t0"`
	T1             types.String                                                                   `tfsdk:"t1"`
	Timespan       types.Float64                                                                  `tfsdk:"timespan"`
	Interval       types.Int64                                                                    `tfsdk:"interval"`
	NetworkIDs     types.List                                                                     `tfsdk:"network_ids"`
	Serials        types.List                                                                     `tfsdk:"serials"`
	ProductTypes   types.List                                                                     `tfsdk:"product_types"`
	Item           *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByInterval `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByInterval struct {
	Items *[]ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItems `tfsdk:"items"`
	Meta  *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalMeta    `tfsdk:"meta"`
}

type ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItems struct {
	Free        *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsFree        `tfsdk:"free"`
	Intervals   *[]ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervals `tfsdk:"intervals"`
	Mac         types.String                                                                                   `tfsdk:"mac"`
	Model       types.String                                                                                   `tfsdk:"model"`
	Name        types.String                                                                                   `tfsdk:"name"`
	Network     *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsNetwork     `tfsdk:"network"`
	Provisioned types.Int64                                                                                    `tfsdk:"provisioned"`
	Serial      types.String                                                                                   `tfsdk:"serial"`
	Tags        types.List                                                                                     `tfsdk:"tags"`
	Used        *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsUsed        `tfsdk:"used"`
}

type ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsFree struct {
	Median types.Int64 `tfsdk:"median"`
}

type ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervals struct {
	EndTs   types.String                                                                                       `tfsdk:"end_ts"`
	Memory  *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervalsMemory `tfsdk:"memory"`
	StartTs types.String                                                                                       `tfsdk:"start_ts"`
}

type ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervalsMemory struct {
	Free *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervalsMemoryFree `tfsdk:"free"`
	Used *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervalsMemoryUsed `tfsdk:"used"`
}

type ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervalsMemoryFree struct {
	Maximum types.Int64 `tfsdk:"maximum"`
	Median  types.Int64 `tfsdk:"median"`
	Minimum types.Int64 `tfsdk:"minimum"`
}

type ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervalsMemoryUsed struct {
	Maximum     types.Int64                                                                                                       `tfsdk:"maximum"`
	Median      types.Int64                                                                                                       `tfsdk:"median"`
	Minimum     types.Int64                                                                                                       `tfsdk:"minimum"`
	Percentages *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervalsMemoryUsedPercentages `tfsdk:"percentages"`
}

type ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervalsMemoryUsedPercentages struct {
	Maximum types.Int64 `tfsdk:"maximum"`
}

type ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Tags types.List   `tfsdk:"tags"`
}

type ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsUsed struct {
	Median types.Int64 `tfsdk:"median"`
}

type ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalMeta struct {
	Counts *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalMetaCounts `tfsdk:"counts"`
}

type ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalMetaCounts struct {
	Items *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalMetaCountsItems `tfsdk:"items"`
}

type ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemToBody(state OrganizationsDevicesSystemMemoryUsageHistoryByInterval, response *merakigosdk.ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByInterval) OrganizationsDevicesSystemMemoryUsageHistoryByInterval {
	itemState := ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByInterval{
		Items: func() *[]ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItems {
			if response.Items != nil {
				result := make([]ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItems{
						Free: func() *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsFree {
							if items.Free != nil {
								return &ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsFree{
									Median: func() types.Int64 {
										if items.Free.Median != nil {
											return types.Int64Value(int64(*items.Free.Median))
										}
										return types.Int64{}
									}(),
								}
							}
							return nil
						}(),
						Intervals: func() *[]ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervals {
							if items.Intervals != nil {
								result := make([]ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervals, len(*items.Intervals))
								for i, intervals := range *items.Intervals {
									result[i] = ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervals{
										EndTs: types.StringValue(intervals.EndTs),
										Memory: func() *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervalsMemory {
											if intervals.Memory != nil {
												return &ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervalsMemory{
													Free: func() *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervalsMemoryFree {
														if intervals.Memory.Free != nil {
															return &ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervalsMemoryFree{
																Maximum: func() types.Int64 {
																	if intervals.Memory.Free.Maximum != nil {
																		return types.Int64Value(int64(*intervals.Memory.Free.Maximum))
																	}
																	return types.Int64{}
																}(),
																Median: func() types.Int64 {
																	if intervals.Memory.Free.Median != nil {
																		return types.Int64Value(int64(*intervals.Memory.Free.Median))
																	}
																	return types.Int64{}
																}(),
																Minimum: func() types.Int64 {
																	if intervals.Memory.Free.Minimum != nil {
																		return types.Int64Value(int64(*intervals.Memory.Free.Minimum))
																	}
																	return types.Int64{}
																}(),
															}
														}
														return nil
													}(),
													Used: func() *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervalsMemoryUsed {
														if intervals.Memory.Used != nil {
															return &ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervalsMemoryUsed{
																Maximum: func() types.Int64 {
																	if intervals.Memory.Used.Maximum != nil {
																		return types.Int64Value(int64(*intervals.Memory.Used.Maximum))
																	}
																	return types.Int64{}
																}(),
																Median: func() types.Int64 {
																	if intervals.Memory.Used.Median != nil {
																		return types.Int64Value(int64(*intervals.Memory.Used.Median))
																	}
																	return types.Int64{}
																}(),
																Minimum: func() types.Int64 {
																	if intervals.Memory.Used.Minimum != nil {
																		return types.Int64Value(int64(*intervals.Memory.Used.Minimum))
																	}
																	return types.Int64{}
																}(),
																Percentages: func() *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervalsMemoryUsedPercentages {
																	if intervals.Memory.Used.Percentages != nil {
																		return &ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsIntervalsMemoryUsedPercentages{
																			Maximum: func() types.Int64 {
																				if intervals.Memory.Used.Percentages.Maximum != nil {
																					return types.Int64Value(int64(*intervals.Memory.Used.Percentages.Maximum))
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
										StartTs: types.StringValue(intervals.StartTs),
									}
								}
								return &result
							}
							return nil
						}(),
						Mac:   types.StringValue(items.Mac),
						Model: types.StringValue(items.Model),
						Name:  types.StringValue(items.Name),
						Network: func() *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsNetwork {
							if items.Network != nil {
								return &ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsNetwork{
									ID:   types.StringValue(items.Network.ID),
									Name: types.StringValue(items.Network.Name),
									Tags: StringSliceToList(items.Network.Tags),
								}
							}
							return nil
						}(),
						Provisioned: func() types.Int64 {
							if items.Provisioned != nil {
								return types.Int64Value(int64(*items.Provisioned))
							}
							return types.Int64{}
						}(),
						Serial: types.StringValue(items.Serial),
						Tags:   StringSliceToList(items.Tags),
						Used: func() *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsUsed {
							if items.Used != nil {
								return &ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalItemsUsed{
									Median: func() types.Int64 {
										if items.Used.Median != nil {
											return types.Int64Value(int64(*items.Used.Median))
										}
										return types.Int64{}
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
		Meta: func() *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalMeta {
			if response.Meta != nil {
				return &ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalMeta{
					Counts: func() *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalMetaCounts{
								Items: func() *ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseOrganizationsGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalMetaCountsItems{
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
