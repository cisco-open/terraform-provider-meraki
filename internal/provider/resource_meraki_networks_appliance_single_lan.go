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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceSingleLanResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceSingleLanResource{}
)

func NewNetworksApplianceSingleLanResource() resource.Resource {
	return &NetworksApplianceSingleLanResource{}
}

type NetworksApplianceSingleLanResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceSingleLanResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceSingleLanResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_single_lan"
}

func (r *NetworksApplianceSingleLanResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"appliance_ip": schema.StringAttribute{
				MarkdownDescription: `The local IP of the appliance on the single LAN`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ipv6": schema.SingleNestedAttribute{
				MarkdownDescription: `IPv6 configuration on the single LAN`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Enable IPv6 on single LAN`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"prefix_assignments": schema.SetNestedAttribute{
						MarkdownDescription: `Prefix assignments on the single LAN`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"autonomous": schema.BoolAttribute{
									MarkdownDescription: `Auto assign a /64 prefix from the origin to the single LAN`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Bool{
										boolplanmodifier.UseStateForUnknown(),
									},
								},
								"origin": schema.SingleNestedAttribute{
									MarkdownDescription: `The origin of the prefix`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"interfaces": schema.SetAttribute{
											MarkdownDescription: `Interfaces associated with the prefix`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.Set{
												setplanmodifier.UseStateForUnknown(),
											},

											ElementType: types.StringType,
										},
										"type": schema.StringAttribute{
											MarkdownDescription: `Type of the origin
                                                    Allowed values: [independent,internet]`,
											Computed: true,
											Optional: true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.UseStateForUnknown(),
											},
											Validators: []validator.String{
												stringvalidator.OneOf(
													"independent",
													"internet",
												),
											},
										},
									},
								},
								"static_appliance_ip6": schema.StringAttribute{
									MarkdownDescription: `Manual configuration of the IPv6 Appliance IP`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"static_prefix": schema.StringAttribute{
									MarkdownDescription: `Manual configuration of a /64 prefix on the single LAN`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
					},
				},
			},
			"mandatory_dhcp": schema.SingleNestedAttribute{
				MarkdownDescription: `Mandatory DHCP will enforce that clients connecting to this single LAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Enable Mandatory DHCP on single LAN.`,
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
			"subnet": schema.StringAttribute{
				MarkdownDescription: `The subnet of the single LAN`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *NetworksApplianceSingleLanResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceSingleLanRs

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
	responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceSingleLan(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceSingleLan only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceSingleLan only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceSingleLan(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceSingleLan",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceSingleLan",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceSingleLan(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceSingleLan",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceSingleLan",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceSingleLanItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceSingleLanResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceSingleLanRs

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
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceSingleLan(vvNetworkID)
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
				"Failure when executing GetNetworkApplianceSingleLan",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceSingleLan",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceSingleLanItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksApplianceSingleLanResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceSingleLanResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceSingleLanRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceSingleLan(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceSingleLan",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceSingleLan",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceSingleLanResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceSingleLan", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceSingleLanRs struct {
	NetworkID     types.String                                                  `tfsdk:"network_id"`
	ApplianceIP   types.String                                                  `tfsdk:"appliance_ip"`
	IPv6          *ResponseApplianceGetNetworkApplianceSingleLanIpv6Rs          `tfsdk:"ipv6"`
	MandatoryDhcp *ResponseApplianceGetNetworkApplianceSingleLanMandatoryDhcpRs `tfsdk:"mandatory_dhcp"`
	Subnet        types.String                                                  `tfsdk:"subnet"`
}

type ResponseApplianceGetNetworkApplianceSingleLanIpv6Rs struct {
	Enabled           types.Bool                                                              `tfsdk:"enabled"`
	PrefixAssignments *[]ResponseApplianceGetNetworkApplianceSingleLanIpv6PrefixAssignmentsRs `tfsdk:"prefix_assignments"`
}

type ResponseApplianceGetNetworkApplianceSingleLanIpv6PrefixAssignmentsRs struct {
	Autonomous         types.Bool                                                                  `tfsdk:"autonomous"`
	Origin             *ResponseApplianceGetNetworkApplianceSingleLanIpv6PrefixAssignmentsOriginRs `tfsdk:"origin"`
	StaticApplianceIP6 types.String                                                                `tfsdk:"static_appliance_ip6"`
	StaticPrefix       types.String                                                                `tfsdk:"static_prefix"`
}

type ResponseApplianceGetNetworkApplianceSingleLanIpv6PrefixAssignmentsOriginRs struct {
	Interfaces types.Set    `tfsdk:"interfaces"`
	Type       types.String `tfsdk:"type"`
}

type ResponseApplianceGetNetworkApplianceSingleLanMandatoryDhcpRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// FromBody
func (r *NetworksApplianceSingleLanRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceSingleLan {
	emptyString := ""
	applianceIP := new(string)
	if !r.ApplianceIP.IsUnknown() && !r.ApplianceIP.IsNull() {
		*applianceIP = r.ApplianceIP.ValueString()
	} else {
		applianceIP = &emptyString
	}

	var requestApplianceUpdateNetworkApplianceSingleLanIPv6 *merakigosdk.RequestApplianceUpdateNetworkApplianceSingleLanIPv6
	if r.IPv6 != nil {
		enabled := func() *bool {
			if !r.IPv6.Enabled.IsUnknown() && !r.IPv6.Enabled.IsNull() {
				return r.IPv6.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		var requestApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignments []merakigosdk.RequestApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignments
		if r.IPv6.PrefixAssignments != nil {
			for _, rItem1 := range *r.IPv6.PrefixAssignments {
				autonomous := func() *bool {
					if !rItem1.Autonomous.IsUnknown() && !rItem1.Autonomous.IsNull() {
						return rItem1.Autonomous.ValueBoolPointer()
					}
					return nil
				}()

				var requestApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignmentsOrigin *merakigosdk.RequestApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignmentsOrigin
				if rItem1.Origin != nil {
					var interfaces []string
					rItem1.Origin.Interfaces.ElementsAs(ctx, &interfaces, false)
					typeR := rItem1.Origin.Type.ValueString()

					requestApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignmentsOrigin = &merakigosdk.RequestApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignmentsOrigin{
						Interfaces: interfaces,
						Type:       typeR,
					}
				}

				staticApplianceIP6 := rItem1.StaticApplianceIP6.ValueString()
				staticPrefix := rItem1.StaticPrefix.ValueString()

				requestApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignments = append(requestApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignments, merakigosdk.RequestApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignments{
					Autonomous:         autonomous,
					Origin:             requestApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignmentsOrigin,
					StaticApplianceIP6: staticApplianceIP6,
					StaticPrefix:       staticPrefix,
				})
			}
		}

		requestApplianceUpdateNetworkApplianceSingleLanIPv6 = &merakigosdk.RequestApplianceUpdateNetworkApplianceSingleLanIPv6{
			Enabled: enabled,
			PrefixAssignments: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignments {
				if len(requestApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignments) > 0 {
					return &requestApplianceUpdateNetworkApplianceSingleLanIPv6PrefixAssignments
				}
				return nil
			}(),
		}
	}

	var requestApplianceUpdateNetworkApplianceSingleLanMandatoryDhcp *merakigosdk.RequestApplianceUpdateNetworkApplianceSingleLanMandatoryDhcp
	if r.MandatoryDhcp != nil {
		enabled := func() *bool {
			if !r.MandatoryDhcp.Enabled.IsUnknown() && !r.MandatoryDhcp.Enabled.IsNull() {
				return r.MandatoryDhcp.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		requestApplianceUpdateNetworkApplianceSingleLanMandatoryDhcp = &merakigosdk.RequestApplianceUpdateNetworkApplianceSingleLanMandatoryDhcp{
			Enabled: enabled,
		}
	}

	subnet := new(string)
	if !r.Subnet.IsUnknown() && !r.Subnet.IsNull() {
		*subnet = r.Subnet.ValueString()
	} else {
		subnet = &emptyString
	}

	out := merakigosdk.RequestApplianceUpdateNetworkApplianceSingleLan{
		ApplianceIP:   *applianceIP,
		IPv6:          requestApplianceUpdateNetworkApplianceSingleLanIPv6,
		MandatoryDhcp: requestApplianceUpdateNetworkApplianceSingleLanMandatoryDhcp,
		Subnet:        *subnet,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceSingleLanItemToBodyRs(state NetworksApplianceSingleLanRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceSingleLan, is_read bool) NetworksApplianceSingleLanRs {
	itemState := NetworksApplianceSingleLanRs{
		ApplianceIP: types.StringValue(response.ApplianceIP),
		IPv6: func() *ResponseApplianceGetNetworkApplianceSingleLanIpv6Rs {
			if response.IPv6 != nil {
				return &ResponseApplianceGetNetworkApplianceSingleLanIpv6Rs{
					Enabled: func() types.Bool {
						if response.IPv6.Enabled != nil {
							return types.BoolValue(*response.IPv6.Enabled)
						}
						return types.Bool{}
					}(),
					PrefixAssignments: func() *[]ResponseApplianceGetNetworkApplianceSingleLanIpv6PrefixAssignmentsRs {
						if response.IPv6.PrefixAssignments != nil {
							result := make([]ResponseApplianceGetNetworkApplianceSingleLanIpv6PrefixAssignmentsRs, len(*response.IPv6.PrefixAssignments))
							for i, prefixAssignments := range *response.IPv6.PrefixAssignments {
								result[i] = ResponseApplianceGetNetworkApplianceSingleLanIpv6PrefixAssignmentsRs{
									Autonomous: func() types.Bool {
										if prefixAssignments.Autonomous != nil {
											return types.BoolValue(*prefixAssignments.Autonomous)
										}
										return types.Bool{}
									}(),
									Origin: func() *ResponseApplianceGetNetworkApplianceSingleLanIpv6PrefixAssignmentsOriginRs {
										if prefixAssignments.Origin != nil {
											return &ResponseApplianceGetNetworkApplianceSingleLanIpv6PrefixAssignmentsOriginRs{
												Interfaces: StringSliceToSet(prefixAssignments.Origin.Interfaces),
												Type:       types.StringValue(prefixAssignments.Origin.Type),
											}
										}
										return nil
									}(),
									StaticApplianceIP6: types.StringValue(prefixAssignments.StaticApplianceIP6),
									StaticPrefix:       types.StringValue(prefixAssignments.StaticPrefix),
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
		MandatoryDhcp: func() *ResponseApplianceGetNetworkApplianceSingleLanMandatoryDhcpRs {
			if response.MandatoryDhcp != nil {
				return &ResponseApplianceGetNetworkApplianceSingleLanMandatoryDhcpRs{
					Enabled: func() types.Bool {
						if response.MandatoryDhcp.Enabled != nil {
							return types.BoolValue(*response.MandatoryDhcp.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
		}(),
		Subnet: types.StringValue(response.Subnet),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceSingleLanRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceSingleLanRs)
}
