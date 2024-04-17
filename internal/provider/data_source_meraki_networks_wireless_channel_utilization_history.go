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
	_ datasource.DataSource              = &NetworksWirelessChannelUtilizationHistoryDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessChannelUtilizationHistoryDataSource{}
)

func NewNetworksWirelessChannelUtilizationHistoryDataSource() datasource.DataSource {
	return &NetworksWirelessChannelUtilizationHistoryDataSource{}
}

type NetworksWirelessChannelUtilizationHistoryDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessChannelUtilizationHistoryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessChannelUtilizationHistoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_channel_utilization_history"
}

func (d *NetworksWirelessChannelUtilizationHistoryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ap_tag": schema.StringAttribute{
				MarkdownDescription: `apTag query parameter. Filter results by AP tag to return AP channel utilization metrics for devices labeled with the given tag; either :clientId or :deviceSerial must be jointly specified.`,
				Optional:            true,
			},
			"auto_resolution": schema.BoolAttribute{
				MarkdownDescription: `autoResolution query parameter. Automatically select a data resolution based on the given timespan; this overrides the value specified by the 'resolution' parameter. The default setting is false.`,
				Optional:            true,
			},
			"band": schema.StringAttribute{
				MarkdownDescription: `band query parameter. Filter results by band (either '2.4', '5' or '6').`,
				Optional:            true,
			},
			"client_id": schema.StringAttribute{
				MarkdownDescription: `clientId query parameter. Filter results by network client to return per-device, per-band AP channel utilization metrics inner joined by the queried client's connection history.`,
				Optional:            true,
			},
			"device_serial": schema.StringAttribute{
				MarkdownDescription: `deviceSerial query parameter. Filter results by device to return AP channel utilization metrics for the queried device; either :band or :clientId must be jointly specified.`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"resolution": schema.Int64Attribute{
				MarkdownDescription: `resolution query parameter. The time resolution in seconds for returned data. The valid resolutions are: 600, 1200, 3600, 14400, 86400. The default is 86400.`,
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
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseWirelessGetNetworkWirelessChannelUtilizationHistory`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"end_ts": schema.StringAttribute{
							MarkdownDescription: `The end time of the query range`,
							Computed:            true,
						},
						"start_ts": schema.StringAttribute{
							MarkdownDescription: `The start time of the query range`,
							Computed:            true,
						},
						"utilization80211": schema.Float64Attribute{
							MarkdownDescription: `Average wifi utilization`,
							Computed:            true,
						},
						"utilization_non80211": schema.Float64Attribute{
							MarkdownDescription: `Average signal interference`,
							Computed:            true,
						},
						"utilization_total": schema.Float64Attribute{
							MarkdownDescription: `Total channel utilization`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessChannelUtilizationHistoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessChannelUtilizationHistory NetworksWirelessChannelUtilizationHistory
	diags := req.Config.Get(ctx, &networksWirelessChannelUtilizationHistory)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessChannelUtilizationHistory")
		vvNetworkID := networksWirelessChannelUtilizationHistory.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkWirelessChannelUtilizationHistoryQueryParams{}

		queryParams1.T0 = networksWirelessChannelUtilizationHistory.T0.ValueString()
		queryParams1.T1 = networksWirelessChannelUtilizationHistory.T1.ValueString()
		queryParams1.Timespan = networksWirelessChannelUtilizationHistory.Timespan.ValueFloat64()
		queryParams1.Resolution = int(networksWirelessChannelUtilizationHistory.Resolution.ValueInt64())
		queryParams1.AutoResolution = networksWirelessChannelUtilizationHistory.AutoResolution.ValueBool()
		queryParams1.ClientID = networksWirelessChannelUtilizationHistory.ClientID.ValueString()
		queryParams1.DeviceSerial = networksWirelessChannelUtilizationHistory.DeviceSerial.ValueString()
		queryParams1.ApTag = networksWirelessChannelUtilizationHistory.ApTag.ValueString()
		queryParams1.Band = networksWirelessChannelUtilizationHistory.Band.ValueString()

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessChannelUtilizationHistory(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessChannelUtilizationHistory",
				err.Error(),
			)
			return
		}

		networksWirelessChannelUtilizationHistory = ResponseWirelessGetNetworkWirelessChannelUtilizationHistoryItemsToBody(networksWirelessChannelUtilizationHistory, response1)
		diags = resp.State.Set(ctx, &networksWirelessChannelUtilizationHistory)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessChannelUtilizationHistory struct {
	NetworkID      types.String                                                       `tfsdk:"network_id"`
	T0             types.String                                                       `tfsdk:"t0"`
	T1             types.String                                                       `tfsdk:"t1"`
	Timespan       types.Float64                                                      `tfsdk:"timespan"`
	Resolution     types.Int64                                                        `tfsdk:"resolution"`
	AutoResolution types.Bool                                                         `tfsdk:"auto_resolution"`
	ClientID       types.String                                                       `tfsdk:"client_id"`
	DeviceSerial   types.String                                                       `tfsdk:"device_serial"`
	ApTag          types.String                                                       `tfsdk:"ap_tag"`
	Band           types.String                                                       `tfsdk:"band"`
	Items          *[]ResponseItemWirelessGetNetworkWirelessChannelUtilizationHistory `tfsdk:"items"`
}

type ResponseItemWirelessGetNetworkWirelessChannelUtilizationHistory struct {
	EndTs               types.String  `tfsdk:"end_ts"`
	StartTs             types.String  `tfsdk:"start_ts"`
	Utilization80211    types.Float64 `tfsdk:"utilization80211"`
	UtilizationNon80211 types.Float64 `tfsdk:"utilization_non80211"`
	UtilizationTotal    types.Float64 `tfsdk:"utilization_total"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessChannelUtilizationHistoryItemsToBody(state NetworksWirelessChannelUtilizationHistory, response *merakigosdk.ResponseWirelessGetNetworkWirelessChannelUtilizationHistory) NetworksWirelessChannelUtilizationHistory {
	var items []ResponseItemWirelessGetNetworkWirelessChannelUtilizationHistory
	for _, item := range *response {
		itemState := ResponseItemWirelessGetNetworkWirelessChannelUtilizationHistory{
			EndTs:   types.StringValue(item.EndTs),
			StartTs: types.StringValue(item.StartTs),
			Utilization80211: func() types.Float64 {
				if item.Utilization80211 != nil {
					return types.Float64Value(float64(*item.Utilization80211))
				}
				return types.Float64{}
			}(),
			UtilizationNon80211: func() types.Float64 {
				if item.UtilizationNon80211 != nil {
					return types.Float64Value(float64(*item.UtilizationNon80211))
				}
				return types.Float64{}
			}(),
			UtilizationTotal: func() types.Float64 {
				if item.UtilizationTotal != nil {
					return types.Float64Value(float64(*item.UtilizationTotal))
				}
				return types.Float64{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
