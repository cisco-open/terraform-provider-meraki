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
	_ datasource.DataSource              = &NetworksSwitchSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchSettingsDataSource{}
)

func NewNetworksSwitchSettingsDataSource() datasource.DataSource {
	return &NetworksSwitchSettingsDataSource{}
}

type NetworksSwitchSettingsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchSettingsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_settings"
}

func (d *NetworksSwitchSettingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"mac_blocklist": schema.SingleNestedAttribute{
						MarkdownDescription: `MAC blocklist`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Enable MAC blocklist for switches in the network`,
								Computed:            true,
							},
						},
					},
					"power_exceptions": schema.SetNestedAttribute{
						MarkdownDescription: `Exceptions on a per switch basis to "useCombinedPower"`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"power_type": schema.StringAttribute{
									MarkdownDescription: `Per switch exception (combined, redundant, useNetworkSetting)`,
									Computed:            true,
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `Serial number of the switch`,
									Computed:            true,
								},
							},
						},
					},
					"uplink_client_sampling": schema.SingleNestedAttribute{
						MarkdownDescription: `Uplink client sampling`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Enable client sampling on uplink`,
								Computed:            true,
							},
						},
					},
					"use_combined_power": schema.BoolAttribute{
						MarkdownDescription: `The use Combined Power as the default behavior of secondary power supplies on supported devices.`,
						Computed:            true,
					},
					"vlan": schema.Int64Attribute{
						MarkdownDescription: `Management VLAN`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchSettings NetworksSwitchSettings
	diags := req.Config.Get(ctx, &networksSwitchSettings)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchSettings")
		vvNetworkID := networksSwitchSettings.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchSettings(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchSettings",
				err.Error(),
			)
			return
		}

		networksSwitchSettings = ResponseSwitchGetNetworkSwitchSettingsItemToBody(networksSwitchSettings, response1)
		diags = resp.State.Set(ctx, &networksSwitchSettings)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchSettings struct {
	NetworkID types.String                            `tfsdk:"network_id"`
	Item      *ResponseSwitchGetNetworkSwitchSettings `tfsdk:"item"`
}

type ResponseSwitchGetNetworkSwitchSettings struct {
	MacBlocklist         *ResponseSwitchGetNetworkSwitchSettingsMacBlocklist         `tfsdk:"mac_blocklist"`
	PowerExceptions      *[]ResponseSwitchGetNetworkSwitchSettingsPowerExceptions    `tfsdk:"power_exceptions"`
	UplinkClientSampling *ResponseSwitchGetNetworkSwitchSettingsUplinkClientSampling `tfsdk:"uplink_client_sampling"`
	UseCombinedPower     types.Bool                                                  `tfsdk:"use_combined_power"`
	VLAN                 types.Int64                                                 `tfsdk:"vlan"`
}

type ResponseSwitchGetNetworkSwitchSettingsMacBlocklist struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseSwitchGetNetworkSwitchSettingsPowerExceptions struct {
	PowerType types.String `tfsdk:"power_type"`
	Serial    types.String `tfsdk:"serial"`
}

type ResponseSwitchGetNetworkSwitchSettingsUplinkClientSampling struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchSettingsItemToBody(state NetworksSwitchSettings, response *merakigosdk.ResponseSwitchGetNetworkSwitchSettings) NetworksSwitchSettings {
	itemState := ResponseSwitchGetNetworkSwitchSettings{
		MacBlocklist: func() *ResponseSwitchGetNetworkSwitchSettingsMacBlocklist {
			if response.MacBlocklist != nil {
				return &ResponseSwitchGetNetworkSwitchSettingsMacBlocklist{
					Enabled: func() types.Bool {
						if response.MacBlocklist.Enabled != nil {
							return types.BoolValue(*response.MacBlocklist.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseSwitchGetNetworkSwitchSettingsMacBlocklist{}
		}(),
		PowerExceptions: func() *[]ResponseSwitchGetNetworkSwitchSettingsPowerExceptions {
			if response.PowerExceptions != nil {
				result := make([]ResponseSwitchGetNetworkSwitchSettingsPowerExceptions, len(*response.PowerExceptions))
				for i, powerExceptions := range *response.PowerExceptions {
					result[i] = ResponseSwitchGetNetworkSwitchSettingsPowerExceptions{
						PowerType: types.StringValue(powerExceptions.PowerType),
						Serial:    types.StringValue(powerExceptions.Serial),
					}
				}
				return &result
			}
			return &[]ResponseSwitchGetNetworkSwitchSettingsPowerExceptions{}
		}(),
		UplinkClientSampling: func() *ResponseSwitchGetNetworkSwitchSettingsUplinkClientSampling {
			if response.UplinkClientSampling != nil {
				return &ResponseSwitchGetNetworkSwitchSettingsUplinkClientSampling{
					Enabled: func() types.Bool {
						if response.UplinkClientSampling.Enabled != nil {
							return types.BoolValue(*response.UplinkClientSampling.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseSwitchGetNetworkSwitchSettingsUplinkClientSampling{}
		}(),
		UseCombinedPower: func() types.Bool {
			if response.UseCombinedPower != nil {
				return types.BoolValue(*response.UseCombinedPower)
			}
			return types.Bool{}
		}(),
		VLAN: func() types.Int64 {
			if response.VLAN != nil {
				return types.Int64Value(int64(*response.VLAN))
			}
			return types.Int64{}
		}(),
	}
	state.Item = &itemState
	return state
}
