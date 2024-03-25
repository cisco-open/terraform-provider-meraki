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
	_ datasource.DataSource              = &OrganizationsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsDataSource{}
)

func NewOrganizationsDataSource() datasource.DataSource {
	return &OrganizationsDataSource{}
}

type OrganizationsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations"
}

func (d *OrganizationsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"api": schema.SingleNestedAttribute{
						MarkdownDescription: `API related settings`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Enable API access`,
								Computed:            true,
							},
						},
					},
					"cloud": schema.SingleNestedAttribute{
						MarkdownDescription: `Data for this organization`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"region": schema.SingleNestedAttribute{
								MarkdownDescription: `Region info`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"name": schema.StringAttribute{
										MarkdownDescription: `Name of region`,
										Computed:            true,
									},
								},
							},
						},
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `Organization ID`,
						Computed:            true,
					},
					"licensing": schema.SingleNestedAttribute{
						MarkdownDescription: `Licensing related settings`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"model": schema.StringAttribute{
								MarkdownDescription: `Organization licensing model. Can be 'co-term', 'per-device', or 'subscription'.`,
								Computed:            true,
							},
						},
					},
					"management": schema.SingleNestedAttribute{
						MarkdownDescription: `Information about the organization's management system`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"details": schema.SetNestedAttribute{
								MarkdownDescription: `Details related to organization management, possibly empty`,
								Computed:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{

										"name": schema.StringAttribute{
											MarkdownDescription: `Name of management data`,
											Computed:            true,
										},
										"value": schema.StringAttribute{
											MarkdownDescription: `Value of management data`,
											Computed:            true,
										},
									},
								},
							},
						},
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Organization name`,
						Computed:            true,
					},
					"url": schema.StringAttribute{
						MarkdownDescription: `Organization URL`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizations`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"api": schema.SingleNestedAttribute{
							MarkdownDescription: `API related settings`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"enabled": schema.BoolAttribute{
									MarkdownDescription: `Enable API access`,
									Computed:            true,
								},
							},
						},
						"cloud": schema.SingleNestedAttribute{
							MarkdownDescription: `Data for this organization`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"region": schema.SingleNestedAttribute{
									MarkdownDescription: `Region info`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"name": schema.StringAttribute{
											MarkdownDescription: `Name of region`,
											Computed:            true,
										},
									},
								},
							},
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `Organization ID`,
							Computed:            true,
						},
						"licensing": schema.SingleNestedAttribute{
							MarkdownDescription: `Licensing related settings`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"model": schema.StringAttribute{
									MarkdownDescription: `Organization licensing model. Can be 'co-term', 'per-device', or 'subscription'.`,
									Computed:            true,
								},
							},
						},
						"management": schema.SingleNestedAttribute{
							MarkdownDescription: `Information about the organization's management system`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"details": schema.SetNestedAttribute{
									MarkdownDescription: `Details related to organization management, possibly empty`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"name": schema.StringAttribute{
												MarkdownDescription: `Name of management data`,
												Computed:            true,
											},
											"value": schema.StringAttribute{
												MarkdownDescription: `Value of management data`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Organization name`,
							Computed:            true,
						},
						"url": schema.StringAttribute{
							MarkdownDescription: `Organization URL`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizations Organizations
	diags := req.Config.Get(ctx, &organizations)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizations.OrganizationID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizations")

		response1, restyResp1, err := d.client.Organizations.GetOrganizations()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizations",
				err.Error(),
			)
			return
		}

		organizations = ResponseOrganizationsGetOrganizationsItemsToBody(organizations, response1)
		diags = resp.State.Set(ctx, &organizations)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganization")
		vvOrganizationID := organizations.OrganizationID.ValueString()

		response2, restyResp2, err := d.client.Organizations.GetOrganization(vvOrganizationID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganization",
				err.Error(),
			)
			return
		}

		organizations = ResponseOrganizationsGetOrganizationItemToBody(organizations, response2)
		diags = resp.State.Set(ctx, &organizations)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type Organizations struct {
	OrganizationID types.String                                 `tfsdk:"organization_id"`
	Items          *[]ResponseItemOrganizationsGetOrganizations `tfsdk:"items"`
	Item           *ResponseOrganizationsGetOrganization        `tfsdk:"item"`
}

type ResponseItemOrganizationsGetOrganizations struct {
	API        *ResponseItemOrganizationsGetOrganizationsApi        `tfsdk:"api"`
	Cloud      *ResponseItemOrganizationsGetOrganizationsCloud      `tfsdk:"cloud"`
	ID         types.String                                         `tfsdk:"id"`
	Licensing  *ResponseItemOrganizationsGetOrganizationsLicensing  `tfsdk:"licensing"`
	Management *ResponseItemOrganizationsGetOrganizationsManagement `tfsdk:"management"`
	Name       types.String                                         `tfsdk:"name"`
	URL        types.String                                         `tfsdk:"url"`
}

type ResponseItemOrganizationsGetOrganizationsApi struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseItemOrganizationsGetOrganizationsCloud struct {
	Region *ResponseItemOrganizationsGetOrganizationsCloudRegion `tfsdk:"region"`
}

type ResponseItemOrganizationsGetOrganizationsCloudRegion struct {
	Name types.String `tfsdk:"name"`
}

type ResponseItemOrganizationsGetOrganizationsLicensing struct {
	Model types.String `tfsdk:"model"`
}

type ResponseItemOrganizationsGetOrganizationsManagement struct {
	Details *[]ResponseItemOrganizationsGetOrganizationsManagementDetails `tfsdk:"details"`
}

type ResponseItemOrganizationsGetOrganizationsManagementDetails struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type ResponseOrganizationsGetOrganization struct {
	API        *ResponseOrganizationsGetOrganizationApi        `tfsdk:"api"`
	Cloud      *ResponseOrganizationsGetOrganizationCloud      `tfsdk:"cloud"`
	ID         types.String                                    `tfsdk:"id"`
	Licensing  *ResponseOrganizationsGetOrganizationLicensing  `tfsdk:"licensing"`
	Management *ResponseOrganizationsGetOrganizationManagement `tfsdk:"management"`
	Name       types.String                                    `tfsdk:"name"`
	URL        types.String                                    `tfsdk:"url"`
}

type ResponseOrganizationsGetOrganizationApi struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseOrganizationsGetOrganizationCloud struct {
	Region *ResponseOrganizationsGetOrganizationCloudRegion `tfsdk:"region"`
}

type ResponseOrganizationsGetOrganizationCloudRegion struct {
	Name types.String `tfsdk:"name"`
}

type ResponseOrganizationsGetOrganizationLicensing struct {
	Model types.String `tfsdk:"model"`
}

type ResponseOrganizationsGetOrganizationManagement struct {
	Details *[]ResponseOrganizationsGetOrganizationManagementDetails `tfsdk:"details"`
}

type ResponseOrganizationsGetOrganizationManagementDetails struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

// ToBody
func ResponseOrganizationsGetOrganizationsItemsToBody(state Organizations, response *merakigosdk.ResponseOrganizationsGetOrganizations) Organizations {
	var items []ResponseItemOrganizationsGetOrganizations
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizations{
			API: func() *ResponseItemOrganizationsGetOrganizationsApi {
				if item.API != nil {
					return &ResponseItemOrganizationsGetOrganizationsApi{
						Enabled: func() types.Bool {
							if item.API.Enabled != nil {
								return types.BoolValue(*item.API.Enabled)
							}
							return types.Bool{}
						}(),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationsApi{}
			}(),
			Cloud: func() *ResponseItemOrganizationsGetOrganizationsCloud {
				if item.Cloud != nil {
					return &ResponseItemOrganizationsGetOrganizationsCloud{
						Region: func() *ResponseItemOrganizationsGetOrganizationsCloudRegion {
							if item.Cloud.Region != nil {
								return &ResponseItemOrganizationsGetOrganizationsCloudRegion{
									Name: types.StringValue(item.Cloud.Region.Name),
								}
							}
							return &ResponseItemOrganizationsGetOrganizationsCloudRegion{}
						}(),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationsCloud{}
			}(),
			ID: types.StringValue(item.ID),
			Licensing: func() *ResponseItemOrganizationsGetOrganizationsLicensing {
				if item.Licensing != nil {
					return &ResponseItemOrganizationsGetOrganizationsLicensing{
						Model: types.StringValue(item.Licensing.Model),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationsLicensing{}
			}(),
			Management: func() *ResponseItemOrganizationsGetOrganizationsManagement {
				if item.Management != nil {
					return &ResponseItemOrganizationsGetOrganizationsManagement{
						Details: func() *[]ResponseItemOrganizationsGetOrganizationsManagementDetails {
							if item.Management.Details != nil {
								result := make([]ResponseItemOrganizationsGetOrganizationsManagementDetails, len(*item.Management.Details))
								for i, details := range *item.Management.Details {
									result[i] = ResponseItemOrganizationsGetOrganizationsManagementDetails{
										Name:  types.StringValue(details.Name),
										Value: types.StringValue(details.Value),
									}
								}
								return &result
							}
							return &[]ResponseItemOrganizationsGetOrganizationsManagementDetails{}
						}(),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationsManagement{}
			}(),
			Name: types.StringValue(item.Name),
			URL:  types.StringValue(item.URL),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseOrganizationsGetOrganizationItemToBody(state Organizations, response *merakigosdk.ResponseOrganizationsGetOrganization) Organizations {
	itemState := ResponseOrganizationsGetOrganization{
		API: func() *ResponseOrganizationsGetOrganizationApi {
			if response.API != nil {
				return &ResponseOrganizationsGetOrganizationApi{
					Enabled: func() types.Bool {
						if response.API.Enabled != nil {
							return types.BoolValue(*response.API.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseOrganizationsGetOrganizationApi{}
		}(),
		Cloud: func() *ResponseOrganizationsGetOrganizationCloud {
			if response.Cloud != nil {
				return &ResponseOrganizationsGetOrganizationCloud{
					Region: func() *ResponseOrganizationsGetOrganizationCloudRegion {
						if response.Cloud.Region != nil {
							return &ResponseOrganizationsGetOrganizationCloudRegion{
								Name: types.StringValue(response.Cloud.Region.Name),
							}
						}
						return &ResponseOrganizationsGetOrganizationCloudRegion{}
					}(),
				}
			}
			return &ResponseOrganizationsGetOrganizationCloud{}
		}(),
		ID: types.StringValue(response.ID),
		Licensing: func() *ResponseOrganizationsGetOrganizationLicensing {
			if response.Licensing != nil {
				return &ResponseOrganizationsGetOrganizationLicensing{
					Model: types.StringValue(response.Licensing.Model),
				}
			}
			return &ResponseOrganizationsGetOrganizationLicensing{}
		}(),
		Management: func() *ResponseOrganizationsGetOrganizationManagement {
			if response.Management != nil {
				return &ResponseOrganizationsGetOrganizationManagement{
					Details: func() *[]ResponseOrganizationsGetOrganizationManagementDetails {
						if response.Management.Details != nil {
							result := make([]ResponseOrganizationsGetOrganizationManagementDetails, len(*response.Management.Details))
							for i, details := range *response.Management.Details {
								result[i] = ResponseOrganizationsGetOrganizationManagementDetails{
									Name:  types.StringValue(details.Name),
									Value: types.StringValue(details.Value),
								}
							}
							return &result
						}
						return &[]ResponseOrganizationsGetOrganizationManagementDetails{}
					}(),
				}
			}
			return &ResponseOrganizationsGetOrganizationManagement{}
		}(),
		Name: types.StringValue(response.Name),
		URL:  types.StringValue(response.URL),
	}
	state.Item = &itemState
	return state
}
