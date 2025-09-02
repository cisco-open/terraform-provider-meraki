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
	_ datasource.DataSource              = &OrganizationsApplianceDNSLocalProfilesAssignmentsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsApplianceDNSLocalProfilesAssignmentsDataSource{}
)

func NewOrganizationsApplianceDNSLocalProfilesAssignmentsDataSource() datasource.DataSource {
	return &OrganizationsApplianceDNSLocalProfilesAssignmentsDataSource{}
}

type OrganizationsApplianceDNSLocalProfilesAssignmentsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsApplianceDNSLocalProfilesAssignmentsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsApplianceDNSLocalProfilesAssignmentsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_appliance_dns_local_profiles_assignments"
}

func (d *OrganizationsApplianceDNSLocalProfilesAssignmentsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter the results by network IDs`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"profile_ids": schema.ListAttribute{
				MarkdownDescription: `profileIds query parameter. Optional parameter to filter the results by profile IDs`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"items": schema.ListNestedAttribute{
						MarkdownDescription: `List of local DNS profile assignment`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"assignment_id": schema.StringAttribute{
									MarkdownDescription: `ID of the assignment`,
									Computed:            true,
								},
								"network": schema.SingleNestedAttribute{
									MarkdownDescription: `The network attached to the profile`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `ID of the network`,
											Computed:            true,
										},
									},
								},
								"profile": schema.SingleNestedAttribute{
									MarkdownDescription: `The profile the network is attached to`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `ID of the profile`,
											Computed:            true,
										},
									},
								},
							},
						},
					},
					"meta": schema.SingleNestedAttribute{
						MarkdownDescription: `Metadata relevant to the paginated dataset`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"counts": schema.SingleNestedAttribute{
								MarkdownDescription: `Counts relating to the paginated dataset`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"items": schema.SingleNestedAttribute{
										MarkdownDescription: `Counts relating to the paginated items`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"remaining": schema.Int64Attribute{
												MarkdownDescription: `The number of items in the dataset that are available on subsequent pages`,
												Computed:            true,
											},
											"total": schema.Int64Attribute{
												MarkdownDescription: `The total number of items in the dataset`,
												Computed:            true,
											},
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

func (d *OrganizationsApplianceDNSLocalProfilesAssignmentsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsApplianceDNSLocalProfilesAssignments OrganizationsApplianceDNSLocalProfilesAssignments
	diags := req.Config.Get(ctx, &organizationsApplianceDNSLocalProfilesAssignments)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationApplianceDNSLocalProfilesAssignments")
		vvOrganizationID := organizationsApplianceDNSLocalProfilesAssignments.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationApplianceDNSLocalProfilesAssignmentsQueryParams{}

		queryParams1.ProfileIDs = elementsToStrings(ctx, organizationsApplianceDNSLocalProfilesAssignments.ProfileIDs)
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsApplianceDNSLocalProfilesAssignments.NetworkIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetOrganizationApplianceDNSLocalProfilesAssignments(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationApplianceDNSLocalProfilesAssignments",
				err.Error(),
			)
			return
		}

		organizationsApplianceDNSLocalProfilesAssignments = ResponseApplianceGetOrganizationApplianceDNSLocalProfilesAssignmentsItemToBody(organizationsApplianceDNSLocalProfilesAssignments, response1)
		diags = resp.State.Set(ctx, &organizationsApplianceDNSLocalProfilesAssignments)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsApplianceDNSLocalProfilesAssignments struct {
	OrganizationID types.String                                                          `tfsdk:"organization_id"`
	ProfileIDs     types.List                                                            `tfsdk:"profile_ids"`
	NetworkIDs     types.List                                                            `tfsdk:"network_ids"`
	Item           *ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignments `tfsdk:"item"`
}

type ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignments struct {
	Items *[]ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsItems `tfsdk:"items"`
	Meta  *ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsMeta    `tfsdk:"meta"`
}

type ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsItems struct {
	AssignmentID types.String                                                                      `tfsdk:"assignment_id"`
	Network      *ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsItemsNetwork `tfsdk:"network"`
	Profile      *ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsItemsProfile `tfsdk:"profile"`
}

type ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsItemsNetwork struct {
	ID types.String `tfsdk:"id"`
}

type ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsItemsProfile struct {
	ID types.String `tfsdk:"id"`
}

type ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsMeta struct {
	Counts *ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsMetaCounts `tfsdk:"counts"`
}

type ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsMetaCounts struct {
	Items *ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsMetaCountsItems `tfsdk:"items"`
}

type ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseApplianceGetOrganizationApplianceDNSLocalProfilesAssignmentsItemToBody(state OrganizationsApplianceDNSLocalProfilesAssignments, response *merakigosdk.ResponseApplianceGetOrganizationApplianceDNSLocalProfilesAssignments) OrganizationsApplianceDNSLocalProfilesAssignments {
	itemState := ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignments{
		Items: func() *[]ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsItems {
			if response.Items != nil {
				result := make([]ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsItems{
						AssignmentID: func() types.String {
							if items.AssignmentID != "" {
								return types.StringValue(items.AssignmentID)
							}
							return types.String{}
						}(),
						Network: func() *ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsItemsNetwork {
							if items.Network != nil {
								return &ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsItemsNetwork{
									ID: func() types.String {
										if items.Network.ID != "" {
											return types.StringValue(items.Network.ID)
										}
										return types.String{}
									}(),
								}
							}
							return nil
						}(),
						Profile: func() *ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsItemsProfile {
							if items.Profile != nil {
								return &ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsItemsProfile{
									ID: func() types.String {
										if items.Profile.ID != "" {
											return types.StringValue(items.Profile.ID)
										}
										return types.String{}
									}(),
								}
							}
							return nil
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		Meta: func() *ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsMeta {
			if response.Meta != nil {
				return &ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsMeta{
					Counts: func() *ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsMetaCounts{
								Items: func() *ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseApplianceGetOrganizationApplianceDnsLocalProfilesAssignmentsMetaCountsItems{
											Remaining: func() types.Int64 {
												if response.Meta.Counts.Items.Remaining != nil {
													return types.Int64Value(int64(*response.Meta.Counts.Items.Remaining))
												}
												return types.Int64{}
											}(),
											Total: func() types.Int64 {
												if response.Meta.Counts.Items.Total != nil {
													return types.Int64Value(int64(*response.Meta.Counts.Items.Total))
												}
												return types.Int64{}
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
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
