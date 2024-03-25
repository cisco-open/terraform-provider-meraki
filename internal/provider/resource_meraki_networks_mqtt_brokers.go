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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksMqttBrokersResource{}
	_ resource.ResourceWithConfigure = &NetworksMqttBrokersResource{}
)

func NewNetworksMqttBrokersResource() resource.Resource {
	return &NetworksMqttBrokersResource{}
}

type NetworksMqttBrokersResource struct {
	client *merakigosdk.Client
}

func (r *NetworksMqttBrokersResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksMqttBrokersResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_mqtt_brokers"
}

// resourceAction
func (r *NetworksMqttBrokersResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{

			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"host": schema.StringAttribute{
						MarkdownDescription: `Host name/IP address where the MQTT broker runs.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of the MQTT broker.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"authentication": schema.StringAttribute{
						MarkdownDescription: `Name of the Auth.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"port": schema.Int64Attribute{
						MarkdownDescription: `Host port though which the MQTT broker can be reached.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplace(),
						},
					},
					"security": schema.SingleNestedAttribute{
						MarkdownDescription: `Security settings of the MQTT broker.`,
						Optional:            true,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"mode": schema.StringAttribute{
								MarkdownDescription: `Security protocol of the MQTT broker.`,
								Optional:            true,
								Computed:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
							},
							"security": schema.SingleNestedAttribute{
								MarkdownDescription: `TLS settings of the MQTT broker.`,
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"ca_certificate": schema.StringAttribute{
										MarkdownDescription: `CA Certificate of the MQTT broker.`,
										Optional:            true,
										Computed:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
									"verify_hostnames": schema.BoolAttribute{
										MarkdownDescription: `Whether the TLS hostname verification is enabled for the MQTT broker.`,
										Optional:            true,
										Computed:            true,
										PlanModifiers: []planmodifier.Bool{
											boolplanmodifier.RequiresReplace(),
										},
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
func (r *NetworksMqttBrokersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksMqttBrokers

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
	dataRequest := data.toSdkApiRequestCreate(ctx)
	restyResp1, err := r.client.Networks.CreateNetworkMqttBroker(vvNetworkID, dataRequest)

	if err != nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkMqttBroker",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkMqttBroker",
			err.Error(),
		)
		return
	}
	//Item

	// data2 := ResponseNetworksCreateNetworkMqttBroker(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksMqttBrokersResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksMqttBrokersResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksMqttBrokersResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksMqttBrokers struct {
	NetworkID  types.String                              `tfsdk:"network_id"`
	Parameters *RequestNetworksCreateNetworkMqttBrokerRs `tfsdk:"parameters"`
}

type RequestNetworksCreateNetworkMqttBrokerRs struct {
	Authentication types.String                                      `tfsdk:"authentication"`
	Host           types.String                                      `tfsdk:"host"`
	Name           types.String                                      `tfsdk:"name"`
	Port           types.Int64                                       `tfsdk:"port"`
	Security       *RequestNetworksCreateNetworkMqttBrokerSecurityRs `tfsdk:"security"`
}

// type RequestNetworksCreateNetworkMqttBrokerAuthenticationRs interface{}

type RequestNetworksCreateNetworkMqttBrokerSecurityRs struct {
	Mode     types.String                                              `tfsdk:"mode"`
	Security *RequestNetworksCreateNetworkMqttBrokerSecuritySecurityRs `tfsdk:"security"`
}

type RequestNetworksCreateNetworkMqttBrokerSecuritySecurityRs struct {
	CaCertificate   types.String `tfsdk:"ca_certificate"`
	VerifyHostnames types.Bool   `tfsdk:"verify_hostnames"`
}

// FromBody
func (r *NetworksMqttBrokers) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksCreateNetworkMqttBroker {
	emptyString := ""
	re := *r.Parameters
	// var requestNetworksCreateNetworkMqttBrokerAuthentication2 merakigosdk.RequestNetworksCreateNetworkMqttBrokerAuthentication
	// requestApplianceUpdateNetworkApplianceVLANFixedIPAssignments := r.FixedIPAssignments.ValueString()
	// var intf interface{} = requestApplianceUpdateNetworkApplianceVLANFixedIPAssignments
	// requestApplianceUpdateNetworkApplianceVLANFixedIPAssignments2, ok := intf.(merakigosdk.RequestApplianceUpdateNetworkApplianceVLANFixedIPAssignments)
	// if !ok {
	// 	requestApplianceUpdateNetworkApplianceVLANFixedIPAssignments2 = nil
	// }
	// requestNetworksCreateNetworkMqttBrokerAuthentication := re.Authentication.ValueString()
	// var intf interface{} = requestNetworksCreateNetworkMqttBrokerAuthentication
	// requestNetworksCreateNetworkMqttBrokerAuthentication2, ok := intf.(merakigosdk.RequestNetworksCreateNetworkMqttBrokerAuthentication)
	// if !ok {
	// 	requestNetworksCreateNetworkMqttBrokerAuthentication2 = nil
	// }
	host := new(string)
	if !re.Host.IsUnknown() && !re.Host.IsNull() {
		*host = re.Host.ValueString()
	} else {
		host = &emptyString
	}
	name := new(string)
	if !re.Name.IsUnknown() && !re.Name.IsNull() {
		*name = re.Name.ValueString()
	} else {
		name = &emptyString
	}
	port := new(int64)
	if !re.Port.IsUnknown() && !re.Port.IsNull() {
		*port = re.Port.ValueInt64()
	} else {
		port = nil
	}
	var requestNetworksCreateNetworkMqttBrokerSecurity *merakigosdk.RequestNetworksCreateNetworkMqttBrokerSecurity
	if re.Security != nil {
		mode := re.Security.Mode.ValueString()
		var requestNetworksCreateNetworkMqttBrokerSecuritySecurity *merakigosdk.RequestNetworksCreateNetworkMqttBrokerSecuritySecurity
		if re.Security.Security != nil {
			caCertificate := re.Security.Security.CaCertificate.ValueString()
			verifyHostnames := func() *bool {
				if !re.Security.Security.VerifyHostnames.IsUnknown() && !re.Security.Security.VerifyHostnames.IsNull() {
					return re.Security.Security.VerifyHostnames.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksCreateNetworkMqttBrokerSecuritySecurity = &merakigosdk.RequestNetworksCreateNetworkMqttBrokerSecuritySecurity{
				CaCertificate:   caCertificate,
				VerifyHostnames: verifyHostnames,
			}
		}
		requestNetworksCreateNetworkMqttBrokerSecurity = &merakigosdk.RequestNetworksCreateNetworkMqttBrokerSecurity{
			Mode:     mode,
			Security: requestNetworksCreateNetworkMqttBrokerSecuritySecurity,
		}
	}
	out := merakigosdk.RequestNetworksCreateNetworkMqttBroker{
		// Authentication: func() *[]merakigosdk.RequestNetworksCreateNetworkMqttBrokerAuthentication2 {
		// 	if len(requestNetworksCreateNetworkMqttBrokerAuthentication2) > 0 {
		// 		return &requestNetworksCreateNetworkMqttBrokerAuthentication2
		// 	}
		// 	return nil
		// }(),
		Host:     *host,
		Name:     *name,
		Port:     int64ToIntPointer(port),
		Security: requestNetworksCreateNetworkMqttBrokerSecurity,
	}
	return &out
}

//ToBody
