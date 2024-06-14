package provider

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsAdaptivePolicyGroupsResource{}
	_ resource.ResourceWithConfigure = &OrganizationsAdaptivePolicyGroupsResource{}
)

func NewOrganizationsAdaptivePolicyGroupsResource() resource.Resource {
	return &OrganizationsAdaptivePolicyGroupsResource{}
}

type OrganizationsAdaptivePolicyGroupsResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsAdaptivePolicyGroupsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsAdaptivePolicyGroupsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_adaptive_policy_groups"
}

func (r *OrganizationsAdaptivePolicyGroupsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: `Description of the group (default: "")`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"group_id": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: `id path parameter.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"is_default_group": schema.BoolAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Name of the group`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"policy_objects": schema.SetNestedAttribute{
				MarkdownDescription: `The policy objects that belong to this group; traffic from addresses specified by these policy objects will be tagged with this group's SGT value if no other tagging scheme is being used (each requires one unique attribute) ()`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"id": schema.StringAttribute{
							MarkdownDescription: `The ID of the policy object`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the policy object`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"required_ip_mappings": schema.SetAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Default:     setdefault.StaticValue(types.SetNull(types.StringType)),
			},
			"sgt": schema.Int64Attribute{
				MarkdownDescription: `SGT value of the group`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

//path params to set ['id']

func (r *OrganizationsAdaptivePolicyGroupsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsAdaptivePolicyGroupsRs

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
	responseVerifyItem, restyResp1, err := r.client.Organizations.GetOrganizationAdaptivePolicyGroups(vvOrganizationID)
	//Have Create
	if err != nil || restyResp1 == nil {
		if restyResp1.StatusCode() != 404 {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAdaptivePolicyGroups",
				err.Error(),
			)
			return
		}
	}
	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvID, ok := result2["GroupID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter ID",
					err.Error(),
				)
				return
			}
			r.client.Organizations.UpdateOrganizationAdaptivePolicyGroup(vvOrganizationID, vvID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Organizations.GetOrganizationAdaptivePolicyGroup(vvOrganizationID, vvID)
			if responseVerifyItem2 != nil {
				data = ResponseOrganizationsGetOrganizationAdaptivePolicyGroupItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	restyResp2, err := r.client.Organizations.CreateOrganizationAdaptivePolicyGroup(vvOrganizationID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationAdaptivePolicyGroup",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationAdaptivePolicyGroup",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationAdaptivePolicyGroups(vvOrganizationID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAdaptivePolicyGroups",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationAdaptivePolicyGroups",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvID, ok := result2["GroupID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter ID",
				err.Error(),
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Organizations.GetOrganizationAdaptivePolicyGroup(vvOrganizationID, vvID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseOrganizationsGetOrganizationAdaptivePolicyGroupItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationAdaptivePolicyGroup",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAdaptivePolicyGroup",
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

func (r *OrganizationsAdaptivePolicyGroupsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsAdaptivePolicyGroupsRs

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
	// organization_id
	vvID := data.GroupID.ValueString()
	// id
	responseGet, restyRespGet, err := r.client.Organizations.GetOrganizationAdaptivePolicyGroup(vvOrganizationID, vvID)
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
				"Failure when executing GetOrganizationAdaptivePolicyGroup",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationAdaptivePolicyGroup",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseOrganizationsGetOrganizationAdaptivePolicyGroupItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *OrganizationsAdaptivePolicyGroupsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
}

func (r *OrganizationsAdaptivePolicyGroupsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsAdaptivePolicyGroupsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	// organization_id
	vvID := data.GroupID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Organizations.UpdateOrganizationAdaptivePolicyGroup(vvOrganizationID, vvID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationAdaptivePolicyGroup",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationAdaptivePolicyGroup",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsAdaptivePolicyGroupsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationsAdaptivePolicyGroupsRs
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
	vvID := state.GroupID.ValueString()
	_, err := r.client.Organizations.DeleteOrganizationAdaptivePolicyGroup(vvOrganizationID, vvID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganizationAdaptivePolicyGroup", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type OrganizationsAdaptivePolicyGroupsRs struct {
	OrganizationID     types.String                                                              `tfsdk:"organization_id"`
	ID                 types.String                                                              `tfsdk:"id"`
	CreatedAt          types.String                                                              `tfsdk:"created_at"`
	Description        types.String                                                              `tfsdk:"description"`
	GroupID            types.String                                                              `tfsdk:"group_id"`
	IsDefaultGroup     types.Bool                                                                `tfsdk:"is_default_group"`
	Name               types.String                                                              `tfsdk:"name"`
	PolicyObjects      *[]ResponseOrganizationsGetOrganizationAdaptivePolicyGroupPolicyObjectsRs `tfsdk:"policy_objects"`
	RequiredIPMappings types.Set                                                                 `tfsdk:"required_ip_mappings"`
	Sgt                types.Int64                                                               `tfsdk:"sgt"`
	UpdatedAt          types.String                                                              `tfsdk:"updated_at"`
}

type ResponseOrganizationsGetOrganizationAdaptivePolicyGroupPolicyObjectsRs struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// FromBody
func (r *OrganizationsAdaptivePolicyGroupsRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganizationAdaptivePolicyGroup {
	emptyString := ""
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestOrganizationsCreateOrganizationAdaptivePolicyGroupPolicyObjects []merakigosdk.RequestOrganizationsCreateOrganizationAdaptivePolicyGroupPolicyObjects
	if r.PolicyObjects != nil {
		for _, rItem1 := range *r.PolicyObjects {
			iD := rItem1.ID.ValueString()
			name := rItem1.Name.ValueString()
			requestOrganizationsCreateOrganizationAdaptivePolicyGroupPolicyObjects = append(requestOrganizationsCreateOrganizationAdaptivePolicyGroupPolicyObjects, merakigosdk.RequestOrganizationsCreateOrganizationAdaptivePolicyGroupPolicyObjects{
				ID:   iD,
				Name: name,
			})
		}
	}
	sgt := new(int64)
	if !r.Sgt.IsUnknown() && !r.Sgt.IsNull() {
		*sgt = r.Sgt.ValueInt64()
	} else {
		sgt = nil
	}
	out := merakigosdk.RequestOrganizationsCreateOrganizationAdaptivePolicyGroup{
		Description: *description,
		Name:        *name,
		PolicyObjects: func() *[]merakigosdk.RequestOrganizationsCreateOrganizationAdaptivePolicyGroupPolicyObjects {
			if len(requestOrganizationsCreateOrganizationAdaptivePolicyGroupPolicyObjects) > 0 {
				return &requestOrganizationsCreateOrganizationAdaptivePolicyGroupPolicyObjects
			}
			return nil
		}(),
		Sgt: int64ToIntPointer(sgt),
	}
	return &out
}
func (r *OrganizationsAdaptivePolicyGroupsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestOrganizationsUpdateOrganizationAdaptivePolicyGroup {
	emptyString := ""
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestOrganizationsUpdateOrganizationAdaptivePolicyGroupPolicyObjects []merakigosdk.RequestOrganizationsUpdateOrganizationAdaptivePolicyGroupPolicyObjects
	if r.PolicyObjects != nil {
		for _, rItem1 := range *r.PolicyObjects {
			iD := rItem1.ID.ValueString()
			name := rItem1.Name.ValueString()
			requestOrganizationsUpdateOrganizationAdaptivePolicyGroupPolicyObjects = append(requestOrganizationsUpdateOrganizationAdaptivePolicyGroupPolicyObjects, merakigosdk.RequestOrganizationsUpdateOrganizationAdaptivePolicyGroupPolicyObjects{
				ID:   iD,
				Name: name,
			})
		}
	}
	sgt := new(int64)
	if !r.Sgt.IsUnknown() && !r.Sgt.IsNull() {
		*sgt = r.Sgt.ValueInt64()
	} else {
		sgt = nil
	}
	out := merakigosdk.RequestOrganizationsUpdateOrganizationAdaptivePolicyGroup{
		Description: *description,
		Name:        *name,
		PolicyObjects: func() *[]merakigosdk.RequestOrganizationsUpdateOrganizationAdaptivePolicyGroupPolicyObjects {
			if len(requestOrganizationsUpdateOrganizationAdaptivePolicyGroupPolicyObjects) > 0 {
				return &requestOrganizationsUpdateOrganizationAdaptivePolicyGroupPolicyObjects
			}
			return nil
		}(),
		Sgt: int64ToIntPointer(sgt),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseOrganizationsGetOrganizationAdaptivePolicyGroupItemToBodyRs(state OrganizationsAdaptivePolicyGroupsRs, response *merakigosdk.ResponseOrganizationsGetOrganizationAdaptivePolicyGroup, is_read bool) OrganizationsAdaptivePolicyGroupsRs {
	itemState := OrganizationsAdaptivePolicyGroupsRs{
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
		PolicyObjects: func() *[]ResponseOrganizationsGetOrganizationAdaptivePolicyGroupPolicyObjectsRs {
			if response.PolicyObjects != nil {
				result := make([]ResponseOrganizationsGetOrganizationAdaptivePolicyGroupPolicyObjectsRs, len(*response.PolicyObjects))
				for i, policyObjects := range *response.PolicyObjects {
					result[i] = ResponseOrganizationsGetOrganizationAdaptivePolicyGroupPolicyObjectsRs{
						ID:   types.StringValue(policyObjects.ID),
						Name: types.StringValue(policyObjects.Name),
					}
				}
				return &result
			}
			return &[]ResponseOrganizationsGetOrganizationAdaptivePolicyGroupPolicyObjectsRs{}
		}(),
		RequiredIPMappings: StringSliceToSet(response.RequiredIPMappings),
		Sgt: func() types.Int64 {
			if response.Sgt != nil {
				return types.Int64Value(int64(*response.Sgt))
			}
			return types.Int64{}
		}(),
		UpdatedAt: types.StringValue(response.UpdatedAt),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsAdaptivePolicyGroupsRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsAdaptivePolicyGroupsRs)
}
