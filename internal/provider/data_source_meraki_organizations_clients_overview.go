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
	_ datasource.DataSource              = &OrganizationsClientsOverviewDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsClientsOverviewDataSource{}
)

func NewOrganizationsClientsOverviewDataSource() datasource.DataSource {
	return &OrganizationsClientsOverviewDataSource{}
}

type OrganizationsClientsOverviewDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsClientsOverviewDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsClientsOverviewDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_clients_overview"
}

func (d *OrganizationsClientsOverviewDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"counts": schema.SingleNestedAttribute{
						MarkdownDescription: `Client count information`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"total": schema.Int64Attribute{
								MarkdownDescription: `Total number of clients with data usage in organization`,
								Computed:            true,
							},
						},
					},
					"usage": schema.SingleNestedAttribute{
						MarkdownDescription: `Usage information of all clients across organization`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"average": schema.Float64Attribute{
								MarkdownDescription: `Average data usage (in kb) of each client in organization`,
								Computed:            true,
							},
							"overall": schema.SingleNestedAttribute{
								MarkdownDescription: `Overall data usage of all clients across organization`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"downstream": schema.Float64Attribute{
										MarkdownDescription: `Downstream data usage (in kb) of all clients across organization`,
										Computed:            true,
									},
									"total": schema.Float64Attribute{
										MarkdownDescription: `Total data usage (in kb) of all clients across organization`,
										Computed:            true,
									},
									"upstream": schema.Float64Attribute{
										MarkdownDescription: `Upstream data usage (in kb) of all clients across organization`,
										Computed:            true,
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

func (d *OrganizationsClientsOverviewDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsClientsOverview OrganizationsClientsOverview
	diags := req.Config.Get(ctx, &organizationsClientsOverview)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationClientsOverview")
		vvOrganizationID := organizationsClientsOverview.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationClientsOverviewQueryParams{}

		queryParams1.T0 = organizationsClientsOverview.T0.ValueString()
		queryParams1.T1 = organizationsClientsOverview.T1.ValueString()
		queryParams1.Timespan = organizationsClientsOverview.Timespan.ValueFloat64()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationClientsOverview(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationClientsOverview",
				err.Error(),
			)
			return
		}

		organizationsClientsOverview = ResponseOrganizationsGetOrganizationClientsOverviewItemToBody(organizationsClientsOverview, response1)
		diags = resp.State.Set(ctx, &organizationsClientsOverview)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsClientsOverview struct {
	OrganizationID types.String                                         `tfsdk:"organization_id"`
	T0             types.String                                         `tfsdk:"t0"`
	T1             types.String                                         `tfsdk:"t1"`
	Timespan       types.Float64                                        `tfsdk:"timespan"`
	Item           *ResponseOrganizationsGetOrganizationClientsOverview `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationClientsOverview struct {
	Counts *ResponseOrganizationsGetOrganizationClientsOverviewCounts `tfsdk:"counts"`
	Usage  *ResponseOrganizationsGetOrganizationClientsOverviewUsage  `tfsdk:"usage"`
}

type ResponseOrganizationsGetOrganizationClientsOverviewCounts struct {
	Total types.Int64 `tfsdk:"total"`
}

type ResponseOrganizationsGetOrganizationClientsOverviewUsage struct {
	Average types.Float64                                                    `tfsdk:"average"`
	Overall *ResponseOrganizationsGetOrganizationClientsOverviewUsageOverall `tfsdk:"overall"`
}

type ResponseOrganizationsGetOrganizationClientsOverviewUsageOverall struct {
	Downstream types.Float64 `tfsdk:"downstream"`
	Total      types.Float64 `tfsdk:"total"`
	Upstream   types.Float64 `tfsdk:"upstream"`
}

// ToBody
func ResponseOrganizationsGetOrganizationClientsOverviewItemToBody(state OrganizationsClientsOverview, response *merakigosdk.ResponseOrganizationsGetOrganizationClientsOverview) OrganizationsClientsOverview {
	itemState := ResponseOrganizationsGetOrganizationClientsOverview{
		Counts: func() *ResponseOrganizationsGetOrganizationClientsOverviewCounts {
			if response.Counts != nil {
				return &ResponseOrganizationsGetOrganizationClientsOverviewCounts{
					Total: func() types.Int64 {
						if response.Counts.Total != nil {
							return types.Int64Value(int64(*response.Counts.Total))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
		Usage: func() *ResponseOrganizationsGetOrganizationClientsOverviewUsage {
			if response.Usage != nil {
				return &ResponseOrganizationsGetOrganizationClientsOverviewUsage{
					Average: func() types.Float64 {
						if response.Usage.Average != nil {
							return types.Float64Value(float64(*response.Usage.Average))
						}
						return types.Float64{}
					}(),
					Overall: func() *ResponseOrganizationsGetOrganizationClientsOverviewUsageOverall {
						if response.Usage.Overall != nil {
							return &ResponseOrganizationsGetOrganizationClientsOverviewUsageOverall{
								Downstream: func() types.Float64 {
									if response.Usage.Overall.Downstream != nil {
										return types.Float64Value(float64(*response.Usage.Overall.Downstream))
									}
									return types.Float64{}
								}(),
								Total: func() types.Float64 {
									if response.Usage.Overall.Total != nil {
										return types.Float64Value(float64(*response.Usage.Overall.Total))
									}
									return types.Float64{}
								}(),
								Upstream: func() types.Float64 {
									if response.Usage.Overall.Upstream != nil {
										return types.Float64Value(float64(*response.Usage.Overall.Upstream))
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
	state.Item = &itemState
	return state
}
