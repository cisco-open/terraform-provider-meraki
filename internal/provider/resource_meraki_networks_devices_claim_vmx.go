package provider

// RESOURCE ACTION

import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksDevicesClaimVmxResource{}
	_ resource.ResourceWithConfigure = &NetworksDevicesClaimVmxResource{}
)

func NewNetworksDevicesClaimVmxResource() resource.Resource {
	return &NetworksDevicesClaimVmxResource{}
}

type NetworksDevicesClaimVmxResource struct {
	client *merakigosdk.Client
}

func (r *NetworksDevicesClaimVmxResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksDevicesClaimVmxResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_devices_claim_vmx"
}

// resourceAction
func (r *NetworksDevicesClaimVmxResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"address": schema.StringAttribute{
						MarkdownDescription: `Physical address of the device`,
						Computed:            true,
					},
					"details": schema.SetNestedAttribute{
						MarkdownDescription: `Additional device information`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"name": schema.StringAttribute{
									MarkdownDescription: `Additional property name`,
									Computed:            true,
								},
								"value": schema.StringAttribute{
									MarkdownDescription: `Additional property value`,
									Computed:            true,
								},
							},
						},
					},
					"firmware": schema.StringAttribute{
						MarkdownDescription: `Firmware version of the device`,
						Computed:            true,
					},
					"imei": schema.StringAttribute{
						MarkdownDescription: `IMEI of the device, if applicable`,
						Computed:            true,
					},
					"lan_ip": schema.StringAttribute{
						MarkdownDescription: `LAN IP address of the device`,
						Computed:            true,
					},
					"lat": schema.Float64Attribute{
						MarkdownDescription: `Latitude of the device`,
						Computed:            true,
					},
					"lng": schema.Float64Attribute{
						MarkdownDescription: `Longitude of the device`,
						Computed:            true,
					},
					"mac": schema.StringAttribute{
						MarkdownDescription: `MAC address of the device`,
						Computed:            true,
					},
					"model": schema.StringAttribute{
						MarkdownDescription: `Model of the device`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of the device`,
						Computed:            true,
					},
					"network_id": schema.StringAttribute{
						MarkdownDescription: `ID of the network the device belongs to`,
						Computed:            true,
					},
					"notes": schema.StringAttribute{
						MarkdownDescription: `Notes for the device, limited to 255 characters`,
						Computed:            true,
					},
					"product_type": schema.StringAttribute{
						MarkdownDescription: `Product type of the device`,
						Computed:            true,
					},
					"serial": schema.StringAttribute{
						MarkdownDescription: `Serial number of the device`,
						Computed:            true,
					},
					"tags": schema.SetAttribute{
						MarkdownDescription: `List of tags assigned to the device`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"size": schema.StringAttribute{
						MarkdownDescription: `The size of the vMX you claim. It can be one of: small, medium, large, xlarge, 100`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
				},
			},
		},
	}
}
func (r *NetworksDevicesClaimVmxResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksDevicesClaimVmx

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
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Networks.VmxNetworkDevicesClaim(vvNetworkID, dataRequest)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing VmxNetworkDevicesClaim",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing VmxNetworkDevicesClaim",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseNetworksVmxNetworkDevicesClaimItemToBody(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksDevicesClaimVmxResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksDevicesClaimVmxResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksDevicesClaimVmxResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksDevicesClaimVmx struct {
	NetworkID  types.String                             `tfsdk:"network_id"`
	Item       *ResponseNetworksVmxNetworkDevicesClaim  `tfsdk:"item"`
	Parameters *RequestNetworksVmxNetworkDevicesClaimRs `tfsdk:"parameters"`
}

type ResponseNetworksVmxNetworkDevicesClaim struct {
	Address     types.String                                     `tfsdk:"address"`
	Details     *[]ResponseNetworksVmxNetworkDevicesClaimDetails `tfsdk:"details"`
	Firmware    types.String                                     `tfsdk:"firmware"`
	Imei        types.String                                     `tfsdk:"imei"`
	LanIP       types.String                                     `tfsdk:"lan_ip"`
	Lat         types.Float64                                    `tfsdk:"lat"`
	Lng         types.Float64                                    `tfsdk:"lng"`
	Mac         types.String                                     `tfsdk:"mac"`
	Model       types.String                                     `tfsdk:"model"`
	Name        types.String                                     `tfsdk:"name"`
	NetworkID   types.String                                     `tfsdk:"network_id"`
	Notes       types.String                                     `tfsdk:"notes"`
	ProductType types.String                                     `tfsdk:"product_type"`
	Serial      types.String                                     `tfsdk:"serial"`
	Tags        types.Set                                        `tfsdk:"tags"`
}

type ResponseNetworksVmxNetworkDevicesClaimDetails struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type RequestNetworksVmxNetworkDevicesClaimRs struct {
	Size types.String `tfsdk:"size"`
}

// FromBody
func (r *NetworksDevicesClaimVmx) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksVmxNetworkDevicesClaim {
	emptyString := ""
	re := *r.Parameters
	size := new(string)
	if !re.Size.IsUnknown() && !re.Size.IsNull() {
		*size = re.Size.ValueString()
	} else {
		size = &emptyString
	}
	out := merakigosdk.RequestNetworksVmxNetworkDevicesClaim{
		Size: *size,
	}
	return &out
}

// ToBody
func ResponseNetworksVmxNetworkDevicesClaimItemToBody(state NetworksDevicesClaimVmx, response *merakigosdk.ResponseNetworksVmxNetworkDevicesClaim) NetworksDevicesClaimVmx {
	itemState := ResponseNetworksVmxNetworkDevicesClaim{
		Address: types.StringValue(response.Address),
		Details: func() *[]ResponseNetworksVmxNetworkDevicesClaimDetails {
			if response.Details != nil {
				result := make([]ResponseNetworksVmxNetworkDevicesClaimDetails, len(*response.Details))
				for i, details := range *response.Details {
					result[i] = ResponseNetworksVmxNetworkDevicesClaimDetails{
						Name:  types.StringValue(details.Name),
						Value: types.StringValue(details.Value),
					}
				}
				return &result
			}
			return &[]ResponseNetworksVmxNetworkDevicesClaimDetails{}
		}(),
		Firmware: types.StringValue(response.Firmware),
		Imei:     types.StringValue(response.Imei),
		LanIP:    types.StringValue(response.LanIP),
		Lat: func() types.Float64 {
			if response.Lat != nil {
				return types.Float64Value(float64(*response.Lat))
			}
			return types.Float64{}
		}(),
		Lng: func() types.Float64 {
			if response.Lng != nil {
				return types.Float64Value(float64(*response.Lng))
			}
			return types.Float64{}
		}(),
		Mac:         types.StringValue(response.Mac),
		Model:       types.StringValue(response.Model),
		Name:        types.StringValue(response.Name),
		NetworkID:   types.StringValue(response.NetworkID),
		Notes:       types.StringValue(response.Notes),
		ProductType: types.StringValue(response.ProductType),
		Serial:      types.StringValue(response.Serial),
		Tags:        StringSliceToSet(response.Tags),
	}
	state.Item = &itemState
	return state
}
