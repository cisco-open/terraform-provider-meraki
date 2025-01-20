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
	_ datasource.DataSource              = &OrganizationsFloorPlansAutoLocateDevicesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsFloorPlansAutoLocateDevicesDataSource{}
)

func NewOrganizationsFloorPlansAutoLocateDevicesDataSource() datasource.DataSource {
	return &OrganizationsFloorPlansAutoLocateDevicesDataSource{}
}

type OrganizationsFloorPlansAutoLocateDevicesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsFloorPlansAutoLocateDevicesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsFloorPlansAutoLocateDevicesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_floor_plans_auto_locate_devices"
}

func (d *OrganizationsFloorPlansAutoLocateDevicesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"floor_plan_ids": schema.ListAttribute{
				MarkdownDescription: `floorPlanIds query parameter. Optional parameter to filter devices by one or more floorplan IDs`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter devices by one or more network IDs`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 10000. Default is 1000.`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationFloorPlansAutoLocateDevices`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"items": schema.ListNestedAttribute{
							MarkdownDescription: `Items in the paginated dataset`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"auto_locate": schema.SingleNestedAttribute{
										MarkdownDescription: `The auto locate position for this device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"lat": schema.Float64Attribute{
												MarkdownDescription: `Latitude`,
												Computed:            true,
											},
											"lng": schema.Float64Attribute{
												MarkdownDescription: `Longitude`,
												Computed:            true,
											},
										},
									},
									"floor_plan": schema.SingleNestedAttribute{
										MarkdownDescription: `The assigned floor plan for this device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"id": schema.StringAttribute{
												MarkdownDescription: `Floor plan ID`,
												Computed:            true,
											},
											"status": schema.StringAttribute{
												MarkdownDescription: `Floor plan name`,
												Computed:            true,
											},
										},
									},
									"is_anchor": schema.BoolAttribute{
										MarkdownDescription: `Whether or not this auto locate position is an anchor`,
										Computed:            true,
									},
									"lat": schema.Float64Attribute{
										MarkdownDescription: `Latitude`,
										Computed:            true,
									},
									"lng": schema.Float64Attribute{
										MarkdownDescription: `Longitude`,
										Computed:            true,
									},
									"mac": schema.StringAttribute{
										MarkdownDescription: `MAC Address`,
										Computed:            true,
									},
									"model": schema.StringAttribute{
										MarkdownDescription: `Model`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `Device Name`,
										Computed:            true,
									},
									"network": schema.SingleNestedAttribute{
										MarkdownDescription: `Network info`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"id": schema.StringAttribute{
												MarkdownDescription: `ID for the network containing this device`,
												Computed:            true,
											},
										},
									},
									"serial": schema.StringAttribute{
										MarkdownDescription: `Device Serial Number`,
										Computed:            true,
									},
									"status": schema.StringAttribute{
										MarkdownDescription: `Device Status`,
										Computed:            true,
									},
									"tags": schema.ListAttribute{
										MarkdownDescription: `Tags`,
										Computed:            true,
										ElementType:         types.StringType,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: `The type of auto locate position. Possible values: 'user', 'gnss', and 'calculated'`,
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
		},
	}
}

func (d *OrganizationsFloorPlansAutoLocateDevicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsFloorPlansAutoLocateDevices OrganizationsFloorPlansAutoLocateDevices
	diags := req.Config.Get(ctx, &organizationsFloorPlansAutoLocateDevices)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationFloorPlansAutoLocateDevices")
		vvOrganizationID := organizationsFloorPlansAutoLocateDevices.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationFloorPlansAutoLocateDevicesQueryParams{}

		queryParams1.PerPage = int(organizationsFloorPlansAutoLocateDevices.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsFloorPlansAutoLocateDevices.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsFloorPlansAutoLocateDevices.EndingBefore.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsFloorPlansAutoLocateDevices.NetworkIDs)
		queryParams1.FloorPlanIDs = elementsToStrings(ctx, organizationsFloorPlansAutoLocateDevices.FloorPlanIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationFloorPlansAutoLocateDevices(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationFloorPlansAutoLocateDevices",
				err.Error(),
			)
			return
		}

		organizationsFloorPlansAutoLocateDevices = ResponseOrganizationsGetOrganizationFloorPlansAutoLocateDevicesItemsToBody(organizationsFloorPlansAutoLocateDevices, response1)
		diags = resp.State.Set(ctx, &organizationsFloorPlansAutoLocateDevices)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsFloorPlansAutoLocateDevices struct {
	OrganizationID types.String                                                           `tfsdk:"organization_id"`
	PerPage        types.Int64                                                            `tfsdk:"per_page"`
	StartingAfter  types.String                                                           `tfsdk:"starting_after"`
	EndingBefore   types.String                                                           `tfsdk:"ending_before"`
	NetworkIDs     types.List                                                             `tfsdk:"network_ids"`
	FloorPlanIDs   types.List                                                             `tfsdk:"floor_plan_ids"`
	Items          *[]ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevices `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevices struct {
	Items *[]ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesItems `tfsdk:"items"`
	Meta  *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesMeta    `tfsdk:"meta"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesItems struct {
	AutoLocate *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesItemsAutoLocate `tfsdk:"auto_locate"`
	FloorPlan  *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesItemsFloorPlan  `tfsdk:"floor_plan"`
	IsAnchor   types.Bool                                                                          `tfsdk:"is_anchor"`
	Lat        types.Float64                                                                       `tfsdk:"lat"`
	Lng        types.Float64                                                                       `tfsdk:"lng"`
	Mac        types.String                                                                        `tfsdk:"mac"`
	Model      types.String                                                                        `tfsdk:"model"`
	Name       types.String                                                                        `tfsdk:"name"`
	Network    *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesItemsNetwork    `tfsdk:"network"`
	Serial     types.String                                                                        `tfsdk:"serial"`
	Status     types.String                                                                        `tfsdk:"status"`
	Tags       types.List                                                                          `tfsdk:"tags"`
	Type       types.String                                                                        `tfsdk:"type"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesItemsAutoLocate struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesItemsFloorPlan struct {
	ID     types.String `tfsdk:"id"`
	Status types.String `tfsdk:"status"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesItemsNetwork struct {
	ID types.String `tfsdk:"id"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesMeta struct {
	Counts *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesMetaCounts `tfsdk:"counts"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesMetaCounts struct {
	Items *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesMetaCountsItems `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseOrganizationsGetOrganizationFloorPlansAutoLocateDevicesItemsToBody(state OrganizationsFloorPlansAutoLocateDevices, response *merakigosdk.ResponseOrganizationsGetOrganizationFloorPlansAutoLocateDevices) OrganizationsFloorPlansAutoLocateDevices {
	var items []ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevices
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevices{
			Items: func() *[]ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesItems {
				if item.Items != nil {
					result := make([]ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesItems, len(*item.Items))
					for i, items := range *item.Items {
						result[i] = ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesItems{
							AutoLocate: func() *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesItemsAutoLocate {
								if items.AutoLocate != nil {
									return &ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesItemsAutoLocate{
										Lat: func() types.Float64 {
											if items.AutoLocate.Lat != nil {
												return types.Float64Value(float64(*items.AutoLocate.Lat))
											}
											return types.Float64{}
										}(),
										Lng: func() types.Float64 {
											if items.AutoLocate.Lng != nil {
												return types.Float64Value(float64(*items.AutoLocate.Lng))
											}
											return types.Float64{}
										}(),
									}
								}
								return nil
							}(),
							FloorPlan: func() *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesItemsFloorPlan {
								if items.FloorPlan != nil {
									return &ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesItemsFloorPlan{
										ID:     types.StringValue(items.FloorPlan.ID),
										Status: types.StringValue(items.FloorPlan.Status),
									}
								}
								return nil
							}(),
							IsAnchor: func() types.Bool {
								if items.IsAnchor != nil {
									return types.BoolValue(*items.IsAnchor)
								}
								return types.Bool{}
							}(),
							Lat: func() types.Float64 {
								if items.Lat != nil {
									return types.Float64Value(float64(*items.Lat))
								}
								return types.Float64{}
							}(),
							Lng: func() types.Float64 {
								if items.Lng != nil {
									return types.Float64Value(float64(*items.Lng))
								}
								return types.Float64{}
							}(),
							Mac:   types.StringValue(items.Mac),
							Model: types.StringValue(items.Model),
							Name:  types.StringValue(items.Name),
							Network: func() *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesItemsNetwork {
								if items.Network != nil {
									return &ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesItemsNetwork{
										ID: types.StringValue(items.Network.ID),
									}
								}
								return nil
							}(),
							Serial: types.StringValue(items.Serial),
							Status: types.StringValue(items.Status),
							Tags:   StringSliceToList(items.Tags),
							Type:   types.StringValue(items.Type),
						}
					}
					return &result
				}
				return nil
			}(),
			Meta: func() *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesMeta {
				if item.Meta != nil {
					return &ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesMeta{
						Counts: func() *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesMetaCounts {
							if item.Meta.Counts != nil {
								return &ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesMetaCounts{
									Items: func() *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesMetaCountsItems {
										if item.Meta.Counts.Items != nil {
											return &ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateDevicesMetaCountsItems{
												Remaining: func() types.Int64 {
													if item.Meta.Counts.Items.Remaining != nil {
														return types.Int64Value(int64(*item.Meta.Counts.Items.Remaining))
													}
													return types.Int64{}
												}(),
												Total: func() types.Int64 {
													if item.Meta.Counts.Items.Total != nil {
														return types.Int64Value(int64(*item.Meta.Counts.Items.Total))
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
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
