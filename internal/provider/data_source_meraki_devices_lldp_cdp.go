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
	_ datasource.DataSource              = &DevicesLldpCdpDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesLldpCdpDataSource{}
)

func NewDevicesLldpCdpDataSource() datasource.DataSource {
	return &DevicesLldpCdpDataSource{}
}

type DevicesLldpCdpDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesLldpCdpDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesLldpCdpDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_lldp_cdp"
}

func (d *DevicesLldpCdpDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"ports": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"status_12": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"cdp": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{

											"address": schema.StringAttribute{
												Computed: true,
											},
											"device_id": schema.StringAttribute{
												Computed: true,
											},
											"port_id": schema.StringAttribute{
												Computed: true,
											},
											"source_port": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"lldp": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{

											"management_address": schema.StringAttribute{
												Computed: true,
											},
											"port_id": schema.StringAttribute{
												Computed: true,
											},
											"source_port": schema.StringAttribute{
												Computed: true,
											},
											"system_name": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
							},
							"status_8": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"cdp": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{

											"address": schema.StringAttribute{
												Computed: true,
											},
											"device_id": schema.StringAttribute{
												Computed: true,
											},
											"port_id": schema.StringAttribute{
												Computed: true,
											},
											"source_port": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
							},
						},
					},
					"source_mac": schema.StringAttribute{
						MarkdownDescription: `Source MAC address`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *DevicesLldpCdpDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesLldpCdp DevicesLldpCdp
	diags := req.Config.Get(ctx, &devicesLldpCdp)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceLldpCdp")
		vvSerial := devicesLldpCdp.Serial.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Devices.GetDeviceLldpCdp(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceLldpCdp",
				err.Error(),
			)
			return
		}

		devicesLldpCdp = ResponseDevicesGetDeviceLldpCdpItemToBody(devicesLldpCdp, response1)
		diags = resp.State.Set(ctx, &devicesLldpCdp)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesLldpCdp struct {
	Serial types.String                     `tfsdk:"serial"`
	Item   *ResponseDevicesGetDeviceLldpCdp `tfsdk:"item"`
}

type ResponseDevicesGetDeviceLldpCdp struct {
	Ports     *ResponseDevicesGetDeviceLldpCdpPorts `tfsdk:"ports"`
	SourceMac types.String                          `tfsdk:"source_mac"`
}

type ResponseDevicesGetDeviceLldpCdpPorts struct {
	Status12 *ResponseDevicesGetDeviceLldpCdpPorts12 `tfsdk:"status_12"`
	Status8  *ResponseDevicesGetDeviceLldpCdpPorts8  `tfsdk:"status_8"`
}

type ResponseDevicesGetDeviceLldpCdpPorts12 struct {
	Cdp  *ResponseDevicesGetDeviceLldpCdpPorts12Cdp  `tfsdk:"cdp"`
	Lldp *ResponseDevicesGetDeviceLldpCdpPorts12Lldp `tfsdk:"lldp"`
}

type ResponseDevicesGetDeviceLldpCdpPorts12Cdp struct {
	Address    types.String `tfsdk:"address"`
	DeviceID   types.String `tfsdk:"device_id"`
	PortID     types.String `tfsdk:"port_id"`
	SourcePort types.String `tfsdk:"source_port"`
}

type ResponseDevicesGetDeviceLldpCdpPorts12Lldp struct {
	ManagementAddress types.String `tfsdk:"management_address"`
	PortID            types.String `tfsdk:"port_id"`
	SourcePort        types.String `tfsdk:"source_port"`
	SystemName        types.String `tfsdk:"system_name"`
}

type ResponseDevicesGetDeviceLldpCdpPorts8 struct {
	Cdp *ResponseDevicesGetDeviceLldpCdpPorts8Cdp `tfsdk:"cdp"`
}

type ResponseDevicesGetDeviceLldpCdpPorts8Cdp struct {
	Address    types.String `tfsdk:"address"`
	DeviceID   types.String `tfsdk:"device_id"`
	PortID     types.String `tfsdk:"port_id"`
	SourcePort types.String `tfsdk:"source_port"`
}

// ToBody
func ResponseDevicesGetDeviceLldpCdpItemToBody(state DevicesLldpCdp, response *merakigosdk.ResponseDevicesGetDeviceLldpCdp) DevicesLldpCdp {
	itemState := ResponseDevicesGetDeviceLldpCdp{
		Ports: func() *ResponseDevicesGetDeviceLldpCdpPorts {
			if response.Ports != nil {
				return &ResponseDevicesGetDeviceLldpCdpPorts{
					Status12: func() *ResponseDevicesGetDeviceLldpCdpPorts12 {
						if response.Ports.Status12 != nil {
							return &ResponseDevicesGetDeviceLldpCdpPorts12{
								Cdp: func() *ResponseDevicesGetDeviceLldpCdpPorts12Cdp {
									if response.Ports.Status12.Cdp != nil {
										return &ResponseDevicesGetDeviceLldpCdpPorts12Cdp{
											Address: func() types.String {
												if response.Ports.Status12.Cdp.Address != "" {
													return types.StringValue(response.Ports.Status12.Cdp.Address)
												}
												return types.String{}
											}(),
											DeviceID: func() types.String {
												if response.Ports.Status12.Cdp.DeviceID != "" {
													return types.StringValue(response.Ports.Status12.Cdp.DeviceID)
												}
												return types.String{}
											}(),
											PortID: func() types.String {
												if response.Ports.Status12.Cdp.PortID != "" {
													return types.StringValue(response.Ports.Status12.Cdp.PortID)
												}
												return types.String{}
											}(),
											SourcePort: func() types.String {
												if response.Ports.Status12.Cdp.SourcePort != "" {
													return types.StringValue(response.Ports.Status12.Cdp.SourcePort)
												}
												return types.String{}
											}(),
										}
									}
									return &ResponseDevicesGetDeviceLldpCdpPorts12Cdp{}
								}(),
								Lldp: func() *ResponseDevicesGetDeviceLldpCdpPorts12Lldp {
									if response.Ports.Status12.Lldp != nil {
										return &ResponseDevicesGetDeviceLldpCdpPorts12Lldp{
											ManagementAddress: func() types.String {
												if response.Ports.Status12.Lldp.ManagementAddress != "" {
													return types.StringValue(response.Ports.Status12.Lldp.ManagementAddress)
												}
												return types.String{}
											}(),
											PortID: func() types.String {
												if response.Ports.Status12.Lldp.PortID != "" {
													return types.StringValue(response.Ports.Status12.Lldp.PortID)
												}
												return types.String{}
											}(),
											SourcePort: func() types.String {
												if response.Ports.Status12.Lldp.SourcePort != "" {
													return types.StringValue(response.Ports.Status12.Lldp.SourcePort)
												}
												return types.String{}
											}(),
											SystemName: func() types.String {
												if response.Ports.Status12.Lldp.SystemName != "" {
													return types.StringValue(response.Ports.Status12.Lldp.SystemName)
												}
												return types.String{}
											}(),
										}
									}
									return &ResponseDevicesGetDeviceLldpCdpPorts12Lldp{}
								}(),
							}
						}
						return &ResponseDevicesGetDeviceLldpCdpPorts12{}
					}(),
					Status8: func() *ResponseDevicesGetDeviceLldpCdpPorts8 {
						if response.Ports.Status8 != nil {
							return &ResponseDevicesGetDeviceLldpCdpPorts8{
								Cdp: func() *ResponseDevicesGetDeviceLldpCdpPorts8Cdp {
									if response.Ports.Status8.Cdp != nil {
										return &ResponseDevicesGetDeviceLldpCdpPorts8Cdp{
											Address: func() types.String {
												if response.Ports.Status8.Cdp.Address != "" {
													return types.StringValue(response.Ports.Status8.Cdp.Address)
												}
												return types.String{}
											}(),
											DeviceID: func() types.String {
												if response.Ports.Status8.Cdp.DeviceID != "" {
													return types.StringValue(response.Ports.Status8.Cdp.DeviceID)
												}
												return types.String{}
											}(),
											PortID: func() types.String {
												if response.Ports.Status8.Cdp.PortID != "" {
													return types.StringValue(response.Ports.Status8.Cdp.PortID)
												}
												return types.String{}
											}(),
											SourcePort: func() types.String {
												if response.Ports.Status8.Cdp.SourcePort != "" {
													return types.StringValue(response.Ports.Status8.Cdp.SourcePort)
												}
												return types.String{}
											}(),
										}
									}
									return &ResponseDevicesGetDeviceLldpCdpPorts8Cdp{}
								}(),
							}
						}
						return &ResponseDevicesGetDeviceLldpCdpPorts8{}
					}(),
				}
			}
			return &ResponseDevicesGetDeviceLldpCdpPorts{}
		}(),
		SourceMac: func() types.String {
			if response.SourceMac != "" {
				return types.StringValue(response.SourceMac)
			}
			return types.String{}
		}(),
	}
	state.Item = &itemState
	return state
}
