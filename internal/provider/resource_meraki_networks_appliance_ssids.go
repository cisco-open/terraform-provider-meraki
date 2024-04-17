package provider

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"strconv"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceSSIDsResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceSSIDsResource{}
)

func NewNetworksApplianceSSIDsResource() resource.Resource {
	return &NetworksApplianceSSIDsResource{}
}

type NetworksApplianceSSIDsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceSSIDsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceSSIDsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_ssids"
}

func (r *NetworksApplianceSSIDsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"auth_mode": schema.StringAttribute{
				MarkdownDescription: `The association control method for the SSID.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"8021x-meraki",
						"8021x-radius",
						"open",
						"psk",
					),
				},
			},
			"default_vlan_id": schema.Int64Attribute{
				MarkdownDescription: `The VLAN ID of the VLAN associated to this SSID.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"dhcp_enforced_deauthentication": schema.SingleNestedAttribute{
				MarkdownDescription: `DHCP Enforced Deauthentication enables the disassociation of wireless clients in addition to Mandatory DHCP. This param is only valid on firmware versions >= MX 17.0 where the associated LAN has Mandatory DHCP Enabled `,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Enable DCHP Enforced Deauthentication on the SSID.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"dot11w": schema.SingleNestedAttribute{
				MarkdownDescription: `The current setting for Protected Management Frames (802.11w).`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether 802.11w is enabled or not.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"required": schema.BoolAttribute{
						MarkdownDescription: `(Optional) Whether 802.11w is required or not.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `Whether or not the SSID is enabled.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"encryption_mode": schema.StringAttribute{
				MarkdownDescription: `The psk encryption mode for the SSID.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"wep",
						"wpa",
					),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name of the SSID.`,
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
			"number": schema.Int64Attribute{
				MarkdownDescription: `The number of the SSID.`,
				Required:            true,
				//            Differents_types: `   parameter: schema.TypeString, item: schema.TypeInt`,
			},
			"psk": schema.StringAttribute{
				MarkdownDescription: `The passkey for the SSID. This param is only valid if the authMode is 'psk'.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"radius_servers": schema.SetNestedAttribute{
				MarkdownDescription: `The RADIUS 802.1x servers to be used for authentication.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"host": schema.StringAttribute{
							MarkdownDescription: `The IP address of your RADIUS server.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: `The UDP port your RADIUS servers listens on for Access-requests.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: `The RADIUS client shared secret.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"radius_servers_response": schema.SetNestedAttribute{
				MarkdownDescription: `The RADIUS 802.1x servers to be used for authentication.`,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"host": schema.StringAttribute{
							MarkdownDescription: `The IP address of your RADIUS server.`,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: `The UDP port your RADIUS servers listens on for Access-requests.`,
							Computed:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: `The RADIUS client shared secret.`,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"visible": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating whether the MX should advertise or hide this SSID.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"wpa_encryption_mode": schema.StringAttribute{
				MarkdownDescription: `WPA encryption mode for the SSID.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"WPA1 and WPA2",
						"WPA2 only",
						"WPA3 Transition Mode",
						"WPA3 only",
					),
				},
			},
		},
	}
}

//path params to set ['number']

func (r *NetworksApplianceSSIDsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceSSIDsRs

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
	a := int(data.Number.ValueInt64())
	vvNumber := strconv.Itoa(a)
	//Item
	responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceSSID(vvNetworkID, vvNumber)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceSSIDs only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceSSIDs only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceSSID(vvNetworkID, vvNumber, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceSSID",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceSSID",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceSSID(vvNetworkID, vvNumber)
	// Has only items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceSSIDs",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceSSIDs",
			err.Error(),
		)
		return
	}
	data = ResponseApplianceGetNetworkApplianceSSIDItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceSSIDsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceSSIDsRs

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
	a := int(data.Number.ValueInt64())
	vvNumber := strconv.Itoa(a)
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceSSID(vvNetworkID, vvNumber)
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
				"Failure when executing GetNetworkApplianceSSID",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceSSID",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceSSIDItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceSSIDsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("number"), idParts[1])...)
}

func (r *NetworksApplianceSSIDsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceSSIDsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	a := int(data.Number.ValueInt64())
	vvNumber := strconv.Itoa(a)
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceSSID(vvNetworkID, vvNumber, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceSSID",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceSSID",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceSSIDsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceSSIDs", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceSSIDsRs struct {
	NetworkID                    types.String                                                              `tfsdk:"network_id"`
	Number                       types.Int64                                                               `tfsdk:"number"`
	AuthMode                     types.String                                                              `tfsdk:"auth_mode"`
	DefaultVLANID                types.Int64                                                               `tfsdk:"default_vlan_id"`
	Enabled                      types.Bool                                                                `tfsdk:"enabled"`
	EncryptionMode               types.String                                                              `tfsdk:"encryption_mode"`
	Name                         types.String                                                              `tfsdk:"name"`
	RadiusServers                *[]ResponseApplianceGetNetworkApplianceSsidRadiusServersRs                `tfsdk:"radius_servers"`
	RadiusServersResponse        *[]ResponseApplianceGetNetworkApplianceSsidRadiusServersRs                `tfsdk:"radius_servers_response"`
	Visible                      types.Bool                                                                `tfsdk:"visible"`
	WpaEncryptionMode            types.String                                                              `tfsdk:"wpa_encryption_mode"`
	DhcpEnforcedDeauthentication *RequestApplianceUpdateNetworkApplianceSsidDhcpEnforcedDeauthenticationRs `tfsdk:"dhcp_enforced_deauthentication"`
	Dot11W                       *RequestApplianceUpdateNetworkApplianceSsidDot11WRs                       `tfsdk:"dot11w"`
	Psk                          types.String                                                              `tfsdk:"psk"`
}

type ResponseApplianceGetNetworkApplianceSsidRadiusServersRs struct {
	Host   types.String `tfsdk:"host"`
	Port   types.Int64  `tfsdk:"port"`
	Secret types.String `tfsdk:"secret"`
}

type RequestApplianceUpdateNetworkApplianceSsidDhcpEnforcedDeauthenticationRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type RequestApplianceUpdateNetworkApplianceSsidDot11WRs struct {
	Enabled  types.Bool `tfsdk:"enabled"`
	Required types.Bool `tfsdk:"required"`
}

// FromBody
func (r *NetworksApplianceSSIDsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceSSID {
	emptyString := ""
	authMode := new(string)
	if !r.AuthMode.IsUnknown() && !r.AuthMode.IsNull() {
		*authMode = r.AuthMode.ValueString()
	} else {
		authMode = &emptyString
	}
	defaultVLANID := new(int64)
	if !r.DefaultVLANID.IsUnknown() && !r.DefaultVLANID.IsNull() {
		*defaultVLANID = r.DefaultVLANID.ValueInt64()
	} else {
		defaultVLANID = nil
	}
	var requestApplianceUpdateNetworkApplianceSSIDDhcpEnforcedDeauthentication *merakigosdk.RequestApplianceUpdateNetworkApplianceSSIDDhcpEnforcedDeauthentication
	if r.DhcpEnforcedDeauthentication != nil {
		enabled := func() *bool {
			if !r.DhcpEnforcedDeauthentication.Enabled.IsUnknown() && !r.DhcpEnforcedDeauthentication.Enabled.IsNull() {
				return r.DhcpEnforcedDeauthentication.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		requestApplianceUpdateNetworkApplianceSSIDDhcpEnforcedDeauthentication = &merakigosdk.RequestApplianceUpdateNetworkApplianceSSIDDhcpEnforcedDeauthentication{
			Enabled: enabled,
		}
	}
	var requestApplianceUpdateNetworkApplianceSSIDDot11W *merakigosdk.RequestApplianceUpdateNetworkApplianceSSIDDot11W
	if r.Dot11W != nil {
		enabled := func() *bool {
			if !r.Dot11W.Enabled.IsUnknown() && !r.Dot11W.Enabled.IsNull() {
				return r.Dot11W.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		required := func() *bool {
			if !r.Dot11W.Required.IsUnknown() && !r.Dot11W.Required.IsNull() {
				return r.Dot11W.Required.ValueBoolPointer()
			}
			return nil
		}()
		requestApplianceUpdateNetworkApplianceSSIDDot11W = &merakigosdk.RequestApplianceUpdateNetworkApplianceSSIDDot11W{
			Enabled:  enabled,
			Required: required,
		}
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	encryptionMode := new(string)
	if !r.EncryptionMode.IsUnknown() && !r.EncryptionMode.IsNull() {
		*encryptionMode = r.EncryptionMode.ValueString()
	} else {
		encryptionMode = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	psk := new(string)
	if !r.Psk.IsUnknown() && !r.Psk.IsNull() {
		*psk = r.Psk.ValueString()
	} else {
		psk = &emptyString
	}
	var requestApplianceUpdateNetworkApplianceSSIDRadiusServers []merakigosdk.RequestApplianceUpdateNetworkApplianceSSIDRadiusServers
	if r.RadiusServers != nil {
		for _, rItem1 := range *r.RadiusServers {
			host := rItem1.Host.ValueString()
			port := func() *int64 {
				if !rItem1.Port.IsUnknown() && !rItem1.Port.IsNull() {
					return rItem1.Port.ValueInt64Pointer()
				}
				return nil
			}()
			secret := rItem1.Secret.ValueString()
			requestApplianceUpdateNetworkApplianceSSIDRadiusServers = append(requestApplianceUpdateNetworkApplianceSSIDRadiusServers, merakigosdk.RequestApplianceUpdateNetworkApplianceSSIDRadiusServers{
				Host:   host,
				Port:   int64ToIntPointer(port),
				Secret: secret,
			})
		}
	}
	visible := new(bool)
	if !r.Visible.IsUnknown() && !r.Visible.IsNull() {
		*visible = r.Visible.ValueBool()
	} else {
		visible = nil
	}
	wpaEncryptionMode := new(string)
	if !r.WpaEncryptionMode.IsUnknown() && !r.WpaEncryptionMode.IsNull() {
		*wpaEncryptionMode = r.WpaEncryptionMode.ValueString()
	} else {
		wpaEncryptionMode = &emptyString
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceSSID{
		AuthMode:                     *authMode,
		DefaultVLANID:                int64ToIntPointer(defaultVLANID),
		DhcpEnforcedDeauthentication: requestApplianceUpdateNetworkApplianceSSIDDhcpEnforcedDeauthentication,
		Dot11W:                       requestApplianceUpdateNetworkApplianceSSIDDot11W,
		Enabled:                      enabled,
		EncryptionMode:               *encryptionMode,
		Name:                         *name,
		Psk:                          *psk,
		RadiusServers: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceSSIDRadiusServers {
			if len(requestApplianceUpdateNetworkApplianceSSIDRadiusServers) > 0 {
				return &requestApplianceUpdateNetworkApplianceSSIDRadiusServers
			}
			return nil
		}(),
		Visible:           visible,
		WpaEncryptionMode: *wpaEncryptionMode,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceSSIDItemToBodyRs(state NetworksApplianceSSIDsRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceSSID, is_read bool) NetworksApplianceSSIDsRs {
	itemState := NetworksApplianceSSIDsRs{
		AuthMode: types.StringValue(response.AuthMode),
		DefaultVLANID: func() types.Int64 {
			if response.DefaultVLANID != nil {
				return types.Int64Value(int64(*response.DefaultVLANID))
			}
			return types.Int64{}
		}(),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		EncryptionMode: types.StringValue(response.EncryptionMode),
		Name:           types.StringValue(response.Name),
		Number: func() types.Int64 {
			if response.Number != nil {
				return types.Int64Value(int64(*response.Number))
			}
			return types.Int64{}
		}(),
		RadiusServersResponse: func() *[]ResponseApplianceGetNetworkApplianceSsidRadiusServersRs {
			if response.RadiusServers != nil {
				result := make([]ResponseApplianceGetNetworkApplianceSsidRadiusServersRs, len(*response.RadiusServers))
				for i, radiusServers := range *response.RadiusServers {
					result[i] = ResponseApplianceGetNetworkApplianceSsidRadiusServersRs{
						Host: types.StringValue(radiusServers.Host),
						Port: func() types.Int64 {
							if radiusServers.Port != nil {
								return types.Int64Value(int64(*radiusServers.Port))
							}
							return types.Int64{}
						}(),
					}
				}
				return &result
			}
			return &[]ResponseApplianceGetNetworkApplianceSsidRadiusServersRs{}
		}(),
		Visible: func() types.Bool {
			if response.Visible != nil {
				return types.BoolValue(*response.Visible)
			}
			return types.Bool{}
		}(),
		WpaEncryptionMode: types.StringValue(response.WpaEncryptionMode),
	}
	itemState.RadiusServers = state.RadiusServers
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceSSIDsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceSSIDsRs)
}
