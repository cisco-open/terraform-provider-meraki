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
	_ datasource.DataSource              = &OrganizationsInsightApplicationsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsInsightApplicationsDataSource{}
)

func NewOrganizationsInsightApplicationsDataSource() datasource.DataSource {
	return &OrganizationsInsightApplicationsDataSource{}
}

type OrganizationsInsightApplicationsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsInsightApplicationsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsInsightApplicationsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_insight_applications"
}

func (d *OrganizationsInsightApplicationsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseInsightGetOrganizationInsightApplications`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"application_id": schema.StringAttribute{
							MarkdownDescription: `Application identifier`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Application name`,
							Computed:            true,
						},
						"thresholds": schema.SingleNestedAttribute{
							MarkdownDescription: `Thresholds defined by a user or Meraki models for each application`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"by_network": schema.SetNestedAttribute{
									MarkdownDescription: `Threshold for each network`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"goodput": schema.Int64Attribute{
												MarkdownDescription: `Number of useful information bits delivered over a network per unit of time`,
												Computed:            true,
											},
											"network_id": schema.StringAttribute{
												MarkdownDescription: `Network identifier`,
												Computed:            true,
											},
											"response_duration": schema.Int64Attribute{
												MarkdownDescription: `Duration of the response, in milliseconds`,
												Computed:            true,
											},
										},
									},
								},
								"type": schema.StringAttribute{
									MarkdownDescription: `Threshold type (static or smart)`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsInsightApplicationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsInsightApplications OrganizationsInsightApplications
	diags := req.Config.Get(ctx, &organizationsInsightApplications)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationInsightApplications")
		vvOrganizationID := organizationsInsightApplications.OrganizationID.ValueString()

		response1, restyResp1, err := d.client.Insight.GetOrganizationInsightApplications(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationInsightApplications",
				err.Error(),
			)
			return
		}

		organizationsInsightApplications = ResponseInsightGetOrganizationInsightApplicationsItemsToBody(organizationsInsightApplications, response1)
		diags = resp.State.Set(ctx, &organizationsInsightApplications)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsInsightApplications struct {
	OrganizationID types.String                                             `tfsdk:"organization_id"`
	Items          *[]ResponseItemInsightGetOrganizationInsightApplications `tfsdk:"items"`
}

type ResponseItemInsightGetOrganizationInsightApplications struct {
	ApplicationID types.String                                                     `tfsdk:"application_id"`
	Name          types.String                                                     `tfsdk:"name"`
	Thresholds    *ResponseItemInsightGetOrganizationInsightApplicationsThresholds `tfsdk:"thresholds"`
}

type ResponseItemInsightGetOrganizationInsightApplicationsThresholds struct {
	ByNetwork *[]ResponseItemInsightGetOrganizationInsightApplicationsThresholdsByNetwork `tfsdk:"by_network"`
	Type      types.String                                                                `tfsdk:"type"`
}

type ResponseItemInsightGetOrganizationInsightApplicationsThresholdsByNetwork struct {
	Goodput          types.Int64  `tfsdk:"goodput"`
	NetworkID        types.String `tfsdk:"network_id"`
	ResponseDuration types.Int64  `tfsdk:"response_duration"`
}

// ToBody
func ResponseInsightGetOrganizationInsightApplicationsItemsToBody(state OrganizationsInsightApplications, response *merakigosdk.ResponseInsightGetOrganizationInsightApplications) OrganizationsInsightApplications {
	var items []ResponseItemInsightGetOrganizationInsightApplications
	for _, item := range *response {
		itemState := ResponseItemInsightGetOrganizationInsightApplications{
			ApplicationID: types.StringValue(item.ApplicationID),
			Name:          types.StringValue(item.Name),
			Thresholds: func() *ResponseItemInsightGetOrganizationInsightApplicationsThresholds {
				if item.Thresholds != nil {
					return &ResponseItemInsightGetOrganizationInsightApplicationsThresholds{
						ByNetwork: func() *[]ResponseItemInsightGetOrganizationInsightApplicationsThresholdsByNetwork {
							if item.Thresholds.ByNetwork != nil {
								result := make([]ResponseItemInsightGetOrganizationInsightApplicationsThresholdsByNetwork, len(*item.Thresholds.ByNetwork))
								for i, byNetwork := range *item.Thresholds.ByNetwork {
									result[i] = ResponseItemInsightGetOrganizationInsightApplicationsThresholdsByNetwork{
										Goodput: func() types.Int64 {
											if byNetwork.Goodput != nil {
												return types.Int64Value(int64(*byNetwork.Goodput))
											}
											return types.Int64{}
										}(),
										NetworkID: types.StringValue(byNetwork.NetworkID),
										ResponseDuration: func() types.Int64 {
											if byNetwork.ResponseDuration != nil {
												return types.Int64Value(int64(*byNetwork.ResponseDuration))
											}
											return types.Int64{}
										}(),
									}
								}
								return &result
							}
							return &[]ResponseItemInsightGetOrganizationInsightApplicationsThresholdsByNetwork{}
						}(),
						Type: types.StringValue(item.Thresholds.Type),
					}
				}
				return &ResponseItemInsightGetOrganizationInsightApplicationsThresholds{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
