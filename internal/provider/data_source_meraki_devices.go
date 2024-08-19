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
	_ datasource.DataSource              = &DevicesDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesDataSource{}
)

func NewDevicesDataSource() datasource.DataSource {
	return &DevicesDataSource{}
}

type DevicesDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices"
}

func (d *DevicesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				MarkdownDescription: `productTypes query parameter. Optional parameter to filter devices by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, and sensor.`,
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
			"item": schema.SingleNestedAttribute{
				Computed: true,
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
						"imei": schema.Float64Attribute{
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

func (d *DevicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devices Devices
	diags := req.Config.Get(ctx, &devices)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!devices.Serial.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!devices.OrganizationID.IsNull(), !devices.PerPage.IsNull(), !devices.StartingAfter.IsNull(), !devices.EndingBefore.IsNull(), !devices.ConfigurationUpdatedAfter.IsNull(), !devices.NetworkIDs.IsNull(), !devices.ProductTypes.IsNull(), !devices.Tags.IsNull(), !devices.TagsFilterType.IsNull(), !devices.Name.IsNull(), !devices.Mac.IsNull(), !devices.Serial.IsNull(), !devices.Model.IsNull(), !devices.Macs.IsNull(), !devices.Serials.IsNull(), !devices.SensorMetrics.IsNull(), !devices.SensorAlertProfileIDs.IsNull(), !devices.Models.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDevice")
		vvSerial := devices.Serial.ValueString()

		response1, restyResp1, err := d.client.Devices.GetDevice(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDevice",
				err.Error(),
			)
			return
		}

		devices = ResponseDevicesGetDeviceItemToBody(devices, response1)
		diags = resp.State.Set(ctx, &devices)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationDevices")
		vvOrganizationID := devices.OrganizationID.ValueString()
		queryParams2 := merakigosdk.GetOrganizationDevicesQueryParams{}

		queryParams2.PerPage = int(devices.PerPage.ValueInt64())
		queryParams2.StartingAfter = devices.StartingAfter.ValueString()
		queryParams2.EndingBefore = devices.EndingBefore.ValueString()
		queryParams2.ConfigurationUpdatedAfter = devices.ConfigurationUpdatedAfter.ValueString()
		queryParams2.NetworkIDs = elementsToStrings(ctx, devices.NetworkIDs)
		queryParams2.ProductTypes = elementsToStrings(ctx, devices.ProductTypes)
		queryParams2.Tags = elementsToStrings(ctx, devices.Tags)
		queryParams2.TagsFilterType = devices.TagsFilterType.ValueString()
		queryParams2.Name = devices.Name.ValueString()
		queryParams2.Mac = devices.Mac.ValueString()
		queryParams2.Serial = devices.Serial.ValueString()
		queryParams2.Model = devices.Model.ValueString()
		queryParams2.Macs = elementsToStrings(ctx, devices.Macs)
		queryParams2.Serials = elementsToStrings(ctx, devices.Serials)
		queryParams2.SensorMetrics = elementsToStrings(ctx, devices.SensorMetrics)
		queryParams2.SensorAlertProfileIDs = elementsToStrings(ctx, devices.SensorAlertProfileIDs)
		queryParams2.Models = elementsToStrings(ctx, devices.Models)

		response2, restyResp2, err := d.client.Organizations.GetOrganizationDevices(vvOrganizationID, &queryParams2)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationDevices",
				err.Error(),
			)
			return
		}

		devices = ResponseDevicesGetOrganizationDevicesItemsToBody(devices, response2)
		diags = resp.State.Set(ctx, &devices)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type Devices struct {
	Serial                    types.String                                       `tfsdk:"serial"`
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
	Model                     types.String                                       `tfsdk:"model"`
	Macs                      types.List                                         `tfsdk:"macs"`
	Serials                   types.List                                         `tfsdk:"serials"`
	SensorMetrics             types.List                                         `tfsdk:"sensor_metrics"`
	SensorAlertProfileIDs     types.List                                         `tfsdk:"sensor_alert_profile_ids"`
	Models                    types.List                                         `tfsdk:"models"`
	Items                     *[]ResponseItemOrganizationsGetOrganizationDevices `tfsdk:"items"`
	Item                      *ResponseDevicesGetDevice                          `tfsdk:"item"`
}

type ResponseDevicesGetDevice struct {
	Address     types.String                       `tfsdk:"address"`
	Details     *[]ResponseDevicesGetDeviceDetails `tfsdk:"details"`
	Firmware    types.String                       `tfsdk:"firmware"`
	Imei        types.String                       `tfsdk:"imei"`
	LanIP       types.String                       `tfsdk:"lan_ip"`
	Lat         types.Float64                      `tfsdk:"lat"`
	Lng         types.Float64                      `tfsdk:"lng"`
	Mac         types.String                       `tfsdk:"mac"`
	Model       types.String                       `tfsdk:"model"`
	Name        types.String                       `tfsdk:"name"`
	NetworkID   types.String                       `tfsdk:"network_id"`
	Notes       types.String                       `tfsdk:"notes"`
	ProductType types.String                       `tfsdk:"product_type"`
	Serial      types.String                       `tfsdk:"serial"`
	Tags        types.List                         `tfsdk:"tags"`
}

type ResponseDevicesGetDeviceDetails struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

// ToBody
func ResponseDevicesGetOrganizationDevicesItemsToBody(state Devices, response *merakigosdk.ResponseOrganizationsGetOrganizationDevices) Devices {
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
				return &[]ResponseItemOrganizationsGetOrganizationDevicesDetails{}
			}(),
			Firmware: types.StringValue(item.Firmware),
			Imei: func() types.Float64 {
				if item.Imei != nil {
					return types.Float64Value(float64(*item.Imei))
				}
				return types.Float64{}
			}(),
			LanIP: types.StringValue(item.LanIP),
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

func ResponseDevicesGetDeviceItemToBody(state Devices, response *merakigosdk.ResponseDevicesGetDevice) Devices {
	itemState := ResponseDevicesGetDevice{
		Address: types.StringValue(response.Address),
		Details: func() *[]ResponseDevicesGetDeviceDetails {
			if response.Details != nil {
				result := make([]ResponseDevicesGetDeviceDetails, len(*response.Details))
				for i, details := range *response.Details {
					result[i] = ResponseDevicesGetDeviceDetails{
						Name:  types.StringValue(details.Name),
						Value: types.StringValue(details.Value),
					}
				}
				return &result
			}
			return &[]ResponseDevicesGetDeviceDetails{}
		}(),
		Firmware: types.StringValue(response.Firmware),
		Imei:     types.StringValue(response.Imei),
		LanIP:    types.StringValue(response.LanIP),
		Lat: func() types.Float64 {
			if response.Lat != nil {
				return types.Float64Value(float64(*response.Lat))
			}
			return types.Float64{}
		}(),
		Lng: func() types.Float64 {
			if response.Lng != nil {
				return types.Float64Value(float64(*response.Lng))
			}
			return types.Float64{}
		}(),
		Mac:         types.StringValue(response.Mac),
		Model:       types.StringValue(response.Model),
		Name:        types.StringValue(response.Name),
		NetworkID:   types.StringValue(response.NetworkID),
		Notes:       types.StringValue(response.Notes),
		ProductType: types.StringValue(response.ProductType),
		Serial:      types.StringValue(response.Serial),
		Tags:        StringSliceToList(response.Tags),
	}
	state.Item = &itemState
	return state
}
