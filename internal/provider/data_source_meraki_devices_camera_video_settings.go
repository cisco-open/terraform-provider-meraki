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

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &DevicesCameraVideoSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesCameraVideoSettingsDataSource{}
)

func NewDevicesCameraVideoSettingsDataSource() datasource.DataSource {
	return &DevicesCameraVideoSettingsDataSource{}
}

type DevicesCameraVideoSettingsDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesCameraVideoSettingsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesCameraVideoSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_camera_video_settings"
}

func (d *DevicesCameraVideoSettingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"external_rtsp_enabled": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating if external rtsp stream is exposed`,
						Computed:            true,
					},
					"rtsp_url": schema.StringAttribute{
						MarkdownDescription: `External rstp url. Will only be returned if external rtsp stream is exposed`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *DevicesCameraVideoSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesCameraVideoSettings DevicesCameraVideoSettings
	diags := req.Config.Get(ctx, &devicesCameraVideoSettings)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceCameraVideoSettings")
		vvSerial := devicesCameraVideoSettings.Serial.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Camera.GetDeviceCameraVideoSettings(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCameraVideoSettings",
				err.Error(),
			)
			return
		}

		devicesCameraVideoSettings = ResponseCameraGetDeviceCameraVideoSettingsItemToBody(devicesCameraVideoSettings, response1)
		diags = resp.State.Set(ctx, &devicesCameraVideoSettings)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesCameraVideoSettings struct {
	Serial types.String                                `tfsdk:"serial"`
	Item   *ResponseCameraGetDeviceCameraVideoSettings `tfsdk:"item"`
}

type ResponseCameraGetDeviceCameraVideoSettings struct {
	ExternalRtspEnabled types.Bool   `tfsdk:"external_rtsp_enabled"`
	RtspURL             types.String `tfsdk:"rtsp_url"`
}

// ToBody
func ResponseCameraGetDeviceCameraVideoSettingsItemToBody(state DevicesCameraVideoSettings, response *merakigosdk.ResponseCameraGetDeviceCameraVideoSettings) DevicesCameraVideoSettings {
	itemState := ResponseCameraGetDeviceCameraVideoSettings{
		ExternalRtspEnabled: func() types.Bool {
			if response.ExternalRtspEnabled != nil {
				return types.BoolValue(*response.ExternalRtspEnabled)
			}
			return types.Bool{}
		}(),
		RtspURL: types.StringValue(response.RtspURL),
	}
	state.Item = &itemState
	return state
}
