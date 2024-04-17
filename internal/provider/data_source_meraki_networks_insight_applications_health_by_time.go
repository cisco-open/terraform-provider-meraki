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
	_ datasource.DataSource              = &NetworksInsightApplicationsHealthByTimeDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksInsightApplicationsHealthByTimeDataSource{}
)

func NewNetworksInsightApplicationsHealthByTimeDataSource() datasource.DataSource {
	return &NetworksInsightApplicationsHealthByTimeDataSource{}
}

type NetworksInsightApplicationsHealthByTimeDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksInsightApplicationsHealthByTimeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksInsightApplicationsHealthByTimeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_insight_applications_health_by_time"
}

func (d *NetworksInsightApplicationsHealthByTimeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"application_id": schema.StringAttribute{
				MarkdownDescription: `applicationId path parameter. Application ID`,
				Required:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"resolution": schema.Int64Attribute{
				MarkdownDescription: `resolution query parameter. The time resolution in seconds for returned data. The valid resolutions are: 60, 300, 3600, 86400. The default is 300.`,
				Optional:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 7 days from today.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 7 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 2 hours.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseInsightGetNetworkInsightApplicationHealthByTime`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"end_ts": schema.StringAttribute{
							MarkdownDescription: `The end time of the query range`,
							Computed:            true,
						},
						"lan_goodput": schema.Int64Attribute{
							MarkdownDescription: `LAN goodput (Number of useful information bits delivered over a LAN per unit of time)`,
							Computed:            true,
						},
						"lan_latency_ms": schema.Float64Attribute{
							MarkdownDescription: `LAN latency in milliseconds`,
							Computed:            true,
						},
						"lan_loss_percent": schema.Float64Attribute{
							MarkdownDescription: `LAN loss percentage`,
							Computed:            true,
						},
						"num_clients": schema.Int64Attribute{
							MarkdownDescription: `Number of clients`,
							Computed:            true,
						},
						"recv": schema.Int64Attribute{
							MarkdownDescription: `Received kilobytes-per-second`,
							Computed:            true,
						},
						"response_duration": schema.Int64Attribute{
							MarkdownDescription: `Duration of the response, in milliseconds`,
							Computed:            true,
						},
						"sent": schema.Int64Attribute{
							MarkdownDescription: `Sent kilobytes-per-second`,
							Computed:            true,
						},
						"start_ts": schema.StringAttribute{
							MarkdownDescription: `The start time of the query range`,
							Computed:            true,
						},
						"wan_goodput": schema.Int64Attribute{
							MarkdownDescription: `WAN goodput (Number of useful information bits delivered over a WAN per unit of time)`,
							Computed:            true,
						},
						"wan_latency_ms": schema.Float64Attribute{
							MarkdownDescription: `WAN latency in milliseconds`,
							Computed:            true,
						},
						"wan_loss_percent": schema.Float64Attribute{
							MarkdownDescription: `WAN loss percentage`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksInsightApplicationsHealthByTimeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksInsightApplicationsHealthByTime NetworksInsightApplicationsHealthByTime
	diags := req.Config.Get(ctx, &networksInsightApplicationsHealthByTime)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkInsightApplicationHealthByTime")
		vvNetworkID := networksInsightApplicationsHealthByTime.NetworkID.ValueString()
		vvApplicationID := networksInsightApplicationsHealthByTime.ApplicationID.ValueString()
		queryParams1 := merakigosdk.GetNetworkInsightApplicationHealthByTimeQueryParams{}

		queryParams1.T0 = networksInsightApplicationsHealthByTime.T0.ValueString()
		queryParams1.T1 = networksInsightApplicationsHealthByTime.T1.ValueString()
		queryParams1.Timespan = networksInsightApplicationsHealthByTime.Timespan.ValueFloat64()
		queryParams1.Resolution = int(networksInsightApplicationsHealthByTime.Resolution.ValueInt64())

		response1, restyResp1, err := d.client.Insight.GetNetworkInsightApplicationHealthByTime(vvNetworkID, vvApplicationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkInsightApplicationHealthByTime",
				err.Error(),
			)
			return
		}

		networksInsightApplicationsHealthByTime = ResponseInsightGetNetworkInsightApplicationHealthByTimeItemsToBody(networksInsightApplicationsHealthByTime, response1)
		diags = resp.State.Set(ctx, &networksInsightApplicationsHealthByTime)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksInsightApplicationsHealthByTime struct {
	NetworkID     types.String                                                   `tfsdk:"network_id"`
	ApplicationID types.String                                                   `tfsdk:"application_id"`
	T0            types.String                                                   `tfsdk:"t0"`
	T1            types.String                                                   `tfsdk:"t1"`
	Timespan      types.Float64                                                  `tfsdk:"timespan"`
	Resolution    types.Int64                                                    `tfsdk:"resolution"`
	Items         *[]ResponseItemInsightGetNetworkInsightApplicationHealthByTime `tfsdk:"items"`
}

type ResponseItemInsightGetNetworkInsightApplicationHealthByTime struct {
	EndTs            types.String  `tfsdk:"end_ts"`
	LanGoodput       types.Int64   `tfsdk:"lan_goodput"`
	LanLatencyMs     types.Float64 `tfsdk:"lan_latency_ms"`
	LanLossPercent   types.Float64 `tfsdk:"lan_loss_percent"`
	NumClients       types.Int64   `tfsdk:"num_clients"`
	Recv             types.Int64   `tfsdk:"recv"`
	ResponseDuration types.Int64   `tfsdk:"response_duration"`
	Sent             types.Int64   `tfsdk:"sent"`
	StartTs          types.String  `tfsdk:"start_ts"`
	WanGoodput       types.Int64   `tfsdk:"wan_goodput"`
	WanLatencyMs     types.Float64 `tfsdk:"wan_latency_ms"`
	WanLossPercent   types.Float64 `tfsdk:"wan_loss_percent"`
}

// ToBody
func ResponseInsightGetNetworkInsightApplicationHealthByTimeItemsToBody(state NetworksInsightApplicationsHealthByTime, response *merakigosdk.ResponseInsightGetNetworkInsightApplicationHealthByTime) NetworksInsightApplicationsHealthByTime {
	var items []ResponseItemInsightGetNetworkInsightApplicationHealthByTime
	for _, item := range *response {
		itemState := ResponseItemInsightGetNetworkInsightApplicationHealthByTime{
			EndTs: types.StringValue(item.EndTs),
			LanGoodput: func() types.Int64 {
				if item.LanGoodput != nil {
					return types.Int64Value(int64(*item.LanGoodput))
				}
				return types.Int64{}
			}(),
			LanLatencyMs: func() types.Float64 {
				if item.LanLatencyMs != nil {
					return types.Float64Value(float64(*item.LanLatencyMs))
				}
				return types.Float64{}
			}(),
			LanLossPercent: func() types.Float64 {
				if item.LanLossPercent != nil {
					return types.Float64Value(float64(*item.LanLossPercent))
				}
				return types.Float64{}
			}(),
			NumClients: func() types.Int64 {
				if item.NumClients != nil {
					return types.Int64Value(int64(*item.NumClients))
				}
				return types.Int64{}
			}(),
			Recv: func() types.Int64 {
				if item.Recv != nil {
					return types.Int64Value(int64(*item.Recv))
				}
				return types.Int64{}
			}(),
			ResponseDuration: func() types.Int64 {
				if item.ResponseDuration != nil {
					return types.Int64Value(int64(*item.ResponseDuration))
				}
				return types.Int64{}
			}(),
			Sent: func() types.Int64 {
				if item.Sent != nil {
					return types.Int64Value(int64(*item.Sent))
				}
				return types.Int64{}
			}(),
			StartTs: types.StringValue(item.StartTs),
			WanGoodput: func() types.Int64 {
				if item.WanGoodput != nil {
					return types.Int64Value(int64(*item.WanGoodput))
				}
				return types.Int64{}
			}(),
			WanLatencyMs: func() types.Float64 {
				if item.WanLatencyMs != nil {
					return types.Float64Value(float64(*item.WanLatencyMs))
				}
				return types.Float64{}
			}(),
			WanLossPercent: func() types.Float64 {
				if item.WanLossPercent != nil {
					return types.Float64Value(float64(*item.WanLossPercent))
				}
				return types.Float64{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
