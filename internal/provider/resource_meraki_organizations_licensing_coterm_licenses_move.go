package provider

// RESOURCE ACTION

import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsLicensingCotermLicensesMoveResource{}
	_ resource.ResourceWithConfigure = &OrganizationsLicensingCotermLicensesMoveResource{}
)

func NewOrganizationsLicensingCotermLicensesMoveResource() resource.Resource {
	return &OrganizationsLicensingCotermLicensesMoveResource{}
}

type OrganizationsLicensingCotermLicensesMoveResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsLicensingCotermLicensesMoveResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsLicensingCotermLicensesMoveResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_licensing_coterm_licenses_move"
}

// resourceAction
func (r *OrganizationsLicensingCotermLicensesMoveResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"moved_licenses": schema.SetNestedAttribute{
						MarkdownDescription: `Newly moved licenses created in the destination organization of the license move operation`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"claimed_at": schema.StringAttribute{
									MarkdownDescription: `When the license was claimed into the organization`,
									Computed:            true,
								},
								"counts": schema.SetNestedAttribute{
									MarkdownDescription: `The counts of the license by model type`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"count": schema.Int64Attribute{
												MarkdownDescription: `The number of counts the license contains of this model`,
												Computed:            true,
											},
											"model": schema.StringAttribute{
												MarkdownDescription: `The license model type`,
												Computed:            true,
											},
										},
									},
								},
								"duration": schema.Int64Attribute{
									MarkdownDescription: `The duration (term length) of the license, measured in days`,
									Computed:            true,
								},
								"editions": schema.SetNestedAttribute{
									MarkdownDescription: `The editions of the license for each relevant product type`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"edition": schema.StringAttribute{
												MarkdownDescription: `The name of the license edition`,
												Computed:            true,
											},
											"product_type": schema.StringAttribute{
												MarkdownDescription: `The product type of the license edition`,
												Computed:            true,
											},
										},
									},
								},
								"expired": schema.BoolAttribute{
									MarkdownDescription: `Flag to indicate if the license is expired`,
									Computed:            true,
								},
								"invalidated": schema.BoolAttribute{
									MarkdownDescription: `Flag to indicated that the license is invalidated`,
									Computed:            true,
								},
								"invalidated_at": schema.StringAttribute{
									MarkdownDescription: `When the license was invalidated. Will be null for active licenses`,
									Computed:            true,
								},
								"key": schema.StringAttribute{
									MarkdownDescription: `The key of the license`,
									Computed:            true,
								},
								"mode": schema.StringAttribute{
									MarkdownDescription: `The operation mode of the license when it was claimed`,
									Computed:            true,
								},
								"organization_id": schema.StringAttribute{
									MarkdownDescription: `The ID of the organization that the license is claimed in`,
									Computed:            true,
								},
								"started_at": schema.StringAttribute{
									MarkdownDescription: `When the license's term began (approximately the date when the license was created)`,
									Computed:            true,
								},
							},
						},
					},
					"remainder_licenses": schema.SetNestedAttribute{
						MarkdownDescription: `Remainder licenses created in the source organization as a result of moving a subset of the counts of a license`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"claimed_at": schema.StringAttribute{
									MarkdownDescription: `When the license was claimed into the organization`,
									Computed:            true,
								},
								"counts": schema.SetNestedAttribute{
									MarkdownDescription: `The counts of the license by model type`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"count": schema.Int64Attribute{
												MarkdownDescription: `The number of counts the license contains of this model`,
												Computed:            true,
											},
											"model": schema.StringAttribute{
												MarkdownDescription: `The license model type`,
												Computed:            true,
											},
										},
									},
								},
								"duration": schema.Int64Attribute{
									MarkdownDescription: `The duration (term length) of the license, measured in days`,
									Computed:            true,
								},
								"editions": schema.SetNestedAttribute{
									MarkdownDescription: `The editions of the license for each relevant product type`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"edition": schema.StringAttribute{
												MarkdownDescription: `The name of the license edition`,
												Computed:            true,
											},
											"product_type": schema.StringAttribute{
												MarkdownDescription: `The product type of the license edition`,
												Computed:            true,
											},
										},
									},
								},
								"expired": schema.BoolAttribute{
									MarkdownDescription: `Flag to indicate if the license is expired`,
									Computed:            true,
								},
								"invalidated": schema.BoolAttribute{
									MarkdownDescription: `Flag to indicated that the license is invalidated`,
									Computed:            true,
								},
								"invalidated_at": schema.StringAttribute{
									MarkdownDescription: `When the license was invalidated. Will be null for active licenses`,
									Computed:            true,
								},
								"key": schema.StringAttribute{
									MarkdownDescription: `The key of the license`,
									Computed:            true,
								},
								"mode": schema.StringAttribute{
									MarkdownDescription: `The operation mode of the license when it was claimed`,
									Computed:            true,
								},
								"organization_id": schema.StringAttribute{
									MarkdownDescription: `The ID of the organization that the license is claimed in`,
									Computed:            true,
								},
								"started_at": schema.StringAttribute{
									MarkdownDescription: `When the license's term began (approximately the date when the license was created)`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"destination": schema.SingleNestedAttribute{
						MarkdownDescription: `Destination data for the license move`,
						Optional:            true,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"mode": schema.StringAttribute{
								MarkdownDescription: `The claim mode of the moved license`,
								Optional:            true,
								Computed:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
							},
							"organization_id": schema.StringAttribute{
								MarkdownDescription: `The organization to move the license to`,
								Optional:            true,
								Computed:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
							},
						},
					},
					"licenses": schema.SetNestedAttribute{
						MarkdownDescription: `The list of licenses to move`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"counts": schema.SetNestedAttribute{
									MarkdownDescription: `The counts to move from the license by model type`,
									Optional:            true,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"count": schema.Int64Attribute{
												MarkdownDescription: `The number of counts to move`,
												Optional:            true,
												Computed:            true,
												PlanModifiers: []planmodifier.Int64{
													int64planmodifier.RequiresReplace(),
												},
											},
											"model": schema.StringAttribute{
												MarkdownDescription: `The license model type to move counts of`,
												Optional:            true,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
										},
									},
								},
								"key": schema.StringAttribute{
									MarkdownDescription: `The license key to move counts from`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
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
func (r *OrganizationsLicensingCotermLicensesMoveResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsLicensingCotermLicensesMove

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
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Licensing.MoveOrganizationLicensingCotermLicenses(vvOrganizationID, dataRequest)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing MoveOrganizationLicensingCotermLicenses",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing MoveOrganizationLicensingCotermLicenses",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseLicensingMoveOrganizationLicensingCotermLicensesItemToBody(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsLicensingCotermLicensesMoveResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsLicensingCotermLicensesMoveResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsLicensingCotermLicensesMoveResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsLicensingCotermLicensesMove struct {
	OrganizationID types.String                                               `tfsdk:"organization_id"`
	Item           *ResponseLicensingMoveOrganizationLicensingCotermLicenses  `tfsdk:"item"`
	Parameters     *RequestLicensingMoveOrganizationLicensingCotermLicensesRs `tfsdk:"parameters"`
}

type ResponseLicensingMoveOrganizationLicensingCotermLicenses struct {
	MovedLicenses     *[]ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicenses     `tfsdk:"moved_licenses"`
	RemainderLicenses *[]ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicenses `tfsdk:"remainder_licenses"`
}

type ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicenses struct {
	ClaimedAt      types.String                                                                     `tfsdk:"claimed_at"`
	Counts         *[]ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicensesCounts   `tfsdk:"counts"`
	Duration       types.Int64                                                                      `tfsdk:"duration"`
	Editions       *[]ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicensesEditions `tfsdk:"editions"`
	Expired        types.Bool                                                                       `tfsdk:"expired"`
	Invalidated    types.Bool                                                                       `tfsdk:"invalidated"`
	InvalidatedAt  types.String                                                                     `tfsdk:"invalidated_at"`
	Key            types.String                                                                     `tfsdk:"key"`
	Mode           types.String                                                                     `tfsdk:"mode"`
	OrganizationID types.String                                                                     `tfsdk:"organization_id"`
	StartedAt      types.String                                                                     `tfsdk:"started_at"`
}

type ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicensesCounts struct {
	Count types.Int64  `tfsdk:"count"`
	Model types.String `tfsdk:"model"`
}

type ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicensesEditions struct {
	Edition     types.String `tfsdk:"edition"`
	ProductType types.String `tfsdk:"product_type"`
}

type ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicenses struct {
	ClaimedAt      types.String                                                                         `tfsdk:"claimed_at"`
	Counts         *[]ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicensesCounts   `tfsdk:"counts"`
	Duration       types.Int64                                                                          `tfsdk:"duration"`
	Editions       *[]ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicensesEditions `tfsdk:"editions"`
	Expired        types.Bool                                                                           `tfsdk:"expired"`
	Invalidated    types.Bool                                                                           `tfsdk:"invalidated"`
	InvalidatedAt  types.String                                                                         `tfsdk:"invalidated_at"`
	Key            types.String                                                                         `tfsdk:"key"`
	Mode           types.String                                                                         `tfsdk:"mode"`
	OrganizationID types.String                                                                         `tfsdk:"organization_id"`
	StartedAt      types.String                                                                         `tfsdk:"started_at"`
}

type ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicensesCounts struct {
	Count types.Int64  `tfsdk:"count"`
	Model types.String `tfsdk:"model"`
}

type ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicensesEditions struct {
	Edition     types.String `tfsdk:"edition"`
	ProductType types.String `tfsdk:"product_type"`
}

type RequestLicensingMoveOrganizationLicensingCotermLicensesRs struct {
	Destination *RequestLicensingMoveOrganizationLicensingCotermLicensesDestinationRs `tfsdk:"destination"`
	Licenses    *[]RequestLicensingMoveOrganizationLicensingCotermLicensesLicensesRs  `tfsdk:"licenses"`
}

type RequestLicensingMoveOrganizationLicensingCotermLicensesDestinationRs struct {
	Mode           types.String `tfsdk:"mode"`
	OrganizationID types.String `tfsdk:"organization_id"`
}

type RequestLicensingMoveOrganizationLicensingCotermLicensesLicensesRs struct {
	Counts *[]RequestLicensingMoveOrganizationLicensingCotermLicensesLicensesCountsRs `tfsdk:"counts"`
	Key    types.String                                                               `tfsdk:"key"`
}

type RequestLicensingMoveOrganizationLicensingCotermLicensesLicensesCountsRs struct {
	Count types.Int64  `tfsdk:"count"`
	Model types.String `tfsdk:"model"`
}

// FromBody
func (r *OrganizationsLicensingCotermLicensesMove) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestLicensingMoveOrganizationLicensingCotermLicenses {
	re := *r.Parameters
	var requestLicensingMoveOrganizationLicensingCotermLicensesDestination *merakigosdk.RequestLicensingMoveOrganizationLicensingCotermLicensesDestination
	if re.Destination != nil {
		mode := re.Destination.Mode.ValueString()
		organizationID := re.Destination.OrganizationID.ValueString()
		requestLicensingMoveOrganizationLicensingCotermLicensesDestination = &merakigosdk.RequestLicensingMoveOrganizationLicensingCotermLicensesDestination{
			Mode:           mode,
			OrganizationID: organizationID,
		}
	}
	var requestLicensingMoveOrganizationLicensingCotermLicensesLicenses []merakigosdk.RequestLicensingMoveOrganizationLicensingCotermLicensesLicenses
	if re.Licenses != nil {
		for _, rItem1 := range *re.Licenses {
			var requestLicensingMoveOrganizationLicensingCotermLicensesLicensesCounts []merakigosdk.RequestLicensingMoveOrganizationLicensingCotermLicensesLicensesCounts
			if rItem1.Counts != nil {
				for _, rItem2 := range *rItem1.Counts { //Counts// name: counts
					count := func() *int64 {
						if !rItem2.Count.IsUnknown() && !rItem2.Count.IsNull() {
							return rItem2.Count.ValueInt64Pointer()
						}
						return nil
					}()
					model := rItem2.Model.ValueString()
					requestLicensingMoveOrganizationLicensingCotermLicensesLicensesCounts = append(requestLicensingMoveOrganizationLicensingCotermLicensesLicensesCounts, merakigosdk.RequestLicensingMoveOrganizationLicensingCotermLicensesLicensesCounts{
						Count: int64ToIntPointer(count),
						Model: model,
					})
				}
			}
			key := rItem1.Key.ValueString()
			requestLicensingMoveOrganizationLicensingCotermLicensesLicenses = append(requestLicensingMoveOrganizationLicensingCotermLicensesLicenses, merakigosdk.RequestLicensingMoveOrganizationLicensingCotermLicensesLicenses{
				Counts: func() *[]merakigosdk.RequestLicensingMoveOrganizationLicensingCotermLicensesLicensesCounts {
					if len(requestLicensingMoveOrganizationLicensingCotermLicensesLicensesCounts) > 0 {
						return &requestLicensingMoveOrganizationLicensingCotermLicensesLicensesCounts
					}
					return nil
				}(),
				Key: key,
			})
		}
	}
	out := merakigosdk.RequestLicensingMoveOrganizationLicensingCotermLicenses{
		Destination: requestLicensingMoveOrganizationLicensingCotermLicensesDestination,
		Licenses: func() *[]merakigosdk.RequestLicensingMoveOrganizationLicensingCotermLicensesLicenses {
			if len(requestLicensingMoveOrganizationLicensingCotermLicensesLicenses) > 0 {
				return &requestLicensingMoveOrganizationLicensingCotermLicensesLicenses
			}
			return nil
		}(),
	}
	return &out
}

// ToBody
func ResponseLicensingMoveOrganizationLicensingCotermLicensesItemToBody(state OrganizationsLicensingCotermLicensesMove, response *merakigosdk.ResponseLicensingMoveOrganizationLicensingCotermLicenses) OrganizationsLicensingCotermLicensesMove {
	itemState := ResponseLicensingMoveOrganizationLicensingCotermLicenses{
		MovedLicenses: func() *[]ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicenses {
			if response.MovedLicenses != nil {
				result := make([]ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicenses, len(*response.MovedLicenses))
				for i, movedLicenses := range *response.MovedLicenses {
					result[i] = ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicenses{
						ClaimedAt: types.StringValue(movedLicenses.ClaimedAt),
						Counts: func() *[]ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicensesCounts {
							if movedLicenses.Counts != nil {
								result := make([]ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicensesCounts, len(*movedLicenses.Counts))
								for i, counts := range *movedLicenses.Counts {
									result[i] = ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicensesCounts{
										Count: func() types.Int64 {
											if counts.Count != nil {
												return types.Int64Value(int64(*counts.Count))
											}
											return types.Int64{}
										}(),
										Model: types.StringValue(counts.Model),
									}
								}
								return &result
							}
							return &[]ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicensesCounts{}
						}(),
						Duration: func() types.Int64 {
							if movedLicenses.Duration != nil {
								return types.Int64Value(int64(*movedLicenses.Duration))
							}
							return types.Int64{}
						}(),
						Editions: func() *[]ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicensesEditions {
							if movedLicenses.Editions != nil {
								result := make([]ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicensesEditions, len(*movedLicenses.Editions))
								for i, editions := range *movedLicenses.Editions {
									result[i] = ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicensesEditions{
										Edition:     types.StringValue(editions.Edition),
										ProductType: types.StringValue(editions.ProductType),
									}
								}
								return &result
							}
							return &[]ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicensesEditions{}
						}(),
						Expired: func() types.Bool {
							if movedLicenses.Expired != nil {
								return types.BoolValue(*movedLicenses.Expired)
							}
							return types.Bool{}
						}(),
						Invalidated: func() types.Bool {
							if movedLicenses.Invalidated != nil {
								return types.BoolValue(*movedLicenses.Invalidated)
							}
							return types.Bool{}
						}(),
						InvalidatedAt:  types.StringValue(movedLicenses.InvalidatedAt),
						Key:            types.StringValue(movedLicenses.Key),
						Mode:           types.StringValue(movedLicenses.Mode),
						OrganizationID: types.StringValue(movedLicenses.OrganizationID),
						StartedAt:      types.StringValue(movedLicenses.StartedAt),
					}
				}
				return &result
			}
			return &[]ResponseLicensingMoveOrganizationLicensingCotermLicensesMovedLicenses{}
		}(),
		RemainderLicenses: func() *[]ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicenses {
			if response.RemainderLicenses != nil {
				result := make([]ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicenses, len(*response.RemainderLicenses))
				for i, remainderLicenses := range *response.RemainderLicenses {
					result[i] = ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicenses{
						ClaimedAt: types.StringValue(remainderLicenses.ClaimedAt),
						Counts: func() *[]ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicensesCounts {
							if remainderLicenses.Counts != nil {
								result := make([]ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicensesCounts, len(*remainderLicenses.Counts))
								for i, counts := range *remainderLicenses.Counts {
									result[i] = ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicensesCounts{
										Count: func() types.Int64 {
											if counts.Count != nil {
												return types.Int64Value(int64(*counts.Count))
											}
											return types.Int64{}
										}(),
										Model: types.StringValue(counts.Model),
									}
								}
								return &result
							}
							return &[]ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicensesCounts{}
						}(),
						Duration: func() types.Int64 {
							if remainderLicenses.Duration != nil {
								return types.Int64Value(int64(*remainderLicenses.Duration))
							}
							return types.Int64{}
						}(),
						Editions: func() *[]ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicensesEditions {
							if remainderLicenses.Editions != nil {
								result := make([]ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicensesEditions, len(*remainderLicenses.Editions))
								for i, editions := range *remainderLicenses.Editions {
									result[i] = ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicensesEditions{
										Edition:     types.StringValue(editions.Edition),
										ProductType: types.StringValue(editions.ProductType),
									}
								}
								return &result
							}
							return &[]ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicensesEditions{}
						}(),
						Expired: func() types.Bool {
							if remainderLicenses.Expired != nil {
								return types.BoolValue(*remainderLicenses.Expired)
							}
							return types.Bool{}
						}(),
						Invalidated: func() types.Bool {
							if remainderLicenses.Invalidated != nil {
								return types.BoolValue(*remainderLicenses.Invalidated)
							}
							return types.Bool{}
						}(),
						InvalidatedAt:  types.StringValue(remainderLicenses.InvalidatedAt),
						Key:            types.StringValue(remainderLicenses.Key),
						Mode:           types.StringValue(remainderLicenses.Mode),
						OrganizationID: types.StringValue(remainderLicenses.OrganizationID),
						StartedAt:      types.StringValue(remainderLicenses.StartedAt),
					}
				}
				return &result
			}
			return &[]ResponseLicensingMoveOrganizationLicensingCotermLicensesRemainderLicenses{}
		}(),
	}
	state.Item = &itemState
	return state
}
