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
	_ datasource.DataSource              = &NetworksApplianceTrafficShapingUplinkBandwidthDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceTrafficShapingUplinkBandwidthDataSource{}
)

func NewNetworksApplianceTrafficShapingUplinkBandwidthDataSource() datasource.DataSource {
	return &NetworksApplianceTrafficShapingUplinkBandwidthDataSource{}
}

type NetworksApplianceTrafficShapingUplinkBandwidthDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceTrafficShapingUplinkBandwidthDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceTrafficShapingUplinkBandwidthDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_traffic_shaping_uplink_bandwidth"
}

func (d *NetworksApplianceTrafficShapingUplinkBandwidthDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
						MarkdownDescription: `A hash uplink keys and their configured settings for the Appliance`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"cellular": schema.SingleNestedAttribute{
								MarkdownDescription: `uplink cellular configured limits [optional]`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"limit_down": schema.Int64Attribute{
										MarkdownDescription: `configured DOWN limit for the uplink (in Kbps).  Null indicated unlimited`,
										Computed:            true,
									},
									"limit_up": schema.Int64Attribute{
										MarkdownDescription: `configured UP limit for the uplink (in Kbps).  Null indicated unlimited`,
										Computed:            true,
									},
								},
							},
							"wan1": schema.SingleNestedAttribute{
								MarkdownDescription: `uplink wan1 configured limits [optional]`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"limit_down": schema.Int64Attribute{
										MarkdownDescription: `configured DOWN limit for the uplink (in Kbps).  Null indicated unlimited`,
										Computed:            true,
									},
									"limit_up": schema.Int64Attribute{
										MarkdownDescription: `configured UP limit for the uplink (in Kbps).  Null indicated unlimited`,
										Computed:            true,
									},
								},
							},
							"wan2": schema.SingleNestedAttribute{
								MarkdownDescription: `uplink wan2 configured limits [optional]`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"limit_down": schema.Int64Attribute{
										MarkdownDescription: `configured DOWN limit for the uplink (in Kbps).  Null indicated unlimited`,
										Computed:            true,
									},
									"limit_up": schema.Int64Attribute{
										MarkdownDescription: `configured UP limit for the uplink (in Kbps).  Null indicated unlimited`,
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksApplianceTrafficShapingUplinkBandwidthDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceTrafficShapingUplinkBandwidth NetworksApplianceTrafficShapingUplinkBandwidth
	diags := req.Config.Get(ctx, &networksApplianceTrafficShapingUplinkBandwidth)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceTrafficShapingUplinkBandwidth")
		vvNetworkID := networksApplianceTrafficShapingUplinkBandwidth.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceTrafficShapingUplinkBandwidth(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceTrafficShapingUplinkBandwidth",
				err.Error(),
			)
			return
		}

		networksApplianceTrafficShapingUplinkBandwidth = ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthItemToBody(networksApplianceTrafficShapingUplinkBandwidth, response1)
		diags = resp.State.Set(ctx, &networksApplianceTrafficShapingUplinkBandwidth)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceTrafficShapingUplinkBandwidth struct {
	NetworkID types.String                                                       `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidth `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidth struct {
	BandwidthLimits *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimits `tfsdk:"bandwidth_limits"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimits struct {
	Cellular *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular `tfsdk:"cellular"`
	Wan1     *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1     `tfsdk:"wan1"`
	Wan2     *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2     `tfsdk:"wan2"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1 struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2 struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthItemToBody(state NetworksApplianceTrafficShapingUplinkBandwidth, response *merakigosdk.ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidth) NetworksApplianceTrafficShapingUplinkBandwidth {
	itemState := ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidth{
		BandwidthLimits: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimits {
			if response.BandwidthLimits != nil {
				return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimits{
					Cellular: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular {
						if response.BandwidthLimits.Cellular != nil {
							return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular{
								LimitDown: func() types.Int64 {
									if response.BandwidthLimits.Cellular.LimitDown != nil {
										return types.Int64Value(int64(*response.BandwidthLimits.Cellular.LimitDown))
									}
									return types.Int64{}
								}(),
								LimitUp: func() types.Int64 {
									if response.BandwidthLimits.Cellular.LimitUp != nil {
										return types.Int64Value(int64(*response.BandwidthLimits.Cellular.LimitUp))
									}
									return types.Int64{}
								}(),
							}
						}
						return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular{}
					}(),
					Wan1: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1 {
						if response.BandwidthLimits.Wan1 != nil {
							return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1{
								LimitDown: func() types.Int64 {
									if response.BandwidthLimits.Wan1.LimitDown != nil {
										return types.Int64Value(int64(*response.BandwidthLimits.Wan1.LimitDown))
									}
									return types.Int64{}
								}(),
								LimitUp: func() types.Int64 {
									if response.BandwidthLimits.Wan1.LimitUp != nil {
										return types.Int64Value(int64(*response.BandwidthLimits.Wan1.LimitUp))
									}
									return types.Int64{}
								}(),
							}
						}
						return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1{}
					}(),
					Wan2: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2 {
						if response.BandwidthLimits.Wan2 != nil {
							return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2{
								LimitDown: func() types.Int64 {
									if response.BandwidthLimits.Wan2.LimitDown != nil {
										return types.Int64Value(int64(*response.BandwidthLimits.Wan2.LimitDown))
									}
									return types.Int64{}
								}(),
								LimitUp: func() types.Int64 {
									if response.BandwidthLimits.Wan2.LimitUp != nil {
										return types.Int64Value(int64(*response.BandwidthLimits.Wan2.LimitUp))
									}
									return types.Int64{}
								}(),
							}
						}
						return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2{}
					}(),
				}
			}
			return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimits{}
		}(),
	}
	state.Item = &itemState
	return state
}
