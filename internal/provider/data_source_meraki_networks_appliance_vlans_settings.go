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
	_ datasource.DataSource              = &NetworksApplianceVLANsSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceVLANsSettingsDataSource{}
)

func NewNetworksApplianceVLANsSettingsDataSource() datasource.DataSource {
	return &NetworksApplianceVLANsSettingsDataSource{}
}

type NetworksApplianceVLANsSettingsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceVLANsSettingsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceVLANsSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_vlans_settings"
}

func (d *NetworksApplianceVLANsSettingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"vlans_enabled": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating whether VLANs are enabled (true) or disabled (false) for the network`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *NetworksApplianceVLANsSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceVLANsSettings NetworksApplianceVLANsSettings
	diags := req.Config.Get(ctx, &networksApplianceVLANsSettings)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceVLANsSettings")
		vvNetworkID := networksApplianceVLANsSettings.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceVLANsSettings(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceVLANsSettings",
				err.Error(),
			)
			return
		}

		networksApplianceVLANsSettings = ResponseApplianceGetNetworkApplianceVLANsSettingsItemToBody(networksApplianceVLANsSettings, response1)
		diags = resp.State.Set(ctx, &networksApplianceVLANsSettings)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceVLANsSettings struct {
	NetworkID types.String                                       `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceVlansSettings `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceVlansSettings struct {
	VLANsEnabled types.Bool `tfsdk:"vlans_enabled"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceVLANsSettingsItemToBody(state NetworksApplianceVLANsSettings, response *merakigosdk.ResponseApplianceGetNetworkApplianceVLANsSettings) NetworksApplianceVLANsSettings {
	itemState := ResponseApplianceGetNetworkApplianceVlansSettings{
		VLANsEnabled: func() types.Bool {
			if response.VLANsEnabled != nil {
				return types.BoolValue(*response.VLANsEnabled)
			}
			return types.Bool{}
		}(),
	}
	state.Item = &itemState
	return state
}
