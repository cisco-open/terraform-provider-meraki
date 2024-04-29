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

// RESOURCE ACTION

import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksVLANProfilesAssignmentsReassignResource{}
	_ resource.ResourceWithConfigure = &NetworksVLANProfilesAssignmentsReassignResource{}
)

func NewNetworksVLANProfilesAssignmentsReassignResource() resource.Resource {
	return &NetworksVLANProfilesAssignmentsReassignResource{}
}

type NetworksVLANProfilesAssignmentsReassignResource struct {
	client *merakigosdk.Client
}

func (r *NetworksVLANProfilesAssignmentsReassignResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksVLANProfilesAssignmentsReassignResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_vlan_profiles_assignments_reassign"
}

// resourceAction
func (r *NetworksVLANProfilesAssignmentsReassignResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"serials": schema.ListAttribute{
						MarkdownDescription: `Array of Device Serials`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"stack_ids": schema.ListAttribute{
						MarkdownDescription: `Array of Switch Stack IDs`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"vlan_profile": schema.SingleNestedAttribute{
						MarkdownDescription: `The VLAN Profile`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"iname": schema.StringAttribute{
								MarkdownDescription: `IName of the VLAN Profile`,
								Computed:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `Name of the VLAN Profile`,
								Computed:            true,
							},
						},
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"serials": schema.ListAttribute{
						MarkdownDescription: `Array of Device Serials`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"stack_ids": schema.ListAttribute{
						MarkdownDescription: `Array of Switch Stack IDs`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"vlan_profile": schema.SingleNestedAttribute{
						MarkdownDescription: `The VLAN Profile`,
						Optional:            true,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"iname": schema.StringAttribute{
								MarkdownDescription: `IName of the VLAN Profile`,
								Optional:            true,
								Computed:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
							},
						},
					},
				},
			},
		},
	}
}
func (r *NetworksVLANProfilesAssignmentsReassignResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksVLANProfilesAssignmentsReassign

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
	//Has Paths
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Networks.ReassignNetworkVLANProfilesAssignments(vvNetworkID, dataRequest)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing ReassignNetworkVLANProfilesAssignments",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing ReassignNetworkVLANProfilesAssignments",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseNetworksReassignNetworkVLANProfilesAssignmentsItemToBody(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksVLANProfilesAssignmentsReassignResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksVLANProfilesAssignmentsReassignResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksVLANProfilesAssignmentsReassignResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksVLANProfilesAssignmentsReassign struct {
	NetworkID  types.String                                             `tfsdk:"network_id"`
	Item       *ResponseNetworksReassignNetworkVlanProfilesAssignments  `tfsdk:"item"`
	Parameters *RequestNetworksReassignNetworkVlanProfilesAssignmentsRs `tfsdk:"parameters"`
}

type ResponseNetworksReassignNetworkVlanProfilesAssignments struct {
	Serials     types.Set                                                          `tfsdk:"serials"`
	StackIDs    types.Set                                                          `tfsdk:"stack_ids"`
	VLANProfile *ResponseNetworksReassignNetworkVlanProfilesAssignmentsVlanProfile `tfsdk:"vlan_profile"`
}

type ResponseNetworksReassignNetworkVlanProfilesAssignmentsVlanProfile struct {
	Iname types.String `tfsdk:"iname"`
	Name  types.String `tfsdk:"name"`
}

type RequestNetworksReassignNetworkVlanProfilesAssignmentsRs struct {
	Serials     types.Set                                                           `tfsdk:"serials"`
	StackIDs    types.Set                                                           `tfsdk:"stack_ids"`
	VLANProfile *RequestNetworksReassignNetworkVlanProfilesAssignmentsVlanProfileRs `tfsdk:"vlan_profile"`
}

type RequestNetworksReassignNetworkVlanProfilesAssignmentsVlanProfileRs struct {
	Iname types.String `tfsdk:"iname"`
}

// FromBody
func (r *NetworksVLANProfilesAssignmentsReassign) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksReassignNetworkVLANProfilesAssignments {
	re := *r.Parameters
	var serials []string = nil
	re.Serials.ElementsAs(ctx, &serials, false)
	var stackIDs []string = nil
	re.StackIDs.ElementsAs(ctx, &stackIDs, false)
	var requestNetworksReassignNetworkVLANProfilesAssignmentsVLANProfile *merakigosdk.RequestNetworksReassignNetworkVLANProfilesAssignmentsVLANProfile
	if re.VLANProfile != nil {
		iname := re.VLANProfile.Iname.ValueString()
		requestNetworksReassignNetworkVLANProfilesAssignmentsVLANProfile = &merakigosdk.RequestNetworksReassignNetworkVLANProfilesAssignmentsVLANProfile{
			Iname: iname,
		}
	}
	out := merakigosdk.RequestNetworksReassignNetworkVLANProfilesAssignments{
		Serials:     serials,
		StackIDs:    stackIDs,
		VLANProfile: requestNetworksReassignNetworkVLANProfilesAssignmentsVLANProfile,
	}
	return &out
}

// ToBody
func ResponseNetworksReassignNetworkVLANProfilesAssignmentsItemToBody(state NetworksVLANProfilesAssignmentsReassign, response *merakigosdk.ResponseNetworksReassignNetworkVLANProfilesAssignments) NetworksVLANProfilesAssignmentsReassign {
	itemState := ResponseNetworksReassignNetworkVlanProfilesAssignments{
		Serials:  StringSliceToSet(response.Serials),
		StackIDs: StringSliceToSet(response.StackIDs),
		VLANProfile: func() *ResponseNetworksReassignNetworkVlanProfilesAssignmentsVlanProfile {
			if response.VLANProfile != nil {
				return &ResponseNetworksReassignNetworkVlanProfilesAssignmentsVlanProfile{
					Iname: types.StringValue(response.VLANProfile.Iname),
					Name:  types.StringValue(response.VLANProfile.Name),
				}
			}
			return &ResponseNetworksReassignNetworkVlanProfilesAssignmentsVlanProfile{}
		}(),
	}
	state.Item = &itemState
	return state
}
