package provider

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesSwitchRoutingStaticRoutesResource{}
	_ resource.ResourceWithConfigure = &DevicesSwitchRoutingStaticRoutesResource{}
)

func NewDevicesSwitchRoutingStaticRoutesResource() resource.Resource {
	return &DevicesSwitchRoutingStaticRoutesResource{}
}

type DevicesSwitchRoutingStaticRoutesResource struct {
	client *merakigosdk.Client
}

func (r *DevicesSwitchRoutingStaticRoutesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesSwitchRoutingStaticRoutesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_switch_routing_static_routes"
}

func (r *DevicesSwitchRoutingStaticRoutesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"advertise_via_ospf_enabled": schema.BoolAttribute{
				MarkdownDescription: `Option to advertise static routes via OSPF`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name or description of the layer 3 static route`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"next_hop_ip": schema.StringAttribute{
				MarkdownDescription: ` The IP address of the router to which traffic for this destination network should be sent`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"prefer_over_ospf_routes_enabled": schema.BoolAttribute{
				MarkdownDescription: `Option to prefer static routes over OSPF routes`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"static_route_id": schema.StringAttribute{
				MarkdownDescription: `The identifier of a layer 3 static route`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"subnet": schema.StringAttribute{
				MarkdownDescription: `The IP address of the subnetwork specified in CIDR notation (ex. 1.2.3.0/24)`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

//path params to set ['staticRouteId']

func (r *DevicesSwitchRoutingStaticRoutesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesSwitchRoutingStaticRoutesRs

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
	vvName := data.Name.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Switch.GetDeviceSwitchRoutingStaticRoutes(vvSerial)
	//Have Create
	if err != nil || restyResp1 == nil {
		if restyResp1.StatusCode() != 404 {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceSwitchRoutingStaticRoutes",
				err.Error(),
			)
			return
		}
	}
	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvStaticRouteID, ok := result2["StaticRouteID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter StaticRouteID",
					err.Error(),
				)
				return
			}
			r.client.Switch.UpdateDeviceSwitchRoutingStaticRoute(vvSerial, vvStaticRouteID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Switch.GetDeviceSwitchRoutingStaticRoute(vvSerial, vvStaticRouteID)
			if responseVerifyItem2 != nil {
				data = ResponseSwitchGetDeviceSwitchRoutingStaticRouteItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Switch.CreateDeviceSwitchRoutingStaticRoute(vvSerial, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateDeviceSwitchRoutingStaticRoute",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateDeviceSwitchRoutingStaticRoute",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Switch.GetDeviceSwitchRoutingStaticRoutes(vvSerial)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceSwitchRoutingStaticRoutes",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceSwitchRoutingStaticRoutes",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvStaticRouteID, ok := result2["StaticRouteID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter StaticRouteID",
				err.Error(),
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Switch.GetDeviceSwitchRoutingStaticRoute(vvSerial, vvStaticRouteID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseSwitchGetDeviceSwitchRoutingStaticRouteItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetDeviceSwitchRoutingStaticRoute",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceSwitchRoutingStaticRoute",
				err.Error(),
			)
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"Error in result.",
			"Error in result.",
		)
		return
	}
}

func (r *DevicesSwitchRoutingStaticRoutesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesSwitchRoutingStaticRoutesRs

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
	vvStaticRouteID := data.StaticRouteID.ValueString()
	responseGet, restyRespGet, err := r.client.Switch.GetDeviceSwitchRoutingStaticRoute(vvSerial, vvStaticRouteID)
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
				"Failure when executing GetDeviceSwitchRoutingStaticRoute",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceSwitchRoutingStaticRoute",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseSwitchGetDeviceSwitchRoutingStaticRouteItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *DevicesSwitchRoutingStaticRoutesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("static_route_id"), idParts[1])...)
}

func (r *DevicesSwitchRoutingStaticRoutesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesSwitchRoutingStaticRoutesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvSerial := data.Serial.ValueString()
	vvStaticRouteID := data.StaticRouteID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateDeviceSwitchRoutingStaticRoute(vvSerial, vvStaticRouteID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceSwitchRoutingStaticRoute",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceSwitchRoutingStaticRoute",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesSwitchRoutingStaticRoutesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DevicesSwitchRoutingStaticRoutesRs
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &state, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}

	vvSerial := state.Serial.ValueString()
	vvStaticRouteID := state.StaticRouteID.ValueString()
	_, err := r.client.Switch.DeleteDeviceSwitchRoutingStaticRoute(vvSerial, vvStaticRouteID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteDeviceSwitchRoutingStaticRoute", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type DevicesSwitchRoutingStaticRoutesRs struct {
	Serial                      types.String `tfsdk:"serial"`
	StaticRouteID               types.String `tfsdk:"static_route_id"`
	AdvertiseViaOspfEnabled     types.Bool   `tfsdk:"advertise_via_ospf_enabled"`
	Name                        types.String `tfsdk:"name"`
	NextHopIP                   types.String `tfsdk:"next_hop_ip"`
	PreferOverOspfRoutesEnabled types.Bool   `tfsdk:"prefer_over_ospf_routes_enabled"`
	Subnet                      types.String `tfsdk:"subnet"`
}

// FromBody
func (r *DevicesSwitchRoutingStaticRoutesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSwitchCreateDeviceSwitchRoutingStaticRoute {
	emptyString := ""
	advertiseViaOspfEnabled := new(bool)
	if !r.AdvertiseViaOspfEnabled.IsUnknown() && !r.AdvertiseViaOspfEnabled.IsNull() {
		*advertiseViaOspfEnabled = r.AdvertiseViaOspfEnabled.ValueBool()
	} else {
		advertiseViaOspfEnabled = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	nextHopIP := new(string)
	if !r.NextHopIP.IsUnknown() && !r.NextHopIP.IsNull() {
		*nextHopIP = r.NextHopIP.ValueString()
	} else {
		nextHopIP = &emptyString
	}
	preferOverOspfRoutesEnabled := new(bool)
	if !r.PreferOverOspfRoutesEnabled.IsUnknown() && !r.PreferOverOspfRoutesEnabled.IsNull() {
		*preferOverOspfRoutesEnabled = r.PreferOverOspfRoutesEnabled.ValueBool()
	} else {
		preferOverOspfRoutesEnabled = nil
	}
	subnet := new(string)
	if !r.Subnet.IsUnknown() && !r.Subnet.IsNull() {
		*subnet = r.Subnet.ValueString()
	} else {
		subnet = &emptyString
	}
	out := merakigosdk.RequestSwitchCreateDeviceSwitchRoutingStaticRoute{
		AdvertiseViaOspfEnabled:     advertiseViaOspfEnabled,
		Name:                        *name,
		NextHopIP:                   *nextHopIP,
		PreferOverOspfRoutesEnabled: preferOverOspfRoutesEnabled,
		Subnet:                      *subnet,
	}
	return &out
}
func (r *DevicesSwitchRoutingStaticRoutesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingStaticRoute {
	emptyString := ""
	advertiseViaOspfEnabled := new(bool)
	if !r.AdvertiseViaOspfEnabled.IsUnknown() && !r.AdvertiseViaOspfEnabled.IsNull() {
		*advertiseViaOspfEnabled = r.AdvertiseViaOspfEnabled.ValueBool()
	} else {
		advertiseViaOspfEnabled = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	nextHopIP := new(string)
	if !r.NextHopIP.IsUnknown() && !r.NextHopIP.IsNull() {
		*nextHopIP = r.NextHopIP.ValueString()
	} else {
		nextHopIP = &emptyString
	}
	preferOverOspfRoutesEnabled := new(bool)
	if !r.PreferOverOspfRoutesEnabled.IsUnknown() && !r.PreferOverOspfRoutesEnabled.IsNull() {
		*preferOverOspfRoutesEnabled = r.PreferOverOspfRoutesEnabled.ValueBool()
	} else {
		preferOverOspfRoutesEnabled = nil
	}
	subnet := new(string)
	if !r.Subnet.IsUnknown() && !r.Subnet.IsNull() {
		*subnet = r.Subnet.ValueString()
	} else {
		subnet = &emptyString
	}
	out := merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingStaticRoute{
		AdvertiseViaOspfEnabled:     advertiseViaOspfEnabled,
		Name:                        *name,
		NextHopIP:                   *nextHopIP,
		PreferOverOspfRoutesEnabled: preferOverOspfRoutesEnabled,
		Subnet:                      *subnet,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetDeviceSwitchRoutingStaticRouteItemToBodyRs(state DevicesSwitchRoutingStaticRoutesRs, response *merakigosdk.ResponseSwitchGetDeviceSwitchRoutingStaticRoute, is_read bool) DevicesSwitchRoutingStaticRoutesRs {
	itemState := DevicesSwitchRoutingStaticRoutesRs{
		AdvertiseViaOspfEnabled: func() types.Bool {
			if response.AdvertiseViaOspfEnabled != nil {
				return types.BoolValue(*response.AdvertiseViaOspfEnabled)
			}
			return types.Bool{}
		}(),
		Name:      types.StringValue(response.Name),
		NextHopIP: types.StringValue(response.NextHopIP),
		PreferOverOspfRoutesEnabled: func() types.Bool {
			if response.PreferOverOspfRoutesEnabled != nil {
				return types.BoolValue(*response.PreferOverOspfRoutesEnabled)
			}
			return types.Bool{}
		}(),
		StaticRouteID: types.StringValue(response.StaticRouteID),
		Subnet:        types.StringValue(response.Subnet),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesSwitchRoutingStaticRoutesRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesSwitchRoutingStaticRoutesRs)
}
