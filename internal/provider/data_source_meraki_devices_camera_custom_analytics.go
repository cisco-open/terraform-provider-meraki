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
	_ datasource.DataSource              = &DevicesCameraCustomAnalyticsDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesCameraCustomAnalyticsDataSource{}
)

func NewDevicesCameraCustomAnalyticsDataSource() datasource.DataSource {
	return &DevicesCameraCustomAnalyticsDataSource{}
}

type DevicesCameraCustomAnalyticsDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesCameraCustomAnalyticsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesCameraCustomAnalyticsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_camera_custom_analytics"
}

func (d *DevicesCameraCustomAnalyticsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"artifact_id": schema.StringAttribute{
						MarkdownDescription: `Custom analytics artifact ID`,
						Computed:            true,
					},
					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether custom analytics is enabled`,
						Computed:            true,
					},
					"parameters": schema.SetNestedAttribute{
						MarkdownDescription: `Parameters for the custom analytics workload`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"name": schema.StringAttribute{
									MarkdownDescription: `Name of the parameter`,
									Computed:            true,
								},
								"value": schema.Float64Attribute{
									MarkdownDescription: `Value of the parameter`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *DevicesCameraCustomAnalyticsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesCameraCustomAnalytics DevicesCameraCustomAnalytics
	diags := req.Config.Get(ctx, &devicesCameraCustomAnalytics)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceCameraCustomAnalytics")
		vvSerial := devicesCameraCustomAnalytics.Serial.ValueString()

		response1, restyResp1, err := d.client.Camera.GetDeviceCameraCustomAnalytics(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCameraCustomAnalytics",
				err.Error(),
			)
			return
		}

		devicesCameraCustomAnalytics = ResponseCameraGetDeviceCameraCustomAnalyticsItemToBody(devicesCameraCustomAnalytics, response1)
		diags = resp.State.Set(ctx, &devicesCameraCustomAnalytics)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesCameraCustomAnalytics struct {
	Serial types.String                                  `tfsdk:"serial"`
	Item   *ResponseCameraGetDeviceCameraCustomAnalytics `tfsdk:"item"`
}

type ResponseCameraGetDeviceCameraCustomAnalytics struct {
	ArtifactID types.String                                              `tfsdk:"artifact_id"`
	Enabled    types.Bool                                                `tfsdk:"enabled"`
	Parameters *[]ResponseCameraGetDeviceCameraCustomAnalyticsParameters `tfsdk:"parameters"`
}

type ResponseCameraGetDeviceCameraCustomAnalyticsParameters struct {
	Name  types.String  `tfsdk:"name"`
	Value types.Float64 `tfsdk:"value"`
}

// ToBody
func ResponseCameraGetDeviceCameraCustomAnalyticsItemToBody(state DevicesCameraCustomAnalytics, response *merakigosdk.ResponseCameraGetDeviceCameraCustomAnalytics) DevicesCameraCustomAnalytics {
	itemState := ResponseCameraGetDeviceCameraCustomAnalytics{
		ArtifactID: types.StringValue(response.ArtifactID),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		Parameters: func() *[]ResponseCameraGetDeviceCameraCustomAnalyticsParameters {
			if response.Parameters != nil {
				result := make([]ResponseCameraGetDeviceCameraCustomAnalyticsParameters, len(*response.Parameters))
				for i, parameters := range *response.Parameters {
					result[i] = ResponseCameraGetDeviceCameraCustomAnalyticsParameters{
						Name: types.StringValue(parameters.Name),
						Value: func() types.Float64 {
							if parameters.Value != nil {
								return types.Float64Value(float64(*parameters.Value))
							}
							return types.Float64{}
						}(),
					}
				}
				return &result
			}
			return &[]ResponseCameraGetDeviceCameraCustomAnalyticsParameters{}
		}(),
	}
	state.Item = &itemState
	return state
}
