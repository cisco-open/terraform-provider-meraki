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
	_ datasource.DataSource              = &OrganizationsWebhooksCallbacksStatusesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWebhooksCallbacksStatusesDataSource{}
)

func NewOrganizationsWebhooksCallbacksStatusesDataSource() datasource.DataSource {
	return &OrganizationsWebhooksCallbacksStatusesDataSource{}
}

type OrganizationsWebhooksCallbacksStatusesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWebhooksCallbacksStatusesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWebhooksCallbacksStatusesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_webhooks_callbacks_statuses"
}

func (d *OrganizationsWebhooksCallbacksStatusesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"callback_id": schema.StringAttribute{
				MarkdownDescription: `callbackId path parameter. Callback ID`,
				Required:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"callback_id": schema.StringAttribute{
						MarkdownDescription: `The ID of the callback`,
						Computed:            true,
					},
					"created_by": schema.SingleNestedAttribute{
						MarkdownDescription: `Information around who initiated the callback`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"admin_id": schema.StringAttribute{
								MarkdownDescription: `The ID of the user who initiated the callback`,
								Computed:            true,
							},
						},
					},
					"errors": schema.ListAttribute{
						MarkdownDescription: `The errors returned by the callback`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"status": schema.StringAttribute{
						MarkdownDescription: `The status of the callback`,
						Computed:            true,
					},
					"webhook": schema.SingleNestedAttribute{
						MarkdownDescription: `The webhook receiver used by the callback to send results`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"http_server": schema.SingleNestedAttribute{
								MarkdownDescription: `The webhook receiver used for the callback webhook`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"id": schema.StringAttribute{
										MarkdownDescription: `The webhook receiver ID that will receive information`,
										Computed:            true,
									},
								},
							},
							"payload_template": schema.SingleNestedAttribute{
								MarkdownDescription: `The payload template of the webhook used for the callback`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"id": schema.StringAttribute{
										MarkdownDescription: `The ID of the payload template`,
										Computed:            true,
									},
								},
							},
							"sent_at": schema.StringAttribute{
								MarkdownDescription: `The timestamp the callback was sent to the webhook receiver`,
								Computed:            true,
							},
							"url": schema.StringAttribute{
								MarkdownDescription: `The webhook receiver URL where the callback will be sent`,
								Computed:            true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsWebhooksCallbacksStatusesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWebhooksCallbacksStatuses OrganizationsWebhooksCallbacksStatuses
	diags := req.Config.Get(ctx, &organizationsWebhooksCallbacksStatuses)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWebhooksCallbacksStatus")
		vvOrganizationID := organizationsWebhooksCallbacksStatuses.OrganizationID.ValueString()
		vvCallbackID := organizationsWebhooksCallbacksStatuses.CallbackID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationWebhooksCallbacksStatus(vvOrganizationID, vvCallbackID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWebhooksCallbacksStatus",
				err.Error(),
			)
			return
		}

		organizationsWebhooksCallbacksStatuses = ResponseOrganizationsGetOrganizationWebhooksCallbacksStatusItemToBody(organizationsWebhooksCallbacksStatuses, response1)
		diags = resp.State.Set(ctx, &organizationsWebhooksCallbacksStatuses)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWebhooksCallbacksStatuses struct {
	OrganizationID types.String                                                 `tfsdk:"organization_id"`
	CallbackID     types.String                                                 `tfsdk:"callback_id"`
	Item           *ResponseOrganizationsGetOrganizationWebhooksCallbacksStatus `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationWebhooksCallbacksStatus struct {
	CallbackID types.String                                                          `tfsdk:"callback_id"`
	CreatedBy  *ResponseOrganizationsGetOrganizationWebhooksCallbacksStatusCreatedBy `tfsdk:"created_by"`
	Errors     types.List                                                            `tfsdk:"errors"`
	Status     types.String                                                          `tfsdk:"status"`
	Webhook    *ResponseOrganizationsGetOrganizationWebhooksCallbacksStatusWebhook   `tfsdk:"webhook"`
}

type ResponseOrganizationsGetOrganizationWebhooksCallbacksStatusCreatedBy struct {
	AdminID types.String `tfsdk:"admin_id"`
}

type ResponseOrganizationsGetOrganizationWebhooksCallbacksStatusWebhook struct {
	HTTPServer      *ResponseOrganizationsGetOrganizationWebhooksCallbacksStatusWebhookHttpServer      `tfsdk:"http_server"`
	PayloadTemplate *ResponseOrganizationsGetOrganizationWebhooksCallbacksStatusWebhookPayloadTemplate `tfsdk:"payload_template"`
	SentAt          types.String                                                                       `tfsdk:"sent_at"`
	URL             types.String                                                                       `tfsdk:"url"`
}

type ResponseOrganizationsGetOrganizationWebhooksCallbacksStatusWebhookHttpServer struct {
	ID types.String `tfsdk:"id"`
}

type ResponseOrganizationsGetOrganizationWebhooksCallbacksStatusWebhookPayloadTemplate struct {
	ID types.String `tfsdk:"id"`
}

// ToBody
func ResponseOrganizationsGetOrganizationWebhooksCallbacksStatusItemToBody(state OrganizationsWebhooksCallbacksStatuses, response *merakigosdk.ResponseOrganizationsGetOrganizationWebhooksCallbacksStatus) OrganizationsWebhooksCallbacksStatuses {
	itemState := ResponseOrganizationsGetOrganizationWebhooksCallbacksStatus{
		CallbackID: types.StringValue(response.CallbackID),
		CreatedBy: func() *ResponseOrganizationsGetOrganizationWebhooksCallbacksStatusCreatedBy {
			if response.CreatedBy != nil {
				return &ResponseOrganizationsGetOrganizationWebhooksCallbacksStatusCreatedBy{
					AdminID: types.StringValue(response.CreatedBy.AdminID),
				}
			}
			return nil
		}(),
		Errors: StringSliceToList(response.Errors),
		Status: types.StringValue(response.Status),
		Webhook: func() *ResponseOrganizationsGetOrganizationWebhooksCallbacksStatusWebhook {
			if response.Webhook != nil {
				return &ResponseOrganizationsGetOrganizationWebhooksCallbacksStatusWebhook{
					HTTPServer: func() *ResponseOrganizationsGetOrganizationWebhooksCallbacksStatusWebhookHttpServer {
						if response.Webhook.HTTPServer != nil {
							return &ResponseOrganizationsGetOrganizationWebhooksCallbacksStatusWebhookHttpServer{
								ID: types.StringValue(response.Webhook.HTTPServer.ID),
							}
						}
						return nil
					}(),
					PayloadTemplate: func() *ResponseOrganizationsGetOrganizationWebhooksCallbacksStatusWebhookPayloadTemplate {
						if response.Webhook.PayloadTemplate != nil {
							return &ResponseOrganizationsGetOrganizationWebhooksCallbacksStatusWebhookPayloadTemplate{
								ID: types.StringValue(response.Webhook.PayloadTemplate.ID),
							}
						}
						return nil
					}(),
					SentAt: types.StringValue(response.Webhook.SentAt),
					URL:    types.StringValue(response.Webhook.URL),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
