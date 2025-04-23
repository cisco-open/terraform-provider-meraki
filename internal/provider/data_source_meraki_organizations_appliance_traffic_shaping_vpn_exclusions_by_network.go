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
	_ datasource.DataSource              = &OrganizationsApplianceTrafficShapingVpnExclusionsByNetworkDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsApplianceTrafficShapingVpnExclusionsByNetworkDataSource{}
)

func NewOrganizationsApplianceTrafficShapingVpnExclusionsByNetworkDataSource() datasource.DataSource {
	return &OrganizationsApplianceTrafficShapingVpnExclusionsByNetworkDataSource{}
}

type OrganizationsApplianceTrafficShapingVpnExclusionsByNetworkDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsApplianceTrafficShapingVpnExclusionsByNetworkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsApplianceTrafficShapingVpnExclusionsByNetworkDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_appliance_traffic_shaping_vpn_exclusions_by_network"
}

func (d *OrganizationsApplianceTrafficShapingVpnExclusionsByNetworkDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 50.`,
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
						MarkdownDescription: `VPN exclusion rules by network`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"custom": schema.SetNestedAttribute{
									MarkdownDescription: `Custom VPN exclusion rules.`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"destination": schema.StringAttribute{
												MarkdownDescription: `Destination address; hostname required for DNS, IPv4 otherwise.`,
												Computed:            true,
											},
											"port": schema.StringAttribute{
												MarkdownDescription: `Destination port.`,
												Computed:            true,
											},
											"protocol": schema.StringAttribute{
												MarkdownDescription: `Protocol.`,
												Computed:            true,
											},
										},
									},
								},
								"major_applications": schema.SetNestedAttribute{
									MarkdownDescription: `Major Application based VPN exclusion rules.`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"id": schema.StringAttribute{
												MarkdownDescription: `Application's Meraki ID.`,
												Computed:            true,
											},
											"name": schema.StringAttribute{
												MarkdownDescription: `Application's name.`,
												Computed:            true,
											},
										},
									},
								},
								"network_id": schema.StringAttribute{
									MarkdownDescription: `ID of the network whose VPN exclusion rules are returned.`,
									Computed:            true,
								},
								"network_name": schema.StringAttribute{
									MarkdownDescription: `Name of the network whose VPN exclusion rules are returned.`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsApplianceTrafficShapingVpnExclusionsByNetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsApplianceTrafficShapingVpnExclusionsByNetwork OrganizationsApplianceTrafficShapingVpnExclusionsByNetwork
	diags := req.Config.Get(ctx, &organizationsApplianceTrafficShapingVpnExclusionsByNetwork)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork")
		vvOrganizationID := organizationsApplianceTrafficShapingVpnExclusionsByNetwork.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkQueryParams{}

		queryParams1.PerPage = int(organizationsApplianceTrafficShapingVpnExclusionsByNetwork.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsApplianceTrafficShapingVpnExclusionsByNetwork.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsApplianceTrafficShapingVpnExclusionsByNetwork.EndingBefore.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsApplianceTrafficShapingVpnExclusionsByNetwork.NetworkIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork",
				err.Error(),
			)
			return
		}

		organizationsApplianceTrafficShapingVpnExclusionsByNetwork = ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItemToBody(organizationsApplianceTrafficShapingVpnExclusionsByNetwork, response1)
		diags = resp.State.Set(ctx, &organizationsApplianceTrafficShapingVpnExclusionsByNetwork)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsApplianceTrafficShapingVpnExclusionsByNetwork struct {
	OrganizationID types.String                                                                   `tfsdk:"organization_id"`
	PerPage        types.Int64                                                                    `tfsdk:"per_page"`
	StartingAfter  types.String                                                                   `tfsdk:"starting_after"`
	EndingBefore   types.String                                                                   `tfsdk:"ending_before"`
	NetworkIDs     types.List                                                                     `tfsdk:"network_ids"`
	Item           *ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork `tfsdk:"item"`
}

type ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork struct {
	Items *[]ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItems `tfsdk:"items"`
}

type ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItems struct {
	Custom            *[]ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItemsCustom            `tfsdk:"custom"`
	MajorApplications *[]ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItemsMajorApplications `tfsdk:"major_applications"`
	NetworkID         types.String                                                                                           `tfsdk:"network_id"`
	NetworkName       types.String                                                                                           `tfsdk:"network_name"`
}

type ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItemsCustom struct {
	Destination types.String `tfsdk:"destination"`
	Port        types.String `tfsdk:"port"`
	Protocol    types.String `tfsdk:"protocol"`
}

type ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItemsMajorApplications struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// ToBody
func ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItemToBody(state OrganizationsApplianceTrafficShapingVpnExclusionsByNetwork, response *merakigosdk.ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork) OrganizationsApplianceTrafficShapingVpnExclusionsByNetwork {
	itemState := ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork{
		Items: func() *[]ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItems {
			if response.Items != nil {
				result := make([]ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItems{
						Custom: func() *[]ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItemsCustom {
							if items.Custom != nil {
								result := make([]ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItemsCustom, len(*items.Custom))
								for i, custom := range *items.Custom {
									result[i] = ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItemsCustom{
										Destination: types.StringValue(custom.Destination),
										Port:        types.StringValue(custom.Port),
										Protocol:    types.StringValue(custom.Protocol),
									}
								}
								return &result
							}
							return nil
						}(),
						MajorApplications: func() *[]ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItemsMajorApplications {
							if items.MajorApplications != nil {
								result := make([]ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItemsMajorApplications, len(*items.MajorApplications))
								for i, majorApplications := range *items.MajorApplications {
									result[i] = ResponseApplianceGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkItemsMajorApplications{
										ID:   types.StringValue(majorApplications.ID),
										Name: types.StringValue(majorApplications.Name),
									}
								}
								return &result
							}
							return nil
						}(),
						NetworkID:   types.StringValue(items.NetworkID),
						NetworkName: types.StringValue(items.NetworkName),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
