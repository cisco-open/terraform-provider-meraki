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

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsSummaryTopNetworksByStatusDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSummaryTopNetworksByStatusDataSource{}
)

func NewOrganizationsSummaryTopNetworksByStatusDataSource() datasource.DataSource {
	return &OrganizationsSummaryTopNetworksByStatusDataSource{}
}

type OrganizationsSummaryTopNetworksByStatusDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSummaryTopNetworksByStatusDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSummaryTopNetworksByStatusDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_summary_top_networks_by_status"
}

func (d *OrganizationsSummaryTopNetworksByStatusDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 5000.`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationSummaryTopNetworksByStatus`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"clients": schema.SingleNestedAttribute{
							MarkdownDescription: `Network clients data`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"counts": schema.SingleNestedAttribute{
									MarkdownDescription: `Network client counts`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"total": schema.Int64Attribute{
											MarkdownDescription: `Total count of clients in network`,
											Computed:            true,
										},
									},
								},
								"usage": schema.SingleNestedAttribute{
									MarkdownDescription: `Network client usage data`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"downstream": schema.Float64Attribute{
											MarkdownDescription: `Total downstream usage in network, in KB`,
											Computed:            true,
										},
										"upstream": schema.Float64Attribute{
											MarkdownDescription: `Total upstream usage in network, in KB`,
											Computed:            true,
										},
									},
								},
							},
						},
						"devices": schema.SingleNestedAttribute{
							MarkdownDescription: `Network device information`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"by_product_type": schema.SetNestedAttribute{
									MarkdownDescription: `URLs by product type`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"product_type": schema.StringAttribute{
												MarkdownDescription: `Product type`,
												Computed:            true,
											},
											"url": schema.StringAttribute{
												MarkdownDescription: `URL to clients list for the relevant product type`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Network name`,
							Computed:            true,
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: `Network identifier`,
							Computed:            true,
						},
						"product_types": schema.ListAttribute{
							MarkdownDescription: `Product types in network`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"statuses": schema.SingleNestedAttribute{
							MarkdownDescription: `Network device statuses`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"by_product_type": schema.SetNestedAttribute{
									MarkdownDescription: `List of status counts by product type`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"counts": schema.SingleNestedAttribute{
												MarkdownDescription: `Counts of devices by status`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"alerting": schema.Int64Attribute{
														MarkdownDescription: `Count of alerting devices`,
														Computed:            true,
													},
													"dormant": schema.Int64Attribute{
														MarkdownDescription: `Count of dormant devices`,
														Computed:            true,
													},
													"offline": schema.Int64Attribute{
														MarkdownDescription: `Count of offline devices`,
														Computed:            true,
													},
													"online": schema.Int64Attribute{
														MarkdownDescription: `Count of online devices`,
														Computed:            true,
													},
												},
											},
											"product_type": schema.StringAttribute{
												MarkdownDescription: `Product type`,
												Computed:            true,
											},
										},
									},
								},
								"overall": schema.StringAttribute{
									MarkdownDescription: `Overall status of network`,
									Computed:            true,
								},
							},
						},
						"tags": schema.ListAttribute{
							MarkdownDescription: `Network tags`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"url": schema.StringAttribute{
							MarkdownDescription: `Network clients list URL`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsSummaryTopNetworksByStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSummaryTopNetworksByStatus OrganizationsSummaryTopNetworksByStatus
	diags := req.Config.Get(ctx, &organizationsSummaryTopNetworksByStatus)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSummaryTopNetworksByStatus")
		vvOrganizationID := organizationsSummaryTopNetworksByStatus.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationSummaryTopNetworksByStatusQueryParams{}

		queryParams1.PerPage = int(organizationsSummaryTopNetworksByStatus.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsSummaryTopNetworksByStatus.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsSummaryTopNetworksByStatus.EndingBefore.ValueString()

		response1, restyResp1, err := d.client.Organizations.GetOrganizationSummaryTopNetworksByStatus(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSummaryTopNetworksByStatus",
				err.Error(),
			)
			return
		}

		organizationsSummaryTopNetworksByStatus = ResponseOrganizationsGetOrganizationSummaryTopNetworksByStatusItemsToBody(organizationsSummaryTopNetworksByStatus, response1)
		diags = resp.State.Set(ctx, &organizationsSummaryTopNetworksByStatus)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSummaryTopNetworksByStatus struct {
	OrganizationID types.String                                                          `tfsdk:"organization_id"`
	PerPage        types.Int64                                                           `tfsdk:"per_page"`
	StartingAfter  types.String                                                          `tfsdk:"starting_after"`
	EndingBefore   types.String                                                          `tfsdk:"ending_before"`
	Items          *[]ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatus `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatus struct {
	Clients      *ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusClients  `tfsdk:"clients"`
	Devices      *ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusDevices  `tfsdk:"devices"`
	Name         types.String                                                                `tfsdk:"name"`
	NetworkID    types.String                                                                `tfsdk:"network_id"`
	ProductTypes types.List                                                                  `tfsdk:"product_types"`
	Statuses     *ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusStatuses `tfsdk:"statuses"`
	Tags         types.List                                                                  `tfsdk:"tags"`
	URL          types.String                                                                `tfsdk:"url"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusClients struct {
	Counts *ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusClientsCounts `tfsdk:"counts"`
	Usage  *ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusClientsUsage  `tfsdk:"usage"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusClientsCounts struct {
	Total types.Int64 `tfsdk:"total"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusClientsUsage struct {
	Downstream types.Float64 `tfsdk:"downstream"`
	Upstream   types.Float64 `tfsdk:"upstream"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusDevices struct {
	ByProductType *[]ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusDevicesByProductType `tfsdk:"by_product_type"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusDevicesByProductType struct {
	ProductType types.String `tfsdk:"product_type"`
	URL         types.String `tfsdk:"url"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusStatuses struct {
	ByProductType *[]ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusStatusesByProductType `tfsdk:"by_product_type"`
	Overall       types.String                                                                               `tfsdk:"overall"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusStatusesByProductType struct {
	Counts      *ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusStatusesByProductTypeCounts `tfsdk:"counts"`
	ProductType types.String                                                                                   `tfsdk:"product_type"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusStatusesByProductTypeCounts struct {
	Alerting types.Int64 `tfsdk:"alerting"`
	Dormant  types.Int64 `tfsdk:"dormant"`
	Offline  types.Int64 `tfsdk:"offline"`
	Online   types.Int64 `tfsdk:"online"`
}

// ToBody
func ResponseOrganizationsGetOrganizationSummaryTopNetworksByStatusItemsToBody(state OrganizationsSummaryTopNetworksByStatus, response *merakigosdk.ResponseOrganizationsGetOrganizationSummaryTopNetworksByStatus) OrganizationsSummaryTopNetworksByStatus {
	var items []ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatus
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatus{
			Clients: func() *ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusClients {
				if item.Clients != nil {
					return &ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusClients{
						Counts: func() *ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusClientsCounts {
							if item.Clients.Counts != nil {
								return &ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusClientsCounts{
									Total: func() types.Int64 {
										if item.Clients.Counts.Total != nil {
											return types.Int64Value(int64(*item.Clients.Counts.Total))
										}
										return types.Int64{}
									}(),
								}
							}
							return &ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusClientsCounts{}
						}(),
						Usage: func() *ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusClientsUsage {
							if item.Clients.Usage != nil {
								return &ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusClientsUsage{
									Downstream: func() types.Float64 {
										if item.Clients.Usage.Downstream != nil {
											return types.Float64Value(float64(*item.Clients.Usage.Downstream))
										}
										return types.Float64{}
									}(),
									Upstream: func() types.Float64 {
										if item.Clients.Usage.Upstream != nil {
											return types.Float64Value(float64(*item.Clients.Usage.Upstream))
										}
										return types.Float64{}
									}(),
								}
							}
							return &ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusClientsUsage{}
						}(),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusClients{}
			}(),
			Devices: func() *ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusDevices {
				if item.Devices != nil {
					return &ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusDevices{
						ByProductType: func() *[]ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusDevicesByProductType {
							if item.Devices.ByProductType != nil {
								result := make([]ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusDevicesByProductType, len(*item.Devices.ByProductType))
								for i, byProductType := range *item.Devices.ByProductType {
									result[i] = ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusDevicesByProductType{
										ProductType: types.StringValue(byProductType.ProductType),
										URL:         types.StringValue(byProductType.URL),
									}
								}
								return &result
							}
							return &[]ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusDevicesByProductType{}
						}(),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusDevices{}
			}(),
			Name:         types.StringValue(item.Name),
			NetworkID:    types.StringValue(item.NetworkID),
			ProductTypes: StringSliceToList(item.ProductTypes),
			Statuses: func() *ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusStatuses {
				if item.Statuses != nil {
					return &ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusStatuses{
						ByProductType: func() *[]ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusStatusesByProductType {
							if item.Statuses.ByProductType != nil {
								result := make([]ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusStatusesByProductType, len(*item.Statuses.ByProductType))
								for i, byProductType := range *item.Statuses.ByProductType {
									result[i] = ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusStatusesByProductType{
										Counts: func() *ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusStatusesByProductTypeCounts {
											if byProductType.Counts != nil {
												return &ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusStatusesByProductTypeCounts{
													Alerting: func() types.Int64 {
														if byProductType.Counts.Alerting != nil {
															return types.Int64Value(int64(*byProductType.Counts.Alerting))
														}
														return types.Int64{}
													}(),
													Dormant: func() types.Int64 {
														if byProductType.Counts.Dormant != nil {
															return types.Int64Value(int64(*byProductType.Counts.Dormant))
														}
														return types.Int64{}
													}(),
													Offline: func() types.Int64 {
														if byProductType.Counts.Offline != nil {
															return types.Int64Value(int64(*byProductType.Counts.Offline))
														}
														return types.Int64{}
													}(),
													Online: func() types.Int64 {
														if byProductType.Counts.Online != nil {
															return types.Int64Value(int64(*byProductType.Counts.Online))
														}
														return types.Int64{}
													}(),
												}
											}
											return &ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusStatusesByProductTypeCounts{}
										}(),
										ProductType: types.StringValue(byProductType.ProductType),
									}
								}
								return &result
							}
							return &[]ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusStatusesByProductType{}
						}(),
						Overall: types.StringValue(item.Statuses.Overall),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationSummaryTopNetworksByStatusStatuses{}
			}(),
			Tags: StringSliceToList(item.Tags),
			URL:  types.StringValue(item.URL),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
