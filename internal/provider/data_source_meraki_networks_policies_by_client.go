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
	_ datasource.DataSource              = &NetworksPoliciesByClientDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksPoliciesByClientDataSource{}
)

func NewNetworksPoliciesByClientDataSource() datasource.DataSource {
	return &NetworksPoliciesByClientDataSource{}
}

type NetworksPoliciesByClientDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksPoliciesByClientDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksPoliciesByClientDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_policies_by_client"
}

func (d *NetworksPoliciesByClientDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
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
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 31 days from today.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseNetworksGetNetworkPoliciesByClient`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"assigned": schema.SetNestedAttribute{
							MarkdownDescription: `Assigned policies`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"group_policy_id": schema.StringAttribute{
										MarkdownDescription: `id of policy`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `name of policy`,
										Computed:            true,
									},
									"ssid": schema.SetNestedAttribute{
										MarkdownDescription: `ssid`,
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{

												"ssid_number": schema.Int64Attribute{
													MarkdownDescription: `number of ssid`,
													Computed:            true,
												},
											},
										},
									},
									"type": schema.StringAttribute{
										MarkdownDescription: `type of policy`,
										Computed:            true,
									},
								},
							},
						},
						"client_id": schema.StringAttribute{
							MarkdownDescription: `ID of client`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of client`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksPoliciesByClientDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksPoliciesByClient NetworksPoliciesByClient
	diags := req.Config.Get(ctx, &networksPoliciesByClient)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkPoliciesByClient")
		vvNetworkID := networksPoliciesByClient.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkPoliciesByClientQueryParams{}

		queryParams1.PerPage = int(networksPoliciesByClient.PerPage.ValueInt64())
		queryParams1.StartingAfter = networksPoliciesByClient.StartingAfter.ValueString()
		queryParams1.EndingBefore = networksPoliciesByClient.EndingBefore.ValueString()
		queryParams1.T0 = networksPoliciesByClient.T0.ValueString()
		queryParams1.Timespan = networksPoliciesByClient.Timespan.ValueFloat64()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkPoliciesByClient(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkPoliciesByClient",
				err.Error(),
			)
			return
		}

		networksPoliciesByClient = ResponseNetworksGetNetworkPoliciesByClientItemsToBody(networksPoliciesByClient, response1)
		diags = resp.State.Set(ctx, &networksPoliciesByClient)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksPoliciesByClient struct {
	NetworkID     types.String                                      `tfsdk:"network_id"`
	PerPage       types.Int64                                       `tfsdk:"per_page"`
	StartingAfter types.String                                      `tfsdk:"starting_after"`
	EndingBefore  types.String                                      `tfsdk:"ending_before"`
	T0            types.String                                      `tfsdk:"t0"`
	Timespan      types.Float64                                     `tfsdk:"timespan"`
	Items         *[]ResponseItemNetworksGetNetworkPoliciesByClient `tfsdk:"items"`
}

type ResponseItemNetworksGetNetworkPoliciesByClient struct {
	Assigned *[]ResponseItemNetworksGetNetworkPoliciesByClientAssigned `tfsdk:"assigned"`
	ClientID types.String                                              `tfsdk:"client_id"`
	Name     types.String                                              `tfsdk:"name"`
}

type ResponseItemNetworksGetNetworkPoliciesByClientAssigned struct {
	GroupPolicyID types.String                                                  `tfsdk:"group_policy_id"`
	Name          types.String                                                  `tfsdk:"name"`
	SSID          *[]ResponseItemNetworksGetNetworkPoliciesByClientAssignedSsid `tfsdk:"ssid"`
	Type          types.String                                                  `tfsdk:"type"`
}

type ResponseItemNetworksGetNetworkPoliciesByClientAssignedSsid struct {
	SSIDNumber types.Int64 `tfsdk:"ssid_number"`
}

// ToBody
func ResponseNetworksGetNetworkPoliciesByClientItemsToBody(state NetworksPoliciesByClient, response *merakigosdk.ResponseNetworksGetNetworkPoliciesByClient) NetworksPoliciesByClient {
	var items []ResponseItemNetworksGetNetworkPoliciesByClient
	for _, item := range *response {
		itemState := ResponseItemNetworksGetNetworkPoliciesByClient{
			Assigned: func() *[]ResponseItemNetworksGetNetworkPoliciesByClientAssigned {
				if item.Assigned != nil {
					result := make([]ResponseItemNetworksGetNetworkPoliciesByClientAssigned, len(*item.Assigned))
					for i, assigned := range *item.Assigned {
						result[i] = ResponseItemNetworksGetNetworkPoliciesByClientAssigned{
							GroupPolicyID: types.StringValue(assigned.GroupPolicyID),
							Name:          types.StringValue(assigned.Name),
							SSID: func() *[]ResponseItemNetworksGetNetworkPoliciesByClientAssignedSsid {
								if assigned.SSID != nil {
									result := make([]ResponseItemNetworksGetNetworkPoliciesByClientAssignedSsid, len(*assigned.SSID))
									for i, sSID := range *assigned.SSID {
										result[i] = ResponseItemNetworksGetNetworkPoliciesByClientAssignedSsid{
											SSIDNumber: func() types.Int64 {
												if sSID.SSIDNumber != nil {
													return types.Int64Value(int64(*sSID.SSIDNumber))
												}
												return types.Int64{}
											}(),
										}
									}
									return &result
								}
								return nil
							}(),
							Type: types.StringValue(assigned.Type),
						}
					}
					return &result
				}
				return nil
			}(),
			ClientID: types.StringValue(item.ClientID),
			Name:     types.StringValue(item.Name),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
