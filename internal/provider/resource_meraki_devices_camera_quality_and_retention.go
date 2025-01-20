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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesCameraQualityAndRetentionResource{}
	_ resource.ResourceWithConfigure = &DevicesCameraQualityAndRetentionResource{}
)

func NewDevicesCameraQualityAndRetentionResource() resource.Resource {
	return &DevicesCameraQualityAndRetentionResource{}
}

type DevicesCameraQualityAndRetentionResource struct {
	client *merakigosdk.Client
}

func (r *DevicesCameraQualityAndRetentionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesCameraQualityAndRetentionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_camera_quality_and_retention"
}

func (r *DevicesCameraQualityAndRetentionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"audio_recording_enabled": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating if audio recording is enabled(true) or disabled(false) on the camera`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"motion_based_retention_enabled": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating if motion-based retention is enabled(true) or disabled(false) on the camera.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"motion_detector_version": schema.Int64Attribute{
				MarkdownDescription: `The version of the motion detector that will be used by the camera. Only applies to Gen 2 cameras. Defaults to v2.
                                  Allowed values: [1,2]`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"profile_id": schema.StringAttribute{
				MarkdownDescription: `The ID of a quality and retention profile to assign to the camera. The profile's settings will override all of the per-camera quality and retention settings. If the value of this parameter is null, any existing profile will be unassigned from the camera.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"quality": schema.StringAttribute{
				MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'High' or 'Enhanced'. Not all qualities are supported by every camera model.
                                  Allowed values: [Enhanced,High,Standard]`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Enhanced",
						"High",
						"Standard",
					),
				},
			},
			"resolution": schema.StringAttribute{
				MarkdownDescription: `Resolution of the camera. Can be one of '1280x720', '1920x1080', '1080x1080', '2112x2112', '2880x2880', '2688x1512' or '3840x2160'.Not all resolutions are supported by every camera model.
                                  Allowed values: [1080x1080,1280x720,1920x1080,2112x2112,2688x1512,2880x2880,3840x2160]`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"1080x1080",
						"1280x720",
						"1920x1080",
						"2112x2112",
						"2688x1512",
						"2880x2880",
						"3840x2160",
					),
				},
			},
			"restricted_bandwidth_mode_enabled": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating if restricted bandwidth is enabled(true) or disabled(false) on the camera. This setting does not apply to MV2 cameras.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
		},
	}
}

func (r *DevicesCameraQualityAndRetentionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesCameraQualityAndRetentionRs

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
	//Item
	responseVerifyItem, restyResp1, err := r.client.Camera.GetDeviceCameraQualityAndRetention(vvSerial)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource DevicesCameraQualityAndRetention only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource DevicesCameraQualityAndRetention only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Camera.UpdateDeviceCameraQualityAndRetention(vvSerial, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceCameraQualityAndRetention",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceCameraQualityAndRetention",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Camera.GetDeviceCameraQualityAndRetention(vvSerial)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCameraQualityAndRetention",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceCameraQualityAndRetention",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseCameraGetDeviceCameraQualityAndRetentionItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesCameraQualityAndRetentionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesCameraQualityAndRetentionRs

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
	responseGet, restyRespGet, err := r.client.Camera.GetDeviceCameraQualityAndRetention(vvSerial)
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
				"Failure when executing GetDeviceCameraQualityAndRetention",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceCameraQualityAndRetention",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseCameraGetDeviceCameraQualityAndRetentionItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *DevicesCameraQualityAndRetentionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), req.ID)...)
}

func (r *DevicesCameraQualityAndRetentionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesCameraQualityAndRetentionRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvSerial := data.Serial.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Camera.UpdateDeviceCameraQualityAndRetention(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceCameraQualityAndRetention",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceCameraQualityAndRetention",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesCameraQualityAndRetentionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting DevicesCameraQualityAndRetention", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesCameraQualityAndRetentionRs struct {
	Serial                         types.String `tfsdk:"serial"`
	AudioRecordingEnabled          types.Bool   `tfsdk:"audio_recording_enabled"`
	MotionBasedRetentionEnabled    types.Bool   `tfsdk:"motion_based_retention_enabled"`
	MotionDetectorVersion          types.Int64  `tfsdk:"motion_detector_version"`
	ProfileID                      types.String `tfsdk:"profile_id"`
	Quality                        types.String `tfsdk:"quality"`
	Resolution                     types.String `tfsdk:"resolution"`
	RestrictedBandwidthModeEnabled types.Bool   `tfsdk:"restricted_bandwidth_mode_enabled"`
}

// FromBody
func (r *DevicesCameraQualityAndRetentionRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestCameraUpdateDeviceCameraQualityAndRetention {
	emptyString := ""
	audioRecordingEnabled := new(bool)
	if !r.AudioRecordingEnabled.IsUnknown() && !r.AudioRecordingEnabled.IsNull() {
		*audioRecordingEnabled = r.AudioRecordingEnabled.ValueBool()
	} else {
		audioRecordingEnabled = nil
	}
	motionBasedRetentionEnabled := new(bool)
	if !r.MotionBasedRetentionEnabled.IsUnknown() && !r.MotionBasedRetentionEnabled.IsNull() {
		*motionBasedRetentionEnabled = r.MotionBasedRetentionEnabled.ValueBool()
	} else {
		motionBasedRetentionEnabled = nil
	}
	motionDetectorVersion := new(int64)
	if !r.MotionDetectorVersion.IsUnknown() && !r.MotionDetectorVersion.IsNull() {
		*motionDetectorVersion = r.MotionDetectorVersion.ValueInt64()
	} else {
		motionDetectorVersion = nil
	}
	profileID := new(string)
	if !r.ProfileID.IsUnknown() && !r.ProfileID.IsNull() {
		*profileID = r.ProfileID.ValueString()
	} else {
		profileID = &emptyString
	}
	quality := new(string)
	if !r.Quality.IsUnknown() && !r.Quality.IsNull() {
		*quality = r.Quality.ValueString()
	} else {
		quality = &emptyString
	}
	resolution := new(string)
	if !r.Resolution.IsUnknown() && !r.Resolution.IsNull() {
		*resolution = r.Resolution.ValueString()
	} else {
		resolution = &emptyString
	}
	restrictedBandwidthModeEnabled := new(bool)
	if !r.RestrictedBandwidthModeEnabled.IsUnknown() && !r.RestrictedBandwidthModeEnabled.IsNull() {
		*restrictedBandwidthModeEnabled = r.RestrictedBandwidthModeEnabled.ValueBool()
	} else {
		restrictedBandwidthModeEnabled = nil
	}
	out := merakigosdk.RequestCameraUpdateDeviceCameraQualityAndRetention{
		AudioRecordingEnabled:          audioRecordingEnabled,
		MotionBasedRetentionEnabled:    motionBasedRetentionEnabled,
		MotionDetectorVersion:          int64ToIntPointer(motionDetectorVersion),
		ProfileID:                      *profileID,
		Quality:                        *quality,
		Resolution:                     *resolution,
		RestrictedBandwidthModeEnabled: restrictedBandwidthModeEnabled,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseCameraGetDeviceCameraQualityAndRetentionItemToBodyRs(state DevicesCameraQualityAndRetentionRs, response *merakigosdk.ResponseCameraGetDeviceCameraQualityAndRetention, is_read bool) DevicesCameraQualityAndRetentionRs {
	itemState := DevicesCameraQualityAndRetentionRs{
		AudioRecordingEnabled: func() types.Bool {
			if response.AudioRecordingEnabled != nil {
				return types.BoolValue(*response.AudioRecordingEnabled)
			}
			return types.Bool{}
		}(),
		MotionBasedRetentionEnabled: func() types.Bool {
			if response.MotionBasedRetentionEnabled != nil {
				return types.BoolValue(*response.MotionBasedRetentionEnabled)
			}
			return types.Bool{}
		}(),
		MotionDetectorVersion: func() types.Int64 {
			if response.MotionDetectorVersion != nil {
				return types.Int64Value(int64(*response.MotionDetectorVersion))
			}
			return types.Int64{}
		}(),
		ProfileID:  types.StringValue(response.ProfileID),
		Quality:    types.StringValue(response.Quality),
		Resolution: types.StringValue(response.Resolution),
		RestrictedBandwidthModeEnabled: func() types.Bool {
			if response.RestrictedBandwidthModeEnabled != nil {
				return types.BoolValue(*response.RestrictedBandwidthModeEnabled)
			}
			return types.Bool{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesCameraQualityAndRetentionRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesCameraQualityAndRetentionRs)
}
