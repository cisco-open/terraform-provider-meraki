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
	_ datasource.DataSource              = &OrganizationsApplianceSecurityIntrusionDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsApplianceSecurityIntrusionDataSource{}
)

func NewOrganizationsApplianceSecurityIntrusionDataSource() datasource.DataSource {
	return &OrganizationsApplianceSecurityIntrusionDataSource{}
}

type OrganizationsApplianceSecurityIntrusionDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsApplianceSecurityIntrusionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsApplianceSecurityIntrusionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_appliance_security_intrusion"
}

func (d *OrganizationsApplianceSecurityIntrusionDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"allowed_rules": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"message": schema.StringAttribute{
									Computed: true,
								},
								"rule_id": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsApplianceSecurityIntrusionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsApplianceSecurityIntrusion OrganizationsApplianceSecurityIntrusion
	diags := req.Config.Get(ctx, &organizationsApplianceSecurityIntrusion)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationApplianceSecurityIntrusion")
		vvOrganizationID := organizationsApplianceSecurityIntrusion.OrganizationID.ValueString()

		response1, restyResp1, err := d.client.Appliance.GetOrganizationApplianceSecurityIntrusion(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationApplianceSecurityIntrusion",
				err.Error(),
			)
			return
		}

		organizationsApplianceSecurityIntrusion = ResponseApplianceGetOrganizationApplianceSecurityIntrusionItemToBody(organizationsApplianceSecurityIntrusion, response1)
		diags = resp.State.Set(ctx, &organizationsApplianceSecurityIntrusion)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsApplianceSecurityIntrusion struct {
	OrganizationID types.String                                                `tfsdk:"organization_id"`
	Item           *ResponseApplianceGetOrganizationApplianceSecurityIntrusion `tfsdk:"item"`
}

type ResponseApplianceGetOrganizationApplianceSecurityIntrusion struct {
	AllowedRules *[]ResponseApplianceGetOrganizationApplianceSecurityIntrusionAllowedRules `tfsdk:"allowed_rules"`
}

type ResponseApplianceGetOrganizationApplianceSecurityIntrusionAllowedRules struct {
	Message types.String `tfsdk:"message"`
	RuleID  types.String `tfsdk:"rule_id"`
}

// ToBody
func ResponseApplianceGetOrganizationApplianceSecurityIntrusionItemToBody(state OrganizationsApplianceSecurityIntrusion, response *merakigosdk.ResponseApplianceGetOrganizationApplianceSecurityIntrusion) OrganizationsApplianceSecurityIntrusion {
	itemState := ResponseApplianceGetOrganizationApplianceSecurityIntrusion{
		AllowedRules: func() *[]ResponseApplianceGetOrganizationApplianceSecurityIntrusionAllowedRules {
			if response.AllowedRules != nil {
				result := make([]ResponseApplianceGetOrganizationApplianceSecurityIntrusionAllowedRules, len(*response.AllowedRules))
				for i, allowedRules := range *response.AllowedRules {
					result[i] = ResponseApplianceGetOrganizationApplianceSecurityIntrusionAllowedRules{
						Message: types.StringValue(allowedRules.Message),
						RuleID:  types.StringValue(allowedRules.RuleID),
					}
				}
				return &result
			}
			return &[]ResponseApplianceGetOrganizationApplianceSecurityIntrusionAllowedRules{}
		}(),
	}
	state.Item = &itemState
	return state
}
