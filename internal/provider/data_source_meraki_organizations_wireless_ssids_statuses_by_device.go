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
	_ datasource.DataSource              = &OrganizationsWirelessSSIDsStatusesByDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessSSIDsStatusesByDeviceDataSource{}
)

func NewOrganizationsWirelessSSIDsStatusesByDeviceDataSource() datasource.DataSource {
	return &OrganizationsWirelessSSIDsStatusesByDeviceDataSource{}
}

type OrganizationsWirelessSSIDsStatusesByDeviceDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessSSIDsStatusesByDeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessSSIDsStatusesByDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_ssids_statuses_by_device"
}

func (d *OrganizationsWirelessSSIDsStatusesByDeviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"bssids": schema.ListAttribute{
				MarkdownDescription: `bssids query parameter. A list of BSSIDs. The returned devices will be filtered to only include these BSSIDs.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"hide_disabled": schema.BoolAttribute{
				MarkdownDescription: `hideDisabled query parameter. If true, the returned devices will not include disabled SSIDs. (default: true)`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter the result set by the included set of network IDs`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 500. Default is 100.`,
				Optional:            true,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. A list of serial numbers. The returned devices will be filtered to only include these serials.`,
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
						MarkdownDescription: `The top-level propery containing all status data.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"basic_service_sets": schema.SetNestedAttribute{
									MarkdownDescription: `Status information for wireless access points.`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"bssid": schema.StringAttribute{
												MarkdownDescription: `Unique identifier for wireless access point.`,
												Computed:            true,
											},
											"radio": schema.SingleNestedAttribute{
												MarkdownDescription: `Wireless access point radio identifier.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"band": schema.StringAttribute{
														MarkdownDescription: `Frequency range used for wireless communication.`,
														Computed:            true,
													},
													"channel": schema.Int64Attribute{
														MarkdownDescription: `Frequency channel used for wireless communication.`,
														Computed:            true,
													},
													"channel_width": schema.Int64Attribute{
														MarkdownDescription: `Width of frequency channel used for wireless communication.`,
														Computed:            true,
													},
													"index": schema.StringAttribute{
														MarkdownDescription: `The radio index.`,
														Computed:            true,
													},
													"is_broadcasting": schema.BoolAttribute{
														MarkdownDescription: `Indicates whether or not this radio is currently broadcasting.`,
														Computed:            true,
													},
													"power": schema.Int64Attribute{
														MarkdownDescription: `Strength of wireless signal.`,
														Computed:            true,
													},
												},
											},
											"ssid": schema.SingleNestedAttribute{
												MarkdownDescription: `Wireless access point and network identifier.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"advertised": schema.BoolAttribute{
														MarkdownDescription: `Availability of wireless network for devices to connect to.`,
														Computed:            true,
													},
													"enabled": schema.BoolAttribute{
														MarkdownDescription: `Status of wireless network.`,
														Computed:            true,
													},
													"name": schema.StringAttribute{
														MarkdownDescription: `Name of wireless network.`,
														Computed:            true,
													},
													"number": schema.Int64Attribute{
														MarkdownDescription: `Unique identifier for wireless network.`,
														Computed:            true,
													},
												},
											},
										},
									},
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Name of device.`,
									Computed:            true,
								},
								"network": schema.SingleNestedAttribute{
									MarkdownDescription: `Group of devices and settings.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `Unique identifier for network.`,
											Computed:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: `Name of network.`,
											Computed:            true,
										},
									},
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `Unique serial number for device.`,
									Computed:            true,
								},
							},
						},
					},
					"meta": schema.SingleNestedAttribute{
						MarkdownDescription: `Other metadata related to this result set.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"counts": schema.SingleNestedAttribute{
								MarkdownDescription: `Count metadata related to this result set.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"items": schema.SingleNestedAttribute{
										MarkdownDescription: `The count metadata.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"remaining": schema.Int64Attribute{
												MarkdownDescription: `The number of items remaining based on current pagination location within the dataset.`,
												Computed:            true,
											},
											"total": schema.Int64Attribute{
												MarkdownDescription: `The total number of items.`,
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

func (d *OrganizationsWirelessSSIDsStatusesByDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessSSIDsStatusesByDevice OrganizationsWirelessSSIDsStatusesByDevice
	diags := req.Config.Get(ctx, &organizationsWirelessSSIDsStatusesByDevice)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessSSIDsStatusesByDevice")
		vvOrganizationID := organizationsWirelessSSIDsStatusesByDevice.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessSSIDsStatusesByDeviceQueryParams{}

		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsWirelessSSIDsStatusesByDevice.NetworkIDs)
		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessSSIDsStatusesByDevice.Serials)
		queryParams1.Bssids = elementsToStrings(ctx, organizationsWirelessSSIDsStatusesByDevice.Bssids)
		queryParams1.HideDisabled = organizationsWirelessSSIDsStatusesByDevice.HideDisabled.ValueBool()
		queryParams1.PerPage = int(organizationsWirelessSSIDsStatusesByDevice.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessSSIDsStatusesByDevice.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessSSIDsStatusesByDevice.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetOrganizationWirelessSSIDsStatusesByDevice(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessSSIDsStatusesByDevice",
				err.Error(),
			)
			return
		}

		organizationsWirelessSSIDsStatusesByDevice = ResponseWirelessGetOrganizationWirelessSSIDsStatusesByDeviceItemToBody(organizationsWirelessSSIDsStatusesByDevice, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessSSIDsStatusesByDevice)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessSSIDsStatusesByDevice struct {
	OrganizationID types.String                                                  `tfsdk:"organization_id"`
	NetworkIDs     types.List                                                    `tfsdk:"network_ids"`
	Serials        types.List                                                    `tfsdk:"serials"`
	Bssids         types.List                                                    `tfsdk:"bssids"`
	HideDisabled   types.Bool                                                    `tfsdk:"hide_disabled"`
	PerPage        types.Int64                                                   `tfsdk:"per_page"`
	StartingAfter  types.String                                                  `tfsdk:"starting_after"`
	EndingBefore   types.String                                                  `tfsdk:"ending_before"`
	Item           *ResponseWirelessGetOrganizationWirelessSsidsStatusesByDevice `tfsdk:"item"`
}

type ResponseWirelessGetOrganizationWirelessSsidsStatusesByDevice struct {
	Items *[]ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItems `tfsdk:"items"`
	Meta  *ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceMeta    `tfsdk:"meta"`
}

type ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItems struct {
	BasicServiceSets *[]ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItemsBasicServiceSets `tfsdk:"basic_service_sets"`
	Name             types.String                                                                         `tfsdk:"name"`
	Network          *ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItemsNetwork            `tfsdk:"network"`
	Serial           types.String                                                                         `tfsdk:"serial"`
}

type ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItemsBasicServiceSets struct {
	Bssid types.String                                                                            `tfsdk:"bssid"`
	Radio *ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItemsBasicServiceSetsRadio `tfsdk:"radio"`
	SSID  *ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItemsBasicServiceSetsSsid  `tfsdk:"ssid"`
}

type ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItemsBasicServiceSetsRadio struct {
	Band           types.String `tfsdk:"band"`
	Channel        types.Int64  `tfsdk:"channel"`
	ChannelWidth   types.Int64  `tfsdk:"channel_width"`
	Index          types.String `tfsdk:"index"`
	IsBroadcasting types.Bool   `tfsdk:"is_broadcasting"`
	Power          types.Int64  `tfsdk:"power"`
}

type ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItemsBasicServiceSetsSsid struct {
	Advertised types.Bool   `tfsdk:"advertised"`
	Enabled    types.Bool   `tfsdk:"enabled"`
	Name       types.String `tfsdk:"name"`
	Number     types.Int64  `tfsdk:"number"`
}

type ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItemsNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceMeta struct {
	Counts *ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceMetaCounts `tfsdk:"counts"`
}

type ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceMetaCounts struct {
	Items *ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceMetaCountsItems `tfsdk:"items"`
}

type ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseWirelessGetOrganizationWirelessSSIDsStatusesByDeviceItemToBody(state OrganizationsWirelessSSIDsStatusesByDevice, response *merakigosdk.ResponseWirelessGetOrganizationWirelessSSIDsStatusesByDevice) OrganizationsWirelessSSIDsStatusesByDevice {
	itemState := ResponseWirelessGetOrganizationWirelessSsidsStatusesByDevice{
		Items: func() *[]ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItems {
			if response.Items != nil {
				result := make([]ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItems{
						BasicServiceSets: func() *[]ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItemsBasicServiceSets {
							if items.BasicServiceSets != nil {
								result := make([]ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItemsBasicServiceSets, len(*items.BasicServiceSets))
								for i, basicServiceSets := range *items.BasicServiceSets {
									result[i] = ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItemsBasicServiceSets{
										Bssid: types.StringValue(basicServiceSets.Bssid),
										Radio: func() *ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItemsBasicServiceSetsRadio {
											if basicServiceSets.Radio != nil {
												return &ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItemsBasicServiceSetsRadio{
													Band: types.StringValue(basicServiceSets.Radio.Band),
													Channel: func() types.Int64 {
														if basicServiceSets.Radio.Channel != nil {
															return types.Int64Value(int64(*basicServiceSets.Radio.Channel))
														}
														return types.Int64{}
													}(),
													ChannelWidth: func() types.Int64 {
														if basicServiceSets.Radio.ChannelWidth != nil {
															return types.Int64Value(int64(*basicServiceSets.Radio.ChannelWidth))
														}
														return types.Int64{}
													}(),
													Index: types.StringValue(basicServiceSets.Radio.Index),
													IsBroadcasting: func() types.Bool {
														if basicServiceSets.Radio.IsBroadcasting != nil {
															return types.BoolValue(*basicServiceSets.Radio.IsBroadcasting)
														}
														return types.Bool{}
													}(),
													Power: func() types.Int64 {
														if basicServiceSets.Radio.Power != nil {
															return types.Int64Value(int64(*basicServiceSets.Radio.Power))
														}
														return types.Int64{}
													}(),
												}
											}
											return nil
										}(),
										SSID: func() *ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItemsBasicServiceSetsSsid {
											if basicServiceSets.SSID != nil {
												return &ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItemsBasicServiceSetsSsid{
													Advertised: func() types.Bool {
														if basicServiceSets.SSID.Advertised != nil {
															return types.BoolValue(*basicServiceSets.SSID.Advertised)
														}
														return types.Bool{}
													}(),
													Enabled: func() types.Bool {
														if basicServiceSets.SSID.Enabled != nil {
															return types.BoolValue(*basicServiceSets.SSID.Enabled)
														}
														return types.Bool{}
													}(),
													Name: types.StringValue(basicServiceSets.SSID.Name),
													Number: func() types.Int64 {
														if basicServiceSets.SSID.Number != nil {
															return types.Int64Value(int64(*basicServiceSets.SSID.Number))
														}
														return types.Int64{}
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
						Name: types.StringValue(items.Name),
						Network: func() *ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItemsNetwork {
							if items.Network != nil {
								return &ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceItemsNetwork{
									ID:   types.StringValue(items.Network.ID),
									Name: types.StringValue(items.Network.Name),
								}
							}
							return nil
						}(),
						Serial: types.StringValue(items.Serial),
					}
				}
				return &result
			}
			return nil
		}(),
		Meta: func() *ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceMeta {
			if response.Meta != nil {
				return &ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceMeta{
					Counts: func() *ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceMetaCounts{
								Items: func() *ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseWirelessGetOrganizationWirelessSsidsStatusesByDeviceMetaCountsItems{
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
