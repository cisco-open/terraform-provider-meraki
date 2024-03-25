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
	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesLiveToolsPingResource{}
	_ resource.ResourceWithConfigure = &DevicesLiveToolsPingResource{}
)

func NewDevicesLiveToolsPingResource() resource.Resource {
	return &DevicesLiveToolsPingResource{}
}

type DevicesLiveToolsPingResource struct {
	client *merakigosdk.Client
}

func (r *DevicesLiveToolsPingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesLiveToolsPingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_live_tools_ping"
}

// resourceAction
func (r *DevicesLiveToolsPingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

					"ping_id": schema.StringAttribute{
						MarkdownDescription: `Id to check the status of your ping request.`,
						Computed:            true,
					},
					"request": schema.SingleNestedAttribute{
						MarkdownDescription: `Ping request parameters`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"count": schema.Int64Attribute{
								MarkdownDescription: `Number of pings to send`,
								Computed:            true,
							},
							"serial": schema.StringAttribute{
								MarkdownDescription: `Device serial number`,
								Computed:            true,
							},
							"target": schema.StringAttribute{
								MarkdownDescription: `IP address or FQDN to ping`,
								Computed:            true,
							},
						},
					},
					"status": schema.StringAttribute{
						MarkdownDescription: `Status of the ping request.`,
						Computed:            true,
					},
					"url": schema.StringAttribute{
						MarkdownDescription: `GET this url to check the status of your ping request.`,
						Computed:            true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"count": schema.Int64Attribute{
						MarkdownDescription: `Count parameter to pass to ping. [1..5], default 5`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplace(),
						},
					},
					"target": schema.StringAttribute{
						MarkdownDescription: `FQDN, IPv4 or IPv6 address`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
				},
			},
		},
	}
}

func (r *DevicesLiveToolsPingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesLiveToolsPingRs
	var item types.Object
	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("[DEBUG] Despues 1")
	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	//Has Paths
	vvSerial := data.Serial.ValueString()
	// serial
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Devices.CreateDeviceLiveToolsPing(vvSerial, dataRequest)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateDeviceLiveToolsPing",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateDeviceLiveToolsPing",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseDevicesCreateDeviceLiveToolsPingItemToBody(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesLiveToolsPingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *DevicesLiveToolsPingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *DevicesLiveToolsPingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesLiveToolsPingRs struct {
	Serial     types.String                               `tfsdk:"serial"`
	Item       *ResponseDevicesCreateDeviceLiveToolsPing  `tfsdk:"item"`
	Parameters *RequestDevicesCreateDeviceLiveToolsPingRs `tfsdk:"parameters"`
}

type ResponseDevicesCreateDeviceLiveToolsPing struct {
	PingID  types.String                                     `tfsdk:"ping_id"`
	Request *ResponseDevicesCreateDeviceLiveToolsPingRequest `tfsdk:"request"`
	Status  types.String                                     `tfsdk:"status"`
	URL     types.String                                     `tfsdk:"url"`
}

type ResponseDevicesCreateDeviceLiveToolsPingRequest struct {
	Count  types.Int64  `tfsdk:"count"`
	Serial types.String `tfsdk:"serial"`
	Target types.String `tfsdk:"target"`
}

type RequestDevicesCreateDeviceLiveToolsPingRs struct {
	Count  types.Int64  `tfsdk:"count"`
	Target types.String `tfsdk:"target"`
}

// FromBody
func (r *DevicesLiveToolsPingRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestDevicesCreateDeviceLiveToolsPing {
	emptyString := ""
	re := *r.Parameters
	count := new(int64)
	if !re.Count.IsUnknown() && !re.Count.IsNull() {
		*count = re.Count.ValueInt64()
	} else {
		count = nil
	}
	target := new(string)
	if !re.Target.IsUnknown() && !re.Target.IsNull() {
		*target = re.Target.ValueString()
	} else {
		target = &emptyString
	}
	out := merakigosdk.RequestDevicesCreateDeviceLiveToolsPing{
		Count:  int64ToIntPointer(count),
		Target: *target,
	}
	return &out
}

// ToBody
func ResponseDevicesCreateDeviceLiveToolsPingItemToBody(state DevicesLiveToolsPingRs, response *merakigosdk.ResponseDevicesCreateDeviceLiveToolsPing) DevicesLiveToolsPingRs {
	itemState := ResponseDevicesCreateDeviceLiveToolsPing{
		PingID: types.StringValue(response.PingID),
		Request: func() *ResponseDevicesCreateDeviceLiveToolsPingRequest {
			if response.Request != nil {
				return &ResponseDevicesCreateDeviceLiveToolsPingRequest{
					Count: func() types.Int64 {
						if response.Request.Count != nil {
							return types.Int64Value(int64(*response.Request.Count))
						}
						return types.Int64{}
					}(),
					Serial: types.StringValue(response.Request.Serial),
					Target: types.StringValue(response.Request.Target),
				}
			}
			return &ResponseDevicesCreateDeviceLiveToolsPingRequest{}
		}(),
		Status: types.StringValue(response.Status),
		URL:    types.StringValue(response.URL),
	}
	state.Item = &itemState
	return state
}
