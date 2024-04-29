package provider

// RESOURCE NORMAL
import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceFirewallSettingsResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceFirewallSettingsResource{}
)

func NewNetworksApplianceFirewallSettingsResource() resource.Resource {
	return &NetworksApplianceFirewallSettingsResource{}
}

type NetworksApplianceFirewallSettingsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceFirewallSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceFirewallSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_firewall_settings"
}

func (r *NetworksApplianceFirewallSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"spoofing_protection": schema.SingleNestedAttribute{
				MarkdownDescription: `Spoofing protection settings`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"ip_source_guard": schema.SingleNestedAttribute{
						MarkdownDescription: `IP source address spoofing settings`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"mode": schema.StringAttribute{
								MarkdownDescription: `Mode of protection`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"block",
										"log",
									),
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *NetworksApplianceFirewallSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceFirewallSettingsRs

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
	responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceFirewallSettings(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceFirewallSettings only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceFirewallSettings only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceFirewallSettings(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceFirewallSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceFirewallSettings",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceFirewallSettings(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceFirewallSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceFirewallSettings",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceFirewallSettingsItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceFirewallSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceFirewallSettingsRs

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
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceFirewallSettings(vvNetworkID)
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
				"Failure when executing GetNetworkApplianceFirewallSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceFirewallSettings",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceFirewallSettingsItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksApplianceFirewallSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceFirewallSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceFirewallSettingsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceFirewallSettings(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceFirewallSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceFirewallSettings",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceFirewallSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceFirewallSettings", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceFirewallSettingsRs struct {
	NetworkID          types.String                                                              `tfsdk:"network_id"`
	SpoofingProtection *ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionRs `tfsdk:"spoofing_protection"`
}

type ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionRs struct {
	IPSourceGuard *ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionIpSourceGuardRs `tfsdk:"ip_source_guard"`
}

type ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionIpSourceGuardRs struct {
	Mode types.String `tfsdk:"mode"`
}

// FromBody
func (r *NetworksApplianceFirewallSettingsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallSettings {
	var requestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtection *merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtection
	if r.SpoofingProtection != nil {
		var requestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtectionIPSourceGuard *merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtectionIPSourceGuard
		if r.SpoofingProtection.IPSourceGuard != nil {
			mode := r.SpoofingProtection.IPSourceGuard.Mode.ValueString()
			requestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtectionIPSourceGuard = &merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtectionIPSourceGuard{
				Mode: mode,
			}
		}
		requestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtection = &merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtection{
			IPSourceGuard: requestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtectionIPSourceGuard,
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallSettings{
		SpoofingProtection: requestApplianceUpdateNetworkApplianceFirewallSettingsSpoofingProtection,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceFirewallSettingsItemToBodyRs(state NetworksApplianceFirewallSettingsRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceFirewallSettings, is_read bool) NetworksApplianceFirewallSettingsRs {
	itemState := NetworksApplianceFirewallSettingsRs{
		SpoofingProtection: func() *ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionRs {
			if response.SpoofingProtection != nil {
				return &ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionRs{
					IPSourceGuard: func() *ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionIpSourceGuardRs {
						if response.SpoofingProtection.IPSourceGuard != nil {
							return &ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionIpSourceGuardRs{
								Mode: types.StringValue(response.SpoofingProtection.IPSourceGuard.Mode),
							}
						}
						return &ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionIpSourceGuardRs{}
					}(),
				}
			}
			return &ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceFirewallSettingsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceFirewallSettingsRs)
}
