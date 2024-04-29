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
	_ datasource.DataSource              = &OrganizationsInventoryDevicesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsInventoryDevicesDataSource{}
)

func NewOrganizationsInventoryDevicesDataSource() datasource.DataSource {
	return &OrganizationsInventoryDevicesDataSource{}
}

type OrganizationsInventoryDevicesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsInventoryDevicesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsInventoryDevicesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_inventory_devices"
}

func (d *OrganizationsInventoryDevicesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"macs": schema.ListAttribute{
				MarkdownDescription: `macs query parameter. Search for devices in inventory based on mac addresses.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"models": schema.ListAttribute{
				MarkdownDescription: `models query parameter. Search for devices in inventory based on model.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Search for devices in inventory based on network ids. Use explicit 'null' value to get available devices only.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"order_numbers": schema.ListAttribute{
				MarkdownDescription: `orderNumbers query parameter. Search for devices in inventory based on order numbers.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.`,
				Optional:            true,
			},
			"product_types": schema.ListAttribute{
				MarkdownDescription: `productTypes query parameter. Filter devices by product type. Accepted values are appliance, camera, cellularGateway, sensor, switch, systemsManager, and wireless.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"search": schema.StringAttribute{
				MarkdownDescription: `search query parameter. Search for devices in inventory based on serial number, mac address, or model.`,
				Optional:            true,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Optional:            true,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Search for devices in inventory based on serials.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"tags": schema.ListAttribute{
				MarkdownDescription: `tags query parameter. Filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below).`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"tags_filter_type": schema.StringAttribute{
				MarkdownDescription: `tagsFilterType query parameter. To use with 'tags' parameter, to filter devices which contain ANY or ALL given tags. Accepted values are 'withAnyTags' or 'withAllTags', default is 'withAnyTags'.`,
				Optional:            true,
			},
			"used_state": schema.StringAttribute{
				MarkdownDescription: `usedState query parameter. Filter results by used or unused inventory. Accepted values are 'used' or 'unused'.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"claimed_at": schema.StringAttribute{
						MarkdownDescription: `Claimed time of the device`,
						Computed:            true,
					},
					"country_code": schema.StringAttribute{
						MarkdownDescription: `Country/region code from device, network, or store order`,
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
					"license_expiration_date": schema.StringAttribute{
						MarkdownDescription: `License expiration date of the device`,
						Computed:            true,
					},
					"mac": schema.StringAttribute{
						MarkdownDescription: `MAC address of the device`,
						Computed:            true,
					},
					"model": schema.StringAttribute{
						MarkdownDescription: `Model type of the device`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of the device`,
						Computed:            true,
					},
					"network_id": schema.StringAttribute{
						MarkdownDescription: `Network Id of the device`,
						Computed:            true,
					},
					"order_number": schema.StringAttribute{
						MarkdownDescription: `Order number of the device`,
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
						MarkdownDescription: `Device tags`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationInventoryDevices`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"claimed_at": schema.StringAttribute{
							MarkdownDescription: `Claimed time of the device`,
							Computed:            true,
						},
						"country_code": schema.StringAttribute{
							MarkdownDescription: `Country/region code from device, network, or store order`,
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
						"license_expiration_date": schema.StringAttribute{
							MarkdownDescription: `License expiration date of the device`,
							Computed:            true,
						},
						"mac": schema.StringAttribute{
							MarkdownDescription: `MAC address of the device`,
							Computed:            true,
						},
						"model": schema.StringAttribute{
							MarkdownDescription: `Model type of the device`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the device`,
							Computed:            true,
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: `Network Id of the device`,
							Computed:            true,
						},
						"order_number": schema.StringAttribute{
							MarkdownDescription: `Order number of the device`,
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
							MarkdownDescription: `Device tags`,
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsInventoryDevicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsInventoryDevices OrganizationsInventoryDevices
	diags := req.Config.Get(ctx, &organizationsInventoryDevices)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsInventoryDevices.OrganizationID.IsNull(), !organizationsInventoryDevices.PerPage.IsNull(), !organizationsInventoryDevices.StartingAfter.IsNull(), !organizationsInventoryDevices.EndingBefore.IsNull(), !organizationsInventoryDevices.UsedState.IsNull(), !organizationsInventoryDevices.Search.IsNull(), !organizationsInventoryDevices.Macs.IsNull(), !organizationsInventoryDevices.NetworkIDs.IsNull(), !organizationsInventoryDevices.Serials.IsNull(), !organizationsInventoryDevices.Models.IsNull(), !organizationsInventoryDevices.OrderNumbers.IsNull(), !organizationsInventoryDevices.Tags.IsNull(), !organizationsInventoryDevices.TagsFilterType.IsNull(), !organizationsInventoryDevices.ProductTypes.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsInventoryDevices.OrganizationID.IsNull(), !organizationsInventoryDevices.Serial.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationInventoryDevices")
		vvOrganizationID := organizationsInventoryDevices.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationInventoryDevicesQueryParams{}

		queryParams1.PerPage = int(organizationsInventoryDevices.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsInventoryDevices.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsInventoryDevices.EndingBefore.ValueString()
		queryParams1.UsedState = organizationsInventoryDevices.UsedState.ValueString()
		queryParams1.Search = organizationsInventoryDevices.Search.ValueString()
		queryParams1.Macs = elementsToStrings(ctx, organizationsInventoryDevices.Macs)
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsInventoryDevices.NetworkIDs)
		queryParams1.Serials = elementsToStrings(ctx, organizationsInventoryDevices.Serials)
		queryParams1.Models = elementsToStrings(ctx, organizationsInventoryDevices.Models)
		queryParams1.OrderNumbers = elementsToStrings(ctx, organizationsInventoryDevices.OrderNumbers)
		queryParams1.Tags = elementsToStrings(ctx, organizationsInventoryDevices.Tags)
		queryParams1.TagsFilterType = organizationsInventoryDevices.TagsFilterType.ValueString()
		queryParams1.ProductTypes = elementsToStrings(ctx, organizationsInventoryDevices.ProductTypes)

		response1, restyResp1, err := d.client.Organizations.GetOrganizationInventoryDevices(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationInventoryDevices",
				err.Error(),
			)
			return
		}

		organizationsInventoryDevices = ResponseOrganizationsGetOrganizationInventoryDevicesItemsToBody(organizationsInventoryDevices, response1)
		diags = resp.State.Set(ctx, &organizationsInventoryDevices)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationInventoryDevice")
		vvOrganizationID := organizationsInventoryDevices.OrganizationID.ValueString()
		vvSerial := organizationsInventoryDevices.Serial.ValueString()

		response2, restyResp2, err := d.client.Organizations.GetOrganizationInventoryDevice(vvOrganizationID, vvSerial)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationInventoryDevice",
				err.Error(),
			)
			return
		}

		organizationsInventoryDevices = ResponseOrganizationsGetOrganizationInventoryDeviceItemToBody(organizationsInventoryDevices, response2)
		diags = resp.State.Set(ctx, &organizationsInventoryDevices)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsInventoryDevices struct {
	OrganizationID types.String                                                `tfsdk:"organization_id"`
	PerPage        types.Int64                                                 `tfsdk:"per_page"`
	StartingAfter  types.String                                                `tfsdk:"starting_after"`
	EndingBefore   types.String                                                `tfsdk:"ending_before"`
	UsedState      types.String                                                `tfsdk:"used_state"`
	Search         types.String                                                `tfsdk:"search"`
	Macs           types.List                                                  `tfsdk:"macs"`
	NetworkIDs     types.List                                                  `tfsdk:"network_ids"`
	Serials        types.List                                                  `tfsdk:"serials"`
	Models         types.List                                                  `tfsdk:"models"`
	OrderNumbers   types.List                                                  `tfsdk:"order_numbers"`
	Tags           types.List                                                  `tfsdk:"tags"`
	TagsFilterType types.String                                                `tfsdk:"tags_filter_type"`
	ProductTypes   types.List                                                  `tfsdk:"product_types"`
	Serial         types.String                                                `tfsdk:"serial"`
	Items          *[]ResponseItemOrganizationsGetOrganizationInventoryDevices `tfsdk:"items"`
	Item           *ResponseOrganizationsGetOrganizationInventoryDevice        `tfsdk:"item"`
}

type ResponseItemOrganizationsGetOrganizationInventoryDevices struct {
	ClaimedAt             types.String                                                       `tfsdk:"claimed_at"`
	CountryCode           types.String                                                       `tfsdk:"country_code"`
	Details               *[]ResponseItemOrganizationsGetOrganizationInventoryDevicesDetails `tfsdk:"details"`
	LicenseExpirationDate types.String                                                       `tfsdk:"license_expiration_date"`
	Mac                   types.String                                                       `tfsdk:"mac"`
	Model                 types.String                                                       `tfsdk:"model"`
	Name                  types.String                                                       `tfsdk:"name"`
	NetworkID             types.String                                                       `tfsdk:"network_id"`
	OrderNumber           types.String                                                       `tfsdk:"order_number"`
	ProductType           types.String                                                       `tfsdk:"product_type"`
	Serial                types.String                                                       `tfsdk:"serial"`
	Tags                  types.List                                                         `tfsdk:"tags"`
}

type ResponseItemOrganizationsGetOrganizationInventoryDevicesDetails struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type ResponseOrganizationsGetOrganizationInventoryDevice struct {
	ClaimedAt             types.String                                                  `tfsdk:"claimed_at"`
	CountryCode           types.String                                                  `tfsdk:"country_code"`
	Details               *[]ResponseOrganizationsGetOrganizationInventoryDeviceDetails `tfsdk:"details"`
	LicenseExpirationDate types.String                                                  `tfsdk:"license_expiration_date"`
	Mac                   types.String                                                  `tfsdk:"mac"`
	Model                 types.String                                                  `tfsdk:"model"`
	Name                  types.String                                                  `tfsdk:"name"`
	NetworkID             types.String                                                  `tfsdk:"network_id"`
	OrderNumber           types.String                                                  `tfsdk:"order_number"`
	ProductType           types.String                                                  `tfsdk:"product_type"`
	Serial                types.String                                                  `tfsdk:"serial"`
	Tags                  types.List                                                    `tfsdk:"tags"`
}

type ResponseOrganizationsGetOrganizationInventoryDeviceDetails struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

// ToBody
func ResponseOrganizationsGetOrganizationInventoryDevicesItemsToBody(state OrganizationsInventoryDevices, response *merakigosdk.ResponseOrganizationsGetOrganizationInventoryDevices) OrganizationsInventoryDevices {
	var items []ResponseItemOrganizationsGetOrganizationInventoryDevices
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationInventoryDevices{
			ClaimedAt:   types.StringValue(item.ClaimedAt),
			CountryCode: types.StringValue(item.CountryCode),
			Details: func() *[]ResponseItemOrganizationsGetOrganizationInventoryDevicesDetails {
				if item.Details != nil {
					result := make([]ResponseItemOrganizationsGetOrganizationInventoryDevicesDetails, len(*item.Details))
					for i, details := range *item.Details {
						result[i] = ResponseItemOrganizationsGetOrganizationInventoryDevicesDetails{
							Name:  types.StringValue(details.Name),
							Value: types.StringValue(details.Value),
						}
					}
					return &result
				}
				return &[]ResponseItemOrganizationsGetOrganizationInventoryDevicesDetails{}
			}(),
			LicenseExpirationDate: types.StringValue(item.LicenseExpirationDate),
			Mac:                   types.StringValue(item.Mac),
			Model:                 types.StringValue(item.Model),
			Name:                  types.StringValue(item.Name),
			NetworkID:             types.StringValue(item.NetworkID),
			OrderNumber:           types.StringValue(item.OrderNumber),
			ProductType:           types.StringValue(item.ProductType),
			Serial:                types.StringValue(item.Serial),
			Tags:                  StringSliceToList(item.Tags),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseOrganizationsGetOrganizationInventoryDeviceItemToBody(state OrganizationsInventoryDevices, response *merakigosdk.ResponseOrganizationsGetOrganizationInventoryDevice) OrganizationsInventoryDevices {
	itemState := ResponseOrganizationsGetOrganizationInventoryDevice{
		ClaimedAt:   types.StringValue(response.ClaimedAt),
		CountryCode: types.StringValue(response.CountryCode),
		Details: func() *[]ResponseOrganizationsGetOrganizationInventoryDeviceDetails {
			if response.Details != nil {
				result := make([]ResponseOrganizationsGetOrganizationInventoryDeviceDetails, len(*response.Details))
				for i, details := range *response.Details {
					result[i] = ResponseOrganizationsGetOrganizationInventoryDeviceDetails{
						Name:  types.StringValue(details.Name),
						Value: types.StringValue(details.Value),
					}
				}
				return &result
			}
			return &[]ResponseOrganizationsGetOrganizationInventoryDeviceDetails{}
		}(),
		LicenseExpirationDate: types.StringValue(response.LicenseExpirationDate),
		Mac:                   types.StringValue(response.Mac),
		Model:                 types.StringValue(response.Model),
		Name:                  types.StringValue(response.Name),
		NetworkID:             types.StringValue(response.NetworkID),
		OrderNumber:           types.StringValue(response.OrderNumber),
		ProductType:           types.StringValue(response.ProductType),
		Serial:                types.StringValue(response.Serial),
		Tags:                  StringSliceToList(response.Tags),
	}
	state.Item = &itemState
	return state
}
