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
	_ datasource.DataSource              = &OrganizationsApplianceFirewallMulticastForwardingByNetworkDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsApplianceFirewallMulticastForwardingByNetworkDataSource{}
)

func NewOrganizationsApplianceFirewallMulticastForwardingByNetworkDataSource() datasource.DataSource {
	return &OrganizationsApplianceFirewallMulticastForwardingByNetworkDataSource{}
}

type OrganizationsApplianceFirewallMulticastForwardingByNetworkDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsApplianceFirewallMulticastForwardingByNetworkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsApplianceFirewallMulticastForwardingByNetworkDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_appliance_firewall_multicast_forwarding_by_network"
}

func (d *OrganizationsApplianceFirewallMulticastForwardingByNetworkDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
						MarkdownDescription: `List of networks with multicast static forwarding rules`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"network": schema.SingleNestedAttribute{
									MarkdownDescription: `Network details`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `ID of the network whose multicast forwarding settings are returned.`,
											Computed:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: `Name of the network whose multicast forwarding settings are returned.`,
											Computed:            true,
										},
									},
								},
								"rules": schema.SetNestedAttribute{
									MarkdownDescription: `Static multicast forwarding rules.`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"address": schema.StringAttribute{
												MarkdownDescription: `IP address`,
												Computed:            true,
											},
											"description": schema.StringAttribute{
												MarkdownDescription: `Forwarding rule description.`,
												Computed:            true,
											},
											"vlan_ids": schema.ListAttribute{
												MarkdownDescription: `List of VLAN IDs`,
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
	}
}

func (d *OrganizationsApplianceFirewallMulticastForwardingByNetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsApplianceFirewallMulticastForwardingByNetwork OrganizationsApplianceFirewallMulticastForwardingByNetwork
	diags := req.Config.Get(ctx, &organizationsApplianceFirewallMulticastForwardingByNetwork)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationApplianceFirewallMulticastForwardingByNetwork")
		vvOrganizationID := organizationsApplianceFirewallMulticastForwardingByNetwork.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationApplianceFirewallMulticastForwardingByNetworkQueryParams{}

		queryParams1.PerPage = int(organizationsApplianceFirewallMulticastForwardingByNetwork.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsApplianceFirewallMulticastForwardingByNetwork.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsApplianceFirewallMulticastForwardingByNetwork.EndingBefore.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsApplianceFirewallMulticastForwardingByNetwork.NetworkIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetOrganizationApplianceFirewallMulticastForwardingByNetwork(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationApplianceFirewallMulticastForwardingByNetwork",
				err.Error(),
			)
			return
		}

		organizationsApplianceFirewallMulticastForwardingByNetwork = ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItemToBody(organizationsApplianceFirewallMulticastForwardingByNetwork, response1)
		diags = resp.State.Set(ctx, &organizationsApplianceFirewallMulticastForwardingByNetwork)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsApplianceFirewallMulticastForwardingByNetwork struct {
	OrganizationID types.String                                                                   `tfsdk:"organization_id"`
	PerPage        types.Int64                                                                    `tfsdk:"per_page"`
	StartingAfter  types.String                                                                   `tfsdk:"starting_after"`
	EndingBefore   types.String                                                                   `tfsdk:"ending_before"`
	NetworkIDs     types.List                                                                     `tfsdk:"network_ids"`
	Item           *ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetwork `tfsdk:"item"`
}

type ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetwork struct {
	Items *[]ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItems `tfsdk:"items"`
	Meta  *ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkMeta    `tfsdk:"meta"`
}

type ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItems struct {
	Network *ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItemsNetwork `tfsdk:"network"`
	Rules   *[]ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItemsRules `tfsdk:"rules"`
}

type ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItemsNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItemsRules struct {
	Address     types.String `tfsdk:"address"`
	Description types.String `tfsdk:"description"`
	VLANIDs     types.List   `tfsdk:"vlan_ids"`
}

type ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkMeta struct {
	Counts *ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkMetaCounts `tfsdk:"counts"`
}

type ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkMetaCounts struct {
	Items *ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkMetaCountsItems `tfsdk:"items"`
}

type ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItemToBody(state OrganizationsApplianceFirewallMulticastForwardingByNetwork, response *merakigosdk.ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetwork) OrganizationsApplianceFirewallMulticastForwardingByNetwork {
	itemState := ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetwork{
		Items: func() *[]ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItems {
			if response.Items != nil {
				result := make([]ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItems{
						Network: func() *ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItemsNetwork {
							if items.Network != nil {
								return &ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItemsNetwork{
									ID:   types.StringValue(items.Network.ID),
									Name: types.StringValue(items.Network.Name),
								}
							}
							return nil
						}(),
						Rules: func() *[]ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItemsRules {
							if items.Rules != nil {
								result := make([]ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItemsRules, len(*items.Rules))
								for i, rules := range *items.Rules {
									result[i] = ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkItemsRules{
										Address:     types.StringValue(rules.Address),
										Description: types.StringValue(rules.Description),
										VLANIDs:     StringSliceToList(rules.VLANIDs),
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
		Meta: func() *ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkMeta {
			if response.Meta != nil {
				return &ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkMeta{
					Counts: func() *ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkMetaCounts{
								Items: func() *ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseApplianceGetOrganizationApplianceFirewallMulticastForwardingByNetworkMetaCountsItems{
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
