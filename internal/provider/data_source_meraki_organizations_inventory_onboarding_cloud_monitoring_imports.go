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
	_ datasource.DataSource              = &OrganizationsInventoryOnboardingCloudMonitoringImportsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsInventoryOnboardingCloudMonitoringImportsDataSource{}
)

func NewOrganizationsInventoryOnboardingCloudMonitoringImportsDataSource() datasource.DataSource {
	return &OrganizationsInventoryOnboardingCloudMonitoringImportsDataSource{}
}

type OrganizationsInventoryOnboardingCloudMonitoringImportsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsInventoryOnboardingCloudMonitoringImportsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsInventoryOnboardingCloudMonitoringImportsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_inventory_onboarding_cloud_monitoring_imports"
}

func (d *OrganizationsInventoryOnboardingCloudMonitoringImportsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"import_ids": schema.ListAttribute{
				MarkdownDescription: `importIds query parameter. import ids from an imports`,
				Required:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImports`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"device": schema.SingleNestedAttribute{
							MarkdownDescription: `Represents the details of an imported device.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"created": schema.BoolAttribute{
									MarkdownDescription: `Whether or not the device was successfully created in dashboard.`,
									Computed:            true,
								},
								"status": schema.StringAttribute{
									MarkdownDescription: `Represents the current state of importing the device.`,
									Computed:            true,
								},
								"url": schema.StringAttribute{
									MarkdownDescription: `The url to the device details page within dashboard.`,
									Computed:            true,
								},
							},
						},
						"import_id": schema.StringAttribute{
							MarkdownDescription: `Database ID for the new entity entry.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsInventoryOnboardingCloudMonitoringImportsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsInventoryOnboardingCloudMonitoringImports OrganizationsInventoryOnboardingCloudMonitoringImports
	diags := req.Config.Get(ctx, &organizationsInventoryOnboardingCloudMonitoringImports)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationInventoryOnboardingCloudMonitoringImports")
		vvOrganizationID := organizationsInventoryOnboardingCloudMonitoringImports.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationInventoryOnboardingCloudMonitoringImportsQueryParams{}

		queryParams1.ImportIDs = elementsToStrings(ctx, organizationsInventoryOnboardingCloudMonitoringImports.ImportIDs)

		response1, restyResp1, err := d.client.Organizations.GetOrganizationInventoryOnboardingCloudMonitoringImports(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationInventoryOnboardingCloudMonitoringImports",
				err.Error(),
			)
			return
		}

		organizationsInventoryOnboardingCloudMonitoringImports = ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImportsItemsToBody(organizationsInventoryOnboardingCloudMonitoringImports, response1)
		diags = resp.State.Set(ctx, &organizationsInventoryOnboardingCloudMonitoringImports)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsInventoryOnboardingCloudMonitoringImports struct {
	OrganizationID types.String                                                                         `tfsdk:"organization_id"`
	ImportIDs      types.List                                                                           `tfsdk:"import_ids"`
	Items          *[]ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImports `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImports struct {
	Device   *ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImportsDevice `tfsdk:"device"`
	ImportID types.String                                                                             `tfsdk:"import_id"`
}

type ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImportsDevice struct {
	Created types.Bool   `tfsdk:"created"`
	Status  types.String `tfsdk:"status"`
	URL     types.String `tfsdk:"url"`
}

// ToBody
func ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImportsItemsToBody(state OrganizationsInventoryOnboardingCloudMonitoringImports, response *merakigosdk.ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImports) OrganizationsInventoryOnboardingCloudMonitoringImports {
	var items []ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImports
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImports{
			Device: func() *ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImportsDevice {
				if item.Device != nil {
					return &ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImportsDevice{
						Created: func() types.Bool {
							if item.Device.Created != nil {
								return types.BoolValue(*item.Device.Created)
							}
							return types.Bool{}
						}(),
						Status: types.StringValue(item.Device.Status),
						URL:    types.StringValue(item.Device.URL),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImportsDevice{}
			}(),
			ImportID: types.StringValue(item.ImportID),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
