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
	_ datasource.DataSource              = &OrganizationsWebhooksLogsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWebhooksLogsDataSource{}
)

func NewOrganizationsWebhooksLogsDataSource() datasource.DataSource {
	return &OrganizationsWebhooksLogsDataSource{}
}

type OrganizationsWebhooksLogsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWebhooksLogsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWebhooksLogsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_webhooks_logs"
}

func (d *OrganizationsWebhooksLogsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 50.`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 90 days from today.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 31 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.`,
				Optional:            true,
			},
			"url": schema.StringAttribute{
				MarkdownDescription: `url query parameter. The URL the webhook was sent to`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationWebhooksLogs`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"alert_type": schema.StringAttribute{
							MarkdownDescription: `Type of alert that the webhook is delivering`,
							Computed:            true,
						},
						"logged_at": schema.StringAttribute{
							MarkdownDescription: `When the webhook log was created, in ISO8601 format`,
							Computed:            true,
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: `Network ID for the webhook log`,
							Computed:            true,
						},
						"organization_id": schema.StringAttribute{
							MarkdownDescription: `ID for the webhook log's organization`,
							Computed:            true,
						},
						"response_code": schema.Int64Attribute{
							MarkdownDescription: `Response code from the webhook`,
							Computed:            true,
						},
						"response_duration": schema.Int64Attribute{
							MarkdownDescription: `Duration of the response, in milliseconds`,
							Computed:            true,
						},
						"sent_at": schema.StringAttribute{
							MarkdownDescription: `When the webhook was sent, in ISO8601 format`,
							Computed:            true,
						},
						"url": schema.StringAttribute{
							MarkdownDescription: `URL where the webhook was sent`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsWebhooksLogsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWebhooksLogs OrganizationsWebhooksLogs
	diags := req.Config.Get(ctx, &organizationsWebhooksLogs)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWebhooksLogs")
		vvOrganizationID := organizationsWebhooksLogs.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWebhooksLogsQueryParams{}

		queryParams1.T0 = organizationsWebhooksLogs.T0.ValueString()
		queryParams1.T1 = organizationsWebhooksLogs.T1.ValueString()
		queryParams1.Timespan = organizationsWebhooksLogs.Timespan.ValueFloat64()
		queryParams1.PerPage = int(organizationsWebhooksLogs.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWebhooksLogs.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWebhooksLogs.EndingBefore.ValueString()
		queryParams1.URL = organizationsWebhooksLogs.URL.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationWebhooksLogs(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWebhooksLogs",
				err.Error(),
			)
			return
		}

		organizationsWebhooksLogs = ResponseOrganizationsGetOrganizationWebhooksLogsItemsToBody(organizationsWebhooksLogs, response1)
		diags = resp.State.Set(ctx, &organizationsWebhooksLogs)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWebhooksLogs struct {
	OrganizationID types.String                                            `tfsdk:"organization_id"`
	T0             types.String                                            `tfsdk:"t0"`
	T1             types.String                                            `tfsdk:"t1"`
	Timespan       types.Float64                                           `tfsdk:"timespan"`
	PerPage        types.Int64                                             `tfsdk:"per_page"`
	StartingAfter  types.String                                            `tfsdk:"starting_after"`
	EndingBefore   types.String                                            `tfsdk:"ending_before"`
	URL            types.String                                            `tfsdk:"url"`
	Items          *[]ResponseItemOrganizationsGetOrganizationWebhooksLogs `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationWebhooksLogs struct {
	AlertType        types.String `tfsdk:"alert_type"`
	LoggedAt         types.String `tfsdk:"logged_at"`
	NetworkID        types.String `tfsdk:"network_id"`
	OrganizationID   types.String `tfsdk:"organization_id"`
	ResponseCode     types.Int64  `tfsdk:"response_code"`
	ResponseDuration types.Int64  `tfsdk:"response_duration"`
	SentAt           types.String `tfsdk:"sent_at"`
	URL              types.String `tfsdk:"url"`
}

// ToBody
func ResponseOrganizationsGetOrganizationWebhooksLogsItemsToBody(state OrganizationsWebhooksLogs, response *merakigosdk.ResponseOrganizationsGetOrganizationWebhooksLogs) OrganizationsWebhooksLogs {
	var items []ResponseItemOrganizationsGetOrganizationWebhooksLogs
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationWebhooksLogs{
			AlertType:      types.StringValue(item.AlertType),
			LoggedAt:       types.StringValue(item.LoggedAt),
			NetworkID:      types.StringValue(item.NetworkID),
			OrganizationID: types.StringValue(item.OrganizationID),
			ResponseCode: func() types.Int64 {
				if item.ResponseCode != nil {
					return types.Int64Value(int64(*item.ResponseCode))
				}
				return types.Int64{}
			}(),
			ResponseDuration: func() types.Int64 {
				if item.ResponseDuration != nil {
					return types.Int64Value(int64(*item.ResponseDuration))
				}
				return types.Int64{}
			}(),
			SentAt: types.StringValue(item.SentAt),
			URL:    types.StringValue(item.URL),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
