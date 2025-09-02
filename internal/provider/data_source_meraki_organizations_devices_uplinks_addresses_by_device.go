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
	_ datasource.DataSource              = &OrganizationsDevicesUplinksAddressesByDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsDevicesUplinksAddressesByDeviceDataSource{}
)

func NewOrganizationsDevicesUplinksAddressesByDeviceDataSource() datasource.DataSource {
	return &OrganizationsDevicesUplinksAddressesByDeviceDataSource{}
}

type OrganizationsDevicesUplinksAddressesByDeviceDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsDevicesUplinksAddressesByDeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsDevicesUplinksAddressesByDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_devices_uplinks_addresses_by_device"
}

func (d *OrganizationsDevicesUplinksAddressesByDeviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter device uplinks by network ID. This filter uses multiple exact matches.`,
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
			"product_types": schema.ListAttribute{
				MarkdownDescription: `productTypes query parameter. Optional parameter to filter device uplinks by device product types. This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"tags": schema.ListAttribute{
				MarkdownDescription: `tags query parameter. An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"tags_filter_type": schema.StringAttribute{
				MarkdownDescription: `tagsFilterType query parameter. An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationDevicesUplinksAddressesByDevice`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"mac": schema.StringAttribute{
							MarkdownDescription: `The device MAC address.`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The device name.`,
							Computed:            true,
						},
						"network": schema.SingleNestedAttribute{
							MarkdownDescription: `Network info.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `ID for the network containing the device.`,
									Computed:            true,
								},
							},
						},
						"product_type": schema.StringAttribute{
							MarkdownDescription: `Device product type.`,
							Computed:            true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `The device serial number.`,
							Computed:            true,
						},
						"tags": schema.ListAttribute{
							MarkdownDescription: `List of custom tags for the device.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"uplinks": schema.SetNestedAttribute{
							MarkdownDescription: `List of device uplink addresses information.`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"addresses": schema.SetNestedAttribute{
										MarkdownDescription: `Available addresses for the interface.`,
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{

												"address": schema.StringAttribute{
													MarkdownDescription: `Device uplink address.`,
													Computed:            true,
												},
												"assignment_mode": schema.StringAttribute{
													MarkdownDescription: `Indicates how the device uplink address is assigned. Available options are: static, dynamic.`,
													Computed:            true,
												},
												"gateway": schema.StringAttribute{
													MarkdownDescription: `Device uplink gateway address.`,
													Computed:            true,
												},
												"nameservers": schema.SingleNestedAttribute{
													MarkdownDescription: `Device DNS nameserver information.`,
													Computed:            true,
													Attributes: map[string]schema.Attribute{

														"addresses": schema.ListAttribute{
															MarkdownDescription: `Device DNS nameserver address.`,
															Computed:            true,
															ElementType:         types.StringType,
														},
													},
												},
												"protocol": schema.StringAttribute{
													MarkdownDescription: `Type of address for the device uplink. Available options are: ipv4, ipv6.`,
													Computed:            true,
												},
												"public": schema.SingleNestedAttribute{
													MarkdownDescription: `Public interface information.`,
													Computed:            true,
													Attributes: map[string]schema.Attribute{

														"address": schema.StringAttribute{
															MarkdownDescription: `The device uplink public IP address.`,
															Computed:            true,
														},
													},
												},
												"vlan": schema.SingleNestedAttribute{
													MarkdownDescription: `VLAN information of the uplink interface`,
													Computed:            true,
													Attributes: map[string]schema.Attribute{

														"id": schema.StringAttribute{
															MarkdownDescription: `VLAN ID of the uplink interface`,
															Computed:            true,
														},
													},
												},
											},
										},
									},
									"interface": schema.StringAttribute{
										MarkdownDescription: `Interface for the device uplink. Available options are: cellular, man1, man2, wan1, wan2`,
										Computed:            true,
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

func (d *OrganizationsDevicesUplinksAddressesByDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsDevicesUplinksAddressesByDevice OrganizationsDevicesUplinksAddressesByDevice
	diags := req.Config.Get(ctx, &organizationsDevicesUplinksAddressesByDevice)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationDevicesUplinksAddressesByDevice")
		vvOrganizationID := organizationsDevicesUplinksAddressesByDevice.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationDevicesUplinksAddressesByDeviceQueryParams{}

		queryParams1.PerPage = int(organizationsDevicesUplinksAddressesByDevice.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsDevicesUplinksAddressesByDevice.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsDevicesUplinksAddressesByDevice.EndingBefore.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsDevicesUplinksAddressesByDevice.NetworkIDs)
		queryParams1.ProductTypes = elementsToStrings(ctx, organizationsDevicesUplinksAddressesByDevice.ProductTypes)
		queryParams1.Serials = elementsToStrings(ctx, organizationsDevicesUplinksAddressesByDevice.Serials)
		queryParams1.Tags = elementsToStrings(ctx, organizationsDevicesUplinksAddressesByDevice.Tags)
		queryParams1.TagsFilterType = organizationsDevicesUplinksAddressesByDevice.TagsFilterType.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationDevicesUplinksAddressesByDevice(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationDevicesUplinksAddressesByDevice",
				err.Error(),
			)
			return
		}

		organizationsDevicesUplinksAddressesByDevice = ResponseOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceItemsToBody(organizationsDevicesUplinksAddressesByDevice, response1)
		diags = resp.State.Set(ctx, &organizationsDevicesUplinksAddressesByDevice)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsDevicesUplinksAddressesByDevice struct {
	OrganizationID types.String                                                               `tfsdk:"organization_id"`
	PerPage        types.Int64                                                                `tfsdk:"per_page"`
	StartingAfter  types.String                                                               `tfsdk:"starting_after"`
	EndingBefore   types.String                                                               `tfsdk:"ending_before"`
	NetworkIDs     types.List                                                                 `tfsdk:"network_ids"`
	ProductTypes   types.List                                                                 `tfsdk:"product_types"`
	Serials        types.List                                                                 `tfsdk:"serials"`
	Tags           types.List                                                                 `tfsdk:"tags"`
	TagsFilterType types.String                                                               `tfsdk:"tags_filter_type"`
	Items          *[]ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDevice `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDevice struct {
	Mac         types.String                                                                      `tfsdk:"mac"`
	Name        types.String                                                                      `tfsdk:"name"`
	Network     *ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceNetwork   `tfsdk:"network"`
	ProductType types.String                                                                      `tfsdk:"product_type"`
	Serial      types.String                                                                      `tfsdk:"serial"`
	Tags        types.List                                                                        `tfsdk:"tags"`
	Uplinks     *[]ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinks `tfsdk:"uplinks"`
}

type ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceNetwork struct {
	ID types.String `tfsdk:"id"`
}

type ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinks struct {
	Addresses *[]ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddresses `tfsdk:"addresses"`
	Interface types.String                                                                               `tfsdk:"interface"`
}

type ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddresses struct {
	Address        types.String                                                                                        `tfsdk:"address"`
	AssignmentMode types.String                                                                                        `tfsdk:"assignment_mode"`
	Gateway        types.String                                                                                        `tfsdk:"gateway"`
	Nameservers    *ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddressesNameservers `tfsdk:"nameservers"`
	Protocol       types.String                                                                                        `tfsdk:"protocol"`
	Public         *ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddressesPublic      `tfsdk:"public"`
	VLAN           *ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddressesVlan        `tfsdk:"vlan"`
}

type ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddressesNameservers struct {
	Addresses types.List `tfsdk:"addresses"`
}

type ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddressesPublic struct {
	Address types.String `tfsdk:"address"`
}

type ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddressesVlan struct {
	ID types.String `tfsdk:"id"`
}

// ToBody
func ResponseOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceItemsToBody(state OrganizationsDevicesUplinksAddressesByDevice, response *merakigosdk.ResponseOrganizationsGetOrganizationDevicesUplinksAddressesByDevice) OrganizationsDevicesUplinksAddressesByDevice {
	var items []ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDevice
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDevice{
			Mac: func() types.String {
				if item.Mac != "" {
					return types.StringValue(item.Mac)
				}
				return types.String{}
			}(),
			Name: func() types.String {
				if item.Name != "" {
					return types.StringValue(item.Name)
				}
				return types.String{}
			}(),
			Network: func() *ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceNetwork {
				if item.Network != nil {
					return &ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceNetwork{
						ID: func() types.String {
							if item.Network.ID != "" {
								return types.StringValue(item.Network.ID)
							}
							return types.String{}
						}(),
					}
				}
				return nil
			}(),
			ProductType: func() types.String {
				if item.ProductType != "" {
					return types.StringValue(item.ProductType)
				}
				return types.String{}
			}(),
			Serial: func() types.String {
				if item.Serial != "" {
					return types.StringValue(item.Serial)
				}
				return types.String{}
			}(),
			Tags: StringSliceToList(item.Tags),
			Uplinks: func() *[]ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinks {
				if item.Uplinks != nil {
					result := make([]ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinks, len(*item.Uplinks))
					for i, uplinks := range *item.Uplinks {
						result[i] = ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinks{
							Addresses: func() *[]ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddresses {
								if uplinks.Addresses != nil {
									result := make([]ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddresses, len(*uplinks.Addresses))
									for i, addresses := range *uplinks.Addresses {
										result[i] = ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddresses{
											Address: func() types.String {
												if addresses.Address != "" {
													return types.StringValue(addresses.Address)
												}
												return types.String{}
											}(),
											AssignmentMode: func() types.String {
												if addresses.AssignmentMode != "" {
													return types.StringValue(addresses.AssignmentMode)
												}
												return types.String{}
											}(),
											Gateway: func() types.String {
												if addresses.Gateway != "" {
													return types.StringValue(addresses.Gateway)
												}
												return types.String{}
											}(),
											Nameservers: func() *ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddressesNameservers {
												if addresses.Nameservers != nil {
													return &ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddressesNameservers{
														Addresses: StringSliceToList(addresses.Nameservers.Addresses),
													}
												}
												return nil
											}(),
											Protocol: func() types.String {
												if addresses.Protocol != "" {
													return types.StringValue(addresses.Protocol)
												}
												return types.String{}
											}(),
											Public: func() *ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddressesPublic {
												if addresses.Public != nil {
													return &ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddressesPublic{
														Address: func() types.String {
															if addresses.Public.Address != "" {
																return types.StringValue(addresses.Public.Address)
															}
															return types.String{}
														}(),
													}
												}
												return nil
											}(),
											VLAN: func() *ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddressesVlan {
												if addresses.VLAN != nil {
													return &ResponseItemOrganizationsGetOrganizationDevicesUplinksAddressesByDeviceUplinksAddressesVlan{
														ID: func() types.String {
															if addresses.VLAN.ID != "" {
																return types.StringValue(addresses.VLAN.ID)
															}
															return types.String{}
														}(),
													}
												}
												return nil
											}(),
										}
									}
									return &result
								}
								return nil
							}(),
							Interface: func() types.String {
								if uplinks.Interface != "" {
									return types.StringValue(uplinks.Interface)
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
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
