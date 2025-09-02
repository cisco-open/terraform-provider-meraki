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
	_ datasource.DataSource              = &OrganizationsWirelessControllerDevicesRedundancyStatusesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessControllerDevicesRedundancyStatusesDataSource{}
)

func NewOrganizationsWirelessControllerDevicesRedundancyStatusesDataSource() datasource.DataSource {
	return &OrganizationsWirelessControllerDevicesRedundancyStatusesDataSource{}
}

type OrganizationsWirelessControllerDevicesRedundancyStatusesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessControllerDevicesRedundancyStatusesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessControllerDevicesRedundancyStatusesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_controller_devices_redundancy_statuses"
}

func (d *OrganizationsWirelessControllerDevicesRedundancyStatusesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				MarkdownDescription: `serials query parameter. Optional parameter to filter wireless LAN controller by its cloud IDs. This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"items": schema.ListNestedAttribute{
						MarkdownDescription: `Wireless LAN controller redundancy statuses`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"enabled": schema.BoolAttribute{
									MarkdownDescription: `Wireless LAN controller redundancy enablement`,
									Computed:            true,
								},
								"failover": schema.SingleNestedAttribute{
									MarkdownDescription: `Wireless LAN controller failover information`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"counts": schema.SingleNestedAttribute{
											MarkdownDescription: `Wireless LAN controller switchover counts`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"total": schema.Int64Attribute{
													MarkdownDescription: `Total number of failovers`,
													Computed:            true,
												},
											},
										},
										"last": schema.SingleNestedAttribute{
											MarkdownDescription: `Wireless LAN controller last failover information`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"reason": schema.StringAttribute{
													MarkdownDescription: `Wireless LAN controller last redundancy switchover reason`,
													Computed:            true,
												},
												"ts": schema.StringAttribute{
													MarkdownDescription: `Wireless LAN controller last redundancy switchover time`,
													Computed:            true,
												},
											},
										},
									},
								},
								"mobility_mac": schema.StringAttribute{
									MarkdownDescription: `Wireless LAN controller redundancy mobility mac `,
									Computed:            true,
								},
								"mode": schema.StringAttribute{
									MarkdownDescription: `Wireless LAN controller redundancy SSO (stateful switchover)`,
									Computed:            true,
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

func (d *OrganizationsWirelessControllerDevicesRedundancyStatusesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessControllerDevicesRedundancyStatuses OrganizationsWirelessControllerDevicesRedundancyStatuses
	diags := req.Config.Get(ctx, &organizationsWirelessControllerDevicesRedundancyStatuses)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessControllerDevicesRedundancyStatuses")
		vvOrganizationID := organizationsWirelessControllerDevicesRedundancyStatuses.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessControllerDevicesRedundancyStatusesQueryParams{}

		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessControllerDevicesRedundancyStatuses.Serials)
		queryParams1.PerPage = int(organizationsWirelessControllerDevicesRedundancyStatuses.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessControllerDevicesRedundancyStatuses.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessControllerDevicesRedundancyStatuses.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.WirelessController.GetOrganizationWirelessControllerDevicesRedundancyStatuses(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessControllerDevicesRedundancyStatuses",
				err.Error(),
			)
			return
		}

		organizationsWirelessControllerDevicesRedundancyStatuses = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItemToBody(organizationsWirelessControllerDevicesRedundancyStatuses, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessControllerDevicesRedundancyStatuses)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessControllerDevicesRedundancyStatuses struct {
	OrganizationID types.String                                                                          `tfsdk:"organization_id"`
	Serials        types.List                                                                            `tfsdk:"serials"`
	PerPage        types.Int64                                                                           `tfsdk:"per_page"`
	StartingAfter  types.String                                                                          `tfsdk:"starting_after"`
	EndingBefore   types.String                                                                          `tfsdk:"ending_before"`
	Item           *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatuses `tfsdk:"item"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatuses struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItems `tfsdk:"items"`
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesMeta    `tfsdk:"meta"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItems struct {
	Enabled     types.Bool                                                                                         `tfsdk:"enabled"`
	Failover    *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailover `tfsdk:"failover"`
	MobilityMac types.String                                                                                       `tfsdk:"mobility_mac"`
	Mode        types.String                                                                                       `tfsdk:"mode"`
	Serial      types.String                                                                                       `tfsdk:"serial"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailover struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailoverCounts `tfsdk:"counts"`
	Last   *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailoverLast   `tfsdk:"last"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailoverCounts struct {
	Total types.Int64 `tfsdk:"total"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailoverLast struct {
	Reason types.String `tfsdk:"reason"`
	Ts     types.String `tfsdk:"ts"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesMetaCounts `tfsdk:"counts"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesMetaCountsItems `tfsdk:"items"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItemToBody(state OrganizationsWirelessControllerDevicesRedundancyStatuses, response *merakigosdk.ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatuses) OrganizationsWirelessControllerDevicesRedundancyStatuses {
	itemState := ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatuses{
		Items: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItems {
			if response.Items != nil {
				result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItems{
						Enabled: func() types.Bool {
							if items.Enabled != nil {
								return types.BoolValue(*items.Enabled)
							}
							return types.Bool{}
						}(),
						Failover: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailover {
							if items.Failover != nil {
								return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailover{
									Counts: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailoverCounts {
										if items.Failover.Counts != nil {
											return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailoverCounts{
												Total: func() types.Int64 {
													if items.Failover.Counts.Total != nil {
														return types.Int64Value(int64(*items.Failover.Counts.Total))
													}
													return types.Int64{}
												}(),
											}
										}
										return nil
									}(),
									Last: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailoverLast {
										if items.Failover.Last != nil {
											return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesItemsFailoverLast{
												Reason: func() types.String {
													if items.Failover.Last.Reason != "" {
														return types.StringValue(items.Failover.Last.Reason)
													}
													return types.String{}
												}(),
												Ts: func() types.String {
													if items.Failover.Last.Ts != "" {
														return types.StringValue(items.Failover.Last.Ts)
													}
													return types.String{}
												}(),
											}
										}
										return nil
									}(),
								}
							}
							return nil
						}(),
						MobilityMac: func() types.String {
							if items.MobilityMac != "" {
								return types.StringValue(items.MobilityMac)
							}
							return types.String{}
						}(),
						Mode: func() types.String {
							if items.Mode != "" {
								return types.StringValue(items.Mode)
							}
							return types.String{}
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
		Meta: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesMeta {
			if response.Meta != nil {
				return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesMeta{
					Counts: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesMetaCounts{
								Items: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyStatusesMetaCountsItems{
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
