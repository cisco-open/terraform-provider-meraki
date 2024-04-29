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
	_ datasource.DataSource              = &NetworksSmUsersDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSmUsersDataSource{}
)

func NewNetworksSmUsersDataSource() datasource.DataSource {
	return &NetworksSmUsersDataSource{}
}

type NetworksSmUsersDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSmUsersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSmUsersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_users"
}

func (d *NetworksSmUsersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"emails": schema.ListAttribute{
				MarkdownDescription: `emails query parameter. Filter users by email(s).`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"ids": schema.ListAttribute{
				MarkdownDescription: `ids query parameter. Filter users by id(s).`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"scope": schema.ListAttribute{
				MarkdownDescription: `scope query parameter. Specifiy a scope (one of all, none, withAny, withAll, withoutAny, withoutAll) and a set of tags.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"usernames": schema.ListAttribute{
				MarkdownDescription: `usernames query parameter. Filter users by username(s).`,
				Optional:            true,
				ElementType:         types.StringType,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSmGetNetworkSmUsers`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"ad_groups": schema.ListAttribute{
							MarkdownDescription: `Active Directory Groups the user belongs to.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"asm_groups": schema.ListAttribute{
							MarkdownDescription: `Apple School Manager Groups the user belongs to.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"azure_ad_groups": schema.ListAttribute{
							MarkdownDescription: `Azure Active Directory Groups the user belongs to.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"display_name": schema.StringAttribute{
							MarkdownDescription: `The user display name.`,
							Computed:            true,
						},
						"email": schema.StringAttribute{
							MarkdownDescription: `User email.`,
							Computed:            true,
						},
						"full_name": schema.StringAttribute{
							MarkdownDescription: `User full name.`,
							Computed:            true,
						},
						"has_identity_certificate": schema.BoolAttribute{
							MarkdownDescription: `A boolean indicating if the user has an associated identity certificate..`,
							Computed:            true,
						},
						"has_password": schema.BoolAttribute{
							MarkdownDescription: `A boolean denoting if the user has a password associated with the record.`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `The Meraki managed Id of the user record.`,
							Computed:            true,
						},
						"is_external": schema.BoolAttribute{
							MarkdownDescription: `Whether the user was created using an external integration, or via the Meraki Dashboard.`,
							Computed:            true,
						},
						"saml_groups": schema.ListAttribute{
							MarkdownDescription: `SAML Groups the user belongs to.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"tags": schema.StringAttribute{
							MarkdownDescription: `The set of tags the user is scoped to.`,
							Computed:            true,
						},
						"user_thumbnail": schema.StringAttribute{
							MarkdownDescription: `The url where the users thumbnail is hosted.`,
							Computed:            true,
						},
						"username": schema.StringAttribute{
							MarkdownDescription: `The users username.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSmUsersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSmUsers NetworksSmUsers
	diags := req.Config.Get(ctx, &networksSmUsers)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSmUsers")
		vvNetworkID := networksSmUsers.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkSmUsersQueryParams{}

		queryParams1.IDs = elementsToStrings(ctx, networksSmUsers.IDs)
		queryParams1.Usernames = elementsToStrings(ctx, networksSmUsers.Usernames)
		queryParams1.Emails = elementsToStrings(ctx, networksSmUsers.Emails)
		queryParams1.Scope = elementsToStrings(ctx, networksSmUsers.Scope)

		response1, restyResp1, err := d.client.Sm.GetNetworkSmUsers(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmUsers",
				err.Error(),
			)
			return
		}

		networksSmUsers = ResponseSmGetNetworkSmUsersItemsToBody(networksSmUsers, response1)
		diags = resp.State.Set(ctx, &networksSmUsers)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSmUsers struct {
	NetworkID types.String                       `tfsdk:"network_id"`
	IDs       types.List                         `tfsdk:"ids"`
	Usernames types.List                         `tfsdk:"usernames"`
	Emails    types.List                         `tfsdk:"emails"`
	Scope     types.List                         `tfsdk:"scope"`
	Items     *[]ResponseItemSmGetNetworkSmUsers `tfsdk:"items"`
}

type ResponseItemSmGetNetworkSmUsers struct {
	AdGroups               types.List   `tfsdk:"ad_groups"`
	AsmGroups              types.List   `tfsdk:"asm_groups"`
	AzureAdGroups          types.List   `tfsdk:"azure_ad_groups"`
	DisplayName            types.String `tfsdk:"display_name"`
	Email                  types.String `tfsdk:"email"`
	FullName               types.String `tfsdk:"full_name"`
	HasIDentityCertificate types.Bool   `tfsdk:"has_identity_certificate"`
	HasPassword            types.Bool   `tfsdk:"has_password"`
	ID                     types.String `tfsdk:"id"`
	IsExternal             types.Bool   `tfsdk:"is_external"`
	SamlGroups             types.List   `tfsdk:"saml_groups"`
	Tags                   types.String `tfsdk:"tags"`
	UserThumbnail          types.String `tfsdk:"user_thumbnail"`
	Username               types.String `tfsdk:"username"`
}

// ToBody
func ResponseSmGetNetworkSmUsersItemsToBody(state NetworksSmUsers, response *merakigosdk.ResponseSmGetNetworkSmUsers) NetworksSmUsers {
	var items []ResponseItemSmGetNetworkSmUsers
	for _, item := range *response {
		itemState := ResponseItemSmGetNetworkSmUsers{
			AdGroups:      StringSliceToList(item.AdGroups),
			AsmGroups:     StringSliceToList(item.AsmGroups),
			AzureAdGroups: StringSliceToList(item.AzureAdGroups),
			DisplayName:   types.StringValue(item.DisplayName),
			Email:         types.StringValue(item.Email),
			FullName:      types.StringValue(item.FullName),
			HasIDentityCertificate: func() types.Bool {
				if item.HasIDentityCertificate != nil {
					return types.BoolValue(*item.HasIDentityCertificate)
				}
				return types.Bool{}
			}(),
			HasPassword: func() types.Bool {
				if item.HasPassword != nil {
					return types.BoolValue(*item.HasPassword)
				}
				return types.Bool{}
			}(),
			ID: types.StringValue(item.ID),
			IsExternal: func() types.Bool {
				if item.IsExternal != nil {
					return types.BoolValue(*item.IsExternal)
				}
				return types.Bool{}
			}(),
			SamlGroups:    StringSliceToList(item.SamlGroups),
			Tags:          types.StringValue(item.Tags),
			UserThumbnail: types.StringValue(item.UserThumbnail),
			Username:      types.StringValue(item.Username),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
