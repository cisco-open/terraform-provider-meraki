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
	_ datasource.DataSource              = &OrganizationsWirelessDevicesChannelUtilizationByNetworkDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessDevicesChannelUtilizationByNetworkDataSource{}
)

func NewOrganizationsWirelessDevicesChannelUtilizationByNetworkDataSource() datasource.DataSource {
	return &OrganizationsWirelessDevicesChannelUtilizationByNetworkDataSource{}
}

type OrganizationsWirelessDevicesChannelUtilizationByNetworkDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessDevicesChannelUtilizationByNetworkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessDevicesChannelUtilizationByNetworkDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_devices_channel_utilization_by_network"
}

func (d *OrganizationsWirelessDevicesChannelUtilizationByNetworkDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"interval": schema.Int64Attribute{
				MarkdownDescription: `interval query parameter. The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600.`,
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
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 90 days. The default is 7 days.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetwork`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"by_band": schema.SetNestedAttribute{
							MarkdownDescription: `Channel utilization broken down by band.`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"band": schema.StringAttribute{
										MarkdownDescription: `The band for the given metrics.`,
										Computed:            true,
									},
									"non_wifi": schema.SingleNestedAttribute{
										MarkdownDescription: `An object containing non-wifi utilization.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"percentage": schema.Float64Attribute{
												MarkdownDescription: `Percentage of non-wifi channel utiliation for the given band.`,
												Computed:            true,
											},
										},
									},
									"total": schema.SingleNestedAttribute{
										MarkdownDescription: `An object containing total channel utilization.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"percentage": schema.Float64Attribute{
												MarkdownDescription: `Percentage of total channel utiliation for the given band.`,
												Computed:            true,
											},
										},
									},
									"wifi": schema.SingleNestedAttribute{
										MarkdownDescription: `An object containing wifi utilization.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"percentage": schema.Float64Attribute{
												MarkdownDescription: `Percentage of wifi channel utiliation for the given band.`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
						"network": schema.SingleNestedAttribute{
							MarkdownDescription: `Network for the given utilization metrics.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `Network ID of the given utilization metrics.`,
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

func (d *OrganizationsWirelessDevicesChannelUtilizationByNetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessDevicesChannelUtilizationByNetwork OrganizationsWirelessDevicesChannelUtilizationByNetwork
	diags := req.Config.Get(ctx, &organizationsWirelessDevicesChannelUtilizationByNetwork)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessDevicesChannelUtilizationByNetwork")
		vvOrganizationID := organizationsWirelessDevicesChannelUtilizationByNetwork.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessDevicesChannelUtilizationByNetworkQueryParams{}

		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsWirelessDevicesChannelUtilizationByNetwork.NetworkIDs)
		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessDevicesChannelUtilizationByNetwork.Serials)
		queryParams1.PerPage = int(organizationsWirelessDevicesChannelUtilizationByNetwork.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessDevicesChannelUtilizationByNetwork.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessDevicesChannelUtilizationByNetwork.EndingBefore.ValueString()
		queryParams1.T0 = organizationsWirelessDevicesChannelUtilizationByNetwork.T0.ValueString()
		queryParams1.T1 = organizationsWirelessDevicesChannelUtilizationByNetwork.T1.ValueString()
		queryParams1.Timespan = organizationsWirelessDevicesChannelUtilizationByNetwork.Timespan.ValueFloat64()
		queryParams1.Interval = int(organizationsWirelessDevicesChannelUtilizationByNetwork.Interval.ValueInt64())

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetOrganizationWirelessDevicesChannelUtilizationByNetwork(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessDevicesChannelUtilizationByNetwork",
				err.Error(),
			)
			return
		}

		organizationsWirelessDevicesChannelUtilizationByNetwork = ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkItemsToBody(organizationsWirelessDevicesChannelUtilizationByNetwork, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessDevicesChannelUtilizationByNetwork)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessDevicesChannelUtilizationByNetwork struct {
	OrganizationID types.String                                                                     `tfsdk:"organization_id"`
	NetworkIDs     types.List                                                                       `tfsdk:"network_ids"`
	Serials        types.List                                                                       `tfsdk:"serials"`
	PerPage        types.Int64                                                                      `tfsdk:"per_page"`
	StartingAfter  types.String                                                                     `tfsdk:"starting_after"`
	EndingBefore   types.String                                                                     `tfsdk:"ending_before"`
	T0             types.String                                                                     `tfsdk:"t0"`
	T1             types.String                                                                     `tfsdk:"t1"`
	Timespan       types.Float64                                                                    `tfsdk:"timespan"`
	Interval       types.Int64                                                                      `tfsdk:"interval"`
	Items          *[]ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetwork `tfsdk:"items"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetwork struct {
	ByBand  *[]ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBand `tfsdk:"by_band"`
	Network *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkNetwork  `tfsdk:"network"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBand struct {
	Band    types.String                                                                                `tfsdk:"band"`
	NonWifi *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandNonWifi `tfsdk:"non_wifi"`
	Total   *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandTotal   `tfsdk:"total"`
	Wifi    *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandWifi    `tfsdk:"wifi"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandNonWifi struct {
	Percentage types.Float64 `tfsdk:"percentage"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandTotal struct {
	Percentage types.Float64 `tfsdk:"percentage"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandWifi struct {
	Percentage types.Float64 `tfsdk:"percentage"`
}

type ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkNetwork struct {
	ID types.String `tfsdk:"id"`
}

// ToBody
func ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkItemsToBody(state OrganizationsWirelessDevicesChannelUtilizationByNetwork, response *merakigosdk.ResponseWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetwork) OrganizationsWirelessDevicesChannelUtilizationByNetwork {
	var items []ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetwork
	for _, item := range *response {
		itemState := ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetwork{
			ByBand: func() *[]ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBand {
				if item.ByBand != nil {
					result := make([]ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBand, len(*item.ByBand))
					for i, byBand := range *item.ByBand {
						result[i] = ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBand{
							Band: types.StringValue(byBand.Band),
							NonWifi: func() *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandNonWifi {
								if byBand.NonWifi != nil {
									return &ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandNonWifi{
										Percentage: func() types.Float64 {
											if byBand.NonWifi.Percentage != nil {
												return types.Float64Value(float64(*byBand.NonWifi.Percentage))
											}
											return types.Float64{}
										}(),
									}
								}
								return nil
							}(),
							Total: func() *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandTotal {
								if byBand.Total != nil {
									return &ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandTotal{
										Percentage: func() types.Float64 {
											if byBand.Total.Percentage != nil {
												return types.Float64Value(float64(*byBand.Total.Percentage))
											}
											return types.Float64{}
										}(),
									}
								}
								return nil
							}(),
							Wifi: func() *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandWifi {
								if byBand.Wifi != nil {
									return &ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkByBandWifi{
										Percentage: func() types.Float64 {
											if byBand.Wifi.Percentage != nil {
												return types.Float64Value(float64(*byBand.Wifi.Percentage))
											}
											return types.Float64{}
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
			Network: func() *ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkNetwork {
				if item.Network != nil {
					return &ResponseItemWirelessGetOrganizationWirelessDevicesChannelUtilizationByNetworkNetwork{
						ID: types.StringValue(item.Network.ID),
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
