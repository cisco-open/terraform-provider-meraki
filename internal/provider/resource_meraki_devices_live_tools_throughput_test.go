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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesLiveToolsThroughputTestResource{}
	_ resource.ResourceWithConfigure = &DevicesLiveToolsThroughputTestResource{}
)

func NewDevicesLiveToolsThroughputTestResource() resource.Resource {
	return &DevicesLiveToolsThroughputTestResource{}
}

type DevicesLiveToolsThroughputTestResource struct {
	client *merakigosdk.Client
}

func (r *DevicesLiveToolsThroughputTestResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesLiveToolsThroughputTestResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_live_tools_throughput_test"
}

func (r *DevicesLiveToolsThroughputTestResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				MarkdownDescription: `Description of the error.`,
				Computed:            true,
			},
			"request": schema.SingleNestedAttribute{
				MarkdownDescription: `The parameters of the throughput test request`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{

					"serial": schema.StringAttribute{
						MarkdownDescription: `Device serial number`,
						Computed:            true,
					},
				},
			},
			"result": schema.SingleNestedAttribute{
				MarkdownDescription: `Result of the throughput test request`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{

					"speeds": schema.SingleNestedAttribute{
						MarkdownDescription: `Shows the speeds (Mbps)`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"downstream": schema.Float64Attribute{
								MarkdownDescription: `Shows the download speed from shard (Mbps)`,
								Computed:            true,
							},
						},
					},
				},
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: `Status of the throughput test request
                                  Allowed values: [complete,failed,new,ready,running,scheduled]`,
				Computed: true,
			},
			"throughput_test_id": schema.StringAttribute{
				MarkdownDescription: `ID of throughput test job`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"url": schema.StringAttribute{
				MarkdownDescription: `GET this url to check the status of your throughput test request`,
				Computed:            true,
			},
		},
	}
}

func (r *DevicesLiveToolsThroughputTestResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesLiveToolsThroughputTestRs

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
	vvThroughputTestID := data.ThroughputTestID.ValueString()
	//Has Item and not has items

	if vvSerial != "" && vvThroughputTestID != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Devices.GetDeviceLiveToolsThroughputTest(vvSerial, vvThroughputTestID)
		//Has Post
		if err != nil {
			if restyResp1 != nil {
				if restyResp1.StatusCode() != 404 {
					resp.Diagnostics.AddError(
						"Failure when executing GetDeviceLiveToolsThroughputTest",
						err.Error(),
					)
					return
				}
			}
		}

		if responseVerifyItem != nil {
			data = ResponseDevicesGetDeviceLiveToolsThroughputTestItemToBodyRs(data, responseVerifyItem, false)
			//Path params in update assigned
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Devices.CreateDeviceLiveToolsThroughputTest(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateDeviceLiveToolsThroughputTest",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateDeviceLiveToolsThroughputTest",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Devices.GetDeviceLiveToolsThroughputTest(vvSerial, vvThroughputTestID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceLiveToolsThroughputTest",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceLiveToolsThroughputTest",
			err.Error(),
		)
		return
	}

	data = ResponseDevicesGetDeviceLiveToolsThroughputTestItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *DevicesLiveToolsThroughputTestResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesLiveToolsThroughputTestRs

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
	vvThroughputTestID := data.ThroughputTestID.ValueString()
	responseGet, restyRespGet, err := r.client.Devices.GetDeviceLiveToolsThroughputTest(vvSerial, vvThroughputTestID)
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
				"Failure when executing GetDeviceLiveToolsThroughputTest",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceLiveToolsThroughputTest",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseDevicesGetDeviceLiveToolsThroughputTestItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesLiveToolsThroughputTestResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("throughput_test_id"), idParts[1])...)
}

func (r *DevicesLiveToolsThroughputTestResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesLiveToolsThroughputTestRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update
	// No update
	resp.Diagnostics.AddError(
		"Update operation not supported in DevicesLiveToolsThroughputTest",
		"Update operation not supported in DevicesLiveToolsThroughputTest",
	)
	return
}

func (r *DevicesLiveToolsThroughputTestResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting DevicesLiveToolsThroughputTest", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesLiveToolsThroughputTestRs struct {
	Serial           types.String                                                 `tfsdk:"serial"`
	ThroughputTestID types.String                                                 `tfsdk:"throughput_test_id"`
	Error            types.String                                                 `tfsdk:"error"`
	Request          *ResponseDevicesGetDeviceLiveToolsThroughputTestRequestRs    `tfsdk:"request"`
	Result           *ResponseDevicesGetDeviceLiveToolsThroughputTestResultRs     `tfsdk:"result"`
	Status           types.String                                                 `tfsdk:"status"`
	URL              types.String                                                 `tfsdk:"url"`
	Callback         *RequestDevicesCreateDeviceLiveToolsThroughputTestCallbackRs `tfsdk:"callback"`
}

type ResponseDevicesGetDeviceLiveToolsThroughputTestRequestRs struct {
	Serial types.String `tfsdk:"serial"`
}

type ResponseDevicesGetDeviceLiveToolsThroughputTestResultRs struct {
	Speeds *ResponseDevicesGetDeviceLiveToolsThroughputTestResultSpeedsRs `tfsdk:"speeds"`
}

type ResponseDevicesGetDeviceLiveToolsThroughputTestResultSpeedsRs struct {
	Downstream types.Float64 `tfsdk:"downstream"`
}

type RequestDevicesCreateDeviceLiveToolsThroughputTestCallbackRs struct {
	HTTPServer      *RequestDevicesCreateDeviceLiveToolsThroughputTestCallbackHttpServerRs      `tfsdk:"http_server"`
	PayloadTemplate *RequestDevicesCreateDeviceLiveToolsThroughputTestCallbackPayloadTemplateRs `tfsdk:"payload_template"`
	SharedSecret    types.String                                                                `tfsdk:"shared_secret"`
	URL             types.String                                                                `tfsdk:"url"`
}

type RequestDevicesCreateDeviceLiveToolsThroughputTestCallbackHttpServerRs struct {
	ID types.String `tfsdk:"id"`
}

type RequestDevicesCreateDeviceLiveToolsThroughputTestCallbackPayloadTemplateRs struct {
	ID types.String `tfsdk:"id"`
}

// FromBody
func (r *DevicesLiveToolsThroughputTestRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestDevicesCreateDeviceLiveToolsThroughputTest {
	var requestDevicesCreateDeviceLiveToolsThroughputTestCallback *merakigosdk.RequestDevicesCreateDeviceLiveToolsThroughputTestCallback

	if r.Callback != nil {
		var requestDevicesCreateDeviceLiveToolsThroughputTestCallbackHTTPServer *merakigosdk.RequestDevicesCreateDeviceLiveToolsThroughputTestCallbackHTTPServer

		if r.Callback.HTTPServer != nil {
			id := r.Callback.HTTPServer.ID.ValueString()
			requestDevicesCreateDeviceLiveToolsThroughputTestCallbackHTTPServer = &merakigosdk.RequestDevicesCreateDeviceLiveToolsThroughputTestCallbackHTTPServer{
				ID: id,
			}
			//[debug] Is Array: False
		}
		var requestDevicesCreateDeviceLiveToolsThroughputTestCallbackPayloadTemplate *merakigosdk.RequestDevicesCreateDeviceLiveToolsThroughputTestCallbackPayloadTemplate

		if r.Callback.PayloadTemplate != nil {
			id := r.Callback.PayloadTemplate.ID.ValueString()
			requestDevicesCreateDeviceLiveToolsThroughputTestCallbackPayloadTemplate = &merakigosdk.RequestDevicesCreateDeviceLiveToolsThroughputTestCallbackPayloadTemplate{
				ID: id,
			}
			//[debug] Is Array: False
		}
		sharedSecret := r.Callback.SharedSecret.ValueString()
		url := r.Callback.URL.ValueString()
		requestDevicesCreateDeviceLiveToolsThroughputTestCallback = &merakigosdk.RequestDevicesCreateDeviceLiveToolsThroughputTestCallback{
			HTTPServer:      requestDevicesCreateDeviceLiveToolsThroughputTestCallbackHTTPServer,
			PayloadTemplate: requestDevicesCreateDeviceLiveToolsThroughputTestCallbackPayloadTemplate,
			SharedSecret:    sharedSecret,
			URL:             url,
		}
		//[debug] Is Array: False
	}
	out := merakigosdk.RequestDevicesCreateDeviceLiveToolsThroughputTest{
		Callback: requestDevicesCreateDeviceLiveToolsThroughputTestCallback,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseDevicesGetDeviceLiveToolsThroughputTestItemToBodyRs(state DevicesLiveToolsThroughputTestRs, response *merakigosdk.ResponseDevicesGetDeviceLiveToolsThroughputTest, is_read bool) DevicesLiveToolsThroughputTestRs {
	itemState := DevicesLiveToolsThroughputTestRs{
		Error: types.StringValue(response.Error),
		Request: func() *ResponseDevicesGetDeviceLiveToolsThroughputTestRequestRs {
			if response.Request != nil {
				return &ResponseDevicesGetDeviceLiveToolsThroughputTestRequestRs{
					Serial: types.StringValue(response.Request.Serial),
				}
			}
			return nil
		}(),
		Result: func() *ResponseDevicesGetDeviceLiveToolsThroughputTestResultRs {
			if response.Result != nil {
				return &ResponseDevicesGetDeviceLiveToolsThroughputTestResultRs{
					Speeds: func() *ResponseDevicesGetDeviceLiveToolsThroughputTestResultSpeedsRs {
						if response.Result.Speeds != nil {
							return &ResponseDevicesGetDeviceLiveToolsThroughputTestResultSpeedsRs{
								Downstream: func() types.Float64 {
									if response.Result.Speeds.Downstream != nil {
										return types.Float64Value(float64(*response.Result.Speeds.Downstream))
									}
									return types.Float64{}
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		Status:           types.StringValue(response.Status),
		ThroughputTestID: types.StringValue(response.ThroughputTestID),
		URL:              types.StringValue(response.URL),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesLiveToolsThroughputTestRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesLiveToolsThroughputTestRs)
}
