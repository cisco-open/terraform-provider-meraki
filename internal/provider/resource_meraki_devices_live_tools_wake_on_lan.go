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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesLiveToolsWakeOnLanResource{}
	_ resource.ResourceWithConfigure = &DevicesLiveToolsWakeOnLanResource{}
)

func NewDevicesLiveToolsWakeOnLanResource() resource.Resource {
	return &DevicesLiveToolsWakeOnLanResource{}
}

type DevicesLiveToolsWakeOnLanResource struct {
	client *merakigosdk.Client
}

func (r *DevicesLiveToolsWakeOnLanResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesLiveToolsWakeOnLanResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_live_tools_wake_on_lan"
}

func (r *DevicesLiveToolsWakeOnLanResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"callback": schema.SingleNestedAttribute{
				MarkdownDescription: `Details for the callback. Please include either an httpServerId OR url and sharedSecret`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"http_server": schema.SingleNestedAttribute{
						MarkdownDescription: `The webhook receiver used for the callback webhook.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"id": schema.StringAttribute{
								MarkdownDescription: `The webhook receiver ID that will receive information. If specifying this, please leave the url and sharedSecret fields blank.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"payload_template": schema.SingleNestedAttribute{
						MarkdownDescription: `The payload template of the webhook used for the callback`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"id": schema.StringAttribute{
								MarkdownDescription: `The ID of the payload template. Defaults to 'wpt_00005' for the Callback (included) template.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"shared_secret": schema.StringAttribute{
						MarkdownDescription: `A shared secret that will be included in the requests sent to the callback URL. It can be used to verify that the request was sent by Meraki. If using this field, please also specify an url.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"url": schema.StringAttribute{
						MarkdownDescription: `The callback URL for the webhook target. If using this field, please also specify a sharedSecret.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"error": schema.StringAttribute{
				MarkdownDescription: `An error message for a failed execution`,
				Computed:            true,
			},
			"mac": schema.StringAttribute{
				MarkdownDescription: `The target's MAC address`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"request": schema.SingleNestedAttribute{
				MarkdownDescription: `The parameters of the Wake-on-LAN request`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{

					"mac": schema.StringAttribute{
						MarkdownDescription: `The target's MAC address`,
						Computed:            true,
					},
					"serial": schema.StringAttribute{
						MarkdownDescription: `Device serial number`,
						Computed:            true,
					},
					"vlan_id": schema.Int64Attribute{
						MarkdownDescription: `The target's VLAN (1 to 4094)`,
						Computed:            true,
					},
				},
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: `Status of the Wake-on-LAN request
                                  Allowed values: [complete,failed,new,ready,running,scheduled]`,
				Computed: true,
			},
			"url": schema.StringAttribute{
				MarkdownDescription: `GET this url to check the status of your ping request`,
				Computed:            true,
			},
			"vlan_id": schema.Int64Attribute{
				MarkdownDescription: `The target's VLAN (1 to 4094)`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"wake_on_lan_id": schema.StringAttribute{
				MarkdownDescription: `ID of the Wake-on-LAN job`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *DevicesLiveToolsWakeOnLanResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesLiveToolsWakeOnLanRs

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
	vvSerial := data.Serial.ValueString()
	vvWakeOnLanID := data.WakeOnLanID.ValueString()
	//Has Item and not has items

	if vvSerial != "" && vvWakeOnLanID != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Devices.GetDeviceLiveToolsWakeOnLan(vvSerial, vvWakeOnLanID)
		//Has Post
		if err != nil {
			if restyResp1 != nil {
				if restyResp1.StatusCode() != 404 {
					resp.Diagnostics.AddError(
						"Failure when executing GetDeviceLiveToolsWakeOnLan",
						err.Error(),
					)
					return
				}
			}
		}

		if responseVerifyItem != nil {
			data = ResponseDevicesGetDeviceLiveToolsWakeOnLanItemToBodyRs(data, responseVerifyItem, false)
			//Path params in update assigned
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Devices.CreateDeviceLiveToolsWakeOnLan(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateDeviceLiveToolsWakeOnLan",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateDeviceLiveToolsWakeOnLan",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Devices.GetDeviceLiveToolsWakeOnLan(vvSerial, vvWakeOnLanID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceLiveToolsWakeOnLan",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceLiveToolsWakeOnLan",
			err.Error(),
		)
		return
	}

	data = ResponseDevicesGetDeviceLiveToolsWakeOnLanItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *DevicesLiveToolsWakeOnLanResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesLiveToolsWakeOnLanRs

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

	vvSerial := data.Serial.ValueString()
	vvWakeOnLanID := data.WakeOnLanID.ValueString()
	responseGet, restyRespGet, err := r.client.Devices.GetDeviceLiveToolsWakeOnLan(vvSerial, vvWakeOnLanID)
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
				"Failure when executing GetDeviceLiveToolsWakeOnLan",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceLiveToolsWakeOnLan",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseDevicesGetDeviceLiveToolsWakeOnLanItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesLiveToolsWakeOnLanResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("wake_on_lan_id"), idParts[1])...)
}

func (r *DevicesLiveToolsWakeOnLanResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesLiveToolsWakeOnLanRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update
	// No update
	resp.Diagnostics.AddError(
		"Update operation not supported in DevicesLiveToolsWakeOnLan",
		"Update operation not supported in DevicesLiveToolsWakeOnLan",
	)
	return
}

func (r *DevicesLiveToolsWakeOnLanResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting DevicesLiveToolsWakeOnLan", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesLiveToolsWakeOnLanRs struct {
	Serial      types.String                                            `tfsdk:"serial"`
	WakeOnLanID types.String                                            `tfsdk:"wake_on_lan_id"`
	Error       types.String                                            `tfsdk:"error"`
	Request     *ResponseDevicesGetDeviceLiveToolsWakeOnLanRequestRs    `tfsdk:"request"`
	Status      types.String                                            `tfsdk:"status"`
	URL         types.String                                            `tfsdk:"url"`
	Callback    *RequestDevicesCreateDeviceLiveToolsWakeOnLanCallbackRs `tfsdk:"callback"`
	Mac         types.String                                            `tfsdk:"mac"`
	VLANID      types.Int64                                             `tfsdk:"vlan_id"`
}

type ResponseDevicesGetDeviceLiveToolsWakeOnLanRequestRs struct {
	Mac    types.String `tfsdk:"mac"`
	Serial types.String `tfsdk:"serial"`
	VLANID types.Int64  `tfsdk:"vlan_id"`
}

type RequestDevicesCreateDeviceLiveToolsWakeOnLanCallbackRs struct {
	HTTPServer      *RequestDevicesCreateDeviceLiveToolsWakeOnLanCallbackHttpServerRs      `tfsdk:"http_server"`
	PayloadTemplate *RequestDevicesCreateDeviceLiveToolsWakeOnLanCallbackPayloadTemplateRs `tfsdk:"payload_template"`
	SharedSecret    types.String                                                           `tfsdk:"shared_secret"`
	URL             types.String                                                           `tfsdk:"url"`
}

type RequestDevicesCreateDeviceLiveToolsWakeOnLanCallbackHttpServerRs struct {
	ID types.String `tfsdk:"id"`
}

type RequestDevicesCreateDeviceLiveToolsWakeOnLanCallbackPayloadTemplateRs struct {
	ID types.String `tfsdk:"id"`
}

// FromBody
func (r *DevicesLiveToolsWakeOnLanRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestDevicesCreateDeviceLiveToolsWakeOnLan {
	emptyString := ""
	var requestDevicesCreateDeviceLiveToolsWakeOnLanCallback *merakigosdk.RequestDevicesCreateDeviceLiveToolsWakeOnLanCallback

	if r.Callback != nil {
		var requestDevicesCreateDeviceLiveToolsWakeOnLanCallbackHTTPServer *merakigosdk.RequestDevicesCreateDeviceLiveToolsWakeOnLanCallbackHTTPServer

		if r.Callback.HTTPServer != nil {
			id := r.Callback.HTTPServer.ID.ValueString()
			requestDevicesCreateDeviceLiveToolsWakeOnLanCallbackHTTPServer = &merakigosdk.RequestDevicesCreateDeviceLiveToolsWakeOnLanCallbackHTTPServer{
				ID: id,
			}
			//[debug] Is Array: False
		}
		var requestDevicesCreateDeviceLiveToolsWakeOnLanCallbackPayloadTemplate *merakigosdk.RequestDevicesCreateDeviceLiveToolsWakeOnLanCallbackPayloadTemplate

		if r.Callback.PayloadTemplate != nil {
			id := r.Callback.PayloadTemplate.ID.ValueString()
			requestDevicesCreateDeviceLiveToolsWakeOnLanCallbackPayloadTemplate = &merakigosdk.RequestDevicesCreateDeviceLiveToolsWakeOnLanCallbackPayloadTemplate{
				ID: id,
			}
			//[debug] Is Array: False
		}
		sharedSecret := r.Callback.SharedSecret.ValueString()
		url := r.Callback.URL.ValueString()
		requestDevicesCreateDeviceLiveToolsWakeOnLanCallback = &merakigosdk.RequestDevicesCreateDeviceLiveToolsWakeOnLanCallback{
			HTTPServer:      requestDevicesCreateDeviceLiveToolsWakeOnLanCallbackHTTPServer,
			PayloadTemplate: requestDevicesCreateDeviceLiveToolsWakeOnLanCallbackPayloadTemplate,
			SharedSecret:    sharedSecret,
			URL:             url,
		}
		//[debug] Is Array: False
	}
	mac := new(string)
	if !r.Mac.IsUnknown() && !r.Mac.IsNull() {
		*mac = r.Mac.ValueString()
	} else {
		mac = &emptyString
	}
	vLANID := new(int64)
	if !r.VLANID.IsUnknown() && !r.VLANID.IsNull() {
		*vLANID = r.VLANID.ValueInt64()
	} else {
		vLANID = nil
	}
	out := merakigosdk.RequestDevicesCreateDeviceLiveToolsWakeOnLan{
		Callback: requestDevicesCreateDeviceLiveToolsWakeOnLanCallback,
		Mac:      *mac,
		VLANID:   int64ToIntPointer(vLANID),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseDevicesGetDeviceLiveToolsWakeOnLanItemToBodyRs(state DevicesLiveToolsWakeOnLanRs, response *merakigosdk.ResponseDevicesGetDeviceLiveToolsWakeOnLan, is_read bool) DevicesLiveToolsWakeOnLanRs {
	itemState := DevicesLiveToolsWakeOnLanRs{
		Error: types.StringValue(response.Error),
		Request: func() *ResponseDevicesGetDeviceLiveToolsWakeOnLanRequestRs {
			if response.Request != nil {
				return &ResponseDevicesGetDeviceLiveToolsWakeOnLanRequestRs{
					Mac:    types.StringValue(response.Request.Mac),
					Serial: types.StringValue(response.Request.Serial),
					VLANID: func() types.Int64 {
						if response.Request.VLANID != nil {
							return types.Int64Value(int64(*response.Request.VLANID))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
		Status:      types.StringValue(response.Status),
		URL:         types.StringValue(response.URL),
		WakeOnLanID: types.StringValue(response.WakeOnLanID),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesLiveToolsWakeOnLanRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesLiveToolsWakeOnLanRs)
}
