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
	"strconv"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSNMPResource{}
	_ resource.ResourceWithConfigure = &NetworksSNMPResource{}
)

func NewNetworksSNMPResource() resource.Resource {
	return &NetworksSNMPResource{}
}

type NetworksSNMPResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSNMPResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSNMPResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_snmp"
}

func (r *NetworksSNMPResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"access": schema.StringAttribute{
				MarkdownDescription: `The type of SNMP access. Can be one of 'none' (disabled), 'community' (V1/V2c), or 'users' (V3).
                                  Allowed values: [community,none,users]`,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"community",
						"none",
						"users",
					),
				},
			},
			"community_string": schema.StringAttribute{
				MarkdownDescription: `SNMP community string if access is 'community'.`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"users": schema.ListNestedAttribute{
				MarkdownDescription: `SNMP settings if access is 'users'.`,
				Optional:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"passphrase": schema.StringAttribute{
							MarkdownDescription: `The passphrase for the SNMP user.`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"username": schema.StringAttribute{
							MarkdownDescription: `The username for the SNMP user.`,
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

func (r *NetworksSNMPResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSNMPRs

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
	//Has Item and not has items

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkSNMP(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSNMP",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSNMP",
			err.Error(),
		)
		return
	}

	// Assign data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *NetworksSNMPResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSNMPRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	responseGet, restyRespGet, err := r.client.Networks.GetNetworkSNMP(vvNetworkID)
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
				"Failure when executing GetNetworkSNMP",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSNMP",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseNetworksGetNetworkSNMPItemToBodyRs(data, responseGet, true)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *NetworksSNMPResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksSNMPResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworksSNMPRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvNetworkID := plan.NetworkID.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkSNMP(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSNMP",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSNMP",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworksSNMPResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksSNMP", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSNMPRs struct {
	NetworkID       types.String                             `tfsdk:"network_id"`
	Access          types.String                             `tfsdk:"access"`
	CommunityString types.String                             `tfsdk:"community_string"`
	Users           *[]ResponseNetworksGetNetworkSnmpUsersRs `tfsdk:"users"`
}

type ResponseNetworksGetNetworkSnmpUsersRs struct {
	Passphrase types.String `tfsdk:"passphrase"`
	Username   types.String `tfsdk:"username"`
}

// FromBody
func (r *NetworksSNMPRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestNetworksUpdateNetworkSNMP {
	emptyString := ""
	access := new(string)
	if !r.Access.IsUnknown() && !r.Access.IsNull() {
		*access = r.Access.ValueString()
	} else {
		access = &emptyString
	}
	communityString := new(string)
	if !r.CommunityString.IsUnknown() && !r.CommunityString.IsNull() {
		*communityString = r.CommunityString.ValueString()
	} else {
		communityString = &emptyString
	}
	var requestNetworksUpdateNetworkSNMPUsers []merakigosdk.RequestNetworksUpdateNetworkSNMPUsers

	if r.Users != nil {
		for _, rItem1 := range *r.Users {
			passphrase := rItem1.Passphrase.ValueString()
			username := rItem1.Username.ValueString()
			requestNetworksUpdateNetworkSNMPUsers = append(requestNetworksUpdateNetworkSNMPUsers, merakigosdk.RequestNetworksUpdateNetworkSNMPUsers{
				Passphrase: passphrase,
				Username:   username,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestNetworksUpdateNetworkSNMP{
		Access:          *access,
		CommunityString: *communityString,
		Users: func() *[]merakigosdk.RequestNetworksUpdateNetworkSNMPUsers {
			if len(requestNetworksUpdateNetworkSNMPUsers) > 0 {
				return &requestNetworksUpdateNetworkSNMPUsers
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseNetworksGetNetworkSNMPItemToBodyRs(state NetworksSNMPRs, response *merakigosdk.ResponseNetworksGetNetworkSNMP, is_read bool) NetworksSNMPRs {
	itemState := NetworksSNMPRs{
		Access: func() types.String {
			if response.Access != "" {
				return types.StringValue(response.Access)
			}
			return types.String{}
		}(),
		CommunityString: func() types.String {
			if response.CommunityString != "" {
				return types.StringValue(response.CommunityString)
			}
			return types.String{}
		}(),
		Users: func() *[]ResponseNetworksGetNetworkSnmpUsersRs {
			if response.Users != nil {
				result := make([]ResponseNetworksGetNetworkSnmpUsersRs, len(*response.Users))
				for i, users := range *response.Users {
					result[i] = ResponseNetworksGetNetworkSnmpUsersRs{
						Passphrase: func() types.String {
							if users.Passphrase != "" {
								return types.StringValue(users.Passphrase)
							}
							return types.String{}
						}(),
						Username: func() types.String {
							if users.Username != "" {
								return types.StringValue(users.Username)
							}
							return types.String{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSNMPRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSNMPRs)
}
