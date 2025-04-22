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
	"strings"

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSwitchStacksRoutingInterfacesDhcpResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchStacksRoutingInterfacesDhcpResource{}
)

func NewNetworksSwitchStacksRoutingInterfacesDhcpResource() resource.Resource {
	return &NetworksSwitchStacksRoutingInterfacesDhcpResource{}
}

type NetworksSwitchStacksRoutingInterfacesDhcpResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchStacksRoutingInterfacesDhcpResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchStacksRoutingInterfacesDhcpResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_stacks_routing_interfaces_dhcp"
}

func (r *NetworksSwitchStacksRoutingInterfacesDhcpResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"boot_file_name": schema.StringAttribute{
				MarkdownDescription: `The PXE boot server file name for the DHCP server running on the switch stack interface`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"boot_next_server": schema.StringAttribute{
				MarkdownDescription: `The PXE boot server IP for the DHCP server running on the switch stack interface`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"boot_options_enabled": schema.BoolAttribute{
				MarkdownDescription: `Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch stack interface`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"dhcp_lease_time": schema.StringAttribute{
				MarkdownDescription: `The DHCP lease time config for the dhcp server running on the switch stack interface ('30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week')
                                  Allowed values: [1 day,1 hour,1 week,12 hours,30 minutes,4 hours]`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"1 day",
						"1 hour",
						"1 week",
						"12 hours",
						"30 minutes",
						"4 hours",
					),
				},
			},
			"dhcp_mode": schema.StringAttribute{
				MarkdownDescription: `The DHCP mode options for the switch stack interface ('dhcpDisabled', 'dhcpRelay' or 'dhcpServer')
                                  Allowed values: [dhcpDisabled,dhcpRelay,dhcpServer]`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"dhcpDisabled",
						"dhcpRelay",
						"dhcpServer",
					),
				},
			},
			"dhcp_options": schema.SetNestedAttribute{
				MarkdownDescription: `Array of DHCP options consisting of code, type and value for the DHCP server running on the switch stack interface`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"code": schema.StringAttribute{
							MarkdownDescription: `The code for DHCP option which should be from 2 to 254`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `The type of the DHCP option which should be one of ('text', 'ip', 'integer' or 'hex')
                                        Allowed values: [hex,integer,ip,text]`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"hex",
									"integer",
									"ip",
									"text",
								),
							},
						},
						"value": schema.StringAttribute{
							MarkdownDescription: `The value of the DHCP option`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"dhcp_relay_server_ips": schema.SetAttribute{
				MarkdownDescription: `The DHCP relay server IPs to which DHCP packets would get relayed for the switch stack interface`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"dns_custom_nameservers": schema.SetAttribute{
				MarkdownDescription: `The DHCP name server IPs when DHCP name server option is 'custom'`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"dns_nameservers_option": schema.StringAttribute{
				MarkdownDescription: `The DHCP name server option for the dhcp server running on the switch stack interface ('googlePublicDns', 'openDns' or 'custom')
                                  Allowed values: [custom,googlePublicDns,openDns]`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"custom",
						"googlePublicDns",
						"openDns",
					),
				},
			},
			"fixed_ip_assignments": schema.SetNestedAttribute{
				MarkdownDescription: `Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"ip": schema.StringAttribute{
							MarkdownDescription: `The IP address of the client which has fixed IP address assigned to it`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"mac": schema.StringAttribute{
							MarkdownDescription: `The MAC address of the client which has fixed IP address`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the client which has fixed IP address`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: `interfaceId path parameter. Interface ID`,
				Required:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"reserved_ip_ranges": schema.SetNestedAttribute{
				MarkdownDescription: `Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"comment": schema.StringAttribute{
							MarkdownDescription: `The comment for the reserved IP range`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"end": schema.StringAttribute{
							MarkdownDescription: `The ending IP address of the reserved IP range`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"start": schema.StringAttribute{
							MarkdownDescription: `The starting IP address of the reserved IP range`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"switch_stack_id": schema.StringAttribute{
				MarkdownDescription: `switchStackId path parameter. Switch stack ID`,
				Required:            true,
			},
		},
	}
}

func (r *NetworksSwitchStacksRoutingInterfacesDhcpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchStacksRoutingInterfacesDhcpRs

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
	vvNetworkID := data.NetworkID.ValueString()
	vvSwitchStackID := data.SwitchStackID.ValueString()
	vvInterfaceID := data.InterfaceID.ValueString()
	//Has Item and not has items

	if vvNetworkID != "" && vvSwitchStackID != "" && vvInterfaceID != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Switch.GetNetworkSwitchStackRoutingInterfaceDhcp(vvNetworkID, vvSwitchStackID, vvInterfaceID)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksSwitchStacksRoutingInterfacesDhcp  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksSwitchStacksRoutingInterfacesDhcp only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchStackRoutingInterfaceDhcp(vvNetworkID, vvSwitchStackID, vvInterfaceID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchStackRoutingInterfaceDhcp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchStackRoutingInterfaceDhcp",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Switch.GetNetworkSwitchStackRoutingInterfaceDhcp(vvNetworkID, vvSwitchStackID, vvInterfaceID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStackRoutingInterfaceDhcp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchStackRoutingInterfaceDhcp",
			err.Error(),
		)
		return
	}

	data = ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *NetworksSwitchStacksRoutingInterfacesDhcpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchStacksRoutingInterfacesDhcpRs

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
	vvSwitchStackID := data.SwitchStackID.ValueString()
	vvInterfaceID := data.InterfaceID.ValueString()
	responseGet, restyRespGet, err := r.client.Switch.GetNetworkSwitchStackRoutingInterfaceDhcp(vvNetworkID, vvSwitchStackID, vvInterfaceID)
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
				"Failure when executing GetNetworkSwitchStackRoutingInterfaceDhcp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchStackRoutingInterfaceDhcp",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchStacksRoutingInterfacesDhcpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 3 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("switch_stack_id"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("interface_id"), idParts[2])...)
}

func (r *NetworksSwitchStacksRoutingInterfacesDhcpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSwitchStacksRoutingInterfacesDhcpRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvSwitchStackID := data.SwitchStackID.ValueString()
	vvInterfaceID := data.InterfaceID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchStackRoutingInterfaceDhcp(vvNetworkID, vvSwitchStackID, vvInterfaceID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchStackRoutingInterfaceDhcp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchStackRoutingInterfaceDhcp",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchStacksRoutingInterfacesDhcpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksSwitchStacksRoutingInterfacesDhcp", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSwitchStacksRoutingInterfacesDhcpRs struct {
	NetworkID            types.String                                                                   `tfsdk:"network_id"`
	SwitchStackID        types.String                                                                   `tfsdk:"switch_stack_id"`
	InterfaceID          types.String                                                                   `tfsdk:"interface_id"`
	BootFileName         types.String                                                                   `tfsdk:"boot_file_name"`
	BootNextServer       types.String                                                                   `tfsdk:"boot_next_server"`
	BootOptionsEnabled   types.Bool                                                                     `tfsdk:"boot_options_enabled"`
	DhcpLeaseTime        types.String                                                                   `tfsdk:"dhcp_lease_time"`
	DhcpMode             types.String                                                                   `tfsdk:"dhcp_mode"`
	DhcpOptions          *[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpDhcpOptionsRs        `tfsdk:"dhcp_options"`
	DhcpRelayServerIPs   types.Set                                                                      `tfsdk:"dhcp_relay_server_ips"`
	DNSCustomNameservers types.Set                                                                      `tfsdk:"dns_custom_nameservers"`
	DNSNameserversOption types.String                                                                   `tfsdk:"dns_nameservers_option"`
	FixedIPAssignments   *[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpFixedIpAssignmentsRs `tfsdk:"fixed_ip_assignments"`
	ReservedIPRanges     *[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpReservedIpRangesRs   `tfsdk:"reserved_ip_ranges"`
}

type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpDhcpOptionsRs struct {
	Code  types.String `tfsdk:"code"`
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpFixedIpAssignmentsRs struct {
	IP   types.String `tfsdk:"ip"`
	Mac  types.String `tfsdk:"mac"`
	Name types.String `tfsdk:"name"`
}

type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpReservedIpRangesRs struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

// FromBody
func (r *NetworksSwitchStacksRoutingInterfacesDhcpRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcp {
	emptyString := ""
	bootFileName := new(string)
	if !r.BootFileName.IsUnknown() && !r.BootFileName.IsNull() {
		*bootFileName = r.BootFileName.ValueString()
	} else {
		bootFileName = &emptyString
	}
	bootNextServer := new(string)
	if !r.BootNextServer.IsUnknown() && !r.BootNextServer.IsNull() {
		*bootNextServer = r.BootNextServer.ValueString()
	} else {
		bootNextServer = &emptyString
	}
	bootOptionsEnabled := new(bool)
	if !r.BootOptionsEnabled.IsUnknown() && !r.BootOptionsEnabled.IsNull() {
		*bootOptionsEnabled = r.BootOptionsEnabled.ValueBool()
	} else {
		bootOptionsEnabled = nil
	}
	dhcpLeaseTime := new(string)
	if !r.DhcpLeaseTime.IsUnknown() && !r.DhcpLeaseTime.IsNull() {
		*dhcpLeaseTime = r.DhcpLeaseTime.ValueString()
	} else {
		dhcpLeaseTime = &emptyString
	}
	dhcpMode := new(string)
	if !r.DhcpMode.IsUnknown() && !r.DhcpMode.IsNull() {
		*dhcpMode = r.DhcpMode.ValueString()
	} else {
		dhcpMode = &emptyString
	}
	var requestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions []merakigosdk.RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions

	if r.DhcpOptions != nil {
		for _, rItem1 := range *r.DhcpOptions {
			code := rItem1.Code.ValueString()
			typeR := rItem1.Type.ValueString()
			value := rItem1.Value.ValueString()
			requestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions = append(requestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions, merakigosdk.RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions{
				Code:  code,
				Type:  typeR,
				Value: value,
			})
			//[debug] Is Array: True
		}
	}
	var dhcpRelayServerIPs []string = nil
	r.DhcpRelayServerIPs.ElementsAs(ctx, &dhcpRelayServerIPs, false)
	var dNSCustomNameservers []string = nil
	r.DNSCustomNameservers.ElementsAs(ctx, &dNSCustomNameservers, false)
	dNSNameserversOption := new(string)
	if !r.DNSNameserversOption.IsUnknown() && !r.DNSNameserversOption.IsNull() {
		*dNSNameserversOption = r.DNSNameserversOption.ValueString()
	} else {
		dNSNameserversOption = &emptyString
	}
	var requestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpFixedIPAssignments []merakigosdk.RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpFixedIPAssignments

	if r.FixedIPAssignments != nil {
		for _, rItem1 := range *r.FixedIPAssignments {
			ip := rItem1.IP.ValueString()
			mac := rItem1.Mac.ValueString()
			name := rItem1.Name.ValueString()
			requestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpFixedIPAssignments = append(requestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpFixedIPAssignments, merakigosdk.RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpFixedIPAssignments{
				IP:   ip,
				Mac:  mac,
				Name: name,
			})
			//[debug] Is Array: True
		}
	}
	var requestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpReservedIPRanges []merakigosdk.RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpReservedIPRanges

	if r.ReservedIPRanges != nil {
		for _, rItem1 := range *r.ReservedIPRanges {
			comment := rItem1.Comment.ValueString()
			end := rItem1.End.ValueString()
			start := rItem1.Start.ValueString()
			requestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpReservedIPRanges = append(requestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpReservedIPRanges, merakigosdk.RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpReservedIPRanges{
				Comment: comment,
				End:     end,
				Start:   start,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcp{
		BootFileName:       *bootFileName,
		BootNextServer:     *bootNextServer,
		BootOptionsEnabled: bootOptionsEnabled,
		DhcpLeaseTime:      *dhcpLeaseTime,
		DhcpMode:           *dhcpMode,
		DhcpOptions: func() *[]merakigosdk.RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions {
			if len(requestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions) > 0 {
				return &requestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpDhcpOptions
			}
			return nil
		}(),
		DhcpRelayServerIPs:   dhcpRelayServerIPs,
		DNSCustomNameservers: dNSCustomNameservers,
		DNSNameserversOption: *dNSNameserversOption,
		FixedIPAssignments: func() *[]merakigosdk.RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpFixedIPAssignments {
			if len(requestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpFixedIPAssignments) > 0 {
				return &requestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpFixedIPAssignments
			}
			return nil
		}(),
		ReservedIPRanges: func() *[]merakigosdk.RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpReservedIPRanges {
			if len(requestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpReservedIPRanges) > 0 {
				return &requestSwitchUpdateNetworkSwitchStackRoutingInterfaceDhcpReservedIPRanges
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpItemToBodyRs(state NetworksSwitchStacksRoutingInterfacesDhcpRs, response *merakigosdk.ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcp, is_read bool) NetworksSwitchStacksRoutingInterfacesDhcpRs {
	itemState := NetworksSwitchStacksRoutingInterfacesDhcpRs{
		BootFileName:   types.StringValue(response.BootFileName),
		BootNextServer: types.StringValue(response.BootNextServer),
		BootOptionsEnabled: func() types.Bool {
			if response.BootOptionsEnabled != nil {
				return types.BoolValue(*response.BootOptionsEnabled)
			}
			return types.Bool{}
		}(),
		DhcpLeaseTime: types.StringValue(response.DhcpLeaseTime),
		DhcpMode:      types.StringValue(response.DhcpMode),
		DhcpOptions: func() *[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpDhcpOptionsRs {
			if response.DhcpOptions != nil {
				result := make([]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpDhcpOptionsRs, len(*response.DhcpOptions))
				for i, dhcpOptions := range *response.DhcpOptions {
					result[i] = ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpDhcpOptionsRs{
						Code:  types.StringValue(dhcpOptions.Code),
						Type:  types.StringValue(dhcpOptions.Type),
						Value: types.StringValue(dhcpOptions.Value),
					}
				}
				return &result
			}
			return nil
		}(),
		DhcpRelayServerIPs:   StringSliceToSet(response.DhcpRelayServerIPs),
		DNSCustomNameservers: StringSliceToSet(response.DNSCustomNameservers),
		DNSNameserversOption: types.StringValue(response.DNSNameserversOption),
		FixedIPAssignments: func() *[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpFixedIpAssignmentsRs {
			if response.FixedIPAssignments != nil {
				result := make([]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpFixedIpAssignmentsRs, len(*response.FixedIPAssignments))
				for i, fixedIPAssignments := range *response.FixedIPAssignments {
					result[i] = ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpFixedIpAssignmentsRs{
						IP:   types.StringValue(fixedIPAssignments.IP),
						Mac:  types.StringValue(fixedIPAssignments.Mac),
						Name: types.StringValue(fixedIPAssignments.Name),
					}
				}
				return &result
			}
			return nil
		}(),
		ReservedIPRanges: func() *[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpReservedIpRangesRs {
			if response.ReservedIPRanges != nil {
				result := make([]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpReservedIpRangesRs, len(*response.ReservedIPRanges))
				for i, reservedIPRanges := range *response.ReservedIPRanges {
					result[i] = ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpReservedIpRangesRs{
						Comment: types.StringValue(reservedIPRanges.Comment),
						End:     types.StringValue(reservedIPRanges.End),
						Start:   types.StringValue(reservedIPRanges.Start),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSwitchStacksRoutingInterfacesDhcpRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSwitchStacksRoutingInterfacesDhcpRs)
}
