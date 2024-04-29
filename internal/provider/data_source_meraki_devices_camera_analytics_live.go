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
	_ datasource.DataSource              = &DevicesCameraAnalyticsLiveDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesCameraAnalyticsLiveDataSource{}
)

func NewDevicesCameraAnalyticsLiveDataSource() datasource.DataSource {
	return &DevicesCameraAnalyticsLiveDataSource{}
}

type DevicesCameraAnalyticsLiveDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesCameraAnalyticsLiveDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesCameraAnalyticsLiveDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_camera_analytics_live"
}

func (d *DevicesCameraAnalyticsLiveDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"ts": schema.StringAttribute{
						Computed: true,
					},
					"zones": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"status_0": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"person": schema.Int64Attribute{
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *DevicesCameraAnalyticsLiveDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesCameraAnalyticsLive DevicesCameraAnalyticsLive
	diags := req.Config.Get(ctx, &devicesCameraAnalyticsLive)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceCameraAnalyticsLive")
		vvSerial := devicesCameraAnalyticsLive.Serial.ValueString()

		response1, restyResp1, err := d.client.Camera.GetDeviceCameraAnalyticsLive(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCameraAnalyticsLive",
				err.Error(),
			)
			return
		}

		devicesCameraAnalyticsLive = ResponseCameraGetDeviceCameraAnalyticsLiveItemToBody(devicesCameraAnalyticsLive, response1)
		diags = resp.State.Set(ctx, &devicesCameraAnalyticsLive)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesCameraAnalyticsLive struct {
	Serial types.String                                `tfsdk:"serial"`
	Item   *ResponseCameraGetDeviceCameraAnalyticsLive `tfsdk:"item"`
}

type ResponseCameraGetDeviceCameraAnalyticsLive struct {
	Ts    types.String                                     `tfsdk:"ts"`
	Zones *ResponseCameraGetDeviceCameraAnalyticsLiveZones `tfsdk:"zones"`
}

type ResponseCameraGetDeviceCameraAnalyticsLiveZones struct {
	Status0 *ResponseCameraGetDeviceCameraAnalyticsLiveZones0 `tfsdk:"status_0"`
}

type ResponseCameraGetDeviceCameraAnalyticsLiveZones0 struct {
	Person types.Int64 `tfsdk:"person"`
}

// ToBody
func ResponseCameraGetDeviceCameraAnalyticsLiveItemToBody(state DevicesCameraAnalyticsLive, response *merakigosdk.ResponseCameraGetDeviceCameraAnalyticsLive) DevicesCameraAnalyticsLive {
	itemState := ResponseCameraGetDeviceCameraAnalyticsLive{
		Ts: types.StringValue(response.Ts),
		Zones: func() *ResponseCameraGetDeviceCameraAnalyticsLiveZones {
			if response.Zones != nil {
				return &ResponseCameraGetDeviceCameraAnalyticsLiveZones{
					Status0: func() *ResponseCameraGetDeviceCameraAnalyticsLiveZones0 {
						if response.Zones.Status0 != nil {
							return &ResponseCameraGetDeviceCameraAnalyticsLiveZones0{
								Person: func() types.Int64 {
									if response.Zones.Status0.Person != nil {
										return types.Int64Value(int64(*response.Zones.Status0.Person))
									}
									return types.Int64{}
								}(),
							}
						}
						return &ResponseCameraGetDeviceCameraAnalyticsLiveZones0{}
					}(),
				}
			}
			return &ResponseCameraGetDeviceCameraAnalyticsLiveZones{}
		}(),
	}
	state.Item = &itemState
	return state
}
