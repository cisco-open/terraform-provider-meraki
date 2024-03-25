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
	_ datasource.DataSource              = &DevicesSensorRelationshipsDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesSensorRelationshipsDataSource{}
)

func NewDevicesSensorRelationshipsDataSource() datasource.DataSource {
	return &DevicesSensorRelationshipsDataSource{}
}

type DevicesSensorRelationshipsDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesSensorRelationshipsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesSensorRelationshipsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_sensor_relationships"
}

func (d *DevicesSensorRelationshipsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSensorGetDeviceSensorRelationships`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"livestream": schema.SingleNestedAttribute{
							MarkdownDescription: `A role defined between an MT sensor and an MV camera that adds the camera's livestream to the sensor's details page. Snapshots from the camera will also appear in alert notifications that the sensor triggers.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"related_devices": schema.SetNestedAttribute{
									MarkdownDescription: `An array of the related devices for the role`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"product_type": schema.StringAttribute{
												MarkdownDescription: `The product type of the related device`,
												Computed:            true,
											},
											"serial": schema.StringAttribute{
												MarkdownDescription: `The serial of the related device`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *DevicesSensorRelationshipsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesSensorRelationships DevicesSensorRelationships
	diags := req.Config.Get(ctx, &devicesSensorRelationships)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceSensorRelationships")
		vvSerial := devicesSensorRelationships.Serial.ValueString()

		response1, restyResp1, err := d.client.Sensor.GetDeviceSensorRelationships(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceSensorRelationships",
				err.Error(),
			)
			return
		}

		devicesSensorRelationships = ResponseSensorGetDeviceSensorRelationshipsItemsToBody(devicesSensorRelationships, response1)
		diags = resp.State.Set(ctx, &devicesSensorRelationships)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesSensorRelationships struct {
	Serial types.String                                      `tfsdk:"serial"`
	Items  *[]ResponseItemSensorGetDeviceSensorRelationships `tfsdk:"items"`
}

type ResponseItemSensorGetDeviceSensorRelationships struct {
	Livestream *ResponseItemSensorGetDeviceSensorRelationshipsLivestream `tfsdk:"livestream"`
}

type ResponseItemSensorGetDeviceSensorRelationshipsLivestream struct {
	RelatedDevices *[]ResponseItemSensorGetDeviceSensorRelationshipsLivestreamRelatedDevices `tfsdk:"related_devices"`
}

type ResponseItemSensorGetDeviceSensorRelationshipsLivestreamRelatedDevices struct {
	ProductType types.String `tfsdk:"product_type"`
	Serial      types.String `tfsdk:"serial"`
}

// ToBody
func ResponseSensorGetDeviceSensorRelationshipsItemsToBody(state DevicesSensorRelationships, response *merakigosdk.ResponseSensorGetDeviceSensorRelationships) DevicesSensorRelationships {
	var items []ResponseItemSensorGetDeviceSensorRelationships
	for _, item := range *response {
		itemState := ResponseItemSensorGetDeviceSensorRelationships{
			Livestream: func() *ResponseItemSensorGetDeviceSensorRelationshipsLivestream {
				if item.Livestream != nil {
					return &ResponseItemSensorGetDeviceSensorRelationshipsLivestream{
						RelatedDevices: func() *[]ResponseItemSensorGetDeviceSensorRelationshipsLivestreamRelatedDevices {
							if item.Livestream.RelatedDevices != nil {
								result := make([]ResponseItemSensorGetDeviceSensorRelationshipsLivestreamRelatedDevices, len(*item.Livestream.RelatedDevices))
								for i, relatedDevices := range *item.Livestream.RelatedDevices {
									result[i] = ResponseItemSensorGetDeviceSensorRelationshipsLivestreamRelatedDevices{
										ProductType: types.StringValue(relatedDevices.ProductType),
										Serial:      types.StringValue(relatedDevices.Serial),
									}
								}
								return &result
							}
							return &[]ResponseItemSensorGetDeviceSensorRelationshipsLivestreamRelatedDevices{}
						}(),
					}
				}
				return &ResponseItemSensorGetDeviceSensorRelationshipsLivestream{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
