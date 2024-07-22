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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsCameraCustomAnalyticsArtifactsResource{}
	_ resource.ResourceWithConfigure = &OrganizationsCameraCustomAnalyticsArtifactsResource{}
)

func NewOrganizationsCameraCustomAnalyticsArtifactsResource() resource.Resource {
	return &OrganizationsCameraCustomAnalyticsArtifactsResource{}
}

type OrganizationsCameraCustomAnalyticsArtifactsResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsCameraCustomAnalyticsArtifactsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsCameraCustomAnalyticsArtifactsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_camera_custom_analytics_artifacts"
}

func (r *OrganizationsCameraCustomAnalyticsArtifactsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"artifact_id": schema.StringAttribute{
				MarkdownDescription: `Custom analytics artifact ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Custom analytics artifact name`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `Organization ID`,
				Required:            true,
			},
			"status": schema.SingleNestedAttribute{
				MarkdownDescription: `Custom analytics artifact status`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{

					"message": schema.StringAttribute{
						MarkdownDescription: `Status message`,
						Computed:            true,
					},
					"type": schema.StringAttribute{
						MarkdownDescription: `Status type`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (r *OrganizationsCameraCustomAnalyticsArtifactsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsCameraCustomAnalyticsArtifactsRs

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
	// organization_id
	vvName := data.Name.ValueString()
	// name
	responseVerifyItem, restyResp1, err := r.client.Camera.GetOrganizationCameraCustomAnalyticsArtifacts(vvOrganizationID)
	//Have Create
	if err != nil {
		if restyResp1 != nil {
			if restyResp1 == nil {
				resp.Diagnostics.AddError(
					"Failure when executing Get",
					err.Error(),
				)
				return
			}
			if restyResp1.StatusCode() != 404 {
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCameraCustomAnalyticsArtifacts",
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
			vvArtifactID, ok := result2["ArtifactID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter ArtifactID",
					"Fail Parsing ArtifactID",
				)
				return
			}
			// r.client..( data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Camera.GetOrganizationCameraCustomAnalyticsArtifact(vvOrganizationID, vvArtifactID)
			if responseVerifyItem2 != nil {
				data = ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	_, restyResp2, err := r.client.Camera.CreateOrganizationCameraCustomAnalyticsArtifact(vvOrganizationID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationCameraCustomAnalyticsArtifact",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationCameraCustomAnalyticsArtifact",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Camera.GetOrganizationCameraCustomAnalyticsArtifacts(vvOrganizationID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCameraCustomAnalyticsArtifacts",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationCameraCustomAnalyticsArtifacts",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvArtifactID, ok := result2["ArtifactID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter ArtifactID",
				err.Error(),
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Camera.GetOrganizationCameraCustomAnalyticsArtifact(vvOrganizationID, vvArtifactID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationCameraCustomAnalyticsArtifact",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCameraCustomAnalyticsArtifact",
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

func (r *OrganizationsCameraCustomAnalyticsArtifactsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsCameraCustomAnalyticsArtifactsRs

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
	vvArtifactID := data.ArtifactID.ValueString()
	responseGet, restyRespGet, err := r.client.Camera.GetOrganizationCameraCustomAnalyticsArtifact(vvOrganizationID, vvArtifactID)
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
				"Failure when executing GetOrganizationCameraCustomAnalyticsArtifact",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationCameraCustomAnalyticsArtifact",
			err.Error(),
		)
		return
	}

	data = ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsCameraCustomAnalyticsArtifactsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("artifact_id"), idParts[1])...)
}

func (r *OrganizationsCameraCustomAnalyticsArtifactsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsCameraCustomAnalyticsArtifactsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update
	// No update
	resp.Diagnostics.AddError(
		"Update operation not supported in OrganizationsCameraCustomAnalyticsArtifacts",
		"Update operation not supported in OrganizationsCameraCustomAnalyticsArtifacts",
	)
	return
}

func (r *OrganizationsCameraCustomAnalyticsArtifactsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationsCameraCustomAnalyticsArtifactsRs
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
	vvArtifactID := state.ArtifactID.ValueString()
	_, err := r.client.Camera.DeleteOrganizationCameraCustomAnalyticsArtifact(vvOrganizationID, vvArtifactID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganizationCameraCustomAnalyticsArtifact", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type OrganizationsCameraCustomAnalyticsArtifactsRs struct {
	OrganizationID types.String                                                        `tfsdk:"organization_id"`
	ArtifactID     types.String                                                        `tfsdk:"artifact_id"`
	Name           types.String                                                        `tfsdk:"name"`
	Status         *ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactStatusRs `tfsdk:"status"`
}

type ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactStatusRs struct {
	Message types.String `tfsdk:"message"`
	Type    types.String `tfsdk:"type"`
}

// FromBody
func (r *OrganizationsCameraCustomAnalyticsArtifactsRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestCameraCreateOrganizationCameraCustomAnalyticsArtifact {
	emptyString := ""
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	out := merakigosdk.RequestCameraCreateOrganizationCameraCustomAnalyticsArtifact{
		Name: *name,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactItemToBodyRs(state OrganizationsCameraCustomAnalyticsArtifactsRs, response *merakigosdk.ResponseCameraGetOrganizationCameraCustomAnalyticsArtifact, is_read bool) OrganizationsCameraCustomAnalyticsArtifactsRs {
	itemState := OrganizationsCameraCustomAnalyticsArtifactsRs{
		ArtifactID:     types.StringValue(response.ArtifactID),
		Name:           types.StringValue(response.Name),
		OrganizationID: types.StringValue(response.OrganizationID),
		Status: func() *ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactStatusRs {
			if response.Status != nil {
				return &ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactStatusRs{
					Message: types.StringValue(response.Status.Message),
					Type:    types.StringValue(response.Status.Type),
				}
			}
			return &ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactStatusRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsCameraCustomAnalyticsArtifactsRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsCameraCustomAnalyticsArtifactsRs)
}
