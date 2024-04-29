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
	_ datasource.DataSource              = &NetworksSmDevicesConnectivityDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSmDevicesConnectivityDataSource{}
)

func NewNetworksSmDevicesConnectivityDataSource() datasource.DataSource {
	return &NetworksSmDevicesConnectivityDataSource{}
}

type NetworksSmDevicesConnectivityDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSmDevicesConnectivityDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSmDevicesConnectivityDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_devices_connectivity"
}

func (d *NetworksSmDevicesConnectivityDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"device_id": schema.StringAttribute{
				MarkdownDescription: `deviceId path parameter. Device ID`,
				Required:            true,
			},
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSmGetNetworkSmDeviceConnectivity`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"first_seen_at": schema.StringAttribute{
							MarkdownDescription: `When the device was first seen as connected to the internet in each connection.`,
							Computed:            true,
						},
						"last_seen_at": schema.StringAttribute{
							MarkdownDescription: `When the device was last seen as connected to the internet in each connection.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSmDevicesConnectivityDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSmDevicesConnectivity NetworksSmDevicesConnectivity
	diags := req.Config.Get(ctx, &networksSmDevicesConnectivity)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSmDeviceConnectivity")
		vvNetworkID := networksSmDevicesConnectivity.NetworkID.ValueString()
		vvDeviceID := networksSmDevicesConnectivity.DeviceID.ValueString()
		queryParams1 := merakigosdk.GetNetworkSmDeviceConnectivityQueryParams{}

		queryParams1.PerPage = int(networksSmDevicesConnectivity.PerPage.ValueInt64())
		queryParams1.StartingAfter = networksSmDevicesConnectivity.StartingAfter.ValueString()
		queryParams1.EndingBefore = networksSmDevicesConnectivity.EndingBefore.ValueString()

		response1, restyResp1, err := d.client.Sm.GetNetworkSmDeviceConnectivity(vvNetworkID, vvDeviceID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmDeviceConnectivity",
				err.Error(),
			)
			return
		}

		networksSmDevicesConnectivity = ResponseSmGetNetworkSmDeviceConnectivityItemsToBody(networksSmDevicesConnectivity, response1)
		diags = resp.State.Set(ctx, &networksSmDevicesConnectivity)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSmDevicesConnectivity struct {
	NetworkID     types.String                                    `tfsdk:"network_id"`
	DeviceID      types.String                                    `tfsdk:"device_id"`
	PerPage       types.Int64                                     `tfsdk:"per_page"`
	StartingAfter types.String                                    `tfsdk:"starting_after"`
	EndingBefore  types.String                                    `tfsdk:"ending_before"`
	Items         *[]ResponseItemSmGetNetworkSmDeviceConnectivity `tfsdk:"items"`
}

type ResponseItemSmGetNetworkSmDeviceConnectivity struct {
	FirstSeenAt types.String `tfsdk:"first_seen_at"`
	LastSeenAt  types.String `tfsdk:"last_seen_at"`
}

// ToBody
func ResponseSmGetNetworkSmDeviceConnectivityItemsToBody(state NetworksSmDevicesConnectivity, response *merakigosdk.ResponseSmGetNetworkSmDeviceConnectivity) NetworksSmDevicesConnectivity {
	var items []ResponseItemSmGetNetworkSmDeviceConnectivity
	for _, item := range *response {
		itemState := ResponseItemSmGetNetworkSmDeviceConnectivity{
			FirstSeenAt: types.StringValue(item.FirstSeenAt),
			LastSeenAt:  types.StringValue(item.LastSeenAt),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
