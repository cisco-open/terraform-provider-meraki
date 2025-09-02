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
	_ datasource.DataSource              = &DevicesSwitchWarmSpareDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesSwitchWarmSpareDataSource{}
)

func NewDevicesSwitchWarmSpareDataSource() datasource.DataSource {
	return &DevicesSwitchWarmSpareDataSource{}
}

type DevicesSwitchWarmSpareDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesSwitchWarmSpareDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesSwitchWarmSpareDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_switch_warm_spare"
}

func (d *DevicesSwitchWarmSpareDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Enable or disable warm spare for a switch`,
						Computed:            true,
					},
					"primary_serial": schema.StringAttribute{
						MarkdownDescription: `Serial number of the primary switch`,
						Computed:            true,
					},
					"spare_serial": schema.StringAttribute{
						MarkdownDescription: `Serial number of the warm spare switch`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *DevicesSwitchWarmSpareDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesSwitchWarmSpare DevicesSwitchWarmSpare
	diags := req.Config.Get(ctx, &devicesSwitchWarmSpare)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceSwitchWarmSpare")
		vvSerial := devicesSwitchWarmSpare.Serial.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetDeviceSwitchWarmSpare(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceSwitchWarmSpare",
				err.Error(),
			)
			return
		}

		devicesSwitchWarmSpare = ResponseSwitchGetDeviceSwitchWarmSpareItemToBody(devicesSwitchWarmSpare, response1)
		diags = resp.State.Set(ctx, &devicesSwitchWarmSpare)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesSwitchWarmSpare struct {
	Serial types.String                            `tfsdk:"serial"`
	Item   *ResponseSwitchGetDeviceSwitchWarmSpare `tfsdk:"item"`
}

type ResponseSwitchGetDeviceSwitchWarmSpare struct {
	Enabled       types.Bool   `tfsdk:"enabled"`
	PrimarySerial types.String `tfsdk:"primary_serial"`
	SpareSerial   types.String `tfsdk:"spare_serial"`
}

// ToBody
func ResponseSwitchGetDeviceSwitchWarmSpareItemToBody(state DevicesSwitchWarmSpare, response *merakigosdk.ResponseSwitchGetDeviceSwitchWarmSpare) DevicesSwitchWarmSpare {
	itemState := ResponseSwitchGetDeviceSwitchWarmSpare{
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		PrimarySerial: func() types.String {
			if response.PrimarySerial != "" {
				return types.StringValue(response.PrimarySerial)
			}
			return types.String{}
		}(),
		SpareSerial: func() types.String {
			if response.SpareSerial != "" {
				return types.StringValue(response.SpareSerial)
			}
			return types.String{}
		}(),
	}
	state.Item = &itemState
	return state
}
