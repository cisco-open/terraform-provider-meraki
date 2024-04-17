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
	_ datasource.DataSource              = &NetworksApplianceFirewallCellularFirewallRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceFirewallCellularFirewallRulesDataSource{}
)

func NewNetworksApplianceFirewallCellularFirewallRulesDataSource() datasource.DataSource {
	return &NetworksApplianceFirewallCellularFirewallRulesDataSource{}
}

type NetworksApplianceFirewallCellularFirewallRulesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceFirewallCellularFirewallRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceFirewallCellularFirewallRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_firewall_cellular_firewall_rules"
}

func (d *NetworksApplianceFirewallCellularFirewallRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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

								"comment": schema.StringAttribute{
									Computed: true,
								},
								"dest_cidr": schema.StringAttribute{
									Computed: true,
								},
								"dest_port": schema.StringAttribute{
									Computed: true,
								},
								"policy": schema.StringAttribute{
									Computed: true,
								},
								"protocol": schema.StringAttribute{
									Computed: true,
								},
								"src_cidr": schema.StringAttribute{
									Computed: true,
								},
								"src_port": schema.StringAttribute{
									Computed: true,
								},
								"syslog_enabled": schema.BoolAttribute{
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

func (d *NetworksApplianceFirewallCellularFirewallRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceFirewallCellularFirewallRules NetworksApplianceFirewallCellularFirewallRules
	diags := req.Config.Get(ctx, &networksApplianceFirewallCellularFirewallRules)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceFirewallCellularFirewallRules")
		vvNetworkID := networksApplianceFirewallCellularFirewallRules.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceFirewallCellularFirewallRules(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceFirewallCellularFirewallRules",
				err.Error(),
			)
			return
		}

		networksApplianceFirewallCellularFirewallRules = ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRulesItemToBody(networksApplianceFirewallCellularFirewallRules, response1)
		diags = resp.State.Set(ctx, &networksApplianceFirewallCellularFirewallRules)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceFirewallCellularFirewallRules struct {
	NetworkID types.String                                                       `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRules `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRules struct {
	Rules *[]ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRulesRules `tfsdk:"rules"`
}

type ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRulesRules struct {
	Comment       types.String `tfsdk:"comment"`
	DestCidr      types.String `tfsdk:"dest_cidr"`
	DestPort      types.String `tfsdk:"dest_port"`
	Policy        types.String `tfsdk:"policy"`
	Protocol      types.String `tfsdk:"protocol"`
	SrcCidr       types.String `tfsdk:"src_cidr"`
	SrcPort       types.String `tfsdk:"src_port"`
	SyslogEnabled types.Bool   `tfsdk:"syslog_enabled"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRulesItemToBody(state NetworksApplianceFirewallCellularFirewallRules, response *merakigosdk.ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRules) NetworksApplianceFirewallCellularFirewallRules {
	itemState := ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRules{
		Rules: func() *[]ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRulesRules {
			if response.Rules != nil {
				result := make([]ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRulesRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRulesRules{
						Comment:  types.StringValue(rules.Comment),
						DestCidr: types.StringValue(rules.DestCidr),
						DestPort: types.StringValue(rules.DestPort),
						Policy:   types.StringValue(rules.Policy),
						Protocol: types.StringValue(rules.Protocol),
						SrcCidr:  types.StringValue(rules.SrcCidr),
						SrcPort:  types.StringValue(rules.SrcPort),
						SyslogEnabled: func() types.Bool {
							if rules.SyslogEnabled != nil {
								return types.BoolValue(*rules.SyslogEnabled)
							}
							return types.Bool{}
						}(),
					}
				}
				return &result
			}
			return &[]ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRulesRules{}
		}(),
	}
	state.Item = &itemState
	return state
}
