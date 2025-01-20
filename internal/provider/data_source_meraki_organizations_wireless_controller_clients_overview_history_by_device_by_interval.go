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
	_ datasource.DataSource              = &OrganizationsWirelessControllerClientsOverviewHistoryByDeviceByIntervalDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessControllerClientsOverviewHistoryByDeviceByIntervalDataSource{}
)

func NewOrganizationsWirelessControllerClientsOverviewHistoryByDeviceByIntervalDataSource() datasource.DataSource {
	return &OrganizationsWirelessControllerClientsOverviewHistoryByDeviceByIntervalDataSource{}
}

type OrganizationsWirelessControllerClientsOverviewHistoryByDeviceByIntervalDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessControllerClientsOverviewHistoryByDeviceByIntervalDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessControllerClientsOverviewHistoryByDeviceByIntervalDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_controller_clients_overview_history_by_device_by_interval"
}

func (d *OrganizationsWirelessControllerClientsOverviewHistoryByDeviceByIntervalDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter wireless LAN controllers by network ID. This filter uses multiple exact matches.`,
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
			"resolution": schema.Int64Attribute{
				MarkdownDescription: `resolution query parameter. The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400.`,
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
						MarkdownDescription: `Overview history of wireless LAN controllers`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"network": schema.SingleNestedAttribute{
									MarkdownDescription: `Wireless LAN controller network`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `Wireless LAN controller network ID`,
											Computed:            true,
										},
									},
								},
								"readings": schema.SetNestedAttribute{
									MarkdownDescription: `Overview history of a wireless LAN controller`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"counts": schema.SingleNestedAttribute{
												MarkdownDescription: `Client counts`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"by_status": schema.SingleNestedAttribute{
														MarkdownDescription: `Client counts by its status`,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"online": schema.Int64Attribute{
																MarkdownDescription: `Number of connected clients`,
																Computed:            true,
															},
														},
													},
												},
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
								"serial": schema.StringAttribute{
									MarkdownDescription: `Wireless LAN controller cloud ID`,
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

func (d *OrganizationsWirelessControllerClientsOverviewHistoryByDeviceByIntervalDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessControllerClientsOverviewHistoryByDeviceByInterval OrganizationsWirelessControllerClientsOverviewHistoryByDeviceByInterval
	diags := req.Config.Get(ctx, &organizationsWirelessControllerClientsOverviewHistoryByDeviceByInterval)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval")
		vvOrganizationID := organizationsWirelessControllerClientsOverviewHistoryByDeviceByInterval.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalQueryParams{}

		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsWirelessControllerClientsOverviewHistoryByDeviceByInterval.NetworkIDs)
		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessControllerClientsOverviewHistoryByDeviceByInterval.Serials)
		queryParams1.T0 = organizationsWirelessControllerClientsOverviewHistoryByDeviceByInterval.T0.ValueString()
		queryParams1.T1 = organizationsWirelessControllerClientsOverviewHistoryByDeviceByInterval.T1.ValueString()
		queryParams1.Timespan = organizationsWirelessControllerClientsOverviewHistoryByDeviceByInterval.Timespan.ValueFloat64()
		queryParams1.PerPage = int(organizationsWirelessControllerClientsOverviewHistoryByDeviceByInterval.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessControllerClientsOverviewHistoryByDeviceByInterval.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessControllerClientsOverviewHistoryByDeviceByInterval.EndingBefore.ValueString()
		queryParams1.Resolution = int(organizationsWirelessControllerClientsOverviewHistoryByDeviceByInterval.Resolution.ValueInt64())

		// has_unknown_response: None

		response1, restyResp1, err := d.client.WirelessController.GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval",
				err.Error(),
			)
			return
		}

		organizationsWirelessControllerClientsOverviewHistoryByDeviceByInterval = ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemToBody(organizationsWirelessControllerClientsOverviewHistoryByDeviceByInterval, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessControllerClientsOverviewHistoryByDeviceByInterval)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessControllerClientsOverviewHistoryByDeviceByInterval struct {
	OrganizationID types.String                                                                                         `tfsdk:"organization_id"`
	NetworkIDs     types.List                                                                                           `tfsdk:"network_ids"`
	Serials        types.List                                                                                           `tfsdk:"serials"`
	T0             types.String                                                                                         `tfsdk:"t0"`
	T1             types.String                                                                                         `tfsdk:"t1"`
	Timespan       types.Float64                                                                                        `tfsdk:"timespan"`
	PerPage        types.Int64                                                                                          `tfsdk:"per_page"`
	StartingAfter  types.String                                                                                         `tfsdk:"starting_after"`
	EndingBefore   types.String                                                                                         `tfsdk:"ending_before"`
	Resolution     types.Int64                                                                                          `tfsdk:"resolution"`
	Item           *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval `tfsdk:"item"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItems `tfsdk:"items"`
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMeta    `tfsdk:"meta"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItems struct {
	Network  *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsNetwork    `tfsdk:"network"`
	Readings *[]ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadings `tfsdk:"readings"`
	Serial   types.String                                                                                                        `tfsdk:"serial"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsNetwork struct {
	ID types.String `tfsdk:"id"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadings struct {
	Counts  *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadingsCounts `tfsdk:"counts"`
	EndTs   types.String                                                                                                            `tfsdk:"end_ts"`
	StartTs types.String                                                                                                            `tfsdk:"start_ts"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadingsCounts struct {
	ByStatus *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadingsCountsByStatus `tfsdk:"by_status"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadingsCountsByStatus struct {
	Online types.Int64 `tfsdk:"online"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMetaCounts `tfsdk:"counts"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMetaCountsItems `tfsdk:"items"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemToBody(state OrganizationsWirelessControllerClientsOverviewHistoryByDeviceByInterval, response *merakigosdk.ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval) OrganizationsWirelessControllerClientsOverviewHistoryByDeviceByInterval {
	itemState := ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval{
		Items: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItems {
			if response.Items != nil {
				result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItems{
						Network: func() *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsNetwork {
							if items.Network != nil {
								return &ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsNetwork{
									ID: types.StringValue(items.Network.ID),
								}
							}
							return &ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsNetwork{}
						}(),
						Readings: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadings {
							if items.Readings != nil {
								result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadings, len(*items.Readings))
								for i, readings := range *items.Readings {
									result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadings{
										Counts: func() *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadingsCounts {
											if readings.Counts != nil {
												return &ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadingsCounts{
													ByStatus: func() *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadingsCountsByStatus {
														if readings.Counts.ByStatus != nil {
															return &ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadingsCountsByStatus{
																Online: func() types.Int64 {
																	if readings.Counts.ByStatus.Online != nil {
																		return types.Int64Value(int64(*readings.Counts.ByStatus.Online))
																	}
																	return types.Int64{}
																}(),
															}
														}
														return &ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadingsCountsByStatus{}
													}(),
												}
											}
											return &ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadingsCounts{}
										}(),
										EndTs:   types.StringValue(readings.EndTs),
										StartTs: types.StringValue(readings.StartTs),
									}
								}
								return &result
							}
							return &[]ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItemsReadings{}
						}(),
						Serial: types.StringValue(items.Serial),
					}
				}
				return &result
			}
			return &[]ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalItems{}
		}(),
		Meta: func() *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMeta {
			if response.Meta != nil {
				return &ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMeta{
					Counts: func() *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMetaCounts{
								Items: func() *ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMetaCountsItems{
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
									return &ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMetaCountsItems{}
								}(),
							}
						}
						return &ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMetaCounts{}
					}(),
				}
			}
			return &ResponseWirelessControllerGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalMeta{}
		}(),
	}
	state.Item = &itemState
	return state
}
