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

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"strings"

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWebhooksPayloadTemplatesResource{}
	_ resource.ResourceWithConfigure = &NetworksWebhooksPayloadTemplatesResource{}
)

func NewNetworksWebhooksPayloadTemplatesResource() resource.Resource {
	return &NetworksWebhooksPayloadTemplatesResource{}
}

type NetworksWebhooksPayloadTemplatesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWebhooksPayloadTemplatesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWebhooksPayloadTemplatesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_webhooks_payload_templates"
}

func (r *NetworksWebhooksPayloadTemplatesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"body": schema.StringAttribute{
				MarkdownDescription: `The body of the payload template, in liquid template`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"body_file": schema.StringAttribute{
				MarkdownDescription: `A Base64 encoded file containing liquid template used for the body of the webhook message. Either **body** or **bodyFile** must be specified.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"headers": schema.SetNestedAttribute{
				MarkdownDescription: `The payload template headers, will be rendered as a key-value pair in the webhook.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the header attribute`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"template": schema.StringAttribute{
							MarkdownDescription: `The value returned in the header attribute, in liquid template`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"headers_file": schema.StringAttribute{
				MarkdownDescription: `A Base64 encoded file containing the liquid template used with the webhook headers.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name of the payload template`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"payload_template_id": schema.StringAttribute{
				MarkdownDescription: `Webhook payload template Id`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
	}
}

//path params to set ['payloadTemplateId']

func (r *NetworksWebhooksPayloadTemplatesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWebhooksPayloadTemplatesRs

	var item types.Object
	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Has Paths
	vvNetworkID := data.NetworkID.ValueString()
	//Has Item and has items and post

	vvName := data.Name.ValueString()

	responseVerifyItem, restyResp1, err := r.client.Networks.GetNetworkWebhooksPayloadTemplates(vvNetworkID)
	//Has Post
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkWebhooksPayloadTemplates",
					err.Error(),
				)
				return
			}
		}
	}

	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvPayloadTemplateID, ok := result2["PayloadTemplateID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter PayloadTemplateID",
					"Fail Parsing PayloadTemplateID",
				)
				return
			}
			r.client.Networks.UpdateNetworkWebhooksPayloadTemplate(vvNetworkID, vvPayloadTemplateID, data.toSdkApiRequestUpdate(ctx))

			responseVerifyItem2, _, _ := r.client.Networks.GetNetworkWebhooksPayloadTemplate(vvNetworkID, vvPayloadTemplateID)
			if responseVerifyItem2 != nil {
				data = ResponseNetworksGetNetworkWebhooksPayloadTemplateItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Networks.CreateNetworkWebhooksPayloadTemplate(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkWebhooksPayloadTemplate",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkWebhooksPayloadTemplate",
			err.Error(),
		)
		return
	}

	responseGet, restyResp1, err := r.client.Networks.GetNetworkWebhooksPayloadTemplates(vvNetworkID)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWebhooksPayloadTemplates",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWebhooksPayloadTemplates",
			err.Error(),
		)
		return
	}

	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvPayloadTemplateID, ok := result2["PayloadTemplateID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter PayloadTemplateID",
				"Fail Parsing PayloadTemplateID",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Networks.GetNetworkWebhooksPayloadTemplate(vvNetworkID, vvPayloadTemplateID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseNetworksGetNetworkWebhooksPayloadTemplateItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkWebhooksPayloadTemplate",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWebhooksPayloadTemplate",
				err.Error(),
			)
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"Error in result.",
			"Error in result.",
		)
		return
	}

}

func (r *NetworksWebhooksPayloadTemplatesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWebhooksPayloadTemplatesRs

	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	vvPayloadTemplateID := data.PayloadTemplateID.ValueString()
	responseGet, restyRespGet, err := r.client.Networks.GetNetworkWebhooksPayloadTemplate(vvNetworkID, vvPayloadTemplateID)
	if err != nil || restyRespGet == nil {
		if restyRespGet != nil {
			if restyRespGet.StatusCode() == 404 {
				resp.Diagnostics.AddWarning(
					"Resource not found",
					"Deleting resource",
				)
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWebhooksPayloadTemplate",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWebhooksPayloadTemplate",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseNetworksGetNetworkWebhooksPayloadTemplateItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksWebhooksPayloadTemplatesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("payload_template_id"), idParts[1])...)
}

func (r *NetworksWebhooksPayloadTemplatesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksWebhooksPayloadTemplatesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvPayloadTemplateID := data.PayloadTemplateID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkWebhooksPayloadTemplate(vvNetworkID, vvPayloadTemplateID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWebhooksPayloadTemplate",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWebhooksPayloadTemplate",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWebhooksPayloadTemplatesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksWebhooksPayloadTemplatesRs
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &state, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}

	vvNetworkID := state.NetworkID.ValueString()
	vvPayloadTemplateID := state.PayloadTemplateID.ValueString()
	_, err := r.client.Networks.DeleteNetworkWebhooksPayloadTemplate(vvNetworkID, vvPayloadTemplateID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkWebhooksPayloadTemplate", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksWebhooksPayloadTemplatesRs struct {
	NetworkID         types.String                                                  `tfsdk:"network_id"`
	PayloadTemplateID types.String                                                  `tfsdk:"payload_template_id"`
	Body              types.String                                                  `tfsdk:"body"`
	Headers           *[]ResponseNetworksGetNetworkWebhooksPayloadTemplateHeadersRs `tfsdk:"headers"`
	Name              types.String                                                  `tfsdk:"name"`
	Sharing           *ResponseNetworksGetNetworkWebhooksPayloadTemplateSharingRs   `tfsdk:"sharing"`
	Type              types.String                                                  `tfsdk:"type"`
	BodyFile          types.String                                                  `tfsdk:"body_file"`
	HeadersFile       types.String                                                  `tfsdk:"headers_file"`
}

type ResponseNetworksGetNetworkWebhooksPayloadTemplateHeadersRs struct {
	Name     types.String `tfsdk:"name"`
	Template types.String `tfsdk:"template"`
}

type ResponseNetworksGetNetworkWebhooksPayloadTemplateSharingRs struct {
	ByNetwork *ResponseNetworksGetNetworkWebhooksPayloadTemplateSharingByNetworkRs `tfsdk:"by_network"`
}

type ResponseNetworksGetNetworkWebhooksPayloadTemplateSharingByNetworkRs struct {
	AdminsCanModify types.Bool `tfsdk:"admins_can_modify"`
}

// FromBody
func (r *NetworksWebhooksPayloadTemplatesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksCreateNetworkWebhooksPayloadTemplate {
	emptyString := ""
	body := new(string)
	if !r.Body.IsUnknown() && !r.Body.IsNull() {
		*body = r.Body.ValueString()
	} else {
		body = &emptyString
	}
	bodyFile := new(string)
	if !r.BodyFile.IsUnknown() && !r.BodyFile.IsNull() {
		*bodyFile = r.BodyFile.ValueString()
	} else {
		bodyFile = &emptyString
	}
	var requestNetworksCreateNetworkWebhooksPayloadTemplateHeaders []merakigosdk.RequestNetworksCreateNetworkWebhooksPayloadTemplateHeaders

	if r.Headers != nil {
		for _, rItem1 := range *r.Headers {
			name := rItem1.Name.ValueString()
			template := rItem1.Template.ValueString()
			requestNetworksCreateNetworkWebhooksPayloadTemplateHeaders = append(requestNetworksCreateNetworkWebhooksPayloadTemplateHeaders, merakigosdk.RequestNetworksCreateNetworkWebhooksPayloadTemplateHeaders{
				Name:     name,
				Template: template,
			})
			//[debug] Is Array: True
		}
	}
	headersFile := new(string)
	if !r.HeadersFile.IsUnknown() && !r.HeadersFile.IsNull() {
		*headersFile = r.HeadersFile.ValueString()
	} else {
		headersFile = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	out := merakigosdk.RequestNetworksCreateNetworkWebhooksPayloadTemplate{
		Body:     *body,
		BodyFile: *bodyFile,
		Headers: func() *[]merakigosdk.RequestNetworksCreateNetworkWebhooksPayloadTemplateHeaders {
			if len(requestNetworksCreateNetworkWebhooksPayloadTemplateHeaders) > 0 {
				return &requestNetworksCreateNetworkWebhooksPayloadTemplateHeaders
			}
			return nil
		}(),
		HeadersFile: *headersFile,
		Name:        *name,
	}
	return &out
}
func (r *NetworksWebhooksPayloadTemplatesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestNetworksUpdateNetworkWebhooksPayloadTemplate {
	emptyString := ""
	body := new(string)
	if !r.Body.IsUnknown() && !r.Body.IsNull() {
		*body = r.Body.ValueString()
	} else {
		body = &emptyString
	}
	bodyFile := new(string)
	if !r.BodyFile.IsUnknown() && !r.BodyFile.IsNull() {
		*bodyFile = r.BodyFile.ValueString()
	} else {
		bodyFile = &emptyString
	}
	var requestNetworksUpdateNetworkWebhooksPayloadTemplateHeaders []merakigosdk.RequestNetworksUpdateNetworkWebhooksPayloadTemplateHeaders

	if r.Headers != nil {
		for _, rItem1 := range *r.Headers {
			name := rItem1.Name.ValueString()
			template := rItem1.Template.ValueString()
			requestNetworksUpdateNetworkWebhooksPayloadTemplateHeaders = append(requestNetworksUpdateNetworkWebhooksPayloadTemplateHeaders, merakigosdk.RequestNetworksUpdateNetworkWebhooksPayloadTemplateHeaders{
				Name:     name,
				Template: template,
			})
			//[debug] Is Array: True
		}
	}
	headersFile := new(string)
	if !r.HeadersFile.IsUnknown() && !r.HeadersFile.IsNull() {
		*headersFile = r.HeadersFile.ValueString()
	} else {
		headersFile = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	out := merakigosdk.RequestNetworksUpdateNetworkWebhooksPayloadTemplate{
		Body:     *body,
		BodyFile: *bodyFile,
		Headers: func() *[]merakigosdk.RequestNetworksUpdateNetworkWebhooksPayloadTemplateHeaders {
			if len(requestNetworksUpdateNetworkWebhooksPayloadTemplateHeaders) > 0 {
				return &requestNetworksUpdateNetworkWebhooksPayloadTemplateHeaders
			}
			return nil
		}(),
		HeadersFile: *headersFile,
		Name:        *name,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseNetworksGetNetworkWebhooksPayloadTemplateItemToBodyRs(state NetworksWebhooksPayloadTemplatesRs, response *merakigosdk.ResponseNetworksGetNetworkWebhooksPayloadTemplate, is_read bool) NetworksWebhooksPayloadTemplatesRs {
	itemState := NetworksWebhooksPayloadTemplatesRs{
		Body: types.StringValue(response.Body),
		Headers: func() *[]ResponseNetworksGetNetworkWebhooksPayloadTemplateHeadersRs {
			if response.Headers != nil {
				result := make([]ResponseNetworksGetNetworkWebhooksPayloadTemplateHeadersRs, len(*response.Headers))
				for i, headers := range *response.Headers {
					result[i] = ResponseNetworksGetNetworkWebhooksPayloadTemplateHeadersRs{
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
		Sharing: func() *ResponseNetworksGetNetworkWebhooksPayloadTemplateSharingRs {
			if response.Sharing != nil {
				return &ResponseNetworksGetNetworkWebhooksPayloadTemplateSharingRs{
					ByNetwork: func() *ResponseNetworksGetNetworkWebhooksPayloadTemplateSharingByNetworkRs {
						if response.Sharing.ByNetwork != nil {
							return &ResponseNetworksGetNetworkWebhooksPayloadTemplateSharingByNetworkRs{
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
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWebhooksPayloadTemplatesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWebhooksPayloadTemplatesRs)
}
