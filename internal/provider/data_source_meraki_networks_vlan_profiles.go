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
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
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

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseNetworksGetNetworkVlanProfiles`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
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
		},
	}
}

func (d *NetworksVLANProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksVLANProfiles NetworksVLANProfiles
	diags := req.Config.Get(ctx, &networksVLANProfiles)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksVLANProfiles.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksVLANProfiles.NetworkID.IsNull(), !networksVLANProfiles.Iname.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkVLANProfiles")
		vvNetworkID := networksVLANProfiles.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkVLANProfiles(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkVLANProfiles",
				err.Error(),
			)
			return
		}

		networksVLANProfiles = ResponseNetworksGetNetworkVLANProfilesItemsToBody(networksVLANProfiles, response1)
		diags = resp.State.Set(ctx, &networksVLANProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkVLANProfile")
		vvNetworkID := networksVLANProfiles.NetworkID.ValueString()
		vvIname := networksVLANProfiles.Iname.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Networks.GetNetworkVLANProfile(vvNetworkID, vvIname)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkVLANProfile",
				err.Error(),
			)
			return
		}

		networksVLANProfiles = ResponseNetworksGetNetworkVLANProfileItemToBody(networksVLANProfiles, response2)
		diags = resp.State.Set(ctx, &networksVLANProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksVLANProfiles struct {
	NetworkID types.String                                  `tfsdk:"network_id"`
	Iname     types.String                                  `tfsdk:"iname"`
	Items     *[]ResponseItemNetworksGetNetworkVlanProfiles `tfsdk:"items"`
	Item      *ResponseNetworksGetNetworkVlanProfile        `tfsdk:"item"`
}

type ResponseItemNetworksGetNetworkVlanProfiles struct {
	Iname      types.String                                            `tfsdk:"iname"`
	IsDefault  types.Bool                                              `tfsdk:"is_default"`
	Name       types.String                                            `tfsdk:"name"`
	VLANGroups *[]ResponseItemNetworksGetNetworkVlanProfilesVlanGroups `tfsdk:"vlan_groups"`
	VLANNames  *[]ResponseItemNetworksGetNetworkVlanProfilesVlanNames  `tfsdk:"vlan_names"`
}

type ResponseItemNetworksGetNetworkVlanProfilesVlanGroups struct {
	Name    types.String `tfsdk:"name"`
	VLANIDs types.String `tfsdk:"vlan_ids"`
}

type ResponseItemNetworksGetNetworkVlanProfilesVlanNames struct {
	AdaptivePolicyGroup *ResponseItemNetworksGetNetworkVlanProfilesVlanNamesAdaptivePolicyGroup `tfsdk:"adaptive_policy_group"`
	Name                types.String                                                            `tfsdk:"name"`
	VLANID              types.String                                                            `tfsdk:"vlan_id"`
}

type ResponseItemNetworksGetNetworkVlanProfilesVlanNamesAdaptivePolicyGroup struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
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
func ResponseNetworksGetNetworkVLANProfilesItemsToBody(state NetworksVLANProfiles, response *merakigosdk.ResponseNetworksGetNetworkVLANProfiles) NetworksVLANProfiles {
	var items []ResponseItemNetworksGetNetworkVlanProfiles
	for _, item := range *response {
		itemState := ResponseItemNetworksGetNetworkVlanProfiles{
			Iname: types.StringValue(item.Iname),
			IsDefault: func() types.Bool {
				if item.IsDefault != nil {
					return types.BoolValue(*item.IsDefault)
				}
				return types.Bool{}
			}(),
			Name: types.StringValue(item.Name),
			VLANGroups: func() *[]ResponseItemNetworksGetNetworkVlanProfilesVlanGroups {
				if item.VLANGroups != nil {
					result := make([]ResponseItemNetworksGetNetworkVlanProfilesVlanGroups, len(*item.VLANGroups))
					for i, vLANGroups := range *item.VLANGroups {
						result[i] = ResponseItemNetworksGetNetworkVlanProfilesVlanGroups{
							Name:    types.StringValue(vLANGroups.Name),
							VLANIDs: types.StringValue(vLANGroups.VLANIDs),
						}
					}
					return &result
				}
				return nil
			}(),
			VLANNames: func() *[]ResponseItemNetworksGetNetworkVlanProfilesVlanNames {
				if item.VLANNames != nil {
					result := make([]ResponseItemNetworksGetNetworkVlanProfilesVlanNames, len(*item.VLANNames))
					for i, vLANNames := range *item.VLANNames {
						result[i] = ResponseItemNetworksGetNetworkVlanProfilesVlanNames{
							AdaptivePolicyGroup: func() *ResponseItemNetworksGetNetworkVlanProfilesVlanNamesAdaptivePolicyGroup {
								if vLANNames.AdaptivePolicyGroup != nil {
									return &ResponseItemNetworksGetNetworkVlanProfilesVlanNamesAdaptivePolicyGroup{
										ID:   types.StringValue(vLANNames.AdaptivePolicyGroup.ID),
										Name: types.StringValue(vLANNames.AdaptivePolicyGroup.Name),
									}
								}
								return nil
							}(),
							Name:   types.StringValue(vLANNames.Name),
							VLANID: types.StringValue(vLANNames.VLANID),
						}
					}
					return &result
				}
				return nil
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

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
			return nil
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
							return nil
						}(),
						Name:   types.StringValue(vLANNames.Name),
						VLANID: types.StringValue(vLANNames.VLANID),
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
