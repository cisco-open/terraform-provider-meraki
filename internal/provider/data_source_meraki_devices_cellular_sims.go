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
	_ datasource.DataSource              = &DevicesCellularSimsDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesCellularSimsDataSource{}
)

func NewDevicesCellularSimsDataSource() datasource.DataSource {
	return &DevicesCellularSimsDataSource{}
}

type DevicesCellularSimsDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesCellularSimsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesCellularSimsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_cellular_sims"
}

func (d *DevicesCellularSimsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"sim_failover": schema.SingleNestedAttribute{
						MarkdownDescription: `SIM Failover settings.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Failover to secondary SIM`,
								Computed:            true,
							},
							"timeout": schema.Int64Attribute{
								MarkdownDescription: `Failover timeout in seconds`,
								Computed:            true,
							},
						},
					},
					"sim_ordering": schema.ListAttribute{
						MarkdownDescription: `Specifies the ordering of all SIMs for an MG: primary, secondary, and not-in-use (when applicable). It's required for devices with 3 or more SIMs and can be used in place of 'isPrimary' for dual-SIM devices. To indicate eSIM, use 'sim3'. Sim failover will occur only between primary and secondary sim slots.`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"sims": schema.SetNestedAttribute{
						MarkdownDescription: `List of SIMs. If a SIM was previously configured and not specified in this request, it will remain unchanged.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"apns": schema.SetNestedAttribute{
									MarkdownDescription: `APN configurations. If empty, the default APN will be used.`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"allowed_ip_types": schema.ListAttribute{
												MarkdownDescription: `IP versions to support (permitted values include 'ipv4', 'ipv6').`,
												Computed:            true,
												ElementType:         types.StringType,
											},
											"authentication": schema.SingleNestedAttribute{
												MarkdownDescription: `APN authentication configurations.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"password": schema.StringAttribute{
														MarkdownDescription: `APN password, if type is set (if APN password is not supplied, the password is left unchanged).`,
														Sensitive:           true,
														Computed:            true,
													},
													"type": schema.StringAttribute{
														MarkdownDescription: `APN auth type.`,
														Computed:            true,
													},
													"username": schema.StringAttribute{
														MarkdownDescription: `APN username, if type is set.`,
														Computed:            true,
													},
												},
											},
											"name": schema.StringAttribute{
												MarkdownDescription: `APN name.`,
												Computed:            true,
											},
										},
									},
								},
								"is_primary": schema.BoolAttribute{
									MarkdownDescription: `If true, this SIM is activated on platform bootup. It must be true on single-SIM devices and is a required field for dual-SIM MGs unless it is being configured using 'simOrdering'.`,
									Computed:            true,
								},
								"slot": schema.StringAttribute{
									MarkdownDescription: `SIM slot being configured. Must be 'sim1' on single-sim devices. Use 'sim3' for eSIM.`,
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

func (d *DevicesCellularSimsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesCellularSims DevicesCellularSims
	diags := req.Config.Get(ctx, &devicesCellularSims)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceCellularSims")
		vvSerial := devicesCellularSims.Serial.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Devices.GetDeviceCellularSims(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCellularSims",
				err.Error(),
			)
			return
		}

		devicesCellularSims = ResponseDevicesGetDeviceCellularSimsItemToBody(devicesCellularSims, response1)
		diags = resp.State.Set(ctx, &devicesCellularSims)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesCellularSims struct {
	Serial types.String                          `tfsdk:"serial"`
	Item   *ResponseDevicesGetDeviceCellularSims `tfsdk:"item"`
}

type ResponseDevicesGetDeviceCellularSims struct {
	SimFailover *ResponseDevicesGetDeviceCellularSimsSimFailover `tfsdk:"sim_failover"`
	SimOrdering types.List                                       `tfsdk:"sim_ordering"`
	Sims        *[]ResponseDevicesGetDeviceCellularSimsSims      `tfsdk:"sims"`
}

type ResponseDevicesGetDeviceCellularSimsSimFailover struct {
	Enabled types.Bool  `tfsdk:"enabled"`
	Timeout types.Int64 `tfsdk:"timeout"`
}

type ResponseDevicesGetDeviceCellularSimsSims struct {
	Apns      *[]ResponseDevicesGetDeviceCellularSimsSimsApns `tfsdk:"apns"`
	IsPrimary types.Bool                                      `tfsdk:"is_primary"`
	Slot      types.String                                    `tfsdk:"slot"`
}

type ResponseDevicesGetDeviceCellularSimsSimsApns struct {
	AllowedIPTypes types.List                                                  `tfsdk:"allowed_ip_types"`
	Authentication *ResponseDevicesGetDeviceCellularSimsSimsApnsAuthentication `tfsdk:"authentication"`
	Name           types.String                                                `tfsdk:"name"`
}

type ResponseDevicesGetDeviceCellularSimsSimsApnsAuthentication struct {
	Password types.String `tfsdk:"password"`
	Type     types.String `tfsdk:"type"`
	Username types.String `tfsdk:"username"`
}

// ToBody
func ResponseDevicesGetDeviceCellularSimsItemToBody(state DevicesCellularSims, response *merakigosdk.ResponseDevicesGetDeviceCellularSims) DevicesCellularSims {
	itemState := ResponseDevicesGetDeviceCellularSims{
		SimFailover: func() *ResponseDevicesGetDeviceCellularSimsSimFailover {
			if response.SimFailover != nil {
				return &ResponseDevicesGetDeviceCellularSimsSimFailover{
					Enabled: func() types.Bool {
						if response.SimFailover.Enabled != nil {
							return types.BoolValue(*response.SimFailover.Enabled)
						}
						return types.Bool{}
					}(),
					Timeout: func() types.Int64 {
						if response.SimFailover.Timeout != nil {
							return types.Int64Value(int64(*response.SimFailover.Timeout))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
		SimOrdering: StringSliceToList(response.SimOrdering),
		Sims: func() *[]ResponseDevicesGetDeviceCellularSimsSims {
			if response.Sims != nil {
				result := make([]ResponseDevicesGetDeviceCellularSimsSims, len(*response.Sims))
				for i, sims := range *response.Sims {
					result[i] = ResponseDevicesGetDeviceCellularSimsSims{
						Apns: func() *[]ResponseDevicesGetDeviceCellularSimsSimsApns {
							if sims.Apns != nil {
								result := make([]ResponseDevicesGetDeviceCellularSimsSimsApns, len(*sims.Apns))
								for i, apns := range *sims.Apns {
									result[i] = ResponseDevicesGetDeviceCellularSimsSimsApns{
										AllowedIPTypes: StringSliceToList(apns.AllowedIPTypes),
										Authentication: func() *ResponseDevicesGetDeviceCellularSimsSimsApnsAuthentication {
											if apns.Authentication != nil {
												return &ResponseDevicesGetDeviceCellularSimsSimsApnsAuthentication{
													Password: func() types.String {
														if apns.Authentication.Password != "" {
															return types.StringValue(apns.Authentication.Password)
														}
														return types.String{}
													}(),
													Type: func() types.String {
														if apns.Authentication.Type != "" {
															return types.StringValue(apns.Authentication.Type)
														}
														return types.String{}
													}(),
													Username: func() types.String {
														if apns.Authentication.Username != "" {
															return types.StringValue(apns.Authentication.Username)
														}
														return types.String{}
													}(),
												}
											}
											return nil
										}(),
										Name: func() types.String {
											if apns.Name != "" {
												return types.StringValue(apns.Name)
											}
											return types.String{}
										}(),
									}
								}
								return &result
							}
							return nil
						}(),
						IsPrimary: func() types.Bool {
							if sims.IsPrimary != nil {
								return types.BoolValue(*sims.IsPrimary)
							}
							return types.Bool{}
						}(),
						Slot: func() types.String {
							if sims.Slot != "" {
								return types.StringValue(sims.Slot)
							}
							return types.String{}
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
