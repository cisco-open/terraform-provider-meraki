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
	_ datasource.DataSource              = &OrganizationsSwitchPortsTopologyDiscoveryByDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSwitchPortsTopologyDiscoveryByDeviceDataSource{}
)

func NewOrganizationsSwitchPortsTopologyDiscoveryByDeviceDataSource() datasource.DataSource {
	return &OrganizationsSwitchPortsTopologyDiscoveryByDeviceDataSource{}
}

type OrganizationsSwitchPortsTopologyDiscoveryByDeviceDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSwitchPortsTopologyDiscoveryByDeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSwitchPortsTopologyDiscoveryByDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_switch_ports_topology_discovery_by_device"
}

func (d *OrganizationsSwitchPortsTopologyDiscoveryByDeviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
									MarkdownDescription: `Ports belonging to the switch with LLDP/CDP discovery info.`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"cdp": schema.SetNestedAttribute{
												MarkdownDescription: `The Cisco Discovery Protocol (CDP) information of the connected device.`,
												Computed:            true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{

														"name": schema.StringAttribute{
															MarkdownDescription: `CDP RFC/official name of TLV`,
															Computed:            true,
														},
														"value": schema.StringAttribute{
															MarkdownDescription: `Value of the named TLV.`,
															Computed:            true,
														},
													},
												},
											},
											"last_updated_at": schema.StringAttribute{
												MarkdownDescription: `Timestamp for most recent discovery info on this port.`,
												Computed:            true,
											},
											"lldp": schema.SetNestedAttribute{
												MarkdownDescription: `The Link Layer Discovery Protocol (LLDP) information of the connected device.`,
												Computed:            true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{

														"name": schema.StringAttribute{
															MarkdownDescription: `LLDP RFC/official name of TLV`,
															Computed:            true,
														},
														"value": schema.StringAttribute{
															MarkdownDescription: `Value of the named TLV.`,
															Computed:            true,
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

func (d *OrganizationsSwitchPortsTopologyDiscoveryByDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSwitchPortsTopologyDiscoveryByDevice OrganizationsSwitchPortsTopologyDiscoveryByDevice
	diags := req.Config.Get(ctx, &organizationsSwitchPortsTopologyDiscoveryByDevice)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSwitchPortsTopologyDiscoveryByDevice")
		vvOrganizationID := organizationsSwitchPortsTopologyDiscoveryByDevice.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationSwitchPortsTopologyDiscoveryByDeviceQueryParams{}

		queryParams1.T0 = organizationsSwitchPortsTopologyDiscoveryByDevice.T0.ValueString()
		queryParams1.Timespan = organizationsSwitchPortsTopologyDiscoveryByDevice.Timespan.ValueFloat64()
		queryParams1.PerPage = int(organizationsSwitchPortsTopologyDiscoveryByDevice.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsSwitchPortsTopologyDiscoveryByDevice.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsSwitchPortsTopologyDiscoveryByDevice.EndingBefore.ValueString()
		queryParams1.ConfigurationUpdatedAfter = organizationsSwitchPortsTopologyDiscoveryByDevice.ConfigurationUpdatedAfter.ValueString()
		queryParams1.Mac = organizationsSwitchPortsTopologyDiscoveryByDevice.Mac.ValueString()
		queryParams1.Macs = elementsToStrings(ctx, organizationsSwitchPortsTopologyDiscoveryByDevice.Macs)
		queryParams1.Name = organizationsSwitchPortsTopologyDiscoveryByDevice.Name.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsSwitchPortsTopologyDiscoveryByDevice.NetworkIDs)
		queryParams1.PortProfileIDs = elementsToStrings(ctx, organizationsSwitchPortsTopologyDiscoveryByDevice.PortProfileIDs)
		queryParams1.Serial = organizationsSwitchPortsTopologyDiscoveryByDevice.Serial.ValueString()
		queryParams1.Serials = elementsToStrings(ctx, organizationsSwitchPortsTopologyDiscoveryByDevice.Serials)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetOrganizationSwitchPortsTopologyDiscoveryByDevice(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSwitchPortsTopologyDiscoveryByDevice",
				err.Error(),
			)
			return
		}

		organizationsSwitchPortsTopologyDiscoveryByDevice = ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemToBody(organizationsSwitchPortsTopologyDiscoveryByDevice, response1)
		diags = resp.State.Set(ctx, &organizationsSwitchPortsTopologyDiscoveryByDevice)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSwitchPortsTopologyDiscoveryByDevice struct {
	OrganizationID            types.String                                                       `tfsdk:"organization_id"`
	T0                        types.String                                                       `tfsdk:"t0"`
	Timespan                  types.Float64                                                      `tfsdk:"timespan"`
	PerPage                   types.Int64                                                        `tfsdk:"per_page"`
	StartingAfter             types.String                                                       `tfsdk:"starting_after"`
	EndingBefore              types.String                                                       `tfsdk:"ending_before"`
	ConfigurationUpdatedAfter types.String                                                       `tfsdk:"configuration_updated_after"`
	Mac                       types.String                                                       `tfsdk:"mac"`
	Macs                      types.List                                                         `tfsdk:"macs"`
	Name                      types.String                                                       `tfsdk:"name"`
	NetworkIDs                types.List                                                         `tfsdk:"network_ids"`
	PortProfileIDs            types.List                                                         `tfsdk:"port_profile_ids"`
	Serial                    types.String                                                       `tfsdk:"serial"`
	Serials                   types.List                                                         `tfsdk:"serials"`
	Item                      *ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDevice `tfsdk:"item"`
}

type ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDevice struct {
	Items *[]ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItems `tfsdk:"items"`
	Meta  *ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceMeta    `tfsdk:"meta"`
}

type ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItems struct {
	Mac     types.String                                                                   `tfsdk:"mac"`
	Model   types.String                                                                   `tfsdk:"model"`
	Name    types.String                                                                   `tfsdk:"name"`
	Network *ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsNetwork `tfsdk:"network"`
	Ports   *[]ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPorts `tfsdk:"ports"`
	Serial  types.String                                                                   `tfsdk:"serial"`
}

type ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPorts struct {
	Cdp           *[]ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPortsCdp  `tfsdk:"cdp"`
	LastUpdatedAt types.String                                                                       `tfsdk:"last_updated_at"`
	Lldp          *[]ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPortsLldp `tfsdk:"lldp"`
	PortID        types.String                                                                       `tfsdk:"port_id"`
}

type ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPortsCdp struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPortsLldp struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceMeta struct {
	Counts *ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceMetaCounts `tfsdk:"counts"`
}

type ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceMetaCounts struct {
	Items *ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceMetaCountsItems `tfsdk:"items"`
}

type ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemToBody(state OrganizationsSwitchPortsTopologyDiscoveryByDevice, response *merakigosdk.ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDevice) OrganizationsSwitchPortsTopologyDiscoveryByDevice {
	itemState := ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDevice{
		Items: func() *[]ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItems {
			if response.Items != nil {
				result := make([]ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItems{
						Mac:   types.StringValue(items.Mac),
						Model: types.StringValue(items.Model),
						Name:  types.StringValue(items.Name),
						Network: func() *ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsNetwork {
							if items.Network != nil {
								return &ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsNetwork{
									ID:   types.StringValue(items.Network.ID),
									Name: types.StringValue(items.Network.Name),
								}
							}
							return nil
						}(),
						Ports: func() *[]ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPorts {
							if items.Ports != nil {
								result := make([]ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPorts, len(*items.Ports))
								for i, ports := range *items.Ports {
									result[i] = ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPorts{
										Cdp: func() *[]ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPortsCdp {
											if ports.Cdp != nil {
												result := make([]ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPortsCdp, len(*ports.Cdp))
												for i, cdp := range *ports.Cdp {
													result[i] = ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPortsCdp{
														Name:  types.StringValue(cdp.Name),
														Value: types.StringValue(cdp.Value),
													}
												}
												return &result
											}
											return nil
										}(),
										LastUpdatedAt: types.StringValue(ports.LastUpdatedAt),
										Lldp: func() *[]ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPortsLldp {
											if ports.Lldp != nil {
												result := make([]ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPortsLldp, len(*ports.Lldp))
												for i, lldp := range *ports.Lldp {
													result[i] = ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceItemsPortsLldp{
														Name:  types.StringValue(lldp.Name),
														Value: types.StringValue(lldp.Value),
													}
												}
												return &result
											}
											return nil
										}(),
										PortID: types.StringValue(ports.PortID),
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
		Meta: func() *ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceMeta {
			if response.Meta != nil {
				return &ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceMeta{
					Counts: func() *ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceMetaCounts{
								Items: func() *ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseSwitchGetOrganizationSwitchPortsTopologyDiscoveryByDeviceMetaCountsItems{
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
