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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSwitchAccessPoliciesResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchAccessPoliciesResource{}
)

func NewNetworksSwitchAccessPoliciesResource() resource.Resource {
	return &NetworksSwitchAccessPoliciesResource{}
}

type NetworksSwitchAccessPoliciesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchAccessPoliciesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchAccessPoliciesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_access_policies"
}

func (r *NetworksSwitchAccessPoliciesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"access_policy_number": schema.StringAttribute{
				MarkdownDescription: `accessPolicyNumber path parameter. Access policy number`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"access_policy_type": schema.StringAttribute{
				MarkdownDescription: `Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.
                                  Allowed values: [802.1x,Hybrid authentication,MAC authentication bypass]`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"802.1x",
						"Hybrid authentication",
						"MAC authentication bypass",
					),
				},
			},
			"counts": schema.SingleNestedAttribute{
				MarkdownDescription: `Counts associated with the access policy`,
				Computed:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"ports": schema.SingleNestedAttribute{
						MarkdownDescription: `Counts associated with ports`,
						Computed:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"with_this_policy": schema.Int64Attribute{
								MarkdownDescription: `Number of ports in the network with this policy. For template networks, this is the number of template ports (not child ports) with this policy.`,
								Computed:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
							},
						},
					},
				},
			},
			"dot1x": schema.SingleNestedAttribute{
				MarkdownDescription: `802.1x Settings`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"control_direction": schema.StringAttribute{
						MarkdownDescription: `Supports either 'both' or 'inbound'. Set to 'inbound' to allow unauthorized egress on the switchport. Set to 'both' to control both traffic directions with authorization. Defaults to 'both'
                                        Allowed values: [both,inbound]`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"both",
								"inbound",
							),
						},
					},
				},
			},
			"guest_port_bouncing": schema.BoolAttribute{
				MarkdownDescription: `If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"guest_vlan_id": schema.Int64Attribute{
				MarkdownDescription: `ID for the guest VLAN allow unauthorized devices access to limited network resources`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"host_mode": schema.StringAttribute{
				MarkdownDescription: `Choose the Host Mode for the access policy.
                                  Allowed values: [Multi-Auth,Multi-Domain,Multi-Host,Single-Host]`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Multi-Auth",
						"Multi-Domain",
						"Multi-Host",
						"Single-Host",
					),
				},
			},
			"increase_access_speed": schema.BoolAttribute{
				MarkdownDescription: `Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is 'Hybrid Authentication.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Name of the access policy`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"radius": schema.SingleNestedAttribute{
				MarkdownDescription: `Object for RADIUS Settings`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"cache": schema.SingleNestedAttribute{
						MarkdownDescription: `Object for RADIUS Cache Settings`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Enable to cache authorization and authentication responses on the RADIUS server`,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"timeout": schema.Int64Attribute{
								MarkdownDescription: `If RADIUS caching is enabled, this value dictates how long the cache will remain in the RADIUS server, in hours, to allow network access without authentication`,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"critical_auth": schema.SingleNestedAttribute{
						MarkdownDescription: `Critical auth settings for when authentication is rejected by the RADIUS server`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"data_vlan_id": schema.Int64Attribute{
								MarkdownDescription: `VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
							},
							"suspend_port_bounce": schema.BoolAttribute{
								MarkdownDescription: `Enable to suspend port bounce when RADIUS servers are unreachable`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"voice_vlan_id": schema.Int64Attribute{
								MarkdownDescription: `VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth`,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"failed_auth_vlan_id": schema.Int64Attribute{
						MarkdownDescription: `VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"re_authentication_interval": schema.Int64Attribute{
						MarkdownDescription: `Re-authentication period in seconds. Will be null if hostMode is Multi-Auth`,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"radius_accounting_enabled": schema.BoolAttribute{
				MarkdownDescription: `Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"radius_accounting_servers": schema.SetNestedAttribute{
				MarkdownDescription: `List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access`,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"host": schema.StringAttribute{
							MarkdownDescription: `Public IP address of the RADIUS accounting server`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"organization_radius_server_id": schema.StringAttribute{
							MarkdownDescription: `Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: `UDP port that the RADIUS Accounting server listens on for access requests`,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: `RADIUS client shared secret`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"server_id": schema.StringAttribute{
							MarkdownDescription: `Unique ID of the RADIUS accounting server`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"radius_accounting_servers_response": schema.SetNestedAttribute{
				MarkdownDescription: `List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access`,
				Computed:            true,

				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"host": schema.StringAttribute{
							MarkdownDescription: `Public IP address of the RADIUS accounting server`,
							Computed:            true,

							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: `UDP port that the RADIUS Accounting server listens on for access requests`,
							Computed:            true,

							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"organization_radius_server_id": schema.StringAttribute{
							MarkdownDescription: `Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server`,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: `RADIUS client shared secret`,
							Computed:            true,
							Default:             stringdefault.StaticString(""),
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"server_id": schema.StringAttribute{
							MarkdownDescription: `Unique ID of the RADIUS accounting server`,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"radius_coa_support_enabled": schema.BoolAttribute{
				MarkdownDescription: `Change of authentication for RADIUS re-authentication and disconnection`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"radius_group_attribute": schema.StringAttribute{
				MarkdownDescription: `Acceptable values are *""* for None, or *"11"* for Group Policies ACL`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"radius_servers": schema.SetNestedAttribute{
				MarkdownDescription: `List of RADIUS servers to require connecting devices to authenticate against before granting network access`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"host": schema.StringAttribute{
							MarkdownDescription: `Public IP address of the RADIUS server`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"organization_radius_server_id": schema.StringAttribute{
							MarkdownDescription: `Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: `UDP port that the RADIUS server listens on for access requests`,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: `RADIUS client shared secret`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"server_id": schema.StringAttribute{
							MarkdownDescription: `Unique ID of the RADIUS server`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"radius_servers_response": schema.SetNestedAttribute{
				MarkdownDescription: `List of RADIUS servers to require connecting devices to authenticate against before granting network access`,
				Computed:            true,

				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"host": schema.StringAttribute{
							MarkdownDescription: `Public IP address of the RADIUS server`,
							Computed:            true,

							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: `UDP port that the RADIUS server listens on for access requests`,
							Computed:            true,

							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: `RADIUS client shared secret`,
							Computed:            true,
							Default:             stringdefault.StaticString(""),
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"organization_radius_server_id": schema.StringAttribute{
							MarkdownDescription: `Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"server_id": schema.StringAttribute{
							MarkdownDescription: `Unique ID of the RADIUS server`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"radius_testing_enabled": schema.BoolAttribute{
				MarkdownDescription: `If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"url_redirect_walled_garden_enabled": schema.BoolAttribute{
				MarkdownDescription: `Enable to restrict access for clients to a response_objectific set of IP addresses or hostnames prior to authentication`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"url_redirect_walled_garden_ranges": schema.SetAttribute{
				MarkdownDescription: `IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication`,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"voice_vlan_clients": schema.BoolAttribute{
				MarkdownDescription: `CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

//path params to set ['accessPolicyNumber']

func (r *NetworksSwitchAccessPoliciesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchAccessPoliciesRs

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
	vvName := data.Name.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Switch.GetNetworkSwitchAccessPolicies(vvNetworkID)
	//Have Create
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkSwitchAccessPolicies",
					err.Error(),
				)
				return
			}
		}
	}
	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvAccessPolicyNumber, ok := result2["AccessPolicyNumber"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter AccessPolicyNumber",
					"Fail Parsing AccessPolicyNumber",
				)
				return
			}
			r.client.Switch.UpdateNetworkSwitchAccessPolicy(vvNetworkID, vvAccessPolicyNumber, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Switch.GetNetworkSwitchAccessPolicy(vvNetworkID, vvAccessPolicyNumber)
			if responseVerifyItem2 != nil {
				data = ResponseSwitchGetNetworkSwitchAccessPolicyItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Switch.CreateNetworkSwitchAccessPolicy(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkSwitchAccessPolicy",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkSwitchAccessPolicy",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Switch.GetNetworkSwitchAccessPolicies(vvNetworkID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchAccessPolicies",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchAccessPolicies",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvAccessPolicyNumber, ok := result2["AccessPolicyNumber"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter AccessPolicyNumber",
				"Error",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Switch.GetNetworkSwitchAccessPolicy(vvNetworkID, vvAccessPolicyNumber)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseSwitchGetNetworkSwitchAccessPolicyItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkSwitchAccessPolicy",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchAccessPolicy",
				err.Error(),
			)
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"Error in result.",
			"Error in result.",
		)
		return
	}
}

func (r *NetworksSwitchAccessPoliciesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchAccessPoliciesRs

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
	vvAccessPolicyNumber := data.AccessPolicyNumber.ValueString()
	responseGet, restyRespGet, err := r.client.Switch.GetNetworkSwitchAccessPolicy(vvNetworkID, vvAccessPolicyNumber)
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
				"Failure when executing GetNetworkSwitchAccessPolicy",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchAccessPolicy",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseSwitchGetNetworkSwitchAccessPolicyItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksSwitchAccessPoliciesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("access_policy_number"), idParts[1])...)
}

func (r *NetworksSwitchAccessPoliciesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSwitchAccessPoliciesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvAccessPolicyNumber := data.AccessPolicyNumber.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchAccessPolicy(vvNetworkID, vvAccessPolicyNumber, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchAccessPolicy",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchAccessPolicy",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchAccessPoliciesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksSwitchAccessPoliciesRs
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &state, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}

	vvNetworkID := state.NetworkID.ValueString()
	vvAccessPolicyNumber := state.AccessPolicyNumber.ValueString()
	_, err := r.client.Switch.DeleteNetworkSwitchAccessPolicy(vvNetworkID, vvAccessPolicyNumber)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkSwitchAccessPolicy", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksSwitchAccessPoliciesRs struct {
	NetworkID                       types.String                                                           `tfsdk:"network_id"`
	AccessPolicyNumber              types.String                                                           `tfsdk:"access_policy_number"`
	AccessPolicyType                types.String                                                           `tfsdk:"access_policy_type"`
	Counts                          *ResponseSwitchGetNetworkSwitchAccessPolicyCountsRs                    `tfsdk:"counts"`
	Dot1X                           *ResponseSwitchGetNetworkSwitchAccessPolicyDot1XRs                     `tfsdk:"dot1x"`
	GuestPortBouncing               types.Bool                                                             `tfsdk:"guest_port_bouncing"`
	GuestVLANID                     types.Int64                                                            `tfsdk:"guest_vlan_id"`
	HostMode                        types.String                                                           `tfsdk:"host_mode"`
	IncreaseAccessSpeed             types.Bool                                                             `tfsdk:"increase_access_speed"`
	Name                            types.String                                                           `tfsdk:"name"`
	Radius                          *ResponseSwitchGetNetworkSwitchAccessPolicyRadiusRs                    `tfsdk:"radius"`
	RadiusAccountingEnabled         types.Bool                                                             `tfsdk:"radius_accounting_enabled"`
	RadiusAccountingServers         *[]ResponseSwitchGetNetworkSwitchAccessPolicyRadiusAccountingServersRs `tfsdk:"radius_accounting_servers"`
	RadiusAccountingServersResponse *[]ResponseSwitchGetNetworkSwitchAccessPolicyRadiusAccountingServersRs `tfsdk:"radius_accounting_servers_response"`
	RadiusCoaSupportEnabled         types.Bool                                                             `tfsdk:"radius_coa_support_enabled"`
	RadiusGroupAttribute            types.String                                                           `tfsdk:"radius_group_attribute"`
	RadiusServers                   *[]ResponseSwitchGetNetworkSwitchAccessPolicyRadiusServersRs           `tfsdk:"radius_servers"`
	RadiusServersResponse           *[]ResponseSwitchGetNetworkSwitchAccessPolicyRadiusServersRs           `tfsdk:"radius_servers_response"`
	RadiusTestingEnabled            types.Bool                                                             `tfsdk:"radius_testing_enabled"`
	URLRedirectWalledGardenEnabled  types.Bool                                                             `tfsdk:"url_redirect_walled_garden_enabled"`
	URLRedirectWalledGardenRanges   types.Set                                                              `tfsdk:"url_redirect_walled_garden_ranges"`
	VoiceVLANClients                types.Bool                                                             `tfsdk:"voice_vlan_clients"`
}

type ResponseSwitchGetNetworkSwitchAccessPolicyCountsRs struct {
	Ports *ResponseSwitchGetNetworkSwitchAccessPolicyCountsPortsRs `tfsdk:"ports"`
}

type ResponseSwitchGetNetworkSwitchAccessPolicyCountsPortsRs struct {
	WithThisPolicy types.Int64 `tfsdk:"with_this_policy"`
}

type ResponseSwitchGetNetworkSwitchAccessPolicyDot1XRs struct {
	ControlDirection types.String `tfsdk:"control_direction"`
}

type ResponseSwitchGetNetworkSwitchAccessPolicyRadiusRs struct {
	Cache                    *ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCacheRs        `tfsdk:"cache"`
	CriticalAuth             *ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCriticalAuthRs `tfsdk:"critical_auth"`
	FailedAuthVLANID         types.Int64                                                     `tfsdk:"failed_auth_vlan_id"`
	ReAuthenticationInterval types.Int64                                                     `tfsdk:"re_authentication_interval"`
}

type ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCacheRs struct {
	Enabled types.Bool  `tfsdk:"enabled"`
	Timeout types.Int64 `tfsdk:"timeout"`
}

type ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCriticalAuthRs struct {
	DataVLANID        types.Int64 `tfsdk:"data_vlan_id"`
	SuspendPortBounce types.Bool  `tfsdk:"suspend_port_bounce"`
	VoiceVLANID       types.Int64 `tfsdk:"voice_vlan_id"`
}

type ResponseSwitchGetNetworkSwitchAccessPolicyRadiusAccountingServersRs struct {
	Host                       types.String `tfsdk:"host"`
	OrganizationRadiusServerID types.String `tfsdk:"organization_radius_server_id"`
	Port                       types.Int64  `tfsdk:"port"`
	ServerID                   types.String `tfsdk:"server_id"`
	Secret                     types.String `tfsdk:"secret"`
}

type ResponseSwitchGetNetworkSwitchAccessPolicyRadiusServersRs struct {
	Host                       types.String `tfsdk:"host"`
	OrganizationRadiusServerID types.String `tfsdk:"organization_radius_server_id"`
	Port                       types.Int64  `tfsdk:"port"`
	ServerID                   types.String `tfsdk:"server_id"`
	Secret                     types.String `tfsdk:"secret"`
}

// FromBody
func (r *NetworksSwitchAccessPoliciesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSwitchCreateNetworkSwitchAccessPolicy {
	emptyString := ""
	accessPolicyType := new(string)
	if !r.AccessPolicyType.IsUnknown() && !r.AccessPolicyType.IsNull() {
		*accessPolicyType = r.AccessPolicyType.ValueString()
	} else {
		accessPolicyType = &emptyString
	}
	var requestSwitchCreateNetworkSwitchAccessPolicyDot1X *merakigosdk.RequestSwitchCreateNetworkSwitchAccessPolicyDot1X
	if r.Dot1X != nil {
		controlDirection := r.Dot1X.ControlDirection.ValueString()
		requestSwitchCreateNetworkSwitchAccessPolicyDot1X = &merakigosdk.RequestSwitchCreateNetworkSwitchAccessPolicyDot1X{
			ControlDirection: controlDirection,
		}
	}
	guestPortBouncing := new(bool)
	if !r.GuestPortBouncing.IsUnknown() && !r.GuestPortBouncing.IsNull() {
		*guestPortBouncing = r.GuestPortBouncing.ValueBool()
	} else {
		guestPortBouncing = nil
	}
	guestVLANID := new(int64)
	if !r.GuestVLANID.IsUnknown() && !r.GuestVLANID.IsNull() {
		*guestVLANID = r.GuestVLANID.ValueInt64()
	} else {
		guestVLANID = nil
	}
	hostMode := new(string)
	if !r.HostMode.IsUnknown() && !r.HostMode.IsNull() {
		*hostMode = r.HostMode.ValueString()
	} else {
		hostMode = &emptyString
	}
	increaseAccessSpeed := new(bool)
	if !r.IncreaseAccessSpeed.IsUnknown() && !r.IncreaseAccessSpeed.IsNull() {
		*increaseAccessSpeed = r.IncreaseAccessSpeed.ValueBool()
	} else {
		increaseAccessSpeed = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestSwitchCreateNetworkSwitchAccessPolicyRadius *merakigosdk.RequestSwitchCreateNetworkSwitchAccessPolicyRadius
	if r.Radius != nil {
		var requestSwitchCreateNetworkSwitchAccessPolicyRadiusCache *merakigosdk.RequestSwitchCreateNetworkSwitchAccessPolicyRadiusCache
		if r.Radius.Cache != nil {
			enabled := func() *bool {
				if !r.Radius.Cache.Enabled.IsUnknown() && !r.Radius.Cache.Enabled.IsNull() {
					return r.Radius.Cache.Enabled.ValueBoolPointer()
				}
				return nil
			}()
			timeout := func() *int64 {
				if !r.Radius.Cache.Timeout.IsUnknown() && !r.Radius.Cache.Timeout.IsNull() {
					return r.Radius.Cache.Timeout.ValueInt64Pointer()
				}
				return nil
			}()
			requestSwitchCreateNetworkSwitchAccessPolicyRadiusCache = &merakigosdk.RequestSwitchCreateNetworkSwitchAccessPolicyRadiusCache{
				Enabled: enabled,
				Timeout: int64ToIntPointer(timeout),
			}
		}
		var requestSwitchCreateNetworkSwitchAccessPolicyRadiusCriticalAuth *merakigosdk.RequestSwitchCreateNetworkSwitchAccessPolicyRadiusCriticalAuth
		if r.Radius.CriticalAuth != nil {
			dataVLANID := func() *int64 {
				if !r.Radius.CriticalAuth.DataVLANID.IsUnknown() && !r.Radius.CriticalAuth.DataVLANID.IsNull() {
					return r.Radius.CriticalAuth.DataVLANID.ValueInt64Pointer()
				}
				return nil
			}()
			suspendPortBounce := func() *bool {
				if !r.Radius.CriticalAuth.SuspendPortBounce.IsUnknown() && !r.Radius.CriticalAuth.SuspendPortBounce.IsNull() {
					return r.Radius.CriticalAuth.SuspendPortBounce.ValueBoolPointer()
				}
				return nil
			}()
			voiceVLANID := func() *int64 {
				if !r.Radius.CriticalAuth.VoiceVLANID.IsUnknown() && !r.Radius.CriticalAuth.VoiceVLANID.IsNull() {
					return r.Radius.CriticalAuth.VoiceVLANID.ValueInt64Pointer()
				}
				return nil
			}()
			requestSwitchCreateNetworkSwitchAccessPolicyRadiusCriticalAuth = &merakigosdk.RequestSwitchCreateNetworkSwitchAccessPolicyRadiusCriticalAuth{
				DataVLANID:        int64ToIntPointer(dataVLANID),
				SuspendPortBounce: suspendPortBounce,
				VoiceVLANID:       int64ToIntPointer(voiceVLANID),
			}
		}
		failedAuthVLANID := func() *int64 {
			if !r.Radius.FailedAuthVLANID.IsUnknown() && !r.Radius.FailedAuthVLANID.IsNull() {
				return r.Radius.FailedAuthVLANID.ValueInt64Pointer()
			}
			return nil
		}()
		reAuthenticationInterval := func() *int64 {
			if !r.Radius.ReAuthenticationInterval.IsUnknown() && !r.Radius.ReAuthenticationInterval.IsNull() {
				return r.Radius.ReAuthenticationInterval.ValueInt64Pointer()
			}
			return nil
		}()
		requestSwitchCreateNetworkSwitchAccessPolicyRadius = &merakigosdk.RequestSwitchCreateNetworkSwitchAccessPolicyRadius{
			Cache:                    requestSwitchCreateNetworkSwitchAccessPolicyRadiusCache,
			CriticalAuth:             requestSwitchCreateNetworkSwitchAccessPolicyRadiusCriticalAuth,
			FailedAuthVLANID:         int64ToIntPointer(failedAuthVLANID),
			ReAuthenticationInterval: int64ToIntPointer(reAuthenticationInterval),
		}
	}
	radiusAccountingEnabled := new(bool)
	if !r.RadiusAccountingEnabled.IsUnknown() && !r.RadiusAccountingEnabled.IsNull() {
		*radiusAccountingEnabled = r.RadiusAccountingEnabled.ValueBool()
	} else {
		radiusAccountingEnabled = nil
	}
	var requestSwitchCreateNetworkSwitchAccessPolicyRadiusAccountingServers []merakigosdk.RequestSwitchCreateNetworkSwitchAccessPolicyRadiusAccountingServers
	if r.RadiusAccountingServers != nil {
		for _, rItem1 := range *r.RadiusAccountingServers {
			host := rItem1.Host.ValueString()
			organizationRadiusServerID := rItem1.OrganizationRadiusServerID.ValueString()
			port := func() *int64 {
				if !rItem1.Port.IsUnknown() && !rItem1.Port.IsNull() {
					return rItem1.Port.ValueInt64Pointer()
				}
				return nil
			}()
			secret := rItem1.Secret.ValueString()
			requestSwitchCreateNetworkSwitchAccessPolicyRadiusAccountingServers = append(requestSwitchCreateNetworkSwitchAccessPolicyRadiusAccountingServers, merakigosdk.RequestSwitchCreateNetworkSwitchAccessPolicyRadiusAccountingServers{
				Host:                       host,
				OrganizationRadiusServerID: organizationRadiusServerID,
				Port:                       int64ToIntPointer(port),
				Secret:                     secret,
			})
		}
	}
	radiusCoaSupportEnabled := new(bool)
	if !r.RadiusCoaSupportEnabled.IsUnknown() && !r.RadiusCoaSupportEnabled.IsNull() {
		*radiusCoaSupportEnabled = r.RadiusCoaSupportEnabled.ValueBool()
	} else {
		radiusCoaSupportEnabled = nil
	}
	radiusGroupAttribute := new(string)
	if !r.RadiusGroupAttribute.IsUnknown() && !r.RadiusGroupAttribute.IsNull() {
		*radiusGroupAttribute = r.RadiusGroupAttribute.ValueString()
	} else {
		radiusGroupAttribute = &emptyString
	}
	var requestSwitchCreateNetworkSwitchAccessPolicyRadiusServers []merakigosdk.RequestSwitchCreateNetworkSwitchAccessPolicyRadiusServers
	if r.RadiusServers != nil {
		for _, rItem1 := range *r.RadiusServers {
			host := rItem1.Host.ValueString()
			organizationRadiusServerID := rItem1.OrganizationRadiusServerID.ValueString()
			port := func() *int64 {
				if !rItem1.Port.IsUnknown() && !rItem1.Port.IsNull() {
					return rItem1.Port.ValueInt64Pointer()
				}
				return nil
			}()
			secret := rItem1.Secret.ValueString()
			requestSwitchCreateNetworkSwitchAccessPolicyRadiusServers = append(requestSwitchCreateNetworkSwitchAccessPolicyRadiusServers, merakigosdk.RequestSwitchCreateNetworkSwitchAccessPolicyRadiusServers{
				Host:                       host,
				OrganizationRadiusServerID: organizationRadiusServerID,
				Port:                       int64ToIntPointer(port),
				Secret:                     secret,
			})
		}
	}
	radiusTestingEnabled := new(bool)
	if !r.RadiusTestingEnabled.IsUnknown() && !r.RadiusTestingEnabled.IsNull() {
		*radiusTestingEnabled = r.RadiusTestingEnabled.ValueBool()
	} else {
		radiusTestingEnabled = nil
	}
	uRLRedirectWalledGardenEnabled := new(bool)
	if !r.URLRedirectWalledGardenEnabled.IsUnknown() && !r.URLRedirectWalledGardenEnabled.IsNull() {
		*uRLRedirectWalledGardenEnabled = r.URLRedirectWalledGardenEnabled.ValueBool()
	} else {
		uRLRedirectWalledGardenEnabled = nil
	}
	var uRLRedirectWalledGardenRanges []string = nil
	r.URLRedirectWalledGardenRanges.ElementsAs(ctx, &uRLRedirectWalledGardenRanges, false)
	voiceVLANClients := new(bool)
	if !r.VoiceVLANClients.IsUnknown() && !r.VoiceVLANClients.IsNull() {
		*voiceVLANClients = r.VoiceVLANClients.ValueBool()
	} else {
		voiceVLANClients = nil
	}
	out := merakigosdk.RequestSwitchCreateNetworkSwitchAccessPolicy{
		AccessPolicyType:        *accessPolicyType,
		Dot1X:                   requestSwitchCreateNetworkSwitchAccessPolicyDot1X,
		GuestPortBouncing:       guestPortBouncing,
		GuestVLANID:             int64ToIntPointer(guestVLANID),
		HostMode:                *hostMode,
		IncreaseAccessSpeed:     increaseAccessSpeed,
		Name:                    *name,
		Radius:                  requestSwitchCreateNetworkSwitchAccessPolicyRadius,
		RadiusAccountingEnabled: radiusAccountingEnabled,
		RadiusAccountingServers: func() *[]merakigosdk.RequestSwitchCreateNetworkSwitchAccessPolicyRadiusAccountingServers {
			if len(requestSwitchCreateNetworkSwitchAccessPolicyRadiusAccountingServers) > 0 {
				return &requestSwitchCreateNetworkSwitchAccessPolicyRadiusAccountingServers
			}
			return nil
		}(),
		RadiusCoaSupportEnabled: radiusCoaSupportEnabled,
		RadiusGroupAttribute:    *radiusGroupAttribute,
		RadiusServers: func() *[]merakigosdk.RequestSwitchCreateNetworkSwitchAccessPolicyRadiusServers {
			if len(requestSwitchCreateNetworkSwitchAccessPolicyRadiusServers) > 0 {
				return &requestSwitchCreateNetworkSwitchAccessPolicyRadiusServers
			}
			return nil
		}(),
		RadiusTestingEnabled:           radiusTestingEnabled,
		URLRedirectWalledGardenEnabled: uRLRedirectWalledGardenEnabled,
		URLRedirectWalledGardenRanges:  uRLRedirectWalledGardenRanges,
		VoiceVLANClients:               voiceVLANClients,
	}
	return &out
}
func (r *NetworksSwitchAccessPoliciesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateNetworkSwitchAccessPolicy {
	emptyString := ""
	accessPolicyType := new(string)
	if !r.AccessPolicyType.IsUnknown() && !r.AccessPolicyType.IsNull() {
		*accessPolicyType = r.AccessPolicyType.ValueString()
	} else {
		accessPolicyType = &emptyString
	}
	var requestSwitchUpdateNetworkSwitchAccessPolicyDot1X *merakigosdk.RequestSwitchUpdateNetworkSwitchAccessPolicyDot1X
	if r.Dot1X != nil {
		controlDirection := r.Dot1X.ControlDirection.ValueString()
		requestSwitchUpdateNetworkSwitchAccessPolicyDot1X = &merakigosdk.RequestSwitchUpdateNetworkSwitchAccessPolicyDot1X{
			ControlDirection: controlDirection,
		}
	}
	guestPortBouncing := new(bool)
	if !r.GuestPortBouncing.IsUnknown() && !r.GuestPortBouncing.IsNull() {
		*guestPortBouncing = r.GuestPortBouncing.ValueBool()
	} else {
		guestPortBouncing = nil
	}
	guestVLANID := new(int64)
	if !r.GuestVLANID.IsUnknown() && !r.GuestVLANID.IsNull() {
		*guestVLANID = r.GuestVLANID.ValueInt64()
	} else {
		guestVLANID = nil
	}
	hostMode := new(string)
	if !r.HostMode.IsUnknown() && !r.HostMode.IsNull() {
		*hostMode = r.HostMode.ValueString()
	} else {
		hostMode = &emptyString
	}
	increaseAccessSpeed := new(bool)
	if !r.IncreaseAccessSpeed.IsUnknown() && !r.IncreaseAccessSpeed.IsNull() {
		*increaseAccessSpeed = r.IncreaseAccessSpeed.ValueBool()
	} else {
		increaseAccessSpeed = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestSwitchUpdateNetworkSwitchAccessPolicyRadius *merakigosdk.RequestSwitchUpdateNetworkSwitchAccessPolicyRadius
	if r.Radius != nil {
		var requestSwitchUpdateNetworkSwitchAccessPolicyRadiusCache *merakigosdk.RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusCache
		if r.Radius.Cache != nil {
			enabled := func() *bool {
				if !r.Radius.Cache.Enabled.IsUnknown() && !r.Radius.Cache.Enabled.IsNull() {
					return r.Radius.Cache.Enabled.ValueBoolPointer()
				}
				return nil
			}()
			timeout := func() *int64 {
				if !r.Radius.Cache.Timeout.IsUnknown() && !r.Radius.Cache.Timeout.IsNull() {
					return r.Radius.Cache.Timeout.ValueInt64Pointer()
				}
				return nil
			}()
			requestSwitchUpdateNetworkSwitchAccessPolicyRadiusCache = &merakigosdk.RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusCache{
				Enabled: enabled,
				Timeout: int64ToIntPointer(timeout),
			}
		}
		var requestSwitchUpdateNetworkSwitchAccessPolicyRadiusCriticalAuth *merakigosdk.RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusCriticalAuth
		if r.Radius.CriticalAuth != nil {
			dataVLANID := func() *int64 {
				if !r.Radius.CriticalAuth.DataVLANID.IsUnknown() && !r.Radius.CriticalAuth.DataVLANID.IsNull() {
					return r.Radius.CriticalAuth.DataVLANID.ValueInt64Pointer()
				}
				return nil
			}()
			suspendPortBounce := func() *bool {
				if !r.Radius.CriticalAuth.SuspendPortBounce.IsUnknown() && !r.Radius.CriticalAuth.SuspendPortBounce.IsNull() {
					return r.Radius.CriticalAuth.SuspendPortBounce.ValueBoolPointer()
				}
				return nil
			}()
			voiceVLANID := func() *int64 {
				if !r.Radius.CriticalAuth.VoiceVLANID.IsUnknown() && !r.Radius.CriticalAuth.VoiceVLANID.IsNull() {
					return r.Radius.CriticalAuth.VoiceVLANID.ValueInt64Pointer()
				}
				return nil
			}()
			requestSwitchUpdateNetworkSwitchAccessPolicyRadiusCriticalAuth = &merakigosdk.RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusCriticalAuth{
				DataVLANID:        int64ToIntPointer(dataVLANID),
				SuspendPortBounce: suspendPortBounce,
				VoiceVLANID:       int64ToIntPointer(voiceVLANID),
			}
		}
		failedAuthVLANID := func() *int64 {
			if !r.Radius.FailedAuthVLANID.IsUnknown() && !r.Radius.FailedAuthVLANID.IsNull() {
				return r.Radius.FailedAuthVLANID.ValueInt64Pointer()
			}
			return nil
		}()
		reAuthenticationInterval := func() *int64 {
			if !r.Radius.ReAuthenticationInterval.IsUnknown() && !r.Radius.ReAuthenticationInterval.IsNull() {
				return r.Radius.ReAuthenticationInterval.ValueInt64Pointer()
			}
			return nil
		}()
		requestSwitchUpdateNetworkSwitchAccessPolicyRadius = &merakigosdk.RequestSwitchUpdateNetworkSwitchAccessPolicyRadius{
			Cache:                    requestSwitchUpdateNetworkSwitchAccessPolicyRadiusCache,
			CriticalAuth:             requestSwitchUpdateNetworkSwitchAccessPolicyRadiusCriticalAuth,
			FailedAuthVLANID:         int64ToIntPointer(failedAuthVLANID),
			ReAuthenticationInterval: int64ToIntPointer(reAuthenticationInterval),
		}
	}
	radiusAccountingEnabled := new(bool)
	if !r.RadiusAccountingEnabled.IsUnknown() && !r.RadiusAccountingEnabled.IsNull() {
		*radiusAccountingEnabled = r.RadiusAccountingEnabled.ValueBool()
	} else {
		radiusAccountingEnabled = nil
	}
	var requestSwitchUpdateNetworkSwitchAccessPolicyRadiusAccountingServers []merakigosdk.RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusAccountingServers
	if r.RadiusAccountingServers != nil {
		for _, rItem1 := range *r.RadiusAccountingServers {
			host := rItem1.Host.ValueString()
			organizationRadiusServerID := rItem1.OrganizationRadiusServerID.ValueString()
			port := func() *int64 {
				if !rItem1.Port.IsUnknown() && !rItem1.Port.IsNull() {
					return rItem1.Port.ValueInt64Pointer()
				}
				return nil
			}()
			secret := rItem1.Secret.ValueString()
			serverID := rItem1.ServerID.ValueString()
			requestSwitchUpdateNetworkSwitchAccessPolicyRadiusAccountingServers = append(requestSwitchUpdateNetworkSwitchAccessPolicyRadiusAccountingServers, merakigosdk.RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusAccountingServers{
				Host:                       host,
				OrganizationRadiusServerID: organizationRadiusServerID,
				Port:                       int64ToIntPointer(port),
				Secret:                     secret,
				ServerID:                   serverID,
			})
		}
	}
	radiusCoaSupportEnabled := new(bool)
	if !r.RadiusCoaSupportEnabled.IsUnknown() && !r.RadiusCoaSupportEnabled.IsNull() {
		*radiusCoaSupportEnabled = r.RadiusCoaSupportEnabled.ValueBool()
	} else {
		radiusCoaSupportEnabled = nil
	}
	radiusGroupAttribute := new(string)
	if !r.RadiusGroupAttribute.IsUnknown() && !r.RadiusGroupAttribute.IsNull() {
		*radiusGroupAttribute = r.RadiusGroupAttribute.ValueString()
	} else {
		radiusGroupAttribute = &emptyString
	}
	var requestSwitchUpdateNetworkSwitchAccessPolicyRadiusServers []merakigosdk.RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusServers
	if r.RadiusServers != nil {
		for _, rItem1 := range *r.RadiusServers {
			host := rItem1.Host.ValueString()
			organizationRadiusServerID := rItem1.OrganizationRadiusServerID.ValueString()
			port := func() *int64 {
				if !rItem1.Port.IsUnknown() && !rItem1.Port.IsNull() {
					return rItem1.Port.ValueInt64Pointer()
				}
				return nil
			}()
			secret := rItem1.Secret.ValueString()
			serverID := rItem1.ServerID.ValueString()
			requestSwitchUpdateNetworkSwitchAccessPolicyRadiusServers = append(requestSwitchUpdateNetworkSwitchAccessPolicyRadiusServers, merakigosdk.RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusServers{
				Host:                       host,
				OrganizationRadiusServerID: organizationRadiusServerID,
				Port:                       int64ToIntPointer(port),
				Secret:                     secret,
				ServerID:                   serverID,
			})
		}
	}
	radiusTestingEnabled := new(bool)
	if !r.RadiusTestingEnabled.IsUnknown() && !r.RadiusTestingEnabled.IsNull() {
		*radiusTestingEnabled = r.RadiusTestingEnabled.ValueBool()
	} else {
		radiusTestingEnabled = nil
	}
	uRLRedirectWalledGardenEnabled := new(bool)
	if !r.URLRedirectWalledGardenEnabled.IsUnknown() && !r.URLRedirectWalledGardenEnabled.IsNull() {
		*uRLRedirectWalledGardenEnabled = r.URLRedirectWalledGardenEnabled.ValueBool()
	} else {
		uRLRedirectWalledGardenEnabled = nil
	}
	var uRLRedirectWalledGardenRanges []string = nil
	r.URLRedirectWalledGardenRanges.ElementsAs(ctx, &uRLRedirectWalledGardenRanges, false)
	voiceVLANClients := new(bool)
	if !r.VoiceVLANClients.IsUnknown() && !r.VoiceVLANClients.IsNull() {
		*voiceVLANClients = r.VoiceVLANClients.ValueBool()
	} else {
		voiceVLANClients = nil
	}
	out := merakigosdk.RequestSwitchUpdateNetworkSwitchAccessPolicy{
		AccessPolicyType:        *accessPolicyType,
		Dot1X:                   requestSwitchUpdateNetworkSwitchAccessPolicyDot1X,
		GuestPortBouncing:       guestPortBouncing,
		GuestVLANID:             int64ToIntPointer(guestVLANID),
		HostMode:                *hostMode,
		IncreaseAccessSpeed:     increaseAccessSpeed,
		Name:                    *name,
		Radius:                  requestSwitchUpdateNetworkSwitchAccessPolicyRadius,
		RadiusAccountingEnabled: radiusAccountingEnabled,
		RadiusAccountingServers: func() *[]merakigosdk.RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusAccountingServers {
			if len(requestSwitchUpdateNetworkSwitchAccessPolicyRadiusAccountingServers) > 0 {
				return &requestSwitchUpdateNetworkSwitchAccessPolicyRadiusAccountingServers
			}
			return nil
		}(),
		RadiusCoaSupportEnabled: radiusCoaSupportEnabled,
		RadiusGroupAttribute:    *radiusGroupAttribute,
		RadiusServers: func() *[]merakigosdk.RequestSwitchUpdateNetworkSwitchAccessPolicyRadiusServers {
			if len(requestSwitchUpdateNetworkSwitchAccessPolicyRadiusServers) > 0 {
				return &requestSwitchUpdateNetworkSwitchAccessPolicyRadiusServers
			}
			return nil
		}(),
		RadiusTestingEnabled:           radiusTestingEnabled,
		URLRedirectWalledGardenEnabled: uRLRedirectWalledGardenEnabled,
		URLRedirectWalledGardenRanges:  uRLRedirectWalledGardenRanges,
		VoiceVLANClients:               voiceVLANClients,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchAccessPolicyItemToBodyRs(state NetworksSwitchAccessPoliciesRs, response *merakigosdk.ResponseSwitchGetNetworkSwitchAccessPolicy, is_read bool) NetworksSwitchAccessPoliciesRs {
	itemState := NetworksSwitchAccessPoliciesRs{
		AccessPolicyNumber: types.StringValue(response.AccessPolicyNumber),
		AccessPolicyType:   types.StringValue(response.AccessPolicyType),
		Counts: func() *ResponseSwitchGetNetworkSwitchAccessPolicyCountsRs {
			if response.Counts != nil {
				return &ResponseSwitchGetNetworkSwitchAccessPolicyCountsRs{
					Ports: func() *ResponseSwitchGetNetworkSwitchAccessPolicyCountsPortsRs {
						if response.Counts.Ports != nil {
							return &ResponseSwitchGetNetworkSwitchAccessPolicyCountsPortsRs{
								WithThisPolicy: func() types.Int64 {
									if response.Counts.Ports.WithThisPolicy != nil {
										return types.Int64Value(int64(*response.Counts.Ports.WithThisPolicy))
									}
									return types.Int64{}
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		Dot1X: func() *ResponseSwitchGetNetworkSwitchAccessPolicyDot1XRs {
			if response.Dot1X != nil {
				return &ResponseSwitchGetNetworkSwitchAccessPolicyDot1XRs{
					ControlDirection: types.StringValue(response.Dot1X.ControlDirection),
				}
			}
			return nil
		}(),
		GuestPortBouncing: func() types.Bool {
			if response.GuestPortBouncing != nil {
				return types.BoolValue(*response.GuestPortBouncing)
			}
			return types.Bool{}
		}(),
		GuestVLANID: func() types.Int64 {
			if response.GuestVLANID != nil {
				return types.Int64Value(int64(*response.GuestVLANID))
			}
			return types.Int64{}
		}(),
		HostMode: types.StringValue(response.HostMode),
		IncreaseAccessSpeed: func() types.Bool {
			if response.IncreaseAccessSpeed != nil {
				return types.BoolValue(*response.IncreaseAccessSpeed)
			}
			return types.Bool{}
		}(),
		Name: types.StringValue(response.Name),
		Radius: func() *ResponseSwitchGetNetworkSwitchAccessPolicyRadiusRs {
			if response.Radius != nil {
				return &ResponseSwitchGetNetworkSwitchAccessPolicyRadiusRs{
					// Cache: func() *ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCacheRs {
					// 	if response.Radius.Cache != nil {
					// 		return &ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCacheRs{
					// 			Enabled: func() types.Bool {
					// 				if response.Radius.Cache.Enabled != nil {
					// 					return types.BoolValue(*response.Radius.Cache.Enabled)
					// 				}
					// 				return types.Bool{}
					// 			}(),
					// 			Timeout: func() types.Int64 {
					// 				if response.Radius.Cache.Timeout != nil {
					// 					return types.Int64Value(int64(*response.Radius.Cache.Timeout))
					// 				}
					// 				return types.Int64{}
					// 			}(),
					// 		}
					// 	}
					// 	return nil
					// }(),
					CriticalAuth: func() *ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCriticalAuthRs {
						if response.Radius.CriticalAuth != nil {
							return &ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCriticalAuthRs{
								DataVLANID: func() types.Int64 {
									if response.Radius.CriticalAuth.DataVLANID != nil {
										return types.Int64Value(int64(*response.Radius.CriticalAuth.DataVLANID))
									}
									return types.Int64{}
								}(),
								SuspendPortBounce: func() types.Bool {
									if response.Radius.CriticalAuth.SuspendPortBounce != nil {
										return types.BoolValue(*response.Radius.CriticalAuth.SuspendPortBounce)
									}
									return types.Bool{}
								}(),
								VoiceVLANID: func() types.Int64 {
									if response.Radius.CriticalAuth.VoiceVLANID != nil {
										return types.Int64Value(int64(*response.Radius.CriticalAuth.VoiceVLANID))
									}
									return types.Int64{}
								}(),
							}
						}
						return nil
					}(),
					FailedAuthVLANID: func() types.Int64 {
						if response.Radius.FailedAuthVLANID != nil {
							return types.Int64Value(int64(*response.Radius.FailedAuthVLANID))
						}
						return types.Int64{}
					}(),
					ReAuthenticationInterval: func() types.Int64 {
						if response.Radius.ReAuthenticationInterval != nil {
							return types.Int64Value(int64(*response.Radius.ReAuthenticationInterval))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
		RadiusAccountingEnabled: func() types.Bool {
			if response.RadiusAccountingEnabled != nil {
				return types.BoolValue(*response.RadiusAccountingEnabled)
			}
			return types.Bool{}
		}(),
		RadiusAccountingServersResponse: func() *[]ResponseSwitchGetNetworkSwitchAccessPolicyRadiusAccountingServersRs {
			if response.RadiusAccountingServers != nil {
				result := make([]ResponseSwitchGetNetworkSwitchAccessPolicyRadiusAccountingServersRs, len(*response.RadiusAccountingServers))
				for i, radiusAccountingServers := range *response.RadiusAccountingServers {
					result[i] = ResponseSwitchGetNetworkSwitchAccessPolicyRadiusAccountingServersRs{
						Host:                       types.StringValue(radiusAccountingServers.Host),
						OrganizationRadiusServerID: types.StringValue(radiusAccountingServers.OrganizationRadiusServerID),
						Port: func() types.Int64 {
							if radiusAccountingServers.Port != nil {
								return types.Int64Value(int64(*radiusAccountingServers.Port))
							}
							return types.Int64{}
						}(),
						ServerID: types.StringValue(radiusAccountingServers.ServerID),
					}
				}
				return &result
			}
			return nil
		}(),
		RadiusCoaSupportEnabled: func() types.Bool {
			if response.RadiusCoaSupportEnabled != nil {
				return types.BoolValue(*response.RadiusCoaSupportEnabled)
			}
			return types.Bool{}
		}(),
		RadiusGroupAttribute: types.StringValue(response.RadiusGroupAttribute),
		RadiusServersResponse: func() *[]ResponseSwitchGetNetworkSwitchAccessPolicyRadiusServersRs {
			if response.RadiusServers != nil {
				result := make([]ResponseSwitchGetNetworkSwitchAccessPolicyRadiusServersRs, len(*response.RadiusServers))
				for i, radiusServers := range *response.RadiusServers {
					result[i] = ResponseSwitchGetNetworkSwitchAccessPolicyRadiusServersRs{
						Host:                       types.StringValue(radiusServers.Host),
						OrganizationRadiusServerID: types.StringValue(radiusServers.OrganizationRadiusServerID),
						Port: func() types.Int64 {
							if radiusServers.Port != nil {
								return types.Int64Value(int64(*radiusServers.Port))
							}
							return types.Int64{}
						}(),
						ServerID: types.StringValue(radiusServers.ServerID),
					}
				}
				return &result
			}
			return nil
		}(),
		RadiusTestingEnabled: func() types.Bool {
			if response.RadiusTestingEnabled != nil {
				return types.BoolValue(*response.RadiusTestingEnabled)
			}
			return types.Bool{}
		}(),
		URLRedirectWalledGardenEnabled: func() types.Bool {
			if response.URLRedirectWalledGardenEnabled != nil {
				return types.BoolValue(*response.URLRedirectWalledGardenEnabled)
			}
			return types.Bool{}
		}(),
		URLRedirectWalledGardenRanges: StringSliceToSet(response.URLRedirectWalledGardenRanges),
		VoiceVLANClients: func() types.Bool {
			if response.VoiceVLANClients != nil {
				return types.BoolValue(*response.VoiceVLANClients)
			}
			return types.Bool{}
		}(),
	}

	itemState.RadiusServers = state.RadiusServers
	itemState.RadiusAccountingServers = state.RadiusAccountingServers
	// itemState.Name = state.Name
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSwitchAccessPoliciesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSwitchAccessPoliciesRs)
}
