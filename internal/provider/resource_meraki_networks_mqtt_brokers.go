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
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"authentication": schema.SingleNestedAttribute{
						MarkdownDescription: `Authentication settings of the MQTT broker`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"username": schema.StringAttribute{
								MarkdownDescription: `Username for the MQTT broker.`,
								Computed:            true,
							},
						},
					},
					"host": schema.StringAttribute{
						MarkdownDescription: `Host name/IP address where the MQTT broker runs.`,
						Computed:            true,
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `ID of the MQTT Broker.`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of the MQTT Broker.`,
						Computed:            true,
					},
					"port": schema.Int64Attribute{
						MarkdownDescription: `Host port though which the MQTT broker can be reached.`,
						Computed:            true,
					},
					"security": schema.SingleNestedAttribute{
						MarkdownDescription: `Security settings of the MQTT broker.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"mode": schema.StringAttribute{
								MarkdownDescription: `Security protocol of the MQTT broker.`,
								Computed:            true,
							},
							"tls": schema.SingleNestedAttribute{
								MarkdownDescription: `TLS settings of the MQTT broker.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"has_ca_certificate": schema.BoolAttribute{
										MarkdownDescription: `Indicates whether the CA certificate is set`,
										Computed:            true,
									},
									"verify_hostnames": schema.BoolAttribute{
										MarkdownDescription: `Whether the TLS hostname verification is enabled for the MQTT broker.`,
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"authentication": schema.SingleNestedAttribute{
						MarkdownDescription: `Authentication settings of the MQTT broker`,
						Optional:            true,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"password": schema.StringAttribute{
								MarkdownDescription: `Password for the MQTT broker.`,
								Optional:            true,
								Sensitive:           true,
								Computed:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
							},
							"username": schema.StringAttribute{
								MarkdownDescription: `Username for the MQTT broker.`,
								Optional:            true,
								Computed:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
							},
						},
					},
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
							"tls": schema.SingleNestedAttribute{
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
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Networks.CreateNetworkMqttBroker(vvNetworkID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkMqttBroker",
				restyResp1.String(),
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
	data = ResponseNetworksCreateNetworkMqttBrokerItemToBody(data, response)
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
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksMqttBrokers struct {
	NetworkID  types.String                              `tfsdk:"network_id"`
	Item       *ResponseNetworksCreateNetworkMqttBroker  `tfsdk:"item"`
	Parameters *RequestNetworksCreateNetworkMqttBrokerRs `tfsdk:"parameters"`
}

type ResponseNetworksCreateNetworkMqttBroker struct {
	Authentication *ResponseNetworksCreateNetworkMqttBrokerAuthentication `tfsdk:"authentication"`
	Host           types.String                                           `tfsdk:"host"`
	ID             types.String                                           `tfsdk:"id"`
	Name           types.String                                           `tfsdk:"name"`
	Port           types.Int64                                            `tfsdk:"port"`
	Security       *ResponseNetworksCreateNetworkMqttBrokerSecurity       `tfsdk:"security"`
}

type ResponseNetworksCreateNetworkMqttBrokerAuthentication struct {
	Username types.String `tfsdk:"username"`
}

type ResponseNetworksCreateNetworkMqttBrokerSecurity struct {
	Mode types.String                                        `tfsdk:"mode"`
	Tls  *ResponseNetworksCreateNetworkMqttBrokerSecurityTls `tfsdk:"tls"`
}

type ResponseNetworksCreateNetworkMqttBrokerSecurityTls struct {
	HasCaCertificate types.Bool `tfsdk:"has_ca_certificate"`
	VerifyHostnames  types.Bool `tfsdk:"verify_hostnames"`
}

type RequestNetworksCreateNetworkMqttBrokerRs struct {
	Authentication *RequestNetworksCreateNetworkMqttBrokerAuthenticationRs `tfsdk:"authentication"`
	Host           types.String                                            `tfsdk:"host"`
	Name           types.String                                            `tfsdk:"name"`
	Port           types.Int64                                             `tfsdk:"port"`
	Security       *RequestNetworksCreateNetworkMqttBrokerSecurityRs       `tfsdk:"security"`
}

type RequestNetworksCreateNetworkMqttBrokerAuthenticationRs struct {
	Password types.String `tfsdk:"password"`
	Username types.String `tfsdk:"username"`
}

type RequestNetworksCreateNetworkMqttBrokerSecurityRs struct {
	Mode types.String                                         `tfsdk:"mode"`
	Tls  *RequestNetworksCreateNetworkMqttBrokerSecurityTlsRs `tfsdk:"tls"`
}

type RequestNetworksCreateNetworkMqttBrokerSecurityTlsRs struct {
	CaCertificate   types.String `tfsdk:"ca_certificate"`
	VerifyHostnames types.Bool   `tfsdk:"verify_hostnames"`
}

// FromBody
func (r *NetworksMqttBrokers) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksCreateNetworkMqttBroker {
	emptyString := ""
	re := *r.Parameters
	var requestNetworksCreateNetworkMqttBrokerAuthentication *merakigosdk.RequestNetworksCreateNetworkMqttBrokerAuthentication

	if re.Authentication != nil {
		password := re.Authentication.Password.ValueString()
		username := re.Authentication.Username.ValueString()
		requestNetworksCreateNetworkMqttBrokerAuthentication = &merakigosdk.RequestNetworksCreateNetworkMqttBrokerAuthentication{
			Password: password,
			Username: username,
		}
		//[debug] Is Array: False
	}
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
		var requestNetworksCreateNetworkMqttBrokerSecurityTls *merakigosdk.RequestNetworksCreateNetworkMqttBrokerSecurityTls

		if re.Security.Tls != nil {
			caCertificate := re.Security.Tls.CaCertificate.ValueString()
			verifyHostnames := func() *bool {
				if !re.Security.Tls.VerifyHostnames.IsUnknown() && !re.Security.Tls.VerifyHostnames.IsNull() {
					return re.Security.Tls.VerifyHostnames.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksCreateNetworkMqttBrokerSecurityTls = &merakigosdk.RequestNetworksCreateNetworkMqttBrokerSecurityTls{
				CaCertificate:   caCertificate,
				VerifyHostnames: verifyHostnames,
			}
			//[debug] Is Array: False
		}
		requestNetworksCreateNetworkMqttBrokerSecurity = &merakigosdk.RequestNetworksCreateNetworkMqttBrokerSecurity{
			Mode: mode,
			Tls:  requestNetworksCreateNetworkMqttBrokerSecurityTls,
		}
		//[debug] Is Array: False
	}
	out := merakigosdk.RequestNetworksCreateNetworkMqttBroker{
		Authentication: requestNetworksCreateNetworkMqttBrokerAuthentication,
		Host:           *host,
		Name:           *name,
		Port:           int64ToIntPointer(port),
		Security:       requestNetworksCreateNetworkMqttBrokerSecurity,
	}
	return &out
}

// ToBody
func ResponseNetworksCreateNetworkMqttBrokerItemToBody(state NetworksMqttBrokers, response *merakigosdk.ResponseNetworksCreateNetworkMqttBroker) NetworksMqttBrokers {
	itemState := ResponseNetworksCreateNetworkMqttBroker{
		Authentication: func() *ResponseNetworksCreateNetworkMqttBrokerAuthentication {
			if response.Authentication != nil {
				return &ResponseNetworksCreateNetworkMqttBrokerAuthentication{
					Username: types.StringValue(response.Authentication.Username),
				}
			}
			return nil
		}(),
		Host: types.StringValue(response.Host),
		ID:   types.StringValue(response.ID),
		Name: types.StringValue(response.Name),
		Port: func() types.Int64 {
			if response.Port != nil {
				return types.Int64Value(int64(*response.Port))
			}
			return types.Int64{}
		}(),
		Security: func() *ResponseNetworksCreateNetworkMqttBrokerSecurity {
			if response.Security != nil {
				return &ResponseNetworksCreateNetworkMqttBrokerSecurity{
					Mode: types.StringValue(response.Security.Mode),
					Tls: func() *ResponseNetworksCreateNetworkMqttBrokerSecurityTls {
						if response.Security.Tls != nil {
							return &ResponseNetworksCreateNetworkMqttBrokerSecurityTls{
								HasCaCertificate: func() types.Bool {
									if response.Security.Tls.HasCaCertificate != nil {
										return types.BoolValue(*response.Security.Tls.HasCaCertificate)
									}
									return types.Bool{}
								}(),
								VerifyHostnames: func() types.Bool {
									if response.Security.Tls.VerifyHostnames != nil {
										return types.BoolValue(*response.Security.Tls.VerifyHostnames)
									}
									return types.Bool{}
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
