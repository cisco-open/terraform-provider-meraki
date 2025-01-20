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
	_ datasource.DataSource              = &NetworksApplianceSingleLanDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceSingleLanDataSource{}
)

func NewNetworksApplianceSingleLanDataSource() datasource.DataSource {
	return &NetworksApplianceSingleLanDataSource{}
}

type NetworksApplianceSingleLanDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceSingleLanDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceSingleLanDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_single_lan"
}

func (d *NetworksApplianceSingleLanDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"appliance_ip": schema.StringAttribute{
						MarkdownDescription: `The local IP of the appliance on the single LAN`,
						Computed:            true,
					},
					"ipv6": schema.SingleNestedAttribute{
						MarkdownDescription: `IPv6 configuration on the single LAN`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Enable IPv6 on single LAN`,
								Computed:            true,
							},
							"prefix_assignments": schema.SetNestedAttribute{
								MarkdownDescription: `Prefix assignments on the single LAN`,
								Computed:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{

										"autonomous": schema.BoolAttribute{
											MarkdownDescription: `Auto assign a /64 prefix from the origin to the single LAN`,
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
											MarkdownDescription: `Manual configuration of a /64 prefix on the single LAN`,
											Computed:            true,
										},
									},
								},
							},
						},
					},
					"mandatory_dhcp": schema.SingleNestedAttribute{
						MarkdownDescription: `Mandatory DHCP will enforce that clients connecting to this single LAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Enable Mandatory DHCP on single LAN.`,
								Computed:            true,
							},
						},
					},
					"subnet": schema.StringAttribute{
						MarkdownDescription: `The subnet of the single LAN`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *NetworksApplianceSingleLanDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceSingleLan NetworksApplianceSingleLan
	diags := req.Config.Get(ctx, &networksApplianceSingleLan)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceSingleLan")
		vvNetworkID := networksApplianceSingleLan.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceSingleLan(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceSingleLan",
				err.Error(),
			)
			return
		}

		networksApplianceSingleLan = ResponseApplianceGetNetworkApplianceSingleLanItemToBody(networksApplianceSingleLan, response1)
		diags = resp.State.Set(ctx, &networksApplianceSingleLan)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceSingleLan struct {
	NetworkID types.String                                   `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceSingleLan `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceSingleLan struct {
	ApplianceIP   types.String                                                `tfsdk:"appliance_ip"`
	IPv6          *ResponseApplianceGetNetworkApplianceSingleLanIpv6          `tfsdk:"ipv6"`
	MandatoryDhcp *ResponseApplianceGetNetworkApplianceSingleLanMandatoryDhcp `tfsdk:"mandatory_dhcp"`
	Subnet        types.String                                                `tfsdk:"subnet"`
}

type ResponseApplianceGetNetworkApplianceSingleLanIpv6 struct {
	Enabled           types.Bool                                                            `tfsdk:"enabled"`
	PrefixAssignments *[]ResponseApplianceGetNetworkApplianceSingleLanIpv6PrefixAssignments `tfsdk:"prefix_assignments"`
}

type ResponseApplianceGetNetworkApplianceSingleLanIpv6PrefixAssignments struct {
	Autonomous         types.Bool                                                                `tfsdk:"autonomous"`
	Origin             *ResponseApplianceGetNetworkApplianceSingleLanIpv6PrefixAssignmentsOrigin `tfsdk:"origin"`
	StaticApplianceIP6 types.String                                                              `tfsdk:"static_appliance_ip6"`
	StaticPrefix       types.String                                                              `tfsdk:"static_prefix"`
}

type ResponseApplianceGetNetworkApplianceSingleLanIpv6PrefixAssignmentsOrigin struct {
	Interfaces types.List   `tfsdk:"interfaces"`
	Type       types.String `tfsdk:"type"`
}

type ResponseApplianceGetNetworkApplianceSingleLanMandatoryDhcp struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceSingleLanItemToBody(state NetworksApplianceSingleLan, response *merakigosdk.ResponseApplianceGetNetworkApplianceSingleLan) NetworksApplianceSingleLan {
	itemState := ResponseApplianceGetNetworkApplianceSingleLan{
		ApplianceIP: types.StringValue(response.ApplianceIP),
		IPv6: func() *ResponseApplianceGetNetworkApplianceSingleLanIpv6 {
			if response.IPv6 != nil {
				return &ResponseApplianceGetNetworkApplianceSingleLanIpv6{
					Enabled: func() types.Bool {
						if response.IPv6.Enabled != nil {
							return types.BoolValue(*response.IPv6.Enabled)
						}
						return types.Bool{}
					}(),
					PrefixAssignments: func() *[]ResponseApplianceGetNetworkApplianceSingleLanIpv6PrefixAssignments {
						if response.IPv6.PrefixAssignments != nil {
							result := make([]ResponseApplianceGetNetworkApplianceSingleLanIpv6PrefixAssignments, len(*response.IPv6.PrefixAssignments))
							for i, prefixAssignments := range *response.IPv6.PrefixAssignments {
								result[i] = ResponseApplianceGetNetworkApplianceSingleLanIpv6PrefixAssignments{
									Autonomous: func() types.Bool {
										if prefixAssignments.Autonomous != nil {
											return types.BoolValue(*prefixAssignments.Autonomous)
										}
										return types.Bool{}
									}(),
									Origin: func() *ResponseApplianceGetNetworkApplianceSingleLanIpv6PrefixAssignmentsOrigin {
										if prefixAssignments.Origin != nil {
											return &ResponseApplianceGetNetworkApplianceSingleLanIpv6PrefixAssignmentsOrigin{
												Interfaces: StringSliceToList(prefixAssignments.Origin.Interfaces),
												Type:       types.StringValue(prefixAssignments.Origin.Type),
											}
										}
										return nil
									}(),
									StaticApplianceIP6: types.StringValue(prefixAssignments.StaticApplianceIP6),
									StaticPrefix:       types.StringValue(prefixAssignments.StaticPrefix),
								}
							}
							return &result
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		MandatoryDhcp: func() *ResponseApplianceGetNetworkApplianceSingleLanMandatoryDhcp {
			if response.MandatoryDhcp != nil {
				return &ResponseApplianceGetNetworkApplianceSingleLanMandatoryDhcp{
					Enabled: func() types.Bool {
						if response.MandatoryDhcp.Enabled != nil {
							return types.BoolValue(*response.MandatoryDhcp.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
		}(),
		Subnet: types.StringValue(response.Subnet),
	}
	state.Item = &itemState
	return state
}
