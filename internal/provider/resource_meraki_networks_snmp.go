package provider

// RESOURCE NORMAL
import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSNMPResource{}
	_ resource.ResourceWithConfigure = &NetworksSNMPResource{}
)

func NewNetworksSNMPResource() resource.Resource {
	return &NetworksSNMPResource{}
}

type NetworksSNMPResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSNMPResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSNMPResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_snmp"
}

func (r *NetworksSNMPResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"access": schema.StringAttribute{
				MarkdownDescription: `The type of SNMP access. Can be one of 'none' (disabled), 'community' (V1/V2c), or 'users' (V3).`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"community",
						"none",
						"users",
					),
				},
			},
			"community_string": schema.StringAttribute{
				MarkdownDescription: `SNMP community string if access is 'community'.`,
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
			"users": schema.SetNestedAttribute{
				MarkdownDescription: `SNMP settings if access is 'users'.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"passphrase": schema.StringAttribute{
							MarkdownDescription: `The passphrase for the SNMP user.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"username": schema.StringAttribute{
							MarkdownDescription: `The username for the SNMP user.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
		},
	}
}

func (r *NetworksSNMPResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSNMPRs

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
	//Item
	responseVerifyItem, restyResp1, err := r.client.Networks.GetNetworkSNMP(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksSNMP only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksSNMP only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkSNMP(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSNMP",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSNMP",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Networks.GetNetworkSNMP(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSNMP",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSNMP",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseNetworksGetNetworkSNMPItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSNMPResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSNMPRs

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
	responseGet, restyRespGet, err := r.client.Networks.GetNetworkSNMP(vvNetworkID)
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
				"Failure when executing GetNetworkSNMP",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSNMP",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseNetworksGetNetworkSNMPItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksSNMPResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksSNMPResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSNMPRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkSNMP(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSNMP",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSNMP",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSNMPResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksSNMP", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSNMPRs struct {
	NetworkID       types.String                             `tfsdk:"network_id"`
	Access          types.String                             `tfsdk:"access"`
	CommunityString types.String                             `tfsdk:"community_string"`
	Users           *[]ResponseNetworksGetNetworkSnmpUsersRs `tfsdk:"users"`
}

type ResponseNetworksGetNetworkSnmpUsersRs struct {
	Passphrase types.String `tfsdk:"passphrase"`
	Username   types.String `tfsdk:"username"`
}

// FromBody
func (r *NetworksSNMPRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestNetworksUpdateNetworkSNMP {
	emptyString := ""
	access := new(string)
	if !r.Access.IsUnknown() && !r.Access.IsNull() {
		*access = r.Access.ValueString()
	} else {
		access = &emptyString
	}
	communityString := new(string)
	if !r.CommunityString.IsUnknown() && !r.CommunityString.IsNull() {
		*communityString = r.CommunityString.ValueString()
	} else {
		communityString = &emptyString
	}
	var requestNetworksUpdateNetworkSNMPUsers []merakigosdk.RequestNetworksUpdateNetworkSNMPUsers
	if r.Users != nil {
		for _, rItem1 := range *r.Users {
			passphrase := rItem1.Passphrase.ValueString()
			username := rItem1.Username.ValueString()
			requestNetworksUpdateNetworkSNMPUsers = append(requestNetworksUpdateNetworkSNMPUsers, merakigosdk.RequestNetworksUpdateNetworkSNMPUsers{
				Passphrase: passphrase,
				Username:   username,
			})
		}
	}
	out := merakigosdk.RequestNetworksUpdateNetworkSNMP{
		Access:          *access,
		CommunityString: *communityString,
		Users: func() *[]merakigosdk.RequestNetworksUpdateNetworkSNMPUsers {
			if len(requestNetworksUpdateNetworkSNMPUsers) > 0 {
				return &requestNetworksUpdateNetworkSNMPUsers
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseNetworksGetNetworkSNMPItemToBodyRs(state NetworksSNMPRs, response *merakigosdk.ResponseNetworksGetNetworkSNMP, is_read bool) NetworksSNMPRs {
	itemState := NetworksSNMPRs{
		Access:          types.StringValue(response.Access),
		CommunityString: types.StringValue(response.CommunityString),
		Users: func() *[]ResponseNetworksGetNetworkSnmpUsersRs {
			if response.Users != nil {
				result := make([]ResponseNetworksGetNetworkSnmpUsersRs, len(*response.Users))
				for i, users := range *response.Users {
					result[i] = ResponseNetworksGetNetworkSnmpUsersRs{
						Passphrase: types.StringValue(users.Passphrase),
						Username:   types.StringValue(users.Username),
					}
				}
				return &result
			}
			return &[]ResponseNetworksGetNetworkSnmpUsersRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSNMPRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSNMPRs)
}
