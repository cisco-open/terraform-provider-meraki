package provider

// DATA SOURCE NORMAL
import (
	"context"
	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksSensorAlertsProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSensorAlertsProfilesDataSource{}
)

func NewNetworksSensorAlertsProfilesDataSource() datasource.DataSource {
	return &NetworksSensorAlertsProfilesDataSource{}
}

type NetworksSensorAlertsProfilesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSensorAlertsProfilesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSensorAlertsProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sensor_alerts_profiles"
}

func (d *NetworksSensorAlertsProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: `id path parameter.`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"conditions": schema.SetNestedAttribute{
						MarkdownDescription: `List of conditions that will cause the profile to send an alert.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"direction": schema.StringAttribute{
									MarkdownDescription: `If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature, humidity, realPower, apparentPower, powerFactor, voltage, current, and frequency thresholds.`,
									Computed:            true,
								},
								"duration": schema.Int64Attribute{
									MarkdownDescription: `Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, 1 hour, 2 hours, 4 hours, and 8 hours. Default is 0.`,
									Computed:            true,
								},
								"metric": schema.StringAttribute{
									MarkdownDescription: `The type of sensor metric that will be monitored for changes. Available metrics are apparentPower, co2, current, door, frequency, humidity, indoorAirQuality, noise, pm25, powerFactor, realPower, temperature, tvoc, upstreamPower, voltage, and water.`,
									Computed:            true,
								},
								"threshold": schema.SingleNestedAttribute{
									MarkdownDescription: `Threshold for sensor readings that will cause an alert to be sent. This object should contain a single property key matching the condition's 'metric' value.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"apparent_power": schema.SingleNestedAttribute{
											MarkdownDescription: `Apparent power threshold. 'draw' must be provided.`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"draw": schema.Float64Attribute{
													MarkdownDescription: `Alerting threshold in volt-amps. Must be between 0 and 3750.`,
													Computed:            true,
												},
											},
										},
										"current": schema.SingleNestedAttribute{
											MarkdownDescription: `Electrical current threshold. 'level' must be provided.`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"draw": schema.Float64Attribute{
													MarkdownDescription: `Alerting threshold in amps. Must be between 0 and 15.`,
													Computed:            true,
												},
											},
										},
										"door": schema.SingleNestedAttribute{
											MarkdownDescription: `Door open threshold. 'open' must be provided and set to true.`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"open": schema.BoolAttribute{
													MarkdownDescription: `Alerting threshold for a door open event. Must be set to true.`,
													Computed:            true,
												},
											},
										},
										"frequency": schema.SingleNestedAttribute{
											MarkdownDescription: `Electrical frequency threshold. 'level' must be provided.`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"level": schema.Float64Attribute{
													MarkdownDescription: `Alerting threshold in hertz. Must be between 0 and 60.`,
													Computed:            true,
												},
											},
										},
										"humidity": schema.SingleNestedAttribute{
											MarkdownDescription: `Humidity threshold. One of 'relativePercentage' or 'quality' must be provided.`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"quality": schema.StringAttribute{
													MarkdownDescription: `Alerting threshold as a qualitative humidity level.`,
													Computed:            true,
												},
												"relative_percentage": schema.Int64Attribute{
													MarkdownDescription: `Alerting threshold in %RH.`,
													Computed:            true,
												},
											},
										},
										"indoor_air_quality": schema.SingleNestedAttribute{
											MarkdownDescription: `Indoor air quality score threshold. One of 'score' or 'quality' must be provided.`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"quality": schema.StringAttribute{
													MarkdownDescription: `Alerting threshold as a qualitative indoor air quality level.`,
													Computed:            true,
												},
												"score": schema.Int64Attribute{
													MarkdownDescription: `Alerting threshold as indoor air quality score.`,
													Computed:            true,
												},
											},
										},
										"noise": schema.SingleNestedAttribute{
											MarkdownDescription: `Noise threshold. 'ambient' must be provided.`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"ambient": schema.SingleNestedAttribute{
													MarkdownDescription: `Ambient noise threshold. One of 'level' or 'quality' must be provided.`,
													Computed:            true,
													Attributes: map[string]schema.Attribute{

														"level": schema.Int64Attribute{
															MarkdownDescription: `Alerting threshold as adjusted decibels.`,
															Computed:            true,
														},
														"quality": schema.StringAttribute{
															MarkdownDescription: `Alerting threshold as a qualitative ambient noise level.`,
															Computed:            true,
														},
													},
												},
											},
										},
										"pm25": schema.SingleNestedAttribute{
											MarkdownDescription: `PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided.`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"concentration": schema.Int64Attribute{
													MarkdownDescription: `Alerting threshold as PM2.5 parts per million.`,
													Computed:            true,
												},
												"quality": schema.StringAttribute{
													MarkdownDescription: `Alerting threshold as a qualitative PM2.5 level.`,
													Computed:            true,
												},
											},
										},
										"power_factor": schema.SingleNestedAttribute{
											MarkdownDescription: `Power factor threshold. 'percentage' must be provided.`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"percentage": schema.Int64Attribute{
													MarkdownDescription: `Alerting threshold as the ratio of active power to apparent power. Must be between 0 and 100.`,
													Computed:            true,
												},
											},
										},
										"real_power": schema.SingleNestedAttribute{
											MarkdownDescription: `Real power threshold. 'draw' must be provided.`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"draw": schema.Float64Attribute{
													MarkdownDescription: `Alerting threshold in watts. Must be between 0 and 3750.`,
													Computed:            true,
												},
											},
										},
										"temperature": schema.SingleNestedAttribute{
											MarkdownDescription: `Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided.`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"celsius": schema.Float64Attribute{
													MarkdownDescription: `Alerting threshold in degrees Celsius.`,
													Computed:            true,
												},
												"fahrenheit": schema.Float64Attribute{
													MarkdownDescription: `Alerting threshold in degrees Fahrenheit.`,
													Computed:            true,
												},
												"quality": schema.StringAttribute{
													MarkdownDescription: `Alerting threshold as a qualitative temperature level.`,
													Computed:            true,
												},
											},
										},
										"tvoc": schema.SingleNestedAttribute{
											MarkdownDescription: `TVOC concentration threshold. One of 'concentration' or 'quality' must be provided.`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"concentration": schema.Int64Attribute{
													MarkdownDescription: `Alerting threshold as TVOC micrograms per cubic meter.`,
													Computed:            true,
												},
												"quality": schema.StringAttribute{
													MarkdownDescription: `Alerting threshold as a qualitative TVOC level.`,
													Computed:            true,
												},
											},
										},
										"upstream_power": schema.SingleNestedAttribute{
											MarkdownDescription: `Upstream power threshold. 'outageDetected' must be provided and set to true.`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"outage_detected": schema.BoolAttribute{
													MarkdownDescription: `Alerting threshold for an upstream power event. Must be set to true.`,
													Computed:            true,
												},
											},
										},
										"voltage": schema.SingleNestedAttribute{
											MarkdownDescription: `Voltage threshold. 'level' must be provided.`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"level": schema.Float64Attribute{
													MarkdownDescription: `Alerting threshold in volts. Must be between 0 and 250.`,
													Computed:            true,
												},
											},
										},
										"water": schema.SingleNestedAttribute{
											MarkdownDescription: `Water detection threshold. 'present' must be provided and set to true.`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"present": schema.BoolAttribute{
													MarkdownDescription: `Alerting threshold for a water detection event. Must be set to true.`,
													Computed:            true,
												},
											},
										},
									},
								},
							},
						},
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of the sensor alert profile.`,
						Computed:            true,
					},
					"profile_id": schema.StringAttribute{
						MarkdownDescription: `ID of the sensor alert profile.`,
						Computed:            true,
					},
					"recipients": schema.SingleNestedAttribute{
						MarkdownDescription: `List of recipients that will receive the alert.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"emails": schema.ListAttribute{
								MarkdownDescription: `A list of emails that will receive information about the alert.`,
								Computed:            true,
								ElementType:         types.StringType,
							},
							"http_server_ids": schema.ListAttribute{
								MarkdownDescription: `A list of webhook endpoint IDs that will receive information about the alert.`,
								Computed:            true,
								ElementType:         types.StringType,
							},
							"sms_numbers": schema.ListAttribute{
								MarkdownDescription: `A list of SMS numbers that will receive information about the alert.`,
								Computed:            true,
								ElementType:         types.StringType,
							},
						},
					},
					"schedule": schema.SingleNestedAttribute{
						MarkdownDescription: `The sensor schedule to use with the alert profile.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"id": schema.StringAttribute{
								MarkdownDescription: `ID of the sensor schedule to use with the alert profile. If not defined, the alert profile will be active at all times.`,
								Computed:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `Name of the sensor schedule to use with the alert profile.`,
								Computed:            true,
							},
						},
					},
					"serials": schema.ListAttribute{
						MarkdownDescription: `List of device serials assigned to this sensor alert profile.`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSensorGetNetworkSensorAlertsProfiles`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"conditions": schema.SetNestedAttribute{
							MarkdownDescription: `List of conditions that will cause the profile to send an alert.`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"direction": schema.StringAttribute{
										MarkdownDescription: `If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature, humidity, realPower, apparentPower, powerFactor, voltage, current, and frequency thresholds.`,
										Computed:            true,
									},
									"duration": schema.Int64Attribute{
										MarkdownDescription: `Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, 1 hour, 2 hours, 4 hours, and 8 hours. Default is 0.`,
										Computed:            true,
									},
									"metric": schema.StringAttribute{
										MarkdownDescription: `The type of sensor metric that will be monitored for changes. Available metrics are apparentPower, co2, current, door, frequency, humidity, indoorAirQuality, noise, pm25, powerFactor, realPower, temperature, tvoc, upstreamPower, voltage, and water.`,
										Computed:            true,
									},
									"threshold": schema.SingleNestedAttribute{
										MarkdownDescription: `Threshold for sensor readings that will cause an alert to be sent. This object should contain a single property key matching the condition's 'metric' value.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"apparent_power": schema.SingleNestedAttribute{
												MarkdownDescription: `Apparent power threshold. 'draw' must be provided.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"draw": schema.Float64Attribute{
														MarkdownDescription: `Alerting threshold in volt-amps. Must be between 0 and 3750.`,
														Computed:            true,
													},
												},
											},
											"current": schema.SingleNestedAttribute{
												MarkdownDescription: `Electrical current threshold. 'level' must be provided.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"draw": schema.Float64Attribute{
														MarkdownDescription: `Alerting threshold in amps. Must be between 0 and 15.`,
														Computed:            true,
													},
												},
											},
											"door": schema.SingleNestedAttribute{
												MarkdownDescription: `Door open threshold. 'open' must be provided and set to true.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"open": schema.BoolAttribute{
														MarkdownDescription: `Alerting threshold for a door open event. Must be set to true.`,
														Computed:            true,
													},
												},
											},
											"frequency": schema.SingleNestedAttribute{
												MarkdownDescription: `Electrical frequency threshold. 'level' must be provided.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"level": schema.Float64Attribute{
														MarkdownDescription: `Alerting threshold in hertz. Must be between 0 and 60.`,
														Computed:            true,
													},
												},
											},
											"humidity": schema.SingleNestedAttribute{
												MarkdownDescription: `Humidity threshold. One of 'relativePercentage' or 'quality' must be provided.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"quality": schema.StringAttribute{
														MarkdownDescription: `Alerting threshold as a qualitative humidity level.`,
														Computed:            true,
													},
													"relative_percentage": schema.Int64Attribute{
														MarkdownDescription: `Alerting threshold in %RH.`,
														Computed:            true,
													},
												},
											},
											"indoor_air_quality": schema.SingleNestedAttribute{
												MarkdownDescription: `Indoor air quality score threshold. One of 'score' or 'quality' must be provided.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"quality": schema.StringAttribute{
														MarkdownDescription: `Alerting threshold as a qualitative indoor air quality level.`,
														Computed:            true,
													},
													"score": schema.Int64Attribute{
														MarkdownDescription: `Alerting threshold as indoor air quality score.`,
														Computed:            true,
													},
												},
											},
											"noise": schema.SingleNestedAttribute{
												MarkdownDescription: `Noise threshold. 'ambient' must be provided.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"ambient": schema.SingleNestedAttribute{
														MarkdownDescription: `Ambient noise threshold. One of 'level' or 'quality' must be provided.`,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"level": schema.Int64Attribute{
																MarkdownDescription: `Alerting threshold as adjusted decibels.`,
																Computed:            true,
															},
															"quality": schema.StringAttribute{
																MarkdownDescription: `Alerting threshold as a qualitative ambient noise level.`,
																Computed:            true,
															},
														},
													},
												},
											},
											"pm25": schema.SingleNestedAttribute{
												MarkdownDescription: `PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"concentration": schema.Int64Attribute{
														MarkdownDescription: `Alerting threshold as PM2.5 parts per million.`,
														Computed:            true,
													},
													"quality": schema.StringAttribute{
														MarkdownDescription: `Alerting threshold as a qualitative PM2.5 level.`,
														Computed:            true,
													},
												},
											},
											"power_factor": schema.SingleNestedAttribute{
												MarkdownDescription: `Power factor threshold. 'percentage' must be provided.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"percentage": schema.Int64Attribute{
														MarkdownDescription: `Alerting threshold as the ratio of active power to apparent power. Must be between 0 and 100.`,
														Computed:            true,
													},
												},
											},
											"real_power": schema.SingleNestedAttribute{
												MarkdownDescription: `Real power threshold. 'draw' must be provided.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"draw": schema.Float64Attribute{
														MarkdownDescription: `Alerting threshold in watts. Must be between 0 and 3750.`,
														Computed:            true,
													},
												},
											},
											"temperature": schema.SingleNestedAttribute{
												MarkdownDescription: `Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"celsius": schema.Float64Attribute{
														MarkdownDescription: `Alerting threshold in degrees Celsius.`,
														Computed:            true,
													},
													"fahrenheit": schema.Float64Attribute{
														MarkdownDescription: `Alerting threshold in degrees Fahrenheit.`,
														Computed:            true,
													},
													"quality": schema.StringAttribute{
														MarkdownDescription: `Alerting threshold as a qualitative temperature level.`,
														Computed:            true,
													},
												},
											},
											"tvoc": schema.SingleNestedAttribute{
												MarkdownDescription: `TVOC concentration threshold. One of 'concentration' or 'quality' must be provided.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"concentration": schema.Int64Attribute{
														MarkdownDescription: `Alerting threshold as TVOC micrograms per cubic meter.`,
														Computed:            true,
													},
													"quality": schema.StringAttribute{
														MarkdownDescription: `Alerting threshold as a qualitative TVOC level.`,
														Computed:            true,
													},
												},
											},
											"upstream_power": schema.SingleNestedAttribute{
												MarkdownDescription: `Upstream power threshold. 'outageDetected' must be provided and set to true.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"outage_detected": schema.BoolAttribute{
														MarkdownDescription: `Alerting threshold for an upstream power event. Must be set to true.`,
														Computed:            true,
													},
												},
											},
											"voltage": schema.SingleNestedAttribute{
												MarkdownDescription: `Voltage threshold. 'level' must be provided.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"level": schema.Float64Attribute{
														MarkdownDescription: `Alerting threshold in volts. Must be between 0 and 250.`,
														Computed:            true,
													},
												},
											},
											"water": schema.SingleNestedAttribute{
												MarkdownDescription: `Water detection threshold. 'present' must be provided and set to true.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"present": schema.BoolAttribute{
														MarkdownDescription: `Alerting threshold for a water detection event. Must be set to true.`,
														Computed:            true,
													},
												},
											},
										},
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the sensor alert profile.`,
							Computed:            true,
						},
						"profile_id": schema.StringAttribute{
							MarkdownDescription: `ID of the sensor alert profile.`,
							Computed:            true,
						},
						"recipients": schema.SingleNestedAttribute{
							MarkdownDescription: `List of recipients that will receive the alert.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"emails": schema.ListAttribute{
									MarkdownDescription: `A list of emails that will receive information about the alert.`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"http_server_ids": schema.ListAttribute{
									MarkdownDescription: `A list of webhook endpoint IDs that will receive information about the alert.`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"sms_numbers": schema.ListAttribute{
									MarkdownDescription: `A list of SMS numbers that will receive information about the alert.`,
									Computed:            true,
									ElementType:         types.StringType,
								},
							},
						},
						"schedule": schema.SingleNestedAttribute{
							MarkdownDescription: `The sensor schedule to use with the alert profile.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `ID of the sensor schedule to use with the alert profile. If not defined, the alert profile will be active at all times.`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Name of the sensor schedule to use with the alert profile.`,
									Computed:            true,
								},
							},
						},
						"serials": schema.ListAttribute{
							MarkdownDescription: `List of device serials assigned to this sensor alert profile.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSensorAlertsProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSensorAlertsProfiles NetworksSensorAlertsProfiles
	diags := req.Config.Get(ctx, &networksSensorAlertsProfiles)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksSensorAlertsProfiles.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksSensorAlertsProfiles.NetworkID.IsNull(), !networksSensorAlertsProfiles.ID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSensorAlertsProfiles")
		vvNetworkID := networksSensorAlertsProfiles.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Sensor.GetNetworkSensorAlertsProfiles(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSensorAlertsProfiles",
				err.Error(),
			)
			return
		}

		networksSensorAlertsProfiles = ResponseSensorGetNetworkSensorAlertsProfilesItemsToBody(networksSensorAlertsProfiles, response1)
		diags = resp.State.Set(ctx, &networksSensorAlertsProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkSensorAlertsProfile")
		vvNetworkID := networksSensorAlertsProfiles.NetworkID.ValueString()
		vvID := networksSensorAlertsProfiles.ID.ValueString()

		response2, restyResp2, err := d.client.Sensor.GetNetworkSensorAlertsProfile(vvNetworkID, vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSensorAlertsProfile",
				err.Error(),
			)
			return
		}

		networksSensorAlertsProfiles = ResponseSensorGetNetworkSensorAlertsProfileItemToBody(networksSensorAlertsProfiles, response2)
		diags = resp.State.Set(ctx, &networksSensorAlertsProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSensorAlertsProfiles struct {
	NetworkID types.String                                        `tfsdk:"network_id"`
	ID        types.String                                        `tfsdk:"id"`
	Items     *[]ResponseItemSensorGetNetworkSensorAlertsProfiles `tfsdk:"items"`
	Item      *ResponseSensorGetNetworkSensorAlertsProfile        `tfsdk:"item"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfiles struct {
	Conditions *[]ResponseItemSensorGetNetworkSensorAlertsProfilesConditions `tfsdk:"conditions"`
	Name       types.String                                                  `tfsdk:"name"`
	ProfileID  types.String                                                  `tfsdk:"profile_id"`
	Recipients *ResponseItemSensorGetNetworkSensorAlertsProfilesRecipients   `tfsdk:"recipients"`
	Schedule   *ResponseItemSensorGetNetworkSensorAlertsProfilesSchedule     `tfsdk:"schedule"`
	Serials    types.List                                                    `tfsdk:"serials"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfilesConditions struct {
	Direction types.String                                                         `tfsdk:"direction"`
	Duration  types.Int64                                                          `tfsdk:"duration"`
	Metric    types.String                                                         `tfsdk:"metric"`
	Threshold *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThreshold `tfsdk:"threshold"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThreshold struct {
	ApparentPower    *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdApparentPower    `tfsdk:"apparent_power"`
	Current          *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdCurrent          `tfsdk:"current"`
	Door             *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdDoor             `tfsdk:"door"`
	Frequency        *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdFrequency        `tfsdk:"frequency"`
	Humidity         *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdHumidity         `tfsdk:"humidity"`
	IndoorAirQuality *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdIndoorAirQuality `tfsdk:"indoor_air_quality"`
	Noise            *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdNoise            `tfsdk:"noise"`
	Pm25             *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdPm25             `tfsdk:"pm25"`
	PowerFactor      *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdPowerFactor      `tfsdk:"power_factor"`
	RealPower        *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdRealPower        `tfsdk:"real_power"`
	Temperature      *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdTemperature      `tfsdk:"temperature"`
	Tvoc             *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdTvoc             `tfsdk:"tvoc"`
	UpstreamPower    *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdUpstreamPower    `tfsdk:"upstream_power"`
	Voltage          *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdVoltage          `tfsdk:"voltage"`
	Water            *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdWater            `tfsdk:"water"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdApparentPower struct {
	Draw types.Float64 `tfsdk:"draw"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdCurrent struct {
	Draw types.Float64 `tfsdk:"draw"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdDoor struct {
	Open types.Bool `tfsdk:"open"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdFrequency struct {
	Level types.Float64 `tfsdk:"level"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdHumidity struct {
	Quality            types.String `tfsdk:"quality"`
	RelativePercentage types.Int64  `tfsdk:"relative_percentage"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdIndoorAirQuality struct {
	Quality types.String `tfsdk:"quality"`
	Score   types.Int64  `tfsdk:"score"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdNoise struct {
	Ambient *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdNoiseAmbient `tfsdk:"ambient"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdNoiseAmbient struct {
	Level   types.Int64  `tfsdk:"level"`
	Quality types.String `tfsdk:"quality"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdPm25 struct {
	Concentration types.Int64  `tfsdk:"concentration"`
	Quality       types.String `tfsdk:"quality"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdPowerFactor struct {
	Percentage types.Int64 `tfsdk:"percentage"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdRealPower struct {
	Draw types.Float64 `tfsdk:"draw"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdTemperature struct {
	Celsius    types.Float64 `tfsdk:"celsius"`
	Fahrenheit types.Float64 `tfsdk:"fahrenheit"`
	Quality    types.String  `tfsdk:"quality"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdTvoc struct {
	Concentration types.Int64  `tfsdk:"concentration"`
	Quality       types.String `tfsdk:"quality"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdUpstreamPower struct {
	OutageDetected types.Bool `tfsdk:"outage_detected"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdVoltage struct {
	Level types.Float64 `tfsdk:"level"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdWater struct {
	Present types.Bool `tfsdk:"present"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfilesRecipients struct {
	Emails        types.List `tfsdk:"emails"`
	HTTPServerIDs types.List `tfsdk:"http_server_ids"`
	SmsNumbers    types.List `tfsdk:"sms_numbers"`
}

type ResponseItemSensorGetNetworkSensorAlertsProfilesSchedule struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseSensorGetNetworkSensorAlertsProfile struct {
	Conditions *[]ResponseSensorGetNetworkSensorAlertsProfileConditions `tfsdk:"conditions"`
	Name       types.String                                             `tfsdk:"name"`
	ProfileID  types.String                                             `tfsdk:"profile_id"`
	Recipients *ResponseSensorGetNetworkSensorAlertsProfileRecipients   `tfsdk:"recipients"`
	Schedule   *ResponseSensorGetNetworkSensorAlertsProfileSchedule     `tfsdk:"schedule"`
	Serials    types.List                                               `tfsdk:"serials"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditions struct {
	Direction types.String                                                    `tfsdk:"direction"`
	Duration  types.Int64                                                     `tfsdk:"duration"`
	Metric    types.String                                                    `tfsdk:"metric"`
	Threshold *ResponseSensorGetNetworkSensorAlertsProfileConditionsThreshold `tfsdk:"threshold"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThreshold struct {
	ApparentPower    *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdApparentPower    `tfsdk:"apparent_power"`
	Current          *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdCurrent          `tfsdk:"current"`
	Door             *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdDoor             `tfsdk:"door"`
	Frequency        *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdFrequency        `tfsdk:"frequency"`
	Humidity         *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdHumidity         `tfsdk:"humidity"`
	IndoorAirQuality *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality `tfsdk:"indoor_air_quality"`
	Noise            *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoise            `tfsdk:"noise"`
	Pm25             *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPm25             `tfsdk:"pm25"`
	PowerFactor      *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPowerFactor      `tfsdk:"power_factor"`
	RealPower        *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdRealPower        `tfsdk:"real_power"`
	Temperature      *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTemperature      `tfsdk:"temperature"`
	Tvoc             *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTvoc             `tfsdk:"tvoc"`
	UpstreamPower    *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdUpstreamPower    `tfsdk:"upstream_power"`
	Voltage          *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdVoltage          `tfsdk:"voltage"`
	Water            *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdWater            `tfsdk:"water"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdApparentPower struct {
	Draw types.Float64 `tfsdk:"draw"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdCurrent struct {
	Draw types.Float64 `tfsdk:"draw"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdDoor struct {
	Open types.Bool `tfsdk:"open"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdFrequency struct {
	Level types.Float64 `tfsdk:"level"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdHumidity struct {
	Quality            types.String `tfsdk:"quality"`
	RelativePercentage types.Int64  `tfsdk:"relative_percentage"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality struct {
	Quality types.String `tfsdk:"quality"`
	Score   types.Int64  `tfsdk:"score"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoise struct {
	Ambient *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient `tfsdk:"ambient"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient struct {
	Level   types.Int64  `tfsdk:"level"`
	Quality types.String `tfsdk:"quality"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPm25 struct {
	Concentration types.Int64  `tfsdk:"concentration"`
	Quality       types.String `tfsdk:"quality"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPowerFactor struct {
	Percentage types.Int64 `tfsdk:"percentage"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdRealPower struct {
	Draw types.Float64 `tfsdk:"draw"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTemperature struct {
	Celsius    types.Float64 `tfsdk:"celsius"`
	Fahrenheit types.Float64 `tfsdk:"fahrenheit"`
	Quality    types.String  `tfsdk:"quality"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTvoc struct {
	Concentration types.Int64  `tfsdk:"concentration"`
	Quality       types.String `tfsdk:"quality"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdUpstreamPower struct {
	OutageDetected types.Bool `tfsdk:"outage_detected"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdVoltage struct {
	Level types.Float64 `tfsdk:"level"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdWater struct {
	Present types.Bool `tfsdk:"present"`
}

type ResponseSensorGetNetworkSensorAlertsProfileRecipients struct {
	Emails        types.List `tfsdk:"emails"`
	HTTPServerIDs types.List `tfsdk:"http_server_ids"`
	SmsNumbers    types.List `tfsdk:"sms_numbers"`
}

type ResponseSensorGetNetworkSensorAlertsProfileSchedule struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// ToBody
func ResponseSensorGetNetworkSensorAlertsProfilesItemsToBody(state NetworksSensorAlertsProfiles, response *merakigosdk.ResponseSensorGetNetworkSensorAlertsProfiles) NetworksSensorAlertsProfiles {
	var items []ResponseItemSensorGetNetworkSensorAlertsProfiles
	for _, item := range *response {
		itemState := ResponseItemSensorGetNetworkSensorAlertsProfiles{
			Conditions: func() *[]ResponseItemSensorGetNetworkSensorAlertsProfilesConditions {
				if item.Conditions != nil {
					result := make([]ResponseItemSensorGetNetworkSensorAlertsProfilesConditions, len(*item.Conditions))
					for i, conditions := range *item.Conditions {
						result[i] = ResponseItemSensorGetNetworkSensorAlertsProfilesConditions{
							Direction: types.StringValue(conditions.Direction),
							Duration: func() types.Int64 {
								if conditions.Duration != nil {
									return types.Int64Value(int64(*conditions.Duration))
								}
								return types.Int64{}
							}(),
							Metric: types.StringValue(conditions.Metric),
							Threshold: func() *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThreshold {
								if conditions.Threshold != nil {
									return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThreshold{
										ApparentPower: func() *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdApparentPower {
											if conditions.Threshold.ApparentPower != nil {
												return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdApparentPower{
													Draw: func() types.Float64 {
														if conditions.Threshold.ApparentPower.Draw != nil {
															return types.Float64Value(float64(*conditions.Threshold.ApparentPower.Draw))
														}
														return types.Float64{}
													}(),
												}
											}
											return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdApparentPower{}
										}(),
										Current: func() *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdCurrent {
											if conditions.Threshold.Current != nil {
												return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdCurrent{
													Draw: func() types.Float64 {
														if conditions.Threshold.Current.Draw != nil {
															return types.Float64Value(float64(*conditions.Threshold.Current.Draw))
														}
														return types.Float64{}
													}(),
												}
											}
											return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdCurrent{}
										}(),
										Door: func() *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdDoor {
											if conditions.Threshold.Door != nil {
												return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdDoor{
													Open: func() types.Bool {
														if conditions.Threshold.Door.Open != nil {
															return types.BoolValue(*conditions.Threshold.Door.Open)
														}
														return types.Bool{}
													}(),
												}
											}
											return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdDoor{}
										}(),
										Frequency: func() *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdFrequency {
											if conditions.Threshold.Frequency != nil {
												return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdFrequency{
													Level: func() types.Float64 {
														if conditions.Threshold.Frequency.Level != nil {
															return types.Float64Value(float64(*conditions.Threshold.Frequency.Level))
														}
														return types.Float64{}
													}(),
												}
											}
											return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdFrequency{}
										}(),
										Humidity: func() *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdHumidity {
											if conditions.Threshold.Humidity != nil {
												return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdHumidity{
													Quality: types.StringValue(conditions.Threshold.Humidity.Quality),
													RelativePercentage: func() types.Int64 {
														if conditions.Threshold.Humidity.RelativePercentage != nil {
															return types.Int64Value(int64(*conditions.Threshold.Humidity.RelativePercentage))
														}
														return types.Int64{}
													}(),
												}
											}
											return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdHumidity{}
										}(),
										IndoorAirQuality: func() *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdIndoorAirQuality {
											if conditions.Threshold.IndoorAirQuality != nil {
												return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdIndoorAirQuality{
													Quality: types.StringValue(conditions.Threshold.IndoorAirQuality.Quality),
													Score: func() types.Int64 {
														if conditions.Threshold.IndoorAirQuality.Score != nil {
															return types.Int64Value(int64(*conditions.Threshold.IndoorAirQuality.Score))
														}
														return types.Int64{}
													}(),
												}
											}
											return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdIndoorAirQuality{}
										}(),
										Noise: func() *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdNoise {
											if conditions.Threshold.Noise != nil {
												return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdNoise{
													Ambient: func() *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdNoiseAmbient {
														if conditions.Threshold.Noise.Ambient != nil {
															return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdNoiseAmbient{
																Level: func() types.Int64 {
																	if conditions.Threshold.Noise.Ambient.Level != nil {
																		return types.Int64Value(int64(*conditions.Threshold.Noise.Ambient.Level))
																	}
																	return types.Int64{}
																}(),
																Quality: types.StringValue(conditions.Threshold.Noise.Ambient.Quality),
															}
														}
														return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdNoiseAmbient{}
													}(),
												}
											}
											return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdNoise{}
										}(),
										Pm25: func() *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdPm25 {
											if conditions.Threshold.Pm25 != nil {
												return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdPm25{
													Concentration: func() types.Int64 {
														if conditions.Threshold.Pm25.Concentration != nil {
															return types.Int64Value(int64(*conditions.Threshold.Pm25.Concentration))
														}
														return types.Int64{}
													}(),
													Quality: types.StringValue(conditions.Threshold.Pm25.Quality),
												}
											}
											return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdPm25{}
										}(),
										PowerFactor: func() *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdPowerFactor {
											if conditions.Threshold.PowerFactor != nil {
												return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdPowerFactor{
													Percentage: func() types.Int64 {
														if conditions.Threshold.PowerFactor.Percentage != nil {
															return types.Int64Value(int64(*conditions.Threshold.PowerFactor.Percentage))
														}
														return types.Int64{}
													}(),
												}
											}
											return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdPowerFactor{}
										}(),
										RealPower: func() *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdRealPower {
											if conditions.Threshold.RealPower != nil {
												return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdRealPower{
													Draw: func() types.Float64 {
														if conditions.Threshold.RealPower.Draw != nil {
															return types.Float64Value(float64(*conditions.Threshold.RealPower.Draw))
														}
														return types.Float64{}
													}(),
												}
											}
											return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdRealPower{}
										}(),
										Temperature: func() *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdTemperature {
											if conditions.Threshold.Temperature != nil {
												return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdTemperature{
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
											return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdTemperature{}
										}(),
										Tvoc: func() *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdTvoc {
											if conditions.Threshold.Tvoc != nil {
												return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdTvoc{
													Concentration: func() types.Int64 {
														if conditions.Threshold.Tvoc.Concentration != nil {
															return types.Int64Value(int64(*conditions.Threshold.Tvoc.Concentration))
														}
														return types.Int64{}
													}(),
													Quality: types.StringValue(conditions.Threshold.Tvoc.Quality),
												}
											}
											return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdTvoc{}
										}(),
										UpstreamPower: func() *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdUpstreamPower {
											if conditions.Threshold.UpstreamPower != nil {
												return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdUpstreamPower{
													OutageDetected: func() types.Bool {
														if conditions.Threshold.UpstreamPower.OutageDetected != nil {
															return types.BoolValue(*conditions.Threshold.UpstreamPower.OutageDetected)
														}
														return types.Bool{}
													}(),
												}
											}
											return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdUpstreamPower{}
										}(),
										Voltage: func() *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdVoltage {
											if conditions.Threshold.Voltage != nil {
												return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdVoltage{
													Level: func() types.Float64 {
														if conditions.Threshold.Voltage.Level != nil {
															return types.Float64Value(float64(*conditions.Threshold.Voltage.Level))
														}
														return types.Float64{}
													}(),
												}
											}
											return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdVoltage{}
										}(),
										Water: func() *ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdWater {
											if conditions.Threshold.Water != nil {
												return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdWater{
													Present: func() types.Bool {
														if conditions.Threshold.Water.Present != nil {
															return types.BoolValue(*conditions.Threshold.Water.Present)
														}
														return types.Bool{}
													}(),
												}
											}
											return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThresholdWater{}
										}(),
									}
								}
								return &ResponseItemSensorGetNetworkSensorAlertsProfilesConditionsThreshold{}
							}(),
						}
					}
					return &result
				}
				return &[]ResponseItemSensorGetNetworkSensorAlertsProfilesConditions{}
			}(),
			Name:      types.StringValue(item.Name),
			ProfileID: types.StringValue(item.ProfileID),
			Recipients: func() *ResponseItemSensorGetNetworkSensorAlertsProfilesRecipients {
				if item.Recipients != nil {
					return &ResponseItemSensorGetNetworkSensorAlertsProfilesRecipients{
						Emails:        StringSliceToList(item.Recipients.Emails),
						HTTPServerIDs: StringSliceToList(item.Recipients.HTTPServerIDs),
						SmsNumbers:    StringSliceToList(item.Recipients.SmsNumbers),
					}
				}
				return &ResponseItemSensorGetNetworkSensorAlertsProfilesRecipients{}
			}(),
			Schedule: func() *ResponseItemSensorGetNetworkSensorAlertsProfilesSchedule {
				if item.Schedule != nil {
					return &ResponseItemSensorGetNetworkSensorAlertsProfilesSchedule{
						ID:   types.StringValue(item.Schedule.ID),
						Name: types.StringValue(item.Schedule.Name),
					}
				}
				return &ResponseItemSensorGetNetworkSensorAlertsProfilesSchedule{}
			}(),
			Serials: StringSliceToList(item.Serials),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseSensorGetNetworkSensorAlertsProfileItemToBody(state NetworksSensorAlertsProfiles, response *merakigosdk.ResponseSensorGetNetworkSensorAlertsProfile) NetworksSensorAlertsProfiles {
	itemState := ResponseSensorGetNetworkSensorAlertsProfile{
		Conditions: func() *[]ResponseSensorGetNetworkSensorAlertsProfileConditions {
			if response.Conditions != nil {
				result := make([]ResponseSensorGetNetworkSensorAlertsProfileConditions, len(*response.Conditions))
				for i, conditions := range *response.Conditions {
					result[i] = ResponseSensorGetNetworkSensorAlertsProfileConditions{
						Direction: types.StringValue(conditions.Direction),
						Duration: func() types.Int64 {
							if conditions.Duration != nil {
								return types.Int64Value(int64(*conditions.Duration))
							}
							return types.Int64{}
						}(),
						Metric: types.StringValue(conditions.Metric),
						Threshold: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThreshold {
							if conditions.Threshold != nil {
								return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThreshold{
									ApparentPower: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdApparentPower {
										if conditions.Threshold.ApparentPower != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdApparentPower{
												Draw: func() types.Float64 {
													if conditions.Threshold.ApparentPower.Draw != nil {
														return types.Float64Value(float64(*conditions.Threshold.ApparentPower.Draw))
													}
													return types.Float64{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdApparentPower{}
									}(),
									Current: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdCurrent {
										if conditions.Threshold.Current != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdCurrent{
												Draw: func() types.Float64 {
													if conditions.Threshold.Current.Draw != nil {
														return types.Float64Value(float64(*conditions.Threshold.Current.Draw))
													}
													return types.Float64{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdCurrent{}
									}(),
									Door: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdDoor {
										if conditions.Threshold.Door != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdDoor{
												Open: func() types.Bool {
													if conditions.Threshold.Door.Open != nil {
														return types.BoolValue(*conditions.Threshold.Door.Open)
													}
													return types.Bool{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdDoor{}
									}(),
									Frequency: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdFrequency {
										if conditions.Threshold.Frequency != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdFrequency{
												Level: func() types.Float64 {
													if conditions.Threshold.Frequency.Level != nil {
														return types.Float64Value(float64(*conditions.Threshold.Frequency.Level))
													}
													return types.Float64{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdFrequency{}
									}(),
									Humidity: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdHumidity {
										if conditions.Threshold.Humidity != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdHumidity{
												Quality: types.StringValue(conditions.Threshold.Humidity.Quality),
												RelativePercentage: func() types.Int64 {
													if conditions.Threshold.Humidity.RelativePercentage != nil {
														return types.Int64Value(int64(*conditions.Threshold.Humidity.RelativePercentage))
													}
													return types.Int64{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdHumidity{}
									}(),
									IndoorAirQuality: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality {
										if conditions.Threshold.IndoorAirQuality != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality{
												Quality: types.StringValue(conditions.Threshold.IndoorAirQuality.Quality),
												Score: func() types.Int64 {
													if conditions.Threshold.IndoorAirQuality.Score != nil {
														return types.Int64Value(int64(*conditions.Threshold.IndoorAirQuality.Score))
													}
													return types.Int64{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality{}
									}(),
									Noise: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoise {
										if conditions.Threshold.Noise != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoise{
												Ambient: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient {
													if conditions.Threshold.Noise.Ambient != nil {
														return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient{
															Level: func() types.Int64 {
																if conditions.Threshold.Noise.Ambient.Level != nil {
																	return types.Int64Value(int64(*conditions.Threshold.Noise.Ambient.Level))
																}
																return types.Int64{}
															}(),
															Quality: types.StringValue(conditions.Threshold.Noise.Ambient.Quality),
														}
													}
													return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoiseAmbient{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoise{}
									}(),
									Pm25: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPm25 {
										if conditions.Threshold.Pm25 != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPm25{
												Concentration: func() types.Int64 {
													if conditions.Threshold.Pm25.Concentration != nil {
														return types.Int64Value(int64(*conditions.Threshold.Pm25.Concentration))
													}
													return types.Int64{}
												}(),
												Quality: types.StringValue(conditions.Threshold.Pm25.Quality),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPm25{}
									}(),
									PowerFactor: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPowerFactor {
										if conditions.Threshold.PowerFactor != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPowerFactor{
												Percentage: func() types.Int64 {
													if conditions.Threshold.PowerFactor.Percentage != nil {
														return types.Int64Value(int64(*conditions.Threshold.PowerFactor.Percentage))
													}
													return types.Int64{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPowerFactor{}
									}(),
									RealPower: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdRealPower {
										if conditions.Threshold.RealPower != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdRealPower{
												Draw: func() types.Float64 {
													if conditions.Threshold.RealPower.Draw != nil {
														return types.Float64Value(float64(*conditions.Threshold.RealPower.Draw))
													}
													return types.Float64{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdRealPower{}
									}(),
									Temperature: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTemperature {
										if conditions.Threshold.Temperature != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTemperature{
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
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTemperature{}
									}(),
									Tvoc: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTvoc {
										if conditions.Threshold.Tvoc != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTvoc{
												Concentration: func() types.Int64 {
													if conditions.Threshold.Tvoc.Concentration != nil {
														return types.Int64Value(int64(*conditions.Threshold.Tvoc.Concentration))
													}
													return types.Int64{}
												}(),
												Quality: types.StringValue(conditions.Threshold.Tvoc.Quality),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTvoc{}
									}(),
									UpstreamPower: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdUpstreamPower {
										if conditions.Threshold.UpstreamPower != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdUpstreamPower{
												OutageDetected: func() types.Bool {
													if conditions.Threshold.UpstreamPower.OutageDetected != nil {
														return types.BoolValue(*conditions.Threshold.UpstreamPower.OutageDetected)
													}
													return types.Bool{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdUpstreamPower{}
									}(),
									Voltage: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdVoltage {
										if conditions.Threshold.Voltage != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdVoltage{
												Level: func() types.Float64 {
													if conditions.Threshold.Voltage.Level != nil {
														return types.Float64Value(float64(*conditions.Threshold.Voltage.Level))
													}
													return types.Float64{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdVoltage{}
									}(),
									Water: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdWater {
										if conditions.Threshold.Water != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdWater{
												Present: func() types.Bool {
													if conditions.Threshold.Water.Present != nil {
														return types.BoolValue(*conditions.Threshold.Water.Present)
													}
													return types.Bool{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdWater{}
									}(),
								}
							}
							return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThreshold{}
						}(),
					}
				}
				return &result
			}
			return &[]ResponseSensorGetNetworkSensorAlertsProfileConditions{}
		}(),
		Name:      types.StringValue(response.Name),
		ProfileID: types.StringValue(response.ProfileID),
		Recipients: func() *ResponseSensorGetNetworkSensorAlertsProfileRecipients {
			if response.Recipients != nil {
				return &ResponseSensorGetNetworkSensorAlertsProfileRecipients{
					Emails:        StringSliceToList(response.Recipients.Emails),
					HTTPServerIDs: StringSliceToList(response.Recipients.HTTPServerIDs),
					SmsNumbers:    StringSliceToList(response.Recipients.SmsNumbers),
				}
			}
			return &ResponseSensorGetNetworkSensorAlertsProfileRecipients{}
		}(),
		Schedule: func() *ResponseSensorGetNetworkSensorAlertsProfileSchedule {
			if response.Schedule != nil {
				return &ResponseSensorGetNetworkSensorAlertsProfileSchedule{
					ID:   types.StringValue(response.Schedule.ID),
					Name: types.StringValue(response.Schedule.Name),
				}
			}
			return &ResponseSensorGetNetworkSensorAlertsProfileSchedule{}
		}(),
		Serials: StringSliceToList(response.Serials),
	}
	state.Item = &itemState
	return state
}
