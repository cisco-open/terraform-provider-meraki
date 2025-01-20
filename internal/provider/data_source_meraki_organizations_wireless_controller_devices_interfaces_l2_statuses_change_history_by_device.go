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
	_ datasource.DataSource              = &OrganizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceDataSource{}
)

func NewOrganizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceDataSource() datasource.DataSource {
	return &OrganizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceDataSource{}
}

type OrganizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_controller_devices_interfaces_l2_statuses_change_history_by_device"
}

func (d *OrganizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
						MarkdownDescription: `Wireless LAN controller layer 2 interfaces historical status`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"interfaces": schema.SetNestedAttribute{
									MarkdownDescription: `layer 2 interfaces belongs to the wireless LAN controller`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"changes": schema.SetNestedAttribute{
												MarkdownDescription: `The statuses of layer 2 interfaces of the wireless LAN controller`,
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

func (d *OrganizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice OrganizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice
	diags := req.Config.Get(ctx, &organizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice")
		vvOrganizationID := organizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceQueryParams{}

		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice.Serials)
		queryParams1.IncludeInterfacesWithoutChanges = organizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice.IncludeInterfacesWithoutChanges.ValueBool()
		queryParams1.T0 = organizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice.T0.ValueString()
		queryParams1.T1 = organizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice.T1.ValueString()
		queryParams1.Timespan = organizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice.Timespan.ValueFloat64()
		queryParams1.PerPage = int(organizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.WirelessController.GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice",
				err.Error(),
			)
			return
		}

		organizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemToBody(organizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice struct {
	OrganizationID                  types.String                                                                                                 `tfsdk:"organization_id"`
	Serials                         types.List                                                                                                   `tfsdk:"serials"`
	IncludeInterfacesWithoutChanges types.Bool                                                                                                   `tfsdk:"include_interfaces_without_changes"`
	T0                              types.String                                                                                                 `tfsdk:"t0"`
	T1                              types.String                                                                                                 `tfsdk:"t1"`
	Timespan                        types.Float64                                                                                                `tfsdk:"timespan"`
	PerPage                         types.Int64                                                                                                  `tfsdk:"per_page"`
	StartingAfter                   types.String                                                                                                 `tfsdk:"starting_after"`
	EndingBefore                    types.String                                                                                                 `tfsdk:"ending_before"`
	Item                            *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice `tfsdk:"item"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItems `tfsdk:"items"`
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMeta    `tfsdk:"meta"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItems struct {
	Interfaces *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemsInterfaces `tfsdk:"interfaces"`
	Serial     types.String                                                                                                                  `tfsdk:"serial"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemsInterfaces struct {
	Changes *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemsInterfacesChanges `tfsdk:"changes"`
	Mac     types.String                                                                                                                         `tfsdk:"mac"`
	Name    types.String                                                                                                                         `tfsdk:"name"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemsInterfacesChanges struct {
	Errors   types.List   `tfsdk:"errors"`
	Status   types.String `tfsdk:"status"`
	Ts       types.String `tfsdk:"ts"`
	Warnings types.List   `tfsdk:"warnings"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMetaCounts `tfsdk:"counts"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMetaCountsItems `tfsdk:"items"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemToBody(state OrganizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice, response *merakigosdk.ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice) OrganizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice {
	itemState := ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice{
		Items: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItems {
			if response.Items != nil {
				result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItems{
						Interfaces: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemsInterfaces {
							if items.Interfaces != nil {
								result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemsInterfaces, len(*items.Interfaces))
								for i, interfaces := range *items.Interfaces {
									result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemsInterfaces{
										Changes: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemsInterfacesChanges {
											if interfaces.Changes != nil {
												result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemsInterfacesChanges, len(*interfaces.Changes))
												for i, changes := range *interfaces.Changes {
													result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemsInterfacesChanges{
														Errors:   StringSliceToList(changes.Errors),
														Status:   types.StringValue(changes.Status),
														Ts:       types.StringValue(changes.Ts),
														Warnings: StringSliceToList(changes.Warnings),
													}
												}
												return &result
											}
											return &[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemsInterfacesChanges{}
										}(),
										Mac:  types.StringValue(interfaces.Mac),
										Name: types.StringValue(interfaces.Name),
									}
								}
								return &result
							}
							return &[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItemsInterfaces{}
						}(),
						Serial: types.StringValue(items.Serial),
					}
				}
				return &result
			}
			return &[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceItems{}
		}(),
		Meta: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMeta {
			if response.Meta != nil {
				return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMeta{
					Counts: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMetaCounts{
								Items: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMetaCountsItems{
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
									return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMetaCountsItems{}
								}(),
							}
						}
						return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMetaCounts{}
					}(),
				}
			}
			return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceMeta{}
		}(),
	}
	state.Item = &itemState
	return state
}
