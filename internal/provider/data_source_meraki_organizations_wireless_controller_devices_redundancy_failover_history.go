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
	_ datasource.DataSource              = &OrganizationsWirelessControllerDevicesRedundancyFailoverHistoryDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessControllerDevicesRedundancyFailoverHistoryDataSource{}
)

func NewOrganizationsWirelessControllerDevicesRedundancyFailoverHistoryDataSource() datasource.DataSource {
	return &OrganizationsWirelessControllerDevicesRedundancyFailoverHistoryDataSource{}
}

type OrganizationsWirelessControllerDevicesRedundancyFailoverHistoryDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessControllerDevicesRedundancyFailoverHistoryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessControllerDevicesRedundancyFailoverHistoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_controller_devices_redundancy_failover_history"
}

func (d *OrganizationsWirelessControllerDevicesRedundancyFailoverHistoryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"items": schema.ListNestedAttribute{
							MarkdownDescription: `Wireless LAN controller HA failover events`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"active": schema.SingleNestedAttribute{
										MarkdownDescription: `Details about the active unit`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"chassis": schema.SingleNestedAttribute{
												MarkdownDescription: `Details about the active unit chassis`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"name": schema.StringAttribute{
														MarkdownDescription: `The name of the active chassis unit`,
														Computed:            true,
													},
												},
											},
										},
									},
									"failed": schema.SingleNestedAttribute{
										MarkdownDescription: `Details about the failed unit`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"chassis": schema.SingleNestedAttribute{
												MarkdownDescription: `Details about the failed unit chassis`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"name": schema.StringAttribute{
														MarkdownDescription: `The name of the failed chassis unit`,
														Computed:            true,
													},
												},
											},
										},
									},
									"reason": schema.StringAttribute{
										MarkdownDescription: `Failover reason`,
										Computed:            true,
									},
									"serial": schema.StringAttribute{
										MarkdownDescription: `Wireless LAN controller cloud ID`,
										Computed:            true,
									},
									"ts": schema.StringAttribute{
										MarkdownDescription: `Failover time`,
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
		},
	}
}

func (d *OrganizationsWirelessControllerDevicesRedundancyFailoverHistoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessControllerDevicesRedundancyFailoverHistory OrganizationsWirelessControllerDevicesRedundancyFailoverHistory
	diags := req.Config.Get(ctx, &organizationsWirelessControllerDevicesRedundancyFailoverHistory)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessControllerDevicesRedundancyFailoverHistory")
		vvOrganizationID := organizationsWirelessControllerDevicesRedundancyFailoverHistory.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryQueryParams{}

		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessControllerDevicesRedundancyFailoverHistory.Serials)
		queryParams1.T0 = organizationsWirelessControllerDevicesRedundancyFailoverHistory.T0.ValueString()
		queryParams1.T1 = organizationsWirelessControllerDevicesRedundancyFailoverHistory.T1.ValueString()
		queryParams1.Timespan = organizationsWirelessControllerDevicesRedundancyFailoverHistory.Timespan.ValueFloat64()
		queryParams1.PerPage = int(organizationsWirelessControllerDevicesRedundancyFailoverHistory.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessControllerDevicesRedundancyFailoverHistory.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessControllerDevicesRedundancyFailoverHistory.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.WirelessController.GetOrganizationWirelessControllerDevicesRedundancyFailoverHistory(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessControllerDevicesRedundancyFailoverHistory",
				err.Error(),
			)
			return
		}

		organizationsWirelessControllerDevicesRedundancyFailoverHistory = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsToBody(organizationsWirelessControllerDevicesRedundancyFailoverHistory, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessControllerDevicesRedundancyFailoverHistory)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessControllerDevicesRedundancyFailoverHistory struct {
	OrganizationID types.String                                                                                       `tfsdk:"organization_id"`
	Serials        types.List                                                                                         `tfsdk:"serials"`
	T0             types.String                                                                                       `tfsdk:"t0"`
	T1             types.String                                                                                       `tfsdk:"t1"`
	Timespan       types.Float64                                                                                      `tfsdk:"timespan"`
	PerPage        types.Int64                                                                                        `tfsdk:"per_page"`
	StartingAfter  types.String                                                                                       `tfsdk:"starting_after"`
	EndingBefore   types.String                                                                                       `tfsdk:"ending_before"`
	Items          *[]ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory `tfsdk:"items"`
}

type ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory struct {
	Items *[]ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItems `tfsdk:"items"`
	Meta  *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMeta    `tfsdk:"meta"`
}

type ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItems struct {
	Active *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsActive `tfsdk:"active"`
	Failed *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsFailed `tfsdk:"failed"`
	Reason types.String                                                                                                `tfsdk:"reason"`
	Serial types.String                                                                                                `tfsdk:"serial"`
	Ts     types.String                                                                                                `tfsdk:"ts"`
}

type ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsActive struct {
	Chassis *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsActiveChassis `tfsdk:"chassis"`
}

type ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsActiveChassis struct {
	Name types.String `tfsdk:"name"`
}

type ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsFailed struct {
	Chassis *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsFailedChassis `tfsdk:"chassis"`
}

type ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsFailedChassis struct {
	Name types.String `tfsdk:"name"`
}

type ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMeta struct {
	Counts *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMetaCounts `tfsdk:"counts"`
}

type ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMetaCounts struct {
	Items *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMetaCountsItems `tfsdk:"items"`
}

type ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsToBody(state OrganizationsWirelessControllerDevicesRedundancyFailoverHistory, response *merakigosdk.ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory) OrganizationsWirelessControllerDevicesRedundancyFailoverHistory {
	var items []ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory
	for _, item := range *response {
		itemState := ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory{
			Items: func() *[]ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItems {
				if item.Items != nil {
					result := make([]ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItems, len(*item.Items))
					for i, items := range *item.Items {
						result[i] = ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItems{
							Active: func() *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsActive {
								if items.Active != nil {
									return &ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsActive{
										Chassis: func() *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsActiveChassis {
											if items.Active.Chassis != nil {
												return &ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsActiveChassis{
													Name: types.StringValue(items.Active.Chassis.Name),
												}
											}
											return nil
										}(),
									}
								}
								return nil
							}(),
							Failed: func() *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsFailed {
								if items.Failed != nil {
									return &ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsFailed{
										Chassis: func() *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsFailedChassis {
											if items.Failed.Chassis != nil {
												return &ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryItemsFailedChassis{
													Name: types.StringValue(items.Failed.Chassis.Name),
												}
											}
											return nil
										}(),
									}
								}
								return nil
							}(),
							Reason: types.StringValue(items.Reason),
							Serial: types.StringValue(items.Serial),
							Ts:     types.StringValue(items.Ts),
						}
					}
					return &result
				}
				return nil
			}(),
			Meta: func() *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMeta {
				if item.Meta != nil {
					return &ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMeta{
						Counts: func() *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMetaCounts {
							if item.Meta.Counts != nil {
								return &ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMetaCounts{
									Items: func() *ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMetaCountsItems {
										if item.Meta.Counts.Items != nil {
											return &ResponseItemWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryMetaCountsItems{
												Remaining: func() types.Int64 {
													if item.Meta.Counts.Items.Remaining != nil {
														return types.Int64Value(int64(*item.Meta.Counts.Items.Remaining))
													}
													return types.Int64{}
												}(),
												Total: func() types.Int64 {
													if item.Meta.Counts.Items.Total != nil {
														return types.Int64Value(int64(*item.Meta.Counts.Items.Total))
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
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
