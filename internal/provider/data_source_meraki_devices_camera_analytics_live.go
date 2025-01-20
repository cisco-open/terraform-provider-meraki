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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &DevicesCameraAnalyticsLiveDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesCameraAnalyticsLiveDataSource{}
)

func NewDevicesCameraAnalyticsLiveDataSource() datasource.DataSource {
	return &DevicesCameraAnalyticsLiveDataSource{}
}

type DevicesCameraAnalyticsLiveDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesCameraAnalyticsLiveDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesCameraAnalyticsLiveDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_camera_analytics_live"
}

func (d *DevicesCameraAnalyticsLiveDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"ts": schema.StringAttribute{
						MarkdownDescription: `The current time`,
						Computed:            true,
					},
					"zones": schema.SingleNestedAttribute{
						MarkdownDescription: `The zones state`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"zone_id": schema.SingleNestedAttribute{
								MarkdownDescription: `The zone state, dynamic`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"person": schema.Int64Attribute{
										MarkdownDescription: `The count per type, dynamic`,
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *DevicesCameraAnalyticsLiveDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesCameraAnalyticsLive DevicesCameraAnalyticsLive
	diags := req.Config.Get(ctx, &devicesCameraAnalyticsLive)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceCameraAnalyticsLive")
		vvSerial := devicesCameraAnalyticsLive.Serial.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Camera.GetDeviceCameraAnalyticsLive(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCameraAnalyticsLive",
				err.Error(),
			)
			return
		}

		devicesCameraAnalyticsLive = ResponseCameraGetDeviceCameraAnalyticsLiveItemToBody(devicesCameraAnalyticsLive, response1)
		diags = resp.State.Set(ctx, &devicesCameraAnalyticsLive)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesCameraAnalyticsLive struct {
	Serial types.String                                `tfsdk:"serial"`
	Item   *ResponseCameraGetDeviceCameraAnalyticsLive `tfsdk:"item"`
}

type ResponseCameraGetDeviceCameraAnalyticsLive struct {
	Ts    types.String                                     `tfsdk:"ts"`
	Zones *ResponseCameraGetDeviceCameraAnalyticsLiveZones `tfsdk:"zones"`
}

type ResponseCameraGetDeviceCameraAnalyticsLiveZones struct {
	ZoneID *ResponseCameraGetDeviceCameraAnalyticsLiveZonesZoneId `tfsdk:"zone_id"`
}

type ResponseCameraGetDeviceCameraAnalyticsLiveZonesZoneId struct {
	Person types.Int64 `tfsdk:"person"`
}

// ToBody
func ResponseCameraGetDeviceCameraAnalyticsLiveItemToBody(state DevicesCameraAnalyticsLive, response *merakigosdk.ResponseCameraGetDeviceCameraAnalyticsLive) DevicesCameraAnalyticsLive {
	itemState := ResponseCameraGetDeviceCameraAnalyticsLive{
		Ts: types.StringValue(response.Ts),
		Zones: func() *ResponseCameraGetDeviceCameraAnalyticsLiveZones {
			if response.Zones != nil {
				return &ResponseCameraGetDeviceCameraAnalyticsLiveZones{
					ZoneID: func() *ResponseCameraGetDeviceCameraAnalyticsLiveZonesZoneId {
						if response.Zones.ZoneID != nil {
							return &ResponseCameraGetDeviceCameraAnalyticsLiveZonesZoneId{
								Person: func() types.Int64 {
									if response.Zones.ZoneID.Person != nil {
										return types.Int64Value(int64(*response.Zones.ZoneID.Person))
									}
									return types.Int64{}
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
