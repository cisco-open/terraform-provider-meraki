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
	_ datasource.DataSource              = &OrganizationsClientsSearchDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsClientsSearchDataSource{}
)

func NewOrganizationsClientsSearchDataSource() datasource.DataSource {
	return &OrganizationsClientsSearchDataSource{}
}

type OrganizationsClientsSearchDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsClientsSearchDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsClientsSearchDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_clients_search"
}

func (d *OrganizationsClientsSearchDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"mac": schema.StringAttribute{
				MarkdownDescription: `mac query parameter. The MAC address of the client. Required.`,
				Required:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 5. Default is 5.`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"client_id": schema.StringAttribute{
						MarkdownDescription: `The ID of the client`,
						Computed:            true,
					},
					"mac": schema.StringAttribute{
						MarkdownDescription: `The MAC address of the client`,
						Computed:            true,
					},
					"manufacturer": schema.StringAttribute{
						MarkdownDescription: `Manufacturer of the client`,
						Computed:            true,
					},
					"records": schema.SetNestedAttribute{
						MarkdownDescription: `The clients that appear on any networks within an organization`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"cdp": schema.SetNestedAttribute{
									MarkdownDescription: `The Cisco discover protocol settings for the client`,
									Computed:            true,
								},
								"client_vpn_connections": schema.SetNestedAttribute{
									MarkdownDescription: `VPN connections associated with the client`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"connected_at": schema.Int64Attribute{
												MarkdownDescription: `The time the client last connected to the VPN`,
												Computed:            true,
											},
											"disconnected_at": schema.Int64Attribute{
												MarkdownDescription: `The time the client last disconnected from the VPN`,
												Computed:            true,
											},
											"remote_ip": schema.StringAttribute{
												MarkdownDescription: `The IP address of the VPN the client last connected to`,
												Computed:            true,
											},
										},
									},
								},
								"description": schema.StringAttribute{
									MarkdownDescription: `Short description of the client`,
									Computed:            true,
								},
								"first_seen": schema.Int64Attribute{
									MarkdownDescription: `Timestamp client was first seen in the network`,
									Computed:            true,
								},
								"ip": schema.StringAttribute{
									MarkdownDescription: `The IP address of the client`,
									Computed:            true,
								},
								"ip6": schema.StringAttribute{
									MarkdownDescription: `The IPv6 address of the client`,
									Computed:            true,
								},
								"last_seen": schema.Int64Attribute{
									MarkdownDescription: `Timestamp client was last seen in the network`,
									Computed:            true,
								},
								"lldp": schema.SetNestedAttribute{
									MarkdownDescription: `The link layer discover protocol settings for the client`,
									Computed:            true,
								},
								"network": schema.SingleNestedAttribute{
									MarkdownDescription: `The network upon which a client with the given MAC address was found`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"enrollment_string": schema.StringAttribute{
											MarkdownDescription: `The network enrollment string`,
											Computed:            true,
										},
										"id": schema.StringAttribute{
											MarkdownDescription: `The network identifier`,
											Computed:            true,
										},
										"is_bound_to_config_template": schema.BoolAttribute{
											MarkdownDescription: `If the network is bound to a config template`,
											Computed:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: `The network name`,
											Computed:            true,
										},
										"notes": schema.StringAttribute{
											MarkdownDescription: `The notes for the network`,
											Computed:            true,
										},
										"organization_id": schema.StringAttribute{
											MarkdownDescription: `The organization identifier`,
											Computed:            true,
										},
										"product_types": schema.ListAttribute{
											Computed:    true,
											ElementType: types.StringType,
										},
										"tags": schema.ListAttribute{
											MarkdownDescription: `The network tags`,
											Computed:            true,
											ElementType:         types.StringType,
										},
										"time_zone": schema.StringAttribute{
											MarkdownDescription: `The network's timezone`,
											Computed:            true,
										},
										"url": schema.StringAttribute{
											MarkdownDescription: `The network URL`,
											Computed:            true,
										},
									},
								},
								"os": schema.StringAttribute{
									MarkdownDescription: `The operating system of the client`,
									Computed:            true,
								},
								"recent_device_mac": schema.StringAttribute{
									MarkdownDescription: `The MAC address of the node that the device was last connected to`,
									Computed:            true,
								},
								"sm_installed": schema.BoolAttribute{
									MarkdownDescription: `Status of SM for the client`,
									Computed:            true,
								},
								"ssid": schema.StringAttribute{
									MarkdownDescription: `The name of the SSID that the client is connected to`,
									Computed:            true,
								},
								"status": schema.StringAttribute{
									MarkdownDescription: `The connection status of the client`,
									Computed:            true,
								},
								"switchport": schema.StringAttribute{
									MarkdownDescription: `The switch port the client is connected to`,
									Computed:            true,
								},
								"user": schema.StringAttribute{
									MarkdownDescription: `The username of the user of the client`,
									Computed:            true,
								},
								"vlan": schema.StringAttribute{
									MarkdownDescription: `The name of the VLAN that the client is connected to`,
									Computed:            true,
								},
								"wireless_capabilities": schema.StringAttribute{
									MarkdownDescription: `Wireless capabilities of the client`,
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

func (d *OrganizationsClientsSearchDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsClientsSearch OrganizationsClientsSearch
	diags := req.Config.Get(ctx, &organizationsClientsSearch)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationClientsSearch")
		vvOrganizationID := organizationsClientsSearch.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationClientsSearchQueryParams{}

		queryParams1.PerPage = int(organizationsClientsSearch.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsClientsSearch.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsClientsSearch.EndingBefore.ValueString()
		queryParams1.Mac = organizationsClientsSearch.Mac.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationClientsSearch(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationClientsSearch",
				err.Error(),
			)
			return
		}

		organizationsClientsSearch = ResponseOrganizationsGetOrganizationClientsSearchItemToBody(organizationsClientsSearch, response1)
		diags = resp.State.Set(ctx, &organizationsClientsSearch)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsClientsSearch struct {
	OrganizationID types.String                                       `tfsdk:"organization_id"`
	PerPage        types.Int64                                        `tfsdk:"per_page"`
	StartingAfter  types.String                                       `tfsdk:"starting_after"`
	EndingBefore   types.String                                       `tfsdk:"ending_before"`
	Mac            types.String                                       `tfsdk:"mac"`
	Item           *ResponseOrganizationsGetOrganizationClientsSearch `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationClientsSearch struct {
	ClientID     types.String                                                `tfsdk:"client_id"`
	Mac          types.String                                                `tfsdk:"mac"`
	Manufacturer types.String                                                `tfsdk:"manufacturer"`
	Records      *[]ResponseOrganizationsGetOrganizationClientsSearchRecords `tfsdk:"records"`
}

type ResponseOrganizationsGetOrganizationClientsSearchRecords struct {
	Cdp                  *[][]string                                                                     `tfsdk:"cdp"`
	ClientVpnConnections *[]ResponseOrganizationsGetOrganizationClientsSearchRecordsClientVpnConnections `tfsdk:"client_vpn_connections"`
	Description          types.String                                                                    `tfsdk:"description"`
	FirstSeen            types.Int64                                                                     `tfsdk:"first_seen"`
	IP                   types.String                                                                    `tfsdk:"ip"`
	IP6                  types.String                                                                    `tfsdk:"ip6"`
	LastSeen             types.Int64                                                                     `tfsdk:"last_seen"`
	Lldp                 *[][]string                                                                     `tfsdk:"lldp"`
	Network              *ResponseOrganizationsGetOrganizationClientsSearchRecordsNetwork                `tfsdk:"network"`
	Os                   types.String                                                                    `tfsdk:"os"`
	RecentDeviceMac      types.String                                                                    `tfsdk:"recent_device_mac"`
	SmInstalled          types.Bool                                                                      `tfsdk:"sm_installed"`
	SSID                 types.String                                                                    `tfsdk:"ssid"`
	Status               types.String                                                                    `tfsdk:"status"`
	Switchport           types.String                                                                    `tfsdk:"switchport"`
	User                 types.String                                                                    `tfsdk:"user"`
	VLAN                 types.String                                                                    `tfsdk:"vlan"`
	WirelessCapabilities types.String                                                                    `tfsdk:"wireless_capabilities"`
}

type ResponseOrganizationsGetOrganizationClientsSearchRecordsClientVpnConnections struct {
	ConnectedAt    types.Int64  `tfsdk:"connected_at"`
	DisconnectedAt types.Int64  `tfsdk:"disconnected_at"`
	RemoteIP       types.String `tfsdk:"remote_ip"`
}

type ResponseOrganizationsGetOrganizationClientsSearchRecordsNetwork struct {
	EnrollmentString        types.String `tfsdk:"enrollment_string"`
	ID                      types.String `tfsdk:"id"`
	IsBoundToConfigTemplate types.Bool   `tfsdk:"is_bound_to_config_template"`
	Name                    types.String `tfsdk:"name"`
	Notes                   types.String `tfsdk:"notes"`
	OrganizationID          types.String `tfsdk:"organization_id"`
	ProductTypes            types.List   `tfsdk:"product_types"`
	Tags                    types.List   `tfsdk:"tags"`
	TimeZone                types.String `tfsdk:"time_zone"`
	URL                     types.String `tfsdk:"url"`
}

// ToBody
func ResponseOrganizationsGetOrganizationClientsSearchItemToBody(state OrganizationsClientsSearch, response *merakigosdk.ResponseOrganizationsGetOrganizationClientsSearch) OrganizationsClientsSearch {
	itemState := ResponseOrganizationsGetOrganizationClientsSearch{
		ClientID: func() types.String {
			if response.ClientID != "" {
				return types.StringValue(response.ClientID)
			}
			return types.String{}
		}(),
		Mac: func() types.String {
			if response.Mac != "" {
				return types.StringValue(response.Mac)
			}
			return types.String{}
		}(),
		Manufacturer: func() types.String {
			if response.Manufacturer != "" {
				return types.StringValue(response.Manufacturer)
			}
			return types.String{}
		}(),
		Records: func() *[]ResponseOrganizationsGetOrganizationClientsSearchRecords {
			if response.Records != nil {
				result := make([]ResponseOrganizationsGetOrganizationClientsSearchRecords, len(*response.Records))
				for i, records := range *response.Records {
					result[i] = ResponseOrganizationsGetOrganizationClientsSearchRecords{
						//TODO [][]
						ClientVpnConnections: func() *[]ResponseOrganizationsGetOrganizationClientsSearchRecordsClientVpnConnections {
							if records.ClientVpnConnections != nil {
								result := make([]ResponseOrganizationsGetOrganizationClientsSearchRecordsClientVpnConnections, len(*records.ClientVpnConnections))
								for i, clientVpnConnections := range *records.ClientVpnConnections {
									result[i] = ResponseOrganizationsGetOrganizationClientsSearchRecordsClientVpnConnections{
										ConnectedAt: func() types.Int64 {
											if clientVpnConnections.ConnectedAt != nil {
												return types.Int64Value(int64(*clientVpnConnections.ConnectedAt))
											}
											return types.Int64{}
										}(),
										DisconnectedAt: func() types.Int64 {
											if clientVpnConnections.DisconnectedAt != nil {
												return types.Int64Value(int64(*clientVpnConnections.DisconnectedAt))
											}
											return types.Int64{}
										}(),
										RemoteIP: func() types.String {
											if clientVpnConnections.RemoteIP != "" {
												return types.StringValue(clientVpnConnections.RemoteIP)
											}
											return types.String{}
										}(),
									}
								}
								return &result
							}
							return nil
						}(),
						Description: func() types.String {
							if records.Description != "" {
								return types.StringValue(records.Description)
							}
							return types.String{}
						}(),
						FirstSeen: func() types.Int64 {
							if records.FirstSeen != nil {
								return types.Int64Value(int64(*records.FirstSeen))
							}
							return types.Int64{}
						}(),
						IP: func() types.String {
							if records.IP != "" {
								return types.StringValue(records.IP)
							}
							return types.String{}
						}(),
						IP6: func() types.String {
							if records.IP6 != "" {
								return types.StringValue(records.IP6)
							}
							return types.String{}
						}(),
						LastSeen: func() types.Int64 {
							if records.LastSeen != nil {
								return types.Int64Value(int64(*records.LastSeen))
							}
							return types.Int64{}
						}(),
						//TODO [][]
						Network: func() *ResponseOrganizationsGetOrganizationClientsSearchRecordsNetwork {
							if records.Network != nil {
								return &ResponseOrganizationsGetOrganizationClientsSearchRecordsNetwork{
									EnrollmentString: func() types.String {
										if records.Network.EnrollmentString != "" {
											return types.StringValue(records.Network.EnrollmentString)
										}
										return types.String{}
									}(),
									ID: func() types.String {
										if records.Network.ID != "" {
											return types.StringValue(records.Network.ID)
										}
										return types.String{}
									}(),
									IsBoundToConfigTemplate: func() types.Bool {
										if records.Network.IsBoundToConfigTemplate != nil {
											return types.BoolValue(*records.Network.IsBoundToConfigTemplate)
										}
										return types.Bool{}
									}(),
									Name: func() types.String {
										if records.Network.Name != "" {
											return types.StringValue(records.Network.Name)
										}
										return types.String{}
									}(),
									Notes: func() types.String {
										if records.Network.Notes != "" {
											return types.StringValue(records.Network.Notes)
										}
										return types.String{}
									}(),
									OrganizationID: func() types.String {
										if records.Network.OrganizationID != "" {
											return types.StringValue(records.Network.OrganizationID)
										}
										return types.String{}
									}(),
									ProductTypes: StringSliceToList(records.Network.ProductTypes),
									Tags:         StringSliceToList(records.Network.Tags),
									TimeZone: func() types.String {
										if records.Network.TimeZone != "" {
											return types.StringValue(records.Network.TimeZone)
										}
										return types.String{}
									}(),
									URL: func() types.String {
										if records.Network.URL != "" {
											return types.StringValue(records.Network.URL)
										}
										return types.String{}
									}(),
								}
							}
							return nil
						}(),
						Os: func() types.String {
							if records.Os != "" {
								return types.StringValue(records.Os)
							}
							return types.String{}
						}(),
						RecentDeviceMac: func() types.String {
							if records.RecentDeviceMac != "" {
								return types.StringValue(records.RecentDeviceMac)
							}
							return types.String{}
						}(),
						SmInstalled: func() types.Bool {
							if records.SmInstalled != nil {
								return types.BoolValue(*records.SmInstalled)
							}
							return types.Bool{}
						}(),
						SSID: func() types.String {
							if records.SSID != "" {
								return types.StringValue(records.SSID)
							}
							return types.String{}
						}(),
						Status: func() types.String {
							if records.Status != "" {
								return types.StringValue(records.Status)
							}
							return types.String{}
						}(),
						Switchport: func() types.String {
							if records.Switchport != "" {
								return types.StringValue(records.Switchport)
							}
							return types.String{}
						}(),
						User: func() types.String {
							if records.User != "" {
								return types.StringValue(records.User)
							}
							return types.String{}
						}(),
						VLAN: func() types.String {
							if records.VLAN != "" {
								return types.StringValue(records.VLAN)
							}
							return types.String{}
						}(),
						WirelessCapabilities: func() types.String {
							if records.WirelessCapabilities != "" {
								return types.StringValue(records.WirelessCapabilities)
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
