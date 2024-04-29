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
	_ datasource.DataSource              = &NetworksWirelessMeshStatusesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessMeshStatusesDataSource{}
)

func NewNetworksWirelessMeshStatusesDataSource() datasource.DataSource {
	return &NetworksWirelessMeshStatusesDataSource{}
}

type NetworksWirelessMeshStatusesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessMeshStatusesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessMeshStatusesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_mesh_statuses"
}

func (d *NetworksWirelessMeshStatusesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 500. Default is 50.`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseWirelessGetNetworkWirelessMeshStatuses`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"latest_mesh_performance": schema.SingleNestedAttribute{
							MarkdownDescription: `Current metrics on how the mesh is performing.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"mbps": schema.Int64Attribute{
									MarkdownDescription: `Average Mbps.`,
									Computed:            true,
								},
								"metric": schema.Int64Attribute{
									MarkdownDescription: `Represents the quality of the entire route from the repeater access point to its gateway access point.`,
									Computed:            true,
								},
								"usage_percentage": schema.StringAttribute{
									MarkdownDescription: `Mesh utilization as a percentage.`,
									Computed:            true,
								},
							},
						},
						"mesh_route": schema.ListAttribute{
							MarkdownDescription: `List of device serials that make up the mesh.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `The serial number for the device.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessMeshStatusesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessMeshStatuses NetworksWirelessMeshStatuses
	diags := req.Config.Get(ctx, &networksWirelessMeshStatuses)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessMeshStatuses")
		vvNetworkID := networksWirelessMeshStatuses.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkWirelessMeshStatusesQueryParams{}

		queryParams1.PerPage = int(networksWirelessMeshStatuses.PerPage.ValueInt64())
		queryParams1.StartingAfter = networksWirelessMeshStatuses.StartingAfter.ValueString()
		queryParams1.EndingBefore = networksWirelessMeshStatuses.EndingBefore.ValueString()

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessMeshStatuses(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessMeshStatuses",
				err.Error(),
			)
			return
		}

		networksWirelessMeshStatuses = ResponseWirelessGetNetworkWirelessMeshStatusesItemsToBody(networksWirelessMeshStatuses, response1)
		diags = resp.State.Set(ctx, &networksWirelessMeshStatuses)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessMeshStatuses struct {
	NetworkID     types.String                                          `tfsdk:"network_id"`
	PerPage       types.Int64                                           `tfsdk:"per_page"`
	StartingAfter types.String                                          `tfsdk:"starting_after"`
	EndingBefore  types.String                                          `tfsdk:"ending_before"`
	Items         *[]ResponseItemWirelessGetNetworkWirelessMeshStatuses `tfsdk:"items"`
}

type ResponseItemWirelessGetNetworkWirelessMeshStatuses struct {
	LatestMeshPerformance *ResponseItemWirelessGetNetworkWirelessMeshStatusesLatestMeshPerformance `tfsdk:"latest_mesh_performance"`
	MeshRoute             types.List                                                               `tfsdk:"mesh_route"`
	Serial                types.String                                                             `tfsdk:"serial"`
}

type ResponseItemWirelessGetNetworkWirelessMeshStatusesLatestMeshPerformance struct {
	Mbps            types.Int64  `tfsdk:"mbps"`
	Metric          types.Int64  `tfsdk:"metric"`
	UsagePercentage types.String `tfsdk:"usage_percentage"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessMeshStatusesItemsToBody(state NetworksWirelessMeshStatuses, response *merakigosdk.ResponseWirelessGetNetworkWirelessMeshStatuses) NetworksWirelessMeshStatuses {
	var items []ResponseItemWirelessGetNetworkWirelessMeshStatuses
	for _, item := range *response {
		itemState := ResponseItemWirelessGetNetworkWirelessMeshStatuses{
			LatestMeshPerformance: func() *ResponseItemWirelessGetNetworkWirelessMeshStatusesLatestMeshPerformance {
				if item.LatestMeshPerformance != nil {
					return &ResponseItemWirelessGetNetworkWirelessMeshStatusesLatestMeshPerformance{
						Mbps: func() types.Int64 {
							if item.LatestMeshPerformance.Mbps != nil {
								return types.Int64Value(int64(*item.LatestMeshPerformance.Mbps))
							}
							return types.Int64{}
						}(),
						Metric: func() types.Int64 {
							if item.LatestMeshPerformance.Metric != nil {
								return types.Int64Value(int64(*item.LatestMeshPerformance.Metric))
							}
							return types.Int64{}
						}(),
						UsagePercentage: types.StringValue(item.LatestMeshPerformance.UsagePercentage),
					}
				}
				return &ResponseItemWirelessGetNetworkWirelessMeshStatusesLatestMeshPerformance{}
			}(),
			MeshRoute: StringSliceToList(item.MeshRoute),
			Serial:    types.StringValue(item.Serial),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
