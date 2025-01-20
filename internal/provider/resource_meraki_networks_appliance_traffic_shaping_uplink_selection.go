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
	_ resource.Resource              = &NetworksApplianceTrafficShapingUplinkSelectionResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceTrafficShapingUplinkSelectionResource{}
)

func NewNetworksApplianceTrafficShapingUplinkSelectionResource() resource.Resource {
	return &NetworksApplianceTrafficShapingUplinkSelectionResource{}
}

type NetworksApplianceTrafficShapingUplinkSelectionResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceTrafficShapingUplinkSelectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceTrafficShapingUplinkSelectionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_traffic_shaping_uplink_selection"
}

func (r *NetworksApplianceTrafficShapingUplinkSelectionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"active_active_auto_vpn_enabled": schema.BoolAttribute{
				MarkdownDescription: `Whether active-active AutoVPN is enabled`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"default_uplink": schema.StringAttribute{
				MarkdownDescription: `The default uplink. Must be one of: 'wan1' or 'wan2'
                                  Allowed values: [wan1,wan2]`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"wan1",
						"wan2",
					),
				},
			},
			"failover_and_failback": schema.SingleNestedAttribute{
				MarkdownDescription: `WAN failover and failback`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"immediate": schema.SingleNestedAttribute{
						MarkdownDescription: `Immediate WAN failover and failback`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Whether immediate WAN failover and failback is enabled`,
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
			"load_balancing_enabled": schema.BoolAttribute{
				MarkdownDescription: `Whether load balancing is enabled`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"vpn_traffic_uplink_preferences": schema.SetNestedAttribute{
				MarkdownDescription: `Uplink preference rules for VPN traffic`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"fail_over_criterion": schema.StringAttribute{
							MarkdownDescription: `Fail over criterion for uplink preference rule. Must be one of: 'poorPerformance' or 'uplinkDown'
                                        Allowed values: [poorPerformance,uplinkDown]`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"poorPerformance",
									"uplinkDown",
								),
							},
						},
						"performance_class": schema.SingleNestedAttribute{
							MarkdownDescription: `Performance class setting for uplink preference rule`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"builtin_performance_class_name": schema.StringAttribute{
									MarkdownDescription: `Name of builtin performance class. Must be present when performanceClass type is 'builtin' and value must be one of: 'VoIP'
                                              Allowed values: [VoIP]`,
									Computed: true,
									Optional: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
									Validators: []validator.String{
										stringvalidator.OneOf(
											"VoIP",
										),
									},
								},
								"custom_performance_class_id": schema.StringAttribute{
									MarkdownDescription: `ID of created custom performance class, must be present when performanceClass type is "custom"`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"type": schema.StringAttribute{
									MarkdownDescription: `Type of this performance class. Must be one of: 'builtin' or 'custom'
                                              Allowed values: [builtin,custom]`,
									Computed: true,
									Optional: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
									Validators: []validator.String{
										stringvalidator.OneOf(
											"builtin",
											"custom",
										),
									},
								},
							},
						},
						"preferred_uplink": schema.StringAttribute{
							MarkdownDescription: `Preferred uplink for uplink preference rule. Must be one of: 'wan1', 'wan2', 'bestForVoIP', 'loadBalancing' or 'defaultUplink'
                                        Allowed values: [bestForVoIP,defaultUplink,loadBalancing,wan1,wan2]`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"bestForVoIP",
									"defaultUplink",
									"loadBalancing",
									"wan1",
									"wan2",
								),
							},
						},
						"traffic_filters": schema.SetNestedAttribute{
							MarkdownDescription: `Traffic filters`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"type": schema.StringAttribute{
										MarkdownDescription: `Traffic filter type. Must be one of: 'applicationCategory', 'application' or 'custom'
                                              Allowed values: [application,applicationCategory,custom]`,
										Computed: true,
										Optional: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										Validators: []validator.String{
											stringvalidator.OneOf(
												"application",
												"applicationCategory",
												"custom",
											),
										},
									},
									"value": schema.SingleNestedAttribute{
										MarkdownDescription: `Value of traffic filter`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"destination": schema.SingleNestedAttribute{
												MarkdownDescription: `Destination of 'custom' type traffic filter`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.Object{
													objectplanmodifier.UseStateForUnknown(),
												},
												Attributes: map[string]schema.Attribute{

													"cidr": schema.StringAttribute{
														MarkdownDescription: `CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any". Cannot be used in combination with the "vlan" or "fqdn" property`,
														Computed:            true,
														Optional:            true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.UseStateForUnknown(),
														},
													},
													"fqdn": schema.StringAttribute{
														MarkdownDescription: `FQDN format address. Cannot be used in combination with the "cidr" or "fqdn" property and is currently only available in the "destination" object of the "vpnTrafficUplinkPreference" object. E.g.: "www.google.com"`,
														Computed:            true,
														Optional:            true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.UseStateForUnknown(),
														},
													},
													"host": schema.Int64Attribute{
														MarkdownDescription: `Host ID in the VLAN. Should not exceed the VLAN subnet capacity. Must be used along with the "vlan" property and is currently only available under a template network.`,
														Computed:            true,
														Optional:            true,
														PlanModifiers: []planmodifier.Int64{
															int64planmodifier.UseStateForUnknown(),
														},
													},
													"network": schema.StringAttribute{
														MarkdownDescription: `Meraki network ID. Currently only available under a template network, and the value should be ID of either same template network, or another template network currently. E.g.: "L_12345678".`,
														Computed:            true,
														Optional:            true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.UseStateForUnknown(),
														},
													},
													"port": schema.StringAttribute{
														MarkdownDescription: `E.g.: "any", "0" (also means "any"), "8080", "1-1024"`,
														Computed:            true,
														Optional:            true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.UseStateForUnknown(),
														},
													},
													"vlan": schema.Int64Attribute{
														MarkdownDescription: `VLAN ID of the configured VLAN in the Meraki network. Cannot be used in combination with the "cidr" or "fqdn" property and is currently only available under a template network.`,
														Computed:            true,
														Optional:            true,
														PlanModifiers: []planmodifier.Int64{
															int64planmodifier.UseStateForUnknown(),
														},
													},
												},
											},
											"id": schema.StringAttribute{
												MarkdownDescription: `ID of 'applicationCategory' or 'application' type traffic filter`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
											"protocol": schema.StringAttribute{
												MarkdownDescription: `Protocol of 'custom' type traffic filter. Must be one of: 'tcp', 'udp', 'icmp', 'icmp6' or 'any'
                                                    Allowed values: [any,icmp,icmp6,tcp,udp]`,
												Computed: true,
												Optional: true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
												Validators: []validator.String{
													stringvalidator.OneOf(
														"any",
														"icmp",
														"icmp6",
														"tcp",
														"udp",
													),
												},
											},
											"source": schema.SingleNestedAttribute{
												MarkdownDescription: `Source of 'custom' type traffic filter`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.Object{
													objectplanmodifier.UseStateForUnknown(),
												},
												Attributes: map[string]schema.Attribute{

													"cidr": schema.StringAttribute{
														MarkdownDescription: `CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any". Cannot be used in combination with the "vlan" property`,
														Computed:            true,
														Optional:            true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.UseStateForUnknown(),
														},
													},
													"host": schema.Int64Attribute{
														MarkdownDescription: `Host ID in the VLAN. Should not exceed the VLAN subnet capacity. Must be used along with the "vlan" property and is currently only available under a template network.`,
														Computed:            true,
														Optional:            true,
														PlanModifiers: []planmodifier.Int64{
															int64planmodifier.UseStateForUnknown(),
														},
													},
													"network": schema.StringAttribute{
														MarkdownDescription: `Meraki network ID. Currently only available under a template network, and the value should be ID of either same template network, or another template network currently. E.g.: "L_12345678".`,
														Computed:            true,
														Optional:            true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.UseStateForUnknown(),
														},
													},
													"port": schema.StringAttribute{
														MarkdownDescription: `E.g.: "any", "0" (also means "any"), "8080", "1-1024"`,
														Computed:            true,
														Optional:            true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.UseStateForUnknown(),
														},
													},
													"vlan": schema.Int64Attribute{
														MarkdownDescription: `VLAN ID of the configured VLAN in the Meraki network. Cannot be used in combination with the "cidr" property and is currently only available under a template network.`,
														Computed:            true,
														Optional:            true,
														PlanModifiers: []planmodifier.Int64{
															int64planmodifier.UseStateForUnknown(),
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"wan_traffic_uplink_preferences": schema.SetNestedAttribute{
				MarkdownDescription: `Uplink preference rules for WAN traffic`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"preferred_uplink": schema.StringAttribute{
							MarkdownDescription: `Preferred uplink for uplink preference rule. Must be one of: 'wan1' or 'wan2'
                                        Allowed values: [wan1,wan2]`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"wan1",
									"wan2",
								),
							},
						},
						"traffic_filters": schema.SetNestedAttribute{
							MarkdownDescription: `Traffic filters`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"type": schema.StringAttribute{
										MarkdownDescription: `Traffic filter type. Must be "custom"
                                              Allowed values: [custom]`,
										Computed: true,
										Optional: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										Validators: []validator.String{
											stringvalidator.OneOf(
												"custom",
											),
										},
									},
									"value": schema.SingleNestedAttribute{
										MarkdownDescription: `Value of traffic filter`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"destination": schema.SingleNestedAttribute{
												MarkdownDescription: `Destination of 'custom' type traffic filter`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.Object{
													objectplanmodifier.UseStateForUnknown(),
												},
												Attributes: map[string]schema.Attribute{

													"applications": schema.SetNestedAttribute{
														MarkdownDescription: `list of application objects (either majorApplication or nbar)`,
														Computed:            true,
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{

																"id": schema.StringAttribute{
																	MarkdownDescription: `Id of the major application, or a list of NBAR Application Category or Application selections`,
																	Computed:            true,
																},
																"name": schema.StringAttribute{
																	MarkdownDescription: `Name of the major application or application category selected`,
																	Computed:            true,
																},
																"type": schema.StringAttribute{
																	MarkdownDescription: `app type (major or nbar)`,
																	Computed:            true,
																},
															},
														},
													},
													"cidr": schema.StringAttribute{
														MarkdownDescription: `CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any"`,
														Computed:            true,
														Optional:            true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.UseStateForUnknown(),
														},
													},
													"port": schema.StringAttribute{
														MarkdownDescription: `E.g.: "any", "0" (also means "any"), "8080", "1-1024"`,
														Computed:            true,
														Optional:            true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.UseStateForUnknown(),
														},
													},
												},
											},
											"protocol": schema.StringAttribute{
												MarkdownDescription: `Protocol of 'custom' type traffic filter. Must be one of: 'tcp', 'udp', 'icmp6' or 'any'
                                                    Allowed values: [any,icmp6,tcp,udp]`,
												Computed: true,
												Optional: true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
												Validators: []validator.String{
													stringvalidator.OneOf(
														"any",
														"icmp6",
														"tcp",
														"udp",
													),
												},
											},
											"source": schema.SingleNestedAttribute{
												MarkdownDescription: `Source of 'custom' type traffic filter`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.Object{
													objectplanmodifier.UseStateForUnknown(),
												},
												Attributes: map[string]schema.Attribute{

													"cidr": schema.StringAttribute{
														MarkdownDescription: `CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any". Cannot be used in combination with the "vlan" property`,
														Computed:            true,
														Optional:            true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.UseStateForUnknown(),
														},
													},
													"host": schema.Int64Attribute{
														MarkdownDescription: `Host ID in the VLAN. Should not exceed the VLAN subnet capacity. Must be used along with the "vlan" property and is currently only available under a template network.`,
														Computed:            true,
														Optional:            true,
														PlanModifiers: []planmodifier.Int64{
															int64planmodifier.UseStateForUnknown(),
														},
													},
													"port": schema.StringAttribute{
														MarkdownDescription: `E.g.: "any", "0" (also means "any"), "8080", "1-1024"`,
														Computed:            true,
														Optional:            true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.UseStateForUnknown(),
														},
													},
													"vlan": schema.Int64Attribute{
														MarkdownDescription: `VLAN ID of the configured VLAN in the Meraki network. Cannot be used in combination with the "cidr" property and is currently only available under a template network.`,
														Computed:            true,
														Optional:            true,
														PlanModifiers: []planmodifier.Int64{
															int64planmodifier.UseStateForUnknown(),
														},
													},
												},
											},
										},
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

func (r *NetworksApplianceTrafficShapingUplinkSelectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceTrafficShapingUplinkSelectionRs

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
	responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceTrafficShapingUplinkSelection(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceTrafficShapingUplinkSelection only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceTrafficShapingUplinkSelection only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceTrafficShapingUplinkSelection(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceTrafficShapingUplinkSelection",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceTrafficShapingUplinkSelection",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceTrafficShapingUplinkSelection(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceTrafficShapingUplinkSelection",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceTrafficShapingUplinkSelection",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceTrafficShapingUplinkSelectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceTrafficShapingUplinkSelectionRs

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
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceTrafficShapingUplinkSelection(vvNetworkID)
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
				"Failure when executing GetNetworkApplianceTrafficShapingUplinkSelection",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceTrafficShapingUplinkSelection",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksApplianceTrafficShapingUplinkSelectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceTrafficShapingUplinkSelectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceTrafficShapingUplinkSelectionRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceTrafficShapingUplinkSelection(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceTrafficShapingUplinkSelection",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceTrafficShapingUplinkSelection",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceTrafficShapingUplinkSelectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceTrafficShapingUplinkSelection", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceTrafficShapingUplinkSelectionRs struct {
	NetworkID                   types.String                                                                                      `tfsdk:"network_id"`
	ActiveActiveAutoVpnEnabled  types.Bool                                                                                        `tfsdk:"active_active_auto_vpn_enabled"`
	DefaultUplink               types.String                                                                                      `tfsdk:"default_uplink"`
	FailoverAndFailback         *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackRs           `tfsdk:"failover_and_failback"`
	LoadBalancingEnabled        types.Bool                                                                                        `tfsdk:"load_balancing_enabled"`
	VpnTrafficUplinkPreferences *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesRs `tfsdk:"vpn_traffic_uplink_preferences"`
	WanTrafficUplinkPreferences *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesRs `tfsdk:"wan_traffic_uplink_preferences"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackRs struct {
	Immediate *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediateRs `tfsdk:"immediate"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediateRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesRs struct {
	FailOverCriterion types.String                                                                                                    `tfsdk:"fail_over_criterion"`
	PerformanceClass  *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesPerformanceClassRs `tfsdk:"performance_class"`
	PreferredUplink   types.String                                                                                                    `tfsdk:"preferred_uplink"`
	TrafficFilters    *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersRs `tfsdk:"traffic_filters"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesPerformanceClassRs struct {
	BuiltinPerformanceClassName types.String `tfsdk:"builtin_performance_class_name"`
	CustomPerformanceClassID    types.String `tfsdk:"custom_performance_class_id"`
	Type                        types.String `tfsdk:"type"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersRs struct {
	Type  types.String                                                                                                       `tfsdk:"type"`
	Value *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueRs `tfsdk:"value"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueRs struct {
	Destination *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueDestinationRs `tfsdk:"destination"`
	ID          types.String                                                                                                                  `tfsdk:"id"`
	Protocol    types.String                                                                                                                  `tfsdk:"protocol"`
	Source      *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueSourceRs      `tfsdk:"source"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueDestinationRs struct {
	Cidr    types.String `tfsdk:"cidr"`
	Fqdn    types.String `tfsdk:"fqdn"`
	Host    types.Int64  `tfsdk:"host"`
	Network types.String `tfsdk:"network"`
	Port    types.String `tfsdk:"port"`
	VLAN    types.Int64  `tfsdk:"vlan"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueSourceRs struct {
	Cidr    types.String `tfsdk:"cidr"`
	Host    types.Int64  `tfsdk:"host"`
	Network types.String `tfsdk:"network"`
	Port    types.String `tfsdk:"port"`
	VLAN    types.Int64  `tfsdk:"vlan"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesRs struct {
	PreferredUplink types.String                                                                                                    `tfsdk:"preferred_uplink"`
	TrafficFilters  *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersRs `tfsdk:"traffic_filters"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersRs struct {
	Type  types.String                                                                                                       `tfsdk:"type"`
	Value *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueRs `tfsdk:"value"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueRs struct {
	Destination *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestinationRs `tfsdk:"destination"`
	Protocol    types.String                                                                                                                  `tfsdk:"protocol"`
	Source      *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueSourceRs      `tfsdk:"source"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestinationRs struct {
	Applications *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplicationsRs `tfsdk:"applications"`
	Cidr         types.String                                                                                                                                `tfsdk:"cidr"`
	Port         types.String                                                                                                                                `tfsdk:"port"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplicationsRs struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueSourceRs struct {
	Cidr types.String `tfsdk:"cidr"`
	Host types.Int64  `tfsdk:"host"`
	Port types.String `tfsdk:"port"`
	VLAN types.Int64  `tfsdk:"vlan"`
}

// FromBody
func (r *NetworksApplianceTrafficShapingUplinkSelectionRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelection {
	emptyString := ""
	activeActiveAutoVpnEnabled := new(bool)
	if !r.ActiveActiveAutoVpnEnabled.IsUnknown() && !r.ActiveActiveAutoVpnEnabled.IsNull() {
		*activeActiveAutoVpnEnabled = r.ActiveActiveAutoVpnEnabled.ValueBool()
	} else {
		activeActiveAutoVpnEnabled = nil
	}
	defaultUplink := new(string)
	if !r.DefaultUplink.IsUnknown() && !r.DefaultUplink.IsNull() {
		*defaultUplink = r.DefaultUplink.ValueString()
	} else {
		defaultUplink = &emptyString
	}
	var requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailback *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailback
	if r.FailoverAndFailback != nil {
		var requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate
		if r.FailoverAndFailback.Immediate != nil {
			enabled := func() *bool {
				if !r.FailoverAndFailback.Immediate.Enabled.IsUnknown() && !r.FailoverAndFailback.Immediate.Enabled.IsNull() {
					return r.FailoverAndFailback.Immediate.Enabled.ValueBoolPointer()
				}
				return nil
			}()
			requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate = &merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate{
				Enabled: enabled,
			}
		}
		requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailback = &merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailback{
			Immediate: requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate,
		}
	}
	loadBalancingEnabled := new(bool)
	if !r.LoadBalancingEnabled.IsUnknown() && !r.LoadBalancingEnabled.IsNull() {
		*loadBalancingEnabled = r.LoadBalancingEnabled.ValueBool()
	} else {
		loadBalancingEnabled = nil
	}
	var requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences []merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences
	if r.VpnTrafficUplinkPreferences != nil {
		for _, rItem1 := range *r.VpnTrafficUplinkPreferences {
			failOverCriterion := rItem1.FailOverCriterion.ValueString()
			var requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesPerformanceClass *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesPerformanceClass
			if rItem1.PerformanceClass != nil {
				builtinPerformanceClassName := rItem1.PerformanceClass.BuiltinPerformanceClassName.ValueString()
				customPerformanceClassID := rItem1.PerformanceClass.CustomPerformanceClassID.ValueString()
				typeR := rItem1.PerformanceClass.Type.ValueString()
				requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesPerformanceClass = &merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesPerformanceClass{
					BuiltinPerformanceClassName: builtinPerformanceClassName,
					CustomPerformanceClassID:    customPerformanceClassID,
					Type:                        typeR,
				}
			}
			preferredUplink := rItem1.PreferredUplink.ValueString()
			var requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters []merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters
			if rItem1.TrafficFilters != nil {
				for _, rItem2 := range *rItem1.TrafficFilters { //TrafficFilters// name: trafficFilters
					typeR := rItem2.Type.ValueString()
					var requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValue *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValue
					if rItem2.Value != nil {
						var requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueDestination *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueDestination
						if rItem2.Value.Destination != nil {
							cidr := rItem2.Value.Destination.Cidr.ValueString()
							fqdn := rItem2.Value.Destination.Fqdn.ValueString()
							host := func() *int64 {
								if !rItem2.Value.Destination.Host.IsUnknown() && !rItem2.Value.Destination.Host.IsNull() {
									return rItem2.Value.Destination.Host.ValueInt64Pointer()
								}
								return nil
							}()
							network := rItem2.Value.Destination.Network.ValueString()
							port := rItem2.Value.Destination.Port.ValueString()
							vLAN := func() *int64 {
								if !rItem2.Value.Destination.VLAN.IsUnknown() && !rItem2.Value.Destination.VLAN.IsNull() {
									return rItem2.Value.Destination.VLAN.ValueInt64Pointer()
								}
								return nil
							}()
							requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueDestination = &merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueDestination{
								Cidr:    cidr,
								Fqdn:    fqdn,
								Host:    int64ToIntPointer(host),
								Network: network,
								Port:    port,
								VLAN:    int64ToIntPointer(vLAN),
							}
						}
						iD := rItem2.Value.ID.ValueString()
						protocol := rItem2.Value.Protocol.ValueString()
						var requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueSource *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueSource
						if rItem2.Value.Source != nil {
							cidr := rItem2.Value.Source.Cidr.ValueString()
							host := func() *int64 {
								if !rItem2.Value.Source.Host.IsUnknown() && !rItem2.Value.Source.Host.IsNull() {
									return rItem2.Value.Source.Host.ValueInt64Pointer()
								}
								return nil
							}()
							network := rItem2.Value.Source.Network.ValueString()
							port := rItem2.Value.Source.Port.ValueString()
							vLAN := func() *int64 {
								if !rItem2.Value.Source.VLAN.IsUnknown() && !rItem2.Value.Source.VLAN.IsNull() {
									return rItem2.Value.Source.VLAN.ValueInt64Pointer()
								}
								return nil
							}()
							requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueSource = &merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueSource{
								Cidr:    cidr,
								Host:    int64ToIntPointer(host),
								Network: network,
								Port:    port,
								VLAN:    int64ToIntPointer(vLAN),
							}
						}
						requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValue = &merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValue{
							Destination: requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueDestination,
							ID:          iD,
							Protocol:    protocol,
							Source:      requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueSource,
						}
					}
					requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters = append(requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters, merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters{
						Type:  typeR,
						Value: requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValue,
					})
				}
			}
			requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences = append(requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences, merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences{
				FailOverCriterion: failOverCriterion,
				PerformanceClass:  requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesPerformanceClass,
				PreferredUplink:   preferredUplink,
				TrafficFilters: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters {
					if len(requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters) > 0 {
						return &requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters
					}
					return nil
				}(),
			})
		}
	}
	var requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences []merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences
	if r.WanTrafficUplinkPreferences != nil {
		for _, rItem1 := range *r.WanTrafficUplinkPreferences {
			preferredUplink := rItem1.PreferredUplink.ValueString()
			var requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters []merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters
			if rItem1.TrafficFilters != nil {
				for _, rItem2 := range *rItem1.TrafficFilters { //TrafficFilters// name: trafficFilters
					typeR := rItem2.Type.ValueString()
					var requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValue *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValue
					if rItem2.Value != nil {
						var requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestination *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestination
						if rItem2.Value.Destination != nil {
							cidr := rItem2.Value.Destination.Cidr.ValueString()
							port := rItem2.Value.Destination.Port.ValueString()
							requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestination = &merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestination{
								Cidr: cidr,
								Port: port,
							}
						}
						protocol := rItem2.Value.Protocol.ValueString()
						var requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueSource *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueSource
						if rItem2.Value.Source != nil {
							cidr := rItem2.Value.Source.Cidr.ValueString()
							host := func() *int64 {
								if !rItem2.Value.Source.Host.IsUnknown() && !rItem2.Value.Source.Host.IsNull() {
									return rItem2.Value.Source.Host.ValueInt64Pointer()
								}
								return nil
							}()
							port := rItem2.Value.Source.Port.ValueString()
							vLAN := func() *int64 {
								if !rItem2.Value.Source.VLAN.IsUnknown() && !rItem2.Value.Source.VLAN.IsNull() {
									return rItem2.Value.Source.VLAN.ValueInt64Pointer()
								}
								return nil
							}()
							requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueSource = &merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueSource{
								Cidr: cidr,
								Host: int64ToIntPointer(host),
								Port: port,
								VLAN: int64ToIntPointer(vLAN),
							}
						}
						requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValue = &merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValue{
							Destination: requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestination,
							Protocol:    protocol,
							Source:      requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueSource,
						}
					}
					requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters = append(requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters, merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters{
						Type:  typeR,
						Value: requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValue,
					})
				}
			}
			requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences = append(requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences, merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences{
				PreferredUplink: preferredUplink,
				TrafficFilters: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters {
					if len(requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters) > 0 {
						return &requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters
					}
					return nil
				}(),
			})
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelection{
		ActiveActiveAutoVpnEnabled: activeActiveAutoVpnEnabled,
		DefaultUplink:              *defaultUplink,
		FailoverAndFailback:        requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailback,
		LoadBalancingEnabled:       loadBalancingEnabled,
		VpnTrafficUplinkPreferences: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences {
			if len(requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences) > 0 {
				return &requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences
			}
			return nil
		}(),
		WanTrafficUplinkPreferences: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences {
			if len(requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences) > 0 {
				return &requestApplianceUpdateNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionItemToBodyRs(state NetworksApplianceTrafficShapingUplinkSelectionRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelection, is_read bool) NetworksApplianceTrafficShapingUplinkSelectionRs {
	itemState := NetworksApplianceTrafficShapingUplinkSelectionRs{
		ActiveActiveAutoVpnEnabled: func() types.Bool {
			if response.ActiveActiveAutoVpnEnabled != nil {
				return types.BoolValue(*response.ActiveActiveAutoVpnEnabled)
			}
			if !state.ActiveActiveAutoVpnEnabled.IsNull() {
				return state.ActiveActiveAutoVpnEnabled
			}
			return types.Bool{}
		}(),
		DefaultUplink: types.StringValue(response.DefaultUplink),
		FailoverAndFailback: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackRs {
			if response.FailoverAndFailback != nil {
				return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackRs{
					Immediate: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediateRs {
						if response.FailoverAndFailback.Immediate != nil {
							return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediateRs{
								Enabled: func() types.Bool {
									if response.FailoverAndFailback.Immediate.Enabled != nil {
										return types.BoolValue(*response.FailoverAndFailback.Immediate.Enabled)
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
		LoadBalancingEnabled: func() types.Bool {
			if response.LoadBalancingEnabled != nil {
				return types.BoolValue(*response.LoadBalancingEnabled)
			}
			return types.Bool{}
		}(),
		VpnTrafficUplinkPreferences: func() *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesRs {
			if response.VpnTrafficUplinkPreferences != nil {
				result := make([]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesRs, len(*response.VpnTrafficUplinkPreferences))
				for i, vpnTrafficUplinkPreferences := range *response.VpnTrafficUplinkPreferences {
					result[i] = ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesRs{
						FailOverCriterion: types.StringValue(vpnTrafficUplinkPreferences.FailOverCriterion),
						PerformanceClass: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesPerformanceClassRs {
							if vpnTrafficUplinkPreferences.PerformanceClass != nil {
								return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesPerformanceClassRs{
									BuiltinPerformanceClassName: types.StringValue(vpnTrafficUplinkPreferences.PerformanceClass.BuiltinPerformanceClassName),
									CustomPerformanceClassID:    types.StringValue(vpnTrafficUplinkPreferences.PerformanceClass.CustomPerformanceClassID),
									Type:                        types.StringValue(vpnTrafficUplinkPreferences.PerformanceClass.Type),
								}
							}
							return nil
						}(),
						PreferredUplink: types.StringValue(vpnTrafficUplinkPreferences.PreferredUplink),
						TrafficFilters: func() *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersRs {
							if vpnTrafficUplinkPreferences.TrafficFilters != nil {
								result := make([]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersRs, len(*vpnTrafficUplinkPreferences.TrafficFilters))
								for i, trafficFilters := range *vpnTrafficUplinkPreferences.TrafficFilters {
									result[i] = ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersRs{
										Type: types.StringValue(trafficFilters.Type),
										Value: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueRs {
											if trafficFilters.Value != nil {
												return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueRs{
													Destination: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueDestinationRs {
														if trafficFilters.Value.Destination != nil {
															return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueDestinationRs{
																Cidr: types.StringValue(trafficFilters.Value.Destination.Cidr),
																Fqdn: types.StringValue(trafficFilters.Value.Destination.Fqdn),
																Host: func() types.Int64 {
																	if trafficFilters.Value.Destination.Host != nil {
																		return types.Int64Value(int64(*trafficFilters.Value.Destination.Host))
																	}
																	return types.Int64{}
																}(),
																Network: types.StringValue(trafficFilters.Value.Destination.Network),
																Port:    types.StringValue(trafficFilters.Value.Destination.Port),
																VLAN: func() types.Int64 {
																	if trafficFilters.Value.Destination.VLAN != nil {
																		return types.Int64Value(int64(*trafficFilters.Value.Destination.VLAN))
																	}
																	return types.Int64{}
																}(),
															}
														}
														return nil
													}(),
													ID:       types.StringValue(trafficFilters.Value.ID),
													Protocol: types.StringValue(trafficFilters.Value.Protocol),
													Source: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueSourceRs {
														if trafficFilters.Value.Source != nil {
															return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueSourceRs{
																Cidr: types.StringValue(trafficFilters.Value.Source.Cidr),
																Host: func() types.Int64 {
																	if trafficFilters.Value.Source.Host != nil {
																		return types.Int64Value(int64(*trafficFilters.Value.Source.Host))
																	}
																	return types.Int64{}
																}(),
																Network: types.StringValue(trafficFilters.Value.Source.Network),
																Port:    types.StringValue(trafficFilters.Value.Source.Port),
																VLAN: func() types.Int64 {
																	if trafficFilters.Value.Source.VLAN != nil {
																		return types.Int64Value(int64(*trafficFilters.Value.Source.VLAN))
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
									}
								}
								return &result
							}
							return nil
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		WanTrafficUplinkPreferences: func() *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesRs {
			if response.WanTrafficUplinkPreferences != nil {
				result := make([]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesRs, len(*response.WanTrafficUplinkPreferences))
				for i, wanTrafficUplinkPreferences := range *response.WanTrafficUplinkPreferences {
					result[i] = ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesRs{
						PreferredUplink: types.StringValue(wanTrafficUplinkPreferences.PreferredUplink),
						TrafficFilters: func() *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersRs {
							if wanTrafficUplinkPreferences.TrafficFilters != nil {
								result := make([]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersRs, len(*wanTrafficUplinkPreferences.TrafficFilters))
								for i, trafficFilters := range *wanTrafficUplinkPreferences.TrafficFilters {
									result[i] = ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersRs{
										Type: types.StringValue(trafficFilters.Type),
										Value: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueRs {
											if trafficFilters.Value != nil {
												return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueRs{
													Destination: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestinationRs {
														if trafficFilters.Value.Destination != nil {
															return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestinationRs{
																Applications: func() *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplicationsRs {
																	if trafficFilters.Value.Destination.Applications != nil {
																		result := make([]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplicationsRs, len(*trafficFilters.Value.Destination.Applications))
																		for i, applications := range *trafficFilters.Value.Destination.Applications {
																			result[i] = ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplicationsRs{
																				ID:   types.StringValue(applications.ID),
																				Name: types.StringValue(applications.Name),
																				Type: types.StringValue(applications.Type),
																			}
																		}
																		return &result
																	}
																	return nil
																}(),
																Cidr: types.StringValue(trafficFilters.Value.Destination.Cidr),
																Port: types.StringValue(trafficFilters.Value.Destination.Port),
															}
														}
														return nil
													}(),
													Protocol: types.StringValue(trafficFilters.Value.Protocol),
													Source: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueSourceRs {
														if trafficFilters.Value.Source != nil {
															return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueSourceRs{
																Cidr: types.StringValue(trafficFilters.Value.Source.Cidr),
																Host: func() types.Int64 {
																	if trafficFilters.Value.Source.Host != nil {
																		return types.Int64Value(int64(*trafficFilters.Value.Source.Host))
																	}
																	return types.Int64{}
																}(),
																Port: types.StringValue(trafficFilters.Value.Source.Port),
																VLAN: func() types.Int64 {
																	if trafficFilters.Value.Source.VLAN != nil {
																		return types.Int64Value(int64(*trafficFilters.Value.Source.VLAN))
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
									}
								}
								return &result
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
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceTrafficShapingUplinkSelectionRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceTrafficShapingUplinkSelectionRs)
}
