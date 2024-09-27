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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSwitchLinkAggregationsResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchLinkAggregationsResource{}
)

func NewNetworksSwitchLinkAggregationsResource() resource.Resource {
	return &NetworksSwitchLinkAggregationsResource{}
}

type NetworksSwitchLinkAggregationsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchLinkAggregationsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchLinkAggregationsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_link_aggregations"
}

func (r *NetworksSwitchLinkAggregationsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"link_aggregation_id": schema.StringAttribute{
				MarkdownDescription: `linkAggregationId path parameter. Link aggregation ID`,
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
			"switch_ports": schema.SetNestedAttribute{
				MarkdownDescription: `Array of switch or stack ports for creating aggregation group. Minimum 2 and maximum 8 ports are supported.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"port_id": schema.StringAttribute{
							MarkdownDescription: `Port identifier of switch port. For modules, the identifier is "SlotNumber_ModuleType_PortNumber" (Ex: "1_8X10G_1"), otherwise it is just the port number (Ex: "8").`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `Serial number of the switch.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"switch_profile_ports": schema.SetNestedAttribute{
				MarkdownDescription: `Array of switch profile ports for creating aggregation group. Minimum 2 and maximum 8 ports are supported.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"port_id": schema.StringAttribute{
							MarkdownDescription: `Port identifier of switch port. For modules, the identifier is "SlotNumber_ModuleType_PortNumber" (Ex: "1_8X10G_1"), otherwise it is just the port number (Ex: "8").`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"profile": schema.StringAttribute{
							MarkdownDescription: `Profile identifier.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
		},
	}
}

//path params to set ['linkAggregationId']

func (r *NetworksSwitchLinkAggregationsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchLinkAggregationsRs

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

	// Create
	// Review This One
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Switch.CreateNetworkSwitchLinkAggregation(vvNetworkID, dataRequest)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkSwitchLinkAggregation",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkSwitchLinkAggregation",
			err.Error(),
		)
		return
	}
	vvID := response.ID
	//Items
	responseGet, restyResp1, err := r.client.Switch.GetNetworkSwitchLinkAggregations(vvNetworkID)
	// Has only items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchLinkAggregations",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchLinkAggregations",
			err.Error(),
		)
		return
	}
	responseStruct2 := structToMap(responseGet)
	result2 := getDictResult(responseStruct2, "ID", vvID, simpleCmp)
	var responseVerifyItem2 merakigosdk.ResponseItemSwitchGetNetworkSwitchLinkAggregations
	if result2 != nil {
		err := mapToStruct(result2.(map[string]interface{}), &responseVerifyItem2)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing mapToStruct in resource",
				err.Error(),
			)
			return
		}
		data = ResponseSwitchGetNetworkSwitchLinkAggregationsItemToBodyRs(data, &responseVerifyItem2, false)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	} else {
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchLinkAggregations Result",
			"Not Found",
		)
		return
	}
}

func (r *NetworksSwitchLinkAggregationsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchLinkAggregationsRs

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
	// Not has Item

	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvID := data.ID.ValueString()
	// id

	responseGet, restyResp1, err := r.client.Switch.GetNetworkSwitchLinkAggregations(vvNetworkID)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchLinkAggregations",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchLinkAggregations",
			err.Error(),
		)
		return
	}
	responseStruct2 := structToMap(responseGet)
	result2 := getDictResult(responseStruct2, "ID", vvID, simpleCmp)
	var responseVerifyItem2 merakigosdk.ResponseItemSwitchGetNetworkSwitchLinkAggregations
	if result2 != nil {
		err := mapToStruct(result2.(map[string]interface{}), &responseVerifyItem2)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing mapToStruct in resource",
				err.Error(),
			)
			return
		}
		//entro aqui
		data = ResponseSwitchGetNetworkSwitchLinkAggregationsItemToBodyRs(data, &responseVerifyItem2, true)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	} else {
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchLinkAggregations Result",
			"Not Found",
		)
		return
	}
}

func (r *NetworksSwitchLinkAggregationsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
}

func (r *NetworksSwitchLinkAggregationsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSwitchLinkAggregationsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvLinkAggregationID := data.ID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Switch.UpdateNetworkSwitchLinkAggregation(vvNetworkID, vvLinkAggregationID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchLinkAggregation",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchLinkAggregation",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchLinkAggregationsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksSwitchLinkAggregationsRs
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
	vvLinkAggregationID := state.ID.ValueString()
	_, err := r.client.Switch.DeleteNetworkSwitchLinkAggregation(vvNetworkID, vvLinkAggregationID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkSwitchLinkAggregation", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksSwitchLinkAggregationsRs struct {
	NetworkID         types.String `tfsdk:"network_id"`
	LinkAggregationID types.String `tfsdk:"link_aggregation_id"`
	//TIENE ITEMS
	ID                 types.String                                                            `tfsdk:"id"`
	SwitchPorts        *[]ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPortsRs      `tfsdk:"switch_ports"`
	SwitchProfilePorts *[]ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchProfilePorts `tfsdk:"switch_profile_ports"`
}

type ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPortsRs struct {
	PortID types.String `tfsdk:"port_id"`
	Serial types.String `tfsdk:"serial"`
}

type ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchProfilePorts struct {
	PortID  types.String `tfsdk:"port_id"`
	Profile types.String `tfsdk:"profile"`
}

// FromBody
func (r *NetworksSwitchLinkAggregationsRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSwitchCreateNetworkSwitchLinkAggregation {
	var requestSwitchCreateNetworkSwitchLinkAggregationSwitchPorts []merakigosdk.RequestSwitchCreateNetworkSwitchLinkAggregationSwitchPorts
	if r.SwitchPorts != nil {
		for _, rItem1 := range *r.SwitchPorts {
			portID := rItem1.PortID.ValueString()
			serial := rItem1.Serial.ValueString()
			requestSwitchCreateNetworkSwitchLinkAggregationSwitchPorts = append(requestSwitchCreateNetworkSwitchLinkAggregationSwitchPorts, merakigosdk.RequestSwitchCreateNetworkSwitchLinkAggregationSwitchPorts{
				PortID: portID,
				Serial: serial,
			})
		}
	}
	var requestSwitchCreateNetworkSwitchLinkAggregationSwitchProfilePorts []merakigosdk.RequestSwitchCreateNetworkSwitchLinkAggregationSwitchProfilePorts
	if r.SwitchProfilePorts != nil {
		for _, rItem1 := range *r.SwitchProfilePorts {
			portID := rItem1.PortID.ValueString()
			profile := rItem1.Profile.ValueString()
			requestSwitchCreateNetworkSwitchLinkAggregationSwitchProfilePorts = append(requestSwitchCreateNetworkSwitchLinkAggregationSwitchProfilePorts, merakigosdk.RequestSwitchCreateNetworkSwitchLinkAggregationSwitchProfilePorts{
				PortID:  portID,
				Profile: profile,
			})
		}
	}
	out := merakigosdk.RequestSwitchCreateNetworkSwitchLinkAggregation{
		SwitchPorts: func() *[]merakigosdk.RequestSwitchCreateNetworkSwitchLinkAggregationSwitchPorts {
			if len(requestSwitchCreateNetworkSwitchLinkAggregationSwitchPorts) > 0 {
				return &requestSwitchCreateNetworkSwitchLinkAggregationSwitchPorts
			}
			return nil
		}(),
		SwitchProfilePorts: func() *[]merakigosdk.RequestSwitchCreateNetworkSwitchLinkAggregationSwitchProfilePorts {
			if len(requestSwitchCreateNetworkSwitchLinkAggregationSwitchProfilePorts) > 0 {
				return &requestSwitchCreateNetworkSwitchLinkAggregationSwitchProfilePorts
			}
			return nil
		}(),
	}
	return &out
}
func (r *NetworksSwitchLinkAggregationsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateNetworkSwitchLinkAggregation {
	var requestSwitchUpdateNetworkSwitchLinkAggregationSwitchPorts []merakigosdk.RequestSwitchUpdateNetworkSwitchLinkAggregationSwitchPorts
	if r.SwitchPorts != nil {
		for _, rItem1 := range *r.SwitchPorts {
			portID := rItem1.PortID.ValueString()
			serial := rItem1.Serial.ValueString()
			requestSwitchUpdateNetworkSwitchLinkAggregationSwitchPorts = append(requestSwitchUpdateNetworkSwitchLinkAggregationSwitchPorts, merakigosdk.RequestSwitchUpdateNetworkSwitchLinkAggregationSwitchPorts{
				PortID: portID,
				Serial: serial,
			})
		}
	}
	var requestSwitchUpdateNetworkSwitchLinkAggregationSwitchProfilePorts []merakigosdk.RequestSwitchUpdateNetworkSwitchLinkAggregationSwitchProfilePorts
	if r.SwitchProfilePorts != nil {
		for _, rItem1 := range *r.SwitchProfilePorts {
			portID := rItem1.PortID.ValueString()
			profile := rItem1.Profile.ValueString()
			requestSwitchUpdateNetworkSwitchLinkAggregationSwitchProfilePorts = append(requestSwitchUpdateNetworkSwitchLinkAggregationSwitchProfilePorts, merakigosdk.RequestSwitchUpdateNetworkSwitchLinkAggregationSwitchProfilePorts{
				PortID:  portID,
				Profile: profile,
			})
		}
	}
	out := merakigosdk.RequestSwitchUpdateNetworkSwitchLinkAggregation{
		SwitchPorts: func() *[]merakigosdk.RequestSwitchUpdateNetworkSwitchLinkAggregationSwitchPorts {
			if len(requestSwitchUpdateNetworkSwitchLinkAggregationSwitchPorts) > 0 {
				return &requestSwitchUpdateNetworkSwitchLinkAggregationSwitchPorts
			}
			return nil
		}(),
		SwitchProfilePorts: func() *[]merakigosdk.RequestSwitchUpdateNetworkSwitchLinkAggregationSwitchProfilePorts {
			if len(requestSwitchUpdateNetworkSwitchLinkAggregationSwitchProfilePorts) > 0 {
				return &requestSwitchUpdateNetworkSwitchLinkAggregationSwitchProfilePorts
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchLinkAggregationsItemToBodyRs(state NetworksSwitchLinkAggregationsRs, response *merakigosdk.ResponseItemSwitchGetNetworkSwitchLinkAggregations, is_read bool) NetworksSwitchLinkAggregationsRs {
	itemState := NetworksSwitchLinkAggregationsRs{
		ID: types.StringValue(response.ID),
		SwitchPorts: func() *[]ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPortsRs {
			if response.SwitchPorts != nil {
				result := make([]ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPortsRs, len(*response.SwitchPorts))
				for i, switchPorts := range *response.SwitchPorts {
					result[i] = ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPortsRs{
						PortID: types.StringValue(switchPorts.PortID),
						Serial: types.StringValue(switchPorts.Serial),
					}
				}
				return &result
			}
			return &[]ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPortsRs{}
		}(),
	}
	itemState.NetworkID = state.NetworkID
	state = itemState
	return state
}
