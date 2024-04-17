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
	_ datasource.DataSource              = &OrganizationsInventoryOnboardingCloudMonitoringNetworksDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsInventoryOnboardingCloudMonitoringNetworksDataSource{}
)

func NewOrganizationsInventoryOnboardingCloudMonitoringNetworksDataSource() datasource.DataSource {
	return &OrganizationsInventoryOnboardingCloudMonitoringNetworksDataSource{}
}

type OrganizationsInventoryOnboardingCloudMonitoringNetworksDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsInventoryOnboardingCloudMonitoringNetworksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsInventoryOnboardingCloudMonitoringNetworksDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_inventory_onboarding_cloud_monitoring_networks"
}

func (d *OrganizationsInventoryOnboardingCloudMonitoringNetworksDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"device_type": schema.StringAttribute{
				MarkdownDescription: `deviceType query parameter. Device Type switch or wireless controller`,
				Required:            true,
			},
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 100000. Default is 1000.`,
				Optional:            true,
			},
			"search": schema.StringAttribute{
				MarkdownDescription: `search query parameter. Optional parameter to search on network name`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringNetworks`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"enrollment_string": schema.StringAttribute{
							MarkdownDescription: `Enrollment string for the network`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `Network ID`,
							Computed:            true,
						},
						"is_bound_to_config_template": schema.BoolAttribute{
							MarkdownDescription: `If the network is bound to a config template`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Network name`,
							Computed:            true,
						},
						"notes": schema.StringAttribute{
							MarkdownDescription: `Notes for the network`,
							Computed:            true,
						},
						"organization_id": schema.StringAttribute{
							MarkdownDescription: `Organization ID`,
							Computed:            true,
						},
						"product_types": schema.ListAttribute{
							MarkdownDescription: `List of the product types that the network supports`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"tags": schema.ListAttribute{
							MarkdownDescription: `Network tags`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"time_zone": schema.StringAttribute{
							MarkdownDescription: `Timezone of the network`,
							Computed:            true,
						},
						"url": schema.StringAttribute{
							MarkdownDescription: `URL to the network Dashboard UI`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsInventoryOnboardingCloudMonitoringNetworksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsInventoryOnboardingCloudMonitoringNetworks OrganizationsInventoryOnboardingCloudMonitoringNetworks
	diags := req.Config.Get(ctx, &organizationsInventoryOnboardingCloudMonitoringNetworks)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationInventoryOnboardingCloudMonitoringNetworks")
		vvOrganizationID := organizationsInventoryOnboardingCloudMonitoringNetworks.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationInventoryOnboardingCloudMonitoringNetworksQueryParams{}

		queryParams1.DeviceType = organizationsInventoryOnboardingCloudMonitoringNetworks.DeviceType.ValueString()

		queryParams1.Search = organizationsInventoryOnboardingCloudMonitoringNetworks.Search.ValueString()
		queryParams1.PerPage = int(organizationsInventoryOnboardingCloudMonitoringNetworks.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsInventoryOnboardingCloudMonitoringNetworks.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsInventoryOnboardingCloudMonitoringNetworks.EndingBefore.ValueString()

		response1, restyResp1, err := d.client.Organizations.GetOrganizationInventoryOnboardingCloudMonitoringNetworks(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationInventoryOnboardingCloudMonitoringNetworks",
				err.Error(),
			)
			return
		}

		organizationsInventoryOnboardingCloudMonitoringNetworks = ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringNetworksItemsToBody(organizationsInventoryOnboardingCloudMonitoringNetworks, response1)
		diags = resp.State.Set(ctx, &organizationsInventoryOnboardingCloudMonitoringNetworks)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsInventoryOnboardingCloudMonitoringNetworks struct {
	OrganizationID types.String                                                                          `tfsdk:"organization_id"`
	DeviceType     types.String                                                                          `tfsdk:"device_type"`
	Search         types.String                                                                          `tfsdk:"search"`
	PerPage        types.Int64                                                                           `tfsdk:"per_page"`
	StartingAfter  types.String                                                                          `tfsdk:"starting_after"`
	EndingBefore   types.String                                                                          `tfsdk:"ending_before"`
	Items          *[]ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringNetworks `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringNetworks struct {
	EnrollmentString        types.String `tfsdk:"enrollment_string"`
	ID                      types.String `tfsdk:"id"`
	IsBoundToConfigTemplate types.Bool   `tfsdk:"is_bound_to_config_template"`
	Name                    types.String `tfsdk:"name"`
	Notes                   types.String `tfsdk:"notes"`
	OrganizationID          types.String `tfsdk:"organization_id"`
	ProductTypes            types.List   `tfsdk:"product_types"`
	Tags                    types.List   `tfsdk:"tags"`
	TimeZone                types.String `tfsdk:"time_zone"`
	URL                     types.String `tfsdk:"url"`
}

// ToBody
func ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringNetworksItemsToBody(state OrganizationsInventoryOnboardingCloudMonitoringNetworks, response *merakigosdk.ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringNetworks) OrganizationsInventoryOnboardingCloudMonitoringNetworks {
	var items []ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringNetworks
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringNetworks{
			EnrollmentString: types.StringValue(item.EnrollmentString),
			ID:               types.StringValue(item.ID),
			IsBoundToConfigTemplate: func() types.Bool {
				if item.IsBoundToConfigTemplate != nil {
					return types.BoolValue(*item.IsBoundToConfigTemplate)
				}
				return types.Bool{}
			}(),
			Name:           types.StringValue(item.Name),
			Notes:          types.StringValue(item.Notes),
			OrganizationID: types.StringValue(item.OrganizationID),
			ProductTypes:   StringSliceToList(item.ProductTypes),
			Tags:           StringSliceToList(item.Tags),
			TimeZone:       types.StringValue(item.TimeZone),
			URL:            types.StringValue(item.URL),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
