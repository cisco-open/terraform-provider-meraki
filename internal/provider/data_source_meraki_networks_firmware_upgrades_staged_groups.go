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
	_ datasource.DataSource              = &NetworksFirmwareUpgradesStagedGroupsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksFirmwareUpgradesStagedGroupsDataSource{}
)

func NewNetworksFirmwareUpgradesStagedGroupsDataSource() datasource.DataSource {
	return &NetworksFirmwareUpgradesStagedGroupsDataSource{}
}

type NetworksFirmwareUpgradesStagedGroupsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksFirmwareUpgradesStagedGroupsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksFirmwareUpgradesStagedGroupsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_firmware_upgrades_staged_groups"
}

func (d *NetworksFirmwareUpgradesStagedGroupsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"group_id": schema.StringAttribute{
				MarkdownDescription: `groupId path parameter. Group ID`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"assigned_devices": schema.SingleNestedAttribute{
						MarkdownDescription: `The devices and Switch Stacks assigned to the Group`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"devices": schema.SetNestedAttribute{
								MarkdownDescription: `Data Array of Devices containing the name and serial`,
								Computed:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{

										"name": schema.StringAttribute{
											MarkdownDescription: `Name of the device`,
											Computed:            true,
										},
										"serial": schema.StringAttribute{
											MarkdownDescription: `Serial of the device`,
											Computed:            true,
										},
									},
								},
							},
							"switch_stacks": schema.SetNestedAttribute{
								MarkdownDescription: `Data Array of Switch Stacks containing the name and id`,
								Computed:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `ID of the Switch Stack`,
											Computed:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: `Name of the Switch Stack`,
											Computed:            true,
										},
									},
								},
							},
						},
					},
					"description": schema.StringAttribute{
						MarkdownDescription: `Description of the Staged Upgrade Group`,
						Computed:            true,
					},
					"group_id": schema.StringAttribute{
						MarkdownDescription: `Id of staged upgrade group`,
						Computed:            true,
					},
					"is_default": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating the default Group. Any device that does not have a group explicitly assigned will upgrade with this group`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of the Staged Upgrade Group`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseNetworksGetNetworkFirmwareUpgradesStagedGroups`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"assigned_devices": schema.SingleNestedAttribute{
							MarkdownDescription: `The devices and Switch Stacks assigned to the Group`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"devices": schema.SetNestedAttribute{
									MarkdownDescription: `Data Array of Devices containing the name and serial`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"name": schema.StringAttribute{
												MarkdownDescription: `Name of the device`,
												Computed:            true,
											},
											"serial": schema.StringAttribute{
												MarkdownDescription: `Serial of the device`,
												Computed:            true,
											},
										},
									},
								},
								"switch_stacks": schema.SetNestedAttribute{
									MarkdownDescription: `Data Array of Switch Stacks containing the name and id`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"id": schema.StringAttribute{
												MarkdownDescription: `ID of the Switch Stack`,
												Computed:            true,
											},
											"name": schema.StringAttribute{
												MarkdownDescription: `Name of the Switch Stack`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
						"description": schema.StringAttribute{
							MarkdownDescription: `Description of the Staged Upgrade Group`,
							Computed:            true,
						},
						"group_id": schema.StringAttribute{
							MarkdownDescription: `Id of staged upgrade group`,
							Computed:            true,
						},
						"is_default": schema.BoolAttribute{
							MarkdownDescription: `Boolean indicating the default Group. Any device that does not have a group explicitly assigned will upgrade with this group`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the Staged Upgrade Group`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksFirmwareUpgradesStagedGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksFirmwareUpgradesStagedGroups NetworksFirmwareUpgradesStagedGroups
	diags := req.Config.Get(ctx, &networksFirmwareUpgradesStagedGroups)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksFirmwareUpgradesStagedGroups.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksFirmwareUpgradesStagedGroups.NetworkID.IsNull(), !networksFirmwareUpgradesStagedGroups.GroupID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkFirmwareUpgradesStagedGroups")
		vvNetworkID := networksFirmwareUpgradesStagedGroups.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Networks.GetNetworkFirmwareUpgradesStagedGroups(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkFirmwareUpgradesStagedGroups",
				err.Error(),
			)
			return
		}

		networksFirmwareUpgradesStagedGroups = ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupsItemsToBody(networksFirmwareUpgradesStagedGroups, response1)
		diags = resp.State.Set(ctx, &networksFirmwareUpgradesStagedGroups)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkFirmwareUpgradesStagedGroup")
		vvNetworkID := networksFirmwareUpgradesStagedGroups.NetworkID.ValueString()
		vvGroupID := networksFirmwareUpgradesStagedGroups.GroupID.ValueString()

		response2, restyResp2, err := d.client.Networks.GetNetworkFirmwareUpgradesStagedGroup(vvNetworkID, vvGroupID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkFirmwareUpgradesStagedGroup",
				err.Error(),
			)
			return
		}

		networksFirmwareUpgradesStagedGroups = ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupItemToBody(networksFirmwareUpgradesStagedGroups, response2)
		diags = resp.State.Set(ctx, &networksFirmwareUpgradesStagedGroups)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksFirmwareUpgradesStagedGroups struct {
	NetworkID types.String                                                  `tfsdk:"network_id"`
	GroupID   types.String                                                  `tfsdk:"group_id"`
	Items     *[]ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroups `tfsdk:"items"`
	Item      *ResponseNetworksGetNetworkFirmwareUpgradesStagedGroup        `tfsdk:"item"`
}

type ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroups struct {
	AssignedDevices *ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevices `tfsdk:"assigned_devices"`
	Description     types.String                                                               `tfsdk:"description"`
	GroupID         types.String                                                               `tfsdk:"group_id"`
	IsDefault       types.Bool                                                                 `tfsdk:"is_default"`
	Name            types.String                                                               `tfsdk:"name"`
}

type ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevices struct {
	Devices      *[]ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevicesDevices      `tfsdk:"devices"`
	SwitchStacks *[]ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks `tfsdk:"switch_stacks"`
}

type ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevicesDevices struct {
	Name   types.String `tfsdk:"name"`
	Serial types.String `tfsdk:"serial"`
}

type ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedGroup struct {
	AssignedDevices *ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevices `tfsdk:"assigned_devices"`
	Description     types.String                                                          `tfsdk:"description"`
	GroupID         types.String                                                          `tfsdk:"group_id"`
	IsDefault       types.Bool                                                            `tfsdk:"is_default"`
	Name            types.String                                                          `tfsdk:"name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevices struct {
	Devices      *[]ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices      `tfsdk:"devices"`
	SwitchStacks *[]ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks `tfsdk:"switch_stacks"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices struct {
	Name   types.String `tfsdk:"name"`
	Serial types.String `tfsdk:"serial"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// ToBody
func ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupsItemsToBody(state NetworksFirmwareUpgradesStagedGroups, response *merakigosdk.ResponseNetworksGetNetworkFirmwareUpgradesStagedGroups) NetworksFirmwareUpgradesStagedGroups {
	var items []ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroups
	for _, item := range *response {
		itemState := ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroups{
			AssignedDevices: func() *ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevices {
				if item.AssignedDevices != nil {
					return &ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevices{
						Devices: func() *[]ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevicesDevices {
							if item.AssignedDevices.Devices != nil {
								result := make([]ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevicesDevices, len(*item.AssignedDevices.Devices))
								for i, devices := range *item.AssignedDevices.Devices {
									result[i] = ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevicesDevices{
										Name:   types.StringValue(devices.Name),
										Serial: types.StringValue(devices.Serial),
									}
								}
								return &result
							}
							return &[]ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevicesDevices{}
						}(),
						SwitchStacks: func() *[]ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks {
							if item.AssignedDevices.SwitchStacks != nil {
								result := make([]ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks, len(*item.AssignedDevices.SwitchStacks))
								for i, switchStacks := range *item.AssignedDevices.SwitchStacks {
									result[i] = ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks{
										ID:   types.StringValue(switchStacks.ID),
										Name: types.StringValue(switchStacks.Name),
									}
								}
								return &result
							}
							return &[]ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks{}
						}(),
					}
				}
				return &ResponseItemNetworksGetNetworkFirmwareUpgradesStagedGroupsAssignedDevices{}
			}(),
			Description: types.StringValue(item.Description),
			GroupID:     types.StringValue(item.GroupID),
			IsDefault: func() types.Bool {
				if item.IsDefault != nil {
					return types.BoolValue(*item.IsDefault)
				}
				return types.Bool{}
			}(),
			Name: types.StringValue(item.Name),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupItemToBody(state NetworksFirmwareUpgradesStagedGroups, response *merakigosdk.ResponseNetworksGetNetworkFirmwareUpgradesStagedGroup) NetworksFirmwareUpgradesStagedGroups {
	itemState := ResponseNetworksGetNetworkFirmwareUpgradesStagedGroup{
		AssignedDevices: func() *ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevices {
			if response.AssignedDevices != nil {
				return &ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevices{
					Devices: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices {
						if response.AssignedDevices.Devices != nil {
							result := make([]ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices, len(*response.AssignedDevices.Devices))
							for i, devices := range *response.AssignedDevices.Devices {
								result[i] = ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices{
									Name:   types.StringValue(devices.Name),
									Serial: types.StringValue(devices.Serial),
								}
							}
							return &result
						}
						return &[]ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices{}
					}(),
					SwitchStacks: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks {
						if response.AssignedDevices.SwitchStacks != nil {
							result := make([]ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks, len(*response.AssignedDevices.SwitchStacks))
							for i, switchStacks := range *response.AssignedDevices.SwitchStacks {
								result[i] = ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks{
									ID:   types.StringValue(switchStacks.ID),
									Name: types.StringValue(switchStacks.Name),
								}
							}
							return &result
						}
						return &[]ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevices{}
		}(),
		Description: types.StringValue(response.Description),
		GroupID:     types.StringValue(response.GroupID),
		IsDefault: func() types.Bool {
			if response.IsDefault != nil {
				return types.BoolValue(*response.IsDefault)
			}
			return types.Bool{}
		}(),
		Name: types.StringValue(response.Name),
	}
	state.Item = &itemState
	return state
}
