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
	_ datasource.DataSource              = &OrganizationsInsightApplicationsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsInsightApplicationsDataSource{}
)

func NewOrganizationsInsightApplicationsDataSource() datasource.DataSource {
	return &OrganizationsInsightApplicationsDataSource{}
}

type OrganizationsInsightApplicationsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsInsightApplicationsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsInsightApplicationsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_insight_applications"
}

func (d *OrganizationsInsightApplicationsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseInsightGetOrganizationInsightApplications`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"application_id": schema.StringAttribute{
							MarkdownDescription: `Application identifier`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Application name`,
							Computed:            true,
						},
						"thresholds": schema.SingleNestedAttribute{
							MarkdownDescription: `Thresholds defined by a user or Meraki models for each application`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"by_network": schema.SetNestedAttribute{
									MarkdownDescription: `Threshold for each network`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"goodput": schema.Int64Attribute{
												MarkdownDescription: `Number of useful information bits delivered over a network per unit of time`,
												Computed:            true,
											},
											"network_id": schema.StringAttribute{
												MarkdownDescription: `Network identifier`,
												Computed:            true,
											},
											"response_duration": schema.Int64Attribute{
												MarkdownDescription: `Duration of the response, in milliseconds`,
												Computed:            true,
											},
										},
									},
								},
								"type": schema.StringAttribute{
									MarkdownDescription: `Threshold type (static or smart)`,
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

func (d *OrganizationsInsightApplicationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsInsightApplications OrganizationsInsightApplications
	diags := req.Config.Get(ctx, &organizationsInsightApplications)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationInsightApplications")
		vvOrganizationID := organizationsInsightApplications.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Insight.GetOrganizationInsightApplications(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationInsightApplications",
				err.Error(),
			)
			return
		}

		organizationsInsightApplications = ResponseInsightGetOrganizationInsightApplicationsItemsToBody(organizationsInsightApplications, response1)
		diags = resp.State.Set(ctx, &organizationsInsightApplications)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsInsightApplications struct {
	OrganizationID types.String                                             `tfsdk:"organization_id"`
	Items          *[]ResponseItemInsightGetOrganizationInsightApplications `tfsdk:"items"`
}

type ResponseItemInsightGetOrganizationInsightApplications struct {
	ApplicationID types.String                                                     `tfsdk:"application_id"`
	Name          types.String                                                     `tfsdk:"name"`
	Thresholds    *ResponseItemInsightGetOrganizationInsightApplicationsThresholds `tfsdk:"thresholds"`
}

type ResponseItemInsightGetOrganizationInsightApplicationsThresholds struct {
	ByNetwork *[]ResponseItemInsightGetOrganizationInsightApplicationsThresholdsByNetwork `tfsdk:"by_network"`
	Type      types.String                                                                `tfsdk:"type"`
}

type ResponseItemInsightGetOrganizationInsightApplicationsThresholdsByNetwork struct {
	Goodput          types.Int64  `tfsdk:"goodput"`
	NetworkID        types.String `tfsdk:"network_id"`
	ResponseDuration types.Int64  `tfsdk:"response_duration"`
}

// ToBody
func ResponseInsightGetOrganizationInsightApplicationsItemsToBody(state OrganizationsInsightApplications, response *merakigosdk.ResponseInsightGetOrganizationInsightApplications) OrganizationsInsightApplications {
	var items []ResponseItemInsightGetOrganizationInsightApplications
	for _, item := range *response {
		itemState := ResponseItemInsightGetOrganizationInsightApplications{
			ApplicationID: func() types.String {
				if item.ApplicationID != "" {
					return types.StringValue(item.ApplicationID)
				}
				return types.String{}
			}(),
			Name: func() types.String {
				if item.Name != "" {
					return types.StringValue(item.Name)
				}
				return types.String{}
			}(),
			Thresholds: func() *ResponseItemInsightGetOrganizationInsightApplicationsThresholds {
				if item.Thresholds != nil {
					return &ResponseItemInsightGetOrganizationInsightApplicationsThresholds{
						ByNetwork: func() *[]ResponseItemInsightGetOrganizationInsightApplicationsThresholdsByNetwork {
							if item.Thresholds.ByNetwork != nil {
								result := make([]ResponseItemInsightGetOrganizationInsightApplicationsThresholdsByNetwork, len(*item.Thresholds.ByNetwork))
								for i, byNetwork := range *item.Thresholds.ByNetwork {
									result[i] = ResponseItemInsightGetOrganizationInsightApplicationsThresholdsByNetwork{
										Goodput: func() types.Int64 {
											if byNetwork.Goodput != nil {
												return types.Int64Value(int64(*byNetwork.Goodput))
											}
											return types.Int64{}
										}(),
										NetworkID: func() types.String {
											if byNetwork.NetworkID != "" {
												return types.StringValue(byNetwork.NetworkID)
											}
											return types.String{}
										}(),
										ResponseDuration: func() types.Int64 {
											if byNetwork.ResponseDuration != nil {
												return types.Int64Value(int64(*byNetwork.ResponseDuration))
											}
											return types.Int64{}
										}(),
									}
								}
								return &result
							}
							return nil
						}(),
						Type: func() types.String {
							if item.Thresholds.Type != "" {
								return types.StringValue(item.Thresholds.Type)
							}
							return types.String{}
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
