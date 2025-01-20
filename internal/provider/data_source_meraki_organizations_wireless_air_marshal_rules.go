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
	_ datasource.DataSource              = &OrganizationsWirelessAirMarshalRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessAirMarshalRulesDataSource{}
)

func NewOrganizationsWirelessAirMarshalRulesDataSource() datasource.DataSource {
	return &OrganizationsWirelessAirMarshalRulesDataSource{}
}

type OrganizationsWirelessAirMarshalRulesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessAirMarshalRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessAirMarshalRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_air_marshal_rules"
}

func (d *OrganizationsWirelessAirMarshalRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. (optional) The set of network IDs to include.`,
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
						MarkdownDescription: `List of rules`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"created_at": schema.StringAttribute{
									MarkdownDescription: `Created at timestamp`,
									Computed:            true,
								},
								"match": schema.SingleNestedAttribute{
									MarkdownDescription: `Indicates whether or not clients are allowed to        connect to rogue SSIDs by default. (blocked by default)`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"string": schema.StringAttribute{
											MarkdownDescription: `Indicates whether or not clients are allowed to        connect to rogue SSIDs by default. (blocked by default)`,
											Computed:            true,
										},
										"type": schema.StringAttribute{
											MarkdownDescription: `Indicates whether or not clients are allowed to        connect to rogue SSIDs by default. (blocked by default)`,
											Computed:            true,
										},
									},
								},
								"network": schema.SingleNestedAttribute{
									MarkdownDescription: `Network details`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `Network ID`,
											Computed:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: `Network name`,
											Computed:            true,
										},
									},
								},
								"rule_id": schema.StringAttribute{
									MarkdownDescription: `Indicates whether or not clients are allowed to        connect to rogue SSIDs by default. (blocked by default)`,
									Computed:            true,
								},
								"type": schema.StringAttribute{
									MarkdownDescription: `Indicates whether or not clients are allowed to        connect to rogue SSIDs by default. (blocked by default)`,
									Computed:            true,
								},
								"updated_at": schema.StringAttribute{
									MarkdownDescription: `Updated at timestamp`,
									Computed:            true,
								},
							},
						},
					},
					"meta": schema.SingleNestedAttribute{
						MarkdownDescription: `Meta details about the result`,
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

											"total": schema.Int64Attribute{
												MarkdownDescription: `Count of rules`,
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

func (d *OrganizationsWirelessAirMarshalRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessAirMarshalRules OrganizationsWirelessAirMarshalRules
	diags := req.Config.Get(ctx, &organizationsWirelessAirMarshalRules)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessAirMarshalRules")
		vvOrganizationID := organizationsWirelessAirMarshalRules.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessAirMarshalRulesQueryParams{}

		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsWirelessAirMarshalRules.NetworkIDs)
		queryParams1.PerPage = int(organizationsWirelessAirMarshalRules.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessAirMarshalRules.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessAirMarshalRules.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetOrganizationWirelessAirMarshalRules(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessAirMarshalRules",
				err.Error(),
			)
			return
		}

		organizationsWirelessAirMarshalRules = ResponseWirelessGetOrganizationWirelessAirMarshalRulesItemToBody(organizationsWirelessAirMarshalRules, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessAirMarshalRules)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessAirMarshalRules struct {
	OrganizationID types.String                                            `tfsdk:"organization_id"`
	NetworkIDs     types.List                                              `tfsdk:"network_ids"`
	PerPage        types.Int64                                             `tfsdk:"per_page"`
	StartingAfter  types.String                                            `tfsdk:"starting_after"`
	EndingBefore   types.String                                            `tfsdk:"ending_before"`
	Item           *ResponseWirelessGetOrganizationWirelessAirMarshalRules `tfsdk:"item"`
}

type ResponseWirelessGetOrganizationWirelessAirMarshalRules struct {
	Items *[]ResponseWirelessGetOrganizationWirelessAirMarshalRulesItems `tfsdk:"items"`
	Meta  *ResponseWirelessGetOrganizationWirelessAirMarshalRulesMeta    `tfsdk:"meta"`
}

type ResponseWirelessGetOrganizationWirelessAirMarshalRulesItems struct {
	CreatedAt types.String                                                        `tfsdk:"created_at"`
	Match     *ResponseWirelessGetOrganizationWirelessAirMarshalRulesItemsMatch   `tfsdk:"match"`
	Network   *ResponseWirelessGetOrganizationWirelessAirMarshalRulesItemsNetwork `tfsdk:"network"`
	RuleID    types.String                                                        `tfsdk:"rule_id"`
	Type      types.String                                                        `tfsdk:"type"`
	UpdatedAt types.String                                                        `tfsdk:"updated_at"`
}

type ResponseWirelessGetOrganizationWirelessAirMarshalRulesItemsMatch struct {
	String types.String `tfsdk:"string"`
	Type   types.String `tfsdk:"type"`
}

type ResponseWirelessGetOrganizationWirelessAirMarshalRulesItemsNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseWirelessGetOrganizationWirelessAirMarshalRulesMeta struct {
	Counts *ResponseWirelessGetOrganizationWirelessAirMarshalRulesMetaCounts `tfsdk:"counts"`
}

type ResponseWirelessGetOrganizationWirelessAirMarshalRulesMetaCounts struct {
	Items *ResponseWirelessGetOrganizationWirelessAirMarshalRulesMetaCountsItems `tfsdk:"items"`
}

type ResponseWirelessGetOrganizationWirelessAirMarshalRulesMetaCountsItems struct {
	Total types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseWirelessGetOrganizationWirelessAirMarshalRulesItemToBody(state OrganizationsWirelessAirMarshalRules, response *merakigosdk.ResponseWirelessGetOrganizationWirelessAirMarshalRules) OrganizationsWirelessAirMarshalRules {
	itemState := ResponseWirelessGetOrganizationWirelessAirMarshalRules{
		Items: func() *[]ResponseWirelessGetOrganizationWirelessAirMarshalRulesItems {
			if response.Items != nil {
				result := make([]ResponseWirelessGetOrganizationWirelessAirMarshalRulesItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseWirelessGetOrganizationWirelessAirMarshalRulesItems{
						CreatedAt: types.StringValue(items.CreatedAt),
						Match: func() *ResponseWirelessGetOrganizationWirelessAirMarshalRulesItemsMatch {
							if items.Match != nil {
								return &ResponseWirelessGetOrganizationWirelessAirMarshalRulesItemsMatch{
									String: types.StringValue(items.Match.String),
									Type:   types.StringValue(items.Match.Type),
								}
							}
							return nil
						}(),
						Network: func() *ResponseWirelessGetOrganizationWirelessAirMarshalRulesItemsNetwork {
							if items.Network != nil {
								return &ResponseWirelessGetOrganizationWirelessAirMarshalRulesItemsNetwork{
									ID:   types.StringValue(items.Network.ID),
									Name: types.StringValue(items.Network.Name),
								}
							}
							return nil
						}(),
						RuleID:    types.StringValue(items.RuleID),
						Type:      types.StringValue(items.Type),
						UpdatedAt: types.StringValue(items.UpdatedAt),
					}
				}
				return &result
			}
			return nil
		}(),
		Meta: func() *ResponseWirelessGetOrganizationWirelessAirMarshalRulesMeta {
			if response.Meta != nil {
				return &ResponseWirelessGetOrganizationWirelessAirMarshalRulesMeta{
					Counts: func() *ResponseWirelessGetOrganizationWirelessAirMarshalRulesMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseWirelessGetOrganizationWirelessAirMarshalRulesMetaCounts{
								Items: func() *ResponseWirelessGetOrganizationWirelessAirMarshalRulesMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseWirelessGetOrganizationWirelessAirMarshalRulesMetaCountsItems{
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
