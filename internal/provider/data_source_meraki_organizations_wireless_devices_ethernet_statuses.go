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
	_ datasource.DataSource              = &OrganizationsWirelessDevicesEthernetStatusesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessDevicesEthernetStatusesDataSource{}
)

func NewOrganizationsWirelessDevicesEthernetStatusesDataSource() datasource.DataSource {
	return &OrganizationsWirelessDevicesEthernetStatusesDataSource{}
}

type OrganizationsWirelessDevicesEthernetStatusesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessDevicesEthernetStatusesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessDevicesEthernetStatusesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_devices_ethernet_statuses"
}

func (d *OrganizationsWirelessDevicesEthernetStatusesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]=N_12345678&networkIds[]=L_3456`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 100.`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseWirelessGetOrganizationWirelessDevicesEthernetStatuses`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"aggregation": schema.SingleNestedAttribute{
							MarkdownDescription: `Aggregation details object`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"enabled": schema.BoolAttribute{
									MarkdownDescription: `Link Aggregation enabled flag will return null on Catalyst devices`,
									Computed:            true,
								},
								"speed": schema.Int64Attribute{
									MarkdownDescription: `Link Aggregation speed will return null on Catalyst devices`,
									Computed:            true,
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the AP`,
							Computed:            true,
						},
						"network": schema.SingleNestedAttribute{
							MarkdownDescription: `Network details object`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `The network ID the AP is associated to`,
									Computed:            true,
								},
							},
						},
						"ports": schema.SetNestedAttribute{
							MarkdownDescription: `List of port details`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"link_negotiation": schema.SingleNestedAttribute{
										MarkdownDescription: `Link negotiation details object for the port`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"duplex": schema.StringAttribute{
												MarkdownDescription: `The duplex mode of the port. Can be 'full' or 'half' will return null on Catalyst devices`,
												Computed:            true,
											},
											"speed": schema.Int64Attribute{
												MarkdownDescription: `Show the speed of the port. The port speed will return null on Catalyst devices`,
												Computed:            true,
											},
										},
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `Label of the port`,
										Computed:            true,
									},
									"poe": schema.SingleNestedAttribute{
										MarkdownDescription: `PoE details object for the port`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"standard": schema.StringAttribute{
												MarkdownDescription: `The PoE Standard for the port. Can be '802.3at', '802.3af', '802.3bt', or null`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
						"power": schema.SingleNestedAttribute{
							MarkdownDescription: `Power details object`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"ac": schema.SingleNestedAttribute{
									MarkdownDescription: `AC power details object`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"is_connected": schema.BoolAttribute{
											MarkdownDescription: `AC power connected`,
											Computed:            true,
										},
									},
								},
								"mode": schema.StringAttribute{
									MarkdownDescription: `The PoE power mode for the AP. Can be 'full' or 'low'`,
									Computed:            true,
								},
								"poe": schema.SingleNestedAttribute{
									MarkdownDescription: `PoE power details object`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"is_connected": schema.BoolAttribute{
											MarkdownDescription: `PoE power connected`,
											Computed:            true,
										},
									},
								},
							},
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `The serial number of the AP`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsWirelessDevicesEthernetStatusesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessDevicesEthernetStatuses OrganizationsWirelessDevicesEthernetStatuses
	diags := req.Config.Get(ctx, &organizationsWirelessDevicesEthernetStatuses)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessDevicesEthernetStatuses")
		vvOrganizationID := organizationsWirelessDevicesEthernetStatuses.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessDevicesEthernetStatusesQueryParams{}

		queryParams1.PerPage = int(organizationsWirelessDevicesEthernetStatuses.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessDevicesEthernetStatuses.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessDevicesEthernetStatuses.EndingBefore.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsWirelessDevicesEthernetStatuses.NetworkIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetOrganizationWirelessDevicesEthernetStatuses(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessDevicesEthernetStatuses",
				err.Error(),
			)
			return
		}

		organizationsWirelessDevicesEthernetStatuses = ResponseWirelessGetOrganizationWirelessDevicesEthernetStatusesItemsToBody(organizationsWirelessDevicesEthernetStatuses, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessDevicesEthernetStatuses)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessDevicesEthernetStatuses struct {
	OrganizationID types.String                                                          `tfsdk:"organization_id"`
	PerPage        types.Int64                                                           `tfsdk:"per_page"`
	StartingAfter  types.String                                                          `tfsdk:"starting_after"`
	EndingBefore   types.String                                                          `tfsdk:"ending_before"`
	NetworkIDs     types.List                                                            `tfsdk:"network_ids"`
	Items          *[]ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatuses `tfsdk:"items"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatuses struct {
	Aggregation *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesAggregation `tfsdk:"aggregation"`
	Name        types.String                                                                   `tfsdk:"name"`
	Network     *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesNetwork     `tfsdk:"network"`
	Ports       *[]ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPorts     `tfsdk:"ports"`
	Power       *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPower       `tfsdk:"power"`
	Serial      types.String                                                                   `tfsdk:"serial"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesAggregation struct {
	Enabled types.Bool  `tfsdk:"enabled"`
	Speed   types.Int64 `tfsdk:"speed"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesNetwork struct {
	ID types.String `tfsdk:"id"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPorts struct {
	LinkNegotiation *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPortsLinkNegotiation `tfsdk:"link_negotiation"`
	Name            types.String                                                                            `tfsdk:"name"`
	Poe             *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPortsPoe             `tfsdk:"poe"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPortsLinkNegotiation struct {
	Duplex types.String `tfsdk:"duplex"`
	Speed  types.Int64  `tfsdk:"speed"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPortsPoe struct {
	Standard types.String `tfsdk:"standard"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPower struct {
	Ac   *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPowerAc  `tfsdk:"ac"`
	Mode types.String                                                                `tfsdk:"mode"`
	Poe  *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPowerPoe `tfsdk:"poe"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPowerAc struct {
	IsConnected types.Bool `tfsdk:"is_connected"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPowerPoe struct {
	IsConnected types.Bool `tfsdk:"is_connected"`
}

// ToBody
func ResponseWirelessGetOrganizationWirelessDevicesEthernetStatusesItemsToBody(state OrganizationsWirelessDevicesEthernetStatuses, response *merakigosdk.ResponseWirelessGetOrganizationWirelessDevicesEthernetStatuses) OrganizationsWirelessDevicesEthernetStatuses {
	var items []ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatuses
	for _, item := range *response {
		itemState := ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatuses{
			Aggregation: func() *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesAggregation {
				if item.Aggregation != nil {
					return &ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesAggregation{
						Enabled: func() types.Bool {
							if item.Aggregation.Enabled != nil {
								return types.BoolValue(*item.Aggregation.Enabled)
							}
							return types.Bool{}
						}(),
						Speed: func() types.Int64 {
							if item.Aggregation.Speed != nil {
								return types.Int64Value(int64(*item.Aggregation.Speed))
							}
							return types.Int64{}
						}(),
					}
				}
				return nil
			}(),
			Name: types.StringValue(item.Name),
			Network: func() *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesNetwork {
				if item.Network != nil {
					return &ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesNetwork{
						ID: types.StringValue(item.Network.ID),
					}
				}
				return nil
			}(),
			Ports: func() *[]ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPorts {
				if item.Ports != nil {
					result := make([]ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPorts, len(*item.Ports))
					for i, ports := range *item.Ports {
						result[i] = ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPorts{
							LinkNegotiation: func() *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPortsLinkNegotiation {
								if ports.LinkNegotiation != nil {
									return &ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPortsLinkNegotiation{
										Duplex: types.StringValue(ports.LinkNegotiation.Duplex),
										Speed: func() types.Int64 {
											if ports.LinkNegotiation.Speed != nil {
												return types.Int64Value(int64(*ports.LinkNegotiation.Speed))
											}
											return types.Int64{}
										}(),
									}
								}
								return nil
							}(),
							Name: types.StringValue(ports.Name),
							Poe: func() *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPortsPoe {
								if ports.Poe != nil {
									return &ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPortsPoe{
										Standard: types.StringValue(ports.Poe.Standard),
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
			Power: func() *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPower {
				if item.Power != nil {
					return &ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPower{
						Ac: func() *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPowerAc {
							if item.Power.Ac != nil {
								return &ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPowerAc{
									IsConnected: func() types.Bool {
										if item.Power.Ac.IsConnected != nil {
											return types.BoolValue(*item.Power.Ac.IsConnected)
										}
										return types.Bool{}
									}(),
								}
							}
							return nil
						}(),
						Mode: types.StringValue(item.Power.Mode),
						Poe: func() *ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPowerPoe {
							if item.Power.Poe != nil {
								return &ResponseItemWirelessGetOrganizationWirelessDevicesEthernetStatusesPowerPoe{
									IsConnected: func() types.Bool {
										if item.Power.Poe.IsConnected != nil {
											return types.BoolValue(*item.Power.Poe.IsConnected)
										}
										return types.Bool{}
									}(),
								}
							}
							return nil
						}(),
					}
				}
				return nil
			}(),
			Serial: types.StringValue(item.Serial),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
