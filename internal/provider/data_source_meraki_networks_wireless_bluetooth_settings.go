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
	_ datasource.DataSource              = &NetworksWirelessBluetoothSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessBluetoothSettingsDataSource{}
)

func NewNetworksWirelessBluetoothSettingsDataSource() datasource.DataSource {
	return &NetworksWirelessBluetoothSettingsDataSource{}
}

type NetworksWirelessBluetoothSettingsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessBluetoothSettingsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessBluetoothSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_bluetooth_settings"
}

func (d *NetworksWirelessBluetoothSettingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"advertising_enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether APs will advertise beacons.`,
						Computed:            true,
					},
					"esl_enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether ESL is enabled on this network.`,
						Computed:            true,
					},
					"major": schema.Int64Attribute{
						MarkdownDescription: `The major number to be used in the beacon identifier. Only valid in 'Non-unique' mode.`,
						Computed:            true,
					},
					"major_minor_assignment_mode": schema.StringAttribute{
						MarkdownDescription: `The way major and minor number should be assigned to nodes in the network. ('Unique', 'Non-unique')`,
						Computed:            true,
					},
					"minor": schema.Int64Attribute{
						MarkdownDescription: `The minor number to be used in the beacon identifier. Only valid in 'Non-unique' mode.`,
						Computed:            true,
					},
					"scanning_enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether APs will scan for Bluetooth enabled clients.`,
						Computed:            true,
					},
					"uuid": schema.StringAttribute{
						MarkdownDescription: `The UUID to be used in the beacon identifier.`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessBluetoothSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessBluetoothSettings NetworksWirelessBluetoothSettings
	diags := req.Config.Get(ctx, &networksWirelessBluetoothSettings)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessBluetoothSettings")
		vvNetworkID := networksWirelessBluetoothSettings.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessBluetoothSettings(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessBluetoothSettings",
				err.Error(),
			)
			return
		}

		networksWirelessBluetoothSettings = ResponseWirelessGetNetworkWirelessBluetoothSettingsItemToBody(networksWirelessBluetoothSettings, response1)
		diags = resp.State.Set(ctx, &networksWirelessBluetoothSettings)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessBluetoothSettings struct {
	NetworkID types.String                                         `tfsdk:"network_id"`
	Item      *ResponseWirelessGetNetworkWirelessBluetoothSettings `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessBluetoothSettings struct {
	AdvertisingEnabled       types.Bool   `tfsdk:"advertising_enabled"`
	EslEnabled               types.Bool   `tfsdk:"esl_enabled"`
	Major                    types.Int64  `tfsdk:"major"`
	MajorMinorAssignmentMode types.String `tfsdk:"major_minor_assignment_mode"`
	Minor                    types.Int64  `tfsdk:"minor"`
	ScanningEnabled          types.Bool   `tfsdk:"scanning_enabled"`
	UUID                     types.String `tfsdk:"uuid"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessBluetoothSettingsItemToBody(state NetworksWirelessBluetoothSettings, response *merakigosdk.ResponseWirelessGetNetworkWirelessBluetoothSettings) NetworksWirelessBluetoothSettings {
	itemState := ResponseWirelessGetNetworkWirelessBluetoothSettings{
		AdvertisingEnabled: func() types.Bool {
			if response.AdvertisingEnabled != nil {
				return types.BoolValue(*response.AdvertisingEnabled)
			}
			return types.Bool{}
		}(),
		EslEnabled: func() types.Bool {
			if response.EslEnabled != nil {
				return types.BoolValue(*response.EslEnabled)
			}
			return types.Bool{}
		}(),
		Major: func() types.Int64 {
			if response.Major != nil {
				return types.Int64Value(int64(*response.Major))
			}
			return types.Int64{}
		}(),
		MajorMinorAssignmentMode: types.StringValue(response.MajorMinorAssignmentMode),
		Minor: func() types.Int64 {
			if response.Minor != nil {
				return types.Int64Value(int64(*response.Minor))
			}
			return types.Int64{}
		}(),
		ScanningEnabled: func() types.Bool {
			if response.ScanningEnabled != nil {
				return types.BoolValue(*response.ScanningEnabled)
			}
			return types.Bool{}
		}(),
		UUID: types.StringValue(response.UUID),
	}
	state.Item = &itemState
	return state
}
