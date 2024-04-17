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
	_ datasource.DataSource              = &NetworksFloorPlansDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksFloorPlansDataSource{}
)

func NewNetworksFloorPlansDataSource() datasource.DataSource {
	return &NetworksFloorPlansDataSource{}
}

type NetworksFloorPlansDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksFloorPlansDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksFloorPlansDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_floor_plans"
}

func (d *NetworksFloorPlansDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"floor_plan_id": schema.StringAttribute{
				MarkdownDescription: `floorPlanId path parameter. Floor plan ID`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"bottom_left_corner": schema.SingleNestedAttribute{
						MarkdownDescription: `The longitude and latitude of the bottom left corner of your floor plan.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"lat": schema.Float64Attribute{
								MarkdownDescription: `Latitude`,
								Computed:            true,
							},
							"lng": schema.Float64Attribute{
								MarkdownDescription: `Longitude`,
								Computed:            true,
							},
						},
					},
					"bottom_right_corner": schema.SingleNestedAttribute{
						MarkdownDescription: `The longitude and latitude of the bottom right corner of your floor plan.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"lat": schema.Float64Attribute{
								MarkdownDescription: `Latitude`,
								Computed:            true,
							},
							"lng": schema.Float64Attribute{
								MarkdownDescription: `Longitude`,
								Computed:            true,
							},
						},
					},
					"center": schema.SingleNestedAttribute{
						MarkdownDescription: `The longitude and latitude of the center of your floor plan. The 'center' or two adjacent corners (e.g. 'topLeftCorner' and 'bottomLeftCorner') must be specified. If 'center' is specified, the floor plan is placed over that point with no rotation. If two adjacent corners are specified, the floor plan is rotated to line up with the two specified points. The aspect ratio of the floor plan's image is preserved regardless of which corners/center are specified. (This means if that more than two corners are specified, only two corners may be used to preserve the floor plan's aspect ratio.). No two points can have the same latitude, longitude pair.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"lat": schema.Float64Attribute{
								MarkdownDescription: `Latitude`,
								Computed:            true,
							},
							"lng": schema.Float64Attribute{
								MarkdownDescription: `Longitude`,
								Computed:            true,
							},
						},
					},
					"devices": schema.SetNestedAttribute{
						MarkdownDescription: `List of devices for the floorplan`,
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
					"floor_plan_id": schema.StringAttribute{
						MarkdownDescription: `Floor plan ID`,
						Computed:            true,
					},
					"height": schema.Float64Attribute{
						MarkdownDescription: `The height of your floor plan.`,
						Computed:            true,
					},
					"image_extension": schema.StringAttribute{
						MarkdownDescription: `The format type of the image.`,
						Computed:            true,
					},
					"image_md5": schema.StringAttribute{
						MarkdownDescription: `The file contents (a base 64 encoded string) of your new image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in. If you upload a new image, and you do NOT specify any new geolocation fields ('center, 'topLeftCorner', etc), the floor plan will be recentered with no rotation in order to maintain the aspect ratio of your new image.`,
						Computed:            true,
					},
					"image_url": schema.StringAttribute{
						MarkdownDescription: `The url link for the floor plan image.`,
						Computed:            true,
					},
					"image_url_expires_at": schema.StringAttribute{
						MarkdownDescription: `The time the image url link will expire.`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name of your floor plan.`,
						Computed:            true,
					},
					"top_left_corner": schema.SingleNestedAttribute{
						MarkdownDescription: `The longitude and latitude of the top left corner of your floor plan.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"lat": schema.Float64Attribute{
								MarkdownDescription: `Latitude`,
								Computed:            true,
							},
							"lng": schema.Float64Attribute{
								MarkdownDescription: `Longitude`,
								Computed:            true,
							},
						},
					},
					"top_right_corner": schema.SingleNestedAttribute{
						MarkdownDescription: `The longitude and latitude of the top right corner of your floor plan.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"lat": schema.Float64Attribute{
								MarkdownDescription: `Latitude`,
								Computed:            true,
							},
							"lng": schema.Float64Attribute{
								MarkdownDescription: `Longitude`,
								Computed:            true,
							},
						},
					},
					"width": schema.Float64Attribute{
						MarkdownDescription: `The width of your floor plan.`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseNetworksGetNetworkFloorPlans`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"bottom_left_corner": schema.SingleNestedAttribute{
							MarkdownDescription: `The longitude and latitude of the bottom left corner of your floor plan.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"lat": schema.Float64Attribute{
									MarkdownDescription: `Latitude`,
									Computed:            true,
								},
								"lng": schema.Float64Attribute{
									MarkdownDescription: `Longitude`,
									Computed:            true,
								},
							},
						},
						"bottom_right_corner": schema.SingleNestedAttribute{
							MarkdownDescription: `The longitude and latitude of the bottom right corner of your floor plan.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"lat": schema.Float64Attribute{
									MarkdownDescription: `Latitude`,
									Computed:            true,
								},
								"lng": schema.Float64Attribute{
									MarkdownDescription: `Longitude`,
									Computed:            true,
								},
							},
						},
						"center": schema.SingleNestedAttribute{
							MarkdownDescription: `The longitude and latitude of the center of your floor plan. The 'center' or two adjacent corners (e.g. 'topLeftCorner' and 'bottomLeftCorner') must be specified. If 'center' is specified, the floor plan is placed over that point with no rotation. If two adjacent corners are specified, the floor plan is rotated to line up with the two specified points. The aspect ratio of the floor plan's image is preserved regardless of which corners/center are specified. (This means if that more than two corners are specified, only two corners may be used to preserve the floor plan's aspect ratio.). No two points can have the same latitude, longitude pair.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"lat": schema.Float64Attribute{
									MarkdownDescription: `Latitude`,
									Computed:            true,
								},
								"lng": schema.Float64Attribute{
									MarkdownDescription: `Longitude`,
									Computed:            true,
								},
							},
						},
						"devices": schema.SetNestedAttribute{
							MarkdownDescription: `List of devices for the floorplan`,
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
						"floor_plan_id": schema.StringAttribute{
							MarkdownDescription: `Floor plan ID`,
							Computed:            true,
						},
						"height": schema.Float64Attribute{
							MarkdownDescription: `The height of your floor plan.`,
							Computed:            true,
						},
						"image_extension": schema.StringAttribute{
							MarkdownDescription: `The format type of the image.`,
							Computed:            true,
						},
						"image_md5": schema.StringAttribute{
							MarkdownDescription: `The file contents (a base 64 encoded string) of your new image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in. If you upload a new image, and you do NOT specify any new geolocation fields ('center, 'topLeftCorner', etc), the floor plan will be recentered with no rotation in order to maintain the aspect ratio of your new image.`,
							Computed:            true,
						},
						"image_url": schema.StringAttribute{
							MarkdownDescription: `The url link for the floor plan image.`,
							Computed:            true,
						},
						"image_url_expires_at": schema.StringAttribute{
							MarkdownDescription: `The time the image url link will expire.`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of your floor plan.`,
							Computed:            true,
						},
						"top_left_corner": schema.SingleNestedAttribute{
							MarkdownDescription: `The longitude and latitude of the top left corner of your floor plan.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"lat": schema.Float64Attribute{
									MarkdownDescription: `Latitude`,
									Computed:            true,
								},
								"lng": schema.Float64Attribute{
									MarkdownDescription: `Longitude`,
									Computed:            true,
								},
							},
						},
						"top_right_corner": schema.SingleNestedAttribute{
							MarkdownDescription: `The longitude and latitude of the top right corner of your floor plan.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"lat": schema.Float64Attribute{
									MarkdownDescription: `Latitude`,
									Computed:            true,
								},
								"lng": schema.Float64Attribute{
									MarkdownDescription: `Longitude`,
									Computed:            true,
								},
							},
						},
						"width": schema.Float64Attribute{
							MarkdownDescription: `The width of your floor plan.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksFloorPlansDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksFloorPlans NetworksFloorPlans
	diags := req.Config.Get(ctx, &networksFloorPlans)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksFloorPlans.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksFloorPlans.NetworkID.IsNull(), !networksFloorPlans.FloorPlanID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkFloorPlans")
		vvNetworkID := networksFloorPlans.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Networks.GetNetworkFloorPlans(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkFloorPlans",
				err.Error(),
			)
			return
		}

		networksFloorPlans = ResponseNetworksGetNetworkFloorPlansItemsToBody(networksFloorPlans, response1)
		diags = resp.State.Set(ctx, &networksFloorPlans)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkFloorPlan")
		vvNetworkID := networksFloorPlans.NetworkID.ValueString()
		vvFloorPlanID := networksFloorPlans.FloorPlanID.ValueString()

		response2, restyResp2, err := d.client.Networks.GetNetworkFloorPlan(vvNetworkID, vvFloorPlanID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkFloorPlan",
				err.Error(),
			)
			return
		}

		networksFloorPlans = ResponseNetworksGetNetworkFloorPlanItemToBody(networksFloorPlans, response2)
		diags = resp.State.Set(ctx, &networksFloorPlans)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksFloorPlans struct {
	NetworkID   types.String                                `tfsdk:"network_id"`
	FloorPlanID types.String                                `tfsdk:"floor_plan_id"`
	Items       *[]ResponseItemNetworksGetNetworkFloorPlans `tfsdk:"items"`
	Item        *ResponseNetworksGetNetworkFloorPlan        `tfsdk:"item"`
}

type ResponseItemNetworksGetNetworkFloorPlans struct {
	BottomLeftCorner  *ResponseItemNetworksGetNetworkFloorPlansBottomLeftCorner  `tfsdk:"bottom_left_corner"`
	BottomRightCorner *ResponseItemNetworksGetNetworkFloorPlansBottomRightCorner `tfsdk:"bottom_right_corner"`
	Center            *ResponseItemNetworksGetNetworkFloorPlansCenter            `tfsdk:"center"`
	Devices           *[]ResponseItemNetworksGetNetworkFloorPlansDevices         `tfsdk:"devices"`
	FloorPlanID       types.String                                               `tfsdk:"floor_plan_id"`
	Height            types.Float64                                              `tfsdk:"height"`
	ImageExtension    types.String                                               `tfsdk:"image_extension"`
	ImageMd5          types.String                                               `tfsdk:"image_md5"`
	ImageURL          types.String                                               `tfsdk:"image_url"`
	ImageURLExpiresAt types.String                                               `tfsdk:"image_url_expires_at"`
	Name              types.String                                               `tfsdk:"name"`
	TopLeftCorner     *ResponseItemNetworksGetNetworkFloorPlansTopLeftCorner     `tfsdk:"top_left_corner"`
	TopRightCorner    *ResponseItemNetworksGetNetworkFloorPlansTopRightCorner    `tfsdk:"top_right_corner"`
	Width             types.Float64                                              `tfsdk:"width"`
}

type ResponseItemNetworksGetNetworkFloorPlansBottomLeftCorner struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseItemNetworksGetNetworkFloorPlansBottomRightCorner struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseItemNetworksGetNetworkFloorPlansCenter struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseItemNetworksGetNetworkFloorPlansDevices struct {
	Address     types.String                                              `tfsdk:"address"`
	Details     *[]ResponseItemNetworksGetNetworkFloorPlansDevicesDetails `tfsdk:"details"`
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

type ResponseItemNetworksGetNetworkFloorPlansDevicesDetails struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type ResponseItemNetworksGetNetworkFloorPlansTopLeftCorner struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseItemNetworksGetNetworkFloorPlansTopRightCorner struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseNetworksGetNetworkFloorPlan struct {
	BottomLeftCorner  *ResponseNetworksGetNetworkFloorPlanBottomLeftCorner  `tfsdk:"bottom_left_corner"`
	BottomRightCorner *ResponseNetworksGetNetworkFloorPlanBottomRightCorner `tfsdk:"bottom_right_corner"`
	Center            *ResponseNetworksGetNetworkFloorPlanCenter            `tfsdk:"center"`
	Devices           *[]ResponseNetworksGetNetworkFloorPlanDevices         `tfsdk:"devices"`
	FloorPlanID       types.String                                          `tfsdk:"floor_plan_id"`
	Height            types.Float64                                         `tfsdk:"height"`
	ImageExtension    types.String                                          `tfsdk:"image_extension"`
	ImageMd5          types.String                                          `tfsdk:"image_md5"`
	ImageURL          types.String                                          `tfsdk:"image_url"`
	ImageURLExpiresAt types.String                                          `tfsdk:"image_url_expires_at"`
	Name              types.String                                          `tfsdk:"name"`
	TopLeftCorner     *ResponseNetworksGetNetworkFloorPlanTopLeftCorner     `tfsdk:"top_left_corner"`
	TopRightCorner    *ResponseNetworksGetNetworkFloorPlanTopRightCorner    `tfsdk:"top_right_corner"`
	Width             types.Float64                                         `tfsdk:"width"`
}

type ResponseNetworksGetNetworkFloorPlanBottomLeftCorner struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseNetworksGetNetworkFloorPlanBottomRightCorner struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseNetworksGetNetworkFloorPlanCenter struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseNetworksGetNetworkFloorPlanDevices struct {
	Address     types.String                                         `tfsdk:"address"`
	Details     *[]ResponseNetworksGetNetworkFloorPlanDevicesDetails `tfsdk:"details"`
	Firmware    types.String                                         `tfsdk:"firmware"`
	Imei        types.String                                         `tfsdk:"imei"`
	LanIP       types.String                                         `tfsdk:"lan_ip"`
	Lat         types.Float64                                        `tfsdk:"lat"`
	Lng         types.Float64                                        `tfsdk:"lng"`
	Mac         types.String                                         `tfsdk:"mac"`
	Model       types.String                                         `tfsdk:"model"`
	Name        types.String                                         `tfsdk:"name"`
	NetworkID   types.String                                         `tfsdk:"network_id"`
	Notes       types.String                                         `tfsdk:"notes"`
	ProductType types.String                                         `tfsdk:"product_type"`
	Serial      types.String                                         `tfsdk:"serial"`
	Tags        types.List                                           `tfsdk:"tags"`
}

type ResponseNetworksGetNetworkFloorPlanDevicesDetails struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type ResponseNetworksGetNetworkFloorPlanTopLeftCorner struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseNetworksGetNetworkFloorPlanTopRightCorner struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

// ToBody
func ResponseNetworksGetNetworkFloorPlansItemsToBody(state NetworksFloorPlans, response *merakigosdk.ResponseNetworksGetNetworkFloorPlans) NetworksFloorPlans {
	var items []ResponseItemNetworksGetNetworkFloorPlans
	for _, item := range *response {
		itemState := ResponseItemNetworksGetNetworkFloorPlans{
			BottomLeftCorner: func() *ResponseItemNetworksGetNetworkFloorPlansBottomLeftCorner {
				if item.BottomLeftCorner != nil {
					return &ResponseItemNetworksGetNetworkFloorPlansBottomLeftCorner{
						Lat: func() types.Float64 {
							if item.BottomLeftCorner.Lat != nil {
								return types.Float64Value(float64(*item.BottomLeftCorner.Lat))
							}
							return types.Float64{}
						}(),
						Lng: func() types.Float64 {
							if item.BottomLeftCorner.Lng != nil {
								return types.Float64Value(float64(*item.BottomLeftCorner.Lng))
							}
							return types.Float64{}
						}(),
					}
				}
				return &ResponseItemNetworksGetNetworkFloorPlansBottomLeftCorner{}
			}(),
			BottomRightCorner: func() *ResponseItemNetworksGetNetworkFloorPlansBottomRightCorner {
				if item.BottomRightCorner != nil {
					return &ResponseItemNetworksGetNetworkFloorPlansBottomRightCorner{
						Lat: func() types.Float64 {
							if item.BottomRightCorner.Lat != nil {
								return types.Float64Value(float64(*item.BottomRightCorner.Lat))
							}
							return types.Float64{}
						}(),
						Lng: func() types.Float64 {
							if item.BottomRightCorner.Lng != nil {
								return types.Float64Value(float64(*item.BottomRightCorner.Lng))
							}
							return types.Float64{}
						}(),
					}
				}
				return &ResponseItemNetworksGetNetworkFloorPlansBottomRightCorner{}
			}(),
			Center: func() *ResponseItemNetworksGetNetworkFloorPlansCenter {
				if item.Center != nil {
					return &ResponseItemNetworksGetNetworkFloorPlansCenter{
						Lat: func() types.Float64 {
							if item.Center.Lat != nil {
								return types.Float64Value(float64(*item.Center.Lat))
							}
							return types.Float64{}
						}(),
						Lng: func() types.Float64 {
							if item.Center.Lng != nil {
								return types.Float64Value(float64(*item.Center.Lng))
							}
							return types.Float64{}
						}(),
					}
				}
				return &ResponseItemNetworksGetNetworkFloorPlansCenter{}
			}(),
			Devices: func() *[]ResponseItemNetworksGetNetworkFloorPlansDevices {
				if item.Devices != nil {
					result := make([]ResponseItemNetworksGetNetworkFloorPlansDevices, len(*item.Devices))
					for i, devices := range *item.Devices {
						result[i] = ResponseItemNetworksGetNetworkFloorPlansDevices{
							Address: types.StringValue(devices.Address),
							Details: func() *[]ResponseItemNetworksGetNetworkFloorPlansDevicesDetails {
								if devices.Details != nil {
									result := make([]ResponseItemNetworksGetNetworkFloorPlansDevicesDetails, len(*devices.Details))
									for i, details := range *devices.Details {
										result[i] = ResponseItemNetworksGetNetworkFloorPlansDevicesDetails{
											Name:  types.StringValue(details.Name),
											Value: types.StringValue(details.Value),
										}
									}
									return &result
								}
								return &[]ResponseItemNetworksGetNetworkFloorPlansDevicesDetails{}
							}(),
							Firmware: types.StringValue(devices.Firmware),
							Imei:     types.StringValue(devices.Imei),
							LanIP:    types.StringValue(devices.LanIP),
							Lat: func() types.Float64 {
								if devices.Lat != nil {
									return types.Float64Value(float64(*devices.Lat))
								}
								return types.Float64{}
							}(),
							Lng: func() types.Float64 {
								if devices.Lng != nil {
									return types.Float64Value(float64(*devices.Lng))
								}
								return types.Float64{}
							}(),
							Mac:         types.StringValue(devices.Mac),
							Model:       types.StringValue(devices.Model),
							Name:        types.StringValue(devices.Name),
							NetworkID:   types.StringValue(devices.NetworkID),
							Notes:       types.StringValue(devices.Notes),
							ProductType: types.StringValue(devices.ProductType),
							Serial:      types.StringValue(devices.Serial),
							Tags:        StringSliceToList(devices.Tags),
						}
					}
					return &result
				}
				return &[]ResponseItemNetworksGetNetworkFloorPlansDevices{}
			}(),
			FloorPlanID: types.StringValue(item.FloorPlanID),
			Height: func() types.Float64 {
				if item.Height != nil {
					return types.Float64Value(float64(*item.Height))
				}
				return types.Float64{}
			}(),
			ImageExtension:    types.StringValue(item.ImageExtension),
			ImageMd5:          types.StringValue(item.ImageMd5),
			ImageURL:          types.StringValue(item.ImageURL),
			ImageURLExpiresAt: types.StringValue(item.ImageURLExpiresAt),
			Name:              types.StringValue(item.Name),
			TopLeftCorner: func() *ResponseItemNetworksGetNetworkFloorPlansTopLeftCorner {
				if item.TopLeftCorner != nil {
					return &ResponseItemNetworksGetNetworkFloorPlansTopLeftCorner{
						Lat: func() types.Float64 {
							if item.TopLeftCorner.Lat != nil {
								return types.Float64Value(float64(*item.TopLeftCorner.Lat))
							}
							return types.Float64{}
						}(),
						Lng: func() types.Float64 {
							if item.TopLeftCorner.Lng != nil {
								return types.Float64Value(float64(*item.TopLeftCorner.Lng))
							}
							return types.Float64{}
						}(),
					}
				}
				return &ResponseItemNetworksGetNetworkFloorPlansTopLeftCorner{}
			}(),
			TopRightCorner: func() *ResponseItemNetworksGetNetworkFloorPlansTopRightCorner {
				if item.TopRightCorner != nil {
					return &ResponseItemNetworksGetNetworkFloorPlansTopRightCorner{
						Lat: func() types.Float64 {
							if item.TopRightCorner.Lat != nil {
								return types.Float64Value(float64(*item.TopRightCorner.Lat))
							}
							return types.Float64{}
						}(),
						Lng: func() types.Float64 {
							if item.TopRightCorner.Lng != nil {
								return types.Float64Value(float64(*item.TopRightCorner.Lng))
							}
							return types.Float64{}
						}(),
					}
				}
				return &ResponseItemNetworksGetNetworkFloorPlansTopRightCorner{}
			}(),
			Width: func() types.Float64 {
				if item.Width != nil {
					return types.Float64Value(float64(*item.Width))
				}
				return types.Float64{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseNetworksGetNetworkFloorPlanItemToBody(state NetworksFloorPlans, response *merakigosdk.ResponseNetworksGetNetworkFloorPlan) NetworksFloorPlans {
	itemState := ResponseNetworksGetNetworkFloorPlan{
		BottomLeftCorner: func() *ResponseNetworksGetNetworkFloorPlanBottomLeftCorner {
			if response.BottomLeftCorner != nil {
				return &ResponseNetworksGetNetworkFloorPlanBottomLeftCorner{
					Lat: func() types.Float64 {
						if response.BottomLeftCorner.Lat != nil {
							return types.Float64Value(float64(*response.BottomLeftCorner.Lat))
						}
						return types.Float64{}
					}(),
					Lng: func() types.Float64 {
						if response.BottomLeftCorner.Lng != nil {
							return types.Float64Value(float64(*response.BottomLeftCorner.Lng))
						}
						return types.Float64{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkFloorPlanBottomLeftCorner{}
		}(),
		BottomRightCorner: func() *ResponseNetworksGetNetworkFloorPlanBottomRightCorner {
			if response.BottomRightCorner != nil {
				return &ResponseNetworksGetNetworkFloorPlanBottomRightCorner{
					Lat: func() types.Float64 {
						if response.BottomRightCorner.Lat != nil {
							return types.Float64Value(float64(*response.BottomRightCorner.Lat))
						}
						return types.Float64{}
					}(),
					Lng: func() types.Float64 {
						if response.BottomRightCorner.Lng != nil {
							return types.Float64Value(float64(*response.BottomRightCorner.Lng))
						}
						return types.Float64{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkFloorPlanBottomRightCorner{}
		}(),
		Center: func() *ResponseNetworksGetNetworkFloorPlanCenter {
			if response.Center != nil {
				return &ResponseNetworksGetNetworkFloorPlanCenter{
					Lat: func() types.Float64 {
						if response.Center.Lat != nil {
							return types.Float64Value(float64(*response.Center.Lat))
						}
						return types.Float64{}
					}(),
					Lng: func() types.Float64 {
						if response.Center.Lng != nil {
							return types.Float64Value(float64(*response.Center.Lng))
						}
						return types.Float64{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkFloorPlanCenter{}
		}(),
		Devices: func() *[]ResponseNetworksGetNetworkFloorPlanDevices {
			if response.Devices != nil {
				result := make([]ResponseNetworksGetNetworkFloorPlanDevices, len(*response.Devices))
				for i, devices := range *response.Devices {
					result[i] = ResponseNetworksGetNetworkFloorPlanDevices{
						Address: types.StringValue(devices.Address),
						Details: func() *[]ResponseNetworksGetNetworkFloorPlanDevicesDetails {
							if devices.Details != nil {
								result := make([]ResponseNetworksGetNetworkFloorPlanDevicesDetails, len(*devices.Details))
								for i, details := range *devices.Details {
									result[i] = ResponseNetworksGetNetworkFloorPlanDevicesDetails{
										Name:  types.StringValue(details.Name),
										Value: types.StringValue(details.Value),
									}
								}
								return &result
							}
							return &[]ResponseNetworksGetNetworkFloorPlanDevicesDetails{}
						}(),
						Firmware: types.StringValue(devices.Firmware),
						Imei:     types.StringValue(devices.Imei),
						LanIP:    types.StringValue(devices.LanIP),
						Lat: func() types.Float64 {
							if devices.Lat != nil {
								return types.Float64Value(float64(*devices.Lat))
							}
							return types.Float64{}
						}(),
						Lng: func() types.Float64 {
							if devices.Lng != nil {
								return types.Float64Value(float64(*devices.Lng))
							}
							return types.Float64{}
						}(),
						Mac:         types.StringValue(devices.Mac),
						Model:       types.StringValue(devices.Model),
						Name:        types.StringValue(devices.Name),
						NetworkID:   types.StringValue(devices.NetworkID),
						Notes:       types.StringValue(devices.Notes),
						ProductType: types.StringValue(devices.ProductType),
						Serial:      types.StringValue(devices.Serial),
						Tags:        StringSliceToList(devices.Tags),
					}
				}
				return &result
			}
			return &[]ResponseNetworksGetNetworkFloorPlanDevices{}
		}(),
		FloorPlanID: types.StringValue(response.FloorPlanID),
		Height: func() types.Float64 {
			if response.Height != nil {
				return types.Float64Value(float64(*response.Height))
			}
			return types.Float64{}
		}(),
		ImageExtension:    types.StringValue(response.ImageExtension),
		ImageMd5:          types.StringValue(response.ImageMd5),
		ImageURL:          types.StringValue(response.ImageURL),
		ImageURLExpiresAt: types.StringValue(response.ImageURLExpiresAt),
		Name:              types.StringValue(response.Name),
		TopLeftCorner: func() *ResponseNetworksGetNetworkFloorPlanTopLeftCorner {
			if response.TopLeftCorner != nil {
				return &ResponseNetworksGetNetworkFloorPlanTopLeftCorner{
					Lat: func() types.Float64 {
						if response.TopLeftCorner.Lat != nil {
							return types.Float64Value(float64(*response.TopLeftCorner.Lat))
						}
						return types.Float64{}
					}(),
					Lng: func() types.Float64 {
						if response.TopLeftCorner.Lng != nil {
							return types.Float64Value(float64(*response.TopLeftCorner.Lng))
						}
						return types.Float64{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkFloorPlanTopLeftCorner{}
		}(),
		TopRightCorner: func() *ResponseNetworksGetNetworkFloorPlanTopRightCorner {
			if response.TopRightCorner != nil {
				return &ResponseNetworksGetNetworkFloorPlanTopRightCorner{
					Lat: func() types.Float64 {
						if response.TopRightCorner.Lat != nil {
							return types.Float64Value(float64(*response.TopRightCorner.Lat))
						}
						return types.Float64{}
					}(),
					Lng: func() types.Float64 {
						if response.TopRightCorner.Lng != nil {
							return types.Float64Value(float64(*response.TopRightCorner.Lng))
						}
						return types.Float64{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkFloorPlanTopRightCorner{}
		}(),
		Width: func() types.Float64 {
			if response.Width != nil {
				return types.Float64Value(float64(*response.Width))
			}
			return types.Float64{}
		}(),
	}
	state.Item = &itemState
	return state
}
