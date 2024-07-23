package provider

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"direction": schema.StringAttribute{
							MarkdownDescription: `If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature, humidity, realPower, apparentPower, powerFactor, voltage, current, and frequency thresholds.`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"above",
									"below",
								),
							},
						},
						"duration": schema.Int64Attribute{
							MarkdownDescription: `Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, 1 hour, 2 hours, 4 hours, and 8 hours. Default is 0.`,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"metric": schema.StringAttribute{
							MarkdownDescription: `The type of sensor metric that will be monitored for changes. Available metrics are apparentPower, co2, current, door, frequency, humidity, indoorAirQuality, noise, pm25, powerFactor, realPower, temperature, tvoc, upstreamPower, voltage, and water.`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"threshold": schema.SingleNestedAttribute{
							MarkdownDescription: `Threshold for sensor readings that will cause an alert to be sent. This object should contain a single property key matching the condition's 'metric' value.`,
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"apparent_power": schema.SingleNestedAttribute{
									MarkdownDescription: `Apparent power threshold. 'draw' must be provided.`,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"draw": schema.Float64Attribute{
											MarkdownDescription: `Alerting threshold in volt-amps. Must be between 0 and 3750.`,
											Optional:            true,
											PlanModifiers: []planmodifier.Float64{
												float64planmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"current": schema.SingleNestedAttribute{
									MarkdownDescription: `Electrical current threshold. 'level' must be provided.`,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"draw": schema.Float64Attribute{
											MarkdownDescription: `Alerting threshold in amps. Must be between 0 and 15.`,
											Optional:            true,
											PlanModifiers: []planmodifier.Float64{
												float64planmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"door": schema.SingleNestedAttribute{
									MarkdownDescription: `Door open threshold. 'open' must be provided and set to true.`,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"open": schema.BoolAttribute{
											MarkdownDescription: `Alerting threshold for a door open event. Must be set to true.`,
											Optional:            true,
											PlanModifiers: []planmodifier.Bool{
												boolplanmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"frequency": schema.SingleNestedAttribute{
									MarkdownDescription: `Electrical frequency threshold. 'level' must be provided.`,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"level": schema.Float64Attribute{
											MarkdownDescription: `Alerting threshold in hertz. Must be between 0 and 60.`,
											Optional:            true,
											PlanModifiers: []planmodifier.Float64{
												float64planmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"humidity": schema.SingleNestedAttribute{
									MarkdownDescription: `Humidity threshold. One of 'relativePercentage' or 'quality' must be provided.`,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"quality": schema.StringAttribute{
											MarkdownDescription: `Alerting threshold as a qualitative humidity level.`,
											Optional:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.UseStateForUnknown(),
											},
											Validators: []validator.String{
												stringvalidator.OneOf(
													"fair",
													"good",
													"inadequate",
													"poor",
												),
											},
										},
										"relative_percentage": schema.Int64Attribute{
											MarkdownDescription: `Alerting threshold in %RH.`,
											Optional:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"indoor_air_quality": schema.SingleNestedAttribute{
									MarkdownDescription: `Indoor air quality score threshold. One of 'score' or 'quality' must be provided.`,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"quality": schema.StringAttribute{
											MarkdownDescription: `Alerting threshold as a qualitative indoor air quality level.`,
											Optional:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.UseStateForUnknown(),
											},
											Validators: []validator.String{
												stringvalidator.OneOf(
													"fair",
													"good",
													"inadequate",
													"poor",
												),
											},
										},
										"score": schema.Int64Attribute{
											MarkdownDescription: `Alerting threshold as indoor air quality score.`,
											Optional:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"noise": schema.SingleNestedAttribute{
									MarkdownDescription: `Noise threshold. 'ambient' must be provided.`,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"ambient": schema.SingleNestedAttribute{
											MarkdownDescription: `Ambient noise threshold. One of 'level' or 'quality' must be provided.`,
											Optional:            true,
											PlanModifiers: []planmodifier.Object{
												objectplanmodifier.UseStateForUnknown(),
											},
											Attributes: map[string]schema.Attribute{

												"level": schema.Int64Attribute{
													MarkdownDescription: `Alerting threshold as adjusted decibels.`,
													Optional:            true,
													PlanModifiers: []planmodifier.Int64{
														int64planmodifier.UseStateForUnknown(),
													},
												},
												"quality": schema.StringAttribute{
													MarkdownDescription: `Alerting threshold as a qualitative ambient noise level.`,
													Optional:            true,
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
													Validators: []validator.String{
														stringvalidator.OneOf(
															"fair",
															"good",
															"inadequate",
															"poor",
														),
													},
												},
											},
										},
									},
								},
								"pm25": schema.SingleNestedAttribute{
									MarkdownDescription: `PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided.`,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"concentration": schema.Int64Attribute{
											MarkdownDescription: `Alerting threshold as PM2.5 parts per million.`,
											Optional:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.UseStateForUnknown(),
											},
										},
										"quality": schema.StringAttribute{
											MarkdownDescription: `Alerting threshold as a qualitative PM2.5 level.`,
											Optional:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.UseStateForUnknown(),
											},
											Validators: []validator.String{
												stringvalidator.OneOf(
													"fair",
													"good",
													"inadequate",
													"poor",
												),
											},
										},
									},
								},
								"power_factor": schema.SingleNestedAttribute{
									MarkdownDescription: `Power factor threshold. 'percentage' must be provided.`,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"percentage": schema.Int64Attribute{
											MarkdownDescription: `Alerting threshold as the ratio of active power to apparent power. Must be between 0 and 100.`,
											Optional:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"real_power": schema.SingleNestedAttribute{
									MarkdownDescription: `Real power threshold. 'draw' must be provided.`,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"draw": schema.Float64Attribute{
											MarkdownDescription: `Alerting threshold in watts. Must be between 0 and 3750.`,
											Optional:            true,
											PlanModifiers: []planmodifier.Float64{
												float64planmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"temperature": schema.SingleNestedAttribute{
									MarkdownDescription: `Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided.`,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"celsius": schema.Float64Attribute{
											MarkdownDescription: `Alerting threshold in degrees Celsius.`,
											Optional:            true,
											PlanModifiers: []planmodifier.Float64{
												float64planmodifier.UseStateForUnknown(),
											},
										},
										"fahrenheit": schema.Float64Attribute{
											MarkdownDescription: `Alerting threshold in degrees Fahrenheit.`,
											Optional:            true,
											PlanModifiers: []planmodifier.Float64{
												float64planmodifier.UseStateForUnknown(),
											},
										},
										"quality": schema.StringAttribute{
											MarkdownDescription: `Alerting threshold as a qualitative temperature level.`,
											Optional:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.UseStateForUnknown(),
											},
											Validators: []validator.String{
												stringvalidator.OneOf(
													"fair",
													"good",
													"inadequate",
													"poor",
												),
											},
										},
									},
								},
								"tvoc": schema.SingleNestedAttribute{
									MarkdownDescription: `TVOC concentration threshold. One of 'concentration' or 'quality' must be provided.`,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"concentration": schema.Int64Attribute{
											MarkdownDescription: `Alerting threshold as TVOC micrograms per cubic meter.`,
											Optional:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.UseStateForUnknown(),
											},
										},
										"quality": schema.StringAttribute{
											MarkdownDescription: `Alerting threshold as a qualitative TVOC level.`,
											Optional:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.UseStateForUnknown(),
											},
											Validators: []validator.String{
												stringvalidator.OneOf(
													"fair",
													"good",
													"inadequate",
													"poor",
												),
											},
										},
									},
								},
								"upstream_power": schema.SingleNestedAttribute{
									MarkdownDescription: `Upstream power threshold. 'outageDetected' must be provided and set to true.`,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"outage_detected": schema.BoolAttribute{
											MarkdownDescription: `Alerting threshold for an upstream power event. Must be set to true.`,
											Optional:            true,
											PlanModifiers: []planmodifier.Bool{
												boolplanmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"voltage": schema.SingleNestedAttribute{
									MarkdownDescription: `Voltage threshold. 'level' must be provided.`,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"level": schema.Float64Attribute{
											MarkdownDescription: `Alerting threshold in volts. Must be between 0 and 250.`,
											Optional:            true,
											PlanModifiers: []planmodifier.Float64{
												float64planmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"water": schema.SingleNestedAttribute{
									MarkdownDescription: `Water detection threshold. 'present' must be provided and set to true.`,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"present": schema.BoolAttribute{
											MarkdownDescription: `Alerting threshold for a water detection event. Must be set to true.`,
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
			"conditions_response": schema.SetNestedAttribute{
				MarkdownDescription: `List of conditions that will cause the profile to send an alert.`,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"direction": schema.StringAttribute{
							MarkdownDescription: `If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature, humidity, realPower, apparentPower, powerFactor, voltage, current, and frequency thresholds.`,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"above",
									"below",
								),
							},
						},
						"duration": schema.Int64Attribute{
							MarkdownDescription: `Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, 1 hour, 2 hours, 4 hours, and 8 hours. Default is 0.`,
							Computed:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"metric": schema.StringAttribute{
							MarkdownDescription: `The type of sensor metric that will be monitored for changes. Available metrics are apparentPower, co2, current, door, frequency, humidity, indoorAirQuality, noise, pm25, powerFactor, realPower, temperature, tvoc, upstreamPower, voltage, and water.`,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"threshold": schema.SingleNestedAttribute{
							MarkdownDescription: `Threshold for sensor readings that will cause an alert to be sent. This object should contain a single property key matching the condition's 'metric' value.`,
							Computed:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"apparent_power": schema.SingleNestedAttribute{
									MarkdownDescription: `Apparent power threshold. 'draw' must be provided.`,
									Computed:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"draw": schema.Float64Attribute{
											MarkdownDescription: `Alerting threshold in volt-amps. Must be between 0 and 3750.`,
											Computed:            true,
											PlanModifiers: []planmodifier.Float64{
												float64planmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"current": schema.SingleNestedAttribute{
									MarkdownDescription: `Electrical current threshold. 'level' must be provided.`,
									Computed:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"draw": schema.Float64Attribute{
											MarkdownDescription: `Alerting threshold in amps. Must be between 0 and 15.`,
											Computed:            true,
											PlanModifiers: []planmodifier.Float64{
												float64planmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"door": schema.SingleNestedAttribute{
									MarkdownDescription: `Door open threshold. 'open' must be provided and set to true.`,
									Computed:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"open": schema.BoolAttribute{
											MarkdownDescription: `Alerting threshold for a door open event. Must be set to true.`,
											Computed:            true,
											PlanModifiers: []planmodifier.Bool{
												boolplanmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"frequency": schema.SingleNestedAttribute{
									MarkdownDescription: `Electrical frequency threshold. 'level' must be provided.`,
									Computed:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"level": schema.Float64Attribute{
											MarkdownDescription: `Alerting threshold in hertz. Must be between 0 and 60.`,
											Computed:            true,
											PlanModifiers: []planmodifier.Float64{
												float64planmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"humidity": schema.SingleNestedAttribute{
									MarkdownDescription: `Humidity threshold. One of 'relativePercentage' or 'quality' must be provided.`,
									Computed:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"quality": schema.StringAttribute{
											MarkdownDescription: `Alerting threshold as a qualitative humidity level.`,
											Computed:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.UseStateForUnknown(),
											},
											Validators: []validator.String{
												stringvalidator.OneOf(
													"fair",
													"good",
													"inadequate",
													"poor",
												),
											},
										},
										"relative_percentage": schema.Int64Attribute{
											MarkdownDescription: `Alerting threshold in %RH.`,
											Computed:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"indoor_air_quality": schema.SingleNestedAttribute{
									MarkdownDescription: `Indoor air quality score threshold. One of 'score' or 'quality' must be provided.`,
									Computed:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"quality": schema.StringAttribute{
											MarkdownDescription: `Alerting threshold as a qualitative indoor air quality level.`,
											Computed:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.UseStateForUnknown(),
											},
											Validators: []validator.String{
												stringvalidator.OneOf(
													"fair",
													"good",
													"inadequate",
													"poor",
												),
											},
										},
										"score": schema.Int64Attribute{
											MarkdownDescription: `Alerting threshold as indoor air quality score.`,
											Computed:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"noise": schema.SingleNestedAttribute{
									MarkdownDescription: `Noise threshold. 'ambient' must be provided.`,
									Computed:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"ambient": schema.SingleNestedAttribute{
											MarkdownDescription: `Ambient noise threshold. One of 'level' or 'quality' must be provided.`,
											Computed:            true,
											PlanModifiers: []planmodifier.Object{
												objectplanmodifier.UseStateForUnknown(),
											},
											Attributes: map[string]schema.Attribute{

												"level": schema.Int64Attribute{
													MarkdownDescription: `Alerting threshold as adjusted decibels.`,
													Computed:            true,
													PlanModifiers: []planmodifier.Int64{
														int64planmodifier.UseStateForUnknown(),
													},
												},
												"quality": schema.StringAttribute{
													MarkdownDescription: `Alerting threshold as a qualitative ambient noise level.`,
													Computed:            true,
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
													Validators: []validator.String{
														stringvalidator.OneOf(
															"fair",
															"good",
															"inadequate",
															"poor",
														),
													},
												},
											},
										},
									},
								},
								"pm25": schema.SingleNestedAttribute{
									MarkdownDescription: `PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided.`,
									Computed:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"concentration": schema.Int64Attribute{
											MarkdownDescription: `Alerting threshold as PM2.5 parts per million.`,
											Computed:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.UseStateForUnknown(),
											},
										},
										"quality": schema.StringAttribute{
											MarkdownDescription: `Alerting threshold as a qualitative PM2.5 level.`,
											Computed:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.UseStateForUnknown(),
											},
											Validators: []validator.String{
												stringvalidator.OneOf(
													"fair",
													"good",
													"inadequate",
													"poor",
												),
											},
										},
									},
								},
								"power_factor": schema.SingleNestedAttribute{
									MarkdownDescription: `Power factor threshold. 'percentage' must be provided.`,
									Computed:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"percentage": schema.Int64Attribute{
											MarkdownDescription: `Alerting threshold as the ratio of active power to apparent power. Must be between 0 and 100.`,
											Computed:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"real_power": schema.SingleNestedAttribute{
									MarkdownDescription: `Real power threshold. 'draw' must be provided.`,
									Computed:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"draw": schema.Float64Attribute{
											MarkdownDescription: `Alerting threshold in watts. Must be between 0 and 3750.`,
											Computed:            true,
											PlanModifiers: []planmodifier.Float64{
												float64planmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"temperature": schema.SingleNestedAttribute{
									MarkdownDescription: `Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided.`,
									Computed:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"celsius": schema.Float64Attribute{
											MarkdownDescription: `Alerting threshold in degrees Celsius.`,
											Computed:            true,
											PlanModifiers: []planmodifier.Float64{
												float64planmodifier.UseStateForUnknown(),
											},
										},
										"fahrenheit": schema.Float64Attribute{
											MarkdownDescription: `Alerting threshold in degrees Fahrenheit.`,
											Computed:            true,
											PlanModifiers: []planmodifier.Float64{
												float64planmodifier.UseStateForUnknown(),
											},
										},
										"quality": schema.StringAttribute{
											MarkdownDescription: `Alerting threshold as a qualitative temperature level.`,
											Computed:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.UseStateForUnknown(),
											},
											Validators: []validator.String{
												stringvalidator.OneOf(
													"fair",
													"good",
													"inadequate",
													"poor",
												),
											},
										},
									},
								},
								"tvoc": schema.SingleNestedAttribute{
									MarkdownDescription: `TVOC concentration threshold. One of 'concentration' or 'quality' must be provided.`,
									Computed:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"concentration": schema.Int64Attribute{
											MarkdownDescription: `Alerting threshold as TVOC micrograms per cubic meter.`,
											Computed:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.UseStateForUnknown(),
											},
										},
										"quality": schema.StringAttribute{
											MarkdownDescription: `Alerting threshold as a qualitative TVOC level.`,
											Computed:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.UseStateForUnknown(),
											},
											Validators: []validator.String{
												stringvalidator.OneOf(
													"fair",
													"good",
													"inadequate",
													"poor",
												),
											},
										},
									},
								},
								"upstream_power": schema.SingleNestedAttribute{
									MarkdownDescription: `Upstream power threshold. 'outageDetected' must be provided and set to true.`,
									Computed:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"outage_detected": schema.BoolAttribute{
											MarkdownDescription: `Alerting threshold for an upstream power event. Must be set to true.`,
											Computed:            true,
											PlanModifiers: []planmodifier.Bool{
												boolplanmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"voltage": schema.SingleNestedAttribute{
									MarkdownDescription: `Voltage threshold. 'level' must be provided.`,
									Computed:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"level": schema.Float64Attribute{
											MarkdownDescription: `Alerting threshold in volts. Must be between 0 and 250.`,
											Computed:            true,
											PlanModifiers: []planmodifier.Float64{
												float64planmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"water": schema.SingleNestedAttribute{
									MarkdownDescription: `Water detection threshold. 'present' must be provided and set to true.`,
									Computed:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"present": schema.BoolAttribute{
											MarkdownDescription: `Alerting threshold for a water detection event. Must be set to true.`,
											Computed:            true,
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
				MarkdownDescription: `List of recipients that will receive the alert.`,
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
	vvName := data.Name.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Sensor.GetNetworkSensorAlertsProfiles(vvNetworkID)
	//Have Create
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkSensorAlertsProfiles",
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
			vvID, ok := result2["ProfileID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter ID",
					"Fail Parsing ID",
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
		vvID, ok := result2["ProfileID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter ID",
				err.Error(),
			)
			return
		}
		// time.Sleep(5 * time.Second)
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
	vvID := data.ProfileID.ValueString()
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
	//entro aqui 2
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
	vvID := data.ProfileID.ValueString()
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
	vvID := state.ProfileID.ValueString()
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
	NetworkID          types.String                                               `tfsdk:"network_id"`
	ID                 types.String                                               `tfsdk:"id"`
	Conditions         *[]ResponseSensorGetNetworkSensorAlertsProfileConditionsRs `tfsdk:"conditions"`
	ConditionsResponse *[]ResponseSensorGetNetworkSensorAlertsProfileConditionsRs `tfsdk:"conditions_response"`
	Name               types.String                                               `tfsdk:"name"`
	ProfileID          types.String                                               `tfsdk:"profile_id"`
	Recipients         *ResponseSensorGetNetworkSensorAlertsProfileRecipientsRs   `tfsdk:"recipients"`
	Schedule           *ResponseSensorGetNetworkSensorAlertsProfileScheduleRs     `tfsdk:"schedule"`
	Serials            types.Set                                                  `tfsdk:"serials"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsRs struct {
	Direction types.String                                                      `tfsdk:"direction"`
	Duration  types.Int64                                                       `tfsdk:"duration"`
	Metric    types.String                                                      `tfsdk:"metric"`
	Threshold *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdRs `tfsdk:"threshold"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdRs struct {
	ApparentPower    *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdApparentPowerRs    `tfsdk:"apparent_power"`
	Current          *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdCurrentRs          `tfsdk:"current"`
	Door             *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdDoorRs             `tfsdk:"door"`
	Frequency        *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdFrequencyRs        `tfsdk:"frequency"`
	Humidity         *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdHumidityRs         `tfsdk:"humidity"`
	IndoorAirQuality *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdIndoorAirQualityRs `tfsdk:"indoor_air_quality"`
	Noise            *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdNoiseRs            `tfsdk:"noise"`
	Pm25             *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPm25Rs             `tfsdk:"pm25"`
	PowerFactor      *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPowerFactorRs      `tfsdk:"power_factor"`
	RealPower        *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdRealPowerRs        `tfsdk:"real_power"`
	Temperature      *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTemperatureRs      `tfsdk:"temperature"`
	Tvoc             *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdTvocRs             `tfsdk:"tvoc"`
	UpstreamPower    *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdUpstreamPowerRs    `tfsdk:"upstream_power"`
	Voltage          *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdVoltageRs          `tfsdk:"voltage"`
	Water            *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdWaterRs            `tfsdk:"water"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdApparentPowerRs struct {
	Draw types.Float64 `tfsdk:"draw"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdCurrentRs struct {
	Draw types.Float64 `tfsdk:"draw"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdDoorRs struct {
	Open types.Bool `tfsdk:"open"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdFrequencyRs struct {
	Level types.Float64 `tfsdk:"level"`
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

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPowerFactorRs struct {
	Percentage types.Int64 `tfsdk:"percentage"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdRealPowerRs struct {
	Draw types.Float64 `tfsdk:"draw"`
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

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdUpstreamPowerRs struct {
	OutageDetected types.Bool `tfsdk:"outage_detected"`
}

type ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdVoltageRs struct {
	Level types.Float64 `tfsdk:"level"`
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
				var requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdApparentPower *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdApparentPower
				if rItem1.Threshold.ApparentPower != nil {
					draw := rItem1.Threshold.ApparentPower.Draw.ValueFloat64Pointer()
					requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdApparentPower = &merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdApparentPower{
						Draw: draw,
					}
				}
				var requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdCurrent *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdCurrent
				if rItem1.Threshold.Current != nil {
					draw := rItem1.Threshold.Current.Draw.ValueFloat64Pointer()
					requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdCurrent = &merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdCurrent{
						Draw: draw,
					}
				}
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
				var requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdFrequency *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdFrequency
				if rItem1.Threshold.Frequency != nil {
					level := rItem1.Threshold.Frequency.Level.ValueFloat64Pointer()
					requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdFrequency = &merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdFrequency{
						Level: level,
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
				var requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdPowerFactor *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdPowerFactor
				if rItem1.Threshold.PowerFactor != nil {
					percentage := func() *int64 {
						if !rItem1.Threshold.PowerFactor.Percentage.IsUnknown() && !rItem1.Threshold.PowerFactor.Percentage.IsNull() {
							return rItem1.Threshold.PowerFactor.Percentage.ValueInt64Pointer()
						}
						return nil
					}()
					requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdPowerFactor = &merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdPowerFactor{
						Percentage: int64ToIntPointer(percentage),
					}
				}
				var requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdRealPower *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdRealPower
				if rItem1.Threshold.RealPower != nil {
					draw := rItem1.Threshold.RealPower.Draw.ValueFloat64Pointer()
					requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdRealPower = &merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdRealPower{
						Draw: draw,
					}
				}
				var requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTemperature *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTemperature
				if rItem1.Threshold.Temperature != nil {
					// celsius := rItem1.Threshold.Temperature.Celsius.ValueFloat64Pointer()
					// fahrenheit := rItem1.Threshold.Temperature.Fahrenheit.ValueFloat64Pointer()
					quality := rItem1.Threshold.Temperature.Quality.ValueString()
					requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTemperature = &merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTemperature{
						Celsius: func() *float64 {
							if !rItem1.Threshold.Temperature.Celsius.IsUnknown() && !rItem1.Threshold.Temperature.Celsius.IsNull() {
								return rItem1.Threshold.Temperature.Celsius.ValueFloat64Pointer()
							}
							return nil
						}(),
						Fahrenheit: func() *float64 {
							if !rItem1.Threshold.Temperature.Fahrenheit.IsUnknown() && !rItem1.Threshold.Temperature.Fahrenheit.IsNull() {
								return rItem1.Threshold.Temperature.Fahrenheit.ValueFloat64Pointer()
							}
							return nil
						}(),
						Quality: quality,
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
				var requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdUpstreamPower *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdUpstreamPower
				if rItem1.Threshold.UpstreamPower != nil {
					outageDetected := func() *bool {
						if !rItem1.Threshold.UpstreamPower.OutageDetected.IsUnknown() && !rItem1.Threshold.UpstreamPower.OutageDetected.IsNull() {
							return rItem1.Threshold.UpstreamPower.OutageDetected.ValueBoolPointer()
						}
						return nil
					}()
					requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdUpstreamPower = &merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdUpstreamPower{
						OutageDetected: outageDetected,
					}
				}
				var requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdVoltage *merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdVoltage
				if rItem1.Threshold.Voltage != nil {
					level := rItem1.Threshold.Voltage.Level.ValueFloat64Pointer()
					requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdVoltage = &merakigosdk.RequestSensorCreateNetworkSensorAlertsProfileConditionsThresholdVoltage{
						Level: level,
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
					ApparentPower:    requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdApparentPower,
					Current:          requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdCurrent,
					Door:             requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdDoor,
					Frequency:        requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdFrequency,
					Humidity:         requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdHumidity,
					IndoorAirQuality: requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality,
					Noise:            requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdNoise,
					Pm25:             requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdPm25,
					PowerFactor:      requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdPowerFactor,
					RealPower:        requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdRealPower,
					Temperature:      requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTemperature,
					Tvoc:             requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdTvoc,
					UpstreamPower:    requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdUpstreamPower,
					Voltage:          requestSensorCreateNetworkSensorAlertsProfileConditionsThresholdVoltage,
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
		//Hoola aqui
		r.Recipients.Emails.ElementsAs(ctx, &emails, false)
		var httpServerIDs []string = nil
		//Hoola aqui
		r.Recipients.HTTPServerIDs.ElementsAs(ctx, &httpServerIDs, false)
		var smsNumbers []string = nil
		//Hoola aqui
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
				var requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdApparentPower *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdApparentPower
				if rItem1.Threshold.ApparentPower != nil {
					draw := rItem1.Threshold.ApparentPower.Draw.ValueFloat64Pointer()
					requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdApparentPower = &merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdApparentPower{
						Draw: draw,
					}
				}
				var requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdCurrent *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdCurrent
				if rItem1.Threshold.Current != nil {
					draw := rItem1.Threshold.Current.Draw.ValueFloat64Pointer()
					requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdCurrent = &merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdCurrent{
						Draw: draw,
					}
				}
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
				var requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdFrequency *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdFrequency
				if rItem1.Threshold.Frequency != nil {
					level := rItem1.Threshold.Frequency.Level.ValueFloat64Pointer()
					requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdFrequency = &merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdFrequency{
						Level: level,
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
				var requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPowerFactor *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPowerFactor
				if rItem1.Threshold.PowerFactor != nil {
					percentage := func() *int64 {
						if !rItem1.Threshold.PowerFactor.Percentage.IsUnknown() && !rItem1.Threshold.PowerFactor.Percentage.IsNull() {
							return rItem1.Threshold.PowerFactor.Percentage.ValueInt64Pointer()
						}
						return nil
					}()
					requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPowerFactor = &merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPowerFactor{
						Percentage: int64ToIntPointer(percentage),
					}
				}
				var requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdRealPower *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdRealPower
				if rItem1.Threshold.RealPower != nil {
					draw := rItem1.Threshold.RealPower.Draw.ValueFloat64Pointer()
					requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdRealPower = &merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdRealPower{
						Draw: draw,
					}
				}
				var requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTemperature *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTemperature
				if rItem1.Threshold.Temperature != nil {
					// celsius := rItem1.Threshold.Temperature.Celsius.ValueFloat64Pointer()
					// fahrenheit := rItem1.Threshold.Temperature.Fahrenheit.ValueFloat64Pointer()
					quality := rItem1.Threshold.Temperature.Quality.ValueString()
					requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTemperature = &merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTemperature{
						Celsius: func() *float64 {
							if !rItem1.Threshold.Temperature.Celsius.IsUnknown() && !rItem1.Threshold.Temperature.Celsius.IsNull() {
								return rItem1.Threshold.Temperature.Celsius.ValueFloat64Pointer()
							}
							return nil
						}(),
						Fahrenheit: func() *float64 {
							if !rItem1.Threshold.Temperature.Fahrenheit.IsUnknown() && !rItem1.Threshold.Temperature.Fahrenheit.IsNull() {
								return rItem1.Threshold.Temperature.Fahrenheit.ValueFloat64Pointer()
							}
							return nil
						}(),
						Quality: quality,
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
				var requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdUpstreamPower *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdUpstreamPower
				if rItem1.Threshold.UpstreamPower != nil {
					outageDetected := func() *bool {
						if !rItem1.Threshold.UpstreamPower.OutageDetected.IsUnknown() && !rItem1.Threshold.UpstreamPower.OutageDetected.IsNull() {
							return rItem1.Threshold.UpstreamPower.OutageDetected.ValueBoolPointer()
						}
						return nil
					}()
					requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdUpstreamPower = &merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdUpstreamPower{
						OutageDetected: outageDetected,
					}
				}
				var requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdVoltage *merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdVoltage
				if rItem1.Threshold.Voltage != nil {
					level := rItem1.Threshold.Voltage.Level.ValueFloat64Pointer()
					requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdVoltage = &merakigosdk.RequestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdVoltage{
						Level: level,
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
					ApparentPower:    requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdApparentPower,
					Current:          requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdCurrent,
					Door:             requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdDoor,
					Frequency:        requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdFrequency,
					Humidity:         requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdHumidity,
					IndoorAirQuality: requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdIndoorAirQuality,
					Noise:            requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdNoise,
					Pm25:             requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPm25,
					PowerFactor:      requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdPowerFactor,
					RealPower:        requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdRealPower,
					Temperature:      requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTemperature,
					Tvoc:             requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdTvoc,
					UpstreamPower:    requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdUpstreamPower,
					Voltage:          requestSensorUpdateNetworkSensorAlertsProfileConditionsThresholdVoltage,
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
		//Hoola aqui
		r.Recipients.Emails.ElementsAs(ctx, &emails, false)
		var httpServerIDs []string = nil
		//Hoola aqui
		r.Recipients.HTTPServerIDs.ElementsAs(ctx, &httpServerIDs, false)
		var smsNumbers []string = nil
		//Hoola aqui
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
		ConditionsResponse: func() *[]ResponseSensorGetNetworkSensorAlertsProfileConditionsRs {
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
									ApparentPower: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdApparentPowerRs {
										if conditions.Threshold.ApparentPower != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdApparentPowerRs{
												Draw: func() types.Float64 {
													if conditions.Threshold.ApparentPower.Draw != nil {
														return types.Float64Value(float64(*conditions.Threshold.ApparentPower.Draw))
													}
													return types.Float64{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdApparentPowerRs{}
									}(),
									Current: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdCurrentRs {
										if conditions.Threshold.Current != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdCurrentRs{
												Draw: func() types.Float64 {
													if conditions.Threshold.Current.Draw != nil {
														return types.Float64Value(float64(*conditions.Threshold.Current.Draw))
													}
													return types.Float64{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdCurrentRs{}
									}(),
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
									Frequency: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdFrequencyRs {
										if conditions.Threshold.Frequency != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdFrequencyRs{
												Level: func() types.Float64 {
													if conditions.Threshold.Frequency.Level != nil {
														return types.Float64Value(float64(*conditions.Threshold.Frequency.Level))
													}
													return types.Float64{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdFrequencyRs{}
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
									PowerFactor: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPowerFactorRs {
										if conditions.Threshold.PowerFactor != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPowerFactorRs{
												Percentage: func() types.Int64 {
													if conditions.Threshold.PowerFactor.Percentage != nil {
														return types.Int64Value(int64(*conditions.Threshold.PowerFactor.Percentage))
													}
													return types.Int64{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdPowerFactorRs{}
									}(),
									RealPower: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdRealPowerRs {
										if conditions.Threshold.RealPower != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdRealPowerRs{
												Draw: func() types.Float64 {
													if conditions.Threshold.RealPower.Draw != nil {
														return types.Float64Value(float64(*conditions.Threshold.RealPower.Draw))
													}
													return types.Float64{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdRealPowerRs{}
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
									UpstreamPower: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdUpstreamPowerRs {
										if conditions.Threshold.UpstreamPower != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdUpstreamPowerRs{
												OutageDetected: func() types.Bool {
													if conditions.Threshold.UpstreamPower.OutageDetected != nil {
														return types.BoolValue(*conditions.Threshold.UpstreamPower.OutageDetected)
													}
													return types.Bool{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdUpstreamPowerRs{}
									}(),
									Voltage: func() *ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdVoltageRs {
										if conditions.Threshold.Voltage != nil {
											return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdVoltageRs{
												Level: func() types.Float64 {
													if conditions.Threshold.Voltage.Level != nil {
														return types.Float64Value(float64(*conditions.Threshold.Voltage.Level))
													}
													return types.Float64{}
												}(),
											}
										}
										return &ResponseSensorGetNetworkSensorAlertsProfileConditionsThresholdVoltageRs{}
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
		Serials:    StringSliceToSet(response.Serials),
		Conditions: state.Conditions,
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSensorAlertsProfilesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSensorAlertsProfilesRs)
}
