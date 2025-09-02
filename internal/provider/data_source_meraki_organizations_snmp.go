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
	_ datasource.DataSource              = &OrganizationsSNMPDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSNMPDataSource{}
)

func NewOrganizationsSNMPDataSource() datasource.DataSource {
	return &OrganizationsSNMPDataSource{}
}

type OrganizationsSNMPDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSNMPDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSNMPDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_snmp"
}

func (d *OrganizationsSNMPDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"hostname": schema.StringAttribute{
						MarkdownDescription: `The hostname of the SNMP server.`,
						Computed:            true,
					},
					"peer_ips": schema.ListAttribute{
						MarkdownDescription: `The list of IPv4 addresses that are allowed to access the SNMP server.`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"port": schema.Int64Attribute{
						MarkdownDescription: `The port of the SNMP server.`,
						Computed:            true,
					},
					"v2_community_string": schema.StringAttribute{
						MarkdownDescription: `The community string for SNMP version 2c, if enabled.`,
						Computed:            true,
					},
					"v2c_enabled": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating whether SNMP version 2c is enabled for the organization.`,
						Computed:            true,
					},
					"v3_auth_mode": schema.StringAttribute{
						MarkdownDescription: `The SNMP version 3 authentication mode. Can be either 'MD5' or 'SHA'.`,
						Computed:            true,
					},
					"v3_enabled": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating whether SNMP version 3 is enabled for the organization.`,
						Computed:            true,
					},
					"v3_priv_mode": schema.StringAttribute{
						MarkdownDescription: `The SNMP version 3 privacy mode. Can be either 'DES' or 'AES128'.`,
						Computed:            true,
					},
					"v3_user": schema.StringAttribute{
						MarkdownDescription: `The user for SNMP version 3, if enabled.`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *OrganizationsSNMPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSNMP OrganizationsSNMP
	diags := req.Config.Get(ctx, &organizationsSNMP)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSNMP")
		vvOrganizationID := organizationsSNMP.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationSNMP(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSNMP",
				err.Error(),
			)
			return
		}

		organizationsSNMP = ResponseOrganizationsGetOrganizationSNMPItemToBody(organizationsSNMP, response1)
		diags = resp.State.Set(ctx, &organizationsSNMP)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSNMP struct {
	OrganizationID types.String                              `tfsdk:"organization_id"`
	Item           *ResponseOrganizationsGetOrganizationSnmp `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationSnmp struct {
	Hostname          types.String `tfsdk:"hostname"`
	PeerIPs           types.List   `tfsdk:"peer_ips"`
	Port              types.Int64  `tfsdk:"port"`
	V2CommunityString types.String `tfsdk:"v2_community_string"`
	V2CEnabled        types.Bool   `tfsdk:"v2c_enabled"`
	V3AuthMode        types.String `tfsdk:"v3_auth_mode"`
	V3Enabled         types.Bool   `tfsdk:"v3_enabled"`
	V3PrivMode        types.String `tfsdk:"v3_priv_mode"`
	V3User            types.String `tfsdk:"v3_user"`
}

// ToBody
func ResponseOrganizationsGetOrganizationSNMPItemToBody(state OrganizationsSNMP, response *merakigosdk.ResponseOrganizationsGetOrganizationSNMP) OrganizationsSNMP {
	itemState := ResponseOrganizationsGetOrganizationSnmp{
		Hostname: func() types.String {
			if response.Hostname != "" {
				return types.StringValue(response.Hostname)
			}
			return types.String{}
		}(),
		PeerIPs: StringSliceToList(response.PeerIPs),
		Port: func() types.Int64 {
			if response.Port != nil {
				return types.Int64Value(int64(*response.Port))
			}
			return types.Int64{}
		}(),
		V2CommunityString: func() types.String {
			if response.V2CommunityString != "" {
				return types.StringValue(response.V2CommunityString)
			}
			return types.String{}
		}(),
		V2CEnabled: func() types.Bool {
			if response.V2CEnabled != nil {
				return types.BoolValue(*response.V2CEnabled)
			}
			return types.Bool{}
		}(),
		V3AuthMode: func() types.String {
			if response.V3AuthMode != "" {
				return types.StringValue(response.V3AuthMode)
			}
			return types.String{}
		}(),
		V3Enabled: func() types.Bool {
			if response.V3Enabled != nil {
				return types.BoolValue(*response.V3Enabled)
			}
			return types.Bool{}
		}(),
		V3PrivMode: func() types.String {
			if response.V3PrivMode != "" {
				return types.StringValue(response.V3PrivMode)
			}
			return types.String{}
		}(),
		V3User: func() types.String {
			if response.V3User != "" {
				return types.StringValue(response.V3User)
			}
			return types.String{}
		}(),
	}
	state.Item = &itemState
	return state
}
