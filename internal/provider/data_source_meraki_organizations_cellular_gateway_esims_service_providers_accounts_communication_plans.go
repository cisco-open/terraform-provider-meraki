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
	_ datasource.DataSource              = &OrganizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansDataSource{}
)

func NewOrganizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansDataSource() datasource.DataSource {
	return &OrganizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansDataSource{}
}

type OrganizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_cellular_gateway_esims_service_providers_accounts_communication_plans"
}

func (d *OrganizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_ids": schema.ListAttribute{
				MarkdownDescription: `accountIds query parameter. Account IDs that communication plans will be fetched for`,
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
						MarkdownDescription: `List of Cellular Service Provider Communication Plans`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"account_id": schema.StringAttribute{
									MarkdownDescription: `Account ID of plans to be fetched`,
									Computed:            true,
								},
								"apns": schema.SetNestedAttribute{
									MarkdownDescription: `Available APNs`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"name": schema.StringAttribute{
												MarkdownDescription: `APN name`,
												Computed:            true,
											},
										},
									},
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Communication plan name`,
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
										MarkdownDescription: `Count of Communication Plans available`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"remaining": schema.Int64Attribute{
												MarkdownDescription: `Remaining number of Communication Plans`,
												Computed:            true,
											},
											"total": schema.Int64Attribute{
												MarkdownDescription: `Total number of Communication Plans`,
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

func (d *OrganizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans OrganizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans
	diags := req.Config.Get(ctx, &organizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans")
		vvOrganizationID := organizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansQueryParams{}

		queryParams1.AccountIDs = elementsToStrings(ctx, organizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans.AccountIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.CellularGateway.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans",
				err.Error(),
			)
			return
		}

		organizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans = ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansItemToBody(organizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans, response1)
		diags = resp.State.Set(ctx, &organizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans struct {
	OrganizationID types.String                                                                                          `tfsdk:"organization_id"`
	AccountIDs     types.List                                                                                            `tfsdk:"account_ids"`
	Item           *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans `tfsdk:"item"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans struct {
	Items *[]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansItems `tfsdk:"items"`
	Meta  *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansMeta    `tfsdk:"meta"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansItems struct {
	AccountID types.String                                                                                                     `tfsdk:"account_id"`
	Apns      *[]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansItemsApns `tfsdk:"apns"`
	Name      types.String                                                                                                     `tfsdk:"name"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansItemsApns struct {
	Name types.String `tfsdk:"name"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansMeta struct {
	Counts *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansMetaCounts `tfsdk:"counts"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansMetaCounts struct {
	Items *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansMetaCountsItems `tfsdk:"items"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansItemToBody(state OrganizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans, response *merakigosdk.ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans) OrganizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans {
	itemState := ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans{
		Items: func() *[]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansItems {
			if response.Items != nil {
				result := make([]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansItems{
						AccountID: func() types.String {
							if items.AccountID != "" {
								return types.StringValue(items.AccountID)
							}
							return types.String{}
						}(),
						Apns: func() *[]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansItemsApns {
							if items.Apns != nil {
								result := make([]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansItemsApns, len(*items.Apns))
								for i, apns := range *items.Apns {
									result[i] = ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansItemsApns{
										Name: func() types.String {
											if apns.Name != "" {
												return types.StringValue(apns.Name)
											}
											return types.String{}
										}(),
									}
								}
								return &result
							}
							return nil
						}(),
						Name: func() types.String {
							if items.Name != "" {
								return types.StringValue(items.Name)
							}
							return types.String{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		Meta: func() *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansMeta {
			if response.Meta != nil {
				return &ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansMeta{
					Counts: func() *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansMetaCounts{
								Items: func() *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansMetaCountsItems{
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
