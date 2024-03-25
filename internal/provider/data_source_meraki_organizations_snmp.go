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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

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
						Computed: true,
					},
					"peer_ips": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"port": schema.Int64Attribute{
						Computed: true,
					},
					"v2c_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"v3_auth_mode": schema.StringAttribute{
						Computed: true,
					},
					"v3_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"v3_priv_mode": schema.StringAttribute{
						Computed: true,
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
	Hostname   types.String `tfsdk:"hostname"`
	PeerIPs    types.List   `tfsdk:"peer_ips"`
	Port       types.Int64  `tfsdk:"port"`
	V2CEnabled types.Bool   `tfsdk:"v2c_enabled"`
	V3AuthMode types.String `tfsdk:"v3_auth_mode"`
	V3Enabled  types.Bool   `tfsdk:"v3_enabled"`
	V3PrivMode types.String `tfsdk:"v3_priv_mode"`
}

// ToBody
func ResponseOrganizationsGetOrganizationSNMPItemToBody(state OrganizationsSNMP, response *merakigosdk.ResponseOrganizationsGetOrganizationSNMP) OrganizationsSNMP {
	itemState := ResponseOrganizationsGetOrganizationSnmp{
		Hostname: types.StringValue(response.Hostname),
		PeerIPs:  StringSliceToList(response.PeerIPs),
		Port: func() types.Int64 {
			if response.Port != nil {
				return types.Int64Value(int64(*response.Port))
			}
			return types.Int64{}
		}(),
		V2CEnabled: func() types.Bool {
			if response.V2CEnabled != nil {
				return types.BoolValue(*response.V2CEnabled)
			}
			return types.Bool{}
		}(),
		V3AuthMode: types.StringValue(response.V3AuthMode),
		V3Enabled: func() types.Bool {
			if response.V3Enabled != nil {
				return types.BoolValue(*response.V3Enabled)
			}
			return types.Bool{}
		}(),
		V3PrivMode: types.StringValue(response.V3PrivMode),
	}
	state.Item = &itemState
	return state
}
