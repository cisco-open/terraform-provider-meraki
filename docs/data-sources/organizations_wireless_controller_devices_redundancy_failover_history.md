---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_wireless_controller_devices_redundancy_failover_history Data Source - terraform-provider-meraki"
subcategory: "wirelessController"
description: |-
  
---

# meraki_organizations_wireless_controller_devices_redundancy_failover_history (Data Source)



## Example Usage

```terraform
data "meraki_organizations_wireless_controller_devices_redundancy_failover_history" "example" {

  ending_before   = "string"
  organization_id = "string"
  per_page        = 1
  serials         = ["string"]
  starting_after  = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
}

output "meraki_organizations_wireless_controller_devices_redundancy_failover_history_example" {
  value = data.meraki_organizations_wireless_controller_devices_redundancy_failover_history.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Optional

- `ending_before` (String) endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `per_page` (Number) perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.
- `serials` (List of String) serials query parameter. Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches.
- `starting_after` (String) startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `t0` (String) t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
- `t1` (String) t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
- `timespan` (Number) timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.

### Read-Only

- `items` (Attributes List) Array of ResponseWirelessControllerGetOrganizationWirelessControllerDevicesRedundancyFailoverHistory (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `items` (Attributes List) Wireless LAN controller HA failover events (see [below for nested schema](#nestedatt--items--items))
- `meta` (Attributes) Metadata relevant to the paginated dataset (see [below for nested schema](#nestedatt--items--meta))

<a id="nestedatt--items--items"></a>
### Nested Schema for `items.items`

Read-Only:

- `active` (Attributes) Details about the active unit (see [below for nested schema](#nestedatt--items--items--active))
- `failed` (Attributes) Details about the failed unit (see [below for nested schema](#nestedatt--items--items--failed))
- `reason` (String) Failover reason
- `serial` (String) Wireless LAN controller cloud ID
- `ts` (String) Failover time

<a id="nestedatt--items--items--active"></a>
### Nested Schema for `items.items.active`

Read-Only:

- `chassis` (Attributes) Details about the active unit chassis (see [below for nested schema](#nestedatt--items--items--active--chassis))

<a id="nestedatt--items--items--active--chassis"></a>
### Nested Schema for `items.items.active.chassis`

Read-Only:

- `name` (String) The name of the active chassis unit



<a id="nestedatt--items--items--failed"></a>
### Nested Schema for `items.items.failed`

Read-Only:

- `chassis` (Attributes) Details about the failed unit chassis (see [below for nested schema](#nestedatt--items--items--failed--chassis))

<a id="nestedatt--items--items--failed--chassis"></a>
### Nested Schema for `items.items.failed.chassis`

Read-Only:

- `name` (String) The name of the failed chassis unit




<a id="nestedatt--items--meta"></a>
### Nested Schema for `items.meta`

Read-Only:

- `counts` (Attributes) Counts relating to the paginated dataset (see [below for nested schema](#nestedatt--items--meta--counts))

<a id="nestedatt--items--meta--counts"></a>
### Nested Schema for `items.meta.counts`

Read-Only:

- `items` (Attributes) Counts relating to the paginated items (see [below for nested schema](#nestedatt--items--meta--counts--items))

<a id="nestedatt--items--meta--counts--items"></a>
### Nested Schema for `items.meta.counts.items`

Read-Only:

- `remaining` (Number) The number of items in the dataset that are available on subsequent pages
- `total` (Number) The total number of items in the dataset
