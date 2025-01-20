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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSettingsResource{}
	_ resource.ResourceWithConfigure = &NetworksSettingsResource{}
)

func NewNetworksSettingsResource() resource.Resource {
	return &NetworksSettingsResource{}
}

type NetworksSettingsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_settings"
}

func (r *NetworksSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"fips": schema.SingleNestedAttribute{
				MarkdownDescription: `A hash of FIPS options applied to the Network`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Enables / disables FIPS on the network.`,
						Computed:            true,
					},
				},
			},
			"local_status_page": schema.SingleNestedAttribute{
				MarkdownDescription: `A hash of Local Status page(s)' authentication options applied to the Network.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"authentication": schema.SingleNestedAttribute{
						MarkdownDescription: `A hash of Local Status page(s)' authentication options applied to the Network.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Enables / disables the authentication on Local Status page(s).`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"password": schema.StringAttribute{
								MarkdownDescription: `The password used for Local Status Page(s). Set this to null to clear the password.`,
								Sensitive:           true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"username": schema.StringAttribute{
								MarkdownDescription: `The username used for Local Status Page(s).`,
								Computed:            true,
							},
						},
					},
				},
			},
			"local_status_page_enabled": schema.BoolAttribute{
				MarkdownDescription: `Enables / disables the local device status pages (<a target='_blank' href='http://my.meraki.com/'>my.meraki.com, </a><a target='_blank' href='http://ap.meraki.com/'>ap.meraki.com, </a><a target='_blank' href='http://switch.meraki.com/'>switch.meraki.com, </a><a target='_blank' href='http://wired.meraki.com/'>wired.meraki.com</a>). Optional (defaults to false)`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"named_vlans": schema.SingleNestedAttribute{
				MarkdownDescription: `A hash of Named VLANs options applied to the Network.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Enables / disables Named VLANs on the Network.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"remote_status_page_enabled": schema.BoolAttribute{
				MarkdownDescription: `Enables / disables access to the device status page (<a target='_blank'>http://[device's LAN IP])</a>. Optional. Can only be set if localStatusPageEnabled is set to true`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"secure_port": schema.SingleNestedAttribute{
				MarkdownDescription: `A hash of SecureConnect options applied to the Network.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Enables / disables SecureConnect on the network. Optional.`,
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

func (r *NetworksSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSettingsRs

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
	//Item
	responseVerifyItem, restyResp1, err := r.client.Networks.GetNetworkSettings(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksSettings only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksSettings only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkSettings(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSettings",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Networks.GetNetworkSettings(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSettings",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseNetworksGetNetworkSettingsItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSettingsRs

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
	responseGet, restyRespGet, err := r.client.Networks.GetNetworkSettings(vvNetworkID)
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
				"Failure when executing GetNetworkSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSettings",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseNetworksGetNetworkSettingsItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSettingsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkSettings(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSettings",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksSettings", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSettingsRs struct {
	NetworkID               types.String                                         `tfsdk:"network_id"`
	Fips                    *ResponseNetworksGetNetworkSettingsFipsRs            `tfsdk:"fips"`
	LocalStatusPage         *ResponseNetworksGetNetworkSettingsLocalStatusPageRs `tfsdk:"local_status_page"`
	LocalStatusPageEnabled  types.Bool                                           `tfsdk:"local_status_page_enabled"`
	NamedVLANs              *ResponseNetworksGetNetworkSettingsNamedVlansRs      `tfsdk:"named_vlans"`
	RemoteStatusPageEnabled types.Bool                                           `tfsdk:"remote_status_page_enabled"`
	SecurePort              *ResponseNetworksGetNetworkSettingsSecurePortRs      `tfsdk:"secure_port"`
}

type ResponseNetworksGetNetworkSettingsFipsRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseNetworksGetNetworkSettingsLocalStatusPageRs struct {
	Authentication *ResponseNetworksGetNetworkSettingsLocalStatusPageAuthenticationRs `tfsdk:"authentication"`
}

type ResponseNetworksGetNetworkSettingsLocalStatusPageAuthenticationRs struct {
	Enabled  types.Bool   `tfsdk:"enabled"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

type ResponseNetworksGetNetworkSettingsNamedVlansRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseNetworksGetNetworkSettingsSecurePortRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// FromBody
func (r *NetworksSettingsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestNetworksUpdateNetworkSettings {
	var requestNetworksUpdateNetworkSettingsLocalStatusPage *merakigosdk.RequestNetworksUpdateNetworkSettingsLocalStatusPage
	if r.LocalStatusPage != nil {
		var requestNetworksUpdateNetworkSettingsLocalStatusPageAuthentication *merakigosdk.RequestNetworksUpdateNetworkSettingsLocalStatusPageAuthentication
		if r.LocalStatusPage.Authentication != nil {
			enabled := func() *bool {
				if !r.LocalStatusPage.Authentication.Enabled.IsUnknown() && !r.LocalStatusPage.Authentication.Enabled.IsNull() {
					return r.LocalStatusPage.Authentication.Enabled.ValueBoolPointer()
				}
				return nil
			}()
			password := r.LocalStatusPage.Authentication.Password.ValueString()
			requestNetworksUpdateNetworkSettingsLocalStatusPageAuthentication = &merakigosdk.RequestNetworksUpdateNetworkSettingsLocalStatusPageAuthentication{
				Enabled:  enabled,
				Password: password,
			}
		}
		requestNetworksUpdateNetworkSettingsLocalStatusPage = &merakigosdk.RequestNetworksUpdateNetworkSettingsLocalStatusPage{
			Authentication: requestNetworksUpdateNetworkSettingsLocalStatusPageAuthentication,
		}
	}
	localStatusPageEnabled := new(bool)
	if !r.LocalStatusPageEnabled.IsUnknown() && !r.LocalStatusPageEnabled.IsNull() {
		*localStatusPageEnabled = r.LocalStatusPageEnabled.ValueBool()
	} else {
		localStatusPageEnabled = nil
	}
	var requestNetworksUpdateNetworkSettingsNamedVLANs *merakigosdk.RequestNetworksUpdateNetworkSettingsNamedVLANs
	if r.NamedVLANs != nil {
		enabled := func() *bool {
			if !r.NamedVLANs.Enabled.IsUnknown() && !r.NamedVLANs.Enabled.IsNull() {
				return r.NamedVLANs.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		requestNetworksUpdateNetworkSettingsNamedVLANs = &merakigosdk.RequestNetworksUpdateNetworkSettingsNamedVLANs{
			Enabled: enabled,
		}
	}
	remoteStatusPageEnabled := new(bool)
	if !r.RemoteStatusPageEnabled.IsUnknown() && !r.RemoteStatusPageEnabled.IsNull() {
		*remoteStatusPageEnabled = r.RemoteStatusPageEnabled.ValueBool()
	} else {
		remoteStatusPageEnabled = nil
	}
	var requestNetworksUpdateNetworkSettingsSecurePort *merakigosdk.RequestNetworksUpdateNetworkSettingsSecurePort
	if r.SecurePort != nil {
		enabled := func() *bool {
			if !r.SecurePort.Enabled.IsUnknown() && !r.SecurePort.Enabled.IsNull() {
				return r.SecurePort.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		requestNetworksUpdateNetworkSettingsSecurePort = &merakigosdk.RequestNetworksUpdateNetworkSettingsSecurePort{
			Enabled: enabled,
		}
	}
	out := merakigosdk.RequestNetworksUpdateNetworkSettings{
		LocalStatusPage:         requestNetworksUpdateNetworkSettingsLocalStatusPage,
		LocalStatusPageEnabled:  localStatusPageEnabled,
		NamedVLANs:              requestNetworksUpdateNetworkSettingsNamedVLANs,
		RemoteStatusPageEnabled: remoteStatusPageEnabled,
		SecurePort:              requestNetworksUpdateNetworkSettingsSecurePort,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseNetworksGetNetworkSettingsItemToBodyRs(state NetworksSettingsRs, response *merakigosdk.ResponseNetworksGetNetworkSettings, is_read bool) NetworksSettingsRs {
	itemState := NetworksSettingsRs{
		Fips: func() *ResponseNetworksGetNetworkSettingsFipsRs {
			if response.Fips != nil {
				return &ResponseNetworksGetNetworkSettingsFipsRs{
					Enabled: func() types.Bool {
						if response.Fips.Enabled != nil {
							return types.BoolValue(*response.Fips.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
		}(),
		LocalStatusPage: func() *ResponseNetworksGetNetworkSettingsLocalStatusPageRs {
			if response.LocalStatusPage != nil {
				return &ResponseNetworksGetNetworkSettingsLocalStatusPageRs{
					Authentication: func() *ResponseNetworksGetNetworkSettingsLocalStatusPageAuthenticationRs {
						if response.LocalStatusPage.Authentication != nil {
							return &ResponseNetworksGetNetworkSettingsLocalStatusPageAuthenticationRs{
								Enabled: func() types.Bool {
									if response.LocalStatusPage.Authentication.Enabled != nil {
										return types.BoolValue(*response.LocalStatusPage.Authentication.Enabled)
									}
									return types.Bool{}
								}(),
								Username: types.StringValue(response.LocalStatusPage.Authentication.Username),
								Password: func() types.String {
									if state.LocalStatusPage != nil {
										if state.LocalStatusPage.Authentication != nil {
											return state.LocalStatusPage.Authentication.Password
										}
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
		LocalStatusPageEnabled: func() types.Bool {
			if response.LocalStatusPageEnabled != nil {
				return types.BoolValue(*response.LocalStatusPageEnabled)
			}
			return types.Bool{}
		}(),
		NamedVLANs: func() *ResponseNetworksGetNetworkSettingsNamedVlansRs {
			if response.NamedVLANs != nil {
				return &ResponseNetworksGetNetworkSettingsNamedVlansRs{
					Enabled: func() types.Bool {
						if response.NamedVLANs.Enabled != nil {
							return types.BoolValue(*response.NamedVLANs.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
		}(),
		RemoteStatusPageEnabled: func() types.Bool {
			if response.RemoteStatusPageEnabled != nil {
				return types.BoolValue(*response.RemoteStatusPageEnabled)
			}
			return types.Bool{}
		}(),
		SecurePort: func() *ResponseNetworksGetNetworkSettingsSecurePortRs {
			if response.SecurePort != nil {
				return &ResponseNetworksGetNetworkSettingsSecurePortRs{
					Enabled: func() types.Bool {
						if response.SecurePort.Enabled != nil {
							return types.BoolValue(*response.SecurePort.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSettingsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSettingsRs)
}
