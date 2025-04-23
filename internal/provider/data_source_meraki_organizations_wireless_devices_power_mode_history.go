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
	_ datasource.DataSource              = &OrganizationsWirelessDevicesPowerModeHistoryDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessDevicesPowerModeHistoryDataSource{}
)

func NewOrganizationsWirelessDevicesPowerModeHistoryDataSource() datasource.DataSource {
	return &OrganizationsWirelessDevicesPowerModeHistoryDataSource{}
}

type OrganizationsWirelessDevicesPowerModeHistoryDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessDevicesPowerModeHistoryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessDevicesPowerModeHistoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_devices_power_mode_history"
}

func (d *OrganizationsWirelessDevicesPowerModeHistoryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
						MarkdownDescription: `The top-level property containing all power mode data.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"events": schema.SetNestedAttribute{
									MarkdownDescription: `Events indicating power mode changes for the device`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"power_mode": schema.StringAttribute{
												MarkdownDescription: `The power mode of the device`,
												Computed:            true,
											},
											"ts": schema.StringAttribute{
												MarkdownDescription: `Timestamp of the event`,
												Computed:            true,
											},
										},
									},
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

func (d *OrganizationsWirelessDevicesPowerModeHistoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessDevicesPowerModeHistory OrganizationsWirelessDevicesPowerModeHistory
	diags := req.Config.Get(ctx, &organizationsWirelessDevicesPowerModeHistory)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessDevicesPowerModeHistory")
		vvOrganizationID := organizationsWirelessDevicesPowerModeHistory.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessDevicesPowerModeHistoryQueryParams{}

		queryParams1.T0 = organizationsWirelessDevicesPowerModeHistory.T0.ValueString()
		queryParams1.T1 = organizationsWirelessDevicesPowerModeHistory.T1.ValueString()
		queryParams1.Timespan = organizationsWirelessDevicesPowerModeHistory.Timespan.ValueFloat64()
		queryParams1.PerPage = int(organizationsWirelessDevicesPowerModeHistory.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessDevicesPowerModeHistory.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessDevicesPowerModeHistory.EndingBefore.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsWirelessDevicesPowerModeHistory.NetworkIDs)
		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessDevicesPowerModeHistory.Serials)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetOrganizationWirelessDevicesPowerModeHistory(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessDevicesPowerModeHistory",
				err.Error(),
			)
			return
		}

		organizationsWirelessDevicesPowerModeHistory = ResponseWirelessGetOrganizationWirelessDevicesPowerModeHistoryItemToBody(organizationsWirelessDevicesPowerModeHistory, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessDevicesPowerModeHistory)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessDevicesPowerModeHistory struct {
	OrganizationID types.String                                                    `tfsdk:"organization_id"`
	T0             types.String                                                    `tfsdk:"t0"`
	T1             types.String                                                    `tfsdk:"t1"`
	Timespan       types.Float64                                                   `tfsdk:"timespan"`
	PerPage        types.Int64                                                     `tfsdk:"per_page"`
	StartingAfter  types.String                                                    `tfsdk:"starting_after"`
	EndingBefore   types.String                                                    `tfsdk:"ending_before"`
	NetworkIDs     types.List                                                      `tfsdk:"network_ids"`
	Serials        types.List                                                      `tfsdk:"serials"`
	Item           *ResponseWirelessGetOrganizationWirelessDevicesPowerModeHistory `tfsdk:"item"`
}

type ResponseWirelessGetOrganizationWirelessDevicesPowerModeHistory struct {
	Items *[]ResponseWirelessGetOrganizationWirelessDevicesPowerModeHistoryItems `tfsdk:"items"`
}

type ResponseWirelessGetOrganizationWirelessDevicesPowerModeHistoryItems struct {
	Events  *[]ResponseWirelessGetOrganizationWirelessDevicesPowerModeHistoryItemsEvents `tfsdk:"events"`
	Mac     types.String                                                                 `tfsdk:"mac"`
	Model   types.String                                                                 `tfsdk:"model"`
	Name    types.String                                                                 `tfsdk:"name"`
	Network *ResponseWirelessGetOrganizationWirelessDevicesPowerModeHistoryItemsNetwork  `tfsdk:"network"`
	Serial  types.String                                                                 `tfsdk:"serial"`
	Tags    types.List                                                                   `tfsdk:"tags"`
}

type ResponseWirelessGetOrganizationWirelessDevicesPowerModeHistoryItemsEvents struct {
	PowerMode types.String `tfsdk:"power_mode"`
	Ts        types.String `tfsdk:"ts"`
}

type ResponseWirelessGetOrganizationWirelessDevicesPowerModeHistoryItemsNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Tags types.List   `tfsdk:"tags"`
}

// ToBody
func ResponseWirelessGetOrganizationWirelessDevicesPowerModeHistoryItemToBody(state OrganizationsWirelessDevicesPowerModeHistory, response *merakigosdk.ResponseWirelessGetOrganizationWirelessDevicesPowerModeHistory) OrganizationsWirelessDevicesPowerModeHistory {
	itemState := ResponseWirelessGetOrganizationWirelessDevicesPowerModeHistory{
		Items: func() *[]ResponseWirelessGetOrganizationWirelessDevicesPowerModeHistoryItems {
			if response.Items != nil {
				result := make([]ResponseWirelessGetOrganizationWirelessDevicesPowerModeHistoryItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseWirelessGetOrganizationWirelessDevicesPowerModeHistoryItems{
						Events: func() *[]ResponseWirelessGetOrganizationWirelessDevicesPowerModeHistoryItemsEvents {
							if items.Events != nil {
								result := make([]ResponseWirelessGetOrganizationWirelessDevicesPowerModeHistoryItemsEvents, len(*items.Events))
								for i, events := range *items.Events {
									result[i] = ResponseWirelessGetOrganizationWirelessDevicesPowerModeHistoryItemsEvents{
										PowerMode: types.StringValue(events.PowerMode),
										Ts:        types.StringValue(events.Ts),
									}
								}
								return &result
							}
							return nil
						}(),
						Mac:   types.StringValue(items.Mac),
						Model: types.StringValue(items.Model),
						Name:  types.StringValue(items.Name),
						Network: func() *ResponseWirelessGetOrganizationWirelessDevicesPowerModeHistoryItemsNetwork {
							if items.Network != nil {
								return &ResponseWirelessGetOrganizationWirelessDevicesPowerModeHistoryItemsNetwork{
									ID:   types.StringValue(items.Network.ID),
									Name: types.StringValue(items.Network.Name),
									Tags: StringSliceToList(items.Network.Tags),
								}
							}
							return nil
						}(),
						Serial: types.StringValue(items.Serial),
						Tags:   StringSliceToList(items.Tags),
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
