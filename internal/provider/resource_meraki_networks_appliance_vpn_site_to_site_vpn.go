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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
			"hubs": schema.SetNestedAttribute{
				MarkdownDescription: `The list of VPN hubs, in order of preference.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
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
				MarkdownDescription: `The site-to-site VPN mode.`,
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
	//Has Paths
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	//Item
	responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceVpnSiteToSiteVpn(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceVpnSiteToSiteVpn only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceVpnSiteToSiteVpn only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceVpnSiteToSiteVpn(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
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
	//Item
	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceVpnSiteToSiteVpn(vvNetworkID)
	// Has item and not has items
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
	// network_id
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
	// network_id
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
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceVpnSiteToSiteVpnRs struct {
	NetworkID types.String                                                     `tfsdk:"network_id"`
	Hubs      *[]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnHubsRs    `tfsdk:"hubs"`
	Mode      types.String                                                     `tfsdk:"mode"`
	Subnets   *[]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsRs `tfsdk:"subnets"`
}

type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnHubsRs struct {
	HubID           types.String `tfsdk:"hub_id"`
	UseDefaultRoute types.Bool   `tfsdk:"use_default_route"`
}

type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsRs struct {
	LocalSubnet types.String `tfsdk:"local_subnet"`
	UseVpn      types.Bool   `tfsdk:"use_vpn"`
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
		}
	}
	mode := new(string)
	if !r.Mode.IsUnknown() && !r.Mode.IsNull() {
		*mode = r.Mode.ValueString()
	} else {
		mode = &emptyString
	}
	var requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets []merakigosdk.RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets
	if r.Subnets != nil {
		for _, rItem1 := range *r.Subnets {
			localSubnet := rItem1.LocalSubnet.ValueString()
			useVpn := func() *bool {
				if !rItem1.UseVpn.IsUnknown() && !rItem1.UseVpn.IsNull() {
					return rItem1.UseVpn.ValueBoolPointer()
				}
				return nil
			}()
			requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets = append(requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets, merakigosdk.RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnSubnets{
				LocalSubnet: localSubnet,
				UseVpn:      useVpn,
			})
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpn{
		Hubs: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnHubs {
			if len(requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnHubs) > 0 {
				return &requestApplianceUpdateNetworkApplianceVpnSiteToSiteVpnHubs
			}
			return nil
		}(),
		Mode: *mode,
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
			return &[]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnHubsRs{}
		}(),
		Mode: types.StringValue(response.Mode),
		Subnets: func() *[]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsRs {
			if response.Subnets != nil {
				result := make([]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsRs, len(*response.Subnets))
				for i, subnets := range *response.Subnets {
					result[i] = ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsRs{
						LocalSubnet: types.StringValue(subnets.LocalSubnet),
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
			return &[]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceVpnSiteToSiteVpnRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceVpnSiteToSiteVpnRs)
}
