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
	_ datasource.DataSource              = &DevicesWirelessBluetoothSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesWirelessBluetoothSettingsDataSource{}
)

func NewDevicesWirelessBluetoothSettingsDataSource() datasource.DataSource {
	return &DevicesWirelessBluetoothSettingsDataSource{}
}

type DevicesWirelessBluetoothSettingsDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesWirelessBluetoothSettingsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesWirelessBluetoothSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_wireless_bluetooth_settings"
}

func (d *DevicesWirelessBluetoothSettingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"major": schema.Int64Attribute{
						MarkdownDescription: `Desired major value of the beacon. If the value is set to null it will reset to
          Dashboard's automatically generated value.`,
						Computed: true,
					},
					"minor": schema.Int64Attribute{
						MarkdownDescription: `Desired minor value of the beacon. If the value is set to null it will reset to
          Dashboard's automatically generated value.`,
						Computed: true,
					},
					"uuid": schema.StringAttribute{
						MarkdownDescription: `Desired UUID of the beacon. If the value is set to null it will reset to Dashboard's
          automatically generated value.`,
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *DevicesWirelessBluetoothSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesWirelessBluetoothSettings DevicesWirelessBluetoothSettings
	diags := req.Config.Get(ctx, &devicesWirelessBluetoothSettings)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceWirelessBluetoothSettings")
		vvSerial := devicesWirelessBluetoothSettings.Serial.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetDeviceWirelessBluetoothSettings(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceWirelessBluetoothSettings",
				err.Error(),
			)
			return
		}

		devicesWirelessBluetoothSettings = ResponseWirelessGetDeviceWirelessBluetoothSettingsItemToBody(devicesWirelessBluetoothSettings, response1)
		diags = resp.State.Set(ctx, &devicesWirelessBluetoothSettings)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesWirelessBluetoothSettings struct {
	Serial types.String                                        `tfsdk:"serial"`
	Item   *ResponseWirelessGetDeviceWirelessBluetoothSettings `tfsdk:"item"`
}

type ResponseWirelessGetDeviceWirelessBluetoothSettings struct {
	Major types.Int64  `tfsdk:"major"`
	Minor types.Int64  `tfsdk:"minor"`
	UUID  types.String `tfsdk:"uuid"`
}

// ToBody
func ResponseWirelessGetDeviceWirelessBluetoothSettingsItemToBody(state DevicesWirelessBluetoothSettings, response *merakigosdk.ResponseWirelessGetDeviceWirelessBluetoothSettings) DevicesWirelessBluetoothSettings {
	itemState := ResponseWirelessGetDeviceWirelessBluetoothSettings{
		Major: func() types.Int64 {
			if response.Major != nil {
				return types.Int64Value(int64(*response.Major))
			}
			return types.Int64{}
		}(),
		Minor: func() types.Int64 {
			if response.Minor != nil {
				return types.Int64Value(int64(*response.Minor))
			}
			return types.Int64{}
		}(),
		UUID: func() types.String {
			if response.UUID != "" {
				return types.StringValue(response.UUID)
			}
			return types.String{}
		}(),
	}
	state.Item = &itemState
	return state
}
