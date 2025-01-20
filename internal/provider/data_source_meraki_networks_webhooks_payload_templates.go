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
	_ datasource.DataSource              = &NetworksWebhooksPayloadTemplatesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWebhooksPayloadTemplatesDataSource{}
)

func NewNetworksWebhooksPayloadTemplatesDataSource() datasource.DataSource {
	return &NetworksWebhooksPayloadTemplatesDataSource{}
}

type NetworksWebhooksPayloadTemplatesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWebhooksPayloadTemplatesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWebhooksPayloadTemplatesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_webhooks_payload_templates"
}

func (d *NetworksWebhooksPayloadTemplatesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"payload_template_id": schema.StringAttribute{
				MarkdownDescription: `payloadTemplateId path parameter. Payload template ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"body": schema.StringAttribute{
						MarkdownDescription: `The body of the payload template, in liquid template`,
						Computed:            true,
					},
					"headers": schema.SetNestedAttribute{
						MarkdownDescription: `The payload template headers, will be rendered as a key-value pair in the webhook.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"name": schema.StringAttribute{
									MarkdownDescription: `The name of the header attribute`,
									Computed:            true,
								},
								"template": schema.StringAttribute{
									MarkdownDescription: `The value returned in the header attribute, in liquid template`,
									Computed:            true,
								},
							},
						},
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name of the payload template`,
						Computed:            true,
					},
					"payload_template_id": schema.StringAttribute{
						MarkdownDescription: `Webhook payload template Id`,
						Computed:            true,
					},
					"sharing": schema.SingleNestedAttribute{
						MarkdownDescription: `Information on which entities have access to the template`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"by_network": schema.SingleNestedAttribute{
								MarkdownDescription: `Information on network access to the template`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"admins_can_modify": schema.BoolAttribute{
										MarkdownDescription: `Indicates whether network admins may modify this template`,
										Computed:            true,
									},
								},
							},
						},
					},
					"type": schema.StringAttribute{
						MarkdownDescription: `The type of the payload template`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseNetworksGetNetworkWebhooksPayloadTemplates`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"body": schema.StringAttribute{
							MarkdownDescription: `The body of the payload template, in liquid template`,
							Computed:            true,
						},
						"headers": schema.SetNestedAttribute{
							MarkdownDescription: `The payload template headers, will be rendered as a key-value pair in the webhook.`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"name": schema.StringAttribute{
										MarkdownDescription: `The name of the header attribute`,
										Computed:            true,
									},
									"template": schema.StringAttribute{
										MarkdownDescription: `The value returned in the header attribute, in liquid template`,
										Computed:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the payload template`,
							Computed:            true,
						},
						"payload_template_id": schema.StringAttribute{
							MarkdownDescription: `Webhook payload template Id`,
							Computed:            true,
						},
						"sharing": schema.SingleNestedAttribute{
							MarkdownDescription: `Information on which entities have access to the template`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"by_network": schema.SingleNestedAttribute{
									MarkdownDescription: `Information on network access to the template`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"admins_can_modify": schema.BoolAttribute{
											MarkdownDescription: `Indicates whether network admins may modify this template`,
											Computed:            true,
										},
									},
								},
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `The type of the payload template`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksWebhooksPayloadTemplatesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWebhooksPayloadTemplates NetworksWebhooksPayloadTemplates
	diags := req.Config.Get(ctx, &networksWebhooksPayloadTemplates)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksWebhooksPayloadTemplates.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksWebhooksPayloadTemplates.NetworkID.IsNull(), !networksWebhooksPayloadTemplates.PayloadTemplateID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWebhooksPayloadTemplates")
		vvNetworkID := networksWebhooksPayloadTemplates.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkWebhooksPayloadTemplates(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWebhooksPayloadTemplates",
				err.Error(),
			)
			return
		}

		networksWebhooksPayloadTemplates = ResponseNetworksGetNetworkWebhooksPayloadTemplatesItemsToBody(networksWebhooksPayloadTemplates, response1)
		diags = resp.State.Set(ctx, &networksWebhooksPayloadTemplates)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkWebhooksPayloadTemplate")
		vvNetworkID := networksWebhooksPayloadTemplates.NetworkID.ValueString()
		vvPayloadTemplateID := networksWebhooksPayloadTemplates.PayloadTemplateID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Networks.GetNetworkWebhooksPayloadTemplate(vvNetworkID, vvPayloadTemplateID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWebhooksPayloadTemplate",
				err.Error(),
			)
			return
		}

		networksWebhooksPayloadTemplates = ResponseNetworksGetNetworkWebhooksPayloadTemplateItemToBody(networksWebhooksPayloadTemplates, response2)
		diags = resp.State.Set(ctx, &networksWebhooksPayloadTemplates)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWebhooksPayloadTemplates struct {
	NetworkID         types.String                                              `tfsdk:"network_id"`
	PayloadTemplateID types.String                                              `tfsdk:"payload_template_id"`
	Items             *[]ResponseItemNetworksGetNetworkWebhooksPayloadTemplates `tfsdk:"items"`
	Item              *ResponseNetworksGetNetworkWebhooksPayloadTemplate        `tfsdk:"item"`
}

type ResponseItemNetworksGetNetworkWebhooksPayloadTemplates struct {
	Body              types.String                                                     `tfsdk:"body"`
	Headers           *[]ResponseItemNetworksGetNetworkWebhooksPayloadTemplatesHeaders `tfsdk:"headers"`
	Name              types.String                                                     `tfsdk:"name"`
	PayloadTemplateID types.String                                                     `tfsdk:"payload_template_id"`
	Sharing           *ResponseItemNetworksGetNetworkWebhooksPayloadTemplatesSharing   `tfsdk:"sharing"`
	Type              types.String                                                     `tfsdk:"type"`
}

type ResponseItemNetworksGetNetworkWebhooksPayloadTemplatesHeaders struct {
	Name     types.String `tfsdk:"name"`
	Template types.String `tfsdk:"template"`
}

type ResponseItemNetworksGetNetworkWebhooksPayloadTemplatesSharing struct {
	ByNetwork *ResponseItemNetworksGetNetworkWebhooksPayloadTemplatesSharingByNetwork `tfsdk:"by_network"`
}

type ResponseItemNetworksGetNetworkWebhooksPayloadTemplatesSharingByNetwork struct {
	AdminsCanModify types.Bool `tfsdk:"admins_can_modify"`
}

type ResponseNetworksGetNetworkWebhooksPayloadTemplate struct {
	Body              types.String                                                `tfsdk:"body"`
	Headers           *[]ResponseNetworksGetNetworkWebhooksPayloadTemplateHeaders `tfsdk:"headers"`
	Name              types.String                                                `tfsdk:"name"`
	PayloadTemplateID types.String                                                `tfsdk:"payload_template_id"`
	Sharing           *ResponseNetworksGetNetworkWebhooksPayloadTemplateSharing   `tfsdk:"sharing"`
	Type              types.String                                                `tfsdk:"type"`
}

type ResponseNetworksGetNetworkWebhooksPayloadTemplateHeaders struct {
	Name     types.String `tfsdk:"name"`
	Template types.String `tfsdk:"template"`
}

type ResponseNetworksGetNetworkWebhooksPayloadTemplateSharing struct {
	ByNetwork *ResponseNetworksGetNetworkWebhooksPayloadTemplateSharingByNetwork `tfsdk:"by_network"`
}

type ResponseNetworksGetNetworkWebhooksPayloadTemplateSharingByNetwork struct {
	AdminsCanModify types.Bool `tfsdk:"admins_can_modify"`
}

// ToBody
func ResponseNetworksGetNetworkWebhooksPayloadTemplatesItemsToBody(state NetworksWebhooksPayloadTemplates, response *merakigosdk.ResponseNetworksGetNetworkWebhooksPayloadTemplates) NetworksWebhooksPayloadTemplates {
	var items []ResponseItemNetworksGetNetworkWebhooksPayloadTemplates
	for _, item := range *response {
		itemState := ResponseItemNetworksGetNetworkWebhooksPayloadTemplates{
			Body: types.StringValue(item.Body),
			Headers: func() *[]ResponseItemNetworksGetNetworkWebhooksPayloadTemplatesHeaders {
				if item.Headers != nil {
					result := make([]ResponseItemNetworksGetNetworkWebhooksPayloadTemplatesHeaders, len(*item.Headers))
					for i, headers := range *item.Headers {
						result[i] = ResponseItemNetworksGetNetworkWebhooksPayloadTemplatesHeaders{
							Name:     types.StringValue(headers.Name),
							Template: types.StringValue(headers.Template),
						}
					}
					return &result
				}
				return nil
			}(),
			Name:              types.StringValue(item.Name),
			PayloadTemplateID: types.StringValue(item.PayloadTemplateID),
			Sharing: func() *ResponseItemNetworksGetNetworkWebhooksPayloadTemplatesSharing {
				if item.Sharing != nil {
					return &ResponseItemNetworksGetNetworkWebhooksPayloadTemplatesSharing{
						ByNetwork: func() *ResponseItemNetworksGetNetworkWebhooksPayloadTemplatesSharingByNetwork {
							if item.Sharing.ByNetwork != nil {
								return &ResponseItemNetworksGetNetworkWebhooksPayloadTemplatesSharingByNetwork{
									AdminsCanModify: func() types.Bool {
										if item.Sharing.ByNetwork.AdminsCanModify != nil {
											return types.BoolValue(*item.Sharing.ByNetwork.AdminsCanModify)
										}
										return types.Bool{}
									}(),
								}
							}
							return nil
						}(),
					}
				}
				return nil
			}(),
			Type: types.StringValue(item.Type),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseNetworksGetNetworkWebhooksPayloadTemplateItemToBody(state NetworksWebhooksPayloadTemplates, response *merakigosdk.ResponseNetworksGetNetworkWebhooksPayloadTemplate) NetworksWebhooksPayloadTemplates {
	itemState := ResponseNetworksGetNetworkWebhooksPayloadTemplate{
		Body: types.StringValue(response.Body),
		Headers: func() *[]ResponseNetworksGetNetworkWebhooksPayloadTemplateHeaders {
			if response.Headers != nil {
				result := make([]ResponseNetworksGetNetworkWebhooksPayloadTemplateHeaders, len(*response.Headers))
				for i, headers := range *response.Headers {
					result[i] = ResponseNetworksGetNetworkWebhooksPayloadTemplateHeaders{
						Name:     types.StringValue(headers.Name),
						Template: types.StringValue(headers.Template),
					}
				}
				return &result
			}
			return nil
		}(),
		Name:              types.StringValue(response.Name),
		PayloadTemplateID: types.StringValue(response.PayloadTemplateID),
		Sharing: func() *ResponseNetworksGetNetworkWebhooksPayloadTemplateSharing {
			if response.Sharing != nil {
				return &ResponseNetworksGetNetworkWebhooksPayloadTemplateSharing{
					ByNetwork: func() *ResponseNetworksGetNetworkWebhooksPayloadTemplateSharingByNetwork {
						if response.Sharing.ByNetwork != nil {
							return &ResponseNetworksGetNetworkWebhooksPayloadTemplateSharingByNetwork{
								AdminsCanModify: func() types.Bool {
									if response.Sharing.ByNetwork.AdminsCanModify != nil {
										return types.BoolValue(*response.Sharing.ByNetwork.AdminsCanModify)
									}
									return types.Bool{}
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		Type: types.StringValue(response.Type),
	}
	state.Item = &itemState
	return state
}
