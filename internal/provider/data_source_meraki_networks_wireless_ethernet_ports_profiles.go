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
	_ datasource.DataSource              = &NetworksWirelessEthernetPortsProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessEthernetPortsProfilesDataSource{}
)

func NewNetworksWirelessEthernetPortsProfilesDataSource() datasource.DataSource {
	return &NetworksWirelessEthernetPortsProfilesDataSource{}
}

type NetworksWirelessEthernetPortsProfilesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessEthernetPortsProfilesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessEthernetPortsProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ethernet_ports_profiles"
}

func (d *NetworksWirelessEthernetPortsProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"profile_id": schema.StringAttribute{
				MarkdownDescription: `profileId path parameter. Profile ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"is_default": schema.BoolAttribute{
						MarkdownDescription: `Is default profile`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `AP port profile name`,
						Computed:            true,
					},
					"ports": schema.SetNestedAttribute{
						MarkdownDescription: `Ports config`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"enabled": schema.BoolAttribute{
									MarkdownDescription: `Enabled`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Name`,
									Computed:            true,
								},
								"number": schema.Int64Attribute{
									MarkdownDescription: `Number`,
									Computed:            true,
								},
								"psk_group_id": schema.StringAttribute{
									MarkdownDescription: `PSK Group number`,
									Computed:            true,
								},
								"ssid": schema.Int64Attribute{
									MarkdownDescription: `Ssid number`,
									Computed:            true,
								},
							},
						},
					},
					"profile_id": schema.StringAttribute{
						MarkdownDescription: `AP port profile ID`,
						Computed:            true,
					},
					"usb_ports": schema.SetNestedAttribute{
						MarkdownDescription: `Usb ports config`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"enabled": schema.BoolAttribute{
									MarkdownDescription: `Enabled`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Name`,
									Computed:            true,
								},
								"ssid": schema.Int64Attribute{
									MarkdownDescription: `Ssid number`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessEthernetPortsProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessEthernetPortsProfiles NetworksWirelessEthernetPortsProfiles
	diags := req.Config.Get(ctx, &networksWirelessEthernetPortsProfiles)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessEthernetPortsProfile")
		vvNetworkID := networksWirelessEthernetPortsProfiles.NetworkID.ValueString()
		vvProfileID := networksWirelessEthernetPortsProfiles.ProfileID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessEthernetPortsProfile(vvNetworkID, vvProfileID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessEthernetPortsProfile",
				err.Error(),
			)
			return
		}

		networksWirelessEthernetPortsProfiles = ResponseWirelessGetNetworkWirelessEthernetPortsProfileItemToBody(networksWirelessEthernetPortsProfiles, response1)
		diags = resp.State.Set(ctx, &networksWirelessEthernetPortsProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessEthernetPortsProfiles struct {
	NetworkID types.String                                            `tfsdk:"network_id"`
	ProfileID types.String                                            `tfsdk:"profile_id"`
	Item      *ResponseWirelessGetNetworkWirelessEthernetPortsProfile `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessEthernetPortsProfile struct {
	IsDefault types.Bool                                                        `tfsdk:"is_default"`
	Name      types.String                                                      `tfsdk:"name"`
	Ports     *[]ResponseWirelessGetNetworkWirelessEthernetPortsProfilePorts    `tfsdk:"ports"`
	ProfileID types.String                                                      `tfsdk:"profile_id"`
	UsbPorts  *[]ResponseWirelessGetNetworkWirelessEthernetPortsProfileUsbPorts `tfsdk:"usb_ports"`
}

type ResponseWirelessGetNetworkWirelessEthernetPortsProfilePorts struct {
	Enabled    types.Bool   `tfsdk:"enabled"`
	Name       types.String `tfsdk:"name"`
	Number     types.Int64  `tfsdk:"number"`
	PskGroupID types.String `tfsdk:"psk_group_id"`
	SSID       types.Int64  `tfsdk:"ssid"`
}

type ResponseWirelessGetNetworkWirelessEthernetPortsProfileUsbPorts struct {
	Enabled types.Bool   `tfsdk:"enabled"`
	Name    types.String `tfsdk:"name"`
	SSID    types.Int64  `tfsdk:"ssid"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessEthernetPortsProfileItemToBody(state NetworksWirelessEthernetPortsProfiles, response *merakigosdk.ResponseWirelessGetNetworkWirelessEthernetPortsProfile) NetworksWirelessEthernetPortsProfiles {
	itemState := ResponseWirelessGetNetworkWirelessEthernetPortsProfile{
		IsDefault: func() types.Bool {
			if response.IsDefault != nil {
				return types.BoolValue(*response.IsDefault)
			}
			return types.Bool{}
		}(),
		Name: types.StringValue(response.Name),
		Ports: func() *[]ResponseWirelessGetNetworkWirelessEthernetPortsProfilePorts {
			if response.Ports != nil {
				result := make([]ResponseWirelessGetNetworkWirelessEthernetPortsProfilePorts, len(*response.Ports))
				for i, ports := range *response.Ports {
					result[i] = ResponseWirelessGetNetworkWirelessEthernetPortsProfilePorts{
						Enabled: func() types.Bool {
							if ports.Enabled != nil {
								return types.BoolValue(*ports.Enabled)
							}
							return types.Bool{}
						}(),
						Name: types.StringValue(ports.Name),
						Number: func() types.Int64 {
							if ports.Number != nil {
								return types.Int64Value(int64(*ports.Number))
							}
							return types.Int64{}
						}(),
						PskGroupID: types.StringValue(ports.PskGroupID),
						SSID: func() types.Int64 {
							if ports.SSID != nil {
								return types.Int64Value(int64(*ports.SSID))
							}
							return types.Int64{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		ProfileID: types.StringValue(response.ProfileID),
		UsbPorts: func() *[]ResponseWirelessGetNetworkWirelessEthernetPortsProfileUsbPorts {
			if response.UsbPorts != nil {
				result := make([]ResponseWirelessGetNetworkWirelessEthernetPortsProfileUsbPorts, len(*response.UsbPorts))
				for i, usbPorts := range *response.UsbPorts {
					result[i] = ResponseWirelessGetNetworkWirelessEthernetPortsProfileUsbPorts{
						Enabled: func() types.Bool {
							if usbPorts.Enabled != nil {
								return types.BoolValue(*usbPorts.Enabled)
							}
							return types.Bool{}
						}(),
						Name: types.StringValue(usbPorts.Name),
						SSID: func() types.Int64 {
							if usbPorts.SSID != nil {
								return types.Int64Value(int64(*usbPorts.SSID))
							}
							return types.Int64{}
						}(),
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
