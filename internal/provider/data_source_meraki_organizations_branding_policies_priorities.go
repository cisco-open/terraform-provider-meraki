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
	_ datasource.DataSource              = &OrganizationsBrandingPoliciesPrioritiesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsBrandingPoliciesPrioritiesDataSource{}
)

func NewOrganizationsBrandingPoliciesPrioritiesDataSource() datasource.DataSource {
	return &OrganizationsBrandingPoliciesPrioritiesDataSource{}
}

type OrganizationsBrandingPoliciesPrioritiesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsBrandingPoliciesPrioritiesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsBrandingPoliciesPrioritiesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_branding_policies_priorities"
}

func (d *OrganizationsBrandingPoliciesPrioritiesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"branding_policy_ids": schema.ListAttribute{
						MarkdownDescription: `      An ordered list of branding policy IDs that determines the priority order of how to apply the policies
`,
						Computed:    true,
						ElementType: types.StringType,
					},
				},
			},
		},
	}
}

func (d *OrganizationsBrandingPoliciesPrioritiesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsBrandingPoliciesPriorities OrganizationsBrandingPoliciesPriorities
	diags := req.Config.Get(ctx, &organizationsBrandingPoliciesPriorities)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationBrandingPoliciesPriorities")
		vvOrganizationID := organizationsBrandingPoliciesPriorities.OrganizationID.ValueString()

		response1, restyResp1, err := d.client.Organizations.GetOrganizationBrandingPoliciesPriorities(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationBrandingPoliciesPriorities",
				err.Error(),
			)
			return
		}

		organizationsBrandingPoliciesPriorities = ResponseOrganizationsGetOrganizationBrandingPoliciesPrioritiesItemToBody(organizationsBrandingPoliciesPriorities, response1)
		diags = resp.State.Set(ctx, &organizationsBrandingPoliciesPriorities)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsBrandingPoliciesPriorities struct {
	OrganizationID types.String                                                    `tfsdk:"organization_id"`
	Item           *ResponseOrganizationsGetOrganizationBrandingPoliciesPriorities `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationBrandingPoliciesPriorities struct {
	BrandingPolicyIDs types.List `tfsdk:"branding_policy_ids"`
}

// ToBody
func ResponseOrganizationsGetOrganizationBrandingPoliciesPrioritiesItemToBody(state OrganizationsBrandingPoliciesPriorities, response *merakigosdk.ResponseOrganizationsGetOrganizationBrandingPoliciesPriorities) OrganizationsBrandingPoliciesPriorities {
	itemState := ResponseOrganizationsGetOrganizationBrandingPoliciesPriorities{
		BrandingPolicyIDs: StringSliceToList(response.BrandingPolicyIDs),
	}
	state.Item = &itemState
	return state
}
