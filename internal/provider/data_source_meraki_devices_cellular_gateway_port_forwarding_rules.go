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
	_ datasource.DataSource              = &DevicesCellularGatewayPortForwardingRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesCellularGatewayPortForwardingRulesDataSource{}
)

func NewDevicesCellularGatewayPortForwardingRulesDataSource() datasource.DataSource {
	return &DevicesCellularGatewayPortForwardingRulesDataSource{}
}

type DevicesCellularGatewayPortForwardingRulesDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesCellularGatewayPortForwardingRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesCellularGatewayPortForwardingRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_cellular_gateway_port_forwarding_rules"
}

func (d *DevicesCellularGatewayPortForwardingRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"rules": schema.SetNestedAttribute{
						MarkdownDescription: `An array of port forwarding params`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"access": schema.StringAttribute{
									MarkdownDescription: `**any** or **restricted**. Specify the right to make inbound connections on the specified ports or port ranges. If **restricted**, a list of allowed IPs is mandatory.`,
									Computed:            true,
								},
								"allowed_ips": schema.ListAttribute{
									MarkdownDescription: `An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges.`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"lan_ip": schema.StringAttribute{
									MarkdownDescription: `The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN`,
									Computed:            true,
								},
								"local_port": schema.StringAttribute{
									MarkdownDescription: `A port or port ranges that will receive the forwarded traffic from the WAN`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `A descriptive name for the rule`,
									Computed:            true,
								},
								"protocol": schema.StringAttribute{
									MarkdownDescription: `TCP or UDP`,
									Computed:            true,
								},
								"public_port": schema.StringAttribute{
									MarkdownDescription: `A port or port ranges that will be forwarded to the host on the LAN`,
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

func (d *DevicesCellularGatewayPortForwardingRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesCellularGatewayPortForwardingRules DevicesCellularGatewayPortForwardingRules
	diags := req.Config.Get(ctx, &devicesCellularGatewayPortForwardingRules)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceCellularGatewayPortForwardingRules")
		vvSerial := devicesCellularGatewayPortForwardingRules.Serial.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.CellularGateway.GetDeviceCellularGatewayPortForwardingRules(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCellularGatewayPortForwardingRules",
				err.Error(),
			)
			return
		}

		devicesCellularGatewayPortForwardingRules = ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesItemToBody(devicesCellularGatewayPortForwardingRules, response1)
		diags = resp.State.Set(ctx, &devicesCellularGatewayPortForwardingRules)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesCellularGatewayPortForwardingRules struct {
	Serial types.String                                                        `tfsdk:"serial"`
	Item   *ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRules `tfsdk:"item"`
}

type ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRules struct {
	Rules *[]ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRules `tfsdk:"rules"`
}

type ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRules struct {
	Access     types.String `tfsdk:"access"`
	AllowedIPs types.List   `tfsdk:"allowed_ips"`
	LanIP      types.String `tfsdk:"lan_ip"`
	LocalPort  types.String `tfsdk:"local_port"`
	Name       types.String `tfsdk:"name"`
	Protocol   types.String `tfsdk:"protocol"`
	PublicPort types.String `tfsdk:"public_port"`
}

// ToBody
func ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesItemToBody(state DevicesCellularGatewayPortForwardingRules, response *merakigosdk.ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRules) DevicesCellularGatewayPortForwardingRules {
	itemState := ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRules{
		Rules: func() *[]ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRules {
			if response.Rules != nil {
				result := make([]ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRules{
						Access: func() types.String {
							if rules.Access != "" {
								return types.StringValue(rules.Access)
							}
							return types.String{}
						}(),
						AllowedIPs: StringSliceToList(rules.AllowedIPs),
						LanIP: func() types.String {
							if rules.LanIP != "" {
								return types.StringValue(rules.LanIP)
							}
							return types.String{}
						}(),
						LocalPort: func() types.String {
							if rules.LocalPort != "" {
								return types.StringValue(rules.LocalPort)
							}
							return types.String{}
						}(),
						Name: func() types.String {
							if rules.Name != "" {
								return types.StringValue(rules.Name)
							}
							return types.String{}
						}(),
						Protocol: func() types.String {
							if rules.Protocol != "" {
								return types.StringValue(rules.Protocol)
							}
							return types.String{}
						}(),
						PublicPort: func() types.String {
							if rules.PublicPort != "" {
								return types.StringValue(rules.PublicPort)
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
