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
	_ datasource.DataSource              = &NetworksWirelessSSIDsFirewallL3FirewallRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessSSIDsFirewallL3FirewallRulesDataSource{}
)

func NewNetworksWirelessSSIDsFirewallL3FirewallRulesDataSource() datasource.DataSource {
	return &NetworksWirelessSSIDsFirewallL3FirewallRulesDataSource{}
}

type NetworksWirelessSSIDsFirewallL3FirewallRulesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessSSIDsFirewallL3FirewallRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessSSIDsFirewallL3FirewallRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_firewall_l3_firewall_rules"
}

func (d *NetworksWirelessSSIDsFirewallL3FirewallRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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

					"rules": schema.SetNestedAttribute{
						MarkdownDescription: `An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule).`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"comment": schema.StringAttribute{
									MarkdownDescription: `Description of the rule (optional)`,
									Computed:            true,
								},
								"dest_cidr": schema.StringAttribute{
									MarkdownDescription: `Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'`,
									Computed:            true,
								},
								"dest_port": schema.StringAttribute{
									MarkdownDescription: `Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'`,
									Computed:            true,
								},
								"policy": schema.StringAttribute{
									MarkdownDescription: `'allow' or 'deny' traffic specified by this rule`,
									Computed:            true,
								},
								"protocol": schema.StringAttribute{
									MarkdownDescription: `The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')`,
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

func (d *NetworksWirelessSSIDsFirewallL3FirewallRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessSSIDsFirewallL3FirewallRules NetworksWirelessSSIDsFirewallL3FirewallRules
	diags := req.Config.Get(ctx, &networksWirelessSSIDsFirewallL3FirewallRules)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessSSIDFirewallL3FirewallRules")
		vvNetworkID := networksWirelessSSIDsFirewallL3FirewallRules.NetworkID.ValueString()
		vvNumber := networksWirelessSSIDsFirewallL3FirewallRules.Number.ValueString()

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessSSIDFirewallL3FirewallRules(vvNetworkID, vvNumber)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDFirewallL3FirewallRules",
				err.Error(),
			)
			return
		}

		networksWirelessSSIDsFirewallL3FirewallRules = ResponseWirelessGetNetworkWirelessSSIDFirewallL3FirewallRulesItemToBody(networksWirelessSSIDsFirewallL3FirewallRules, response1)
		diags = resp.State.Set(ctx, &networksWirelessSSIDsFirewallL3FirewallRules)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessSSIDsFirewallL3FirewallRules struct {
	NetworkID types.String                                                   `tfsdk:"network_id"`
	Number    types.String                                                   `tfsdk:"number"`
	Item      *ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRules `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRules struct {
	Rules *[]ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRulesRules `tfsdk:"rules"`
}

type ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRulesRules struct {
	Comment  types.String `tfsdk:"comment"`
	DestCidr types.String `tfsdk:"dest_cidr"`
	DestPort types.String `tfsdk:"dest_port"`
	Policy   types.String `tfsdk:"policy"`
	Protocol types.String `tfsdk:"protocol"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessSSIDFirewallL3FirewallRulesItemToBody(state NetworksWirelessSSIDsFirewallL3FirewallRules, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDFirewallL3FirewallRules) NetworksWirelessSSIDsFirewallL3FirewallRules {
	itemState := ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRules{
		Rules: func() *[]ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRulesRules {
			if response.Rules != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRulesRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRulesRules{
						Comment:  types.StringValue(rules.Comment),
						DestCidr: types.StringValue(rules.DestCidr),
						DestPort: types.StringValue(rules.DestPort),
						Policy:   types.StringValue(rules.Policy),
						Protocol: types.StringValue(rules.Protocol),
					}
				}
				return &result
			}
			return &[]ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRulesRules{}
		}(),
	}
	state.Item = &itemState
	return state
}
