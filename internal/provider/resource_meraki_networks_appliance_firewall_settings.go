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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceFirewallSettingsResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceFirewallSettingsResource{}
)

func NewNetworksApplianceFirewallSettingsResource() resource.Resource {
	return &NetworksApplianceFirewallSettingsResource{}
}

type NetworksApplianceFirewallSettingsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceFirewallSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceFirewallSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_firewall_settings"
}

func (r *NetworksApplianceFirewallSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"spoofing_protection": schema.SingleNestedAttribute{
				MarkdownDescription: `Spoofing protection settings`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"ip_source_guard": schema.SingleNestedAttribute{
						MarkdownDescription: `IP source address spoofing settings`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"mode": schema.StringAttribute{
								MarkdownDescription: `Mode of protection
                                              Allowed values: [block,log]`,
								Optional: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"block",
										"log",
									),
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *NetworksApplianceFirewallSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceFirewallSettingsRs

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
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceFirewallSettings(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceFirewallSettings",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceFirewallSettings",
			err.Error(),
		)
		return
	}

	// Assign data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *NetworksApplianceFirewallSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceFirewallSettingsRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceFirewallSettings(vvNetworkID)
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
				"Failure when executing GetNetworkApplianceFirewallSettings",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceFirewallSettings",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceFirewallSettingsItemToBodyRs(data, responseGet, true)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *NetworksApplianceFirewallSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceFirewallSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworksApplianceFirewallSettingsRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvNetworkID := plan.NetworkID.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceFirewallSettings(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceFirewallSettings",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceFirewallSettings",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworksApplianceFirewallSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceFirewallSettings", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceFirewallSettingsRs struct {
	NetworkID          types.String                                                              `tfsdk:"network_id"`
	SpoofingProtection *ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionRs `tfsdk:"spoofing_protection"`
}

type ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionRs struct {
	IPSourceGuard *ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionIpSourceGuardRs `tfsdk:"ip_source_guard"`
}

type ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionIpSourceGuardRs struct {
	Mode types.String `tfsdk:"mode"`
}

// FromBody
func (r *NetworksApplianceFirewallSettingsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallSettings {
	var requestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtection *merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtection

	if r.SpoofingProtection != nil {
		var requestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtectionIPSourceGuard *merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtectionIPSourceGuard

		if r.SpoofingProtection.IPSourceGuard != nil {
			mode := r.SpoofingProtection.IPSourceGuard.Mode.ValueString()
			requestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtectionIPSourceGuard = &merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtectionIPSourceGuard{
				Mode: mode,
			}
			//[debug] Is Array: False
		}
		requestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtection = &merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtection{
			IPSourceGuard: requestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtectionIPSourceGuard,
		}
		//[debug] Is Array: False
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallSettings{
		SpoofingProtection: requestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtection,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceFirewallSettingsItemToBodyRs(state NetworksApplianceFirewallSettingsRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceFirewallSettings, is_read bool) NetworksApplianceFirewallSettingsRs {
	itemState := NetworksApplianceFirewallSettingsRs{
		SpoofingProtection: func() *ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionRs {
			if response.SpoofingProtection != nil {
				return &ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionRs{
					IPSourceGuard: func() *ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionIpSourceGuardRs {
						if response.SpoofingProtection.IPSourceGuard != nil {
							return &ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionIpSourceGuardRs{
								Mode: func() types.String {
									if response.SpoofingProtection.IPSourceGuard.Mode != "" {
										return types.StringValue(response.SpoofingProtection.IPSourceGuard.Mode)
									}
									return types.String{}
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
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceFirewallSettingsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceFirewallSettingsRs)
}
