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

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksVLANProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksVLANProfilesDataSource{}
)

func NewNetworksVLANProfilesDataSource() datasource.DataSource {
	return &NetworksVLANProfilesDataSource{}
}

type NetworksVLANProfilesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksVLANProfilesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksVLANProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_vlan_profiles"
}

func (d *NetworksVLANProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"iname": schema.StringAttribute{
				MarkdownDescription: `iname path parameter.`,
				Required:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"iname": schema.StringAttribute{
						MarkdownDescription: `IName of the VLAN profile`,
						Computed:            true,
					},
					"is_default": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating the default VLAN Profile for any device that does not have a profile explicitly assigned`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of the profile, string length must be from 1 to 255 characters`,
						Computed:            true,
					},
					"vlan_groups": schema.SetNestedAttribute{
						MarkdownDescription: `An array of named VLANs`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"name": schema.StringAttribute{
									MarkdownDescription: `Name of the VLAN, string length must be from 1 to 32 characters`,
									Computed:            true,
								},
								"vlan_ids": schema.StringAttribute{
									MarkdownDescription: `Comma-separated VLAN IDs or ID ranges`,
									Computed:            true,
								},
							},
						},
					},
					"vlan_names": schema.SetNestedAttribute{
						MarkdownDescription: `An array of named VLANs`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"adaptive_policy_group": schema.SingleNestedAttribute{
									MarkdownDescription: `Adaptive Policy Group assigned to Vlan ID`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `Adaptive Policy Group ID`,
											Computed:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: `Adaptive Policy Group name`,
											Computed:            true,
										},
									},
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Name of the VLAN, string length must be from 1 to 32 characters`,
									Computed:            true,
								},
								"vlan_id": schema.StringAttribute{
									MarkdownDescription: `VLAN ID`,
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

func (d *NetworksVLANProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksVLANProfiles NetworksVLANProfiles
	diags := req.Config.Get(ctx, &networksVLANProfiles)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkVLANProfile")
		vvNetworkID := networksVLANProfiles.NetworkID.ValueString()
		vvIname := networksVLANProfiles.Iname.ValueString()

		response1, restyResp1, err := d.client.Networks.GetNetworkVLANProfile(vvNetworkID, vvIname)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkVLANProfile",
				err.Error(),
			)
			return
		}

		networksVLANProfiles = ResponseNetworksGetNetworkVLANProfileItemToBody(networksVLANProfiles, response1)
		diags = resp.State.Set(ctx, &networksVLANProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksVLANProfiles struct {
	NetworkID types.String                           `tfsdk:"network_id"`
	Iname     types.String                           `tfsdk:"iname"`
	Item      *ResponseNetworksGetNetworkVlanProfile `tfsdk:"item"`
}

type ResponseNetworksGetNetworkVlanProfile struct {
	Iname      types.String                                       `tfsdk:"iname"`
	IsDefault  types.Bool                                         `tfsdk:"is_default"`
	Name       types.String                                       `tfsdk:"name"`
	VLANGroups *[]ResponseNetworksGetNetworkVlanProfileVlanGroups `tfsdk:"vlan_groups"`
	VLANNames  *[]ResponseNetworksGetNetworkVlanProfileVlanNames  `tfsdk:"vlan_names"`
}

type ResponseNetworksGetNetworkVlanProfileVlanGroups struct {
	Name    types.String `tfsdk:"name"`
	VLANIDs types.String `tfsdk:"vlan_ids"`
}

type ResponseNetworksGetNetworkVlanProfileVlanNames struct {
	AdaptivePolicyGroup *ResponseNetworksGetNetworkVlanProfileVlanNamesAdaptivePolicyGroup `tfsdk:"adaptive_policy_group"`
	Name                types.String                                                       `tfsdk:"name"`
	VLANID              types.String                                                       `tfsdk:"vlan_id"`
}

type ResponseNetworksGetNetworkVlanProfileVlanNamesAdaptivePolicyGroup struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// ToBody
func ResponseNetworksGetNetworkVLANProfileItemToBody(state NetworksVLANProfiles, response *merakigosdk.ResponseNetworksGetNetworkVLANProfile) NetworksVLANProfiles {
	itemState := ResponseNetworksGetNetworkVlanProfile{
		Iname: types.StringValue(response.Iname),
		IsDefault: func() types.Bool {
			if response.IsDefault != nil {
				return types.BoolValue(*response.IsDefault)
			}
			return types.Bool{}
		}(),
		Name: types.StringValue(response.Name),
		VLANGroups: func() *[]ResponseNetworksGetNetworkVlanProfileVlanGroups {
			if response.VLANGroups != nil {
				result := make([]ResponseNetworksGetNetworkVlanProfileVlanGroups, len(*response.VLANGroups))
				for i, vLANGroups := range *response.VLANGroups {
					result[i] = ResponseNetworksGetNetworkVlanProfileVlanGroups{
						Name:    types.StringValue(vLANGroups.Name),
						VLANIDs: types.StringValue(vLANGroups.VLANIDs),
					}
				}
				return &result
			}
			return &[]ResponseNetworksGetNetworkVlanProfileVlanGroups{}
		}(),
		VLANNames: func() *[]ResponseNetworksGetNetworkVlanProfileVlanNames {
			if response.VLANNames != nil {
				result := make([]ResponseNetworksGetNetworkVlanProfileVlanNames, len(*response.VLANNames))
				for i, vLANNames := range *response.VLANNames {
					result[i] = ResponseNetworksGetNetworkVlanProfileVlanNames{
						AdaptivePolicyGroup: func() *ResponseNetworksGetNetworkVlanProfileVlanNamesAdaptivePolicyGroup {
							if vLANNames.AdaptivePolicyGroup != nil {
								return &ResponseNetworksGetNetworkVlanProfileVlanNamesAdaptivePolicyGroup{
									ID:   types.StringValue(vLANNames.AdaptivePolicyGroup.ID),
									Name: types.StringValue(vLANNames.AdaptivePolicyGroup.Name),
								}
							}
							return &ResponseNetworksGetNetworkVlanProfileVlanNamesAdaptivePolicyGroup{}
						}(),
						Name:   types.StringValue(vLANNames.Name),
						VLANID: types.StringValue(vLANNames.VLANID),
					}
				}
				return &result
			}
			return &[]ResponseNetworksGetNetworkVlanProfileVlanNames{}
		}(),
	}
	state.Item = &itemState
	return state
}
