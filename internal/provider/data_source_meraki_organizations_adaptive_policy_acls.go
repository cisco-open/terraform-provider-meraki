package provider

// DATA SOURCE NORMAL
import (
	"context"
	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsAdaptivePolicyACLsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsAdaptivePolicyACLsDataSource{}
)

func NewOrganizationsAdaptivePolicyACLsDataSource() datasource.DataSource {
	return &OrganizationsAdaptivePolicyACLsDataSource{}
}

type OrganizationsAdaptivePolicyACLsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsAdaptivePolicyACLsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsAdaptivePolicyACLsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_adaptive_policy_acls"
}

func (d *OrganizationsAdaptivePolicyACLsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"acl_id": schema.StringAttribute{
				MarkdownDescription: `aclId path parameter. Acl ID`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"acl_id": schema.StringAttribute{
						MarkdownDescription: `ID of the adaptive policy ACL`,
						Computed:            true,
					},
					"created_at": schema.StringAttribute{
						MarkdownDescription: `When the adaptive policy ACL was created`,
						Computed:            true,
					},
					"description": schema.StringAttribute{
						MarkdownDescription: `Description of the adaptive policy ACL`,
						Computed:            true,
					},
					"ip_version": schema.StringAttribute{
						MarkdownDescription: `IP version of adpative policy ACL`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of the adaptive policy ACL`,
						Computed:            true,
					},
					"rules": schema.SetNestedAttribute{
						MarkdownDescription: `An ordered array of the adaptive policy ACL rules`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"dst_port": schema.StringAttribute{
									MarkdownDescription: `Destination port`,
									Computed:            true,
								},
								"policy": schema.StringAttribute{
									MarkdownDescription: `'allow' or 'deny' traffic specified by this rule`,
									Computed:            true,
								},
								"protocol": schema.StringAttribute{
									MarkdownDescription: `The type of protocol`,
									Computed:            true,
								},
								"src_port": schema.StringAttribute{
									MarkdownDescription: `Source port`,
									Computed:            true,
								},
							},
						},
					},
					"updated_at": schema.StringAttribute{
						MarkdownDescription: `When the adaptive policy ACL was last updated`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationAdaptivePolicyAcls`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"acl_id": schema.StringAttribute{
							MarkdownDescription: `ID of the adaptive policy ACL`,
							Computed:            true,
						},
						"created_at": schema.StringAttribute{
							MarkdownDescription: `When the adaptive policy ACL was created`,
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: `Description of the adaptive policy ACL`,
							Computed:            true,
						},
						"ip_version": schema.StringAttribute{
							MarkdownDescription: `IP version of adpative policy ACL`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the adaptive policy ACL`,
							Computed:            true,
						},
						"rules": schema.SetNestedAttribute{
							MarkdownDescription: `An ordered array of the adaptive policy ACL rules`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"dst_port": schema.StringAttribute{
										MarkdownDescription: `Destination port`,
										Computed:            true,
									},
									"policy": schema.StringAttribute{
										MarkdownDescription: `'allow' or 'deny' traffic specified by this rule`,
										Computed:            true,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: `The type of protocol`,
										Computed:            true,
									},
									"src_port": schema.StringAttribute{
										MarkdownDescription: `Source port`,
										Computed:            true,
									},
								},
							},
						},
						"updated_at": schema.StringAttribute{
							MarkdownDescription: `When the adaptive policy ACL was last updated`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsAdaptivePolicyACLsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsAdaptivePolicyACLs OrganizationsAdaptivePolicyACLs
	diags := req.Config.Get(ctx, &organizationsAdaptivePolicyACLs)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsAdaptivePolicyACLs.OrganizationID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsAdaptivePolicyACLs.OrganizationID.IsNull(), !organizationsAdaptivePolicyACLs.ACLID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationAdaptivePolicyACLs")
		vvOrganizationID := organizationsAdaptivePolicyACLs.OrganizationID.ValueString()

		response1, restyResp1, err := d.client.Organizations.GetOrganizationAdaptivePolicyACLs(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAdaptivePolicyACLs",
				err.Error(),
			)
			return
		}

		organizationsAdaptivePolicyACLs = ResponseOrganizationsGetOrganizationAdaptivePolicyACLsItemsToBody(organizationsAdaptivePolicyACLs, response1)
		diags = resp.State.Set(ctx, &organizationsAdaptivePolicyACLs)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationAdaptivePolicyACL")
		vvOrganizationID := organizationsAdaptivePolicyACLs.OrganizationID.ValueString()
		vvACLID := organizationsAdaptivePolicyACLs.ACLID.ValueString()

		response2, restyResp2, err := d.client.Organizations.GetOrganizationAdaptivePolicyACL(vvOrganizationID, vvACLID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAdaptivePolicyACL",
				err.Error(),
			)
			return
		}

		organizationsAdaptivePolicyACLs = ResponseOrganizationsGetOrganizationAdaptivePolicyACLItemToBody(organizationsAdaptivePolicyACLs, response2)
		diags = resp.State.Set(ctx, &organizationsAdaptivePolicyACLs)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsAdaptivePolicyACLs struct {
	OrganizationID types.String                                                  `tfsdk:"organization_id"`
	ACLID          types.String                                                  `tfsdk:"acl_id"`
	Items          *[]ResponseItemOrganizationsGetOrganizationAdaptivePolicyAcls `tfsdk:"items"`
	Item           *ResponseOrganizationsGetOrganizationAdaptivePolicyAcl        `tfsdk:"item"`
}

type ResponseItemOrganizationsGetOrganizationAdaptivePolicyAcls struct {
	ACLID       types.String                                                       `tfsdk:"acl_id"`
	CreatedAt   types.String                                                       `tfsdk:"created_at"`
	Description types.String                                                       `tfsdk:"description"`
	IPVersion   types.String                                                       `tfsdk:"ip_version"`
	Name        types.String                                                       `tfsdk:"name"`
	Rules       *[]ResponseItemOrganizationsGetOrganizationAdaptivePolicyAclsRules `tfsdk:"rules"`
	UpdatedAt   types.String                                                       `tfsdk:"updated_at"`
}

type ResponseItemOrganizationsGetOrganizationAdaptivePolicyAclsRules struct {
	DstPort  types.String `tfsdk:"dst_port"`
	Policy   types.String `tfsdk:"policy"`
	Protocol types.String `tfsdk:"protocol"`
	SrcPort  types.String `tfsdk:"src_port"`
}

type ResponseOrganizationsGetOrganizationAdaptivePolicyAcl struct {
	ACLID       types.String                                                  `tfsdk:"acl_id"`
	CreatedAt   types.String                                                  `tfsdk:"created_at"`
	Description types.String                                                  `tfsdk:"description"`
	IPVersion   types.String                                                  `tfsdk:"ip_version"`
	Name        types.String                                                  `tfsdk:"name"`
	Rules       *[]ResponseOrganizationsGetOrganizationAdaptivePolicyAclRules `tfsdk:"rules"`
	UpdatedAt   types.String                                                  `tfsdk:"updated_at"`
}

type ResponseOrganizationsGetOrganizationAdaptivePolicyAclRules struct {
	DstPort  types.String `tfsdk:"dst_port"`
	Policy   types.String `tfsdk:"policy"`
	Protocol types.String `tfsdk:"protocol"`
	SrcPort  types.String `tfsdk:"src_port"`
}

// ToBody
func ResponseOrganizationsGetOrganizationAdaptivePolicyACLsItemsToBody(state OrganizationsAdaptivePolicyACLs, response *merakigosdk.ResponseOrganizationsGetOrganizationAdaptivePolicyACLs) OrganizationsAdaptivePolicyACLs {
	var items []ResponseItemOrganizationsGetOrganizationAdaptivePolicyAcls
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationAdaptivePolicyAcls{
			ACLID:       types.StringValue(item.ACLID),
			CreatedAt:   types.StringValue(item.CreatedAt),
			Description: types.StringValue(item.Description),
			IPVersion:   types.StringValue(item.IPVersion),
			Name:        types.StringValue(item.Name),
			Rules: func() *[]ResponseItemOrganizationsGetOrganizationAdaptivePolicyAclsRules {
				if item.Rules != nil {
					result := make([]ResponseItemOrganizationsGetOrganizationAdaptivePolicyAclsRules, len(*item.Rules))
					for i, rules := range *item.Rules {
						result[i] = ResponseItemOrganizationsGetOrganizationAdaptivePolicyAclsRules{
							DstPort:  types.StringValue(rules.DstPort),
							Policy:   types.StringValue(rules.Policy),
							Protocol: types.StringValue(rules.Protocol),
							SrcPort:  types.StringValue(rules.SrcPort),
						}
					}
					return &result
				}
				return &[]ResponseItemOrganizationsGetOrganizationAdaptivePolicyAclsRules{}
			}(),
			UpdatedAt: types.StringValue(item.UpdatedAt),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseOrganizationsGetOrganizationAdaptivePolicyACLItemToBody(state OrganizationsAdaptivePolicyACLs, response *merakigosdk.ResponseOrganizationsGetOrganizationAdaptivePolicyACL) OrganizationsAdaptivePolicyACLs {
	itemState := ResponseOrganizationsGetOrganizationAdaptivePolicyAcl{
		ACLID:       types.StringValue(response.ACLID),
		CreatedAt:   types.StringValue(response.CreatedAt),
		Description: types.StringValue(response.Description),
		IPVersion:   types.StringValue(response.IPVersion),
		Name:        types.StringValue(response.Name),
		Rules: func() *[]ResponseOrganizationsGetOrganizationAdaptivePolicyAclRules {
			if response.Rules != nil {
				result := make([]ResponseOrganizationsGetOrganizationAdaptivePolicyAclRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseOrganizationsGetOrganizationAdaptivePolicyAclRules{
						DstPort:  types.StringValue(rules.DstPort),
						Policy:   types.StringValue(rules.Policy),
						Protocol: types.StringValue(rules.Protocol),
						SrcPort:  types.StringValue(rules.SrcPort),
					}
				}
				return &result
			}
			return &[]ResponseOrganizationsGetOrganizationAdaptivePolicyAclRules{}
		}(),
		UpdatedAt: types.StringValue(response.UpdatedAt),
	}
	state.Item = &itemState
	return state
}
