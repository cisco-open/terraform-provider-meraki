package provider

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksAppliancePrefixesDelegatedStaticsResource{}
	_ resource.ResourceWithConfigure = &NetworksAppliancePrefixesDelegatedStaticsResource{}
)

func NewNetworksAppliancePrefixesDelegatedStaticsResource() resource.Resource {
	return &NetworksAppliancePrefixesDelegatedStaticsResource{}
}

type NetworksAppliancePrefixesDelegatedStaticsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksAppliancePrefixesDelegatedStaticsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksAppliancePrefixesDelegatedStaticsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_prefixes_delegated_statics"
}

func (r *NetworksAppliancePrefixesDelegatedStaticsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"created_at": schema.StringAttribute{
				MarkdownDescription: `Prefix creation time.`,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: `Identifying description for the prefix.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"origin": schema.SingleNestedAttribute{
				MarkdownDescription: `WAN1/WAN2/Independent prefix.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"interfaces": schema.SetAttribute{
						MarkdownDescription: `Uplink provided or independent`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},

						ElementType: types.StringType,
					},
					"type": schema.StringAttribute{
						MarkdownDescription: `Origin type`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"independent",
								"internet",
							),
						},
					},
				},
			},
			"prefix": schema.StringAttribute{
				MarkdownDescription: `IPv6 prefix/prefix length.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"static_delegated_prefix_id": schema.StringAttribute{
				MarkdownDescription: `Static delegated prefix id.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: `Prefix Updated time.`,
				Computed:            true,
			},
		},
	}
}

//path params to set ['staticDelegatedPrefixId']

func (r *NetworksAppliancePrefixesDelegatedStaticsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksAppliancePrefixesDelegatedStaticsRs

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
	vvNetworkID := data.NetworkID.ValueString()
	vvPrefix := data.Prefix.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkAppliancePrefixesDelegatedStatics(vvNetworkID)
	//Have Create
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkAppliancePrefixesDelegatedStatics",
					err.Error(),
				)
				return
			}
		}
	}
	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Prefix", vvPrefix, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvStaticDelegatedPrefixID, ok := result2["StaticDelegatedPrefixID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter StaticDelegatedPrefixID",
					"Fail Parsing StaticDelegatedPrefixID",
				)
				return
			}
			r.client.Appliance.UpdateNetworkAppliancePrefixesDelegatedStatic(vvNetworkID, vvStaticDelegatedPrefixID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Appliance.GetNetworkAppliancePrefixesDelegatedStatic(vvNetworkID, vvStaticDelegatedPrefixID)
			if responseVerifyItem2 != nil {
				data = ResponseApplianceGetNetworkAppliancePrefixesDelegatedStaticItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	restyResp2, err := r.client.Appliance.CreateNetworkAppliancePrefixesDelegatedStatic(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkAppliancePrefixesDelegatedStatic",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkAppliancePrefixesDelegatedStatic",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Appliance.GetNetworkAppliancePrefixesDelegatedStatics(vvNetworkID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkAppliancePrefixesDelegatedStatics",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkAppliancePrefixesDelegatedStatics",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Prefix", vvPrefix, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvStaticDelegatedPrefixID, ok := result2["Prefix"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter StaticDelegatedPrefixID",
				err.Error(),
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Appliance.GetNetworkAppliancePrefixesDelegatedStatic(vvNetworkID, vvStaticDelegatedPrefixID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseApplianceGetNetworkAppliancePrefixesDelegatedStaticItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkAppliancePrefixesDelegatedStatic",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkAppliancePrefixesDelegatedStatic",
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

func (r *NetworksAppliancePrefixesDelegatedStaticsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksAppliancePrefixesDelegatedStaticsRs

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

	vvNetworkID := data.NetworkID.ValueString()
	vvStaticDelegatedPrefixID := data.StaticDelegatedPrefixID.ValueString()
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkAppliancePrefixesDelegatedStatic(vvNetworkID, vvStaticDelegatedPrefixID)
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
				"Failure when executing GetNetworkAppliancePrefixesDelegatedStatic",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkAppliancePrefixesDelegatedStatic",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkAppliancePrefixesDelegatedStaticItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksAppliancePrefixesDelegatedStaticsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("static_delegated_prefix_id"), idParts[1])...)
}

func (r *NetworksAppliancePrefixesDelegatedStaticsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksAppliancePrefixesDelegatedStaticsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvStaticDelegatedPrefixID := data.StaticDelegatedPrefixID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkAppliancePrefixesDelegatedStatic(vvNetworkID, vvStaticDelegatedPrefixID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkAppliancePrefixesDelegatedStatic",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkAppliancePrefixesDelegatedStatic",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksAppliancePrefixesDelegatedStaticsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksAppliancePrefixesDelegatedStaticsRs
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

	vvNetworkID := state.NetworkID.ValueString()
	vvStaticDelegatedPrefixID := state.StaticDelegatedPrefixID.ValueString()
	_, err := r.client.Appliance.DeleteNetworkAppliancePrefixesDelegatedStatic(vvNetworkID, vvStaticDelegatedPrefixID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkAppliancePrefixesDelegatedStatic", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksAppliancePrefixesDelegatedStaticsRs struct {
	NetworkID               types.String                                                         `tfsdk:"network_id"`
	StaticDelegatedPrefixID types.String                                                         `tfsdk:"static_delegated_prefix_id"`
	CreatedAt               types.String                                                         `tfsdk:"created_at"`
	Description             types.String                                                         `tfsdk:"description"`
	Origin                  *ResponseApplianceGetNetworkAppliancePrefixesDelegatedStaticOriginRs `tfsdk:"origin"`
	Prefix                  types.String                                                         `tfsdk:"prefix"`
	UpdatedAt               types.String                                                         `tfsdk:"updated_at"`
}

type ResponseApplianceGetNetworkAppliancePrefixesDelegatedStaticOriginRs struct {
	Interfaces types.Set    `tfsdk:"interfaces"`
	Type       types.String `tfsdk:"type"`
}

// FromBody
func (r *NetworksAppliancePrefixesDelegatedStaticsRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestApplianceCreateNetworkAppliancePrefixesDelegatedStatic {
	emptyString := ""
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = &emptyString
	}
	var requestApplianceCreateNetworkAppliancePrefixesDelegatedStaticOrigin *merakigosdk.RequestApplianceCreateNetworkAppliancePrefixesDelegatedStaticOrigin
	if r.Origin != nil {
		var interfaces []string = nil
		//Hoola aqui
		r.Origin.Interfaces.ElementsAs(ctx, &interfaces, false)
		typeR := r.Origin.Type.ValueString()
		requestApplianceCreateNetworkAppliancePrefixesDelegatedStaticOrigin = &merakigosdk.RequestApplianceCreateNetworkAppliancePrefixesDelegatedStaticOrigin{
			Interfaces: interfaces,
			Type:       typeR,
		}
	}
	prefix := new(string)
	if !r.Prefix.IsUnknown() && !r.Prefix.IsNull() {
		*prefix = r.Prefix.ValueString()
	} else {
		prefix = &emptyString
	}
	out := merakigosdk.RequestApplianceCreateNetworkAppliancePrefixesDelegatedStatic{
		Description: *description,
		Origin:      requestApplianceCreateNetworkAppliancePrefixesDelegatedStaticOrigin,
		Prefix:      *prefix,
	}
	return &out
}
func (r *NetworksAppliancePrefixesDelegatedStaticsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkAppliancePrefixesDelegatedStatic {
	emptyString := ""
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = &emptyString
	}
	var requestApplianceUpdateNetworkAppliancePrefixesDelegatedStaticOrigin *merakigosdk.RequestApplianceUpdateNetworkAppliancePrefixesDelegatedStaticOrigin
	if r.Origin != nil {
		var interfaces []string = nil
		//Hoola aqui
		r.Origin.Interfaces.ElementsAs(ctx, &interfaces, false)
		typeR := r.Origin.Type.ValueString()
		requestApplianceUpdateNetworkAppliancePrefixesDelegatedStaticOrigin = &merakigosdk.RequestApplianceUpdateNetworkAppliancePrefixesDelegatedStaticOrigin{
			Interfaces: interfaces,
			Type:       typeR,
		}
	}
	prefix := new(string)
	if !r.Prefix.IsUnknown() && !r.Prefix.IsNull() {
		*prefix = r.Prefix.ValueString()
	} else {
		prefix = &emptyString
	}
	out := merakigosdk.RequestApplianceUpdateNetworkAppliancePrefixesDelegatedStatic{
		Description: *description,
		Origin:      requestApplianceUpdateNetworkAppliancePrefixesDelegatedStaticOrigin,
		Prefix:      *prefix,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkAppliancePrefixesDelegatedStaticItemToBodyRs(state NetworksAppliancePrefixesDelegatedStaticsRs, response *merakigosdk.ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatic, is_read bool) NetworksAppliancePrefixesDelegatedStaticsRs {
	itemState := NetworksAppliancePrefixesDelegatedStaticsRs{
		CreatedAt:   types.StringValue(response.CreatedAt),
		Description: types.StringValue(response.Description),
		Origin: func() *ResponseApplianceGetNetworkAppliancePrefixesDelegatedStaticOriginRs {
			if response.Origin != nil {
				return &ResponseApplianceGetNetworkAppliancePrefixesDelegatedStaticOriginRs{
					Interfaces: StringSliceToSet(response.Origin.Interfaces),
					Type:       types.StringValue(response.Origin.Type),
				}
			}
			return &ResponseApplianceGetNetworkAppliancePrefixesDelegatedStaticOriginRs{}
		}(),
		Prefix:                  types.StringValue(response.Prefix),
		StaticDelegatedPrefixID: types.StringValue(response.StaticDelegatedPrefixID),
		UpdatedAt:               types.StringValue(response.UpdatedAt),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksAppliancePrefixesDelegatedStaticsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksAppliancePrefixesDelegatedStaticsRs)
}
