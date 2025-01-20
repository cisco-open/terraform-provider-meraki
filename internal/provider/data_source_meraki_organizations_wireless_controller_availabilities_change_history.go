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
	_ datasource.DataSource              = &OrganizationsWirelessControllerAvailabilitiesChangeHistoryDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessControllerAvailabilitiesChangeHistoryDataSource{}
)

func NewOrganizationsWirelessControllerAvailabilitiesChangeHistoryDataSource() datasource.DataSource {
	return &OrganizationsWirelessControllerAvailabilitiesChangeHistoryDataSource{}
}

type OrganizationsWirelessControllerAvailabilitiesChangeHistoryDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessControllerAvailabilitiesChangeHistoryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessControllerAvailabilitiesChangeHistoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_controller_availabilities_change_history"
}

func (d *OrganizationsWirelessControllerAvailabilitiesChangeHistoryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
						MarkdownDescription: `Wireless LAN controller connectivity information`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"changes": schema.SetNestedAttribute{
									MarkdownDescription: `Connectivity information of a wireless LAN controller`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"end_ts": schema.StringAttribute{
												MarkdownDescription: `The end time(UTC seconds) of the wireless LAN controller connectivity status change. This attribute is set to be null by default if there's no need to assign.`,
												Computed:            true,
											},
											"start_ts": schema.StringAttribute{
												MarkdownDescription: `The start time(UTC seconds) of the wireless LAN controller connectivity status change`,
												Computed:            true,
											},
											"status": schema.StringAttribute{
												MarkdownDescription: `The wireless LAN controller connectivity status`,
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

func (d *OrganizationsWirelessControllerAvailabilitiesChangeHistoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessControllerAvailabilitiesChangeHistory OrganizationsWirelessControllerAvailabilitiesChangeHistory
	diags := req.Config.Get(ctx, &organizationsWirelessControllerAvailabilitiesChangeHistory)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessControllerAvailabilitiesChangeHistory")
		vvOrganizationID := organizationsWirelessControllerAvailabilitiesChangeHistory.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessControllerAvailabilitiesChangeHistoryQueryParams{}

		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessControllerAvailabilitiesChangeHistory.Serials)
		queryParams1.T0 = organizationsWirelessControllerAvailabilitiesChangeHistory.T0.ValueString()
		queryParams1.T1 = organizationsWirelessControllerAvailabilitiesChangeHistory.T1.ValueString()
		queryParams1.Timespan = organizationsWirelessControllerAvailabilitiesChangeHistory.Timespan.ValueFloat64()
		queryParams1.PerPage = int(organizationsWirelessControllerAvailabilitiesChangeHistory.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessControllerAvailabilitiesChangeHistory.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessControllerAvailabilitiesChangeHistory.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.WirelessController.GetOrganizationWirelessControllerAvailabilitiesChangeHistory(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessControllerAvailabilitiesChangeHistory",
				err.Error(),
			)
			return
		}

		organizationsWirelessControllerAvailabilitiesChangeHistory = ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryItemToBody(organizationsWirelessControllerAvailabilitiesChangeHistory, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessControllerAvailabilitiesChangeHistory)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessControllerAvailabilitiesChangeHistory struct {
	OrganizationID types.String                                                                            `tfsdk:"organization_id"`
	Serials        types.List                                                                              `tfsdk:"serials"`
	T0             types.String                                                                            `tfsdk:"t0"`
	T1             types.String                                                                            `tfsdk:"t1"`
	Timespan       types.Float64                                                                           `tfsdk:"timespan"`
	PerPage        types.Int64                                                                             `tfsdk:"per_page"`
	StartingAfter  types.String                                                                            `tfsdk:"starting_after"`
	EndingBefore   types.String                                                                            `tfsdk:"ending_before"`
	Item           *ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistory `tfsdk:"item"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistory struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryItems `tfsdk:"items"`
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMeta    `tfsdk:"meta"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryItems struct {
	Changes *[]ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryItemsChanges `tfsdk:"changes"`
	Serial  types.String                                                                                          `tfsdk:"serial"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryItemsChanges struct {
	EndTs   types.String `tfsdk:"end_ts"`
	StartTs types.String `tfsdk:"start_ts"`
	Status  types.String `tfsdk:"status"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMetaCounts `tfsdk:"counts"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMetaCountsItems `tfsdk:"items"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryItemToBody(state OrganizationsWirelessControllerAvailabilitiesChangeHistory, response *merakigosdk.ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistory) OrganizationsWirelessControllerAvailabilitiesChangeHistory {
	itemState := ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistory{
		Items: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryItems {
			if response.Items != nil {
				result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryItems{
						Changes: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryItemsChanges {
							if items.Changes != nil {
								result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryItemsChanges, len(*items.Changes))
								for i, changes := range *items.Changes {
									result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryItemsChanges{
										EndTs:   types.StringValue(changes.EndTs),
										StartTs: types.StringValue(changes.StartTs),
										Status:  types.StringValue(changes.Status),
									}
								}
								return &result
							}
							return &[]ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryItemsChanges{}
						}(),
						Serial: types.StringValue(items.Serial),
					}
				}
				return &result
			}
			return &[]ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryItems{}
		}(),
		Meta: func() *ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMeta {
			if response.Meta != nil {
				return &ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMeta{
					Counts: func() *ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMetaCounts{
								Items: func() *ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMetaCountsItems{
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
									return &ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMetaCountsItems{}
								}(),
							}
						}
						return &ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMetaCounts{}
					}(),
				}
			}
			return &ResponseWirelessControllerGetOrganizationWirelessControllerAvailabilitiesChangeHistoryMeta{}
		}(),
	}
	state.Item = &itemState
	return state
}
