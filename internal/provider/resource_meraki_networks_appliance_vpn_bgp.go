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

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceVpnBgpResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceVpnBgpResource{}
)

func NewNetworksApplianceVpnBgpResource() resource.Resource {
	return &NetworksApplianceVpnBgpResource{}
}

type NetworksApplianceVpnBgpResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceVpnBgpResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceVpnBgpResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_vpn_bgp"
}

func (r *NetworksApplianceVpnBgpResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"as_number": schema.Int64Attribute{
				MarkdownDescription: `The number of the Autonomous System to which the appliance belongs`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `Whether BGP is enabled on the appliance`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"ibgp_hold_timer": schema.Int64Attribute{
				MarkdownDescription: `The iBGP hold time in seconds`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"neighbors": schema.ListNestedAttribute{
				MarkdownDescription: `List of eBGP neighbor configurations`,
				Optional:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"allow_transit": schema.BoolAttribute{
							MarkdownDescription: `Whether the appliance will advertise routes learned from other Autonomous Systems`,
							Optional:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
						"authentication": schema.SingleNestedAttribute{
							MarkdownDescription: `Authentication settings between BGP peers`,
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"password": schema.StringAttribute{
									MarkdownDescription: `Password to configure MD5 authentication between BGP peers`,
									Sensitive:           true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
						"ebgp_hold_timer": schema.Int64Attribute{
							MarkdownDescription: `The eBGP hold time in seconds for the neighbor`,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"ebgp_multihop": schema.Int64Attribute{
							MarkdownDescription: `The number of hops the appliance must traverse to establish a peering relationship with the neighbor`,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"ip": schema.StringAttribute{
							MarkdownDescription: `The IPv4 address of the neighbor`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"ipv6": schema.SingleNestedAttribute{
							MarkdownDescription: `Information regarding IPv6 address of the neighbor`,
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"address": schema.StringAttribute{
									MarkdownDescription: `The IPv6 address of the neighbor`,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
						"next_hop_ip": schema.StringAttribute{
							MarkdownDescription: `The IPv4 address of the neighbor that will establish a TCP session with the appliance`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"receive_limit": schema.Int64Attribute{
							MarkdownDescription: `The maximum number of routes that the appliance can receive from the neighbor`,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"remote_as_number": schema.Int64Attribute{
							MarkdownDescription: `Remote AS number of the neighbor`,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: `The output interface the appliance uses to establish a peering relationship with the neighbor`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"ttl_security": schema.SingleNestedAttribute{
							MarkdownDescription: `Settings for BGP TTL security to protect BGP peering sessions from forged IP attacks`,
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"enabled": schema.BoolAttribute{
									MarkdownDescription: `Whether BGP TTL security is enabled`,
									Optional:            true,
									PlanModifiers: []planmodifier.Bool{
										boolplanmodifier.UseStateForUnknown(),
									},
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
		},
	}
}

func (r *NetworksApplianceVpnBgpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceVpnBgpRs

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
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceVpnBgp(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceVpnBgp",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceVpnBgp",
			err.Error(),
		)
		return
	}

	// Assign data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *NetworksApplianceVpnBgpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceVpnBgpRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceVpnBgp(vvNetworkID)
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
				"Failure when executing GetNetworkApplianceVpnBgp",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceVpnBgp",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceVpnBgpItemToBodyRs(data, responseGet, true)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *NetworksApplianceVpnBgpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceVpnBgpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworksApplianceVpnBgpRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvNetworkID := plan.NetworkID.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceVpnBgp(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceVpnBgp",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceVpnBgp",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworksApplianceVpnBgpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceVpnBgp", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceVpnBgpRs struct {
	NetworkID     types.String                                             `tfsdk:"network_id"`
	AsNumber      types.Int64                                              `tfsdk:"as_number"`
	Enabled       types.Bool                                               `tfsdk:"enabled"`
	IbgpHoldTimer types.Int64                                              `tfsdk:"ibgp_hold_timer"`
	Neighbors     *[]ResponseApplianceGetNetworkApplianceVpnBgpNeighborsRs `tfsdk:"neighbors"`
}

type ResponseApplianceGetNetworkApplianceVpnBgpNeighborsRs struct {
	AllowTransit    types.Bool                                                           `tfsdk:"allow_transit"`
	Authentication  *ResponseApplianceGetNetworkApplianceVpnBgpNeighborsAuthenticationRs `tfsdk:"authentication"`
	EbgpHoldTimer   types.Int64                                                          `tfsdk:"ebgp_hold_timer"`
	EbgpMultihop    types.Int64                                                          `tfsdk:"ebgp_multihop"`
	IP              types.String                                                         `tfsdk:"ip"`
	IPv6            *ResponseApplianceGetNetworkApplianceVpnBgpNeighborsIpv6Rs           `tfsdk:"ipv6"`
	NextHopIP       types.String                                                         `tfsdk:"next_hop_ip"`
	ReceiveLimit    types.Int64                                                          `tfsdk:"receive_limit"`
	RemoteAsNumber  types.Int64                                                          `tfsdk:"remote_as_number"`
	SourceInterface types.String                                                         `tfsdk:"source_interface"`
	TtlSecurity     *ResponseApplianceGetNetworkApplianceVpnBgpNeighborsTtlSecurityRs    `tfsdk:"ttl_security"`
}

type ResponseApplianceGetNetworkApplianceVpnBgpNeighborsAuthenticationRs struct {
	Password types.String `tfsdk:"password"`
}

type ResponseApplianceGetNetworkApplianceVpnBgpNeighborsIpv6Rs struct {
	Address types.String `tfsdk:"address"`
}

type ResponseApplianceGetNetworkApplianceVpnBgpNeighborsTtlSecurityRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// FromBody
func (r *NetworksApplianceVpnBgpRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgp {
	asNumber := new(int64)
	if !r.AsNumber.IsUnknown() && !r.AsNumber.IsNull() {
		*asNumber = r.AsNumber.ValueInt64()
	} else {
		asNumber = nil
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	ibgpHoldTimer := new(int64)
	if !r.IbgpHoldTimer.IsUnknown() && !r.IbgpHoldTimer.IsNull() {
		*ibgpHoldTimer = r.IbgpHoldTimer.ValueInt64()
	} else {
		ibgpHoldTimer = nil
	}
	var requestApplianceUpdateNetworkApplianceVpnBgpNeighbors []merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgpNeighbors

	if r.Neighbors != nil {
		for _, rItem1 := range *r.Neighbors {
			allowTransit := func() *bool {
				if !rItem1.AllowTransit.IsUnknown() && !rItem1.AllowTransit.IsNull() {
					return rItem1.AllowTransit.ValueBoolPointer()
				}
				return nil
			}()
			var requestApplianceUpdateNetworkApplianceVpnBgpNeighborsAuthentication *merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsAuthentication

			if rItem1.Authentication != nil {
				password := rItem1.Authentication.Password.ValueString()
				requestApplianceUpdateNetworkApplianceVpnBgpNeighborsAuthentication = &merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsAuthentication{
					Password: password,
				}
				//[debug] Is Array: False
			}
			ebgpHoldTimer := func() *int64 {
				if !rItem1.EbgpHoldTimer.IsUnknown() && !rItem1.EbgpHoldTimer.IsNull() {
					return rItem1.EbgpHoldTimer.ValueInt64Pointer()
				}
				return nil
			}()
			ebgpMultihop := func() *int64 {
				if !rItem1.EbgpMultihop.IsUnknown() && !rItem1.EbgpMultihop.IsNull() {
					return rItem1.EbgpMultihop.ValueInt64Pointer()
				}
				return nil
			}()
			ip := rItem1.IP.ValueString()
			var requestApplianceUpdateNetworkApplianceVpnBgpNeighborsIPv6 *merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsIPv6

			if rItem1.IPv6 != nil {
				address := rItem1.IPv6.Address.ValueString()
				requestApplianceUpdateNetworkApplianceVpnBgpNeighborsIPv6 = &merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsIPv6{
					Address: address,
				}
				//[debug] Is Array: False
			}
			nextHopIP := rItem1.NextHopIP.ValueString()
			receiveLimit := func() *int64 {
				if !rItem1.ReceiveLimit.IsUnknown() && !rItem1.ReceiveLimit.IsNull() {
					return rItem1.ReceiveLimit.ValueInt64Pointer()
				}
				return nil
			}()
			remoteAsNumber := func() *int64 {
				if !rItem1.RemoteAsNumber.IsUnknown() && !rItem1.RemoteAsNumber.IsNull() {
					return rItem1.RemoteAsNumber.ValueInt64Pointer()
				}
				return nil
			}()
			sourceInterface := rItem1.SourceInterface.ValueString()
			var requestApplianceUpdateNetworkApplianceVpnBgpNeighborsTtlSecurity *merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsTtlSecurity

			if rItem1.TtlSecurity != nil {
				enabled := func() *bool {
					if !rItem1.TtlSecurity.Enabled.IsUnknown() && !rItem1.TtlSecurity.Enabled.IsNull() {
						return rItem1.TtlSecurity.Enabled.ValueBoolPointer()
					}
					return nil
				}()
				requestApplianceUpdateNetworkApplianceVpnBgpNeighborsTtlSecurity = &merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsTtlSecurity{
					Enabled: enabled,
				}
				//[debug] Is Array: False
			}
			requestApplianceUpdateNetworkApplianceVpnBgpNeighbors = append(requestApplianceUpdateNetworkApplianceVpnBgpNeighbors, merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgpNeighbors{
				AllowTransit:    allowTransit,
				Authentication:  requestApplianceUpdateNetworkApplianceVpnBgpNeighborsAuthentication,
				EbgpHoldTimer:   int64ToIntPointer(ebgpHoldTimer),
				EbgpMultihop:    int64ToIntPointer(ebgpMultihop),
				IP:              ip,
				IPv6:            requestApplianceUpdateNetworkApplianceVpnBgpNeighborsIPv6,
				NextHopIP:       nextHopIP,
				ReceiveLimit:    int64ToIntPointer(receiveLimit),
				RemoteAsNumber:  int64ToIntPointer(remoteAsNumber),
				SourceInterface: sourceInterface,
				TtlSecurity:     requestApplianceUpdateNetworkApplianceVpnBgpNeighborsTtlSecurity,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgp{
		AsNumber:      int64ToIntPointer(asNumber),
		Enabled:       enabled,
		IbgpHoldTimer: int64ToIntPointer(ibgpHoldTimer),
		Neighbors: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgpNeighbors {
			if len(requestApplianceUpdateNetworkApplianceVpnBgpNeighbors) > 0 {
				return &requestApplianceUpdateNetworkApplianceVpnBgpNeighbors
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceVpnBgpItemToBodyRs(state NetworksApplianceVpnBgpRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceVpnBgp, is_read bool) NetworksApplianceVpnBgpRs {
	itemState := NetworksApplianceVpnBgpRs{
		AsNumber: func() types.Int64 {
			if response.AsNumber != nil {
				return types.Int64Value(int64(*response.AsNumber))
			}
			return types.Int64{}
		}(),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		IbgpHoldTimer: func() types.Int64 {
			if response.IbgpHoldTimer != nil {
				return types.Int64Value(int64(*response.IbgpHoldTimer))
			}
			return types.Int64{}
		}(),
		Neighbors: func() *[]ResponseApplianceGetNetworkApplianceVpnBgpNeighborsRs {
			if response.Neighbors != nil {
				result := make([]ResponseApplianceGetNetworkApplianceVpnBgpNeighborsRs, len(*response.Neighbors))
				for i, neighbors := range *response.Neighbors {
					result[i] = ResponseApplianceGetNetworkApplianceVpnBgpNeighborsRs{
						AllowTransit: func() types.Bool {
							if neighbors.AllowTransit != nil {
								return types.BoolValue(*neighbors.AllowTransit)
							}
							return types.Bool{}
						}(),
						Authentication: func() *ResponseApplianceGetNetworkApplianceVpnBgpNeighborsAuthenticationRs {
							if neighbors.Authentication != nil {
								return &ResponseApplianceGetNetworkApplianceVpnBgpNeighborsAuthenticationRs{
									Password: func() types.String {
										if neighbors.Authentication.Password != "" {
											return types.StringValue(neighbors.Authentication.Password)
										}
										return types.String{}
									}(),
								}
							}
							return nil
						}(),
						EbgpHoldTimer: func() types.Int64 {
							if neighbors.EbgpHoldTimer != nil {
								return types.Int64Value(int64(*neighbors.EbgpHoldTimer))
							}
							return types.Int64{}
						}(),
						EbgpMultihop: func() types.Int64 {
							if neighbors.EbgpMultihop != nil {
								return types.Int64Value(int64(*neighbors.EbgpMultihop))
							}
							return types.Int64{}
						}(),
						IP: func() types.String {
							if neighbors.IP != "" {
								return types.StringValue(neighbors.IP)
							}
							return types.String{}
						}(),
						IPv6: func() *ResponseApplianceGetNetworkApplianceVpnBgpNeighborsIpv6Rs {
							if neighbors.IPv6 != nil {
								return &ResponseApplianceGetNetworkApplianceVpnBgpNeighborsIpv6Rs{
									Address: func() types.String {
										if neighbors.IPv6.Address != "" {
											return types.StringValue(neighbors.IPv6.Address)
										}
										return types.String{}
									}(),
								}
							}
							return nil
						}(),
						NextHopIP: func() types.String {
							if neighbors.NextHopIP != "" {
								return types.StringValue(neighbors.NextHopIP)
							}
							return types.String{}
						}(),
						ReceiveLimit: func() types.Int64 {
							if neighbors.ReceiveLimit != nil {
								return types.Int64Value(int64(*neighbors.ReceiveLimit))
							}
							return types.Int64{}
						}(),
						RemoteAsNumber: func() types.Int64 {
							if neighbors.RemoteAsNumber != nil {
								return types.Int64Value(int64(*neighbors.RemoteAsNumber))
							}
							return types.Int64{}
						}(),
						SourceInterface: func() types.String {
							if neighbors.SourceInterface != "" {
								return types.StringValue(neighbors.SourceInterface)
							}
							return types.String{}
						}(),
						TtlSecurity: func() *ResponseApplianceGetNetworkApplianceVpnBgpNeighborsTtlSecurityRs {
							if neighbors.TtlSecurity != nil {
								return &ResponseApplianceGetNetworkApplianceVpnBgpNeighborsTtlSecurityRs{
									Enabled: func() types.Bool {
										if neighbors.TtlSecurity.Enabled != nil {
											return types.BoolValue(*neighbors.TtlSecurity.Enabled)
										}
										return types.Bool{}
									}(),
								}
							}
							return nil
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceVpnBgpRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceVpnBgpRs)
}
