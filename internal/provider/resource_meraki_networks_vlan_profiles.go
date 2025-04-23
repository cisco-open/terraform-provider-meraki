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

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksVLANProfilesResource{}
	_ resource.ResourceWithConfigure = &NetworksVLANProfilesResource{}
)

func NewNetworksVLANProfilesResource() resource.Resource {
	return &NetworksVLANProfilesResource{}
}

type NetworksVLANProfilesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksVLANProfilesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksVLANProfilesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_vlan_profiles"
}

func (r *NetworksVLANProfilesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"iname": schema.StringAttribute{
				MarkdownDescription: `IName of the VLAN profile`,
				Required:            true,
			},
			"is_default": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating the default VLAN Profile for any device that does not have a profile explicitly assigned`,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Name of the profile, string length must be from 1 to 255 characters`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"vlan_groups": schema.SetNestedAttribute{
				MarkdownDescription: `An array of named VLANs`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the VLAN, string length must be from 1 to 32 characters`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"vlan_ids": schema.StringAttribute{
							MarkdownDescription: `Comma-separated VLAN IDs or ID ranges`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"vlan_names": schema.SetNestedAttribute{
				MarkdownDescription: `An array of named VLANs`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"adaptive_policy_group": schema.SingleNestedAttribute{
							MarkdownDescription: `Adaptive Policy Group assigned to Vlan ID`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `Adaptive Policy Group ID`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
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
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"vlan_id": schema.StringAttribute{
							MarkdownDescription: `VLAN ID`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
		},
	}
}

//path params to set ['iname']
//path params to assign NOT EDITABLE ['iname']

func (r *NetworksVLANProfilesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksVLANProfilesRs

	var item types.Object
	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Has Paths
	vvNetworkID := data.NetworkID.ValueString()
	//Has Item and has items and post

	vvName := data.Name.ValueString()

	responseVerifyItem, restyResp1, err := r.client.Networks.GetNetworkVLANProfiles(vvNetworkID)
	//Has Post
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkVLANProfiles",
					err.Error(),
				)
				return
			}
		}
	}

	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvIname, ok := result2["Iname"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter Iname",
					"Fail Parsing Iname",
				)
				return
			}
			r.client.Networks.UpdateNetworkVLANProfile(vvNetworkID, vvIname, data.toSdkApiRequestUpdate(ctx))

			responseVerifyItem2, _, _ := r.client.Networks.GetNetworkVLANProfile(vvNetworkID, vvIname)
			if responseVerifyItem2 != nil {
				data = ResponseNetworksGetNetworkVLANProfileItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Networks.CreateNetworkVLANProfile(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkVLANProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkVLANProfile",
			err.Error(),
		)
		return
	}

	responseGet, restyResp1, err := r.client.Networks.GetNetworkVLANProfiles(vvNetworkID)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkVLANProfiles",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkVLANProfiles",
			err.Error(),
		)
		return
	}

	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvIname, ok := result2["Iname"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter Iname",
				"Fail Parsing Iname",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Networks.GetNetworkVLANProfile(vvNetworkID, vvIname)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseNetworksGetNetworkVLANProfileItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkVLANProfile",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkVLANProfile",
				err.Error(),
			)
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"Error in result.",
			"Error in result.",
		)
		return
	}

}

func (r *NetworksVLANProfilesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksVLANProfilesRs

	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	vvIname := data.Iname.ValueString()
	responseGet, restyRespGet, err := r.client.Networks.GetNetworkVLANProfile(vvNetworkID, vvIname)
	if err != nil || restyRespGet == nil {
		if restyRespGet != nil {
			if restyRespGet.StatusCode() == 404 {
				resp.Diagnostics.AddWarning(
					"Resource not found",
					"Deleting resource",
				)
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkVLANProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkVLANProfile",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseNetworksGetNetworkVLANProfileItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksVLANProfilesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("iname"), idParts[1])...)
}

func (r *NetworksVLANProfilesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksVLANProfilesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvIname := data.Iname.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkVLANProfile(vvNetworkID, vvIname, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkVLANProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkVLANProfile",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksVLANProfilesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksVLANProfilesRs
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &state, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}

	vvNetworkID := state.NetworkID.ValueString()
	vvIname := state.Iname.ValueString()
	_, err := r.client.Networks.DeleteNetworkVLANProfile(vvNetworkID, vvIname)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkVLANProfile", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksVLANProfilesRs struct {
	NetworkID  types.String                                         `tfsdk:"network_id"`
	Iname      types.String                                         `tfsdk:"iname"`
	IsDefault  types.Bool                                           `tfsdk:"is_default"`
	Name       types.String                                         `tfsdk:"name"`
	VLANGroups *[]ResponseNetworksGetNetworkVlanProfileVlanGroupsRs `tfsdk:"vlan_groups"`
	VLANNames  *[]ResponseNetworksGetNetworkVlanProfileVlanNamesRs  `tfsdk:"vlan_names"`
}

type ResponseNetworksGetNetworkVlanProfileVlanGroupsRs struct {
	Name    types.String `tfsdk:"name"`
	VLANIDs types.String `tfsdk:"vlan_ids"`
}

type ResponseNetworksGetNetworkVlanProfileVlanNamesRs struct {
	AdaptivePolicyGroup *ResponseNetworksGetNetworkVlanProfileVlanNamesAdaptivePolicyGroupRs `tfsdk:"adaptive_policy_group"`
	Name                types.String                                                         `tfsdk:"name"`
	VLANID              types.String                                                         `tfsdk:"vlan_id"`
}

type ResponseNetworksGetNetworkVlanProfileVlanNamesAdaptivePolicyGroupRs struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// FromBody
func (r *NetworksVLANProfilesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksCreateNetworkVLANProfile {
	emptyString := ""
	iname := new(string)
	if !r.Iname.IsUnknown() && !r.Iname.IsNull() {
		*iname = r.Iname.ValueString()
	} else {
		iname = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestNetworksCreateNetworkVLANProfileVLANGroups []merakigosdk.RequestNetworksCreateNetworkVLANProfileVLANGroups

	if r.VLANGroups != nil {
		for _, rItem1 := range *r.VLANGroups {
			name := rItem1.Name.ValueString()
			vlanIDs := rItem1.VLANIDs.ValueString()
			requestNetworksCreateNetworkVLANProfileVLANGroups = append(requestNetworksCreateNetworkVLANProfileVLANGroups, merakigosdk.RequestNetworksCreateNetworkVLANProfileVLANGroups{
				Name:    name,
				VLANIDs: vlanIDs,
			})
			//[debug] Is Array: True
		}
	}
	var requestNetworksCreateNetworkVLANProfileVLANNames []merakigosdk.RequestNetworksCreateNetworkVLANProfileVLANNames

	if r.VLANNames != nil {
		for _, rItem1 := range *r.VLANNames {
			var requestNetworksCreateNetworkVLANProfileVLANNamesAdaptivePolicyGroup *merakigosdk.RequestNetworksCreateNetworkVLANProfileVLANNamesAdaptivePolicyGroup

			if rItem1.AdaptivePolicyGroup != nil {
				id := rItem1.AdaptivePolicyGroup.ID.ValueString()
				requestNetworksCreateNetworkVLANProfileVLANNamesAdaptivePolicyGroup = &merakigosdk.RequestNetworksCreateNetworkVLANProfileVLANNamesAdaptivePolicyGroup{
					ID: id,
				}
				//[debug] Is Array: False
			}
			name := rItem1.Name.ValueString()
			vlanID := rItem1.VLANID.ValueString()
			requestNetworksCreateNetworkVLANProfileVLANNames = append(requestNetworksCreateNetworkVLANProfileVLANNames, merakigosdk.RequestNetworksCreateNetworkVLANProfileVLANNames{
				AdaptivePolicyGroup: requestNetworksCreateNetworkVLANProfileVLANNamesAdaptivePolicyGroup,
				Name:                name,
				VLANID:              vlanID,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestNetworksCreateNetworkVLANProfile{
		Iname: *iname,
		Name:  *name,
		VLANGroups: func() *[]merakigosdk.RequestNetworksCreateNetworkVLANProfileVLANGroups {
			if len(requestNetworksCreateNetworkVLANProfileVLANGroups) > 0 {
				return &requestNetworksCreateNetworkVLANProfileVLANGroups
			}
			return nil
		}(),
		VLANNames: func() *[]merakigosdk.RequestNetworksCreateNetworkVLANProfileVLANNames {
			if len(requestNetworksCreateNetworkVLANProfileVLANNames) > 0 {
				return &requestNetworksCreateNetworkVLANProfileVLANNames
			}
			return nil
		}(),
	}
	return &out
}
func (r *NetworksVLANProfilesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestNetworksUpdateNetworkVLANProfile {
	emptyString := ""
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestNetworksUpdateNetworkVLANProfileVLANGroups []merakigosdk.RequestNetworksUpdateNetworkVLANProfileVLANGroups

	if r.VLANGroups != nil {
		for _, rItem1 := range *r.VLANGroups {
			name := rItem1.Name.ValueString()
			vlanIDs := rItem1.VLANIDs.ValueString()
			requestNetworksUpdateNetworkVLANProfileVLANGroups = append(requestNetworksUpdateNetworkVLANProfileVLANGroups, merakigosdk.RequestNetworksUpdateNetworkVLANProfileVLANGroups{
				Name:    name,
				VLANIDs: vlanIDs,
			})
			//[debug] Is Array: True
		}
	}
	var requestNetworksUpdateNetworkVLANProfileVLANNames []merakigosdk.RequestNetworksUpdateNetworkVLANProfileVLANNames

	if r.VLANNames != nil {
		for _, rItem1 := range *r.VLANNames {
			var requestNetworksUpdateNetworkVLANProfileVLANNamesAdaptivePolicyGroup *merakigosdk.RequestNetworksUpdateNetworkVLANProfileVLANNamesAdaptivePolicyGroup

			if rItem1.AdaptivePolicyGroup != nil {
				id := rItem1.AdaptivePolicyGroup.ID.ValueString()
				requestNetworksUpdateNetworkVLANProfileVLANNamesAdaptivePolicyGroup = &merakigosdk.RequestNetworksUpdateNetworkVLANProfileVLANNamesAdaptivePolicyGroup{
					ID: id,
				}
				//[debug] Is Array: False
			}
			name := rItem1.Name.ValueString()
			vlanID := rItem1.VLANID.ValueString()
			requestNetworksUpdateNetworkVLANProfileVLANNames = append(requestNetworksUpdateNetworkVLANProfileVLANNames, merakigosdk.RequestNetworksUpdateNetworkVLANProfileVLANNames{
				AdaptivePolicyGroup: requestNetworksUpdateNetworkVLANProfileVLANNamesAdaptivePolicyGroup,
				Name:                name,
				VLANID:              vlanID,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestNetworksUpdateNetworkVLANProfile{
		Name: *name,
		VLANGroups: func() *[]merakigosdk.RequestNetworksUpdateNetworkVLANProfileVLANGroups {
			if len(requestNetworksUpdateNetworkVLANProfileVLANGroups) > 0 {
				return &requestNetworksUpdateNetworkVLANProfileVLANGroups
			}
			return nil
		}(),
		VLANNames: func() *[]merakigosdk.RequestNetworksUpdateNetworkVLANProfileVLANNames {
			if len(requestNetworksUpdateNetworkVLANProfileVLANNames) > 0 {
				return &requestNetworksUpdateNetworkVLANProfileVLANNames
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseNetworksGetNetworkVLANProfileItemToBodyRs(state NetworksVLANProfilesRs, response *merakigosdk.ResponseNetworksGetNetworkVLANProfile, is_read bool) NetworksVLANProfilesRs {
	itemState := NetworksVLANProfilesRs{
		Iname: types.StringValue(response.Iname),
		IsDefault: func() types.Bool {
			if response.IsDefault != nil {
				return types.BoolValue(*response.IsDefault)
			}
			return types.Bool{}
		}(),
		Name: types.StringValue(response.Name),
		VLANGroups: func() *[]ResponseNetworksGetNetworkVlanProfileVlanGroupsRs {
			if response.VLANGroups != nil {
				result := make([]ResponseNetworksGetNetworkVlanProfileVlanGroupsRs, len(*response.VLANGroups))
				for i, vLANGroups := range *response.VLANGroups {
					result[i] = ResponseNetworksGetNetworkVlanProfileVlanGroupsRs{
						Name:    types.StringValue(vLANGroups.Name),
						VLANIDs: types.StringValue(vLANGroups.VLANIDs),
					}
				}
				return &result
			}
			return nil
		}(),
		VLANNames: func() *[]ResponseNetworksGetNetworkVlanProfileVlanNamesRs {
			if response.VLANNames != nil {
				result := make([]ResponseNetworksGetNetworkVlanProfileVlanNamesRs, len(*response.VLANNames))
				for i, vLANNames := range *response.VLANNames {
					result[i] = ResponseNetworksGetNetworkVlanProfileVlanNamesRs{
						AdaptivePolicyGroup: func() *ResponseNetworksGetNetworkVlanProfileVlanNamesAdaptivePolicyGroupRs {
							if vLANNames.AdaptivePolicyGroup != nil {
								return &ResponseNetworksGetNetworkVlanProfileVlanNamesAdaptivePolicyGroupRs{
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
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksVLANProfilesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksVLANProfilesRs)
}
