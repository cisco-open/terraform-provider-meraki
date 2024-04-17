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
	_ datasource.DataSource              = &NetworksApplianceFirewallOneToOneNatRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceFirewallOneToOneNatRulesDataSource{}
)

func NewNetworksApplianceFirewallOneToOneNatRulesDataSource() datasource.DataSource {
	return &NetworksApplianceFirewallOneToOneNatRulesDataSource{}
}

type NetworksApplianceFirewallOneToOneNatRulesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceFirewallOneToOneNatRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceFirewallOneToOneNatRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_firewall_one_to_one_nat_rules"
}

func (d *NetworksApplianceFirewallOneToOneNatRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"rules": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"allowed_inbound": schema.SetNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"allowed_ips": schema.ListAttribute{
												Computed:    true,
												ElementType: types.StringType,
											},
											"destination_ports": schema.ListAttribute{
												Computed:    true,
												ElementType: types.StringType,
											},
											"protocol": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
								"lan_ip": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"public_ip": schema.StringAttribute{
									Computed: true,
								},
								"uplink": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksApplianceFirewallOneToOneNatRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceFirewallOneToOneNatRules NetworksApplianceFirewallOneToOneNatRules
	diags := req.Config.Get(ctx, &networksApplianceFirewallOneToOneNatRules)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceFirewallOneToOneNatRules")
		vvNetworkID := networksApplianceFirewallOneToOneNatRules.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceFirewallOneToOneNatRules(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceFirewallOneToOneNatRules",
				err.Error(),
			)
			return
		}

		networksApplianceFirewallOneToOneNatRules = ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesItemToBody(networksApplianceFirewallOneToOneNatRules, response1)
		diags = resp.State.Set(ctx, &networksApplianceFirewallOneToOneNatRules)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceFirewallOneToOneNatRules struct {
	NetworkID types.String                                                  `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRules `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRules struct {
	Rules *[]ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRules `tfsdk:"rules"`
}

type ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRules struct {
	AllowedInbound *[]ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInbound `tfsdk:"allowed_inbound"`
	LanIP          types.String                                                                       `tfsdk:"lan_ip"`
	Name           types.String                                                                       `tfsdk:"name"`
	PublicIP       types.String                                                                       `tfsdk:"public_ip"`
	Uplink         types.String                                                                       `tfsdk:"uplink"`
}

type ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInbound struct {
	AllowedIPs       types.List   `tfsdk:"allowed_ips"`
	DestinationPorts types.List   `tfsdk:"destination_ports"`
	Protocol         types.String `tfsdk:"protocol"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesItemToBody(state NetworksApplianceFirewallOneToOneNatRules, response *merakigosdk.ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRules) NetworksApplianceFirewallOneToOneNatRules {
	itemState := ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRules{
		Rules: func() *[]ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRules {
			if response.Rules != nil {
				result := make([]ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRules{
						AllowedInbound: func() *[]ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInbound {
							if rules.AllowedInbound != nil {
								result := make([]ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInbound, len(*rules.AllowedInbound))
								for i, allowedInbound := range *rules.AllowedInbound {
									result[i] = ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInbound{
										AllowedIPs:       StringSliceToList(allowedInbound.AllowedIPs),
										DestinationPorts: StringSliceToList(allowedInbound.DestinationPorts),
										Protocol:         types.StringValue(allowedInbound.Protocol),
									}
								}
								return &result
							}
							return &[]ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInbound{}
						}(),
						LanIP:    types.StringValue(rules.LanIP),
						Name:     types.StringValue(rules.Name),
						PublicIP: types.StringValue(rules.PublicIP),
						Uplink:   types.StringValue(rules.Uplink),
					}
				}
				return &result
			}
			return &[]ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRules{}
		}(),
	}
	state.Item = &itemState
	return state
}
