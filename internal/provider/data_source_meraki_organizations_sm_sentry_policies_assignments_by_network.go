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
	_ datasource.DataSource              = &OrganizationsSmSentryPoliciesAssignmentsByNetworkDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSmSentryPoliciesAssignmentsByNetworkDataSource{}
)

func NewOrganizationsSmSentryPoliciesAssignmentsByNetworkDataSource() datasource.DataSource {
	return &OrganizationsSmSentryPoliciesAssignmentsByNetworkDataSource{}
}

type OrganizationsSmSentryPoliciesAssignmentsByNetworkDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSmSentryPoliciesAssignmentsByNetworkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSmSentryPoliciesAssignmentsByNetworkDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_sm_sentry_policies_assignments_by_network"
}

func (d *OrganizationsSmSentryPoliciesAssignmentsByNetworkDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter Sentry Policies by Network Id`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 50.`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSmGetOrganizationSmSentryPoliciesAssignmentsByNetwork`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"items": schema.ListNestedAttribute{
							MarkdownDescription: `Sentry Group Policies for the Organization keyed by the Network or Locale Id the Policy belongs to`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"network_id": schema.StringAttribute{
										MarkdownDescription: `The Id of the Network`,
										Computed:            true,
									},
									"policies": schema.SetNestedAttribute{
										MarkdownDescription: `Array of Sentry Group Policies for the Network`,
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{

												"created_at": schema.StringAttribute{
													MarkdownDescription: `The creation time of the Sentry Policy`,
													Computed:            true,
												},
												"group_number": schema.StringAttribute{
													MarkdownDescription: `The number of the Group Policy`,
													Computed:            true,
												},
												"group_policy_id": schema.StringAttribute{
													MarkdownDescription: `The Id of the Group Policy. This is associated with the network specified by the networkId.`,
													Computed:            true,
												},
												"last_updated_at": schema.StringAttribute{
													MarkdownDescription: `The last update time of the Sentry Policy`,
													Computed:            true,
												},
												"network_id": schema.StringAttribute{
													MarkdownDescription: `The Id of the Network the Sentry Policy is associated with. In a locale, this should be the Wireless Group if present, otherwise the Wired Group.`,
													Computed:            true,
												},
												"policy_id": schema.StringAttribute{
													MarkdownDescription: `The Id of the Sentry Policy`,
													Computed:            true,
												},
												"priority": schema.StringAttribute{
													MarkdownDescription: `The priority of the Sentry Policy`,
													Computed:            true,
												},
												"scope": schema.StringAttribute{
													MarkdownDescription: `The scope of the Sentry Policy`,
													Computed:            true,
												},
												"sm_network_id": schema.StringAttribute{
													MarkdownDescription: `The Id of the Systems Manager Network the Sentry Policy is assigned to`,
													Computed:            true,
												},
												"tags": schema.ListAttribute{
													MarkdownDescription: `The tags of the Sentry Policy`,
													Computed:            true,
													ElementType:         types.StringType,
												},
											},
										},
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

func (d *OrganizationsSmSentryPoliciesAssignmentsByNetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSmSentryPoliciesAssignmentsByNetwork OrganizationsSmSentryPoliciesAssignmentsByNetwork
	diags := req.Config.Get(ctx, &organizationsSmSentryPoliciesAssignmentsByNetwork)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSmSentryPoliciesAssignmentsByNetwork")
		vvOrganizationID := organizationsSmSentryPoliciesAssignmentsByNetwork.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationSmSentryPoliciesAssignmentsByNetworkQueryParams{}

		queryParams1.PerPage = int(organizationsSmSentryPoliciesAssignmentsByNetwork.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsSmSentryPoliciesAssignmentsByNetwork.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsSmSentryPoliciesAssignmentsByNetwork.EndingBefore.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsSmSentryPoliciesAssignmentsByNetwork.NetworkIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Sm.GetOrganizationSmSentryPoliciesAssignmentsByNetwork(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSmSentryPoliciesAssignmentsByNetwork",
				err.Error(),
			)
			return
		}

		organizationsSmSentryPoliciesAssignmentsByNetwork = ResponseSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkItemsToBody(organizationsSmSentryPoliciesAssignmentsByNetwork, response1)
		diags = resp.State.Set(ctx, &organizationsSmSentryPoliciesAssignmentsByNetwork)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSmSentryPoliciesAssignmentsByNetwork struct {
	OrganizationID types.String                                                         `tfsdk:"organization_id"`
	PerPage        types.Int64                                                          `tfsdk:"per_page"`
	StartingAfter  types.String                                                         `tfsdk:"starting_after"`
	EndingBefore   types.String                                                         `tfsdk:"ending_before"`
	NetworkIDs     types.List                                                           `tfsdk:"network_ids"`
	Items          *[]ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetwork `tfsdk:"items"`
}

type ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetwork struct {
	Items *[]ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkItems `tfsdk:"items"`
	Meta  *ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkMeta    `tfsdk:"meta"`
}

type ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkItems struct {
	NetworkID types.String                                                                      `tfsdk:"network_id"`
	Policies  *[]ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkItemsPolicies `tfsdk:"policies"`
}

type ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkItemsPolicies struct {
	CreatedAt     types.String `tfsdk:"created_at"`
	GroupNumber   types.String `tfsdk:"group_number"`
	GroupPolicyID types.String `tfsdk:"group_policy_id"`
	LastUpdatedAt types.String `tfsdk:"last_updated_at"`
	NetworkID     types.String `tfsdk:"network_id"`
	PolicyID      types.String `tfsdk:"policy_id"`
	Priority      types.String `tfsdk:"priority"`
	Scope         types.String `tfsdk:"scope"`
	SmNetworkID   types.String `tfsdk:"sm_network_id"`
	Tags          types.List   `tfsdk:"tags"`
}

type ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkMeta struct {
	Counts *ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkMetaCounts `tfsdk:"counts"`
}

type ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkMetaCounts struct {
	Items *ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkMetaCountsItems `tfsdk:"items"`
}

type ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkItemsToBody(state OrganizationsSmSentryPoliciesAssignmentsByNetwork, response *merakigosdk.ResponseSmGetOrganizationSmSentryPoliciesAssignmentsByNetwork) OrganizationsSmSentryPoliciesAssignmentsByNetwork {
	var items []ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetwork
	for _, item := range *response {
		itemState := ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetwork{
			Items: func() *[]ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkItems {
				if item.Items != nil {
					result := make([]ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkItems, len(*item.Items))
					for i, items := range *item.Items {
						result[i] = ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkItems{
							NetworkID: types.StringValue(items.NetworkID),
							Policies: func() *[]ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkItemsPolicies {
								if items.Policies != nil {
									result := make([]ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkItemsPolicies, len(*items.Policies))
									for i, policies := range *items.Policies {
										result[i] = ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkItemsPolicies{
											CreatedAt:     types.StringValue(policies.CreatedAt),
											GroupNumber:   types.StringValue(policies.GroupNumber),
											GroupPolicyID: types.StringValue(policies.GroupPolicyID),
											LastUpdatedAt: types.StringValue(policies.LastUpdatedAt),
											NetworkID:     types.StringValue(policies.NetworkID),
											PolicyID:      types.StringValue(policies.PolicyID),
											Priority:      types.StringValue(policies.Priority),
											Scope:         types.StringValue(policies.Scope),
											SmNetworkID:   types.StringValue(policies.SmNetworkID),
											Tags:          StringSliceToList(policies.Tags),
										}
									}
									return &result
								}
								return nil
							}(),
						}
					}
					return &result
				}
				return nil
			}(),
			Meta: func() *ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkMeta {
				if item.Meta != nil {
					return &ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkMeta{
						Counts: func() *ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkMetaCounts {
							if item.Meta.Counts != nil {
								return &ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkMetaCounts{
									Items: func() *ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkMetaCountsItems {
										if item.Meta.Counts.Items != nil {
											return &ResponseItemSmGetOrganizationSmSentryPoliciesAssignmentsByNetworkMetaCountsItems{
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
