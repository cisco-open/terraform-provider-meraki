package provider

// RESOURCE ACTION

import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsInventoryClaimResource{}
	_ resource.ResourceWithConfigure = &OrganizationsInventoryClaimResource{}
)

func NewOrganizationsInventoryClaimResource() resource.Resource {
	return &OrganizationsInventoryClaimResource{}
}

type OrganizationsInventoryClaimResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsInventoryClaimResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsInventoryClaimResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_inventory_claim"
}

// resourceAction
func (r *OrganizationsInventoryClaimResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

					"licenses": schema.SetNestedAttribute{
						MarkdownDescription: `The licenses claimed`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"key": schema.StringAttribute{
									MarkdownDescription: `The key of the license`,
									Computed:            true,
								},
								"mode": schema.StringAttribute{
									MarkdownDescription: `The mode of the license`,
									Computed:            true,
								},
							},
						},
					},
					"orders": schema.SetAttribute{
						MarkdownDescription: `The numbers of the orders claimed`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"serials": schema.SetAttribute{
						MarkdownDescription: `The serials of the devices claimed`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"licenses": schema.SetNestedAttribute{
						MarkdownDescription: `The licenses that should be claimed`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"key": schema.StringAttribute{
									MarkdownDescription: `The key of the license`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"mode": schema.StringAttribute{
									MarkdownDescription: `Co-term licensing only: either 'renew' or 'addDevices'. 'addDevices' will increase the license limit, while 'renew' will extend the amount of time until expiration. Defaults to 'addDevices'. All licenses must be claimed with the same mode, and at most one renewal can be claimed at a time. Does not apply to organizations using per-device licensing model.`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
							},
						},
					},
					"orders": schema.SetAttribute{
						MarkdownDescription: `The numbers of the orders that should be claimed`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"serials": schema.SetAttribute{
						MarkdownDescription: `The serials of the devices that should be claimed`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}
func (r *OrganizationsInventoryClaimResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsInventoryClaim

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
	response, restyResp1, err := r.client.Organizations.ClaimIntoOrganizationInventory(vvOrganizationID, dataRequest)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing ClaimIntoOrganizationInventory",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing ClaimIntoOrganizationInventory",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseOrganizationsClaimIntoOrganizationInventoryItemToBody(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsInventoryClaimResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsInventoryClaimResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsInventoryClaimResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsInventoryClaim struct {
	OrganizationID types.String                                          `tfsdk:"organization_id"`
	Item           *ResponseOrganizationsClaimIntoOrganizationInventory  `tfsdk:"item"`
	Parameters     *RequestOrganizationsClaimIntoOrganizationInventoryRs `tfsdk:"parameters"`
}

type ResponseOrganizationsClaimIntoOrganizationInventory struct {
	Licenses *[]ResponseOrganizationsClaimIntoOrganizationInventoryLicenses `tfsdk:"licenses"`
	Orders   types.Set                                                      `tfsdk:"orders"`
	Serials  types.Set                                                      `tfsdk:"serials"`
}

type ResponseOrganizationsClaimIntoOrganizationInventoryLicenses struct {
	Key  types.String `tfsdk:"key"`
	Mode types.String `tfsdk:"mode"`
}

type RequestOrganizationsClaimIntoOrganizationInventoryRs struct {
	Licenses *[]RequestOrganizationsClaimIntoOrganizationInventoryLicensesRs `tfsdk:"licenses"`
	Orders   types.Set                                                       `tfsdk:"orders"`
	Serials  types.Set                                                       `tfsdk:"serials"`
}

type RequestOrganizationsClaimIntoOrganizationInventoryLicensesRs struct {
	Key  types.String `tfsdk:"key"`
	Mode types.String `tfsdk:"mode"`
}

// FromBody
func (r *OrganizationsInventoryClaim) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsClaimIntoOrganizationInventory {
	re := *r.Parameters
	var requestOrganizationsClaimIntoOrganizationInventoryLicenses []merakigosdk.RequestOrganizationsClaimIntoOrganizationInventoryLicenses
	if re.Licenses != nil {
		for _, rItem1 := range *re.Licenses {
			key := rItem1.Key.ValueString()
			mode := rItem1.Mode.ValueString()
			requestOrganizationsClaimIntoOrganizationInventoryLicenses = append(requestOrganizationsClaimIntoOrganizationInventoryLicenses, merakigosdk.RequestOrganizationsClaimIntoOrganizationInventoryLicenses{
				Key:  key,
				Mode: mode,
			})
		}
	}
	var orders []string = nil
	re.Orders.ElementsAs(ctx, &orders, false)
	var serials []string = nil
	re.Serials.ElementsAs(ctx, &serials, false)
	out := merakigosdk.RequestOrganizationsClaimIntoOrganizationInventory{
		Licenses: func() *[]merakigosdk.RequestOrganizationsClaimIntoOrganizationInventoryLicenses {
			if len(requestOrganizationsClaimIntoOrganizationInventoryLicenses) > 0 {
				return &requestOrganizationsClaimIntoOrganizationInventoryLicenses
			}
			return nil
		}(),
		Orders:  orders,
		Serials: serials,
	}
	return &out
}

// ToBody
func ResponseOrganizationsClaimIntoOrganizationInventoryItemToBody(state OrganizationsInventoryClaim, response *merakigosdk.ResponseOrganizationsClaimIntoOrganizationInventory) OrganizationsInventoryClaim {
	itemState := ResponseOrganizationsClaimIntoOrganizationInventory{
		Licenses: func() *[]ResponseOrganizationsClaimIntoOrganizationInventoryLicenses {
			if response.Licenses != nil {
				result := make([]ResponseOrganizationsClaimIntoOrganizationInventoryLicenses, len(*response.Licenses))
				for i, licenses := range *response.Licenses {
					result[i] = ResponseOrganizationsClaimIntoOrganizationInventoryLicenses{
						Key:  types.StringValue(licenses.Key),
						Mode: types.StringValue(licenses.Mode),
					}
				}
				return &result
			}
			return &[]ResponseOrganizationsClaimIntoOrganizationInventoryLicenses{}
		}(),
		Orders:  StringSliceToSet(response.Orders),
		Serials: StringSliceToSet(response.Serials),
	}
	state.Item = &itemState
	return state
}
