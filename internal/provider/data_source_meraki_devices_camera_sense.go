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
	_ datasource.DataSource              = &DevicesCameraSenseDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesCameraSenseDataSource{}
)

func NewDevicesCameraSenseDataSource() datasource.DataSource {
	return &DevicesCameraSenseDataSource{}
}

type DevicesCameraSenseDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesCameraSenseDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesCameraSenseDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_camera_sense"
}

func (d *DevicesCameraSenseDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"audio_detection": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								Computed: true,
							},
						},
					},
					"mqtt_broker_id": schema.StringAttribute{
						Computed: true,
					},
					"mqtt_topics": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"sense_enabled": schema.BoolAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *DevicesCameraSenseDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesCameraSense DevicesCameraSense
	diags := req.Config.Get(ctx, &devicesCameraSense)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceCameraSense")
		vvSerial := devicesCameraSense.Serial.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Camera.GetDeviceCameraSense(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCameraSense",
				err.Error(),
			)
			return
		}

		devicesCameraSense = ResponseCameraGetDeviceCameraSenseItemToBody(devicesCameraSense, response1)
		diags = resp.State.Set(ctx, &devicesCameraSense)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesCameraSense struct {
	Serial types.String                        `tfsdk:"serial"`
	Item   *ResponseCameraGetDeviceCameraSense `tfsdk:"item"`
}

type ResponseCameraGetDeviceCameraSense struct {
	AudioDetection *ResponseCameraGetDeviceCameraSenseAudioDetection `tfsdk:"audio_detection"`
	MqttBrokerID   types.String                                      `tfsdk:"mqtt_broker_id"`
	MqttTopics     types.List                                        `tfsdk:"mqtt_topics"`
	SenseEnabled   types.Bool                                        `tfsdk:"sense_enabled"`
}

type ResponseCameraGetDeviceCameraSenseAudioDetection struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// ToBody
func ResponseCameraGetDeviceCameraSenseItemToBody(state DevicesCameraSense, response *merakigosdk.ResponseCameraGetDeviceCameraSense) DevicesCameraSense {
	itemState := ResponseCameraGetDeviceCameraSense{
		AudioDetection: func() *ResponseCameraGetDeviceCameraSenseAudioDetection {
			if response.AudioDetection != nil {
				return &ResponseCameraGetDeviceCameraSenseAudioDetection{
					Enabled: func() types.Bool {
						if response.AudioDetection.Enabled != nil {
							return types.BoolValue(*response.AudioDetection.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
		}(),
		MqttBrokerID: func() types.String {
			if response.MqttBrokerID != "" {
				return types.StringValue(response.MqttBrokerID)
			}
			return types.String{}
		}(),
		MqttTopics: StringSliceToList(response.MqttTopics),
		SenseEnabled: func() types.Bool {
			if response.SenseEnabled != nil {
				return types.BoolValue(*response.SenseEnabled)
			}
			return types.Bool{}
		}(),
	}
	state.Item = &itemState
	return state
}
