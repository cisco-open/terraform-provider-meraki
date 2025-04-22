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
	_ datasource.DataSource              = &OrganizationsSummaryTopApplicationsByUsageDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSummaryTopApplicationsByUsageDataSource{}
)

func NewOrganizationsSummaryTopApplicationsByUsageDataSource() datasource.DataSource {
	return &OrganizationsSummaryTopApplicationsByUsageDataSource{}
}

type OrganizationsSummaryTopApplicationsByUsageDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSummaryTopApplicationsByUsageDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSummaryTopApplicationsByUsageDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_summary_top_applications_by_usage"
}

func (d *OrganizationsSummaryTopApplicationsByUsageDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: `device query parameter. Match result to an exact device tag`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId query parameter. Match result to an exact network id`,
				Optional:            true,
			},
			"network_tag": schema.StringAttribute{
				MarkdownDescription: `networkTag query parameter. Match result to an exact network tag`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"quantity": schema.Int64Attribute{
				MarkdownDescription: `quantity query parameter. Set number of desired results to return. Default is 10.`,
				Optional:            true,
			},
			"ssid_name": schema.StringAttribute{
				MarkdownDescription: `ssidName query parameter. Filter results by ssid name`,
				Optional:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 186 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 25 minutes and be less than or equal to 186 days. The default is 1 day.`,
				Optional:            true,
			},
			"usage_uplink": schema.StringAttribute{
				MarkdownDescription: `usageUplink query parameter. Filter results by usage uplink`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationSummaryTopApplicationsByUsage`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"application": schema.StringAttribute{
							MarkdownDescription: `Name of the Application`,
							Computed:            true,
						},
						"downstream": schema.Float64Attribute{
							MarkdownDescription: `Downstream usage of the Application, in megabytes`,
							Computed:            true,
						},
						"percentage": schema.Float64Attribute{
							MarkdownDescription: `Percent usage of the Application`,
							Computed:            true,
						},
						"total": schema.Float64Attribute{
							MarkdownDescription: `Total usage of the Application, in megabytes`,
							Computed:            true,
						},
						"upstream": schema.Float64Attribute{
							MarkdownDescription: `Upstream usage of the Application, in megabytes`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsSummaryTopApplicationsByUsageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSummaryTopApplicationsByUsage OrganizationsSummaryTopApplicationsByUsage
	diags := req.Config.Get(ctx, &organizationsSummaryTopApplicationsByUsage)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSummaryTopApplicationsByUsage")
		vvOrganizationID := organizationsSummaryTopApplicationsByUsage.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationSummaryTopApplicationsByUsageQueryParams{}

		queryParams1.NetworkTag = organizationsSummaryTopApplicationsByUsage.NetworkTag.ValueString()
		queryParams1.Device = organizationsSummaryTopApplicationsByUsage.Device.ValueString()
		queryParams1.NetworkID = organizationsSummaryTopApplicationsByUsage.NetworkID.ValueString()
		queryParams1.Quantity = int(organizationsSummaryTopApplicationsByUsage.Quantity.ValueInt64())
		queryParams1.SSIDName = organizationsSummaryTopApplicationsByUsage.SSIDName.ValueString()
		queryParams1.UsageUplink = organizationsSummaryTopApplicationsByUsage.UsageUplink.ValueString()
		queryParams1.T0 = organizationsSummaryTopApplicationsByUsage.T0.ValueString()
		queryParams1.T1 = organizationsSummaryTopApplicationsByUsage.T1.ValueString()
		queryParams1.Timespan = organizationsSummaryTopApplicationsByUsage.Timespan.ValueFloat64()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationSummaryTopApplicationsByUsage(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSummaryTopApplicationsByUsage",
				err.Error(),
			)
			return
		}

		organizationsSummaryTopApplicationsByUsage = ResponseOrganizationsGetOrganizationSummaryTopApplicationsByUsageItemsToBody(organizationsSummaryTopApplicationsByUsage, response1)
		diags = resp.State.Set(ctx, &organizationsSummaryTopApplicationsByUsage)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSummaryTopApplicationsByUsage struct {
	OrganizationID types.String                                                             `tfsdk:"organization_id"`
	NetworkTag     types.String                                                             `tfsdk:"network_tag"`
	Device         types.String                                                             `tfsdk:"device"`
	NetworkID      types.String                                                             `tfsdk:"network_id"`
	Quantity       types.Int64                                                              `tfsdk:"quantity"`
	SSIDName       types.String                                                             `tfsdk:"ssid_name"`
	UsageUplink    types.String                                                             `tfsdk:"usage_uplink"`
	T0             types.String                                                             `tfsdk:"t0"`
	T1             types.String                                                             `tfsdk:"t1"`
	Timespan       types.Float64                                                            `tfsdk:"timespan"`
	Items          *[]ResponseItemOrganizationsGetOrganizationSummaryTopApplicationsByUsage `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopApplicationsByUsage struct {
	Application types.String  `tfsdk:"application"`
	Downstream  types.Float64 `tfsdk:"downstream"`
	Percentage  types.Float64 `tfsdk:"percentage"`
	Total       types.Float64 `tfsdk:"total"`
	Upstream    types.Float64 `tfsdk:"upstream"`
}

// ToBody
func ResponseOrganizationsGetOrganizationSummaryTopApplicationsByUsageItemsToBody(state OrganizationsSummaryTopApplicationsByUsage, response *merakigosdk.ResponseOrganizationsGetOrganizationSummaryTopApplicationsByUsage) OrganizationsSummaryTopApplicationsByUsage {
	var items []ResponseItemOrganizationsGetOrganizationSummaryTopApplicationsByUsage
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationSummaryTopApplicationsByUsage{
			Application: types.StringValue(item.Application),
			Downstream: func() types.Float64 {
				if item.Downstream != nil {
					return types.Float64Value(float64(*item.Downstream))
				}
				return types.Float64{}
			}(),
			Percentage: func() types.Float64 {
				if item.Percentage != nil {
					return types.Float64Value(float64(*item.Percentage))
				}
				return types.Float64{}
			}(),
			Total: func() types.Float64 {
				if item.Total != nil {
					return types.Float64Value(float64(*item.Total))
				}
				return types.Float64{}
			}(),
			Upstream: func() types.Float64 {
				if item.Upstream != nil {
					return types.Float64Value(float64(*item.Upstream))
				}
				return types.Float64{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
