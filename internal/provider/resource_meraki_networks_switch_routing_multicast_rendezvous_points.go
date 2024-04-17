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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSwitchRoutingMulticastRendezvousPointsResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchRoutingMulticastRendezvousPointsResource{}
)

func NewNetworksSwitchRoutingMulticastRendezvousPointsResource() resource.Resource {
	return &NetworksSwitchRoutingMulticastRendezvousPointsResource{}
}

type NetworksSwitchRoutingMulticastRendezvousPointsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchRoutingMulticastRendezvousPointsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchRoutingMulticastRendezvousPointsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_routing_multicast_rendezvous_points"
}

func (r *NetworksSwitchRoutingMulticastRendezvousPointsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"interface_ip": schema.StringAttribute{
				MarkdownDescription: `The IP address of the interface where the RP needs to be created.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"interface_name": schema.StringAttribute{
				Computed: true,
			},
			"multicast_group": schema.StringAttribute{
				MarkdownDescription: `'Any', or the IP address of a multicast group`,
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
			"rendezvous_point_id": schema.StringAttribute{
				MarkdownDescription: `rendezvousPointId path parameter. Rendezvous point ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"serial": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

//path params to set ['rendezvousPointId']

func (r *NetworksSwitchRoutingMulticastRendezvousPointsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchRoutingMulticastRendezvousPointsRs

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
	// network_id
	//Reviw This  Has Item and item
	//HAS CREATE

	vvRendezvousPointID := data.RendezvousPointID.ValueString()
	if vvRendezvousPointID != "" {
		responseVerifyItem, restyRespGet, err := r.client.Switch.GetNetworkSwitchRoutingMulticastRendezvousPoint(vvNetworkID, vvRendezvousPointID)
		if err != nil || responseVerifyItem == nil {
			if restyRespGet != nil {
				if restyRespGet.StatusCode() != 404 {

					resp.Diagnostics.AddError(
						"Failure when executing GetNetworkSwitchRoutingMulticastRendezvousPoint",
						err.Error(),
					)
					return
				}
			}
		}

		if responseVerifyItem != nil {
			data = ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPointItemToBodyRs(data, responseVerifyItem, false)
			diags := resp.State.Set(ctx, &data)
			resp.Diagnostics.Append(diags...)
			return
		}
	}

	response, restyResp2, err := r.client.Switch.CreateNetworkSwitchRoutingMulticastRendezvousPoint(vvNetworkID, data.toSdkApiRequestCreate(ctx))

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing ",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing ",
			err.Error(),
		)
		return
	}
	//Items
	vvRendezvousPointID = response.RendezvousPointID
	responseGet, restyResp1, err := r.client.Switch.GetNetworkSwitchRoutingMulticastRendezvousPoint(vvNetworkID, vvRendezvousPointID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchRoutingMulticastRendezvousPoints",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchRoutingMulticastRendezvousPoints",
			err.Error(),
		)
		return
	} else {
		data = ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPointItemToBodyRs(data, responseGet, false)
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}
}

func (r *NetworksSwitchRoutingMulticastRendezvousPointsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchRoutingMulticastRendezvousPointsRs

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
	// network_id
	vvRendezvousPointID := data.RendezvousPointID.ValueString()
	// rendezvous_point_id
	responseGet, restyRespGet, err := r.client.Switch.GetNetworkSwitchRoutingMulticastRendezvousPoint(vvNetworkID, vvRendezvousPointID)
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
				"Failure when executing GetNetworkSwitchRoutingMulticastRendezvousPoint",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchRoutingMulticastRendezvousPoint",
			err.Error(),
		)
		return
	}

	data = ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPointItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksSwitchRoutingMulticastRendezvousPointsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("rendezvous_point_id"), idParts[1])...)
}

func (r *NetworksSwitchRoutingMulticastRendezvousPointsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSwitchRoutingMulticastRendezvousPointsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvRendezvousPointID := data.RendezvousPointID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	_, restyResp2, err := r.client.Switch.UpdateNetworkSwitchRoutingMulticastRendezvousPoint(vvNetworkID, vvRendezvousPointID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchRoutingMulticastRendezvousPoint",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchRoutingMulticastRendezvousPoint",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchRoutingMulticastRendezvousPointsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksSwitchRoutingMulticastRendezvousPointsRs
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
	vvRendezvousPointID := state.RendezvousPointID.ValueString()
	_, err := r.client.Switch.DeleteNetworkSwitchRoutingMulticastRendezvousPoint(vvNetworkID, vvRendezvousPointID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkSwitchRoutingMulticastRendezvousPoint", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksSwitchRoutingMulticastRendezvousPointsRs struct {
	NetworkID         types.String `tfsdk:"network_id"`
	RendezvousPointID types.String `tfsdk:"rendezvous_point_id"`
	InterfaceIP       types.String `tfsdk:"interface_ip"`
	InterfaceName     types.String `tfsdk:"interface_name"`
	MulticastGroup    types.String `tfsdk:"multicast_group"`
	Serial            types.String `tfsdk:"serial"`
}

// FromBody
func (r *NetworksSwitchRoutingMulticastRendezvousPointsRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSwitchCreateNetworkSwitchRoutingMulticastRendezvousPoint {
	emptyString := ""
	interfaceIP := new(string)
	if !r.InterfaceIP.IsUnknown() && !r.InterfaceIP.IsNull() {
		*interfaceIP = r.InterfaceIP.ValueString()
	} else {
		interfaceIP = &emptyString
	}
	multicastGroup := new(string)
	if !r.MulticastGroup.IsUnknown() && !r.MulticastGroup.IsNull() {
		*multicastGroup = r.MulticastGroup.ValueString()
	} else {
		multicastGroup = &emptyString
	}
	out := merakigosdk.RequestSwitchCreateNetworkSwitchRoutingMulticastRendezvousPoint{
		InterfaceIP:    *interfaceIP,
		MulticastGroup: *multicastGroup,
	}
	return &out
}
func (r *NetworksSwitchRoutingMulticastRendezvousPointsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingMulticastRendezvousPoint {
	emptyString := ""
	interfaceIP := new(string)
	if !r.InterfaceIP.IsUnknown() && !r.InterfaceIP.IsNull() {
		*interfaceIP = r.InterfaceIP.ValueString()
	} else {
		interfaceIP = &emptyString
	}
	multicastGroup := new(string)
	if !r.MulticastGroup.IsUnknown() && !r.MulticastGroup.IsNull() {
		*multicastGroup = r.MulticastGroup.ValueString()
	} else {
		multicastGroup = &emptyString
	}
	out := merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingMulticastRendezvousPoint{
		InterfaceIP:    *interfaceIP,
		MulticastGroup: *multicastGroup,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPointItemToBodyRs(state NetworksSwitchRoutingMulticastRendezvousPointsRs, response *merakigosdk.ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoint, is_read bool) NetworksSwitchRoutingMulticastRendezvousPointsRs {
	itemState := NetworksSwitchRoutingMulticastRendezvousPointsRs{
		InterfaceIP:       types.StringValue(response.InterfaceIP),
		InterfaceName:     types.StringValue(response.InterfaceName),
		MulticastGroup:    types.StringValue(response.MulticastGroup),
		RendezvousPointID: types.StringValue(response.RendezvousPointID),
		Serial:            types.StringValue(response.Serial),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSwitchRoutingMulticastRendezvousPointsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSwitchRoutingMulticastRendezvousPointsRs)
}
