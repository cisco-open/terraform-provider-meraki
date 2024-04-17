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
	_ datasource.DataSource              = &OrganizationsAPIRequestsOverviewResponseCodesByIntervalDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsAPIRequestsOverviewResponseCodesByIntervalDataSource{}
)

func NewOrganizationsAPIRequestsOverviewResponseCodesByIntervalDataSource() datasource.DataSource {
	return &OrganizationsAPIRequestsOverviewResponseCodesByIntervalDataSource{}
}

type OrganizationsAPIRequestsOverviewResponseCodesByIntervalDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsAPIRequestsOverviewResponseCodesByIntervalDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsAPIRequestsOverviewResponseCodesByIntervalDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_api_requests_overview_response_codes_by_interval"
}

func (d *OrganizationsAPIRequestsOverviewResponseCodesByIntervalDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"admin_ids": schema.ListAttribute{
				MarkdownDescription: `adminIds query parameter. Filter by admin ID of user that made the API request`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"interval": schema.Int64Attribute{
				MarkdownDescription: `interval query parameter. The time interval in seconds for returned data. The valid intervals are: 120, 3600, 14400, 21600. The default is 21600. Interval is calculated if time params are provided.`,
				Optional:            true,
			},
			"operation_ids": schema.ListAttribute{
				MarkdownDescription: `operationIds query parameter. Filter by operation ID of the endpoint`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"source_ips": schema.ListAttribute{
				MarkdownDescription: `sourceIps query parameter. Filter by source IP that made the API request`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 31 days from today.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 31 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. If interval is provided, the timespan will be autocalculated.`,
				Optional:            true,
			},
			"user_agent": schema.StringAttribute{
				MarkdownDescription: `userAgent query parameter. Filter by user agent string for API request. This will filter by a complete or partial match.`,
				Optional:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: `version query parameter. Filter by API version of the endpoint. Allowable values are: [0, 1]`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationApiRequestsOverviewResponseCodesByInterval`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"counts": schema.SetNestedAttribute{
							MarkdownDescription: `list of response codes and a count of how many requests had that code in the given time period`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"code": schema.Int64Attribute{
										MarkdownDescription: `Response status code of the API response`,
										Computed:            true,
									},
									"count": schema.Int64Attribute{
										MarkdownDescription: `Number of records that match the status code`,
										Computed:            true,
									},
								},
							},
						},
						"end_ts": schema.StringAttribute{
							MarkdownDescription: `The end time of the access period`,
							Computed:            true,
						},
						"start_ts": schema.StringAttribute{
							MarkdownDescription: `The start time of the access period`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsAPIRequestsOverviewResponseCodesByIntervalDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsAPIRequestsOverviewResponseCodesByInterval OrganizationsAPIRequestsOverviewResponseCodesByInterval
	diags := req.Config.Get(ctx, &organizationsAPIRequestsOverviewResponseCodesByInterval)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationAPIRequestsOverviewResponseCodesByInterval")
		vvOrganizationID := organizationsAPIRequestsOverviewResponseCodesByInterval.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationAPIRequestsOverviewResponseCodesByIntervalQueryParams{}

		queryParams1.T0 = organizationsAPIRequestsOverviewResponseCodesByInterval.T0.ValueString()
		queryParams1.T1 = organizationsAPIRequestsOverviewResponseCodesByInterval.T1.ValueString()
		queryParams1.Timespan = organizationsAPIRequestsOverviewResponseCodesByInterval.Timespan.ValueFloat64()
		queryParams1.Interval = int(organizationsAPIRequestsOverviewResponseCodesByInterval.Interval.ValueInt64())
		queryParams1.Version = int(organizationsAPIRequestsOverviewResponseCodesByInterval.Version.ValueInt64())
		queryParams1.OperationIDs = elementsToStrings(ctx, organizationsAPIRequestsOverviewResponseCodesByInterval.OperationIDs)
		queryParams1.SourceIPs = elementsToStrings(ctx, organizationsAPIRequestsOverviewResponseCodesByInterval.SourceIPs)
		queryParams1.AdminIDs = elementsToStrings(ctx, organizationsAPIRequestsOverviewResponseCodesByInterval.AdminIDs)
		queryParams1.UserAgent = organizationsAPIRequestsOverviewResponseCodesByInterval.UserAgent.ValueString()

		response1, restyResp1, err := d.client.Organizations.GetOrganizationAPIRequestsOverviewResponseCodesByInterval(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAPIRequestsOverviewResponseCodesByInterval",
				err.Error(),
			)
			return
		}

		organizationsAPIRequestsOverviewResponseCodesByInterval = ResponseOrganizationsGetOrganizationAPIRequestsOverviewResponseCodesByIntervalItemsToBody(organizationsAPIRequestsOverviewResponseCodesByInterval, response1)
		diags = resp.State.Set(ctx, &organizationsAPIRequestsOverviewResponseCodesByInterval)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsAPIRequestsOverviewResponseCodesByInterval struct {
	OrganizationID types.String                                                                          `tfsdk:"organization_id"`
	T0             types.String                                                                          `tfsdk:"t0"`
	T1             types.String                                                                          `tfsdk:"t1"`
	Timespan       types.Float64                                                                         `tfsdk:"timespan"`
	Interval       types.Int64                                                                           `tfsdk:"interval"`
	Version        types.Int64                                                                           `tfsdk:"version"`
	OperationIDs   types.List                                                                            `tfsdk:"operation_ids"`
	SourceIPs      types.List                                                                            `tfsdk:"source_ips"`
	AdminIDs       types.List                                                                            `tfsdk:"admin_ids"`
	UserAgent      types.String                                                                          `tfsdk:"user_agent"`
	Items          *[]ResponseItemOrganizationsGetOrganizationApiRequestsOverviewResponseCodesByInterval `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationApiRequestsOverviewResponseCodesByInterval struct {
	Counts  *[]ResponseItemOrganizationsGetOrganizationApiRequestsOverviewResponseCodesByIntervalCounts `tfsdk:"counts"`
	EndTs   types.String                                                                                `tfsdk:"end_ts"`
	StartTs types.String                                                                                `tfsdk:"start_ts"`
}

type ResponseItemOrganizationsGetOrganizationApiRequestsOverviewResponseCodesByIntervalCounts struct {
	Code  types.Int64 `tfsdk:"code"`
	Count types.Int64 `tfsdk:"count"`
}

// ToBody
func ResponseOrganizationsGetOrganizationAPIRequestsOverviewResponseCodesByIntervalItemsToBody(state OrganizationsAPIRequestsOverviewResponseCodesByInterval, response *merakigosdk.ResponseOrganizationsGetOrganizationAPIRequestsOverviewResponseCodesByInterval) OrganizationsAPIRequestsOverviewResponseCodesByInterval {
	var items []ResponseItemOrganizationsGetOrganizationApiRequestsOverviewResponseCodesByInterval
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationApiRequestsOverviewResponseCodesByInterval{
			Counts: func() *[]ResponseItemOrganizationsGetOrganizationApiRequestsOverviewResponseCodesByIntervalCounts {
				if item.Counts != nil {
					result := make([]ResponseItemOrganizationsGetOrganizationApiRequestsOverviewResponseCodesByIntervalCounts, len(*item.Counts))
					for i, counts := range *item.Counts {
						result[i] = ResponseItemOrganizationsGetOrganizationApiRequestsOverviewResponseCodesByIntervalCounts{
							Code: func() types.Int64 {
								if counts.Code != nil {
									return types.Int64Value(int64(*counts.Code))
								}
								return types.Int64{}
							}(),
							Count: func() types.Int64 {
								if counts.Count != nil {
									return types.Int64Value(int64(*counts.Count))
								}
								return types.Int64{}
							}(),
						}
					}
					return &result
				}
				return &[]ResponseItemOrganizationsGetOrganizationApiRequestsOverviewResponseCodesByIntervalCounts{}
			}(),
			EndTs:   types.StringValue(item.EndTs),
			StartTs: types.StringValue(item.StartTs),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
