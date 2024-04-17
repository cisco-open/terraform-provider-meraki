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
	_ datasource.DataSource              = &NetworksSwitchRoutingMulticastDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchRoutingMulticastDataSource{}
)

func NewNetworksSwitchRoutingMulticastDataSource() datasource.DataSource {
	return &NetworksSwitchRoutingMulticastDataSource{}
}

type NetworksSwitchRoutingMulticastDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchRoutingMulticastDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchRoutingMulticastDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_routing_multicast"
}

func (d *NetworksSwitchRoutingMulticastDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"default_settings": schema.SingleNestedAttribute{
						MarkdownDescription: `Default multicast setting for entire network. IGMP snooping and Flood unknown
      multicast traffic settings are enabled by default.`,
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"flood_unknown_multicast_traffic_enabled": schema.BoolAttribute{
								MarkdownDescription: `Flood unknown multicast traffic enabled for the entire network`,
								Computed:            true,
							},
							"igmp_snooping_enabled": schema.BoolAttribute{
								MarkdownDescription: `IGMP snooping enabled for the entire network`,
								Computed:            true,
							},
						},
					},
					"overrides": schema.SetNestedAttribute{
						MarkdownDescription: `Array of paired switches/stacks/profiles and corresponding multicast settings.
      An empty array will clear the multicast settings.`,
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"flood_unknown_multicast_traffic_enabled": schema.BoolAttribute{
									MarkdownDescription: `Flood unknown multicast traffic enabled for switches, switch stacks or switch templates`,
									Computed:            true,
								},
								"igmp_snooping_enabled": schema.BoolAttribute{
									MarkdownDescription: `IGMP snooping enabled for switches, switch stacks or switch templates`,
									Computed:            true,
								},
								"stacks": schema.ListAttribute{
									MarkdownDescription: `(optional) List of switch stack ids for non-template network`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"switch_profiles": schema.ListAttribute{
									MarkdownDescription: `(optional) List of switch templates ids for template network`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"switches": schema.ListAttribute{
									MarkdownDescription: `(optional) List of switch serials for non-template network`,
									Computed:            true,
									ElementType:         types.StringType,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchRoutingMulticastDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchRoutingMulticast NetworksSwitchRoutingMulticast
	diags := req.Config.Get(ctx, &networksSwitchRoutingMulticast)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchRoutingMulticast")
		vvNetworkID := networksSwitchRoutingMulticast.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchRoutingMulticast(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchRoutingMulticast",
				err.Error(),
			)
			return
		}

		networksSwitchRoutingMulticast = ResponseSwitchGetNetworkSwitchRoutingMulticastItemToBody(networksSwitchRoutingMulticast, response1)
		diags = resp.State.Set(ctx, &networksSwitchRoutingMulticast)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchRoutingMulticast struct {
	NetworkID types.String                                    `tfsdk:"network_id"`
	Item      *ResponseSwitchGetNetworkSwitchRoutingMulticast `tfsdk:"item"`
}

type ResponseSwitchGetNetworkSwitchRoutingMulticast struct {
	DefaultSettings *ResponseSwitchGetNetworkSwitchRoutingMulticastDefaultSettings `tfsdk:"default_settings"`
	Overrides       *[]ResponseSwitchGetNetworkSwitchRoutingMulticastOverrides     `tfsdk:"overrides"`
}

type ResponseSwitchGetNetworkSwitchRoutingMulticastDefaultSettings struct {
	FloodUnknownMulticastTrafficEnabled types.Bool `tfsdk:"flood_unknown_multicast_traffic_enabled"`
	IgmpSnoopingEnabled                 types.Bool `tfsdk:"igmp_snooping_enabled"`
}

type ResponseSwitchGetNetworkSwitchRoutingMulticastOverrides struct {
	FloodUnknownMulticastTrafficEnabled types.Bool `tfsdk:"flood_unknown_multicast_traffic_enabled"`
	IgmpSnoopingEnabled                 types.Bool `tfsdk:"igmp_snooping_enabled"`
	Stacks                              types.List `tfsdk:"stacks"`
	SwitchProfiles                      types.List `tfsdk:"switch_profiles"`
	Switches                            types.List `tfsdk:"switches"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchRoutingMulticastItemToBody(state NetworksSwitchRoutingMulticast, response *merakigosdk.ResponseSwitchGetNetworkSwitchRoutingMulticast) NetworksSwitchRoutingMulticast {
	itemState := ResponseSwitchGetNetworkSwitchRoutingMulticast{
		DefaultSettings: func() *ResponseSwitchGetNetworkSwitchRoutingMulticastDefaultSettings {
			if response.DefaultSettings != nil {
				return &ResponseSwitchGetNetworkSwitchRoutingMulticastDefaultSettings{
					FloodUnknownMulticastTrafficEnabled: func() types.Bool {
						if response.DefaultSettings.FloodUnknownMulticastTrafficEnabled != nil {
							return types.BoolValue(*response.DefaultSettings.FloodUnknownMulticastTrafficEnabled)
						}
						return types.Bool{}
					}(),
					IgmpSnoopingEnabled: func() types.Bool {
						if response.DefaultSettings.IgmpSnoopingEnabled != nil {
							return types.BoolValue(*response.DefaultSettings.IgmpSnoopingEnabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseSwitchGetNetworkSwitchRoutingMulticastDefaultSettings{}
		}(),
		Overrides: func() *[]ResponseSwitchGetNetworkSwitchRoutingMulticastOverrides {
			if response.Overrides != nil {
				result := make([]ResponseSwitchGetNetworkSwitchRoutingMulticastOverrides, len(*response.Overrides))
				for i, overrides := range *response.Overrides {
					result[i] = ResponseSwitchGetNetworkSwitchRoutingMulticastOverrides{
						FloodUnknownMulticastTrafficEnabled: func() types.Bool {
							if overrides.FloodUnknownMulticastTrafficEnabled != nil {
								return types.BoolValue(*overrides.FloodUnknownMulticastTrafficEnabled)
							}
							return types.Bool{}
						}(),
						IgmpSnoopingEnabled: func() types.Bool {
							if overrides.IgmpSnoopingEnabled != nil {
								return types.BoolValue(*overrides.IgmpSnoopingEnabled)
							}
							return types.Bool{}
						}(),
						Stacks:         StringSliceToList(overrides.Stacks),
						SwitchProfiles: StringSliceToList(overrides.SwitchProfiles),
						Switches:       StringSliceToList(overrides.Switches),
					}
				}
				return &result
			}
			return &[]ResponseSwitchGetNetworkSwitchRoutingMulticastOverrides{}
		}(),
	}
	state.Item = &itemState
	return state
}
