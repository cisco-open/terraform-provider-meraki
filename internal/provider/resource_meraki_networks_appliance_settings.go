package provider

// RESOURCE NORMAL
import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceSettingsResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceSettingsResource{}
)

func NewNetworksApplianceSettingsResource() resource.Resource {
	return &NetworksApplianceSettingsResource{}
}

type NetworksApplianceSettingsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_settings"
}

func (r *NetworksApplianceSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"client_tracking_method": schema.StringAttribute{
				MarkdownDescription: `Client tracking method of a network`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"IP address",
						"MAC address",
						"Unique client identifier",
					),
				},
			},
			"deployment_mode": schema.StringAttribute{
				MarkdownDescription: `Deployment mode of a network`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"passthrough",
						"routed",
					),
				},
			},
			"dynamic_dns": schema.SingleNestedAttribute{
				MarkdownDescription: `Dynamic DNS settings for a network`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Dynamic DNS enabled`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"prefix": schema.StringAttribute{
						MarkdownDescription: `Dynamic DNS url prefix. DDNS must be enabled to update`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"url": schema.StringAttribute{
						MarkdownDescription: `Dynamic DNS url. DDNS must be enabled to update`,
						Computed:            true,
					},
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
		},
	}
}

func (r *NetworksApplianceSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceSettingsRs

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
	//Item
	responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceSettings(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceSettings only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceSettings only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceSettings(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceSettings",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceSettings(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceSettings",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceSettingsItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceSettingsRs

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
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceSettings(vvNetworkID)
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
				"Failure when executing GetNetworkApplianceSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceSettings",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceSettingsItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksApplianceSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceSettingsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceSettings(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceSettings",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceSettings", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceSettingsRs struct {
	NetworkID            types.String                                              `tfsdk:"network_id"`
	ClientTrackingMethod types.String                                              `tfsdk:"client_tracking_method"`
	DeploymentMode       types.String                                              `tfsdk:"deployment_mode"`
	DynamicDNS           *ResponseApplianceGetNetworkApplianceSettingsDynamicDnsRs `tfsdk:"dynamic_dns"`
}

type ResponseApplianceGetNetworkApplianceSettingsDynamicDnsRs struct {
	Enabled types.Bool   `tfsdk:"enabled"`
	Prefix  types.String `tfsdk:"prefix"`
	URL     types.String `tfsdk:"url"`
}

// FromBody
func (r *NetworksApplianceSettingsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceSettings {
	emptyString := ""
	clientTrackingMethod := new(string)
	if !r.ClientTrackingMethod.IsUnknown() && !r.ClientTrackingMethod.IsNull() {
		*clientTrackingMethod = r.ClientTrackingMethod.ValueString()
	} else {
		clientTrackingMethod = &emptyString
	}
	deploymentMode := new(string)
	if !r.DeploymentMode.IsUnknown() && !r.DeploymentMode.IsNull() {
		*deploymentMode = r.DeploymentMode.ValueString()
	} else {
		deploymentMode = &emptyString
	}
	var requestApplianceUpdateNetworkApplianceSettingsDynamicDNS *merakigosdk.RequestApplianceUpdateNetworkApplianceSettingsDynamicDNS
	if r.DynamicDNS != nil {
		enabled := func() *bool {
			if !r.DynamicDNS.Enabled.IsUnknown() && !r.DynamicDNS.Enabled.IsNull() {
				return r.DynamicDNS.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		prefix := r.DynamicDNS.Prefix.ValueString()
		requestApplianceUpdateNetworkApplianceSettingsDynamicDNS = &merakigosdk.RequestApplianceUpdateNetworkApplianceSettingsDynamicDNS{
			Enabled: enabled,
			Prefix:  prefix,
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceSettings{
		ClientTrackingMethod: *clientTrackingMethod,
		DeploymentMode:       *deploymentMode,
		DynamicDNS:           requestApplianceUpdateNetworkApplianceSettingsDynamicDNS,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceSettingsItemToBodyRs(state NetworksApplianceSettingsRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceSettings, is_read bool) NetworksApplianceSettingsRs {
	itemState := NetworksApplianceSettingsRs{
		ClientTrackingMethod: types.StringValue(response.ClientTrackingMethod),
		DeploymentMode:       types.StringValue(response.DeploymentMode),
		DynamicDNS: func() *ResponseApplianceGetNetworkApplianceSettingsDynamicDnsRs {
			if response.DynamicDNS != nil {
				return &ResponseApplianceGetNetworkApplianceSettingsDynamicDnsRs{
					Enabled: func() types.Bool {
						if response.DynamicDNS.Enabled != nil {
							return types.BoolValue(*response.DynamicDNS.Enabled)
						}
						return types.Bool{}
					}(),
					Prefix: types.StringValue(response.DynamicDNS.Prefix),
					URL:    types.StringValue(response.DynamicDNS.URL),
				}
			}
			return &ResponseApplianceGetNetworkApplianceSettingsDynamicDnsRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceSettingsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceSettingsRs)
}
