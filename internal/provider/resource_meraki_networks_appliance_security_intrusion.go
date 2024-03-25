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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceSecurityIntrusionResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceSecurityIntrusionResource{}
)

func NewNetworksApplianceSecurityIntrusionResource() resource.Resource {
	return &NetworksApplianceSecurityIntrusionResource{}
}

type NetworksApplianceSecurityIntrusionResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceSecurityIntrusionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceSecurityIntrusionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_security_intrusion"
}

func (r *NetworksApplianceSecurityIntrusionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ids_rulesets": schema.StringAttribute{
				MarkdownDescription: `Set the detection ruleset 'connectivity'/'balanced'/'security' (optional - omitting will leave current config unchanged). Default value is 'balanced' if none currently saved`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: `Set mode to 'disabled'/'detection'/'prevention' (optional - omitting will leave current config unchanged)`,
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
			"protected_networks": schema.SingleNestedAttribute{
				MarkdownDescription: `Set the included/excluded networks from the intrusion engine (optional - omitting will leave current config unchanged). This is available only in 'passthrough' mode`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"excluded_cidr": schema.SetAttribute{
						MarkdownDescription: `list of IP addresses or subnets being excluded from protection (required if 'useDefault' is false)`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},

						ElementType: types.StringType,
					},
					"included_cidr": schema.SetAttribute{
						MarkdownDescription: `list of IP addresses or subnets being protected (required if 'useDefault' is false)`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},

						ElementType: types.StringType,
					},
					"use_default": schema.BoolAttribute{
						MarkdownDescription: `true/false whether to use special IPv4 addresses: https://tools.ietf.org/html/rfc5735 (required). Default value is true if none currently saved`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
		},
	}
}

func (r *NetworksApplianceSecurityIntrusionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceSecurityIntrusionRs

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
	responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceSecurityIntrusion(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceSecurityIntrusion only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceSecurityIntrusion only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceSecurityIntrusion(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceSecurityIntrusion",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceSecurityIntrusion",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceSecurityIntrusion(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceSecurityIntrusion",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceSecurityIntrusion",
			err.Error(),
		)
		return
	}

	data = ResponseApplianceGetNetworkApplianceSecurityIntrusionItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceSecurityIntrusionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceSecurityIntrusionRs

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
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceSecurityIntrusion(vvNetworkID)
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
				"Failure when executing GetNetworkApplianceSecurityIntrusion",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceSecurityIntrusion",
			err.Error(),
		)
		return
	}

	data = ResponseApplianceGetNetworkApplianceSecurityIntrusionItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksApplianceSecurityIntrusionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceSecurityIntrusionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceSecurityIntrusionRs
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
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceSecurityIntrusion(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceSecurityIntrusion",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceSecurityIntrusion",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceSecurityIntrusionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceSecurityIntrusionRs struct {
	NetworkID         types.String                                                              `tfsdk:"network_id"`
	IDsRulesets       types.String                                                              `tfsdk:"ids_rulesets"`
	Mode              types.String                                                              `tfsdk:"mode"`
	ProtectedNetworks *ResponseApplianceGetNetworkApplianceSecurityIntrusionProtectedNetworksRs `tfsdk:"protected_networks"`
}

type ResponseApplianceGetNetworkApplianceSecurityIntrusionProtectedNetworksRs struct {
	ExcludedCidr types.Set  `tfsdk:"excluded_cidr"`
	IncludedCidr types.Set  `tfsdk:"included_cidr"`
	UseDefault   types.Bool `tfsdk:"use_default"`
}

// FromBody
func (r *NetworksApplianceSecurityIntrusionRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceSecurityIntrusion {
	emptyString := ""
	iDsRulesets := new(string)
	if !r.IDsRulesets.IsUnknown() && !r.IDsRulesets.IsNull() {
		*iDsRulesets = r.IDsRulesets.ValueString()
	} else {
		iDsRulesets = &emptyString
	}
	mode := new(string)
	if !r.Mode.IsUnknown() && !r.Mode.IsNull() {
		*mode = r.Mode.ValueString()
	} else {
		mode = &emptyString
	}
	var requestApplianceUpdateNetworkApplianceSecurityIntrusionProtectedNetworks *merakigosdk.RequestApplianceUpdateNetworkApplianceSecurityIntrusionProtectedNetworks
	if r.ProtectedNetworks != nil {
		var excludedCidr []string = nil

		r.ProtectedNetworks.ExcludedCidr.ElementsAs(ctx, &excludedCidr, false)
		var includedCidr []string = nil

		r.ProtectedNetworks.IncludedCidr.ElementsAs(ctx, &includedCidr, false)
		useDefault := func() *bool {
			if !r.ProtectedNetworks.UseDefault.IsUnknown() && !r.ProtectedNetworks.UseDefault.IsNull() {
				return r.ProtectedNetworks.UseDefault.ValueBoolPointer()
			}
			return nil
		}()
		requestApplianceUpdateNetworkApplianceSecurityIntrusionProtectedNetworks = &merakigosdk.RequestApplianceUpdateNetworkApplianceSecurityIntrusionProtectedNetworks{
			ExcludedCidr: excludedCidr,
			IncludedCidr: includedCidr,
			UseDefault:   useDefault,
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceSecurityIntrusion{
		IDsRulesets:       *iDsRulesets,
		Mode:              *mode,
		ProtectedNetworks: requestApplianceUpdateNetworkApplianceSecurityIntrusionProtectedNetworks,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceSecurityIntrusionItemToBodyRs(state NetworksApplianceSecurityIntrusionRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceSecurityIntrusion, is_read bool) NetworksApplianceSecurityIntrusionRs {
	itemState := NetworksApplianceSecurityIntrusionRs{
		IDsRulesets: types.StringValue(response.IDsRulesets),
		Mode:        types.StringValue(response.Mode),
		ProtectedNetworks: func() *ResponseApplianceGetNetworkApplianceSecurityIntrusionProtectedNetworksRs {
			if response.ProtectedNetworks != nil {
				return &ResponseApplianceGetNetworkApplianceSecurityIntrusionProtectedNetworksRs{
					ExcludedCidr: StringSliceToSet(response.ProtectedNetworks.ExcludedCidr),
					IncludedCidr: StringSliceToSet(response.ProtectedNetworks.IncludedCidr),
					UseDefault: func() types.Bool {
						if response.ProtectedNetworks.UseDefault != nil {
							return types.BoolValue(*response.ProtectedNetworks.UseDefault)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseApplianceGetNetworkApplianceSecurityIntrusionProtectedNetworksRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceSecurityIntrusionRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceSecurityIntrusionRs)
}
