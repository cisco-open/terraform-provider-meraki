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
	_ datasource.DataSource              = &OrganizationsWirelessDevicesPacketLossByNetworkDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessDevicesPacketLossByNetworkDataSource{}
)

func NewOrganizationsWirelessDevicesPacketLossByNetworkDataSource() datasource.DataSource {
	return &OrganizationsWirelessDevicesPacketLossByNetworkDataSource{}
}

type OrganizationsWirelessDevicesPacketLossByNetworkDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessDevicesPacketLossByNetworkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessDevicesPacketLossByNetworkDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_devices_packet_loss_by_network"
}

func (d *OrganizationsWirelessDevicesPacketLossByNetworkDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"bands": schema.ListAttribute{
				MarkdownDescription: `bands query parameter. Filter results by band. Valid bands are: 2.4, 5, and 6.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Filter results by network.`,
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
				MarkdownDescription: `serials query parameter. Filter results by device.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"ssids": schema.ListAttribute{
				MarkdownDescription: `ssids query parameter. Filter results by SSID number.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 90 days from today.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 90 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseWirelessGetOrganizationWirelessDevicesPacketLossByNetwork`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"downstream": schema.SingleNestedAttribute{
							MarkdownDescription: `Packets sent from an AP to a client.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"loss_percentage": schema.Float64Attribute{
									MarkdownDescription: `Percentage of lost packets.`,
									Computed:            true,
								},
								"lost": schema.Int64Attribute{
									MarkdownDescription: `Total packets sent by an AP that did not reach the client.`,
									Computed:            true,
								},
								"total": schema.Int64Attribute{
									MarkdownDescription: `Total packets received by a client.`,
									Computed:            true,
								},
							},
						},
						"network": schema.SingleNestedAttribute{
							MarkdownDescription: `Network.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `Network ID.`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Name of the network.`,
									Computed:            true,
								},
							},
						},
						"upstream": schema.SingleNestedAttribute{
							MarkdownDescription: `Packets sent from a client to an AP.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"loss_percentage": schema.Float64Attribute{
									MarkdownDescription: `Percentage of lost packets.`,
									Computed:            true,
								},
								"lost": schema.Int64Attribute{
									MarkdownDescription: `Total packets sent by a client and did not reach the AP.`,
									Computed:            true,
								},
								"total": schema.Int64Attribute{
									MarkdownDescription: `Total packets sent by a client to an AP.`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsWirelessDevicesPacketLossByNetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessDevicesPacketLossByNetwork OrganizationsWirelessDevicesPacketLossByNetwork
	diags := req.Config.Get(ctx, &organizationsWirelessDevicesPacketLossByNetwork)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessDevicesPacketLossByNetwork")
		vvOrganizationID := organizationsWirelessDevicesPacketLossByNetwork.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessDevicesPacketLossByNetworkQueryParams{}

		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsWirelessDevicesPacketLossByNetwork.NetworkIDs)
		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessDevicesPacketLossByNetwork.Serials)
		queryParams1.SSIDs = elementsToStrings(ctx, organizationsWirelessDevicesPacketLossByNetwork.SSIDs)
		queryParams1.Bands = elementsToStrings(ctx, organizationsWirelessDevicesPacketLossByNetwork.Bands)
		queryParams1.PerPage = int(organizationsWirelessDevicesPacketLossByNetwork.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessDevicesPacketLossByNetwork.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessDevicesPacketLossByNetwork.EndingBefore.ValueString()
		queryParams1.T0 = organizationsWirelessDevicesPacketLossByNetwork.T0.ValueString()
		queryParams1.T1 = organizationsWirelessDevicesPacketLossByNetwork.T1.ValueString()
		queryParams1.Timespan = organizationsWirelessDevicesPacketLossByNetwork.Timespan.ValueFloat64()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetOrganizationWirelessDevicesPacketLossByNetwork(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessDevicesPacketLossByNetwork",
				err.Error(),
			)
			return
		}

		organizationsWirelessDevicesPacketLossByNetwork = ResponseWirelessGetOrganizationWirelessDevicesPacketLossByNetworkItemsToBody(organizationsWirelessDevicesPacketLossByNetwork, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessDevicesPacketLossByNetwork)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessDevicesPacketLossByNetwork struct {
	OrganizationID types.String                                                             `tfsdk:"organization_id"`
	NetworkIDs     types.List                                                               `tfsdk:"network_ids"`
	Serials        types.List                                                               `tfsdk:"serials"`
	SSIDs          types.List                                                               `tfsdk:"ssids"`
	Bands          types.List                                                               `tfsdk:"bands"`
	PerPage        types.Int64                                                              `tfsdk:"per_page"`
	StartingAfter  types.String                                                             `tfsdk:"starting_after"`
	EndingBefore   types.String                                                             `tfsdk:"ending_before"`
	T0             types.String                                                             `tfsdk:"t0"`
	T1             types.String                                                             `tfsdk:"t1"`
	Timespan       types.Float64                                                            `tfsdk:"timespan"`
	Items          *[]ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetwork `tfsdk:"items"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetwork struct {
	Downstream *ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetworkDownstream `tfsdk:"downstream"`
	Network    *ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetworkNetwork    `tfsdk:"network"`
	Upstream   *ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetworkUpstream   `tfsdk:"upstream"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetworkDownstream struct {
	LossPercentage types.Float64 `tfsdk:"loss_percentage"`
	Lost           types.Int64   `tfsdk:"lost"`
	Total          types.Int64   `tfsdk:"total"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetworkNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetworkUpstream struct {
	LossPercentage types.Float64 `tfsdk:"loss_percentage"`
	Lost           types.Int64   `tfsdk:"lost"`
	Total          types.Int64   `tfsdk:"total"`
}

// ToBody
func ResponseWirelessGetOrganizationWirelessDevicesPacketLossByNetworkItemsToBody(state OrganizationsWirelessDevicesPacketLossByNetwork, response *merakigosdk.ResponseWirelessGetOrganizationWirelessDevicesPacketLossByNetwork) OrganizationsWirelessDevicesPacketLossByNetwork {
	var items []ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetwork
	for _, item := range *response {
		itemState := ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetwork{
			Downstream: func() *ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetworkDownstream {
				if item.Downstream != nil {
					return &ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetworkDownstream{
						LossPercentage: func() types.Float64 {
							if item.Downstream.LossPercentage != nil {
								return types.Float64Value(float64(*item.Downstream.LossPercentage))
							}
							return types.Float64{}
						}(),
						Lost: func() types.Int64 {
							if item.Downstream.Lost != nil {
								return types.Int64Value(int64(*item.Downstream.Lost))
							}
							return types.Int64{}
						}(),
						Total: func() types.Int64 {
							if item.Downstream.Total != nil {
								return types.Int64Value(int64(*item.Downstream.Total))
							}
							return types.Int64{}
						}(),
					}
				}
				return nil
			}(),
			Network: func() *ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetworkNetwork {
				if item.Network != nil {
					return &ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetworkNetwork{
						ID:   types.StringValue(item.Network.ID),
						Name: types.StringValue(item.Network.Name),
					}
				}
				return nil
			}(),
			Upstream: func() *ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetworkUpstream {
				if item.Upstream != nil {
					return &ResponseItemWirelessGetOrganizationWirelessDevicesPacketLossByNetworkUpstream{
						LossPercentage: func() types.Float64 {
							if item.Upstream.LossPercentage != nil {
								return types.Float64Value(float64(*item.Upstream.LossPercentage))
							}
							return types.Float64{}
						}(),
						Lost: func() types.Int64 {
							if item.Upstream.Lost != nil {
								return types.Int64Value(int64(*item.Upstream.Lost))
							}
							return types.Int64{}
						}(),
						Total: func() types.Int64 {
							if item.Upstream.Total != nil {
								return types.Int64Value(int64(*item.Upstream.Total))
							}
							return types.Int64{}
						}(),
					}
				}
				return nil
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
