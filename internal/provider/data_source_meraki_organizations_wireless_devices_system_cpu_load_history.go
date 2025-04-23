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
	_ datasource.DataSource              = &OrganizationsWirelessDevicesSystemCPULoadHistoryDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessDevicesSystemCPULoadHistoryDataSource{}
)

func NewOrganizationsWirelessDevicesSystemCPULoadHistoryDataSource() datasource.DataSource {
	return &OrganizationsWirelessDevicesSystemCPULoadHistoryDataSource{}
}

type OrganizationsWirelessDevicesSystemCPULoadHistoryDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessDevicesSystemCPULoadHistoryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessDevicesSystemCPULoadHistoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_devices_system_cpu_load_history"
}

func (d *OrganizationsWirelessDevicesSystemCPULoadHistoryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
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
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 20. Default is 10.`,
				Optional:            true,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Optional parameter to filter device availabilities history by device serial numbers`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 1 day from today.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 1 day after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 1 day. The default is 1 day.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"items": schema.ListNestedAttribute{
						MarkdownDescription: `The top-level property containing all cpu load data.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"cpu_count": schema.Int64Attribute{
									MarkdownDescription: `Number of CPU cores on the device`,
									Computed:            true,
								},
								"mac": schema.StringAttribute{
									MarkdownDescription: `MAC address of the device`,
									Computed:            true,
								},
								"model": schema.StringAttribute{
									MarkdownDescription: `Model of the device`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Name of the device`,
									Computed:            true,
								},
								"network": schema.SingleNestedAttribute{
									MarkdownDescription: `Information regarding the network the device belongs to`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `The network ID`,
											Computed:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: `The name of the network`,
											Computed:            true,
										},
										"tags": schema.ListAttribute{
											MarkdownDescription: `List of custom tags for the network`,
											Computed:            true,
											ElementType:         types.StringType,
										},
									},
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `Unique serial number for the device`,
									Computed:            true,
								},
								"series": schema.SetNestedAttribute{
									MarkdownDescription: `Series of cpu load average measurements on the device`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"cpu_load5": schema.Int64Attribute{
												MarkdownDescription: `The 5 minutes cpu load average of the device`,
												Computed:            true,
											},
											"ts": schema.StringAttribute{
												MarkdownDescription: `Timestamp of the cpu load measurement`,
												Computed:            true,
											},
										},
									},
								},
								"tags": schema.ListAttribute{
									MarkdownDescription: `List of custom tags for the device`,
									Computed:            true,
									ElementType:         types.StringType,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsWirelessDevicesSystemCPULoadHistoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessDevicesSystemCPULoadHistory OrganizationsWirelessDevicesSystemCPULoadHistory
	diags := req.Config.Get(ctx, &organizationsWirelessDevicesSystemCPULoadHistory)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessDevicesSystemCPULoadHistory")
		vvOrganizationID := organizationsWirelessDevicesSystemCPULoadHistory.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessDevicesSystemCPULoadHistoryQueryParams{}

		queryParams1.T0 = organizationsWirelessDevicesSystemCPULoadHistory.T0.ValueString()
		queryParams1.T1 = organizationsWirelessDevicesSystemCPULoadHistory.T1.ValueString()
		queryParams1.Timespan = organizationsWirelessDevicesSystemCPULoadHistory.Timespan.ValueFloat64()
		queryParams1.PerPage = int(organizationsWirelessDevicesSystemCPULoadHistory.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessDevicesSystemCPULoadHistory.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessDevicesSystemCPULoadHistory.EndingBefore.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsWirelessDevicesSystemCPULoadHistory.NetworkIDs)
		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessDevicesSystemCPULoadHistory.Serials)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetOrganizationWirelessDevicesSystemCPULoadHistory(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessDevicesSystemCPULoadHistory",
				err.Error(),
			)
			return
		}

		organizationsWirelessDevicesSystemCPULoadHistory = ResponseWirelessGetOrganizationWirelessDevicesSystemCPULoadHistoryItemToBody(organizationsWirelessDevicesSystemCPULoadHistory, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessDevicesSystemCPULoadHistory)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessDevicesSystemCPULoadHistory struct {
	OrganizationID types.String                                                        `tfsdk:"organization_id"`
	T0             types.String                                                        `tfsdk:"t0"`
	T1             types.String                                                        `tfsdk:"t1"`
	Timespan       types.Float64                                                       `tfsdk:"timespan"`
	PerPage        types.Int64                                                         `tfsdk:"per_page"`
	StartingAfter  types.String                                                        `tfsdk:"starting_after"`
	EndingBefore   types.String                                                        `tfsdk:"ending_before"`
	NetworkIDs     types.List                                                          `tfsdk:"network_ids"`
	Serials        types.List                                                          `tfsdk:"serials"`
	Item           *ResponseWirelessGetOrganizationWirelessDevicesSystemCpuLoadHistory `tfsdk:"item"`
}

type ResponseWirelessGetOrganizationWirelessDevicesSystemCpuLoadHistory struct {
	Items *[]ResponseWirelessGetOrganizationWirelessDevicesSystemCpuLoadHistoryItems `tfsdk:"items"`
}

type ResponseWirelessGetOrganizationWirelessDevicesSystemCpuLoadHistoryItems struct {
	CPUCount types.Int64                                                                      `tfsdk:"cpu_count"`
	Mac      types.String                                                                     `tfsdk:"mac"`
	Model    types.String                                                                     `tfsdk:"model"`
	Name     types.String                                                                     `tfsdk:"name"`
	Network  *ResponseWirelessGetOrganizationWirelessDevicesSystemCpuLoadHistoryItemsNetwork  `tfsdk:"network"`
	Serial   types.String                                                                     `tfsdk:"serial"`
	Series   *[]ResponseWirelessGetOrganizationWirelessDevicesSystemCpuLoadHistoryItemsSeries `tfsdk:"series"`
	Tags     types.List                                                                       `tfsdk:"tags"`
}

type ResponseWirelessGetOrganizationWirelessDevicesSystemCpuLoadHistoryItemsNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Tags types.List   `tfsdk:"tags"`
}

type ResponseWirelessGetOrganizationWirelessDevicesSystemCpuLoadHistoryItemsSeries struct {
	CPULoad5 types.Int64  `tfsdk:"cpu_load5"`
	Ts       types.String `tfsdk:"ts"`
}

// ToBody
func ResponseWirelessGetOrganizationWirelessDevicesSystemCPULoadHistoryItemToBody(state OrganizationsWirelessDevicesSystemCPULoadHistory, response *merakigosdk.ResponseWirelessGetOrganizationWirelessDevicesSystemCPULoadHistory) OrganizationsWirelessDevicesSystemCPULoadHistory {
	itemState := ResponseWirelessGetOrganizationWirelessDevicesSystemCpuLoadHistory{
		Items: func() *[]ResponseWirelessGetOrganizationWirelessDevicesSystemCpuLoadHistoryItems {
			if response.Items != nil {
				result := make([]ResponseWirelessGetOrganizationWirelessDevicesSystemCpuLoadHistoryItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseWirelessGetOrganizationWirelessDevicesSystemCpuLoadHistoryItems{
						CPUCount: func() types.Int64 {
							if items.CPUCount != nil {
								return types.Int64Value(int64(*items.CPUCount))
							}
							return types.Int64{}
						}(),
						Mac:   types.StringValue(items.Mac),
						Model: types.StringValue(items.Model),
						Name:  types.StringValue(items.Name),
						Network: func() *ResponseWirelessGetOrganizationWirelessDevicesSystemCpuLoadHistoryItemsNetwork {
							if items.Network != nil {
								return &ResponseWirelessGetOrganizationWirelessDevicesSystemCpuLoadHistoryItemsNetwork{
									ID:   types.StringValue(items.Network.ID),
									Name: types.StringValue(items.Network.Name),
									Tags: StringSliceToList(items.Network.Tags),
								}
							}
							return nil
						}(),
						Serial: types.StringValue(items.Serial),
						Series: func() *[]ResponseWirelessGetOrganizationWirelessDevicesSystemCpuLoadHistoryItemsSeries {
							if items.Series != nil {
								result := make([]ResponseWirelessGetOrganizationWirelessDevicesSystemCpuLoadHistoryItemsSeries, len(*items.Series))
								for i, series := range *items.Series {
									result[i] = ResponseWirelessGetOrganizationWirelessDevicesSystemCpuLoadHistoryItemsSeries{
										CPULoad5: func() types.Int64 {
											if series.CPULoad5 != nil {
												return types.Int64Value(int64(*series.CPULoad5))
											}
											return types.Int64{}
										}(),
										Ts: types.StringValue(series.Ts),
									}
								}
								return &result
							}
							return nil
						}(),
						Tags: StringSliceToList(items.Tags),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
