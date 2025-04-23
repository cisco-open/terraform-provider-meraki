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

// RESOURCE ACTION

import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesLiveToolsLedsBlinkResource{}
	_ resource.ResourceWithConfigure = &DevicesLiveToolsLedsBlinkResource{}
)

func NewDevicesLiveToolsLedsBlinkResource() resource.Resource {
	return &DevicesLiveToolsLedsBlinkResource{}
}

type DevicesLiveToolsLedsBlinkResource struct {
	client *merakigosdk.Client
}

func (r *DevicesLiveToolsLedsBlinkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesLiveToolsLedsBlinkResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_live_tools_leds_blink"
}

// resourceAction
func (r *DevicesLiveToolsLedsBlinkResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"callback": schema.SingleNestedAttribute{
						MarkdownDescription: `Information for callback used to send back results`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"id": schema.StringAttribute{
								MarkdownDescription: `The ID of the callback. To check the status of the callback, use this ID in a request to /webhooks/callbacks/statuses/{id}`,
								Computed:            true,
							},
							"status": schema.StringAttribute{
								MarkdownDescription: `The status of the callback`,
								Computed:            true,
							},
							"url": schema.StringAttribute{
								MarkdownDescription: `The callback URL for the webhook target. This was either provided in the original request or comes from a configured webhook receiver`,
								Computed:            true,
							},
						},
					},
					"error": schema.StringAttribute{
						MarkdownDescription: `An error message for a failed Blink LEDs execution, if present`,
						Computed:            true,
					},
					"leds_blink_id": schema.StringAttribute{
						MarkdownDescription: `ID of led blink job`,
						Computed:            true,
					},
					"request": schema.SingleNestedAttribute{
						MarkdownDescription: `The parameters of the leds blink request`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"duration": schema.Int64Attribute{
								MarkdownDescription: `The duration to blink leds in seconds`,
								Computed:            true,
							},
							"serial": schema.StringAttribute{
								MarkdownDescription: `Device serial number`,
								Computed:            true,
							},
						},
					},
					"status": schema.StringAttribute{
						MarkdownDescription: `Status of the leds blink request
                                          Allowed values: [complete,failed,new,ready,running,scheduled]`,
						Computed: true,
					},
					"url": schema.StringAttribute{
						MarkdownDescription: `GET this url to check the status of your leds blink request`,
						Computed:            true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"callback": schema.SingleNestedAttribute{
						MarkdownDescription: `Details for the callback. Please include either an httpServerId OR url and sharedSecret`,
						Optional:            true,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"http_server": schema.SingleNestedAttribute{
								MarkdownDescription: `The webhook receiver used for the callback webhook.`,
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"id": schema.StringAttribute{
										MarkdownDescription: `The webhook receiver ID that will receive information. If specifying this, please leave the url and sharedSecret fields blank.`,
										Optional:            true,
										Computed:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
								},
							},
							"payload_template": schema.SingleNestedAttribute{
								MarkdownDescription: `The payload template of the webhook used for the callback`,
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"id": schema.StringAttribute{
										MarkdownDescription: `The ID of the payload template. Defaults to 'wpt_00005' for the Callback (included) template.`,
										Optional:            true,
										Computed:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
								},
							},
							"shared_secret": schema.StringAttribute{
								MarkdownDescription: `A shared secret that will be included in the requests sent to the callback URL. It can be used to verify that the request was sent by Meraki. If using this field, please also specify an url.`,
								Optional:            true,
								Computed:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
							},
							"url": schema.StringAttribute{
								MarkdownDescription: `The callback URL for the webhook target. If using this field, please also specify a sharedSecret.`,
								Optional:            true,
								Computed:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
							},
						},
					},
					"duration": schema.Int64Attribute{
						MarkdownDescription: `The duration in seconds to blink LEDs.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplace(),
						},
					},
				},
			},
		},
	}
}
func (r *DevicesLiveToolsLedsBlinkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesLiveToolsLedsBlink

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
	//Has Paths
	vvSerial := data.Serial.ValueString()
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Devices.CreateDeviceLiveToolsLedsBlink(vvSerial, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateDeviceLiveToolsLedsBlink",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateDeviceLiveToolsLedsBlink",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseDevicesCreateDeviceLiveToolsLedsBlinkItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesLiveToolsLedsBlinkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *DevicesLiveToolsLedsBlinkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *DevicesLiveToolsLedsBlinkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesLiveToolsLedsBlink struct {
	Serial     types.String                                    `tfsdk:"serial"`
	Item       *ResponseDevicesCreateDeviceLiveToolsLedsBlink  `tfsdk:"item"`
	Parameters *RequestDevicesCreateDeviceLiveToolsLedsBlinkRs `tfsdk:"parameters"`
}

type ResponseDevicesCreateDeviceLiveToolsLedsBlink struct {
	Callback    *ResponseDevicesCreateDeviceLiveToolsLedsBlinkCallback `tfsdk:"callback"`
	Error       types.String                                           `tfsdk:"error"`
	LedsBlinkID types.String                                           `tfsdk:"leds_blink_id"`
	Request     *ResponseDevicesCreateDeviceLiveToolsLedsBlinkRequest  `tfsdk:"request"`
	Status      types.String                                           `tfsdk:"status"`
	URL         types.String                                           `tfsdk:"url"`
}

type ResponseDevicesCreateDeviceLiveToolsLedsBlinkCallback struct {
	ID     types.String `tfsdk:"id"`
	Status types.String `tfsdk:"status"`
	URL    types.String `tfsdk:"url"`
}

type ResponseDevicesCreateDeviceLiveToolsLedsBlinkRequest struct {
	Duration types.Int64  `tfsdk:"duration"`
	Serial   types.String `tfsdk:"serial"`
}

type RequestDevicesCreateDeviceLiveToolsLedsBlinkRs struct {
	Callback *RequestDevicesCreateDeviceLiveToolsLedsBlinkCallbackRs `tfsdk:"callback"`
	Duration types.Int64                                             `tfsdk:"duration"`
}

type RequestDevicesCreateDeviceLiveToolsLedsBlinkCallbackRs struct {
	HTTPServer      *RequestDevicesCreateDeviceLiveToolsLedsBlinkCallbackHttpServerRs      `tfsdk:"http_server"`
	PayloadTemplate *RequestDevicesCreateDeviceLiveToolsLedsBlinkCallbackPayloadTemplateRs `tfsdk:"payload_template"`
	SharedSecret    types.String                                                           `tfsdk:"shared_secret"`
	URL             types.String                                                           `tfsdk:"url"`
}

type RequestDevicesCreateDeviceLiveToolsLedsBlinkCallbackHttpServerRs struct {
	ID types.String `tfsdk:"id"`
}

type RequestDevicesCreateDeviceLiveToolsLedsBlinkCallbackPayloadTemplateRs struct {
	ID types.String `tfsdk:"id"`
}

// FromBody
func (r *DevicesLiveToolsLedsBlink) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestDevicesCreateDeviceLiveToolsLedsBlink {
	re := *r.Parameters
	var requestDevicesCreateDeviceLiveToolsLedsBlinkCallback *merakigosdk.RequestDevicesCreateDeviceLiveToolsLedsBlinkCallback

	if re.Callback != nil {
		var requestDevicesCreateDeviceLiveToolsLedsBlinkCallbackHTTPServer *merakigosdk.RequestDevicesCreateDeviceLiveToolsLedsBlinkCallbackHTTPServer

		if re.Callback.HTTPServer != nil {
			id := re.Callback.HTTPServer.ID.ValueString()
			requestDevicesCreateDeviceLiveToolsLedsBlinkCallbackHTTPServer = &merakigosdk.RequestDevicesCreateDeviceLiveToolsLedsBlinkCallbackHTTPServer{
				ID: id,
			}
			//[debug] Is Array: False
		}
		var requestDevicesCreateDeviceLiveToolsLedsBlinkCallbackPayloadTemplate *merakigosdk.RequestDevicesCreateDeviceLiveToolsLedsBlinkCallbackPayloadTemplate

		if re.Callback.PayloadTemplate != nil {
			id := re.Callback.PayloadTemplate.ID.ValueString()
			requestDevicesCreateDeviceLiveToolsLedsBlinkCallbackPayloadTemplate = &merakigosdk.RequestDevicesCreateDeviceLiveToolsLedsBlinkCallbackPayloadTemplate{
				ID: id,
			}
			//[debug] Is Array: False
		}
		sharedSecret := re.Callback.SharedSecret.ValueString()
		url := re.Callback.URL.ValueString()
		requestDevicesCreateDeviceLiveToolsLedsBlinkCallback = &merakigosdk.RequestDevicesCreateDeviceLiveToolsLedsBlinkCallback{
			HTTPServer:      requestDevicesCreateDeviceLiveToolsLedsBlinkCallbackHTTPServer,
			PayloadTemplate: requestDevicesCreateDeviceLiveToolsLedsBlinkCallbackPayloadTemplate,
			SharedSecret:    sharedSecret,
			URL:             url,
		}
		//[debug] Is Array: False
	}
	duration := new(int64)
	if !re.Duration.IsUnknown() && !re.Duration.IsNull() {
		*duration = re.Duration.ValueInt64()
	} else {
		duration = nil
	}
	out := merakigosdk.RequestDevicesCreateDeviceLiveToolsLedsBlink{
		Callback: requestDevicesCreateDeviceLiveToolsLedsBlinkCallback,
		Duration: int64ToIntPointer(duration),
	}
	return &out
}

// ToBody
func ResponseDevicesCreateDeviceLiveToolsLedsBlinkItemToBody(state DevicesLiveToolsLedsBlink, response *merakigosdk.ResponseDevicesCreateDeviceLiveToolsLedsBlink) DevicesLiveToolsLedsBlink {
	itemState := ResponseDevicesCreateDeviceLiveToolsLedsBlink{
		Callback: func() *ResponseDevicesCreateDeviceLiveToolsLedsBlinkCallback {
			if response.Callback != nil {
				return &ResponseDevicesCreateDeviceLiveToolsLedsBlinkCallback{
					ID:     types.StringValue(response.Callback.ID),
					Status: types.StringValue(response.Callback.Status),
					URL:    types.StringValue(response.Callback.URL),
				}
			}
			return nil
		}(),
		Error:       types.StringValue(response.Error),
		LedsBlinkID: types.StringValue(response.LedsBlinkID),
		Request: func() *ResponseDevicesCreateDeviceLiveToolsLedsBlinkRequest {
			if response.Request != nil {
				return &ResponseDevicesCreateDeviceLiveToolsLedsBlinkRequest{
					Duration: func() types.Int64 {
						if response.Request.Duration != nil {
							return types.Int64Value(int64(*response.Request.Duration))
						}
						return types.Int64{}
					}(),
					Serial: types.StringValue(response.Request.Serial),
				}
			}
			return nil
		}(),
		Status: types.StringValue(response.Status),
		URL:    types.StringValue(response.URL),
	}
	state.Item = &itemState
	return state
}
