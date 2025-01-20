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
	_ datasource.DataSource              = &OrganizationsWirelessControllerDevicesSystemUtilizationHistoryByIntervalDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessControllerDevicesSystemUtilizationHistoryByIntervalDataSource{}
)

func NewOrganizationsWirelessControllerDevicesSystemUtilizationHistoryByIntervalDataSource() datasource.DataSource {
	return &OrganizationsWirelessControllerDevicesSystemUtilizationHistoryByIntervalDataSource{}
}

type OrganizationsWirelessControllerDevicesSystemUtilizationHistoryByIntervalDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessControllerDevicesSystemUtilizationHistoryByIntervalDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessControllerDevicesSystemUtilizationHistoryByIntervalDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_controller_devices_system_utilization_history_by_interval"
}

func (d *OrganizationsWirelessControllerDevicesSystemUtilizationHistoryByIntervalDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
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
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"items": schema.ListNestedAttribute{
						MarkdownDescription: `Wireless LAN controller CPU usage data`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"intervals": schema.SetNestedAttribute{
									MarkdownDescription: `Time interval snapshots of CPU usage data of the wireless LAN controller`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"by_core": schema.SetNestedAttribute{
												MarkdownDescription: `The CPU usage per core of the wireless LAN controller`,
												Computed:            true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{

														"name": schema.StringAttribute{
															MarkdownDescription: `The CPU core name`,
															Computed:            true,
														},
														"usage": schema.SingleNestedAttribute{
															MarkdownDescription: `The specific core CPU usage of the wireless LAN controller`,
															Computed:            true,
															Attributes: map[string]schema.Attribute{

																"average": schema.SingleNestedAttribute{
																	MarkdownDescription: `The specific core average CPU usage of the wireless LAN controller`,
																	Computed:            true,
																	Attributes: map[string]schema.Attribute{

																		"percentage": schema.Float64Attribute{
																			MarkdownDescription: `The specific core CPU usage percentage of the wireless LAN controller`,
																			Computed:            true,
																		},
																	},
																},
															},
														},
													},
												},
											},
											"end_ts": schema.StringAttribute{
												MarkdownDescription: `The end time of the query range  with iso8601 format`,
												Computed:            true,
											},
											"overall": schema.SingleNestedAttribute{
												MarkdownDescription: `The overall CPU usage of the wireless LAN controller`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"usage": schema.SingleNestedAttribute{
														MarkdownDescription: `The CPU usage of the wireless LAN controller`,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"average": schema.SingleNestedAttribute{
																MarkdownDescription: `The average CPU usage of the wireless LAN controller`,
																Computed:            true,
																Attributes: map[string]schema.Attribute{

																	"percentage": schema.Float64Attribute{
																		MarkdownDescription: `The CPU usage percentage of the wireless LAN controller`,
																		Computed:            true,
																	},
																},
															},
														},
													},
												},
											},
											"start_ts": schema.StringAttribute{
												MarkdownDescription: `The start time of the query range with iso8601 format`,
												Computed:            true,
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

func (d *OrganizationsWirelessControllerDevicesSystemUtilizationHistoryByIntervalDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessControllerDevicesSystemUtilizationHistoryByInterval OrganizationsWirelessControllerDevicesSystemUtilizationHistoryByInterval
	diags := req.Config.Get(ctx, &organizationsWirelessControllerDevicesSystemUtilizationHistoryByInterval)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval")
		vvOrganizationID := organizationsWirelessControllerDevicesSystemUtilizationHistoryByInterval.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalQueryParams{}

		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessControllerDevicesSystemUtilizationHistoryByInterval.Serials)
		queryParams1.T0 = organizationsWirelessControllerDevicesSystemUtilizationHistoryByInterval.T0.ValueString()
		queryParams1.T1 = organizationsWirelessControllerDevicesSystemUtilizationHistoryByInterval.T1.ValueString()
		queryParams1.Timespan = organizationsWirelessControllerDevicesSystemUtilizationHistoryByInterval.Timespan.ValueFloat64()
		queryParams1.PerPage = int(organizationsWirelessControllerDevicesSystemUtilizationHistoryByInterval.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessControllerDevicesSystemUtilizationHistoryByInterval.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessControllerDevicesSystemUtilizationHistoryByInterval.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.WirelessController.GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval",
				err.Error(),
			)
			return
		}

		organizationsWirelessControllerDevicesSystemUtilizationHistoryByInterval = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemToBody(organizationsWirelessControllerDevicesSystemUtilizationHistoryByInterval, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessControllerDevicesSystemUtilizationHistoryByInterval)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessControllerDevicesSystemUtilizationHistoryByInterval struct {
	OrganizationID types.String                                                                                          `tfsdk:"organization_id"`
	Serials        types.List                                                                                            `tfsdk:"serials"`
	T0             types.String                                                                                          `tfsdk:"t0"`
	T1             types.String                                                                                          `tfsdk:"t1"`
	Timespan       types.Float64                                                                                         `tfsdk:"timespan"`
	PerPage        types.Int64                                                                                           `tfsdk:"per_page"`
	StartingAfter  types.String                                                                                          `tfsdk:"starting_after"`
	EndingBefore   types.String                                                                                          `tfsdk:"ending_before"`
	Item           *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval `tfsdk:"item"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItems `tfsdk:"items"`
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMeta    `tfsdk:"meta"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItems struct {
	Intervals *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervals `tfsdk:"intervals"`
	Serial    types.String                                                                                                          `tfsdk:"serial"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervals struct {
	ByCore  *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCore `tfsdk:"by_core"`
	EndTs   types.String                                                                                                                `tfsdk:"end_ts"`
	Overall *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverall  `tfsdk:"overall"`
	StartTs types.String                                                                                                                `tfsdk:"start_ts"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCore struct {
	Name  types.String                                                                                                                   `tfsdk:"name"`
	Usage *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCoreUsage `tfsdk:"usage"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCoreUsage struct {
	Average *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCoreUsageAverage `tfsdk:"average"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCoreUsageAverage struct {
	Percentage types.Float64 `tfsdk:"percentage"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverall struct {
	Usage *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverallUsage `tfsdk:"usage"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverallUsage struct {
	Average *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverallUsageAverage `tfsdk:"average"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverallUsageAverage struct {
	Percentage types.Float64 `tfsdk:"percentage"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMetaCounts `tfsdk:"counts"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMetaCountsItems `tfsdk:"items"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemToBody(state OrganizationsWirelessControllerDevicesSystemUtilizationHistoryByInterval, response *merakigosdk.ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval) OrganizationsWirelessControllerDevicesSystemUtilizationHistoryByInterval {
	itemState := ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval{
		Items: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItems {
			if response.Items != nil {
				result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItems{
						Intervals: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervals {
							if items.Intervals != nil {
								result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervals, len(*items.Intervals))
								for i, intervals := range *items.Intervals {
									result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervals{
										ByCore: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCore {
											if intervals.ByCore != nil {
												result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCore, len(*intervals.ByCore))
												for i, byCore := range *intervals.ByCore {
													result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCore{
														Name: types.StringValue(byCore.Name),
														Usage: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCoreUsage {
															if byCore.Usage != nil {
																return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCoreUsage{
																	Average: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCoreUsageAverage {
																		if byCore.Usage.Average != nil {
																			return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsByCoreUsageAverage{
																				Percentage: func() types.Float64 {
																					if byCore.Usage.Average.Percentage != nil {
																						return types.Float64Value(float64(*byCore.Usage.Average.Percentage))
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
													}
												}
												return &result
											}
											return nil
										}(),
										EndTs: types.StringValue(intervals.EndTs),
										Overall: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverall {
											if intervals.Overall != nil {
												return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverall{
													Usage: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverallUsage {
														if intervals.Overall.Usage != nil {
															return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverallUsage{
																Average: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverallUsageAverage {
																	if intervals.Overall.Usage.Average != nil {
																		return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalItemsIntervalsOverallUsageAverage{
																			Percentage: func() types.Float64 {
																				if intervals.Overall.Usage.Average.Percentage != nil {
																					return types.Float64Value(float64(*intervals.Overall.Usage.Average.Percentage))
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
						Serial: types.StringValue(items.Serial),
					}
				}
				return &result
			}
			return nil
		}(),
		Meta: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMeta {
			if response.Meta != nil {
				return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMeta{
					Counts: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMetaCounts{
								Items: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalMetaCountsItems{
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
