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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

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
	_ resource.Resource              = &DevicesSwitchRoutingInterfacesDhcpResource{}
	_ resource.ResourceWithConfigure = &DevicesSwitchRoutingInterfacesDhcpResource{}
)

func NewDevicesSwitchRoutingInterfacesDhcpResource() resource.Resource {
	return &DevicesSwitchRoutingInterfacesDhcpResource{}
}

type DevicesSwitchRoutingInterfacesDhcpResource struct {
	client *merakigosdk.Client
}

func (r *DevicesSwitchRoutingInterfacesDhcpResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesSwitchRoutingInterfacesDhcpResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_switch_routing_interfaces_dhcp"
}

func (r *DevicesSwitchRoutingInterfacesDhcpResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
		},
	}
}

func (r *DevicesSwitchRoutingInterfacesDhcpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesSwitchRoutingInterfacesDhcpRs

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
	vvInterfaceID := data.InterfaceID.ValueString()
	//Item
	responseVerifyItem, restyResp1, err := r.client.Switch.GetDeviceSwitchRoutingInterfaceDhcp(vvSerial, vvInterfaceID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource DevicesSwitchRoutingInterfacesDhcp only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource DevicesSwitchRoutingInterfacesDhcp only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateDeviceSwitchRoutingInterfaceDhcp(vvSerial, vvInterfaceID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceSwitchRoutingInterfaceDhcp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceSwitchRoutingInterfaceDhcp",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Switch.GetDeviceSwitchRoutingInterfaceDhcp(vvSerial, vvInterfaceID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceSwitchRoutingInterfaceDhcp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceSwitchRoutingInterfaceDhcp",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesSwitchRoutingInterfacesDhcpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesSwitchRoutingInterfacesDhcpRs

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

	vvSerial := data.Serial.ValueString()
	vvInterfaceID := data.InterfaceID.ValueString()
	responseGet, restyRespGet, err := r.client.Switch.GetDeviceSwitchRoutingInterfaceDhcp(vvSerial, vvInterfaceID)
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
				"Failure when executing GetDeviceSwitchRoutingInterfaceDhcp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceSwitchRoutingInterfaceDhcp",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesSwitchRoutingInterfacesDhcpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("interface_id"), idParts[1])...)
}

func (r *DevicesSwitchRoutingInterfacesDhcpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesSwitchRoutingInterfacesDhcpRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvSerial := data.Serial.ValueString()
	vvInterfaceID := data.InterfaceID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateDeviceSwitchRoutingInterfaceDhcp(vvSerial, vvInterfaceID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceSwitchRoutingInterfaceDhcp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceSwitchRoutingInterfaceDhcp",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesSwitchRoutingInterfacesDhcpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting DevicesSwitchRoutingInterfacesDhcp", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesSwitchRoutingInterfacesDhcpRs struct {
	Serial               types.String                                                             `tfsdk:"serial"`
	InterfaceID          types.String                                                             `tfsdk:"interface_id"`
	BootFileName         types.String                                                             `tfsdk:"boot_file_name"`
	BootNextServer       types.String                                                             `tfsdk:"boot_next_server"`
	BootOptionsEnabled   types.Bool                                                               `tfsdk:"boot_options_enabled"`
	DhcpLeaseTime        types.String                                                             `tfsdk:"dhcp_lease_time"`
	DhcpMode             types.String                                                             `tfsdk:"dhcp_mode"`
	DhcpOptions          *[]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpDhcpOptionsRs        `tfsdk:"dhcp_options"`
	DhcpRelayServerIPs   types.Set                                                                `tfsdk:"dhcp_relay_server_ips"`
	DNSCustomNameservers types.Set                                                                `tfsdk:"dns_custom_nameservers"`
	DNSNameserversOption types.String                                                             `tfsdk:"dns_nameservers_option"`
	FixedIPAssignments   *[]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpFixedIpAssignmentsRs `tfsdk:"fixed_ip_assignments"`
	ReservedIPRanges     *[]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpReservedIpRangesRs   `tfsdk:"reserved_ip_ranges"`
}

type ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpDhcpOptionsRs struct {
	Code  types.String `tfsdk:"code"`
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpFixedIpAssignmentsRs struct {
	IP   types.String `tfsdk:"ip"`
	Mac  types.String `tfsdk:"mac"`
	Name types.String `tfsdk:"name"`
}

type ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpReservedIpRangesRs struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

// FromBody
func (r *DevicesSwitchRoutingInterfacesDhcpRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcp {
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
	var requestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpDhcpOptions []merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpDhcpOptions
	if r.DhcpOptions != nil {
		for _, rItem1 := range *r.DhcpOptions {
			code := rItem1.Code.ValueString()
			typeR := rItem1.Type.ValueString()
			value := rItem1.Value.ValueString()
			requestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpDhcpOptions = append(requestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpDhcpOptions, merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpDhcpOptions{
				Code:  code,
				Type:  typeR,
				Value: value,
			})
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
	var requestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpFixedIPAssignments []merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpFixedIPAssignments
	if r.FixedIPAssignments != nil {
		for _, rItem1 := range *r.FixedIPAssignments {
			iP := rItem1.IP.ValueString()
			mac := rItem1.Mac.ValueString()
			name := rItem1.Name.ValueString()
			requestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpFixedIPAssignments = append(requestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpFixedIPAssignments, merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpFixedIPAssignments{
				IP:   iP,
				Mac:  mac,
				Name: name,
			})
		}
	}
	var requestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpReservedIPRanges []merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpReservedIPRanges
	if r.ReservedIPRanges != nil {
		for _, rItem1 := range *r.ReservedIPRanges {
			comment := rItem1.Comment.ValueString()
			end := rItem1.End.ValueString()
			start := rItem1.Start.ValueString()
			requestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpReservedIPRanges = append(requestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpReservedIPRanges, merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpReservedIPRanges{
				Comment: comment,
				End:     end,
				Start:   start,
			})
		}
	}
	out := merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcp{
		BootFileName:       *bootFileName,
		BootNextServer:     *bootNextServer,
		BootOptionsEnabled: bootOptionsEnabled,
		DhcpLeaseTime:      *dhcpLeaseTime,
		DhcpMode:           *dhcpMode,
		DhcpOptions: func() *[]merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpDhcpOptions {
			if len(requestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpDhcpOptions) > 0 {
				return &requestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpDhcpOptions
			}
			return nil
		}(),
		DhcpRelayServerIPs:   dhcpRelayServerIPs,
		DNSCustomNameservers: dNSCustomNameservers,
		DNSNameserversOption: *dNSNameserversOption,
		FixedIPAssignments: func() *[]merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpFixedIPAssignments {
			if len(requestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpFixedIPAssignments) > 0 {
				return &requestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpFixedIPAssignments
			}
			return nil
		}(),
		ReservedIPRanges: func() *[]merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpReservedIPRanges {
			if len(requestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpReservedIPRanges) > 0 {
				return &requestSwitchUpdateDeviceSwitchRoutingInterfaceDhcpReservedIPRanges
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpItemToBodyRs(state DevicesSwitchRoutingInterfacesDhcpRs, response *merakigosdk.ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcp, is_read bool) DevicesSwitchRoutingInterfacesDhcpRs {
	itemState := DevicesSwitchRoutingInterfacesDhcpRs{
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
		DhcpOptions: func() *[]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpDhcpOptionsRs {
			if response.DhcpOptions != nil {
				result := make([]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpDhcpOptionsRs, len(*response.DhcpOptions))
				for i, dhcpOptions := range *response.DhcpOptions {
					result[i] = ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpDhcpOptionsRs{
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
		FixedIPAssignments: func() *[]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpFixedIpAssignmentsRs {
			if response.FixedIPAssignments != nil {
				result := make([]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpFixedIpAssignmentsRs, len(*response.FixedIPAssignments))
				for i, fixedIPAssignments := range *response.FixedIPAssignments {
					result[i] = ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpFixedIpAssignmentsRs{
						IP:   types.StringValue(fixedIPAssignments.IP),
						Mac:  types.StringValue(fixedIPAssignments.Mac),
						Name: types.StringValue(fixedIPAssignments.Name),
					}
				}
				return &result
			}
			return nil
		}(),
		ReservedIPRanges: func() *[]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpReservedIpRangesRs {
			if response.ReservedIPRanges != nil {
				result := make([]ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpReservedIpRangesRs, len(*response.ReservedIPRanges))
				for i, reservedIPRanges := range *response.ReservedIPRanges {
					result[i] = ResponseSwitchGetDeviceSwitchRoutingInterfaceDhcpReservedIpRangesRs{
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
	if !state.DhcpRelayServerIPs.IsUnknown() {
		itemState.DhcpRelayServerIPs = state.DhcpRelayServerIPs
	} else {
		itemState.DhcpRelayServerIPs = types.SetNull(types.StringType)
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesSwitchRoutingInterfacesDhcpRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesSwitchRoutingInterfacesDhcpRs)
}
