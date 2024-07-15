package provider

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsBrandingPoliciesResource{}
	_ resource.ResourceWithConfigure = &OrganizationsBrandingPoliciesResource{}
)

func NewOrganizationsBrandingPoliciesResource() resource.Resource {
	return &OrganizationsBrandingPoliciesResource{}
}

type OrganizationsBrandingPoliciesResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsBrandingPoliciesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsBrandingPoliciesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_branding_policies"
}

func (r *OrganizationsBrandingPoliciesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"admin_settings": schema.SingleNestedAttribute{
				MarkdownDescription: `Settings for describing which kinds of admins this policy applies to.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"applies_to": schema.StringAttribute{
						MarkdownDescription: `Which kinds of admins this policy applies to. Can be one of 'All organization admins', 'All enterprise admins', 'All network admins', 'All admins of networks...', 'All admins of networks tagged...', 'Specific admins...', 'All admins' or 'All SAML admins'.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"All SAML admins",
								"All admins",
								"All admins of networks tagged...",
								"All admins of networks...",
								"All enterprise admins",
								"All network admins",
								"All organization admins",
								"Specific admins...",
							),
						},
					},
					"values": schema.SetAttribute{
						MarkdownDescription: `      If 'appliesTo' is set to one of 'Specific admins...', 'All admins of networks...' or 'All admins of networks tagged...', then you must specify this 'values' property to provide the set of
      entities to apply the branding policy to. For 'Specific admins...', specify an array of admin IDs. For 'All admins of
      networks...', specify an array of network IDs and/or configuration template IDs. For 'All admins of networks tagged...',
      specify an array of tag names.
`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},

						ElementType: types.StringType,
					},
				},
			},
			"branding_policy_id": schema.StringAttribute{
				MarkdownDescription: `brandingPolicyId path parameter. Branding policy ID`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"custom_logo": schema.SingleNestedAttribute{
				MarkdownDescription: `Properties describing the custom logo attached to the branding policy.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether or not there is a custom logo enabled.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"image": schema.SingleNestedAttribute{
						MarkdownDescription: `Properties of the image.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"contents": schema.StringAttribute{
								MarkdownDescription: `The file contents (a base 64 encoded string) of your new logo.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"format": schema.StringAttribute{
								MarkdownDescription: `The format of the encoded contents.  Supported formats are 'png', 'gif', and jpg'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"gif",
										"jpg",
										"png",
									),
								},
							},
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
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"help_settings": schema.SingleNestedAttribute{
				MarkdownDescription: `      Settings for describing the modifications to various Help page features. Each property in this object accepts one of
      'default or inherit' (do not modify functionality), 'hide' (remove the section from Dashboard), or 'show' (always show
      the section on Dashboard). Some properties in this object also accept custom HTML used to replace the section on
      Dashboard; see the documentation for each property to see the allowed values.
`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"api_docs_subtab": schema.StringAttribute{
						MarkdownDescription: `      The 'Help -> API docs' subtab where a detailed description of the Dashboard API is listed. Can be one of
      'default or inherit', 'hide' or 'show'.
`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"default or inherit",
								"hide",
								"show",
							),
						},
					},
					"cases_subtab": schema.StringAttribute{
						MarkdownDescription: `      The 'Help -> Cases' Dashboard subtab on which Cisco Meraki support cases for this organization can be managed. Can be one
      of 'default or inherit', 'hide' or 'show'.
`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"default or inherit",
								"hide",
								"show",
							),
						},
					},
					"cisco_meraki_product_documentation": schema.StringAttribute{
						MarkdownDescription: `      The 'Product Manuals' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"community_subtab": schema.StringAttribute{
						MarkdownDescription: `      The 'Help -> Community' subtab which provides a link to Meraki Community. Can be one of 'default or inherit', 'hide' or 'show'.
`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"default or inherit",
								"hide",
								"show",
							),
						},
					},
					"data_protection_requests_subtab": schema.StringAttribute{
						MarkdownDescription: `      The 'Help -> Data protection requests' Dashboard subtab on which requests to delete, restrict, or export end-user data can
      be audited. Can be one of 'default or inherit', 'hide' or 'show'.
`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"default or inherit",
								"hide",
								"show",
							),
						},
					},
					"firewall_info_subtab": schema.StringAttribute{
						MarkdownDescription: `      The 'Help -> Firewall info' subtab where necessary upstream firewall rules for communication to the Cisco Meraki cloud are
      listed. Can be one of 'default or inherit', 'hide' or 'show'.
`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"default or inherit",
								"hide",
								"show",
							),
						},
					},
					"get_help_subtab": schema.StringAttribute{
						MarkdownDescription: `      The 'Help -> Get Help' subtab on which Cisco Meraki KB, Product Manuals, and Support/Case Information are displayed. Note
      that if this subtab is hidden, branding customizations for the KB on 'Get help', Cisco Meraki product documentation,
      and support contact info will not be visible. Can be one of 'default or inherit', 'hide' or 'show'.
`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"default or inherit",
								"hide",
								"show",
							),
						},
					},
					"get_help_subtab_knowledge_base_search": schema.StringAttribute{
						MarkdownDescription: `      The KB search box which appears on the Help page. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"hardware_replacements_subtab": schema.StringAttribute{
						MarkdownDescription: `      The 'Help -> Replacement info' subtab where important information regarding device replacements is detailed. Can be one of
      'default or inherit', 'hide' or 'show'.
`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"default or inherit",
								"hide",
								"show",
							),
						},
					},
					"help_tab": schema.StringAttribute{
						MarkdownDescription: `      The Help tab, under which all support information resides. If this tab is hidden, no other 'Help' branding
      customizations will be visible. Can be one of 'default or inherit', 'hide' or 'show'.
`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"default or inherit",
								"hide",
								"show",
							),
						},
					},
					"help_widget": schema.StringAttribute{
						MarkdownDescription: `      The 'Help Widget' is a support widget which provides access to live chat, documentation links, Sales contact info,
      and other contact avenues to reach Meraki Support. Can be one of 'default or inherit', 'hide' or 'show'.
`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"default or inherit",
								"hide",
								"show",
							),
						},
					},
					"new_features_subtab": schema.StringAttribute{
						MarkdownDescription: `      The 'Help -> New features' subtab where new Dashboard features are detailed. Can be one of 'default or inherit', 'hide' or 'show'.
`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"default or inherit",
								"hide",
								"show",
							),
						},
					},
					"sm_forums": schema.StringAttribute{
						MarkdownDescription: `      The 'SM Forums' subtab which links to community-based support for Cisco Meraki Systems Manager. Only configurable for
      organizations that contain Systems Manager networks. Can be one of 'default or inherit', 'hide' or 'show'.
`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"default or inherit",
								"hide",
								"show",
							),
						},
					},
					"support_contact_info": schema.StringAttribute{
						MarkdownDescription: `      The 'Contact Meraki Support' section of the 'Help -> Get Help' subtab. Can be one of 'default or inherit', 'hide', 'show', or a replacement custom HTML string.
`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"universal_search_knowledge_base_search": schema.StringAttribute{
						MarkdownDescription: `      The universal search box always visible on Dashboard will, by default, present results from the Meraki KB. This configures
      whether these Meraki KB results should be returned. Can be one of 'default or inherit', 'hide' or 'show'.
`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"default or inherit",
								"hide",
								"show",
							),
						},
					},
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Name of the Dashboard branding policy.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
		},
	}
}

//path params to set ['brandingPolicyId']

func (r *OrganizationsBrandingPoliciesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsBrandingPoliciesRs

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
	vvOrganizationID := data.OrganizationID.ValueString()
	vvName := data.Name.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Organizations.GetOrganizationBrandingPolicies(vvOrganizationID)
	//Have Create
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationBrandingPolicies",
					err.Error(),
				)
				return
			}
		}
	}
	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvBrandingPolicyID, ok := result2["BrandingPolicyID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter BrandingPolicyID",
					err.Error(),
				)
				return
			}
			r.client.Organizations.UpdateOrganizationBrandingPolicy(vvOrganizationID, vvBrandingPolicyID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Organizations.GetOrganizationBrandingPolicy(vvOrganizationID, vvBrandingPolicyID)
			if responseVerifyItem2 != nil {
				data = ResponseOrganizationsGetOrganizationBrandingPolicyItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Organizations.CreateOrganizationBrandingPolicy(vvOrganizationID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationBrandingPolicy",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationBrandingPolicy",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationBrandingPolicies(vvOrganizationID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationBrandingPolicies",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationBrandingPolicies",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvBrandingPolicyID, ok := result2["BrandingPolicyID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter BrandingPolicyID",
				err.Error(),
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Organizations.GetOrganizationBrandingPolicy(vvOrganizationID, vvBrandingPolicyID)
		if responseVerifyItem2 != nil && err == nil {
			data3 := ResponseOrganizationsGetOrganizationBrandingPolicyItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data3)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationBrandingPolicy",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationBrandingPolicy",
				err.Error(),
			)
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"Error in result.",
			"Error in result.",
		)
		return
	}
}

func (r *OrganizationsBrandingPoliciesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsBrandingPoliciesRs

	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	// Has Item2

	vvOrganizationID := data.OrganizationID.ValueString()
	vvBrandingPolicyID := data.BrandingPolicyID.ValueString()
	responseGet, restyRespGet, err := r.client.Organizations.GetOrganizationBrandingPolicy(vvOrganizationID, vvBrandingPolicyID)
	if err != nil || restyRespGet == nil {
		if restyRespGet != nil {
			if restyRespGet.StatusCode() == 404 {
				resp.Diagnostics.AddWarning(
					"Resource not found",
					"Deleting resource",
				)
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationBrandingPolicy",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationBrandingPolicy",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseOrganizationsGetOrganizationBrandingPolicyItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *OrganizationsBrandingPoliciesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("branding_policy_id"), idParts[1])...)
}

func (r *OrganizationsBrandingPoliciesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsBrandingPoliciesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	vvBrandingPolicyID := data.BrandingPolicyID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Organizations.UpdateOrganizationBrandingPolicy(vvOrganizationID, vvBrandingPolicyID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationBrandingPolicy",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationBrandingPolicy",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsBrandingPoliciesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationsBrandingPoliciesRs
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &state, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}

	vvOrganizationID := state.OrganizationID.ValueString()
	vvBrandingPolicyID := state.BrandingPolicyID.ValueString()
	_, err := r.client.Organizations.DeleteOrganizationBrandingPolicy(vvOrganizationID, vvBrandingPolicyID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganizationBrandingPolicy", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type OrganizationsBrandingPoliciesRs struct {
	OrganizationID   types.String                                                       `tfsdk:"organization_id"`
	BrandingPolicyID types.String                                                       `tfsdk:"branding_policy_id"`
	AdminSettings    *ResponseOrganizationsGetOrganizationBrandingPolicyAdminSettingsRs `tfsdk:"admin_settings"`
	CustomLogo       *ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoRs    `tfsdk:"custom_logo"`
	Enabled          types.Bool                                                         `tfsdk:"enabled"`
	HelpSettings     *ResponseOrganizationsGetOrganizationBrandingPolicyHelpSettingsRs  `tfsdk:"help_settings"`
	Name             types.String                                                       `tfsdk:"name"`
}

type ResponseOrganizationsGetOrganizationBrandingPolicyAdminSettingsRs struct {
	AppliesTo types.String `tfsdk:"applies_to"`
	Values    types.Set    `tfsdk:"values"`
}

type ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoRs struct {
	Enabled types.Bool                                                           `tfsdk:"enabled"`
	Image   *ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImageRs `tfsdk:"image"`
}

type ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImageRs struct {
	Preview  *ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImagePreviewRs `tfsdk:"preview"`
	Contents types.String                                                                `tfsdk:"contents"`
	Format   types.String                                                                `tfsdk:"format"`
}

type ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImagePreviewRs struct {
	ExpiresAt types.String `tfsdk:"expires_at"`
	URL       types.String `tfsdk:"url"`
}

type ResponseOrganizationsGetOrganizationBrandingPolicyHelpSettingsRs struct {
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

// FromBody
func (r *OrganizationsBrandingPoliciesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganizationBrandingPolicy {
	emptyString := ""
	var requestOrganizationsCreateOrganizationBrandingPolicyAdminSettings *merakigosdk.RequestOrganizationsCreateOrganizationBrandingPolicyAdminSettings
	if r.AdminSettings != nil {
		appliesTo := r.AdminSettings.AppliesTo.ValueString()
		var values []string = nil
		//Hoola aqui
		r.AdminSettings.Values.ElementsAs(ctx, &values, false)
		requestOrganizationsCreateOrganizationBrandingPolicyAdminSettings = &merakigosdk.RequestOrganizationsCreateOrganizationBrandingPolicyAdminSettings{
			AppliesTo: appliesTo,
			Values:    values,
		}
	}
	var requestOrganizationsCreateOrganizationBrandingPolicyCustomLogo *merakigosdk.RequestOrganizationsCreateOrganizationBrandingPolicyCustomLogo
	if r.CustomLogo != nil {
		enabled := func() *bool {
			if !r.CustomLogo.Enabled.IsUnknown() && !r.CustomLogo.Enabled.IsNull() {
				return r.CustomLogo.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		var requestOrganizationsCreateOrganizationBrandingPolicyCustomLogoImage *merakigosdk.RequestOrganizationsCreateOrganizationBrandingPolicyCustomLogoImage
		if r.CustomLogo.Image != nil {
			contents := r.CustomLogo.Image.Contents.ValueString()
			format := r.CustomLogo.Image.Format.ValueString()
			requestOrganizationsCreateOrganizationBrandingPolicyCustomLogoImage = &merakigosdk.RequestOrganizationsCreateOrganizationBrandingPolicyCustomLogoImage{
				Contents: contents,
				Format:   format,
			}
		}
		requestOrganizationsCreateOrganizationBrandingPolicyCustomLogo = &merakigosdk.RequestOrganizationsCreateOrganizationBrandingPolicyCustomLogo{
			Enabled: enabled,
			Image:   requestOrganizationsCreateOrganizationBrandingPolicyCustomLogoImage,
		}
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var requestOrganizationsCreateOrganizationBrandingPolicyHelpSettings *merakigosdk.RequestOrganizationsCreateOrganizationBrandingPolicyHelpSettings
	if r.HelpSettings != nil {
		aPIDocsSubtab := r.HelpSettings.APIDocsSubtab.ValueString()
		casesSubtab := r.HelpSettings.CasesSubtab.ValueString()
		ciscoMerakiProductDocumentation := r.HelpSettings.CiscoMerakiProductDocumentation.ValueString()
		communitySubtab := r.HelpSettings.CommunitySubtab.ValueString()
		dataProtectionRequestsSubtab := r.HelpSettings.DataProtectionRequestsSubtab.ValueString()
		firewallInfoSubtab := r.HelpSettings.FirewallInfoSubtab.ValueString()
		getHelpSubtab := r.HelpSettings.GetHelpSubtab.ValueString()
		getHelpSubtabKnowledgeBaseSearch := r.HelpSettings.GetHelpSubtabKnowledgeBaseSearch.ValueString()
		hardwareReplacementsSubtab := r.HelpSettings.HardwareReplacementsSubtab.ValueString()
		helpTab := r.HelpSettings.HelpTab.ValueString()
		helpWidget := r.HelpSettings.HelpWidget.ValueString()
		newFeaturesSubtab := r.HelpSettings.NewFeaturesSubtab.ValueString()
		smForums := r.HelpSettings.SmForums.ValueString()
		supportContactInfo := r.HelpSettings.SupportContactInfo.ValueString()
		universalSearchKnowledgeBaseSearch := r.HelpSettings.UniversalSearchKnowledgeBaseSearch.ValueString()
		requestOrganizationsCreateOrganizationBrandingPolicyHelpSettings = &merakigosdk.RequestOrganizationsCreateOrganizationBrandingPolicyHelpSettings{
			APIDocsSubtab:                      aPIDocsSubtab,
			CasesSubtab:                        casesSubtab,
			CiscoMerakiProductDocumentation:    ciscoMerakiProductDocumentation,
			CommunitySubtab:                    communitySubtab,
			DataProtectionRequestsSubtab:       dataProtectionRequestsSubtab,
			FirewallInfoSubtab:                 firewallInfoSubtab,
			GetHelpSubtab:                      getHelpSubtab,
			GetHelpSubtabKnowledgeBaseSearch:   getHelpSubtabKnowledgeBaseSearch,
			HardwareReplacementsSubtab:         hardwareReplacementsSubtab,
			HelpTab:                            helpTab,
			HelpWidget:                         helpWidget,
			NewFeaturesSubtab:                  newFeaturesSubtab,
			SmForums:                           smForums,
			SupportContactInfo:                 supportContactInfo,
			UniversalSearchKnowledgeBaseSearch: universalSearchKnowledgeBaseSearch,
		}
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	out := merakigosdk.RequestOrganizationsCreateOrganizationBrandingPolicy{
		AdminSettings: requestOrganizationsCreateOrganizationBrandingPolicyAdminSettings,
		CustomLogo:    requestOrganizationsCreateOrganizationBrandingPolicyCustomLogo,
		Enabled:       enabled,
		HelpSettings:  requestOrganizationsCreateOrganizationBrandingPolicyHelpSettings,
		Name:          *name,
	}
	return &out
}
func (r *OrganizationsBrandingPoliciesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestOrganizationsUpdateOrganizationBrandingPolicy {
	emptyString := ""
	var requestOrganizationsUpdateOrganizationBrandingPolicyAdminSettings *merakigosdk.RequestOrganizationsUpdateOrganizationBrandingPolicyAdminSettings
	if r.AdminSettings != nil {
		appliesTo := r.AdminSettings.AppliesTo.ValueString()
		var values []string = nil
		//Hoola aqui
		r.AdminSettings.Values.ElementsAs(ctx, &values, false)
		requestOrganizationsUpdateOrganizationBrandingPolicyAdminSettings = &merakigosdk.RequestOrganizationsUpdateOrganizationBrandingPolicyAdminSettings{
			AppliesTo: appliesTo,
			Values:    values,
		}
	}
	var requestOrganizationsUpdateOrganizationBrandingPolicyCustomLogo *merakigosdk.RequestOrganizationsUpdateOrganizationBrandingPolicyCustomLogo
	if r.CustomLogo != nil {
		enabled := func() *bool {
			if !r.CustomLogo.Enabled.IsUnknown() && !r.CustomLogo.Enabled.IsNull() {
				return r.CustomLogo.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		var requestOrganizationsUpdateOrganizationBrandingPolicyCustomLogoImage *merakigosdk.RequestOrganizationsUpdateOrganizationBrandingPolicyCustomLogoImage
		if r.CustomLogo.Image != nil {
			contents := r.CustomLogo.Image.Contents.ValueString()
			format := r.CustomLogo.Image.Format.ValueString()
			requestOrganizationsUpdateOrganizationBrandingPolicyCustomLogoImage = &merakigosdk.RequestOrganizationsUpdateOrganizationBrandingPolicyCustomLogoImage{
				Contents: contents,
				Format:   format,
			}
		}
		requestOrganizationsUpdateOrganizationBrandingPolicyCustomLogo = &merakigosdk.RequestOrganizationsUpdateOrganizationBrandingPolicyCustomLogo{
			Enabled: enabled,
			Image:   requestOrganizationsUpdateOrganizationBrandingPolicyCustomLogoImage,
		}
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var requestOrganizationsUpdateOrganizationBrandingPolicyHelpSettings *merakigosdk.RequestOrganizationsUpdateOrganizationBrandingPolicyHelpSettings
	if r.HelpSettings != nil {
		aPIDocsSubtab := r.HelpSettings.APIDocsSubtab.ValueString()
		casesSubtab := r.HelpSettings.CasesSubtab.ValueString()
		ciscoMerakiProductDocumentation := r.HelpSettings.CiscoMerakiProductDocumentation.ValueString()
		communitySubtab := r.HelpSettings.CommunitySubtab.ValueString()
		dataProtectionRequestsSubtab := r.HelpSettings.DataProtectionRequestsSubtab.ValueString()
		firewallInfoSubtab := r.HelpSettings.FirewallInfoSubtab.ValueString()
		getHelpSubtab := r.HelpSettings.GetHelpSubtab.ValueString()
		getHelpSubtabKnowledgeBaseSearch := r.HelpSettings.GetHelpSubtabKnowledgeBaseSearch.ValueString()
		hardwareReplacementsSubtab := r.HelpSettings.HardwareReplacementsSubtab.ValueString()
		helpTab := r.HelpSettings.HelpTab.ValueString()
		helpWidget := r.HelpSettings.HelpWidget.ValueString()
		newFeaturesSubtab := r.HelpSettings.NewFeaturesSubtab.ValueString()
		smForums := r.HelpSettings.SmForums.ValueString()
		supportContactInfo := r.HelpSettings.SupportContactInfo.ValueString()
		universalSearchKnowledgeBaseSearch := r.HelpSettings.UniversalSearchKnowledgeBaseSearch.ValueString()
		requestOrganizationsUpdateOrganizationBrandingPolicyHelpSettings = &merakigosdk.RequestOrganizationsUpdateOrganizationBrandingPolicyHelpSettings{
			APIDocsSubtab:                      aPIDocsSubtab,
			CasesSubtab:                        casesSubtab,
			CiscoMerakiProductDocumentation:    ciscoMerakiProductDocumentation,
			CommunitySubtab:                    communitySubtab,
			DataProtectionRequestsSubtab:       dataProtectionRequestsSubtab,
			FirewallInfoSubtab:                 firewallInfoSubtab,
			GetHelpSubtab:                      getHelpSubtab,
			GetHelpSubtabKnowledgeBaseSearch:   getHelpSubtabKnowledgeBaseSearch,
			HardwareReplacementsSubtab:         hardwareReplacementsSubtab,
			HelpTab:                            helpTab,
			HelpWidget:                         helpWidget,
			NewFeaturesSubtab:                  newFeaturesSubtab,
			SmForums:                           smForums,
			SupportContactInfo:                 supportContactInfo,
			UniversalSearchKnowledgeBaseSearch: universalSearchKnowledgeBaseSearch,
		}
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	out := merakigosdk.RequestOrganizationsUpdateOrganizationBrandingPolicy{
		AdminSettings: requestOrganizationsUpdateOrganizationBrandingPolicyAdminSettings,
		CustomLogo:    requestOrganizationsUpdateOrganizationBrandingPolicyCustomLogo,
		Enabled:       enabled,
		HelpSettings:  requestOrganizationsUpdateOrganizationBrandingPolicyHelpSettings,
		Name:          *name,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseOrganizationsGetOrganizationBrandingPolicyItemToBodyRs(state OrganizationsBrandingPoliciesRs, response *merakigosdk.ResponseOrganizationsGetOrganizationBrandingPolicy, is_read bool) OrganizationsBrandingPoliciesRs {
	itemState := OrganizationsBrandingPoliciesRs{
		AdminSettings: func() *ResponseOrganizationsGetOrganizationBrandingPolicyAdminSettingsRs {
			if response.AdminSettings != nil {
				return &ResponseOrganizationsGetOrganizationBrandingPolicyAdminSettingsRs{
					AppliesTo: types.StringValue(response.AdminSettings.AppliesTo),
					Values:    StringSliceToSet(response.AdminSettings.Values),
				}
			}
			return &ResponseOrganizationsGetOrganizationBrandingPolicyAdminSettingsRs{}
		}(),
		CustomLogo: func() *ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoRs {
			if response.CustomLogo != nil {
				return &ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoRs{
					Enabled: func() types.Bool {
						if response.CustomLogo.Enabled != nil {
							return types.BoolValue(*response.CustomLogo.Enabled)
						}
						return types.Bool{}
					}(),
					Image: func() *ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImageRs {
						if response.CustomLogo.Image != nil {
							return &ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImageRs{
								Preview: func() *ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImagePreviewRs {
									if response.CustomLogo.Image.Preview != nil {
										return &ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImagePreviewRs{
											ExpiresAt: types.StringValue(response.CustomLogo.Image.Preview.ExpiresAt),
											URL:       types.StringValue(response.CustomLogo.Image.Preview.URL),
										}
									}
									return &ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImagePreviewRs{}
								}(),
							}
						}
						return &ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoImageRs{}
					}(),
				}
			}
			return &ResponseOrganizationsGetOrganizationBrandingPolicyCustomLogoRs{}
		}(),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		HelpSettings: func() *ResponseOrganizationsGetOrganizationBrandingPolicyHelpSettingsRs {
			if response.HelpSettings != nil {
				return &ResponseOrganizationsGetOrganizationBrandingPolicyHelpSettingsRs{
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
			return &ResponseOrganizationsGetOrganizationBrandingPolicyHelpSettingsRs{}
		}(),
		Name: types.StringValue(response.Name),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsBrandingPoliciesRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsBrandingPoliciesRs)
}
