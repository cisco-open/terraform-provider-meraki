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
	_ resource.Resource              = &NetworksSmDevicesUnenrollResource{}
	_ resource.ResourceWithConfigure = &NetworksSmDevicesUnenrollResource{}
)

func NewNetworksSmDevicesUnenrollResource() resource.Resource {
	return &NetworksSmDevicesUnenrollResource{}
}

type NetworksSmDevicesUnenrollResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSmDevicesUnenrollResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSmDevicesUnenrollResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_devices_unenroll"
}

// resourceAction
func (r *NetworksSmDevicesUnenrollResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"device_id": schema.StringAttribute{
				MarkdownDescription: `deviceId path parameter. Device ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"success": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating whether the operation was completed successfully.`,
						Computed:            true,
					},
				},
			},
		},
	}
}
func (r *NetworksSmDevicesUnenrollResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSmDevicesUnenroll

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
	vvDeviceID := data.DeviceID.ValueString()
	response, restyResp1, err := r.client.Sm.UnenrollNetworkSmDevice(vvNetworkID, vvDeviceID)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UnenrollNetworkSmDevice",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UnenrollNetworkSmDevice",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseSmUnenrollNetworkSmDeviceItemToBody(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSmDevicesUnenrollResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSmDevicesUnenrollResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSmDevicesUnenrollResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSmDevicesUnenroll struct {
	NetworkID  types.String                        `tfsdk:"network_id"`
	DeviceID   types.String                        `tfsdk:"device_id"`
	Item       *ResponseSmUnenrollNetworkSmDevice  `tfsdk:"item"`
	Parameters *RequestSmUnenrollNetworkSmDeviceRs `tfsdk:"parameters"`
}

type ResponseSmUnenrollNetworkSmDevice struct {
	Success types.Bool `tfsdk:"success"`
}

type RequestSmUnenrollNetworkSmDeviceRs interface{}

// FromBody
// ToBody
func ResponseSmUnenrollNetworkSmDeviceItemToBody(state NetworksSmDevicesUnenroll, response *merakigosdk.ResponseSmUnenrollNetworkSmDevice) NetworksSmDevicesUnenroll {
	itemState := ResponseSmUnenrollNetworkSmDevice{
		Success: func() types.Bool {
			if response.Success != nil {
				return types.BoolValue(*response.Success)
			}
			return types.Bool{}
		}(),
	}
	state.Item = &itemState
	return state
}
