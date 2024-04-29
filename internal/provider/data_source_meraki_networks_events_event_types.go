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
	_ datasource.DataSource              = &NetworksEventsEventTypesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksEventsEventTypesDataSource{}
)

func NewNetworksEventsEventTypesDataSource() datasource.DataSource {
	return &NetworksEventsEventTypesDataSource{}
}

type NetworksEventsEventTypesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksEventsEventTypesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksEventsEventTypesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_events_event_types"
}

func (d *NetworksEventsEventTypesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseNetworksGetNetworkEventsEventTypes`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"category": schema.StringAttribute{
							MarkdownDescription: `Event category`,
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: `Description of the event`,
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `Event type`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksEventsEventTypesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksEventsEventTypes NetworksEventsEventTypes
	diags := req.Config.Get(ctx, &networksEventsEventTypes)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkEventsEventTypes")
		vvNetworkID := networksEventsEventTypes.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Networks.GetNetworkEventsEventTypes(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkEventsEventTypes",
				err.Error(),
			)
			return
		}

		networksEventsEventTypes = ResponseNetworksGetNetworkEventsEventTypesItemsToBody(networksEventsEventTypes, response1)
		diags = resp.State.Set(ctx, &networksEventsEventTypes)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksEventsEventTypes struct {
	NetworkID types.String                                      `tfsdk:"network_id"`
	Items     *[]ResponseItemNetworksGetNetworkEventsEventTypes `tfsdk:"items"`
}

type ResponseItemNetworksGetNetworkEventsEventTypes struct {
	Category    types.String `tfsdk:"category"`
	Description types.String `tfsdk:"description"`
	Type        types.String `tfsdk:"type"`
}

// ToBody
func ResponseNetworksGetNetworkEventsEventTypesItemsToBody(state NetworksEventsEventTypes, response *merakigosdk.ResponseNetworksGetNetworkEventsEventTypes) NetworksEventsEventTypes {
	var items []ResponseItemNetworksGetNetworkEventsEventTypes
	for _, item := range *response {
		itemState := ResponseItemNetworksGetNetworkEventsEventTypes{
			Category:    types.StringValue(item.Category),
			Description: types.StringValue(item.Description),
			Type:        types.StringValue(item.Type),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
