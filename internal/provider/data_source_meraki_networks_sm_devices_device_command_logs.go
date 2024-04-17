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
	_ datasource.DataSource              = &NetworksSmDevicesDeviceCommandLogsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSmDevicesDeviceCommandLogsDataSource{}
)

func NewNetworksSmDevicesDeviceCommandLogsDataSource() datasource.DataSource {
	return &NetworksSmDevicesDeviceCommandLogsDataSource{}
}

type NetworksSmDevicesDeviceCommandLogsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSmDevicesDeviceCommandLogsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSmDevicesDeviceCommandLogsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_devices_device_command_logs"
}

func (d *NetworksSmDevicesDeviceCommandLogsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				MarkdownDescription: `Array of ResponseSmGetNetworkSmDeviceDeviceCommandLogs`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"action": schema.StringAttribute{
							MarkdownDescription: `The type of command sent to the device.`,
							Computed:            true,
						},
						"dashboard_user": schema.StringAttribute{
							MarkdownDescription: `The Meraki dashboard user who initiated the command.`,
							Computed:            true,
						},
						"details": schema.StringAttribute{
							MarkdownDescription: `A JSON string object containing command details.`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the device to which the command is sent.`,
							Computed:            true,
						},
						"ts": schema.StringAttribute{
							MarkdownDescription: `The time the command was sent to the device.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSmDevicesDeviceCommandLogsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSmDevicesDeviceCommandLogs NetworksSmDevicesDeviceCommandLogs
	diags := req.Config.Get(ctx, &networksSmDevicesDeviceCommandLogs)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSmDeviceDeviceCommandLogs")
		vvNetworkID := networksSmDevicesDeviceCommandLogs.NetworkID.ValueString()
		vvDeviceID := networksSmDevicesDeviceCommandLogs.DeviceID.ValueString()
		queryParams1 := merakigosdk.GetNetworkSmDeviceDeviceCommandLogsQueryParams{}

		queryParams1.PerPage = int(networksSmDevicesDeviceCommandLogs.PerPage.ValueInt64())
		queryParams1.StartingAfter = networksSmDevicesDeviceCommandLogs.StartingAfter.ValueString()
		queryParams1.EndingBefore = networksSmDevicesDeviceCommandLogs.EndingBefore.ValueString()

		response1, restyResp1, err := d.client.Sm.GetNetworkSmDeviceDeviceCommandLogs(vvNetworkID, vvDeviceID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmDeviceDeviceCommandLogs",
				err.Error(),
			)
			return
		}

		networksSmDevicesDeviceCommandLogs = ResponseSmGetNetworkSmDeviceDeviceCommandLogsItemsToBody(networksSmDevicesDeviceCommandLogs, response1)
		diags = resp.State.Set(ctx, &networksSmDevicesDeviceCommandLogs)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSmDevicesDeviceCommandLogs struct {
	NetworkID     types.String                                         `tfsdk:"network_id"`
	DeviceID      types.String                                         `tfsdk:"device_id"`
	PerPage       types.Int64                                          `tfsdk:"per_page"`
	StartingAfter types.String                                         `tfsdk:"starting_after"`
	EndingBefore  types.String                                         `tfsdk:"ending_before"`
	Items         *[]ResponseItemSmGetNetworkSmDeviceDeviceCommandLogs `tfsdk:"items"`
}

type ResponseItemSmGetNetworkSmDeviceDeviceCommandLogs struct {
	Action        types.String `tfsdk:"action"`
	DashboardUser types.String `tfsdk:"dashboard_user"`
	Details       types.String `tfsdk:"details"`
	Name          types.String `tfsdk:"name"`
	Ts            types.String `tfsdk:"ts"`
}

// ToBody
func ResponseSmGetNetworkSmDeviceDeviceCommandLogsItemsToBody(state NetworksSmDevicesDeviceCommandLogs, response *merakigosdk.ResponseSmGetNetworkSmDeviceDeviceCommandLogs) NetworksSmDevicesDeviceCommandLogs {
	var items []ResponseItemSmGetNetworkSmDeviceDeviceCommandLogs
	for _, item := range *response {
		itemState := ResponseItemSmGetNetworkSmDeviceDeviceCommandLogs{
			Action:        types.StringValue(item.Action),
			DashboardUser: types.StringValue(item.DashboardUser),
			Details:       types.StringValue(item.Details),
			Name:          types.StringValue(item.Name),
			Ts:            types.StringValue(item.Ts),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
