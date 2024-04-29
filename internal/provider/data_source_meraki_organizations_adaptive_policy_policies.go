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
	_ datasource.DataSource              = &OrganizationsAdaptivePolicyPoliciesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsAdaptivePolicyPoliciesDataSource{}
)

func NewOrganizationsAdaptivePolicyPoliciesDataSource() datasource.DataSource {
	return &OrganizationsAdaptivePolicyPoliciesDataSource{}
}

type OrganizationsAdaptivePolicyPoliciesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsAdaptivePolicyPoliciesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsAdaptivePolicyPoliciesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_adaptive_policy_policies"
}

func (d *OrganizationsAdaptivePolicyPoliciesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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

					"acls": schema.SetNestedAttribute{
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
					"adaptive_policy_id": schema.StringAttribute{
						Computed: true,
					},
					"created_at": schema.StringAttribute{
						Computed: true,
					},
					"destination_group": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"id": schema.StringAttribute{
								Computed: true,
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
							"sgt": schema.Int64Attribute{
								Computed: true,
							},
						},
					},
					"last_entry_rule": schema.StringAttribute{
						Computed: true,
					},
					"source_group": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"id": schema.StringAttribute{
								Computed: true,
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
							"sgt": schema.Int64Attribute{
								Computed: true,
							},
						},
					},
					"updated_at": schema.StringAttribute{
						Computed: true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationAdaptivePolicyPolicies`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"acls": schema.SetNestedAttribute{
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
						"adaptive_policy_id": schema.StringAttribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Computed: true,
						},
						"destination_group": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"sgt": schema.Int64Attribute{
									Computed: true,
								},
							},
						},
						"last_entry_rule": schema.StringAttribute{
							Computed: true,
						},
						"source_group": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"sgt": schema.Int64Attribute{
									Computed: true,
								},
							},
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

func (d *OrganizationsAdaptivePolicyPoliciesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsAdaptivePolicyPolicies OrganizationsAdaptivePolicyPolicies
	diags := req.Config.Get(ctx, &organizationsAdaptivePolicyPolicies)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsAdaptivePolicyPolicies.OrganizationID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsAdaptivePolicyPolicies.OrganizationID.IsNull(), !organizationsAdaptivePolicyPolicies.ID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationAdaptivePolicyPolicies")
		vvOrganizationID := organizationsAdaptivePolicyPolicies.OrganizationID.ValueString()

		response1, restyResp1, err := d.client.Organizations.GetOrganizationAdaptivePolicyPolicies(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAdaptivePolicyPolicies",
				err.Error(),
			)
			return
		}

		organizationsAdaptivePolicyPolicies = ResponseOrganizationsGetOrganizationAdaptivePolicyPoliciesItemsToBody(organizationsAdaptivePolicyPolicies, response1)
		diags = resp.State.Set(ctx, &organizationsAdaptivePolicyPolicies)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationAdaptivePolicyPolicy")
		vvOrganizationID := organizationsAdaptivePolicyPolicies.OrganizationID.ValueString()
		vvID := organizationsAdaptivePolicyPolicies.ID.ValueString()

		response2, restyResp2, err := d.client.Organizations.GetOrganizationAdaptivePolicyPolicy(vvOrganizationID, vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAdaptivePolicyPolicy",
				err.Error(),
			)
			return
		}

		organizationsAdaptivePolicyPolicies = ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyItemToBody(organizationsAdaptivePolicyPolicies, response2)
		diags = resp.State.Set(ctx, &organizationsAdaptivePolicyPolicies)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsAdaptivePolicyPolicies struct {
	OrganizationID types.String                                                      `tfsdk:"organization_id"`
	ID             types.String                                                      `tfsdk:"id"`
	Items          *[]ResponseItemOrganizationsGetOrganizationAdaptivePolicyPolicies `tfsdk:"items"`
	Item           *ResponseOrganizationsGetOrganizationAdaptivePolicyPolicy         `tfsdk:"item"`
}

type ResponseItemOrganizationsGetOrganizationAdaptivePolicyPolicies struct {
	ACLs             *[]ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesAcls           `tfsdk:"acls"`
	AdaptivePolicyID types.String                                                                    `tfsdk:"adaptive_policy_id"`
	CreatedAt        types.String                                                                    `tfsdk:"created_at"`
	DestinationGroup *ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesDestinationGroup `tfsdk:"destination_group"`
	LastEntryRule    types.String                                                                    `tfsdk:"last_entry_rule"`
	SourceGroup      *ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesSourceGroup      `tfsdk:"source_group"`
	UpdatedAt        types.String                                                                    `tfsdk:"updated_at"`
}

type ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesAcls struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesDestinationGroup struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Sgt  types.Int64  `tfsdk:"sgt"`
}

type ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesSourceGroup struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Sgt  types.Int64  `tfsdk:"sgt"`
}

type ResponseOrganizationsGetOrganizationAdaptivePolicyPolicy struct {
	ACLs             *[]ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyAcls           `tfsdk:"acls"`
	AdaptivePolicyID types.String                                                              `tfsdk:"adaptive_policy_id"`
	CreatedAt        types.String                                                              `tfsdk:"created_at"`
	DestinationGroup *ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyDestinationGroup `tfsdk:"destination_group"`
	LastEntryRule    types.String                                                              `tfsdk:"last_entry_rule"`
	SourceGroup      *ResponseOrganizationsGetOrganizationAdaptivePolicyPolicySourceGroup      `tfsdk:"source_group"`
	UpdatedAt        types.String                                                              `tfsdk:"updated_at"`
}

type ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyAcls struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyDestinationGroup struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Sgt  types.Int64  `tfsdk:"sgt"`
}

type ResponseOrganizationsGetOrganizationAdaptivePolicyPolicySourceGroup struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Sgt  types.Int64  `tfsdk:"sgt"`
}

// ToBody
func ResponseOrganizationsGetOrganizationAdaptivePolicyPoliciesItemsToBody(state OrganizationsAdaptivePolicyPolicies, response *merakigosdk.ResponseOrganizationsGetOrganizationAdaptivePolicyPolicies) OrganizationsAdaptivePolicyPolicies {
	var items []ResponseItemOrganizationsGetOrganizationAdaptivePolicyPolicies
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationAdaptivePolicyPolicies{
			ACLs: func() *[]ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesAcls {
				if item.ACLs != nil {
					result := make([]ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesAcls, len(*item.ACLs))
					for i, aCLs := range *item.ACLs {
						result[i] = ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesAcls{
							ID:   types.StringValue(aCLs.ID),
							Name: types.StringValue(aCLs.Name),
						}
					}
					return &result
				}
				return &[]ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesAcls{}
			}(),
			AdaptivePolicyID: types.StringValue(item.AdaptivePolicyID),
			CreatedAt:        types.StringValue(item.CreatedAt),
			DestinationGroup: func() *ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesDestinationGroup {
				if item.DestinationGroup != nil {
					return &ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesDestinationGroup{
						ID:   types.StringValue(item.DestinationGroup.ID),
						Name: types.StringValue(item.DestinationGroup.Name),
						Sgt: func() types.Int64 {
							if item.DestinationGroup.Sgt != nil {
								return types.Int64Value(int64(*item.DestinationGroup.Sgt))
							}
							return types.Int64{}
						}(),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesDestinationGroup{}
			}(),
			LastEntryRule: types.StringValue(item.LastEntryRule),
			SourceGroup: func() *ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesSourceGroup {
				if item.SourceGroup != nil {
					return &ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesSourceGroup{
						ID:   types.StringValue(item.SourceGroup.ID),
						Name: types.StringValue(item.SourceGroup.Name),
						Sgt: func() types.Int64 {
							if item.SourceGroup.Sgt != nil {
								return types.Int64Value(int64(*item.SourceGroup.Sgt))
							}
							return types.Int64{}
						}(),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationAdaptivePolicyPoliciesSourceGroup{}
			}(),
			UpdatedAt: types.StringValue(item.UpdatedAt),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyItemToBody(state OrganizationsAdaptivePolicyPolicies, response *merakigosdk.ResponseOrganizationsGetOrganizationAdaptivePolicyPolicy) OrganizationsAdaptivePolicyPolicies {
	itemState := ResponseOrganizationsGetOrganizationAdaptivePolicyPolicy{
		ACLs: func() *[]ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyAcls {
			if response.ACLs != nil {
				result := make([]ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyAcls, len(*response.ACLs))
				for i, aCLs := range *response.ACLs {
					result[i] = ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyAcls{
						ID:   types.StringValue(aCLs.ID),
						Name: types.StringValue(aCLs.Name),
					}
				}
				return &result
			}
			return &[]ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyAcls{}
		}(),
		AdaptivePolicyID: types.StringValue(response.AdaptivePolicyID),
		CreatedAt:        types.StringValue(response.CreatedAt),
		DestinationGroup: func() *ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyDestinationGroup {
			if response.DestinationGroup != nil {
				return &ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyDestinationGroup{
					ID:   types.StringValue(response.DestinationGroup.ID),
					Name: types.StringValue(response.DestinationGroup.Name),
					Sgt: func() types.Int64 {
						if response.DestinationGroup.Sgt != nil {
							return types.Int64Value(int64(*response.DestinationGroup.Sgt))
						}
						return types.Int64{}
					}(),
				}
			}
			return &ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyDestinationGroup{}
		}(),
		LastEntryRule: types.StringValue(response.LastEntryRule),
		SourceGroup: func() *ResponseOrganizationsGetOrganizationAdaptivePolicyPolicySourceGroup {
			if response.SourceGroup != nil {
				return &ResponseOrganizationsGetOrganizationAdaptivePolicyPolicySourceGroup{
					ID:   types.StringValue(response.SourceGroup.ID),
					Name: types.StringValue(response.SourceGroup.Name),
					Sgt: func() types.Int64 {
						if response.SourceGroup.Sgt != nil {
							return types.Int64Value(int64(*response.SourceGroup.Sgt))
						}
						return types.Int64{}
					}(),
				}
			}
			return &ResponseOrganizationsGetOrganizationAdaptivePolicyPolicySourceGroup{}
		}(),
		UpdatedAt: types.StringValue(response.UpdatedAt),
	}
	state.Item = &itemState
	return state
}
