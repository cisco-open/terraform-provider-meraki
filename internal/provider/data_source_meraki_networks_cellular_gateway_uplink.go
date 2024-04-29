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
	_ datasource.DataSource              = &NetworksCellularGatewayUplinkDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksCellularGatewayUplinkDataSource{}
)

func NewNetworksCellularGatewayUplinkDataSource() datasource.DataSource {
	return &NetworksCellularGatewayUplinkDataSource{}
}

type NetworksCellularGatewayUplinkDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksCellularGatewayUplinkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksCellularGatewayUplinkDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_cellular_gateway_uplink"
}

func (d *NetworksCellularGatewayUplinkDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"bandwidth_limits": schema.SingleNestedAttribute{
						MarkdownDescription: `The bandwidth settings for the 'cellular' uplink`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"limit_down": schema.Int64Attribute{
								MarkdownDescription: `The maximum download limit (integer, in Kbps). 'null' indicates no limit.`,
								Computed:            true,
							},
							"limit_up": schema.Int64Attribute{
								MarkdownDescription: `The maximum upload limit (integer, in Kbps). 'null' indicates no limit.`,
								Computed:            true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksCellularGatewayUplinkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksCellularGatewayUplink NetworksCellularGatewayUplink
	diags := req.Config.Get(ctx, &networksCellularGatewayUplink)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkCellularGatewayUplink")
		vvNetworkID := networksCellularGatewayUplink.NetworkID.ValueString()

		response1, restyResp1, err := d.client.CellularGateway.GetNetworkCellularGatewayUplink(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkCellularGatewayUplink",
				err.Error(),
			)
			return
		}

		networksCellularGatewayUplink = ResponseCellularGatewayGetNetworkCellularGatewayUplinkItemToBody(networksCellularGatewayUplink, response1)
		diags = resp.State.Set(ctx, &networksCellularGatewayUplink)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksCellularGatewayUplink struct {
	NetworkID types.String                                            `tfsdk:"network_id"`
	Item      *ResponseCellularGatewayGetNetworkCellularGatewayUplink `tfsdk:"item"`
}

type ResponseCellularGatewayGetNetworkCellularGatewayUplink struct {
	BandwidthLimits *ResponseCellularGatewayGetNetworkCellularGatewayUplinkBandwidthLimits `tfsdk:"bandwidth_limits"`
}

type ResponseCellularGatewayGetNetworkCellularGatewayUplinkBandwidthLimits struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

// ToBody
func ResponseCellularGatewayGetNetworkCellularGatewayUplinkItemToBody(state NetworksCellularGatewayUplink, response *merakigosdk.ResponseCellularGatewayGetNetworkCellularGatewayUplink) NetworksCellularGatewayUplink {
	itemState := ResponseCellularGatewayGetNetworkCellularGatewayUplink{
		BandwidthLimits: func() *ResponseCellularGatewayGetNetworkCellularGatewayUplinkBandwidthLimits {
			if response.BandwidthLimits != nil {
				return &ResponseCellularGatewayGetNetworkCellularGatewayUplinkBandwidthLimits{
					LimitDown: func() types.Int64 {
						if response.BandwidthLimits.LimitDown != nil {
							return types.Int64Value(int64(*response.BandwidthLimits.LimitDown))
						}
						return types.Int64{}
					}(),
					LimitUp: func() types.Int64 {
						if response.BandwidthLimits.LimitUp != nil {
							return types.Int64Value(int64(*response.BandwidthLimits.LimitUp))
						}
						return types.Int64{}
					}(),
				}
			}
			return &ResponseCellularGatewayGetNetworkCellularGatewayUplinkBandwidthLimits{}
		}(),
	}
	state.Item = &itemState
	return state
}
