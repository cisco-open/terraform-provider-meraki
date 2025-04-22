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
	_ datasource.DataSource              = &OrganizationsCellularGatewayEsimsServiceProvidersDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsCellularGatewayEsimsServiceProvidersDataSource{}
)

func NewOrganizationsCellularGatewayEsimsServiceProvidersDataSource() datasource.DataSource {
	return &OrganizationsCellularGatewayEsimsServiceProvidersDataSource{}
}

type OrganizationsCellularGatewayEsimsServiceProvidersDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsCellularGatewayEsimsServiceProvidersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsCellularGatewayEsimsServiceProvidersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_cellular_gateway_esims_service_providers"
}

func (d *OrganizationsCellularGatewayEsimsServiceProvidersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"items": schema.ListNestedAttribute{
						MarkdownDescription: `List Cellular Service Providers`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"is_bootstrap": schema.BoolAttribute{
									MarkdownDescription: `Indicates if service provider is the bootstrap provider.`,
									Computed:            true,
								},
								"logo": schema.SingleNestedAttribute{
									MarkdownDescription: `Service Provider logo data.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"url": schema.StringAttribute{
											MarkdownDescription: `URL of service provider's logo.`,
											Computed:            true,
										},
									},
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Service provider name.`,
									Computed:            true,
								},
								"terms": schema.SingleNestedAttribute{
									MarkdownDescription: `Service provider terms.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"content": schema.StringAttribute{
											MarkdownDescription: `URL of service provider's terms.`,
											Computed:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: `Label for service provider's terms.`,
											Computed:            true,
										},
									},
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
										MarkdownDescription: `Service Providers available`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"remaining": schema.Int64Attribute{
												MarkdownDescription: `Remaining number of Service Providers`,
												Computed:            true,
											},
											"total": schema.Int64Attribute{
												MarkdownDescription: `Total number of Service Providers`,
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

func (d *OrganizationsCellularGatewayEsimsServiceProvidersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsCellularGatewayEsimsServiceProviders OrganizationsCellularGatewayEsimsServiceProviders
	diags := req.Config.Get(ctx, &organizationsCellularGatewayEsimsServiceProviders)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationCellularGatewayEsimsServiceProviders")
		vvOrganizationID := organizationsCellularGatewayEsimsServiceProviders.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.CellularGateway.GetOrganizationCellularGatewayEsimsServiceProviders(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCellularGatewayEsimsServiceProviders",
				err.Error(),
			)
			return
		}

		organizationsCellularGatewayEsimsServiceProviders = ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItemToBody(organizationsCellularGatewayEsimsServiceProviders, response1)
		diags = resp.State.Set(ctx, &organizationsCellularGatewayEsimsServiceProviders)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsCellularGatewayEsimsServiceProviders struct {
	OrganizationID types.String                                                                `tfsdk:"organization_id"`
	Item           *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProviders `tfsdk:"item"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProviders struct {
	Items *[]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItems `tfsdk:"items"`
	Meta  *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersMeta    `tfsdk:"meta"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItems struct {
	IsBootstrap types.Bool                                                                            `tfsdk:"is_bootstrap"`
	Logo        *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItemsLogo  `tfsdk:"logo"`
	Name        types.String                                                                          `tfsdk:"name"`
	Terms       *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItemsTerms `tfsdk:"terms"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItemsLogo struct {
	URL types.String `tfsdk:"url"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItemsTerms struct {
	Content types.String `tfsdk:"content"`
	Name    types.String `tfsdk:"name"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersMeta struct {
	Counts *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersMetaCounts `tfsdk:"counts"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersMetaCounts struct {
	Items *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersMetaCountsItems `tfsdk:"items"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItemToBody(state OrganizationsCellularGatewayEsimsServiceProviders, response *merakigosdk.ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProviders) OrganizationsCellularGatewayEsimsServiceProviders {
	itemState := ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProviders{
		Items: func() *[]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItems {
			if response.Items != nil {
				result := make([]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItems{
						IsBootstrap: func() types.Bool {
							if items.IsBootstrap != nil {
								return types.BoolValue(*items.IsBootstrap)
							}
							return types.Bool{}
						}(),
						Logo: func() *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItemsLogo {
							if items.Logo != nil {
								return &ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItemsLogo{
									URL: types.StringValue(items.Logo.URL),
								}
							}
							return nil
						}(),
						Name: types.StringValue(items.Name),
						Terms: func() *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItemsTerms {
							if items.Terms != nil {
								return &ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersItemsTerms{
									Content: types.StringValue(items.Terms.Content),
									Name:    types.StringValue(items.Terms.Name),
								}
							}
							return nil
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		Meta: func() *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersMeta {
			if response.Meta != nil {
				return &ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersMeta{
					Counts: func() *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersMetaCounts{
								Items: func() *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersMetaCountsItems{
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
