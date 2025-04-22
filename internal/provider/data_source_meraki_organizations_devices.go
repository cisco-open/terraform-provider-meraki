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
	_ datasource.DataSource              = &OrganizationsDevicesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsDevicesDataSource{}
)

func NewOrganizationsDevicesDataSource() datasource.DataSource {
	return &OrganizationsDevicesDataSource{}
}

type OrganizationsDevicesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsDevicesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsDevicesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_devices"
}

func (d *OrganizationsDevicesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"configuration_updated_after": schema.StringAttribute{
				MarkdownDescription: `configurationUpdatedAfter query parameter. Filter results by whether or not the device's configuration has been updated after the given timestamp`,
				Optional:            true,
			},
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"mac": schema.StringAttribute{
				MarkdownDescription: `mac query parameter. Optional parameter to filter devices by MAC address. All returned devices will have a MAC address that contains the search term or is an exact match.`,
				Optional:            true,
			},
			"macs": schema.ListAttribute{
				MarkdownDescription: `macs query parameter. Optional parameter to filter devices by one or more MAC addresses. All returned devices will have a MAC address that is an exact match.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"model": schema.StringAttribute{
				MarkdownDescription: `model query parameter. Optional parameter to filter devices by model. All returned devices will have a model that contains the search term or is an exact match.`,
				Optional:            true,
			},
			"models": schema.ListAttribute{
				MarkdownDescription: `models query parameter. Optional parameter to filter devices by one or more models. All returned devices will have a model that is an exact match.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `name query parameter. Optional parameter to filter devices by name. All returned devices will have a name that contains the search term or is an exact match.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter devices by network.`,
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
			"product_types": schema.ListAttribute{
				MarkdownDescription: `productTypes query parameter. Optional parameter to filter devices by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, and secureConnect.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"sensor_alert_profile_ids": schema.ListAttribute{
				MarkdownDescription: `sensorAlertProfileIds query parameter. Optional parameter to filter devices by the alert profiles that are bound to them. Only applies to sensor devices.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"sensor_metrics": schema.ListAttribute{
				MarkdownDescription: `sensorMetrics query parameter. Optional parameter to filter devices by the metrics that they provide. Only applies to sensor devices.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial query parameter. Optional parameter to filter devices by serial number. All returned devices will have a serial number that contains the search term or is an exact match.`,
				Optional:            true,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Optional parameter to filter devices by one or more serial numbers. All returned devices will have a serial number that is an exact match.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"tags": schema.ListAttribute{
				MarkdownDescription: `tags query parameter. Optional parameter to filter devices by tags.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"tags_filter_type": schema.StringAttribute{
				MarkdownDescription: `tagsFilterType query parameter. Optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationDevices`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"address": schema.StringAttribute{
							MarkdownDescription: `Physical address of the device`,
							Computed:            true,
						},
						"details": schema.SetNestedAttribute{
							MarkdownDescription: `Additional device information`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"name": schema.StringAttribute{
										MarkdownDescription: `Additional property name`,
										Computed:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: `Additional property value`,
										Computed:            true,
									},
								},
							},
						},
						"firmware": schema.StringAttribute{
							MarkdownDescription: `Firmware version of the device`,
							Computed:            true,
						},
						"imei": schema.StringAttribute{
							MarkdownDescription: `IMEI of the device, if applicable`,
							Computed:            true,
						},
						"lan_ip": schema.StringAttribute{
							MarkdownDescription: `LAN IP address of the device`,
							Computed:            true,
						},
						"lat": schema.Float64Attribute{
							MarkdownDescription: `Latitude of the device`,
							Computed:            true,
						},
						"lng": schema.Float64Attribute{
							MarkdownDescription: `Longitude of the device`,
							Computed:            true,
						},
						"mac": schema.StringAttribute{
							MarkdownDescription: `MAC address of the device`,
							Computed:            true,
						},
						"model": schema.StringAttribute{
							MarkdownDescription: `Model of the device`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the device`,
							Computed:            true,
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: `ID of the network the device belongs to`,
							Computed:            true,
						},
						"notes": schema.StringAttribute{
							MarkdownDescription: `Notes for the device, limited to 255 characters`,
							Computed:            true,
						},
						"product_type": schema.StringAttribute{
							MarkdownDescription: `Product type of the device`,
							Computed:            true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `Serial number of the device`,
							Computed:            true,
						},
						"tags": schema.ListAttribute{
							MarkdownDescription: `List of tags assigned to the device`,
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsDevicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsDevices OrganizationsDevices
	diags := req.Config.Get(ctx, &organizationsDevices)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationDevices")
		vvOrganizationID := organizationsDevices.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationDevicesQueryParams{}

		queryParams1.PerPage = int(organizationsDevices.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsDevices.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsDevices.EndingBefore.ValueString()
		queryParams1.ConfigurationUpdatedAfter = organizationsDevices.ConfigurationUpdatedAfter.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsDevices.NetworkIDs)
		queryParams1.ProductTypes = elementsToStrings(ctx, organizationsDevices.ProductTypes)
		queryParams1.Tags = elementsToStrings(ctx, organizationsDevices.Tags)
		queryParams1.TagsFilterType = organizationsDevices.TagsFilterType.ValueString()
		queryParams1.Name = organizationsDevices.Name.ValueString()
		queryParams1.Mac = organizationsDevices.Mac.ValueString()
		queryParams1.Serial = organizationsDevices.Serial.ValueString()
		queryParams1.Model = organizationsDevices.Model.ValueString()
		queryParams1.Macs = elementsToStrings(ctx, organizationsDevices.Macs)
		queryParams1.Serials = elementsToStrings(ctx, organizationsDevices.Serials)
		queryParams1.SensorMetrics = elementsToStrings(ctx, organizationsDevices.SensorMetrics)
		queryParams1.SensorAlertProfileIDs = elementsToStrings(ctx, organizationsDevices.SensorAlertProfileIDs)
		queryParams1.Models = elementsToStrings(ctx, organizationsDevices.Models)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationDevices(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationDevices",
				err.Error(),
			)
			return
		}

		organizationsDevices = ResponseOrganizationsGetOrganizationDevicesItemsToBody(organizationsDevices, response1)
		diags = resp.State.Set(ctx, &organizationsDevices)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsDevices struct {
	OrganizationID            types.String                                       `tfsdk:"organization_id"`
	PerPage                   types.Int64                                        `tfsdk:"per_page"`
	StartingAfter             types.String                                       `tfsdk:"starting_after"`
	EndingBefore              types.String                                       `tfsdk:"ending_before"`
	ConfigurationUpdatedAfter types.String                                       `tfsdk:"configuration_updated_after"`
	NetworkIDs                types.List                                         `tfsdk:"network_ids"`
	ProductTypes              types.List                                         `tfsdk:"product_types"`
	Tags                      types.List                                         `tfsdk:"tags"`
	TagsFilterType            types.String                                       `tfsdk:"tags_filter_type"`
	Name                      types.String                                       `tfsdk:"name"`
	Mac                       types.String                                       `tfsdk:"mac"`
	Serial                    types.String                                       `tfsdk:"serial"`
	Model                     types.String                                       `tfsdk:"model"`
	Macs                      types.List                                         `tfsdk:"macs"`
	Serials                   types.List                                         `tfsdk:"serials"`
	SensorMetrics             types.List                                         `tfsdk:"sensor_metrics"`
	SensorAlertProfileIDs     types.List                                         `tfsdk:"sensor_alert_profile_ids"`
	Models                    types.List                                         `tfsdk:"models"`
	Items                     *[]ResponseItemOrganizationsGetOrganizationDevices `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationDevices struct {
	Address     types.String                                              `tfsdk:"address"`
	Details     *[]ResponseItemOrganizationsGetOrganizationDevicesDetails `tfsdk:"details"`
	Firmware    types.String                                              `tfsdk:"firmware"`
	Imei        types.String                                              `tfsdk:"imei"`
	LanIP       types.String                                              `tfsdk:"lan_ip"`
	Lat         types.Float64                                             `tfsdk:"lat"`
	Lng         types.Float64                                             `tfsdk:"lng"`
	Mac         types.String                                              `tfsdk:"mac"`
	Model       types.String                                              `tfsdk:"model"`
	Name        types.String                                              `tfsdk:"name"`
	NetworkID   types.String                                              `tfsdk:"network_id"`
	Notes       types.String                                              `tfsdk:"notes"`
	ProductType types.String                                              `tfsdk:"product_type"`
	Serial      types.String                                              `tfsdk:"serial"`
	Tags        types.List                                                `tfsdk:"tags"`
}

type ResponseItemOrganizationsGetOrganizationDevicesDetails struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

// ToBody
func ResponseOrganizationsGetOrganizationDevicesItemsToBody(state OrganizationsDevices, response *merakigosdk.ResponseOrganizationsGetOrganizationDevices) OrganizationsDevices {
	var items []ResponseItemOrganizationsGetOrganizationDevices
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationDevices{
			Address: types.StringValue(item.Address),
			Details: func() *[]ResponseItemOrganizationsGetOrganizationDevicesDetails {
				if item.Details != nil {
					result := make([]ResponseItemOrganizationsGetOrganizationDevicesDetails, len(*item.Details))
					for i, details := range *item.Details {
						result[i] = ResponseItemOrganizationsGetOrganizationDevicesDetails{
							Name:  types.StringValue(details.Name),
							Value: types.StringValue(details.Value),
						}
					}
					return &result
				}
				return nil
			}(),
			Firmware: types.StringValue(item.Firmware),
			Imei:     types.StringValue(item.Imei),
			LanIP:    types.StringValue(item.LanIP),
			Lat: func() types.Float64 {
				if item.Lat != nil {
					return types.Float64Value(float64(*item.Lat))
				}
				return types.Float64{}
			}(),
			Lng: func() types.Float64 {
				if item.Lng != nil {
					return types.Float64Value(float64(*item.Lng))
				}
				return types.Float64{}
			}(),
			Mac:         types.StringValue(item.Mac),
			Model:       types.StringValue(item.Model),
			Name:        types.StringValue(item.Name),
			NetworkID:   types.StringValue(item.NetworkID),
			Notes:       types.StringValue(item.Notes),
			ProductType: types.StringValue(item.ProductType),
			Serial:      types.StringValue(item.Serial),
			Tags:        StringSliceToList(item.Tags),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
