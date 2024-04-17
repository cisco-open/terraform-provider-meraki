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
	_ datasource.DataSource              = &DevicesCameraWirelessProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesCameraWirelessProfilesDataSource{}
)

func NewDevicesCameraWirelessProfilesDataSource() datasource.DataSource {
	return &DevicesCameraWirelessProfilesDataSource{}
}

type DevicesCameraWirelessProfilesDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesCameraWirelessProfilesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesCameraWirelessProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_camera_wireless_profiles"
}

func (d *DevicesCameraWirelessProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"ids": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"backup": schema.StringAttribute{
								Computed: true,
							},
							"primary": schema.StringAttribute{
								Computed: true,
							},
							"secondary": schema.StringAttribute{
								Computed: true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *DevicesCameraWirelessProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesCameraWirelessProfiles DevicesCameraWirelessProfiles
	diags := req.Config.Get(ctx, &devicesCameraWirelessProfiles)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceCameraWirelessProfiles")
		vvSerial := devicesCameraWirelessProfiles.Serial.ValueString()

		response1, restyResp1, err := d.client.Camera.GetDeviceCameraWirelessProfiles(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCameraWirelessProfiles",
				err.Error(),
			)
			return
		}

		devicesCameraWirelessProfiles = ResponseCameraGetDeviceCameraWirelessProfilesItemToBody(devicesCameraWirelessProfiles, response1)
		diags = resp.State.Set(ctx, &devicesCameraWirelessProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesCameraWirelessProfiles struct {
	Serial types.String                                   `tfsdk:"serial"`
	Item   *ResponseCameraGetDeviceCameraWirelessProfiles `tfsdk:"item"`
}

type ResponseCameraGetDeviceCameraWirelessProfiles struct {
	IDs *ResponseCameraGetDeviceCameraWirelessProfilesIds `tfsdk:"ids"`
}

type ResponseCameraGetDeviceCameraWirelessProfilesIds struct {
	Backup    types.String `tfsdk:"backup"`
	Primary   types.String `tfsdk:"primary"`
	Secondary types.String `tfsdk:"secondary"`
}

// ToBody
func ResponseCameraGetDeviceCameraWirelessProfilesItemToBody(state DevicesCameraWirelessProfiles, response *merakigosdk.ResponseCameraGetDeviceCameraWirelessProfiles) DevicesCameraWirelessProfiles {
	itemState := ResponseCameraGetDeviceCameraWirelessProfiles{
		IDs: func() *ResponseCameraGetDeviceCameraWirelessProfilesIds {
			if response.IDs != nil {
				return &ResponseCameraGetDeviceCameraWirelessProfilesIds{
					Backup:    types.StringValue(response.IDs.Backup),
					Primary:   types.StringValue(response.IDs.Primary),
					Secondary: types.StringValue(response.IDs.Secondary),
				}
			}
			return &ResponseCameraGetDeviceCameraWirelessProfilesIds{}
		}(),
	}
	state.Item = &itemState
	return state
}
