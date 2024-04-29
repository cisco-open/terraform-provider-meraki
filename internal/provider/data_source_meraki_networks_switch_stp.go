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
	_ datasource.DataSource              = &NetworksSwitchStpDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchStpDataSource{}
)

func NewNetworksSwitchStpDataSource() datasource.DataSource {
	return &NetworksSwitchStpDataSource{}
}

type NetworksSwitchStpDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchStpDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchStpDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_stp"
}

func (d *NetworksSwitchStpDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"rstp_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"stp_bridge_priority": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"stp_priority": schema.Int64Attribute{
									Computed: true,
								},
								"switches": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchStpDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchStp NetworksSwitchStp
	diags := req.Config.Get(ctx, &networksSwitchStp)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchStp")
		vvNetworkID := networksSwitchStp.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchStp(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStp",
				err.Error(),
			)
			return
		}

		networksSwitchStp = ResponseSwitchGetNetworkSwitchStpItemToBody(networksSwitchStp, response1)
		diags = resp.State.Set(ctx, &networksSwitchStp)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchStp struct {
	NetworkID types.String                       `tfsdk:"network_id"`
	Item      *ResponseSwitchGetNetworkSwitchStp `tfsdk:"item"`
}

type ResponseSwitchGetNetworkSwitchStp struct {
	RstpEnabled       types.Bool                                            `tfsdk:"rstp_enabled"`
	StpBridgePriority *[]ResponseSwitchGetNetworkSwitchStpStpBridgePriority `tfsdk:"stp_bridge_priority"`
}

type ResponseSwitchGetNetworkSwitchStpStpBridgePriority struct {
	StpPriority types.Int64 `tfsdk:"stp_priority"`
	Switches    types.List  `tfsdk:"switches"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchStpItemToBody(state NetworksSwitchStp, response *merakigosdk.ResponseSwitchGetNetworkSwitchStp) NetworksSwitchStp {
	itemState := ResponseSwitchGetNetworkSwitchStp{
		RstpEnabled: func() types.Bool {
			if response.RstpEnabled != nil {
				return types.BoolValue(*response.RstpEnabled)
			}
			return types.Bool{}
		}(),
		StpBridgePriority: func() *[]ResponseSwitchGetNetworkSwitchStpStpBridgePriority {
			if response.StpBridgePriority != nil {
				result := make([]ResponseSwitchGetNetworkSwitchStpStpBridgePriority, len(*response.StpBridgePriority))
				for i, stpBridgePriority := range *response.StpBridgePriority {
					result[i] = ResponseSwitchGetNetworkSwitchStpStpBridgePriority{
						StpPriority: func() types.Int64 {
							if stpBridgePriority.StpPriority != nil {
								return types.Int64Value(int64(*stpBridgePriority.StpPriority))
							}
							return types.Int64{}
						}(),
						Switches: StringSliceToList(stpBridgePriority.Switches),
					}
				}
				return &result
			}
			return &[]ResponseSwitchGetNetworkSwitchStpStpBridgePriority{}
		}(),
	}
	state.Item = &itemState
	return state
}
