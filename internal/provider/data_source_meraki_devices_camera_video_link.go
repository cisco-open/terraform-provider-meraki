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
	_ datasource.DataSource              = &DevicesCameraVideoLinkDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesCameraVideoLinkDataSource{}
)

func NewDevicesCameraVideoLinkDataSource() datasource.DataSource {
	return &DevicesCameraVideoLinkDataSource{}
}

type DevicesCameraVideoLinkDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesCameraVideoLinkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesCameraVideoLinkDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_camera_video_link"
}

func (d *DevicesCameraVideoLinkDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"timestamp": schema.StringAttribute{
				MarkdownDescription: `timestamp query parameter. [optional] The video link will start at this time. The timestamp should be a string in ISO8601 format. If no timestamp is specified, we will assume current time.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"url": schema.StringAttribute{
						Computed: true,
					},
					"vision_url": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *DevicesCameraVideoLinkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesCameraVideoLink DevicesCameraVideoLink
	diags := req.Config.Get(ctx, &devicesCameraVideoLink)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceCameraVideoLink")
		vvSerial := devicesCameraVideoLink.Serial.ValueString()
		queryParams1 := merakigosdk.GetDeviceCameraVideoLinkQueryParams{}

		queryParams1.Timestamp = devicesCameraVideoLink.Timestamp.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Camera.GetDeviceCameraVideoLink(vvSerial, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCameraVideoLink",
				err.Error(),
			)
			return
		}

		devicesCameraVideoLink = ResponseCameraGetDeviceCameraVideoLinkItemToBody(devicesCameraVideoLink, response1)
		diags = resp.State.Set(ctx, &devicesCameraVideoLink)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesCameraVideoLink struct {
	Serial    types.String                            `tfsdk:"serial"`
	Timestamp types.String                            `tfsdk:"timestamp"`
	Item      *ResponseCameraGetDeviceCameraVideoLink `tfsdk:"item"`
}

type ResponseCameraGetDeviceCameraVideoLink struct {
	URL       types.String `tfsdk:"url"`
	VisionURL types.String `tfsdk:"vision_url"`
}

// ToBody
func ResponseCameraGetDeviceCameraVideoLinkItemToBody(state DevicesCameraVideoLink, response *merakigosdk.ResponseCameraGetDeviceCameraVideoLink) DevicesCameraVideoLink {
	itemState := ResponseCameraGetDeviceCameraVideoLink{
		URL: func() types.String {
			if response.URL != "" {
				return types.StringValue(response.URL)
			}
			return types.String{}
		}(),
		VisionURL: func() types.String {
			if response.VisionURL != "" {
				return types.StringValue(response.VisionURL)
			}
			return types.String{}
		}(),
	}
	state.Item = &itemState
	return state
}
