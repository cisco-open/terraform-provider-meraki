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
	_ datasource.DataSource              = &OrganizationsSummaryTopAppliancesByUtilizationDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSummaryTopAppliancesByUtilizationDataSource{}
)

func NewOrganizationsSummaryTopAppliancesByUtilizationDataSource() datasource.DataSource {
	return &OrganizationsSummaryTopAppliancesByUtilizationDataSource{}
}

type OrganizationsSummaryTopAppliancesByUtilizationDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSummaryTopAppliancesByUtilizationDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSummaryTopAppliancesByUtilizationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_summary_top_appliances_by_utilization"
}

func (d *OrganizationsSummaryTopAppliancesByUtilizationDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationSummaryTopAppliancesByUtilization`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"mac": schema.StringAttribute{
							MarkdownDescription: `Mac address of the appliance`,
							Computed:            true,
						},
						"model": schema.StringAttribute{
							MarkdownDescription: `Model of the appliance`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the appliance`,
							Computed:            true,
						},
						"network": schema.SingleNestedAttribute{
							MarkdownDescription: `Network info`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `Network id`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Network name`,
									Computed:            true,
								},
							},
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `Serial number of the appliance`,
							Computed:            true,
						},
						"utilization": schema.SingleNestedAttribute{
							MarkdownDescription: `Utilization of the appliance`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"average": schema.SingleNestedAttribute{
									MarkdownDescription: `Average utilization of the appliance`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"percentage": schema.Float64Attribute{
											MarkdownDescription: `Average percentage utilization of the appliance`,
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
	}
}

func (d *OrganizationsSummaryTopAppliancesByUtilizationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSummaryTopAppliancesByUtilization OrganizationsSummaryTopAppliancesByUtilization
	diags := req.Config.Get(ctx, &organizationsSummaryTopAppliancesByUtilization)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSummaryTopAppliancesByUtilization")
		vvOrganizationID := organizationsSummaryTopAppliancesByUtilization.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationSummaryTopAppliancesByUtilizationQueryParams{}

		queryParams1.T0 = organizationsSummaryTopAppliancesByUtilization.T0.ValueString()
		queryParams1.T1 = organizationsSummaryTopAppliancesByUtilization.T1.ValueString()
		queryParams1.Timespan = organizationsSummaryTopAppliancesByUtilization.Timespan.ValueFloat64()

		response1, restyResp1, err := d.client.Organizations.GetOrganizationSummaryTopAppliancesByUtilization(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSummaryTopAppliancesByUtilization",
				err.Error(),
			)
			return
		}

		organizationsSummaryTopAppliancesByUtilization = ResponseOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationItemsToBody(organizationsSummaryTopAppliancesByUtilization, response1)
		diags = resp.State.Set(ctx, &organizationsSummaryTopAppliancesByUtilization)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSummaryTopAppliancesByUtilization struct {
	OrganizationID types.String                                                                 `tfsdk:"organization_id"`
	T0             types.String                                                                 `tfsdk:"t0"`
	T1             types.String                                                                 `tfsdk:"t1"`
	Timespan       types.Float64                                                                `tfsdk:"timespan"`
	Items          *[]ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilization `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilization struct {
	Mac         types.String                                                                          `tfsdk:"mac"`
	Model       types.String                                                                          `tfsdk:"model"`
	Name        types.String                                                                          `tfsdk:"name"`
	Network     *ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationNetwork     `tfsdk:"network"`
	Serial      types.String                                                                          `tfsdk:"serial"`
	Utilization *ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationUtilization `tfsdk:"utilization"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationUtilization struct {
	Average *ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationUtilizationAverage `tfsdk:"average"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationUtilizationAverage struct {
	Percentage types.Float64 `tfsdk:"percentage"`
}

// ToBody
func ResponseOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationItemsToBody(state OrganizationsSummaryTopAppliancesByUtilization, response *merakigosdk.ResponseOrganizationsGetOrganizationSummaryTopAppliancesByUtilization) OrganizationsSummaryTopAppliancesByUtilization {
	var items []ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilization
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilization{
			Mac:   types.StringValue(item.Mac),
			Model: types.StringValue(item.Model),
			Name:  types.StringValue(item.Name),
			Network: func() *ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationNetwork {
				if item.Network != nil {
					return &ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationNetwork{
						ID:   types.StringValue(item.Network.ID),
						Name: types.StringValue(item.Network.Name),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationNetwork{}
			}(),
			Serial: types.StringValue(item.Serial),
			Utilization: func() *ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationUtilization {
				if item.Utilization != nil {
					return &ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationUtilization{
						Average: func() *ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationUtilizationAverage {
							if item.Utilization.Average != nil {
								return &ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationUtilizationAverage{
									Percentage: func() types.Float64 {
										if item.Utilization.Average.Percentage != nil {
											return types.Float64Value(float64(*item.Utilization.Average.Percentage))
										}
										return types.Float64{}
									}(),
								}
							}
							return &ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationUtilizationAverage{}
						}(),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationSummaryTopAppliancesByUtilizationUtilization{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
