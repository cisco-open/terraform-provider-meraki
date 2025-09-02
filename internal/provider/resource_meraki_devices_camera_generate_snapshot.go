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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesCameraGenerateSnapshotResource{}
	_ resource.ResourceWithConfigure = &DevicesCameraGenerateSnapshotResource{}
)

func NewDevicesCameraGenerateSnapshotResource() resource.Resource {
	return &DevicesCameraGenerateSnapshotResource{}
}

type DevicesCameraGenerateSnapshotResource struct {
	client *merakigosdk.Client
}

func (r *DevicesCameraGenerateSnapshotResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesCameraGenerateSnapshotResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_camera_generate_snapshot"
}

// resourceAction
func (r *DevicesCameraGenerateSnapshotResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

					"expiry": schema.StringAttribute{
						MarkdownDescription: `Expiration details for snapshot image access`,
						Computed:            true,
					},
					"url": schema.StringAttribute{
						MarkdownDescription: `Url for the snapshot`,
						Computed:            true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"fullframe": schema.BoolAttribute{
						MarkdownDescription: `[optional] If set to "true" the snapshot will be taken at full sensor resolution. This will error if used with timestamp.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.RequiresReplace(),
						},
					},
					"timestamp": schema.StringAttribute{
						MarkdownDescription: `[optional] The snapshot will be taken from this time on the camera. The timestamp is expected to be in ISO 8601 format. If no timestamp is specified, we will assume current time.`,
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
func (r *DevicesCameraGenerateSnapshotResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesCameraGenerateSnapshot

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
	response, restyResp1, err := r.client.Camera.GenerateDeviceCameraSnapshot(vvSerial, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GenerateDeviceCameraSnapshot",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GenerateDeviceCameraSnapshot",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseCameraGenerateDeviceCameraSnapshotItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesCameraGenerateSnapshotResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *DevicesCameraGenerateSnapshotResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *DevicesCameraGenerateSnapshotResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesCameraGenerateSnapshot struct {
	Serial     types.String                                 `tfsdk:"serial"`
	Item       *ResponseCameraGenerateDeviceCameraSnapshot  `tfsdk:"item"`
	Parameters *RequestCameraGenerateDeviceCameraSnapshotRs `tfsdk:"parameters"`
}

type ResponseCameraGenerateDeviceCameraSnapshot struct {
	Expiry types.String `tfsdk:"expiry"`
	URL    types.String `tfsdk:"url"`
}

type RequestCameraGenerateDeviceCameraSnapshotRs struct {
	Fullframe types.Bool   `tfsdk:"fullframe"`
	Timestamp types.String `tfsdk:"timestamp"`
}

// FromBody
func (r *DevicesCameraGenerateSnapshot) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestCameraGenerateDeviceCameraSnapshot {
	emptyString := ""
	re := *r.Parameters
	fullframe := new(bool)
	if !re.Fullframe.IsUnknown() && !re.Fullframe.IsNull() {
		*fullframe = re.Fullframe.ValueBool()
	} else {
		fullframe = nil
	}
	timestamp := new(string)
	if !re.Timestamp.IsUnknown() && !re.Timestamp.IsNull() {
		*timestamp = re.Timestamp.ValueString()
	} else {
		timestamp = &emptyString
	}
	out := merakigosdk.RequestCameraGenerateDeviceCameraSnapshot{
		Fullframe: fullframe,
		Timestamp: *timestamp,
	}
	return &out
}

// ToBody
func ResponseCameraGenerateDeviceCameraSnapshotItemToBody(state DevicesCameraGenerateSnapshot, response *merakigosdk.ResponseCameraGenerateDeviceCameraSnapshot) DevicesCameraGenerateSnapshot {
	itemState := ResponseCameraGenerateDeviceCameraSnapshot{
		Expiry: func() types.String {
			if response.Expiry != "" {
				return types.StringValue(response.Expiry)
			}
			return types.String{}
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
