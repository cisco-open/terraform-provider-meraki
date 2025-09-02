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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSwitchDhcpServerPolicyResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchDhcpServerPolicyResource{}
)

func NewNetworksSwitchDhcpServerPolicyResource() resource.Resource {
	return &NetworksSwitchDhcpServerPolicyResource{}
}

type NetworksSwitchDhcpServerPolicyResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchDhcpServerPolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchDhcpServerPolicyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_dhcp_server_policy"
}

func (r *NetworksSwitchDhcpServerPolicyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"alerts": schema.SingleNestedAttribute{
				MarkdownDescription: `Email alert settings for DHCP servers`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"email": schema.SingleNestedAttribute{
						MarkdownDescription: `Alert settings for DHCP servers`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `When enabled, send an email if a new DHCP server is seen. Default value is false.`,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
				},
			},
			"allowed_servers": schema.ListAttribute{
				MarkdownDescription: `List the MAC addresses of DHCP servers to permit on the network when defaultPolicy is set
      to block.An empty array will clear the entries.`,
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
				Default:     listdefault.StaticValue(types.ListNull(types.StringType)),
			},
			"arp_inspection": schema.SingleNestedAttribute{
				MarkdownDescription: `Dynamic ARP Inspection settings`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Enable or disable Dynamic ARP Inspection on the network. Default value is false.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"unsupported_models": schema.ListAttribute{
						MarkdownDescription: `List of switch models that does not support dynamic ARP inspection`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
			"blocked_servers": schema.ListAttribute{
				MarkdownDescription: `List the MAC addresses of DHCP servers to block on the network when defaultPolicy is set
      to allow.An empty array will clear the entries.`,
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
				Default:     listdefault.StaticValue(types.ListNull(types.StringType)),
			},
			"default_policy": schema.StringAttribute{
				MarkdownDescription: `'allow' or 'block' new DHCP servers. Default value is 'allow'.
                                  Allowed values: [allow,block]`,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"allow",
						"block",
					),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
		},
	}
}

func (r *NetworksSwitchDhcpServerPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchDhcpServerPolicyRs

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
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchDhcpServerPolicy(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchDhcpServerPolicy",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchDhcpServerPolicy",
			err.Error(),
		)
		return
	}

	// Assign data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *NetworksSwitchDhcpServerPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchDhcpServerPolicyRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	responseGet, restyRespGet, err := r.client.Switch.GetNetworkSwitchDhcpServerPolicy(vvNetworkID)
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
				"Failure when executing GetNetworkSwitchDhcpServerPolicy",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchDhcpServerPolicy",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseSwitchGetNetworkSwitchDhcpServerPolicyItemToBodyRs(data, responseGet, true)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *NetworksSwitchDhcpServerPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksSwitchDhcpServerPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworksSwitchDhcpServerPolicyRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvNetworkID := plan.NetworkID.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchDhcpServerPolicy(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchDhcpServerPolicy",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchDhcpServerPolicy",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworksSwitchDhcpServerPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksSwitchDhcpServerPolicy", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSwitchDhcpServerPolicyRs struct {
	NetworkID      types.String                                                   `tfsdk:"network_id"`
	Alerts         *ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlertsRs        `tfsdk:"alerts"`
	AllowedServers types.List                                                     `tfsdk:"allowed_servers"`
	ArpInspection  *ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionRs `tfsdk:"arp_inspection"`
	BlockedServers types.List                                                     `tfsdk:"blocked_servers"`
	DefaultPolicy  types.String                                                   `tfsdk:"default_policy"`
}

type ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlertsRs struct {
	Email *ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlertsEmailRs `tfsdk:"email"`
}

type ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlertsEmailRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionRs struct {
	Enabled           types.Bool `tfsdk:"enabled"`
	UnsupportedModels types.List `tfsdk:"unsupported_models"`
}

// FromBody
func (r *NetworksSwitchDhcpServerPolicyRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateNetworkSwitchDhcpServerPolicy {
	emptyString := ""
	var requestSwitchUpdateNetworkSwitchDhcpServerPolicyAlerts *merakigosdk.RequestSwitchUpdateNetworkSwitchDhcpServerPolicyAlerts

	if r.Alerts != nil {
		var requestSwitchUpdateNetworkSwitchDhcpServerPolicyAlertsEmail *merakigosdk.RequestSwitchUpdateNetworkSwitchDhcpServerPolicyAlertsEmail

		if r.Alerts.Email != nil {
			enabled := func() *bool {
				if !r.Alerts.Email.Enabled.IsUnknown() && !r.Alerts.Email.Enabled.IsNull() {
					return r.Alerts.Email.Enabled.ValueBoolPointer()
				}
				return nil
			}()
			requestSwitchUpdateNetworkSwitchDhcpServerPolicyAlertsEmail = &merakigosdk.RequestSwitchUpdateNetworkSwitchDhcpServerPolicyAlertsEmail{
				Enabled: enabled,
			}
			//[debug] Is Array: False
		}
		requestSwitchUpdateNetworkSwitchDhcpServerPolicyAlerts = &merakigosdk.RequestSwitchUpdateNetworkSwitchDhcpServerPolicyAlerts{
			Email: requestSwitchUpdateNetworkSwitchDhcpServerPolicyAlertsEmail,
		}
		//[debug] Is Array: False
	}
	var allowedServers []string = nil
	r.AllowedServers.ElementsAs(ctx, &allowedServers, false)
	var requestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspection *merakigosdk.RequestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspection

	if r.ArpInspection != nil {
		enabled := func() *bool {
			if !r.ArpInspection.Enabled.IsUnknown() && !r.ArpInspection.Enabled.IsNull() {
				return r.ArpInspection.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		requestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspection = &merakigosdk.RequestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspection{
			Enabled: enabled,
		}
		//[debug] Is Array: False
	}
	var blockedServers []string = nil
	r.BlockedServers.ElementsAs(ctx, &blockedServers, false)
	defaultPolicy := new(string)
	if !r.DefaultPolicy.IsUnknown() && !r.DefaultPolicy.IsNull() {
		*defaultPolicy = r.DefaultPolicy.ValueString()
	} else {
		defaultPolicy = &emptyString
	}
	out := merakigosdk.RequestSwitchUpdateNetworkSwitchDhcpServerPolicy{
		Alerts:         requestSwitchUpdateNetworkSwitchDhcpServerPolicyAlerts,
		AllowedServers: allowedServers,
		ArpInspection:  requestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspection,
		BlockedServers: blockedServers,
		DefaultPolicy:  *defaultPolicy,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchDhcpServerPolicyItemToBodyRs(state NetworksSwitchDhcpServerPolicyRs, response *merakigosdk.ResponseSwitchGetNetworkSwitchDhcpServerPolicy, is_read bool) NetworksSwitchDhcpServerPolicyRs {
	itemState := NetworksSwitchDhcpServerPolicyRs{
		Alerts: func() *ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlertsRs {
			if response.Alerts != nil {
				return &ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlertsRs{
					Email: func() *ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlertsEmailRs {
						if response.Alerts.Email != nil {
							return &ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlertsEmailRs{
								Enabled: func() types.Bool {
									if response.Alerts.Email.Enabled != nil {
										return types.BoolValue(*response.Alerts.Email.Enabled)
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
		AllowedServers: StringSliceToList(response.AllowedServers),
		ArpInspection: func() *ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionRs {
			if response.ArpInspection != nil {
				return &ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionRs{
					Enabled: func() types.Bool {
						if response.ArpInspection.Enabled != nil {
							return types.BoolValue(*response.ArpInspection.Enabled)
						}
						return types.Bool{}
					}(),
					UnsupportedModels: StringSliceToList(response.ArpInspection.UnsupportedModels),
				}
			}
			return nil
		}(),
		BlockedServers: StringSliceToList(response.BlockedServers),
		DefaultPolicy: func() types.String {
			if response.DefaultPolicy != "" {
				return types.StringValue(response.DefaultPolicy)
			}
			return types.String{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSwitchDhcpServerPolicyRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSwitchDhcpServerPolicyRs)
}
