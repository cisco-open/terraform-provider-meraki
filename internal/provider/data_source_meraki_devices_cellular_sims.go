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
	_ datasource.DataSource              = &DevicesCellularSimsDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesCellularSimsDataSource{}
)

func NewDevicesCellularSimsDataSource() datasource.DataSource {
	return &DevicesCellularSimsDataSource{}
}

type DevicesCellularSimsDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesCellularSimsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesCellularSimsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_cellular_sims"
}

func (d *DevicesCellularSimsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"sims": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"apns": schema.SetNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"allowed_ip_types": schema.ListAttribute{
												Computed:    true,
												ElementType: types.StringType,
											},
											"authentication": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{

													"type": schema.StringAttribute{
														Computed: true,
													},
													"username": schema.StringAttribute{
														Computed: true,
													},
												},
											},
											"name": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
								"is_primary": schema.BoolAttribute{
									Computed: true,
								},
								"slot": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *DevicesCellularSimsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesCellularSims DevicesCellularSims
	diags := req.Config.Get(ctx, &devicesCellularSims)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceCellularSims")
		vvSerial := devicesCellularSims.Serial.ValueString()

		response1, restyResp1, err := d.client.Devices.GetDeviceCellularSims(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCellularSims",
				err.Error(),
			)
			return
		}

		devicesCellularSims = ResponseDevicesGetDeviceCellularSimsItemToBody(devicesCellularSims, response1)
		diags = resp.State.Set(ctx, &devicesCellularSims)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesCellularSims struct {
	Serial types.String                          `tfsdk:"serial"`
	Item   *ResponseDevicesGetDeviceCellularSims `tfsdk:"item"`
}

type ResponseDevicesGetDeviceCellularSims struct {
	Sims *[]ResponseDevicesGetDeviceCellularSimsSims `tfsdk:"sims"`
}

type ResponseDevicesGetDeviceCellularSimsSims struct {
	Apns      *[]ResponseDevicesGetDeviceCellularSimsSimsApns `tfsdk:"apns"`
	IsPrimary types.Bool                                      `tfsdk:"is_primary"`
	Slot      types.String                                    `tfsdk:"slot"`
}

type ResponseDevicesGetDeviceCellularSimsSimsApns struct {
	AllowedIPTypes types.List                                                  `tfsdk:"allowed_ip_types"`
	Authentication *ResponseDevicesGetDeviceCellularSimsSimsApnsAuthentication `tfsdk:"authentication"`
	Name           types.String                                                `tfsdk:"name"`
}

type ResponseDevicesGetDeviceCellularSimsSimsApnsAuthentication struct {
	Type     types.String `tfsdk:"type"`
	Username types.String `tfsdk:"username"`
}

// ToBody
func ResponseDevicesGetDeviceCellularSimsItemToBody(state DevicesCellularSims, response *merakigosdk.ResponseDevicesGetDeviceCellularSims) DevicesCellularSims {
	itemState := ResponseDevicesGetDeviceCellularSims{
		Sims: func() *[]ResponseDevicesGetDeviceCellularSimsSims {
			if response.Sims != nil {
				result := make([]ResponseDevicesGetDeviceCellularSimsSims, len(*response.Sims))
				for i, sims := range *response.Sims {
					result[i] = ResponseDevicesGetDeviceCellularSimsSims{
						Apns: func() *[]ResponseDevicesGetDeviceCellularSimsSimsApns {
							if sims.Apns != nil {
								result := make([]ResponseDevicesGetDeviceCellularSimsSimsApns, len(*sims.Apns))
								for i, apns := range *sims.Apns {
									result[i] = ResponseDevicesGetDeviceCellularSimsSimsApns{
										AllowedIPTypes: StringSliceToList(apns.AllowedIPTypes),
										Authentication: func() *ResponseDevicesGetDeviceCellularSimsSimsApnsAuthentication {
											if apns.Authentication != nil {
												return &ResponseDevicesGetDeviceCellularSimsSimsApnsAuthentication{
													Type:     types.StringValue(apns.Authentication.Type),
													Username: types.StringValue(apns.Authentication.Username),
												}
											}
											return &ResponseDevicesGetDeviceCellularSimsSimsApnsAuthentication{}
										}(),
										Name: types.StringValue(apns.Name),
									}
								}
								return &result
							}
							return &[]ResponseDevicesGetDeviceCellularSimsSimsApns{}
						}(),
						IsPrimary: func() types.Bool {
							if sims.IsPrimary != nil {
								return types.BoolValue(*sims.IsPrimary)
							}
							return types.Bool{}
						}(),
						Slot: types.StringValue(sims.Slot),
					}
				}
				return &result
			}
			return &[]ResponseDevicesGetDeviceCellularSimsSims{}
		}(),
	}
	state.Item = &itemState
	return state
}
