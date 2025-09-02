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

// DATA SOURCE NORMAL
import (
	"context"
	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsConfigTemplatesSwitchProfilesPortsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsConfigTemplatesSwitchProfilesPortsDataSource{}
)

func NewOrganizationsConfigTemplatesSwitchProfilesPortsDataSource() datasource.DataSource {
	return &OrganizationsConfigTemplatesSwitchProfilesPortsDataSource{}
}

type OrganizationsConfigTemplatesSwitchProfilesPortsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsConfigTemplatesSwitchProfilesPortsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsConfigTemplatesSwitchProfilesPortsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_config_templates_switch_profiles_ports"
}

func (d *OrganizationsConfigTemplatesSwitchProfilesPortsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"config_template_id": schema.StringAttribute{
				MarkdownDescription: `configTemplateId path parameter. Config template ID`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"port_id": schema.StringAttribute{
				MarkdownDescription: `portId path parameter. Port ID`,
				Optional:            true,
			},
			"profile_id": schema.StringAttribute{
				MarkdownDescription: `profileId path parameter. Profile ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"access_policy_number": schema.Int64Attribute{
						MarkdownDescription: `The number of a custom access policy to configure on the switch template port. Only applicable when 'accessPolicyType' is 'Custom access policy'.`,
						Computed:            true,
					},
					"access_policy_type": schema.StringAttribute{
						MarkdownDescription: `The type of the access policy of the switch template port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.`,
						Computed:            true,
					},
					"allowed_vlans": schema.StringAttribute{
						MarkdownDescription: `The VLANs allowed on the switch template port. Only applicable to trunk ports.`,
						Computed:            true,
					},
					"dai_trusted": schema.BoolAttribute{
						MarkdownDescription: `If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.`,
						Computed:            true,
					},
					"dot3az": schema.SingleNestedAttribute{
						MarkdownDescription: `dot3az settings for the port`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `The Energy Efficient Ethernet status of the switch template port.`,
								Computed:            true,
							},
						},
					},
					"enabled": schema.BoolAttribute{
						MarkdownDescription: `The status of the switch template port.`,
						Computed:            true,
					},
					"flexible_stacking_enabled": schema.BoolAttribute{
						MarkdownDescription: `For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.`,
						Computed:            true,
					},
					"isolation_enabled": schema.BoolAttribute{
						MarkdownDescription: `The isolation status of the switch template port.`,
						Computed:            true,
					},
					"link_negotiation": schema.StringAttribute{
						MarkdownDescription: `The link speed for the switch template port.`,
						Computed:            true,
					},
					"link_negotiation_capabilities": schema.ListAttribute{
						MarkdownDescription: `Available link speeds for the switch template port.`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"mac_allow_list": schema.ListAttribute{
						MarkdownDescription: `Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.`,
						Computed:            true,
						ElementType:         types.StringType,
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
					},
					"poe_enabled": schema.BoolAttribute{
						MarkdownDescription: `The PoE status of the switch template port.`,
						Computed:            true,
					},
					"port_id": schema.StringAttribute{
						MarkdownDescription: `The identifier of the switch template port.`,
						Computed:            true,
					},
					"port_schedule_id": schema.StringAttribute{
						MarkdownDescription: `The ID of the port schedule. A value of null will clear the port schedule.`,
						Computed:            true,
					},
					"profile": schema.SingleNestedAttribute{
						MarkdownDescription: `Profile attributes`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `When enabled, override this port's configuration with a port profile.`,
								Computed:            true,
							},
							"id": schema.StringAttribute{
								MarkdownDescription: `When enabled, the ID of the port profile used to override the port's configuration.`,
								Computed:            true,
							},
							"iname": schema.StringAttribute{
								MarkdownDescription: `When enabled, the IName of the profile.`,
								Computed:            true,
							},
						},
					},
					"rstp_enabled": schema.BoolAttribute{
						MarkdownDescription: `The rapid spanning tree protocol status.`,
						Computed:            true,
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
						Computed:            true,
						ElementType:         types.StringType,
					},
					"sticky_mac_allow_list_limit": schema.Int64Attribute{
						MarkdownDescription: `The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.`,
						Computed:            true,
					},
					"storm_control_enabled": schema.BoolAttribute{
						MarkdownDescription: `The storm control status of the switch template port.`,
						Computed:            true,
					},
					"stp_guard": schema.StringAttribute{
						MarkdownDescription: `The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').`,
						Computed:            true,
					},
					"tags": schema.ListAttribute{
						MarkdownDescription: `The list of tags of the switch template port.`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"type": schema.StringAttribute{
						MarkdownDescription: `The type of the switch template port ('trunk', 'access', 'stack' or 'routed').`,
						Computed:            true,
					},
					"udld": schema.StringAttribute{
						MarkdownDescription: `The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.`,
						Computed:            true,
					},
					"vlan": schema.Int64Attribute{
						MarkdownDescription: `The VLAN of the switch template port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.`,
						Computed:            true,
					},
					"voice_vlan": schema.Int64Attribute{
						MarkdownDescription: `The voice VLAN of the switch template port. Only applicable to access ports.`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePorts`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"access_policy_number": schema.Int64Attribute{
							MarkdownDescription: `The number of a custom access policy to configure on the switch template port. Only applicable when 'accessPolicyType' is 'Custom access policy'.`,
							Computed:            true,
						},
						"access_policy_type": schema.StringAttribute{
							MarkdownDescription: `The type of the access policy of the switch template port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.`,
							Computed:            true,
						},
						"allowed_vlans": schema.StringAttribute{
							MarkdownDescription: `The VLANs allowed on the switch template port. Only applicable to trunk ports.`,
							Computed:            true,
						},
						"dai_trusted": schema.BoolAttribute{
							MarkdownDescription: `If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.`,
							Computed:            true,
						},
						"dot3az": schema.SingleNestedAttribute{
							MarkdownDescription: `dot3az settings for the port`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"enabled": schema.BoolAttribute{
									MarkdownDescription: `The Energy Efficient Ethernet status of the switch template port.`,
									Computed:            true,
								},
							},
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: `The status of the switch template port.`,
							Computed:            true,
						},
						"flexible_stacking_enabled": schema.BoolAttribute{
							MarkdownDescription: `For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.`,
							Computed:            true,
						},
						"isolation_enabled": schema.BoolAttribute{
							MarkdownDescription: `The isolation status of the switch template port.`,
							Computed:            true,
						},
						"link_negotiation": schema.StringAttribute{
							MarkdownDescription: `The link speed for the switch template port.`,
							Computed:            true,
						},
						"link_negotiation_capabilities": schema.ListAttribute{
							MarkdownDescription: `Available link speeds for the switch template port.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"mac_allow_list": schema.ListAttribute{
							MarkdownDescription: `Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.`,
							Computed:            true,
							ElementType:         types.StringType,
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
						},
						"poe_enabled": schema.BoolAttribute{
							MarkdownDescription: `The PoE status of the switch template port.`,
							Computed:            true,
						},
						"port_id": schema.StringAttribute{
							MarkdownDescription: `The identifier of the switch template port.`,
							Computed:            true,
						},
						"port_schedule_id": schema.StringAttribute{
							MarkdownDescription: `The ID of the port schedule. A value of null will clear the port schedule.`,
							Computed:            true,
						},
						"profile": schema.SingleNestedAttribute{
							MarkdownDescription: `Profile attributes`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"enabled": schema.BoolAttribute{
									MarkdownDescription: `When enabled, override this port's configuration with a port profile.`,
									Computed:            true,
								},
								"id": schema.StringAttribute{
									MarkdownDescription: `When enabled, the ID of the port profile used to override the port's configuration.`,
									Computed:            true,
								},
								"iname": schema.StringAttribute{
									MarkdownDescription: `When enabled, the IName of the profile.`,
									Computed:            true,
								},
							},
						},
						"rstp_enabled": schema.BoolAttribute{
							MarkdownDescription: `The rapid spanning tree protocol status.`,
							Computed:            true,
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
							Computed:            true,
							ElementType:         types.StringType,
						},
						"sticky_mac_allow_list_limit": schema.Int64Attribute{
							MarkdownDescription: `The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.`,
							Computed:            true,
						},
						"storm_control_enabled": schema.BoolAttribute{
							MarkdownDescription: `The storm control status of the switch template port.`,
							Computed:            true,
						},
						"stp_guard": schema.StringAttribute{
							MarkdownDescription: `The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').`,
							Computed:            true,
						},
						"tags": schema.ListAttribute{
							MarkdownDescription: `The list of tags of the switch template port.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `The type of the switch template port ('trunk', 'access', 'stack' or 'routed').`,
							Computed:            true,
						},
						"udld": schema.StringAttribute{
							MarkdownDescription: `The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.`,
							Computed:            true,
						},
						"vlan": schema.Int64Attribute{
							MarkdownDescription: `The VLAN of the switch template port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.`,
							Computed:            true,
						},
						"voice_vlan": schema.Int64Attribute{
							MarkdownDescription: `The voice VLAN of the switch template port. Only applicable to access ports.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsConfigTemplatesSwitchProfilesPortsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsConfigTemplatesSwitchProfilesPorts OrganizationsConfigTemplatesSwitchProfilesPorts
	diags := req.Config.Get(ctx, &organizationsConfigTemplatesSwitchProfilesPorts)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsConfigTemplatesSwitchProfilesPorts.OrganizationID.IsNull(), !organizationsConfigTemplatesSwitchProfilesPorts.ConfigTemplateID.IsNull(), !organizationsConfigTemplatesSwitchProfilesPorts.ProfileID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsConfigTemplatesSwitchProfilesPorts.OrganizationID.IsNull(), !organizationsConfigTemplatesSwitchProfilesPorts.ConfigTemplateID.IsNull(), !organizationsConfigTemplatesSwitchProfilesPorts.ProfileID.IsNull(), !organizationsConfigTemplatesSwitchProfilesPorts.PortID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationConfigTemplateSwitchProfilePorts")
		vvOrganizationID := organizationsConfigTemplatesSwitchProfilesPorts.OrganizationID.ValueString()
		vvConfigTemplateID := organizationsConfigTemplatesSwitchProfilesPorts.ConfigTemplateID.ValueString()
		vvProfileID := organizationsConfigTemplatesSwitchProfilesPorts.ProfileID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetOrganizationConfigTemplateSwitchProfilePorts(vvOrganizationID, vvConfigTemplateID, vvProfileID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationConfigTemplateSwitchProfilePorts",
				err.Error(),
			)
			return
		}

		organizationsConfigTemplatesSwitchProfilesPorts = ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortsItemsToBody(organizationsConfigTemplatesSwitchProfilesPorts, response1)
		diags = resp.State.Set(ctx, &organizationsConfigTemplatesSwitchProfilesPorts)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationConfigTemplateSwitchProfilePort")
		vvOrganizationID := organizationsConfigTemplatesSwitchProfilesPorts.OrganizationID.ValueString()
		vvConfigTemplateID := organizationsConfigTemplatesSwitchProfilesPorts.ConfigTemplateID.ValueString()
		vvProfileID := organizationsConfigTemplatesSwitchProfilesPorts.ProfileID.ValueString()
		vvPortID := organizationsConfigTemplatesSwitchProfilesPorts.PortID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Switch.GetOrganizationConfigTemplateSwitchProfilePort(vvOrganizationID, vvConfigTemplateID, vvProfileID, vvPortID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationConfigTemplateSwitchProfilePort",
				err.Error(),
			)
			return
		}

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationConfigTemplateSwitchProfilePort",
				err.Error(),
			)
			return
		}

		organizationsConfigTemplatesSwitchProfilesPorts = ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortItemToBody(organizationsConfigTemplatesSwitchProfilesPorts, response2)
		diags = resp.State.Set(ctx, &organizationsConfigTemplatesSwitchProfilesPorts)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsConfigTemplatesSwitchProfilesPorts struct {
	OrganizationID   types.String                                                         `tfsdk:"organization_id"`
	ConfigTemplateID types.String                                                         `tfsdk:"config_template_id"`
	ProfileID        types.String                                                         `tfsdk:"profile_id"`
	PortID           types.String                                                         `tfsdk:"port_id"`
	Items            *[]ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePorts `tfsdk:"items"`
	Item             *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePort        `tfsdk:"item"`
}

type ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePorts struct {
	AccessPolicyNumber          types.Int64                                                                        `tfsdk:"access_policy_number"`
	AccessPolicyType            types.String                                                                       `tfsdk:"access_policy_type"`
	AllowedVLANs                types.String                                                                       `tfsdk:"allowed_vlans"`
	DaiTrusted                  types.Bool                                                                         `tfsdk:"dai_trusted"`
	Dot3Az                      *ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsDot3Az           `tfsdk:"dot3az"`
	Enabled                     types.Bool                                                                         `tfsdk:"enabled"`
	FlexibleStackingEnabled     types.Bool                                                                         `tfsdk:"flexible_stacking_enabled"`
	IsolationEnabled            types.Bool                                                                         `tfsdk:"isolation_enabled"`
	LinkNegotiation             types.String                                                                       `tfsdk:"link_negotiation"`
	LinkNegotiationCapabilities types.List                                                                         `tfsdk:"link_negotiation_capabilities"`
	MacAllowList                types.List                                                                         `tfsdk:"mac_allow_list"`
	Mirror                      *ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsMirror           `tfsdk:"mirror"`
	Module                      *ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsModule           `tfsdk:"module"`
	Name                        types.String                                                                       `tfsdk:"name"`
	PoeEnabled                  types.Bool                                                                         `tfsdk:"poe_enabled"`
	PortID                      types.String                                                                       `tfsdk:"port_id"`
	PortScheduleID              types.String                                                                       `tfsdk:"port_schedule_id"`
	Profile                     *ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsProfile          `tfsdk:"profile"`
	RstpEnabled                 types.Bool                                                                         `tfsdk:"rstp_enabled"`
	Schedule                    *ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsSchedule         `tfsdk:"schedule"`
	StackwiseVirtual            *ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsStackwiseVirtual `tfsdk:"stackwise_virtual"`
	StickyMacAllowList          types.List                                                                         `tfsdk:"sticky_mac_allow_list"`
	StickyMacAllowListLimit     types.Int64                                                                        `tfsdk:"sticky_mac_allow_list_limit"`
	StormControlEnabled         types.Bool                                                                         `tfsdk:"storm_control_enabled"`
	StpGuard                    types.String                                                                       `tfsdk:"stp_guard"`
	Tags                        types.List                                                                         `tfsdk:"tags"`
	Type                        types.String                                                                       `tfsdk:"type"`
	Udld                        types.String                                                                       `tfsdk:"udld"`
	VLAN                        types.Int64                                                                        `tfsdk:"vlan"`
	VoiceVLAN                   types.Int64                                                                        `tfsdk:"voice_vlan"`
}

type ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsDot3Az struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsMirror struct {
	Mode types.String `tfsdk:"mode"`
}

type ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsModule struct {
	Model types.String `tfsdk:"model"`
}

type ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsProfile struct {
	Enabled types.Bool   `tfsdk:"enabled"`
	ID      types.String `tfsdk:"id"`
	Iname   types.String `tfsdk:"iname"`
}

type ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsSchedule struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsStackwiseVirtual struct {
	IsDualActiveDetector   types.Bool `tfsdk:"is_dual_active_detector"`
	IsStackWiseVirtualLink types.Bool `tfsdk:"is_stack_wise_virtual_link"`
}

type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePort struct {
	AccessPolicyNumber          types.Int64                                                                   `tfsdk:"access_policy_number"`
	AccessPolicyType            types.String                                                                  `tfsdk:"access_policy_type"`
	AllowedVLANs                types.String                                                                  `tfsdk:"allowed_vlans"`
	DaiTrusted                  types.Bool                                                                    `tfsdk:"dai_trusted"`
	Dot3Az                      *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortDot3Az           `tfsdk:"dot3az"`
	Enabled                     types.Bool                                                                    `tfsdk:"enabled"`
	FlexibleStackingEnabled     types.Bool                                                                    `tfsdk:"flexible_stacking_enabled"`
	IsolationEnabled            types.Bool                                                                    `tfsdk:"isolation_enabled"`
	LinkNegotiation             types.String                                                                  `tfsdk:"link_negotiation"`
	LinkNegotiationCapabilities types.List                                                                    `tfsdk:"link_negotiation_capabilities"`
	MacAllowList                types.List                                                                    `tfsdk:"mac_allow_list"`
	Mirror                      *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortMirror           `tfsdk:"mirror"`
	Module                      *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortModule           `tfsdk:"module"`
	Name                        types.String                                                                  `tfsdk:"name"`
	PoeEnabled                  types.Bool                                                                    `tfsdk:"poe_enabled"`
	PortID                      types.String                                                                  `tfsdk:"port_id"`
	PortScheduleID              types.String                                                                  `tfsdk:"port_schedule_id"`
	Profile                     *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortProfile          `tfsdk:"profile"`
	RstpEnabled                 types.Bool                                                                    `tfsdk:"rstp_enabled"`
	Schedule                    *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortSchedule         `tfsdk:"schedule"`
	StackwiseVirtual            *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortStackwiseVirtual `tfsdk:"stackwise_virtual"`
	StickyMacAllowList          types.List                                                                    `tfsdk:"sticky_mac_allow_list"`
	StickyMacAllowListLimit     types.Int64                                                                   `tfsdk:"sticky_mac_allow_list_limit"`
	StormControlEnabled         types.Bool                                                                    `tfsdk:"storm_control_enabled"`
	StpGuard                    types.String                                                                  `tfsdk:"stp_guard"`
	Tags                        types.List                                                                    `tfsdk:"tags"`
	Type                        types.String                                                                  `tfsdk:"type"`
	Udld                        types.String                                                                  `tfsdk:"udld"`
	VLAN                        types.Int64                                                                   `tfsdk:"vlan"`
	VoiceVLAN                   types.Int64                                                                   `tfsdk:"voice_vlan"`
}

type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortDot3Az struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortMirror struct {
	Mode types.String `tfsdk:"mode"`
}

type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortModule struct {
	Model types.String `tfsdk:"model"`
}

type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortProfile struct {
	Enabled types.Bool   `tfsdk:"enabled"`
	ID      types.String `tfsdk:"id"`
	Iname   types.String `tfsdk:"iname"`
}

type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortSchedule struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortStackwiseVirtual struct {
	IsDualActiveDetector   types.Bool `tfsdk:"is_dual_active_detector"`
	IsStackWiseVirtualLink types.Bool `tfsdk:"is_stack_wise_virtual_link"`
}

// ToBody
func ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortsItemsToBody(state OrganizationsConfigTemplatesSwitchProfilesPorts, response *merakigosdk.ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePorts) OrganizationsConfigTemplatesSwitchProfilesPorts {
	var items []ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePorts
	for _, item := range *response {
		itemState := ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePorts{
			AccessPolicyNumber: func() types.Int64 {
				if item.AccessPolicyNumber != nil {
					return types.Int64Value(int64(*item.AccessPolicyNumber))
				}
				return types.Int64{}
			}(),
			AccessPolicyType: func() types.String {
				if item.AccessPolicyType != "" {
					return types.StringValue(item.AccessPolicyType)
				}
				return types.String{}
			}(),
			AllowedVLANs: func() types.String {
				if item.AllowedVLANs != "" {
					return types.StringValue(item.AllowedVLANs)
				}
				return types.String{}
			}(),
			DaiTrusted: func() types.Bool {
				if item.DaiTrusted != nil {
					return types.BoolValue(*item.DaiTrusted)
				}
				return types.Bool{}
			}(),
			Dot3Az: func() *ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsDot3Az {
				if item.Dot3Az != nil {
					return &ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsDot3Az{
						Enabled: func() types.Bool {
							if item.Dot3Az.Enabled != nil {
								return types.BoolValue(*item.Dot3Az.Enabled)
							}
							return types.Bool{}
						}(),
					}
				}
				return nil
			}(),
			Enabled: func() types.Bool {
				if item.Enabled != nil {
					return types.BoolValue(*item.Enabled)
				}
				return types.Bool{}
			}(),
			FlexibleStackingEnabled: func() types.Bool {
				if item.FlexibleStackingEnabled != nil {
					return types.BoolValue(*item.FlexibleStackingEnabled)
				}
				return types.Bool{}
			}(),
			IsolationEnabled: func() types.Bool {
				if item.IsolationEnabled != nil {
					return types.BoolValue(*item.IsolationEnabled)
				}
				return types.Bool{}
			}(),
			LinkNegotiation: func() types.String {
				if item.LinkNegotiation != "" {
					return types.StringValue(item.LinkNegotiation)
				}
				return types.String{}
			}(),
			LinkNegotiationCapabilities: StringSliceToList(item.LinkNegotiationCapabilities),
			MacAllowList:                StringSliceToList(item.MacAllowList),
			Mirror: func() *ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsMirror {
				if item.Mirror != nil {
					return &ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsMirror{
						Mode: func() types.String {
							if item.Mirror.Mode != "" {
								return types.StringValue(item.Mirror.Mode)
							}
							return types.String{}
						}(),
					}
				}
				return nil
			}(),
			Module: func() *ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsModule {
				if item.Module != nil {
					return &ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsModule{
						Model: func() types.String {
							if item.Module.Model != "" {
								return types.StringValue(item.Module.Model)
							}
							return types.String{}
						}(),
					}
				}
				return nil
			}(),
			Name: func() types.String {
				if item.Name != "" {
					return types.StringValue(item.Name)
				}
				return types.String{}
			}(),
			PoeEnabled: func() types.Bool {
				if item.PoeEnabled != nil {
					return types.BoolValue(*item.PoeEnabled)
				}
				return types.Bool{}
			}(),
			PortID: func() types.String {
				if item.PortID != "" {
					return types.StringValue(item.PortID)
				}
				return types.String{}
			}(),
			PortScheduleID: func() types.String {
				if item.PortScheduleID != "" {
					return types.StringValue(item.PortScheduleID)
				}
				return types.String{}
			}(),
			Profile: func() *ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsProfile {
				if item.Profile != nil {
					return &ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsProfile{
						Enabled: func() types.Bool {
							if item.Profile.Enabled != nil {
								return types.BoolValue(*item.Profile.Enabled)
							}
							return types.Bool{}
						}(),
						ID: func() types.String {
							if item.Profile.ID != "" {
								return types.StringValue(item.Profile.ID)
							}
							return types.String{}
						}(),
						Iname: func() types.String {
							if item.Profile.Iname != "" {
								return types.StringValue(item.Profile.Iname)
							}
							return types.String{}
						}(),
					}
				}
				return nil
			}(),
			RstpEnabled: func() types.Bool {
				if item.RstpEnabled != nil {
					return types.BoolValue(*item.RstpEnabled)
				}
				return types.Bool{}
			}(),
			Schedule: func() *ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsSchedule {
				if item.Schedule != nil {
					return &ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsSchedule{
						ID: func() types.String {
							if item.Schedule.ID != "" {
								return types.StringValue(item.Schedule.ID)
							}
							return types.String{}
						}(),
						Name: func() types.String {
							if item.Schedule.Name != "" {
								return types.StringValue(item.Schedule.Name)
							}
							return types.String{}
						}(),
					}
				}
				return nil
			}(),
			StackwiseVirtual: func() *ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsStackwiseVirtual {
				if item.StackwiseVirtual != nil {
					return &ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfilePortsStackwiseVirtual{
						IsDualActiveDetector: func() types.Bool {
							if item.StackwiseVirtual.IsDualActiveDetector != nil {
								return types.BoolValue(*item.StackwiseVirtual.IsDualActiveDetector)
							}
							return types.Bool{}
						}(),
						IsStackWiseVirtualLink: func() types.Bool {
							if item.StackwiseVirtual.IsStackWiseVirtualLink != nil {
								return types.BoolValue(*item.StackwiseVirtual.IsStackWiseVirtualLink)
							}
							return types.Bool{}
						}(),
					}
				}
				return nil
			}(),
			StickyMacAllowList: StringSliceToList(item.StickyMacAllowList),
			StickyMacAllowListLimit: func() types.Int64 {
				if item.StickyMacAllowListLimit != nil {
					return types.Int64Value(int64(*item.StickyMacAllowListLimit))
				}
				return types.Int64{}
			}(),
			StormControlEnabled: func() types.Bool {
				if item.StormControlEnabled != nil {
					return types.BoolValue(*item.StormControlEnabled)
				}
				return types.Bool{}
			}(),
			StpGuard: func() types.String {
				if item.StpGuard != "" {
					return types.StringValue(item.StpGuard)
				}
				return types.String{}
			}(),
			Tags: StringSliceToList(item.Tags),
			Type: func() types.String {
				if item.Type != "" {
					return types.StringValue(item.Type)
				}
				return types.String{}
			}(),
			Udld: func() types.String {
				if item.Udld != "" {
					return types.StringValue(item.Udld)
				}
				return types.String{}
			}(),
			VLAN: func() types.Int64 {
				if item.VLAN != nil {
					return types.Int64Value(int64(*item.VLAN))
				}
				return types.Int64{}
			}(),
			VoiceVLAN: func() types.Int64 {
				if item.VoiceVLAN != nil {
					return types.Int64Value(int64(*item.VoiceVLAN))
				}
				return types.Int64{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortItemToBody(state OrganizationsConfigTemplatesSwitchProfilesPorts, response *merakigosdk.ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePort) OrganizationsConfigTemplatesSwitchProfilesPorts {
	itemState := ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePort{
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
		Dot3Az: func() *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortDot3Az {
			if response.Dot3Az != nil {
				return &ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortDot3Az{
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
		Mirror: func() *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortMirror {
			if response.Mirror != nil {
				return &ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortMirror{
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
		Module: func() *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortModule {
			if response.Module != nil {
				return &ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortModule{
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
		Profile: func() *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortProfile {
			if response.Profile != nil {
				return &ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortProfile{
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
		Schedule: func() *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortSchedule {
			if response.Schedule != nil {
				return &ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortSchedule{
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
		StackwiseVirtual: func() *ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortStackwiseVirtual {
			if response.StackwiseVirtual != nil {
				return &ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePortStackwiseVirtual{
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
	state.Item = &itemState
	return state
}
