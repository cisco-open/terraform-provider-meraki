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
	_ datasource.DataSource              = &DevicesCellularGatewayLanDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesCellularGatewayLanDataSource{}
)

func NewDevicesCellularGatewayLanDataSource() datasource.DataSource {
	return &DevicesCellularGatewayLanDataSource{}
}

type DevicesCellularGatewayLanDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesCellularGatewayLanDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesCellularGatewayLanDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_cellular_gateway_lan"
}

func (d *DevicesCellularGatewayLanDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"device_lan_ip": schema.StringAttribute{
						MarkdownDescription: `Lan IP of the MG`,
						Computed:            true,
					},
					"device_name": schema.StringAttribute{
						MarkdownDescription: `Name of the MG.`,
						Computed:            true,
					},
					"device_subnet": schema.StringAttribute{
						MarkdownDescription: `Subnet configuration of the MG.`,
						Computed:            true,
					},
					"fixed_ip_assignments": schema.SetNestedAttribute{
						MarkdownDescription: `list of all fixed IP assignments for a single MG`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"ip": schema.StringAttribute{
									MarkdownDescription: `The IP address you want to assign to a specific server or device`,
									Computed:            true,
								},
								"mac": schema.StringAttribute{
									MarkdownDescription: `The MAC address of the server or device that hosts the internal resource that you wish to receive the specified IP address`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `A descriptive name of the assignment`,
									Computed:            true,
								},
							},
						},
					},
					"reserved_ip_ranges": schema.SetNestedAttribute{
						MarkdownDescription: `list of all reserved IP ranges for a single MG`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"comment": schema.StringAttribute{
									MarkdownDescription: `Comment explaining the reserved IP range`,
									Computed:            true,
								},
								"end": schema.StringAttribute{
									MarkdownDescription: `Ending IP included in the reserved range of IPs`,
									Computed:            true,
								},
								"start": schema.StringAttribute{
									MarkdownDescription: `Starting IP included in the reserved range of IPs`,
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

func (d *DevicesCellularGatewayLanDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesCellularGatewayLan DevicesCellularGatewayLan
	diags := req.Config.Get(ctx, &devicesCellularGatewayLan)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceCellularGatewayLan")
		vvSerial := devicesCellularGatewayLan.Serial.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.CellularGateway.GetDeviceCellularGatewayLan(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCellularGatewayLan",
				err.Error(),
			)
			return
		}

		devicesCellularGatewayLan = ResponseCellularGatewayGetDeviceCellularGatewayLanItemToBody(devicesCellularGatewayLan, response1)
		diags = resp.State.Set(ctx, &devicesCellularGatewayLan)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesCellularGatewayLan struct {
	Serial types.String                                        `tfsdk:"serial"`
	Item   *ResponseCellularGatewayGetDeviceCellularGatewayLan `tfsdk:"item"`
}

type ResponseCellularGatewayGetDeviceCellularGatewayLan struct {
	DeviceLanIP        types.String                                                            `tfsdk:"device_lan_ip"`
	DeviceName         types.String                                                            `tfsdk:"device_name"`
	DeviceSubnet       types.String                                                            `tfsdk:"device_subnet"`
	FixedIPAssignments *[]ResponseCellularGatewayGetDeviceCellularGatewayLanFixedIpAssignments `tfsdk:"fixed_ip_assignments"`
	ReservedIPRanges   *[]ResponseCellularGatewayGetDeviceCellularGatewayLanReservedIpRanges   `tfsdk:"reserved_ip_ranges"`
}

type ResponseCellularGatewayGetDeviceCellularGatewayLanFixedIpAssignments struct {
	IP   types.String `tfsdk:"ip"`
	Mac  types.String `tfsdk:"mac"`
	Name types.String `tfsdk:"name"`
}

type ResponseCellularGatewayGetDeviceCellularGatewayLanReservedIpRanges struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

// ToBody
func ResponseCellularGatewayGetDeviceCellularGatewayLanItemToBody(state DevicesCellularGatewayLan, response *merakigosdk.ResponseCellularGatewayGetDeviceCellularGatewayLan) DevicesCellularGatewayLan {
	itemState := ResponseCellularGatewayGetDeviceCellularGatewayLan{
		DeviceLanIP:  types.StringValue(response.DeviceLanIP),
		DeviceName:   types.StringValue(response.DeviceName),
		DeviceSubnet: types.StringValue(response.DeviceSubnet),
		FixedIPAssignments: func() *[]ResponseCellularGatewayGetDeviceCellularGatewayLanFixedIpAssignments {
			if response.FixedIPAssignments != nil {
				result := make([]ResponseCellularGatewayGetDeviceCellularGatewayLanFixedIpAssignments, len(*response.FixedIPAssignments))
				for i, fixedIPAssignments := range *response.FixedIPAssignments {
					result[i] = ResponseCellularGatewayGetDeviceCellularGatewayLanFixedIpAssignments{
						IP:   types.StringValue(fixedIPAssignments.IP),
						Mac:  types.StringValue(fixedIPAssignments.Mac),
						Name: types.StringValue(fixedIPAssignments.Name),
					}
				}
				return &result
			}
			return nil
		}(),
		ReservedIPRanges: func() *[]ResponseCellularGatewayGetDeviceCellularGatewayLanReservedIpRanges {
			if response.ReservedIPRanges != nil {
				result := make([]ResponseCellularGatewayGetDeviceCellularGatewayLanReservedIpRanges, len(*response.ReservedIPRanges))
				for i, reservedIPRanges := range *response.ReservedIPRanges {
					result[i] = ResponseCellularGatewayGetDeviceCellularGatewayLanReservedIpRanges{
						Comment: types.StringValue(reservedIPRanges.Comment),
						End:     types.StringValue(reservedIPRanges.End),
						Start:   types.StringValue(reservedIPRanges.Start),
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
