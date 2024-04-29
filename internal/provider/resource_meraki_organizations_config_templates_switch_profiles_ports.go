package provider

// RESOURCE NORMAL
import (
	"context"
	"fmt"
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
	_ resource.Resource              = &OrganizationsConfigTemplatesSwitchProfilesPortsResource{}
	_ resource.ResourceWithConfigure = &OrganizationsConfigTemplatesSwitchProfilesPortsResource{}
)

func NewOrganizationsConfigTemplatesSwitchProfilesPortsResource() resource.Resource {
	return &OrganizationsConfigTemplatesSwitchProfilesPortsResource{}
}

type OrganizationsConfigTemplatesSwitchProfilesPortsResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsConfigTemplatesSwitchProfilesPortsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsConfigTemplatesSwitchProfilesPortsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_config_templates_switch_profiles_ports"
}

func (r *OrganizationsConfigTemplatesSwitchProfilesPortsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"access_policy_number": schema.Int64Attribute{
				MarkdownDescription: `The number of a custom access policy to configure on the switch template port. Only applicable when 'accessPolicyType' is 'Custom access policy'.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"access_policy_type": schema.StringAttribute{
				MarkdownDescription: `The type of the access policy of the switch template port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Custom access policy",
						"MAC allow list",
						"Open",
						"Sticky MAC allow list",
					),
				},
			},
			"allowed_vlans": schema.StringAttribute{
				MarkdownDescription: `The VLANs allowed on the switch template port. Only applicable to trunk ports.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"config_template_id": schema.StringAttribute{
				MarkdownDescription: `configTemplateId path parameter. Config template ID`,
				Required:            true,
			},
			"dai_trusted": schema.BoolAttribute{
				MarkdownDescription: `If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `The status of the switch template port.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"flexible_stacking_enabled": schema.BoolAttribute{
				MarkdownDescription: `For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"isolation_enabled": schema.BoolAttribute{
				MarkdownDescription: `The isolation status of the switch template port.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"link_negotiation": schema.StringAttribute{
				MarkdownDescription: `The link speed for the switch template port.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"link_negotiation_capabilities": schema.SetAttribute{
				MarkdownDescription: `Available link speeds for the switch template port.`,
				Computed:            true,
				ElementType:         types.StringType,
			},
			"mac_allow_list": schema.SetAttribute{
				MarkdownDescription: `Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"mirror": schema.SingleNestedAttribute{
				MarkdownDescription: `Port mirror`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{

					"mode": schema.StringAttribute{
						MarkdownDescription: `The port mirror mode. Can be one of ('Destination port', 'Source port' or 'Not mirroring traffic').`,
						Computed:            true,
					},
				},
			},
			"module": schema.SingleNestedAttribute{
				MarkdownDescription: `Expansion module`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{

					"model": schema.StringAttribute{
						MarkdownDescription: `The model of the expansion module.`,
						Computed:            true,
					},
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name of the switch template port.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"poe_enabled": schema.BoolAttribute{
				MarkdownDescription: `The PoE status of the switch template port.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"port_id": schema.StringAttribute{
				MarkdownDescription: `The identifier of the switch template port.`,
				Required:            true,
			},
			"port_schedule_id": schema.StringAttribute{
				MarkdownDescription: `The ID of the port schedule. A value of null will clear the port schedule.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"profile": schema.SingleNestedAttribute{
				MarkdownDescription: `Profile attributes`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `When enabled, override this port's configuration with a port profile.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `When enabled, the ID of the port profile used to override the port's configuration.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"iname": schema.StringAttribute{
						MarkdownDescription: `When enabled, the IName of the profile.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"profile_id": schema.StringAttribute{
				MarkdownDescription: `profileId path parameter. Profile ID`,
				Required:            true,
			},
			"rstp_enabled": schema.BoolAttribute{
				MarkdownDescription: `The rapid spanning tree protocol status.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"sticky_mac_allow_list": schema.SetAttribute{
				MarkdownDescription: `The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"sticky_mac_allow_list_limit": schema.Int64Attribute{
				MarkdownDescription: `The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"storm_control_enabled": schema.BoolAttribute{
				MarkdownDescription: `The storm control status of the switch template port.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"stp_guard": schema.StringAttribute{
				MarkdownDescription: `The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"bpdu guard",
						"disabled",
						"loop guard",
						"root guard",
					),
				},
			},
			"tags": schema.SetAttribute{
				MarkdownDescription: `The list of tags of the switch template port.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: `The type of the switch template port ('trunk' or 'access').`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"access",
						"trunk",
					),
				},
			},
			"udld": schema.StringAttribute{
				MarkdownDescription: `The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Alert only",
						"Enforce",
					),
				},
			},
			"vlan": schema.Int64Attribute{
				MarkdownDescription: `The VLAN of the switch template port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"voice_vlan": schema.Int64Attribute{
				MarkdownDescription: `The voice VLAN of the switch template port. Only applicable to access ports.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

//path params to set ['portId']

func (r *OrganizationsConfigTemplatesSwitchProfilesPortsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsConfigTemplatesSwitchProfilesPortsRs

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
	vvOrganizationID := data.OrganizationID.ValueString()
	vvConfigTemplateID := data.ConfigTemplateID.ValueString()
	vvProfileID := data.ProfileID.ValueString()
	vvPortID := data.PortID.ValueString()
	//Item
	responseVerifyItem, restyResp1, err := r.client.Switch.GetOrganizationConfigTemplateSwitchProfilePort(vvOrganizationID, vvConfigTemplateID, vvProfileID, vvPortID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource OrganizationsConfigTemplatesSwitchProfilesPorts only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource OrganizationsConfigTemplatesSwitchProfilesPorts only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateOrganizationConfigTemplateSwitchProfilePort(vvOrganizationID, vvConfigTemplateID, vvProfileID, vvPortID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationConfigTemplateSwitchProfilePort",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationConfigTemplateSwitchProfilePort",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Switch.GetOrganizationConfigTemplateSwitchProfilePort(vvOrganizationID, vvConfigTemplateID, vvProfileID, vvPortID)
	// Has only items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationConfigTemplateSwitchProfilePorts",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationConfigTemplateSwitchProfilePorts",
			err.Error(),
		)
		return
	}
	data = ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsConfigTemplatesSwitchProfilesPortsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsConfigTemplatesSwitchProfilesPortsRs

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

	vvOrganizationID := data.OrganizationID.ValueString()
	vvConfigTemplateID := data.ConfigTemplateID.ValueString()
	vvProfileID := data.ProfileID.ValueString()
	vvPortID := data.PortID.ValueString()
	responseGet, restyRespGet, err := r.client.Switch.GetOrganizationConfigTemplateSwitchProfilePort(vvOrganizationID, vvConfigTemplateID, vvProfileID, vvPortID)
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
				"Failure when executing GetOrganizationConfigTemplateSwitchProfilePort",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationConfigTemplateSwitchProfilePort",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsConfigTemplatesSwitchProfilesPortsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 4 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" || idParts[3] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("config_template_id"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("profile_id"), idParts[2])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("port_id"), idParts[3])...)
}

func (r *OrganizationsConfigTemplatesSwitchProfilesPortsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsConfigTemplatesSwitchProfilesPortsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	vvConfigTemplateID := data.ConfigTemplateID.ValueString()
	vvProfileID := data.ProfileID.ValueString()
	vvPortID := data.PortID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateOrganizationConfigTemplateSwitchProfilePort(vvOrganizationID, vvConfigTemplateID, vvProfileID, vvPortID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationConfigTemplateSwitchProfilePort",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationConfigTemplateSwitchProfilePort",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsConfigTemplatesSwitchProfilesPortsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting OrganizationsConfigTemplatesSwitchProfilesPorts", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsConfigTemplatesSwitchProfilesPortsRs struct {
	OrganizationID              types.String                                                           `tfsdk:"organization_id"`
	ConfigTemplateID            types.String                                                           `tfsdk:"config_template_id"`
	ProfileID                   types.String                                                           `tfsdk:"profile_id"`
	PortID                      types.String                                                           `tfsdk:"port_id"`
	AccessPolicyNumber          types.Int64                                                            `tfsdk:"access_policy_number"`
	AccessPolicyType            types.String                                                           `tfsdk:"access_policy_type"`
	AllowedVLANs                types.String                                                           `tfsdk:"allowed_vlans"`
	DaiTrusted                  types.Bool                                                             `tfsdk:"dai_trusted"`
	Enabled                     types.Bool                                                             `tfsdk:"enabled"`
	FlexibleStackingEnabled     types.Bool                                                             `tfsdk:"flexible_stacking_enabled"`
	IsolationEnabled            types.Bool                                                             `tfsdk:"isolation_enabled"`
	LinkNegotiation             types.String                                                           `tfsdk:"link_negotiation"`
	LinkNegotiationCapabilities types.Set                                                              `tfsdk:"link_negotiation_capabilities"`
	MacAllowList                types.Set                                                              `tfsdk:"mac_allow_list"`
	Mirror                      *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortMirrorRs  `tfsdk:"mirror"`
	Module                      *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortModuleRs  `tfsdk:"module"`
	Name                        types.String                                                           `tfsdk:"name"`
	PoeEnabled                  types.Bool                                                             `tfsdk:"poe_enabled"`
	PortScheduleID              types.String                                                           `tfsdk:"port_schedule_id"`
	Profile                     *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortProfileRs `tfsdk:"profile"`
	RstpEnabled                 types.Bool                                                             `tfsdk:"rstp_enabled"`
	StickyMacAllowList          types.Set                                                              `tfsdk:"sticky_mac_allow_list"`
	StickyMacAllowListLimit     types.Int64                                                            `tfsdk:"sticky_mac_allow_list_limit"`
	StormControlEnabled         types.Bool                                                             `tfsdk:"storm_control_enabled"`
	StpGuard                    types.String                                                           `tfsdk:"stp_guard"`
	Tags                        types.Set                                                              `tfsdk:"tags"`
	Type                        types.String                                                           `tfsdk:"type"`
	Udld                        types.String                                                           `tfsdk:"udld"`
	VLAN                        types.Int64                                                            `tfsdk:"vlan"`
	VoiceVLAN                   types.Int64                                                            `tfsdk:"voice_vlan"`
}

type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortMirrorRs struct {
	Mode types.String `tfsdk:"mode"`
}

type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortModuleRs struct {
	Model types.String `tfsdk:"model"`
}

type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortProfileRs struct {
	Enabled types.Bool   `tfsdk:"enabled"`
	ID      types.String `tfsdk:"id"`
	Iname   types.String `tfsdk:"iname"`
}

// FromBody
func (r *OrganizationsConfigTemplatesSwitchProfilesPortsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateOrganizationConfigTemplateSwitchProfilePort {
	emptyString := ""
	accessPolicyNumber := new(int64)
	if !r.AccessPolicyNumber.IsUnknown() && !r.AccessPolicyNumber.IsNull() {
		*accessPolicyNumber = r.AccessPolicyNumber.ValueInt64()
	} else {
		accessPolicyNumber = nil
	}
	accessPolicyType := new(string)
	if !r.AccessPolicyType.IsUnknown() && !r.AccessPolicyType.IsNull() {
		*accessPolicyType = r.AccessPolicyType.ValueString()
	} else {
		accessPolicyType = &emptyString
	}
	allowedVLANs := new(string)
	if !r.AllowedVLANs.IsUnknown() && !r.AllowedVLANs.IsNull() {
		*allowedVLANs = r.AllowedVLANs.ValueString()
	} else {
		allowedVLANs = &emptyString
	}
	daiTrusted := new(bool)
	if !r.DaiTrusted.IsUnknown() && !r.DaiTrusted.IsNull() {
		*daiTrusted = r.DaiTrusted.ValueBool()
	} else {
		daiTrusted = nil
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	flexibleStackingEnabled := new(bool)
	if !r.FlexibleStackingEnabled.IsUnknown() && !r.FlexibleStackingEnabled.IsNull() {
		*flexibleStackingEnabled = r.FlexibleStackingEnabled.ValueBool()
	} else {
		flexibleStackingEnabled = nil
	}
	isolationEnabled := new(bool)
	if !r.IsolationEnabled.IsUnknown() && !r.IsolationEnabled.IsNull() {
		*isolationEnabled = r.IsolationEnabled.ValueBool()
	} else {
		isolationEnabled = nil
	}
	linkNegotiation := new(string)
	if !r.LinkNegotiation.IsUnknown() && !r.LinkNegotiation.IsNull() {
		*linkNegotiation = r.LinkNegotiation.ValueString()
	} else {
		linkNegotiation = &emptyString
	}
	var macAllowList []string = nil
	r.MacAllowList.ElementsAs(ctx, &macAllowList, false)
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	poeEnabled := new(bool)
	if !r.PoeEnabled.IsUnknown() && !r.PoeEnabled.IsNull() {
		*poeEnabled = r.PoeEnabled.ValueBool()
	} else {
		poeEnabled = nil
	}
	portScheduleID := new(string)
	if !r.PortScheduleID.IsUnknown() && !r.PortScheduleID.IsNull() {
		*portScheduleID = r.PortScheduleID.ValueString()
	} else {
		portScheduleID = &emptyString
	}
	var requestSwitchUpdateOrganizationConfigTemplateSwitchProfilePortProfile *merakigosdk.RequestSwitchUpdateOrganizationConfigTemplateSwitchProfilePortProfile
	if r.Profile != nil {
		enabled := func() *bool {
			if !r.Profile.Enabled.IsUnknown() && !r.Profile.Enabled.IsNull() {
				return r.Profile.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		iD := r.Profile.ID.ValueString()
		iname := r.Profile.Iname.ValueString()
		requestSwitchUpdateOrganizationConfigTemplateSwitchProfilePortProfile = &merakigosdk.RequestSwitchUpdateOrganizationConfigTemplateSwitchProfilePortProfile{
			Enabled: enabled,
			ID:      iD,
			Iname:   iname,
		}
	}
	rstpEnabled := new(bool)
	if !r.RstpEnabled.IsUnknown() && !r.RstpEnabled.IsNull() {
		*rstpEnabled = r.RstpEnabled.ValueBool()
	} else {
		rstpEnabled = nil
	}
	var stickyMacAllowList []string = nil
	r.StickyMacAllowList.ElementsAs(ctx, &stickyMacAllowList, false)
	stickyMacAllowListLimit := new(int64)
	if !r.StickyMacAllowListLimit.IsUnknown() && !r.StickyMacAllowListLimit.IsNull() {
		*stickyMacAllowListLimit = r.StickyMacAllowListLimit.ValueInt64()
	} else {
		stickyMacAllowListLimit = nil
	}
	stormControlEnabled := new(bool)
	if !r.StormControlEnabled.IsUnknown() && !r.StormControlEnabled.IsNull() {
		*stormControlEnabled = r.StormControlEnabled.ValueBool()
	} else {
		stormControlEnabled = nil
	}
	stpGuard := new(string)
	if !r.StpGuard.IsUnknown() && !r.StpGuard.IsNull() {
		*stpGuard = r.StpGuard.ValueString()
	} else {
		stpGuard = &emptyString
	}
	var tags []string = nil
	r.Tags.ElementsAs(ctx, &tags, false)
	typeR := new(string)
	if !r.Type.IsUnknown() && !r.Type.IsNull() {
		*typeR = r.Type.ValueString()
	} else {
		typeR = &emptyString
	}
	udld := new(string)
	if !r.Udld.IsUnknown() && !r.Udld.IsNull() {
		*udld = r.Udld.ValueString()
	} else {
		udld = &emptyString
	}
	vLAN := new(int64)
	if !r.VLAN.IsUnknown() && !r.VLAN.IsNull() {
		*vLAN = r.VLAN.ValueInt64()
	} else {
		vLAN = nil
	}
	voiceVLAN := new(int64)
	if !r.VoiceVLAN.IsUnknown() && !r.VoiceVLAN.IsNull() {
		*voiceVLAN = r.VoiceVLAN.ValueInt64()
	} else {
		voiceVLAN = nil
	}
	out := merakigosdk.RequestSwitchUpdateOrganizationConfigTemplateSwitchProfilePort{
		AccessPolicyNumber:      int64ToIntPointer(accessPolicyNumber),
		AccessPolicyType:        *accessPolicyType,
		AllowedVLANs:            *allowedVLANs,
		DaiTrusted:              daiTrusted,
		Enabled:                 enabled,
		FlexibleStackingEnabled: flexibleStackingEnabled,
		IsolationEnabled:        isolationEnabled,
		LinkNegotiation:         *linkNegotiation,
		MacAllowList:            macAllowList,
		Name:                    *name,
		PoeEnabled:              poeEnabled,
		PortScheduleID:          *portScheduleID,
		Profile:                 requestSwitchUpdateOrganizationConfigTemplateSwitchProfilePortProfile,
		RstpEnabled:             rstpEnabled,
		StickyMacAllowList:      stickyMacAllowList,
		StickyMacAllowListLimit: int64ToIntPointer(stickyMacAllowListLimit),
		StormControlEnabled:     stormControlEnabled,
		StpGuard:                *stpGuard,
		Tags:                    tags,
		Type:                    *typeR,
		Udld:                    *udld,
		VLAN:                    int64ToIntPointer(vLAN),
		VoiceVLAN:               int64ToIntPointer(voiceVLAN),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortItemToBodyRs(state OrganizationsConfigTemplatesSwitchProfilesPortsRs, response *merakigosdk.ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePort, is_read bool) OrganizationsConfigTemplatesSwitchProfilesPortsRs {
	itemState := OrganizationsConfigTemplatesSwitchProfilesPortsRs{
		AccessPolicyNumber: func() types.Int64 {
			if response.AccessPolicyNumber != nil {
				return types.Int64Value(int64(*response.AccessPolicyNumber))
			}
			return types.Int64{}
		}(),
		AccessPolicyType: types.StringValue(response.AccessPolicyType),
		AllowedVLANs:     types.StringValue(response.AllowedVLANs),
		DaiTrusted: func() types.Bool {
			if response.DaiTrusted != nil {
				return types.BoolValue(*response.DaiTrusted)
			}
			return types.Bool{}
		}(),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		FlexibleStackingEnabled: func() types.Bool {
			if response.FlexibleStackingEnabled != nil {
				return types.BoolValue(*response.FlexibleStackingEnabled)
			}
			return types.Bool{}
		}(),
		IsolationEnabled: func() types.Bool {
			if response.IsolationEnabled != nil {
				return types.BoolValue(*response.IsolationEnabled)
			}
			return types.Bool{}
		}(),
		LinkNegotiation:             types.StringValue(response.LinkNegotiation),
		LinkNegotiationCapabilities: StringSliceToSet(response.LinkNegotiationCapabilities),
		MacAllowList:                StringSliceToSet(response.MacAllowList),
		Mirror: func() *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortMirrorRs {
			if response.Mirror != nil {
				return &ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortMirrorRs{
					Mode: types.StringValue(response.Mirror.Mode),
				}
			}
			return &ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortMirrorRs{}
		}(),
		Module: func() *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortModuleRs {
			if response.Module != nil {
				return &ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortModuleRs{
					Model: types.StringValue(response.Module.Model),
				}
			}
			return &ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortModuleRs{}
		}(),
		Name: types.StringValue(response.Name),
		PoeEnabled: func() types.Bool {
			if response.PoeEnabled != nil {
				return types.BoolValue(*response.PoeEnabled)
			}
			return types.Bool{}
		}(),
		PortID:         types.StringValue(response.PortID),
		PortScheduleID: types.StringValue(response.PortScheduleID),
		Profile: func() *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortProfileRs {
			if response.Profile != nil {
				return &ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortProfileRs{
					Enabled: func() types.Bool {
						if response.Profile.Enabled != nil {
							return types.BoolValue(*response.Profile.Enabled)
						}
						return types.Bool{}
					}(),
					ID:    types.StringValue(response.Profile.ID),
					Iname: types.StringValue(response.Profile.Iname),
				}
			}
			return &ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortProfileRs{}
		}(),
		RstpEnabled: func() types.Bool {
			if response.RstpEnabled != nil {
				return types.BoolValue(*response.RstpEnabled)
			}
			return types.Bool{}
		}(),
		StickyMacAllowList: StringSliceToSet(response.StickyMacAllowList),
		StickyMacAllowListLimit: func() types.Int64 {
			if response.StickyMacAllowListLimit != nil {
				return types.Int64Value(int64(*response.StickyMacAllowListLimit))
			}
			return types.Int64{}
		}(),
		StormControlEnabled: func() types.Bool {
			if response.StormControlEnabled != nil {
				return types.BoolValue(*response.StormControlEnabled)
			}
			return types.Bool{}
		}(),
		StpGuard: types.StringValue(response.StpGuard),
		Tags:     StringSliceToSet(response.Tags),
		Type:     types.StringValue(response.Type),
		Udld:     types.StringValue(response.Udld),
		VLAN: func() types.Int64 {
			if response.VLAN != nil {
				return types.Int64Value(int64(*response.VLAN))
			}
			return types.Int64{}
		}(),
		VoiceVLAN: func() types.Int64 {
			if response.VoiceVLAN != nil {
				return types.Int64Value(int64(*response.VoiceVLAN))
			}
			return types.Int64{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsConfigTemplatesSwitchProfilesPortsRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsConfigTemplatesSwitchProfilesPortsRs)
}
