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
	_ datasource.DataSource              = &OrganizationsWirelessClientsOverviewByDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessClientsOverviewByDeviceDataSource{}
)

func NewOrganizationsWirelessClientsOverviewByDeviceDataSource() datasource.DataSource {
	return &OrganizationsWirelessClientsOverviewByDeviceDataSource{}
}

type OrganizationsWirelessClientsOverviewByDeviceDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessClientsOverviewByDeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessClientsOverviewByDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_clients_overview_by_device"
}

func (d *OrganizationsWirelessClientsOverviewByDeviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"campus_gateway_cluster_ids": schema.ListAttribute{
				MarkdownDescription: `campusGatewayClusterIds query parameter. Optional parameter to filter access points client counts by MCG cluster IDs. This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter access points client counts by network ID. This filter uses multiple exact matches.`,
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
				MarkdownDescription: `serials query parameter. Optional parameter to filter access points client counts by its serial numbers. This filter uses multiple exact matches.`,
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
						MarkdownDescription: `Access point client count`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"counts": schema.SingleNestedAttribute{
									MarkdownDescription: `Associated client count on access point`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"by_status": schema.SingleNestedAttribute{
											MarkdownDescription: `Associated client count on access point by status`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"online": schema.Int64Attribute{
													MarkdownDescription: `Active client count`,
													Computed:            true,
												},
											},
										},
									},
								},
								"network": schema.SingleNestedAttribute{
									MarkdownDescription: `Access point network`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `Access point network ID`,
											Computed:            true,
										},
									},
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `Access point Serial number`,
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

func (d *OrganizationsWirelessClientsOverviewByDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessClientsOverviewByDevice OrganizationsWirelessClientsOverviewByDevice
	diags := req.Config.Get(ctx, &organizationsWirelessClientsOverviewByDevice)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessClientsOverviewByDevice")
		vvOrganizationID := organizationsWirelessClientsOverviewByDevice.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessClientsOverviewByDeviceQueryParams{}

		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsWirelessClientsOverviewByDevice.NetworkIDs)
		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessClientsOverviewByDevice.Serials)
		queryParams1.CampusGatewayClusterIDs = elementsToStrings(ctx, organizationsWirelessClientsOverviewByDevice.CampusGatewayClusterIDs)
		queryParams1.PerPage = int(organizationsWirelessClientsOverviewByDevice.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessClientsOverviewByDevice.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessClientsOverviewByDevice.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetOrganizationWirelessClientsOverviewByDevice(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessClientsOverviewByDevice",
				err.Error(),
			)
			return
		}

		organizationsWirelessClientsOverviewByDevice = ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceItemToBody(organizationsWirelessClientsOverviewByDevice, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessClientsOverviewByDevice)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessClientsOverviewByDevice struct {
	OrganizationID          types.String                                                    `tfsdk:"organization_id"`
	NetworkIDs              types.List                                                      `tfsdk:"network_ids"`
	Serials                 types.List                                                      `tfsdk:"serials"`
	CampusGatewayClusterIDs types.List                                                      `tfsdk:"campus_gateway_cluster_ids"`
	PerPage                 types.Int64                                                     `tfsdk:"per_page"`
	StartingAfter           types.String                                                    `tfsdk:"starting_after"`
	EndingBefore            types.String                                                    `tfsdk:"ending_before"`
	Item                    *ResponseWirelessGetOrganizationWirelessClientsOverviewByDevice `tfsdk:"item"`
}

type ResponseWirelessGetOrganizationWirelessClientsOverviewByDevice struct {
	Items *[]ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceItems `tfsdk:"items"`
	Meta  *ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceMeta    `tfsdk:"meta"`
}

type ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceItems struct {
	Counts  *ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceItemsCounts  `tfsdk:"counts"`
	Network *ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceItemsNetwork `tfsdk:"network"`
	Serial  types.String                                                                `tfsdk:"serial"`
}

type ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceItemsCounts struct {
	ByStatus *ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceItemsCountsByStatus `tfsdk:"by_status"`
}

type ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceItemsCountsByStatus struct {
	Online types.Int64 `tfsdk:"online"`
}

type ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceItemsNetwork struct {
	ID types.String `tfsdk:"id"`
}

type ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceMeta struct {
	Counts *ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceMetaCounts `tfsdk:"counts"`
}

type ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceMetaCounts struct {
	Items *ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceMetaCountsItems `tfsdk:"items"`
}

type ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceItemToBody(state OrganizationsWirelessClientsOverviewByDevice, response *merakigosdk.ResponseWirelessGetOrganizationWirelessClientsOverviewByDevice) OrganizationsWirelessClientsOverviewByDevice {
	itemState := ResponseWirelessGetOrganizationWirelessClientsOverviewByDevice{
		Items: func() *[]ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceItems {
			if response.Items != nil {
				result := make([]ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceItems{
						Counts: func() *ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceItemsCounts {
							if items.Counts != nil {
								return &ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceItemsCounts{
									ByStatus: func() *ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceItemsCountsByStatus {
										if items.Counts.ByStatus != nil {
											return &ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceItemsCountsByStatus{
												Online: func() types.Int64 {
													if items.Counts.ByStatus.Online != nil {
														return types.Int64Value(int64(*items.Counts.ByStatus.Online))
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
						Network: func() *ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceItemsNetwork {
							if items.Network != nil {
								return &ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceItemsNetwork{
									ID: func() types.String {
										if items.Network.ID != "" {
											return types.StringValue(items.Network.ID)
										}
										return types.String{}
									}(),
								}
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
		Meta: func() *ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceMeta {
			if response.Meta != nil {
				return &ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceMeta{
					Counts: func() *ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceMetaCounts{
								Items: func() *ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseWirelessGetOrganizationWirelessClientsOverviewByDeviceMetaCountsItems{
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
