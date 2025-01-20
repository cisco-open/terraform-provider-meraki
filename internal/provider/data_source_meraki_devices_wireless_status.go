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
	_ datasource.DataSource              = &DevicesWirelessStatusDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesWirelessStatusDataSource{}
)

func NewDevicesWirelessStatusDataSource() datasource.DataSource {
	return &DevicesWirelessStatusDataSource{}
}

type DevicesWirelessStatusDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesWirelessStatusDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesWirelessStatusDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_wireless_status"
}

func (d *DevicesWirelessStatusDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"basic_service_sets": schema.SetNestedAttribute{
						MarkdownDescription: `SSID status list`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"band": schema.StringAttribute{
									MarkdownDescription: `Frequency range used by wireless network`,
									Computed:            true,
								},
								"broadcasting": schema.BoolAttribute{
									MarkdownDescription: `Whether the SSID is broadcasting based on an availability schedule`,
									Computed:            true,
								},
								"bssid": schema.StringAttribute{
									MarkdownDescription: `Unique identifier of wireless access point`,
									Computed:            true,
								},
								"channel": schema.Int64Attribute{
									MarkdownDescription: `Frequency channel used by wireless network`,
									Computed:            true,
								},
								"channel_width": schema.StringAttribute{
									MarkdownDescription: `Width of frequency channel used by wireless network`,
									Computed:            true,
								},
								"enabled": schema.BoolAttribute{
									MarkdownDescription: `Status of wireless network`,
									Computed:            true,
								},
								"power": schema.StringAttribute{
									MarkdownDescription: `Strength of wireless signal`,
									Computed:            true,
								},
								"ssid_name": schema.StringAttribute{
									MarkdownDescription: `Name of wireless network`,
									Computed:            true,
								},
								"ssid_number": schema.Int64Attribute{
									MarkdownDescription: `Unique identifier of wireless network`,
									Computed:            true,
								},
								"visible": schema.BoolAttribute{
									MarkdownDescription: `Whether the SSID is advertised or hidden`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *DevicesWirelessStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesWirelessStatus DevicesWirelessStatus
	diags := req.Config.Get(ctx, &devicesWirelessStatus)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceWirelessStatus")
		vvSerial := devicesWirelessStatus.Serial.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetDeviceWirelessStatus(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceWirelessStatus",
				err.Error(),
			)
			return
		}

		devicesWirelessStatus = ResponseWirelessGetDeviceWirelessStatusItemToBody(devicesWirelessStatus, response1)
		diags = resp.State.Set(ctx, &devicesWirelessStatus)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesWirelessStatus struct {
	Serial types.String                             `tfsdk:"serial"`
	Item   *ResponseWirelessGetDeviceWirelessStatus `tfsdk:"item"`
}

type ResponseWirelessGetDeviceWirelessStatus struct {
	BasicServiceSets *[]ResponseWirelessGetDeviceWirelessStatusBasicServiceSets `tfsdk:"basic_service_sets"`
}

type ResponseWirelessGetDeviceWirelessStatusBasicServiceSets struct {
	Band         types.String `tfsdk:"band"`
	Broadcasting types.Bool   `tfsdk:"broadcasting"`
	Bssid        types.String `tfsdk:"bssid"`
	Channel      types.Int64  `tfsdk:"channel"`
	ChannelWidth types.String `tfsdk:"channel_width"`
	Enabled      types.Bool   `tfsdk:"enabled"`
	Power        types.String `tfsdk:"power"`
	SSIDName     types.String `tfsdk:"ssid_name"`
	SSIDNumber   types.Int64  `tfsdk:"ssid_number"`
	Visible      types.Bool   `tfsdk:"visible"`
}

// ToBody
func ResponseWirelessGetDeviceWirelessStatusItemToBody(state DevicesWirelessStatus, response *merakigosdk.ResponseWirelessGetDeviceWirelessStatus) DevicesWirelessStatus {
	itemState := ResponseWirelessGetDeviceWirelessStatus{
		BasicServiceSets: func() *[]ResponseWirelessGetDeviceWirelessStatusBasicServiceSets {
			if response.BasicServiceSets != nil {
				result := make([]ResponseWirelessGetDeviceWirelessStatusBasicServiceSets, len(*response.BasicServiceSets))
				for i, basicServiceSets := range *response.BasicServiceSets {
					result[i] = ResponseWirelessGetDeviceWirelessStatusBasicServiceSets{
						Band: types.StringValue(basicServiceSets.Band),
						Broadcasting: func() types.Bool {
							if basicServiceSets.Broadcasting != nil {
								return types.BoolValue(*basicServiceSets.Broadcasting)
							}
							return types.Bool{}
						}(),
						Bssid: types.StringValue(basicServiceSets.Bssid),
						Channel: func() types.Int64 {
							if basicServiceSets.Channel != nil {
								return types.Int64Value(int64(*basicServiceSets.Channel))
							}
							return types.Int64{}
						}(),
						ChannelWidth: types.StringValue(basicServiceSets.ChannelWidth),
						Enabled: func() types.Bool {
							if basicServiceSets.Enabled != nil {
								return types.BoolValue(*basicServiceSets.Enabled)
							}
							return types.Bool{}
						}(),
						Power:    types.StringValue(basicServiceSets.Power),
						SSIDName: types.StringValue(basicServiceSets.SSIDName),
						SSIDNumber: func() types.Int64 {
							if basicServiceSets.SSIDNumber != nil {
								return types.Int64Value(int64(*basicServiceSets.SSIDNumber))
							}
							return types.Int64{}
						}(),
						Visible: func() types.Bool {
							if basicServiceSets.Visible != nil {
								return types.BoolValue(*basicServiceSets.Visible)
							}
							return types.Bool{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
