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
	_ datasource.DataSource              = &OrganizationsWirelessControllerConnectionsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessControllerConnectionsDataSource{}
)

func NewOrganizationsWirelessControllerConnectionsDataSource() datasource.DataSource {
	return &OrganizationsWirelessControllerConnectionsDataSource{}
}

type OrganizationsWirelessControllerConnectionsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessControllerConnectionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessControllerConnectionsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_controller_connections"
}

func (d *OrganizationsWirelessControllerConnectionsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"controller_serials": schema.ListAttribute{
				MarkdownDescription: `controllerSerials query parameter. Optional parameter to filter access points by its controller cloud ID. This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter access points by network ID. This filter uses multiple exact matches.`,
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
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"items": schema.ListNestedAttribute{
						MarkdownDescription: `Access points associated with Wireless LAN controllers`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"controller": schema.SingleNestedAttribute{
									MarkdownDescription: `Associated wireless LAN controller`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"serial": schema.StringAttribute{
											MarkdownDescription: `Associated wireless LAN controller cloud ID`,
											Computed:            true,
										},
									},
								},
								"network": schema.SingleNestedAttribute{
									MarkdownDescription: `Access points network`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `Access points network ID`,
											Computed:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: `Access points network name`,
											Computed:            true,
										},
										"url": schema.StringAttribute{
											MarkdownDescription: `Access points network URL`,
											Computed:            true,
										},
									},
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `Access points cloud ID`,
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

func (d *OrganizationsWirelessControllerConnectionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessControllerConnections OrganizationsWirelessControllerConnections
	diags := req.Config.Get(ctx, &organizationsWirelessControllerConnections)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessControllerConnections")
		vvOrganizationID := organizationsWirelessControllerConnections.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessControllerConnectionsQueryParams{}

		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsWirelessControllerConnections.NetworkIDs)
		queryParams1.ControllerSerials = elementsToStrings(ctx, organizationsWirelessControllerConnections.ControllerSerials)
		queryParams1.PerPage = int(organizationsWirelessControllerConnections.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessControllerConnections.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessControllerConnections.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.WirelessController.GetOrganizationWirelessControllerConnections(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessControllerConnections",
				err.Error(),
			)
			return
		}

		organizationsWirelessControllerConnections = ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItemToBody(organizationsWirelessControllerConnections, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessControllerConnections)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessControllerConnections struct {
	OrganizationID    types.String                                                            `tfsdk:"organization_id"`
	NetworkIDs        types.List                                                              `tfsdk:"network_ids"`
	ControllerSerials types.List                                                              `tfsdk:"controller_serials"`
	PerPage           types.Int64                                                             `tfsdk:"per_page"`
	StartingAfter     types.String                                                            `tfsdk:"starting_after"`
	EndingBefore      types.String                                                            `tfsdk:"ending_before"`
	Item              *ResponseWirelessControllerGetOrganizationWirelessControllerConnections `tfsdk:"item"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerConnections struct {
	Items *[]ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItems `tfsdk:"items"`
	Meta  *ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsMeta    `tfsdk:"meta"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItems struct {
	Controller *ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItemsController `tfsdk:"controller"`
	Network    *ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItemsNetwork    `tfsdk:"network"`
	Serial     types.String                                                                           `tfsdk:"serial"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItemsController struct {
	Serial types.String `tfsdk:"serial"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItemsNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	URL  types.String `tfsdk:"url"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsMeta struct {
	Counts *ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsMetaCounts `tfsdk:"counts"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsMetaCounts struct {
	Items *ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsMetaCountsItems `tfsdk:"items"`
}

type ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItemToBody(state OrganizationsWirelessControllerConnections, response *merakigosdk.ResponseWirelessControllerGetOrganizationWirelessControllerConnections) OrganizationsWirelessControllerConnections {
	itemState := ResponseWirelessControllerGetOrganizationWirelessControllerConnections{
		Items: func() *[]ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItems {
			if response.Items != nil {
				result := make([]ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItems{
						Controller: func() *ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItemsController {
							if items.Controller != nil {
								return &ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItemsController{
									Serial: types.StringValue(items.Controller.Serial),
								}
							}
							return nil
						}(),
						Network: func() *ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItemsNetwork {
							if items.Network != nil {
								return &ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsItemsNetwork{
									ID:   types.StringValue(items.Network.ID),
									Name: types.StringValue(items.Network.Name),
									URL:  types.StringValue(items.Network.URL),
								}
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
		Meta: func() *ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsMeta {
			if response.Meta != nil {
				return &ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsMeta{
					Counts: func() *ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsMetaCounts{
								Items: func() *ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseWirelessControllerGetOrganizationWirelessControllerConnectionsMetaCountsItems{
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
