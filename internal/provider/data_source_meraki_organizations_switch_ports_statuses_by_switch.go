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
	_ datasource.DataSource              = &OrganizationsSwitchPortsStatusesBySwitchDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSwitchPortsStatusesBySwitchDataSource{}
)

func NewOrganizationsSwitchPortsStatusesBySwitchDataSource() datasource.DataSource {
	return &OrganizationsSwitchPortsStatusesBySwitchDataSource{}
}

type OrganizationsSwitchPortsStatusesBySwitchDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSwitchPortsStatusesBySwitchDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSwitchPortsStatusesBySwitchDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_switch_ports_statuses_by_switch"
}

func (d *OrganizationsSwitchPortsStatusesBySwitchDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"configuration_updated_after": schema.StringAttribute{
				MarkdownDescription: `configurationUpdatedAfter query parameter. Optional parameter to filter items to switches where the configuration has been updated after the given timestamp.`,
				Optional:            true,
			},
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"mac": schema.StringAttribute{
				MarkdownDescription: `mac query parameter. Optional parameter to filter items to switches with MAC addresses that contain the search term or are an exact match.`,
				Optional:            true,
			},
			"macs": schema.ListAttribute{
				MarkdownDescription: `macs query parameter. Optional parameter to filter items to switches that have one of the provided MAC addresses.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `name query parameter. Optional parameter to filter items to switches with names that contain the search term or are an exact match.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter items to switches in one of the provided networks.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 20. Default is 10.`,
				Optional:            true,
			},
			"port_profile_ids": schema.ListAttribute{
				MarkdownDescription: `portProfileIds query parameter. Optional parameter to filter items to switches that contain switchports belonging to one of the specified port profiles.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial query parameter. Optional parameter to filter items to switches with serial number that contains the search term or are an exact match.`,
				Optional:            true,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Optional parameter to filter items to switches that have one of the provided serials.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"items": schema.ListNestedAttribute{
						MarkdownDescription: `Switches`,
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
									MarkdownDescription: `The statuses of the ports on the switch.`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"duplex": schema.StringAttribute{
												MarkdownDescription: `The current duplex of a connected port.`,
												Computed:            true,
											},
											"enabled": schema.BoolAttribute{
												MarkdownDescription: `Whether the port is configured to be enabled.`,
												Computed:            true,
											},
											"errors": schema.ListAttribute{
												MarkdownDescription: `All errors present on the port.`,
												Computed:            true,
												ElementType:         types.StringType,
											},
											"is_uplink": schema.BoolAttribute{
												MarkdownDescription: `Whether the port is the switch's uplink.`,
												Computed:            true,
											},
											"poe": schema.SingleNestedAttribute{
												MarkdownDescription: `PoE status of the port.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"is_allocated": schema.BoolAttribute{
														MarkdownDescription: `Whether the port is drawing power`,
														Computed:            true,
													},
												},
											},
											"port_id": schema.StringAttribute{
												MarkdownDescription: `The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module.`,
												Computed:            true,
											},
											"secure_port": schema.SingleNestedAttribute{
												MarkdownDescription: `The Secure Port status of the port.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"active": schema.BoolAttribute{
														MarkdownDescription: `Whether Secure Port is currently active for this port.`,
														Computed:            true,
													},
													"authentication_status": schema.StringAttribute{
														MarkdownDescription: `The current Secure Port status.`,
														Computed:            true,
													},
												},
											},
											"spanning_tree": schema.SingleNestedAttribute{
												MarkdownDescription: `The Spanning Tree Protocol (STP) information of the connected device.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"statuses": schema.ListAttribute{
														MarkdownDescription: `The current Spanning Tree Protocol statuses of the port.`,
														Computed:            true,
														ElementType:         types.StringType,
													},
												},
											},
											"speed": schema.StringAttribute{
												MarkdownDescription: `The current data transfer rate which the port is operating at.`,
												Computed:            true,
											},
											"status": schema.StringAttribute{
												MarkdownDescription: `The current connection status of the port.`,
												Computed:            true,
											},
											"warnings": schema.ListAttribute{
												MarkdownDescription: `All warnings present on the port.`,
												Computed:            true,
												ElementType:         types.StringType,
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
					"meta": schema.SingleNestedAttribute{
						MarkdownDescription: `Metadata relevant to the paginated dataset`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"counts": schema.SingleNestedAttribute{
								MarkdownDescription: `Counts relating to the paginated dataset`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"items": schema.SingleNestedAttribute{
										MarkdownDescription: `Counts relating to the paginated items`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"remaining": schema.Int64Attribute{
												MarkdownDescription: `The number of items in the dataset that are available on subsequent pages`,
												Computed:            true,
											},
											"total": schema.Int64Attribute{
												MarkdownDescription: `The total number of items in the dataset`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsSwitchPortsStatusesBySwitchDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSwitchPortsStatusesBySwitch OrganizationsSwitchPortsStatusesBySwitch
	diags := req.Config.Get(ctx, &organizationsSwitchPortsStatusesBySwitch)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSwitchPortsStatusesBySwitch")
		vvOrganizationID := organizationsSwitchPortsStatusesBySwitch.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationSwitchPortsStatusesBySwitchQueryParams{}

		queryParams1.PerPage = int(organizationsSwitchPortsStatusesBySwitch.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsSwitchPortsStatusesBySwitch.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsSwitchPortsStatusesBySwitch.EndingBefore.ValueString()
		queryParams1.ConfigurationUpdatedAfter = organizationsSwitchPortsStatusesBySwitch.ConfigurationUpdatedAfter.ValueString()
		queryParams1.Mac = organizationsSwitchPortsStatusesBySwitch.Mac.ValueString()
		queryParams1.Macs = elementsToStrings(ctx, organizationsSwitchPortsStatusesBySwitch.Macs)
		queryParams1.Name = organizationsSwitchPortsStatusesBySwitch.Name.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsSwitchPortsStatusesBySwitch.NetworkIDs)
		queryParams1.PortProfileIDs = elementsToStrings(ctx, organizationsSwitchPortsStatusesBySwitch.PortProfileIDs)
		queryParams1.Serial = organizationsSwitchPortsStatusesBySwitch.Serial.ValueString()
		queryParams1.Serials = elementsToStrings(ctx, organizationsSwitchPortsStatusesBySwitch.Serials)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetOrganizationSwitchPortsStatusesBySwitch(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSwitchPortsStatusesBySwitch",
				err.Error(),
			)
			return
		}

		organizationsSwitchPortsStatusesBySwitch = ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemToBody(organizationsSwitchPortsStatusesBySwitch, response1)
		diags = resp.State.Set(ctx, &organizationsSwitchPortsStatusesBySwitch)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSwitchPortsStatusesBySwitch struct {
	OrganizationID            types.String                                              `tfsdk:"organization_id"`
	PerPage                   types.Int64                                               `tfsdk:"per_page"`
	StartingAfter             types.String                                              `tfsdk:"starting_after"`
	EndingBefore              types.String                                              `tfsdk:"ending_before"`
	ConfigurationUpdatedAfter types.String                                              `tfsdk:"configuration_updated_after"`
	Mac                       types.String                                              `tfsdk:"mac"`
	Macs                      types.List                                                `tfsdk:"macs"`
	Name                      types.String                                              `tfsdk:"name"`
	NetworkIDs                types.List                                                `tfsdk:"network_ids"`
	PortProfileIDs            types.List                                                `tfsdk:"port_profile_ids"`
	Serial                    types.String                                              `tfsdk:"serial"`
	Serials                   types.List                                                `tfsdk:"serials"`
	Item                      *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitch `tfsdk:"item"`
}

type ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitch struct {
	Items *[]ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItems `tfsdk:"items"`
	Meta  *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchMeta    `tfsdk:"meta"`
}

type ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItems struct {
	Mac     types.String                                                          `tfsdk:"mac"`
	Model   types.String                                                          `tfsdk:"model"`
	Name    types.String                                                          `tfsdk:"name"`
	Network *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsNetwork `tfsdk:"network"`
	Ports   *[]ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPorts `tfsdk:"ports"`
	Serial  types.String                                                          `tfsdk:"serial"`
}

type ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPorts struct {
	Duplex       types.String                                                                    `tfsdk:"duplex"`
	Enabled      types.Bool                                                                      `tfsdk:"enabled"`
	Errors       types.List                                                                      `tfsdk:"errors"`
	IsUplink     types.Bool                                                                      `tfsdk:"is_uplink"`
	Poe          *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPortsPoe          `tfsdk:"poe"`
	PortID       types.String                                                                    `tfsdk:"port_id"`
	SecurePort   *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPortsSecurePort   `tfsdk:"secure_port"`
	SpanningTree *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPortsSpanningTree `tfsdk:"spanning_tree"`
	Speed        types.String                                                                    `tfsdk:"speed"`
	Status       types.String                                                                    `tfsdk:"status"`
	Warnings     types.List                                                                      `tfsdk:"warnings"`
}

type ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPortsPoe struct {
	IsAllocated types.Bool `tfsdk:"is_allocated"`
}

type ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPortsSecurePort struct {
	Active               types.Bool   `tfsdk:"active"`
	AuthenticationStatus types.String `tfsdk:"authentication_status"`
}

type ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPortsSpanningTree struct {
	Statuses types.List `tfsdk:"statuses"`
}

type ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchMeta struct {
	Counts *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchMetaCounts `tfsdk:"counts"`
}

type ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchMetaCounts struct {
	Items *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchMetaCountsItems `tfsdk:"items"`
}

type ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemToBody(state OrganizationsSwitchPortsStatusesBySwitch, response *merakigosdk.ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitch) OrganizationsSwitchPortsStatusesBySwitch {
	itemState := ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitch{
		Items: func() *[]ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItems {
			if response.Items != nil {
				result := make([]ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItems{
						Mac: func() types.String {
							if items.Mac != "" {
								return types.StringValue(items.Mac)
							}
							return types.String{}
						}(),
						Model: func() types.String {
							if items.Model != "" {
								return types.StringValue(items.Model)
							}
							return types.String{}
						}(),
						Name: func() types.String {
							if items.Name != "" {
								return types.StringValue(items.Name)
							}
							return types.String{}
						}(),
						Network: func() *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsNetwork {
							if items.Network != nil {
								return &ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsNetwork{
									ID: func() types.String {
										if items.Network.ID != "" {
											return types.StringValue(items.Network.ID)
										}
										return types.String{}
									}(),
									Name: func() types.String {
										if items.Network.Name != "" {
											return types.StringValue(items.Network.Name)
										}
										return types.String{}
									}(),
								}
							}
							return nil
						}(),
						Ports: func() *[]ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPorts {
							if items.Ports != nil {
								result := make([]ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPorts, len(*items.Ports))
								for i, ports := range *items.Ports {
									result[i] = ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPorts{
										Duplex: func() types.String {
											if ports.Duplex != "" {
												return types.StringValue(ports.Duplex)
											}
											return types.String{}
										}(),
										Enabled: func() types.Bool {
											if ports.Enabled != nil {
												return types.BoolValue(*ports.Enabled)
											}
											return types.Bool{}
										}(),
										Errors: StringSliceToList(ports.Errors),
										IsUplink: func() types.Bool {
											if ports.IsUplink != nil {
												return types.BoolValue(*ports.IsUplink)
											}
											return types.Bool{}
										}(),
										Poe: func() *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPortsPoe {
											if ports.Poe != nil {
												return &ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPortsPoe{
													IsAllocated: func() types.Bool {
														if ports.Poe.IsAllocated != nil {
															return types.BoolValue(*ports.Poe.IsAllocated)
														}
														return types.Bool{}
													}(),
												}
											}
											return nil
										}(),
										PortID: func() types.String {
											if ports.PortID != "" {
												return types.StringValue(ports.PortID)
											}
											return types.String{}
										}(),
										SecurePort: func() *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPortsSecurePort {
											if ports.SecurePort != nil {
												return &ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPortsSecurePort{
													Active: func() types.Bool {
														if ports.SecurePort.Active != nil {
															return types.BoolValue(*ports.SecurePort.Active)
														}
														return types.Bool{}
													}(),
													AuthenticationStatus: func() types.String {
														if ports.SecurePort.AuthenticationStatus != "" {
															return types.StringValue(ports.SecurePort.AuthenticationStatus)
														}
														return types.String{}
													}(),
												}
											}
											return nil
										}(),
										SpanningTree: func() *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPortsSpanningTree {
											if ports.SpanningTree != nil {
												return &ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchItemsPortsSpanningTree{
													Statuses: StringSliceToList(ports.SpanningTree.Statuses),
												}
											}
											return nil
										}(),
										Speed: func() types.String {
											if ports.Speed != "" {
												return types.StringValue(ports.Speed)
											}
											return types.String{}
										}(),
										Status: func() types.String {
											if ports.Status != "" {
												return types.StringValue(ports.Status)
											}
											return types.String{}
										}(),
										Warnings: StringSliceToList(ports.Warnings),
									}
								}
								return &result
							}
							return nil
						}(),
						Serial: func() types.String {
							if items.Serial != "" {
								return types.StringValue(items.Serial)
							}
							return types.String{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		Meta: func() *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchMeta {
			if response.Meta != nil {
				return &ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchMeta{
					Counts: func() *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchMetaCounts{
								Items: func() *ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseSwitchGetOrganizationSwitchPortsStatusesBySwitchMetaCountsItems{
											Remaining: func() types.Int64 {
												if response.Meta.Counts.Items.Remaining != nil {
													return types.Int64Value(int64(*response.Meta.Counts.Items.Remaining))
												}
												return types.Int64{}
											}(),
											Total: func() types.Int64 {
												if response.Meta.Counts.Items.Total != nil {
													return types.Int64Value(int64(*response.Meta.Counts.Items.Total))
												}
												return types.Int64{}
											}(),
										}
									}
									return nil
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
