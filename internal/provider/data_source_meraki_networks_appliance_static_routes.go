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
	_ datasource.DataSource              = &NetworksApplianceStaticRoutesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceStaticRoutesDataSource{}
)

func NewNetworksApplianceStaticRoutesDataSource() datasource.DataSource {
	return &NetworksApplianceStaticRoutesDataSource{}
}

type NetworksApplianceStaticRoutesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceStaticRoutesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceStaticRoutesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_static_routes"
}

func (d *NetworksApplianceStaticRoutesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"static_route_id": schema.StringAttribute{
				MarkdownDescription: `staticRouteId path parameter. Static route ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						Computed: true,
					},
					"fixed_ip_assignments": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"attribute_22_33_44_55_66_77": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"ip": schema.StringAttribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
					},
					"gateway_ip": schema.StringAttribute{
						Computed: true,
					},
					"gateway_vlan_id": schema.Int64Attribute{
						Computed: true,
					},
					"id": schema.StringAttribute{
						Computed: true,
					},
					"ip_version": schema.Int64Attribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"network_id": schema.StringAttribute{
						Computed: true,
					},
					"reserved_ip_ranges": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"comment": schema.StringAttribute{
									Computed: true,
								},
								"end": schema.StringAttribute{
									Computed: true,
								},
								"start": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
					"subnet": schema.StringAttribute{
						Computed: true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseApplianceGetNetworkApplianceStaticRoutes`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"enabled": schema.BoolAttribute{
							Computed: true,
						},
						"fixed_ip_assignments": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"attribute_22_33_44_55_66_77": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"ip": schema.StringAttribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
						},
						"gateway_ip": schema.StringAttribute{
							Computed: true,
						},
						"gateway_vlan_id": schema.Int64Attribute{
							Computed: true,
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"ip_version": schema.Int64Attribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"network_id": schema.StringAttribute{
							Computed: true,
						},
						"reserved_ip_ranges": schema.SetNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"comment": schema.StringAttribute{
										Computed: true,
									},
									"end": schema.StringAttribute{
										Computed: true,
									},
									"start": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"subnet": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksApplianceStaticRoutesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceStaticRoutes NetworksApplianceStaticRoutes
	diags := req.Config.Get(ctx, &networksApplianceStaticRoutes)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksApplianceStaticRoutes.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksApplianceStaticRoutes.NetworkID.IsNull(), !networksApplianceStaticRoutes.StaticRouteID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceStaticRoutes")
		vvNetworkID := networksApplianceStaticRoutes.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceStaticRoutes(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceStaticRoutes",
				err.Error(),
			)
			return
		}

		networksApplianceStaticRoutes = ResponseApplianceGetNetworkApplianceStaticRoutesItemsToBody(networksApplianceStaticRoutes, response1)
		diags = resp.State.Set(ctx, &networksApplianceStaticRoutes)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceStaticRoute")
		vvNetworkID := networksApplianceStaticRoutes.NetworkID.ValueString()
		vvStaticRouteID := networksApplianceStaticRoutes.StaticRouteID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Appliance.GetNetworkApplianceStaticRoute(vvNetworkID, vvStaticRouteID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceStaticRoute",
				err.Error(),
			)
			return
		}

		networksApplianceStaticRoutes = ResponseApplianceGetNetworkApplianceStaticRouteItemToBody(networksApplianceStaticRoutes, response2)
		diags = resp.State.Set(ctx, &networksApplianceStaticRoutes)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceStaticRoutes struct {
	NetworkID     types.String                                            `tfsdk:"network_id"`
	StaticRouteID types.String                                            `tfsdk:"static_route_id"`
	Items         *[]ResponseItemApplianceGetNetworkApplianceStaticRoutes `tfsdk:"items"`
	Item          *ResponseApplianceGetNetworkApplianceStaticRoute        `tfsdk:"item"`
}

type ResponseItemApplianceGetNetworkApplianceStaticRoutes struct {
	Enabled            types.Bool                                                              `tfsdk:"enabled"`
	FixedIPAssignments *ResponseItemApplianceGetNetworkApplianceStaticRoutesFixedIpAssignments `tfsdk:"fixed_ip_assignments"`
	GatewayIP          types.String                                                            `tfsdk:"gateway_ip"`
	GatewayVLANID      types.Int64                                                             `tfsdk:"gateway_vlan_id"`
	ID                 types.String                                                            `tfsdk:"id"`
	IPVersion          types.Int64                                                             `tfsdk:"ip_version"`
	Name               types.String                                                            `tfsdk:"name"`
	NetworkID          types.String                                                            `tfsdk:"network_id"`
	ReservedIPRanges   *[]ResponseItemApplianceGetNetworkApplianceStaticRoutesReservedIpRanges `tfsdk:"reserved_ip_ranges"`
	Subnet             types.String                                                            `tfsdk:"subnet"`
}

type ResponseItemApplianceGetNetworkApplianceStaticRoutesFixedIpAssignments struct {
	Status223344556677 *ResponseItemApplianceGetNetworkApplianceStaticRoutesFixedIpAssignments223344556677 `tfsdk:"attribute_22_33_44_55_66_77"`
}

type ResponseItemApplianceGetNetworkApplianceStaticRoutesFixedIpAssignments223344556677 struct {
	IP   types.String `tfsdk:"ip"`
	Name types.String `tfsdk:"name"`
}

type ResponseItemApplianceGetNetworkApplianceStaticRoutesReservedIpRanges struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

type ResponseApplianceGetNetworkApplianceStaticRoute struct {
	Enabled            types.Bool                                                         `tfsdk:"enabled"`
	FixedIPAssignments *ResponseApplianceGetNetworkApplianceStaticRouteFixedIpAssignments `tfsdk:"fixed_ip_assignments"`
	GatewayIP          types.String                                                       `tfsdk:"gateway_ip"`
	GatewayVLANID      types.Int64                                                        `tfsdk:"gateway_vlan_id"`
	ID                 types.String                                                       `tfsdk:"id"`
	IPVersion          types.Int64                                                        `tfsdk:"ip_version"`
	Name               types.String                                                       `tfsdk:"name"`
	NetworkID          types.String                                                       `tfsdk:"network_id"`
	ReservedIPRanges   *[]ResponseApplianceGetNetworkApplianceStaticRouteReservedIpRanges `tfsdk:"reserved_ip_ranges"`
	Subnet             types.String                                                       `tfsdk:"subnet"`
}

type ResponseApplianceGetNetworkApplianceStaticRouteFixedIpAssignments struct {
	Status223344556677 *ResponseApplianceGetNetworkApplianceStaticRouteFixedIpAssignments223344556677 `tfsdk:"attribute_22_33_44_55_66_77"`
}

type ResponseApplianceGetNetworkApplianceStaticRouteFixedIpAssignments223344556677 struct {
	IP   types.String `tfsdk:"ip"`
	Name types.String `tfsdk:"name"`
}

type ResponseApplianceGetNetworkApplianceStaticRouteReservedIpRanges struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceStaticRoutesItemsToBody(state NetworksApplianceStaticRoutes, response *merakigosdk.ResponseApplianceGetNetworkApplianceStaticRoutes) NetworksApplianceStaticRoutes {
	var items []ResponseItemApplianceGetNetworkApplianceStaticRoutes
	for _, item := range *response {
		itemState := ResponseItemApplianceGetNetworkApplianceStaticRoutes{
			Enabled: func() types.Bool {
				if item.Enabled != nil {
					return types.BoolValue(*item.Enabled)
				}
				return types.Bool{}
			}(),
			FixedIPAssignments: func() *ResponseItemApplianceGetNetworkApplianceStaticRoutesFixedIpAssignments {
				if item.FixedIPAssignments != nil {
					return &ResponseItemApplianceGetNetworkApplianceStaticRoutesFixedIpAssignments{
						Status223344556677: func() *ResponseItemApplianceGetNetworkApplianceStaticRoutesFixedIpAssignments223344556677 {
							if item.FixedIPAssignments.Status223344556677 != nil {
								return &ResponseItemApplianceGetNetworkApplianceStaticRoutesFixedIpAssignments223344556677{
									IP:   types.StringValue(item.FixedIPAssignments.Status223344556677.IP),
									Name: types.StringValue(item.FixedIPAssignments.Status223344556677.Name),
								}
							}
							return &ResponseItemApplianceGetNetworkApplianceStaticRoutesFixedIpAssignments223344556677{}
						}(),
					}
				}
				return &ResponseItemApplianceGetNetworkApplianceStaticRoutesFixedIpAssignments{}
			}(),
			GatewayIP: types.StringValue(item.GatewayIP),
			GatewayVLANID: func() types.Int64 {
				if item.GatewayVLANID != nil {
					return types.Int64Value(int64(*item.GatewayVLANID))
				}
				return types.Int64{}
			}(),
			ID: types.StringValue(item.ID),
			IPVersion: func() types.Int64 {
				if item.IPVersion != nil {
					return types.Int64Value(int64(*item.IPVersion))
				}
				return types.Int64{}
			}(),
			Name:      types.StringValue(item.Name),
			NetworkID: types.StringValue(item.NetworkID),
			ReservedIPRanges: func() *[]ResponseItemApplianceGetNetworkApplianceStaticRoutesReservedIpRanges {
				if item.ReservedIPRanges != nil {
					result := make([]ResponseItemApplianceGetNetworkApplianceStaticRoutesReservedIpRanges, len(*item.ReservedIPRanges))
					for i, reservedIPRanges := range *item.ReservedIPRanges {
						result[i] = ResponseItemApplianceGetNetworkApplianceStaticRoutesReservedIpRanges{
							Comment: types.StringValue(reservedIPRanges.Comment),
							End:     types.StringValue(reservedIPRanges.End),
							Start:   types.StringValue(reservedIPRanges.Start),
						}
					}
					return &result
				}
				return &[]ResponseItemApplianceGetNetworkApplianceStaticRoutesReservedIpRanges{}
			}(),
			Subnet: types.StringValue(item.Subnet),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseApplianceGetNetworkApplianceStaticRouteItemToBody(state NetworksApplianceStaticRoutes, response *merakigosdk.ResponseApplianceGetNetworkApplianceStaticRoute) NetworksApplianceStaticRoutes {
	itemState := ResponseApplianceGetNetworkApplianceStaticRoute{
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		FixedIPAssignments: func() *ResponseApplianceGetNetworkApplianceStaticRouteFixedIpAssignments {
			if response.FixedIPAssignments != nil {
				return &ResponseApplianceGetNetworkApplianceStaticRouteFixedIpAssignments{
					Status223344556677: func() *ResponseApplianceGetNetworkApplianceStaticRouteFixedIpAssignments223344556677 {
						if response.FixedIPAssignments.Status223344556677 != nil {
							return &ResponseApplianceGetNetworkApplianceStaticRouteFixedIpAssignments223344556677{
								IP:   types.StringValue(response.FixedIPAssignments.Status223344556677.IP),
								Name: types.StringValue(response.FixedIPAssignments.Status223344556677.Name),
							}
						}
						return &ResponseApplianceGetNetworkApplianceStaticRouteFixedIpAssignments223344556677{}
					}(),
				}
			}
			return &ResponseApplianceGetNetworkApplianceStaticRouteFixedIpAssignments{}
		}(),
		GatewayIP: types.StringValue(response.GatewayIP),
		GatewayVLANID: func() types.Int64 {
			if response.GatewayVLANID != nil {
				return types.Int64Value(int64(*response.GatewayVLANID))
			}
			return types.Int64{}
		}(),
		ID: types.StringValue(response.ID),
		IPVersion: func() types.Int64 {
			if response.IPVersion != nil {
				return types.Int64Value(int64(*response.IPVersion))
			}
			return types.Int64{}
		}(),
		Name:      types.StringValue(response.Name),
		NetworkID: types.StringValue(response.NetworkID),
		ReservedIPRanges: func() *[]ResponseApplianceGetNetworkApplianceStaticRouteReservedIpRanges {
			if response.ReservedIPRanges != nil {
				result := make([]ResponseApplianceGetNetworkApplianceStaticRouteReservedIpRanges, len(*response.ReservedIPRanges))
				for i, reservedIPRanges := range *response.ReservedIPRanges {
					result[i] = ResponseApplianceGetNetworkApplianceStaticRouteReservedIpRanges{
						Comment: types.StringValue(reservedIPRanges.Comment),
						End:     types.StringValue(reservedIPRanges.End),
						Start:   types.StringValue(reservedIPRanges.Start),
					}
				}
				return &result
			}
			return &[]ResponseApplianceGetNetworkApplianceStaticRouteReservedIpRanges{}
		}(),
		Subnet: types.StringValue(response.Subnet),
	}
	state.Item = &itemState
	return state
}
