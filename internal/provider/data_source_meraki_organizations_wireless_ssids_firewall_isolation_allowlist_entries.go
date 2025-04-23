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
	_ datasource.DataSource              = &OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDataSource{}
)

func NewOrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDataSource() datasource.DataSource {
	return &OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDataSource{}
}

type OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_ssids_firewall_isolation_allowlist_entries"
}

func (d *OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. networkIds array to filter out results`,
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
			"ssids": schema.ListAttribute{
				MarkdownDescription: `ssids query parameter. ssids number array to filter out results`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"items": schema.ListNestedAttribute{
						MarkdownDescription: `L2 isolation allowlist items`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"client": schema.SingleNestedAttribute{
									MarkdownDescription: `The client of allowlist`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"mac": schema.StringAttribute{
											MarkdownDescription: `L2 Isolation mac address`,
											Computed:            true,
										},
									},
								},
								"created_at": schema.StringAttribute{
									MarkdownDescription: `Created at timestamp for the adaptive policy group`,
									Computed:            true,
								},
								"description": schema.StringAttribute{
									MarkdownDescription: `The description of mac address`,
									Computed:            true,
								},
								"entry_id": schema.StringAttribute{
									MarkdownDescription: `The id of entry`,
									Computed:            true,
								},
								"last_updated_at": schema.StringAttribute{
									MarkdownDescription: `Updated at timestamp for the adaptive policy group`,
									Computed:            true,
								},
								"network": schema.SingleNestedAttribute{
									MarkdownDescription: `The network that allowlist SSID belongs to`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `The index of network`,
											Computed:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: `The name of network`,
											Computed:            true,
										},
									},
								},
								"ssid": schema.SingleNestedAttribute{
									MarkdownDescription: `The SSID that allowlist belongs to`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `The index of SSID`,
											Computed:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: `The name of SSID`,
											Computed:            true,
										},
										"number": schema.Int64Attribute{
											MarkdownDescription: `The number of SSID`,
											Computed:            true,
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

func (d *OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessSSIDsFirewallIsolationAllowlistEntries OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntries
	diags := req.Config.Get(ctx, &organizationsWirelessSSIDsFirewallIsolationAllowlistEntries)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessSSIDsFirewallIsolationAllowlistEntries")
		vvOrganizationID := organizationsWirelessSSIDsFirewallIsolationAllowlistEntries.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessSSIDsFirewallIsolationAllowlistEntriesQueryParams{}

		queryParams1.PerPage = int(organizationsWirelessSSIDsFirewallIsolationAllowlistEntries.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessSSIDsFirewallIsolationAllowlistEntries.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessSSIDsFirewallIsolationAllowlistEntries.EndingBefore.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsWirelessSSIDsFirewallIsolationAllowlistEntries.NetworkIDs)
		queryParams1.SSIDs = elementsToStrings(ctx, organizationsWirelessSSIDsFirewallIsolationAllowlistEntries.SSIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetOrganizationWirelessSSIDsFirewallIsolationAllowlistEntries(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessSSIDsFirewallIsolationAllowlistEntries",
				err.Error(),
			)
			return
		}

		organizationsWirelessSSIDsFirewallIsolationAllowlistEntries = ResponseWirelessGetOrganizationWirelessSSIDsFirewallIsolationAllowlistEntriesItemToBody(organizationsWirelessSSIDsFirewallIsolationAllowlistEntries, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessSSIDsFirewallIsolationAllowlistEntries)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntries struct {
	OrganizationID types.String                                                                   `tfsdk:"organization_id"`
	PerPage        types.Int64                                                                    `tfsdk:"per_page"`
	StartingAfter  types.String                                                                   `tfsdk:"starting_after"`
	EndingBefore   types.String                                                                   `tfsdk:"ending_before"`
	NetworkIDs     types.List                                                                     `tfsdk:"network_ids"`
	SSIDs          types.List                                                                     `tfsdk:"ssids"`
	Item           *ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntries `tfsdk:"item"`
}

type ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntries struct {
	Items *[]ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesItems `tfsdk:"items"`
	Meta  *ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesMeta    `tfsdk:"meta"`
}

type ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesItems struct {
	Client        *ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesItemsClient  `tfsdk:"client"`
	CreatedAt     types.String                                                                               `tfsdk:"created_at"`
	Description   types.String                                                                               `tfsdk:"description"`
	EntryID       types.String                                                                               `tfsdk:"entry_id"`
	LastUpdatedAt types.String                                                                               `tfsdk:"last_updated_at"`
	Network       *ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesItemsNetwork `tfsdk:"network"`
	SSID          *ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesItemsSsid    `tfsdk:"ssid"`
}

type ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesItemsClient struct {
	Mac types.String `tfsdk:"mac"`
}

type ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesItemsNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesItemsSsid struct {
	ID     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Number types.Int64  `tfsdk:"number"`
}

type ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesMeta struct {
	Counts *ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesMetaCounts `tfsdk:"counts"`
}

type ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesMetaCounts struct {
	Items *ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesMetaCountsItems `tfsdk:"items"`
}

type ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseWirelessGetOrganizationWirelessSSIDsFirewallIsolationAllowlistEntriesItemToBody(state OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntries, response *merakigosdk.ResponseWirelessGetOrganizationWirelessSSIDsFirewallIsolationAllowlistEntries) OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntries {
	itemState := ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntries{
		Items: func() *[]ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesItems {
			if response.Items != nil {
				result := make([]ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesItems{
						Client: func() *ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesItemsClient {
							if items.Client != nil {
								return &ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesItemsClient{
									Mac: types.StringValue(items.Client.Mac),
								}
							}
							return nil
						}(),
						CreatedAt:     types.StringValue(items.CreatedAt),
						Description:   types.StringValue(items.Description),
						EntryID:       types.StringValue(items.EntryID),
						LastUpdatedAt: types.StringValue(items.LastUpdatedAt),
						Network: func() *ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesItemsNetwork {
							if items.Network != nil {
								return &ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesItemsNetwork{
									ID:   types.StringValue(items.Network.ID),
									Name: types.StringValue(items.Network.Name),
								}
							}
							return nil
						}(),
						SSID: func() *ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesItemsSsid {
							if items.SSID != nil {
								return &ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesItemsSsid{
									ID:   types.StringValue(items.SSID.ID),
									Name: types.StringValue(items.SSID.Name),
									Number: func() types.Int64 {
										if items.SSID.Number != nil {
											return types.Int64Value(int64(*items.SSID.Number))
										}
										return types.Int64{}
									}(),
								}
							}
							return nil
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		Meta: func() *ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesMeta {
			if response.Meta != nil {
				return &ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesMeta{
					Counts: func() *ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesMetaCounts{
								Items: func() *ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseWirelessGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesMetaCountsItems{
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
