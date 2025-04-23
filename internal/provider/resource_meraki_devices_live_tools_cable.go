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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesLiveToolsCableResource{}
	_ resource.ResourceWithConfigure = &DevicesLiveToolsCableResource{}
)

func NewDevicesLiveToolsCableResource() resource.Resource {
	return &DevicesLiveToolsCableResource{}
}

type DevicesLiveToolsCableResource struct {
	client *merakigosdk.Client
}

func (r *DevicesLiveToolsCableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesLiveToolsCableResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_live_tools_cable"
}

func (r *DevicesLiveToolsCableResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"cable_test_id": schema.StringAttribute{
				MarkdownDescription: `Id of the cable test request. Used to check the status of the request.`,
				Computed:            true,
			},
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
			"id": schema.StringAttribute{
				MarkdownDescription: `id path parameter.`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ports": schema.SetAttribute{
				MarkdownDescription: `A list of ports for which to perform the cable test.  For Catalyst switches, IOS interface names are also supported, such as "GigabitEthernet1/0/8", "Gi1/0/8", or even "1/0/8".`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"request": schema.SingleNestedAttribute{
				MarkdownDescription: `Cable test request parameters`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{

					"ports": schema.SetAttribute{
						MarkdownDescription: `A list of ports for which to perform the cable test.`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"serial": schema.StringAttribute{
						MarkdownDescription: `Device serial number`,
						Computed:            true,
					},
				},
			},
			"results": schema.SetNestedAttribute{
				MarkdownDescription: `Results of the cable test request, one for each requested port.`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"error": schema.StringAttribute{
							MarkdownDescription: `If an error occurred during the cable test, the error message will be populated here.`,
							Computed:            true,
						},
						"pairs": schema.SetNestedAttribute{
							MarkdownDescription: `Results for each twisted pair within the cable.`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"index": schema.Int64Attribute{
										MarkdownDescription: `The index of the twisted pair tested.
                                              Allowed values: [0,1,2,3]`,
										Computed: true,
									},
									"length_meters": schema.Int64Attribute{
										MarkdownDescription: `The detected length of the twisted pair.`,
										Computed:            true,
									},
									"status": schema.StringAttribute{
										MarkdownDescription: `The test result of the twisted pair tested.
                                              Allowed values: [abnormal,couplex,fail,forced,in-progress,invalid,not-supported,ok,open,open or short,short,short or abnormal,short or couplex,unknown]`,
										Computed: true,
									},
								},
							},
						},
						"port": schema.StringAttribute{
							MarkdownDescription: `The port for which the test was performed.`,
							Computed:            true,
						},
						"speed_mbps": schema.Int64Attribute{
							MarkdownDescription: `Speed in Mbps.  A speed of 0 indicates the port is down or the port speed is automatic.`,
							Computed:            true,
						},
						"status": schema.StringAttribute{
							MarkdownDescription: `The current status of the port. If the cable test is still being performed on the port, "in-progress" is used. If an error occurred during the cable test, "error" is used and the error property will be populated.
                                        Allowed values: [down,error,in-progress,up]`,
							Computed: true,
						},
					},
				},
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: `Status of the cable test request.
                                  Allowed values: [complete,failed,new,ready,running,scheduled]`,
				Computed: true,
			},
			"url": schema.StringAttribute{
				MarkdownDescription: `GET this url to check the status of your cable test request.`,
				Computed:            true,
			},
		},
	}
}

func (r *DevicesLiveToolsCableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesLiveToolsCableRs

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
	vvID := data.ID.ValueString()
	//Has Item and not has items

	if vvSerial != "" && vvID != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Devices.GetDeviceLiveToolsCableTest(vvSerial, vvID)
		//Has Post
		if err != nil {
			if restyResp1 != nil {
				if restyResp1.StatusCode() != 404 {
					resp.Diagnostics.AddError(
						"Failure when executing GetDeviceLiveToolsCableTest",
						err.Error(),
					)
					return
				}
			}
		}

		if responseVerifyItem != nil {
			data = ResponseDevicesGetDeviceLiveToolsCableTestItemToBodyRs(data, responseVerifyItem, false)
			//Path params in update assigned
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Devices.CreateDeviceLiveToolsCableTest(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateDeviceLiveToolsCableTest",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateDeviceLiveToolsCableTest",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Devices.GetDeviceLiveToolsCableTest(vvSerial, vvID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceLiveToolsCableTest",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceLiveToolsCableTest",
			err.Error(),
		)
		return
	}

	data = ResponseDevicesGetDeviceLiveToolsCableTestItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *DevicesLiveToolsCableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesLiveToolsCableRs

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
	vvID := data.ID.ValueString()
	responseGet, restyRespGet, err := r.client.Devices.GetDeviceLiveToolsCableTest(vvSerial, vvID)
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
				"Failure when executing GetDeviceLiveToolsCableTest",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceLiveToolsCableTest",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseDevicesGetDeviceLiveToolsCableTestItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesLiveToolsCableResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
}

func (r *DevicesLiveToolsCableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesLiveToolsCableRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update
	// No update
	resp.Diagnostics.AddError(
		"Update operation not supported in DevicesLiveToolsCable",
		"Update operation not supported in DevicesLiveToolsCable",
	)
	return
}

func (r *DevicesLiveToolsCableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting DevicesLiveToolsCable", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesLiveToolsCableRs struct {
	Serial      types.String                                            `tfsdk:"serial"`
	ID          types.String                                            `tfsdk:"id"`
	CableTestID types.String                                            `tfsdk:"cable_test_id"`
	Error       types.String                                            `tfsdk:"error"`
	Request     *ResponseDevicesGetDeviceLiveToolsCableTestRequestRs    `tfsdk:"request"`
	Results     *[]ResponseDevicesGetDeviceLiveToolsCableTestResultsRs  `tfsdk:"results"`
	Status      types.String                                            `tfsdk:"status"`
	URL         types.String                                            `tfsdk:"url"`
	Callback    *RequestDevicesCreateDeviceLiveToolsCableTestCallbackRs `tfsdk:"callback"`
	Ports       types.Set                                               `tfsdk:"ports"`
}

type ResponseDevicesGetDeviceLiveToolsCableTestRequestRs struct {
	Ports  types.Set    `tfsdk:"ports"`
	Serial types.String `tfsdk:"serial"`
}

type ResponseDevicesGetDeviceLiveToolsCableTestResultsRs struct {
	Error     types.String                                                `tfsdk:"error"`
	Pairs     *[]ResponseDevicesGetDeviceLiveToolsCableTestResultsPairsRs `tfsdk:"pairs"`
	Port      types.String                                                `tfsdk:"port"`
	SpeedMbps types.Int64                                                 `tfsdk:"speed_mbps"`
	Status    types.String                                                `tfsdk:"status"`
}

type ResponseDevicesGetDeviceLiveToolsCableTestResultsPairsRs struct {
	Index        types.Int64  `tfsdk:"index"`
	LengthMeters types.Int64  `tfsdk:"length_meters"`
	Status       types.String `tfsdk:"status"`
}

type RequestDevicesCreateDeviceLiveToolsCableTestCallbackRs struct {
	HTTPServer      *RequestDevicesCreateDeviceLiveToolsCableTestCallbackHttpServerRs      `tfsdk:"http_server"`
	PayloadTemplate *RequestDevicesCreateDeviceLiveToolsCableTestCallbackPayloadTemplateRs `tfsdk:"payload_template"`
	SharedSecret    types.String                                                           `tfsdk:"shared_secret"`
	URL             types.String                                                           `tfsdk:"url"`
}

type RequestDevicesCreateDeviceLiveToolsCableTestCallbackHttpServerRs struct {
	ID types.String `tfsdk:"id"`
}

type RequestDevicesCreateDeviceLiveToolsCableTestCallbackPayloadTemplateRs struct {
	ID types.String `tfsdk:"id"`
}

// FromBody
func (r *DevicesLiveToolsCableRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestDevicesCreateDeviceLiveToolsCableTest {
	var requestDevicesCreateDeviceLiveToolsCableTestCallback *merakigosdk.RequestDevicesCreateDeviceLiveToolsCableTestCallback

	if r.Callback != nil {
		var requestDevicesCreateDeviceLiveToolsCableTestCallbackHTTPServer *merakigosdk.RequestDevicesCreateDeviceLiveToolsCableTestCallbackHTTPServer

		if r.Callback.HTTPServer != nil {
			id := r.Callback.HTTPServer.ID.ValueString()
			requestDevicesCreateDeviceLiveToolsCableTestCallbackHTTPServer = &merakigosdk.RequestDevicesCreateDeviceLiveToolsCableTestCallbackHTTPServer{
				ID: id,
			}
			//[debug] Is Array: False
		}
		var requestDevicesCreateDeviceLiveToolsCableTestCallbackPayloadTemplate *merakigosdk.RequestDevicesCreateDeviceLiveToolsCableTestCallbackPayloadTemplate

		if r.Callback.PayloadTemplate != nil {
			id := r.Callback.PayloadTemplate.ID.ValueString()
			requestDevicesCreateDeviceLiveToolsCableTestCallbackPayloadTemplate = &merakigosdk.RequestDevicesCreateDeviceLiveToolsCableTestCallbackPayloadTemplate{
				ID: id,
			}
			//[debug] Is Array: False
		}
		sharedSecret := r.Callback.SharedSecret.ValueString()
		url := r.Callback.URL.ValueString()
		requestDevicesCreateDeviceLiveToolsCableTestCallback = &merakigosdk.RequestDevicesCreateDeviceLiveToolsCableTestCallback{
			HTTPServer:      requestDevicesCreateDeviceLiveToolsCableTestCallbackHTTPServer,
			PayloadTemplate: requestDevicesCreateDeviceLiveToolsCableTestCallbackPayloadTemplate,
			SharedSecret:    sharedSecret,
			URL:             url,
		}
		//[debug] Is Array: False
	}
	var ports []string = nil
	r.Ports.ElementsAs(ctx, &ports, false)
	out := merakigosdk.RequestDevicesCreateDeviceLiveToolsCableTest{
		Callback: requestDevicesCreateDeviceLiveToolsCableTestCallback,
		Ports:    ports,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseDevicesGetDeviceLiveToolsCableTestItemToBodyRs(state DevicesLiveToolsCableRs, response *merakigosdk.ResponseDevicesGetDeviceLiveToolsCableTest, is_read bool) DevicesLiveToolsCableRs {
	itemState := DevicesLiveToolsCableRs{
		CableTestID: types.StringValue(response.CableTestID),
		Error:       types.StringValue(response.Error),
		Request: func() *ResponseDevicesGetDeviceLiveToolsCableTestRequestRs {
			if response.Request != nil {
				return &ResponseDevicesGetDeviceLiveToolsCableTestRequestRs{
					Ports:  StringSliceToSet(response.Request.Ports),
					Serial: types.StringValue(response.Request.Serial),
				}
			}
			return nil
		}(),
		Results: func() *[]ResponseDevicesGetDeviceLiveToolsCableTestResultsRs {
			if response.Results != nil {
				result := make([]ResponseDevicesGetDeviceLiveToolsCableTestResultsRs, len(*response.Results))
				for i, results := range *response.Results {
					result[i] = ResponseDevicesGetDeviceLiveToolsCableTestResultsRs{
						Error: types.StringValue(results.Error),
						Pairs: func() *[]ResponseDevicesGetDeviceLiveToolsCableTestResultsPairsRs {
							if results.Pairs != nil {
								result := make([]ResponseDevicesGetDeviceLiveToolsCableTestResultsPairsRs, len(*results.Pairs))
								for i, pairs := range *results.Pairs {
									result[i] = ResponseDevicesGetDeviceLiveToolsCableTestResultsPairsRs{
										Index: func() types.Int64 {
											if pairs.Index != nil {
												return types.Int64Value(int64(*pairs.Index))
											}
											return types.Int64{}
										}(),
										LengthMeters: func() types.Int64 {
											if pairs.LengthMeters != nil {
												return types.Int64Value(int64(*pairs.LengthMeters))
											}
											return types.Int64{}
										}(),
										Status: types.StringValue(pairs.Status),
									}
								}
								return &result
							}
							return nil
						}(),
						Port: types.StringValue(results.Port),
						SpeedMbps: func() types.Int64 {
							if results.SpeedMbps != nil {
								return types.Int64Value(int64(*results.SpeedMbps))
							}
							return types.Int64{}
						}(),
						Status: types.StringValue(results.Status),
					}
				}
				return &result
			}
			return nil
		}(),
		Status: types.StringValue(response.Status),
		URL:    types.StringValue(response.URL),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesLiveToolsCableRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesLiveToolsCableRs)
}
