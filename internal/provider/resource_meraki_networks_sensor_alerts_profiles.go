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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSensorAlertsProfilesResource{}
	_ resource.ResourceWithConfigure = &NetworksSensorAlertsProfilesResource{}
)

func NewNetworksSensorAlertsProfilesResource() resource.Resource {
	return &NetworksSensorAlertsProfilesResource{}
}

type NetworksSensorAlertsProfilesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSensorAlertsProfilesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSensorAlertsProfilesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sensor_alerts_profiles"
}

func (r *NetworksSensorAlertsProfilesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"conditions": schema.SetNestedAttribute{
				MarkdownDescription: `List of conditions that will cause the profile to send an alert.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"direction": schema.StringAttribute{
							MarkdownDescription: `If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature and humidity thresholds.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"duration": schema.Int64Attribute{
							MarkdownDescription: `Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, and 1 hour. Default is 0.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"metric": schema.StringAttribute{
							MarkdownDescription: `The type of sensor metric that will be monitored for changes. Available metrics are door, humidity, indoorAirQuality, noise, pm25, temperature, tvoc, and water.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"threshold": schema.SingleNestedAttribute{
							MarkdownDescription: `Threshold for sensor readings that will cause an alert to be sent. This object should contain a single property key matching the condition's 'metric' value.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"door": schema.SingleNestedAttribute{
									MarkdownDescription: `Door open threshold. 'open' must be provided and set to true.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"open": schema.BoolAttribute{
											MarkdownDescription: `Alerting threshold for a door open event. Must be set to true.`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.Bool{
												boolplanmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"humidity": schema.SingleNestedAttribute{
									MarkdownDescription: `Humidity threshold. One of 'relativePercentage' or 'quality' must be provided.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"quality": schema.StringAttribute{
											MarkdownDescription: `Alerting threshold as a qualitative humidity level.`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.UseStateForUnknown(),
											},
										},
										"relative_percentage": schema.Int64Attribute{
											MarkdownDescription: `Alerting threshold in %RH.`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"indoor_air_quality": schema.SingleNestedAttribute{
									MarkdownDescription: `Indoor air quality score threshold. One of 'score' or 'quality' must be provided.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"quality": schema.StringAttribute{
											MarkdownDescription: `Alerting threshold as a qualitative indoor air quality level.`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.UseStateForUnknown(),
											},
										},
										"score": schema.Int64Attribute{
											MarkdownDescription: `Alerting threshold as indoor air quality score.`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"noise": schema.SingleNestedAttribute{
									MarkdownDescription: `Noise threshold. 'ambient' must be provided.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"ambient": schema.SingleNestedAttribute{
											MarkdownDescription: `Ambient noise threshold. One of 'level' or 'quality' must be provided.`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.Object{
												objectplanmodifier.UseStateForUnknown(),
											},
											Attributes: map[string]schema.Attribute{

												"level": schema.Int64Attribute{
													MarkdownDescription: `Alerting threshold as adjusted decibels.`,
													Computed:            true,
													Optional:            true,
													PlanModifiers: []planmodifier.Int64{
														int64planmodifier.UseStateForUnknown(),
													},
												},
												"quality": schema.StringAttribute{
													MarkdownDescription: `Alerting threshold as a qualitative ambient noise level.`,
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
								"pm25": schema.SingleNestedAttribute{
									MarkdownDescription: `PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"concentration": schema.Int64Attribute{
											MarkdownDescription: `Alerting threshold as PM2.5 parts per million.`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.UseStateForUnknown(),
											},
										},
										"quality": schema.StringAttribute{
											MarkdownDescription: `Alerting threshold as a qualitative PM2.5 level.`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"temperature": schema.SingleNestedAttribute{
									MarkdownDescription: `Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"celsius": schema.Float64Attribute{
											MarkdownDescription: `Alerting threshold in degrees Celsius.`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.Float64{
												float64planmodifier.UseStateForUnknown(),
											},
										},
										"fahrenheit": schema.Float64Attribute{
											MarkdownDescription: `Alerting threshold in degrees Fahrenheit.`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.Float64{
												float64planmodifier.UseStateForUnknown(),
											},
										},
										"quality": schema.StringAttribute{
											MarkdownDescription: `Alerting threshold as a qualitative temperature level.`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"tvoc": schema.SingleNestedAttribute{
									MarkdownDescription: `TVOC concentration threshold. One of 'concentration' or 'quality' must be provided.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"concentration": schema.Int64Attribute{
											MarkdownDescription: `Alerting threshold as TVOC micrograms per cubic meter.`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.UseStateForUnknown(),
											},
										},
										"quality": schema.StringAttribute{
											MarkdownDescription: `Alerting threshold as a qualitative TVOC level.`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"water": schema.SingleNestedAttribute{
									MarkdownDescription: `Water detection threshold. 'present' must be provided and set to true.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"present": schema.BoolAttribute{
											MarkdownDescription: `Alerting threshold for a water detection event. Must be set to true.`,
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
					},
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: `id path parameter.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Name of the sensor alert profile.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"profile_id": schema.StringAttribute{
				MarkdownDescription: `ID of the sensor alert profile.`,
				Computed:            true,
			},
			"recipients": schema.SingleNestedAttribute{
				MarkdownDescription: `List of recipients that will recieve the alert.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"emails": schema.SetAttribute{
						MarkdownDescription: `A list of emails that will receive information about the alert.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},

						ElementType: types.StringType,
					},
					"http_server_ids": schema.SetAttribute{
						MarkdownDescription: `A list of webhook endpoint IDs that will receive information about the alert.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},

						ElementType: types.StringType,
					},
					"sms_numbers": schema.SetAttribute{
						MarkdownDescription: `A list of SMS numbers that will receive information about the alert.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},

						ElementType: types.StringType,
					},
				},
			},
			"schedule": schema.SingleNestedAttribute{
				MarkdownDescription: `The sensor schedule to use with the alert profile.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"id": schema.StringAttribute{
						MarkdownDescription: `ID of the sensor schedule to use with the alert profile. If not defined, the alert profile will be active at all times.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of the sensor schedule to use with the alert profile.`,
						Computed:            true,
					},
				},
			},
			"serials": schema.SetAttribute{
				MarkdownDescription: `List of device serials assigned to this sensor alert profile.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
		},
	}
}

//path params to set ['id']

func (r *NetworksSensorAlertsProfilesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSensorAlertsProfilesRs

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
	// network_id
	vvName := data.Name.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Sensor.GetNetworkSensorAlertsProfiles(vvNetworkID)
	//Have Create
	if err != nil || restyResp1 == nil {
		if restyResp1.StatusCode() != 404 {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSensorAlertsProfiles",
				err.Error(),
			)
			return
		}
	}
	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvID, ok := result2["ProfileID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter ID",
					"Error",
				)
				return
			}
			r.client.Sensor.UpdateNetworkSensorAlertsProfile(vvNetworkID, vvID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Sensor.GetNetworkSensorAlertsProfile(vvNetworkID, vvID)
			if responseVerifyItem2 != nil {
				data = ResponseSensorGetNetworkSensorAlertsProfileItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Sensor.CreateNetworkSensorAlertsProfile(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkSensorAlertsProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkSensorAlertsProfile",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Sensor.GetNetworkSensorAlertsProfiles(vvNetworkID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSensorAlertsProfiles",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSensorAlertsProfiles",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvID, ok := result2["Name"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter ID",
				"Error",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Sensor.GetNetworkSensorAlertsProfile(vvNetworkID, vvID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseSensorGetNetworkSensorAlertsProfileItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkSensorAlertsProfile",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSensorAlertsProfile",
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

func (r *NetworksSensorAlertsProfilesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSensorAlertsProfilesRs

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
	// network_id
	vvID := data.ID.ValueString()
	// id
	responseGet, restyRespGet, err := r.client.Sensor.GetNetworkSensorAlertsProfile(vvNetworkID, vvID)
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
				"Failure when executing GetNetworkSensorAlertsProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSensorAlertsProfile",
			err.Error(),
		)
		return
	}

	data = ResponseSensorGetNetworkSensorAlertsProfileItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksSensorAlertsProfilesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
}

func (r *NetworksSensorAlertsProfilesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSensorAlertsProfilesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvID := data.ID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Sensor.UpdateNetworkSensorAlertsProfile(vvNetworkID, vvID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSensorAlertsProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSensorAlertsProfile",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSensorAlertsProfilesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksSensorAlertsProfilesRs
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
	vvID := state.ID.ValueString()
	_, err := r.client.Sensor.DeleteNetworkSensorAlertsProfile(vvNetworkID, vvID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkSensorAlertsProfile", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksSensorAlertsProfilesRs struct {
	NetworkID  types.String                                               `tfsdk:"network_id"`
	ID         types.String                                               `tfsdk:"id"`
	Conditions *[]ResponseSensorGetNetworkSensorAlertsProfileConditionsRs `tfsdk:"conditions"`
	Name       types.String                                               `tfsdk:"name"`
	ProfileID  types.String                                               `tfsdk:"profile_id"`
	Recipients *ResponseSensorGetNetworkSensorAlertsProfileRecipientsRs   `tfsdk:"recipients"`
	Schedule   *ResponseSensorGetNetworkSensorAlertsProfileScheduleRs     `tfsdk:"schedule"`
	Serials    types.Set                                                  `tfsdk:"serials"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsRs struct {
	Direction types.String                                                      `tfsdk:"direction"`
	Duration  types.Int64                                                       `tfsdk:"duration"`
	Metric    types.String                                                      `tfsdk:"metric"`
	Threshold *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdRs `tfsdk:"threshold"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdRs struct {
	Door             *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdDoorRs             `tfsdk:"door"`
	Humidity         *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdHumidityRs         `tfsdk:"humidity"`
	IndoorAirQuality *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdIndoorAirQualityRs `tfsdk:"indoor_air_quality"`
	Noise            *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoiseRs            `tfsdk:"noise"`
	Pm25             *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPm25Rs             `tfsdk:"pm25"`
	Temperature      *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTemperatureRs      `tfsdk:"temperature"`
	Tvoc             *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTvocRs             `tfsdk:"tvoc"`
	Water            *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdWaterRs            `tfsdk:"water"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdDoorRs struct {
	Open types.Bool `tfsdk:"open"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdHumidityRs struct {
	Quality            types.String `tfsdk:"quality"`
	RelativePercentage types.Int64  `tfsdk:"relative_percentage"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdIndoorAirQualityRs struct {
	Quality types.String `tfsdk:"quality"`
	Score   types.Int64  `tfsdk:"score"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoiseRs struct {
	Ambient *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoiseAmbientRs `tfsdk:"ambient"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoiseAmbientRs struct {
	Level   types.Int64  `tfsdk:"level"`
	Quality types.String `tfsdk:"quality"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPm25Rs struct {
	Concentration types.Int64  `tfsdk:"concentration"`
	Quality       types.String `tfsdk:"quality"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTemperatureRs struct {
	Celsius    types.Float64 `tfsdk:"celsius"`
	Fahrenheit types.Float64 `tfsdk:"fahrenheit"`
	Quality    types.String  `tfsdk:"quality"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTvocRs struct {
	Concentration types.Int64  `tfsdk:"concentration"`
	Quality       types.String `tfsdk:"quality"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdWaterRs struct {
	Present types.Bool `tfsdk:"present"`
}

type ResponseSensorGetNetworkSensorAlertsProfileRecipientsRs struct {
	Emails        types.Set `tfsdk:"emails"`
	HTTPServerIDs types.Set `tfsdk:"http_server_ids"`
	SmsNumbers    types.Set `tfsdk:"sms_numbers"`
}

type ResponseSensorGetNetworkSensorAlertsProfileScheduleRs struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// FromBody
func (r *NetworksSensorAlertsProfilesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfile {
	emptyString := ""
	var requestSensorCreateNetworkSensorAlertsProfileConditions []merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditions
	if r.Conditions != nil {
		for _, rItem1 := range *r.Conditions {
			direction := rItem1.Direction.ValueString()
			duration := func() *int64 {
				if !rItem1.Duration.IsUnknown() && !rItem1.Duration.IsNull() {
					return rItem1.Duration.ValueInt64Pointer()
				}
				return nil
			}()
			metric := rItem1.Metric.ValueString()
			var requestSensorCreateNetworkSensorAlertsProfileConditionsThreshold *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThreshold
			if rItem1.Threshold != nil {
				var requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdDoor *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdDoor
				if rItem1.Threshold.Door != nil {
					open := func() *bool {
						if !rItem1.Threshold.Door.Open.IsUnknown() && !rItem1.Threshold.Door.Open.IsNull() {
							return rItem1.Threshold.Door.Open.ValueBoolPointer()
						}
						return nil
					}()
					requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdDoor = &merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdDoor{
						Open: open,
					}
				}
				var requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdHumidity *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdHumidity
				if rItem1.Threshold.Humidity != nil {
					quality := rItem1.Threshold.Humidity.Quality.ValueString()
					relativePercentage := func() *int64 {
						if !rItem1.Threshold.Humidity.RelativePercentage.IsUnknown() && !rItem1.Threshold.Humidity.RelativePercentage.IsNull() {
							return rItem1.Threshold.Humidity.RelativePercentage.ValueInt64Pointer()
						}
						return nil
					}()
					requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdHumidity = &merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdHumidity{
						Quality:            quality,
						RelativePercentage: int64ToIntPointer(relativePercentage),
					}
				}
				var requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality
				if rItem1.Threshold.IndoorAirQuality != nil {
					quality := rItem1.Threshold.IndoorAirQuality.Quality.ValueString()
					score := func() *int64 {
						if !rItem1.Threshold.IndoorAirQuality.Score.IsUnknown() && !rItem1.Threshold.IndoorAirQuality.Score.IsNull() {
							return rItem1.Threshold.IndoorAirQuality.Score.ValueInt64Pointer()
						}
						return nil
					}()
					requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality = &merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality{
						Quality: quality,
						Score:   int64ToIntPointer(score),
					}
				}
				var requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoise *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoise
				if rItem1.Threshold.Noise != nil {
					var requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient
					if rItem1.Threshold.Noise.Ambient != nil {
						level := func() *int64 {
							if !rItem1.Threshold.Noise.Ambient.Level.IsUnknown() && !rItem1.Threshold.Noise.Ambient.Level.IsNull() {
								return rItem1.Threshold.Noise.Ambient.Level.ValueInt64Pointer()
							}
							return nil
						}()
						quality := rItem1.Threshold.Noise.Ambient.Quality.ValueString()
						requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient = &merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient{
							Level:   int64ToIntPointer(level),
							Quality: quality,
						}
					}
					requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoise = &merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoise{
						Ambient: requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient,
					}
				}
				var requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdPm25 *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdPm25
				if rItem1.Threshold.Pm25 != nil {
					concentration := func() *int64 {
						if !rItem1.Threshold.Pm25.Concentration.IsUnknown() && !rItem1.Threshold.Pm25.Concentration.IsNull() {
							return rItem1.Threshold.Pm25.Concentration.ValueInt64Pointer()
						}
						return nil
					}()
					quality := rItem1.Threshold.Pm25.Quality.ValueString()
					requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdPm25 = &merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdPm25{
						Concentration: int64ToIntPointer(concentration),
						Quality:       quality,
					}
				}
				var requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTemperature *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTemperature
				if rItem1.Threshold.Temperature != nil {
					celsius := func() *float64 {
						if !rItem1.Threshold.Temperature.Celsius.IsUnknown() && !rItem1.Threshold.Temperature.Celsius.IsNull() {
							return rItem1.Threshold.Temperature.Celsius.ValueFloat64Pointer()
						}
						return nil
					}()
					fahrenheit := func() *float64 {
						if !rItem1.Threshold.Temperature.Fahrenheit.IsUnknown() && !rItem1.Threshold.Temperature.Fahrenheit.IsNull() {
							return rItem1.Threshold.Temperature.Fahrenheit.ValueFloat64Pointer()
						}
						return nil
					}()
					quality := rItem1.Threshold.Temperature.Quality.ValueString()
					requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTemperature = &merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTemperature{
						Celsius:    celsius,
						Fahrenheit: fahrenheit,
						Quality:    quality,
					}
				}
				var requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTvoc *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTvoc
				if rItem1.Threshold.Tvoc != nil {
					concentration := func() *int64 {
						if !rItem1.Threshold.Tvoc.Concentration.IsUnknown() && !rItem1.Threshold.Tvoc.Concentration.IsNull() {
							return rItem1.Threshold.Tvoc.Concentration.ValueInt64Pointer()
						}
						return nil
					}()
					quality := rItem1.Threshold.Tvoc.Quality.ValueString()
					requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTvoc = &merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTvoc{
						Concentration: int64ToIntPointer(concentration),
						Quality:       quality,
					}
				}
				var requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdWater *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdWater
				if rItem1.Threshold.Water != nil {
					present := func() *bool {
						if !rItem1.Threshold.Water.Present.IsUnknown() && !rItem1.Threshold.Water.Present.IsNull() {
							return rItem1.Threshold.Water.Present.ValueBoolPointer()
						}
						return nil
					}()
					requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdWater = &merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdWater{
						Present: present,
					}
				}
				requestSensorCreateNetworkSensorAlertsProfileConditionsThreshold = &merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThreshold{
					Door:             requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdDoor,
					Humidity:         requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdHumidity,
					IndoorAirQuality: requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality,
					Noise:            requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoise,
					Pm25:             requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdPm25,
					Temperature:      requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTemperature,
					Tvoc:             requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTvoc,
					Water:            requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdWater,
				}
			}
			requestSensorCreateNetworkSensorAlertsProfileConditions = append(requestSensorCreateNetworkSensorAlertsProfileConditions, merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditions{
				Direction: direction,
				Duration:  int64ToIntPointer(duration),
				Metric:    metric,
				Threshold: requestSensorCreateNetworkSensorAlertsProfileConditionsThreshold,
			})
		}
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestSensorCreateNetworkSensorAlertsProfileRecipients *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileRecipients
	if r.Recipients != nil {
		var emails []string = nil

		r.Recipients.Emails.ElementsAs(ctx, &emails, false)
		var httpServerIDs []string = nil

		r.Recipients.HTTPServerIDs.ElementsAs(ctx, &httpServerIDs, false)
		var smsNumbers []string = nil

		r.Recipients.SmsNumbers.ElementsAs(ctx, &smsNumbers, false)
		requestSensorCreateNetworkSensorAlertsProfileRecipients = &merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileRecipients{
			Emails:        emails,
			HTTPServerIDs: httpServerIDs,
			SmsNumbers:    smsNumbers,
		}
	}
	var requestSensorCreateNetworkSensorAlertsProfileSchedule *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileSchedule
	if r.Schedule != nil {
		iD := r.Schedule.ID.ValueString()
		requestSensorCreateNetworkSensorAlertsProfileSchedule = &merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileSchedule{
			ID: iD,
		}
	}
	var serials []string = nil
	r.Serials.ElementsAs(ctx, &serials, false)
	out := merakigosdk.RequestSensorCreateNetworkSensorAlertsProfile{
		Conditions: func() *[]merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditions {
			if len(requestSensorCreateNetworkSensorAlertsProfileConditions) > 0 {
				return &requestSensorCreateNetworkSensorAlertsProfileConditions
			}
			return nil
		}(),
		Name:       *name,
		Recipients: requestSensorCreateNetworkSensorAlertsProfileRecipients,
		Schedule:   requestSensorCreateNetworkSensorAlertsProfileSchedule,
		Serials:    serials,
	}
	return &out
}
func (r *NetworksSensorAlertsProfilesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfile {
	emptyString := ""
	var requestSensorUpdateNetworkSensorAlertsProfileConditions []merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditions
	if r.Conditions != nil {
		for _, rItem1 := range *r.Conditions {
			direction := rItem1.Direction.ValueString()
			duration := func() *int64 {
				if !rItem1.Duration.IsUnknown() && !rItem1.Duration.IsNull() {
					return rItem1.Duration.ValueInt64Pointer()
				}
				return nil
			}()
			metric := rItem1.Metric.ValueString()
			var requestSensorUpdateNetworkSensorAlertsProfileConditionsThreshold *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThreshold
			if rItem1.Threshold != nil {
				var requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdDoor *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdDoor
				if rItem1.Threshold.Door != nil {
					open := func() *bool {
						if !rItem1.Threshold.Door.Open.IsUnknown() && !rItem1.Threshold.Door.Open.IsNull() {
							return rItem1.Threshold.Door.Open.ValueBoolPointer()
						}
						return nil
					}()
					requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdDoor = &merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdDoor{
						Open: open,
					}
				}
				var requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdHumidity *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdHumidity
				if rItem1.Threshold.Humidity != nil {
					quality := rItem1.Threshold.Humidity.Quality.ValueString()
					relativePercentage := func() *int64 {
						if !rItem1.Threshold.Humidity.RelativePercentage.IsUnknown() && !rItem1.Threshold.Humidity.RelativePercentage.IsNull() {
							return rItem1.Threshold.Humidity.RelativePercentage.ValueInt64Pointer()
						}
						return nil
					}()
					requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdHumidity = &merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdHumidity{
						Quality:            quality,
						RelativePercentage: int64ToIntPointer(relativePercentage),
					}
				}
				var requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality
				if rItem1.Threshold.IndoorAirQuality != nil {
					quality := rItem1.Threshold.IndoorAirQuality.Quality.ValueString()
					score := func() *int64 {
						if !rItem1.Threshold.IndoorAirQuality.Score.IsUnknown() && !rItem1.Threshold.IndoorAirQuality.Score.IsNull() {
							return rItem1.Threshold.IndoorAirQuality.Score.ValueInt64Pointer()
						}
						return nil
					}()
					requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality = &merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality{
						Quality: quality,
						Score:   int64ToIntPointer(score),
					}
				}
				var requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoise *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoise
				if rItem1.Threshold.Noise != nil {
					var requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient
					if rItem1.Threshold.Noise.Ambient != nil {
						level := func() *int64 {
							if !rItem1.Threshold.Noise.Ambient.Level.IsUnknown() && !rItem1.Threshold.Noise.Ambient.Level.IsNull() {
								return rItem1.Threshold.Noise.Ambient.Level.ValueInt64Pointer()
							}
							return nil
						}()
						quality := rItem1.Threshold.Noise.Ambient.Quality.ValueString()
						requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient = &merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient{
							Level:   int64ToIntPointer(level),
							Quality: quality,
						}
					}
					requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoise = &merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoise{
						Ambient: requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient,
					}
				}
				var requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPm25 *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPm25
				if rItem1.Threshold.Pm25 != nil {
					concentration := func() *int64 {
						if !rItem1.Threshold.Pm25.Concentration.IsUnknown() && !rItem1.Threshold.Pm25.Concentration.IsNull() {
							return rItem1.Threshold.Pm25.Concentration.ValueInt64Pointer()
						}
						return nil
					}()
					quality := rItem1.Threshold.Pm25.Quality.ValueString()
					requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPm25 = &merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPm25{
						Concentration: int64ToIntPointer(concentration),
						Quality:       quality,
					}
				}
				var requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTemperature *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTemperature
				if rItem1.Threshold.Temperature != nil {
					celsius := func() *float64 {
						if !rItem1.Threshold.Temperature.Celsius.IsUnknown() && !rItem1.Threshold.Temperature.Celsius.IsNull() {
							return rItem1.Threshold.Temperature.Celsius.ValueFloat64Pointer()
						}
						return nil
					}()
					fahrenheit := func() *float64 {
						if !rItem1.Threshold.Temperature.Fahrenheit.IsUnknown() && !rItem1.Threshold.Temperature.Fahrenheit.IsNull() {
							return rItem1.Threshold.Temperature.Fahrenheit.ValueFloat64Pointer()
						}
						return nil
					}()
					quality := rItem1.Threshold.Temperature.Quality.ValueString()
					requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTemperature = &merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTemperature{
						Celsius:    celsius,
						Fahrenheit: fahrenheit,
						Quality:    quality,
					}
				}
				var requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTvoc *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTvoc
				if rItem1.Threshold.Tvoc != nil {
					concentration := func() *int64 {
						if !rItem1.Threshold.Tvoc.Concentration.IsUnknown() && !rItem1.Threshold.Tvoc.Concentration.IsNull() {
							return rItem1.Threshold.Tvoc.Concentration.ValueInt64Pointer()
						}
						return nil
					}()
					quality := rItem1.Threshold.Tvoc.Quality.ValueString()
					requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTvoc = &merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTvoc{
						Concentration: int64ToIntPointer(concentration),
						Quality:       quality,
					}
				}
				var requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdWater *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdWater
				if rItem1.Threshold.Water != nil {
					present := func() *bool {
						if !rItem1.Threshold.Water.Present.IsUnknown() && !rItem1.Threshold.Water.Present.IsNull() {
							return rItem1.Threshold.Water.Present.ValueBoolPointer()
						}
						return nil
					}()
					requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdWater = &merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdWater{
						Present: present,
					}
				}
				requestSensorUpdateNetworkSensorAlertsProfileConditionsThreshold = &merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThreshold{
					Door:             requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdDoor,
					Humidity:         requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdHumidity,
					IndoorAirQuality: requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality,
					Noise:            requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoise,
					Pm25:             requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPm25,
					Temperature:      requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTemperature,
					Tvoc:             requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTvoc,
					Water:            requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdWater,
				}
			}
			requestSensorUpdateNetworkSensorAlertsProfileConditions = append(requestSensorUpdateNetworkSensorAlertsProfileConditions, merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditions{
				Direction: direction,
				Duration:  int64ToIntPointer(duration),
				Metric:    metric,
				Threshold: requestSensorUpdateNetworkSensorAlertsProfileConditionsThreshold,
			})
		}
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestSensorUpdateNetworkSensorAlertsProfileRecipients *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileRecipients
	if r.Recipients != nil {
		var emails []string = nil

		r.Recipients.Emails.ElementsAs(ctx, &emails, false)
		var httpServerIDs []string = nil

		r.Recipients.HTTPServerIDs.ElementsAs(ctx, &httpServerIDs, false)
		var smsNumbers []string = nil

		r.Recipients.SmsNumbers.ElementsAs(ctx, &smsNumbers, false)
		requestSensorUpdateNetworkSensorAlertsProfileRecipients = &merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileRecipients{
			Emails:        emails,
			HTTPServerIDs: httpServerIDs,
			SmsNumbers:    smsNumbers,
		}
	}
	var requestSensorUpdateNetworkSensorAlertsProfileSchedule *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileSchedule
	if r.Schedule != nil {
		iD := r.Schedule.ID.ValueString()
		requestSensorUpdateNetworkSensorAlertsProfileSchedule = &merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileSchedule{
			ID: iD,
		}
	}
	var serials []string = nil
	r.Serials.ElementsAs(ctx, &serials, false)
	out := merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfile{
		Conditions: func() *[]merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditions {
			if len(requestSensorUpdateNetworkSensorAlertsProfileConditions) > 0 {
				return &requestSensorUpdateNetworkSensorAlertsProfileConditions
			}
			return nil
		}(),
		Name:       *name,
		Recipients: requestSensorUpdateNetworkSensorAlertsProfileRecipients,
		Schedule:   requestSensorUpdateNetworkSensorAlertsProfileSchedule,
		Serials:    serials,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSensorGetNetworkSensorAlertsProfileItemToBodyRs(state NetworksSensorAlertsProfilesRs, response *merakigosdk.ResponseSensorGetNetworkSensorAlertsProfile, is_read bool) NetworksSensorAlertsProfilesRs {
	itemState := NetworksSensorAlertsProfilesRs{
		Conditions: func() *[]ResponseSensorGetNetworkSensorAlertsProfileConditionsRs {
			if response.Conditions != nil {
				result := make([]ResponseSensorGetNetworkSensorAlertsProfileConditionsRs, len(*response.Conditions))
				for i, conditions := range *response.Conditions {
					result[i] = ResponseSensorGetNetworkSensorAlertsProfileConditionsRs{
						Direction: types.StringValue(conditions.Direction),
						Duration: func() types.Int64 {
							if conditions.Duration != nil {
								return types.Int64Value(int64(*conditions.Duration))
							}
							return types.Int64{}
						}(),
						Metric: types.StringValue(conditions.Metric),
						Threshold: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdRs {
							if conditions.Threshold != nil {
								return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdRs{
									Door: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdDoorRs {
										if conditions.Threshold.Door != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdDoorRs{
												Open: func() types.Bool {
													if conditions.Threshold.Door.Open != nil {
														return types.BoolValue(*conditions.Threshold.Door.Open)
													}
													return types.Bool{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdDoorRs{}
									}(),
									Humidity: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdHumidityRs {
										if conditions.Threshold.Humidity != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdHumidityRs{
												Quality: types.StringValue(conditions.Threshold.Humidity.Quality),
												RelativePercentage: func() types.Int64 {
													if conditions.Threshold.Humidity.RelativePercentage != nil {
														return types.Int64Value(int64(*conditions.Threshold.Humidity.RelativePercentage))
													}
													return types.Int64{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdHumidityRs{}
									}(),
									IndoorAirQuality: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdIndoorAirQualityRs {
										if conditions.Threshold.IndoorAirQuality != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdIndoorAirQualityRs{
												Quality: types.StringValue(conditions.Threshold.IndoorAirQuality.Quality),
												Score: func() types.Int64 {
													if conditions.Threshold.IndoorAirQuality.Score != nil {
														return types.Int64Value(int64(*conditions.Threshold.IndoorAirQuality.Score))
													}
													return types.Int64{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdIndoorAirQualityRs{}
									}(),
									Noise: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoiseRs {
										if conditions.Threshold.Noise != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoiseRs{
												Ambient: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoiseAmbientRs {
													if conditions.Threshold.Noise.Ambient != nil {
														return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoiseAmbientRs{
															Level: func() types.Int64 {
																if conditions.Threshold.Noise.Ambient.Level != nil {
																	return types.Int64Value(int64(*conditions.Threshold.Noise.Ambient.Level))
																}
																return types.Int64{}
															}(),
															Quality: types.StringValue(conditions.Threshold.Noise.Ambient.Quality),
														}
													}
													return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoiseAmbientRs{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoiseRs{}
									}(),
									Pm25: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPm25Rs {
										if conditions.Threshold.Pm25 != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPm25Rs{
												Concentration: func() types.Int64 {
													if conditions.Threshold.Pm25.Concentration != nil {
														return types.Int64Value(int64(*conditions.Threshold.Pm25.Concentration))
													}
													return types.Int64{}
												}(),
												Quality: types.StringValue(conditions.Threshold.Pm25.Quality),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPm25Rs{}
									}(),
									Temperature: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTemperatureRs {
										if conditions.Threshold.Temperature != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTemperatureRs{
												Celsius: func() types.Float64 {
													if conditions.Threshold.Temperature.Celsius != nil {
														return types.Float64Value(float64(*conditions.Threshold.Temperature.Celsius))
													}
													return types.Float64{}
												}(),
												Fahrenheit: func() types.Float64 {
													if conditions.Threshold.Temperature.Fahrenheit != nil {
														return types.Float64Value(float64(*conditions.Threshold.Temperature.Fahrenheit))
													}
													return types.Float64{}
												}(),
												Quality: types.StringValue(conditions.Threshold.Temperature.Quality),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTemperatureRs{}
									}(),
									Tvoc: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTvocRs {
										if conditions.Threshold.Tvoc != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTvocRs{
												Concentration: func() types.Int64 {
													if conditions.Threshold.Tvoc.Concentration != nil {
														return types.Int64Value(int64(*conditions.Threshold.Tvoc.Concentration))
													}
													return types.Int64{}
												}(),
												Quality: types.StringValue(conditions.Threshold.Tvoc.Quality),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTvocRs{}
									}(),
									Water: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdWaterRs {
										if conditions.Threshold.Water != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdWaterRs{
												Present: func() types.Bool {
													if conditions.Threshold.Water.Present != nil {
														return types.BoolValue(*conditions.Threshold.Water.Present)
													}
													return types.Bool{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdWaterRs{}
									}(),
								}
							}
							return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdRs{}
						}(),
					}
				}
				return &result
			}
			return &[]ResponseSensorGetNetworkSensorAlertsProfileConditionsRs{}
		}(),
		Name:      types.StringValue(response.Name),
		ProfileID: types.StringValue(response.ProfileID),
		Recipients: func() *ResponseSensorGetNetworkSensorAlertsProfileRecipientsRs {
			if response.Recipients != nil {
				return &ResponseSensorGetNetworkSensorAlertsProfileRecipientsRs{
					Emails:        StringSliceToSet(response.Recipients.Emails),
					HTTPServerIDs: StringSliceToSet(response.Recipients.HTTPServerIDs),
					SmsNumbers:    StringSliceToSet(response.Recipients.SmsNumbers),
				}
			}
			return &ResponseSensorGetNetworkSensorAlertsProfileRecipientsRs{}
		}(),
		Schedule: func() *ResponseSensorGetNetworkSensorAlertsProfileScheduleRs {
			if response.Schedule != nil {
				return &ResponseSensorGetNetworkSensorAlertsProfileScheduleRs{
					ID:   types.StringValue(response.Schedule.ID),
					Name: types.StringValue(response.Schedule.Name),
				}
			}
			return &ResponseSensorGetNetworkSensorAlertsProfileScheduleRs{}
		}(),
		Serials: StringSliceToSet(response.Serials),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSensorAlertsProfilesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSensorAlertsProfilesRs)
}
