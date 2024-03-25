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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsSensorReadingsHistoryDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSensorReadingsHistoryDataSource{}
)

func NewOrganizationsSensorReadingsHistoryDataSource() datasource.DataSource {
	return &OrganizationsSensorReadingsHistoryDataSource{}
}

type OrganizationsSensorReadingsHistoryDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSensorReadingsHistoryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSensorReadingsHistoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_sensor_readings_history"
}

func (d *OrganizationsSensorReadingsHistoryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"metrics": schema.ListAttribute{
				MarkdownDescription: `metrics query parameter. Types of sensor readings to retrieve. If no metrics are supplied, all available types of readings will be retrieved. Allowed values are battery, button, door, humidity, indoorAirQuality, noise, pm25, temperature, tvoc, and water.`,
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
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 365 days and 6 hours from today.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 7 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 2 hours.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSensorGetOrganizationSensorReadingsHistory`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

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
						"serial": schema.StringAttribute{
							MarkdownDescription: `Serial number of the sensor that took the reading.`,
							Computed:            true,
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
		},
	}
}

func (d *OrganizationsSensorReadingsHistoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSensorReadingsHistory OrganizationsSensorReadingsHistory
	diags := req.Config.Get(ctx, &organizationsSensorReadingsHistory)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSensorReadingsHistory")
		vvOrganizationID := organizationsSensorReadingsHistory.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationSensorReadingsHistoryQueryParams{}

		queryParams1.PerPage = int(organizationsSensorReadingsHistory.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsSensorReadingsHistory.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsSensorReadingsHistory.EndingBefore.ValueString()
		queryParams1.T0 = organizationsSensorReadingsHistory.T0.ValueString()
		queryParams1.T1 = organizationsSensorReadingsHistory.T1.ValueString()
		queryParams1.Timespan = organizationsSensorReadingsHistory.Timespan.ValueFloat64()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsSensorReadingsHistory.NetworkIDs)
		queryParams1.Serials = elementsToStrings(ctx, organizationsSensorReadingsHistory.Serials)
		queryParams1.Metrics = elementsToStrings(ctx, organizationsSensorReadingsHistory.Metrics)

		response1, restyResp1, err := d.client.Sensor.GetOrganizationSensorReadingsHistory(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSensorReadingsHistory",
				err.Error(),
			)
			return
		}

		organizationsSensorReadingsHistory = ResponseSensorGetOrganizationSensorReadingsHistoryItemsToBody(organizationsSensorReadingsHistory, response1)
		diags = resp.State.Set(ctx, &organizationsSensorReadingsHistory)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSensorReadingsHistory struct {
	OrganizationID types.String                                              `tfsdk:"organization_id"`
	PerPage        types.Int64                                               `tfsdk:"per_page"`
	StartingAfter  types.String                                              `tfsdk:"starting_after"`
	EndingBefore   types.String                                              `tfsdk:"ending_before"`
	T0             types.String                                              `tfsdk:"t0"`
	T1             types.String                                              `tfsdk:"t1"`
	Timespan       types.Float64                                             `tfsdk:"timespan"`
	NetworkIDs     types.List                                                `tfsdk:"network_ids"`
	Serials        types.List                                                `tfsdk:"serials"`
	Metrics        types.List                                                `tfsdk:"metrics"`
	Items          *[]ResponseItemSensorGetOrganizationSensorReadingsHistory `tfsdk:"items"`
}

type ResponseItemSensorGetOrganizationSensorReadingsHistory struct {
	Battery          *ResponseItemSensorGetOrganizationSensorReadingsHistoryBattery          `tfsdk:"battery"`
	Button           *ResponseItemSensorGetOrganizationSensorReadingsHistoryButton           `tfsdk:"button"`
	Door             *ResponseItemSensorGetOrganizationSensorReadingsHistoryDoor             `tfsdk:"door"`
	Humidity         *ResponseItemSensorGetOrganizationSensorReadingsHistoryHumidity         `tfsdk:"humidity"`
	IndoorAirQuality *ResponseItemSensorGetOrganizationSensorReadingsHistoryIndoorAirQuality `tfsdk:"indoor_air_quality"`
	Metric           types.String                                                            `tfsdk:"metric"`
	Network          *ResponseItemSensorGetOrganizationSensorReadingsHistoryNetwork          `tfsdk:"network"`
	Noise            *ResponseItemSensorGetOrganizationSensorReadingsHistoryNoise            `tfsdk:"noise"`
	Pm25             *ResponseItemSensorGetOrganizationSensorReadingsHistoryPm25             `tfsdk:"pm25"`
	Serial           types.String                                                            `tfsdk:"serial"`
	Temperature      *ResponseItemSensorGetOrganizationSensorReadingsHistoryTemperature      `tfsdk:"temperature"`
	Ts               types.String                                                            `tfsdk:"ts"`
	Tvoc             *ResponseItemSensorGetOrganizationSensorReadingsHistoryTvoc             `tfsdk:"tvoc"`
	Water            *ResponseItemSensorGetOrganizationSensorReadingsHistoryWater            `tfsdk:"water"`
}

type ResponseItemSensorGetOrganizationSensorReadingsHistoryBattery struct {
	Percentage types.Int64 `tfsdk:"percentage"`
}

type ResponseItemSensorGetOrganizationSensorReadingsHistoryButton struct {
	PressType types.String `tfsdk:"press_type"`
}

type ResponseItemSensorGetOrganizationSensorReadingsHistoryDoor struct {
	Open types.Bool `tfsdk:"open"`
}

type ResponseItemSensorGetOrganizationSensorReadingsHistoryHumidity struct {
	RelativePercentage types.Int64 `tfsdk:"relative_percentage"`
}

type ResponseItemSensorGetOrganizationSensorReadingsHistoryIndoorAirQuality struct {
	Score types.Int64 `tfsdk:"score"`
}

type ResponseItemSensorGetOrganizationSensorReadingsHistoryNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseItemSensorGetOrganizationSensorReadingsHistoryNoise struct {
	Ambient *ResponseItemSensorGetOrganizationSensorReadingsHistoryNoiseAmbient `tfsdk:"ambient"`
}

type ResponseItemSensorGetOrganizationSensorReadingsHistoryNoiseAmbient struct {
	Level types.Int64 `tfsdk:"level"`
}

type ResponseItemSensorGetOrganizationSensorReadingsHistoryPm25 struct {
	Concentration types.Int64 `tfsdk:"concentration"`
}

type ResponseItemSensorGetOrganizationSensorReadingsHistoryTemperature struct {
	Celsius    types.Float64 `tfsdk:"celsius"`
	Fahrenheit types.Float64 `tfsdk:"fahrenheit"`
}

type ResponseItemSensorGetOrganizationSensorReadingsHistoryTvoc struct {
	Concentration types.Int64 `tfsdk:"concentration"`
}

type ResponseItemSensorGetOrganizationSensorReadingsHistoryWater struct {
	Present types.Bool `tfsdk:"present"`
}

// ToBody
func ResponseSensorGetOrganizationSensorReadingsHistoryItemsToBody(state OrganizationsSensorReadingsHistory, response *merakigosdk.ResponseSensorGetOrganizationSensorReadingsHistory) OrganizationsSensorReadingsHistory {
	var items []ResponseItemSensorGetOrganizationSensorReadingsHistory
	for _, item := range *response {
		itemState := ResponseItemSensorGetOrganizationSensorReadingsHistory{
			Battery: func() *ResponseItemSensorGetOrganizationSensorReadingsHistoryBattery {
				if item.Battery != nil {
					return &ResponseItemSensorGetOrganizationSensorReadingsHistoryBattery{
						Percentage: func() types.Int64 {
							if item.Battery.Percentage != nil {
								return types.Int64Value(int64(*item.Battery.Percentage))
							}
							return types.Int64{}
						}(),
					}
				}
				return &ResponseItemSensorGetOrganizationSensorReadingsHistoryBattery{}
			}(),
			Button: func() *ResponseItemSensorGetOrganizationSensorReadingsHistoryButton {
				if item.Button != nil {
					return &ResponseItemSensorGetOrganizationSensorReadingsHistoryButton{
						PressType: types.StringValue(item.Button.PressType),
					}
				}
				return &ResponseItemSensorGetOrganizationSensorReadingsHistoryButton{}
			}(),
			Door: func() *ResponseItemSensorGetOrganizationSensorReadingsHistoryDoor {
				if item.Door != nil {
					return &ResponseItemSensorGetOrganizationSensorReadingsHistoryDoor{
						Open: func() types.Bool {
							if item.Door.Open != nil {
								return types.BoolValue(*item.Door.Open)
							}
							return types.Bool{}
						}(),
					}
				}
				return &ResponseItemSensorGetOrganizationSensorReadingsHistoryDoor{}
			}(),
			Humidity: func() *ResponseItemSensorGetOrganizationSensorReadingsHistoryHumidity {
				if item.Humidity != nil {
					return &ResponseItemSensorGetOrganizationSensorReadingsHistoryHumidity{
						RelativePercentage: func() types.Int64 {
							if item.Humidity.RelativePercentage != nil {
								return types.Int64Value(int64(*item.Humidity.RelativePercentage))
							}
							return types.Int64{}
						}(),
					}
				}
				return &ResponseItemSensorGetOrganizationSensorReadingsHistoryHumidity{}
			}(),
			IndoorAirQuality: func() *ResponseItemSensorGetOrganizationSensorReadingsHistoryIndoorAirQuality {
				if item.IndoorAirQuality != nil {
					return &ResponseItemSensorGetOrganizationSensorReadingsHistoryIndoorAirQuality{
						Score: func() types.Int64 {
							if item.IndoorAirQuality.Score != nil {
								return types.Int64Value(int64(*item.IndoorAirQuality.Score))
							}
							return types.Int64{}
						}(),
					}
				}
				return &ResponseItemSensorGetOrganizationSensorReadingsHistoryIndoorAirQuality{}
			}(),
			Metric: types.StringValue(item.Metric),
			Network: func() *ResponseItemSensorGetOrganizationSensorReadingsHistoryNetwork {
				if item.Network != nil {
					return &ResponseItemSensorGetOrganizationSensorReadingsHistoryNetwork{
						ID:   types.StringValue(item.Network.ID),
						Name: types.StringValue(item.Network.Name),
					}
				}
				return &ResponseItemSensorGetOrganizationSensorReadingsHistoryNetwork{}
			}(),
			Noise: func() *ResponseItemSensorGetOrganizationSensorReadingsHistoryNoise {
				if item.Noise != nil {
					return &ResponseItemSensorGetOrganizationSensorReadingsHistoryNoise{
						Ambient: func() *ResponseItemSensorGetOrganizationSensorReadingsHistoryNoiseAmbient {
							if item.Noise.Ambient != nil {
								return &ResponseItemSensorGetOrganizationSensorReadingsHistoryNoiseAmbient{
									Level: func() types.Int64 {
										if item.Noise.Ambient.Level != nil {
											return types.Int64Value(int64(*item.Noise.Ambient.Level))
										}
										return types.Int64{}
									}(),
								}
							}
							return &ResponseItemSensorGetOrganizationSensorReadingsHistoryNoiseAmbient{}
						}(),
					}
				}
				return &ResponseItemSensorGetOrganizationSensorReadingsHistoryNoise{}
			}(),
			Pm25: func() *ResponseItemSensorGetOrganizationSensorReadingsHistoryPm25 {
				if item.Pm25 != nil {
					return &ResponseItemSensorGetOrganizationSensorReadingsHistoryPm25{
						Concentration: func() types.Int64 {
							if item.Pm25.Concentration != nil {
								return types.Int64Value(int64(*item.Pm25.Concentration))
							}
							return types.Int64{}
						}(),
					}
				}
				return &ResponseItemSensorGetOrganizationSensorReadingsHistoryPm25{}
			}(),
			Serial: types.StringValue(item.Serial),
			Temperature: func() *ResponseItemSensorGetOrganizationSensorReadingsHistoryTemperature {
				if item.Temperature != nil {
					return &ResponseItemSensorGetOrganizationSensorReadingsHistoryTemperature{
						Celsius: func() types.Float64 {
							if item.Temperature.Celsius != nil {
								return types.Float64Value(float64(*item.Temperature.Celsius))
							}
							return types.Float64{}
						}(),
						Fahrenheit: func() types.Float64 {
							if item.Temperature.Fahrenheit != nil {
								return types.Float64Value(float64(*item.Temperature.Fahrenheit))
							}
							return types.Float64{}
						}(),
					}
				}
				return &ResponseItemSensorGetOrganizationSensorReadingsHistoryTemperature{}
			}(),
			Ts: types.StringValue(item.Ts),
			Tvoc: func() *ResponseItemSensorGetOrganizationSensorReadingsHistoryTvoc {
				if item.Tvoc != nil {
					return &ResponseItemSensorGetOrganizationSensorReadingsHistoryTvoc{
						Concentration: func() types.Int64 {
							if item.Tvoc.Concentration != nil {
								return types.Int64Value(int64(*item.Tvoc.Concentration))
							}
							return types.Int64{}
						}(),
					}
				}
				return &ResponseItemSensorGetOrganizationSensorReadingsHistoryTvoc{}
			}(),
			Water: func() *ResponseItemSensorGetOrganizationSensorReadingsHistoryWater {
				if item.Water != nil {
					return &ResponseItemSensorGetOrganizationSensorReadingsHistoryWater{
						Present: func() types.Bool {
							if item.Water.Present != nil {
								return types.BoolValue(*item.Water.Present)
							}
							return types.Bool{}
						}(),
					}
				}
				return &ResponseItemSensorGetOrganizationSensorReadingsHistoryWater{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
