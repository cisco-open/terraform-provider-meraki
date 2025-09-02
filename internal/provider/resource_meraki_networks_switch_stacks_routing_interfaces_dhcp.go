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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"boot_next_server": schema.StringAttribute{
				MarkdownDescription: `The PXE boot server IP for the DHCP server running on the switch stack interface`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"boot_options_enabled": schema.BoolAttribute{
				MarkdownDescription: `Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch stack interface`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"dhcp_lease_time": schema.StringAttribute{
				MarkdownDescription: `The DHCP lease time config for the dhcp server running on the switch stack interface ('30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week')
                                  Allowed values: [1 day,1 hour,1 week,12 hours,30 minutes,4 hours]`,
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
			"dhcp_options": schema.ListNestedAttribute{
				MarkdownDescription: `Array of DHCP options consisting of code, type and value for the DHCP server running on the switch stack interface`,
				Optional:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"code": schema.StringAttribute{
							MarkdownDescription: `The code for DHCP option which should be from 2 to 254`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `The type of the DHCP option which should be one of ('text', 'ip', 'integer' or 'hex')
                                        Allowed values: [hex,integer,ip,text]`,
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
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"dhcp_relay_server_ips": schema.ListAttribute{
				MarkdownDescription: `The DHCP relay server IPs to which DHCP packets would get relayed for the switch stack interface`,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
				Default:     listdefault.StaticValue(types.ListNull(types.StringType)),
			},
			"dns_custom_nameservers": schema.ListAttribute{
				MarkdownDescription: `The DHCP name server IPs when DHCP name server option is 'custom'`,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
				Default:     listdefault.StaticValue(types.ListNull(types.StringType)),
			},
			"dns_nameservers_option": schema.StringAttribute{
				MarkdownDescription: `The DHCP name server option for the dhcp server running on the switch stack interface ('googlePublicDns', 'openDns' or 'custom')
                                  Allowed values: [custom,googlePublicDns,openDns]`,
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
			"fixed_ip_assignments": schema.ListNestedAttribute{
				MarkdownDescription: `Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface`,
				Optional:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"ip": schema.StringAttribute{
							MarkdownDescription: `The IP address of the client which has fixed IP address assigned to it`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"mac": schema.StringAttribute{
							MarkdownDescription: `The MAC address of the client which has fixed IP address`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the client which has fixed IP address`,
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
			"reserved_ip_ranges": schema.ListNestedAttribute{
				MarkdownDescription: `Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface`,
				Optional:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"comment": schema.StringAttribute{
							MarkdownDescription: `The comment for the reserved IP range`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"end": schema.StringAttribute{
							MarkdownDescription: `The ending IP address of the reserved IP range`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"start": schema.StringAttribute{
							MarkdownDescription: `The starting IP address of the reserved IP range`,
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

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchStackRoutingInterfaceDhcp(vvNetworkID, vvSwitchStackID, vvInterfaceID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchStackRoutingInterfaceDhcp",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchStackRoutingInterfaceDhcp",
			err.Error(),
		)
		return
	}

	// Assign data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *NetworksSwitchStacksRoutingInterfacesDhcpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchStacksRoutingInterfacesDhcpRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
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
				restyRespGet.String(),
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
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *NetworksSwitchStacksRoutingInterfacesDhcpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 3 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: networkId,switchStackId,interfaceId. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("switch_stack_id"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("interface_id"), idParts[2])...)
}

func (r *NetworksSwitchStacksRoutingInterfacesDhcpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworksSwitchStacksRoutingInterfacesDhcpRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvNetworkID := plan.NetworkID.ValueString()
	vvSwitchStackID := plan.SwitchStackID.ValueString()
	vvInterfaceID := plan.InterfaceID.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchStackRoutingInterfaceDhcp(vvNetworkID, vvSwitchStackID, vvInterfaceID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchStackRoutingInterfaceDhcp",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchStackRoutingInterfaceDhcp",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
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
	DhcpRelayServerIPs   types.List                                                                     `tfsdk:"dhcp_relay_server_ips"`
	DNSCustomNameservers types.List                                                                     `tfsdk:"dns_custom_nameservers"`
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
		BootFileName: func() types.String {
			if response.BootFileName != "" {
				return types.StringValue(response.BootFileName)
			}
			return types.String{}
		}(),
		BootNextServer: func() types.String {
			if response.BootNextServer != "" {
				return types.StringValue(response.BootNextServer)
			}
			return types.String{}
		}(),
		BootOptionsEnabled: func() types.Bool {
			if response.BootOptionsEnabled != nil {
				return types.BoolValue(*response.BootOptionsEnabled)
			}
			return types.Bool{}
		}(),
		DhcpLeaseTime: func() types.String {
			if response.DhcpLeaseTime != "" {
				return types.StringValue(response.DhcpLeaseTime)
			}
			return types.String{}
		}(),
		DhcpMode: func() types.String {
			if response.DhcpMode != "" {
				return types.StringValue(response.DhcpMode)
			}
			return types.String{}
		}(),
		DhcpOptions: func() *[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpDhcpOptionsRs {
			if response.DhcpOptions != nil {
				result := make([]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpDhcpOptionsRs, len(*response.DhcpOptions))
				for i, dhcpOptions := range *response.DhcpOptions {
					result[i] = ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpDhcpOptionsRs{
						Code: func() types.String {
							if dhcpOptions.Code != "" {
								return types.StringValue(dhcpOptions.Code)
							}
							return types.String{}
						}(),
						Type: func() types.String {
							if dhcpOptions.Type != "" {
								return types.StringValue(dhcpOptions.Type)
							}
							return types.String{}
						}(),
						Value: func() types.String {
							if dhcpOptions.Value != "" {
								return types.StringValue(dhcpOptions.Value)
							}
							return types.String{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		DhcpRelayServerIPs:   StringSliceToList(response.DhcpRelayServerIPs),
		DNSCustomNameservers: StringSliceToList(response.DNSCustomNameservers),
		DNSNameserversOption: func() types.String {
			if response.DNSNameserversOption != "" {
				return types.StringValue(response.DNSNameserversOption)
			}
			return types.String{}
		}(),
		FixedIPAssignments: func() *[]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpFixedIpAssignmentsRs {
			if response.FixedIPAssignments != nil {
				result := make([]ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpFixedIpAssignmentsRs, len(*response.FixedIPAssignments))
				for i, fixedIPAssignments := range *response.FixedIPAssignments {
					result[i] = ResponseSwitchGetNetworkSwitchStackRoutingInterfaceDhcpFixedIpAssignmentsRs{
						IP: func() types.String {
							if fixedIPAssignments.IP != "" {
								return types.StringValue(fixedIPAssignments.IP)
							}
							return types.String{}
						}(),
						Mac: func() types.String {
							if fixedIPAssignments.Mac != "" {
								return types.StringValue(fixedIPAssignments.Mac)
							}
							return types.String{}
						}(),
						Name: func() types.String {
							if fixedIPAssignments.Name != "" {
								return types.StringValue(fixedIPAssignments.Name)
							}
							return types.String{}
						}(),
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
						Comment: func() types.String {
							if reservedIPRanges.Comment != "" {
								return types.StringValue(reservedIPRanges.Comment)
							}
							return types.String{}
						}(),
						End: func() types.String {
							if reservedIPRanges.End != "" {
								return types.StringValue(reservedIPRanges.End)
							}
							return types.String{}
						}(),
						Start: func() types.String {
							if reservedIPRanges.Start != "" {
								return types.StringValue(reservedIPRanges.Start)
							}
							return types.String{}
						}(),
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
