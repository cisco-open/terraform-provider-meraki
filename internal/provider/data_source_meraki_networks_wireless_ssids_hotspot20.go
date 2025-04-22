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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksWirelessSSIDsHotspot20DataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessSSIDsHotspot20DataSource{}
)

func NewNetworksWirelessSSIDsHotspot20DataSource() datasource.DataSource {
	return &NetworksWirelessSSIDsHotspot20DataSource{}
}

type NetworksWirelessSSIDsHotspot20DataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessSSIDsHotspot20DataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessSSIDsHotspot20DataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_hotspot20"
}

func (d *NetworksWirelessSSIDsHotspot20DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: `number path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"domains": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"enabled": schema.BoolAttribute{
						Computed: true,
					},
					"mcc_mncs": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"mcc": schema.StringAttribute{
									Computed: true,
								},
								"mnc": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
					"nai_realms": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"format": schema.StringAttribute{
									Computed: true,
								},
								"methods": schema.SetNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"authentication_types": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{

													"credentials": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
													},
													"eapinner_authentication": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
													},
													"non_eapinner_authentication": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
													},
													"tunneled_eap_method_credentials": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
													},
												},
											},
											"id": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
					"network_access_type": schema.StringAttribute{
						Computed: true,
					},
					"operator": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"name": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"roam_consort_ois": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"venue": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"name": schema.StringAttribute{
								Computed: true,
							},
							"type": schema.StringAttribute{
								Computed: true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessSSIDsHotspot20DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessSSIDsHotspot20 NetworksWirelessSSIDsHotspot20
	diags := req.Config.Get(ctx, &networksWirelessSSIDsHotspot20)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessSSIDHotspot20")
		vvNetworkID := networksWirelessSSIDsHotspot20.NetworkID.ValueString()
		vvNumber := networksWirelessSSIDsHotspot20.Number.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessSSIDHotspot20(vvNetworkID, vvNumber)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDHotspot20",
				err.Error(),
			)
			return
		}

		networksWirelessSSIDsHotspot20 = ResponseWirelessGetNetworkWirelessSSIDHotspot20ItemToBody(networksWirelessSSIDsHotspot20, response1)
		diags = resp.State.Set(ctx, &networksWirelessSSIDsHotspot20)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessSSIDsHotspot20 struct {
	NetworkID types.String                                     `tfsdk:"network_id"`
	Number    types.String                                     `tfsdk:"number"`
	Item      *ResponseWirelessGetNetworkWirelessSsidHotspot20 `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessSsidHotspot20 struct {
	Domains           types.List                                                  `tfsdk:"domains"`
	Enabled           types.Bool                                                  `tfsdk:"enabled"`
	MccMncs           *[]ResponseWirelessGetNetworkWirelessSsidHotspot20MccMncs   `tfsdk:"mcc_mncs"`
	NaiRealms         *[]ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealms `tfsdk:"nai_realms"`
	NetworkAccessType types.String                                                `tfsdk:"network_access_type"`
	Operator          *ResponseWirelessGetNetworkWirelessSsidHotspot20Operator    `tfsdk:"operator"`
	RoamConsortOis    types.List                                                  `tfsdk:"roam_consort_ois"`
	Venue             *ResponseWirelessGetNetworkWirelessSsidHotspot20Venue       `tfsdk:"venue"`
}

type ResponseWirelessGetNetworkWirelessSsidHotspot20MccMncs struct {
	Mcc types.String `tfsdk:"mcc"`
	Mnc types.String `tfsdk:"mnc"`
}

type ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealms struct {
	Format  types.String                                                       `tfsdk:"format"`
	Methods *[]ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsMethods `tfsdk:"methods"`
	Name    types.String                                                       `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsMethods struct {
	AuthenticationTypes *ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsMethodsAuthenticationTypes `tfsdk:"authentication_types"`
	ID                  types.String                                                                        `tfsdk:"id"`
}

type ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsMethodsAuthenticationTypes struct {
	Credentials                  types.List `tfsdk:"credentials"`
	EapinnerAuthentication       types.List `tfsdk:"eap_inner_authentication"`
	NonEapinnerAuthentication    types.List `tfsdk:"non_eap_inner_authentication"`
	TunneledEapMethodCredentials types.List `tfsdk:"tunneled_eap_method_credentials"`
}

type ResponseWirelessGetNetworkWirelessSsidHotspot20Operator struct {
	Name types.String `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessSsidHotspot20Venue struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessSSIDHotspot20ItemToBody(state NetworksWirelessSSIDsHotspot20, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDHotspot20) NetworksWirelessSSIDsHotspot20 {
	itemState := ResponseWirelessGetNetworkWirelessSsidHotspot20{
		Domains: StringSliceToList(response.Domains),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		MccMncs: func() *[]ResponseWirelessGetNetworkWirelessSsidHotspot20MccMncs {
			if response.MccMncs != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidHotspot20MccMncs, len(*response.MccMncs))
				for i, mccMncs := range *response.MccMncs {
					result[i] = ResponseWirelessGetNetworkWirelessSsidHotspot20MccMncs{
						Mcc: types.StringValue(mccMncs.Mcc),
						Mnc: types.StringValue(mccMncs.Mnc),
					}
				}
				return &result
			}
			return nil
		}(),
		NaiRealms: func() *[]ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealms {
			if response.NaiRealms != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealms, len(*response.NaiRealms))
				for i, naiRealms := range *response.NaiRealms {
					result[i] = ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealms{
						Format: types.StringValue(naiRealms.Format),
						Methods: func() *[]ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsMethods {
							if naiRealms.Methods != nil {
								result := make([]ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsMethods, len(*naiRealms.Methods))
								for i, methods := range *naiRealms.Methods {
									result[i] = ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsMethods{
										AuthenticationTypes: func() *ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsMethodsAuthenticationTypes {
											if methods.AuthenticationTypes != nil {
												return &ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsMethodsAuthenticationTypes{
													Credentials:                  StringSliceToList(methods.AuthenticationTypes.Credentials),
													EapinnerAuthentication:       StringSliceToList(methods.AuthenticationTypes.EapinnerAuthentication),
													NonEapinnerAuthentication:    StringSliceToList(methods.AuthenticationTypes.NonEapinnerAuthentication),
													TunneledEapMethodCredentials: StringSliceToList(methods.AuthenticationTypes.TunneledEapMethodCredentials),
												}
											}
											return nil
										}(),
										ID: types.StringValue(methods.ID),
									}
								}
								return &result
							}
							return nil
						}(),
						Name: types.StringValue(naiRealms.Name),
					}
				}
				return &result
			}
			return nil
		}(),
		NetworkAccessType: types.StringValue(response.NetworkAccessType),
		Operator: func() *ResponseWirelessGetNetworkWirelessSsidHotspot20Operator {
			if response.Operator != nil {
				return &ResponseWirelessGetNetworkWirelessSsidHotspot20Operator{
					Name: types.StringValue(response.Operator.Name),
				}
			}
			return nil
		}(),
		RoamConsortOis: StringSliceToList(response.RoamConsortOis),
		Venue: func() *ResponseWirelessGetNetworkWirelessSsidHotspot20Venue {
			if response.Venue != nil {
				return &ResponseWirelessGetNetworkWirelessSsidHotspot20Venue{
					Name: types.StringValue(response.Venue.Name),
					Type: types.StringValue(response.Venue.Type),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
