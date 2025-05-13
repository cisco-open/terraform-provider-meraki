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

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWebhooksHTTPServersResource{}
	_ resource.ResourceWithConfigure = &NetworksWebhooksHTTPServersResource{}
)

func NewNetworksWebhooksHTTPServersResource() resource.Resource {
	return &NetworksWebhooksHTTPServersResource{}
}

type NetworksWebhooksHTTPServersResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWebhooksHTTPServersResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWebhooksHTTPServersResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_webhooks_http_servers"
}

func (r *NetworksWebhooksHTTPServersResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"http_server_id": schema.StringAttribute{
				MarkdownDescription: `httpServerId path parameter. Http server ID`,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: `A Base64 encoded ID.`,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `A name for easy reference to the HTTP server`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `A Meraki network ID.`,
				Required:            true,
			},
			"payload_template": schema.SingleNestedAttribute{
				MarkdownDescription: `The payload template to use when posting data to the HTTP server.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"name": schema.StringAttribute{
						MarkdownDescription: `The name of the payload template.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
							SuppressDiffString(),
						},
					},
					"payload_template_id": schema.StringAttribute{
						MarkdownDescription: `The ID of the payload template.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"shared_secret": schema.StringAttribute{
				MarkdownDescription: `A shared secret that will be included in POSTs sent to the HTTP server. This secret can be used to verify that the request was sent by Meraki.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"url": schema.StringAttribute{
				MarkdownDescription: `The URL of the HTTP server.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SuppressDiffString(),
				},
			},
		},
	}
}

//path params to set ['httpServerId']
//path params to assign NOT EDITABLE ['url']

func (r *NetworksWebhooksHTTPServersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWebhooksHTTPServersRs

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

	responseVerifyItem, restyResp1, err := r.client.Networks.GetNetworkWebhooksHTTPServers(vvNetworkID)
	//Has Post
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkWebhooksHTTPServers",
					restyResp1.String(),
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
			vvHTTPServerID, ok := result2["ID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter HTTPServerID",
					"Fail Parsing HTTPServerID",
				)
				return
			}
			r.client.Networks.UpdateNetworkWebhooksHTTPServer(vvNetworkID, vvHTTPServerID, data.toSdkApiRequestUpdate(ctx))

			responseVerifyItem2, _, _ := r.client.Networks.GetNetworkWebhooksHTTPServer(vvNetworkID, vvHTTPServerID)
			if responseVerifyItem2 != nil {
				data = ResponseNetworksGetNetworkWebhooksHTTPServerItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Networks.CreateNetworkWebhooksHTTPServer(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkWebhooksHTTPServer",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkWebhooksHTTPServer",
			err.Error(),
		)
		return
	}

	responseGet, restyResp1, err := r.client.Networks.GetNetworkWebhooksHTTPServers(vvNetworkID)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWebhooksHTTPServers",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWebhooksHTTPServers",
			err.Error(),
		)
		return
	}

	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvHTTPServerID, ok := result2["ID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter HTTPServerID",
				"Fail Parsing HTTPServerID",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Networks.GetNetworkWebhooksHTTPServer(vvNetworkID, vvHTTPServerID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseNetworksGetNetworkWebhooksHTTPServerItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkWebhooksHTTPServer",
					restyRespGet.String(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWebhooksHTTPServer",
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

func (r *NetworksWebhooksHTTPServersResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWebhooksHTTPServersRs

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
	vvHTTPServerID := data.HTTPServerID.ValueString()
	responseGet, restyRespGet, err := r.client.Networks.GetNetworkWebhooksHTTPServer(vvNetworkID, vvHTTPServerID)
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
				"Failure when executing GetNetworkWebhooksHTTPServer",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWebhooksHTTPServer",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseNetworksGetNetworkWebhooksHTTPServerItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksWebhooksHTTPServersResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("http_server_id"), idParts[1])...)
}

func (r *NetworksWebhooksHTTPServersResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksWebhooksHTTPServersRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvHTTPServerID := data.HTTPServerID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkWebhooksHTTPServer(vvNetworkID, vvHTTPServerID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWebhooksHTTPServer",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWebhooksHTTPServer",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWebhooksHTTPServersResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksWebhooksHTTPServersRs
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
	vvHTTPServerID := state.HTTPServerID.ValueString()
	_, err := r.client.Networks.DeleteNetworkWebhooksHTTPServer(vvNetworkID, vvHTTPServerID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkWebhooksHTTPServer", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksWebhooksHTTPServersRs struct {
	NetworkID       types.String                                                   `tfsdk:"network_id"`
	HTTPServerID    types.String                                                   `tfsdk:"http_server_id"`
	ID              types.String                                                   `tfsdk:"id"`
	Name            types.String                                                   `tfsdk:"name"`
	PayloadTemplate *ResponseNetworksGetNetworkWebhooksHttpServerPayloadTemplateRs `tfsdk:"payload_template"`
	URL             types.String                                                   `tfsdk:"url"`
	SharedSecret    types.String                                                   `tfsdk:"shared_secret"`
}

type ResponseNetworksGetNetworkWebhooksHttpServerPayloadTemplateRs struct {
	Name              types.String `tfsdk:"name"`
	PayloadTemplateID types.String `tfsdk:"payload_template_id"`
}

// FromBody
func (r *NetworksWebhooksHTTPServersRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksCreateNetworkWebhooksHTTPServer {
	emptyString := ""
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestNetworksCreateNetworkWebhooksHTTPServerPayloadTemplate *merakigosdk.RequestNetworksCreateNetworkWebhooksHTTPServerPayloadTemplate

	if r.PayloadTemplate != nil {
		name := r.PayloadTemplate.Name.ValueString()
		payloadTemplateID := r.PayloadTemplate.PayloadTemplateID.ValueString()
		requestNetworksCreateNetworkWebhooksHTTPServerPayloadTemplate = &merakigosdk.RequestNetworksCreateNetworkWebhooksHTTPServerPayloadTemplate{
			Name:              name,
			PayloadTemplateID: payloadTemplateID,
		}
		//[debug] Is Array: False
	}
	sharedSecret := new(string)
	if !r.SharedSecret.IsUnknown() && !r.SharedSecret.IsNull() {
		*sharedSecret = r.SharedSecret.ValueString()
	} else {
		sharedSecret = &emptyString
	}
	uRL := new(string)
	if !r.URL.IsUnknown() && !r.URL.IsNull() {
		*uRL = r.URL.ValueString()
	} else {
		uRL = &emptyString
	}
	out := merakigosdk.RequestNetworksCreateNetworkWebhooksHTTPServer{
		Name:            *name,
		PayloadTemplate: requestNetworksCreateNetworkWebhooksHTTPServerPayloadTemplate,
		SharedSecret:    *sharedSecret,
		URL:             *uRL,
	}
	return &out
}
func (r *NetworksWebhooksHTTPServersRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestNetworksUpdateNetworkWebhooksHTTPServer {
	emptyString := ""
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestNetworksUpdateNetworkWebhooksHTTPServerPayloadTemplate *merakigosdk.RequestNetworksUpdateNetworkWebhooksHTTPServerPayloadTemplate

	if r.PayloadTemplate != nil {
		payloadTemplateID := r.PayloadTemplate.PayloadTemplateID.ValueString()
		requestNetworksUpdateNetworkWebhooksHTTPServerPayloadTemplate = &merakigosdk.RequestNetworksUpdateNetworkWebhooksHTTPServerPayloadTemplate{
			PayloadTemplateID: payloadTemplateID,
		}
		//[debug] Is Array: False
	}
	sharedSecret := new(string)
	if !r.SharedSecret.IsUnknown() && !r.SharedSecret.IsNull() {
		*sharedSecret = r.SharedSecret.ValueString()
	} else {
		sharedSecret = &emptyString
	}
	out := merakigosdk.RequestNetworksUpdateNetworkWebhooksHTTPServer{
		Name:            *name,
		PayloadTemplate: requestNetworksUpdateNetworkWebhooksHTTPServerPayloadTemplate,
		SharedSecret:    *sharedSecret,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseNetworksGetNetworkWebhooksHTTPServerItemToBodyRs(state NetworksWebhooksHTTPServersRs, response *merakigosdk.ResponseNetworksGetNetworkWebhooksHTTPServer, is_read bool) NetworksWebhooksHTTPServersRs {
	itemState := NetworksWebhooksHTTPServersRs{
		ID:        types.StringValue(response.ID),
		Name:      types.StringValue(response.Name),
		NetworkID: types.StringValue(response.NetworkID),
		PayloadTemplate: func() *ResponseNetworksGetNetworkWebhooksHttpServerPayloadTemplateRs {
			if response.PayloadTemplate != nil {
				return &ResponseNetworksGetNetworkWebhooksHttpServerPayloadTemplateRs{
					Name:              types.StringValue(response.PayloadTemplate.Name),
					PayloadTemplateID: types.StringValue(response.PayloadTemplate.PayloadTemplateID),
				}
			}
			return nil
		}(),
		URL:          types.StringValue(response.URL),
		SharedSecret: state.SharedSecret,
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWebhooksHTTPServersRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWebhooksHTTPServersRs)
}
