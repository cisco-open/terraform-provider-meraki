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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksSwitchStacksRoutingInterfacesDhcpDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchStacksRoutingInterfacesDhcpDataSource{}
)

func NewNetworksSwitchStacksRoutingInterfacesDhcpDataSource() datasource.DataSource {
	return &NetworksSwitchStacksRoutingInterfacesDhcpDataSource{}
}

type NetworksSwitchStacksRoutingInterfacesDhcpDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchStacksRoutingInterfacesDhcpDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchStacksRoutingInterfacesDhcpDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_stacks_routing_interfaces_dhcp"
}

func (d *NetworksSwitchStacksRoutingInterfacesDhcpDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"interface_id": schema.StringAttribute{
				MarkdownDescription: `interfaceId path parameter. Interface ID`,
				Required:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"switch_stack_id": schema.StringAttribute{
				MarkdownDescription: `switchStackId path parameter. Switch stack ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"boot_file_name": schema.StringAttribute{
						Computed: true,
					},
					"boot_next_server": schema.StringAttribute{
						Computed: true,
					},
					"boot_options_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"dhcp_lease_time": schema.StringAttribute{
						Computed: true,
					},
					"dhcp_mode": schema.StringAttribute{
						Computed: true,
					},
					"dhcp_options": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"code": schema.StringAttribute{
									Computed: true,
								},
								"type": schema.StringAttribute{
									Computed: true,
								},
								"value": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
					"dns_custom_nameservers": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"dns_nameservers_option": schema.StringAttribute{
						Computed: true,
					},
					"fixed_ip_assignments": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"ip": schema.StringAttribute{
									Computed: true,
								},
								"mac": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
					"reserved_ip_ranges": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"comment": schema.StringAttribute{
									Computed: true,
								},
								"end": schema.StringAttribute{
									Computed: true,
								},
								"start": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchStacksRoutingInterfacesDhcpDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchStacksRoutingInterfacesDhcp NetworksSwitchStacksRoutingInterfacesDhcp
	diags := req.Config.Get(ctx, &networksSwitchStacksRoutingInterfacesDhcp)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchStackRoutingInterfaceDhcp")
		vvNetworkID := networksSwitchStacksRoutingInterfacesDhcp.NetworkID.ValueString()
		vvSwitchStackID := networksSwitchStacksRoutingInterfacesDhcp.SwitchStackID.ValueString()
		vvInterfaceID := networksSwitchStacksRoutingInterfacesDhcp.InterfaceID.ValueString()

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchStackRoutingInterfaceDhcp(vvNetworkID, vvSwitchStackID, vvInterfaceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStackRoutingInterfaceDhcp",
				err.Error(),
			)
			return
		}

		networksSwitchStacksRoutingInterfacesDhcp = ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpItemToBody(networksSwitchStacksRoutingInterfacesDhcp, response1)
		diags = resp.State.Set(ctx, &networksSwitchStacksRoutingInterfacesDhcp)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchStacksRoutingInterfacesDhcp struct {
	NetworkID     types.String                                             `tfsdk:"network_id"`
	SwitchStackID types.String                                             `tfsdk:"switch_stack_id"`
	InterfaceID   types.String                                             `tfsdk:"interface_id"`
	Item          *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcp `tfsdk:"item"`
}

type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcp struct {
	BootFileName         types.String                                                                 `tfsdk:"boot_file_name"`
	BootNextServer       types.String                                                                 `tfsdk:"boot_next_server"`
	BootOptionsEnabled   types.Bool                                                                   `tfsdk:"boot_options_enabled"`
	DhcpLeaseTime        types.String                                                                 `tfsdk:"dhcp_lease_time"`
	DhcpMode             types.String                                                                 `tfsdk:"dhcp_mode"`
	DhcpOptions          *[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions        `tfsdk:"dhcp_options"`
	DNSCustomNameservers types.List                                                                   `tfsdk:"dns_custom_nameservers"`
	DNSNameserversOption types.String                                                                 `tfsdk:"dns_nameservers_option"`
	FixedIPAssignments   *[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpFixedIpAssignments `tfsdk:"fixed_ip_assignments"`
	ReservedIPRanges     *[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpReservedIpRanges   `tfsdk:"reserved_ip_ranges"`
}

type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions struct {
	Code  types.String `tfsdk:"code"`
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpFixedIpAssignments struct {
	IP   types.String `tfsdk:"ip"`
	Mac  types.String `tfsdk:"mac"`
	Name types.String `tfsdk:"name"`
}

type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpReservedIpRanges struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpItemToBody(state NetworksSwitchStacksRoutingInterfacesDhcp, response *merakigosdk.ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcp) NetworksSwitchStacksRoutingInterfacesDhcp {
	itemState := ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcp{
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
		DhcpOptions: func() *[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions {
			if response.DhcpOptions != nil {
				result := make([]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions, len(*response.DhcpOptions))
				for i, dhcpOptions := range *response.DhcpOptions {
					result[i] = ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions{
						Code:  types.StringValue(dhcpOptions.Code),
						Type:  types.StringValue(dhcpOptions.Type),
						Value: types.StringValue(dhcpOptions.Value),
					}
				}
				return &result
			}
			return &[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions{}
		}(),
		DNSCustomNameservers: StringSliceToList(response.DNSCustomNameservers),
		DNSNameserversOption: types.StringValue(response.DNSNameserversOption),
		FixedIPAssignments: func() *[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpFixedIpAssignments {
			if response.FixedIPAssignments != nil {
				result := make([]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpFixedIpAssignments, len(*response.FixedIPAssignments))
				for i, fixedIPAssignments := range *response.FixedIPAssignments {
					result[i] = ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpFixedIpAssignments{
						IP:   types.StringValue(fixedIPAssignments.IP),
						Mac:  types.StringValue(fixedIPAssignments.Mac),
						Name: types.StringValue(fixedIPAssignments.Name),
					}
				}
				return &result
			}
			return &[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpFixedIpAssignments{}
		}(),
		ReservedIPRanges: func() *[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpReservedIpRanges {
			if response.ReservedIPRanges != nil {
				result := make([]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpReservedIpRanges, len(*response.ReservedIPRanges))
				for i, reservedIPRanges := range *response.ReservedIPRanges {
					result[i] = ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpReservedIpRanges{
						Comment: types.StringValue(reservedIPRanges.Comment),
						End:     types.StringValue(reservedIPRanges.End),
						Start:   types.StringValue(reservedIPRanges.Start),
					}
				}
				return &result
			}
			return &[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpReservedIpRanges{}
		}(),
	}
	state.Item = &itemState
	return state
}
