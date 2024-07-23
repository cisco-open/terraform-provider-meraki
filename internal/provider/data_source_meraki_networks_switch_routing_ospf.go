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
	_ datasource.DataSource              = &NetworksSwitchRoutingOspfDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchRoutingOspfDataSource{}
)

func NewNetworksSwitchRoutingOspfDataSource() datasource.DataSource {
	return &NetworksSwitchRoutingOspfDataSource{}
}

type NetworksSwitchRoutingOspfDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchRoutingOspfDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchRoutingOspfDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_routing_ospf"
}

func (d *NetworksSwitchRoutingOspfDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"areas": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"area_id": schema.StringAttribute{
									Computed: true,
								},
								"area_name": schema.StringAttribute{
									Computed: true,
								},
								"area_type": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
					"dead_timer_in_seconds": schema.Int64Attribute{
						Computed: true,
					},
					"enabled": schema.BoolAttribute{
						Computed: true,
					},
					"hello_timer_in_seconds": schema.Int64Attribute{
						Computed: true,
					},
					"md5_authentication_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"md5_authentication_key": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"id": schema.Int64Attribute{
								Computed: true,
							},
							"passphrase": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"v3": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"areas": schema.SetNestedAttribute{
								Computed: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{

										"area_id": schema.StringAttribute{
											Computed: true,
										},
										"area_name": schema.StringAttribute{
											Computed: true,
										},
										"area_type": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
							"dead_timer_in_seconds": schema.Int64Attribute{
								Computed: true,
							},
							"enabled": schema.BoolAttribute{
								Computed: true,
							},
							"hello_timer_in_seconds": schema.Int64Attribute{
								Computed: true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchRoutingOspfDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchRoutingOspf NetworksSwitchRoutingOspf
	diags := req.Config.Get(ctx, &networksSwitchRoutingOspf)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchRoutingOspf")
		vvNetworkID := networksSwitchRoutingOspf.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchRoutingOspf(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchRoutingOspf",
				err.Error(),
			)
			return
		}

		networksSwitchRoutingOspf = ResponseSwitchGetNetworkSwitchRoutingOspfItemToBody(networksSwitchRoutingOspf, response1)
		diags = resp.State.Set(ctx, &networksSwitchRoutingOspf)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchRoutingOspf struct {
	NetworkID types.String                               `tfsdk:"network_id"`
	Item      *ResponseSwitchGetNetworkSwitchRoutingOspf `tfsdk:"item"`
}

type ResponseSwitchGetNetworkSwitchRoutingOspf struct {
	Areas                    *[]ResponseSwitchGetNetworkSwitchRoutingOspfAreas              `tfsdk:"areas"`
	DeadTimerInSeconds       types.Int64                                                    `tfsdk:"dead_timer_in_seconds"`
	Enabled                  types.Bool                                                     `tfsdk:"enabled"`
	HelloTimerInSeconds      types.Int64                                                    `tfsdk:"hello_timer_in_seconds"`
	Md5AuthenticationEnabled types.Bool                                                     `tfsdk:"md5_authentication_enabled"`
	Md5AuthenticationKey     *ResponseSwitchGetNetworkSwitchRoutingOspfMd5AuthenticationKey `tfsdk:"md5_authentication_key"`
	V3                       *ResponseSwitchGetNetworkSwitchRoutingOspfV3                   `tfsdk:"v3"`
}

type ResponseSwitchGetNetworkSwitchRoutingOspfAreas struct {
	AreaID   types.Int64  `tfsdk:"area_id"`
	AreaName types.String `tfsdk:"area_name"`
	AreaType types.String `tfsdk:"area_type"`
}

type ResponseSwitchGetNetworkSwitchRoutingOspfMd5AuthenticationKey struct {
	ID         types.Int64  `tfsdk:"id"`
	Passphrase types.String `tfsdk:"passphrase"`
}

type ResponseSwitchGetNetworkSwitchRoutingOspfV3 struct {
	Areas               *[]ResponseSwitchGetNetworkSwitchRoutingOspfV3Areas `tfsdk:"areas"`
	DeadTimerInSeconds  types.Int64                                         `tfsdk:"dead_timer_in_seconds"`
	Enabled             types.Bool                                          `tfsdk:"enabled"`
	HelloTimerInSeconds types.Int64                                         `tfsdk:"hello_timer_in_seconds"`
}

type ResponseSwitchGetNetworkSwitchRoutingOspfV3Areas struct {
	AreaID   types.Int64  `tfsdk:"area_id"`
	AreaName types.String `tfsdk:"area_name"`
	AreaType types.String `tfsdk:"area_type"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchRoutingOspfItemToBody(state NetworksSwitchRoutingOspf, response *merakigosdk.ResponseSwitchGetNetworkSwitchRoutingOspf) NetworksSwitchRoutingOspf {
	itemState := ResponseSwitchGetNetworkSwitchRoutingOspf{
		Areas: func() *[]ResponseSwitchGetNetworkSwitchRoutingOspfAreas {
			if response.Areas != nil {
				result := make([]ResponseSwitchGetNetworkSwitchRoutingOspfAreas, len(*response.Areas))
				for i, areas := range *response.Areas {
					result[i] = ResponseSwitchGetNetworkSwitchRoutingOspfAreas{
						AreaID:   types.Int64Value(int64(*areas.AreaID)),
						AreaName: types.StringValue(areas.AreaName),
						AreaType: types.StringValue(areas.AreaType),
					}
				}
				return &result
			}
			return &[]ResponseSwitchGetNetworkSwitchRoutingOspfAreas{}
		}(),
		DeadTimerInSeconds: func() types.Int64 {
			if response.DeadTimerInSeconds != nil {
				return types.Int64Value(int64(*response.DeadTimerInSeconds))
			}
			return types.Int64{}
		}(),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		HelloTimerInSeconds: func() types.Int64 {
			if response.HelloTimerInSeconds != nil {
				return types.Int64Value(int64(*response.HelloTimerInSeconds))
			}
			return types.Int64{}
		}(),
		Md5AuthenticationEnabled: func() types.Bool {
			if response.Md5AuthenticationEnabled != nil {
				return types.BoolValue(*response.Md5AuthenticationEnabled)
			}
			return types.Bool{}
		}(),
		Md5AuthenticationKey: func() *ResponseSwitchGetNetworkSwitchRoutingOspfMd5AuthenticationKey {
			if response.Md5AuthenticationKey != nil {
				return &ResponseSwitchGetNetworkSwitchRoutingOspfMd5AuthenticationKey{
					ID: func() types.Int64 {
						if response.Md5AuthenticationKey.ID != nil {
							return types.Int64Value(int64(*response.Md5AuthenticationKey.ID))
						}
						return types.Int64{}
					}(),
					Passphrase: types.StringValue(response.Md5AuthenticationKey.Passphrase),
				}
			}
			return nil
		}(),
		V3: func() *ResponseSwitchGetNetworkSwitchRoutingOspfV3 {
			if response.V3 != nil {
				return &ResponseSwitchGetNetworkSwitchRoutingOspfV3{
					Areas: func() *[]ResponseSwitchGetNetworkSwitchRoutingOspfV3Areas {
						if response.V3.Areas != nil {
							result := make([]ResponseSwitchGetNetworkSwitchRoutingOspfV3Areas, len(*response.V3.Areas))
							for i, areas := range *response.V3.Areas {
								result[i] = ResponseSwitchGetNetworkSwitchRoutingOspfV3Areas{
									AreaID:   types.Int64Value(int64(*areas.AreaID)),
									AreaName: types.StringValue(areas.AreaName),
									AreaType: types.StringValue(areas.AreaType),
								}
							}
							return &result
						}
						return &[]ResponseSwitchGetNetworkSwitchRoutingOspfV3Areas{}
					}(),
					DeadTimerInSeconds: func() types.Int64 {
						if response.V3.DeadTimerInSeconds != nil {
							return types.Int64Value(int64(*response.V3.DeadTimerInSeconds))
						}
						return types.Int64{}
					}(),
					Enabled: func() types.Bool {
						if response.V3.Enabled != nil {
							return types.BoolValue(*response.V3.Enabled)
						}
						return types.Bool{}
					}(),
					HelloTimerInSeconds: func() types.Int64 {
						if response.V3.HelloTimerInSeconds != nil {
							return types.Int64Value(int64(*response.V3.HelloTimerInSeconds))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
