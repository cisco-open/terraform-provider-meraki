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
	_ datasource.DataSource              = &OrganizationsCellularGatewayEsimsServiceProvidersAccountsRatePlansDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsCellularGatewayEsimsServiceProvidersAccountsRatePlansDataSource{}
)

func NewOrganizationsCellularGatewayEsimsServiceProvidersAccountsRatePlansDataSource() datasource.DataSource {
	return &OrganizationsCellularGatewayEsimsServiceProvidersAccountsRatePlansDataSource{}
}

type OrganizationsCellularGatewayEsimsServiceProvidersAccountsRatePlansDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsCellularGatewayEsimsServiceProvidersAccountsRatePlansDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsCellularGatewayEsimsServiceProvidersAccountsRatePlansDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_cellular_gateway_esims_service_providers_accounts_rate_plans"
}

func (d *OrganizationsCellularGatewayEsimsServiceProvidersAccountsRatePlansDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_ids": schema.ListAttribute{
				MarkdownDescription: `accountIds query parameter. Account IDs that rate plans will be fetched for`,
				Required:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"items": schema.ListNestedAttribute{
						MarkdownDescription: `List of Cellular Service Provider Rate Plans`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"account_id": schema.StringAttribute{
									MarkdownDescription: `Account ID of plans to be fetched`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Rate plan name`,
									Computed:            true,
								},
							},
						},
					},
					"meta": schema.SingleNestedAttribute{
						MarkdownDescription: `Meta details about the result`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"counts": schema.SingleNestedAttribute{
								MarkdownDescription: `Counts of involved entities`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"items": schema.SingleNestedAttribute{
										MarkdownDescription: `Count of Rate Plans available`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"remaining": schema.Int64Attribute{
												MarkdownDescription: `Remaining number of Rate Plans`,
												Computed:            true,
											},
											"total": schema.Int64Attribute{
												MarkdownDescription: `Total number of Rate Plans`,
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

func (d *OrganizationsCellularGatewayEsimsServiceProvidersAccountsRatePlansDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsCellularGatewayEsimsServiceProvidersAccountsRatePlans OrganizationsCellularGatewayEsimsServiceProvidersAccountsRatePlans
	diags := req.Config.Get(ctx, &organizationsCellularGatewayEsimsServiceProvidersAccountsRatePlans)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans")
		vvOrganizationID := organizationsCellularGatewayEsimsServiceProvidersAccountsRatePlans.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansQueryParams{}

		queryParams1.AccountIDs = elementsToStrings(ctx, organizationsCellularGatewayEsimsServiceProvidersAccountsRatePlans.AccountIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.CellularGateway.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans",
				err.Error(),
			)
			return
		}

		organizationsCellularGatewayEsimsServiceProvidersAccountsRatePlans = ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansItemToBody(organizationsCellularGatewayEsimsServiceProvidersAccountsRatePlans, response1)
		diags = resp.State.Set(ctx, &organizationsCellularGatewayEsimsServiceProvidersAccountsRatePlans)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsCellularGatewayEsimsServiceProvidersAccountsRatePlans struct {
	OrganizationID types.String                                                                                 `tfsdk:"organization_id"`
	AccountIDs     types.List                                                                                   `tfsdk:"account_ids"`
	Item           *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans `tfsdk:"item"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans struct {
	Items *[]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansItems `tfsdk:"items"`
	Meta  *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansMeta    `tfsdk:"meta"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansItems struct {
	AccountID types.String `tfsdk:"account_id"`
	Name      types.String `tfsdk:"name"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansMeta struct {
	Counts *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansMetaCounts `tfsdk:"counts"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansMetaCounts struct {
	Items *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansMetaCountsItems `tfsdk:"items"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansItemToBody(state OrganizationsCellularGatewayEsimsServiceProvidersAccountsRatePlans, response *merakigosdk.ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans) OrganizationsCellularGatewayEsimsServiceProvidersAccountsRatePlans {
	itemState := ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans{
		Items: func() *[]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansItems {
			if response.Items != nil {
				result := make([]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansItems{
						AccountID: types.StringValue(items.AccountID),
						Name:      types.StringValue(items.Name),
					}
				}
				return &result
			}
			return nil
		}(),
		Meta: func() *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansMeta {
			if response.Meta != nil {
				return &ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansMeta{
					Counts: func() *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansMetaCounts{
								Items: func() *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansMetaCountsItems{
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
