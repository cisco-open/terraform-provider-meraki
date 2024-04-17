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
	_ datasource.DataSource              = &NetworksTrafficAnalysisDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksTrafficAnalysisDataSource{}
)

func NewNetworksTrafficAnalysisDataSource() datasource.DataSource {
	return &NetworksTrafficAnalysisDataSource{}
}

type NetworksTrafficAnalysisDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksTrafficAnalysisDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksTrafficAnalysisDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_traffic_analysis"
}

func (d *NetworksTrafficAnalysisDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"custom_pie_chart_items": schema.SetNestedAttribute{
						MarkdownDescription: `The list of items that make up the custom pie chart for traffic reporting.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"name": schema.StringAttribute{
									MarkdownDescription: `The name of the custom pie chart item.`,
									Computed:            true,
								},
								"type": schema.StringAttribute{
									MarkdownDescription: `    The signature type for the custom pie chart item. Can be one of 'host', 'port' or 'ipRange'.
`,
									Computed: true,
								},
								"value": schema.StringAttribute{
									MarkdownDescription: `    The value of the custom pie chart item. Valid syntax depends on the signature type of the chart item
    (see sample request/response for more details).
`,
									Computed: true,
								},
							},
						},
					},
					"mode": schema.StringAttribute{
						MarkdownDescription: `    The traffic analysis mode for the network. Can be one of 'disabled' (do not collect traffic types),
    'basic' (collect generic traffic categories), or 'detailed' (collect destination hostnames).
`,
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *NetworksTrafficAnalysisDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksTrafficAnalysis NetworksTrafficAnalysis
	diags := req.Config.Get(ctx, &networksTrafficAnalysis)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkTrafficAnalysis")
		vvNetworkID := networksTrafficAnalysis.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Networks.GetNetworkTrafficAnalysis(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkTrafficAnalysis",
				err.Error(),
			)
			return
		}

		networksTrafficAnalysis = ResponseNetworksGetNetworkTrafficAnalysisItemToBody(networksTrafficAnalysis, response1)
		diags = resp.State.Set(ctx, &networksTrafficAnalysis)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksTrafficAnalysis struct {
	NetworkID types.String                               `tfsdk:"network_id"`
	Item      *ResponseNetworksGetNetworkTrafficAnalysis `tfsdk:"item"`
}

type ResponseNetworksGetNetworkTrafficAnalysis struct {
	CustomPieChartItems *[]ResponseNetworksGetNetworkTrafficAnalysisCustomPieChartItems `tfsdk:"custom_pie_chart_items"`
	Mode                types.String                                                    `tfsdk:"mode"`
}

type ResponseNetworksGetNetworkTrafficAnalysisCustomPieChartItems struct {
	Name  types.String `tfsdk:"name"`
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

// ToBody
func ResponseNetworksGetNetworkTrafficAnalysisItemToBody(state NetworksTrafficAnalysis, response *merakigosdk.ResponseNetworksGetNetworkTrafficAnalysis) NetworksTrafficAnalysis {
	itemState := ResponseNetworksGetNetworkTrafficAnalysis{
		CustomPieChartItems: func() *[]ResponseNetworksGetNetworkTrafficAnalysisCustomPieChartItems {
			if response.CustomPieChartItems != nil {
				result := make([]ResponseNetworksGetNetworkTrafficAnalysisCustomPieChartItems, len(*response.CustomPieChartItems))
				for i, customPieChartItems := range *response.CustomPieChartItems {
					result[i] = ResponseNetworksGetNetworkTrafficAnalysisCustomPieChartItems{
						Name:  types.StringValue(customPieChartItems.Name),
						Type:  types.StringValue(customPieChartItems.Type),
						Value: types.StringValue(customPieChartItems.Value),
					}
				}
				return &result
			}
			return &[]ResponseNetworksGetNetworkTrafficAnalysisCustomPieChartItems{}
		}(),
		Mode: types.StringValue(response.Mode),
	}
	state.Item = &itemState
	return state
}
