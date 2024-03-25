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
	_ datasource.DataSource              = &NetworksApplianceVLANsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceVLANsDataSource{}
)

func NewNetworksApplianceVLANsDataSource() datasource.DataSource {
	return &NetworksApplianceVLANsDataSource{}
}

type NetworksApplianceVLANsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceVLANsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceVLANsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_vlans"
}

func (d *NetworksApplianceVLANsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"vlan_id": schema.StringAttribute{
				MarkdownDescription: `vlanId path parameter. Vlan ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"appliance_ip": schema.StringAttribute{
						MarkdownDescription: `The local IP of the appliance on the VLAN`,
						Computed:            true,
					},
					"cidr": schema.StringAttribute{
						MarkdownDescription: `CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN.`,
						Computed:            true,
					},
					"dhcp_boot_filename": schema.StringAttribute{
						MarkdownDescription: `DHCP boot option for boot filename`,
						Computed:            true,
					},
					"dhcp_boot_next_server": schema.StringAttribute{
						MarkdownDescription: `DHCP boot option to direct boot clients to the server to load the boot file from`,
						Computed:            true,
					},
					"dhcp_boot_options_enabled": schema.BoolAttribute{
						MarkdownDescription: `Use DHCP boot options specified in other properties`,
						Computed:            true,
					},
					"dhcp_handling": schema.StringAttribute{
						MarkdownDescription: `The appliance's handling of DHCP requests on this VLAN. One of: 'Run a DHCP server', 'Relay DHCP to another server' or 'Do not respond to DHCP requests'`,
						Computed:            true,
					},
					"dhcp_lease_time": schema.StringAttribute{
						MarkdownDescription: `The term of DHCP leases if the appliance is running a DHCP server on this VLAN. One of: '30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week'`,
						Computed:            true,
					},
					"dhcp_options": schema.SetNestedAttribute{
						MarkdownDescription: `The list of DHCP options that will be included in DHCP responses. Each object in the list should have "code", "type", and "value" properties.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"code": schema.StringAttribute{
									MarkdownDescription: `The code for the DHCP option. This should be an integer between 2 and 254.`,
									Computed:            true,
								},
								"type": schema.StringAttribute{
									MarkdownDescription: `The type for the DHCP option. One of: 'text', 'ip', 'hex' or 'integer'`,
									Computed:            true,
								},
								"value": schema.StringAttribute{
									MarkdownDescription: `The value for the DHCP option`,
									Computed:            true,
								},
							},
						},
					},
					"dhcp_relay_server_ips": schema.ListAttribute{
						MarkdownDescription: `The IPs of the DHCP servers that DHCP requests should be relayed to`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"dns_nameservers": schema.StringAttribute{
						MarkdownDescription: `The DNS nameservers used for DHCP responses, either "upstream_dns", "google_dns", "opendns", or a newline seperated string of IP addresses or domain names`,
						Computed:            true,
					},
					"fixed_ip_assignments": schema.StringAttribute{
						MarkdownDescription: `The DHCP fixed IP assignments on the VLAN. This should be an object that contains mappings from MAC addresses to objects that themselves each contain "ip" and "name" string fields. See the sample request/response for more details.`,
						Computed:            true,
					},
					"group_policy_id": schema.StringAttribute{
						MarkdownDescription: `The id of the desired group policy to apply to the VLAN`,
						Computed:            true,
					},
					"id": schema.Int64Attribute{
						MarkdownDescription: `The VLAN ID of the VLAN`,
						Computed:            true,
					},
					"interface_id": schema.StringAttribute{
						MarkdownDescription: `The interface ID of the VLAN`,
						Computed:            true,
					},
					"ipv6": schema.SingleNestedAttribute{
						MarkdownDescription: `IPv6 configuration on the VLAN`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Enable IPv6 on VLAN`,
								Computed:            true,
							},
							"prefix_assignments": schema.SetNestedAttribute{
								MarkdownDescription: `Prefix assignments on the VLAN`,
								Computed:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{

										"autonomous": schema.BoolAttribute{
											MarkdownDescription: `Auto assign a /64 prefix from the origin to the VLAN`,
											Computed:            true,
										},
										"origin": schema.SingleNestedAttribute{
											MarkdownDescription: `The origin of the prefix`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"interfaces": schema.ListAttribute{
													MarkdownDescription: `Interfaces associated with the prefix`,
													Computed:            true,
													ElementType:         types.StringType,
												},
												"type": schema.StringAttribute{
													MarkdownDescription: `Type of the origin`,
													Computed:            true,
												},
											},
										},
										"static_appliance_ip6": schema.StringAttribute{
											MarkdownDescription: `Manual configuration of the IPv6 Appliance IP`,
											Computed:            true,
										},
										"static_prefix": schema.StringAttribute{
											MarkdownDescription: `Manual configuration of a /64 prefix on the VLAN`,
											Computed:            true,
										},
									},
								},
							},
						},
					},
					"mandatory_dhcp": schema.SingleNestedAttribute{
						MarkdownDescription: `Mandatory DHCP will enforce that clients connecting to this VLAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Enable Mandatory DHCP on VLAN.`,
								Computed:            true,
							},
						},
					},
					"mask": schema.Int64Attribute{
						MarkdownDescription: `Mask used for the subnet of all bound to the template networks. Applicable only for template network.`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name of the VLAN`,
						Computed:            true,
					},
					"reserved_ip_ranges": schema.SetNestedAttribute{
						MarkdownDescription: `The DHCP reserved IP ranges on the VLAN`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"comment": schema.StringAttribute{
									MarkdownDescription: `A text comment for the reserved range`,
									Computed:            true,
								},
								"end": schema.StringAttribute{
									MarkdownDescription: `The last IP in the reserved range`,
									Computed:            true,
								},
								"start": schema.StringAttribute{
									MarkdownDescription: `The first IP in the reserved range`,
									Computed:            true,
								},
							},
						},
					},
					"subnet": schema.StringAttribute{
						MarkdownDescription: `The subnet of the VLAN`,
						Computed:            true,
					},
					"template_vlan_type": schema.StringAttribute{
						MarkdownDescription: `Type of subnetting of the VLAN. Applicable only for template network.`,
						Computed:            true,
					},
					"vpn_nat_subnet": schema.StringAttribute{
						MarkdownDescription: `The translated VPN subnet if VPN and VPN subnet translation are enabled on the VLAN`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseApplianceGetNetworkApplianceVlans`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"appliance_ip": schema.StringAttribute{
							MarkdownDescription: `The local IP of the appliance on the VLAN`,
							Computed:            true,
						},
						"cidr": schema.StringAttribute{
							MarkdownDescription: `CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN.`,
							Computed:            true,
						},
						"dhcp_boot_filename": schema.StringAttribute{
							MarkdownDescription: `DHCP boot option for boot filename`,
							Computed:            true,
						},
						"dhcp_boot_next_server": schema.StringAttribute{
							MarkdownDescription: `DHCP boot option to direct boot clients to the server to load the boot file from`,
							Computed:            true,
						},
						"dhcp_boot_options_enabled": schema.BoolAttribute{
							MarkdownDescription: `Use DHCP boot options specified in other properties`,
							Computed:            true,
						},
						"dhcp_handling": schema.StringAttribute{
							MarkdownDescription: `The appliance's handling of DHCP requests on this VLAN. One of: 'Run a DHCP server', 'Relay DHCP to another server' or 'Do not respond to DHCP requests'`,
							Computed:            true,
						},
						"dhcp_lease_time": schema.StringAttribute{
							MarkdownDescription: `The term of DHCP leases if the appliance is running a DHCP server on this VLAN. One of: '30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week'`,
							Computed:            true,
						},
						"dhcp_options": schema.SetNestedAttribute{
							MarkdownDescription: `The list of DHCP options that will be included in DHCP responses. Each object in the list should have "code", "type", and "value" properties.`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"code": schema.StringAttribute{
										MarkdownDescription: `The code for the DHCP option. This should be an integer between 2 and 254.`,
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: `The type for the DHCP option. One of: 'text', 'ip', 'hex' or 'integer'`,
										Computed:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: `The value for the DHCP option`,
										Computed:            true,
									},
								},
							},
						},
						"dhcp_relay_server_ips": schema.ListAttribute{
							MarkdownDescription: `The IPs of the DHCP servers that DHCP requests should be relayed to`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"dns_nameservers": schema.StringAttribute{
							MarkdownDescription: `The DNS nameservers used for DHCP responses, either "upstream_dns", "google_dns", "opendns", or a newline seperated string of IP addresses or domain names`,
							Computed:            true,
						},
						"fixed_ip_assignments": schema.StringAttribute{
							MarkdownDescription: `The DHCP fixed IP assignments on the VLAN. This should be an object that contains mappings from MAC addresses to objects that themselves each contain "ip" and "name" string fields. See the sample request/response for more details.`,
							Computed:            true,
						},
						"group_policy_id": schema.StringAttribute{
							MarkdownDescription: `The id of the desired group policy to apply to the VLAN`,
							Computed:            true,
						},
						"id": schema.Int64Attribute{
							MarkdownDescription: `The VLAN ID of the VLAN`,
							Computed:            true,
						},
						"interface_id": schema.StringAttribute{
							MarkdownDescription: `The interface ID of the VLAN`,
							Computed:            true,
						},
						"ipv6": schema.SingleNestedAttribute{
							MarkdownDescription: `IPv6 configuration on the VLAN`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"enabled": schema.BoolAttribute{
									MarkdownDescription: `Enable IPv6 on VLAN`,
									Computed:            true,
								},
								"prefix_assignments": schema.SetNestedAttribute{
									MarkdownDescription: `Prefix assignments on the VLAN`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"autonomous": schema.BoolAttribute{
												MarkdownDescription: `Auto assign a /64 prefix from the origin to the VLAN`,
												Computed:            true,
											},
											"origin": schema.SingleNestedAttribute{
												MarkdownDescription: `The origin of the prefix`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"interfaces": schema.ListAttribute{
														MarkdownDescription: `Interfaces associated with the prefix`,
														Computed:            true,
														ElementType:         types.StringType,
													},
													"type": schema.StringAttribute{
														MarkdownDescription: `Type of the origin`,
														Computed:            true,
													},
												},
											},
											"static_appliance_ip6": schema.StringAttribute{
												MarkdownDescription: `Manual configuration of the IPv6 Appliance IP`,
												Computed:            true,
											},
											"static_prefix": schema.StringAttribute{
												MarkdownDescription: `Manual configuration of a /64 prefix on the VLAN`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
						"mandatory_dhcp": schema.SingleNestedAttribute{
							MarkdownDescription: `Mandatory DHCP will enforce that clients connecting to this VLAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"enabled": schema.BoolAttribute{
									MarkdownDescription: `Enable Mandatory DHCP on VLAN.`,
									Computed:            true,
								},
							},
						},
						"mask": schema.Int64Attribute{
							MarkdownDescription: `Mask used for the subnet of all bound to the template networks. Applicable only for template network.`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the VLAN`,
							Computed:            true,
						},
						"reserved_ip_ranges": schema.SetNestedAttribute{
							MarkdownDescription: `The DHCP reserved IP ranges on the VLAN`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"comment": schema.StringAttribute{
										MarkdownDescription: `A text comment for the reserved range`,
										Computed:            true,
									},
									"end": schema.StringAttribute{
										MarkdownDescription: `The last IP in the reserved range`,
										Computed:            true,
									},
									"start": schema.StringAttribute{
										MarkdownDescription: `The first IP in the reserved range`,
										Computed:            true,
									},
								},
							},
						},
						"subnet": schema.StringAttribute{
							MarkdownDescription: `The subnet of the VLAN`,
							Computed:            true,
						},
						"template_vlan_type": schema.StringAttribute{
							MarkdownDescription: `Type of subnetting of the VLAN. Applicable only for template network.`,
							Computed:            true,
						},
						"vpn_nat_subnet": schema.StringAttribute{
							MarkdownDescription: `The translated VPN subnet if VPN and VPN subnet translation are enabled on the VLAN`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksApplianceVLANsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceVLANs NetworksApplianceVLANs
	diags := req.Config.Get(ctx, &networksApplianceVLANs)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksApplianceVLANs.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksApplianceVLANs.NetworkID.IsNull(), !networksApplianceVLANs.VLANID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceVLANs")
		vvNetworkID := networksApplianceVLANs.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceVLANs(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceVLANs",
				err.Error(),
			)
			return
		}

		networksApplianceVLANs = ResponseApplianceGetNetworkApplianceVLANsItemsToBody(networksApplianceVLANs, response1)
		diags = resp.State.Set(ctx, &networksApplianceVLANs)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceVLAN")
		vvNetworkID := networksApplianceVLANs.NetworkID.ValueString()
		vvVLANID := networksApplianceVLANs.VLANID.ValueString()

		response2, restyResp2, err := d.client.Appliance.GetNetworkApplianceVLAN(vvNetworkID, vvVLANID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceVLAN",
				err.Error(),
			)
			return
		}

		networksApplianceVLANs = ResponseApplianceGetNetworkApplianceVLANItemToBody(networksApplianceVLANs, response2)
		diags = resp.State.Set(ctx, &networksApplianceVLANs)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceVLANs struct {
	NetworkID types.String                                     `tfsdk:"network_id"`
	VLANID    types.String                                     `tfsdk:"vlan_id"`
	Items     *[]ResponseItemApplianceGetNetworkApplianceVlans `tfsdk:"items"`
	Item      *ResponseApplianceGetNetworkApplianceVlan        `tfsdk:"item"`
}

type ResponseItemApplianceGetNetworkApplianceVlans struct {
	ApplianceIP            types.String                                                     `tfsdk:"appliance_ip"`
	Cidr                   types.String                                                     `tfsdk:"cidr"`
	DhcpBootFilename       types.String                                                     `tfsdk:"dhcp_boot_filename"`
	DhcpBootNextServer     types.String                                                     `tfsdk:"dhcp_boot_next_server"`
	DhcpBootOptionsEnabled types.Bool                                                       `tfsdk:"dhcp_boot_options_enabled"`
	DhcpHandling           types.String                                                     `tfsdk:"dhcp_handling"`
	DhcpLeaseTime          types.String                                                     `tfsdk:"dhcp_lease_time"`
	DhcpOptions            *[]ResponseItemApplianceGetNetworkApplianceVlansDhcpOptions      `tfsdk:"dhcp_options"`
	DhcpRelayServerIPs     types.List                                                       `tfsdk:"dhcp_relay_server_ips"`
	DNSNameservers         types.String                                                     `tfsdk:"dns_nameservers"`
	FixedIPAssignments     types.String                                                     `tfsdk:"fixed_ip_assignments"`
	GroupPolicyID          types.String                                                     `tfsdk:"group_policy_id"`
	ID                     types.Int64                                                      `tfsdk:"id"`
	InterfaceID            types.String                                                     `tfsdk:"interface_id"`
	IPv6                   *ResponseItemApplianceGetNetworkApplianceVlansIpv6               `tfsdk:"ipv6"`
	MandatoryDhcp          *ResponseItemApplianceGetNetworkApplianceVlansMandatoryDhcp      `tfsdk:"mandatory_dhcp"`
	Mask                   types.Int64                                                      `tfsdk:"mask"`
	Name                   types.String                                                     `tfsdk:"name"`
	ReservedIPRanges       *[]ResponseItemApplianceGetNetworkApplianceVlansReservedIpRanges `tfsdk:"reserved_ip_ranges"`
	Subnet                 types.String                                                     `tfsdk:"subnet"`
	TemplateVLANType       types.String                                                     `tfsdk:"template_vlan_type"`
	VpnNatSubnet           types.String                                                     `tfsdk:"vpn_nat_subnet"`
}

type ResponseItemApplianceGetNetworkApplianceVlansDhcpOptions struct {
	Code  types.String `tfsdk:"code"`
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

// type ResponseItemApplianceGetNetworkApplianceVlansFixedIpAssignments interface{}

type ResponseItemApplianceGetNetworkApplianceVlansIpv6 struct {
	Enabled           types.Bool                                                            `tfsdk:"enabled"`
	PrefixAssignments *[]ResponseItemApplianceGetNetworkApplianceVlansIpv6PrefixAssignments `tfsdk:"prefix_assignments"`
}

type ResponseItemApplianceGetNetworkApplianceVlansIpv6PrefixAssignments struct {
	Autonomous         types.Bool                                                                `tfsdk:"autonomous"`
	Origin             *ResponseItemApplianceGetNetworkApplianceVlansIpv6PrefixAssignmentsOrigin `tfsdk:"origin"`
	StaticApplianceIP6 types.String                                                              `tfsdk:"static_appliance_ip6"`
	StaticPrefix       types.String                                                              `tfsdk:"static_prefix"`
}

type ResponseItemApplianceGetNetworkApplianceVlansIpv6PrefixAssignmentsOrigin struct {
	Interfaces types.List   `tfsdk:"interfaces"`
	Type       types.String `tfsdk:"type"`
}

type ResponseItemApplianceGetNetworkApplianceVlansMandatoryDhcp struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseItemApplianceGetNetworkApplianceVlansReservedIpRanges struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

type ResponseApplianceGetNetworkApplianceVlan struct {
	ApplianceIP            types.String                                                `tfsdk:"appliance_ip"`
	Cidr                   types.String                                                `tfsdk:"cidr"`
	DhcpBootFilename       types.String                                                `tfsdk:"dhcp_boot_filename"`
	DhcpBootNextServer     types.String                                                `tfsdk:"dhcp_boot_next_server"`
	DhcpBootOptionsEnabled types.Bool                                                  `tfsdk:"dhcp_boot_options_enabled"`
	DhcpHandling           types.String                                                `tfsdk:"dhcp_handling"`
	DhcpLeaseTime          types.String                                                `tfsdk:"dhcp_lease_time"`
	DhcpOptions            *[]ResponseApplianceGetNetworkApplianceVlanDhcpOptions      `tfsdk:"dhcp_options"`
	DhcpRelayServerIPs     types.List                                                  `tfsdk:"dhcp_relay_server_ips"`
	DNSNameservers         types.String                                                `tfsdk:"dns_nameservers"`
	FixedIPAssignments     *ResponseApplianceGetNetworkApplianceVlanFixedIpAssignments `tfsdk:"fixed_ip_assignments"`
	GroupPolicyID          types.String                                                `tfsdk:"group_policy_id"`
	ID                     types.Int64                                                 `tfsdk:"id"`
	InterfaceID            types.String                                                `tfsdk:"interface_id"`
	IPv6                   *ResponseApplianceGetNetworkApplianceVlanIpv6               `tfsdk:"ipv6"`
	MandatoryDhcp          *ResponseApplianceGetNetworkApplianceVlanMandatoryDhcp      `tfsdk:"mandatory_dhcp"`
	Mask                   types.Int64                                                 `tfsdk:"mask"`
	Name                   types.String                                                `tfsdk:"name"`
	ReservedIPRanges       *[]ResponseApplianceGetNetworkApplianceVlanReservedIpRanges `tfsdk:"reserved_ip_ranges"`
	Subnet                 types.String                                                `tfsdk:"subnet"`
	TemplateVLANType       types.String                                                `tfsdk:"template_vlan_type"`
	VpnNatSubnet           types.String                                                `tfsdk:"vpn_nat_subnet"`
}

type ResponseApplianceGetNetworkApplianceVlanDhcpOptions struct {
	Code  types.String `tfsdk:"code"`
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type ResponseApplianceGetNetworkApplianceVlanFixedIpAssignments interface{}

type ResponseApplianceGetNetworkApplianceVlanIpv6 struct {
	Enabled           types.Bool                                                       `tfsdk:"enabled"`
	PrefixAssignments *[]ResponseApplianceGetNetworkApplianceVlanIpv6PrefixAssignments `tfsdk:"prefix_assignments"`
}

type ResponseApplianceGetNetworkApplianceVlanIpv6PrefixAssignments struct {
	Autonomous         types.Bool                                                           `tfsdk:"autonomous"`
	Origin             *ResponseApplianceGetNetworkApplianceVlanIpv6PrefixAssignmentsOrigin `tfsdk:"origin"`
	StaticApplianceIP6 types.String                                                         `tfsdk:"static_appliance_ip6"`
	StaticPrefix       types.String                                                         `tfsdk:"static_prefix"`
}

type ResponseApplianceGetNetworkApplianceVlanIpv6PrefixAssignmentsOrigin struct {
	Interfaces types.List   `tfsdk:"interfaces"`
	Type       types.String `tfsdk:"type"`
}

type ResponseApplianceGetNetworkApplianceVlanMandatoryDhcp struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseApplianceGetNetworkApplianceVlanReservedIpRanges struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceVLANsItemsToBody(state NetworksApplianceVLANs, response *merakigosdk.ResponseApplianceGetNetworkApplianceVLANs) NetworksApplianceVLANs {
	var items []ResponseItemApplianceGetNetworkApplianceVlans
	for _, item := range *response {
		itemState := ResponseItemApplianceGetNetworkApplianceVlans{
			ApplianceIP:        types.StringValue(item.ApplianceIP),
			Cidr:               types.StringValue(item.Cidr),
			DhcpBootFilename:   types.StringValue(item.DhcpBootFilename),
			DhcpBootNextServer: types.StringValue(item.DhcpBootNextServer),
			DhcpBootOptionsEnabled: func() types.Bool {
				if item.DhcpBootOptionsEnabled != nil {
					return types.BoolValue(*item.DhcpBootOptionsEnabled)
				}
				return types.Bool{}
			}(),
			DhcpHandling:  types.StringValue(item.DhcpHandling),
			DhcpLeaseTime: types.StringValue(item.DhcpLeaseTime),
			DhcpOptions: func() *[]ResponseItemApplianceGetNetworkApplianceVlansDhcpOptions {
				if item.DhcpOptions != nil {
					result := make([]ResponseItemApplianceGetNetworkApplianceVlansDhcpOptions, len(*item.DhcpOptions))
					for i, dhcpOptions := range *item.DhcpOptions {
						result[i] = ResponseItemApplianceGetNetworkApplianceVlansDhcpOptions{
							Code:  types.StringValue(dhcpOptions.Code),
							Type:  types.StringValue(dhcpOptions.Type),
							Value: types.StringValue(dhcpOptions.Value),
						}
					}
					return &result
				}
				return &[]ResponseItemApplianceGetNetworkApplianceVlansDhcpOptions{}
			}(),
			DhcpRelayServerIPs: StringSliceToList(item.DhcpRelayServerIPs),
			DNSNameservers:     types.StringValue(item.DNSNameservers),
			// FixedIPAssignments: types.StringValue(fmt.Sprintf("%v", item.FixedIPAssignments)), //TODO POSIBLE interface
			GroupPolicyID: types.StringValue(item.GroupPolicyID),
			ID: func() types.Int64 {
				if item.ID != nil {
					return types.Int64Value(int64(*item.ID))
				}
				return types.Int64{}
			}(),
			InterfaceID: types.StringValue(item.InterfaceID),
			IPv6: func() *ResponseItemApplianceGetNetworkApplianceVlansIpv6 {
				if item.IPv6 != nil {
					return &ResponseItemApplianceGetNetworkApplianceVlansIpv6{
						Enabled: func() types.Bool {
							if item.IPv6.Enabled != nil {
								return types.BoolValue(*item.IPv6.Enabled)
							}
							return types.Bool{}
						}(),
						PrefixAssignments: func() *[]ResponseItemApplianceGetNetworkApplianceVlansIpv6PrefixAssignments {
							if item.IPv6.PrefixAssignments != nil {
								result := make([]ResponseItemApplianceGetNetworkApplianceVlansIpv6PrefixAssignments, len(*item.IPv6.PrefixAssignments))
								for i, prefixAssignments := range *item.IPv6.PrefixAssignments {
									result[i] = ResponseItemApplianceGetNetworkApplianceVlansIpv6PrefixAssignments{
										Autonomous: func() types.Bool {
											if prefixAssignments.Autonomous != nil {
												return types.BoolValue(*prefixAssignments.Autonomous)
											}
											return types.Bool{}
										}(),
										Origin: func() *ResponseItemApplianceGetNetworkApplianceVlansIpv6PrefixAssignmentsOrigin {
											if prefixAssignments.Origin != nil {
												return &ResponseItemApplianceGetNetworkApplianceVlansIpv6PrefixAssignmentsOrigin{
													Interfaces: StringSliceToList(prefixAssignments.Origin.Interfaces),
													Type:       types.StringValue(prefixAssignments.Origin.Type),
												}
											}
											return &ResponseItemApplianceGetNetworkApplianceVlansIpv6PrefixAssignmentsOrigin{}
										}(),
										StaticApplianceIP6: types.StringValue(prefixAssignments.StaticApplianceIP6),
										StaticPrefix:       types.StringValue(prefixAssignments.StaticPrefix),
									}
								}
								return &result
							}
							return &[]ResponseItemApplianceGetNetworkApplianceVlansIpv6PrefixAssignments{}
						}(),
					}
				}
				return &ResponseItemApplianceGetNetworkApplianceVlansIpv6{}
			}(),
			MandatoryDhcp: func() *ResponseItemApplianceGetNetworkApplianceVlansMandatoryDhcp {
				if item.MandatoryDhcp != nil {
					return &ResponseItemApplianceGetNetworkApplianceVlansMandatoryDhcp{
						Enabled: func() types.Bool {
							if item.MandatoryDhcp.Enabled != nil {
								return types.BoolValue(*item.MandatoryDhcp.Enabled)
							}
							return types.Bool{}
						}(),
					}
				}
				return &ResponseItemApplianceGetNetworkApplianceVlansMandatoryDhcp{}
			}(),
			Mask: func() types.Int64 {
				if item.Mask != nil {
					return types.Int64Value(int64(*item.Mask))
				}
				return types.Int64{}
			}(),
			Name: types.StringValue(item.Name),
			ReservedIPRanges: func() *[]ResponseItemApplianceGetNetworkApplianceVlansReservedIpRanges {
				if item.ReservedIPRanges != nil {
					result := make([]ResponseItemApplianceGetNetworkApplianceVlansReservedIpRanges, len(*item.ReservedIPRanges))
					for i, reservedIPRanges := range *item.ReservedIPRanges {
						result[i] = ResponseItemApplianceGetNetworkApplianceVlansReservedIpRanges{
							Comment: types.StringValue(reservedIPRanges.Comment),
							End:     types.StringValue(reservedIPRanges.End),
							Start:   types.StringValue(reservedIPRanges.Start),
						}
					}
					return &result
				}
				return &[]ResponseItemApplianceGetNetworkApplianceVlansReservedIpRanges{}
			}(),
			Subnet:           types.StringValue(item.Subnet),
			TemplateVLANType: types.StringValue(item.TemplateVLANType),
			VpnNatSubnet:     types.StringValue(item.VpnNatSubnet),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseApplianceGetNetworkApplianceVLANItemToBody(state NetworksApplianceVLANs, response *merakigosdk.ResponseApplianceGetNetworkApplianceVLAN) NetworksApplianceVLANs {
	itemState := ResponseApplianceGetNetworkApplianceVlan{
		ApplianceIP:        types.StringValue(response.ApplianceIP),
		Cidr:               types.StringValue(response.Cidr),
		DhcpBootFilename:   types.StringValue(response.DhcpBootFilename),
		DhcpBootNextServer: types.StringValue(response.DhcpBootNextServer),
		DhcpBootOptionsEnabled: func() types.Bool {
			if response.DhcpBootOptionsEnabled != nil {
				return types.BoolValue(*response.DhcpBootOptionsEnabled)
			}
			return types.Bool{}
		}(),
		DhcpHandling:  types.StringValue(response.DhcpHandling),
		DhcpLeaseTime: types.StringValue(response.DhcpLeaseTime),
		DhcpOptions: func() *[]ResponseApplianceGetNetworkApplianceVlanDhcpOptions {
			if response.DhcpOptions != nil {
				result := make([]ResponseApplianceGetNetworkApplianceVlanDhcpOptions, len(*response.DhcpOptions))
				for i, dhcpOptions := range *response.DhcpOptions {
					result[i] = ResponseApplianceGetNetworkApplianceVlanDhcpOptions{
						Code:  types.StringValue(dhcpOptions.Code),
						Type:  types.StringValue(dhcpOptions.Type),
						Value: types.StringValue(dhcpOptions.Value),
					}
				}
				return &result
			}
			return &[]ResponseApplianceGetNetworkApplianceVlanDhcpOptions{}
		}(),
		DhcpRelayServerIPs: StringSliceToList(response.DhcpRelayServerIPs),
		DNSNameservers:     types.StringValue(response.DNSNameservers),
		// FixedIPAssignments: types.StringValue(response.FixedIPAssignments), //TODO POSIBLE interface
		GroupPolicyID: types.StringValue(response.GroupPolicyID),
		ID: func() types.Int64 {
			if response.ID != nil {
				return types.Int64Value(int64(*response.ID))
			}
			return types.Int64{}
		}(),
		InterfaceID: types.StringValue(response.InterfaceID),
		IPv6: func() *ResponseApplianceGetNetworkApplianceVlanIpv6 {
			if response.IPv6 != nil {
				return &ResponseApplianceGetNetworkApplianceVlanIpv6{
					Enabled: func() types.Bool {
						if response.IPv6.Enabled != nil {
							return types.BoolValue(*response.IPv6.Enabled)
						}
						return types.Bool{}
					}(),
					PrefixAssignments: func() *[]ResponseApplianceGetNetworkApplianceVlanIpv6PrefixAssignments {
						if response.IPv6.PrefixAssignments != nil {
							result := make([]ResponseApplianceGetNetworkApplianceVlanIpv6PrefixAssignments, len(*response.IPv6.PrefixAssignments))
							for i, prefixAssignments := range *response.IPv6.PrefixAssignments {
								result[i] = ResponseApplianceGetNetworkApplianceVlanIpv6PrefixAssignments{
									Autonomous: func() types.Bool {
										if prefixAssignments.Autonomous != nil {
											return types.BoolValue(*prefixAssignments.Autonomous)
										}
										return types.Bool{}
									}(),
									Origin: func() *ResponseApplianceGetNetworkApplianceVlanIpv6PrefixAssignmentsOrigin {
										if prefixAssignments.Origin != nil {
											return &ResponseApplianceGetNetworkApplianceVlanIpv6PrefixAssignmentsOrigin{
												Interfaces: StringSliceToList(prefixAssignments.Origin.Interfaces),
												Type:       types.StringValue(prefixAssignments.Origin.Type),
											}
										}
										return &ResponseApplianceGetNetworkApplianceVlanIpv6PrefixAssignmentsOrigin{}
									}(),
									StaticApplianceIP6: types.StringValue(prefixAssignments.StaticApplianceIP6),
									StaticPrefix:       types.StringValue(prefixAssignments.StaticPrefix),
								}
							}
							return &result
						}
						return &[]ResponseApplianceGetNetworkApplianceVlanIpv6PrefixAssignments{}
					}(),
				}
			}
			return &ResponseApplianceGetNetworkApplianceVlanIpv6{}
		}(),
		MandatoryDhcp: func() *ResponseApplianceGetNetworkApplianceVlanMandatoryDhcp {
			if response.MandatoryDhcp != nil {
				return &ResponseApplianceGetNetworkApplianceVlanMandatoryDhcp{
					Enabled: func() types.Bool {
						if response.MandatoryDhcp.Enabled != nil {
							return types.BoolValue(*response.MandatoryDhcp.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseApplianceGetNetworkApplianceVlanMandatoryDhcp{}
		}(),
		Mask: func() types.Int64 {
			if response.Mask != nil {
				return types.Int64Value(int64(*response.Mask))
			}
			return types.Int64{}
		}(),
		Name: types.StringValue(response.Name),
		ReservedIPRanges: func() *[]ResponseApplianceGetNetworkApplianceVlanReservedIpRanges {
			if response.ReservedIPRanges != nil {
				result := make([]ResponseApplianceGetNetworkApplianceVlanReservedIpRanges, len(*response.ReservedIPRanges))
				for i, reservedIPRanges := range *response.ReservedIPRanges {
					result[i] = ResponseApplianceGetNetworkApplianceVlanReservedIpRanges{
						Comment: types.StringValue(reservedIPRanges.Comment),
						End:     types.StringValue(reservedIPRanges.End),
						Start:   types.StringValue(reservedIPRanges.Start),
					}
				}
				return &result
			}
			return &[]ResponseApplianceGetNetworkApplianceVlanReservedIpRanges{}
		}(),
		Subnet:           types.StringValue(response.Subnet),
		TemplateVLANType: types.StringValue(response.TemplateVLANType),
		VpnNatSubnet:     types.StringValue(response.VpnNatSubnet),
	}
	state.Item = &itemState
	return state
}
