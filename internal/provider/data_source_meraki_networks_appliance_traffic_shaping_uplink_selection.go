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

// DATA SOURCE NORMAL
import (
	"context"
	"log"

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksApplianceTrafficShapingUplinkSelectionDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceTrafficShapingUplinkSelectionDataSource{}
)

func NewNetworksApplianceTrafficShapingUplinkSelectionDataSource() datasource.DataSource {
	return &NetworksApplianceTrafficShapingUplinkSelectionDataSource{}
}

type NetworksApplianceTrafficShapingUplinkSelectionDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceTrafficShapingUplinkSelectionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceTrafficShapingUplinkSelectionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_traffic_shaping_uplink_selection"
}

func (d *NetworksApplianceTrafficShapingUplinkSelectionDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"active_active_auto_vpn_enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether active-active AutoVPN is enabled`,
						Computed:            true,
					},
					"default_uplink": schema.StringAttribute{
						MarkdownDescription: `The default uplink. Must be one of: 'wan1' or 'wan2'`,
						Computed:            true,
					},
					"failover_and_failback": schema.SingleNestedAttribute{
						MarkdownDescription: `WAN failover and failback`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"immediate": schema.SingleNestedAttribute{
								MarkdownDescription: `Immediate WAN failover and failback`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"enabled": schema.BoolAttribute{
										MarkdownDescription: `Whether immediate WAN failover and failback is enabled`,
										Computed:            true,
									},
								},
							},
						},
					},
					"load_balancing_enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether load balancing is enabled`,
						Computed:            true,
					},
					"vpn_traffic_uplink_preferences": schema.SetNestedAttribute{
						MarkdownDescription: `Uplink preference rules for VPN traffic`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"fail_over_criterion": schema.StringAttribute{
									MarkdownDescription: `Fail over criterion for uplink preference rule. Must be one of: 'poorPerformance' or 'uplinkDown'`,
									Computed:            true,
								},
								"performance_class": schema.SingleNestedAttribute{
									MarkdownDescription: `Performance class setting for uplink preference rule`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"builtin_performance_class_name": schema.StringAttribute{
											MarkdownDescription: `Name of builtin performance class. Must be present when performanceClass type is 'builtin' and value must be one of: 'VoIP'`,
											Computed:            true,
										},
										"custom_performance_class_id": schema.StringAttribute{
											MarkdownDescription: `ID of created custom performance class, must be present when performanceClass type is "custom"`,
											Computed:            true,
										},
										"type": schema.StringAttribute{
											MarkdownDescription: `Type of this performance class. Must be one of: 'builtin' or 'custom'`,
											Computed:            true,
										},
									},
								},
								"preferred_uplink": schema.StringAttribute{
									MarkdownDescription: `Preferred uplink for uplink preference rule. Must be one of: 'wan1', 'wan2', 'bestForVoIP', 'loadBalancing' or 'defaultUplink'`,
									Computed:            true,
								},
								"traffic_filters": schema.SetNestedAttribute{
									MarkdownDescription: `Traffic filters`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"type": schema.StringAttribute{
												MarkdownDescription: `Traffic filter type. Must be one of: 'applicationCategory', 'application' or 'custom'`,
												Computed:            true,
											},
											"value": schema.SingleNestedAttribute{
												MarkdownDescription: `Value of traffic filter`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"destination": schema.SingleNestedAttribute{
														MarkdownDescription: `Destination of 'custom' type traffic filter`,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"cidr": schema.StringAttribute{
																MarkdownDescription: `CIDR format address (e.g."192.168.10.1", which is the same as "192.168.10.1/32"), or "any". Cannot be used in combination with the "vlan" or "fqdn" property`,
																Computed:            true,
															},
															"fqdn": schema.StringAttribute{
																MarkdownDescription: `FQDN format address. Cannot be used in combination with the "cidr" or "fqdn" property and is currently only available in the "destination" object of the "vpnTrafficUplinkPreference" object. E.g.: "www.google.com"`,
																Computed:            true,
															},
															"host": schema.Int64Attribute{
																MarkdownDescription: `Host ID in the VLAN. Should not exceed the VLAN subnet capacity. Must be used along with the "vlan" property and is currently only available under a template network.`,
																Computed:            true,
															},
															"network": schema.StringAttribute{
																MarkdownDescription: `Meraki network ID. Currently only available under a template network, and the value should be ID of either same template network, or another template network currently. E.g.: "L_12345678".`,
																Computed:            true,
															},
															"port": schema.StringAttribute{
																MarkdownDescription: `E.g.: "any", "0" (also means "any"), "8080", "1-1024"`,
																Computed:            true,
															},
															"vlan": schema.Int64Attribute{
																MarkdownDescription: `VLAN ID of the configured VLAN in the Meraki network. Cannot be used in combination with the "cidr" or "fqdn" property and is currently only available under a template network.`,
																Computed:            true,
															},
														},
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `ID of 'applicationCategory' or 'application' type traffic filter`,
														Computed:            true,
													},
													"protocol": schema.StringAttribute{
														MarkdownDescription: `Protocol of 'custom' type traffic filter. Must be one of: 'tcp', 'udp', 'icmp', 'icmp6' or 'any'`,
														Computed:            true,
													},
													"source": schema.SingleNestedAttribute{
														MarkdownDescription: `Source of 'custom' type traffic filter`,
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
															"network": schema.StringAttribute{
																MarkdownDescription: `Meraki network ID. Currently only available under a template network, and the value should be ID of either same template network, or another template network currently. E.g.: "L_12345678".`,
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
					"wan_traffic_uplink_preferences": schema.SetNestedAttribute{
						MarkdownDescription: `Uplink preference rules for WAN traffic`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"preferred_uplink": schema.StringAttribute{
									MarkdownDescription: `Preferred uplink for uplink preference rule. Must be one of: 'wan1' or 'wan2'`,
									Computed:            true,
								},
								"traffic_filters": schema.SetNestedAttribute{
									MarkdownDescription: `Traffic filters`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"type": schema.StringAttribute{
												MarkdownDescription: `Traffic filter type. Must be "custom"`,
												Computed:            true,
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
														MarkdownDescription: `Protocol of 'custom' type traffic filter. Must be one of: 'tcp', 'udp', 'icmp6' or 'any'`,
														Computed:            true,
													},
													"source": schema.SingleNestedAttribute{
														MarkdownDescription: `Source of 'custom' type traffic filter`,
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
		},
	}
}

func (d *NetworksApplianceTrafficShapingUplinkSelectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceTrafficShapingUplinkSelection NetworksApplianceTrafficShapingUplinkSelection
	diags := req.Config.Get(ctx, &networksApplianceTrafficShapingUplinkSelection)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceTrafficShapingUplinkSelection")
		vvNetworkID := networksApplianceTrafficShapingUplinkSelection.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceTrafficShapingUplinkSelection(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceTrafficShapingUplinkSelection",
				err.Error(),
			)
			return
		}

		networksApplianceTrafficShapingUplinkSelection = ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionItemToBody(networksApplianceTrafficShapingUplinkSelection, response1)
		diags = resp.State.Set(ctx, &networksApplianceTrafficShapingUplinkSelection)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceTrafficShapingUplinkSelection struct {
	NetworkID types.String                                                       `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelection `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelection struct {
	ActiveActiveAutoVpnEnabled  types.Bool                                                                                      `tfsdk:"active_active_auto_vpn_enabled"`
	DefaultUplink               types.String                                                                                    `tfsdk:"default_uplink"`
	FailoverAndFailback         *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailback           `tfsdk:"failover_and_failback"`
	LoadBalancingEnabled        types.Bool                                                                                      `tfsdk:"load_balancing_enabled"`
	VpnTrafficUplinkPreferences *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences `tfsdk:"vpn_traffic_uplink_preferences"`
	WanTrafficUplinkPreferences *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences `tfsdk:"wan_traffic_uplink_preferences"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailback struct {
	Immediate *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate `tfsdk:"immediate"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences struct {
	FailOverCriterion types.String                                                                                                  `tfsdk:"fail_over_criterion"`
	PerformanceClass  *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesPerformanceClass `tfsdk:"performance_class"`
	PreferredUplink   types.String                                                                                                  `tfsdk:"preferred_uplink"`
	TrafficFilters    *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters `tfsdk:"traffic_filters"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesPerformanceClass struct {
	BuiltinPerformanceClassName types.String `tfsdk:"builtin_performance_class_name"`
	CustomPerformanceClassID    types.String `tfsdk:"custom_performance_class_id"`
	Type                        types.String `tfsdk:"type"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters struct {
	Type  types.String                                                                                                     `tfsdk:"type"`
	Value *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValue `tfsdk:"value"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValue struct {
	Destination *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueDestination `tfsdk:"destination"`
	ID          types.String                                                                                                                `tfsdk:"id"`
	Protocol    types.String                                                                                                                `tfsdk:"protocol"`
	Source      *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueSource      `tfsdk:"source"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueDestination struct {
	Cidr    types.String `tfsdk:"cidr"`
	Fqdn    types.String `tfsdk:"fqdn"`
	Host    types.Int64  `tfsdk:"host"`
	Network types.String `tfsdk:"network"`
	Port    types.String `tfsdk:"port"`
	VLAN    types.Int64  `tfsdk:"vlan"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueSource struct {
	Cidr    types.String `tfsdk:"cidr"`
	Host    types.Int64  `tfsdk:"host"`
	Network types.String `tfsdk:"network"`
	Port    types.String `tfsdk:"port"`
	VLAN    types.Int64  `tfsdk:"vlan"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences struct {
	PreferredUplink types.String                                                                                                  `tfsdk:"preferred_uplink"`
	TrafficFilters  *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters `tfsdk:"traffic_filters"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters struct {
	Type  types.String                                                                                                     `tfsdk:"type"`
	Value *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValue `tfsdk:"value"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValue struct {
	Destination *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestination `tfsdk:"destination"`
	Protocol    types.String                                                                                                                `tfsdk:"protocol"`
	Source      *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueSource      `tfsdk:"source"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestination struct {
	Applications *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications `tfsdk:"applications"`
	Cidr         types.String                                                                                                                              `tfsdk:"cidr"`
	Port         types.String                                                                                                                              `tfsdk:"port"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueSource struct {
	Cidr types.String `tfsdk:"cidr"`
	Host types.Int64  `tfsdk:"host"`
	Port types.String `tfsdk:"port"`
	VLAN types.Int64  `tfsdk:"vlan"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionItemToBody(state NetworksApplianceTrafficShapingUplinkSelection, response *merakigosdk.ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelection) NetworksApplianceTrafficShapingUplinkSelection {
	itemState := ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelection{
		ActiveActiveAutoVpnEnabled: func() types.Bool {
			if response.ActiveActiveAutoVpnEnabled != nil {
				return types.BoolValue(*response.ActiveActiveAutoVpnEnabled)
			}
			return types.Bool{}
		}(),
		DefaultUplink: types.StringValue(response.DefaultUplink),
		FailoverAndFailback: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailback {
			if response.FailoverAndFailback != nil {
				return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailback{
					Immediate: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate {
						if response.FailoverAndFailback.Immediate != nil {
							return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate{
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
		VpnTrafficUplinkPreferences: func() *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences {
			if response.VpnTrafficUplinkPreferences != nil {
				result := make([]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences, len(*response.VpnTrafficUplinkPreferences))
				for i, vpnTrafficUplinkPreferences := range *response.VpnTrafficUplinkPreferences {
					result[i] = ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences{
						FailOverCriterion: types.StringValue(vpnTrafficUplinkPreferences.FailOverCriterion),
						PerformanceClass: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesPerformanceClass {
							if vpnTrafficUplinkPreferences.PerformanceClass != nil {
								return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesPerformanceClass{
									BuiltinPerformanceClassName: types.StringValue(vpnTrafficUplinkPreferences.PerformanceClass.BuiltinPerformanceClassName),
									CustomPerformanceClassID:    types.StringValue(vpnTrafficUplinkPreferences.PerformanceClass.CustomPerformanceClassID),
									Type:                        types.StringValue(vpnTrafficUplinkPreferences.PerformanceClass.Type),
								}
							}
							return nil
						}(),
						PreferredUplink: types.StringValue(vpnTrafficUplinkPreferences.PreferredUplink),
						TrafficFilters: func() *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters {
							if vpnTrafficUplinkPreferences.TrafficFilters != nil {
								result := make([]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters, len(*vpnTrafficUplinkPreferences.TrafficFilters))
								for i, trafficFilters := range *vpnTrafficUplinkPreferences.TrafficFilters {
									result[i] = ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters{
										Type: types.StringValue(trafficFilters.Type),
										Value: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValue {
											if trafficFilters.Value != nil {
												return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValue{
													Destination: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueDestination {
														if trafficFilters.Value.Destination != nil {
															return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueDestination{
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
													Source: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueSource {
														if trafficFilters.Value.Source != nil {
															return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFiltersValueSource{
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
		WanTrafficUplinkPreferences: func() *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences {
			if response.WanTrafficUplinkPreferences != nil {
				result := make([]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences, len(*response.WanTrafficUplinkPreferences))
				for i, wanTrafficUplinkPreferences := range *response.WanTrafficUplinkPreferences {
					result[i] = ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences{
						PreferredUplink: types.StringValue(wanTrafficUplinkPreferences.PreferredUplink),
						TrafficFilters: func() *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters {
							if wanTrafficUplinkPreferences.TrafficFilters != nil {
								result := make([]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters, len(*wanTrafficUplinkPreferences.TrafficFilters))
								for i, trafficFilters := range *wanTrafficUplinkPreferences.TrafficFilters {
									result[i] = ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters{
										Type: types.StringValue(trafficFilters.Type),
										Value: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValue {
											if trafficFilters.Value != nil {
												return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValue{
													Destination: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestination {
														if trafficFilters.Value.Destination != nil {
															return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestination{
																Applications: func() *[]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications {
																	if trafficFilters.Value.Destination.Applications != nil {
																		result := make([]ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications, len(*trafficFilters.Value.Destination.Applications))
																		for i, applications := range *trafficFilters.Value.Destination.Applications {
																			result[i] = ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueDestinationApplications{
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
													Source: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueSource {
														if trafficFilters.Value.Source != nil {
															return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFiltersValueSource{
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
