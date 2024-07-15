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
	_ datasource.DataSource              = &NetworksWirelessSSIDsFirewallL7FirewallRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessSSIDsFirewallL7FirewallRulesDataSource{}
)

func NewNetworksWirelessSSIDsFirewallL7FirewallRulesDataSource() datasource.DataSource {
	return &NetworksWirelessSSIDsFirewallL7FirewallRulesDataSource{}
}

type NetworksWirelessSSIDsFirewallL7FirewallRulesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessSSIDsFirewallL7FirewallRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessSSIDsFirewallL7FirewallRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_firewall_l7_firewall_rules"
}

func (d *NetworksWirelessSSIDsFirewallL7FirewallRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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

								"policy": schema.StringAttribute{
									MarkdownDescription: `'Deny' traffic specified by this rule`,
									Computed:            true,
								},
								"type": schema.StringAttribute{
									MarkdownDescription: `Type of the L7 firewall rule. One of: 'application', 'applicationCategory', 'host', 'port', 'ipRange'`,
									Computed:            true,
								},
								"value": schema.StringAttribute{
									MarkdownDescription: `The value of what needs to get blocked. Format of the value varies depending on type of the firewall rule selected.`,
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

func (d *NetworksWirelessSSIDsFirewallL7FirewallRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessSSIDsFirewallL7FirewallRules NetworksWirelessSSIDsFirewallL7FirewallRules
	diags := req.Config.Get(ctx, &networksWirelessSSIDsFirewallL7FirewallRules)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessSSIDFirewallL7FirewallRules")
		vvNetworkID := networksWirelessSSIDsFirewallL7FirewallRules.NetworkID.ValueString()
		vvNumber := networksWirelessSSIDsFirewallL7FirewallRules.Number.ValueString()

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessSSIDFirewallL7FirewallRules(vvNetworkID, vvNumber)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDFirewallL7FirewallRules",
				err.Error(),
			)
			return
		}

		networksWirelessSSIDsFirewallL7FirewallRules = ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRulesItemToBody(networksWirelessSSIDsFirewallL7FirewallRules, response1)
		diags = resp.State.Set(ctx, &networksWirelessSSIDsFirewallL7FirewallRules)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessSSIDsFirewallL7FirewallRules struct {
	NetworkID types.String                                                   `tfsdk:"network_id"`
	Number    types.String                                                   `tfsdk:"number"`
	Item      *ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRules `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRules struct {
	Rules *[]ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRulesRules `tfsdk:"rules"`
}

type ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRulesRules struct {
	Policy    types.String                                                              `tfsdk:"policy"`
	Type      types.String                                                              `tfsdk:"type"`
	Value     types.String                                                              `tfsdk:"value"`
	ValueList types.Set                                                                 `tfsdk:"value_list"`
	ValueObj  *ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj `tfsdk:"value_obj"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRulesItemToBody(state NetworksWirelessSSIDsFirewallL7FirewallRules, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRules) NetworksWirelessSSIDsFirewallL7FirewallRules {
	itemState := ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRules{
		Rules: func() *[]ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRulesRules {
			if response.Rules != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRulesRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRulesRules{
						Policy: types.StringValue(rules.Policy),
						Type:   types.StringValue(rules.Type),
						Value: func() types.String {
							if rules.Value == nil {
								return types.StringNull()
							}
							return types.StringValue(*rules.Value)
						}(),
						ValueList: func() types.Set {
							if rules.ValueList == nil {
								return types.SetNull(types.StringType)
							}
							return StringSliceToSet(*rules.ValueList)
						}(),
						ValueObj: func() *ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj {
							if rules.ValueObj == nil {
								return nil
							}
							return &ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj{
								ID:   types.StringValue(rules.ValueObj.ID),
								Name: types.StringValue(rules.ValueObj.Name),
							}
						}(),
					}
				}
				return &result
			}
			return &[]ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRulesRules{}
		}(),
	}
	state.Item = &itemState
	return state
}
