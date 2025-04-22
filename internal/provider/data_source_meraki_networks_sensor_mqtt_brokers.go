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
	_ datasource.DataSource              = &NetworksSensorMqttBrokersDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSensorMqttBrokersDataSource{}
)

func NewNetworksSensorMqttBrokersDataSource() datasource.DataSource {
	return &NetworksSensorMqttBrokersDataSource{}
}

type NetworksSensorMqttBrokersDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSensorMqttBrokersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSensorMqttBrokersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sensor_mqtt_brokers"
}

func (d *NetworksSensorMqttBrokersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"mqtt_broker_id": schema.StringAttribute{
				MarkdownDescription: `mqttBrokerId path parameter. Mqtt broker ID`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Specifies whether the broker is enabled for sensor data. Currently, only a single broker may be enabled for sensor data.`,
						Computed:            true,
					},
					"mqtt_broker_id": schema.StringAttribute{
						MarkdownDescription: `ID of the MQTT Broker.`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSensorGetNetworkSensorMqttBrokers`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"enabled": schema.BoolAttribute{
							MarkdownDescription: `Specifies whether the broker is enabled for sensor data. Currently, only a single broker may be enabled for sensor data.`,
							Computed:            true,
						},
						"mqtt_broker_id": schema.StringAttribute{
							MarkdownDescription: `ID of the MQTT Broker.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSensorMqttBrokersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSensorMqttBrokers NetworksSensorMqttBrokers
	diags := req.Config.Get(ctx, &networksSensorMqttBrokers)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksSensorMqttBrokers.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksSensorMqttBrokers.NetworkID.IsNull(), !networksSensorMqttBrokers.MqttBrokerID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSensorMqttBrokers")
		vvNetworkID := networksSensorMqttBrokers.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Sensor.GetNetworkSensorMqttBrokers(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSensorMqttBrokers",
				err.Error(),
			)
			return
		}

		networksSensorMqttBrokers = ResponseSensorGetNetworkSensorMqttBrokersItemsToBody(networksSensorMqttBrokers, response1)
		diags = resp.State.Set(ctx, &networksSensorMqttBrokers)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkSensorMqttBroker")
		vvNetworkID := networksSensorMqttBrokers.NetworkID.ValueString()
		vvMqttBrokerID := networksSensorMqttBrokers.MqttBrokerID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Sensor.GetNetworkSensorMqttBroker(vvNetworkID, vvMqttBrokerID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSensorMqttBroker",
				err.Error(),
			)
			return
		}

		networksSensorMqttBrokers = ResponseSensorGetNetworkSensorMqttBrokerItemToBody(networksSensorMqttBrokers, response2)
		diags = resp.State.Set(ctx, &networksSensorMqttBrokers)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSensorMqttBrokers struct {
	NetworkID    types.String                                     `tfsdk:"network_id"`
	MqttBrokerID types.String                                     `tfsdk:"mqtt_broker_id"`
	Items        *[]ResponseItemSensorGetNetworkSensorMqttBrokers `tfsdk:"items"`
	Item         *ResponseSensorGetNetworkSensorMqttBroker        `tfsdk:"item"`
}

type ResponseItemSensorGetNetworkSensorMqttBrokers struct {
	Enabled      types.Bool   `tfsdk:"enabled"`
	MqttBrokerID types.String `tfsdk:"mqtt_broker_id"`
}

type ResponseSensorGetNetworkSensorMqttBroker struct {
	Enabled      types.Bool   `tfsdk:"enabled"`
	MqttBrokerID types.String `tfsdk:"mqtt_broker_id"`
}

// ToBody
func ResponseSensorGetNetworkSensorMqttBrokersItemsToBody(state NetworksSensorMqttBrokers, response *merakigosdk.ResponseSensorGetNetworkSensorMqttBrokers) NetworksSensorMqttBrokers {
	var items []ResponseItemSensorGetNetworkSensorMqttBrokers
	for _, item := range *response {
		itemState := ResponseItemSensorGetNetworkSensorMqttBrokers{
			Enabled: func() types.Bool {
				if item.Enabled != nil {
					return types.BoolValue(*item.Enabled)
				}
				return types.Bool{}
			}(),
			MqttBrokerID: types.StringValue(item.MqttBrokerID),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseSensorGetNetworkSensorMqttBrokerItemToBody(state NetworksSensorMqttBrokers, response *merakigosdk.ResponseSensorGetNetworkSensorMqttBroker) NetworksSensorMqttBrokers {
	itemState := ResponseSensorGetNetworkSensorMqttBroker{
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		MqttBrokerID: types.StringValue(response.MqttBrokerID),
	}
	state.Item = &itemState
	return state
}
