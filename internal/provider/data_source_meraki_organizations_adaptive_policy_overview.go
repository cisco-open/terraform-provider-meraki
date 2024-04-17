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
	_ datasource.DataSource              = &OrganizationsAdaptivePolicyOverviewDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsAdaptivePolicyOverviewDataSource{}
)

func NewOrganizationsAdaptivePolicyOverviewDataSource() datasource.DataSource {
	return &OrganizationsAdaptivePolicyOverviewDataSource{}
}

type OrganizationsAdaptivePolicyOverviewDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsAdaptivePolicyOverviewDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsAdaptivePolicyOverviewDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_adaptive_policy_overview"
}

func (d *OrganizationsAdaptivePolicyOverviewDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"counts": schema.SingleNestedAttribute{
						MarkdownDescription: `The current amount of various adaptive policy objects.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"allow_policies": schema.Int64Attribute{
								MarkdownDescription: `Number of adaptive policies currently in the organization that allow all traffic.`,
								Computed:            true,
							},
							"custom_acls": schema.Int64Attribute{
								MarkdownDescription: `Number of user-created adaptive policy ACLs currently in the organization.`,
								Computed:            true,
							},
							"custom_groups": schema.Int64Attribute{
								MarkdownDescription: `Number of user-created adaptive policy groups currently in the organization.`,
								Computed:            true,
							},
							"deny_policies": schema.Int64Attribute{
								MarkdownDescription: `Number of adaptive policies currently in the organization that deny all traffic.`,
								Computed:            true,
							},
							"groups": schema.Int64Attribute{
								MarkdownDescription: `Number of adaptive policy groups currently in the organization.`,
								Computed:            true,
							},
							"policies": schema.Int64Attribute{
								MarkdownDescription: `Number of adaptive policies currently in the organization.`,
								Computed:            true,
							},
							"policy_objects": schema.Int64Attribute{
								MarkdownDescription: `Number of policy objects (with the adaptive policy type) currently in the organization.`,
								Computed:            true,
							},
						},
					},
					"limits": schema.SingleNestedAttribute{
						MarkdownDescription: `The current limits of various adaptive policy objects.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"acls_in_a_policy": schema.Int64Attribute{
								MarkdownDescription: `Maximum number of adaptive policy ACLs that can be assigned to an adaptive policy in the organization.`,
								Computed:            true,
							},
							"custom_groups": schema.Int64Attribute{
								MarkdownDescription: `Maximum number of user-created adaptive policy groups allowed in the organization.`,
								Computed:            true,
							},
							"policy_objects": schema.Int64Attribute{
								MarkdownDescription: `Maximum number of policy objects (with the adaptive policy type) allowed in the organization.`,
								Computed:            true,
							},
							"rules_in_an_acl": schema.Int64Attribute{
								MarkdownDescription: `Maximum number of rules allowed in an adaptive policy ACL in the organization.`,
								Computed:            true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsAdaptivePolicyOverviewDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsAdaptivePolicyOverview OrganizationsAdaptivePolicyOverview
	diags := req.Config.Get(ctx, &organizationsAdaptivePolicyOverview)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationAdaptivePolicyOverview")
		vvOrganizationID := organizationsAdaptivePolicyOverview.OrganizationID.ValueString()

		response1, restyResp1, err := d.client.Organizations.GetOrganizationAdaptivePolicyOverview(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAdaptivePolicyOverview",
				err.Error(),
			)
			return
		}

		organizationsAdaptivePolicyOverview = ResponseOrganizationsGetOrganizationAdaptivePolicyOverviewItemToBody(organizationsAdaptivePolicyOverview, response1)
		diags = resp.State.Set(ctx, &organizationsAdaptivePolicyOverview)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsAdaptivePolicyOverview struct {
	OrganizationID types.String                                                `tfsdk:"organization_id"`
	Item           *ResponseOrganizationsGetOrganizationAdaptivePolicyOverview `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationAdaptivePolicyOverview struct {
	Counts *ResponseOrganizationsGetOrganizationAdaptivePolicyOverviewCounts `tfsdk:"counts"`
	Limits *ResponseOrganizationsGetOrganizationAdaptivePolicyOverviewLimits `tfsdk:"limits"`
}

type ResponseOrganizationsGetOrganizationAdaptivePolicyOverviewCounts struct {
	AllowPolicies types.Int64 `tfsdk:"allow_policies"`
	CustomACLs    types.Int64 `tfsdk:"custom_acls"`
	CustomGroups  types.Int64 `tfsdk:"custom_groups"`
	DenyPolicies  types.Int64 `tfsdk:"deny_policies"`
	Groups        types.Int64 `tfsdk:"groups"`
	Policies      types.Int64 `tfsdk:"policies"`
	PolicyObjects types.Int64 `tfsdk:"policy_objects"`
}

type ResponseOrganizationsGetOrganizationAdaptivePolicyOverviewLimits struct {
	ACLsInAPolicy types.Int64 `tfsdk:"acls_in_a_policy"`
	CustomGroups  types.Int64 `tfsdk:"custom_groups"`
	PolicyObjects types.Int64 `tfsdk:"policy_objects"`
	RulesInAnACL  types.Int64 `tfsdk:"rules_in_an_acl"`
}

// ToBody
func ResponseOrganizationsGetOrganizationAdaptivePolicyOverviewItemToBody(state OrganizationsAdaptivePolicyOverview, response *merakigosdk.ResponseOrganizationsGetOrganizationAdaptivePolicyOverview) OrganizationsAdaptivePolicyOverview {
	itemState := ResponseOrganizationsGetOrganizationAdaptivePolicyOverview{
		Counts: func() *ResponseOrganizationsGetOrganizationAdaptivePolicyOverviewCounts {
			if response.Counts != nil {
				return &ResponseOrganizationsGetOrganizationAdaptivePolicyOverviewCounts{
					AllowPolicies: func() types.Int64 {
						if response.Counts.AllowPolicies != nil {
							return types.Int64Value(int64(*response.Counts.AllowPolicies))
						}
						return types.Int64{}
					}(),
					CustomACLs: func() types.Int64 {
						if response.Counts.CustomACLs != nil {
							return types.Int64Value(int64(*response.Counts.CustomACLs))
						}
						return types.Int64{}
					}(),
					CustomGroups: func() types.Int64 {
						if response.Counts.CustomGroups != nil {
							return types.Int64Value(int64(*response.Counts.CustomGroups))
						}
						return types.Int64{}
					}(),
					DenyPolicies: func() types.Int64 {
						if response.Counts.DenyPolicies != nil {
							return types.Int64Value(int64(*response.Counts.DenyPolicies))
						}
						return types.Int64{}
					}(),
					Groups: func() types.Int64 {
						if response.Counts.Groups != nil {
							return types.Int64Value(int64(*response.Counts.Groups))
						}
						return types.Int64{}
					}(),
					Policies: func() types.Int64 {
						if response.Counts.Policies != nil {
							return types.Int64Value(int64(*response.Counts.Policies))
						}
						return types.Int64{}
					}(),
					PolicyObjects: func() types.Int64 {
						if response.Counts.PolicyObjects != nil {
							return types.Int64Value(int64(*response.Counts.PolicyObjects))
						}
						return types.Int64{}
					}(),
				}
			}
			return &ResponseOrganizationsGetOrganizationAdaptivePolicyOverviewCounts{}
		}(),
		Limits: func() *ResponseOrganizationsGetOrganizationAdaptivePolicyOverviewLimits {
			if response.Limits != nil {
				return &ResponseOrganizationsGetOrganizationAdaptivePolicyOverviewLimits{
					ACLsInAPolicy: func() types.Int64 {
						if response.Limits.ACLsInAPolicy != nil {
							return types.Int64Value(int64(*response.Limits.ACLsInAPolicy))
						}
						return types.Int64{}
					}(),
					CustomGroups: func() types.Int64 {
						if response.Limits.CustomGroups != nil {
							return types.Int64Value(int64(*response.Limits.CustomGroups))
						}
						return types.Int64{}
					}(),
					PolicyObjects: func() types.Int64 {
						if response.Limits.PolicyObjects != nil {
							return types.Int64Value(int64(*response.Limits.PolicyObjects))
						}
						return types.Int64{}
					}(),
					RulesInAnACL: func() types.Int64 {
						if response.Limits.RulesInAnACL != nil {
							return types.Int64Value(int64(*response.Limits.RulesInAnACL))
						}
						return types.Int64{}
					}(),
				}
			}
			return &ResponseOrganizationsGetOrganizationAdaptivePolicyOverviewLimits{}
		}(),
	}
	state.Item = &itemState
	return state
}
