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
	_ datasource.DataSource              = &OrganizationsAPIRequestsOverviewDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsAPIRequestsOverviewDataSource{}
)

func NewOrganizationsAPIRequestsOverviewDataSource() datasource.DataSource {
	return &OrganizationsAPIRequestsOverviewDataSource{}
}

type OrganizationsAPIRequestsOverviewDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsAPIRequestsOverviewDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsAPIRequestsOverviewDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_api_requests_overview"
}

func (d *OrganizationsAPIRequestsOverviewDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 31 days from today.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 31 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"response_code_counts": schema.SingleNestedAttribute{
						MarkdownDescription: `object of all supported HTTP response code`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"status_200": schema.Int64Attribute{
								MarkdownDescription: `HTTP 200 response code count.`,
								Computed:            true,
							},
							"status_201": schema.Int64Attribute{
								MarkdownDescription: `HTTP 201 response code count.`,
								Computed:            true,
							},
							"status_202": schema.Int64Attribute{
								MarkdownDescription: `HTTP 202 response code count.`,
								Computed:            true,
							},
							"status_203": schema.Int64Attribute{
								MarkdownDescription: `HTTP 203 response code count.`,
								Computed:            true,
							},
							"status_204": schema.Int64Attribute{
								MarkdownDescription: `HTTP 204 response code count.`,
								Computed:            true,
							},
							"status_205": schema.Int64Attribute{
								MarkdownDescription: `HTTP 205 response code count.`,
								Computed:            true,
							},
							"status_206": schema.Int64Attribute{
								MarkdownDescription: `HTTP 206 response code count.`,
								Computed:            true,
							},
							"status_207": schema.Int64Attribute{
								MarkdownDescription: `HTTP 207 response code count.`,
								Computed:            true,
							},
							"status_208": schema.Int64Attribute{
								MarkdownDescription: `HTTP 208 response code count.`,
								Computed:            true,
							},
							"status_226": schema.Int64Attribute{
								MarkdownDescription: `HTTP 226 response code count.`,
								Computed:            true,
							},
							"status_300": schema.Int64Attribute{
								MarkdownDescription: `HTTP 300 response code count.`,
								Computed:            true,
							},
							"status_301": schema.Int64Attribute{
								MarkdownDescription: `HTTP 301 response code count.`,
								Computed:            true,
							},
							"status_302": schema.Int64Attribute{
								MarkdownDescription: `HTTP 302 response code count.`,
								Computed:            true,
							},
							"status_303": schema.Int64Attribute{
								MarkdownDescription: `HTTP 303 response code count.`,
								Computed:            true,
							},
							"status_304": schema.Int64Attribute{
								MarkdownDescription: `HTTP 304 response code count.`,
								Computed:            true,
							},
							"status_305": schema.Int64Attribute{
								MarkdownDescription: `HTTP 305 response code count.`,
								Computed:            true,
							},
							"status_306": schema.Int64Attribute{
								MarkdownDescription: `HTTP 306 response code count.`,
								Computed:            true,
							},
							"status_307": schema.Int64Attribute{
								MarkdownDescription: `HTTP 307 response code count.`,
								Computed:            true,
							},
							"status_308": schema.Int64Attribute{
								MarkdownDescription: `HTTP 308 response code count.`,
								Computed:            true,
							},
							"status_400": schema.Int64Attribute{
								MarkdownDescription: `HTTP 400 response code count.`,
								Computed:            true,
							},
							"status_401": schema.Int64Attribute{
								MarkdownDescription: `HTTP 401 response code count.`,
								Computed:            true,
							},
							"status_402": schema.Int64Attribute{
								MarkdownDescription: `HTTP 402 response code count.`,
								Computed:            true,
							},
							"status_403": schema.Int64Attribute{
								MarkdownDescription: `HTTP 403 response code count.`,
								Computed:            true,
							},
							"status_404": schema.Int64Attribute{
								MarkdownDescription: `HTTP 404 response code count.`,
								Computed:            true,
							},
							"status_405": schema.Int64Attribute{
								MarkdownDescription: `HTTP 405 response code count.`,
								Computed:            true,
							},
							"status_406": schema.Int64Attribute{
								MarkdownDescription: `HTTP 406 response code count.`,
								Computed:            true,
							},
							"status_407": schema.Int64Attribute{
								MarkdownDescription: `HTTP 407 response code count.`,
								Computed:            true,
							},
							"status_408": schema.Int64Attribute{
								MarkdownDescription: `HTTP 408 response code count.`,
								Computed:            true,
							},
							"status_409": schema.Int64Attribute{
								MarkdownDescription: `HTTP 409 response code count.`,
								Computed:            true,
							},
							"status_410": schema.Int64Attribute{
								MarkdownDescription: `HTTP 410 response code count.`,
								Computed:            true,
							},
							"status_411": schema.Int64Attribute{
								MarkdownDescription: `HTTP 411 response code count.`,
								Computed:            true,
							},
							"status_412": schema.Int64Attribute{
								MarkdownDescription: `HTTP 412 response code count.`,
								Computed:            true,
							},
							"status_413": schema.Int64Attribute{
								MarkdownDescription: `HTTP 413 response code count.`,
								Computed:            true,
							},
							"status_414": schema.Int64Attribute{
								MarkdownDescription: `HTTP 414 response code count.`,
								Computed:            true,
							},
							"status_415": schema.Int64Attribute{
								MarkdownDescription: `HTTP 415 response code count.`,
								Computed:            true,
							},
							"status_416": schema.Int64Attribute{
								MarkdownDescription: `HTTP 416 response code count.`,
								Computed:            true,
							},
							"status_417": schema.Int64Attribute{
								MarkdownDescription: `HTTP 417 response code count.`,
								Computed:            true,
							},
							"status_421": schema.Int64Attribute{
								MarkdownDescription: `HTTP 421 response code count.`,
								Computed:            true,
							},
							"status_422": schema.Int64Attribute{
								MarkdownDescription: `HTTP 422 response code count.`,
								Computed:            true,
							},
							"status_423": schema.Int64Attribute{
								MarkdownDescription: `HTTP 423 response code count.`,
								Computed:            true,
							},
							"status_424": schema.Int64Attribute{
								MarkdownDescription: `HTTP 424 response code count.`,
								Computed:            true,
							},
							"status_425": schema.Int64Attribute{
								MarkdownDescription: `HTTP 425 response code count.`,
								Computed:            true,
							},
							"status_426": schema.Int64Attribute{
								MarkdownDescription: `HTTP 426 response code count.`,
								Computed:            true,
							},
							"status_428": schema.Int64Attribute{
								MarkdownDescription: `HTTP 428 response code count.`,
								Computed:            true,
							},
							"status_429": schema.Int64Attribute{
								MarkdownDescription: `HTTP 429 response code count.`,
								Computed:            true,
							},
							"status_431": schema.Int64Attribute{
								MarkdownDescription: `HTTP 431 response code count.`,
								Computed:            true,
							},
							"status_451": schema.Int64Attribute{
								MarkdownDescription: `HTTP 451 response code count.`,
								Computed:            true,
							},
							"status_500": schema.Int64Attribute{
								MarkdownDescription: `HTTP 500 response code count.`,
								Computed:            true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsAPIRequestsOverviewDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsAPIRequestsOverview OrganizationsAPIRequestsOverview
	diags := req.Config.Get(ctx, &organizationsAPIRequestsOverview)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationAPIRequestsOverview")
		vvOrganizationID := organizationsAPIRequestsOverview.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationAPIRequestsOverviewQueryParams{}

		queryParams1.T0 = organizationsAPIRequestsOverview.T0.ValueString()
		queryParams1.T1 = organizationsAPIRequestsOverview.T1.ValueString()
		queryParams1.Timespan = organizationsAPIRequestsOverview.Timespan.ValueFloat64()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationAPIRequestsOverview(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAPIRequestsOverview",
				err.Error(),
			)
			return
		}

		organizationsAPIRequestsOverview = ResponseOrganizationsGetOrganizationAPIRequestsOverviewItemToBody(organizationsAPIRequestsOverview, response1)
		diags = resp.State.Set(ctx, &organizationsAPIRequestsOverview)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsAPIRequestsOverview struct {
	OrganizationID types.String                                             `tfsdk:"organization_id"`
	T0             types.String                                             `tfsdk:"t0"`
	T1             types.String                                             `tfsdk:"t1"`
	Timespan       types.Float64                                            `tfsdk:"timespan"`
	Item           *ResponseOrganizationsGetOrganizationApiRequestsOverview `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationApiRequestsOverview struct {
	ResponseCodeCounts *ResponseOrganizationsGetOrganizationApiRequestsOverviewResponseCodeCounts `tfsdk:"response_code_counts"`
}

type ResponseOrganizationsGetOrganizationApiRequestsOverviewResponseCodeCounts struct {
	Status200 types.Int64 `tfsdk:"status_200"`
	Status201 types.Int64 `tfsdk:"status_201"`
	Status202 types.Int64 `tfsdk:"status_202"`
	Status203 types.Int64 `tfsdk:"status_203"`
	Status204 types.Int64 `tfsdk:"status_204"`
	Status205 types.Int64 `tfsdk:"status_205"`
	Status206 types.Int64 `tfsdk:"status_206"`
	Status207 types.Int64 `tfsdk:"status_207"`
	Status208 types.Int64 `tfsdk:"status_208"`
	Status226 types.Int64 `tfsdk:"status_226"`
	Status300 types.Int64 `tfsdk:"status_300"`
	Status301 types.Int64 `tfsdk:"status_301"`
	Status302 types.Int64 `tfsdk:"status_302"`
	Status303 types.Int64 `tfsdk:"status_303"`
	Status304 types.Int64 `tfsdk:"status_304"`
	Status305 types.Int64 `tfsdk:"status_305"`
	Status306 types.Int64 `tfsdk:"status_306"`
	Status307 types.Int64 `tfsdk:"status_307"`
	Status308 types.Int64 `tfsdk:"status_308"`
	Status400 types.Int64 `tfsdk:"status_400"`
	Status401 types.Int64 `tfsdk:"status_401"`
	Status402 types.Int64 `tfsdk:"status_402"`
	Status403 types.Int64 `tfsdk:"status_403"`
	Status404 types.Int64 `tfsdk:"status_404"`
	Status405 types.Int64 `tfsdk:"status_405"`
	Status406 types.Int64 `tfsdk:"status_406"`
	Status407 types.Int64 `tfsdk:"status_407"`
	Status408 types.Int64 `tfsdk:"status_408"`
	Status409 types.Int64 `tfsdk:"status_409"`
	Status410 types.Int64 `tfsdk:"status_410"`
	Status411 types.Int64 `tfsdk:"status_411"`
	Status412 types.Int64 `tfsdk:"status_412"`
	Status413 types.Int64 `tfsdk:"status_413"`
	Status414 types.Int64 `tfsdk:"status_414"`
	Status415 types.Int64 `tfsdk:"status_415"`
	Status416 types.Int64 `tfsdk:"status_416"`
	Status417 types.Int64 `tfsdk:"status_417"`
	Status421 types.Int64 `tfsdk:"status_421"`
	Status422 types.Int64 `tfsdk:"status_422"`
	Status423 types.Int64 `tfsdk:"status_423"`
	Status424 types.Int64 `tfsdk:"status_424"`
	Status425 types.Int64 `tfsdk:"status_425"`
	Status426 types.Int64 `tfsdk:"status_426"`
	Status428 types.Int64 `tfsdk:"status_428"`
	Status429 types.Int64 `tfsdk:"status_429"`
	Status431 types.Int64 `tfsdk:"status_431"`
	Status451 types.Int64 `tfsdk:"status_451"`
	Status500 types.Int64 `tfsdk:"status_500"`
}

// ToBody
func ResponseOrganizationsGetOrganizationAPIRequestsOverviewItemToBody(state OrganizationsAPIRequestsOverview, response *merakigosdk.ResponseOrganizationsGetOrganizationAPIRequestsOverview) OrganizationsAPIRequestsOverview {
	itemState := ResponseOrganizationsGetOrganizationApiRequestsOverview{
		ResponseCodeCounts: func() *ResponseOrganizationsGetOrganizationApiRequestsOverviewResponseCodeCounts {
			if response.ResponseCodeCounts != nil {
				return &ResponseOrganizationsGetOrganizationApiRequestsOverviewResponseCodeCounts{
					Status200: func() types.Int64 {
						if response.ResponseCodeCounts.Status200 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status200))
						}
						return types.Int64{}
					}(),
					Status201: func() types.Int64 {
						if response.ResponseCodeCounts.Status201 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status201))
						}
						return types.Int64{}
					}(),
					Status202: func() types.Int64 {
						if response.ResponseCodeCounts.Status202 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status202))
						}
						return types.Int64{}
					}(),
					Status203: func() types.Int64 {
						if response.ResponseCodeCounts.Status203 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status203))
						}
						return types.Int64{}
					}(),
					Status204: func() types.Int64 {
						if response.ResponseCodeCounts.Status204 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status204))
						}
						return types.Int64{}
					}(),
					Status205: func() types.Int64 {
						if response.ResponseCodeCounts.Status205 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status205))
						}
						return types.Int64{}
					}(),
					Status206: func() types.Int64 {
						if response.ResponseCodeCounts.Status206 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status206))
						}
						return types.Int64{}
					}(),
					Status207: func() types.Int64 {
						if response.ResponseCodeCounts.Status207 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status207))
						}
						return types.Int64{}
					}(),
					Status208: func() types.Int64 {
						if response.ResponseCodeCounts.Status208 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status208))
						}
						return types.Int64{}
					}(),
					Status226: func() types.Int64 {
						if response.ResponseCodeCounts.Status226 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status226))
						}
						return types.Int64{}
					}(),
					Status300: func() types.Int64 {
						if response.ResponseCodeCounts.Status300 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status300))
						}
						return types.Int64{}
					}(),
					Status301: func() types.Int64 {
						if response.ResponseCodeCounts.Status301 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status301))
						}
						return types.Int64{}
					}(),
					Status302: func() types.Int64 {
						if response.ResponseCodeCounts.Status302 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status302))
						}
						return types.Int64{}
					}(),
					Status303: func() types.Int64 {
						if response.ResponseCodeCounts.Status303 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status303))
						}
						return types.Int64{}
					}(),
					Status304: func() types.Int64 {
						if response.ResponseCodeCounts.Status304 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status304))
						}
						return types.Int64{}
					}(),
					Status305: func() types.Int64 {
						if response.ResponseCodeCounts.Status305 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status305))
						}
						return types.Int64{}
					}(),
					Status306: func() types.Int64 {
						if response.ResponseCodeCounts.Status306 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status306))
						}
						return types.Int64{}
					}(),
					Status307: func() types.Int64 {
						if response.ResponseCodeCounts.Status307 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status307))
						}
						return types.Int64{}
					}(),
					Status308: func() types.Int64 {
						if response.ResponseCodeCounts.Status308 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status308))
						}
						return types.Int64{}
					}(),
					Status400: func() types.Int64 {
						if response.ResponseCodeCounts.Status400 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status400))
						}
						return types.Int64{}
					}(),
					Status401: func() types.Int64 {
						if response.ResponseCodeCounts.Status401 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status401))
						}
						return types.Int64{}
					}(),
					Status402: func() types.Int64 {
						if response.ResponseCodeCounts.Status402 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status402))
						}
						return types.Int64{}
					}(),
					Status403: func() types.Int64 {
						if response.ResponseCodeCounts.Status403 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status403))
						}
						return types.Int64{}
					}(),
					Status404: func() types.Int64 {
						if response.ResponseCodeCounts.Status404 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status404))
						}
						return types.Int64{}
					}(),
					Status405: func() types.Int64 {
						if response.ResponseCodeCounts.Status405 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status405))
						}
						return types.Int64{}
					}(),
					Status406: func() types.Int64 {
						if response.ResponseCodeCounts.Status406 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status406))
						}
						return types.Int64{}
					}(),
					Status407: func() types.Int64 {
						if response.ResponseCodeCounts.Status407 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status407))
						}
						return types.Int64{}
					}(),
					Status408: func() types.Int64 {
						if response.ResponseCodeCounts.Status408 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status408))
						}
						return types.Int64{}
					}(),
					Status409: func() types.Int64 {
						if response.ResponseCodeCounts.Status409 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status409))
						}
						return types.Int64{}
					}(),
					Status410: func() types.Int64 {
						if response.ResponseCodeCounts.Status410 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status410))
						}
						return types.Int64{}
					}(),
					Status411: func() types.Int64 {
						if response.ResponseCodeCounts.Status411 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status411))
						}
						return types.Int64{}
					}(),
					Status412: func() types.Int64 {
						if response.ResponseCodeCounts.Status412 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status412))
						}
						return types.Int64{}
					}(),
					Status413: func() types.Int64 {
						if response.ResponseCodeCounts.Status413 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status413))
						}
						return types.Int64{}
					}(),
					Status414: func() types.Int64 {
						if response.ResponseCodeCounts.Status414 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status414))
						}
						return types.Int64{}
					}(),
					Status415: func() types.Int64 {
						if response.ResponseCodeCounts.Status415 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status415))
						}
						return types.Int64{}
					}(),
					Status416: func() types.Int64 {
						if response.ResponseCodeCounts.Status416 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status416))
						}
						return types.Int64{}
					}(),
					Status417: func() types.Int64 {
						if response.ResponseCodeCounts.Status417 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status417))
						}
						return types.Int64{}
					}(),
					Status421: func() types.Int64 {
						if response.ResponseCodeCounts.Status421 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status421))
						}
						return types.Int64{}
					}(),
					Status422: func() types.Int64 {
						if response.ResponseCodeCounts.Status422 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status422))
						}
						return types.Int64{}
					}(),
					Status423: func() types.Int64 {
						if response.ResponseCodeCounts.Status423 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status423))
						}
						return types.Int64{}
					}(),
					Status424: func() types.Int64 {
						if response.ResponseCodeCounts.Status424 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status424))
						}
						return types.Int64{}
					}(),
					Status425: func() types.Int64 {
						if response.ResponseCodeCounts.Status425 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status425))
						}
						return types.Int64{}
					}(),
					Status426: func() types.Int64 {
						if response.ResponseCodeCounts.Status426 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status426))
						}
						return types.Int64{}
					}(),
					Status428: func() types.Int64 {
						if response.ResponseCodeCounts.Status428 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status428))
						}
						return types.Int64{}
					}(),
					Status429: func() types.Int64 {
						if response.ResponseCodeCounts.Status429 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status429))
						}
						return types.Int64{}
					}(),
					Status431: func() types.Int64 {
						if response.ResponseCodeCounts.Status431 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status431))
						}
						return types.Int64{}
					}(),
					Status451: func() types.Int64 {
						if response.ResponseCodeCounts.Status451 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status451))
						}
						return types.Int64{}
					}(),
					Status500: func() types.Int64 {
						if response.ResponseCodeCounts.Status500 != nil {
							return types.Int64Value(int64(*response.ResponseCodeCounts.Status500))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
