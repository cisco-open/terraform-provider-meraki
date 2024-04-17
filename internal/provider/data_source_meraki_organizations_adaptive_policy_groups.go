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
	_ datasource.DataSource              = &OrganizationsAdaptivePolicyGroupsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsAdaptivePolicyGroupsDataSource{}
)

func NewOrganizationsAdaptivePolicyGroupsDataSource() datasource.DataSource {
	return &OrganizationsAdaptivePolicyGroupsDataSource{}
}

type OrganizationsAdaptivePolicyGroupsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsAdaptivePolicyGroupsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsAdaptivePolicyGroupsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_adaptive_policy_groups"
}

func (d *OrganizationsAdaptivePolicyGroupsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: `id path parameter.`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"created_at": schema.StringAttribute{
						Computed: true,
					},
					"description": schema.StringAttribute{
						Computed: true,
					},
					"group_id": schema.StringAttribute{
						Computed: true,
					},
					"is_default_group": schema.BoolAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"policy_objects": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
					"required_ip_mappings": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"sgt": schema.Int64Attribute{
						Computed: true,
					},
					"updated_at": schema.StringAttribute{
						Computed: true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationAdaptivePolicyGroups`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"created_at": schema.StringAttribute{
							Computed: true,
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"group_id": schema.StringAttribute{
							Computed: true,
						},
						"is_default_group": schema.BoolAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"policy_objects": schema.SetNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"id": schema.StringAttribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"required_ip_mappings": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"sgt": schema.Int64Attribute{
							Computed: true,
						},
						"updated_at": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsAdaptivePolicyGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsAdaptivePolicyGroups OrganizationsAdaptivePolicyGroups
	diags := req.Config.Get(ctx, &organizationsAdaptivePolicyGroups)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsAdaptivePolicyGroups.OrganizationID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsAdaptivePolicyGroups.OrganizationID.IsNull(), !organizationsAdaptivePolicyGroups.ID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationAdaptivePolicyGroups")
		vvOrganizationID := organizationsAdaptivePolicyGroups.OrganizationID.ValueString()

		response1, restyResp1, err := d.client.Organizations.GetOrganizationAdaptivePolicyGroups(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAdaptivePolicyGroups",
				err.Error(),
			)
			return
		}

		organizationsAdaptivePolicyGroups = ResponseOrganizationsGetOrganizationAdaptivePolicyGroupsItemsToBody(organizationsAdaptivePolicyGroups, response1)
		diags = resp.State.Set(ctx, &organizationsAdaptivePolicyGroups)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationAdaptivePolicyGroup")
		vvOrganizationID := organizationsAdaptivePolicyGroups.OrganizationID.ValueString()
		vvID := organizationsAdaptivePolicyGroups.ID.ValueString()

		response2, restyResp2, err := d.client.Organizations.GetOrganizationAdaptivePolicyGroup(vvOrganizationID, vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAdaptivePolicyGroup",
				err.Error(),
			)
			return
		}

		organizationsAdaptivePolicyGroups = ResponseOrganizationsGetOrganizationAdaptivePolicyGroupItemToBody(organizationsAdaptivePolicyGroups, response2)
		diags = resp.State.Set(ctx, &organizationsAdaptivePolicyGroups)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsAdaptivePolicyGroups struct {
	OrganizationID types.String                                                    `tfsdk:"organization_id"`
	ID             types.String                                                    `tfsdk:"id"`
	Items          *[]ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroups `tfsdk:"items"`
	Item           *ResponseOrganizationsGetOrganizationAdaptivePolicyGroup        `tfsdk:"item"`
}

type ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroups struct {
	CreatedAt          types.String                                                                 `tfsdk:"created_at"`
	Description        types.String                                                                 `tfsdk:"description"`
	GroupID            types.String                                                                 `tfsdk:"group_id"`
	IsDefaultGroup     types.Bool                                                                   `tfsdk:"is_default_group"`
	Name               types.String                                                                 `tfsdk:"name"`
	PolicyObjects      *[]ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroupsPolicyObjects `tfsdk:"policy_objects"`
	RequiredIPMappings types.List                                                                   `tfsdk:"required_ip_mappings"`
	Sgt                types.Int64                                                                  `tfsdk:"sgt"`
	UpdatedAt          types.String                                                                 `tfsdk:"updated_at"`
}

type ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroupsPolicyObjects struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseOrganizationsGetOrganizationAdaptivePolicyGroup struct {
	CreatedAt          types.String                                                            `tfsdk:"created_at"`
	Description        types.String                                                            `tfsdk:"description"`
	GroupID            types.String                                                            `tfsdk:"group_id"`
	IsDefaultGroup     types.Bool                                                              `tfsdk:"is_default_group"`
	Name               types.String                                                            `tfsdk:"name"`
	PolicyObjects      *[]ResponseOrganizationsGetOrganizationAdaptivePolicyGroupPolicyObjects `tfsdk:"policy_objects"`
	RequiredIPMappings types.List                                                              `tfsdk:"required_ip_mappings"`
	Sgt                types.Int64                                                             `tfsdk:"sgt"`
	UpdatedAt          types.String                                                            `tfsdk:"updated_at"`
}

type ResponseOrganizationsGetOrganizationAdaptivePolicyGroupPolicyObjects struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// ToBody
func ResponseOrganizationsGetOrganizationAdaptivePolicyGroupsItemsToBody(state OrganizationsAdaptivePolicyGroups, response *merakigosdk.ResponseOrganizationsGetOrganizationAdaptivePolicyGroups) OrganizationsAdaptivePolicyGroups {
	var items []ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroups
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroups{
			CreatedAt:   types.StringValue(item.CreatedAt),
			Description: types.StringValue(item.Description),
			GroupID:     types.StringValue(item.GroupID),
			IsDefaultGroup: func() types.Bool {
				if item.IsDefaultGroup != nil {
					return types.BoolValue(*item.IsDefaultGroup)
				}
				return types.Bool{}
			}(),
			Name: types.StringValue(item.Name),
			PolicyObjects: func() *[]ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroupsPolicyObjects {
				if item.PolicyObjects != nil {
					result := make([]ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroupsPolicyObjects, len(*item.PolicyObjects))
					for i, policyObjects := range *item.PolicyObjects {
						result[i] = ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroupsPolicyObjects{
							ID:   types.StringValue(policyObjects.ID),
							Name: types.StringValue(policyObjects.Name),
						}
					}
					return &result
				}
				return &[]ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroupsPolicyObjects{}
			}(),
			RequiredIPMappings: StringSliceToList(item.RequiredIPMappings),
			Sgt: func() types.Int64 {
				if item.Sgt != nil {
					return types.Int64Value(int64(*item.Sgt))
				}
				return types.Int64{}
			}(),
			UpdatedAt: types.StringValue(item.UpdatedAt),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseOrganizationsGetOrganizationAdaptivePolicyGroupItemToBody(state OrganizationsAdaptivePolicyGroups, response *merakigosdk.ResponseOrganizationsGetOrganizationAdaptivePolicyGroup) OrganizationsAdaptivePolicyGroups {
	itemState := ResponseOrganizationsGetOrganizationAdaptivePolicyGroup{
		CreatedAt:   types.StringValue(response.CreatedAt),
		Description: types.StringValue(response.Description),
		GroupID:     types.StringValue(response.GroupID),
		IsDefaultGroup: func() types.Bool {
			if response.IsDefaultGroup != nil {
				return types.BoolValue(*response.IsDefaultGroup)
			}
			return types.Bool{}
		}(),
		Name: types.StringValue(response.Name),
		PolicyObjects: func() *[]ResponseOrganizationsGetOrganizationAdaptivePolicyGroupPolicyObjects {
			if response.PolicyObjects != nil {
				result := make([]ResponseOrganizationsGetOrganizationAdaptivePolicyGroupPolicyObjects, len(*response.PolicyObjects))
				for i, policyObjects := range *response.PolicyObjects {
					result[i] = ResponseOrganizationsGetOrganizationAdaptivePolicyGroupPolicyObjects{
						ID:   types.StringValue(policyObjects.ID),
						Name: types.StringValue(policyObjects.Name),
					}
				}
				return &result
			}
			return &[]ResponseOrganizationsGetOrganizationAdaptivePolicyGroupPolicyObjects{}
		}(),
		RequiredIPMappings: StringSliceToList(response.RequiredIPMappings),
		Sgt: func() types.Int64 {
			if response.Sgt != nil {
				return types.Int64Value(int64(*response.Sgt))
			}
			return types.Int64{}
		}(),
		UpdatedAt: types.StringValue(response.UpdatedAt),
	}
	state.Item = &itemState
	return state
}
