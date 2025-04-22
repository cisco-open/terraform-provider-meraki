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
	_ datasource.DataSource              = &OrganizationsIntegrationsXdrNetworksDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsIntegrationsXdrNetworksDataSource{}
)

func NewOrganizationsIntegrationsXdrNetworksDataSource() datasource.DataSource {
	return &OrganizationsIntegrationsXdrNetworksDataSource{}
}

type OrganizationsIntegrationsXdrNetworksDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsIntegrationsXdrNetworksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsIntegrationsXdrNetworksDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_integrations_xdr_networks"
}

func (d *OrganizationsIntegrationsXdrNetworksDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter the results by network IDs`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 100. Default is 20.`,
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
						MarkdownDescription: `List of networks with XDR enabled`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"enabled": schema.BoolAttribute{
									MarkdownDescription: `Represents whether XDR is enabled for the network`,
									Computed:            true,
								},
								"is_eligible": schema.BoolAttribute{
									MarkdownDescription: `Represents whether the network is eligible for XDR`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `The name of the network`,
									Computed:            true,
								},
								"network_id": schema.StringAttribute{
									MarkdownDescription: `Network ID`,
									Computed:            true,
								},
								"product_types": schema.ListAttribute{
									MarkdownDescription: `List of products that have XDR enabled`,
									Computed:            true,
									ElementType:         types.StringType,
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
										MarkdownDescription: `Counts relating to the paginated networks`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"remaining": schema.Int64Attribute{
												MarkdownDescription: `The number of networks in the dataset that are available on subsequent pages`,
												Computed:            true,
											},
											"total": schema.Int64Attribute{
												MarkdownDescription: `The total number of networks in the dataset`,
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

func (d *OrganizationsIntegrationsXdrNetworksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsIntegrationsXdrNetworks OrganizationsIntegrationsXdrNetworks
	diags := req.Config.Get(ctx, &organizationsIntegrationsXdrNetworks)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationIntegrationsXdrNetworks")
		vvOrganizationID := organizationsIntegrationsXdrNetworks.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationIntegrationsXdrNetworksQueryParams{}

		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsIntegrationsXdrNetworks.NetworkIDs)
		queryParams1.PerPage = int(organizationsIntegrationsXdrNetworks.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsIntegrationsXdrNetworks.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsIntegrationsXdrNetworks.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationIntegrationsXdrNetworks(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationIntegrationsXdrNetworks",
				err.Error(),
			)
			return
		}

		organizationsIntegrationsXdrNetworks = ResponseOrganizationsGetOrganizationIntegrationsXdrNetworksItemToBody(organizationsIntegrationsXdrNetworks, response1)
		diags = resp.State.Set(ctx, &organizationsIntegrationsXdrNetworks)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsIntegrationsXdrNetworks struct {
	OrganizationID types.String                                                 `tfsdk:"organization_id"`
	NetworkIDs     types.List                                                   `tfsdk:"network_ids"`
	PerPage        types.Int64                                                  `tfsdk:"per_page"`
	StartingAfter  types.String                                                 `tfsdk:"starting_after"`
	EndingBefore   types.String                                                 `tfsdk:"ending_before"`
	Item           *ResponseOrganizationsGetOrganizationIntegrationsXdrNetworks `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationIntegrationsXdrNetworks struct {
	Items *[]ResponseOrganizationsGetOrganizationIntegrationsXdrNetworksItems `tfsdk:"items"`
	Meta  *ResponseOrganizationsGetOrganizationIntegrationsXdrNetworksMeta    `tfsdk:"meta"`
}

type ResponseOrganizationsGetOrganizationIntegrationsXdrNetworksItems struct {
	Enabled      types.Bool   `tfsdk:"enabled"`
	IsEligible   types.Bool   `tfsdk:"is_eligible"`
	Name         types.String `tfsdk:"name"`
	NetworkID    types.String `tfsdk:"network_id"`
	ProductTypes types.List   `tfsdk:"product_types"`
}

type ResponseOrganizationsGetOrganizationIntegrationsXdrNetworksMeta struct {
	Counts *ResponseOrganizationsGetOrganizationIntegrationsXdrNetworksMetaCounts `tfsdk:"counts"`
}

type ResponseOrganizationsGetOrganizationIntegrationsXdrNetworksMetaCounts struct {
	Items *ResponseOrganizationsGetOrganizationIntegrationsXdrNetworksMetaCountsItems `tfsdk:"items"`
}

type ResponseOrganizationsGetOrganizationIntegrationsXdrNetworksMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseOrganizationsGetOrganizationIntegrationsXdrNetworksItemToBody(state OrganizationsIntegrationsXdrNetworks, response *merakigosdk.ResponseOrganizationsGetOrganizationIntegrationsXdrNetworks) OrganizationsIntegrationsXdrNetworks {
	itemState := ResponseOrganizationsGetOrganizationIntegrationsXdrNetworks{
		Items: func() *[]ResponseOrganizationsGetOrganizationIntegrationsXdrNetworksItems {
			if response.Items != nil {
				result := make([]ResponseOrganizationsGetOrganizationIntegrationsXdrNetworksItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseOrganizationsGetOrganizationIntegrationsXdrNetworksItems{
						Enabled: func() types.Bool {
							if items.Enabled != nil {
								return types.BoolValue(*items.Enabled)
							}
							return types.Bool{}
						}(),
						IsEligible: func() types.Bool {
							if items.IsEligible != nil {
								return types.BoolValue(*items.IsEligible)
							}
							return types.Bool{}
						}(),
						Name:         types.StringValue(items.Name),
						NetworkID:    types.StringValue(items.NetworkID),
						ProductTypes: StringSliceToList(items.ProductTypes),
					}
				}
				return &result
			}
			return nil
		}(),
		Meta: func() *ResponseOrganizationsGetOrganizationIntegrationsXdrNetworksMeta {
			if response.Meta != nil {
				return &ResponseOrganizationsGetOrganizationIntegrationsXdrNetworksMeta{
					Counts: func() *ResponseOrganizationsGetOrganizationIntegrationsXdrNetworksMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseOrganizationsGetOrganizationIntegrationsXdrNetworksMetaCounts{
								Items: func() *ResponseOrganizationsGetOrganizationIntegrationsXdrNetworksMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseOrganizationsGetOrganizationIntegrationsXdrNetworksMetaCountsItems{
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
