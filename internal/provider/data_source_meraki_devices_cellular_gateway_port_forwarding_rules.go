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
	_ datasource.DataSource              = &DevicesCellularGatewayPortForwardingRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesCellularGatewayPortForwardingRulesDataSource{}
)

func NewDevicesCellularGatewayPortForwardingRulesDataSource() datasource.DataSource {
	return &DevicesCellularGatewayPortForwardingRulesDataSource{}
}

type DevicesCellularGatewayPortForwardingRulesDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesCellularGatewayPortForwardingRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesCellularGatewayPortForwardingRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_cellular_gateway_port_forwarding_rules"
}

func (d *DevicesCellularGatewayPortForwardingRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"rules": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"access": schema.StringAttribute{
									Computed: true,
								},
								"lan_ip": schema.StringAttribute{
									Computed: true,
								},
								"local_port": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"protocol": schema.StringAttribute{
									Computed: true,
								},
								"public_port": schema.StringAttribute{
									Computed: true,
								},
								"uplink": schema.StringAttribute{
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

func (d *DevicesCellularGatewayPortForwardingRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesCellularGatewayPortForwardingRules DevicesCellularGatewayPortForwardingRules
	diags := req.Config.Get(ctx, &devicesCellularGatewayPortForwardingRules)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceCellularGatewayPortForwardingRules")
		vvSerial := devicesCellularGatewayPortForwardingRules.Serial.ValueString()

		response1, restyResp1, err := d.client.CellularGateway.GetDeviceCellularGatewayPortForwardingRules(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCellularGatewayPortForwardingRules",
				err.Error(),
			)
			return
		}

		devicesCellularGatewayPortForwardingRules = ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesItemToBody(devicesCellularGatewayPortForwardingRules, response1)
		diags = resp.State.Set(ctx, &devicesCellularGatewayPortForwardingRules)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesCellularGatewayPortForwardingRules struct {
	Serial types.String                                                        `tfsdk:"serial"`
	Item   *ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRules `tfsdk:"item"`
}

type ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRules struct {
	Rules *[]ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRules `tfsdk:"rules"`
}

type ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRules struct {
	Access     types.String `tfsdk:"access"`
	LanIP      types.String `tfsdk:"lan_ip"`
	LocalPort  types.String `tfsdk:"local_port"`
	Name       types.String `tfsdk:"name"`
	Protocol   types.String `tfsdk:"protocol"`
	PublicPort types.String `tfsdk:"public_port"`
	Uplink     types.String `tfsdk:"uplink"`
}

// ToBody
func ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesItemToBody(state DevicesCellularGatewayPortForwardingRules, response *merakigosdk.ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRules) DevicesCellularGatewayPortForwardingRules {
	itemState := ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRules{
		Rules: func() *[]ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRules {
			if response.Rules != nil {
				result := make([]ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRules{
						Access:     types.StringValue(rules.Access),
						LanIP:      types.StringValue(rules.LanIP),
						LocalPort:  types.StringValue(rules.LocalPort),
						Name:       types.StringValue(rules.Name),
						Protocol:   types.StringValue(rules.Protocol),
						PublicPort: types.StringValue(rules.PublicPort),
						Uplink:     types.StringValue(rules.Uplink),
					}
				}
				return &result
			}
			return &[]ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRules{}
		}(),
	}
	state.Item = &itemState
	return state
}
