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

// RESOURCE ACTION

import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesWirelessAlternateManagementInterfaceIPv6Resource{}
	_ resource.ResourceWithConfigure = &DevicesWirelessAlternateManagementInterfaceIPv6Resource{}
)

func NewDevicesWirelessAlternateManagementInterfaceIPv6Resource() resource.Resource {
	return &DevicesWirelessAlternateManagementInterfaceIPv6Resource{}
}

type DevicesWirelessAlternateManagementInterfaceIPv6Resource struct {
	client *merakigosdk.Client
}

func (r *DevicesWirelessAlternateManagementInterfaceIPv6Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesWirelessAlternateManagementInterfaceIPv6Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_wireless_alternate_management_interface_ipv6"
}

// resourceAction
func (r *DevicesWirelessAlternateManagementInterfaceIPv6Resource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

					"addresses": schema.SetNestedAttribute{
						MarkdownDescription: `configured alternate management interface addresses`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"address": schema.StringAttribute{
									MarkdownDescription: `The IP address configured for the alternate management interface`,
									Computed:            true,
								},
								"assignment_mode": schema.StringAttribute{
									MarkdownDescription: `The type of address assignment. Either static or dynamic.
                                                Allowed values: [dynamic,static]`,
									Computed: true,
								},
								"gateway": schema.StringAttribute{
									MarkdownDescription: `The gateway address configured for the alternate managment interface`,
									Computed:            true,
								},
								"nameservers": schema.SingleNestedAttribute{
									MarkdownDescription: `The DNS servers settings for this address.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"addresses": schema.SetAttribute{
											MarkdownDescription: `Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.`,
											Computed:            true,
											ElementType:         types.StringType,
										},
									},
								},
								"prefix": schema.StringAttribute{
									MarkdownDescription: `The IPv6 prefix of the interface. Required if IPv6 object is included.`,
									Computed:            true,
								},
								"protocol": schema.StringAttribute{
									MarkdownDescription: `The IP protocol used for the address
                                                Allowed values: [ipv4,ipv6]`,
									Computed: true,
								},
							},
						},
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"addresses": schema.SetNestedAttribute{
						MarkdownDescription: `configured alternate management interface addresses`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"address": schema.StringAttribute{
									MarkdownDescription: `The IP address configured for the alternate management interface`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"assignment_mode": schema.StringAttribute{
									MarkdownDescription: `The type of address assignment. Either static or dynamic.
                                              Allowed values: [dynamic,static]`,
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"gateway": schema.StringAttribute{
									MarkdownDescription: `The gateway address configured for the alternate managment interface`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"nameservers": schema.SingleNestedAttribute{
									MarkdownDescription: `The DNS servers settings for this address.`,
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"addresses": schema.ListAttribute{
											MarkdownDescription: `Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.`,
											Optional:            true,
											Computed:            true,
											ElementType:         types.StringType,
										},
									},
								},
								"prefix": schema.StringAttribute{
									MarkdownDescription: `The IPv6 prefix length of the IPv6 interface. Required if IPv6 object is included.`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"protocol": schema.StringAttribute{
									MarkdownDescription: `The IP protocol used for the address
                                              Allowed values: [ipv4,ipv6]`,
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
func (r *DevicesWirelessAlternateManagementInterfaceIPv6Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesWirelessAlternateManagementInterfaceIPv6

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
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp1, err := r.client.Wireless.UpdateDeviceWirelessAlternateManagementInterfaceIPv6(vvSerial, dataRequest)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceWirelessAlternateManagementInterfaceIPv6",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceWirelessAlternateManagementInterfaceIPv6",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6ItemToBody(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesWirelessAlternateManagementInterfaceIPv6Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *DevicesWirelessAlternateManagementInterfaceIPv6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *DevicesWirelessAlternateManagementInterfaceIPv6Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesWirelessAlternateManagementInterfaceIPv6 struct {
	Serial     types.String                                                           `tfsdk:"serial"`
	Item       *ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIpv6  `tfsdk:"item"`
	Parameters *RequestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIpv6Rs `tfsdk:"parameters"`
}

type ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIpv6 struct {
	Addresses *[]ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIpv6Addresses `tfsdk:"addresses"`
}

type ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIpv6Addresses struct {
	Address        types.String                                                                              `tfsdk:"address"`
	AssignmentMode types.String                                                                              `tfsdk:"assignment_mode"`
	Gateway        types.String                                                                              `tfsdk:"gateway"`
	Nameservers    *ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIpv6AddressesNameservers `tfsdk:"nameservers"`
	Prefix         types.String                                                                              `tfsdk:"prefix"`
	Protocol       types.String                                                                              `tfsdk:"protocol"`
}

type ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIpv6AddressesNameservers struct {
	Addresses types.List `tfsdk:"addresses"`
}

type RequestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIpv6Rs struct {
	Addresses *[]RequestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIpv6AddressesRs `tfsdk:"addresses"`
}

type RequestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIpv6AddressesRs struct {
	Address        types.String                                                                               `tfsdk:"address"`
	AssignmentMode types.String                                                                               `tfsdk:"assignment_mode"`
	Gateway        types.String                                                                               `tfsdk:"gateway"`
	Nameservers    *RequestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIpv6AddressesNameserversRs `tfsdk:"nameservers"`
	Prefix         types.String                                                                               `tfsdk:"prefix"`
	Protocol       types.String                                                                               `tfsdk:"protocol"`
}

type RequestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIpv6AddressesNameserversRs struct {
	Addresses types.Set `tfsdk:"addresses"`
}

// FromBody
func (r *DevicesWirelessAlternateManagementInterfaceIPv6) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6 {
	re := *r.Parameters
	var requestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6Addresses []merakigosdk.RequestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6Addresses
	if re.Addresses != nil {
		for _, rItem1 := range *re.Addresses {
			address := rItem1.Address.ValueString()
			assignmentMode := rItem1.AssignmentMode.ValueString()
			gateway := rItem1.Gateway.ValueString()
			var requestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6AddressesNameservers *merakigosdk.RequestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6AddressesNameservers
			if rItem1.Nameservers != nil {
				var addresses []string = nil
				//Hoola aqui
				rItem1.Nameservers.Addresses.ElementsAs(ctx, &addresses, false)
				requestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6AddressesNameservers = &merakigosdk.RequestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6AddressesNameservers{
					Addresses: addresses,
				}
			}
			prefix := rItem1.Prefix.ValueString()
			protocol := rItem1.Protocol.ValueString()
			requestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6Addresses = append(requestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6Addresses, merakigosdk.RequestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6Addresses{
				Address:        address,
				AssignmentMode: assignmentMode,
				Gateway:        gateway,
				Nameservers:    requestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6AddressesNameservers,
				Prefix:         prefix,
				Protocol:       protocol,
			})
		}
	}
	out := merakigosdk.RequestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6{
		Addresses: func() *[]merakigosdk.RequestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6Addresses {
			if len(requestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6Addresses) > 0 {
				return &requestWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6Addresses
			}
			return nil
		}(),
	}
	return &out
}

// ToBody
func ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6ItemToBody(state DevicesWirelessAlternateManagementInterfaceIPv6, response *merakigosdk.ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIPv6) DevicesWirelessAlternateManagementInterfaceIPv6 {
	itemState := ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIpv6{
		Addresses: func() *[]ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIpv6Addresses {
			if response.Addresses != nil {
				result := make([]ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIpv6Addresses, len(*response.Addresses))
				for i, addresses := range *response.Addresses {
					result[i] = ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIpv6Addresses{
						Address:        types.StringValue(addresses.Address),
						AssignmentMode: types.StringValue(addresses.AssignmentMode),
						Gateway:        types.StringValue(addresses.Gateway),
						Nameservers: func() *ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIpv6AddressesNameservers {
							if addresses.Nameservers != nil {
								return &ResponseWirelessUpdateDeviceWirelessAlternateManagementInterfaceIpv6AddressesNameservers{
									Addresses: StringSliceToList(addresses.Nameservers.Addresses),
								}
							}
							return nil
						}(),
						Prefix:   types.StringValue(addresses.Prefix),
						Protocol: types.StringValue(addresses.Protocol),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
