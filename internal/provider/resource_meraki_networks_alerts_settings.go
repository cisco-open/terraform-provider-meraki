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

	"log"

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksAlertsSettingsResource{}
	_ resource.ResourceWithConfigure = &NetworksAlertsSettingsResource{}
)

func NewNetworksAlertsSettingsResource() resource.Resource {
	return &NetworksAlertsSettingsResource{}
}

type NetworksAlertsSettingsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksAlertsSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksAlertsSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_alerts_settings"
}

func (r *NetworksAlertsSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"alerts_response": schema.SetNestedAttribute{
				MarkdownDescription: `Alert-specific configuration for each type. Only alerts that pertain to the network can be updated.`,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"alert_destinations": schema.SingleNestedAttribute{
							MarkdownDescription: `A hash of destinations for this specific alert`,
							Computed:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"all_admins": schema.BoolAttribute{
									MarkdownDescription: `If true, then all network admins will receive emails for this alert`,
									Computed:            true,
									PlanModifiers: []planmodifier.Bool{
										boolplanmodifier.UseStateForUnknown(),
									},
								},
								"emails": schema.SetAttribute{
									MarkdownDescription: `A list of emails that will receive information about the alert`,
									Computed:            true,
									PlanModifiers: []planmodifier.Set{
										setplanmodifier.UseStateForUnknown(),
									},

									ElementType: types.StringType,
								},
								"http_server_ids": schema.SetAttribute{
									MarkdownDescription: `A list of HTTP server IDs to send a Webhook to for this alert`,
									Computed:            true,
									PlanModifiers: []planmodifier.Set{
										setplanmodifier.UseStateForUnknown(),
									},
									Default:     setdefault.StaticValue(types.SetNull(types.StringType)),
									ElementType: types.StringType,
								},
								"sms_numbers": schema.SetAttribute{
									MarkdownDescription: `A list of phone numbers that will receive text messages about the alert. Only available for sensors status alerts.`,
									Computed:            true,
									PlanModifiers: []planmodifier.Set{
										setplanmodifier.UseStateForUnknown(),
									},

									ElementType: types.StringType,
								},
								"snmp": schema.BoolAttribute{
									MarkdownDescription: `If true, then an SNMP trap will be sent for this alert if there is an SNMP trap server configured for this network`,
									Computed:            true,
									PlanModifiers: []planmodifier.Bool{
										boolplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: `A boolean depicting if the alert is turned on or off`,
							Computed:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
						"filters": schema.SingleNestedAttribute{
							MarkdownDescription: `A hash of specific configuration data for the alert. Only filters specific to the alert will be updated.`,
							Computed:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"conditions": schema.SetNestedAttribute{
									MarkdownDescription: `Conditions`,
									Computed:            true,
									PlanModifiers: []planmodifier.Set{
										setplanmodifier.UseStateForUnknown(),
									},
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"direction": schema.StringAttribute{
												MarkdownDescription: `Direction
                                                    Allowed values: [+,-]`,
												Computed: true,
												Optional: true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
												Validators: []validator.String{
													stringvalidator.OneOf(
														"+",
														"-",
													),
												},
											},
											"duration": schema.Int64Attribute{
												MarkdownDescription: `Duration`,
												Computed:            true,
												PlanModifiers: []planmodifier.Int64{
													int64planmodifier.UseStateForUnknown(),
												},
											},
											"threshold": schema.Float64Attribute{
												MarkdownDescription: `Threshold`,
												Computed:            true,
												PlanModifiers: []planmodifier.Float64{
													float64planmodifier.UseStateForUnknown(),
												},
											},
											"type": schema.StringAttribute{
												MarkdownDescription: `Type of condition`,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
											"unit": schema.StringAttribute{
												MarkdownDescription: `Unit`,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
										},
									},
								},
								"failure_type": schema.StringAttribute{
									MarkdownDescription: `Failure Type`,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"lookback_window": schema.Int64Attribute{
									MarkdownDescription: `Loopback Window (in sec)`,
									Computed:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"min_duration": schema.Int64Attribute{
									MarkdownDescription: `Min Duration`,
									Computed:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Name`,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"period": schema.Int64Attribute{
									MarkdownDescription: `Period`,
									Computed:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"priority": schema.StringAttribute{
									MarkdownDescription: `Priority`,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"regex": schema.StringAttribute{
									MarkdownDescription: `Regex`,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"selector": schema.StringAttribute{
									MarkdownDescription: `Selector`,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"serials": schema.SetAttribute{
									MarkdownDescription: `Serials`,
									Computed:            true,
									PlanModifiers: []planmodifier.Set{
										setplanmodifier.UseStateForUnknown(),
									},

									ElementType: types.StringType,
								},
								"ssid_num": schema.Int64Attribute{
									MarkdownDescription: `SSID Number`,
									Computed:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"tag": schema.StringAttribute{
									MarkdownDescription: `Tag`,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"threshold": schema.Int64Attribute{
									MarkdownDescription: `Threshold`,
									Computed:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"timeout": schema.Int64Attribute{
									MarkdownDescription: `Timeout`,
									Computed:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `The type of alert`,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"alerts": schema.SetNestedAttribute{
				MarkdownDescription: `Alert-specific configuration for each type. Only alerts that pertain to the network can be updated.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"alert_destinations": schema.SingleNestedAttribute{
							MarkdownDescription: `A hash of destinations for this specific alert`,
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"all_admins": schema.BoolAttribute{
									MarkdownDescription: `If true, then all network admins will receive emails for this alert`,
									Optional:            true,
									PlanModifiers: []planmodifier.Bool{
										boolplanmodifier.UseStateForUnknown(),
									},
								},
								"emails": schema.SetAttribute{
									MarkdownDescription: `A list of emails that will receive information about the alert`,
									Optional:            true,
									PlanModifiers: []planmodifier.Set{
										setplanmodifier.UseStateForUnknown(),
									},

									ElementType: types.StringType,
								},
								"http_server_ids": schema.SetAttribute{
									MarkdownDescription: `A list of HTTP server IDs to send a Webhook to for this alert`,
									Optional:            true,
									PlanModifiers: []planmodifier.Set{
										setplanmodifier.UseStateForUnknown(),
									},

									ElementType: types.StringType,
								},
								"sms_numbers": schema.SetAttribute{
									MarkdownDescription: `A list of phone numbers that will receive text messages about the alert. Only available for sensors status alerts.`,
									Optional:            true,
									PlanModifiers: []planmodifier.Set{
										setplanmodifier.UseStateForUnknown(),
									},

									ElementType: types.StringType,
								},
								"snmp": schema.BoolAttribute{
									MarkdownDescription: `If true, then an SNMP trap will be sent for this alert if there is an SNMP trap server configured for this network`,
									Optional:            true,
									PlanModifiers: []planmodifier.Bool{
										boolplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: `A boolean depicting if the alert is turned on or off`,
							Optional:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
						"filters": schema.SingleNestedAttribute{
							MarkdownDescription: `A hash of specific configuration data for the alert. Only filters specific to the alert will be updated.`,
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"conditions": schema.SetNestedAttribute{
									MarkdownDescription: `Conditions`,
									Optional:            true,
									PlanModifiers: []planmodifier.Set{
										setplanmodifier.UseStateForUnknown(),
									},
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"direction": schema.StringAttribute{
												MarkdownDescription: `Direction
                                                    Allowed values: [+,-]`,
												Optional: true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
												Validators: []validator.String{
													stringvalidator.OneOf(
														"+",
														"-",
													),
												},
											},
											"duration": schema.Int64Attribute{
												MarkdownDescription: `Duration`,
												Optional:            true,
												PlanModifiers: []planmodifier.Int64{
													int64planmodifier.UseStateForUnknown(),
												},
											},
											"threshold": schema.Float64Attribute{
												MarkdownDescription: `Threshold`,
												Optional:            true,
												PlanModifiers: []planmodifier.Float64{
													float64planmodifier.UseStateForUnknown(),
												},
											},
											"type": schema.StringAttribute{
												MarkdownDescription: `Type of condition`,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
											"unit": schema.StringAttribute{
												MarkdownDescription: `Unit`,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
										},
									},
								},
								"failure_type": schema.StringAttribute{
									MarkdownDescription: `Failure Type`,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"lookback_window": schema.Int64Attribute{
									MarkdownDescription: `Loopback Window (in sec)`,
									Optional:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"min_duration": schema.Int64Attribute{
									MarkdownDescription: `Min Duration`,
									Optional:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Name`,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"period": schema.Int64Attribute{
									MarkdownDescription: `Period`,
									Optional:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"priority": schema.StringAttribute{
									MarkdownDescription: `Priority`,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"regex": schema.StringAttribute{
									MarkdownDescription: `Regex`,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"selector": schema.StringAttribute{
									MarkdownDescription: `Selector`,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"serials": schema.SetAttribute{
									MarkdownDescription: `Serials`,
									Optional:            true,
									PlanModifiers: []planmodifier.Set{
										setplanmodifier.UseStateForUnknown(),
									},

									ElementType: types.StringType,
								},
								"ssid_num": schema.Int64Attribute{
									MarkdownDescription: `SSID Number`,
									Optional:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"tag": schema.StringAttribute{
									MarkdownDescription: `Tag`,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"threshold": schema.Int64Attribute{
									MarkdownDescription: `Threshold`,
									Optional:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"timeout": schema.Int64Attribute{
									MarkdownDescription: `Timeout`,
									Optional:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `The type of alert`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"default_destinations": schema.SingleNestedAttribute{
				MarkdownDescription: `The network-wide destinations for all alerts on the network.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"all_admins": schema.BoolAttribute{
						MarkdownDescription: `If true, then all network admins will receive emails.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"emails": schema.SetAttribute{
						MarkdownDescription: `A list of emails that will receive the alert(s).`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						Default:     setdefault.StaticValue(types.SetNull(types.StringType)),
						ElementType: types.StringType,
					},
					"http_server_ids": schema.SetAttribute{
						MarkdownDescription: `A list of HTTP server IDs to send a Webhook to`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						Default:     setdefault.StaticValue(types.SetNull(types.StringType)),
						ElementType: types.StringType,
					},
					"snmp": schema.BoolAttribute{
						MarkdownDescription: `If true, then an SNMP trap will be sent if there is an SNMP trap server configured for this network.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Default: booldefault.StaticBool(false),
					},
				},
			},
			"muting": schema.SingleNestedAttribute{
				MarkdownDescription: `Mute alerts under certain conditions`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"by_port_schedules": schema.SingleNestedAttribute{
						MarkdownDescription: `Mute wireless unreachable alerts based on switch port schedules`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `If true, then wireless unreachable alerts will be muted when caused by a port schedule`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
								Default: booldefault.StaticBool(false),
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

func (r *NetworksAlertsSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksAlertsSettingsRs

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

	if vvNetworkID != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Networks.GetNetworkAlertsSettings(vvNetworkID)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksAlertsSettings  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksAlertsSettings only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkAlertsSettings(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkAlertsSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkAlertsSettings",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Networks.GetNetworkAlertsSettings(vvNetworkID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkAlertsSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkAlertsSettings",
			err.Error(),
		)
		return
	}

	data = ResponseNetworksGetNetworkAlertsSettingsItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *NetworksAlertsSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksAlertsSettingsRs

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
	responseGet, restyRespGet, err := r.client.Networks.GetNetworkAlertsSettings(vvNetworkID)
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
				"Failure when executing GetNetworkAlertsSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkAlertsSettings",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseNetworksGetNetworkAlertsSettingsItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksAlertsSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksAlertsSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksAlertsSettingsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkAlertsSettings(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkAlertsSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkAlertsSettings",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksAlertsSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksAlertsSettings", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksAlertsSettingsRs struct {
	NetworkID           types.String                                                   `tfsdk:"network_id"`
	Alerts              *[]ResponseNetworksGetNetworkAlertsSettingsAlertsRs            `tfsdk:"alerts"`
	AlertsResponse      *[]ResponseNetworksGetNetworkAlertsSettingsAlertsRs            `tfsdk:"alerts_response"`
	DefaultDestinations *ResponseNetworksGetNetworkAlertsSettingsDefaultDestinationsRs `tfsdk:"default_destinations"`
	Muting              *ResponseNetworksGetNetworkAlertsSettingsMutingRs              `tfsdk:"muting"`
}

type ResponseNetworksGetNetworkAlertsSettingsAlertsRs struct {
	AlertDestinations *ResponseNetworksGetNetworkAlertsSettingsAlertsAlertDestinationsRs `tfsdk:"alert_destinations"`
	Enabled           types.Bool                                                         `tfsdk:"enabled"`
	Filters           *ResponseNetworksGetNetworkAlertsSettingsAlertsFiltersRs           `tfsdk:"filters"`
	Type              types.String                                                       `tfsdk:"type"`
}

type ResponseNetworksGetNetworkAlertsSettingsAlertsAlertDestinationsRs struct {
	AllAdmins     types.Bool `tfsdk:"all_admins"`
	Emails        types.Set  `tfsdk:"emails"`
	HTTPServerIDs types.Set  `tfsdk:"http_server_ids"`
	SmsNumbers    types.Set  `tfsdk:"sms_numbers"`
	SNMP          types.Bool `tfsdk:"snmp"`
}

type ResponseNetworksGetNetworkAlertsSettingsAlertsFiltersRs struct {
	Conditions     *[]ResponseNetworksGetNetworkAlertsSettingsAlertsFiltersConditionsRs `tfsdk:"conditions"`
	FailureType    types.String                                                         `tfsdk:"failure_type"`
	LookbackWindow types.Int64                                                          `tfsdk:"lookback_window"`
	MinDuration    types.Int64                                                          `tfsdk:"min_duration"`
	Name           types.String                                                         `tfsdk:"name"`
	Period         types.Int64                                                          `tfsdk:"period"`
	Priority       types.String                                                         `tfsdk:"priority"`
	Regex          types.String                                                         `tfsdk:"regex"`
	Selector       types.String                                                         `tfsdk:"selector"`
	Serials        types.Set                                                            `tfsdk:"serials"`
	SSIDNum        types.Int64                                                          `tfsdk:"ssid_num"`
	Tag            types.String                                                         `tfsdk:"tag"`
	Threshold      types.Int64                                                          `tfsdk:"threshold"`
	Timeout        types.Int64                                                          `tfsdk:"timeout"`
}

type ResponseNetworksGetNetworkAlertsSettingsAlertsFiltersConditionsRs struct {
	Direction types.String  `tfsdk:"direction"`
	Duration  types.Int64   `tfsdk:"duration"`
	Threshold types.Float64 `tfsdk:"threshold"`
	Type      types.String  `tfsdk:"type"`
	Unit      types.String  `tfsdk:"unit"`
}

type ResponseNetworksGetNetworkAlertsSettingsDefaultDestinationsRs struct {
	AllAdmins     types.Bool `tfsdk:"all_admins"`
	Emails        types.Set  `tfsdk:"emails"`
	HTTPServerIDs types.Set  `tfsdk:"http_server_ids"`
	SNMP          types.Bool `tfsdk:"snmp"`
}

type ResponseNetworksGetNetworkAlertsSettingsMutingRs struct {
	ByPortSchedules *ResponseNetworksGetNetworkAlertsSettingsMutingByPortSchedulesRs `tfsdk:"by_port_schedules"`
}

type ResponseNetworksGetNetworkAlertsSettingsMutingByPortSchedulesRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// FromBody
func (r *NetworksAlertsSettingsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestNetworksUpdateNetworkAlertsSettings {
	var requestNetworksUpdateNetworkAlertsSettingsAlerts []merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsAlerts

	if r.Alerts != nil {
		for _, rItem1 := range *r.Alerts {
			var requestNetworksUpdateNetworkAlertsSettingsAlertsAlertDestinations *merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsAlertsAlertDestinations

			if rItem1.AlertDestinations != nil {
				allAdmins := func() *bool {
					if !rItem1.AlertDestinations.AllAdmins.IsUnknown() && !rItem1.AlertDestinations.AllAdmins.IsNull() {
						return rItem1.AlertDestinations.AllAdmins.ValueBoolPointer()
					}
					return nil
				}()

				var emails []string = nil
				rItem1.AlertDestinations.Emails.ElementsAs(ctx, &emails, false)

				var httpServerIDs []string = nil
				rItem1.AlertDestinations.HTTPServerIDs.ElementsAs(ctx, &httpServerIDs, false)

				var smsNumbers []string = nil
				rItem1.AlertDestinations.SmsNumbers.ElementsAs(ctx, &smsNumbers, false)
				snmp := func() *bool {
					if !rItem1.AlertDestinations.SNMP.IsUnknown() && !rItem1.AlertDestinations.SNMP.IsNull() {
						return rItem1.AlertDestinations.SNMP.ValueBoolPointer()
					}
					return nil
				}()
				requestNetworksUpdateNetworkAlertsSettingsAlertsAlertDestinations = &merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsAlertsAlertDestinations{
					AllAdmins:     allAdmins,
					Emails:        emails,
					HTTPServerIDs: httpServerIDs,
					SmsNumbers:    smsNumbers,
					SNMP:          snmp,
				}
				//[debug] Is Array: False
			}
			enabled := func() *bool {
				if !rItem1.Enabled.IsUnknown() && !rItem1.Enabled.IsNull() {
					return rItem1.Enabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestNetworksUpdateNetworkAlertsSettingsAlertsFilters *merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsAlertsFilters

			if rItem1.Filters != nil {

				log.Printf("[DEBUG] #TODO []RequestNetworksUpdateNetworkAlertsSettingsAlertsFiltersConditions")
				var requestNetworksUpdateNetworkAlertsSettingsAlertsFiltersConditions []merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsAlertsFiltersConditions

				if rItem1.Filters.Conditions != nil {
					for _, rItem2 := range *rItem1.Filters.Conditions {
						direction := rItem2.Direction.ValueString()
						duration := func() *int64 {
							if !rItem2.Duration.IsUnknown() && !rItem2.Duration.IsNull() {
								return rItem2.Duration.ValueInt64Pointer()
							}
							return nil
						}()
						threshold := func() *float64 {
							if !rItem2.Threshold.IsUnknown() && !rItem2.Threshold.IsNull() {
								return rItem2.Threshold.ValueFloat64Pointer()
							}
							return nil
						}()
						typeR := rItem2.Type.ValueString()
						unit := rItem2.Unit.ValueString()
						requestNetworksUpdateNetworkAlertsSettingsAlertsFiltersConditions = append(requestNetworksUpdateNetworkAlertsSettingsAlertsFiltersConditions, merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsAlertsFiltersConditions{
							Direction: direction,
							Duration:  int64ToIntPointer(duration),
							Threshold: threshold,
							Type:      typeR,
							Unit:      unit,
						})
						//[debug] Is Array: True
					}
				}
				failureType := rItem1.Filters.FailureType.ValueString()
				lookbackWindow := func() *int64 {
					if !rItem1.Filters.LookbackWindow.IsUnknown() && !rItem1.Filters.LookbackWindow.IsNull() {
						return rItem1.Filters.LookbackWindow.ValueInt64Pointer()
					}
					return nil
				}()
				minDuration := func() *int64 {
					if !rItem1.Filters.MinDuration.IsUnknown() && !rItem1.Filters.MinDuration.IsNull() {
						return rItem1.Filters.MinDuration.ValueInt64Pointer()
					}
					return nil
				}()
				name := rItem1.Filters.Name.ValueString()
				period := func() *int64 {
					if !rItem1.Filters.Period.IsUnknown() && !rItem1.Filters.Period.IsNull() {
						return rItem1.Filters.Period.ValueInt64Pointer()
					}
					return nil
				}()
				priority := rItem1.Filters.Priority.ValueString()
				regex := rItem1.Filters.Regex.ValueString()
				selector := rItem1.Filters.Selector.ValueString()

				var serials []string = nil
				rItem1.Filters.Serials.ElementsAs(ctx, &serials, false)
				ssidNum := func() *int64 {
					if !rItem1.Filters.SSIDNum.IsUnknown() && !rItem1.Filters.SSIDNum.IsNull() {
						return rItem1.Filters.SSIDNum.ValueInt64Pointer()
					}
					return nil
				}()
				tag := rItem1.Filters.Tag.ValueString()
				threshold := func() *int64 {
					if !rItem1.Filters.Threshold.IsUnknown() && !rItem1.Filters.Threshold.IsNull() {
						return rItem1.Filters.Threshold.ValueInt64Pointer()
					}
					return nil
				}()
				timeout := func() *int64 {
					if !rItem1.Filters.Timeout.IsUnknown() && !rItem1.Filters.Timeout.IsNull() {
						return rItem1.Filters.Timeout.ValueInt64Pointer()
					}
					return nil
				}()
				requestNetworksUpdateNetworkAlertsSettingsAlertsFilters = &merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsAlertsFilters{
					Conditions: func() *[]merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsAlertsFiltersConditions {
						if len(requestNetworksUpdateNetworkAlertsSettingsAlertsFiltersConditions) > 0 {
							return &requestNetworksUpdateNetworkAlertsSettingsAlertsFiltersConditions
						}
						return nil
					}(),
					FailureType:    failureType,
					LookbackWindow: int64ToIntPointer(lookbackWindow),
					MinDuration:    int64ToIntPointer(minDuration),
					Name:           name,
					Period:         int64ToIntPointer(period),
					Priority:       priority,
					Regex:          regex,
					Selector:       selector,
					Serials:        serials,
					SSIDNum:        int64ToIntPointer(ssidNum),
					Tag:            tag,
					Threshold:      int64ToIntPointer(threshold),
					Timeout:        int64ToIntPointer(timeout),
				}
				//[debug] Is Array: False
			}
			typeR := rItem1.Type.ValueString()
			requestNetworksUpdateNetworkAlertsSettingsAlerts = append(requestNetworksUpdateNetworkAlertsSettingsAlerts, merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsAlerts{
				AlertDestinations: requestNetworksUpdateNetworkAlertsSettingsAlertsAlertDestinations,
				Enabled:           enabled,
				Filters:           requestNetworksUpdateNetworkAlertsSettingsAlertsFilters,
				Type:              typeR,
			})
			//[debug] Is Array: True
		}
	}
	var requestNetworksUpdateNetworkAlertsSettingsDefaultDestinations *merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsDefaultDestinations

	if r.DefaultDestinations != nil {
		allAdmins := func() *bool {
			if !r.DefaultDestinations.AllAdmins.IsUnknown() && !r.DefaultDestinations.AllAdmins.IsNull() {
				return r.DefaultDestinations.AllAdmins.ValueBoolPointer()
			}
			return nil
		}()

		var emails []string = nil
		r.DefaultDestinations.Emails.ElementsAs(ctx, &emails, false)

		var httpServerIDs []string = nil
		r.DefaultDestinations.HTTPServerIDs.ElementsAs(ctx, &httpServerIDs, false)
		snmp := func() *bool {
			if !r.DefaultDestinations.SNMP.IsUnknown() && !r.DefaultDestinations.SNMP.IsNull() {
				return r.DefaultDestinations.SNMP.ValueBoolPointer()
			}
			return nil
		}()
		requestNetworksUpdateNetworkAlertsSettingsDefaultDestinations = &merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsDefaultDestinations{
			AllAdmins:     allAdmins,
			Emails:        emails,
			HTTPServerIDs: httpServerIDs,
			SNMP:          snmp,
		}
		//[debug] Is Array: False
	}
	var requestNetworksUpdateNetworkAlertsSettingsMuting *merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsMuting

	if r.Muting != nil {
		var requestNetworksUpdateNetworkAlertsSettingsMutingByPortSchedules *merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsMutingByPortSchedules

		if r.Muting.ByPortSchedules != nil {
			enabled := func() *bool {
				if !r.Muting.ByPortSchedules.Enabled.IsUnknown() && !r.Muting.ByPortSchedules.Enabled.IsNull() {
					return r.Muting.ByPortSchedules.Enabled.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkAlertsSettingsMutingByPortSchedules = &merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsMutingByPortSchedules{
				Enabled: enabled,
			}
			//[debug] Is Array: False
		}
		requestNetworksUpdateNetworkAlertsSettingsMuting = &merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsMuting{
			ByPortSchedules: requestNetworksUpdateNetworkAlertsSettingsMutingByPortSchedules,
		}
		//[debug] Is Array: False
	}
	out := merakigosdk.RequestNetworksUpdateNetworkAlertsSettings{
		Alerts: func() *[]merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsAlerts {
			if len(requestNetworksUpdateNetworkAlertsSettingsAlerts) > 0 {
				return &requestNetworksUpdateNetworkAlertsSettingsAlerts
			}
			return nil
		}(),
		DefaultDestinations: requestNetworksUpdateNetworkAlertsSettingsDefaultDestinations,
		Muting:              requestNetworksUpdateNetworkAlertsSettingsMuting,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseNetworksGetNetworkAlertsSettingsItemToBodyRs(state NetworksAlertsSettingsRs, response *merakigosdk.ResponseNetworksGetNetworkAlertsSettings, is_read bool) NetworksAlertsSettingsRs {
	itemState := NetworksAlertsSettingsRs{
		AlertsResponse: func() *[]ResponseNetworksGetNetworkAlertsSettingsAlertsRs {
			if response.Alerts != nil {
				result := make([]ResponseNetworksGetNetworkAlertsSettingsAlertsRs, len(*response.Alerts))
				for i, alerts := range *response.Alerts {
					result[i] = ResponseNetworksGetNetworkAlertsSettingsAlertsRs{
						AlertDestinations: func() *ResponseNetworksGetNetworkAlertsSettingsAlertsAlertDestinationsRs {
							if alerts.AlertDestinations != nil {
								return &ResponseNetworksGetNetworkAlertsSettingsAlertsAlertDestinationsRs{
									AllAdmins: func() types.Bool {
										if alerts.AlertDestinations.AllAdmins != nil {
											return types.BoolValue(*alerts.AlertDestinations.AllAdmins)
										}
										return types.Bool{}
									}(),
									Emails:        StringSliceToSet(alerts.AlertDestinations.Emails),
									HTTPServerIDs: StringSliceToSet(alerts.AlertDestinations.HTTPServerIDs),
									SmsNumbers:    StringSliceToSet(alerts.AlertDestinations.SmsNumbers),
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
						Filters: func() *ResponseNetworksGetNetworkAlertsSettingsAlertsFiltersRs {
							if alerts.Filters != nil {
								return &ResponseNetworksGetNetworkAlertsSettingsAlertsFiltersRs{
									Conditions: func() *[]ResponseNetworksGetNetworkAlertsSettingsAlertsFiltersConditionsRs {
										if alerts.Filters.Conditions != nil {
											result := make([]ResponseNetworksGetNetworkAlertsSettingsAlertsFiltersConditionsRs, len(*alerts.Filters.Conditions))
											for i, conditions := range *alerts.Filters.Conditions {
												result[i] = ResponseNetworksGetNetworkAlertsSettingsAlertsFiltersConditionsRs{
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
									Serials:  StringSliceToSet(alerts.Filters.Serials),
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
		DefaultDestinations: func() *ResponseNetworksGetNetworkAlertsSettingsDefaultDestinationsRs {
			if response.DefaultDestinations != nil {
				return &ResponseNetworksGetNetworkAlertsSettingsDefaultDestinationsRs{
					AllAdmins: func() types.Bool {
						if response.DefaultDestinations.AllAdmins != nil {
							return types.BoolValue(*response.DefaultDestinations.AllAdmins)
						}
						return types.Bool{}
					}(),
					Emails:        StringSliceToSet(response.DefaultDestinations.Emails),
					HTTPServerIDs: StringSliceToSet(response.DefaultDestinations.HTTPServerIDs),
					SNMP: func() types.Bool {
						if response.DefaultDestinations.SNMP != nil {
							return types.BoolValue(*response.DefaultDestinations.SNMP)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkAlertsSettingsDefaultDestinationsRs{}
		}(),
		Muting: state.Muting,
	}
	if itemState.DefaultDestinations != nil && state.DefaultDestinations != nil {
		itemState.DefaultDestinations.SNMP = state.DefaultDestinations.SNMP
	}

	itemState.Alerts = state.Alerts
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksAlertsSettingsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksAlertsSettingsRs)
}
