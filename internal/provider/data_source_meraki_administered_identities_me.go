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
	_ datasource.DataSource              = &AdministeredIDentitiesMeDataSource{}
	_ datasource.DataSourceWithConfigure = &AdministeredIDentitiesMeDataSource{}
)

func NewAdministeredIDentitiesMeDataSource() datasource.DataSource {
	return &AdministeredIDentitiesMeDataSource{}
}

type AdministeredIDentitiesMeDataSource struct {
	client *merakigosdk.Client
}

func (d *AdministeredIDentitiesMeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *AdministeredIDentitiesMeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_administered_identities_me"
}

func (d *AdministeredIDentitiesMeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"authentication": schema.SingleNestedAttribute{
						MarkdownDescription: `Authentication info`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"api": schema.SingleNestedAttribute{
								MarkdownDescription: `API authentication`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"key": schema.SingleNestedAttribute{
										MarkdownDescription: `API key`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"created": schema.BoolAttribute{
												MarkdownDescription: `If API key is created for this user`,
												Computed:            true,
											},
										},
									},
								},
							},
							"mode": schema.StringAttribute{
								MarkdownDescription: `Authentication mode`,
								Computed:            true,
							},
							"saml": schema.SingleNestedAttribute{
								MarkdownDescription: `SAML authentication`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"enabled": schema.BoolAttribute{
										MarkdownDescription: `If SAML authentication is enabled for this user`,
										Computed:            true,
									},
								},
							},
							"two_factor": schema.SingleNestedAttribute{
								MarkdownDescription: `TwoFactor authentication`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"enabled": schema.BoolAttribute{
										MarkdownDescription: `If twoFactor authentication is enabled for this user`,
										Computed:            true,
									},
								},
							},
						},
					},
					"email": schema.StringAttribute{
						MarkdownDescription: `User email`,
						Computed:            true,
					},
					"last_used_dashboard_at": schema.StringAttribute{
						MarkdownDescription: `Last seen active on Dashboard UI`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Username`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *AdministeredIDentitiesMeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var administeredIDentitiesMe AdministeredIDentitiesMe
	diags := req.Config.Get(ctx, &administeredIDentitiesMe)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAdministeredIDentitiesMe")

		response1, restyResp1, err := d.client.Administered.GetAdministeredIDentitiesMe()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetAdministeredIDentitiesMe",
				err.Error(),
			)
			return
		}

		administeredIDentitiesMe = ResponseAdministeredGetAdministeredIDentitiesMeItemToBody(administeredIDentitiesMe, response1)
		diags = resp.State.Set(ctx, &administeredIDentitiesMe)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type AdministeredIDentitiesMe struct {
	Item *ResponseAdministeredGetAdministeredIdentitiesMe `tfsdk:"item"`
}

type ResponseAdministeredGetAdministeredIdentitiesMe struct {
	Authentication      *ResponseAdministeredGetAdministeredIdentitiesMeAuthentication `tfsdk:"authentication"`
	Email               types.String                                                   `tfsdk:"email"`
	LastUsedDashboardAt types.String                                                   `tfsdk:"last_used_dashboard_at"`
	Name                types.String                                                   `tfsdk:"name"`
}

type ResponseAdministeredGetAdministeredIdentitiesMeAuthentication struct {
	API       *ResponseAdministeredGetAdministeredIdentitiesMeAuthenticationApi       `tfsdk:"api"`
	Mode      types.String                                                            `tfsdk:"mode"`
	Saml      *ResponseAdministeredGetAdministeredIdentitiesMeAuthenticationSaml      `tfsdk:"saml"`
	TwoFactor *ResponseAdministeredGetAdministeredIdentitiesMeAuthenticationTwoFactor `tfsdk:"two_factor"`
}

type ResponseAdministeredGetAdministeredIdentitiesMeAuthenticationApi struct {
	Key *ResponseAdministeredGetAdministeredIdentitiesMeAuthenticationApiKey `tfsdk:"key"`
}

type ResponseAdministeredGetAdministeredIdentitiesMeAuthenticationApiKey struct {
	Created types.Bool `tfsdk:"created"`
}

type ResponseAdministeredGetAdministeredIdentitiesMeAuthenticationSaml struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseAdministeredGetAdministeredIdentitiesMeAuthenticationTwoFactor struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// ToBody
func ResponseAdministeredGetAdministeredIDentitiesMeItemToBody(state AdministeredIDentitiesMe, response *merakigosdk.ResponseAdministeredGetAdministeredIDentitiesMe) AdministeredIDentitiesMe {
	itemState := ResponseAdministeredGetAdministeredIdentitiesMe{
		Authentication: func() *ResponseAdministeredGetAdministeredIdentitiesMeAuthentication {
			if response.Authentication != nil {
				return &ResponseAdministeredGetAdministeredIdentitiesMeAuthentication{
					API: func() *ResponseAdministeredGetAdministeredIdentitiesMeAuthenticationApi {
						if response.Authentication.API != nil {
							return &ResponseAdministeredGetAdministeredIdentitiesMeAuthenticationApi{
								Key: func() *ResponseAdministeredGetAdministeredIdentitiesMeAuthenticationApiKey {
									if response.Authentication.API.Key != nil {
										return &ResponseAdministeredGetAdministeredIdentitiesMeAuthenticationApiKey{
											Created: func() types.Bool {
												if response.Authentication.API.Key.Created != nil {
													return types.BoolValue(*response.Authentication.API.Key.Created)
												}
												return types.Bool{}
											}(),
										}
									}
									return &ResponseAdministeredGetAdministeredIdentitiesMeAuthenticationApiKey{}
								}(),
							}
						}
						return &ResponseAdministeredGetAdministeredIdentitiesMeAuthenticationApi{}
					}(),
					Mode: types.StringValue(response.Authentication.Mode),
					Saml: func() *ResponseAdministeredGetAdministeredIdentitiesMeAuthenticationSaml {
						if response.Authentication.Saml != nil {
							return &ResponseAdministeredGetAdministeredIdentitiesMeAuthenticationSaml{
								Enabled: func() types.Bool {
									if response.Authentication.Saml.Enabled != nil {
										return types.BoolValue(*response.Authentication.Saml.Enabled)
									}
									return types.Bool{}
								}(),
							}
						}
						return &ResponseAdministeredGetAdministeredIdentitiesMeAuthenticationSaml{}
					}(),
					TwoFactor: func() *ResponseAdministeredGetAdministeredIdentitiesMeAuthenticationTwoFactor {
						if response.Authentication.TwoFactor != nil {
							return &ResponseAdministeredGetAdministeredIdentitiesMeAuthenticationTwoFactor{
								Enabled: func() types.Bool {
									if response.Authentication.TwoFactor.Enabled != nil {
										return types.BoolValue(*response.Authentication.TwoFactor.Enabled)
									}
									return types.Bool{}
								}(),
							}
						}
						return &ResponseAdministeredGetAdministeredIdentitiesMeAuthenticationTwoFactor{}
					}(),
				}
			}
			return &ResponseAdministeredGetAdministeredIdentitiesMeAuthentication{}
		}(),
		Email:               types.StringValue(response.Email),
		LastUsedDashboardAt: types.StringValue(response.LastUsedDashboardAt),
		Name:                types.StringValue(response.Name),
	}
	state.Item = &itemState
	return state
}
