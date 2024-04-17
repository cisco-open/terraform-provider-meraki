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
	_ datasource.DataSource              = &OrganizationsSensorReadingsLatestDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSensorReadingsLatestDataSource{}
)

func NewOrganizationsSensorReadingsLatestDataSource() datasource.DataSource {
	return &OrganizationsSensorReadingsLatestDataSource{}
}

type OrganizationsSensorReadingsLatestDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSensorReadingsLatestDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSensorReadingsLatestDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_sensor_readings_latest"
}

func (d *OrganizationsSensorReadingsLatestDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"metrics": schema.ListAttribute{
				MarkdownDescription: `metrics query parameter. Types of sensor readings to retrieve. If no metrics are supplied, all available types of readings will be retrieved. Allowed values are apparentPower, battery, button, co2, current, door, downstreamPower, frequency, humidity, indoorAirQuality, noise, pm25, powerFactor, realPower, remoteLockoutSwitch, temperature, tvoc, voltage, and water.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter readings by network.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.`,
				Optional:            true,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Optional parameter to filter readings by sensor.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSensorGetOrganizationSensorReadingsLatest`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"network": schema.SingleNestedAttribute{
							MarkdownDescription: `Network to which the sensor belongs.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `ID of the network.`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Name of the network.`,
									Computed:            true,
								},
							},
						},
						"readings": schema.SetNestedAttribute{
							MarkdownDescription: `Array of latest readings from the sensor. Each object represents a single reading for a single metric.`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"apparent_power": schema.SingleNestedAttribute{
										MarkdownDescription: `Reading for the 'apparentPower' metric. This will only be present if the 'metric' property equals 'apparentPower'.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"draw": schema.Float64Attribute{
												MarkdownDescription: `Apparent power reading in volt-amperes.`,
												Computed:            true,
											},
										},
									},
									"battery": schema.SingleNestedAttribute{
										MarkdownDescription: `Reading for the 'battery' metric. This will only be present if the 'metric' property equals 'battery'.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"percentage": schema.Int64Attribute{
												MarkdownDescription: `Remaining battery life.`,
												Computed:            true,
											},
										},
									},
									"button": schema.SingleNestedAttribute{
										MarkdownDescription: `Reading for the 'button' metric. This will only be present if the 'metric' property equals 'button'.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"press_type": schema.StringAttribute{
												MarkdownDescription: `Type of button press that occurred.`,
												Computed:            true,
											},
										},
									},
									"co2": schema.SingleNestedAttribute{
										MarkdownDescription: `Reading for the 'co2' metric. This will only be present if the 'metric' property equals 'co2'.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"concentration": schema.Int64Attribute{
												MarkdownDescription: `CO2 reading in parts per million.`,
												Computed:            true,
											},
										},
									},
									"current": schema.SingleNestedAttribute{
										MarkdownDescription: `Reading for the 'current' metric. This will only be present if the 'metric' property equals 'current'.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"draw": schema.Float64Attribute{
												MarkdownDescription: `Electrical current reading in amperes.`,
												Computed:            true,
											},
										},
									},
									"door": schema.SingleNestedAttribute{
										MarkdownDescription: `Reading for the 'door' metric. This will only be present if the 'metric' property equals 'door'.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"open": schema.BoolAttribute{
												MarkdownDescription: `True if the door is open.`,
												Computed:            true,
											},
										},
									},
									"downstream_power": schema.SingleNestedAttribute{
										MarkdownDescription: `Reading for the 'downstreamPower' metric. This will only be present if the 'metric' property equals 'downstreamPower'.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"enabled": schema.BoolAttribute{
												MarkdownDescription: `True if power is turned on to the device that is connected downstream of the MT40 power monitor.`,
												Computed:            true,
											},
										},
									},
									"frequency": schema.SingleNestedAttribute{
										MarkdownDescription: `Reading for the 'frequency' metric. This will only be present if the 'metric' property equals 'frequency'.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"level": schema.Float64Attribute{
												MarkdownDescription: `Electrical current frequency reading in hertz.`,
												Computed:            true,
											},
										},
									},
									"humidity": schema.SingleNestedAttribute{
										MarkdownDescription: `Reading for the 'humidity' metric. This will only be present if the 'metric' property equals 'humidity'.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"relative_percentage": schema.Int64Attribute{
												MarkdownDescription: `Humidity reading in %RH.`,
												Computed:            true,
											},
										},
									},
									"indoor_air_quality": schema.SingleNestedAttribute{
										MarkdownDescription: `Reading for the 'indoorAirQuality' metric. This will only be present if the 'metric' property equals 'indoorAirQuality'.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"score": schema.Int64Attribute{
												MarkdownDescription: `Indoor air quality score between 0 and 100.`,
												Computed:            true,
											},
										},
									},
									"metric": schema.StringAttribute{
										MarkdownDescription: `Type of sensor reading.`,
										Computed:            true,
									},
									"noise": schema.SingleNestedAttribute{
										MarkdownDescription: `Reading for the 'noise' metric. This will only be present if the 'metric' property equals 'noise'.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"ambient": schema.SingleNestedAttribute{
												MarkdownDescription: `Ambient noise reading.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"level": schema.Int64Attribute{
														MarkdownDescription: `Ambient noise reading in adjusted decibels.`,
														Computed:            true,
													},
												},
											},
										},
									},
									"pm25": schema.SingleNestedAttribute{
										MarkdownDescription: `Reading for the 'pm25' metric. This will only be present if the 'metric' property equals 'pm25'.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"concentration": schema.Int64Attribute{
												MarkdownDescription: `PM2.5 reading in micrograms per cubic meter.`,
												Computed:            true,
											},
										},
									},
									"power_factor": schema.SingleNestedAttribute{
										MarkdownDescription: `Reading for the 'powerFactor' metric. This will only be present if the 'metric' property equals 'powerFactor'.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"percentage": schema.Int64Attribute{
												MarkdownDescription: `Power factor reading as a percentage.`,
												Computed:            true,
											},
										},
									},
									"real_power": schema.SingleNestedAttribute{
										MarkdownDescription: `Reading for the 'realPower' metric. This will only be present if the 'metric' property equals 'realPower'.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"draw": schema.Float64Attribute{
												MarkdownDescription: `Real power reading in watts.`,
												Computed:            true,
											},
										},
									},
									"remote_lockout_switch": schema.SingleNestedAttribute{
										MarkdownDescription: `Reading for the 'remoteLockoutSwitch' metric. This will only be present if the 'metric' property equals 'remoteLockoutSwitch'.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"locked": schema.BoolAttribute{
												MarkdownDescription: `True if power controls are disabled via the MT40's physical remote lockout switch.`,
												Computed:            true,
											},
										},
									},
									"temperature": schema.SingleNestedAttribute{
										MarkdownDescription: `Reading for the 'temperature' metric. This will only be present if the 'metric' property equals 'temperature'.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"celsius": schema.Float64Attribute{
												MarkdownDescription: `Temperature reading in degrees Celsius.`,
												Computed:            true,
											},
											"fahrenheit": schema.Float64Attribute{
												MarkdownDescription: `Temperature reading in degrees Fahrenheit.`,
												Computed:            true,
											},
										},
									},
									"ts": schema.StringAttribute{
										MarkdownDescription: `Time at which the reading occurred, in ISO8601 format.`,
										Computed:            true,
									},
									"tvoc": schema.SingleNestedAttribute{
										MarkdownDescription: `Reading for the 'tvoc' metric. This will only be present if the 'metric' property equals 'tvoc'.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"concentration": schema.Int64Attribute{
												MarkdownDescription: `TVOC reading in micrograms per cubic meter.`,
												Computed:            true,
											},
										},
									},
									"voltage": schema.SingleNestedAttribute{
										MarkdownDescription: `Reading for the 'voltage' metric. This will only be present if the 'metric' property equals 'voltage'.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"level": schema.Float64Attribute{
												MarkdownDescription: `Voltage reading in volts.`,
												Computed:            true,
											},
										},
									},
									"water": schema.SingleNestedAttribute{
										MarkdownDescription: `Reading for the 'water' metric. This will only be present if the 'metric' property equals 'water'.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"present": schema.BoolAttribute{
												MarkdownDescription: `True if water is detected.`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `Serial number of the sensor that took the readings.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsSensorReadingsLatestDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSensorReadingsLatest OrganizationsSensorReadingsLatest
	diags := req.Config.Get(ctx, &organizationsSensorReadingsLatest)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSensorReadingsLatest")
		vvOrganizationID := organizationsSensorReadingsLatest.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationSensorReadingsLatestQueryParams{}

		queryParams1.PerPage = int(organizationsSensorReadingsLatest.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsSensorReadingsLatest.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsSensorReadingsLatest.EndingBefore.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsSensorReadingsLatest.NetworkIDs)
		queryParams1.Serials = elementsToStrings(ctx, organizationsSensorReadingsLatest.Serials)
		queryParams1.Metrics = elementsToStrings(ctx, organizationsSensorReadingsLatest.Metrics)

		response1, restyResp1, err := d.client.Sensor.GetOrganizationSensorReadingsLatest(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSensorReadingsLatest",
				err.Error(),
			)
			return
		}

		organizationsSensorReadingsLatest = ResponseSensorGetOrganizationSensorReadingsLatestItemsToBody(organizationsSensorReadingsLatest, response1)
		diags = resp.State.Set(ctx, &organizationsSensorReadingsLatest)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSensorReadingsLatest struct {
	OrganizationID types.String                                             `tfsdk:"organization_id"`
	PerPage        types.Int64                                              `tfsdk:"per_page"`
	StartingAfter  types.String                                             `tfsdk:"starting_after"`
	EndingBefore   types.String                                             `tfsdk:"ending_before"`
	NetworkIDs     types.List                                               `tfsdk:"network_ids"`
	Serials        types.List                                               `tfsdk:"serials"`
	Metrics        types.List                                               `tfsdk:"metrics"`
	Items          *[]ResponseItemSensorGetOrganizationSensorReadingsLatest `tfsdk:"items"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatest struct {
	Network  *ResponseItemSensorGetOrganizationSensorReadingsLatestNetwork    `tfsdk:"network"`
	Readings *[]ResponseItemSensorGetOrganizationSensorReadingsLatestReadings `tfsdk:"readings"`
	Serial   types.String                                                     `tfsdk:"serial"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadings struct {
	ApparentPower       *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsApparentPower       `tfsdk:"apparent_power"`
	Battery             *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsBattery             `tfsdk:"battery"`
	Button              *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsButton              `tfsdk:"button"`
	Co2                 *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsCo2                 `tfsdk:"co2"`
	Current             *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsCurrent             `tfsdk:"current"`
	Door                *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsDoor                `tfsdk:"door"`
	DownstreamPower     *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsDownstreamPower     `tfsdk:"downstream_power"`
	Frequency           *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsFrequency           `tfsdk:"frequency"`
	Humidity            *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsHumidity            `tfsdk:"humidity"`
	IndoorAirQuality    *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsIndoorAirQuality    `tfsdk:"indoor_air_quality"`
	Metric              types.String                                                                      `tfsdk:"metric"`
	Noise               *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsNoise               `tfsdk:"noise"`
	Pm25                *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsPm25                `tfsdk:"pm25"`
	PowerFactor         *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsPowerFactor         `tfsdk:"power_factor"`
	RealPower           *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsRealPower           `tfsdk:"real_power"`
	RemoteLockoutSwitch *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsRemoteLockoutSwitch `tfsdk:"remote_lockout_switch"`
	Temperature         *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsTemperature         `tfsdk:"temperature"`
	Ts                  types.String                                                                      `tfsdk:"ts"`
	Tvoc                *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsTvoc                `tfsdk:"tvoc"`
	Voltage             *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsVoltage             `tfsdk:"voltage"`
	Water               *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsWater               `tfsdk:"water"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsApparentPower struct {
	Draw types.Float64 `tfsdk:"draw"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsBattery struct {
	Percentage types.Int64 `tfsdk:"percentage"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsButton struct {
	PressType types.String `tfsdk:"press_type"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsCo2 struct {
	Concentration types.Int64 `tfsdk:"concentration"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsCurrent struct {
	Draw types.Float64 `tfsdk:"draw"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsDoor struct {
	Open types.Bool `tfsdk:"open"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsDownstreamPower struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsFrequency struct {
	Level types.Float64 `tfsdk:"level"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsHumidity struct {
	RelativePercentage types.Int64 `tfsdk:"relative_percentage"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsIndoorAirQuality struct {
	Score types.Int64 `tfsdk:"score"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsNoise struct {
	Ambient *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsNoiseAmbient `tfsdk:"ambient"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsNoiseAmbient struct {
	Level types.Int64 `tfsdk:"level"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsPm25 struct {
	Concentration types.Int64 `tfsdk:"concentration"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsPowerFactor struct {
	Percentage types.Int64 `tfsdk:"percentage"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsRealPower struct {
	Draw types.Float64 `tfsdk:"draw"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsRemoteLockoutSwitch struct {
	Locked types.Bool `tfsdk:"locked"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsTemperature struct {
	Celsius    types.Float64 `tfsdk:"celsius"`
	Fahrenheit types.Float64 `tfsdk:"fahrenheit"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsTvoc struct {
	Concentration types.Int64 `tfsdk:"concentration"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsVoltage struct {
	Level types.Float64 `tfsdk:"level"`
}

type ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsWater struct {
	Present types.Bool `tfsdk:"present"`
}

// ToBody
func ResponseSensorGetOrganizationSensorReadingsLatestItemsToBody(state OrganizationsSensorReadingsLatest, response *merakigosdk.ResponseSensorGetOrganizationSensorReadingsLatest) OrganizationsSensorReadingsLatest {
	var items []ResponseItemSensorGetOrganizationSensorReadingsLatest
	for _, item := range *response {
		itemState := ResponseItemSensorGetOrganizationSensorReadingsLatest{
			Network: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestNetwork {
				if item.Network != nil {
					return &ResponseItemSensorGetOrganizationSensorReadingsLatestNetwork{
						ID:   types.StringValue(item.Network.ID),
						Name: types.StringValue(item.Network.Name),
					}
				}
				return &ResponseItemSensorGetOrganizationSensorReadingsLatestNetwork{}
			}(),
			Readings: func() *[]ResponseItemSensorGetOrganizationSensorReadingsLatestReadings {
				if item.Readings != nil {
					result := make([]ResponseItemSensorGetOrganizationSensorReadingsLatestReadings, len(*item.Readings))
					for i, readings := range *item.Readings {
						result[i] = ResponseItemSensorGetOrganizationSensorReadingsLatestReadings{
							ApparentPower: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsApparentPower {
								if readings.ApparentPower != nil {
									return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsApparentPower{
										Draw: func() types.Float64 {
											if readings.ApparentPower.Draw != nil {
												return types.Float64Value(float64(*readings.ApparentPower.Draw))
											}
											return types.Float64{}
										}(),
									}
								}
								return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsApparentPower{}
							}(),
							Battery: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsBattery {
								if readings.Battery != nil {
									return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsBattery{
										Percentage: func() types.Int64 {
											if readings.Battery.Percentage != nil {
												return types.Int64Value(int64(*readings.Battery.Percentage))
											}
											return types.Int64{}
										}(),
									}
								}
								return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsBattery{}
							}(),
							Button: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsButton {
								if readings.Button != nil {
									return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsButton{
										PressType: types.StringValue(readings.Button.PressType),
									}
								}
								return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsButton{}
							}(),
							Co2: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsCo2 {
								if readings.Co2 != nil {
									return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsCo2{
										Concentration: func() types.Int64 {
											if readings.Co2.Concentration != nil {
												return types.Int64Value(int64(*readings.Co2.Concentration))
											}
											return types.Int64{}
										}(),
									}
								}
								return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsCo2{}
							}(),
							Current: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsCurrent {
								if readings.Current != nil {
									return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsCurrent{
										Draw: func() types.Float64 {
											if readings.Current.Draw != nil {
												return types.Float64Value(float64(*readings.Current.Draw))
											}
											return types.Float64{}
										}(),
									}
								}
								return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsCurrent{}
							}(),
							Door: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsDoor {
								if readings.Door != nil {
									return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsDoor{
										Open: func() types.Bool {
											if readings.Door.Open != nil {
												return types.BoolValue(*readings.Door.Open)
											}
											return types.Bool{}
										}(),
									}
								}
								return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsDoor{}
							}(),
							DownstreamPower: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsDownstreamPower {
								if readings.DownstreamPower != nil {
									return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsDownstreamPower{
										Enabled: func() types.Bool {
											if readings.DownstreamPower.Enabled != nil {
												return types.BoolValue(*readings.DownstreamPower.Enabled)
											}
											return types.Bool{}
										}(),
									}
								}
								return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsDownstreamPower{}
							}(),
							Frequency: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsFrequency {
								if readings.Frequency != nil {
									return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsFrequency{
										Level: func() types.Float64 {
											if readings.Frequency.Level != nil {
												return types.Float64Value(float64(*readings.Frequency.Level))
											}
											return types.Float64{}
										}(),
									}
								}
								return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsFrequency{}
							}(),
							Humidity: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsHumidity {
								if readings.Humidity != nil {
									return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsHumidity{
										RelativePercentage: func() types.Int64 {
											if readings.Humidity.RelativePercentage != nil {
												return types.Int64Value(int64(*readings.Humidity.RelativePercentage))
											}
											return types.Int64{}
										}(),
									}
								}
								return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsHumidity{}
							}(),
							IndoorAirQuality: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsIndoorAirQuality {
								if readings.IndoorAirQuality != nil {
									return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsIndoorAirQuality{
										Score: func() types.Int64 {
											if readings.IndoorAirQuality.Score != nil {
												return types.Int64Value(int64(*readings.IndoorAirQuality.Score))
											}
											return types.Int64{}
										}(),
									}
								}
								return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsIndoorAirQuality{}
							}(),
							Metric: types.StringValue(readings.Metric),
							Noise: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsNoise {
								if readings.Noise != nil {
									return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsNoise{
										Ambient: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsNoiseAmbient {
											if readings.Noise.Ambient != nil {
												return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsNoiseAmbient{
													Level: func() types.Int64 {
														if readings.Noise.Ambient.Level != nil {
															return types.Int64Value(int64(*readings.Noise.Ambient.Level))
														}
														return types.Int64{}
													}(),
												}
											}
											return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsNoiseAmbient{}
										}(),
									}
								}
								return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsNoise{}
							}(),
							Pm25: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsPm25 {
								if readings.Pm25 != nil {
									return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsPm25{
										Concentration: func() types.Int64 {
											if readings.Pm25.Concentration != nil {
												return types.Int64Value(int64(*readings.Pm25.Concentration))
											}
											return types.Int64{}
										}(),
									}
								}
								return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsPm25{}
							}(),
							PowerFactor: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsPowerFactor {
								if readings.PowerFactor != nil {
									return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsPowerFactor{
										Percentage: func() types.Int64 {
											if readings.PowerFactor.Percentage != nil {
												return types.Int64Value(int64(*readings.PowerFactor.Percentage))
											}
											return types.Int64{}
										}(),
									}
								}
								return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsPowerFactor{}
							}(),
							RealPower: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsRealPower {
								if readings.RealPower != nil {
									return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsRealPower{
										Draw: func() types.Float64 {
											if readings.RealPower.Draw != nil {
												return types.Float64Value(float64(*readings.RealPower.Draw))
											}
											return types.Float64{}
										}(),
									}
								}
								return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsRealPower{}
							}(),
							RemoteLockoutSwitch: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsRemoteLockoutSwitch {
								if readings.RemoteLockoutSwitch != nil {
									return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsRemoteLockoutSwitch{
										Locked: func() types.Bool {
											if readings.RemoteLockoutSwitch.Locked != nil {
												return types.BoolValue(*readings.RemoteLockoutSwitch.Locked)
											}
											return types.Bool{}
										}(),
									}
								}
								return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsRemoteLockoutSwitch{}
							}(),
							Temperature: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsTemperature {
								if readings.Temperature != nil {
									return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsTemperature{
										Celsius: func() types.Float64 {
											if readings.Temperature.Celsius != nil {
												return types.Float64Value(float64(*readings.Temperature.Celsius))
											}
											return types.Float64{}
										}(),
										Fahrenheit: func() types.Float64 {
											if readings.Temperature.Fahrenheit != nil {
												return types.Float64Value(float64(*readings.Temperature.Fahrenheit))
											}
											return types.Float64{}
										}(),
									}
								}
								return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsTemperature{}
							}(),
							Ts: types.StringValue(readings.Ts),
							Tvoc: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsTvoc {
								if readings.Tvoc != nil {
									return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsTvoc{
										Concentration: func() types.Int64 {
											if readings.Tvoc.Concentration != nil {
												return types.Int64Value(int64(*readings.Tvoc.Concentration))
											}
											return types.Int64{}
										}(),
									}
								}
								return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsTvoc{}
							}(),
							Voltage: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsVoltage {
								if readings.Voltage != nil {
									return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsVoltage{
										Level: func() types.Float64 {
											if readings.Voltage.Level != nil {
												return types.Float64Value(float64(*readings.Voltage.Level))
											}
											return types.Float64{}
										}(),
									}
								}
								return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsVoltage{}
							}(),
							Water: func() *ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsWater {
								if readings.Water != nil {
									return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsWater{
										Present: func() types.Bool {
											if readings.Water.Present != nil {
												return types.BoolValue(*readings.Water.Present)
											}
											return types.Bool{}
										}(),
									}
								}
								return &ResponseItemSensorGetOrganizationSensorReadingsLatestReadingsWater{}
							}(),
						}
					}
					return &result
				}
				return &[]ResponseItemSensorGetOrganizationSensorReadingsLatestReadings{}
			}(),
			Serial: types.StringValue(item.Serial),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
