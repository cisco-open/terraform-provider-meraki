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
	_ resource.Resource              = &NetworksSwitchStacksRoutingStaticRoutesResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchStacksRoutingStaticRoutesResource{}
)

func NewNetworksSwitchStacksRoutingStaticRoutesResource() resource.Resource {
	return &NetworksSwitchStacksRoutingStaticRoutesResource{}
}

type NetworksSwitchStacksRoutingStaticRoutesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchStacksRoutingStaticRoutesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchStacksRoutingStaticRoutesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_stacks_routing_static_routes"
}

func (r *NetworksSwitchStacksRoutingStaticRoutesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
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
			"switch_stack_id": schema.StringAttribute{
				MarkdownDescription: `switchStackId path parameter. Switch stack ID`,
				Required:            true,
			},
		},
	}
}

//path params to set ['staticRouteId']

func (r *NetworksSwitchStacksRoutingStaticRoutesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchStacksRoutingStaticRoutesRs

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
	vvNetworkID := data.NetworkID.ValueString()
	vvSwitchStackID := data.SwitchStackID.ValueString()
	vvName := data.Name.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Switch.GetNetworkSwitchStackRoutingStaticRoutes(vvNetworkID, vvSwitchStackID)
	//Have Create
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkSwitchStackRoutingStaticRoutes",
					err.Error(),
				)
				return
			}
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
					"Fail Parsing StaticRouteID",
				)
				return
			}
			r.client.Switch.UpdateNetworkSwitchStackRoutingStaticRoute(vvNetworkID, vvSwitchStackID, vvStaticRouteID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Switch.GetNetworkSwitchStackRoutingStaticRoute(vvNetworkID, vvSwitchStackID, vvStaticRouteID)
			if responseVerifyItem2 != nil {
				data = ResponseSwitchGetNetworkSwitchStackRoutingStaticRouteItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Switch.CreateNetworkSwitchStackRoutingStaticRoute(vvNetworkID, vvSwitchStackID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkSwitchStackRoutingStaticRoute",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkSwitchStackRoutingStaticRoute",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Switch.GetNetworkSwitchStackRoutingStaticRoutes(vvNetworkID, vvSwitchStackID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStackRoutingStaticRoutes",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchStackRoutingStaticRoutes",
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
		responseVerifyItem2, restyRespGet, err := r.client.Switch.GetNetworkSwitchStackRoutingStaticRoute(vvNetworkID, vvSwitchStackID, vvStaticRouteID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseSwitchGetNetworkSwitchStackRoutingStaticRouteItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkSwitchStackRoutingStaticRoute",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStackRoutingStaticRoute",
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

func (r *NetworksSwitchStacksRoutingStaticRoutesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchStacksRoutingStaticRoutesRs

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

	vvNetworkID := data.NetworkID.ValueString()
	vvSwitchStackID := data.SwitchStackID.ValueString()
	vvStaticRouteID := data.StaticRouteID.ValueString()
	responseGet, restyRespGet, err := r.client.Switch.GetNetworkSwitchStackRoutingStaticRoute(vvNetworkID, vvSwitchStackID, vvStaticRouteID)
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
				"Failure when executing GetNetworkSwitchStackRoutingStaticRoute",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchStackRoutingStaticRoute",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseSwitchGetNetworkSwitchStackRoutingStaticRouteItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchStacksRoutingStaticRoutesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 3 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("switch_stack_id"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("static_route_id"), idParts[2])...)
}

func (r *NetworksSwitchStacksRoutingStaticRoutesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSwitchStacksRoutingStaticRoutesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvSwitchStackID := data.SwitchStackID.ValueString()
	vvStaticRouteID := data.StaticRouteID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchStackRoutingStaticRoute(vvNetworkID, vvSwitchStackID, vvStaticRouteID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchStackRoutingStaticRoute",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchStackRoutingStaticRoute",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchStacksRoutingStaticRoutesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksSwitchStacksRoutingStaticRoutesRs
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

	vvNetworkID := state.NetworkID.ValueString()
	vvSwitchStackID := state.SwitchStackID.ValueString()
	vvStaticRouteID := state.StaticRouteID.ValueString()
	_, err := r.client.Switch.DeleteNetworkSwitchStackRoutingStaticRoute(vvNetworkID, vvSwitchStackID, vvStaticRouteID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkSwitchStackRoutingStaticRoute", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksSwitchStacksRoutingStaticRoutesRs struct {
	NetworkID                   types.String `tfsdk:"network_id"`
	SwitchStackID               types.String `tfsdk:"switch_stack_id"`
	StaticRouteID               types.String `tfsdk:"static_route_id"`
	AdvertiseViaOspfEnabled     types.Bool   `tfsdk:"advertise_via_ospf_enabled"`
	Name                        types.String `tfsdk:"name"`
	NextHopIP                   types.String `tfsdk:"next_hop_ip"`
	PreferOverOspfRoutesEnabled types.Bool   `tfsdk:"prefer_over_ospf_routes_enabled"`
	Subnet                      types.String `tfsdk:"subnet"`
}

// FromBody
func (r *NetworksSwitchStacksRoutingStaticRoutesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSwitchCreateNetworkSwitchStackRoutingStaticRoute {
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
	out := merakigosdk.RequestSwitchCreateNetworkSwitchStackRoutingStaticRoute{
		AdvertiseViaOspfEnabled:     advertiseViaOspfEnabled,
		Name:                        *name,
		NextHopIP:                   *nextHopIP,
		PreferOverOspfRoutesEnabled: preferOverOspfRoutesEnabled,
		Subnet:                      *subnet,
	}
	return &out
}
func (r *NetworksSwitchStacksRoutingStaticRoutesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateNetworkSwitchStackRoutingStaticRoute {
	emptyString := ""
	// advertiseViaOspfEnabled := new(bool)
	// if !r.AdvertiseViaOspfEnabled.IsUnknown() && !r.AdvertiseViaOspfEnabled.IsNull() {
	// 	*advertiseViaOspfEnabled = r.AdvertiseViaOspfEnabled.ValueBool()
	// } else {
	// 	advertiseViaOspfEnabled = nil
	// }
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
	// preferOverOspfRoutesEnabled := new(bool)
	// if !r.PreferOverOspfRoutesEnabled.IsUnknown() && !r.PreferOverOspfRoutesEnabled.IsNull() {
	// 	*preferOverOspfRoutesEnabled = r.PreferOverOspfRoutesEnabled.ValueBool()
	// } else {
	// 	preferOverOspfRoutesEnabled = nil
	// }
	subnet := new(string)
	if !r.Subnet.IsUnknown() && !r.Subnet.IsNull() {
		*subnet = r.Subnet.ValueString()
	} else {
		subnet = &emptyString
	}
	out := merakigosdk.RequestSwitchUpdateNetworkSwitchStackRoutingStaticRoute{
		// AdvertiseViaOspfEnabled:     advertiseViaOspfEnabled,
		Name:      *name,
		NextHopIP: *nextHopIP,
		// PreferOverOspfRoutesEnabled: preferOverOspfRoutesEnabled,
		Subnet: *subnet,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchStackRoutingStaticRouteItemToBodyRs(state NetworksSwitchStacksRoutingStaticRoutesRs, response *merakigosdk.ResponseSwitchGetNetworkSwitchStackRoutingStaticRoute, is_read bool) NetworksSwitchStacksRoutingStaticRoutesRs {
	itemState := NetworksSwitchStacksRoutingStaticRoutesRs{
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
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSwitchStacksRoutingStaticRoutesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSwitchStacksRoutingStaticRoutesRs)
}
