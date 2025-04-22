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
	_ datasource.DataSource              = &OrganizationsAPIRequestsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsAPIRequestsDataSource{}
)

func NewOrganizationsAPIRequestsDataSource() datasource.DataSource {
	return &OrganizationsAPIRequestsDataSource{}
}

type OrganizationsAPIRequestsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsAPIRequestsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsAPIRequestsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_api_requests"
}

func (d *OrganizationsAPIRequestsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"admin_id": schema.StringAttribute{
				MarkdownDescription: `adminId query parameter. Filter the results by the ID of the admin who made the API requests`,
				Optional:            true,
			},
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"method": schema.StringAttribute{
				MarkdownDescription: `method query parameter. Filter the results by the method of the API requests (must be 'GET', 'PUT', 'POST' or 'DELETE')`,
				Optional:            true,
			},
			"operation_ids": schema.ListAttribute{
				MarkdownDescription: `operationIds query parameter. Filter the results by one or more operation IDs for the API request`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"path": schema.StringAttribute{
				MarkdownDescription: `path query parameter. Filter the results by the path of the API requests`,
				Optional:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 50.`,
				Optional:            true,
			},
			"response_code": schema.Int64Attribute{
				MarkdownDescription: `responseCode query parameter. Filter the results by the response code of the API requests`,
				Optional:            true,
			},
			"source_ip": schema.StringAttribute{
				MarkdownDescription: `sourceIp query parameter. Filter the results by the IP address of the originating API request`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
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
			"user_agent": schema.StringAttribute{
				MarkdownDescription: `userAgent query parameter. Filter the results by the user agent string of the API request`,
				Optional:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: `version query parameter. Filter the results by the API version of the API request`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationApiRequests`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"admin_id": schema.StringAttribute{
							MarkdownDescription: `Database ID for the admin user who made the API request.`,
							Computed:            true,
						},
						"client": schema.SingleNestedAttribute{
							MarkdownDescription: `Client information`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `ID for the client which made the request, if applicable.`,
									Computed:            true,
								},
								"type": schema.StringAttribute{
									MarkdownDescription: `Type of client which made the request, if applicable. Available options are: oauth, api_key`,
									Computed:            true,
								},
							},
						},
						"host": schema.StringAttribute{
							MarkdownDescription: `The host which the API request was directed at.`,
							Computed:            true,
						},
						"method": schema.StringAttribute{
							MarkdownDescription: `HTTP method used in the API request.`,
							Computed:            true,
						},
						"operation_id": schema.StringAttribute{
							MarkdownDescription: `Operation ID for the endpoint.`,
							Computed:            true,
						},
						"path": schema.StringAttribute{
							MarkdownDescription: `The API request path.`,
							Computed:            true,
						},
						"query_string": schema.StringAttribute{
							MarkdownDescription: `The query string sent with the API request.`,
							Computed:            true,
						},
						"response_code": schema.Int64Attribute{
							MarkdownDescription: `API request response code.`,
							Computed:            true,
						},
						"source_ip": schema.StringAttribute{
							MarkdownDescription: `Public IP address from which the API request was made.`,
							Computed:            true,
						},
						"ts": schema.StringAttribute{
							MarkdownDescription: `Timestamp, in iso8601 format, indicating when the API request was made.`,
							Computed:            true,
						},
						"user_agent": schema.StringAttribute{
							MarkdownDescription: `The API request user agent.`,
							Computed:            true,
						},
						"version": schema.Int64Attribute{
							MarkdownDescription: `API version of the endpoint.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsAPIRequestsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsAPIRequests OrganizationsAPIRequests
	diags := req.Config.Get(ctx, &organizationsAPIRequests)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationAPIRequests")
		vvOrganizationID := organizationsAPIRequests.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationAPIRequestsQueryParams{}

		queryParams1.T0 = organizationsAPIRequests.T0.ValueString()
		queryParams1.T1 = organizationsAPIRequests.T1.ValueString()
		queryParams1.Timespan = organizationsAPIRequests.Timespan.ValueFloat64()
		queryParams1.PerPage = int(organizationsAPIRequests.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsAPIRequests.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsAPIRequests.EndingBefore.ValueString()
		queryParams1.AdminID = organizationsAPIRequests.AdminID.ValueString()
		queryParams1.Path = organizationsAPIRequests.Path.ValueString()
		queryParams1.Method = organizationsAPIRequests.Method.ValueString()
		queryParams1.ResponseCode = int(organizationsAPIRequests.ResponseCode.ValueInt64())
		queryParams1.SourceIP = organizationsAPIRequests.SourceIP.ValueString()
		queryParams1.UserAgent = organizationsAPIRequests.UserAgent.ValueString()
		queryParams1.Version = int(organizationsAPIRequests.Version.ValueInt64())
		queryParams1.OperationIDs = elementsToStrings(ctx, organizationsAPIRequests.OperationIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationAPIRequests(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAPIRequests",
				err.Error(),
			)
			return
		}

		organizationsAPIRequests = ResponseOrganizationsGetOrganizationAPIRequestsItemsToBody(organizationsAPIRequests, response1)
		diags = resp.State.Set(ctx, &organizationsAPIRequests)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsAPIRequests struct {
	OrganizationID types.String                                           `tfsdk:"organization_id"`
	T0             types.String                                           `tfsdk:"t0"`
	T1             types.String                                           `tfsdk:"t1"`
	Timespan       types.Float64                                          `tfsdk:"timespan"`
	PerPage        types.Int64                                            `tfsdk:"per_page"`
	StartingAfter  types.String                                           `tfsdk:"starting_after"`
	EndingBefore   types.String                                           `tfsdk:"ending_before"`
	AdminID        types.String                                           `tfsdk:"admin_id"`
	Path           types.String                                           `tfsdk:"path"`
	Method         types.String                                           `tfsdk:"method"`
	ResponseCode   types.Int64                                            `tfsdk:"response_code"`
	SourceIP       types.String                                           `tfsdk:"source_ip"`
	UserAgent      types.String                                           `tfsdk:"user_agent"`
	Version        types.Int64                                            `tfsdk:"version"`
	OperationIDs   types.List                                             `tfsdk:"operation_ids"`
	Items          *[]ResponseItemOrganizationsGetOrganizationApiRequests `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationApiRequests struct {
	AdminID      types.String                                               `tfsdk:"admin_id"`
	Client       *ResponseItemOrganizationsGetOrganizationApiRequestsClient `tfsdk:"client"`
	Host         types.String                                               `tfsdk:"host"`
	Method       types.String                                               `tfsdk:"method"`
	OperationID  types.String                                               `tfsdk:"operation_id"`
	Path         types.String                                               `tfsdk:"path"`
	QueryString  types.String                                               `tfsdk:"query_string"`
	ResponseCode types.Int64                                                `tfsdk:"response_code"`
	SourceIP     types.String                                               `tfsdk:"source_ip"`
	Ts           types.String                                               `tfsdk:"ts"`
	UserAgent    types.String                                               `tfsdk:"user_agent"`
	Version      types.Int64                                                `tfsdk:"version"`
}

type ResponseItemOrganizationsGetOrganizationApiRequestsClient struct {
	ID   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

// ToBody
func ResponseOrganizationsGetOrganizationAPIRequestsItemsToBody(state OrganizationsAPIRequests, response *merakigosdk.ResponseOrganizationsGetOrganizationAPIRequests) OrganizationsAPIRequests {
	var items []ResponseItemOrganizationsGetOrganizationApiRequests
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationApiRequests{
			AdminID: types.StringValue(item.AdminID),
			Client: func() *ResponseItemOrganizationsGetOrganizationApiRequestsClient {
				if item.Client != nil {
					return &ResponseItemOrganizationsGetOrganizationApiRequestsClient{
						ID:   types.StringValue(item.Client.ID),
						Type: types.StringValue(item.Client.Type),
					}
				}
				return nil
			}(),
			Host:        types.StringValue(item.Host),
			Method:      types.StringValue(item.Method),
			OperationID: types.StringValue(item.OperationID),
			Path:        types.StringValue(item.Path),
			QueryString: types.StringValue(item.QueryString),
			ResponseCode: func() types.Int64 {
				if item.ResponseCode != nil {
					return types.Int64Value(int64(*item.ResponseCode))
				}
				return types.Int64{}
			}(),
			SourceIP:  types.StringValue(item.SourceIP),
			Ts:        types.StringValue(item.Ts),
			UserAgent: types.StringValue(item.UserAgent),
			Version: func() types.Int64 {
				if item.Version != nil {
					return types.Int64Value(int64(*item.Version))
				}
				return types.Int64{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
