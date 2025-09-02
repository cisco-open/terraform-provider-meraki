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
	_ datasource.DataSource              = &NetworksSmDevicesCertsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSmDevicesCertsDataSource{}
)

func NewNetworksSmDevicesCertsDataSource() datasource.DataSource {
	return &NetworksSmDevicesCertsDataSource{}
}

type NetworksSmDevicesCertsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSmDevicesCertsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSmDevicesCertsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_devices_certs"
}

func (d *NetworksSmDevicesCertsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"device_id": schema.StringAttribute{
				MarkdownDescription: `deviceId path parameter. Device ID`,
				Required:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSmGetNetworkSmDeviceCerts`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"cert_pem": schema.StringAttribute{
							MarkdownDescription: `The PEM of the certificate.`,
							Computed:            true,
						},
						"device_id": schema.StringAttribute{
							MarkdownDescription: `The Meraki managed device Id.`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `The Meraki Id of the certificate record.`,
							Computed:            true,
						},
						"issuer": schema.StringAttribute{
							MarkdownDescription: `The certificate issuer.`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the certificate.`,
							Computed:            true,
						},
						"not_valid_after": schema.StringAttribute{
							MarkdownDescription: `The date after which the certificate is no longer valid.`,
							Computed:            true,
						},
						"not_valid_before": schema.StringAttribute{
							MarkdownDescription: `The date before which the certificate is not valid.`,
							Computed:            true,
						},
						"subject": schema.StringAttribute{
							MarkdownDescription: `The subject of the certificate.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSmDevicesCertsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSmDevicesCerts NetworksSmDevicesCerts
	diags := req.Config.Get(ctx, &networksSmDevicesCerts)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSmDeviceCerts")
		vvNetworkID := networksSmDevicesCerts.NetworkID.ValueString()
		vvDeviceID := networksSmDevicesCerts.DeviceID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Sm.GetNetworkSmDeviceCerts(vvNetworkID, vvDeviceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmDeviceCerts",
				err.Error(),
			)
			return
		}

		networksSmDevicesCerts = ResponseSmGetNetworkSmDeviceCertsItemsToBody(networksSmDevicesCerts, response1)
		diags = resp.State.Set(ctx, &networksSmDevicesCerts)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSmDevicesCerts struct {
	NetworkID types.String                             `tfsdk:"network_id"`
	DeviceID  types.String                             `tfsdk:"device_id"`
	Items     *[]ResponseItemSmGetNetworkSmDeviceCerts `tfsdk:"items"`
}

type ResponseItemSmGetNetworkSmDeviceCerts struct {
	CertPem        types.String `tfsdk:"cert_pem"`
	DeviceID       types.String `tfsdk:"device_id"`
	ID             types.String `tfsdk:"id"`
	Issuer         types.String `tfsdk:"issuer"`
	Name           types.String `tfsdk:"name"`
	NotValidAfter  types.String `tfsdk:"not_valid_after"`
	NotValidBefore types.String `tfsdk:"not_valid_before"`
	Subject        types.String `tfsdk:"subject"`
}

// ToBody
func ResponseSmGetNetworkSmDeviceCertsItemsToBody(state NetworksSmDevicesCerts, response *merakigosdk.ResponseSmGetNetworkSmDeviceCerts) NetworksSmDevicesCerts {
	var items []ResponseItemSmGetNetworkSmDeviceCerts
	for _, item := range *response {
		itemState := ResponseItemSmGetNetworkSmDeviceCerts{
			CertPem: func() types.String {
				if item.CertPem != "" {
					return types.StringValue(item.CertPem)
				}
				return types.String{}
			}(),
			DeviceID: func() types.String {
				if item.DeviceID != "" {
					return types.StringValue(item.DeviceID)
				}
				return types.String{}
			}(),
			ID: func() types.String {
				if item.ID != "" {
					return types.StringValue(item.ID)
				}
				return types.String{}
			}(),
			Issuer: func() types.String {
				if item.Issuer != "" {
					return types.StringValue(item.Issuer)
				}
				return types.String{}
			}(),
			Name: func() types.String {
				if item.Name != "" {
					return types.StringValue(item.Name)
				}
				return types.String{}
			}(),
			NotValidAfter: func() types.String {
				if item.NotValidAfter != "" {
					return types.StringValue(item.NotValidAfter)
				}
				return types.String{}
			}(),
			NotValidBefore: func() types.String {
				if item.NotValidBefore != "" {
					return types.StringValue(item.NotValidBefore)
				}
				return types.String{}
			}(),
			Subject: func() types.String {
				if item.Subject != "" {
					return types.StringValue(item.Subject)
				}
				return types.String{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
