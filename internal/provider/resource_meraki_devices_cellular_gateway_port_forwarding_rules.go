package provider

// RESOURCE NORMAL
import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesCellularGatewayPortForwardingRulesResource{}
	_ resource.ResourceWithConfigure = &DevicesCellularGatewayPortForwardingRulesResource{}
)

func NewDevicesCellularGatewayPortForwardingRulesResource() resource.Resource {
	return &DevicesCellularGatewayPortForwardingRulesResource{}
}

type DevicesCellularGatewayPortForwardingRulesResource struct {
	client *merakigosdk.Client
}

func (r *DevicesCellularGatewayPortForwardingRulesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesCellularGatewayPortForwardingRulesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_cellular_gateway_port_forwarding_rules"
}

func (r *DevicesCellularGatewayPortForwardingRulesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"rules": schema.SetNestedAttribute{
				MarkdownDescription: `An array of port forwarding params`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"access": schema.StringAttribute{
							MarkdownDescription: `*any* or *restricted*. Specify the right to make inbound connections on the specified ports or port ranges. If *restricted*, a list of allowed IPs is mandatory.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"allowed_ips": schema.SetAttribute{
							MarkdownDescription: `An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
						},
						"lan_ip": schema.StringAttribute{
							MarkdownDescription: `The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"local_port": schema.StringAttribute{
							MarkdownDescription: `A port or port ranges that will receive the forwarded traffic from the WAN`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `A descriptive name for the rule`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: `TCP or UDP`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"public_port": schema.StringAttribute{
							MarkdownDescription: `A port or port ranges that will be forwarded to the host on the LAN`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"uplink": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
		},
	}
}

func (r *DevicesCellularGatewayPortForwardingRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesCellularGatewayPortForwardingRulesRs

	var item types.Object
	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	vvSerial := data.Serial.ValueString()
	//Item
	responseVerifyItem, restyResp1, err := r.client.CellularGateway.GetDeviceCellularGatewayPortForwardingRules(vvSerial)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource DevicesCellularGatewayPortForwardingRules only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource DevicesCellularGatewayPortForwardingRules only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.CellularGateway.UpdateDeviceCellularGatewayPortForwardingRules(vvSerial, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceCellularGatewayPortForwardingRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceCellularGatewayPortForwardingRules",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.CellularGateway.GetDeviceCellularGatewayPortForwardingRules(vvSerial)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCellularGatewayPortForwardingRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceCellularGatewayPortForwardingRules",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesCellularGatewayPortForwardingRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesCellularGatewayPortForwardingRulesRs

	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	// Has Item2

	vvSerial := data.Serial.ValueString()
	responseGet, restyRespGet, err := r.client.CellularGateway.GetDeviceCellularGatewayPortForwardingRules(vvSerial)
	if err != nil || restyRespGet == nil {
		if restyRespGet != nil {
			if restyRespGet.StatusCode() == 404 {
				resp.Diagnostics.AddWarning(
					"Resource not found",
					"Deleting resource",
				)
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCellularGatewayPortForwardingRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceCellularGatewayPortForwardingRules",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *DevicesCellularGatewayPortForwardingRulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), req.ID)...)
}

func (r *DevicesCellularGatewayPortForwardingRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesCellularGatewayPortForwardingRulesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvSerial := data.Serial.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.CellularGateway.UpdateDeviceCellularGatewayPortForwardingRules(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceCellularGatewayPortForwardingRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceCellularGatewayPortForwardingRules",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesCellularGatewayPortForwardingRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting DevicesCellularGatewayPortForwardingRules", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesCellularGatewayPortForwardingRulesRs struct {
	Serial types.String                                                                 `tfsdk:"serial"`
	Rules  *[]ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRulesRs `tfsdk:"rules"`
}

type ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRulesRs struct {
	Access     types.String `tfsdk:"access"`
	LanIP      types.String `tfsdk:"lan_ip"`
	LocalPort  types.String `tfsdk:"local_port"`
	Name       types.String `tfsdk:"name"`
	Protocol   types.String `tfsdk:"protocol"`
	PublicPort types.String `tfsdk:"public_port"`
	Uplink     types.String `tfsdk:"uplink"`
	AllowedIPs types.Set    `tfsdk:"allowed_ips"`
}

// FromBody
func (r *DevicesCellularGatewayPortForwardingRulesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRules {
	var requestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRulesRules []merakigosdk.RequestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRulesRules
	if r.Rules != nil {
		for _, rItem1 := range *r.Rules {
			access := rItem1.Access.ValueString()
			var allowedIPs []string = nil
			//Hoola aqui
			rItem1.AllowedIPs.ElementsAs(ctx, &allowedIPs, false)
			lanIP := rItem1.LanIP.ValueString()
			localPort := rItem1.LocalPort.ValueString()
			name := rItem1.Name.ValueString()
			protocol := rItem1.Protocol.ValueString()
			publicPort := rItem1.PublicPort.ValueString()
			requestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRulesRules = append(requestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRulesRules, merakigosdk.RequestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRulesRules{
				Access:     access,
				AllowedIPs: allowedIPs,
				LanIP:      lanIP,
				LocalPort:  localPort,
				Name:       name,
				Protocol:   protocol,
				PublicPort: publicPort,
			})
		}
	}
	out := merakigosdk.RequestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRules{
		Rules: func() *[]merakigosdk.RequestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRulesRules {
			if len(requestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRulesRules) > 0 {
				return &requestCellularGatewayUpdateDeviceCellularGatewayPortForwardingRulesRules
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesItemToBodyRs(state DevicesCellularGatewayPortForwardingRulesRs, response *merakigosdk.ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRules, is_read bool) DevicesCellularGatewayPortForwardingRulesRs {
	itemState := DevicesCellularGatewayPortForwardingRulesRs{
		Rules: func() *[]ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRulesRs {
			if response.Rules != nil {
				result := make([]ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRulesRs, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRulesRs{
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
			return &[]ResponseCellularGatewayGetDeviceCellularGatewayPortForwardingRulesRulesRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesCellularGatewayPortForwardingRulesRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesCellularGatewayPortForwardingRulesRs)
}
