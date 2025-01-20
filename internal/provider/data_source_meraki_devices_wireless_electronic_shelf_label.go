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
	_ datasource.DataSource              = &DevicesWirelessElectronicShelfLabelDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesWirelessElectronicShelfLabelDataSource{}
)

func NewDevicesWirelessElectronicShelfLabelDataSource() datasource.DataSource {
	return &DevicesWirelessElectronicShelfLabelDataSource{}
}

type DevicesWirelessElectronicShelfLabelDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesWirelessElectronicShelfLabelDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesWirelessElectronicShelfLabelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_wireless_electronic_shelf_label"
}

func (d *DevicesWirelessElectronicShelfLabelDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"ap_esl_id": schema.Int64Attribute{
						MarkdownDescription: `An identifier for the device used by the ESL system`,
						Computed:            true,
					},
					"channel": schema.StringAttribute{
						MarkdownDescription: `Desired ESL channel for the device, or 'Auto' (case insensitive) to use the recommended channel`,
						Computed:            true,
					},
					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Turn ESL features on and off for this device`,
						Computed:            true,
					},
					"hostname": schema.StringAttribute{
						MarkdownDescription: `Hostname of the ESL management service`,
						Computed:            true,
					},
					"network_id": schema.StringAttribute{
						MarkdownDescription: `The identifier for the device's network`,
						Computed:            true,
					},
					"provider_r": schema.StringAttribute{
						MarkdownDescription: `The service providing ESL functionality`,
						Computed:            true,
					},
					"serial": schema.StringAttribute{
						MarkdownDescription: `The serial number of the device`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *DevicesWirelessElectronicShelfLabelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesWirelessElectronicShelfLabel DevicesWirelessElectronicShelfLabel
	diags := req.Config.Get(ctx, &devicesWirelessElectronicShelfLabel)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceWirelessElectronicShelfLabel")
		vvSerial := devicesWirelessElectronicShelfLabel.Serial.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetDeviceWirelessElectronicShelfLabel(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceWirelessElectronicShelfLabel",
				err.Error(),
			)
			return
		}

		devicesWirelessElectronicShelfLabel = ResponseWirelessGetDeviceWirelessElectronicShelfLabelItemToBody(devicesWirelessElectronicShelfLabel, response1)
		diags = resp.State.Set(ctx, &devicesWirelessElectronicShelfLabel)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesWirelessElectronicShelfLabel struct {
	Serial types.String                                           `tfsdk:"serial"`
	Item   *ResponseWirelessGetDeviceWirelessElectronicShelfLabel `tfsdk:"item"`
}

type ResponseWirelessGetDeviceWirelessElectronicShelfLabel struct {
	ApEslID   types.Int64  `tfsdk:"ap_esl_id"`
	Channel   types.String `tfsdk:"channel"`
	Enabled   types.Bool   `tfsdk:"enabled"`
	Hostname  types.String `tfsdk:"hostname"`
	NetworkID types.String `tfsdk:"network_id"`
	Provider  types.String `tfsdk:"provider_r"`
	Serial    types.String `tfsdk:"serial"`
}

// ToBody
func ResponseWirelessGetDeviceWirelessElectronicShelfLabelItemToBody(state DevicesWirelessElectronicShelfLabel, response *merakigosdk.ResponseWirelessGetDeviceWirelessElectronicShelfLabel) DevicesWirelessElectronicShelfLabel {
	itemState := ResponseWirelessGetDeviceWirelessElectronicShelfLabel{
		ApEslID: func() types.Int64 {
			if response.ApEslID != nil {
				return types.Int64Value(int64(*response.ApEslID))
			}
			return types.Int64{}
		}(),
		Channel: types.StringValue(response.Channel),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		Hostname:  types.StringValue(response.Hostname),
		NetworkID: types.StringValue(response.NetworkID),
		Provider:  types.StringValue(response.Provider),
		Serial:    types.StringValue(response.Serial),
	}
	state.Item = &itemState
	return state
}
