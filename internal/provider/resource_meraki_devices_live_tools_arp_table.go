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
	_ resource.Resource              = &DevicesLiveToolsArpTableResource{}
	_ resource.ResourceWithConfigure = &DevicesLiveToolsArpTableResource{}
)

func NewDevicesLiveToolsArpTableResource() resource.Resource {
	return &DevicesLiveToolsArpTableResource{}
}

type DevicesLiveToolsArpTableResource struct {
	client *merakigosdk.Client
}

func (r *DevicesLiveToolsArpTableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesLiveToolsArpTableResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_live_tools_arp_table"
}

func (r *DevicesLiveToolsArpTableResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"arp_table_id": schema.StringAttribute{
				MarkdownDescription: `Id of the ARP table request. Used to check the status of the request.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
			"entries": schema.SetNestedAttribute{
				MarkdownDescription: `The ARP table entries`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"ip": schema.StringAttribute{
							MarkdownDescription: `The IP address of the ARP table entry`,
							Computed:            true,
						},
						"last_updated_at": schema.StringAttribute{
							MarkdownDescription: `Time of the last update of the ARP table entry`,
							Computed:            true,
						},
						"mac": schema.StringAttribute{
							MarkdownDescription: `The MAC address of the ARP table entry`,
							Computed:            true,
						},
						"vlan_id": schema.Int64Attribute{
							MarkdownDescription: `The VLAN ID of the ARP table entry`,
							Computed:            true,
						},
					},
				},
			},
			"error": schema.StringAttribute{
				MarkdownDescription: `An error message for a failed execution`,
				Computed:            true,
			},
			"request": schema.SingleNestedAttribute{
				MarkdownDescription: `ARP table request parameters`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{

					"serial": schema.StringAttribute{
						MarkdownDescription: `Device serial number`,
						Computed:            true,
					},
				},
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: `Status of the ARP table request.
                                  Allowed values: [complete,failed,new,ready,running,scheduled]`,
				Computed: true,
			},
			"url": schema.StringAttribute{
				MarkdownDescription: `GET this url to check the status of your ARP table request.`,
				Computed:            true,
			},
		},
	}
}

func (r *DevicesLiveToolsArpTableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesLiveToolsArpTableRs

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
	vvArpTableID := data.ArpTableID.ValueString()
	//Has Item and not has items

	if vvSerial != "" && vvArpTableID != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Devices.GetDeviceLiveToolsArpTable(vvSerial, vvArpTableID)
		//Has Post
		if err != nil {
			if restyResp1 != nil {
				if restyResp1.StatusCode() != 404 {
					resp.Diagnostics.AddError(
						"Failure when executing GetDeviceLiveToolsArpTable",
						err.Error(),
					)
					return
				}
			}
		}

		if responseVerifyItem != nil {
			data = ResponseDevicesGetDeviceLiveToolsArpTableItemToBodyRs(data, responseVerifyItem, false)
			//Path params in update assigned
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Devices.CreateDeviceLiveToolsArpTable(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateDeviceLiveToolsArpTable",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateDeviceLiveToolsArpTable",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Devices.GetDeviceLiveToolsArpTable(vvSerial, vvArpTableID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceLiveToolsArpTable",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceLiveToolsArpTable",
			err.Error(),
		)
		return
	}

	data = ResponseDevicesGetDeviceLiveToolsArpTableItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *DevicesLiveToolsArpTableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesLiveToolsArpTableRs

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
	vvArpTableID := data.ArpTableID.ValueString()
	responseGet, restyRespGet, err := r.client.Devices.GetDeviceLiveToolsArpTable(vvSerial, vvArpTableID)
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
				"Failure when executing GetDeviceLiveToolsArpTable",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceLiveToolsArpTable",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseDevicesGetDeviceLiveToolsArpTableItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesLiveToolsArpTableResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("arp_table_id"), idParts[1])...)
}

func (r *DevicesLiveToolsArpTableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesLiveToolsArpTableRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update
	// No update
	resp.Diagnostics.AddError(
		"Update operation not supported in DevicesLiveToolsArpTable",
		"Update operation not supported in DevicesLiveToolsArpTable",
	)
	return
}

func (r *DevicesLiveToolsArpTableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting DevicesLiveToolsArpTable", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesLiveToolsArpTableRs struct {
	Serial     types.String                                           `tfsdk:"serial"`
	ArpTableID types.String                                           `tfsdk:"arp_table_id"`
	Entries    *[]ResponseDevicesGetDeviceLiveToolsArpTableEntriesRs  `tfsdk:"entries"`
	Error      types.String                                           `tfsdk:"error"`
	Request    *ResponseDevicesGetDeviceLiveToolsArpTableRequestRs    `tfsdk:"request"`
	Status     types.String                                           `tfsdk:"status"`
	URL        types.String                                           `tfsdk:"url"`
	Callback   *RequestDevicesCreateDeviceLiveToolsArpTableCallbackRs `tfsdk:"callback"`
}

type ResponseDevicesGetDeviceLiveToolsArpTableEntriesRs struct {
	IP            types.String `tfsdk:"ip"`
	LastUpdatedAt types.String `tfsdk:"last_updated_at"`
	Mac           types.String `tfsdk:"mac"`
	VLANID        types.Int64  `tfsdk:"vlan_id"`
}

type ResponseDevicesGetDeviceLiveToolsArpTableRequestRs struct {
	Serial types.String `tfsdk:"serial"`
}

type RequestDevicesCreateDeviceLiveToolsArpTableCallbackRs struct {
	HTTPServer      *RequestDevicesCreateDeviceLiveToolsArpTableCallbackHttpServerRs      `tfsdk:"http_server"`
	PayloadTemplate *RequestDevicesCreateDeviceLiveToolsArpTableCallbackPayloadTemplateRs `tfsdk:"payload_template"`
	SharedSecret    types.String                                                          `tfsdk:"shared_secret"`
	URL             types.String                                                          `tfsdk:"url"`
}

type RequestDevicesCreateDeviceLiveToolsArpTableCallbackHttpServerRs struct {
	ID types.String `tfsdk:"id"`
}

type RequestDevicesCreateDeviceLiveToolsArpTableCallbackPayloadTemplateRs struct {
	ID types.String `tfsdk:"id"`
}

// FromBody
func (r *DevicesLiveToolsArpTableRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestDevicesCreateDeviceLiveToolsArpTable {
	var requestDevicesCreateDeviceLiveToolsArpTableCallback *merakigosdk.RequestDevicesCreateDeviceLiveToolsArpTableCallback

	if r.Callback != nil {
		var requestDevicesCreateDeviceLiveToolsArpTableCallbackHTTPServer *merakigosdk.RequestDevicesCreateDeviceLiveToolsArpTableCallbackHTTPServer

		if r.Callback.HTTPServer != nil {
			id := r.Callback.HTTPServer.ID.ValueString()
			requestDevicesCreateDeviceLiveToolsArpTableCallbackHTTPServer = &merakigosdk.RequestDevicesCreateDeviceLiveToolsArpTableCallbackHTTPServer{
				ID: id,
			}
			//[debug] Is Array: False
		}
		var requestDevicesCreateDeviceLiveToolsArpTableCallbackPayloadTemplate *merakigosdk.RequestDevicesCreateDeviceLiveToolsArpTableCallbackPayloadTemplate

		if r.Callback.PayloadTemplate != nil {
			id := r.Callback.PayloadTemplate.ID.ValueString()
			requestDevicesCreateDeviceLiveToolsArpTableCallbackPayloadTemplate = &merakigosdk.RequestDevicesCreateDeviceLiveToolsArpTableCallbackPayloadTemplate{
				ID: id,
			}
			//[debug] Is Array: False
		}
		sharedSecret := r.Callback.SharedSecret.ValueString()
		url := r.Callback.URL.ValueString()
		requestDevicesCreateDeviceLiveToolsArpTableCallback = &merakigosdk.RequestDevicesCreateDeviceLiveToolsArpTableCallback{
			HTTPServer:      requestDevicesCreateDeviceLiveToolsArpTableCallbackHTTPServer,
			PayloadTemplate: requestDevicesCreateDeviceLiveToolsArpTableCallbackPayloadTemplate,
			SharedSecret:    sharedSecret,
			URL:             url,
		}
		//[debug] Is Array: False
	}
	out := merakigosdk.RequestDevicesCreateDeviceLiveToolsArpTable{
		Callback: requestDevicesCreateDeviceLiveToolsArpTableCallback,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseDevicesGetDeviceLiveToolsArpTableItemToBodyRs(state DevicesLiveToolsArpTableRs, response *merakigosdk.ResponseDevicesGetDeviceLiveToolsArpTable, is_read bool) DevicesLiveToolsArpTableRs {
	itemState := DevicesLiveToolsArpTableRs{
		ArpTableID: types.StringValue(response.ArpTableID),
		Entries: func() *[]ResponseDevicesGetDeviceLiveToolsArpTableEntriesRs {
			if response.Entries != nil {
				result := make([]ResponseDevicesGetDeviceLiveToolsArpTableEntriesRs, len(*response.Entries))
				for i, entries := range *response.Entries {
					result[i] = ResponseDevicesGetDeviceLiveToolsArpTableEntriesRs{
						IP:            types.StringValue(entries.IP),
						LastUpdatedAt: types.StringValue(entries.LastUpdatedAt),
						Mac:           types.StringValue(entries.Mac),
						VLANID: func() types.Int64 {
							if entries.VLANID != nil {
								return types.Int64Value(int64(*entries.VLANID))
							}
							return types.Int64{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		Error: types.StringValue(response.Error),
		Request: func() *ResponseDevicesGetDeviceLiveToolsArpTableRequestRs {
			if response.Request != nil {
				return &ResponseDevicesGetDeviceLiveToolsArpTableRequestRs{
					Serial: types.StringValue(response.Request.Serial),
				}
			}
			return nil
		}(),
		Status: types.StringValue(response.Status),
		URL:    types.StringValue(response.URL),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesLiveToolsArpTableRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesLiveToolsArpTableRs)
}
