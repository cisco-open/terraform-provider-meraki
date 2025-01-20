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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsOpenapiSpecDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsOpenapiSpecDataSource{}
)

func NewOrganizationsOpenapiSpecDataSource() datasource.DataSource {
	return &OrganizationsOpenapiSpecDataSource{}
}

type OrganizationsOpenapiSpecDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsOpenapiSpecDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsOpenapiSpecDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_openapi_spec"
}

func (d *OrganizationsOpenapiSpecDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: `version query parameter. OpenAPI Specification version to return. Default is 2`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"info": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"description": schema.StringAttribute{
								Computed: true,
							},
							"title": schema.StringAttribute{
								Computed: true,
							},
							"version": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"openapi": schema.StringAttribute{
						Computed: true,
					},
					"paths": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"organizations": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"get": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{

											"description": schema.StringAttribute{
												Computed: true,
											},
											"operation_id": schema.StringAttribute{
												Computed: true,
											},
											"responses": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{

													"status_200": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{

															"description": schema.StringAttribute{
																Computed: true,
															},
															"examples": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{

																	"application_json": schema.SetNestedAttribute{
																		Computed: true,
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{

																				"id": schema.StringAttribute{
																					Computed: true,
																				},
																				"name": schema.StringAttribute{
																					Computed: true,
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

func (d *OrganizationsOpenapiSpecDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsOpenapiSpec OrganizationsOpenapiSpec
	diags := req.Config.Get(ctx, &organizationsOpenapiSpec)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationOpenapiSpec")
		vvOrganizationID := organizationsOpenapiSpec.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationOpenapiSpecQueryParams{}

		queryParams1.Version = int(organizationsOpenapiSpec.Version.ValueInt64())

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationOpenapiSpec(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationOpenapiSpec",
				err.Error(),
			)
			return
		}

		organizationsOpenapiSpec = ResponseOrganizationsGetOrganizationOpenapiSpecItemToBody(organizationsOpenapiSpec, response1)
		diags = resp.State.Set(ctx, &organizationsOpenapiSpec)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsOpenapiSpec struct {
	OrganizationID types.String                                     `tfsdk:"organization_id"`
	Version        types.Int64                                      `tfsdk:"version"`
	Item           *ResponseOrganizationsGetOrganizationOpenapiSpec `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationOpenapiSpec struct {
	Info    *ResponseOrganizationsGetOrganizationOpenapiSpecInfo  `tfsdk:"info"`
	Openapi types.String                                          `tfsdk:"openapi"`
	Paths   *ResponseOrganizationsGetOrganizationOpenapiSpecPaths `tfsdk:"paths"`
}

type ResponseOrganizationsGetOrganizationOpenapiSpecInfo struct {
	Description types.String `tfsdk:"description"`
	Title       types.String `tfsdk:"title"`
	Version     types.String `tfsdk:"version"`
}

type ResponseOrganizationsGetOrganizationOpenapiSpecPaths struct {
	Organizations *ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizations `tfsdk:"/organizations"`
}

type ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizations struct {
	Get *ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGet `tfsdk:"get"`
}

type ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGet struct {
	Description types.String                                                                   `tfsdk:"description"`
	OperationID types.String                                                                   `tfsdk:"operation_id"`
	Responses   *ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses `tfsdk:"responses"`
}

type ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses struct {
	Status200 *ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses200 `tfsdk:"status_200"`
}

type ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses200 struct {
	Description types.String                                                                              `tfsdk:"description"`
	Examples    *ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses200Examples `tfsdk:"examples"`
}

type ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses200Examples struct {
	ApplicationJSON *[]ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses200ExamplesApplicationJson `tfsdk:"application/json"`
}

type ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses200ExamplesApplicationJson struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// ToBody
func ResponseOrganizationsGetOrganizationOpenapiSpecItemToBody(state OrganizationsOpenapiSpec, response *merakigosdk.ResponseOrganizationsGetOrganizationOpenapiSpec) OrganizationsOpenapiSpec {
	itemState := ResponseOrganizationsGetOrganizationOpenapiSpec{
		Info: func() *ResponseOrganizationsGetOrganizationOpenapiSpecInfo {
			if response.Info != nil {
				return &ResponseOrganizationsGetOrganizationOpenapiSpecInfo{
					Description: types.StringValue(response.Info.Description),
					Title:       types.StringValue(response.Info.Title),
					Version:     types.StringValue(response.Info.Version),
				}
			}
			return nil
		}(),
		Openapi: types.StringValue(response.Openapi),
		Paths: func() *ResponseOrganizationsGetOrganizationOpenapiSpecPaths {
			if response.Paths != nil {
				return &ResponseOrganizationsGetOrganizationOpenapiSpecPaths{
					Organizations: func() *ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizations {
						if response.Paths.Organizations != nil {
							return &ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizations{
								Get: func() *ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGet {
									if response.Paths.Organizations.Get != nil {
										return &ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGet{
											Description: types.StringValue(response.Paths.Organizations.Get.Description),
											OperationID: types.StringValue(response.Paths.Organizations.Get.OperationID),
											Responses: func() *ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses {
												if response.Paths.Organizations.Get.Responses != nil {
													return &ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses{
														Status200: func() *ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses200 {
															if response.Paths.Organizations.Get.Responses.Status200 != nil {
																return &ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses200{
																	Description: types.StringValue(response.Paths.Organizations.Get.Responses.Status200.Description),
																	Examples: func() *ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses200Examples {
																		if response.Paths.Organizations.Get.Responses.Status200.Examples != nil {
																			return &ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses200Examples{
																				ApplicationJSON: func() *[]ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses200ExamplesApplicationJson {
																					if response.Paths.Organizations.Get.Responses.Status200.Examples.ApplicationJSON != nil {
																						result := make([]ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses200ExamplesApplicationJson, len(*response.Paths.Organizations.Get.Responses.Status200.Examples.ApplicationJSON))
																						for i, applicationJSON := range *response.Paths.Organizations.Get.Responses.Status200.Examples.ApplicationJSON {
																							result[i] = ResponseOrganizationsGetOrganizationOpenapiSpecPathsOrganizationsGetResponses200ExamplesApplicationJson{
																								ID:   types.StringValue(applicationJSON.ID),
																								Name: types.StringValue(applicationJSON.Name),
																							}
																						}
																						return &result
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
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
