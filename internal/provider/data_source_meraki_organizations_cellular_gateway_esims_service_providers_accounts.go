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
	_ datasource.DataSource              = &OrganizationsCellularGatewayEsimsServiceProvidersAccountsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsCellularGatewayEsimsServiceProvidersAccountsDataSource{}
)

func NewOrganizationsCellularGatewayEsimsServiceProvidersAccountsDataSource() datasource.DataSource {
	return &OrganizationsCellularGatewayEsimsServiceProvidersAccountsDataSource{}
}

type OrganizationsCellularGatewayEsimsServiceProvidersAccountsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsCellularGatewayEsimsServiceProvidersAccountsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsCellularGatewayEsimsServiceProvidersAccountsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_cellular_gateway_esims_service_providers_accounts"
}

func (d *OrganizationsCellularGatewayEsimsServiceProvidersAccountsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_ids": schema.ListAttribute{
				MarkdownDescription: `accountIds query parameter. Optional parameter to filter the results by service provider account IDs.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccounts`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"items": schema.ListNestedAttribute{
							MarkdownDescription: `IList of Cellular Service Provider Accounts`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"account_id": schema.StringAttribute{
										MarkdownDescription: `Service provider account ID`,
										Computed:            true,
									},
									"last_updated_at": schema.StringAttribute{
										MarkdownDescription: `Last updated at`,
										Computed:            true,
									},
									"service_provider": schema.SingleNestedAttribute{
										MarkdownDescription: `Service provider data.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"logo": schema.SingleNestedAttribute{
												MarkdownDescription: `Service provider logo data.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"url": schema.StringAttribute{
														MarkdownDescription: `Service Provider logo url.`,
														Computed:            true,
													},
												},
											},
											"name": schema.StringAttribute{
												MarkdownDescription: `Name of the service provider.`,
												Computed:            true,
											},
										},
									},
									"title": schema.StringAttribute{
										MarkdownDescription: `Service provider account name`,
										Computed:            true,
									},
									"username": schema.StringAttribute{
										MarkdownDescription: `Service provider account username`,
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
											MarkdownDescription: `Count of Cellular Service Providers available`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"remaining": schema.Int64Attribute{
													MarkdownDescription: `Remaining number of Cellular Service Providers`,
													Computed:            true,
												},
												"total": schema.Int64Attribute{
													MarkdownDescription: `Total number of Cellular Service Providers`,
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
		},
	}
}

func (d *OrganizationsCellularGatewayEsimsServiceProvidersAccountsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsCellularGatewayEsimsServiceProvidersAccounts OrganizationsCellularGatewayEsimsServiceProvidersAccounts
	diags := req.Config.Get(ctx, &organizationsCellularGatewayEsimsServiceProvidersAccounts)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationCellularGatewayEsimsServiceProvidersAccounts")
		vvOrganizationID := organizationsCellularGatewayEsimsServiceProvidersAccounts.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsQueryParams{}

		queryParams1.AccountIDs = elementsToStrings(ctx, organizationsCellularGatewayEsimsServiceProvidersAccounts.AccountIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.CellularGateway.GetOrganizationCellularGatewayEsimsServiceProvidersAccounts(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCellularGatewayEsimsServiceProvidersAccounts",
				err.Error(),
			)
			return
		}

		organizationsCellularGatewayEsimsServiceProvidersAccounts = ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsToBody(organizationsCellularGatewayEsimsServiceProvidersAccounts, response1)
		diags = resp.State.Set(ctx, &organizationsCellularGatewayEsimsServiceProvidersAccounts)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsCellularGatewayEsimsServiceProvidersAccounts struct {
	OrganizationID types.String                                                                            `tfsdk:"organization_id"`
	AccountIDs     types.List                                                                              `tfsdk:"account_ids"`
	Items          *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccounts `tfsdk:"items"`
}

type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccounts struct {
	Items *[]ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItems `tfsdk:"items"`
	Meta  *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsMeta    `tfsdk:"meta"`
}

type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItems struct {
	AccountID       types.String                                                                                                `tfsdk:"account_id"`
	LastUpdatedAt   types.String                                                                                                `tfsdk:"last_updated_at"`
	ServiceProvider *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsServiceProvider `tfsdk:"service_provider"`
	Title           types.String                                                                                                `tfsdk:"title"`
	Username        types.String                                                                                                `tfsdk:"username"`
}

type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsServiceProvider struct {
	Logo *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsServiceProviderLogo `tfsdk:"logo"`
	Name types.String                                                                                                    `tfsdk:"name"`
}

type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsServiceProviderLogo struct {
	URL types.String `tfsdk:"url"`
}

type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsMeta struct {
	Counts *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsMetaCounts `tfsdk:"counts"`
}

type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsMetaCounts struct {
	Items *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsMetaCountsItems `tfsdk:"items"`
}

type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsToBody(state OrganizationsCellularGatewayEsimsServiceProvidersAccounts, response *merakigosdk.ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccounts) OrganizationsCellularGatewayEsimsServiceProvidersAccounts {
	item := response
	itemState := ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccounts{
		Items: func() *[]ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItems {

			if item.Items != nil {
				result := make([]ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItems, len(*item.Items))
				for i, items := range *item.Items {
					result[i] = ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItems{
						AccountID:     types.StringValue(items.AccountID),
						LastUpdatedAt: types.StringValue(items.LastUpdatedAt),
						ServiceProvider: func() *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsServiceProvider {
							if items.ServiceProvider != nil {
								return &ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsServiceProvider{
									Logo: func() *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsServiceProviderLogo {
										if items.ServiceProvider.Logo != nil {
											return &ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsServiceProviderLogo{
												URL: types.StringValue(items.ServiceProvider.Logo.URL),
											}
										}
										return nil
									}(),
									Name: types.StringValue(items.ServiceProvider.Name),
								}
							}
							return nil
						}(),
						Title:    types.StringValue(items.Title),
						Username: types.StringValue(items.Username),
					}
				}
				return &result
			}
			return nil
		}(),
		Meta: func() *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsMeta {
			if item.Meta != nil {
				return &ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsMeta{
					Counts: func() *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsMetaCounts {
						if item.Meta.Counts != nil {
							return &ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsMetaCounts{
								Items: func() *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsMetaCountsItems {
									if item.Meta.Counts.Items != nil {
										return &ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsMetaCountsItems{
											Remaining: func() types.Int64 {
												if item.Meta.Counts.Items.Remaining != nil {
													return types.Int64Value(int64(*item.Meta.Counts.Items.Remaining))
												}
												return types.Int64{}
											}(),
											Total: func() types.Int64 {
												if item.Meta.Counts.Items.Total != nil {
													return types.Int64Value(int64(*item.Meta.Counts.Items.Total))
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
	state.Items = &itemState
	return state
}
