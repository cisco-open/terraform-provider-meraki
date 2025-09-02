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
	"strconv"

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
	_ resource.Resource              = &DevicesCameraWirelessProfilesResource{}
	_ resource.ResourceWithConfigure = &DevicesCameraWirelessProfilesResource{}
)

func NewDevicesCameraWirelessProfilesResource() resource.Resource {
	return &DevicesCameraWirelessProfilesResource{}
}

type DevicesCameraWirelessProfilesResource struct {
	client *merakigosdk.Client
}

func (r *DevicesCameraWirelessProfilesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesCameraWirelessProfilesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_camera_wireless_profiles"
}

func (r *DevicesCameraWirelessProfilesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ids": schema.SingleNestedAttribute{
				MarkdownDescription: `The ids of the wireless profile to assign to the given camera`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"backup": schema.StringAttribute{
						MarkdownDescription: `The id of the backup wireless profile`,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"primary": schema.StringAttribute{
						MarkdownDescription: `The id of the primary wireless profile`,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"secondary": schema.StringAttribute{
						MarkdownDescription: `The id of the secondary wireless profile`,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
		},
	}
}

func (r *DevicesCameraWirelessProfilesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesCameraWirelessProfilesRs

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
	//Has Item and not has items

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Camera.UpdateDeviceCameraWirelessProfiles(vvSerial, dataRequest)
	//Update
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceCameraWirelessProfiles",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceCameraWirelessProfiles",
			err.Error(),
		)
		return
	}

	// Assign data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *DevicesCameraWirelessProfilesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesCameraWirelessProfilesRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	//Has Paths
	// Has Item2

	vvSerial := data.Serial.ValueString()
	responseGet, restyRespGet, err := r.client.Camera.GetDeviceCameraWirelessProfiles(vvSerial)
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
				"Failure when executing GetDeviceCameraWirelessProfiles",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceCameraWirelessProfiles",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseCameraGetDeviceCameraWirelessProfilesItemToBodyRs(data, responseGet, true)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *DevicesCameraWirelessProfilesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), req.ID)...)
}

func (r *DevicesCameraWirelessProfilesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan DevicesCameraWirelessProfilesRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvSerial := plan.Serial.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Camera.UpdateDeviceCameraWirelessProfiles(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceCameraWirelessProfiles",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceCameraWirelessProfiles",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *DevicesCameraWirelessProfilesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting DevicesCameraWirelessProfiles", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesCameraWirelessProfilesRs struct {
	Serial types.String                                        `tfsdk:"serial"`
	IDs    *ResponseCameraGetDeviceCameraWirelessProfilesIdsRs `tfsdk:"ids"`
}

type ResponseCameraGetDeviceCameraWirelessProfilesIdsRs struct {
	Backup    types.String `tfsdk:"backup"`
	Primary   types.String `tfsdk:"primary"`
	Secondary types.String `tfsdk:"secondary"`
}

// FromBody
func (r *DevicesCameraWirelessProfilesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestCameraUpdateDeviceCameraWirelessProfiles {
	var requestCameraUpdateDeviceCameraWirelessProfilesIDs *merakigosdk.RequestCameraUpdateDeviceCameraWirelessProfilesIDs

	if r.IDs != nil {
		backup := r.IDs.Backup.ValueString()
		primary := r.IDs.Primary.ValueString()
		secondary := r.IDs.Secondary.ValueString()
		requestCameraUpdateDeviceCameraWirelessProfilesIDs = &merakigosdk.RequestCameraUpdateDeviceCameraWirelessProfilesIDs{
			Backup:    backup,
			Primary:   primary,
			Secondary: secondary,
		}
		//[debug] Is Array: False
	}
	out := merakigosdk.RequestCameraUpdateDeviceCameraWirelessProfiles{
		IDs: requestCameraUpdateDeviceCameraWirelessProfilesIDs,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseCameraGetDeviceCameraWirelessProfilesItemToBodyRs(state DevicesCameraWirelessProfilesRs, response *merakigosdk.ResponseCameraGetDeviceCameraWirelessProfiles, is_read bool) DevicesCameraWirelessProfilesRs {
	itemState := DevicesCameraWirelessProfilesRs{
		IDs: func() *ResponseCameraGetDeviceCameraWirelessProfilesIdsRs {
			if response.IDs != nil {
				return &ResponseCameraGetDeviceCameraWirelessProfilesIdsRs{
					Backup: func() types.String {
						if response.IDs.Backup != "" {
							return types.StringValue(response.IDs.Backup)
						}
						return types.String{}
					}(),
					Primary: func() types.String {
						if response.IDs.Primary != "" {
							return types.StringValue(response.IDs.Primary)
						}
						return types.String{}
					}(),
					Secondary: func() types.String {
						if response.IDs.Secondary != "" {
							return types.StringValue(response.IDs.Secondary)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesCameraWirelessProfilesRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesCameraWirelessProfilesRs)
}
