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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceDataSource{}
)

func NewOrganizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceDataSource() datasource.DataSource {
	return &OrganizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceDataSource{}
}

type OrganizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_controller_devices_interfaces_l3_statuses_change_history_by_device"
}

func (d *OrganizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"include_interfaces_without_changes": schema.BoolAttribute{
				MarkdownDescription: `includeInterfacesWithoutChanges query parameter. By default, interfaces without changes are omitted from the response for brevity. If you want to include the interfaces even if they have no changes, set to true. (default: false)`,
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
						MarkdownDescription: `Wireless LAN controller layer 3 interfaces historical status`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"interfaces": schema.SetNestedAttribute{
									MarkdownDescription: `layer 3 interfaces belongs to the wireless LAN controller`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"changes": schema.SetNestedAttribute{
												MarkdownDescription: `The statuses of layer 3 interfaces of the wireless LAN controller`,
												Computed:            true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{

														"errors": schema.ListAttribute{
															MarkdownDescription: `All errors present on the port`,
															Computed:            true,
															ElementType:         types.StringType,
														},
														"status": schema.StringAttribute{
															MarkdownDescription: `The status of the interface`,
															Computed:            true,
														},
														"ts": schema.StringAttribute{
															MarkdownDescription: `The timestamp of current status of the interface`,
															Computed:            true,
														},
														"warnings": schema.ListAttribute{
															MarkdownDescription: `All warnings present on the port`,
															Computed:            true,
															ElementType:         types.StringType,
														},
													},
												},
											},
											"mac": schema.StringAttribute{
												MarkdownDescription: `The MAC address of the wireless LAN controller interface`,
												Computed:            true,
											},
											"name": schema.StringAttribute{
												MarkdownDescription: `The name of the wireless LAN controller interface`,
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

func (d *OrganizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice OrganizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice
	diags := req.Config.Get(ctx, &organizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice")
		vvOrganizationID := organizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceQueryParams{}

		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice.Serials)
		queryParams1.IncludeInterfacesWithoutChanges = organizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice.IncludeInterfacesWithoutChanges.ValueBool()
		queryParams1.T0 = organizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice.T0.ValueString()
		queryParams1.T1 = organizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice.T1.ValueString()
		queryParams1.Timespan = organizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice.Timespan.ValueFloat64()
		queryParams1.PerPage = int(organizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.WirelessController.GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice",
				err.Error(),
			)
			return
		}

		organizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItemToBody(organizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice struct {
	OrganizationID                  types.String                                                                                                 `tfsdk:"organization_id"`
	Serials                         types.List                                                                                                   `tfsdk:"serials"`
	IncludeInterfacesWithoutChanges types.Bool                                                                                                   `tfsdk:"include_interfaces_without_changes"`
	T0                              types.String                                                                                                 `tfsdk:"t0"`
	T1                              types.String                                                                                                 `tfsdk:"t1"`
	Timespan                        types.Float64                                                                                                `tfsdk:"timespan"`
	PerPage                         types.Int64                                                                                                  `tfsdk:"per_page"`
	StartingAfter                   types.String                                                                                                 `tfsdk:"starting_after"`
	EndingBefore                    types.String                                                                                                 `tfsdk:"ending_before"`
	Item                            *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice `tfsdk:"item"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItems `tfsdk:"items"`
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMeta    `tfsdk:"meta"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItems struct {
	Interfaces *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItemsInterfaces `tfsdk:"interfaces"`
	Serial     types.String                                                                                                                  `tfsdk:"serial"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItemsInterfaces struct {
	Changes *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItemsInterfacesChanges `tfsdk:"changes"`
	Mac     types.String                                                                                                                         `tfsdk:"mac"`
	Name    types.String                                                                                                                         `tfsdk:"name"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItemsInterfacesChanges struct {
	Errors   types.List   `tfsdk:"errors"`
	Status   types.String `tfsdk:"status"`
	Ts       types.String `tfsdk:"ts"`
	Warnings types.List   `tfsdk:"warnings"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMetaCounts `tfsdk:"counts"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMetaCountsItems `tfsdk:"items"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItemToBody(state OrganizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice, response *merakigosdk.ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice) OrganizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice {
	itemState := ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice{
		Items: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItems {
			if response.Items != nil {
				result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItems{
						Interfaces: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItemsInterfaces {
							if items.Interfaces != nil {
								result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItemsInterfaces, len(*items.Interfaces))
								for i, interfaces := range *items.Interfaces {
									result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItemsInterfaces{
										Changes: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItemsInterfacesChanges {
											if interfaces.Changes != nil {
												result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItemsInterfacesChanges, len(*interfaces.Changes))
												for i, changes := range *interfaces.Changes {
													result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceItemsInterfacesChanges{
														Errors:   StringSliceToList(changes.Errors),
														Status:   types.StringValue(changes.Status),
														Ts:       types.StringValue(changes.Ts),
														Warnings: StringSliceToList(changes.Warnings),
													}
												}
												return &result
											}
											return nil
										}(),
										Mac:  types.StringValue(interfaces.Mac),
										Name: types.StringValue(interfaces.Name),
									}
								}
								return &result
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
		Meta: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMeta {
			if response.Meta != nil {
				return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMeta{
					Counts: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMetaCounts{
								Items: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceMetaCountsItems{
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
