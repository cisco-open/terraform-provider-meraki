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
	_ datasource.DataSource              = &NetworksSwitchDhcpServerPolicyArpInspectionWarningsByDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchDhcpServerPolicyArpInspectionWarningsByDeviceDataSource{}
)

func NewNetworksSwitchDhcpServerPolicyArpInspectionWarningsByDeviceDataSource() datasource.DataSource {
	return &NetworksSwitchDhcpServerPolicyArpInspectionWarningsByDeviceDataSource{}
}

type NetworksSwitchDhcpServerPolicyArpInspectionWarningsByDeviceDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchDhcpServerPolicyArpInspectionWarningsByDeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchDhcpServerPolicyArpInspectionWarningsByDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_dhcp_server_policy_arp_inspection_warnings_by_device"
}

func (d *NetworksSwitchDhcpServerPolicyArpInspectionWarningsByDeviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"has_trusted_port": schema.BoolAttribute{
							MarkdownDescription: `Whether this switch has a trusted DAI port. Always false if supportsInspection is false.`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Switch name.`,
							Computed:            true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `Switch serial.`,
							Computed:            true,
						},
						"supports_inspection": schema.BoolAttribute{
							MarkdownDescription: `Whether this switch supports Dynamic ARP Inspection.`,
							Computed:            true,
						},
						"url": schema.StringAttribute{
							MarkdownDescription: `Url link to switch.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchDhcpServerPolicyArpInspectionWarningsByDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchDhcpServerPolicyArpInspectionWarningsByDevice NetworksSwitchDhcpServerPolicyArpInspectionWarningsByDevice
	diags := req.Config.Get(ctx, &networksSwitchDhcpServerPolicyArpInspectionWarningsByDevice)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice")
		vvNetworkID := networksSwitchDhcpServerPolicyArpInspectionWarningsByDevice.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceQueryParams{}

		queryParams1.PerPage = int(networksSwitchDhcpServerPolicyArpInspectionWarningsByDevice.PerPage.ValueInt64())
		queryParams1.StartingAfter = networksSwitchDhcpServerPolicyArpInspectionWarningsByDevice.StartingAfter.ValueString()
		queryParams1.EndingBefore = networksSwitchDhcpServerPolicyArpInspectionWarningsByDevice.EndingBefore.ValueString()

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice",
				err.Error(),
			)
			return
		}

		networksSwitchDhcpServerPolicyArpInspectionWarningsByDevice = ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceItemsToBody(networksSwitchDhcpServerPolicyArpInspectionWarningsByDevice, response1)
		diags = resp.State.Set(ctx, &networksSwitchDhcpServerPolicyArpInspectionWarningsByDevice)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchDhcpServerPolicyArpInspectionWarningsByDevice struct {
	NetworkID     types.String                                                                       `tfsdk:"network_id"`
	PerPage       types.Int64                                                                        `tfsdk:"per_page"`
	StartingAfter types.String                                                                       `tfsdk:"starting_after"`
	EndingBefore  types.String                                                                       `tfsdk:"ending_before"`
	Items         *[]ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice `tfsdk:"items"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice struct {
	HasTrustedPort     types.Bool   `tfsdk:"has_trusted_port"`
	Name               types.String `tfsdk:"name"`
	Serial             types.String `tfsdk:"serial"`
	SupportsInspection types.Bool   `tfsdk:"supports_inspection"`
	URL                types.String `tfsdk:"url"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceItemsToBody(state NetworksSwitchDhcpServerPolicyArpInspectionWarningsByDevice, response *merakigosdk.ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice) NetworksSwitchDhcpServerPolicyArpInspectionWarningsByDevice {
	var items []ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice
	for _, item := range *response {
		itemState := ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice{
			HasTrustedPort: func() types.Bool {
				if item.HasTrustedPort != nil {
					return types.BoolValue(*item.HasTrustedPort)
				}
				return types.Bool{}
			}(),
			Name:   types.StringValue(item.Name),
			Serial: types.StringValue(item.Serial),
			SupportsInspection: func() types.Bool {
				if item.SupportsInspection != nil {
					return types.BoolValue(*item.SupportsInspection)
				}
				return types.Bool{}
			}(),
			URL: types.StringValue(item.URL),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
