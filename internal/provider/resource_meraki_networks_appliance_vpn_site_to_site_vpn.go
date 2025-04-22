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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceVpnSiteToSiteVpnResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceVpnSiteToSiteVpnResource{}
)

func NewNetworksApplianceVpnSiteToSiteVpnResource() resource.Resource {
	return &NetworksApplianceVpnSiteToSiteVpnResource{}
}

type NetworksApplianceVpnSiteToSiteVpnResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceVpnSiteToSiteVpnResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceVpnSiteToSiteVpnResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_vpn_site_to_site_vpn"
}

func (r *NetworksApplianceVpnSiteToSiteVpnResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"hubs": schema.ListNestedAttribute{
				MarkdownDescription: `The list of VPN hubs, in order of preference.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"hub_id": schema.StringAttribute{
							MarkdownDescription: `The network ID of the hub.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"use_default_route": schema.BoolAttribute{
							MarkdownDescription: `Indicates whether default route traffic should be sent to this hub.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: `The site-to-site VPN mode.
                                  Allowed values: [hub,none,spoke]`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"hub",
						"none",
						"spoke",
					),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"subnet": schema.SingleNestedAttribute{
				MarkdownDescription: `Configuration of subnet features`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"nat": schema.SingleNestedAttribute{
						MarkdownDescription: `Configuration of NAT for subnets`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"is_allowed": schema.BoolAttribute{
								MarkdownDescription: `If enabled, VPN subnet translation can be used to translate any local subnets that are allowed to use the VPN into a new subnet with the same number of addresses.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
				},
			},
			"subnets": schema.SetNestedAttribute{
				MarkdownDescription: `The list of subnets and their VPN presence.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"local_subnet": schema.StringAttribute{
							MarkdownDescription: `The CIDR notation subnet used within the VPN`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"nat": schema.SingleNestedAttribute{
							MarkdownDescription: `Configuration of NAT for the subnet`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"enabled": schema.BoolAttribute{
									MarkdownDescription: `Whether or not VPN subnet translation is enabled for the subnet`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Bool{
										boolplanmodifier.UseStateForUnknown(),
									},
								},
								"remote_subnet": schema.StringAttribute{
									MarkdownDescription: `The translated subnet to be used in the VPN`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
						"use_vpn": schema.BoolAttribute{
							MarkdownDescription: `Indicates the presence of the subnet in the VPN`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
		},
	}
}

func (r *NetworksApplianceVpnSiteToSiteVpnResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceVpnSiteToSiteVpnRs

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

	if vvNetworkID != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceVpnSiteToSiteVpn(vvNetworkID)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksApplianceVpnSiteToSiteVpn  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksApplianceVpnSiteToSiteVpn only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceVpnSiteToSiteVpn(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceVpnSiteToSiteVpn",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceVpnSiteToSiteVpn",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceVpnSiteToSiteVpn(vvNetworkID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceVpnSiteToSiteVpn",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceVpnSiteToSiteVpn",
			err.Error(),
		)
		return
	}

	data = ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *NetworksApplianceVpnSiteToSiteVpnResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceVpnSiteToSiteVpnRs

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
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceVpnSiteToSiteVpn(vvNetworkID)
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
				"Failure when executing GetNetworkApplianceVpnSiteToSiteVpn",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceVpnSiteToSiteVpn",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksApplianceVpnSiteToSiteVpnResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceVpnSiteToSiteVpnResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceVpnSiteToSiteVpnRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceVpnSiteToSiteVpn(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceVpnSiteToSiteVpn",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceVpnSiteToSiteVpn",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceVpnSiteToSiteVpnResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceVpnSiteToSiteVpn", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceVpnSiteToSiteVpnRs struct {
	NetworkID types.String                                                     `tfsdk:"network_id"`
	Hubs      *[]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnHubsRs    `tfsdk:"hubs"`
	Mode      types.String                                                     `tfsdk:"mode"`
	Subnet    *ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetRs    `tfsdk:"subnet"`
	Subnets   *[]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsRs `tfsdk:"subnets"`
}

type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnHubsRs struct {
	HubID           types.String `tfsdk:"hub_id"`
	UseDefaultRoute types.Bool   `tfsdk:"use_default_route"`
}

type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetRs struct {
	Nat *ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetNatRs `tfsdk:"nat"`
}

type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetNatRs struct {
	IsAllowed types.Bool `tfsdk:"is_allowed"`
}

type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsRs struct {
	LocalSubnet types.String                                                      `tfsdk:"local_subnet"`
	Nat         *ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsNatRs `tfsdk:"nat"`
	UseVpn      types.Bool                                                        `tfsdk:"use_vpn"`
}

type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsNatRs struct {
	Enabled      types.Bool   `tfsdk:"enabled"`
	RemoteSubnet types.String `tfsdk:"remote_subnet"`
}

// FromBody
func (r *NetworksApplianceVpnSiteToSiteVpnRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpn {
	emptyString := ""
	var requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnHubs []merakigosdk.RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnHubs

	if r.Hubs != nil {
		for _, rItem1 := range *r.Hubs {
			hubID := rItem1.HubID.ValueString()
			useDefaultRoute := func() *bool {
				if !rItem1.UseDefaultRoute.IsUnknown() && !rItem1.UseDefaultRoute.IsNull() {
					return rItem1.UseDefaultRoute.ValueBoolPointer()
				}
				return nil
			}()
			requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnHubs = append(requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnHubs, merakigosdk.RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnHubs{
				HubID:           hubID,
				UseDefaultRoute: useDefaultRoute,
			})
			//[debug] Is Array: True
		}
	}
	mode := new(string)
	if !r.Mode.IsUnknown() && !r.Mode.IsNull() {
		*mode = r.Mode.ValueString()
	} else {
		mode = &emptyString
	}
	var requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnet *merakigosdk.RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnet

	if r.Subnet != nil {
		var requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnetNat *merakigosdk.RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnetNat

		if r.Subnet.Nat != nil {
			isAllowed := func() *bool {
				if !r.Subnet.Nat.IsAllowed.IsUnknown() && !r.Subnet.Nat.IsAllowed.IsNull() {
					return r.Subnet.Nat.IsAllowed.ValueBoolPointer()
				}
				return nil
			}()
			requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnetNat = &merakigosdk.RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnetNat{
				IsAllowed: isAllowed,
			}
			//[debug] Is Array: False
		}
		requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnet = &merakigosdk.RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnet{
			Nat: requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnetNat,
		}
		//[debug] Is Array: False
	}
	var requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets []merakigosdk.RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets

	if r.Subnets != nil {
		for _, rItem1 := range *r.Subnets {
			localSubnet := rItem1.LocalSubnet.ValueString()
			var requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnetsNat *merakigosdk.RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnetsNat

			if rItem1.Nat != nil {
				enabled := func() *bool {
					if !rItem1.Nat.Enabled.IsUnknown() && !rItem1.Nat.Enabled.IsNull() {
						return rItem1.Nat.Enabled.ValueBoolPointer()
					}
					return nil
				}()
				remoteSubnet := rItem1.Nat.RemoteSubnet.ValueString()
				requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnetsNat = &merakigosdk.RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnetsNat{
					Enabled:      enabled,
					RemoteSubnet: remoteSubnet,
				}
				//[debug] Is Array: False
			}
			useVpn := func() *bool {
				if !rItem1.UseVpn.IsUnknown() && !rItem1.UseVpn.IsNull() {
					return rItem1.UseVpn.ValueBoolPointer()
				}
				return nil
			}()
			requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets = append(requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets, merakigosdk.RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets{
				LocalSubnet: localSubnet,
				Nat:         requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnetsNat,
				UseVpn:      useVpn,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpn{
		Hubs: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnHubs {
			if len(requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnHubs) > 0 {
				return &requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnHubs
			}
			return nil
		}(),
		Mode:   *mode,
		Subnet: requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnet,
		Subnets: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets {
			if len(requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets) > 0 {
				return &requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnItemToBodyRs(state NetworksApplianceVpnSiteToSiteVpnRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpn, is_read bool) NetworksApplianceVpnSiteToSiteVpnRs {
	itemState := NetworksApplianceVpnSiteToSiteVpnRs{
		Hubs: func() *[]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnHubsRs {
			if response.Hubs != nil {
				result := make([]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnHubsRs, len(*response.Hubs))
				for i, hubs := range *response.Hubs {
					result[i] = ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnHubsRs{
						HubID: types.StringValue(hubs.HubID),
						UseDefaultRoute: func() types.Bool {
							if hubs.UseDefaultRoute != nil {
								return types.BoolValue(*hubs.UseDefaultRoute)
							}
							return types.Bool{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		Mode: types.StringValue(response.Mode),
		Subnet: func() *ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetRs {
			if response.Subnet != nil {
				return &ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetRs{
					Nat: func() *ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetNatRs {
						if response.Subnet.Nat != nil {
							return &ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetNatRs{
								IsAllowed: func() types.Bool {
									if response.Subnet.Nat.IsAllowed != nil {
										return types.BoolValue(*response.Subnet.Nat.IsAllowed)
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
		Subnets: func() *[]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsRs {
			if response.Subnets != nil {
				result := make([]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsRs, len(*response.Subnets))
				for i, subnets := range *response.Subnets {
					result[i] = ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsRs{
						LocalSubnet: types.StringValue(subnets.LocalSubnet),
						Nat: func() *ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsNatRs {
							if subnets.Nat != nil {
								return &ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsNatRs{
									Enabled: func() types.Bool {
										if subnets.Nat.Enabled != nil {
											return types.BoolValue(*subnets.Nat.Enabled)
										}
										return types.Bool{}
									}(),
									RemoteSubnet: types.StringValue(subnets.Nat.RemoteSubnet),
								}
							}
							return nil
						}(),
						UseVpn: func() types.Bool {
							if subnets.UseVpn != nil {
								return types.BoolValue(*subnets.UseVpn)
							}
							return types.Bool{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceVpnSiteToSiteVpnRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceVpnSiteToSiteVpnRs)
}
