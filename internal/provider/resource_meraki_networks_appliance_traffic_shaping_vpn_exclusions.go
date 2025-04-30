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

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceTrafficShapingVpnExclusionsResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceTrafficShapingVpnExclusionsResource{}
)

func NewNetworksApplianceTrafficShapingVpnExclusionsResource() resource.Resource {
	return &NetworksApplianceTrafficShapingVpnExclusionsResource{}
}

type NetworksApplianceTrafficShapingVpnExclusionsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceTrafficShapingVpnExclusionsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceTrafficShapingVpnExclusionsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_traffic_shaping_vpn_exclusions"
}

// resourceAction
func (r *NetworksApplianceTrafficShapingVpnExclusionsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"custom": schema.SetNestedAttribute{
						MarkdownDescription: `Custom VPN exclusion rules.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"destination": schema.StringAttribute{
									MarkdownDescription: `Destination address; hostname required for DNS, IPv4 otherwise.`,
									Computed:            true,
								},
								"port": schema.StringAttribute{
									MarkdownDescription: `Destination port.`,
									Computed:            true,
								},
								"protocol": schema.StringAttribute{
									MarkdownDescription: `Protocol.
                                                Allowed values: [any,dns,icmp,tcp,udp]`,
									Computed: true,
								},
							},
						},
					},
					"major_applications": schema.SetNestedAttribute{
						MarkdownDescription: `Major Application based VPN exclusion rules.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `Application's Meraki ID.`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Application's name.
                                                Allowed values: [AWS,Box,Office 365 Sharepoint,Office 365 Suite,Oracle,SAP,Salesforce,Skype & Teams,Slack,Webex,Webex Calling,Webex Meetings,Zoom]`,
									Computed: true,
								},
							},
						},
					},
					"network_id": schema.StringAttribute{
						MarkdownDescription: `ID of the network whose VPN exclusion rules are returned.`,
						Computed:            true,
					},
					"network_name": schema.StringAttribute{
						MarkdownDescription: `Name of the network whose VPN exclusion rules are returned.`,
						Computed:            true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"custom": schema.SetNestedAttribute{
						MarkdownDescription: `Custom VPN exclusion rules. Pass an empty array to clear existing rules.`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"destination": schema.StringAttribute{
									MarkdownDescription: `Destination address; hostname required for DNS, IPv4 otherwise.`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"port": schema.StringAttribute{
									MarkdownDescription: `Destination port.`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"protocol": schema.StringAttribute{
									MarkdownDescription: `Protocol.
                                              Allowed values: [any,dns,icmp,tcp,udp]`,
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
							},
						},
					},
					"major_applications": schema.SetNestedAttribute{
						MarkdownDescription: `Major Application based VPN exclusion rules. Pass an empty array to clear existing rules.`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `Application's Meraki ID.`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Application's name.
                                              Allowed values: [AWS,Box,Office 365 Sharepoint,Office 365 Suite,Oracle,SAP,Salesforce,Skype & Teams,Slack,Webex,Webex Calling,Webex Meetings,Zoom]`,
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
func (r *NetworksApplianceTrafficShapingVpnExclusionsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceTrafficShapingVpnExclusions

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
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp1, err := r.client.Appliance.UpdateNetworkApplianceTrafficShapingVpnExclusions(vvNetworkID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceTrafficShapingVpnExclusions",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceTrafficShapingVpnExclusions",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceTrafficShapingVpnExclusionsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksApplianceTrafficShapingVpnExclusionsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksApplianceTrafficShapingVpnExclusionsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceTrafficShapingVpnExclusions struct {
	NetworkID  types.String                                                         `tfsdk:"network_id"`
	Item       *ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusions  `tfsdk:"item"`
	Parameters *RequestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsRs `tfsdk:"parameters"`
}

type ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusions struct {
	Custom            *[]ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsCustom            `tfsdk:"custom"`
	MajorApplications *[]ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsMajorApplications `tfsdk:"major_applications"`
	NetworkID         types.String                                                                           `tfsdk:"network_id"`
	NetworkName       types.String                                                                           `tfsdk:"network_name"`
}

type ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsCustom struct {
	Destination types.String `tfsdk:"destination"`
	Port        types.String `tfsdk:"port"`
	Protocol    types.String `tfsdk:"protocol"`
}

type ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsMajorApplications struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type RequestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsRs struct {
	Custom            *[]RequestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsCustomRs            `tfsdk:"custom"`
	MajorApplications *[]RequestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsMajorApplicationsRs `tfsdk:"major_applications"`
}

type RequestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsCustomRs struct {
	Destination types.String `tfsdk:"destination"`
	Port        types.String `tfsdk:"port"`
	Protocol    types.String `tfsdk:"protocol"`
}

type RequestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsMajorApplicationsRs struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// FromBody
func (r *NetworksApplianceTrafficShapingVpnExclusions) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusions {
	re := *r.Parameters
	var requestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsCustom []merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsCustom

	if re.Custom != nil {
		for _, rItem1 := range *re.Custom {
			destination := rItem1.Destination.ValueString()
			port := rItem1.Port.ValueString()
			protocol := rItem1.Protocol.ValueString()
			requestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsCustom = append(requestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsCustom, merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsCustom{
				Destination: destination,
				Port:        port,
				Protocol:    protocol,
			})
			//[debug] Is Array: True
		}
	}
	var requestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsMajorApplications []merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsMajorApplications

	if re.MajorApplications != nil {
		for _, rItem1 := range *re.MajorApplications {
			id := rItem1.ID.ValueString()
			name := rItem1.Name.ValueString()
			requestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsMajorApplications = append(requestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsMajorApplications, merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsMajorApplications{
				ID:   id,
				Name: name,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusions{
		Custom: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsCustom {
			if len(requestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsCustom) > 0 {
				return &requestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsCustom
			}
			return nil
		}(),
		MajorApplications: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsMajorApplications {
			if len(requestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsMajorApplications) > 0 {
				return &requestApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsMajorApplications
			}
			return nil
		}(),
	}
	return &out
}

// ToBody
func ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsItemToBody(state NetworksApplianceTrafficShapingVpnExclusions, response *merakigosdk.ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusions) NetworksApplianceTrafficShapingVpnExclusions {
	itemState := ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusions{
		Custom: func() *[]ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsCustom {
			if response.Custom != nil {
				result := make([]ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsCustom, len(*response.Custom))
				for i, custom := range *response.Custom {
					result[i] = ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsCustom{
						Destination: types.StringValue(custom.Destination),
						Port:        types.StringValue(custom.Port),
						Protocol:    types.StringValue(custom.Protocol),
					}
				}
				return &result
			}
			return nil
		}(),
		MajorApplications: func() *[]ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsMajorApplications {
			if response.MajorApplications != nil {
				result := make([]ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsMajorApplications, len(*response.MajorApplications))
				for i, majorApplications := range *response.MajorApplications {
					result[i] = ResponseApplianceUpdateNetworkApplianceTrafficShapingVpnExclusionsMajorApplications{
						ID:   types.StringValue(majorApplications.ID),
						Name: types.StringValue(majorApplications.Name),
					}
				}
				return &result
			}
			return nil
		}(),
		NetworkID:   types.StringValue(response.NetworkID),
		NetworkName: types.StringValue(response.NetworkName),
	}
	state.Item = &itemState
	return state
}
