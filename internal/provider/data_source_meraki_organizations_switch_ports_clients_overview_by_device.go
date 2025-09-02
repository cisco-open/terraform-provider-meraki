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
	_ datasource.DataSource              = &OrganizationsSwitchPortsClientsOverviewByDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSwitchPortsClientsOverviewByDeviceDataSource{}
)

func NewOrganizationsSwitchPortsClientsOverviewByDeviceDataSource() datasource.DataSource {
	return &OrganizationsSwitchPortsClientsOverviewByDeviceDataSource{}
}

type OrganizationsSwitchPortsClientsOverviewByDeviceDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSwitchPortsClientsOverviewByDeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSwitchPortsClientsOverviewByDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_switch_ports_clients_overview_by_device"
}

func (d *OrganizationsSwitchPortsClientsOverviewByDeviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 20. Default is 20.`,
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
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 31 days from today.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.`,
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
									MarkdownDescription: `The number of online clients of the ports on the switch.`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"counts": schema.SingleNestedAttribute{
												MarkdownDescription: `Number of clients on the port in a given time.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"by_status": schema.SingleNestedAttribute{
														MarkdownDescription: `Associated client count on access point by status.`,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"online": schema.Int64Attribute{
																MarkdownDescription: `Active client count.`,
																Computed:            true,
															},
														},
													},
												},
											},
											"port_id": schema.StringAttribute{
												MarkdownDescription: `The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module.`,
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

func (d *OrganizationsSwitchPortsClientsOverviewByDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSwitchPortsClientsOverviewByDevice OrganizationsSwitchPortsClientsOverviewByDevice
	diags := req.Config.Get(ctx, &organizationsSwitchPortsClientsOverviewByDevice)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSwitchPortsClientsOverviewByDevice")
		vvOrganizationID := organizationsSwitchPortsClientsOverviewByDevice.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationSwitchPortsClientsOverviewByDeviceQueryParams{}

		queryParams1.T0 = organizationsSwitchPortsClientsOverviewByDevice.T0.ValueString()
		queryParams1.Timespan = organizationsSwitchPortsClientsOverviewByDevice.Timespan.ValueFloat64()
		queryParams1.PerPage = int(organizationsSwitchPortsClientsOverviewByDevice.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsSwitchPortsClientsOverviewByDevice.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsSwitchPortsClientsOverviewByDevice.EndingBefore.ValueString()
		queryParams1.ConfigurationUpdatedAfter = organizationsSwitchPortsClientsOverviewByDevice.ConfigurationUpdatedAfter.ValueString()
		queryParams1.Mac = organizationsSwitchPortsClientsOverviewByDevice.Mac.ValueString()
		queryParams1.Macs = elementsToStrings(ctx, organizationsSwitchPortsClientsOverviewByDevice.Macs)
		queryParams1.Name = organizationsSwitchPortsClientsOverviewByDevice.Name.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsSwitchPortsClientsOverviewByDevice.NetworkIDs)
		queryParams1.PortProfileIDs = elementsToStrings(ctx, organizationsSwitchPortsClientsOverviewByDevice.PortProfileIDs)
		queryParams1.Serial = organizationsSwitchPortsClientsOverviewByDevice.Serial.ValueString()
		queryParams1.Serials = elementsToStrings(ctx, organizationsSwitchPortsClientsOverviewByDevice.Serials)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetOrganizationSwitchPortsClientsOverviewByDevice(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSwitchPortsClientsOverviewByDevice",
				err.Error(),
			)
			return
		}

		organizationsSwitchPortsClientsOverviewByDevice = ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemToBody(organizationsSwitchPortsClientsOverviewByDevice, response1)
		diags = resp.State.Set(ctx, &organizationsSwitchPortsClientsOverviewByDevice)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSwitchPortsClientsOverviewByDevice struct {
	OrganizationID            types.String                                                     `tfsdk:"organization_id"`
	T0                        types.String                                                     `tfsdk:"t0"`
	Timespan                  types.Float64                                                    `tfsdk:"timespan"`
	PerPage                   types.Int64                                                      `tfsdk:"per_page"`
	StartingAfter             types.String                                                     `tfsdk:"starting_after"`
	EndingBefore              types.String                                                     `tfsdk:"ending_before"`
	ConfigurationUpdatedAfter types.String                                                     `tfsdk:"configuration_updated_after"`
	Mac                       types.String                                                     `tfsdk:"mac"`
	Macs                      types.List                                                       `tfsdk:"macs"`
	Name                      types.String                                                     `tfsdk:"name"`
	NetworkIDs                types.List                                                       `tfsdk:"network_ids"`
	PortProfileIDs            types.List                                                       `tfsdk:"port_profile_ids"`
	Serial                    types.String                                                     `tfsdk:"serial"`
	Serials                   types.List                                                       `tfsdk:"serials"`
	Item                      *ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDevice `tfsdk:"item"`
}

type ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDevice struct {
	Items *[]ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItems `tfsdk:"items"`
	Meta  *ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceMeta    `tfsdk:"meta"`
}

type ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItems struct {
	Mac     types.String                                                                 `tfsdk:"mac"`
	Model   types.String                                                                 `tfsdk:"model"`
	Name    types.String                                                                 `tfsdk:"name"`
	Network *ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsNetwork `tfsdk:"network"`
	Ports   *[]ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsPorts `tfsdk:"ports"`
	Serial  types.String                                                                 `tfsdk:"serial"`
}

type ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsPorts struct {
	Counts *ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsPortsCounts `tfsdk:"counts"`
	PortID types.String                                                                     `tfsdk:"port_id"`
}

type ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsPortsCounts struct {
	ByStatus *ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsPortsCountsByStatus `tfsdk:"by_status"`
}

type ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsPortsCountsByStatus struct {
	Online types.Int64 `tfsdk:"online"`
}

type ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceMeta struct {
	Counts *ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceMetaCounts `tfsdk:"counts"`
}

type ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceMetaCounts struct {
	Items *ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceMetaCountsItems `tfsdk:"items"`
}

type ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemToBody(state OrganizationsSwitchPortsClientsOverviewByDevice, response *merakigosdk.ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDevice) OrganizationsSwitchPortsClientsOverviewByDevice {
	itemState := ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDevice{
		Items: func() *[]ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItems {
			if response.Items != nil {
				result := make([]ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItems{
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
						Network: func() *ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsNetwork {
							if items.Network != nil {
								return &ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsNetwork{
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
						Ports: func() *[]ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsPorts {
							if items.Ports != nil {
								result := make([]ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsPorts, len(*items.Ports))
								for i, ports := range *items.Ports {
									result[i] = ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsPorts{
										Counts: func() *ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsPortsCounts {
											if ports.Counts != nil {
												return &ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsPortsCounts{
													ByStatus: func() *ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsPortsCountsByStatus {
														if ports.Counts.ByStatus != nil {
															return &ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceItemsPortsCountsByStatus{
																Online: func() types.Int64 {
																	if ports.Counts.ByStatus.Online != nil {
																		return types.Int64Value(int64(*ports.Counts.ByStatus.Online))
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
										PortID: func() types.String {
											if ports.PortID != "" {
												return types.StringValue(ports.PortID)
											}
											return types.String{}
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
		Meta: func() *ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceMeta {
			if response.Meta != nil {
				return &ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceMeta{
					Counts: func() *ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceMetaCounts{
								Items: func() *ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseSwitchGetOrganizationSwitchPortsClientsOverviewByDeviceMetaCountsItems{
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
