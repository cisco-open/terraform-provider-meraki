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
	_ datasource.DataSource              = &DevicesSwitchRoutingInterfacesDhcpDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesSwitchRoutingInterfacesDhcpDataSource{}
)

func NewDevicesSwitchRoutingInterfacesDhcpDataSource() datasource.DataSource {
	return &DevicesSwitchRoutingInterfacesDhcpDataSource{}
}

type DevicesSwitchRoutingInterfacesDhcpDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesSwitchRoutingInterfacesDhcpDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesSwitchRoutingInterfacesDhcpDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_switch_routing_interfaces_dhcp"
}

func (d *DevicesSwitchRoutingInterfacesDhcpDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"interface_id": schema.StringAttribute{
				MarkdownDescription: `interfaceId path parameter. Interface ID`,
				Required:            true,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"boot_file_name": schema.StringAttribute{
						MarkdownDescription: `The PXE boot server file name for the DHCP server running on the switch stack interface`,
						Computed:            true,
					},
					"boot_next_server": schema.StringAttribute{
						MarkdownDescription: `The PXE boot server IP for the DHCP server running on the switch stack interface`,
						Computed:            true,
					},
					"boot_options_enabled": schema.BoolAttribute{
						MarkdownDescription: `Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch stack interface`,
						Computed:            true,
					},
					"dhcp_lease_time": schema.StringAttribute{
						MarkdownDescription: `The DHCP lease time config for the dhcp server running on the switch stack interface ('30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week')`,
						Computed:            true,
					},
					"dhcp_mode": schema.StringAttribute{
						MarkdownDescription: `The DHCP mode options for the switch stack interface ('dhcpDisabled', 'dhcpRelay' or 'dhcpServer')`,
						Computed:            true,
					},
					"dhcp_options": schema.SetNestedAttribute{
						MarkdownDescription: `Array of DHCP options consisting of code, type and value for the DHCP server running on the switch stack interface`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"code": schema.StringAttribute{
									MarkdownDescription: `The code for DHCP option which should be from 2 to 254`,
									Computed:            true,
								},
								"type": schema.StringAttribute{
									MarkdownDescription: `The type of the DHCP option which should be one of ('text', 'ip', 'integer' or 'hex')`,
									Computed:            true,
								},
								"value": schema.StringAttribute{
									MarkdownDescription: `The value of the DHCP option`,
									Computed:            true,
								},
							},
						},
					},
					"dhcp_relay_server_ips": schema.ListAttribute{
						MarkdownDescription: `The DHCP relay server IPs to which DHCP packets would get relayed for the switch stack interface`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"dns_custom_nameservers": schema.ListAttribute{
						MarkdownDescription: `The DHCP name server IPs when DHCP name server option is 'custom'`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"dns_nameservers_option": schema.StringAttribute{
						MarkdownDescription: `The DHCP name server option for the dhcp server running on the switch stack interface ('googlePublicDns', 'openDns' or 'custom')`,
						Computed:            true,
					},
					"fixed_ip_assignments": schema.SetNestedAttribute{
						MarkdownDescription: `Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"ip": schema.StringAttribute{
									MarkdownDescription: `The IP address of the client which has fixed IP address assigned to it`,
									Computed:            true,
								},
								"mac": schema.StringAttribute{
									MarkdownDescription: `The MAC address of the client which has fixed IP address`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `The name of the client which has fixed IP address`,
									Computed:            true,
								},
							},
						},
					},
					"reserved_ip_ranges": schema.SetNestedAttribute{
						MarkdownDescription: `Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"comment": schema.StringAttribute{
									MarkdownDescription: `The comment for the reserved IP range`,
									Computed:            true,
								},
								"end": schema.StringAttribute{
									MarkdownDescription: `The ending IP address of the reserved IP range`,
									Computed:            true,
								},
								"start": schema.StringAttribute{
									MarkdownDescription: `The starting IP address of the reserved IP range`,
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

func (d *DevicesSwitchRoutingInterfacesDhcpDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesSwitchRoutingInterfacesDhcp DevicesSwitchRoutingInterfacesDhcp
	diags := req.Config.Get(ctx, &devicesSwitchRoutingInterfacesDhcp)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceSwitchRoutingInterfaceDhcp")
		vvSerial := devicesSwitchRoutingInterfacesDhcp.Serial.ValueString()
		vvInterfaceID := devicesSwitchRoutingInterfacesDhcp.InterfaceID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetDeviceSwitchRoutingInterfaceDhcp(vvSerial, vvInterfaceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceSwitchRoutingInterfaceDhcp",
				err.Error(),
			)
			return
		}

		devicesSwitchRoutingInterfacesDhcp = ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpItemToBody(devicesSwitchRoutingInterfacesDhcp, response1)
		diags = resp.State.Set(ctx, &devicesSwitchRoutingInterfacesDhcp)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesSwitchRoutingInterfacesDhcp struct {
	Serial      types.String                                       `tfsdk:"serial"`
	InterfaceID types.String                                       `tfsdk:"interface_id"`
	Item        *ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcp `tfsdk:"item"`
}

type ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcp struct {
	BootFileName         types.String                                                           `tfsdk:"boot_file_name"`
	BootNextServer       types.String                                                           `tfsdk:"boot_next_server"`
	BootOptionsEnabled   types.Bool                                                             `tfsdk:"boot_options_enabled"`
	DhcpLeaseTime        types.String                                                           `tfsdk:"dhcp_lease_time"`
	DhcpMode             types.String                                                           `tfsdk:"dhcp_mode"`
	DhcpOptions          *[]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpDhcpOptions        `tfsdk:"dhcp_options"`
	DhcpRelayServerIPs   types.List                                                             `tfsdk:"dhcp_relay_server_ips"`
	DNSCustomNameservers types.List                                                             `tfsdk:"dns_custom_nameservers"`
	DNSNameserversOption types.String                                                           `tfsdk:"dns_nameservers_option"`
	FixedIPAssignments   *[]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpFixedIpAssignments `tfsdk:"fixed_ip_assignments"`
	ReservedIPRanges     *[]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpReservedIpRanges   `tfsdk:"reserved_ip_ranges"`
}

type ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpDhcpOptions struct {
	Code  types.String `tfsdk:"code"`
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpFixedIpAssignments struct {
	IP   types.String `tfsdk:"ip"`
	Mac  types.String `tfsdk:"mac"`
	Name types.String `tfsdk:"name"`
}

type ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpReservedIpRanges struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

// ToBody
func ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpItemToBody(state DevicesSwitchRoutingInterfacesDhcp, response *merakigosdk.ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcp) DevicesSwitchRoutingInterfacesDhcp {
	itemState := ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcp{
		BootFileName:   types.StringValue(response.BootFileName),
		BootNextServer: types.StringValue(response.BootNextServer),
		BootOptionsEnabled: func() types.Bool {
			if response.BootOptionsEnabled != nil {
				return types.BoolValue(*response.BootOptionsEnabled)
			}
			return types.Bool{}
		}(),
		DhcpLeaseTime: types.StringValue(response.DhcpLeaseTime),
		DhcpMode:      types.StringValue(response.DhcpMode),
		DhcpOptions: func() *[]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpDhcpOptions {
			if response.DhcpOptions != nil {
				result := make([]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpDhcpOptions, len(*response.DhcpOptions))
				for i, dhcpOptions := range *response.DhcpOptions {
					result[i] = ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpDhcpOptions{
						Code:  types.StringValue(dhcpOptions.Code),
						Type:  types.StringValue(dhcpOptions.Type),
						Value: types.StringValue(dhcpOptions.Value),
					}
				}
				return &result
			}
			return nil
		}(),
		DhcpRelayServerIPs:   StringSliceToList(response.DhcpRelayServerIPs),
		DNSCustomNameservers: StringSliceToList(response.DNSCustomNameservers),
		DNSNameserversOption: types.StringValue(response.DNSNameserversOption),
		FixedIPAssignments: func() *[]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpFixedIpAssignments {
			if response.FixedIPAssignments != nil {
				result := make([]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpFixedIpAssignments, len(*response.FixedIPAssignments))
				for i, fixedIPAssignments := range *response.FixedIPAssignments {
					result[i] = ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpFixedIpAssignments{
						IP:   types.StringValue(fixedIPAssignments.IP),
						Mac:  types.StringValue(fixedIPAssignments.Mac),
						Name: types.StringValue(fixedIPAssignments.Name),
					}
				}
				return &result
			}
			return nil
		}(),
		ReservedIPRanges: func() *[]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpReservedIpRanges {
			if response.ReservedIPRanges != nil {
				result := make([]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpReservedIpRanges, len(*response.ReservedIPRanges))
				for i, reservedIPRanges := range *response.ReservedIPRanges {
					result[i] = ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpReservedIpRanges{
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
