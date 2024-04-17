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

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsDevicesBootsHistoryDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsDevicesBootsHistoryDataSource{}
)

func NewOrganizationsDevicesBootsHistoryDataSource() datasource.DataSource {
	return &OrganizationsDevicesBootsHistoryDataSource{}
}

type OrganizationsDevicesBootsHistoryDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsDevicesBootsHistoryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsDevicesBootsHistoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_devices_boots_history"
}

func (d *OrganizationsDevicesBootsHistoryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"most_recent_per_device": schema.BoolAttribute{
				MarkdownDescription: `mostRecentPerDevice query parameter. If true, only the most recent boot for each device is returned.`,
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
				MarkdownDescription: `serials query parameter. Optional parameter to filter device by device serial numbers. This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"sort_order": schema.StringAttribute{
				MarkdownDescription: `sortOrder query parameter. Sorted order of entries. Order options are 'ascending' and 'descending'. Default is 'descending'.`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 730 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 730 days.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationDevicesBootsHistory`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"network": schema.SingleNestedAttribute{
							MarkdownDescription: `Device network`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `API-formatted network ID`,
									Computed:            true,
								},
							},
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `Device serial number`,
							Computed:            true,
						},
						"start": schema.SingleNestedAttribute{
							MarkdownDescription: `Device power up`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"booted_at": schema.StringAttribute{
									MarkdownDescription: `Indicates when the device booted`,
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

func (d *OrganizationsDevicesBootsHistoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsDevicesBootsHistory OrganizationsDevicesBootsHistory
	diags := req.Config.Get(ctx, &organizationsDevicesBootsHistory)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationDevicesBootsHistory")
		vvOrganizationID := organizationsDevicesBootsHistory.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationDevicesBootsHistoryQueryParams{}

		queryParams1.T0 = organizationsDevicesBootsHistory.T0.ValueString()
		queryParams1.T1 = organizationsDevicesBootsHistory.T1.ValueString()
		queryParams1.Timespan = organizationsDevicesBootsHistory.Timespan.ValueFloat64()
		queryParams1.Serials = elementsToStrings(ctx, organizationsDevicesBootsHistory.Serials)
		queryParams1.MostRecentPerDevice = organizationsDevicesBootsHistory.MostRecentPerDevice.ValueBool()
		queryParams1.PerPage = int(organizationsDevicesBootsHistory.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsDevicesBootsHistory.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsDevicesBootsHistory.EndingBefore.ValueString()
		queryParams1.SortOrder = organizationsDevicesBootsHistory.SortOrder.ValueString()

		response1, restyResp1, err := d.client.Organizations.GetOrganizationDevicesBootsHistory(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationDevicesBootsHistory",
				err.Error(),
			)
			return
		}

		organizationsDevicesBootsHistory = ResponseOrganizationsGetOrganizationDevicesBootsHistoryItemsToBody(organizationsDevicesBootsHistory, response1)
		diags = resp.State.Set(ctx, &organizationsDevicesBootsHistory)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsDevicesBootsHistory struct {
	OrganizationID      types.String                                                   `tfsdk:"organization_id"`
	T0                  types.String                                                   `tfsdk:"t0"`
	T1                  types.String                                                   `tfsdk:"t1"`
	Timespan            types.Float64                                                  `tfsdk:"timespan"`
	Serials             types.List                                                     `tfsdk:"serials"`
	MostRecentPerDevice types.Bool                                                     `tfsdk:"most_recent_per_device"`
	PerPage             types.Int64                                                    `tfsdk:"per_page"`
	StartingAfter       types.String                                                   `tfsdk:"starting_after"`
	EndingBefore        types.String                                                   `tfsdk:"ending_before"`
	SortOrder           types.String                                                   `tfsdk:"sort_order"`
	Items               *[]ResponseItemOrganizationsGetOrganizationDevicesBootsHistory `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationDevicesBootsHistory struct {
	Network *ResponseItemOrganizationsGetOrganizationDevicesBootsHistoryNetwork `tfsdk:"network"`
	Serial  types.String                                                        `tfsdk:"serial"`
	Start   *ResponseItemOrganizationsGetOrganizationDevicesBootsHistoryStart   `tfsdk:"start"`
}

type ResponseItemOrganizationsGetOrganizationDevicesBootsHistoryNetwork struct {
	ID types.String `tfsdk:"id"`
}

type ResponseItemOrganizationsGetOrganizationDevicesBootsHistoryStart struct {
	BootedAt types.String `tfsdk:"booted_at"`
}

// ToBody
func ResponseOrganizationsGetOrganizationDevicesBootsHistoryItemsToBody(state OrganizationsDevicesBootsHistory, response *merakigosdk.ResponseOrganizationsGetOrganizationDevicesBootsHistory) OrganizationsDevicesBootsHistory {
	var items []ResponseItemOrganizationsGetOrganizationDevicesBootsHistory
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationDevicesBootsHistory{
			Network: func() *ResponseItemOrganizationsGetOrganizationDevicesBootsHistoryNetwork {
				if item.Network != nil {
					return &ResponseItemOrganizationsGetOrganizationDevicesBootsHistoryNetwork{
						ID: types.StringValue(item.Network.ID),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationDevicesBootsHistoryNetwork{}
			}(),
			Serial: types.StringValue(item.Serial),
			Start: func() *ResponseItemOrganizationsGetOrganizationDevicesBootsHistoryStart {
				if item.Start != nil {
					return &ResponseItemOrganizationsGetOrganizationDevicesBootsHistoryStart{
						BootedAt: types.StringValue(item.Start.BootedAt),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationDevicesBootsHistoryStart{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
