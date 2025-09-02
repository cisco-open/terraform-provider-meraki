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
	_ datasource.DataSource              = &OrganizationsWirelessControllerOverviewByDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessControllerOverviewByDeviceDataSource{}
)

func NewOrganizationsWirelessControllerOverviewByDeviceDataSource() datasource.DataSource {
	return &OrganizationsWirelessControllerOverviewByDeviceDataSource{}
}

type OrganizationsWirelessControllerOverviewByDeviceDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessControllerOverviewByDeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessControllerOverviewByDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_controller_overview_by_device"
}

func (d *OrganizationsWirelessControllerOverviewByDeviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter wireless LAN controllers by network ID. This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.`,
				Optional:            true,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches.`,
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
						MarkdownDescription: `Wireless LAN controller overview`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"counts": schema.SingleNestedAttribute{
									MarkdownDescription: `Wireless LAN controller client and access point counts`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"clients": schema.SingleNestedAttribute{
											MarkdownDescription: `Wireless LAN controller client counts`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"by_status": schema.SingleNestedAttribute{
													MarkdownDescription: `Client counts by their status`,
													Computed:            true,
													Attributes: map[string]schema.Attribute{

														"online": schema.Int64Attribute{
															MarkdownDescription: `Wireless LAN controller active client count`,
															Computed:            true,
														},
													},
												},
											},
										},
										"connections": schema.SingleNestedAttribute{
											MarkdownDescription: `Wireless LAN controller associated access point counts`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"by_status": schema.SingleNestedAttribute{
													MarkdownDescription: `Access point counts by their status`,
													Computed:            true,
													Attributes: map[string]schema.Attribute{

														"offline": schema.Int64Attribute{
															MarkdownDescription: `Wireless LAN controller associated offline access point count`,
															Computed:            true,
														},
														"online": schema.Int64Attribute{
															MarkdownDescription: `Wireless LAN controller associated online access point count`,
															Computed:            true,
														},
													},
												},
												"total": schema.Int64Attribute{
													MarkdownDescription: `Wireless LAN controller associated total access point count`,
													Computed:            true,
												},
											},
										},
									},
								},
								"firmware": schema.SingleNestedAttribute{
									MarkdownDescription: `Wireless LAN controller device firmware information`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"version": schema.SingleNestedAttribute{
											MarkdownDescription: `Wireless LAN controller firmware version`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"short_name": schema.StringAttribute{
													MarkdownDescription: `Wireless LAN controller firmware version short name`,
													Computed:            true,
												},
											},
										},
									},
								},
								"network": schema.SingleNestedAttribute{
									MarkdownDescription: `Wireless LAN controller network`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `Wireless LAN controller network ID`,
											Computed:            true,
										},
									},
								},
								"redundancy": schema.SingleNestedAttribute{
									MarkdownDescription: `Wireless LAN controller redundancy information`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"chassis_name": schema.StringAttribute{
											MarkdownDescription: `Wireless LAN controller chassis name`,
											Computed:            true,
										},
										"id": schema.StringAttribute{
											MarkdownDescription: `Wireless LAN controller redundancy ID`,
											Computed:            true,
										},
										"management": schema.SingleNestedAttribute{
											MarkdownDescription: `Wireless LAN controller redundancy management interface information`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"addresses": schema.SetNestedAttribute{
													MarkdownDescription: `Wireless LAN controller redundancy management interface addresses`,
													Computed:            true,
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{

															"address": schema.StringAttribute{
																MarkdownDescription: `Wireless LAN controller redundancy management interface ip address`,
																Computed:            true,
															},
														},
													},
												},
											},
										},
										"redundant_serial": schema.StringAttribute{
											MarkdownDescription: `Wireless LAN controller redundant device serial`,
											Computed:            true,
										},
										"role": schema.StringAttribute{
											MarkdownDescription: `Wireless LAN controller role(Active, Active recovery, Standby hot, Standby recovery and Offline)`,
											Computed:            true,
										},
									},
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `Wireless LAN controller cloud ID`,
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

func (d *OrganizationsWirelessControllerOverviewByDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessControllerOverviewByDevice OrganizationsWirelessControllerOverviewByDevice
	diags := req.Config.Get(ctx, &organizationsWirelessControllerOverviewByDevice)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessControllerOverviewByDevice")
		vvOrganizationID := organizationsWirelessControllerOverviewByDevice.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessControllerOverviewByDeviceQueryParams{}

		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsWirelessControllerOverviewByDevice.NetworkIDs)
		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessControllerOverviewByDevice.Serials)
		queryParams1.PerPage = int(organizationsWirelessControllerOverviewByDevice.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessControllerOverviewByDevice.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessControllerOverviewByDevice.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.WirelessController.GetOrganizationWirelessControllerOverviewByDevice(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessControllerOverviewByDevice",
				err.Error(),
			)
			return
		}

		organizationsWirelessControllerOverviewByDevice = ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemToBody(organizationsWirelessControllerOverviewByDevice, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessControllerOverviewByDevice)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessControllerOverviewByDevice struct {
	OrganizationID types.String                                                                 `tfsdk:"organization_id"`
	NetworkIDs     types.List                                                                   `tfsdk:"network_ids"`
	Serials        types.List                                                                   `tfsdk:"serials"`
	PerPage        types.Int64                                                                  `tfsdk:"per_page"`
	StartingAfter  types.String                                                                 `tfsdk:"starting_after"`
	EndingBefore   types.String                                                                 `tfsdk:"ending_before"`
	Item           *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDevice `tfsdk:"item"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDevice struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItems `tfsdk:"items"`
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceMeta    `tfsdk:"meta"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItems struct {
	Counts     *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCounts     `tfsdk:"counts"`
	Firmware   *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsFirmware   `tfsdk:"firmware"`
	Network    *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsNetwork    `tfsdk:"network"`
	Redundancy *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsRedundancy `tfsdk:"redundancy"`
	Serial     types.String                                                                                `tfsdk:"serial"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCounts struct {
	Clients     *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsClients     `tfsdk:"clients"`
	Connections *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsConnections `tfsdk:"connections"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsClients struct {
	ByStatus *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsClientsByStatus `tfsdk:"by_status"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsClientsByStatus struct {
	Online types.Int64 `tfsdk:"online"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsConnections struct {
	ByStatus *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsConnectionsByStatus `tfsdk:"by_status"`
	Total    types.Int64                                                                                                `tfsdk:"total"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsConnectionsByStatus struct {
	Offline types.Int64 `tfsdk:"offline"`
	Online  types.Int64 `tfsdk:"online"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsFirmware struct {
	Version *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsFirmwareVersion `tfsdk:"version"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsFirmwareVersion struct {
	ShortName types.String `tfsdk:"short_name"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsNetwork struct {
	ID types.String `tfsdk:"id"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsRedundancy struct {
	ChassisName     types.String                                                                                          `tfsdk:"chassis_name"`
	ID              types.String                                                                                          `tfsdk:"id"`
	Management      *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsRedundancyManagement `tfsdk:"management"`
	RedundantSerial types.String                                                                                          `tfsdk:"redundant_serial"`
	Role            types.String                                                                                          `tfsdk:"role"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsRedundancyManagement struct {
	Addresses *[]ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsRedundancyManagementAddresses `tfsdk:"addresses"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsRedundancyManagementAddresses struct {
	Address types.String `tfsdk:"address"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceMetaCounts `tfsdk:"counts"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceMetaCountsItems `tfsdk:"items"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemToBody(state OrganizationsWirelessControllerOverviewByDevice, response *merakigosdk.ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDevice) OrganizationsWirelessControllerOverviewByDevice {
	itemState := ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDevice{
		Items: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItems {
			if response.Items != nil {
				result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItems{
						Counts: func() *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCounts {
							if items.Counts != nil {
								return &ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCounts{
									Clients: func() *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsClients {
										if items.Counts.Clients != nil {
											return &ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsClients{
												ByStatus: func() *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsClientsByStatus {
													if items.Counts.Clients.ByStatus != nil {
														return &ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsClientsByStatus{
															Online: func() types.Int64 {
																if items.Counts.Clients.ByStatus.Online != nil {
																	return types.Int64Value(int64(*items.Counts.Clients.ByStatus.Online))
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
									Connections: func() *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsConnections {
										if items.Counts.Connections != nil {
											return &ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsConnections{
												ByStatus: func() *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsConnectionsByStatus {
													if items.Counts.Connections.ByStatus != nil {
														return &ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsCountsConnectionsByStatus{
															Offline: func() types.Int64 {
																if items.Counts.Connections.ByStatus.Offline != nil {
																	return types.Int64Value(int64(*items.Counts.Connections.ByStatus.Offline))
																}
																return types.Int64{}
															}(),
															Online: func() types.Int64 {
																if items.Counts.Connections.ByStatus.Online != nil {
																	return types.Int64Value(int64(*items.Counts.Connections.ByStatus.Online))
																}
																return types.Int64{}
															}(),
														}
													}
													return nil
												}(),
												Total: func() types.Int64 {
													if items.Counts.Connections.Total != nil {
														return types.Int64Value(int64(*items.Counts.Connections.Total))
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
						Firmware: func() *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsFirmware {
							if items.Firmware != nil {
								return &ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsFirmware{
									Version: func() *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsFirmwareVersion {
										if items.Firmware.Version != nil {
											return &ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsFirmwareVersion{
												ShortName: func() types.String {
													if items.Firmware.Version.ShortName != "" {
														return types.StringValue(items.Firmware.Version.ShortName)
													}
													return types.String{}
												}(),
											}
										}
										return nil
									}(),
								}
							}
							return nil
						}(),
						Network: func() *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsNetwork {
							if items.Network != nil {
								return &ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsNetwork{
									ID: func() types.String {
										if items.Network.ID != "" {
											return types.StringValue(items.Network.ID)
										}
										return types.String{}
									}(),
								}
							}
							return nil
						}(),
						Redundancy: func() *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsRedundancy {
							if items.Redundancy != nil {
								return &ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsRedundancy{
									ChassisName: func() types.String {
										if items.Redundancy.ChassisName != "" {
											return types.StringValue(items.Redundancy.ChassisName)
										}
										return types.String{}
									}(),
									ID: func() types.String {
										if items.Redundancy.ID != "" {
											return types.StringValue(items.Redundancy.ID)
										}
										return types.String{}
									}(),
									Management: func() *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsRedundancyManagement {
										if items.Redundancy.Management != nil {
											return &ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsRedundancyManagement{
												Addresses: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsRedundancyManagementAddresses {
													if items.Redundancy.Management.Addresses != nil {
														result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsRedundancyManagementAddresses, len(*items.Redundancy.Management.Addresses))
														for i, addresses := range *items.Redundancy.Management.Addresses {
															result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceItemsRedundancyManagementAddresses{
																Address: func() types.String {
																	if addresses.Address != "" {
																		return types.StringValue(addresses.Address)
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
										}
										return nil
									}(),
									RedundantSerial: func() types.String {
										if items.Redundancy.RedundantSerial != "" {
											return types.StringValue(items.Redundancy.RedundantSerial)
										}
										return types.String{}
									}(),
									Role: func() types.String {
										if items.Redundancy.Role != "" {
											return types.StringValue(items.Redundancy.Role)
										}
										return types.String{}
									}(),
								}
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
		Meta: func() *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceMeta {
			if response.Meta != nil {
				return &ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceMeta{
					Counts: func() *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceMetaCounts{
								Items: func() *ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseWirelessControllerGetOrganizationWirelessControllerOverviewByDeviceMetaCountsItems{
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
