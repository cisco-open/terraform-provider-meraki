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
	_ datasource.DataSource              = &OrganizationsWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalDataSource{}
)

func NewOrganizationsWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalDataSource() datasource.DataSource {
	return &OrganizationsWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalDataSource{}
}

type OrganizationsWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_controller_devices_interfaces_l2_usage_history_by_interval"
}

func (d *OrganizationsWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
						MarkdownDescription: `Wireless LAN controller layer 2 interfaces usage`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"readings": schema.SetNestedAttribute{
									MarkdownDescription: `The usages of layer 2 interfaces of the wireless LAN controller. Usage is in bytes`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"mac": schema.StringAttribute{
												MarkdownDescription: `The MAC address of the wireless controller interface`,
												Computed:            true,
											},
											"name": schema.StringAttribute{
												MarkdownDescription: `The name of the wireless LAN controller interface`,
												Computed:            true,
											},
											"recv": schema.Int64Attribute{
												MarkdownDescription: `The volume of data, in bytes/sec, received by wireless controller interface`,
												Computed:            true,
											},
											"send": schema.Int64Attribute{
												MarkdownDescription: `The volume of data, in bytes/sec, transmitted by wireless controller interface`,
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

func (d *OrganizationsWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessControllerDevicesInterfacesL2UsageHistoryByInterval OrganizationsWirelessControllerDevicesInterfacesL2UsageHistoryByInterval
	diags := req.Config.Get(ctx, &organizationsWirelessControllerDevicesInterfacesL2UsageHistoryByInterval)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval")
		vvOrganizationID := organizationsWirelessControllerDevicesInterfacesL2UsageHistoryByInterval.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalQueryParams{}

		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessControllerDevicesInterfacesL2UsageHistoryByInterval.Serials)
		queryParams1.T0 = organizationsWirelessControllerDevicesInterfacesL2UsageHistoryByInterval.T0.ValueString()
		queryParams1.T1 = organizationsWirelessControllerDevicesInterfacesL2UsageHistoryByInterval.T1.ValueString()
		queryParams1.Timespan = organizationsWirelessControllerDevicesInterfacesL2UsageHistoryByInterval.Timespan.ValueFloat64()
		queryParams1.PerPage = int(organizationsWirelessControllerDevicesInterfacesL2UsageHistoryByInterval.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessControllerDevicesInterfacesL2UsageHistoryByInterval.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessControllerDevicesInterfacesL2UsageHistoryByInterval.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.WirelessController.GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval",
				err.Error(),
			)
			return
		}

		organizationsWirelessControllerDevicesInterfacesL2UsageHistoryByInterval = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItemToBody(organizationsWirelessControllerDevicesInterfacesL2UsageHistoryByInterval, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessControllerDevicesInterfacesL2UsageHistoryByInterval)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessControllerDevicesInterfacesL2UsageHistoryByInterval struct {
	OrganizationID types.String                                                                                          `tfsdk:"organization_id"`
	Serials        types.List                                                                                            `tfsdk:"serials"`
	T0             types.String                                                                                          `tfsdk:"t0"`
	T1             types.String                                                                                          `tfsdk:"t1"`
	Timespan       types.Float64                                                                                         `tfsdk:"timespan"`
	PerPage        types.Int64                                                                                           `tfsdk:"per_page"`
	StartingAfter  types.String                                                                                          `tfsdk:"starting_after"`
	EndingBefore   types.String                                                                                          `tfsdk:"ending_before"`
	Item           *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval `tfsdk:"item"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItems `tfsdk:"items"`
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMeta    `tfsdk:"meta"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItems struct {
	Readings *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItemsReadings `tfsdk:"readings"`
	Serial   types.String                                                                                                         `tfsdk:"serial"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItemsReadings struct {
	Mac  types.String `tfsdk:"mac"`
	Name types.String `tfsdk:"name"`
	Recv types.Int64  `tfsdk:"recv"`
	Send types.Int64  `tfsdk:"send"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMetaCounts `tfsdk:"counts"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMetaCountsItems `tfsdk:"items"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItemToBody(state OrganizationsWirelessControllerDevicesInterfacesL2UsageHistoryByInterval, response *merakigosdk.ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval) OrganizationsWirelessControllerDevicesInterfacesL2UsageHistoryByInterval {
	itemState := ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval{
		Items: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItems {
			if response.Items != nil {
				result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItems{
						Readings: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItemsReadings {
							if items.Readings != nil {
								result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItemsReadings, len(*items.Readings))
								for i, readings := range *items.Readings {
									result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItemsReadings{
										Mac:  types.StringValue(readings.Mac),
										Name: types.StringValue(readings.Name),
										Recv: func() types.Int64 {
											if readings.Recv != nil {
												return types.Int64Value(int64(*readings.Recv))
											}
											return types.Int64{}
										}(),
										Send: func() types.Int64 {
											if readings.Send != nil {
												return types.Int64Value(int64(*readings.Send))
											}
											return types.Int64{}
										}(),
									}
								}
								return &result
							}
							return &[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItemsReadings{}
						}(),
						Serial: types.StringValue(items.Serial),
					}
				}
				return &result
			}
			return &[]ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalItems{}
		}(),
		Meta: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMeta {
			if response.Meta != nil {
				return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMeta{
					Counts: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMetaCounts{
								Items: func() *ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMetaCountsItems{
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
									return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMetaCountsItems{}
								}(),
							}
						}
						return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMetaCounts{}
					}(),
				}
			}
			return &ResponseWirelessControllerGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalMeta{}
		}(),
	}
	state.Item = &itemState
	return state
}
