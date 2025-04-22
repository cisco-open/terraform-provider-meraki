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
	_ datasource.DataSource              = &DevicesLiveToolsWakeOnLanDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesLiveToolsWakeOnLanDataSource{}
)

func NewDevicesLiveToolsWakeOnLanDataSource() datasource.DataSource {
	return &DevicesLiveToolsWakeOnLanDataSource{}
}

type DevicesLiveToolsWakeOnLanDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesLiveToolsWakeOnLanDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesLiveToolsWakeOnLanDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_live_tools_wake_on_lan"
}

func (d *DevicesLiveToolsWakeOnLanDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"wake_on_lan_id": schema.StringAttribute{
				MarkdownDescription: `wakeOnLanId path parameter. Wake on lan ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"error": schema.StringAttribute{
						MarkdownDescription: `An error message for a failed execution`,
						Computed:            true,
					},
					"request": schema.SingleNestedAttribute{
						MarkdownDescription: `The parameters of the Wake-on-LAN request`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"mac": schema.StringAttribute{
								MarkdownDescription: `The target's MAC address`,
								Computed:            true,
							},
							"serial": schema.StringAttribute{
								MarkdownDescription: `Device serial number`,
								Computed:            true,
							},
							"vlan_id": schema.Int64Attribute{
								MarkdownDescription: `The target's VLAN (1 to 4094)`,
								Computed:            true,
							},
						},
					},
					"status": schema.StringAttribute{
						MarkdownDescription: `Status of the Wake-on-LAN request`,
						Computed:            true,
					},
					"url": schema.StringAttribute{
						MarkdownDescription: `GET this url to check the status of your ping request`,
						Computed:            true,
					},
					"wake_on_lan_id": schema.StringAttribute{
						MarkdownDescription: `ID of the Wake-on-LAN job`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *DevicesLiveToolsWakeOnLanDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesLiveToolsWakeOnLan DevicesLiveToolsWakeOnLan
	diags := req.Config.Get(ctx, &devicesLiveToolsWakeOnLan)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceLiveToolsWakeOnLan")
		vvSerial := devicesLiveToolsWakeOnLan.Serial.ValueString()
		vvWakeOnLanID := devicesLiveToolsWakeOnLan.WakeOnLanID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Devices.GetDeviceLiveToolsWakeOnLan(vvSerial, vvWakeOnLanID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceLiveToolsWakeOnLan",
				err.Error(),
			)
			return
		}

		devicesLiveToolsWakeOnLan = ResponseDevicesGetDeviceLiveToolsWakeOnLanItemToBody(devicesLiveToolsWakeOnLan, response1)
		diags = resp.State.Set(ctx, &devicesLiveToolsWakeOnLan)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesLiveToolsWakeOnLan struct {
	Serial      types.String                                `tfsdk:"serial"`
	WakeOnLanID types.String                                `tfsdk:"wake_on_lan_id"`
	Item        *ResponseDevicesGetDeviceLiveToolsWakeOnLan `tfsdk:"item"`
}

type ResponseDevicesGetDeviceLiveToolsWakeOnLan struct {
	Error       types.String                                       `tfsdk:"error"`
	Request     *ResponseDevicesGetDeviceLiveToolsWakeOnLanRequest `tfsdk:"request"`
	Status      types.String                                       `tfsdk:"status"`
	URL         types.String                                       `tfsdk:"url"`
	WakeOnLanID types.String                                       `tfsdk:"wake_on_lan_id"`
}

type ResponseDevicesGetDeviceLiveToolsWakeOnLanRequest struct {
	Mac    types.String `tfsdk:"mac"`
	Serial types.String `tfsdk:"serial"`
	VLANID types.Int64  `tfsdk:"vlan_id"`
}

// ToBody
func ResponseDevicesGetDeviceLiveToolsWakeOnLanItemToBody(state DevicesLiveToolsWakeOnLan, response *merakigosdk.ResponseDevicesGetDeviceLiveToolsWakeOnLan) DevicesLiveToolsWakeOnLan {
	itemState := ResponseDevicesGetDeviceLiveToolsWakeOnLan{
		Error: types.StringValue(response.Error),
		Request: func() *ResponseDevicesGetDeviceLiveToolsWakeOnLanRequest {
			if response.Request != nil {
				return &ResponseDevicesGetDeviceLiveToolsWakeOnLanRequest{
					Mac:    types.StringValue(response.Request.Mac),
					Serial: types.StringValue(response.Request.Serial),
					VLANID: func() types.Int64 {
						if response.Request.VLANID != nil {
							return types.Int64Value(int64(*response.Request.VLANID))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
		Status:      types.StringValue(response.Status),
		URL:         types.StringValue(response.URL),
		WakeOnLanID: types.StringValue(response.WakeOnLanID),
	}
	state.Item = &itemState
	return state
}
