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
	_ datasource.DataSource              = &NetworksWirelessSSIDsBonjourForwardingDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessSSIDsBonjourForwardingDataSource{}
)

func NewNetworksWirelessSSIDsBonjourForwardingDataSource() datasource.DataSource {
	return &NetworksWirelessSSIDsBonjourForwardingDataSource{}
}

type NetworksWirelessSSIDsBonjourForwardingDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessSSIDsBonjourForwardingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessSSIDsBonjourForwardingDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_bonjour_forwarding"
}

func (d *NetworksWirelessSSIDsBonjourForwardingDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: `number path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `If true, Bonjour forwarding is enabled on the SSID.`,
						Computed:            true,
					},
					"exception": schema.SingleNestedAttribute{
						MarkdownDescription: `Bonjour forwarding exception`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `If true, Bonjour forwarding exception is enabled on this SSID. Exception is required to enable L2 isolation and Bonjour forwarding to work together.`,
								Computed:            true,
							},
						},
					},
					"rules": schema.SetNestedAttribute{
						MarkdownDescription: `Bonjour forwarding rules`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"description": schema.StringAttribute{
									MarkdownDescription: `Desctiption of the bonjour forwarding rule`,
									Computed:            true,
								},
								"services": schema.ListAttribute{
									MarkdownDescription: `A list of Bonjour services. At least one service must be specified. Available services are 'All Services', 'AirPlay', 'AFP', 'BitTorrent', 'FTP', 'iChat', 'iTunes', 'Printers', 'Samba', 'Scanners' and 'SSH'`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"vlan_id": schema.StringAttribute{
									MarkdownDescription: `The ID of the service VLAN. Required`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessSSIDsBonjourForwardingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessSSIDsBonjourForwarding NetworksWirelessSSIDsBonjourForwarding
	diags := req.Config.Get(ctx, &networksWirelessSSIDsBonjourForwarding)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessSSIDBonjourForwarding")
		vvNetworkID := networksWirelessSSIDsBonjourForwarding.NetworkID.ValueString()
		vvNumber := networksWirelessSSIDsBonjourForwarding.Number.ValueString()

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessSSIDBonjourForwarding(vvNetworkID, vvNumber)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDBonjourForwarding",
				err.Error(),
			)
			return
		}

		networksWirelessSSIDsBonjourForwarding = ResponseWirelessGetNetworkWirelessSSIDBonjourForwardingItemToBody(networksWirelessSSIDsBonjourForwarding, response1)
		diags = resp.State.Set(ctx, &networksWirelessSSIDsBonjourForwarding)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessSSIDsBonjourForwarding struct {
	NetworkID types.String                                             `tfsdk:"network_id"`
	Number    types.String                                             `tfsdk:"number"`
	Item      *ResponseWirelessGetNetworkWirelessSsidBonjourForwarding `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessSsidBonjourForwarding struct {
	Enabled   types.Bool                                                        `tfsdk:"enabled"`
	Exception *ResponseWirelessGetNetworkWirelessSsidBonjourForwardingException `tfsdk:"exception"`
	Rules     *[]ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRules   `tfsdk:"rules"`
}

type ResponseWirelessGetNetworkWirelessSsidBonjourForwardingException struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRules struct {
	Description types.String `tfsdk:"description"`
	Services    types.List   `tfsdk:"services"`
	VLANID      types.String `tfsdk:"vlan_id"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessSSIDBonjourForwardingItemToBody(state NetworksWirelessSSIDsBonjourForwarding, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDBonjourForwarding) NetworksWirelessSSIDsBonjourForwarding {
	itemState := ResponseWirelessGetNetworkWirelessSsidBonjourForwarding{
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		Exception: func() *ResponseWirelessGetNetworkWirelessSsidBonjourForwardingException {
			if response.Exception != nil {
				return &ResponseWirelessGetNetworkWirelessSsidBonjourForwardingException{
					Enabled: func() types.Bool {
						if response.Exception.Enabled != nil {
							return types.BoolValue(*response.Exception.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseWirelessGetNetworkWirelessSsidBonjourForwardingException{}
		}(),
		Rules: func() *[]ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRules {
			if response.Rules != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRules{
						Description: types.StringValue(rules.Description),
						Services:    StringSliceToList(rules.Services),
						VLANID:      types.StringValue(rules.VLANID),
					}
				}
				return &result
			}
			return &[]ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRules{}
		}(),
	}
	state.Item = &itemState
	return state
}
