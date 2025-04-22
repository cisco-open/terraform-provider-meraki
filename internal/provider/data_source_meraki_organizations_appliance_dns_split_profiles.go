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
	_ datasource.DataSource              = &OrganizationsApplianceDNSSplitProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsApplianceDNSSplitProfilesDataSource{}
)

func NewOrganizationsApplianceDNSSplitProfilesDataSource() datasource.DataSource {
	return &OrganizationsApplianceDNSSplitProfilesDataSource{}
}

type OrganizationsApplianceDNSSplitProfilesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsApplianceDNSSplitProfilesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsApplianceDNSSplitProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_appliance_dns_split_profiles"
}

func (d *OrganizationsApplianceDNSSplitProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"profile_ids": schema.ListAttribute{
				MarkdownDescription: `profileIds query parameter. Optional parameter to filter the results by profile IDs`,
				Optional:            true,
				ElementType:         types.StringType,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseApplianceGetOrganizationApplianceDnsSplitProfiles`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"hostnames": schema.ListAttribute{
							MarkdownDescription: `The hostname patterns to match for redirection. For more information on Split DNS hostname pattern formatting, please consult the Split DNS KB.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of profile`,
							Computed:            true,
						},
						"nameservers": schema.SingleNestedAttribute{
							MarkdownDescription: `Contains the nameserver information for redirection.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"addresses": schema.ListAttribute{
									MarkdownDescription: `The nameserver address(es) to use for redirection. A maximum of one address is supported.`,
									Computed:            true,
									ElementType:         types.StringType,
								},
							},
						},
						"profile_id": schema.StringAttribute{
							MarkdownDescription: `Profile ID`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsApplianceDNSSplitProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsApplianceDNSSplitProfiles OrganizationsApplianceDNSSplitProfiles
	diags := req.Config.Get(ctx, &organizationsApplianceDNSSplitProfiles)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationApplianceDNSSplitProfiles")
		vvOrganizationID := organizationsApplianceDNSSplitProfiles.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationApplianceDNSSplitProfilesQueryParams{}

		queryParams1.ProfileIDs = elementsToStrings(ctx, organizationsApplianceDNSSplitProfiles.ProfileIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetOrganizationApplianceDNSSplitProfiles(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationApplianceDNSSplitProfiles",
				err.Error(),
			)
			return
		}

		organizationsApplianceDNSSplitProfiles = ResponseApplianceGetOrganizationApplianceDNSSplitProfilesItemsToBody(organizationsApplianceDNSSplitProfiles, response1)
		diags = resp.State.Set(ctx, &organizationsApplianceDNSSplitProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsApplianceDNSSplitProfiles struct {
	OrganizationID types.String                                                     `tfsdk:"organization_id"`
	ProfileIDs     types.List                                                       `tfsdk:"profile_ids"`
	Items          *[]ResponseItemApplianceGetOrganizationApplianceDnsSplitProfiles `tfsdk:"items"`
}

type ResponseItemApplianceGetOrganizationApplianceDnsSplitProfiles struct {
	Hostnames   types.List                                                                `tfsdk:"hostnames"`
	Name        types.String                                                              `tfsdk:"name"`
	Nameservers *ResponseItemApplianceGetOrganizationApplianceDnsSplitProfilesNameservers `tfsdk:"nameservers"`
	ProfileID   types.String                                                              `tfsdk:"profile_id"`
}

type ResponseItemApplianceGetOrganizationApplianceDnsSplitProfilesNameservers struct {
	Addresses types.List `tfsdk:"addresses"`
}

// ToBody
func ResponseApplianceGetOrganizationApplianceDNSSplitProfilesItemsToBody(state OrganizationsApplianceDNSSplitProfiles, response *merakigosdk.ResponseApplianceGetOrganizationApplianceDNSSplitProfiles) OrganizationsApplianceDNSSplitProfiles {
	var items []ResponseItemApplianceGetOrganizationApplianceDnsSplitProfiles
	for _, item := range *response {
		itemState := ResponseItemApplianceGetOrganizationApplianceDnsSplitProfiles{
			Hostnames: StringSliceToList(item.Hostnames),
			Name:      types.StringValue(item.Name),
			Nameservers: func() *ResponseItemApplianceGetOrganizationApplianceDnsSplitProfilesNameservers {
				if item.Nameservers != nil {
					return &ResponseItemApplianceGetOrganizationApplianceDnsSplitProfilesNameservers{
						Addresses: StringSliceToList(item.Nameservers.Addresses),
					}
				}
				return nil
			}(),
			ProfileID: types.StringValue(item.ProfileID),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
