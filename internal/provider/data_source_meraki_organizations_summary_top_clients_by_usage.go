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
	_ datasource.DataSource              = &OrganizationsSummaryTopClientsByUsageDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSummaryTopClientsByUsageDataSource{}
)

func NewOrganizationsSummaryTopClientsByUsageDataSource() datasource.DataSource {
	return &OrganizationsSummaryTopClientsByUsageDataSource{}
}

type OrganizationsSummaryTopClientsByUsageDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSummaryTopClientsByUsageDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSummaryTopClientsByUsageDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_summary_top_clients_by_usage"
}

func (d *OrganizationsSummaryTopClientsByUsageDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"device_tag": schema.StringAttribute{
				MarkdownDescription: `deviceTag query parameter. Match result to an exact device tag`,
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
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 8 hours and be less than or equal to 186 days. The default is 1 day.`,
				Optional:            true,
			},
			"usage_uplink": schema.StringAttribute{
				MarkdownDescription: `usageUplink query parameter. Filter results by usage uplink`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationSummaryTopClientsByUsage`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"id": schema.StringAttribute{
							MarkdownDescription: `ID of client`,
							Computed:            true,
						},
						"mac": schema.StringAttribute{
							MarkdownDescription: `MAC address of client`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of client`,
							Computed:            true,
						},
						"network": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `ID of network`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Name of network`,
									Computed:            true,
								},
							},
						},
						"usage": schema.SingleNestedAttribute{
							MarkdownDescription: `Data usage information`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"downstream": schema.Float64Attribute{
									MarkdownDescription: `Downstream data usage by client`,
									Computed:            true,
								},
								"percentage": schema.Float64Attribute{
									MarkdownDescription: `Percentage of total data usage by client`,
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

func (d *OrganizationsSummaryTopClientsByUsageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSummaryTopClientsByUsage OrganizationsSummaryTopClientsByUsage
	diags := req.Config.Get(ctx, &organizationsSummaryTopClientsByUsage)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSummaryTopClientsByUsage")
		vvOrganizationID := organizationsSummaryTopClientsByUsage.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationSummaryTopClientsByUsageQueryParams{}

		queryParams1.NetworkTag = organizationsSummaryTopClientsByUsage.NetworkTag.ValueString()
		queryParams1.DeviceTag = organizationsSummaryTopClientsByUsage.DeviceTag.ValueString()
		queryParams1.Quantity = int(organizationsSummaryTopClientsByUsage.Quantity.ValueInt64())
		queryParams1.SSIDName = organizationsSummaryTopClientsByUsage.SSIDName.ValueString()
		queryParams1.UsageUplink = organizationsSummaryTopClientsByUsage.UsageUplink.ValueString()
		queryParams1.T0 = organizationsSummaryTopClientsByUsage.T0.ValueString()
		queryParams1.T1 = organizationsSummaryTopClientsByUsage.T1.ValueString()
		queryParams1.Timespan = organizationsSummaryTopClientsByUsage.Timespan.ValueFloat64()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationSummaryTopClientsByUsage(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSummaryTopClientsByUsage",
				err.Error(),
			)
			return
		}

		organizationsSummaryTopClientsByUsage = ResponseOrganizationsGetOrganizationSummaryTopClientsByUsageItemsToBody(organizationsSummaryTopClientsByUsage, response1)
		diags = resp.State.Set(ctx, &organizationsSummaryTopClientsByUsage)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSummaryTopClientsByUsage struct {
	OrganizationID types.String                                                        `tfsdk:"organization_id"`
	NetworkTag     types.String                                                        `tfsdk:"network_tag"`
	DeviceTag      types.String                                                        `tfsdk:"device_tag"`
	Quantity       types.Int64                                                         `tfsdk:"quantity"`
	SSIDName       types.String                                                        `tfsdk:"ssid_name"`
	UsageUplink    types.String                                                        `tfsdk:"usage_uplink"`
	T0             types.String                                                        `tfsdk:"t0"`
	T1             types.String                                                        `tfsdk:"t1"`
	Timespan       types.Float64                                                       `tfsdk:"timespan"`
	Items          *[]ResponseItemOrganizationsGetOrganizationSummaryTopClientsByUsage `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopClientsByUsage struct {
	ID      types.String                                                             `tfsdk:"id"`
	Mac     types.String                                                             `tfsdk:"mac"`
	Name    types.String                                                             `tfsdk:"name"`
	Network *ResponseItemOrganizationsGetOrganizationSummaryTopClientsByUsageNetwork `tfsdk:"network"`
	Usage   *ResponseItemOrganizationsGetOrganizationSummaryTopClientsByUsageUsage   `tfsdk:"usage"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopClientsByUsageNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopClientsByUsageUsage struct {
	Downstream types.Float64 `tfsdk:"downstream"`
	Percentage types.Float64 `tfsdk:"percentage"`
	Total      types.Float64 `tfsdk:"total"`
	Upstream   types.Float64 `tfsdk:"upstream"`
}

// ToBody
func ResponseOrganizationsGetOrganizationSummaryTopClientsByUsageItemsToBody(state OrganizationsSummaryTopClientsByUsage, response *merakigosdk.ResponseOrganizationsGetOrganizationSummaryTopClientsByUsage) OrganizationsSummaryTopClientsByUsage {
	var items []ResponseItemOrganizationsGetOrganizationSummaryTopClientsByUsage
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationSummaryTopClientsByUsage{
			ID:   types.StringValue(item.ID),
			Mac:  types.StringValue(item.Mac),
			Name: types.StringValue(item.Name),
			Network: func() *ResponseItemOrganizationsGetOrganizationSummaryTopClientsByUsageNetwork {
				if item.Network != nil {
					return &ResponseItemOrganizationsGetOrganizationSummaryTopClientsByUsageNetwork{
						ID:   types.StringValue(item.Network.ID),
						Name: types.StringValue(item.Network.Name),
					}
				}
				return nil
			}(),
			Usage: func() *ResponseItemOrganizationsGetOrganizationSummaryTopClientsByUsageUsage {
				if item.Usage != nil {
					return &ResponseItemOrganizationsGetOrganizationSummaryTopClientsByUsageUsage{
						Downstream: func() types.Float64 {
							if item.Usage.Downstream != nil {
								return types.Float64Value(float64(*item.Usage.Downstream))
							}
							return types.Float64{}
						}(),
						Percentage: func() types.Float64 {
							if item.Usage.Percentage != nil {
								return types.Float64Value(float64(*item.Usage.Percentage))
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
				return nil
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
