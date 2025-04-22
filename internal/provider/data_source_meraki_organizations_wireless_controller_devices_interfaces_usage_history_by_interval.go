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
	_ datasource.DataSource              = &OrganizationsWirelessControllerDevicesInterfacesUsageHistoryByIntervalDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessControllerDevicesInterfacesUsageHistoryByIntervalDataSource{}
)

func NewOrganizationsWirelessControllerDevicesInterfacesUsageHistoryByIntervalDataSource() datasource.DataSource {
	return &OrganizationsWirelessControllerDevicesInterfacesUsageHistoryByIntervalDataSource{}
}

type OrganizationsWirelessControllerDevicesInterfacesUsageHistoryByIntervalDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessControllerDevicesInterfacesUsageHistoryByIntervalDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessControllerDevicesInterfacesUsageHistoryByIntervalDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_controller_devices_interfaces_usage_history_by_interval"
}

func (d *OrganizationsWirelessControllerDevicesInterfacesUsageHistoryByIntervalDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
						MarkdownDescription: `Wireless LAN controller interfaces usage data`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"intervals": schema.SetNestedAttribute{
									MarkdownDescription: `Time interval snapshots of interfaces usage data of the wireless LAN controller`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"by_interface": schema.SetNestedAttribute{
												MarkdownDescription: `The usage data on the interfaces of the wireless LAN controller`,
												Computed:            true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{

														"name": schema.StringAttribute{
															MarkdownDescription: `The name of the wireless LAN controller interface`,
															Computed:            true,
														},
														"usage": schema.SingleNestedAttribute{
															MarkdownDescription: `The usage on the interfaces of the wireless LAN controller`,
															Computed:            true,
															Attributes: map[string]schema.Attribute{

																"recv": schema.Int64Attribute{
																	MarkdownDescription: `The received usage on the interface during the interval, unit is bit/sec`,
																	Computed:            true,
																},
																"send": schema.Int64Attribute{
																	MarkdownDescription: `The sent usage on the interface during the interval, unit is bit/sec`,
																	Computed:            true,
																},
																"total": schema.Int64Attribute{
																	MarkdownDescription: `The total usage on the interface during the interval, unit is bit/sec`,
																	Computed:            true,
																},
															},
														},
													},
												},
											},
											"end_ts": schema.StringAttribute{
												MarkdownDescription: `The end time interval snapshots of the query range with iso8601 format`,
												Computed:            true,
											},
											"overall": schema.SingleNestedAttribute{
												MarkdownDescription: `The overall usage of all queried interfaces of the wireless LAN controller`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"recv": schema.Int64Attribute{
														MarkdownDescription: `The received usage of all queried interfaces during the interval, unit is bit/sec`,
														Computed:            true,
													},
													"send": schema.Int64Attribute{
														MarkdownDescription: `The sent usage of all queried interfaces during the interval, unit is bit/sec`,
														Computed:            true,
													},
													"total": schema.Int64Attribute{
														MarkdownDescription: `The total usage of all queried interfaces during the interval, unit is bit/sec`,
														Computed:            true,
													},
												},
											},
											"start_ts": schema.StringAttribute{
												MarkdownDescription: `The start time interval snapshots of the query range with iso8601 format`,
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

func (d *OrganizationsWirelessControllerDevicesInterfacesUsageHistoryByIntervalDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessControllerDevicesInterfacesUsageHistoryByInterval OrganizationsWirelessControllerDevicesInterfacesUsageHistoryByInterval
	diags := req.Config.Get(ctx, &organizationsWirelessControllerDevicesInterfacesUsageHistoryByInterval)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval")
		vvOrganizationID := organizationsWirelessControllerDevicesInterfacesUsageHistoryByInterval.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalQueryParams{}

		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessControllerDevicesInterfacesUsageHistoryByInterval.Serials)
		queryParams1.Names = elementsToStrings(ctx, organizationsWirelessControllerDevicesInterfacesUsageHistoryByInterval.Names)
		queryParams1.T0 = organizationsWirelessControllerDevicesInterfacesUsageHistoryByInterval.T0.ValueString()
		queryParams1.T1 = organizationsWirelessControllerDevicesInterfacesUsageHistoryByInterval.T1.ValueString()
		queryParams1.Timespan = organizationsWirelessControllerDevicesInterfacesUsageHistoryByInterval.Timespan.ValueFloat64()
		queryParams1.PerPage = int(organizationsWirelessControllerDevicesInterfacesUsageHistoryByInterval.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessControllerDevicesInterfacesUsageHistoryByInterval.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessControllerDevicesInterfacesUsageHistoryByInterval.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.WirelessController.GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval",
				err.Error(),
			)
			return
		}

		organizationsWirelessControllerDevicesInterfacesUsageHistoryByInterval = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemToBody(organizationsWirelessControllerDevicesInterfacesUsageHistoryByInterval, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessControllerDevicesInterfacesUsageHistoryByInterval)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessControllerDevicesInterfacesUsageHistoryByInterval struct {
	OrganizationID types.String                                                                                        `tfsdk:"organization_id"`
	Serials        types.List                                                                                          `tfsdk:"serials"`
	Names          types.List                                                                                          `tfsdk:"names"`
	T0             types.String                                                                                        `tfsdk:"t0"`
	T1             types.String                                                                                        `tfsdk:"t1"`
	Timespan       types.Float64                                                                                       `tfsdk:"timespan"`
	PerPage        types.Int64                                                                                         `tfsdk:"per_page"`
	StartingAfter  types.String                                                                                        `tfsdk:"starting_after"`
	EndingBefore   types.String                                                                                        `tfsdk:"ending_before"`
	Item           *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval `tfsdk:"item"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItems `tfsdk:"items"`
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMeta    `tfsdk:"meta"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItems struct {
	Intervals *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervals `tfsdk:"intervals"`
	Serial    types.String                                                                                                        `tfsdk:"serial"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervals struct {
	ByInterface *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsByInterface `tfsdk:"by_interface"`
	EndTs       types.String                                                                                                                   `tfsdk:"end_ts"`
	Overall     *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsOverall       `tfsdk:"overall"`
	StartTs     types.String                                                                                                                   `tfsdk:"start_ts"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsByInterface struct {
	Name  types.String                                                                                                                      `tfsdk:"name"`
	Usage *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsByInterfaceUsage `tfsdk:"usage"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsByInterfaceUsage struct {
	Recv  types.Int64 `tfsdk:"recv"`
	Send  types.Int64 `tfsdk:"send"`
	Total types.Int64 `tfsdk:"total"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsOverall struct {
	Recv  types.Int64 `tfsdk:"recv"`
	Send  types.Int64 `tfsdk:"send"`
	Total types.Int64 `tfsdk:"total"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMetaCounts `tfsdk:"counts"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMetaCountsItems `tfsdk:"items"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemToBody(state OrganizationsWirelessControllerDevicesInterfacesUsageHistoryByInterval, response *merakigosdk.ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval) OrganizationsWirelessControllerDevicesInterfacesUsageHistoryByInterval {
	itemState := ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval{
		Items: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItems {
			if response.Items != nil {
				result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItems{
						Intervals: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervals {
							if items.Intervals != nil {
								result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervals, len(*items.Intervals))
								for i, intervals := range *items.Intervals {
									result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervals{
										ByInterface: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsByInterface {
											if intervals.ByInterface != nil {
												result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsByInterface, len(*intervals.ByInterface))
												for i, byInterface := range *intervals.ByInterface {
													result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsByInterface{
														Name: types.StringValue(byInterface.Name),
														Usage: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsByInterfaceUsage {
															if byInterface.Usage != nil {
																return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsByInterfaceUsage{
																	Recv: func() types.Int64 {
																		if byInterface.Usage.Recv != nil {
																			return types.Int64Value(int64(*byInterface.Usage.Recv))
																		}
																		return types.Int64{}
																	}(),
																	Send: func() types.Int64 {
																		if byInterface.Usage.Send != nil {
																			return types.Int64Value(int64(*byInterface.Usage.Send))
																		}
																		return types.Int64{}
																	}(),
																	Total: func() types.Int64 {
																		if byInterface.Usage.Total != nil {
																			return types.Int64Value(int64(*byInterface.Usage.Total))
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
										EndTs: types.StringValue(intervals.EndTs),
										Overall: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsOverall {
											if intervals.Overall != nil {
												return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalItemsIntervalsOverall{
													Recv: func() types.Int64 {
														if intervals.Overall.Recv != nil {
															return types.Int64Value(int64(*intervals.Overall.Recv))
														}
														return types.Int64{}
													}(),
													Send: func() types.Int64 {
														if intervals.Overall.Send != nil {
															return types.Int64Value(int64(*intervals.Overall.Send))
														}
														return types.Int64{}
													}(),
													Total: func() types.Int64 {
														if intervals.Overall.Total != nil {
															return types.Int64Value(int64(*intervals.Overall.Total))
														}
														return types.Int64{}
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
		Meta: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMeta {
			if response.Meta != nil {
				return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMeta{
					Counts: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMetaCounts{
								Items: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalMetaCountsItems{
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
