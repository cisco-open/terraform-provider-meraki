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
	_ datasource.DataSource              = &OrganizationsPolicyObjectsGroupsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsPolicyObjectsGroupsDataSource{}
)

func NewOrganizationsPolicyObjectsGroupsDataSource() datasource.DataSource {
	return &OrganizationsPolicyObjectsGroupsDataSource{}
}

type OrganizationsPolicyObjectsGroupsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsPolicyObjectsGroupsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsPolicyObjectsGroupsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_policy_objects_groups"
}

func (d *OrganizationsPolicyObjectsGroupsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 10 1000. Default is 1000.`,
				Optional:            true,
			},
			"policy_object_group_id": schema.StringAttribute{
				MarkdownDescription: `policyObjectGroupId path parameter. Policy object group ID`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"category": schema.StringAttribute{
						MarkdownDescription: `Type of object groups. (NetworkObjectGroup, GeoLocationGroup, PortObjectGroup, ApplicationGroup)`,
						Computed:            true,
					},
					"created_at": schema.StringAttribute{
						MarkdownDescription: `Time Stamp of policy object creation.`,
						Computed:            true,
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `Policy object ID`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of the Policy object group.`,
						Computed:            true,
					},
					"network_ids": schema.ListAttribute{
						MarkdownDescription: `Network ID's associated with the policy objects.`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"object_ids": schema.SetAttribute{
						MarkdownDescription: `Policy objects associated with Network Object Group or Port Object Group`,
						Computed:            true,
						ElementType:         types.StringType, //TODO FINAL ELSE param_schema.Elem.Type para revisar
						// {'Type': 'schema.TypeInt'}
					},
					"updated_at": schema.StringAttribute{
						MarkdownDescription: `Time Stamp of policy object updation.`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *OrganizationsPolicyObjectsGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsPolicyObjectsGroups OrganizationsPolicyObjectsGroups
	diags := req.Config.Get(ctx, &organizationsPolicyObjectsGroups)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsPolicyObjectsGroups.OrganizationID.IsNull(), !organizationsPolicyObjectsGroups.PerPage.IsNull(), !organizationsPolicyObjectsGroups.StartingAfter.IsNull(), !organizationsPolicyObjectsGroups.EndingBefore.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsPolicyObjectsGroups.OrganizationID.IsNull(), !organizationsPolicyObjectsGroups.PolicyObjectGroupID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationPolicyObjectsGroups")
		vvOrganizationID := organizationsPolicyObjectsGroups.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationPolicyObjectsGroupsQueryParams{}

		queryParams1.PerPage = int(organizationsPolicyObjectsGroups.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsPolicyObjectsGroups.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsPolicyObjectsGroups.EndingBefore.ValueString()

		response1, restyResp1, err := d.client.Organizations.GetOrganizationPolicyObjectsGroups(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationPolicyObjectsGroups",
				err.Error(),
			)
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationPolicyObjectsGroup")
		vvOrganizationID := organizationsPolicyObjectsGroups.OrganizationID.ValueString()
		vvPolicyObjectGroupID := organizationsPolicyObjectsGroups.PolicyObjectGroupID.ValueString()

		response2, restyResp2, err := d.client.Organizations.GetOrganizationPolicyObjectsGroup(vvOrganizationID, vvPolicyObjectGroupID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationPolicyObjectsGroup",
				err.Error(),
			)
			return
		}

		organizationsPolicyObjectsGroups = ResponseOrganizationsGetOrganizationPolicyObjectsGroupItemToBody(organizationsPolicyObjectsGroups, response2)
		diags = resp.State.Set(ctx, &organizationsPolicyObjectsGroups)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsPolicyObjectsGroups struct {
	OrganizationID      types.String                                            `tfsdk:"organization_id"`
	PerPage             types.Int64                                             `tfsdk:"per_page"`
	StartingAfter       types.String                                            `tfsdk:"starting_after"`
	EndingBefore        types.String                                            `tfsdk:"ending_before"`
	PolicyObjectGroupID types.String                                            `tfsdk:"policy_object_group_id"`
	Item                *ResponseOrganizationsGetOrganizationPolicyObjectsGroup `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationPolicyObjectsGroup struct {
	Category   types.String `tfsdk:"category"`
	CreatedAt  types.String `tfsdk:"created_at"`
	ID         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	NetworkIDs types.Set    `tfsdk:"network_ids"`
	ObjectIDs  types.Set    `tfsdk:"object_ids"`
	UpdatedAt  types.String `tfsdk:"updated_at"`
}

// ToBody
func ResponseOrganizationsGetOrganizationPolicyObjectsGroupItemToBody(state OrganizationsPolicyObjectsGroups, response *merakigosdk.ResponseOrganizationsGetOrganizationPolicyObjectsGroup) OrganizationsPolicyObjectsGroups {
	itemState := ResponseOrganizationsGetOrganizationPolicyObjectsGroup{
		Category:   types.StringValue(response.Category),
		CreatedAt:  types.StringValue(response.CreatedAt),
		ID:         types.StringValue(response.ID),
		Name:       types.StringValue(response.Name),
		NetworkIDs: StringSliceToSet(response.NetworkIDs),
		ObjectIDs:  StringSliceToSet(*response.ObjectIDs),
		UpdatedAt:  types.StringValue(response.UpdatedAt),
	}
	state.Item = &itemState
	return state
}
