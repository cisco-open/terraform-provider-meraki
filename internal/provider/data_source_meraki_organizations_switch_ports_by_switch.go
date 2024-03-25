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
	_ datasource.DataSource              = &OrganizationsSwitchPortsBySwitchDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSwitchPortsBySwitchDataSource{}
)

func NewOrganizationsSwitchPortsBySwitchDataSource() datasource.DataSource {
	return &OrganizationsSwitchPortsBySwitchDataSource{}
}

type OrganizationsSwitchPortsBySwitchDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSwitchPortsBySwitchDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSwitchPortsBySwitchDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_switch_ports_by_switch"
}

func (d *OrganizationsSwitchPortsBySwitchDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"configuration_updated_after": schema.StringAttribute{
				MarkdownDescription: `configurationUpdatedAfter query parameter. Optional parameter to filter results by switches where the configuration has been updated after the given timestamp.`,
				Optional:            true,
			},
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"mac": schema.StringAttribute{
				MarkdownDescription: `mac query parameter. Optional parameter to filter switchports belonging to switches by MAC address. All returned switches will have a MAC address that contains the search term or is an exact match.`,
				Optional:            true,
			},
			"macs": schema.ListAttribute{
				MarkdownDescription: `macs query parameter. Optional parameter to filter switchports by one or more MAC addresses belonging to devices. All switchports returned belong to MAC addresses of switches that are an exact match.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `name query parameter. Optional parameter to filter switchports belonging to switches by name. All returned switches will have a name that contains the search term or is an exact match.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter switchports by network.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 50. Default is 50.`,
				Optional:            true,
			},
			"port_profile_ids": schema.ListAttribute{
				MarkdownDescription: `portProfileIds query parameter. Optional parameter to filter switchports belonging to the specified port profiles.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial query parameter. Optional parameter to filter switchports belonging to switches by serial number. All returned switches will have a serial number that contains the search term or is an exact match.`,
				Optional:            true,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Optional parameter to filter switchports belonging to switches with one or more serial numbers. All switchports returned belong to serial numbers of switches that are an exact match.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSwitchGetOrganizationSwitchPortsBySwitch`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"mac": schema.StringAttribute{
							MarkdownDescription: `The MAC address of the switch.`,
							Computed:            true,
						},
						"model": schema.StringAttribute{
							MarkdownDescription: `The model of the switch.`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the switch.`,
							Computed:            true,
						},
						"network": schema.SingleNestedAttribute{
							MarkdownDescription: `Identifying information of the switch's network.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `The ID of the network.`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `The name of the network.`,
									Computed:            true,
								},
							},
						},
						"ports": schema.SetNestedAttribute{
							MarkdownDescription: `Ports belonging to the switch`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"access_policy_type": schema.StringAttribute{
										MarkdownDescription: `The type of the access policy of the switch port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.`,
										Computed:            true,
									},
									"allowed_vlans": schema.StringAttribute{
										MarkdownDescription: `The VLANs allowed on the switch port. Only applicable to trunk ports.`,
										Computed:            true,
									},
									"enabled": schema.BoolAttribute{
										MarkdownDescription: `The status of the switch port.`,
										Computed:            true,
									},
									"link_negotiation": schema.StringAttribute{
										MarkdownDescription: `The link speed for the switch port.`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `The name of the switch port.`,
										Computed:            true,
									},
									"poe_enabled": schema.BoolAttribute{
										MarkdownDescription: `The PoE status of the switch port.`,
										Computed:            true,
									},
									"port_id": schema.StringAttribute{
										MarkdownDescription: `The identifier of the switch port.`,
										Computed:            true,
									},
									"rstp_enabled": schema.BoolAttribute{
										MarkdownDescription: `The rapid spanning tree protocol status.`,
										Computed:            true,
									},
									"sticky_mac_allow_list": schema.ListAttribute{
										MarkdownDescription: `The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.`,
										Computed:            true,
										ElementType:         types.StringType,
									},
									"sticky_mac_allow_list_limit": schema.Int64Attribute{
										MarkdownDescription: `The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.`,
										Computed:            true,
									},
									"stp_guard": schema.StringAttribute{
										MarkdownDescription: `The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').`,
										Computed:            true,
									},
									"tags": schema.ListAttribute{
										MarkdownDescription: `The list of tags of the switch port.`,
										Computed:            true,
										ElementType:         types.StringType,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: `The type of the switch port ('trunk' or 'access').`,
										Computed:            true,
									},
									"vlan": schema.Int64Attribute{
										MarkdownDescription: `The VLAN of the switch port. A null value will clear the value set for trunk ports.`,
										Computed:            true,
									},
									"voice_vlan": schema.Int64Attribute{
										MarkdownDescription: `The voice VLAN of the switch port. Only applicable to access ports.`,
										Computed:            true,
									},
								},
							},
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `The serial number of the switch.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsSwitchPortsBySwitchDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSwitchPortsBySwitch OrganizationsSwitchPortsBySwitch
	diags := req.Config.Get(ctx, &organizationsSwitchPortsBySwitch)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSwitchPortsBySwitch")
		vvOrganizationID := organizationsSwitchPortsBySwitch.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationSwitchPortsBySwitchQueryParams{}

		queryParams1.PerPage = int(organizationsSwitchPortsBySwitch.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsSwitchPortsBySwitch.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsSwitchPortsBySwitch.EndingBefore.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsSwitchPortsBySwitch.NetworkIDs)
		queryParams1.PortProfileIDs = elementsToStrings(ctx, organizationsSwitchPortsBySwitch.PortProfileIDs)
		queryParams1.Name = organizationsSwitchPortsBySwitch.Name.ValueString()
		queryParams1.Mac = organizationsSwitchPortsBySwitch.Mac.ValueString()
		queryParams1.Macs = elementsToStrings(ctx, organizationsSwitchPortsBySwitch.Macs)
		queryParams1.Serial = organizationsSwitchPortsBySwitch.Serial.ValueString()
		queryParams1.Serials = elementsToStrings(ctx, organizationsSwitchPortsBySwitch.Serials)
		queryParams1.ConfigurationUpdatedAfter = organizationsSwitchPortsBySwitch.ConfigurationUpdatedAfter.ValueString()

		response1, restyResp1, err := d.client.Switch.GetOrganizationSwitchPortsBySwitch(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSwitchPortsBySwitch",
				err.Error(),
			)
			return
		}

		organizationsSwitchPortsBySwitch = ResponseSwitchGetOrganizationSwitchPortsBySwitchItemsToBody(organizationsSwitchPortsBySwitch, response1)
		diags = resp.State.Set(ctx, &organizationsSwitchPortsBySwitch)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSwitchPortsBySwitch struct {
	OrganizationID            types.String                                            `tfsdk:"organization_id"`
	PerPage                   types.Int64                                             `tfsdk:"per_page"`
	StartingAfter             types.String                                            `tfsdk:"starting_after"`
	EndingBefore              types.String                                            `tfsdk:"ending_before"`
	NetworkIDs                types.List                                              `tfsdk:"network_ids"`
	PortProfileIDs            types.List                                              `tfsdk:"port_profile_ids"`
	Name                      types.String                                            `tfsdk:"name"`
	Mac                       types.String                                            `tfsdk:"mac"`
	Macs                      types.List                                              `tfsdk:"macs"`
	Serial                    types.String                                            `tfsdk:"serial"`
	Serials                   types.List                                              `tfsdk:"serials"`
	ConfigurationUpdatedAfter types.String                                            `tfsdk:"configuration_updated_after"`
	Items                     *[]ResponseItemSwitchGetOrganizationSwitchPortsBySwitch `tfsdk:"items"`
}

type ResponseItemSwitchGetOrganizationSwitchPortsBySwitch struct {
	Mac     types.String                                                 `tfsdk:"mac"`
	Model   types.String                                                 `tfsdk:"model"`
	Name    types.String                                                 `tfsdk:"name"`
	Network *ResponseItemSwitchGetOrganizationSwitchPortsBySwitchNetwork `tfsdk:"network"`
	Ports   *[]ResponseItemSwitchGetOrganizationSwitchPortsBySwitchPorts `tfsdk:"ports"`
	Serial  types.String                                                 `tfsdk:"serial"`
}

type ResponseItemSwitchGetOrganizationSwitchPortsBySwitchNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseItemSwitchGetOrganizationSwitchPortsBySwitchPorts struct {
	AccessPolicyType        types.String `tfsdk:"access_policy_type"`
	AllowedVLANs            types.String `tfsdk:"allowed_vlans"`
	Enabled                 types.Bool   `tfsdk:"enabled"`
	LinkNegotiation         types.String `tfsdk:"link_negotiation"`
	Name                    types.String `tfsdk:"name"`
	PoeEnabled              types.Bool   `tfsdk:"poe_enabled"`
	PortID                  types.String `tfsdk:"port_id"`
	RstpEnabled             types.Bool   `tfsdk:"rstp_enabled"`
	StickyMacAllowList      types.List   `tfsdk:"sticky_mac_allow_list"`
	StickyMacAllowListLimit types.Int64  `tfsdk:"sticky_mac_allow_list_limit"`
	StpGuard                types.String `tfsdk:"stp_guard"`
	Tags                    types.List   `tfsdk:"tags"`
	Type                    types.String `tfsdk:"type"`
	VLAN                    types.Int64  `tfsdk:"vlan"`
	VoiceVLAN               types.Int64  `tfsdk:"voice_vlan"`
}

// ToBody
func ResponseSwitchGetOrganizationSwitchPortsBySwitchItemsToBody(state OrganizationsSwitchPortsBySwitch, response *merakigosdk.ResponseSwitchGetOrganizationSwitchPortsBySwitch) OrganizationsSwitchPortsBySwitch {
	var items []ResponseItemSwitchGetOrganizationSwitchPortsBySwitch
	for _, item := range *response {
		itemState := ResponseItemSwitchGetOrganizationSwitchPortsBySwitch{
			Mac:   types.StringValue(item.Mac),
			Model: types.StringValue(item.Model),
			Name:  types.StringValue(item.Name),
			Network: func() *ResponseItemSwitchGetOrganizationSwitchPortsBySwitchNetwork {
				if item.Network != nil {
					return &ResponseItemSwitchGetOrganizationSwitchPortsBySwitchNetwork{
						ID:   types.StringValue(item.Network.ID),
						Name: types.StringValue(item.Network.Name),
					}
				}
				return &ResponseItemSwitchGetOrganizationSwitchPortsBySwitchNetwork{}
			}(),
			Ports: func() *[]ResponseItemSwitchGetOrganizationSwitchPortsBySwitchPorts {
				if item.Ports != nil {
					result := make([]ResponseItemSwitchGetOrganizationSwitchPortsBySwitchPorts, len(*item.Ports))
					for i, ports := range *item.Ports {
						result[i] = ResponseItemSwitchGetOrganizationSwitchPortsBySwitchPorts{
							AccessPolicyType: types.StringValue(ports.AccessPolicyType),
							AllowedVLANs:     types.StringValue(ports.AllowedVLANs),
							Enabled: func() types.Bool {
								if ports.Enabled != nil {
									return types.BoolValue(*ports.Enabled)
								}
								return types.Bool{}
							}(),
							LinkNegotiation: types.StringValue(ports.LinkNegotiation),
							Name:            types.StringValue(ports.Name),
							PoeEnabled: func() types.Bool {
								if ports.PoeEnabled != nil {
									return types.BoolValue(*ports.PoeEnabled)
								}
								return types.Bool{}
							}(),
							PortID: types.StringValue(ports.PortID),
							RstpEnabled: func() types.Bool {
								if ports.RstpEnabled != nil {
									return types.BoolValue(*ports.RstpEnabled)
								}
								return types.Bool{}
							}(),
							StickyMacAllowList: StringSliceToList(ports.StickyMacAllowList),
							StickyMacAllowListLimit: func() types.Int64 {
								if ports.StickyMacAllowListLimit != nil {
									return types.Int64Value(int64(*ports.StickyMacAllowListLimit))
								}
								return types.Int64{}
							}(),
							StpGuard: types.StringValue(ports.StpGuard),
							Tags:     StringSliceToList(ports.Tags),
							Type:     types.StringValue(ports.Type),
							VLAN: func() types.Int64 {
								if ports.VLAN != nil {
									return types.Int64Value(int64(*ports.VLAN))
								}
								return types.Int64{}
							}(),
							VoiceVLAN: func() types.Int64 {
								if ports.VoiceVLAN != nil {
									return types.Int64Value(int64(*ports.VoiceVLAN))
								}
								return types.Int64{}
							}(),
						}
					}
					return &result
				}
				return &[]ResponseItemSwitchGetOrganizationSwitchPortsBySwitchPorts{}
			}(),
			Serial: types.StringValue(item.Serial),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
