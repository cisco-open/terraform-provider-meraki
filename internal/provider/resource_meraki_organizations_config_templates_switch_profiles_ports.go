// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0
package provider

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"strconv"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"access_policy_type": schema.StringAttribute{
				MarkdownDescription: `The type of the access policy of the switch template port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
                                  Allowed values: [Custom access policy,MAC allow list,Open,Sticky MAC allow list]`,
				Optional: true,
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
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"dot3az": schema.SingleNestedAttribute{
				MarkdownDescription: `dot3az settings for the port`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `The Energy Efficient Ethernet status of the switch template port.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `The status of the switch template port.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"flexible_stacking_enabled": schema.BoolAttribute{
				MarkdownDescription: `For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"isolation_enabled": schema.BoolAttribute{
				MarkdownDescription: `The isolation status of the switch template port.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"link_negotiation": schema.StringAttribute{
				MarkdownDescription: `The link speed for the switch template port.`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"link_negotiation_capabilities": schema.ListAttribute{
				MarkdownDescription: `Available link speeds for the switch template port.`,
				Computed:            true,
				ElementType:         types.StringType,
			},
			"mac_allow_list": schema.ListAttribute{
				MarkdownDescription: `Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.`,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
				Default:     listdefault.StaticValue(types.ListNull(types.StringType)),
			},
			"mirror": schema.SingleNestedAttribute{
				MarkdownDescription: `Port mirror`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{

					"mode": schema.StringAttribute{
						MarkdownDescription: `The port mirror mode. Can be one of ('Destination port', 'Source port' or 'Not mirroring traffic').
                                        Allowed values: [Destination port,Not mirroring traffic,Source port]`,
						Computed: true,
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
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"profile": schema.SingleNestedAttribute{
				MarkdownDescription: `Profile attributes`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `When enabled, override this port's configuration with a port profile.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `When enabled, the ID of the port profile used to override the port's configuration.`,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"iname": schema.StringAttribute{
						MarkdownDescription: `When enabled, the IName of the profile.`,
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
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"schedule": schema.SingleNestedAttribute{
				MarkdownDescription: `The port schedule data.`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{

					"id": schema.StringAttribute{
						MarkdownDescription: `The ID of the port schedule.`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name of the port schedule.`,
						Computed:            true,
					},
				},
			},
			"stackwise_virtual": schema.SingleNestedAttribute{
				MarkdownDescription: `Stackwise Virtual settings for the port`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{

					"is_dual_active_detector": schema.BoolAttribute{
						MarkdownDescription: `For SVL devices, whether or not the port is used for Dual Active Detection.`,
						Computed:            true,
					},
					"is_stack_wise_virtual_link": schema.BoolAttribute{
						MarkdownDescription: `For SVL devices, whether or not the port is used for StackWise Virtual Link.`,
						Computed:            true,
					},
				},
			},
			"sticky_mac_allow_list": schema.ListAttribute{
				MarkdownDescription: `The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.`,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
				Default:     listdefault.StaticValue(types.ListNull(types.StringType)),
			},
			"sticky_mac_allow_list_limit": schema.Int64Attribute{
				MarkdownDescription: `The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"storm_control_enabled": schema.BoolAttribute{
				MarkdownDescription: `The storm control status of the switch template port.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"stp_guard": schema.StringAttribute{
				MarkdownDescription: `The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
                                  Allowed values: [bpdu guard,disabled,loop guard,root guard]`,
				Optional: true,
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
			"tags": schema.ListAttribute{
				MarkdownDescription: `The list of tags of the switch template port.`,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
				Default:     listdefault.StaticValue(types.ListNull(types.StringType)),
			},
			"type": schema.StringAttribute{
				MarkdownDescription: `The type of the switch template port ('trunk', 'access', 'stack' or 'routed').
                                  Allowed values: [access,routed,stack,trunk]`,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"access",
						"routed",
						"stack",
						"trunk",
					),
				},
			},
			"udld": schema.StringAttribute{
				MarkdownDescription: `The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
                                  Allowed values: [Alert only,Enforce]`,
				Optional: true,
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
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"voice_vlan": schema.Int64Attribute{
				MarkdownDescription: `The voice VLAN of the switch template port. Only applicable to access ports.`,
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
	// Has Paths
	vvOrganizationID := data.OrganizationID.ValueString()
	vvConfigTemplateID := data.ConfigTemplateID.ValueString()
	vvProfileID := data.ProfileID.ValueString()
	vvPortID := data.PortID.ValueString()
	//Has Item and has items and not post

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateOrganizationConfigTemplateSwitchProfilePort(vvOrganizationID, vvConfigTemplateID, vvProfileID, vvPortID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationConfigTemplateSwitchProfilePort",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationConfigTemplateSwitchProfilePort",
			err.Error(),
		)
		return
	}

	// Assign data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *OrganizationsConfigTemplatesSwitchProfilesPortsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsConfigTemplatesSwitchProfilesPortsRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
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
				restyRespGet.String(),
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
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *OrganizationsConfigTemplatesSwitchProfilesPortsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 4 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" || idParts[3] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: organizationId,configTemplateId,profileId,portId. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("config_template_id"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("profile_id"), idParts[2])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("port_id"), idParts[3])...)
}

func (r *OrganizationsConfigTemplatesSwitchProfilesPortsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan OrganizationsConfigTemplatesSwitchProfilesPortsRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvOrganizationID := plan.OrganizationID.ValueString()
	vvConfigTemplateID := plan.ConfigTemplateID.ValueString()
	vvProfileID := plan.ProfileID.ValueString()
	vvPortID := plan.PortID.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateOrganizationConfigTemplateSwitchProfilePort(vvOrganizationID, vvConfigTemplateID, vvProfileID, vvPortID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationConfigTemplateSwitchProfilePort",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationConfigTemplateSwitchProfilePort",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *OrganizationsConfigTemplatesSwitchProfilesPortsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting OrganizationsConfigTemplatesSwitchProfilesPorts", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsConfigTemplatesSwitchProfilesPortsRs struct {
	OrganizationID              types.String                                                                    `tfsdk:"organization_id"`
	ConfigTemplateID            types.String                                                                    `tfsdk:"config_template_id"`
	ProfileID                   types.String                                                                    `tfsdk:"profile_id"`
	PortID                      types.String                                                                    `tfsdk:"port_id"`
	AccessPolicyNumber          types.Int64                                                                     `tfsdk:"access_policy_number"`
	AccessPolicyType            types.String                                                                    `tfsdk:"access_policy_type"`
	AllowedVLANs                types.String                                                                    `tfsdk:"allowed_vlans"`
	DaiTrusted                  types.Bool                                                                      `tfsdk:"dai_trusted"`
	Dot3Az                      *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortDot3AzRs           `tfsdk:"dot3az"`
	Enabled                     types.Bool                                                                      `tfsdk:"enabled"`
	FlexibleStackingEnabled     types.Bool                                                                      `tfsdk:"flexible_stacking_enabled"`
	IsolationEnabled            types.Bool                                                                      `tfsdk:"isolation_enabled"`
	LinkNegotiation             types.String                                                                    `tfsdk:"link_negotiation"`
	LinkNegotiationCapabilities types.List                                                                      `tfsdk:"link_negotiation_capabilities"`
	MacAllowList                types.List                                                                      `tfsdk:"mac_allow_list"`
	Mirror                      *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortMirrorRs           `tfsdk:"mirror"`
	Module                      *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortModuleRs           `tfsdk:"module"`
	Name                        types.String                                                                    `tfsdk:"name"`
	PoeEnabled                  types.Bool                                                                      `tfsdk:"poe_enabled"`
	PortScheduleID              types.String                                                                    `tfsdk:"port_schedule_id"`
	Profile                     *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortProfileRs          `tfsdk:"profile"`
	RstpEnabled                 types.Bool                                                                      `tfsdk:"rstp_enabled"`
	Schedule                    *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortScheduleRs         `tfsdk:"schedule"`
	StackwiseVirtual            *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortStackwiseVirtualRs `tfsdk:"stackwise_virtual"`
	StickyMacAllowList          types.List                                                                      `tfsdk:"sticky_mac_allow_list"`
	StickyMacAllowListLimit     types.Int64                                                                     `tfsdk:"sticky_mac_allow_list_limit"`
	StormControlEnabled         types.Bool                                                                      `tfsdk:"storm_control_enabled"`
	StpGuard                    types.String                                                                    `tfsdk:"stp_guard"`
	Tags                        types.List                                                                      `tfsdk:"tags"`
	Type                        types.String                                                                    `tfsdk:"type"`
	Udld                        types.String                                                                    `tfsdk:"udld"`
	VLAN                        types.Int64                                                                     `tfsdk:"vlan"`
	VoiceVLAN                   types.Int64                                                                     `tfsdk:"voice_vlan"`
}

type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortDot3AzRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
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

type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortScheduleRs struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortStackwiseVirtualRs struct {
	IsDualActiveDetector   types.Bool `tfsdk:"is_dual_active_detector"`
	IsStackWiseVirtualLink types.Bool `tfsdk:"is_stack_wise_virtual_link"`
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
	var requestSwitchUpdateOrganizationConfigTemplateSwitchProfilePortDot3Az *merakigosdk.RequestSwitchUpdateOrganizationConfigTemplateSwitchProfilePortDot3Az

	if r.Dot3Az != nil {
		enabled := func() *bool {
			if !r.Dot3Az.Enabled.IsUnknown() && !r.Dot3Az.Enabled.IsNull() {
				return r.Dot3Az.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		requestSwitchUpdateOrganizationConfigTemplateSwitchProfilePortDot3Az = &merakigosdk.RequestSwitchUpdateOrganizationConfigTemplateSwitchProfilePortDot3Az{
			Enabled: enabled,
		}
		//[debug] Is Array: False
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
		id := r.Profile.ID.ValueString()
		iname := r.Profile.Iname.ValueString()
		requestSwitchUpdateOrganizationConfigTemplateSwitchProfilePortProfile = &merakigosdk.RequestSwitchUpdateOrganizationConfigTemplateSwitchProfilePortProfile{
			Enabled: enabled,
			ID:      id,
			Iname:   iname,
		}
		//[debug] Is Array: False
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
		Dot3Az:                  requestSwitchUpdateOrganizationConfigTemplateSwitchProfilePortDot3Az,
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
		AccessPolicyType: func() types.String {
			if response.AccessPolicyType != "" {
				return types.StringValue(response.AccessPolicyType)
			}
			return types.String{}
		}(),
		AllowedVLANs: func() types.String {
			if response.AllowedVLANs != "" {
				return types.StringValue(response.AllowedVLANs)
			}
			return types.String{}
		}(),
		DaiTrusted: func() types.Bool {
			if response.DaiTrusted != nil {
				return types.BoolValue(*response.DaiTrusted)
			}
			return types.Bool{}
		}(),
		Dot3Az: func() *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortDot3AzRs {
			if response.Dot3Az != nil {
				return &ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortDot3AzRs{
					Enabled: func() types.Bool {
						if response.Dot3Az.Enabled != nil {
							return types.BoolValue(*response.Dot3Az.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
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
		LinkNegotiation: func() types.String {
			if response.LinkNegotiation != "" {
				return types.StringValue(response.LinkNegotiation)
			}
			return types.String{}
		}(),
		LinkNegotiationCapabilities: StringSliceToList(response.LinkNegotiationCapabilities),
		MacAllowList:                StringSliceToList(response.MacAllowList),
		Mirror: func() *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortMirrorRs {
			if response.Mirror != nil {
				return &ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortMirrorRs{
					Mode: func() types.String {
						if response.Mirror.Mode != "" {
							return types.StringValue(response.Mirror.Mode)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
		Module: func() *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortModuleRs {
			if response.Module != nil {
				return &ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortModuleRs{
					Model: func() types.String {
						if response.Module.Model != "" {
							return types.StringValue(response.Module.Model)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
		Name: func() types.String {
			if response.Name != "" {
				return types.StringValue(response.Name)
			}
			return types.String{}
		}(),
		PoeEnabled: func() types.Bool {
			if response.PoeEnabled != nil {
				return types.BoolValue(*response.PoeEnabled)
			}
			return types.Bool{}
		}(),
		PortID: func() types.String {
			if response.PortID != "" {
				return types.StringValue(response.PortID)
			}
			return types.String{}
		}(),
		PortScheduleID: func() types.String {
			if response.PortScheduleID != "" {
				return types.StringValue(response.PortScheduleID)
			}
			return types.String{}
		}(),
		Profile: func() *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortProfileRs {
			if response.Profile != nil {
				return &ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortProfileRs{
					Enabled: func() types.Bool {
						if response.Profile.Enabled != nil {
							return types.BoolValue(*response.Profile.Enabled)
						}
						return types.Bool{}
					}(),
					ID: func() types.String {
						if response.Profile.ID != "" {
							return types.StringValue(response.Profile.ID)
						}
						return types.String{}
					}(),
					Iname: func() types.String {
						if response.Profile.Iname != "" {
							return types.StringValue(response.Profile.Iname)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
		RstpEnabled: func() types.Bool {
			if response.RstpEnabled != nil {
				return types.BoolValue(*response.RstpEnabled)
			}
			return types.Bool{}
		}(),
		Schedule: func() *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortScheduleRs {
			if response.Schedule != nil {
				return &ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortScheduleRs{
					ID: func() types.String {
						if response.Schedule.ID != "" {
							return types.StringValue(response.Schedule.ID)
						}
						return types.String{}
					}(),
					Name: func() types.String {
						if response.Schedule.Name != "" {
							return types.StringValue(response.Schedule.Name)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
		StackwiseVirtual: func() *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortStackwiseVirtualRs {
			if response.StackwiseVirtual != nil {
				return &ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortStackwiseVirtualRs{
					IsDualActiveDetector: func() types.Bool {
						if response.StackwiseVirtual.IsDualActiveDetector != nil {
							return types.BoolValue(*response.StackwiseVirtual.IsDualActiveDetector)
						}
						return types.Bool{}
					}(),
					IsStackWiseVirtualLink: func() types.Bool {
						if response.StackwiseVirtual.IsStackWiseVirtualLink != nil {
							return types.BoolValue(*response.StackwiseVirtual.IsStackWiseVirtualLink)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
		}(),
		StickyMacAllowList: StringSliceToList(response.StickyMacAllowList),
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
		StpGuard: func() types.String {
			if response.StpGuard != "" {
				return types.StringValue(response.StpGuard)
			}
			return types.String{}
		}(),
		Tags: StringSliceToList(response.Tags),
		Type: func() types.String {
			if response.Type != "" {
				return types.StringValue(response.Type)
			}
			return types.String{}
		}(),
		Udld: func() types.String {
			if response.Udld != "" {
				return types.StringValue(response.Udld)
			}
			return types.String{}
		}(),
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
