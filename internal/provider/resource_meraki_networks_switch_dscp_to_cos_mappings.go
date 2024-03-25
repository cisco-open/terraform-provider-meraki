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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSwitchDscpToCosMappingsResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchDscpToCosMappingsResource{}
)

func NewNetworksSwitchDscpToCosMappingsResource() resource.Resource {
	return &NetworksSwitchDscpToCosMappingsResource{}
}

type NetworksSwitchDscpToCosMappingsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchDscpToCosMappingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchDscpToCosMappingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_dscp_to_cos_mappings"
}

func (r *NetworksSwitchDscpToCosMappingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"mappings": schema.SetNestedAttribute{
				MarkdownDescription: `An array of DSCP to CoS mappings. An empty array will reset the mappings to default.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"cos": schema.Int64Attribute{
							MarkdownDescription: `The actual layer-2 CoS queue the DSCP value is mapped to. These are not bits set on outgoing frames. Value can be in the range of 0 to 5 inclusive.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"dscp": schema.Int64Attribute{
							MarkdownDescription: `The Differentiated Services Code Point (DSCP) tag in the IP header that will be mapped to a particular Class-of-Service (CoS) queue. Value can be in the range of 0 to 63 inclusive.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"title": schema.StringAttribute{
							MarkdownDescription: `Label for the mapping (optional).`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
		},
	}
}

func (r *NetworksSwitchDscpToCosMappingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchDscpToCosMappingsRs

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
	// network_id
	//Item
	responseVerifyItem, restyResp1, err := r.client.Switch.GetNetworkSwitchDscpToCosMappings(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksSwitchDscpToCosMappings only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksSwitchDscpToCosMappings only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Switch.UpdateNetworkSwitchDscpToCosMappings(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchDscpToCosMappings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchDscpToCosMappings",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Switch.GetNetworkSwitchDscpToCosMappings(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchDscpToCosMappings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchDscpToCosMappings",
			err.Error(),
		)
		return
	}

	data = ResponseSwitchGetNetworkSwitchDscpToCosMappingsItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchDscpToCosMappingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchDscpToCosMappingsRs

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
	// network_id
	responseGet, restyRespGet, err := r.client.Switch.GetNetworkSwitchDscpToCosMappings(vvNetworkID)
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
				"Failure when executing GetNetworkSwitchDscpToCosMappings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchDscpToCosMappings",
			err.Error(),
		)
		return
	}

	data = ResponseSwitchGetNetworkSwitchDscpToCosMappingsItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksSwitchDscpToCosMappingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksSwitchDscpToCosMappingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSwitchDscpToCosMappingsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Switch.UpdateNetworkSwitchDscpToCosMappings(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchDscpToCosMappings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchDscpToCosMappings",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchDscpToCosMappingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSwitchDscpToCosMappingsRs struct {
	NetworkID types.String                                                 `tfsdk:"network_id"`
	Mappings  *[]ResponseSwitchGetNetworkSwitchDscpToCosMappingsMappingsRs `tfsdk:"mappings"`
}

type ResponseSwitchGetNetworkSwitchDscpToCosMappingsMappingsRs struct {
	Cos   types.Int64  `tfsdk:"cos"`
	Dscp  types.Int64  `tfsdk:"dscp"`
	Title types.String `tfsdk:"title"`
}

// FromBody
func (r *NetworksSwitchDscpToCosMappingsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateNetworkSwitchDscpToCosMappings {
	var requestSwitchUpdateNetworkSwitchDscpToCosMappingsMappings []merakigosdk.RequestSwitchUpdateNetworkSwitchDscpToCosMappingsMappings
	if r.Mappings != nil {
		for _, rItem1 := range *r.Mappings {
			cos := func() *int64 {
				if !rItem1.Cos.IsUnknown() && !rItem1.Cos.IsNull() {
					return rItem1.Cos.ValueInt64Pointer()
				}
				return nil
			}()
			dscp := func() *int64 {
				if !rItem1.Dscp.IsUnknown() && !rItem1.Dscp.IsNull() {
					return rItem1.Dscp.ValueInt64Pointer()
				}
				return nil
			}()
			title := rItem1.Title.ValueString()
			requestSwitchUpdateNetworkSwitchDscpToCosMappingsMappings = append(requestSwitchUpdateNetworkSwitchDscpToCosMappingsMappings, merakigosdk.RequestSwitchUpdateNetworkSwitchDscpToCosMappingsMappings{
				Cos:   int64ToIntPointer(cos),
				Dscp:  int64ToIntPointer(dscp),
				Title: title,
			})
		}
	}
	out := merakigosdk.RequestSwitchUpdateNetworkSwitchDscpToCosMappings{
		Mappings: func() *[]merakigosdk.RequestSwitchUpdateNetworkSwitchDscpToCosMappingsMappings {
			if len(requestSwitchUpdateNetworkSwitchDscpToCosMappingsMappings) > 0 {
				return &requestSwitchUpdateNetworkSwitchDscpToCosMappingsMappings
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchDscpToCosMappingsItemToBodyRs(state NetworksSwitchDscpToCosMappingsRs, response *merakigosdk.ResponseSwitchGetNetworkSwitchDscpToCosMappings, is_read bool) NetworksSwitchDscpToCosMappingsRs {
	itemState := NetworksSwitchDscpToCosMappingsRs{
		Mappings: func() *[]ResponseSwitchGetNetworkSwitchDscpToCosMappingsMappingsRs {
			if response.Mappings != nil {
				result := make([]ResponseSwitchGetNetworkSwitchDscpToCosMappingsMappingsRs, len(*response.Mappings))
				for i, mappings := range *response.Mappings {
					result[i] = ResponseSwitchGetNetworkSwitchDscpToCosMappingsMappingsRs{
						Cos: func() types.Int64 {
							if mappings.Cos != nil {
								return types.Int64Value(int64(*mappings.Cos))
							}
							return types.Int64{}
						}(),
						Dscp: func() types.Int64 {
							if mappings.Dscp != nil {
								return types.Int64Value(int64(*mappings.Dscp))
							}
							return types.Int64{}
						}(),
						Title: types.StringValue(mappings.Title),
					}
				}
				return &result
			}
			return &[]ResponseSwitchGetNetworkSwitchDscpToCosMappingsMappingsRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSwitchDscpToCosMappingsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSwitchDscpToCosMappingsRs)
}
