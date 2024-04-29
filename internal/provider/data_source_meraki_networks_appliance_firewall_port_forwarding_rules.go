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
	_ datasource.DataSource              = &NetworksApplianceFirewallPortForwardingRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceFirewallPortForwardingRulesDataSource{}
)

func NewNetworksApplianceFirewallPortForwardingRulesDataSource() datasource.DataSource {
	return &NetworksApplianceFirewallPortForwardingRulesDataSource{}
}

type NetworksApplianceFirewallPortForwardingRulesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceFirewallPortForwardingRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceFirewallPortForwardingRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_firewall_port_forwarding_rules"
}

func (d *NetworksApplianceFirewallPortForwardingRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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

								"allowed_ips": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
								"lan_ip": schema.StringAttribute{
									Computed: true,
								},
								"local_port": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"protocol": schema.StringAttribute{
									Computed: true,
								},
								"public_port": schema.StringAttribute{
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

func (d *NetworksApplianceFirewallPortForwardingRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceFirewallPortForwardingRules NetworksApplianceFirewallPortForwardingRules
	diags := req.Config.Get(ctx, &networksApplianceFirewallPortForwardingRules)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceFirewallPortForwardingRules")
		vvNetworkID := networksApplianceFirewallPortForwardingRules.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceFirewallPortForwardingRules(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceFirewallPortForwardingRules",
				err.Error(),
			)
			return
		}

		networksApplianceFirewallPortForwardingRules = ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesItemToBody(networksApplianceFirewallPortForwardingRules, response1)
		diags = resp.State.Set(ctx, &networksApplianceFirewallPortForwardingRules)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceFirewallPortForwardingRules struct {
	NetworkID types.String                                                     `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceFirewallPortForwardingRules `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceFirewallPortForwardingRules struct {
	Rules *[]ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRules `tfsdk:"rules"`
}

type ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRules struct {
	AllowedIPs types.List   `tfsdk:"allowed_ips"`
	LanIP      types.String `tfsdk:"lan_ip"`
	LocalPort  types.String `tfsdk:"local_port"`
	Name       types.String `tfsdk:"name"`
	Protocol   types.String `tfsdk:"protocol"`
	PublicPort types.String `tfsdk:"public_port"`
	Uplink     types.String `tfsdk:"uplink"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesItemToBody(state NetworksApplianceFirewallPortForwardingRules, response *merakigosdk.ResponseApplianceGetNetworkApplianceFirewallPortForwardingRules) NetworksApplianceFirewallPortForwardingRules {
	itemState := ResponseApplianceGetNetworkApplianceFirewallPortForwardingRules{
		Rules: func() *[]ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRules {
			if response.Rules != nil {
				result := make([]ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRules{
						AllowedIPs: StringSliceToList(rules.AllowedIPs),
						LanIP:      types.StringValue(rules.LanIP),
						LocalPort:  types.StringValue(rules.LocalPort),
						Name:       types.StringValue(rules.Name),
						Protocol:   types.StringValue(rules.Protocol),
						PublicPort: types.StringValue(rules.PublicPort),
						Uplink:     types.StringValue(rules.Uplink),
					}
				}
				return &result
			}
			return &[]ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRules{}
		}(),
	}
	state.Item = &itemState
	return state
}
