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
	_ resource.Resource              = &NetworksSmUserAccessDevicesDeleteResource{}
	_ resource.ResourceWithConfigure = &NetworksSmUserAccessDevicesDeleteResource{}
)

func NewNetworksSmUserAccessDevicesDeleteResource() resource.Resource {
	return &NetworksSmUserAccessDevicesDeleteResource{}
}

type NetworksSmUserAccessDevicesDeleteResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSmUserAccessDevicesDeleteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSmUserAccessDevicesDeleteResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_user_access_devices_delete"
}

// resourceAction
func (r *NetworksSmUserAccessDevicesDeleteResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"user_access_device_id": schema.StringAttribute{
				MarkdownDescription: `userAccessDeviceId path parameter. User access device ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}
func (r *NetworksSmUserAccessDevicesDeleteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSmUserAccessDevicesDelete

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
	vvUserAccessDeviceID := data.UserAccessDeviceID.ValueString()
	restyResp1, err := r.client.Sm.DeleteNetworkSmUserAccessDevice(vvNetworkID, vvUserAccessDeviceID)

	if err != nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing DeleteNetworkSmUserAccessDevice",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkSmUserAccessDevice",
			err.Error(),
		)
		return
	}
	//Item
	// //entro aqui 2
	// data = ResponseSmDeleteNetworkSmUserAccessDevice(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSmUserAccessDevicesDeleteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSmUserAccessDevicesDeleteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSmUserAccessDevicesDeleteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSmUserAccessDevicesDelete struct {
	NetworkID          types.String `tfsdk:"network_id"`
	UserAccessDeviceID types.String `tfsdk:"user_access_device_id"`
	// Parameters         *Rs          `tfsdk:"parameters"`
}

//FromBody
