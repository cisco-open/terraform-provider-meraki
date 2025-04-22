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

// DATA SOURCE NORMAL
import (
	"context"
	"log"

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &DevicesCameraQualityAndRetentionDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesCameraQualityAndRetentionDataSource{}
)

func NewDevicesCameraQualityAndRetentionDataSource() datasource.DataSource {
	return &DevicesCameraQualityAndRetentionDataSource{}
}

type DevicesCameraQualityAndRetentionDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesCameraQualityAndRetentionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesCameraQualityAndRetentionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_camera_quality_and_retention"
}

func (d *DevicesCameraQualityAndRetentionDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"audio_recording_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"motion_based_retention_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"motion_detector_version": schema.Int64Attribute{
						Computed: true,
					},
					"profile_id": schema.StringAttribute{
						Computed: true,
					},
					"quality": schema.StringAttribute{
						Computed: true,
					},
					"resolution": schema.StringAttribute{
						Computed: true,
					},
					"restricted_bandwidth_mode_enabled": schema.BoolAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *DevicesCameraQualityAndRetentionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesCameraQualityAndRetention DevicesCameraQualityAndRetention
	diags := req.Config.Get(ctx, &devicesCameraQualityAndRetention)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceCameraQualityAndRetention")
		vvSerial := devicesCameraQualityAndRetention.Serial.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Camera.GetDeviceCameraQualityAndRetention(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCameraQualityAndRetention",
				err.Error(),
			)
			return
		}

		devicesCameraQualityAndRetention = ResponseCameraGetDeviceCameraQualityAndRetentionItemToBody(devicesCameraQualityAndRetention, response1)
		diags = resp.State.Set(ctx, &devicesCameraQualityAndRetention)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesCameraQualityAndRetention struct {
	Serial types.String                                      `tfsdk:"serial"`
	Item   *ResponseCameraGetDeviceCameraQualityAndRetention `tfsdk:"item"`
}

type ResponseCameraGetDeviceCameraQualityAndRetention struct {
	AudioRecordingEnabled          types.Bool   `tfsdk:"audio_recording_enabled"`
	MotionBasedRetentionEnabled    types.Bool   `tfsdk:"motion_based_retention_enabled"`
	MotionDetectorVersion          types.Int64  `tfsdk:"motion_detector_version"`
	ProfileID                      types.String `tfsdk:"profile_id"`
	Quality                        types.String `tfsdk:"quality"`
	Resolution                     types.String `tfsdk:"resolution"`
	RestrictedBandwidthModeEnabled types.Bool   `tfsdk:"restricted_bandwidth_mode_enabled"`
}

// ToBody
func ResponseCameraGetDeviceCameraQualityAndRetentionItemToBody(state DevicesCameraQualityAndRetention, response *merakigosdk.ResponseCameraGetDeviceCameraQualityAndRetention) DevicesCameraQualityAndRetention {
	itemState := ResponseCameraGetDeviceCameraQualityAndRetention{
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
	state.Item = &itemState
	return state
}
