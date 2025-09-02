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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWirelessSettingsResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessSettingsResource{}
)

func NewNetworksWirelessSettingsResource() resource.Resource {
	return &NetworksWirelessSettingsResource{}
}

type NetworksWirelessSettingsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_settings"
}

func (r *NetworksWirelessSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ipv6_bridge_enabled": schema.BoolAttribute{
				MarkdownDescription: `Toggle for enabling or disabling IPv6 bridging in a network (Note: if enabled, SSIDs must also be configured to use bridge mode)`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"led_lights_on": schema.BoolAttribute{
				MarkdownDescription: `Toggle for enabling or disabling LED lights on all APs in the network (making them run dark)`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"location_analytics_enabled": schema.BoolAttribute{
				MarkdownDescription: `Toggle for enabling or disabling location analytics for your network`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"meshing_enabled": schema.BoolAttribute{
				MarkdownDescription: `Toggle for enabling or disabling meshing in a network`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"named_vlans": schema.SingleNestedAttribute{
				MarkdownDescription: `Named VLAN settings for wireless networks.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"pool_dhcp_monitoring": schema.SingleNestedAttribute{
						MarkdownDescription: `Named VLAN Pool DHCP Monitoring settings.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"duration": schema.Int64Attribute{
								MarkdownDescription: `The duration in minutes that devices will refrain from using dirty VLANs before adding them back to the pool.`,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
							},
							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Whether or not devices using named VLAN pools should remove dirty VLANs from the pool, thereby preventing clients from being assigned to VLANs where they would be unable to obtain an IP address via DHCP`,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"regulatory_domain": schema.SingleNestedAttribute{
				MarkdownDescription: `Regulatory domain information for this network.`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{

					"country_code": schema.StringAttribute{
						MarkdownDescription: `The country code of the regulatory domain.`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name of the regulatory domain for this network.`,
						Computed:            true,
					},
					"permits6e": schema.BoolAttribute{
						MarkdownDescription: `Whether or not the regulatory domain for this network permits Wifi 6E.`,
						Computed:            true,
					},
				},
			},
			"upgradestrategy": schema.StringAttribute{
				MarkdownDescription: `The default strategy that network devices will use to perform an upgrade. Requires firmware version MR 26.8 or higher.
                                  Allowed values: [minimizeClientDowntime,minimizeUpgradeTime]`,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"minimizeClientDowntime",
						"minimizeUpgradeTime",
					),
				},
			},
		},
	}
}

func (r *NetworksWirelessSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessSettingsRs

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
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSettings(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSettings",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSettings",
			err.Error(),
		)
		return
	}

	// Assign data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *NetworksWirelessSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessSettingsRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	responseGet, restyRespGet, err := r.client.Wireless.GetNetworkWirelessSettings(vvNetworkID)
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
				"Failure when executing GetNetworkWirelessSettings",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSettings",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseWirelessGetNetworkWirelessSettingsItemToBodyRs(data, responseGet, true)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *NetworksWirelessSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksWirelessSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworksWirelessSettingsRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvNetworkID := plan.NetworkID.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSettings(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSettings",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSettings",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworksWirelessSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksWirelessSettings", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessSettingsRs struct {
	NetworkID                types.String                                                  `tfsdk:"network_id"`
	IPv6BridgeEnabled        types.Bool                                                    `tfsdk:"ipv6_bridge_enabled"`
	LedLightsOn              types.Bool                                                    `tfsdk:"led_lights_on"`
	LocationAnalyticsEnabled types.Bool                                                    `tfsdk:"location_analytics_enabled"`
	MeshingEnabled           types.Bool                                                    `tfsdk:"meshing_enabled"`
	NamedVLANs               *ResponseWirelessGetNetworkWirelessSettingsNamedVlansRs       `tfsdk:"named_vlans"`
	RegulatoryDomain         *ResponseWirelessGetNetworkWirelessSettingsRegulatoryDomainRs `tfsdk:"regulatory_domain"`
	Upgradestrategy          types.String                                                  `tfsdk:"upgradestrategy"`
}

type ResponseWirelessGetNetworkWirelessSettingsNamedVlansRs struct {
	PoolDhcpMonitoring *ResponseWirelessGetNetworkWirelessSettingsNamedVlansPoolDhcpMonitoringRs `tfsdk:"pool_dhcp_monitoring"`
}

type ResponseWirelessGetNetworkWirelessSettingsNamedVlansPoolDhcpMonitoringRs struct {
	Duration types.Int64 `tfsdk:"duration"`
	Enabled  types.Bool  `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessSettingsRegulatoryDomainRs struct {
	CountryCode types.String `tfsdk:"country_code"`
	Name        types.String `tfsdk:"name"`
	Permits6E   types.Bool   `tfsdk:"permits6e"`
}

// FromBody
func (r *NetworksWirelessSettingsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessSettings {
	emptyString := ""
	iPv6BridgeEnabled := new(bool)
	if !r.IPv6BridgeEnabled.IsUnknown() && !r.IPv6BridgeEnabled.IsNull() {
		*iPv6BridgeEnabled = r.IPv6BridgeEnabled.ValueBool()
	} else {
		iPv6BridgeEnabled = nil
	}
	ledLightsOn := new(bool)
	if !r.LedLightsOn.IsUnknown() && !r.LedLightsOn.IsNull() {
		*ledLightsOn = r.LedLightsOn.ValueBool()
	} else {
		ledLightsOn = nil
	}
	locationAnalyticsEnabled := new(bool)
	if !r.LocationAnalyticsEnabled.IsUnknown() && !r.LocationAnalyticsEnabled.IsNull() {
		*locationAnalyticsEnabled = r.LocationAnalyticsEnabled.ValueBool()
	} else {
		locationAnalyticsEnabled = nil
	}
	meshingEnabled := new(bool)
	if !r.MeshingEnabled.IsUnknown() && !r.MeshingEnabled.IsNull() {
		*meshingEnabled = r.MeshingEnabled.ValueBool()
	} else {
		meshingEnabled = nil
	}
	var requestWirelessUpdateNetworkWirelessSettingsNamedVLANs *merakigosdk.RequestWirelessUpdateNetworkWirelessSettingsNamedVLANs

	if r.NamedVLANs != nil {
		var requestWirelessUpdateNetworkWirelessSettingsNamedVLANsPoolDhcpMonitoring *merakigosdk.RequestWirelessUpdateNetworkWirelessSettingsNamedVLANsPoolDhcpMonitoring

		if r.NamedVLANs.PoolDhcpMonitoring != nil {
			duration := func() *int64 {
				if !r.NamedVLANs.PoolDhcpMonitoring.Duration.IsUnknown() && !r.NamedVLANs.PoolDhcpMonitoring.Duration.IsNull() {
					return r.NamedVLANs.PoolDhcpMonitoring.Duration.ValueInt64Pointer()
				}
				return nil
			}()
			enabled := func() *bool {
				if !r.NamedVLANs.PoolDhcpMonitoring.Enabled.IsUnknown() && !r.NamedVLANs.PoolDhcpMonitoring.Enabled.IsNull() {
					return r.NamedVLANs.PoolDhcpMonitoring.Enabled.ValueBoolPointer()
				}
				return nil
			}()
			requestWirelessUpdateNetworkWirelessSettingsNamedVLANsPoolDhcpMonitoring = &merakigosdk.RequestWirelessUpdateNetworkWirelessSettingsNamedVLANsPoolDhcpMonitoring{
				Duration: int64ToIntPointer(duration),
				Enabled:  enabled,
			}
			//[debug] Is Array: False
		}
		requestWirelessUpdateNetworkWirelessSettingsNamedVLANs = &merakigosdk.RequestWirelessUpdateNetworkWirelessSettingsNamedVLANs{
			PoolDhcpMonitoring: requestWirelessUpdateNetworkWirelessSettingsNamedVLANsPoolDhcpMonitoring,
		}
		//[debug] Is Array: False
	}
	upgradestrategy := new(string)
	if !r.Upgradestrategy.IsUnknown() && !r.Upgradestrategy.IsNull() {
		*upgradestrategy = r.Upgradestrategy.ValueString()
	} else {
		upgradestrategy = &emptyString
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessSettings{
		IPv6BridgeEnabled:        iPv6BridgeEnabled,
		LedLightsOn:              ledLightsOn,
		LocationAnalyticsEnabled: locationAnalyticsEnabled,
		MeshingEnabled:           meshingEnabled,
		NamedVLANs:               requestWirelessUpdateNetworkWirelessSettingsNamedVLANs,
		Upgradestrategy:          *upgradestrategy,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetNetworkWirelessSettingsItemToBodyRs(state NetworksWirelessSettingsRs, response *merakigosdk.ResponseWirelessGetNetworkWirelessSettings, is_read bool) NetworksWirelessSettingsRs {
	itemState := NetworksWirelessSettingsRs{
		IPv6BridgeEnabled: func() types.Bool {
			if response.IPv6BridgeEnabled != nil {
				return types.BoolValue(*response.IPv6BridgeEnabled)
			}
			return types.Bool{}
		}(),
		LedLightsOn: func() types.Bool {
			if response.LedLightsOn != nil {
				return types.BoolValue(*response.LedLightsOn)
			}
			return types.Bool{}
		}(),
		LocationAnalyticsEnabled: func() types.Bool {
			if response.LocationAnalyticsEnabled != nil {
				return types.BoolValue(*response.LocationAnalyticsEnabled)
			}
			return types.Bool{}
		}(),
		MeshingEnabled: func() types.Bool {
			if response.MeshingEnabled != nil {
				return types.BoolValue(*response.MeshingEnabled)
			}
			return types.Bool{}
		}(),
		NamedVLANs: func() *ResponseWirelessGetNetworkWirelessSettingsNamedVlansRs {
			if response.NamedVLANs != nil && state.NamedVLANs != nil {
				return &ResponseWirelessGetNetworkWirelessSettingsNamedVlansRs{
					PoolDhcpMonitoring: func() *ResponseWirelessGetNetworkWirelessSettingsNamedVlansPoolDhcpMonitoringRs {
						if response.NamedVLANs.PoolDhcpMonitoring != nil {
							return &ResponseWirelessGetNetworkWirelessSettingsNamedVlansPoolDhcpMonitoringRs{
								Duration: func() types.Int64 {
									if response.NamedVLANs.PoolDhcpMonitoring.Duration != nil {
										return types.Int64Value(int64(*response.NamedVLANs.PoolDhcpMonitoring.Duration))
									}
									return types.Int64{}
								}(),
								Enabled: func() types.Bool {
									if response.NamedVLANs.PoolDhcpMonitoring.Enabled != nil {
										return types.BoolValue(*response.NamedVLANs.PoolDhcpMonitoring.Enabled)
									}
									return types.Bool{}
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		RegulatoryDomain: func() *ResponseWirelessGetNetworkWirelessSettingsRegulatoryDomainRs {
			if response.RegulatoryDomain != nil {
				return &ResponseWirelessGetNetworkWirelessSettingsRegulatoryDomainRs{
					CountryCode: func() types.String {
						if response.RegulatoryDomain.CountryCode != "" {
							return types.StringValue(response.RegulatoryDomain.CountryCode)
						}
						return types.StringNull()
					}(),
					Name: func() types.String {
						if response.RegulatoryDomain.Name != "" {
							return types.StringValue(response.RegulatoryDomain.Name)
						}
						return types.StringNull()
					}(),
					Permits6E: func() types.Bool {
						if response.RegulatoryDomain.Permits6E != nil {
							return types.BoolValue(*response.RegulatoryDomain.Permits6E)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
		}(),
		Upgradestrategy: func() types.String {
			if response.Upgradestrategy != "" {
				return types.StringValue(response.Upgradestrategy)
			}
			return types.StringNull()
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessSettingsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessSettingsRs)
}
