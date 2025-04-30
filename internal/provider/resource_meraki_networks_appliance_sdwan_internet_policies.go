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

// RESOURCE ACTION

import (
	"context"

	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceSdwanInternetPoliciesResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceSdwanInternetPoliciesResource{}
)

func NewNetworksApplianceSdwanInternetPoliciesResource() resource.Resource {
	return &NetworksApplianceSdwanInternetPoliciesResource{}
}

type NetworksApplianceSdwanInternetPoliciesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceSdwanInternetPoliciesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceSdwanInternetPoliciesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_sdwan_internet_policies"
}

// resourceAction
func (r *NetworksApplianceSdwanInternetPoliciesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"wan_traffic_uplink_preferences": schema.SetNestedAttribute{
						MarkdownDescription: `policies with respective traffic filters for an MX network`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"fail_over_criterion": schema.StringAttribute{
									MarkdownDescription: `WAN failover and failback behavior
                                                Allowed values: [poorPerformance,uplinkDown]`,
									Computed: true,
								},
								"performance_class": schema.SingleNestedAttribute{
									MarkdownDescription: `Performance class setting for uplink preference rule`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"builtin_performance_class_name": schema.StringAttribute{
											MarkdownDescription: `Name of builtin performance class. Must be present when performanceClass type is 'builtin' and value must be one of: 'VoIP'
                                                      Allowed values: [VoIP]`,
											Computed: true,
										},
										"custom_performance_class_id": schema.StringAttribute{
											MarkdownDescription: `ID of created custom performance class, must be present when performanceClass type is "custom"`,
											Computed:            true,
										},
										"type": schema.StringAttribute{
											MarkdownDescription: `Type of this performance class. Must be one of: 'builtin' or 'custom'
                                                      Allowed values: [builtin,custom]`,
											Computed: true,
										},
									},
								},
								"preferred_uplink": schema.StringAttribute{
									MarkdownDescription: `Preferred uplink for uplink preference rule. Must be one of: 'wan1', 'wan2', 'bestForVoIP', 'loadBalancing' or 'defaultUplink'
                                                Allowed values: [bestForVoIP,defaultUplink,loadBalancing,wan1,wan2]`,
									Computed: true,
								},
								"traffic_filters": schema.SetNestedAttribute{
									MarkdownDescription: `Traffic filters`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"type": schema.StringAttribute{
												MarkdownDescription: `Traffic filter type. Must be 'custom', 'major_application', 'application (NBAR)', if type is 'application', you can pass either an NBAR App Category or Application
                                                      Allowed values: [application,custom,majorApplication]`,
												Computed: true,
											},
											"value": schema.SingleNestedAttribute{
												MarkdownDescription: `Value of traffic filter`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"destination": schema.SingleNestedAttribute{
														MarkdownDescription: `Destination of 'custom' type traffic filter`,
														Computed:            true,
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
															},
															"port": schema.StringAttribute{
																MarkdownDescription: `E.g.: "any", "0" (also means "any"), "8080", "1-1024"`,
																Computed:            true,
															},
														},
													},
													"protocol": schema.StringAttribute{
														MarkdownDescription: `Protocol of the traffic filter. Must be one of: 'tcp', 'udp', 'icmp6' or 'any'
                                                            Allowed values: [any,icmp6,tcp,udp]`,
														Computed: true,
													},
													"source": schema.SingleNestedAttribute{
														MarkdownDescription: `Source of traffic filter`,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"cidr": schema.StringAttribute{
																MarkdownDescription: `CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any". Cannot be used in combination with the "vlan" property`,
																Computed:            true,
															},
															"host": schema.Int64Attribute{
																MarkdownDescription: `Host ID in the VLAN. Should not exceed the VLAN subnet capacity. Must be used along with the "vlan" property and is currently only available under a template network.`,
																Computed:            true,
															},
															"port": schema.StringAttribute{
																MarkdownDescription: `E.g.: "any", "0" (also means "any"), "8080", "1-1024"`,
																Computed:            true,
															},
															"vlan": schema.Int64Attribute{
																MarkdownDescription: `VLAN ID of the configured VLAN in the Meraki network. Cannot be used in combination with the "cidr" property and is currently only available under a template network.`,
																Computed:            true,
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
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"wan_traffic_uplink_preferences": schema.SetNestedAttribute{
						MarkdownDescription: `policies with respective traffic filters for an MX network`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"fail_over_criterion": schema.StringAttribute{
									MarkdownDescription: `WAN failover and failback behavior
                                              Allowed values: [poorPerformance,uplinkDown]`,
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"performance_class": schema.SingleNestedAttribute{
									MarkdownDescription: `Performance class setting for uplink preference rule`,
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"builtin_performance_class_name": schema.StringAttribute{
											MarkdownDescription: `Name of builtin performance class. Must be present when performanceClass type is 'builtin' and value must be one of: 'VoIP'
                                                    Allowed values: [VoIP]`,
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.RequiresReplace(),
											},
										},
										"custom_performance_class_id": schema.StringAttribute{
											MarkdownDescription: `ID of created custom performance class, must be present when performanceClass type is "custom"`,
											Optional:            true,
											Computed:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.RequiresReplace(),
											},
										},
										"type": schema.StringAttribute{
											MarkdownDescription: `Type of this performance class. Must be one of: 'builtin' or 'custom'
                                                    Allowed values: [builtin,custom]`,
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.RequiresReplace(),
											},
										},
									},
								},
								"preferred_uplink": schema.StringAttribute{
									MarkdownDescription: `Preferred uplink for uplink preference rule. Must be one of: 'wan1', 'wan2', 'bestForVoIP', 'loadBalancing' or 'defaultUplink'
                                              Allowed values: [bestForVoIP,defaultUplink,loadBalancing,wan1,wan2]`,
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"traffic_filters": schema.SetNestedAttribute{
									MarkdownDescription: `Traffic filters`,
									Optional:            true,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"type": schema.StringAttribute{
												MarkdownDescription: `Traffic filter type. Must be 'custom', 'major_application', 'application (NBAR)', if type is 'application', you can pass either an NBAR App Category or Application
                                                    Allowed values: [application,custom,majorApplication]`,
												Optional: true,
												Computed: true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
											"value": schema.SingleNestedAttribute{
												MarkdownDescription: `Value of traffic filter`,
												Optional:            true,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"destination": schema.SingleNestedAttribute{
														MarkdownDescription: `Destination of 'custom' type traffic filter`,
														Optional:            true,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"applications": schema.SetNestedAttribute{
																MarkdownDescription: `list of application objects (either majorApplication or nbar)`,
																Optional:            true,
																Computed:            true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{

																		"id": schema.StringAttribute{
																			MarkdownDescription: `Id of the major application, or a list of NBAR Application Category or Application selections`,
																			Optional:            true,
																			Computed:            true,
																			PlanModifiers: []planmodifier.String{
																				stringplanmodifier.RequiresReplace(),
																			},
																		},
																		"name": schema.StringAttribute{
																			MarkdownDescription: `Name of the major application or application category selected`,
																			Optional:            true,
																			Computed:            true,
																			PlanModifiers: []planmodifier.String{
																				stringplanmodifier.RequiresReplace(),
																			},
																		},
																		"type": schema.StringAttribute{
																			MarkdownDescription: `app type (major or nbar)`,
																			Optional:            true,
																			Computed:            true,
																			PlanModifiers: []planmodifier.String{
																				stringplanmodifier.RequiresReplace(),
																			},
																		},
																	},
																},
															},
															"cidr": schema.StringAttribute{
																MarkdownDescription: `CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any"`,
																Optional:            true,
																Computed:            true,
																PlanModifiers: []planmodifier.String{
																	stringplanmodifier.RequiresReplace(),
																},
															},
															"port": schema.StringAttribute{
																MarkdownDescription: `E.g.: "any", "0" (also means "any"), "8080", "1-1024"`,
																Optional:            true,
																Computed:            true,
																PlanModifiers: []planmodifier.String{
																	stringplanmodifier.RequiresReplace(),
																},
															},
														},
													},
													"protocol": schema.StringAttribute{
														MarkdownDescription: `Protocol of the traffic filter. Must be one of: 'tcp', 'udp', 'icmp6' or 'any'
                                                          Allowed values: [any,icmp6,tcp,udp]`,
														Optional: true,
														Computed: true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.RequiresReplace(),
														},
													},
													"source": schema.SingleNestedAttribute{
														MarkdownDescription: `Source of traffic filter`,
														Optional:            true,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"cidr": schema.StringAttribute{
																MarkdownDescription: `CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any". Cannot be used in combination with the "vlan" property`,
																Optional:            true,
																Computed:            true,
																PlanModifiers: []planmodifier.String{
																	stringplanmodifier.RequiresReplace(),
																},
															},
															"host": schema.Int64Attribute{
																MarkdownDescription: `Host ID in the VLAN. Should not exceed the VLAN subnet capacity. Must be used along with the "vlan" property and is currently only available under a template network.`,
																Optional:            true,
																Computed:            true,
																PlanModifiers: []planmodifier.Int64{
																	int64planmodifier.RequiresReplace(),
																},
															},
															"port": schema.StringAttribute{
																MarkdownDescription: `E.g.: "any", "0" (also means "any"), "8080", "1-1024"`,
																Optional:            true,
																Computed:            true,
																PlanModifiers: []planmodifier.String{
																	stringplanmodifier.RequiresReplace(),
																},
															},
															"vlan": schema.Int64Attribute{
																MarkdownDescription: `VLAN ID of the configured VLAN in the Meraki network. Cannot be used in combination with the "cidr" property and is currently only available under a template network.`,
																Optional:            true,
																Computed:            true,
																PlanModifiers: []planmodifier.Int64{
																	int64planmodifier.RequiresReplace(),
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
			},
		},
	}
}
func (r *NetworksApplianceSdwanInternetPoliciesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceSdwanInternetPolicies

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
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp1, err := r.client.Appliance.UpdateNetworkApplianceSdwanInternetPolicies(vvNetworkID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceSdwanInternetPolicies",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceSdwanInternetPolicies",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceSdwanInternetPoliciesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksApplianceSdwanInternetPoliciesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksApplianceSdwanInternetPoliciesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceSdwanInternetPolicies struct {
	NetworkID  types.String                                                   `tfsdk:"network_id"`
	Item       *ResponseApplianceUpdateNetworkApplianceSdwanInternetPolicies  `tfsdk:"item"`
	Parameters *RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesRs `tfsdk:"parameters"`
}

type ResponseApplianceUpdateNetworkApplianceSdwanInternetPolicies struct {
	WanTrafficUplinkPreferences *[]ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences `tfsdk:"wan_traffic_uplink_preferences"`
}

type ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences struct {
	FailOverCriterion types.String                                                                                             `tfsdk:"fail_over_criterion"`
	PerformanceClass  *ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesPerformanceClass `tfsdk:"performance_class"`
	PreferredUplink   types.String                                                                                             `tfsdk:"preferred_uplink"`
	TrafficFilters    *[]ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters `tfsdk:"traffic_filters"`
}

type ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesPerformanceClass struct {
	BuiltinPerformanceClassName types.String `tfsdk:"builtin_performance_class_name"`
	CustomPerformanceClassID    types.String `tfsdk:"custom_performance_class_id"`
	Type                        types.String `tfsdk:"type"`
}

type ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters struct {
	Type  types.String                                                                                                `tfsdk:"type"`
	Value *ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValue `tfsdk:"value"`
}

type ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValue struct {
	Destination *ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestination `tfsdk:"destination"`
	Protocol    types.String                                                                                                           `tfsdk:"protocol"`
	Source      *ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueSource      `tfsdk:"source"`
}

type ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestination struct {
	Applications *[]ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications `tfsdk:"applications"`
	Cidr         types.String                                                                                                                         `tfsdk:"cidr"`
	Port         types.String                                                                                                                         `tfsdk:"port"`
}

type ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

type ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueSource struct {
	Cidr types.String `tfsdk:"cidr"`
	Host types.Int64  `tfsdk:"host"`
	Port types.String `tfsdk:"port"`
	VLAN types.Int64  `tfsdk:"vlan"`
}

type RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesRs struct {
	WanTrafficUplinkPreferences *[]RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesRs `tfsdk:"wan_traffic_uplink_preferences"`
}

type RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesRs struct {
	FailOverCriterion types.String                                                                                              `tfsdk:"fail_over_criterion"`
	PerformanceClass  *RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesPerformanceClassRs `tfsdk:"performance_class"`
	PreferredUplink   types.String                                                                                              `tfsdk:"preferred_uplink"`
	TrafficFilters    *[]RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersRs `tfsdk:"traffic_filters"`
}

type RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesPerformanceClassRs struct {
	BuiltinPerformanceClassName types.String `tfsdk:"builtin_performance_class_name"`
	CustomPerformanceClassID    types.String `tfsdk:"custom_performance_class_id"`
	Type                        types.String `tfsdk:"type"`
}

type RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersRs struct {
	Type  types.String                                                                                                 `tfsdk:"type"`
	Value *RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueRs `tfsdk:"value"`
}

type RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueRs struct {
	Destination *RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationRs `tfsdk:"destination"`
	Protocol    types.String                                                                                                            `tfsdk:"protocol"`
	Source      *RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueSourceRs      `tfsdk:"source"`
}

type RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationRs struct {
	Applications *[]RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplicationsRs `tfsdk:"applications"`
	Cidr         types.String                                                                                                                          `tfsdk:"cidr"`
	Port         types.String                                                                                                                          `tfsdk:"port"`
}

type RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplicationsRs struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

type RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueSourceRs struct {
	Cidr types.String `tfsdk:"cidr"`
	Host types.Int64  `tfsdk:"host"`
	Port types.String `tfsdk:"port"`
	VLAN types.Int64  `tfsdk:"vlan"`
}

// FromBody
func (r *NetworksApplianceSdwanInternetPolicies) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceSdwanInternetPolicies {
	re := *r.Parameters
	var requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences []merakigosdk.RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences

	if re.WanTrafficUplinkPreferences != nil {
		for _, rItem1 := range *re.WanTrafficUplinkPreferences {
			failOverCriterion := rItem1.FailOverCriterion.ValueString()
			var requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesPerformanceClass *merakigosdk.RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesPerformanceClass

			if rItem1.PerformanceClass != nil {
				builtinPerformanceClassName := rItem1.PerformanceClass.BuiltinPerformanceClassName.ValueString()
				customPerformanceClassID := rItem1.PerformanceClass.CustomPerformanceClassID.ValueString()
				typeR := rItem1.PerformanceClass.Type.ValueString()
				requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesPerformanceClass = &merakigosdk.RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesPerformanceClass{
					BuiltinPerformanceClassName: builtinPerformanceClassName,
					CustomPerformanceClassID:    customPerformanceClassID,
					Type:                        typeR,
				}
				//[debug] Is Array: False
			}
			preferredUplink := rItem1.PreferredUplink.ValueString()

			log.Printf("[DEBUG] #TODO []RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters")
			var requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters []merakigosdk.RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters

			if rItem1.TrafficFilters != nil {
				for _, rItem2 := range *rItem1.TrafficFilters {
					typeR := rItem2.Type.ValueString()
					var requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValue *merakigosdk.RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValue

					if rItem2.Value != nil {
						var requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestination *merakigosdk.RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestination

						if rItem2.Value.Destination != nil {

							log.Printf("[DEBUG] #TODO []RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications")
							var requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications []merakigosdk.RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications

							if rItem2.Value.Destination.Applications != nil {
								for _, rItem3 := range *rItem2.Value.Destination.Applications {
									id := rItem3.ID.ValueString()
									name := rItem3.Name.ValueString()
									typeR := rItem3.Type.ValueString()
									requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications = append(requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications, merakigosdk.RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications{
										ID:   id,
										Name: name,
										Type: typeR,
									})
									//[debug] Is Array: True
								}
							}
							cidr := rItem2.Value.Destination.Cidr.ValueString()
							port := rItem2.Value.Destination.Port.ValueString()
							requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestination = &merakigosdk.RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestination{
								Applications: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications {
									if len(requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications) > 0 {
										return &requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications
									}
									return nil
								}(),
								Cidr: cidr,
								Port: port,
							}
							//[debug] Is Array: False
						}
						protocol := rItem2.Value.Protocol.ValueString()
						var requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueSource *merakigosdk.RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueSource

						if rItem2.Value.Source != nil {
							cidr := rItem2.Value.Source.Cidr.ValueString()
							host := func() *int64 {
								if !rItem2.Value.Source.Host.IsUnknown() && !rItem2.Value.Source.Host.IsNull() {
									return rItem2.Value.Source.Host.ValueInt64Pointer()
								}
								return nil
							}()
							port := rItem2.Value.Source.Port.ValueString()
							vlan := func() *int64 {
								if !rItem2.Value.Source.VLAN.IsUnknown() && !rItem2.Value.Source.VLAN.IsNull() {
									return rItem2.Value.Source.VLAN.ValueInt64Pointer()
								}
								return nil
							}()
							requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueSource = &merakigosdk.RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueSource{
								Cidr: cidr,
								Host: int64ToIntPointer(host),
								Port: port,
								VLAN: int64ToIntPointer(vlan),
							}
							//[debug] Is Array: False
						}
						requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValue = &merakigosdk.RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValue{
							Destination: requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestination,
							Protocol:    protocol,
							Source:      requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueSource,
						}
						//[debug] Is Array: False
					}
					requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters = append(requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters, merakigosdk.RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters{
						Type:  typeR,
						Value: requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValue,
					})
					//[debug] Is Array: True
				}
			}
			requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences = append(requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences, merakigosdk.RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences{
				FailOverCriterion: failOverCriterion,
				PerformanceClass:  requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesPerformanceClass,
				PreferredUplink:   preferredUplink,
				TrafficFilters: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters {
					if len(requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters) > 0 {
						return &requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters
					}
					return nil
				}(),
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceSdwanInternetPolicies{
		WanTrafficUplinkPreferences: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences {
			if len(requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) > 0 {
				return &requestApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences
			}
			return nil
		}(),
	}
	return &out
}

// ToBody
func ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesItemToBody(state NetworksApplianceSdwanInternetPolicies, response *merakigosdk.ResponseApplianceUpdateNetworkApplianceSdwanInternetPolicies) NetworksApplianceSdwanInternetPolicies {
	itemState := ResponseApplianceUpdateNetworkApplianceSdwanInternetPolicies{
		WanTrafficUplinkPreferences: func() *[]ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences {
			if response.WanTrafficUplinkPreferences != nil {
				result := make([]ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences, len(*response.WanTrafficUplinkPreferences))
				for i, wanTrafficUplinkPreferences := range *response.WanTrafficUplinkPreferences {
					result[i] = ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences{
						FailOverCriterion: types.StringValue(wanTrafficUplinkPreferences.FailOverCriterion),
						PerformanceClass: func() *ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesPerformanceClass {
							if wanTrafficUplinkPreferences.PerformanceClass != nil {
								return &ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesPerformanceClass{
									BuiltinPerformanceClassName: types.StringValue(wanTrafficUplinkPreferences.PerformanceClass.BuiltinPerformanceClassName),
									CustomPerformanceClassID:    types.StringValue(wanTrafficUplinkPreferences.PerformanceClass.CustomPerformanceClassID),
									Type:                        types.StringValue(wanTrafficUplinkPreferences.PerformanceClass.Type),
								}
							}
							return nil
						}(),
						PreferredUplink: types.StringValue(wanTrafficUplinkPreferences.PreferredUplink),
						TrafficFilters: func() *[]ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters {
							if wanTrafficUplinkPreferences.TrafficFilters != nil {
								result := make([]ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters, len(*wanTrafficUplinkPreferences.TrafficFilters))
								for i, trafficFilters := range *wanTrafficUplinkPreferences.TrafficFilters {
									result[i] = ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters{
										Type: types.StringValue(trafficFilters.Type),
										Value: func() *ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValue {
											if trafficFilters.Value != nil {
												return &ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValue{
													Destination: func() *ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestination {
														if trafficFilters.Value.Destination != nil {
															return &ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestination{
																Applications: func() *[]ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications {
																	if trafficFilters.Value.Destination.Applications != nil {
																		result := make([]ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications, len(*trafficFilters.Value.Destination.Applications))
																		for i, applications := range *trafficFilters.Value.Destination.Applications {
																			result[i] = ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications{
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
													Source: func() *ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueSource {
														if trafficFilters.Value.Source != nil {
															return &ResponseApplianceUpdateNetworkApplianceSdwanInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersValueSource{
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
	state.Item = &itemState
	return state
}
