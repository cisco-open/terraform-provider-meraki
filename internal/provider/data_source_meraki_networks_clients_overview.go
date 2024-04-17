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
	_ datasource.DataSource              = &NetworksClientsOverviewDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksClientsOverviewDataSource{}
)

func NewNetworksClientsOverviewDataSource() datasource.DataSource {
	return &NetworksClientsOverviewDataSource{}
}

type NetworksClientsOverviewDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksClientsOverviewDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksClientsOverviewDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_clients_overview"
}

func (d *NetworksClientsOverviewDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"resolution": schema.Int64Attribute{
				MarkdownDescription: `resolution query parameter. The time resolution in seconds for returned data. The valid resolutions are: 7200, 86400, 604800, 2592000. The default is 604800.`,
				Optional:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 31 days from today.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 31 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"counts": schema.SingleNestedAttribute{
						MarkdownDescription: `The number of clients on a network over a given time range`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"total": schema.Int64Attribute{
								MarkdownDescription: `The total number of clients on a network`,
								Computed:            true,
							},
							"with_heavy_usage": schema.Int64Attribute{
								MarkdownDescription: `The total number of clients with heavy usage on a network`,
								Computed:            true,
							},
						},
					},
					"usages": schema.SingleNestedAttribute{
						MarkdownDescription: `The average usage of the clients on a network over a given time range`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"average": schema.Int64Attribute{
								MarkdownDescription: `The average usage of all clients on a network`,
								Computed:            true,
							},
							"with_heavy_usage_average": schema.Int64Attribute{
								MarkdownDescription: `The average usage of all clients with heavy usage on a network`,
								Computed:            true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksClientsOverviewDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksClientsOverview NetworksClientsOverview
	diags := req.Config.Get(ctx, &networksClientsOverview)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkClientsOverview")
		vvNetworkID := networksClientsOverview.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkClientsOverviewQueryParams{}

		queryParams1.T0 = networksClientsOverview.T0.ValueString()
		queryParams1.T1 = networksClientsOverview.T1.ValueString()
		queryParams1.Timespan = networksClientsOverview.Timespan.ValueFloat64()
		queryParams1.Resolution = int(networksClientsOverview.Resolution.ValueInt64())

		response1, restyResp1, err := d.client.Networks.GetNetworkClientsOverview(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkClientsOverview",
				err.Error(),
			)
			return
		}

		networksClientsOverview = ResponseNetworksGetNetworkClientsOverviewItemToBody(networksClientsOverview, response1)
		diags = resp.State.Set(ctx, &networksClientsOverview)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksClientsOverview struct {
	NetworkID  types.String                               `tfsdk:"network_id"`
	T0         types.String                               `tfsdk:"t0"`
	T1         types.String                               `tfsdk:"t1"`
	Timespan   types.Float64                              `tfsdk:"timespan"`
	Resolution types.Int64                                `tfsdk:"resolution"`
	Item       *ResponseNetworksGetNetworkClientsOverview `tfsdk:"item"`
}

type ResponseNetworksGetNetworkClientsOverview struct {
	Counts *ResponseNetworksGetNetworkClientsOverviewCounts `tfsdk:"counts"`
	Usages *ResponseNetworksGetNetworkClientsOverviewUsages `tfsdk:"usages"`
}

type ResponseNetworksGetNetworkClientsOverviewCounts struct {
	Total          types.Int64 `tfsdk:"total"`
	WithHeavyUsage types.Int64 `tfsdk:"with_heavy_usage"`
}

type ResponseNetworksGetNetworkClientsOverviewUsages struct {
	Average               types.Int64 `tfsdk:"average"`
	WithHeavyUsageAverage types.Int64 `tfsdk:"with_heavy_usage_average"`
}

// ToBody
func ResponseNetworksGetNetworkClientsOverviewItemToBody(state NetworksClientsOverview, response *merakigosdk.ResponseNetworksGetNetworkClientsOverview) NetworksClientsOverview {
	itemState := ResponseNetworksGetNetworkClientsOverview{
		Counts: func() *ResponseNetworksGetNetworkClientsOverviewCounts {
			if response.Counts != nil {
				return &ResponseNetworksGetNetworkClientsOverviewCounts{
					Total: func() types.Int64 {
						if response.Counts.Total != nil {
							return types.Int64Value(int64(*response.Counts.Total))
						}
						return types.Int64{}
					}(),
					WithHeavyUsage: func() types.Int64 {
						if response.Counts.WithHeavyUsage != nil {
							return types.Int64Value(int64(*response.Counts.WithHeavyUsage))
						}
						return types.Int64{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkClientsOverviewCounts{}
		}(),
		Usages: func() *ResponseNetworksGetNetworkClientsOverviewUsages {
			if response.Usages != nil {
				return &ResponseNetworksGetNetworkClientsOverviewUsages{
					Average: func() types.Int64 {
						if response.Usages.Average != nil {
							return types.Int64Value(int64(*response.Usages.Average))
						}
						return types.Int64{}
					}(),
					WithHeavyUsageAverage: func() types.Int64 {
						if response.Usages.WithHeavyUsageAverage != nil {
							return types.Int64Value(int64(*response.Usages.WithHeavyUsageAverage))
						}
						return types.Int64{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkClientsOverviewUsages{}
		}(),
	}
	state.Item = &itemState
	return state
}
