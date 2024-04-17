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
	_ datasource.DataSource              = &DevicesWirelessRadioSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesWirelessRadioSettingsDataSource{}
)

func NewDevicesWirelessRadioSettingsDataSource() datasource.DataSource {
	return &DevicesWirelessRadioSettingsDataSource{}
}

type DevicesWirelessRadioSettingsDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesWirelessRadioSettingsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesWirelessRadioSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_wireless_radio_settings"
}

func (d *DevicesWirelessRadioSettingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"five_ghz_settings": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"channel": schema.Int64Attribute{
								Computed: true,
							},
							"channel_width": schema.Int64Attribute{
								Computed: true,
							},
							"target_power": schema.Int64Attribute{
								Computed: true,
							},
						},
					},
					"rf_profile_id": schema.StringAttribute{
						Computed: true,
					},
					"serial": schema.StringAttribute{
						Computed: true,
					},
					"two_four_ghz_settings": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"channel": schema.Int64Attribute{
								Computed: true,
							},
							"target_power": schema.Int64Attribute{
								Computed: true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *DevicesWirelessRadioSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesWirelessRadioSettings DevicesWirelessRadioSettings
	diags := req.Config.Get(ctx, &devicesWirelessRadioSettings)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceWirelessRadioSettings")
		vvSerial := devicesWirelessRadioSettings.Serial.ValueString()

		response1, restyResp1, err := d.client.Wireless.GetDeviceWirelessRadioSettings(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceWirelessRadioSettings",
				err.Error(),
			)
			return
		}

		devicesWirelessRadioSettings = ResponseWirelessGetDeviceWirelessRadioSettingsItemToBody(devicesWirelessRadioSettings, response1)
		diags = resp.State.Set(ctx, &devicesWirelessRadioSettings)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesWirelessRadioSettings struct {
	Serial types.String                                    `tfsdk:"serial"`
	Item   *ResponseWirelessGetDeviceWirelessRadioSettings `tfsdk:"item"`
}

type ResponseWirelessGetDeviceWirelessRadioSettings struct {
	FiveGhzSettings    *ResponseWirelessGetDeviceWirelessRadioSettingsFiveGhzSettings    `tfsdk:"five_ghz_settings"`
	RfProfileID        types.String                                                      `tfsdk:"rf_profile_id"`
	Serial             types.String                                                      `tfsdk:"serial"`
	TwoFourGhzSettings *ResponseWirelessGetDeviceWirelessRadioSettingsTwoFourGhzSettings `tfsdk:"two_four_ghz_settings"`
}

type ResponseWirelessGetDeviceWirelessRadioSettingsFiveGhzSettings struct {
	Channel      types.Int64 `tfsdk:"channel"`
	ChannelWidth types.Int64 `tfsdk:"channel_width"`
	TargetPower  types.Int64 `tfsdk:"target_power"`
}

type ResponseWirelessGetDeviceWirelessRadioSettingsTwoFourGhzSettings struct {
	Channel     types.Int64 `tfsdk:"channel"`
	TargetPower types.Int64 `tfsdk:"target_power"`
}

// ToBody
func ResponseWirelessGetDeviceWirelessRadioSettingsItemToBody(state DevicesWirelessRadioSettings, response *merakigosdk.ResponseWirelessGetDeviceWirelessRadioSettings) DevicesWirelessRadioSettings {
	itemState := ResponseWirelessGetDeviceWirelessRadioSettings{
		FiveGhzSettings: func() *ResponseWirelessGetDeviceWirelessRadioSettingsFiveGhzSettings {
			if response.FiveGhzSettings != nil {
				return &ResponseWirelessGetDeviceWirelessRadioSettingsFiveGhzSettings{
					Channel: func() types.Int64 {
						if response.FiveGhzSettings.Channel != nil {
							return types.Int64Value(int64(*response.FiveGhzSettings.Channel))
						}
						return types.Int64{}
					}(),
					ChannelWidth: func() types.Int64 {
						if response.FiveGhzSettings.ChannelWidth != nil {
							return types.Int64Value(int64(*response.FiveGhzSettings.ChannelWidth))
						}
						return types.Int64{}
					}(),
					TargetPower: func() types.Int64 {
						if response.FiveGhzSettings.TargetPower != nil {
							return types.Int64Value(int64(*response.FiveGhzSettings.TargetPower))
						}
						return types.Int64{}
					}(),
				}
			}
			return &ResponseWirelessGetDeviceWirelessRadioSettingsFiveGhzSettings{}
		}(),
		RfProfileID: types.StringValue(response.RfProfileID),
		Serial:      types.StringValue(response.Serial),
		TwoFourGhzSettings: func() *ResponseWirelessGetDeviceWirelessRadioSettingsTwoFourGhzSettings {
			if response.TwoFourGhzSettings != nil {
				return &ResponseWirelessGetDeviceWirelessRadioSettingsTwoFourGhzSettings{
					Channel: func() types.Int64 {
						if response.TwoFourGhzSettings.Channel != nil {
							return types.Int64Value(int64(*response.TwoFourGhzSettings.Channel))
						}
						return types.Int64{}
					}(),
					TargetPower: func() types.Int64 {
						if response.TwoFourGhzSettings.TargetPower != nil {
							return types.Int64Value(int64(*response.TwoFourGhzSettings.TargetPower))
						}
						return types.Int64{}
					}(),
				}
			}
			return &ResponseWirelessGetDeviceWirelessRadioSettingsTwoFourGhzSettings{}
		}(),
	}
	state.Item = &itemState
	return state
}
