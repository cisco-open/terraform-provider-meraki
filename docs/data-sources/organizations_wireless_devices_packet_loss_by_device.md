---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_wireless_devices_packet_loss_by_device Data Source - terraform-provider-meraki"
subcategory: "wireless"
description: |-
  
---

# meraki_organizations_wireless_devices_packet_loss_by_device (Data Source)



## Example Usage

```terraform
data "meraki_organizations_wireless_devices_packet_loss_by_device" "example" {

  bands           = ["string"]
  ending_before   = "string"
  network_ids     = ["string"]
  organization_id = "string"
  per_page        = 1
  serials         = ["string"]
  ssids           = ["string"]
  starting_after  = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
}

output "meraki_organizations_wireless_devices_packet_loss_by_device_example" {
  value = data.meraki_organizations_wireless_devices_packet_loss_by_device.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Optional

- `bands` (List of String) bands query parameter. Filter results by band. Valid bands are: 2.4, 5, and 6.
- `ending_before` (String) endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `network_ids` (List of String) networkIds query parameter. Filter results by network.
- `per_page` (Number) perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.
- `serials` (List of String) serials query parameter. Filter results by device.
- `ssids` (List of String) ssids query parameter. Filter results by SSID number.
- `starting_after` (String) startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `t0` (String) t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 90 days from today.
- `t1` (String) t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 90 days after t0.
- `timespan` (Number) timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days.

### Read-Only

- `items` (Attributes List) Array of ResponseWirelessGetOrganizationWirelessDevicesPacketLossByDevice (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `device` (Attributes) Device. (see [below for nested schema](#nestedatt--items--device))
- `downstream` (Attributes) Packets sent from an AP to a client. (see [below for nested schema](#nestedatt--items--downstream))
- `network` (Attributes) Network. (see [below for nested schema](#nestedatt--items--network))
- `upstream` (Attributes) Packets sent from a client to an AP. (see [below for nested schema](#nestedatt--items--upstream))

<a id="nestedatt--items--device"></a>
### Nested Schema for `items.device`

Read-Only:

- `mac` (String) MAC address
- `name` (String) Name
- `serial` (String) Serial Number


<a id="nestedatt--items--downstream"></a>
### Nested Schema for `items.downstream`

Read-Only:

- `loss_percentage` (Number) Percentage of lost packets.
- `lost` (Number) Total packets sent by an AP that did not reach the client.
- `total` (Number) Total packets received by a client.


<a id="nestedatt--items--network"></a>
### Nested Schema for `items.network`

Read-Only:

- `id` (String) Network ID.
- `name` (String) Name of the network.


<a id="nestedatt--items--upstream"></a>
### Nested Schema for `items.upstream`

Read-Only:

- `loss_percentage` (Number) Percentage of lost packets.
- `lost` (Number) Total packets sent by a client and did not reach the AP.
- `total` (Number) Total packets sent by a client to an AP.
