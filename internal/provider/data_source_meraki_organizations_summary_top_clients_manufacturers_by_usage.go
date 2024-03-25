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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsSummaryTopClientsManufacturersByUsageDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSummaryTopClientsManufacturersByUsageDataSource{}
)

func NewOrganizationsSummaryTopClientsManufacturersByUsageDataSource() datasource.DataSource {
	return &OrganizationsSummaryTopClientsManufacturersByUsageDataSource{}
}

type OrganizationsSummaryTopClientsManufacturersByUsageDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSummaryTopClientsManufacturersByUsageDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSummaryTopClientsManufacturersByUsageDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_summary_top_clients_manufacturers_by_usage"
}

func (d *OrganizationsSummaryTopClientsManufacturersByUsageDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 31 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsage`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"clients": schema.SingleNestedAttribute{
							MarkdownDescription: `Clients info`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"counts": schema.SingleNestedAttribute{
									MarkdownDescription: `Counts of clients`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"total": schema.Int64Attribute{
											MarkdownDescription: `Total counts of clients`,
											Computed:            true,
										},
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the manufacturer`,
							Computed:            true,
						},
						"usage": schema.SingleNestedAttribute{
							MarkdownDescription: `Clients usage`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"downstream": schema.Float64Attribute{
									MarkdownDescription: `Downstream data usage by client`,
									Computed:            true,
								},
								"total": schema.Float64Attribute{
									MarkdownDescription: `Total data usage by client`,
									Computed:            true,
								},
								"upstream": schema.Float64Attribute{
									MarkdownDescription: `Upstream data usage by client`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsSummaryTopClientsManufacturersByUsageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSummaryTopClientsManufacturersByUsage OrganizationsSummaryTopClientsManufacturersByUsage
	diags := req.Config.Get(ctx, &organizationsSummaryTopClientsManufacturersByUsage)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSummaryTopClientsManufacturersByUsage")
		vvOrganizationID := organizationsSummaryTopClientsManufacturersByUsage.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationSummaryTopClientsManufacturersByUsageQueryParams{}

		queryParams1.T0 = organizationsSummaryTopClientsManufacturersByUsage.T0.ValueString()
		queryParams1.T1 = organizationsSummaryTopClientsManufacturersByUsage.T1.ValueString()
		queryParams1.Timespan = organizationsSummaryTopClientsManufacturersByUsage.Timespan.ValueFloat64()

		response1, restyResp1, err := d.client.Organizations.GetOrganizationSummaryTopClientsManufacturersByUsage(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSummaryTopClientsManufacturersByUsage",
				err.Error(),
			)
			return
		}

		organizationsSummaryTopClientsManufacturersByUsage = ResponseOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageItemsToBody(organizationsSummaryTopClientsManufacturersByUsage, response1)
		diags = resp.State.Set(ctx, &organizationsSummaryTopClientsManufacturersByUsage)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSummaryTopClientsManufacturersByUsage struct {
	OrganizationID types.String                                                                     `tfsdk:"organization_id"`
	T0             types.String                                                                     `tfsdk:"t0"`
	T1             types.String                                                                     `tfsdk:"t1"`
	Timespan       types.Float64                                                                    `tfsdk:"timespan"`
	Items          *[]ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsage `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsage struct {
	Clients *ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageClients `tfsdk:"clients"`
	Name    types.String                                                                          `tfsdk:"name"`
	Usage   *ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageUsage   `tfsdk:"usage"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageClients struct {
	Counts *ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageClientsCounts `tfsdk:"counts"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageClientsCounts struct {
	Total types.Int64 `tfsdk:"total"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageUsage struct {
	Downstream types.Float64 `tfsdk:"downstream"`
	Total      types.Float64 `tfsdk:"total"`
	Upstream   types.Float64 `tfsdk:"upstream"`
}

// ToBody
func ResponseOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageItemsToBody(state OrganizationsSummaryTopClientsManufacturersByUsage, response *merakigosdk.ResponseOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsage) OrganizationsSummaryTopClientsManufacturersByUsage {
	var items []ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsage
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsage{
			Clients: func() *ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageClients {
				if item.Clients != nil {
					return &ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageClients{
						Counts: func() *ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageClientsCounts {
							if item.Clients.Counts != nil {
								return &ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageClientsCounts{
									Total: func() types.Int64 {
										if item.Clients.Counts.Total != nil {
											return types.Int64Value(int64(*item.Clients.Counts.Total))
										}
										return types.Int64{}
									}(),
								}
							}
							return &ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageClientsCounts{}
						}(),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageClients{}
			}(),
			Name: types.StringValue(item.Name),
			Usage: func() *ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageUsage {
				if item.Usage != nil {
					return &ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageUsage{
						Downstream: func() types.Float64 {
							if item.Usage.Downstream != nil {
								return types.Float64Value(float64(*item.Usage.Downstream))
							}
							return types.Float64{}
						}(),
						Total: func() types.Float64 {
							if item.Usage.Total != nil {
								return types.Float64Value(float64(*item.Usage.Total))
							}
							return types.Float64{}
						}(),
						Upstream: func() types.Float64 {
							if item.Usage.Upstream != nil {
								return types.Float64Value(float64(*item.Usage.Upstream))
							}
							return types.Float64{}
						}(),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationSummaryTopClientsManufacturersByUsageUsage{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
