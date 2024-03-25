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

							"door": schema.Int64Attribute{
								MarkdownDescription: `Number of sensors that are currently alerting due to an open door`,
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
							"temperature": schema.Int64Attribute{
								MarkdownDescription: `Number of sensors that are currently alerting due to temperature readings`,
								Computed:            true,
							},
							"tvoc": schema.Int64Attribute{
								MarkdownDescription: `Number of sensors that are currently alerting due to TVOC readings`,
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
	Door             types.Int64                                                             `tfsdk:"door"`
	Humidity         types.Int64                                                             `tfsdk:"humidity"`
	IndoorAirQuality types.Int64                                                             `tfsdk:"indoor_air_quality"`
	Noise            *ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetricCountsNoise `tfsdk:"noise"`
	Pm25             types.Int64                                                             `tfsdk:"pm25"`
	Temperature      types.Int64                                                             `tfsdk:"temperature"`
	Tvoc             types.Int64                                                             `tfsdk:"tvoc"`
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
					Door: func() types.Int64 {
						if response.Counts.Door != nil {
							return types.Int64Value(int64(*response.Counts.Door))
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
						return &ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetricCountsNoise{}
					}(),
					Pm25: func() types.Int64 {
						if response.Counts.Pm25 != nil {
							return types.Int64Value(int64(*response.Counts.Pm25))
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
					Water: func() types.Int64 {
						if response.Counts.Water != nil {
							return types.Int64Value(int64(*response.Counts.Water))
						}
						return types.Int64{}
					}(),
				}
			}
			return &ResponseSensorGetNetworkSensorAlertsCurrentOverviewByMetricCounts{}
		}(),
		SupportedMetrics: StringSliceToList(response.SupportedMetrics),
	}
	state.Item = &itemState
	return state
}
