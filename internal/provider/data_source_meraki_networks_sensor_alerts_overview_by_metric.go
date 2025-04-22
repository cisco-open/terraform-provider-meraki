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
	_ datasource.DataSource              = &NetworksSensorAlertsOverviewByMetricDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSensorAlertsOverviewByMetricDataSource{}
)

func NewNetworksSensorAlertsOverviewByMetricDataSource() datasource.DataSource {
	return &NetworksSensorAlertsOverviewByMetricDataSource{}
}

type NetworksSensorAlertsOverviewByMetricDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSensorAlertsOverviewByMetricDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSensorAlertsOverviewByMetricDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sensor_alerts_overview_by_metric"
}

func (d *NetworksSensorAlertsOverviewByMetricDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"interval": schema.Int64Attribute{
				MarkdownDescription: `interval query parameter. The time interval in seconds for returned data. The valid intervals are: 900, 3600, 86400, 604800, 2592000. The default is 604800. Interval is calculated if time params are provided.`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 731 days from today.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 366 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 366 days. The default is 7 days. If interval is provided, the timespan will be autocalculated.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSensorGetNetworkSensorAlertsOverviewByMetric`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"counts": schema.SingleNestedAttribute{
							MarkdownDescription: `Counts of sensor alerts over the timespan, by reading metric`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"apparent_power": schema.Int64Attribute{
									MarkdownDescription: `Number of sensor alerts that occurred due to apparent power readings`,
									Computed:            true,
								},
								"co2": schema.Int64Attribute{
									MarkdownDescription: `Number of sensors that are currently alerting due to CO2 readings`,
									Computed:            true,
								},
								"current": schema.Int64Attribute{
									MarkdownDescription: `Number of sensor alerts that occurred due to electrical current readings`,
									Computed:            true,
								},
								"door": schema.Int64Attribute{
									MarkdownDescription: `Number of sensor alerts that occurred due to an open door`,
									Computed:            true,
								},
								"frequency": schema.Int64Attribute{
									MarkdownDescription: `Number of sensor alerts that occurred due to frequency readings`,
									Computed:            true,
								},
								"humidity": schema.Int64Attribute{
									MarkdownDescription: `Number of sensor alerts that occurred due to humidity readings`,
									Computed:            true,
								},
								"indoor_air_quality": schema.Int64Attribute{
									MarkdownDescription: `Number of sensor alerts that occurred due to indoor air quality readings`,
									Computed:            true,
								},
								"noise": schema.SingleNestedAttribute{
									MarkdownDescription: `Object containing the number of sensor alerts that occurred due to noise readings`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"ambient": schema.Int64Attribute{
											MarkdownDescription: `Number of sensor alerts that occurred due to ambient noise readings`,
											Computed:            true,
										},
									},
								},
								"pm25": schema.Int64Attribute{
									MarkdownDescription: `Number of sensor alerts that occurred due to PM2.5 readings`,
									Computed:            true,
								},
								"power_factor": schema.Int64Attribute{
									MarkdownDescription: `Number of sensor alerts that occurred due to power factor readings`,
									Computed:            true,
								},
								"real_power": schema.Int64Attribute{
									MarkdownDescription: `Number of sensor alerts that occurred due to real power readings`,
									Computed:            true,
								},
								"temperature": schema.Int64Attribute{
									MarkdownDescription: `Number of sensor alerts that occurred due to temperature readings`,
									Computed:            true,
								},
								"tvoc": schema.Int64Attribute{
									MarkdownDescription: `Number of sensor alerts that occurred due to TVOC readings`,
									Computed:            true,
								},
								"upstream_power": schema.Int64Attribute{
									MarkdownDescription: `Number of sensor alerts that occurred due to upstream power outages`,
									Computed:            true,
								},
								"voltage": schema.Int64Attribute{
									MarkdownDescription: `Number of sensor alerts that occurred due to voltage readings`,
									Computed:            true,
								},
								"water": schema.Int64Attribute{
									MarkdownDescription: `Number of sensor alerts that occurred due to the presence of water`,
									Computed:            true,
								},
							},
						},
						"end_ts": schema.StringAttribute{
							MarkdownDescription: `End of the timespan over which sensor alerts are counted`,
							Computed:            true,
						},
						"start_ts": schema.StringAttribute{
							MarkdownDescription: `Start of the timespan over which sensor alerts are counted`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSensorAlertsOverviewByMetricDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSensorAlertsOverviewByMetric NetworksSensorAlertsOverviewByMetric
	diags := req.Config.Get(ctx, &networksSensorAlertsOverviewByMetric)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSensorAlertsOverviewByMetric")
		vvNetworkID := networksSensorAlertsOverviewByMetric.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkSensorAlertsOverviewByMetricQueryParams{}

		queryParams1.T0 = networksSensorAlertsOverviewByMetric.T0.ValueString()
		queryParams1.T1 = networksSensorAlertsOverviewByMetric.T1.ValueString()
		queryParams1.Timespan = networksSensorAlertsOverviewByMetric.Timespan.ValueFloat64()
		queryParams1.Interval = int(networksSensorAlertsOverviewByMetric.Interval.ValueInt64())

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Sensor.GetNetworkSensorAlertsOverviewByMetric(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSensorAlertsOverviewByMetric",
				err.Error(),
			)
			return
		}

		networksSensorAlertsOverviewByMetric = ResponseSensorGetNetworkSensorAlertsOverviewByMetricItemsToBody(networksSensorAlertsOverviewByMetric, response1)
		diags = resp.State.Set(ctx, &networksSensorAlertsOverviewByMetric)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSensorAlertsOverviewByMetric struct {
	NetworkID types.String                                                `tfsdk:"network_id"`
	T0        types.String                                                `tfsdk:"t0"`
	T1        types.String                                                `tfsdk:"t1"`
	Timespan  types.Float64                                               `tfsdk:"timespan"`
	Interval  types.Int64                                                 `tfsdk:"interval"`
	Items     *[]ResponseItemSensorGetNetworkSensorAlertsOverviewByMetric `tfsdk:"items"`
}

type ResponseItemSensorGetNetworkSensorAlertsOverviewByMetric struct {
	Counts  *ResponseItemSensorGetNetworkSensorAlertsOverviewByMetricCounts `tfsdk:"counts"`
	EndTs   types.String                                                    `tfsdk:"end_ts"`
	StartTs types.String                                                    `tfsdk:"start_ts"`
}

type ResponseItemSensorGetNetworkSensorAlertsOverviewByMetricCounts struct {
	ApparentPower    types.Int64                                                          `tfsdk:"apparent_power"`
	Co2              types.Int64                                                          `tfsdk:"co2"`
	Current          types.Int64                                                          `tfsdk:"current"`
	Door             types.Int64                                                          `tfsdk:"door"`
	Frequency        types.Int64                                                          `tfsdk:"frequency"`
	Humidity         types.Int64                                                          `tfsdk:"humidity"`
	IndoorAirQuality types.Int64                                                          `tfsdk:"indoor_air_quality"`
	Noise            *ResponseItemSensorGetNetworkSensorAlertsOverviewByMetricCountsNoise `tfsdk:"noise"`
	Pm25             types.Int64                                                          `tfsdk:"pm25"`
	PowerFactor      types.Int64                                                          `tfsdk:"power_factor"`
	RealPower        types.Int64                                                          `tfsdk:"real_power"`
	Temperature      types.Int64                                                          `tfsdk:"temperature"`
	Tvoc             types.Int64                                                          `tfsdk:"tvoc"`
	UpstreamPower    types.Int64                                                          `tfsdk:"upstream_power"`
	Voltage          types.Int64                                                          `tfsdk:"voltage"`
	Water            types.Int64                                                          `tfsdk:"water"`
}

type ResponseItemSensorGetNetworkSensorAlertsOverviewByMetricCountsNoise struct {
	Ambient types.Int64 `tfsdk:"ambient"`
}

// ToBody
func ResponseSensorGetNetworkSensorAlertsOverviewByMetricItemsToBody(state NetworksSensorAlertsOverviewByMetric, response *merakigosdk.ResponseSensorGetNetworkSensorAlertsOverviewByMetric) NetworksSensorAlertsOverviewByMetric {
	var items []ResponseItemSensorGetNetworkSensorAlertsOverviewByMetric
	for _, item := range *response {
		itemState := ResponseItemSensorGetNetworkSensorAlertsOverviewByMetric{
			Counts: func() *ResponseItemSensorGetNetworkSensorAlertsOverviewByMetricCounts {
				if item.Counts != nil {
					return &ResponseItemSensorGetNetworkSensorAlertsOverviewByMetricCounts{
						ApparentPower: func() types.Int64 {
							if item.Counts.ApparentPower != nil {
								return types.Int64Value(int64(*item.Counts.ApparentPower))
							}
							return types.Int64{}
						}(),
						Co2: func() types.Int64 {
							if item.Counts.Co2 != nil {
								return types.Int64Value(int64(*item.Counts.Co2))
							}
							return types.Int64{}
						}(),
						Current: func() types.Int64 {
							if item.Counts.Current != nil {
								return types.Int64Value(int64(*item.Counts.Current))
							}
							return types.Int64{}
						}(),
						Door: func() types.Int64 {
							if item.Counts.Door != nil {
								return types.Int64Value(int64(*item.Counts.Door))
							}
							return types.Int64{}
						}(),
						Frequency: func() types.Int64 {
							if item.Counts.Frequency != nil {
								return types.Int64Value(int64(*item.Counts.Frequency))
							}
							return types.Int64{}
						}(),
						Humidity: func() types.Int64 {
							if item.Counts.Humidity != nil {
								return types.Int64Value(int64(*item.Counts.Humidity))
							}
							return types.Int64{}
						}(),
						IndoorAirQuality: func() types.Int64 {
							if item.Counts.IndoorAirQuality != nil {
								return types.Int64Value(int64(*item.Counts.IndoorAirQuality))
							}
							return types.Int64{}
						}(),
						Noise: func() *ResponseItemSensorGetNetworkSensorAlertsOverviewByMetricCountsNoise {
							if item.Counts.Noise != nil {
								return &ResponseItemSensorGetNetworkSensorAlertsOverviewByMetricCountsNoise{
									Ambient: func() types.Int64 {
										if item.Counts.Noise.Ambient != nil {
											return types.Int64Value(int64(*item.Counts.Noise.Ambient))
										}
										return types.Int64{}
									}(),
								}
							}
							return nil
						}(),
						Pm25: func() types.Int64 {
							if item.Counts.Pm25 != nil {
								return types.Int64Value(int64(*item.Counts.Pm25))
							}
							return types.Int64{}
						}(),
						PowerFactor: func() types.Int64 {
							if item.Counts.PowerFactor != nil {
								return types.Int64Value(int64(*item.Counts.PowerFactor))
							}
							return types.Int64{}
						}(),
						RealPower: func() types.Int64 {
							if item.Counts.RealPower != nil {
								return types.Int64Value(int64(*item.Counts.RealPower))
							}
							return types.Int64{}
						}(),
						Temperature: func() types.Int64 {
							if item.Counts.Temperature != nil {
								return types.Int64Value(int64(*item.Counts.Temperature))
							}
							return types.Int64{}
						}(),
						Tvoc: func() types.Int64 {
							if item.Counts.Tvoc != nil {
								return types.Int64Value(int64(*item.Counts.Tvoc))
							}
							return types.Int64{}
						}(),
						UpstreamPower: func() types.Int64 {
							if item.Counts.UpstreamPower != nil {
								return types.Int64Value(int64(*item.Counts.UpstreamPower))
							}
							return types.Int64{}
						}(),
						Voltage: func() types.Int64 {
							if item.Counts.Voltage != nil {
								return types.Int64Value(int64(*item.Counts.Voltage))
							}
							return types.Int64{}
						}(),
						Water: func() types.Int64 {
							if item.Counts.Water != nil {
								return types.Int64Value(int64(*item.Counts.Water))
							}
							return types.Int64{}
						}(),
					}
				}
				return nil
			}(),
			EndTs:   types.StringValue(item.EndTs),
			StartTs: types.StringValue(item.StartTs),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
