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
	_ datasource.DataSource              = &NetworksApplianceVpnBgpDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceVpnBgpDataSource{}
)

func NewNetworksApplianceVpnBgpDataSource() datasource.DataSource {
	return &NetworksApplianceVpnBgpDataSource{}
}

type NetworksApplianceVpnBgpDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceVpnBgpDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceVpnBgpDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_vpn_bgp"
}

func (d *NetworksApplianceVpnBgpDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"as_number": schema.Int64Attribute{
						Computed: true,
					},
					"enabled": schema.BoolAttribute{
						Computed: true,
					},
					"ibgp_hold_timer": schema.Int64Attribute{
						Computed: true,
					},
					"neighbors": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"allow_transit": schema.BoolAttribute{
									Computed: true,
								},
								"ebgp_hold_timer": schema.Int64Attribute{
									Computed: true,
								},
								"ebgp_multihop": schema.Int64Attribute{
									Computed: true,
								},
								"ip": schema.StringAttribute{
									Computed: true,
								},
								"receive_limit": schema.Int64Attribute{
									Computed: true,
								},
								"remote_as_number": schema.Int64Attribute{
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

func (d *NetworksApplianceVpnBgpDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceVpnBgp NetworksApplianceVpnBgp
	diags := req.Config.Get(ctx, &networksApplianceVpnBgp)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceVpnBgp")
		vvNetworkID := networksApplianceVpnBgp.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceVpnBgp(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceVpnBgp",
				err.Error(),
			)
			return
		}

		networksApplianceVpnBgp = ResponseApplianceGetNetworkApplianceVpnBgpItemToBody(networksApplianceVpnBgp, response1)
		diags = resp.State.Set(ctx, &networksApplianceVpnBgp)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceVpnBgp struct {
	NetworkID types.String                                `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceVpnBgp `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceVpnBgp struct {
	AsNumber      types.Int64                                            `tfsdk:"as_number"`
	Enabled       types.Bool                                             `tfsdk:"enabled"`
	IbgpHoldTimer types.Int64                                            `tfsdk:"ibgp_hold_timer"`
	Neighbors     *[]ResponseApplianceGetNetworkApplianceVpnBgpNeighbors `tfsdk:"neighbors"`
}

type ResponseApplianceGetNetworkApplianceVpnBgpNeighbors struct {
	AllowTransit   types.Bool   `tfsdk:"allow_transit"`
	EbgpHoldTimer  types.Int64  `tfsdk:"ebgp_hold_timer"`
	EbgpMultihop   types.Int64  `tfsdk:"ebgp_multihop"`
	IP             types.String `tfsdk:"ip"`
	ReceiveLimit   types.Int64  `tfsdk:"receive_limit"`
	RemoteAsNumber types.Int64  `tfsdk:"remote_as_number"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceVpnBgpItemToBody(state NetworksApplianceVpnBgp, response *merakigosdk.ResponseApplianceGetNetworkApplianceVpnBgp) NetworksApplianceVpnBgp {
	itemState := ResponseApplianceGetNetworkApplianceVpnBgp{
		AsNumber: func() types.Int64 {
			if response.AsNumber != nil {
				return types.Int64Value(int64(*response.AsNumber))
			}
			return types.Int64{}
		}(),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		IbgpHoldTimer: func() types.Int64 {
			if response.IbgpHoldTimer != nil {
				return types.Int64Value(int64(*response.IbgpHoldTimer))
			}
			return types.Int64{}
		}(),
		Neighbors: func() *[]ResponseApplianceGetNetworkApplianceVpnBgpNeighbors {
			if response.Neighbors != nil {
				result := make([]ResponseApplianceGetNetworkApplianceVpnBgpNeighbors, len(*response.Neighbors))
				for i, neighbors := range *response.Neighbors {
					result[i] = ResponseApplianceGetNetworkApplianceVpnBgpNeighbors{
						AllowTransit: func() types.Bool {
							if neighbors.AllowTransit != nil {
								return types.BoolValue(*neighbors.AllowTransit)
							}
							return types.Bool{}
						}(),
						EbgpHoldTimer: func() types.Int64 {
							if neighbors.EbgpHoldTimer != nil {
								return types.Int64Value(int64(*neighbors.EbgpHoldTimer))
							}
							return types.Int64{}
						}(),
						EbgpMultihop: func() types.Int64 {
							if neighbors.EbgpMultihop != nil {
								return types.Int64Value(int64(*neighbors.EbgpMultihop))
							}
							return types.Int64{}
						}(),
						IP: types.StringValue(neighbors.IP),
						ReceiveLimit: func() types.Int64 {
							if neighbors.ReceiveLimit != nil {
								return types.Int64Value(int64(*neighbors.ReceiveLimit))
							}
							return types.Int64{}
						}(),
						RemoteAsNumber: func() types.Int64 {
							if neighbors.RemoteAsNumber != nil {
								return types.Int64Value(int64(*neighbors.RemoteAsNumber))
							}
							return types.Int64{}
						}(),
					}
				}
				return &result
			}
			return &[]ResponseApplianceGetNetworkApplianceVpnBgpNeighbors{}
		}(),
	}
	state.Item = &itemState
	return state
}
