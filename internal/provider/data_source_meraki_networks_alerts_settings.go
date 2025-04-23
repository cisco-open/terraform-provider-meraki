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

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksAlertsSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksAlertsSettingsDataSource{}
)

func NewNetworksAlertsSettingsDataSource() datasource.DataSource {
	return &NetworksAlertsSettingsDataSource{}
}

type NetworksAlertsSettingsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksAlertsSettingsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksAlertsSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_alerts_settings"
}

func (d *NetworksAlertsSettingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"alerts": schema.SetNestedAttribute{
						MarkdownDescription: `Alert-specific configuration for each type. Only alerts that pertain to the network can be updated.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"alert_destinations": schema.SingleNestedAttribute{
									MarkdownDescription: `A hash of destinations for this specific alert`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"all_admins": schema.BoolAttribute{
											MarkdownDescription: `If true, then all network admins will receive emails for this alert`,
											Computed:            true,
										},
										"emails": schema.ListAttribute{
											MarkdownDescription: `A list of emails that will receive information about the alert`,
											Computed:            true,
											ElementType:         types.StringType,
										},
										"http_server_ids": schema.ListAttribute{
											MarkdownDescription: `A list of HTTP server IDs to send a Webhook to for this alert`,
											Computed:            true,
											ElementType:         types.StringType,
										},
										"sms_numbers": schema.ListAttribute{
											MarkdownDescription: `A list of phone numbers that will receive text messages about the alert. Only available for sensors status alerts.`,
											Computed:            true,
											ElementType:         types.StringType,
										},
										"snmp": schema.BoolAttribute{
											MarkdownDescription: `If true, then an SNMP trap will be sent for this alert if there is an SNMP trap server configured for this network`,
											Computed:            true,
										},
									},
								},
								"enabled": schema.BoolAttribute{
									MarkdownDescription: `A boolean depicting if the alert is turned on or off`,
									Computed:            true,
								},
								"filters": schema.SingleNestedAttribute{
									MarkdownDescription: `A hash of specific configuration data for the alert. Only filters specific to the alert will be updated.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"conditions": schema.SetNestedAttribute{
											MarkdownDescription: `Conditions`,
											Computed:            true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{

													"direction": schema.StringAttribute{
														MarkdownDescription: `Direction`,
														Computed:            true,
													},
													"duration": schema.Int64Attribute{
														MarkdownDescription: `Duration`,
														Computed:            true,
													},
													"threshold": schema.Float64Attribute{
														MarkdownDescription: `Threshold`,
														Computed:            true,
													},
													"type": schema.StringAttribute{
														MarkdownDescription: `Type of condition`,
														Computed:            true,
													},
													"unit": schema.StringAttribute{
														MarkdownDescription: `Unit`,
														Computed:            true,
													},
												},
											},
										},
										"failure_type": schema.StringAttribute{
											MarkdownDescription: `Failure Type`,
											Computed:            true,
										},
										"lookback_window": schema.Int64Attribute{
											MarkdownDescription: `Loopback Window (in sec)`,
											Computed:            true,
										},
										"min_duration": schema.Int64Attribute{
											MarkdownDescription: `Min Duration`,
											Computed:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: `Name`,
											Computed:            true,
										},
										"period": schema.Int64Attribute{
											MarkdownDescription: `Period`,
											Computed:            true,
										},
										"priority": schema.StringAttribute{
											MarkdownDescription: `Priority`,
											Computed:            true,
										},
										"regex": schema.StringAttribute{
											MarkdownDescription: `Regex`,
											Computed:            true,
										},
										"selector": schema.StringAttribute{
											MarkdownDescription: `Selector`,
											Computed:            true,
										},
										"serials": schema.ListAttribute{
											MarkdownDescription: `Serials`,
											Computed:            true,
											ElementType:         types.StringType,
										},
										"ssid_num": schema.Int64Attribute{
											MarkdownDescription: `SSID Number`,
											Computed:            true,
										},
										"tag": schema.StringAttribute{
											MarkdownDescription: `Tag`,
											Computed:            true,
										},
										"threshold": schema.Int64Attribute{
											MarkdownDescription: `Threshold`,
											Computed:            true,
										},
										"timeout": schema.Int64Attribute{
											MarkdownDescription: `Timeout`,
											Computed:            true,
										},
									},
								},
								"type": schema.StringAttribute{
									MarkdownDescription: `The type of alert`,
									Computed:            true,
								},
							},
						},
					},
					"default_destinations": schema.SingleNestedAttribute{
						MarkdownDescription: `The network-wide destinations for all alerts on the network.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"all_admins": schema.BoolAttribute{
								MarkdownDescription: `If true, then all network admins will receive emails.`,
								Computed:            true,
							},
							"emails": schema.ListAttribute{
								MarkdownDescription: `A list of emails that will receive the alert(s).`,
								Computed:            true,
								ElementType:         types.StringType,
							},
							"http_server_ids": schema.ListAttribute{
								MarkdownDescription: `A list of HTTP server IDs to send a Webhook to`,
								Computed:            true,
								ElementType:         types.StringType,
							},
							"snmp": schema.BoolAttribute{
								MarkdownDescription: `If true, then an SNMP trap will be sent if there is an SNMP trap server configured for this network.`,
								Computed:            true,
							},
						},
					},
					"muting": schema.SingleNestedAttribute{
						MarkdownDescription: `Mute alerts under certain conditions`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"by_port_schedules": schema.SingleNestedAttribute{
								MarkdownDescription: `Mute wireless unreachable alerts based on switch port schedules`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"enabled": schema.BoolAttribute{
										MarkdownDescription: `If true, then wireless unreachable alerts will be muted when caused by a port schedule`,
										Computed:            true,
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

func (d *NetworksAlertsSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksAlertsSettings NetworksAlertsSettings
	diags := req.Config.Get(ctx, &networksAlertsSettings)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkAlertsSettings")
		vvNetworkID := networksAlertsSettings.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkAlertsSettings(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkAlertsSettings",
				err.Error(),
			)
			return
		}

		networksAlertsSettings = ResponseNetworksGetNetworkAlertsSettingsItemToBody(networksAlertsSettings, response1)
		diags = resp.State.Set(ctx, &networksAlertsSettings)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksAlertsSettings struct {
	NetworkID types.String                              `tfsdk:"network_id"`
	Item      *ResponseNetworksGetNetworkAlertsSettings `tfsdk:"item"`
}

type ResponseNetworksGetNetworkAlertsSettings struct {
	Alerts              *[]ResponseNetworksGetNetworkAlertsSettingsAlerts            `tfsdk:"alerts"`
	DefaultDestinations *ResponseNetworksGetNetworkAlertsSettingsDefaultDestinations `tfsdk:"default_destinations"`
	Muting              *ResponseNetworksGetNetworkAlertsSettingsMuting              `tfsdk:"muting"`
}

type ResponseNetworksGetNetworkAlertsSettingsAlerts struct {
	AlertDestinations *ResponseNetworksGetNetworkAlertsSettingsAlertsAlertDestinations `tfsdk:"alert_destinations"`
	Enabled           types.Bool                                                       `tfsdk:"enabled"`
	Filters           *ResponseNetworksGetNetworkAlertsSettingsAlertsFilters           `tfsdk:"filters"`
	Type              types.String                                                     `tfsdk:"type"`
}

type ResponseNetworksGetNetworkAlertsSettingsAlertsAlertDestinations struct {
	AllAdmins     types.Bool `tfsdk:"all_admins"`
	Emails        types.List `tfsdk:"emails"`
	HTTPServerIDs types.List `tfsdk:"http_server_ids"`
	SmsNumbers    types.List `tfsdk:"sms_numbers"`
	SNMP          types.Bool `tfsdk:"snmp"`
}

type ResponseNetworksGetNetworkAlertsSettingsAlertsFilters struct {
	Conditions     *[]ResponseNetworksGetNetworkAlertsSettingsAlertsFiltersConditions `tfsdk:"conditions"`
	FailureType    types.String                                                       `tfsdk:"failure_type"`
	LookbackWindow types.Int64                                                        `tfsdk:"lookback_window"`
	MinDuration    types.Int64                                                        `tfsdk:"min_duration"`
	Name           types.String                                                       `tfsdk:"name"`
	Period         types.Int64                                                        `tfsdk:"period"`
	Priority       types.String                                                       `tfsdk:"priority"`
	Regex          types.String                                                       `tfsdk:"regex"`
	Selector       types.String                                                       `tfsdk:"selector"`
	Serials        types.List                                                         `tfsdk:"serials"`
	SSIDNum        types.Int64                                                        `tfsdk:"ssid_num"`
	Tag            types.String                                                       `tfsdk:"tag"`
	Threshold      types.Int64                                                        `tfsdk:"threshold"`
	Timeout        types.Int64                                                        `tfsdk:"timeout"`
}

type ResponseNetworksGetNetworkAlertsSettingsAlertsFiltersConditions struct {
	Direction types.String  `tfsdk:"direction"`
	Duration  types.Int64   `tfsdk:"duration"`
	Threshold types.Float64 `tfsdk:"threshold"`
	Type      types.String  `tfsdk:"type"`
	Unit      types.String  `tfsdk:"unit"`
}

type ResponseNetworksGetNetworkAlertsSettingsDefaultDestinations struct {
	AllAdmins     types.Bool `tfsdk:"all_admins"`
	Emails        types.List `tfsdk:"emails"`
	HTTPServerIDs types.List `tfsdk:"http_server_ids"`
	SNMP          types.Bool `tfsdk:"snmp"`
}

type ResponseNetworksGetNetworkAlertsSettingsMuting struct {
	ByPortSchedules *ResponseNetworksGetNetworkAlertsSettingsMutingByPortSchedules `tfsdk:"by_port_schedules"`
}

type ResponseNetworksGetNetworkAlertsSettingsMutingByPortSchedules struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// ToBody
func ResponseNetworksGetNetworkAlertsSettingsItemToBody(state NetworksAlertsSettings, response *merakigosdk.ResponseNetworksGetNetworkAlertsSettings) NetworksAlertsSettings {
	itemState := ResponseNetworksGetNetworkAlertsSettings{
		Alerts: func() *[]ResponseNetworksGetNetworkAlertsSettingsAlerts {
			if response.Alerts != nil {
				result := make([]ResponseNetworksGetNetworkAlertsSettingsAlerts, len(*response.Alerts))
				for i, alerts := range *response.Alerts {
					result[i] = ResponseNetworksGetNetworkAlertsSettingsAlerts{
						AlertDestinations: func() *ResponseNetworksGetNetworkAlertsSettingsAlertsAlertDestinations {
							if alerts.AlertDestinations != nil {
								return &ResponseNetworksGetNetworkAlertsSettingsAlertsAlertDestinations{
									AllAdmins: func() types.Bool {
										if alerts.AlertDestinations.AllAdmins != nil {
											return types.BoolValue(*alerts.AlertDestinations.AllAdmins)
										}
										return types.Bool{}
									}(),
									Emails:        StringSliceToList(alerts.AlertDestinations.Emails),
									HTTPServerIDs: StringSliceToList(alerts.AlertDestinations.HTTPServerIDs),
									SmsNumbers:    StringSliceToList(alerts.AlertDestinations.SmsNumbers),
									SNMP: func() types.Bool {
										if alerts.AlertDestinations.SNMP != nil {
											return types.BoolValue(*alerts.AlertDestinations.SNMP)
										}
										return types.Bool{}
									}(),
								}
							}
							return nil
						}(),
						Enabled: func() types.Bool {
							if alerts.Enabled != nil {
								return types.BoolValue(*alerts.Enabled)
							}
							return types.Bool{}
						}(),
						Filters: func() *ResponseNetworksGetNetworkAlertsSettingsAlertsFilters {
							if alerts.Filters != nil {
								return &ResponseNetworksGetNetworkAlertsSettingsAlertsFilters{
									Conditions: func() *[]ResponseNetworksGetNetworkAlertsSettingsAlertsFiltersConditions {
										if alerts.Filters.Conditions != nil {
											result := make([]ResponseNetworksGetNetworkAlertsSettingsAlertsFiltersConditions, len(*alerts.Filters.Conditions))
											for i, conditions := range *alerts.Filters.Conditions {
												result[i] = ResponseNetworksGetNetworkAlertsSettingsAlertsFiltersConditions{
													Direction: types.StringValue(conditions.Direction),
													Duration: func() types.Int64 {
														if conditions.Duration != nil {
															return types.Int64Value(int64(*conditions.Duration))
														}
														return types.Int64{}
													}(),
													Threshold: func() types.Float64 {
														if conditions.Threshold != nil {
															return types.Float64Value(float64(*conditions.Threshold))
														}
														return types.Float64{}
													}(),
													Type: types.StringValue(conditions.Type),
													Unit: types.StringValue(conditions.Unit),
												}
											}
											return &result
										}
										return nil
									}(),
									FailureType: types.StringValue(alerts.Filters.FailureType),
									LookbackWindow: func() types.Int64 {
										if alerts.Filters.LookbackWindow != nil {
											return types.Int64Value(int64(*alerts.Filters.LookbackWindow))
										}
										return types.Int64{}
									}(),
									MinDuration: func() types.Int64 {
										if alerts.Filters.MinDuration != nil {
											return types.Int64Value(int64(*alerts.Filters.MinDuration))
										}
										return types.Int64{}
									}(),
									Name: types.StringValue(alerts.Filters.Name),
									Period: func() types.Int64 {
										if alerts.Filters.Period != nil {
											return types.Int64Value(int64(*alerts.Filters.Period))
										}
										return types.Int64{}
									}(),
									Priority: types.StringValue(alerts.Filters.Priority),
									Regex:    types.StringValue(alerts.Filters.Regex),
									Selector: types.StringValue(alerts.Filters.Selector),
									Serials:  StringSliceToList(alerts.Filters.Serials),
									SSIDNum: func() types.Int64 {
										if alerts.Filters.SSIDNum != nil {
											return types.Int64Value(int64(*alerts.Filters.SSIDNum))
										}
										return types.Int64{}
									}(),
									Tag: types.StringValue(alerts.Filters.Tag),
									Threshold: func() types.Int64 {
										if alerts.Filters.Threshold != nil {
											return types.Int64Value(int64(*alerts.Filters.Threshold))
										}
										return types.Int64{}
									}(),
									Timeout: func() types.Int64 {
										if alerts.Filters.Timeout != nil {
											return types.Int64Value(int64(*alerts.Filters.Timeout))
										}
										return types.Int64{}
									}(),
								}
							}
							return nil
						}(),
						Type: types.StringValue(alerts.Type),
					}
				}
				return &result
			}
			return nil
		}(),
		DefaultDestinations: func() *ResponseNetworksGetNetworkAlertsSettingsDefaultDestinations {
			if response.DefaultDestinations != nil {
				return &ResponseNetworksGetNetworkAlertsSettingsDefaultDestinations{
					AllAdmins: func() types.Bool {
						if response.DefaultDestinations.AllAdmins != nil {
							return types.BoolValue(*response.DefaultDestinations.AllAdmins)
						}
						return types.Bool{}
					}(),
					Emails:        StringSliceToList(response.DefaultDestinations.Emails),
					HTTPServerIDs: StringSliceToList(response.DefaultDestinations.HTTPServerIDs),
					SNMP: func() types.Bool {
						if response.DefaultDestinations.SNMP != nil {
							return types.BoolValue(*response.DefaultDestinations.SNMP)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
		}(),
		Muting: func() *ResponseNetworksGetNetworkAlertsSettingsMuting {
			if response.Muting != nil {
				return &ResponseNetworksGetNetworkAlertsSettingsMuting{
					ByPortSchedules: func() *ResponseNetworksGetNetworkAlertsSettingsMutingByPortSchedules {
						if response.Muting.ByPortSchedules != nil {
							return &ResponseNetworksGetNetworkAlertsSettingsMutingByPortSchedules{
								Enabled: func() types.Bool {
									if response.Muting.ByPortSchedules.Enabled != nil {
										return types.BoolValue(*response.Muting.ByPortSchedules.Enabled)
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
	}
	state.Item = &itemState
	return state
}
