package provider

// RESOURCE ACTION

import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSmDevicesWipeResource{}
	_ resource.ResourceWithConfigure = &NetworksSmDevicesWipeResource{}
)

func NewNetworksSmDevicesWipeResource() resource.Resource {
	return &NetworksSmDevicesWipeResource{}
}

type NetworksSmDevicesWipeResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSmDevicesWipeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSmDevicesWipeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_devices_wipe"
}

// resourceAction
func (r *NetworksSmDevicesWipeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

					"id": schema.StringAttribute{
						MarkdownDescription: `The Meraki Id of the devices.`,
						Computed:            true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						MarkdownDescription: `The id of the device to be wiped.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"pin": schema.Int64Attribute{
						MarkdownDescription: `The pin number (a six digit value) for wiping a macOS device. Required only for macOS devices.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplace(),
						},
					},
					"serial": schema.StringAttribute{
						MarkdownDescription: `The serial of the device to be wiped.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"wifi_mac": schema.StringAttribute{
						MarkdownDescription: `The wifiMac of the device to be wiped.`,
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
func (r *NetworksSmDevicesWipeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSmDevicesWipe

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
	response, restyResp1, err := r.client.Sm.WipeNetworkSmDevices(vvNetworkID, dataRequest)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing WipeNetworkSmDevices",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing WipeNetworkSmDevices",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseSmWipeNetworkSmDevicesItemToBody(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSmDevicesWipeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSmDevicesWipeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSmDevicesWipeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSmDevicesWipe struct {
	NetworkID  types.String                     `tfsdk:"network_id"`
	Item       *ResponseSmWipeNetworkSmDevices  `tfsdk:"item"`
	Parameters *RequestSmWipeNetworkSmDevicesRs `tfsdk:"parameters"`
}

type ResponseSmWipeNetworkSmDevices struct {
	ID types.String `tfsdk:"id"`
}

type RequestSmWipeNetworkSmDevicesRs struct {
	ID      types.String `tfsdk:"id"`
	Pin     types.Int64  `tfsdk:"pin"`
	Serial  types.String `tfsdk:"serial"`
	WifiMac types.String `tfsdk:"wifi_mac"`
}

// FromBody
func (r *NetworksSmDevicesWipe) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSmWipeNetworkSmDevices {
	emptyString := ""
	re := *r.Parameters
	iD := new(string)
	if !re.ID.IsUnknown() && !re.ID.IsNull() {
		*iD = re.ID.ValueString()
	} else {
		iD = &emptyString
	}
	pin := new(int64)
	if !re.Pin.IsUnknown() && !re.Pin.IsNull() {
		*pin = re.Pin.ValueInt64()
	} else {
		pin = nil
	}
	serial := new(string)
	if !re.Serial.IsUnknown() && !re.Serial.IsNull() {
		*serial = re.Serial.ValueString()
	} else {
		serial = &emptyString
	}
	wifiMac := new(string)
	if !re.WifiMac.IsUnknown() && !re.WifiMac.IsNull() {
		*wifiMac = re.WifiMac.ValueString()
	} else {
		wifiMac = &emptyString
	}
	out := merakigosdk.RequestSmWipeNetworkSmDevices{
		ID:      *iD,
		Pin:     int64ToIntPointer(pin),
		Serial:  *serial,
		WifiMac: *wifiMac,
	}
	return &out
}

// ToBody
func ResponseSmWipeNetworkSmDevicesItemToBody(state NetworksSmDevicesWipe, response *merakigosdk.ResponseSmWipeNetworkSmDevices) NetworksSmDevicesWipe {
	itemState := ResponseSmWipeNetworkSmDevices{
		ID: types.StringValue(response.ID),
	}
	state.Item = &itemState
	return state
}
