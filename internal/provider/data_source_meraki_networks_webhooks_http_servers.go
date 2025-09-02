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
	_ datasource.DataSource              = &NetworksWebhooksHTTPServersDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWebhooksHTTPServersDataSource{}
)

func NewNetworksWebhooksHTTPServersDataSource() datasource.DataSource {
	return &NetworksWebhooksHTTPServersDataSource{}
}

type NetworksWebhooksHTTPServersDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWebhooksHTTPServersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWebhooksHTTPServersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_webhooks_http_servers"
}

func (d *NetworksWebhooksHTTPServersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"http_server_id": schema.StringAttribute{
				MarkdownDescription: `httpServerId path parameter. Http server ID`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"id": schema.StringAttribute{
						MarkdownDescription: `A Base64 encoded ID.`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `A name for easy reference to the HTTP server`,
						Computed:            true,
					},
					"network_id": schema.StringAttribute{
						MarkdownDescription: `A Meraki network ID.`,
						Computed:            true,
					},
					"payload_template": schema.SingleNestedAttribute{
						MarkdownDescription: `The payload template to use when posting data to the HTTP server.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"name": schema.StringAttribute{
								MarkdownDescription: `The name of the payload template.`,
								Computed:            true,
							},
							"payload_template_id": schema.StringAttribute{
								MarkdownDescription: `The ID of the payload template.`,
								Computed:            true,
							},
						},
					},
					"url": schema.StringAttribute{
						MarkdownDescription: `The URL of the HTTP server.`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseNetworksGetNetworkWebhooksHttpServers`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"id": schema.StringAttribute{
							MarkdownDescription: `A Base64 encoded ID.`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `A name for easy reference to the HTTP server`,
							Computed:            true,
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: `A Meraki network ID.`,
							Computed:            true,
						},
						"payload_template": schema.SingleNestedAttribute{
							MarkdownDescription: `The payload template to use when posting data to the HTTP server.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"name": schema.StringAttribute{
									MarkdownDescription: `The name of the payload template.`,
									Computed:            true,
								},
								"payload_template_id": schema.StringAttribute{
									MarkdownDescription: `The ID of the payload template.`,
									Computed:            true,
								},
							},
						},
						"url": schema.StringAttribute{
							MarkdownDescription: `The URL of the HTTP server.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksWebhooksHTTPServersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWebhooksHTTPServers NetworksWebhooksHTTPServers
	diags := req.Config.Get(ctx, &networksWebhooksHTTPServers)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksWebhooksHTTPServers.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksWebhooksHTTPServers.NetworkID.IsNull(), !networksWebhooksHTTPServers.HTTPServerID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWebhooksHTTPServers")
		vvNetworkID := networksWebhooksHTTPServers.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkWebhooksHTTPServers(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWebhooksHTTPServers",
				err.Error(),
			)
			return
		}

		networksWebhooksHTTPServers = ResponseNetworksGetNetworkWebhooksHTTPServersItemsToBody(networksWebhooksHTTPServers, response1)
		diags = resp.State.Set(ctx, &networksWebhooksHTTPServers)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkWebhooksHTTPServer")
		vvNetworkID := networksWebhooksHTTPServers.NetworkID.ValueString()
		vvHTTPServerID := networksWebhooksHTTPServers.HTTPServerID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Networks.GetNetworkWebhooksHTTPServer(vvNetworkID, vvHTTPServerID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWebhooksHTTPServer",
				err.Error(),
			)
			return
		}

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWebhooksHTTPServer",
				err.Error(),
			)
			return
		}

		networksWebhooksHTTPServers = ResponseNetworksGetNetworkWebhooksHTTPServerItemToBody(networksWebhooksHTTPServers, response2)
		diags = resp.State.Set(ctx, &networksWebhooksHTTPServers)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWebhooksHTTPServers struct {
	NetworkID    types.String                                         `tfsdk:"network_id"`
	HTTPServerID types.String                                         `tfsdk:"http_server_id"`
	Items        *[]ResponseItemNetworksGetNetworkWebhooksHttpServers `tfsdk:"items"`
	Item         *ResponseNetworksGetNetworkWebhooksHttpServer        `tfsdk:"item"`
}

type ResponseItemNetworksGetNetworkWebhooksHttpServers struct {
	ID              types.String                                                      `tfsdk:"id"`
	Name            types.String                                                      `tfsdk:"name"`
	NetworkID       types.String                                                      `tfsdk:"network_id"`
	PayloadTemplate *ResponseItemNetworksGetNetworkWebhooksHttpServersPayloadTemplate `tfsdk:"payload_template"`
	URL             types.String                                                      `tfsdk:"url"`
}

type ResponseItemNetworksGetNetworkWebhooksHttpServersPayloadTemplate struct {
	Name              types.String `tfsdk:"name"`
	PayloadTemplateID types.String `tfsdk:"payload_template_id"`
}

type ResponseNetworksGetNetworkWebhooksHttpServer struct {
	ID              types.String                                                 `tfsdk:"id"`
	Name            types.String                                                 `tfsdk:"name"`
	NetworkID       types.String                                                 `tfsdk:"network_id"`
	PayloadTemplate *ResponseNetworksGetNetworkWebhooksHttpServerPayloadTemplate `tfsdk:"payload_template"`
	URL             types.String                                                 `tfsdk:"url"`
}

type ResponseNetworksGetNetworkWebhooksHttpServerPayloadTemplate struct {
	Name              types.String `tfsdk:"name"`
	PayloadTemplateID types.String `tfsdk:"payload_template_id"`
}

// ToBody
func ResponseNetworksGetNetworkWebhooksHTTPServersItemsToBody(state NetworksWebhooksHTTPServers, response *merakigosdk.ResponseNetworksGetNetworkWebhooksHTTPServers) NetworksWebhooksHTTPServers {
	var items []ResponseItemNetworksGetNetworkWebhooksHttpServers
	for _, item := range *response {
		itemState := ResponseItemNetworksGetNetworkWebhooksHttpServers{
			ID: func() types.String {
				if item.ID != "" {
					return types.StringValue(item.ID)
				}
				return types.String{}
			}(),
			Name: func() types.String {
				if item.Name != "" {
					return types.StringValue(item.Name)
				}
				return types.String{}
			}(),
			NetworkID: func() types.String {
				if item.NetworkID != "" {
					return types.StringValue(item.NetworkID)
				}
				return types.String{}
			}(),
			PayloadTemplate: func() *ResponseItemNetworksGetNetworkWebhooksHttpServersPayloadTemplate {
				if item.PayloadTemplate != nil {
					return &ResponseItemNetworksGetNetworkWebhooksHttpServersPayloadTemplate{
						Name: func() types.String {
							if item.PayloadTemplate.Name != "" {
								return types.StringValue(item.PayloadTemplate.Name)
							}
							return types.String{}
						}(),
						PayloadTemplateID: func() types.String {
							if item.PayloadTemplate.PayloadTemplateID != "" {
								return types.StringValue(item.PayloadTemplate.PayloadTemplateID)
							}
							return types.String{}
						}(),
					}
				}
				return nil
			}(),
			URL: func() types.String {
				if item.URL != "" {
					return types.StringValue(item.URL)
				}
				return types.String{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseNetworksGetNetworkWebhooksHTTPServerItemToBody(state NetworksWebhooksHTTPServers, response *merakigosdk.ResponseNetworksGetNetworkWebhooksHTTPServer) NetworksWebhooksHTTPServers {
	itemState := ResponseNetworksGetNetworkWebhooksHttpServer{
		ID: func() types.String {
			if response.ID != "" {
				return types.StringValue(response.ID)
			}
			return types.String{}
		}(),
		Name: func() types.String {
			if response.Name != "" {
				return types.StringValue(response.Name)
			}
			return types.String{}
		}(),
		NetworkID: func() types.String {
			if response.NetworkID != "" {
				return types.StringValue(response.NetworkID)
			}
			return types.String{}
		}(),
		PayloadTemplate: func() *ResponseNetworksGetNetworkWebhooksHttpServerPayloadTemplate {
			if response.PayloadTemplate != nil {
				return &ResponseNetworksGetNetworkWebhooksHttpServerPayloadTemplate{
					Name: func() types.String {
						if response.PayloadTemplate.Name != "" {
							return types.StringValue(response.PayloadTemplate.Name)
						}
						return types.String{}
					}(),
					PayloadTemplateID: func() types.String {
						if response.PayloadTemplate.PayloadTemplateID != "" {
							return types.StringValue(response.PayloadTemplate.PayloadTemplateID)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
		URL: func() types.String {
			if response.URL != "" {
				return types.StringValue(response.URL)
			}
			return types.String{}
		}(),
	}
	state.Item = &itemState
	return state
}
