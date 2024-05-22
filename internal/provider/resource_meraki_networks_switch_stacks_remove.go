package provider

// RESOURCE ACTION

import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSwitchStacksRemoveResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchStacksRemoveResource{}
)

func NewNetworksSwitchStacksRemoveResource() resource.Resource {
	return &NetworksSwitchStacksRemoveResource{}
}

type NetworksSwitchStacksRemoveResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchStacksRemoveResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchStacksRemoveResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_stacks_remove"
}

// resourceAction
func (r *NetworksSwitchStacksRemoveResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"switch_stack_id": schema.StringAttribute{
				MarkdownDescription: `switchStackId path parameter. Switch stack ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"id": schema.StringAttribute{
						MarkdownDescription: `ID of the Switch stack`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of the Switch stack`,
						Computed:            true,
					},
					"serials": schema.SetAttribute{
						MarkdownDescription: `Serials of the switches in the switch stack`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"serial": schema.StringAttribute{
						MarkdownDescription: `The serial of the switch to be removed`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
				},
			},
		},
	}
}
func (r *NetworksSwitchStacksRemoveResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchStacksRemove

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
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Switch.RemoveNetworkSwitchStack(vvNetworkID, vvSwitchStackID, dataRequest)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing RemoveNetworkSwitchStack",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing RemoveNetworkSwitchStack",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseSwitchRemoveNetworkSwitchStackItemToBody(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchStacksRemoveResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSwitchStacksRemoveResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSwitchStacksRemoveResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSwitchStacksRemove struct {
	NetworkID     types.String                             `tfsdk:"network_id"`
	SwitchStackID types.String                             `tfsdk:"switch_stack_id"`
	Item          *ResponseSwitchRemoveNetworkSwitchStack  `tfsdk:"item"`
	Parameters    *RequestSwitchRemoveNetworkSwitchStackRs `tfsdk:"parameters"`
}

type ResponseSwitchRemoveNetworkSwitchStack struct {
	ID      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Serials types.Set    `tfsdk:"serials"`
}

type RequestSwitchRemoveNetworkSwitchStackRs struct {
	Serial types.String `tfsdk:"serial"`
}

// FromBody
func (r *NetworksSwitchStacksRemove) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSwitchRemoveNetworkSwitchStack {
	emptyString := ""
	re := *r.Parameters
	serial := new(string)
	if !re.Serial.IsUnknown() && !re.Serial.IsNull() {
		*serial = re.Serial.ValueString()
	} else {
		serial = &emptyString
	}
	out := merakigosdk.RequestSwitchRemoveNetworkSwitchStack{
		Serial: *serial,
	}
	return &out
}

// ToBody
func ResponseSwitchRemoveNetworkSwitchStackItemToBody(state NetworksSwitchStacksRemove, response *merakigosdk.ResponseSwitchRemoveNetworkSwitchStack) NetworksSwitchStacksRemove {
	itemState := ResponseSwitchRemoveNetworkSwitchStack{
		ID:      types.StringValue(response.ID),
		Name:    types.StringValue(response.Name),
		Serials: StringSliceToSet(response.Serials),
	}
	state.Item = &itemState
	return state
}
