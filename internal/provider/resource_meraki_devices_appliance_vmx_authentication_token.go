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
	_ resource.Resource              = &DevicesApplianceVmxAuthenticationTokenResource{}
	_ resource.ResourceWithConfigure = &DevicesApplianceVmxAuthenticationTokenResource{}
)

func NewDevicesApplianceVmxAuthenticationTokenResource() resource.Resource {
	return &DevicesApplianceVmxAuthenticationTokenResource{}
}

type DevicesApplianceVmxAuthenticationTokenResource struct {
	client *merakigosdk.Client
}

func (r *DevicesApplianceVmxAuthenticationTokenResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesApplianceVmxAuthenticationTokenResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_appliance_vmx_authentication_token"
}

// resourceAction
func (r *DevicesApplianceVmxAuthenticationTokenResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"expires_at": schema.StringAttribute{
						MarkdownDescription: `The expiration time for the token, in ISO 8601 format`,
						Computed:            true,
					},
					"token": schema.StringAttribute{
						MarkdownDescription: `The newly generated authentication token for the vMX instance`,
						Computed:            true,
					},
				},
			},
		},
	}
}
func (r *DevicesApplianceVmxAuthenticationTokenResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesApplianceVmxAuthenticationToken

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
	response, restyResp1, err := r.client.Appliance.CreateDeviceApplianceVmxAuthenticationToken(vvSerial)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateDeviceApplianceVmxAuthenticationToken",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateDeviceApplianceVmxAuthenticationToken",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseApplianceCreateDeviceApplianceVmxAuthenticationTokenItemToBody(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesApplianceVmxAuthenticationTokenResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *DevicesApplianceVmxAuthenticationTokenResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *DevicesApplianceVmxAuthenticationTokenResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesApplianceVmxAuthenticationToken struct {
	Serial types.String                                                  `tfsdk:"serial"`
	Item   *ResponseApplianceCreateDeviceApplianceVmxAuthenticationToken `tfsdk:"item"`
}

type ResponseApplianceCreateDeviceApplianceVmxAuthenticationToken struct {
	ExpiresAt types.String `tfsdk:"expires_at"`
	Token     types.String `tfsdk:"token"`
}

type RequestApplianceCreateDeviceApplianceVmxAuthenticationTokenRs interface{}

// FromBody
// ToBody
func ResponseApplianceCreateDeviceApplianceVmxAuthenticationTokenItemToBody(state DevicesApplianceVmxAuthenticationToken, response *merakigosdk.ResponseApplianceCreateDeviceApplianceVmxAuthenticationToken) DevicesApplianceVmxAuthenticationToken {
	itemState := ResponseApplianceCreateDeviceApplianceVmxAuthenticationToken{
		ExpiresAt: types.StringValue(response.ExpiresAt),
		Token:     types.StringValue(response.Token),
	}
	state.Item = &itemState
	return state
}
