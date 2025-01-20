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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksSensorAlertsCurrentOverviewByMetricDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSensorAlertsCurrentOverviewByMetricDataSource{}
)

func NewNetworksSensorAlertsCurrentOverviewByMetricDataSource() datasource.DataSource {
	return &NetworksSensorAlertsCurrentOverviewByMetricDataSource{}
}

type NetworksSensorAlertsCurrentOverviewByMetricDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSensorAlertsCurrentOverviewByMetricDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSensorAlertsCurrentOverviewByMetricDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sensor_alerts_current_overview_by_metric"
}

func (d *NetworksSensorAlertsCurrentOverviewByMetricDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"counts": schema.SingleNestedAttribute{
						MarkdownDescription: `Counts of currently alerting sensors, aggregated by alerting metric`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"apparent_power": schema.Int64Attribute{
								MarkdownDescription: `Number of sensors that are currently alerting due to apparent power readings`,
								Computed:            true,
							},
							"co2": schema.Int64Attribute{
								MarkdownDescription: `Number of sensors that are currently alerting due to CO2 readings`,
								Computed:            true,
							},
							"current": schema.Int64Attribute{
								MarkdownDescription: `Number of sensors that are currently alerting due to electrical current readings`,
								Computed:            true,
							},
							"door": schema.Int64Attribute{
								MarkdownDescription: `Number of sensors that are currently alerting due to an open door`,
								Computed:            true,
							},
							"frequency": schema.Int64Attribute{
								MarkdownDescription: `Number of sensors that are currently alerting due to frequency readings`,
								Computed:            true,
							},
							"humidity": schema.Int64Attribute{
								MarkdownDescription: `Number of sensors that are currently alerting due to humidity readings`,
								Computed:            true,
							},
							"indoor_air_quality": schema.Int64Attribute{
								MarkdownDescription: `Number of sensors that are currently alerting due to indoor air quality readings`,
								Computed:            true,
							},
							"noise": schema.SingleNestedAttribute{
								MarkdownDescription: `Object containing the number of sensors that are currently alerting due to noise readings`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"ambient": schema.Int64Attribute{
										MarkdownDescription: `Number of sensors that are currently alerting due to ambient noise readings`,
										Computed:            true,
									},
								},
							},
							"pm25": schema.Int64Attribute{
								MarkdownDescription: `Number of sensors that are currently alerting due to PM2.5 readings`,
								Computed:            true,
							},
							"power_factor": schema.Int64Attribute{
								MarkdownDescription: `Number of sensors that are currently alerting due to power factor readings`,
								Computed:            true,
							},
							"real_power": schema.Int64Attribute{
								MarkdownDescription: `Number of sensors that are currently alerting due to real power readings`,
								Computed:            true,
							},
							"temperature": schema.Int64Attribute{
								MarkdownDescription: `Number of sensors that are currently alerting due to temperature readings`,
								Computed:            true,
							},
							"tvoc": schema.Int64Attribute{
								MarkdownDescription: `Number of sensors that are currently alerting due to TVOC readings`,
								Computed:            true,
							},
							"upstream_power": schema.Int64Attribute{
								MarkdownDescription: `Number of sensors that are currently alerting due to an upstream power outage`,
								Computed:            true,
							},
							"voltage": schema.Int64Attribute{
								MarkdownDescription: `Number of sensors that are currently alerting due to voltage readings`,
								Computed:            true,
							},
							"water": schema.Int64Attribute{
								MarkdownDescription: `Number of sensors that are currently alerting due to the presence of water`,
								Computed:            true,
							},
						},
					},
					"supported_metrics": schema.ListAttribute{
						MarkdownDescription: `List of metrics that are supported for alerts, based on available sensor devices in the network`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}

func (d *NetworksSensorAlertsCurrentOverviewByMetricDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSensorAlertsCurrentOverviewByMetric NetworksSensorAlertsCurrentOverviewByMetric
	diags := req.Config.Get(ctx, &networksSensorAlertsCurrentOverviewByMetric)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSensorAlertsCurrentOverviewByMetric")
		vvNetworkID := networksSensorAlertsCurrentOverviewByMetric.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Sensor.GetNetworkSensorAlertsCurrentOverviewByMetric(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSensorAlertsCurrentOverviewByMetric",
				err.Error(),
			)
			return
		}

		networksSensorAlertsCurrentOverviewByMetric = ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetricItemToBody(networksSensorAlertsCurrentOverviewByMetric, response1)
		diags = resp.State.Set(ctx, &networksSensorAlertsCurrentOverviewByMetric)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSensorAlertsCurrentOverviewByMetric struct {
	NetworkID types.String                                                 `tfsdk:"network_id"`
	Item      *ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetric `tfsdk:"item"`
}

type ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetric struct {
	Counts           *ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetricCounts `tfsdk:"counts"`
	SupportedMetrics types.List                                                         `tfsdk:"supported_metrics"`
}

type ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetricCounts struct {
	ApparentPower    types.Int64                                                             `tfsdk:"apparent_power"`
	Co2              types.Int64                                                             `tfsdk:"co2"`
	Current          types.Int64                                                             `tfsdk:"current"`
	Door             types.Int64                                                             `tfsdk:"door"`
	Frequency        types.Int64                                                             `tfsdk:"frequency"`
	Humidity         types.Int64                                                             `tfsdk:"humidity"`
	IndoorAirQuality types.Int64                                                             `tfsdk:"indoor_air_quality"`
	Noise            *ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetricCountsNoise `tfsdk:"noise"`
	Pm25             types.Int64                                                             `tfsdk:"pm25"`
	PowerFactor      types.Int64                                                             `tfsdk:"power_factor"`
	RealPower        types.Int64                                                             `tfsdk:"real_power"`
	Temperature      types.Int64                                                             `tfsdk:"temperature"`
	Tvoc             types.Int64                                                             `tfsdk:"tvoc"`
	UpstreamPower    types.Int64                                                             `tfsdk:"upstream_power"`
	Voltage          types.Int64                                                             `tfsdk:"voltage"`
	Water            types.Int64                                                             `tfsdk:"water"`
}

type ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetricCountsNoise struct {
	Ambient types.Int64 `tfsdk:"ambient"`
}

// ToBody
func ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetricItemToBody(state NetworksSensorAlertsCurrentOverviewByMetric, response *merakigosdk.ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetric) NetworksSensorAlertsCurrentOverviewByMetric {
	itemState := ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetric{
		Counts: func() *ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetricCounts {
			if response.Counts != nil {
				return &ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetricCounts{
					ApparentPower: func() types.Int64 {
						if response.Counts.ApparentPower != nil {
							return types.Int64Value(int64(*response.Counts.ApparentPower))
						}
						return types.Int64{}
					}(),
					Co2: func() types.Int64 {
						if response.Counts.Co2 != nil {
							return types.Int64Value(int64(*response.Counts.Co2))
						}
						return types.Int64{}
					}(),
					Current: func() types.Int64 {
						if response.Counts.Current != nil {
							return types.Int64Value(int64(*response.Counts.Current))
						}
						return types.Int64{}
					}(),
					Door: func() types.Int64 {
						if response.Counts.Door != nil {
							return types.Int64Value(int64(*response.Counts.Door))
						}
						return types.Int64{}
					}(),
					Frequency: func() types.Int64 {
						if response.Counts.Frequency != nil {
							return types.Int64Value(int64(*response.Counts.Frequency))
						}
						return types.Int64{}
					}(),
					Humidity: func() types.Int64 {
						if response.Counts.Humidity != nil {
							return types.Int64Value(int64(*response.Counts.Humidity))
						}
						return types.Int64{}
					}(),
					IndoorAirQuality: func() types.Int64 {
						if response.Counts.IndoorAirQuality != nil {
							return types.Int64Value(int64(*response.Counts.IndoorAirQuality))
						}
						return types.Int64{}
					}(),
					Noise: func() *ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetricCountsNoise {
						if response.Counts.Noise != nil {
							return &ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetricCountsNoise{
								Ambient: func() types.Int64 {
									if response.Counts.Noise.Ambient != nil {
										return types.Int64Value(int64(*response.Counts.Noise.Ambient))
									}
									return types.Int64{}
								}(),
							}
						}
						return nil
					}(),
					Pm25: func() types.Int64 {
						if response.Counts.Pm25 != nil {
							return types.Int64Value(int64(*response.Counts.Pm25))
						}
						return types.Int64{}
					}(),
					PowerFactor: func() types.Int64 {
						if response.Counts.PowerFactor != nil {
							return types.Int64Value(int64(*response.Counts.PowerFactor))
						}
						return types.Int64{}
					}(),
					RealPower: func() types.Int64 {
						if response.Counts.RealPower != nil {
							return types.Int64Value(int64(*response.Counts.RealPower))
						}
						return types.Int64{}
					}(),
					Temperature: func() types.Int64 {
						if response.Counts.Temperature != nil {
							return types.Int64Value(int64(*response.Counts.Temperature))
						}
						return types.Int64{}
					}(),
					Tvoc: func() types.Int64 {
						if response.Counts.Tvoc != nil {
							return types.Int64Value(int64(*response.Counts.Tvoc))
						}
						return types.Int64{}
					}(),
					UpstreamPower: func() types.Int64 {
						if response.Counts.UpstreamPower != nil {
							return types.Int64Value(int64(*response.Counts.UpstreamPower))
						}
						return types.Int64{}
					}(),
					Voltage: func() types.Int64 {
						if response.Counts.Voltage != nil {
							return types.Int64Value(int64(*response.Counts.Voltage))
						}
						return types.Int64{}
					}(),
					Water: func() types.Int64 {
						if response.Counts.Water != nil {
							return types.Int64Value(int64(*response.Counts.Water))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
		SupportedMetrics: StringSliceToList(response.SupportedMetrics),
	}
	state.Item = &itemState
	return state
}
