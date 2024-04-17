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
	_ datasource.DataSource              = &NetworksSwitchStacksDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchStacksDataSource{}
)

func NewNetworksSwitchStacksDataSource() datasource.DataSource {
	return &NetworksSwitchStacksDataSource{}
}

type NetworksSwitchStacksDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchStacksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchStacksDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_stacks"
}

func (d *NetworksSwitchStacksDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"switch_stack_id": schema.StringAttribute{
				MarkdownDescription: `switchStackId path parameter. Switch stack ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"id": schema.StringAttribute{
						MarkdownDescription: `ID of the Switch stack`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of the Switch stack`,
						Computed:            true,
					},
					"serials": schema.ListAttribute{
						MarkdownDescription: `Serials of the switches in the switch stack`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSwitchGetNetworkSwitchStacks`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"id": schema.StringAttribute{
							MarkdownDescription: `ID of the Switch stack`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the Switch stack`,
							Computed:            true,
						},
						"serials": schema.ListAttribute{
							MarkdownDescription: `Serials of the switches in the switch stack`,
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchStacksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchStacks NetworksSwitchStacks
	diags := req.Config.Get(ctx, &networksSwitchStacks)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksSwitchStacks.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksSwitchStacks.NetworkID.IsNull(), !networksSwitchStacks.SwitchStackID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchStacks")
		vvNetworkID := networksSwitchStacks.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchStacks(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStacks",
				err.Error(),
			)
			return
		}

		networksSwitchStacks = ResponseSwitchGetNetworkSwitchStacksItemsToBody(networksSwitchStacks, response1)
		diags = resp.State.Set(ctx, &networksSwitchStacks)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchStack")
		vvNetworkID := networksSwitchStacks.NetworkID.ValueString()
		vvSwitchStackID := networksSwitchStacks.SwitchStackID.ValueString()

		response2, restyResp2, err := d.client.Switch.GetNetworkSwitchStack(vvNetworkID, vvSwitchStackID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStack",
				err.Error(),
			)
			return
		}

		networksSwitchStacks = ResponseSwitchGetNetworkSwitchStackItemToBody(networksSwitchStacks, response2)
		diags = resp.State.Set(ctx, &networksSwitchStacks)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchStacks struct {
	NetworkID     types.String                                `tfsdk:"network_id"`
	SwitchStackID types.String                                `tfsdk:"switch_stack_id"`
	Items         *[]ResponseItemSwitchGetNetworkSwitchStacks `tfsdk:"items"`
	Item          *ResponseSwitchGetNetworkSwitchStack        `tfsdk:"item"`
}

type ResponseItemSwitchGetNetworkSwitchStacks struct {
	ID      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Serials types.List   `tfsdk:"serials"`
}

type ResponseSwitchGetNetworkSwitchStack struct {
	ID      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Serials types.List   `tfsdk:"serials"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchStacksItemsToBody(state NetworksSwitchStacks, response *merakigosdk.ResponseSwitchGetNetworkSwitchStacks) NetworksSwitchStacks {
	var items []ResponseItemSwitchGetNetworkSwitchStacks
	for _, item := range *response {
		itemState := ResponseItemSwitchGetNetworkSwitchStacks{
			ID:      types.StringValue(item.ID),
			Name:    types.StringValue(item.Name),
			Serials: StringSliceToList(item.Serials),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseSwitchGetNetworkSwitchStackItemToBody(state NetworksSwitchStacks, response *merakigosdk.ResponseSwitchGetNetworkSwitchStack) NetworksSwitchStacks {
	itemState := ResponseSwitchGetNetworkSwitchStack{
		ID:      types.StringValue(response.ID),
		Name:    types.StringValue(response.Name),
		Serials: StringSliceToList(response.Serials),
	}
	state.Item = &itemState
	return state
}
