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
	"fmt"
	"strings"

	"log"

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWirelessSSIDsVpnResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessSSIDsVpnResource{}
)

func NewNetworksWirelessSSIDsVpnResource() resource.Resource {
	return &NetworksWirelessSSIDsVpnResource{}
}

type NetworksWirelessSSIDsVpnResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessSSIDsVpnResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessSSIDsVpnResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_vpn"
}

func (r *NetworksWirelessSSIDsVpnResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"concentrator": schema.SingleNestedAttribute{
				MarkdownDescription: `The VPN concentrator settings for this SSID.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"name": schema.StringAttribute{
						Computed: true,
					},
					"network_id": schema.StringAttribute{
						MarkdownDescription: `The NAT ID of the concentrator that should be set.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"vlan_id": schema.Int64Attribute{
						MarkdownDescription: `The VLAN that should be tagged for the concentrator.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"failover": schema.SingleNestedAttribute{
				MarkdownDescription: `Secondary VPN concentrator settings. This is only used when two VPN concentrators are configured on the SSID.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"heartbeat_interval": schema.Int64Attribute{
						MarkdownDescription: `Idle timer interval in seconds.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"idle_timeout": schema.Int64Attribute{
						MarkdownDescription: `Idle timer timeout in seconds.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"request_ip": schema.StringAttribute{
						MarkdownDescription: `IP addressed reserved on DHCP server where SSID will terminate.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: `number path parameter.`,
				Required:            true,
			},
			"split_tunnel": schema.SingleNestedAttribute{
				MarkdownDescription: `The VPN split tunnel settings for this SSID.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `If true, VPN split tunnel is enabled.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"rules": schema.SetNestedAttribute{
						MarkdownDescription: `List of VPN split tunnel rules.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"comment": schema.StringAttribute{
									MarkdownDescription: `Description for this split tunnel rule (optional).`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"dest_cidr": schema.StringAttribute{
									MarkdownDescription: `Destination for this split tunnel rule. IP address, fully-qualified domain names (FQDN) or 'any'.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"dest_port": schema.StringAttribute{
									MarkdownDescription: `Destination port for this split tunnel rule, (integer in the range 1-65535), or 'any'.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"policy": schema.StringAttribute{
									MarkdownDescription: `Traffic policy specified for this split tunnel rule, 'allow' or 'deny'.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"protocol": schema.StringAttribute{
									MarkdownDescription: `Protocol for this split tunnel rule.
                                              Allowed values: [Any,TCP,UDP]`,
									Computed: true,
									Optional: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
									Validators: []validator.String{
										stringvalidator.OneOf(
											"Any",
											"TCP",
											"UDP",
										),
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *NetworksWirelessSSIDsVpnResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessSSIDsVpnRs

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
	vvNumber := data.Number.ValueString()
	//Has Item and not has items

	if vvNetworkID != "" && vvNumber != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSIDVpn(vvNetworkID, vvNumber)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksWirelessSsidsVpn  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksWirelessSsidsVpn only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDVpn(vvNetworkID, vvNumber, dataRequest)
	//Update
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDVpn",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDVpn",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSIDVpn(vvNetworkID, vvNumber)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDVpn",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDVpn",
			err.Error(),
		)
		return
	}

	data = ResponseWirelessGetNetworkWirelessSSIDVpnItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *NetworksWirelessSSIDsVpnResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessSSIDsVpnRs

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
	vvNumber := data.Number.ValueString()
	responseGet, restyRespGet, err := r.client.Wireless.GetNetworkWirelessSSIDVpn(vvNetworkID, vvNumber)
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
				"Failure when executing GetNetworkWirelessSSIDVpn",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDVpn",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseWirelessGetNetworkWirelessSSIDVpnItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksWirelessSSIDsVpnResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("number"), idParts[1])...)
}

func (r *NetworksWirelessSSIDsVpnResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksWirelessSSIDsVpnRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvNumber := data.Number.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDVpn(vvNetworkID, vvNumber, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDVpn",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDVpn",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessSSIDsVpnResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksWirelessSSIDsVpn", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessSSIDsVpnRs struct {
	NetworkID    types.String                                             `tfsdk:"network_id"`
	Number       types.String                                             `tfsdk:"number"`
	Concentrator *ResponseWirelessGetNetworkWirelessSsidVpnConcentratorRs `tfsdk:"concentrator"`
	Failover     *ResponseWirelessGetNetworkWirelessSsidVpnFailoverRs     `tfsdk:"failover"`
	SplitTunnel  *ResponseWirelessGetNetworkWirelessSsidVpnSplitTunnelRs  `tfsdk:"split_tunnel"`
}

type ResponseWirelessGetNetworkWirelessSsidVpnConcentratorRs struct {
	Name      types.String `tfsdk:"name"`
	NetworkID types.String `tfsdk:"network_id"`
	VLANID    types.Int64  `tfsdk:"vlan_id"`
}

type ResponseWirelessGetNetworkWirelessSsidVpnFailoverRs struct {
	HeartbeatInterval types.Int64  `tfsdk:"heartbeat_interval"`
	IDleTimeout       types.Int64  `tfsdk:"idle_timeout"`
	RequestIP         types.String `tfsdk:"request_ip"`
}

type ResponseWirelessGetNetworkWirelessSsidVpnSplitTunnelRs struct {
	Enabled types.Bool                                                     `tfsdk:"enabled"`
	Rules   *[]ResponseWirelessGetNetworkWirelessSsidVpnSplitTunnelRulesRs `tfsdk:"rules"`
}

type ResponseWirelessGetNetworkWirelessSsidVpnSplitTunnelRulesRs struct {
	Comment  types.String `tfsdk:"comment"`
	DestCidr types.String `tfsdk:"dest_cidr"`
	DestPort types.String `tfsdk:"dest_port"`
	Policy   types.String `tfsdk:"policy"`
	Protocol types.String `tfsdk:"protocol"`
}

// FromBody
func (r *NetworksWirelessSSIDsVpnRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDVpn {
	var requestWirelessUpdateNetworkWirelessSSIDVpnConcentrator *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDVpnConcentrator

	if r.Concentrator != nil {
		networkID := r.Concentrator.NetworkID.ValueString()
		vlanID := func() *int64 {
			if !r.Concentrator.VLANID.IsUnknown() && !r.Concentrator.VLANID.IsNull() {
				return r.Concentrator.VLANID.ValueInt64Pointer()
			}
			return nil
		}()
		requestWirelessUpdateNetworkWirelessSSIDVpnConcentrator = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDVpnConcentrator{
			NetworkID: networkID,
			VLANID:    int64ToIntPointer(vlanID),
		}
		//[debug] Is Array: False
	}
	var requestWirelessUpdateNetworkWirelessSSIDVpnFailover *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDVpnFailover

	if r.Failover != nil {
		heartbeatInterval := func() *int64 {
			if !r.Failover.HeartbeatInterval.IsUnknown() && !r.Failover.HeartbeatInterval.IsNull() {
				return r.Failover.HeartbeatInterval.ValueInt64Pointer()
			}
			return nil
		}()
		idleTimeout := func() *int64 {
			if !r.Failover.IDleTimeout.IsUnknown() && !r.Failover.IDleTimeout.IsNull() {
				return r.Failover.IDleTimeout.ValueInt64Pointer()
			}
			return nil
		}()
		requestIP := r.Failover.RequestIP.ValueString()
		requestWirelessUpdateNetworkWirelessSSIDVpnFailover = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDVpnFailover{
			HeartbeatInterval: int64ToIntPointer(heartbeatInterval),
			IDleTimeout:       int64ToIntPointer(idleTimeout),
			RequestIP:         requestIP,
		}
		//[debug] Is Array: False
	}
	var requestWirelessUpdateNetworkWirelessSSIDVpnSplitTunnel *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDVpnSplitTunnel

	if r.SplitTunnel != nil {
		enabled := func() *bool {
			if !r.SplitTunnel.Enabled.IsUnknown() && !r.SplitTunnel.Enabled.IsNull() {
				return r.SplitTunnel.Enabled.ValueBoolPointer()
			}
			return nil
		}()

		log.Printf("[DEBUG] #TODO []RequestWirelessUpdateNetworkWirelessSsidVpnSplitTunnelRules")
		var requestWirelessUpdateNetworkWirelessSSIDVpnSplitTunnelRules []merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDVpnSplitTunnelRules

		if r.SplitTunnel.Rules != nil {
			for _, rItem1 := range *r.SplitTunnel.Rules {
				comment := rItem1.Comment.ValueString()
				destCidr := rItem1.DestCidr.ValueString()
				destPort := rItem1.DestPort.ValueString()
				policy := rItem1.Policy.ValueString()
				protocol := rItem1.Protocol.ValueString()
				requestWirelessUpdateNetworkWirelessSSIDVpnSplitTunnelRules = append(requestWirelessUpdateNetworkWirelessSSIDVpnSplitTunnelRules, merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDVpnSplitTunnelRules{
					Comment:  comment,
					DestCidr: destCidr,
					DestPort: destPort,
					Policy:   policy,
					Protocol: protocol,
				})
				//[debug] Is Array: True
			}
		}
		requestWirelessUpdateNetworkWirelessSSIDVpnSplitTunnel = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDVpnSplitTunnel{
			Enabled: enabled,
			Rules: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDVpnSplitTunnelRules {
				if len(requestWirelessUpdateNetworkWirelessSSIDVpnSplitTunnelRules) > 0 {
					return &requestWirelessUpdateNetworkWirelessSSIDVpnSplitTunnelRules
				}
				return nil
			}(),
		}
		//[debug] Is Array: False
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDVpn{
		Concentrator: requestWirelessUpdateNetworkWirelessSSIDVpnConcentrator,
		Failover:     requestWirelessUpdateNetworkWirelessSSIDVpnFailover,
		SplitTunnel:  requestWirelessUpdateNetworkWirelessSSIDVpnSplitTunnel,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetNetworkWirelessSSIDVpnItemToBodyRs(state NetworksWirelessSSIDsVpnRs, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDVpn, is_read bool) NetworksWirelessSSIDsVpnRs {
	itemState := NetworksWirelessSSIDsVpnRs{
		Concentrator: func() *ResponseWirelessGetNetworkWirelessSsidVpnConcentratorRs {
			if response.Concentrator != nil {
				return &ResponseWirelessGetNetworkWirelessSsidVpnConcentratorRs{
					Name:      types.StringValue(response.Concentrator.Name),
					NetworkID: types.StringValue(response.Concentrator.NetworkID),
					VLANID: func() types.Int64 {
						if response.Concentrator.VLANID != nil {
							return types.Int64Value(int64(*response.Concentrator.VLANID))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
		Failover: func() *ResponseWirelessGetNetworkWirelessSsidVpnFailoverRs {
			if response.Failover != nil {
				return &ResponseWirelessGetNetworkWirelessSsidVpnFailoverRs{
					HeartbeatInterval: func() types.Int64 {
						if response.Failover.HeartbeatInterval != nil {
							return types.Int64Value(int64(*response.Failover.HeartbeatInterval))
						}
						return types.Int64{}
					}(),
					IDleTimeout: func() types.Int64 {
						if response.Failover.IDleTimeout != nil {
							return types.Int64Value(int64(*response.Failover.IDleTimeout))
						}
						return types.Int64{}
					}(),
					RequestIP: types.StringValue(response.Failover.RequestIP),
				}
			}
			return nil
		}(),
		SplitTunnel: func() *ResponseWirelessGetNetworkWirelessSsidVpnSplitTunnelRs {
			if response.SplitTunnel != nil {
				return &ResponseWirelessGetNetworkWirelessSsidVpnSplitTunnelRs{
					Enabled: func() types.Bool {
						if response.SplitTunnel.Enabled != nil {
							return types.BoolValue(*response.SplitTunnel.Enabled)
						}
						return types.Bool{}
					}(),
					Rules: func() *[]ResponseWirelessGetNetworkWirelessSsidVpnSplitTunnelRulesRs {
						if response.SplitTunnel.Rules != nil {
							result := make([]ResponseWirelessGetNetworkWirelessSsidVpnSplitTunnelRulesRs, len(*response.SplitTunnel.Rules))
							for i, rules := range *response.SplitTunnel.Rules {
								result[i] = ResponseWirelessGetNetworkWirelessSsidVpnSplitTunnelRulesRs{
									Comment:  types.StringValue(rules.Comment),
									DestCidr: types.StringValue(rules.DestCidr),
									DestPort: types.StringValue(rules.DestPort),
									Policy:   types.StringValue(rules.Policy),
									Protocol: types.StringValue(rules.Protocol),
								}
							}
							return &result
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessSSIDsVpnRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessSSIDsVpnRs)
}
