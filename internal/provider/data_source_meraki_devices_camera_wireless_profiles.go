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
	_ datasource.DataSource              = &DevicesCameraWirelessProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesCameraWirelessProfilesDataSource{}
)

func NewDevicesCameraWirelessProfilesDataSource() datasource.DataSource {
	return &DevicesCameraWirelessProfilesDataSource{}
}

type DevicesCameraWirelessProfilesDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesCameraWirelessProfilesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesCameraWirelessProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_camera_wireless_profiles"
}

func (d *DevicesCameraWirelessProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"ids": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"backup": schema.StringAttribute{
								Computed: true,
							},
							"primary": schema.StringAttribute{
								Computed: true,
							},
							"secondary": schema.StringAttribute{
								Computed: true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *DevicesCameraWirelessProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesCameraWirelessProfiles DevicesCameraWirelessProfiles
	diags := req.Config.Get(ctx, &devicesCameraWirelessProfiles)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceCameraWirelessProfiles")
		vvSerial := devicesCameraWirelessProfiles.Serial.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Camera.GetDeviceCameraWirelessProfiles(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCameraWirelessProfiles",
				err.Error(),
			)
			return
		}

		devicesCameraWirelessProfiles = ResponseCameraGetDeviceCameraWirelessProfilesItemToBody(devicesCameraWirelessProfiles, response1)
		diags = resp.State.Set(ctx, &devicesCameraWirelessProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesCameraWirelessProfiles struct {
	Serial types.String                                   `tfsdk:"serial"`
	Item   *ResponseCameraGetDeviceCameraWirelessProfiles `tfsdk:"item"`
}

type ResponseCameraGetDeviceCameraWirelessProfiles struct {
	IDs *ResponseCameraGetDeviceCameraWirelessProfilesIds `tfsdk:"ids"`
}

type ResponseCameraGetDeviceCameraWirelessProfilesIds struct {
	Backup    types.String `tfsdk:"backup"`
	Primary   types.String `tfsdk:"primary"`
	Secondary types.String `tfsdk:"secondary"`
}

// ToBody
func ResponseCameraGetDeviceCameraWirelessProfilesItemToBody(state DevicesCameraWirelessProfiles, response *merakigosdk.ResponseCameraGetDeviceCameraWirelessProfiles) DevicesCameraWirelessProfiles {
	itemState := ResponseCameraGetDeviceCameraWirelessProfiles{
		IDs: func() *ResponseCameraGetDeviceCameraWirelessProfilesIds {
			if response.IDs != nil {
				return &ResponseCameraGetDeviceCameraWirelessProfilesIds{
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
	state.Item = &itemState
	return state
}
