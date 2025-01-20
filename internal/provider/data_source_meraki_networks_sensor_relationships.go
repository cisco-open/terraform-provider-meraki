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
	_ datasource.DataSource              = &NetworksSensorRelationshipsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSensorRelationshipsDataSource{}
)

func NewNetworksSensorRelationshipsDataSource() datasource.DataSource {
	return &NetworksSensorRelationshipsDataSource{}
}

type NetworksSensorRelationshipsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSensorRelationshipsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSensorRelationshipsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sensor_relationships"
}

func (d *NetworksSensorRelationshipsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSensorGetNetworkSensorRelationships`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"device": schema.SingleNestedAttribute{
							MarkdownDescription: `A sensor or gateway device in the network`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"name": schema.StringAttribute{
									MarkdownDescription: `The name of the device`,
									Computed:            true,
								},
								"product_type": schema.StringAttribute{
									MarkdownDescription: `The product type of the device`,
									Computed:            true,
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `The serial of the device`,
									Computed:            true,
								},
							},
						},
						"relationships": schema.SingleNestedAttribute{
							MarkdownDescription: `An object describing the relationships defined between the device and other devices`,
							Computed:            true,
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
			},
		},
	}
}

func (d *NetworksSensorRelationshipsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSensorRelationships NetworksSensorRelationships
	diags := req.Config.Get(ctx, &networksSensorRelationships)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSensorRelationships")
		vvNetworkID := networksSensorRelationships.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Sensor.GetNetworkSensorRelationships(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSensorRelationships",
				err.Error(),
			)
			return
		}

		networksSensorRelationships = ResponseSensorGetNetworkSensorRelationshipsItemsToBody(networksSensorRelationships, response1)
		diags = resp.State.Set(ctx, &networksSensorRelationships)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSensorRelationships struct {
	NetworkID types.String                                       `tfsdk:"network_id"`
	Items     *[]ResponseItemSensorGetNetworkSensorRelationships `tfsdk:"items"`
}

type ResponseItemSensorGetNetworkSensorRelationships struct {
	Device        *ResponseItemSensorGetNetworkSensorRelationshipsDevice        `tfsdk:"device"`
	Relationships *ResponseItemSensorGetNetworkSensorRelationshipsRelationships `tfsdk:"relationships"`
}

type ResponseItemSensorGetNetworkSensorRelationshipsDevice struct {
	Name        types.String `tfsdk:"name"`
	ProductType types.String `tfsdk:"product_type"`
	Serial      types.String `tfsdk:"serial"`
}

type ResponseItemSensorGetNetworkSensorRelationshipsRelationships struct {
	Livestream *ResponseItemSensorGetNetworkSensorRelationshipsRelationshipsLivestream `tfsdk:"livestream"`
}

type ResponseItemSensorGetNetworkSensorRelationshipsRelationshipsLivestream struct {
	RelatedDevices *[]ResponseItemSensorGetNetworkSensorRelationshipsRelationshipsLivestreamRelatedDevices `tfsdk:"related_devices"`
}

type ResponseItemSensorGetNetworkSensorRelationshipsRelationshipsLivestreamRelatedDevices struct {
	ProductType types.String `tfsdk:"product_type"`
	Serial      types.String `tfsdk:"serial"`
}

// ToBody
func ResponseSensorGetNetworkSensorRelationshipsItemsToBody(state NetworksSensorRelationships, response *merakigosdk.ResponseSensorGetNetworkSensorRelationships) NetworksSensorRelationships {
	var items []ResponseItemSensorGetNetworkSensorRelationships
	for _, item := range *response {
		itemState := ResponseItemSensorGetNetworkSensorRelationships{
			Device: func() *ResponseItemSensorGetNetworkSensorRelationshipsDevice {
				if item.Device != nil {
					return &ResponseItemSensorGetNetworkSensorRelationshipsDevice{
						Name:        types.StringValue(item.Device.Name),
						ProductType: types.StringValue(item.Device.ProductType),
						Serial:      types.StringValue(item.Device.Serial),
					}
				}
				return nil
			}(),
			Relationships: func() *ResponseItemSensorGetNetworkSensorRelationshipsRelationships {
				if item.Relationships != nil {
					return &ResponseItemSensorGetNetworkSensorRelationshipsRelationships{
						Livestream: func() *ResponseItemSensorGetNetworkSensorRelationshipsRelationshipsLivestream {
							if item.Relationships.Livestream != nil {
								return &ResponseItemSensorGetNetworkSensorRelationshipsRelationshipsLivestream{
									RelatedDevices: func() *[]ResponseItemSensorGetNetworkSensorRelationshipsRelationshipsLivestreamRelatedDevices {
										if item.Relationships.Livestream.RelatedDevices != nil {
											result := make([]ResponseItemSensorGetNetworkSensorRelationshipsRelationshipsLivestreamRelatedDevices, len(*item.Relationships.Livestream.RelatedDevices))
											for i, relatedDevices := range *item.Relationships.Livestream.RelatedDevices {
												result[i] = ResponseItemSensorGetNetworkSensorRelationshipsRelationshipsLivestreamRelatedDevices{
													ProductType: types.StringValue(relatedDevices.ProductType),
													Serial:      types.StringValue(relatedDevices.Serial),
												}
											}
											return &result
										}
										return nil
									}(),
								}
							}
							return nil
						}(),
					}
				}
				return nil
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
