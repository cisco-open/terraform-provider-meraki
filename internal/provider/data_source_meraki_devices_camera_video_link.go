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
	_ datasource.DataSource              = &DevicesCameraVideoLinkDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesCameraVideoLinkDataSource{}
)

func NewDevicesCameraVideoLinkDataSource() datasource.DataSource {
	return &DevicesCameraVideoLinkDataSource{}
}

type DevicesCameraVideoLinkDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesCameraVideoLinkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesCameraVideoLinkDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_camera_video_link"
}

func (d *DevicesCameraVideoLinkDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"timestamp": schema.StringAttribute{
				MarkdownDescription: `timestamp query parameter. [optional] The video link will start at this time. The timestamp should be a string in ISO8601 format. If no timestamp is specified, we will assume current time.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"url": schema.StringAttribute{
						Computed: true,
					},
					"vision_url": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *DevicesCameraVideoLinkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesCameraVideoLink DevicesCameraVideoLink
	diags := req.Config.Get(ctx, &devicesCameraVideoLink)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceCameraVideoLink")
		vvSerial := devicesCameraVideoLink.Serial.ValueString()
		queryParams1 := merakigosdk.GetDeviceCameraVideoLinkQueryParams{}

		queryParams1.Timestamp = devicesCameraVideoLink.Timestamp.ValueString()

		response1, restyResp1, err := d.client.Camera.GetDeviceCameraVideoLink(vvSerial, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCameraVideoLink",
				err.Error(),
			)
			return
		}

		devicesCameraVideoLink = ResponseCameraGetDeviceCameraVideoLinkItemToBody(devicesCameraVideoLink, response1)
		diags = resp.State.Set(ctx, &devicesCameraVideoLink)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesCameraVideoLink struct {
	Serial    types.String                            `tfsdk:"serial"`
	Timestamp types.String                            `tfsdk:"timestamp"`
	Item      *ResponseCameraGetDeviceCameraVideoLink `tfsdk:"item"`
}

type ResponseCameraGetDeviceCameraVideoLink struct {
	URL       types.String `tfsdk:"url"`
	VisionURL types.String `tfsdk:"vision_url"`
}

// ToBody
func ResponseCameraGetDeviceCameraVideoLinkItemToBody(state DevicesCameraVideoLink, response *merakigosdk.ResponseCameraGetDeviceCameraVideoLink) DevicesCameraVideoLink {
	itemState := ResponseCameraGetDeviceCameraVideoLink{
		URL:       types.StringValue(response.URL),
		VisionURL: types.StringValue(response.VisionURL),
	}
	state.Item = &itemState
	return state
}
