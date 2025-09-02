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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceSettingsResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceSettingsResource{}
)

func NewNetworksApplianceSettingsResource() resource.Resource {
	return &NetworksApplianceSettingsResource{}
}

type NetworksApplianceSettingsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_settings"
}

func (r *NetworksApplianceSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"client_tracking_method": schema.StringAttribute{
				MarkdownDescription: `Client tracking method of a network
                                  Allowed values: [IP address,MAC address,Unique client identifier]`,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"IP address",
						"MAC address",
						"Unique client identifier",
					),
				},
			},
			"deployment_mode": schema.StringAttribute{
				MarkdownDescription: `Deployment mode of a network
                                  Allowed values: [passthrough,routed]`,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"passthrough",
						"routed",
					),
				},
			},
			"dynamic_dns": schema.SingleNestedAttribute{
				MarkdownDescription: `Dynamic DNS settings for a network`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Dynamic DNS enabled`,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"prefix": schema.StringAttribute{
						MarkdownDescription: `Dynamic DNS url prefix. DDNS must be enabled to update`,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"url": schema.StringAttribute{
						MarkdownDescription: `Dynamic DNS url. DDNS must be enabled to update`,
						Computed:            true,
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

func (r *NetworksApplianceSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceSettingsRs

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
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceSettings(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceSettings",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceSettings",
			err.Error(),
		)
		return
	}

	// Assign data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *NetworksApplianceSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceSettingsRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceSettings(vvNetworkID)
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
				"Failure when executing GetNetworkApplianceSettings",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceSettings",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceSettingsItemToBodyRs(data, responseGet, true)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *NetworksApplianceSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworksApplianceSettingsRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvNetworkID := plan.NetworkID.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceSettings(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceSettings",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceSettings",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworksApplianceSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceSettings", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceSettingsRs struct {
	NetworkID            types.String                                              `tfsdk:"network_id"`
	ClientTrackingMethod types.String                                              `tfsdk:"client_tracking_method"`
	DeploymentMode       types.String                                              `tfsdk:"deployment_mode"`
	DynamicDNS           *ResponseApplianceGetNetworkApplianceSettingsDynamicDnsRs `tfsdk:"dynamic_dns"`
}

type ResponseApplianceGetNetworkApplianceSettingsDynamicDnsRs struct {
	Enabled types.Bool   `tfsdk:"enabled"`
	Prefix  types.String `tfsdk:"prefix"`
	URL     types.String `tfsdk:"url"`
}

// FromBody
func (r *NetworksApplianceSettingsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceSettings {
	emptyString := ""
	clientTrackingMethod := new(string)
	if !r.ClientTrackingMethod.IsUnknown() && !r.ClientTrackingMethod.IsNull() {
		*clientTrackingMethod = r.ClientTrackingMethod.ValueString()
	} else {
		clientTrackingMethod = &emptyString
	}
	deploymentMode := new(string)
	if !r.DeploymentMode.IsUnknown() && !r.DeploymentMode.IsNull() {
		*deploymentMode = r.DeploymentMode.ValueString()
	} else {
		deploymentMode = &emptyString
	}
	var requestApplianceUpdateNetworkApplianceSettingsDynamicDNS *merakigosdk.RequestApplianceUpdateNetworkApplianceSettingsDynamicDNS

	if r.DynamicDNS != nil {
		enabled := func() *bool {
			if !r.DynamicDNS.Enabled.IsUnknown() && !r.DynamicDNS.Enabled.IsNull() {
				return r.DynamicDNS.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		prefix := r.DynamicDNS.Prefix.ValueString()
		requestApplianceUpdateNetworkApplianceSettingsDynamicDNS = &merakigosdk.RequestApplianceUpdateNetworkApplianceSettingsDynamicDNS{
			Enabled: enabled,
			Prefix:  prefix,
		}
		//[debug] Is Array: False
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceSettings{
		ClientTrackingMethod: *clientTrackingMethod,
		DeploymentMode:       *deploymentMode,
		DynamicDNS:           requestApplianceUpdateNetworkApplianceSettingsDynamicDNS,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceSettingsItemToBodyRs(state NetworksApplianceSettingsRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceSettings, is_read bool) NetworksApplianceSettingsRs {
	itemState := NetworksApplianceSettingsRs{
		ClientTrackingMethod: func() types.String {
			if response.ClientTrackingMethod != "" {
				return types.StringValue(response.ClientTrackingMethod)
			}
			return types.String{}
		}(),
		DeploymentMode: func() types.String {
			if response.DeploymentMode != "" {
				return types.StringValue(response.DeploymentMode)
			}
			return types.String{}
		}(),
		DynamicDNS: func() *ResponseApplianceGetNetworkApplianceSettingsDynamicDnsRs {
			if response.DynamicDNS != nil {
				return &ResponseApplianceGetNetworkApplianceSettingsDynamicDnsRs{
					Enabled: func() types.Bool {
						if response.DynamicDNS.Enabled != nil {
							return types.BoolValue(*response.DynamicDNS.Enabled)
						}
						return types.Bool{}
					}(),
					Prefix: func() types.String {
						if response.DynamicDNS.Prefix != "" {
							return types.StringValue(response.DynamicDNS.Prefix)
						}
						return types.String{}
					}(),
					URL: func() types.String {
						if response.DynamicDNS.URL != "" {
							return types.StringValue(response.DynamicDNS.URL)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceSettingsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceSettingsRs)
}
