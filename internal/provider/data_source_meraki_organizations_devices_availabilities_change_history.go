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
	_ datasource.DataSource              = &OrganizationsDevicesAvailabilitiesChangeHistoryDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsDevicesAvailabilitiesChangeHistoryDataSource{}
)

func NewOrganizationsDevicesAvailabilitiesChangeHistoryDataSource() datasource.DataSource {
	return &OrganizationsDevicesAvailabilitiesChangeHistoryDataSource{}
}

type OrganizationsDevicesAvailabilitiesChangeHistoryDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsDevicesAvailabilitiesChangeHistoryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsDevicesAvailabilitiesChangeHistoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_devices_availabilities_change_history"
}

func (d *OrganizationsDevicesAvailabilitiesChangeHistoryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter device availabilities history by network IDs`,
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
			"product_types": schema.ListAttribute{
				MarkdownDescription: `productTypes query parameter. Optional parameter to filter device availabilities history by device product types`,
				Optional:            true,
				ElementType:         types.StringType,
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
			"statuses": schema.ListAttribute{
				MarkdownDescription: `statuses query parameter. Optional parameter to filter device availabilities history by device statuses`,
				Optional:            true,
				ElementType:         types.StringType,
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
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistory`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"details": schema.SingleNestedAttribute{
							MarkdownDescription: `Details about the status changes`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"new": schema.SetNestedAttribute{
									MarkdownDescription: `Details about the new status`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"name": schema.StringAttribute{
												MarkdownDescription: `Name of the detail`,
												Computed:            true,
											},
											"value": schema.StringAttribute{
												MarkdownDescription: `Value of the detail`,
												Computed:            true,
											},
										},
									},
								},
								"old": schema.SetNestedAttribute{
									MarkdownDescription: `Details about the old status`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"name": schema.StringAttribute{
												MarkdownDescription: `Name of the detail`,
												Computed:            true,
											},
											"value": schema.StringAttribute{
												MarkdownDescription: `Value of the detail`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
						"device": schema.SingleNestedAttribute{
							MarkdownDescription: `Device information`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"model": schema.StringAttribute{
									MarkdownDescription: `Device model`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Device name`,
									Computed:            true,
								},
								"product_type": schema.StringAttribute{
									MarkdownDescription: `Device product type.`,
									Computed:            true,
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `Device serial number`,
									Computed:            true,
								},
							},
						},
						"network": schema.SingleNestedAttribute{
							MarkdownDescription: `Network information`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `Network id`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Network name`,
									Computed:            true,
								},
								"tags": schema.ListAttribute{
									MarkdownDescription: `Network tags`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"url": schema.StringAttribute{
									MarkdownDescription: `Network dashboard url`,
									Computed:            true,
								},
							},
						},
						"ts": schema.StringAttribute{
							MarkdownDescription: `Timestamp, in iso8601 format, at which the event happened`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsDevicesAvailabilitiesChangeHistoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsDevicesAvailabilitiesChangeHistory OrganizationsDevicesAvailabilitiesChangeHistory
	diags := req.Config.Get(ctx, &organizationsDevicesAvailabilitiesChangeHistory)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationDevicesAvailabilitiesChangeHistory")
		vvOrganizationID := organizationsDevicesAvailabilitiesChangeHistory.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationDevicesAvailabilitiesChangeHistoryQueryParams{}

		queryParams1.PerPage = int(organizationsDevicesAvailabilitiesChangeHistory.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsDevicesAvailabilitiesChangeHistory.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsDevicesAvailabilitiesChangeHistory.EndingBefore.ValueString()
		queryParams1.T0 = organizationsDevicesAvailabilitiesChangeHistory.T0.ValueString()
		queryParams1.T1 = organizationsDevicesAvailabilitiesChangeHistory.T1.ValueString()
		queryParams1.Timespan = organizationsDevicesAvailabilitiesChangeHistory.Timespan.ValueFloat64()
		queryParams1.Serials = elementsToStrings(ctx, organizationsDevicesAvailabilitiesChangeHistory.Serials)
		queryParams1.ProductTypes = elementsToStrings(ctx, organizationsDevicesAvailabilitiesChangeHistory.ProductTypes)
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsDevicesAvailabilitiesChangeHistory.NetworkIDs)
		queryParams1.Statuses = elementsToStrings(ctx, organizationsDevicesAvailabilitiesChangeHistory.Statuses)

		response1, restyResp1, err := d.client.Organizations.GetOrganizationDevicesAvailabilitiesChangeHistory(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationDevicesAvailabilitiesChangeHistory",
				err.Error(),
			)
			return
		}

		organizationsDevicesAvailabilitiesChangeHistory = ResponseOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryItemsToBody(organizationsDevicesAvailabilitiesChangeHistory, response1)
		diags = resp.State.Set(ctx, &organizationsDevicesAvailabilitiesChangeHistory)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsDevicesAvailabilitiesChangeHistory struct {
	OrganizationID types.String                                                                  `tfsdk:"organization_id"`
	PerPage        types.Int64                                                                   `tfsdk:"per_page"`
	StartingAfter  types.String                                                                  `tfsdk:"starting_after"`
	EndingBefore   types.String                                                                  `tfsdk:"ending_before"`
	T0             types.String                                                                  `tfsdk:"t0"`
	T1             types.String                                                                  `tfsdk:"t1"`
	Timespan       types.Float64                                                                 `tfsdk:"timespan"`
	Serials        types.List                                                                    `tfsdk:"serials"`
	ProductTypes   types.List                                                                    `tfsdk:"product_types"`
	NetworkIDs     types.List                                                                    `tfsdk:"network_ids"`
	Statuses       types.List                                                                    `tfsdk:"statuses"`
	Items          *[]ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistory `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistory struct {
	Details *ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDetails `tfsdk:"details"`
	Device  *ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDevice  `tfsdk:"device"`
	Network *ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryNetwork `tfsdk:"network"`
	Ts      types.String                                                                       `tfsdk:"ts"`
}

type ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDetails struct {
	New *[]ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDetailsNew `tfsdk:"new"`
	Old *[]ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDetailsOld `tfsdk:"old"`
}

type ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDetailsNew struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDetailsOld struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDevice struct {
	Model       types.String `tfsdk:"model"`
	Name        types.String `tfsdk:"name"`
	ProductType types.String `tfsdk:"product_type"`
	Serial      types.String `tfsdk:"serial"`
}

type ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Tags types.List   `tfsdk:"tags"`
	URL  types.String `tfsdk:"url"`
}

// ToBody
func ResponseOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryItemsToBody(state OrganizationsDevicesAvailabilitiesChangeHistory, response *merakigosdk.ResponseOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistory) OrganizationsDevicesAvailabilitiesChangeHistory {
	var items []ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistory
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistory{
			Details: func() *ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDetails {
				if item.Details != nil {
					return &ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDetails{
						New: func() *[]ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDetailsNew {
							if item.Details.New != nil {
								result := make([]ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDetailsNew, len(*item.Details.New))
								for i, new := range *item.Details.New {
									result[i] = ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDetailsNew{
										Name:  types.StringValue(new.Name),
										Value: types.StringValue(new.Value),
									}
								}
								return &result
							}
							return &[]ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDetailsNew{}
						}(),
						Old: func() *[]ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDetailsOld {
							if item.Details.Old != nil {
								result := make([]ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDetailsOld, len(*item.Details.Old))
								for i, old := range *item.Details.Old {
									result[i] = ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDetailsOld{
										Name:  types.StringValue(old.Name),
										Value: types.StringValue(old.Value),
									}
								}
								return &result
							}
							return &[]ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDetailsOld{}
						}(),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDetails{}
			}(),
			Device: func() *ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDevice {
				if item.Device != nil {
					return &ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDevice{
						Model:       types.StringValue(item.Device.Model),
						Name:        types.StringValue(item.Device.Name),
						ProductType: types.StringValue(item.Device.ProductType),
						Serial:      types.StringValue(item.Device.Serial),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryDevice{}
			}(),
			Network: func() *ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryNetwork {
				if item.Network != nil {
					return &ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryNetwork{
						ID:   types.StringValue(item.Network.ID),
						Name: types.StringValue(item.Network.Name),
						Tags: StringSliceToList(item.Network.Tags),
						URL:  types.StringValue(item.Network.URL),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesChangeHistoryNetwork{}
			}(),
			Ts: types.StringValue(item.Ts),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
