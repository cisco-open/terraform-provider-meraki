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
	_ datasource.DataSource              = &OrganizationsWirelessControllerDevicesInterfacesL2ByDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessControllerDevicesInterfacesL2ByDeviceDataSource{}
)

func NewOrganizationsWirelessControllerDevicesInterfacesL2ByDeviceDataSource() datasource.DataSource {
	return &OrganizationsWirelessControllerDevicesInterfacesL2ByDeviceDataSource{}
}

type OrganizationsWirelessControllerDevicesInterfacesL2ByDeviceDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessControllerDevicesInterfacesL2ByDeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessControllerDevicesInterfacesL2ByDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_controller_devices_interfaces_l2_by_device"
}

func (d *OrganizationsWirelessControllerDevicesInterfacesL2ByDeviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
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
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 31 days from today.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 31 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"items": schema.ListNestedAttribute{
						MarkdownDescription: `Wireless LAN controller L2 interfaces`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"interfaces": schema.SetNestedAttribute{
									MarkdownDescription: `Layer 2 interfaces belongs to the wireless LAN controller`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"channel_group": schema.SingleNestedAttribute{
												MarkdownDescription: `The channel group of this wireless LAN controller interface`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"number": schema.Int64Attribute{
														MarkdownDescription: `The interface channel group number`,
														Computed:            true,
													},
												},
											},
											"description": schema.StringAttribute{
												MarkdownDescription: `The description of the wireless LAN controller interface`,
												Computed:            true,
											},
											"enabled": schema.BoolAttribute{
												MarkdownDescription: `The status of the wireless LAN controller interface`,
												Computed:            true,
											},
											"is_redundancy_port": schema.BoolAttribute{
												MarkdownDescription: `Indicate whether the interface is a redundancy port used to perform HA role negotiation`,
												Computed:            true,
											},
											"is_uplink": schema.BoolAttribute{
												MarkdownDescription: `Indicate whether the interface is uplink`,
												Computed:            true,
											},
											"link_negotiation": schema.StringAttribute{
												MarkdownDescription: `The interface negotiation mode`,
												Computed:            true,
											},
											"mac": schema.StringAttribute{
												MarkdownDescription: `The MAC address of the wireless LAN controller interface`,
												Computed:            true,
											},
											"module": schema.SingleNestedAttribute{
												MarkdownDescription: `The module of this wireless LAN controller interface`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"model": schema.StringAttribute{
														MarkdownDescription: `The module type of this wireless LAN controller interface`,
														Computed:            true,
													},
												},
											},
											"name": schema.StringAttribute{
												MarkdownDescription: `The name of the wireless LAN controller interface`,
												Computed:            true,
											},
											"speed": schema.StringAttribute{
												MarkdownDescription: `The current data transfer rate which the interface is operating at. enum = [1 Gbps, 2 Gbps, 5 Gbps, 10 Gbps, 20 Gbps, 40 Gbps, 100 Gbps]`,
												Computed:            true,
											},
											"status": schema.StringAttribute{
												MarkdownDescription: `The status of the wireless LAN controller interface`,
												Computed:            true,
											},
											"vlan": schema.Int64Attribute{
												MarkdownDescription: `The VLAN of the switch port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.`,
												Computed:            true,
											},
										},
									},
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `The cloud ID of the wireless LAN controller`,
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

func (d *OrganizationsWirelessControllerDevicesInterfacesL2ByDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessControllerDevicesInterfacesL2ByDevice OrganizationsWirelessControllerDevicesInterfacesL2ByDevice
	diags := req.Config.Get(ctx, &organizationsWirelessControllerDevicesInterfacesL2ByDevice)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice")
		vvOrganizationID := organizationsWirelessControllerDevicesInterfacesL2ByDevice.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceQueryParams{}

		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessControllerDevicesInterfacesL2ByDevice.Serials)
		queryParams1.T0 = organizationsWirelessControllerDevicesInterfacesL2ByDevice.T0.ValueString()
		queryParams1.T1 = organizationsWirelessControllerDevicesInterfacesL2ByDevice.T1.ValueString()
		queryParams1.Timespan = organizationsWirelessControllerDevicesInterfacesL2ByDevice.Timespan.ValueFloat64()
		queryParams1.PerPage = int(organizationsWirelessControllerDevicesInterfacesL2ByDevice.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessControllerDevicesInterfacesL2ByDevice.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessControllerDevicesInterfacesL2ByDevice.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.WirelessController.GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice",
				err.Error(),
			)
			return
		}

		organizationsWirelessControllerDevicesInterfacesL2ByDevice = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemToBody(organizationsWirelessControllerDevicesInterfacesL2ByDevice, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessControllerDevicesInterfacesL2ByDevice)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessControllerDevicesInterfacesL2ByDevice struct {
	OrganizationID types.String                                                                            `tfsdk:"organization_id"`
	Serials        types.List                                                                              `tfsdk:"serials"`
	T0             types.String                                                                            `tfsdk:"t0"`
	T1             types.String                                                                            `tfsdk:"t1"`
	Timespan       types.Float64                                                                           `tfsdk:"timespan"`
	PerPage        types.Int64                                                                             `tfsdk:"per_page"`
	StartingAfter  types.String                                                                            `tfsdk:"starting_after"`
	EndingBefore   types.String                                                                            `tfsdk:"ending_before"`
	Item           *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDevice `tfsdk:"item"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDevice struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItems `tfsdk:"items"`
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMeta    `tfsdk:"meta"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItems struct {
	Interfaces *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfaces `tfsdk:"interfaces"`
	Serial     types.String                                                                                             `tfsdk:"serial"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfaces struct {
	ChannelGroup     *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfacesChannelGroup `tfsdk:"channel_group"`
	Description      types.String                                                                                                       `tfsdk:"description"`
	Enabled          types.Bool                                                                                                         `tfsdk:"enabled"`
	IsRedundancyPort types.Bool                                                                                                         `tfsdk:"is_redundancy_port"`
	IsUplink         types.Bool                                                                                                         `tfsdk:"is_uplink"`
	LinkNegotiation  types.String                                                                                                       `tfsdk:"link_negotiation"`
	Mac              types.String                                                                                                       `tfsdk:"mac"`
	Module           *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfacesModule       `tfsdk:"module"`
	Name             types.String                                                                                                       `tfsdk:"name"`
	Speed            types.String                                                                                                       `tfsdk:"speed"`
	Status           types.String                                                                                                       `tfsdk:"status"`
	VLAN             types.Int64                                                                                                        `tfsdk:"vlan"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfacesChannelGroup struct {
	Number types.Int64 `tfsdk:"number"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfacesModule struct {
	Model types.String `tfsdk:"model"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMetaCounts `tfsdk:"counts"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMetaCountsItems `tfsdk:"items"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemToBody(state OrganizationsWirelessControllerDevicesInterfacesL2ByDevice, response *merakigosdk.ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDevice) OrganizationsWirelessControllerDevicesInterfacesL2ByDevice {
	itemState := ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDevice{
		Items: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItems {
			if response.Items != nil {
				result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItems{
						Interfaces: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfaces {
							if items.Interfaces != nil {
								result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfaces, len(*items.Interfaces))
								for i, interfaces := range *items.Interfaces {
									result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfaces{
										ChannelGroup: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfacesChannelGroup {
											if interfaces.ChannelGroup != nil {
												return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfacesChannelGroup{
													Number: func() types.Int64 {
														if interfaces.ChannelGroup.Number != nil {
															return types.Int64Value(int64(*interfaces.ChannelGroup.Number))
														}
														return types.Int64{}
													}(),
												}
											}
											return nil
										}(),
										Description: func() types.String {
											if interfaces.Description != "" {
												return types.StringValue(interfaces.Description)
											}
											return types.String{}
										}(),
										Enabled: func() types.Bool {
											if interfaces.Enabled != nil {
												return types.BoolValue(*interfaces.Enabled)
											}
											return types.Bool{}
										}(),
										IsRedundancyPort: func() types.Bool {
											if interfaces.IsRedundancyPort != nil {
												return types.BoolValue(*interfaces.IsRedundancyPort)
											}
											return types.Bool{}
										}(),
										IsUplink: func() types.Bool {
											if interfaces.IsUplink != nil {
												return types.BoolValue(*interfaces.IsUplink)
											}
											return types.Bool{}
										}(),
										LinkNegotiation: func() types.String {
											if interfaces.LinkNegotiation != "" {
												return types.StringValue(interfaces.LinkNegotiation)
											}
											return types.String{}
										}(),
										Mac: func() types.String {
											if interfaces.Mac != "" {
												return types.StringValue(interfaces.Mac)
											}
											return types.String{}
										}(),
										Module: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfacesModule {
											if interfaces.Module != nil {
												return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceItemsInterfacesModule{
													Model: func() types.String {
														if interfaces.Module.Model != "" {
															return types.StringValue(interfaces.Module.Model)
														}
														return types.String{}
													}(),
												}
											}
											return nil
										}(),
										Name: func() types.String {
											if interfaces.Name != "" {
												return types.StringValue(interfaces.Name)
											}
											return types.String{}
										}(),
										Speed: func() types.String {
											if interfaces.Speed != "" {
												return types.StringValue(interfaces.Speed)
											}
											return types.String{}
										}(),
										Status: func() types.String {
											if interfaces.Status != "" {
												return types.StringValue(interfaces.Status)
											}
											return types.String{}
										}(),
										VLAN: func() types.Int64 {
											if interfaces.VLAN != nil {
												return types.Int64Value(int64(*interfaces.VLAN))
											}
											return types.Int64{}
										}(),
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
		Meta: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMeta {
			if response.Meta != nil {
				return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMeta{
					Counts: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMetaCounts{
								Items: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceMetaCountsItems{
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
