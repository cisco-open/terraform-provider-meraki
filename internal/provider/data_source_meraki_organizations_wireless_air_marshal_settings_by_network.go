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
	_ datasource.DataSource              = &OrganizationsWirelessAirMarshalSettingsByNetworkDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessAirMarshalSettingsByNetworkDataSource{}
)

func NewOrganizationsWirelessAirMarshalSettingsByNetworkDataSource() datasource.DataSource {
	return &OrganizationsWirelessAirMarshalSettingsByNetworkDataSource{}
}

type OrganizationsWirelessAirMarshalSettingsByNetworkDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessAirMarshalSettingsByNetworkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessAirMarshalSettingsByNetworkDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_air_marshal_settings_by_network"
}

func (d *OrganizationsWirelessAirMarshalSettingsByNetworkDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. The network IDs to include in the result set.`,
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
						MarkdownDescription: `List of settings`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"default_policy": schema.StringAttribute{
									MarkdownDescription: `Indicates whether or not clients are allowed to       connect to rogue SSIDs. (blocked by default)`,
									Computed:            true,
								},
								"network_id": schema.StringAttribute{
									MarkdownDescription: `The network ID`,
									Computed:            true,
								},
							},
						},
					},
					"meta": schema.SingleNestedAttribute{
						MarkdownDescription: `Metadata`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"counts": schema.SingleNestedAttribute{
								MarkdownDescription: `Counts`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"items": schema.SingleNestedAttribute{
										MarkdownDescription: `Items`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"remaining": schema.Int64Attribute{
												MarkdownDescription: `Remaining number of items`,
												Computed:            true,
											},
											"total": schema.Int64Attribute{
												MarkdownDescription: `Total number of items`,
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

func (d *OrganizationsWirelessAirMarshalSettingsByNetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessAirMarshalSettingsByNetwork OrganizationsWirelessAirMarshalSettingsByNetwork
	diags := req.Config.Get(ctx, &organizationsWirelessAirMarshalSettingsByNetwork)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessAirMarshalSettingsByNetwork")
		vvOrganizationID := organizationsWirelessAirMarshalSettingsByNetwork.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessAirMarshalSettingsByNetworkQueryParams{}

		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsWirelessAirMarshalSettingsByNetwork.NetworkIDs)
		queryParams1.PerPage = int(organizationsWirelessAirMarshalSettingsByNetwork.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessAirMarshalSettingsByNetwork.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessAirMarshalSettingsByNetwork.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetOrganizationWirelessAirMarshalSettingsByNetwork(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessAirMarshalSettingsByNetwork",
				err.Error(),
			)
			return
		}

		organizationsWirelessAirMarshalSettingsByNetwork = ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetworkItemToBody(organizationsWirelessAirMarshalSettingsByNetwork, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessAirMarshalSettingsByNetwork)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessAirMarshalSettingsByNetwork struct {
	OrganizationID types.String                                                        `tfsdk:"organization_id"`
	NetworkIDs     types.List                                                          `tfsdk:"network_ids"`
	PerPage        types.Int64                                                         `tfsdk:"per_page"`
	StartingAfter  types.String                                                        `tfsdk:"starting_after"`
	EndingBefore   types.String                                                        `tfsdk:"ending_before"`
	Item           *ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetwork `tfsdk:"item"`
}

type ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetwork struct {
	Items *[]ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetworkItems `tfsdk:"items"`
	Meta  *ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetworkMeta    `tfsdk:"meta"`
}

type ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetworkItems struct {
	DefaultPolicy types.String `tfsdk:"default_policy"`
	NetworkID     types.String `tfsdk:"network_id"`
}

type ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetworkMeta struct {
	Counts *ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetworkMetaCounts `tfsdk:"counts"`
}

type ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetworkMetaCounts struct {
	Items *ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetworkMetaCountsItems `tfsdk:"items"`
}

type ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetworkMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetworkItemToBody(state OrganizationsWirelessAirMarshalSettingsByNetwork, response *merakigosdk.ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetwork) OrganizationsWirelessAirMarshalSettingsByNetwork {
	itemState := ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetwork{
		Items: func() *[]ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetworkItems {
			if response.Items != nil {
				result := make([]ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetworkItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetworkItems{
						DefaultPolicy: types.StringValue(items.DefaultPolicy),
						NetworkID:     types.StringValue(items.NetworkID),
					}
				}
				return &result
			}
			return nil
		}(),
		Meta: func() *ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetworkMeta {
			if response.Meta != nil {
				return &ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetworkMeta{
					Counts: func() *ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetworkMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetworkMetaCounts{
								Items: func() *ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetworkMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseWirelessGetOrganizationWirelessAirMarshalSettingsByNetworkMetaCountsItems{
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
