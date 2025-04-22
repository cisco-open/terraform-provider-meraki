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
	_ datasource.DataSource              = &OrganizationsBrandingPoliciesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsBrandingPoliciesDataSource{}
)

func NewOrganizationsBrandingPoliciesDataSource() datasource.DataSource {
	return &OrganizationsBrandingPoliciesDataSource{}
}

type OrganizationsBrandingPoliciesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsBrandingPoliciesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsBrandingPoliciesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_branding_policies"
}

func (d *OrganizationsBrandingPoliciesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"branding_policy_id": schema.StringAttribute{
				MarkdownDescription: `brandingPolicyId path parameter. Branding policy ID`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"admin_settings": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for describing which kinds of admins this policy applies to.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"applies_to": schema.StringAttribute{
								MarkdownDescription: `Which kinds of admins this policy applies to. Can be one of 'All organization admins', 'All enterprise admins', 'All network admins', 'All admins of networks...', 'All admins of networks tagged...', 'Specific admins...', 'All admins' or 'All SAML admins'.`,
								Computed:            true,
							},
							"values": schema.ListAttribute{
								MarkdownDescription: `      If 'appliesTo' is set to one of 'Specific admins...', 'All admins of networks...' or 'All admins of networks tagged...', then you must specify this 'values' property to provide the set of
      entities to apply the branding policy to. For 'Specific admins...', specify an array of admin IDs. For 'All admins of
      networks...', specify an array of network IDs and/or configuration template IDs. For 'All admins of networks tagged...',
      specify an array of tag names.
`,
								Computed:    true,
								ElementType: types.StringType,
							},
						},
					},
					"custom_logo": schema.SingleNestedAttribute{
						MarkdownDescription: `Properties describing the custom logo attached to the branding policy.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Whether or not there is a custom logo enabled.`,
								Computed:            true,
							},
							"image": schema.SingleNestedAttribute{
								MarkdownDescription: `Properties of the image.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"preview": schema.SingleNestedAttribute{
										MarkdownDescription: `Preview of the image`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"expires_at": schema.StringAttribute{
												MarkdownDescription: `Timestamp of the preview image`,
												Computed:            true,
											},
											"url": schema.StringAttribute{
												MarkdownDescription: `Url of the preview image`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
					},
					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating whether this policy is enabled.`,
						Computed:            true,
					},
					"help_settings": schema.SingleNestedAttribute{
						MarkdownDescription: `      Settings for describing the modifications to various Help page features. Each property in this object accepts one of
      'default or inherit' (do not modify functionality), 'hide' (remove the section from Dashboard), or 'show' (always show
      the section on Dashboard). Some properties in this object also accept custom HTML used to replace the section on
      Dashboard; see the documentation for each property to see the allowed values.
`,
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"api_docs_subtab": schema.StringAttribute{
								MarkdownDescription: `      The 'Help -> API docs' subtab where a detailed description of the Dashboard API is listed. Can be one of
      'default or inherit', 'hide' or 'show'.
`,
								Computed: true,
							},
							"cases_subtab": schema.StringAttribute{
								MarkdownDescription: `      The 'Help -> Cases' Dashboard subtab on which Cisco Meraki support cases for this organization can be managed. Can be one
      of 'default or inherit', 'hide' or 'show'.
`,
								Computed: true,
							},
							"cisco_meraki_product_documentation": schema.StringAttribute{
								MarkdownDescription: `      The 'Product Manuals' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
`,
								Computed: true,
							},
							"community_subtab": schema.StringAttribute{
								MarkdownDescription: `      The 'Help -> Community' subtab which provides a link to Meraki Community. Can be one of 'default or inherit', 'hide' or 'show'.
`,
								Computed: true,
							},
							"data_protection_requests_subtab": schema.StringAttribute{
								MarkdownDescription: `      The 'Help -> Data protection requests' Dashboard subtab on which requests to delete, restrict, or export end-user data can
      be audited. Can be one of 'default or inherit', 'hide' or 'show'.
`,
								Computed: true,
							},
							"firewall_info_subtab": schema.StringAttribute{
								MarkdownDescription: `      The 'Help -> Firewall info' subtab where necessary upstream firewall rules for communication to the Cisco Meraki cloud are
      listed. Can be one of 'default or inherit', 'hide' or 'show'.
`,
								Computed: true,
							},
							"get_help_subtab": schema.StringAttribute{
								MarkdownDescription: `      The 'Help -> Get Help' subtab on which Cisco Meraki KB, Product Manuals, and Support/Case Information are displayed. Note
      that if this subtab is hidden, branding customizations for the KB on 'Get help', Cisco Meraki product documentation,
      and support contact info will not be visible. Can be one of 'default or inherit', 'hide' or 'show'.
`,
								Computed: true,
							},
							"get_help_subtab_knowledge_base_search": schema.StringAttribute{
								MarkdownDescription: `      The KB search box which appears on the Help page. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
`,
								Computed: true,
							},
							"hardware_replacements_subtab": schema.StringAttribute{
								MarkdownDescription: `      The 'Help -> Replacement info' subtab where important information regarding device replacements is detailed. Can be one of
      'default or inherit', 'hide' or 'show'.
`,
								Computed: true,
							},
							"help_tab": schema.StringAttribute{
								MarkdownDescription: `      The Help tab, under which all support information resides. If this tab is hidden, no other 'Help' branding
      customizations will be visible. Can be one of 'default or inherit', 'hide' or 'show'.
`,
								Computed: true,
							},
							"help_widget": schema.StringAttribute{
								MarkdownDescription: `      The 'Help Widget' is a support widget which provides access to live chat, documentation links, Sales contact info,
      and other contact avenues to reach Meraki Support. Can be one of 'default or inherit', 'hide' or 'show'.
`,
								Computed: true,
							},
							"new_features_subtab": schema.StringAttribute{
								MarkdownDescription: `      The 'Help -> New features' subtab where new Dashboard features are detailed. Can be one of 'default or inherit', 'hide' or 'show'.
`,
								Computed: true,
							},
							"sm_forums": schema.StringAttribute{
								MarkdownDescription: `      The 'SM Forums' subtab which links to community-based support for Cisco Meraki Systems Manager. Only configurable for
      organizations that contain Systems Manager networks. Can be one of 'default or inherit', 'hide' or 'show'.
`,
								Computed: true,
							},
							"support_contact_info": schema.StringAttribute{
								MarkdownDescription: `      The 'Contact Meraki Support' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
`,
								Computed: true,
							},
							"universal_search_knowledge_base_search": schema.StringAttribute{
								MarkdownDescription: `      The universal search box always visible on Dashboard will, by default, present results from the Meraki KB. This configures
      whether these Meraki KB results should be returned. Can be one of 'default or inherit', 'hide' or 'show'.
`,
								Computed: true,
							},
						},
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of the Dashboard branding policy.`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationBrandingPolicies`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"admin_settings": schema.SingleNestedAttribute{
							MarkdownDescription: `Settings for describing which kinds of admins this policy applies to.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"applies_to": schema.StringAttribute{
									MarkdownDescription: `Which kinds of admins this policy applies to. Can be one of 'All organization admins', 'All enterprise admins', 'All network admins', 'All admins of networks...', 'All admins of networks tagged...', 'Specific admins...', 'All admins' or 'All SAML admins'.`,
									Computed:            true,
								},
								"values": schema.ListAttribute{
									MarkdownDescription: `      If 'appliesTo' is set to one of 'Specific admins...', 'All admins of networks...' or 'All admins of networks tagged...', then you must specify this 'values' property to provide the set of
      entities to apply the branding policy to. For 'Specific admins...', specify an array of admin IDs. For 'All admins of
      networks...', specify an array of network IDs and/or configuration template IDs. For 'All admins of networks tagged...',
      specify an array of tag names.
`,
									Computed:    true,
									ElementType: types.StringType,
								},
							},
						},
						"custom_logo": schema.SingleNestedAttribute{
							MarkdownDescription: `Properties describing the custom logo attached to the branding policy.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"enabled": schema.BoolAttribute{
									MarkdownDescription: `Whether or not there is a custom logo enabled.`,
									Computed:            true,
								},
								"image": schema.SingleNestedAttribute{
									MarkdownDescription: `Properties of the image.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"preview": schema.SingleNestedAttribute{
											MarkdownDescription: `Preview of the image`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"expires_at": schema.StringAttribute{
													MarkdownDescription: `Timestamp of the preview image`,
													Computed:            true,
												},
												"url": schema.StringAttribute{
													MarkdownDescription: `Url of the preview image`,
													Computed:            true,
												},
											},
										},
									},
								},
							},
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: `Boolean indicating whether this policy is enabled.`,
							Computed:            true,
						},
						"help_settings": schema.SingleNestedAttribute{
							MarkdownDescription: `      Settings for describing the modifications to various Help page features. Each property in this object accepts one of
      'default or inherit' (do not modify functionality), 'hide' (remove the section from Dashboard), or 'show' (always show
      the section on Dashboard). Some properties in this object also accept custom HTML used to replace the section on
      Dashboard; see the documentation for each property to see the allowed values.
`,
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"api_docs_subtab": schema.StringAttribute{
									MarkdownDescription: `      The 'Help -> API docs' subtab where a detailed description of the Dashboard API is listed. Can be one of
      'default or inherit', 'hide' or 'show'.
`,
									Computed: true,
								},
								"cases_subtab": schema.StringAttribute{
									MarkdownDescription: `      The 'Help -> Cases' Dashboard subtab on which Cisco Meraki support cases for this organization can be managed. Can be one
      of 'default or inherit', 'hide' or 'show'.
`,
									Computed: true,
								},
								"cisco_meraki_product_documentation": schema.StringAttribute{
									MarkdownDescription: `      The 'Product Manuals' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
`,
									Computed: true,
								},
								"community_subtab": schema.StringAttribute{
									MarkdownDescription: `      The 'Help -> Community' subtab which provides a link to Meraki Community. Can be one of 'default or inherit', 'hide' or 'show'.
`,
									Computed: true,
								},
								"data_protection_requests_subtab": schema.StringAttribute{
									MarkdownDescription: `      The 'Help -> Data protection requests' Dashboard subtab on which requests to delete, restrict, or export end-user data can
      be audited. Can be one of 'default or inherit', 'hide' or 'show'.
`,
									Computed: true,
								},
								"firewall_info_subtab": schema.StringAttribute{
									MarkdownDescription: `      The 'Help -> Firewall info' subtab where necessary upstream firewall rules for communication to the Cisco Meraki cloud are
      listed. Can be one of 'default or inherit', 'hide' or 'show'.
`,
									Computed: true,
								},
								"get_help_subtab": schema.StringAttribute{
									MarkdownDescription: `      The 'Help -> Get Help' subtab on which Cisco Meraki KB, Product Manuals, and Support/Case Information are displayed. Note
      that if this subtab is hidden, branding customizations for the KB on 'Get help', Cisco Meraki product documentation,
      and support contact info will not be visible. Can be one of 'default or inherit', 'hide' or 'show'.
`,
									Computed: true,
								},
								"get_help_subtab_knowledge_base_search": schema.StringAttribute{
									MarkdownDescription: `      The KB search box which appears on the Help page. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
`,
									Computed: true,
								},
								"hardware_replacements_subtab": schema.StringAttribute{
									MarkdownDescription: `      The 'Help -> Replacement info' subtab where important information regarding device replacements is detailed. Can be one of
      'default or inherit', 'hide' or 'show'.
`,
									Computed: true,
								},
								"help_tab": schema.StringAttribute{
									MarkdownDescription: `      The Help tab, under which all support information resides. If this tab is hidden, no other 'Help' branding
      customizations will be visible. Can be one of 'default or inherit', 'hide' or 'show'.
`,
									Computed: true,
								},
								"help_widget": schema.StringAttribute{
									MarkdownDescription: `      The 'Help Widget' is a support widget which provides access to live chat, documentation links, Sales contact info,
      and other contact avenues to reach Meraki Support. Can be one of 'default or inherit', 'hide' or 'show'.
`,
									Computed: true,
								},
								"new_features_subtab": schema.StringAttribute{
									MarkdownDescription: `      The 'Help -> New features' subtab where new Dashboard features are detailed. Can be one of 'default or inherit', 'hide' or 'show'.
`,
									Computed: true,
								},
								"sm_forums": schema.StringAttribute{
									MarkdownDescription: `      The 'SM Forums' subtab which links to community-based support for Cisco Meraki Systems Manager. Only configurable for
      organizations that contain Systems Manager networks. Can be one of 'default or inherit', 'hide' or 'show'.
`,
									Computed: true,
								},
								"support_contact_info": schema.StringAttribute{
									MarkdownDescription: `      The 'Contact Meraki Support' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
`,
									Computed: true,
								},
								"universal_search_knowledge_base_search": schema.StringAttribute{
									MarkdownDescription: `      The universal search box always visible on Dashboard will, by default, present results from the Meraki KB. This configures
      whether these Meraki KB results should be returned. Can be one of 'default or inherit', 'hide' or 'show'.
`,
									Computed: true,
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the Dashboard branding policy.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsBrandingPoliciesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsBrandingPolicies OrganizationsBrandingPolicies
	diags := req.Config.Get(ctx, &organizationsBrandingPolicies)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsBrandingPolicies.OrganizationID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsBrandingPolicies.OrganizationID.IsNull(), !organizationsBrandingPolicies.BrandingPolicyID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationBrandingPolicies")
		vvOrganizationID := organizationsBrandingPolicies.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationBrandingPolicies(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationBrandingPolicies",
				err.Error(),
			)
			return
		}

		organizationsBrandingPolicies = ResponseOrganizationsGetOrganizationBrandingPoliciesItemsToBody(organizationsBrandingPolicies, response1)
		diags = resp.State.Set(ctx, &organizationsBrandingPolicies)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationBrandingPolicy")
		vvOrganizationID := organizationsBrandingPolicies.OrganizationID.ValueString()
		vvBrandingPolicyID := organizationsBrandingPolicies.BrandingPolicyID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Organizations.GetOrganizationBrandingPolicy(vvOrganizationID, vvBrandingPolicyID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationBrandingPolicy",
				err.Error(),
			)
			return
		}

		organizationsBrandingPolicies = ResponseOrganizationsGetOrganizationBrandingPolicyItemToBody(organizationsBrandingPolicies, response2)
		diags = resp.State.Set(ctx, &organizationsBrandingPolicies)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsBrandingPolicies struct {
	OrganizationID   types.String                                                `tfsdk:"organization_id"`
	BrandingPolicyID types.String                                                `tfsdk:"branding_policy_id"`
	Items            *[]ResponseItemOrganizationsGetOrganizationBrandingPolicies `tfsdk:"items"`
	Item             *ResponseOrganizationsGetOrganizationBrandingPolicy         `tfsdk:"item"`
}

type ResponseItemOrganizationsGetOrganizationBrandingPolicies struct {
	AdminSettings *ResponseItemOrganizationsGetOrganizationBrandingPoliciesAdminSettings `tfsdk:"admin_settings"`
	CustomLogo    *ResponseItemOrganizationsGetOrganizationBrandingPoliciesCustomLogo    `tfsdk:"custom_logo"`
	Enabled       types.Bool                                                             `tfsdk:"enabled"`
	HelpSettings  *ResponseItemOrganizationsGetOrganizationBrandingPoliciesHelpSettings  `tfsdk:"help_settings"`
	Name          types.String                                                           `tfsdk:"name"`
}

type ResponseItemOrganizationsGetOrganizationBrandingPoliciesAdminSettings struct {
	AppliesTo types.String `tfsdk:"applies_to"`
	Values    types.List   `tfsdk:"values"`
}

type ResponseItemOrganizationsGetOrganizationBrandingPoliciesCustomLogo struct {
	Enabled types.Bool                                                               `tfsdk:"enabled"`
	Image   *ResponseItemOrganizationsGetOrganizationBrandingPoliciesCustomLogoImage `tfsdk:"image"`
}

type ResponseItemOrganizationsGetOrganizationBrandingPoliciesCustomLogoImage struct {
	Preview *ResponseItemOrganizationsGetOrganizationBrandingPoliciesCustomLogoImagePreview `tfsdk:"preview"`
}

type ResponseItemOrganizationsGetOrganizationBrandingPoliciesCustomLogoImagePreview struct {
	ExpiresAt types.String `tfsdk:"expires_at"`
	URL       types.String `tfsdk:"url"`
}

type ResponseItemOrganizationsGetOrganizationBrandingPoliciesHelpSettings struct {
	APIDocsSubtab                      types.String `tfsdk:"api_docs_subtab"`
	CasesSubtab                        types.String `tfsdk:"cases_subtab"`
	CiscoMerakiProductDocumentation    types.String `tfsdk:"cisco_meraki_product_documentation"`
	CommunitySubtab                    types.String `tfsdk:"community_subtab"`
	DataProtectionRequestsSubtab       types.String `tfsdk:"data_protection_requests_subtab"`
	FirewallInfoSubtab                 types.String `tfsdk:"firewall_info_subtab"`
	GetHelpSubtab                      types.String `tfsdk:"get_help_subtab"`
	GetHelpSubtabKnowledgeBaseSearch   types.String `tfsdk:"get_help_subtab_knowledge_base_search"`
	HardwareReplacementsSubtab         types.String `tfsdk:"hardware_replacements_subtab"`
	HelpTab                            types.String `tfsdk:"help_tab"`
	HelpWidget                         types.String `tfsdk:"help_widget"`
	NewFeaturesSubtab                  types.String `tfsdk:"new_features_subtab"`
	SmForums                           types.String `tfsdk:"sm_forums"`
	SupportContactInfo                 types.String `tfsdk:"support_contact_info"`
	UniversalSearchKnowledgeBaseSearch types.String `tfsdk:"universal_search_knowledge_base_search"`
}

type ResponseOrganizationsGetOrganizationBrandingPolicy struct {
	AdminSettings *ResponseOrganizationsGetOrganizationBrandingPolicyAdminSettings `tfsdk:"admin_settings"`
	CustomLogo    *ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogo    `tfsdk:"custom_logo"`
	Enabled       types.Bool                                                       `tfsdk:"enabled"`
	HelpSettings  *ResponseOrganizationsGetOrganizationBrandingPolicyHelpSettings  `tfsdk:"help_settings"`
	Name          types.String                                                     `tfsdk:"name"`
}

type ResponseOrganizationsGetOrganizationBrandingPolicyAdminSettings struct {
	AppliesTo types.String `tfsdk:"applies_to"`
	Values    types.List   `tfsdk:"values"`
}

type ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogo struct {
	Enabled types.Bool                                                         `tfsdk:"enabled"`
	Image   *ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImage `tfsdk:"image"`
}

type ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImage struct {
	Preview *ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImagePreview `tfsdk:"preview"`
}

type ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImagePreview struct {
	ExpiresAt types.String `tfsdk:"expires_at"`
	URL       types.String `tfsdk:"url"`
}

type ResponseOrganizationsGetOrganizationBrandingPolicyHelpSettings struct {
	APIDocsSubtab                      types.String `tfsdk:"api_docs_subtab"`
	CasesSubtab                        types.String `tfsdk:"cases_subtab"`
	CiscoMerakiProductDocumentation    types.String `tfsdk:"cisco_meraki_product_documentation"`
	CommunitySubtab                    types.String `tfsdk:"community_subtab"`
	DataProtectionRequestsSubtab       types.String `tfsdk:"data_protection_requests_subtab"`
	FirewallInfoSubtab                 types.String `tfsdk:"firewall_info_subtab"`
	GetHelpSubtab                      types.String `tfsdk:"get_help_subtab"`
	GetHelpSubtabKnowledgeBaseSearch   types.String `tfsdk:"get_help_subtab_knowledge_base_search"`
	HardwareReplacementsSubtab         types.String `tfsdk:"hardware_replacements_subtab"`
	HelpTab                            types.String `tfsdk:"help_tab"`
	HelpWidget                         types.String `tfsdk:"help_widget"`
	NewFeaturesSubtab                  types.String `tfsdk:"new_features_subtab"`
	SmForums                           types.String `tfsdk:"sm_forums"`
	SupportContactInfo                 types.String `tfsdk:"support_contact_info"`
	UniversalSearchKnowledgeBaseSearch types.String `tfsdk:"universal_search_knowledge_base_search"`
}

// ToBody
func ResponseOrganizationsGetOrganizationBrandingPoliciesItemsToBody(state OrganizationsBrandingPolicies, response *merakigosdk.ResponseOrganizationsGetOrganizationBrandingPolicies) OrganizationsBrandingPolicies {
	var items []ResponseItemOrganizationsGetOrganizationBrandingPolicies
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationBrandingPolicies{
			AdminSettings: func() *ResponseItemOrganizationsGetOrganizationBrandingPoliciesAdminSettings {
				if item.AdminSettings != nil {
					return &ResponseItemOrganizationsGetOrganizationBrandingPoliciesAdminSettings{
						AppliesTo: types.StringValue(item.AdminSettings.AppliesTo),
						Values:    StringSliceToList(item.AdminSettings.Values),
					}
				}
				return nil
			}(),
			CustomLogo: func() *ResponseItemOrganizationsGetOrganizationBrandingPoliciesCustomLogo {
				if item.CustomLogo != nil {
					return &ResponseItemOrganizationsGetOrganizationBrandingPoliciesCustomLogo{
						Enabled: func() types.Bool {
							if item.CustomLogo.Enabled != nil {
								return types.BoolValue(*item.CustomLogo.Enabled)
							}
							return types.Bool{}
						}(),
						Image: func() *ResponseItemOrganizationsGetOrganizationBrandingPoliciesCustomLogoImage {
							if item.CustomLogo.Image != nil {
								return &ResponseItemOrganizationsGetOrganizationBrandingPoliciesCustomLogoImage{
									Preview: func() *ResponseItemOrganizationsGetOrganizationBrandingPoliciesCustomLogoImagePreview {
										if item.CustomLogo.Image.Preview != nil {
											return &ResponseItemOrganizationsGetOrganizationBrandingPoliciesCustomLogoImagePreview{
												ExpiresAt: types.StringValue(item.CustomLogo.Image.Preview.ExpiresAt),
												URL:       types.StringValue(item.CustomLogo.Image.Preview.URL),
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
			Enabled: func() types.Bool {
				if item.Enabled != nil {
					return types.BoolValue(*item.Enabled)
				}
				return types.Bool{}
			}(),
			HelpSettings: func() *ResponseItemOrganizationsGetOrganizationBrandingPoliciesHelpSettings {
				if item.HelpSettings != nil {
					return &ResponseItemOrganizationsGetOrganizationBrandingPoliciesHelpSettings{
						APIDocsSubtab:                      types.StringValue(item.HelpSettings.APIDocsSubtab),
						CasesSubtab:                        types.StringValue(item.HelpSettings.CasesSubtab),
						CiscoMerakiProductDocumentation:    types.StringValue(item.HelpSettings.CiscoMerakiProductDocumentation),
						CommunitySubtab:                    types.StringValue(item.HelpSettings.CommunitySubtab),
						DataProtectionRequestsSubtab:       types.StringValue(item.HelpSettings.DataProtectionRequestsSubtab),
						FirewallInfoSubtab:                 types.StringValue(item.HelpSettings.FirewallInfoSubtab),
						GetHelpSubtab:                      types.StringValue(item.HelpSettings.GetHelpSubtab),
						GetHelpSubtabKnowledgeBaseSearch:   types.StringValue(item.HelpSettings.GetHelpSubtabKnowledgeBaseSearch),
						HardwareReplacementsSubtab:         types.StringValue(item.HelpSettings.HardwareReplacementsSubtab),
						HelpTab:                            types.StringValue(item.HelpSettings.HelpTab),
						HelpWidget:                         types.StringValue(item.HelpSettings.HelpWidget),
						NewFeaturesSubtab:                  types.StringValue(item.HelpSettings.NewFeaturesSubtab),
						SmForums:                           types.StringValue(item.HelpSettings.SmForums),
						SupportContactInfo:                 types.StringValue(item.HelpSettings.SupportContactInfo),
						UniversalSearchKnowledgeBaseSearch: types.StringValue(item.HelpSettings.UniversalSearchKnowledgeBaseSearch),
					}
				}
				return nil
			}(),
			Name: types.StringValue(item.Name),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseOrganizationsGetOrganizationBrandingPolicyItemToBody(state OrganizationsBrandingPolicies, response *merakigosdk.ResponseOrganizationsGetOrganizationBrandingPolicy) OrganizationsBrandingPolicies {
	itemState := ResponseOrganizationsGetOrganizationBrandingPolicy{
		AdminSettings: func() *ResponseOrganizationsGetOrganizationBrandingPolicyAdminSettings {
			if response.AdminSettings != nil {
				return &ResponseOrganizationsGetOrganizationBrandingPolicyAdminSettings{
					AppliesTo: types.StringValue(response.AdminSettings.AppliesTo),
					Values:    StringSliceToList(response.AdminSettings.Values),
				}
			}
			return nil
		}(),
		CustomLogo: func() *ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogo {
			if response.CustomLogo != nil {
				return &ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogo{
					Enabled: func() types.Bool {
						if response.CustomLogo.Enabled != nil {
							return types.BoolValue(*response.CustomLogo.Enabled)
						}
						return types.Bool{}
					}(),
					Image: func() *ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImage {
						if response.CustomLogo.Image != nil {
							return &ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImage{
								Preview: func() *ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImagePreview {
									if response.CustomLogo.Image.Preview != nil {
										return &ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImagePreview{
											ExpiresAt: types.StringValue(response.CustomLogo.Image.Preview.ExpiresAt),
											URL:       types.StringValue(response.CustomLogo.Image.Preview.URL),
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
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		HelpSettings: func() *ResponseOrganizationsGetOrganizationBrandingPolicyHelpSettings {
			if response.HelpSettings != nil {
				return &ResponseOrganizationsGetOrganizationBrandingPolicyHelpSettings{
					APIDocsSubtab:                      types.StringValue(response.HelpSettings.APIDocsSubtab),
					CasesSubtab:                        types.StringValue(response.HelpSettings.CasesSubtab),
					CiscoMerakiProductDocumentation:    types.StringValue(response.HelpSettings.CiscoMerakiProductDocumentation),
					CommunitySubtab:                    types.StringValue(response.HelpSettings.CommunitySubtab),
					DataProtectionRequestsSubtab:       types.StringValue(response.HelpSettings.DataProtectionRequestsSubtab),
					FirewallInfoSubtab:                 types.StringValue(response.HelpSettings.FirewallInfoSubtab),
					GetHelpSubtab:                      types.StringValue(response.HelpSettings.GetHelpSubtab),
					GetHelpSubtabKnowledgeBaseSearch:   types.StringValue(response.HelpSettings.GetHelpSubtabKnowledgeBaseSearch),
					HardwareReplacementsSubtab:         types.StringValue(response.HelpSettings.HardwareReplacementsSubtab),
					HelpTab:                            types.StringValue(response.HelpSettings.HelpTab),
					HelpWidget:                         types.StringValue(response.HelpSettings.HelpWidget),
					NewFeaturesSubtab:                  types.StringValue(response.HelpSettings.NewFeaturesSubtab),
					SmForums:                           types.StringValue(response.HelpSettings.SmForums),
					SupportContactInfo:                 types.StringValue(response.HelpSettings.SupportContactInfo),
					UniversalSearchKnowledgeBaseSearch: types.StringValue(response.HelpSettings.UniversalSearchKnowledgeBaseSearch),
				}
			}
			return nil
		}(),
		Name: types.StringValue(response.Name),
	}
	state.Item = &itemState
	return state
}
