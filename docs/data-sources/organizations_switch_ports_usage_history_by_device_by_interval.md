---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_switch_ports_usage_history_by_device_by_interval Data Source - terraform-provider-meraki"
subcategory: "switch"
description: |-
  
---

# meraki_organizations_switch_ports_usage_history_by_device_by_interval (Data Source)



## Example Usage

```terraform
data "meraki_organizations_switch_ports_usage_history_by_device_by_interval" "example" {

  configuration_updated_after = "string"
  ending_before               = "string"
  interval                    = 1
  mac                         = "string"
  macs                        = ["string"]
  name                        = "string"
  network_ids                 = ["string"]
  organization_id             = "string"
  per_page                    = 1
  port_profile_ids            = ["string"]
  serial                      = "string"
  serials                     = ["string"]
  starting_after              = "string"
  t0                          = "string"
  t1                          = "string"
  timespan                    = 1.0
}

output "meraki_organizations_switch_ports_usage_history_by_device_by_interval_example" {
  value = data.meraki_organizations_switch_ports_usage_history_by_device_by_interval.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Optional

- `configuration_updated_after` (String) configurationUpdatedAfter query parameter. Optional parameter to filter items to switches where the configuration has been updated after the given timestamp.
- `ending_before` (String) endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `interval` (Number) interval query parameter. The time interval in seconds for returned data. The valid intervals are: 300, 1200, 14400, 86400. The default is 1200. Interval is calculated if time params are provided.
- `mac` (String) mac query parameter. Optional parameter to filter items to switches with MAC addresses that contain the search term or are an exact match.
- `macs` (List of String) macs query parameter. Optional parameter to filter items to switches that have one of the provided MAC addresses.
- `name` (String) name query parameter. Optional parameter to filter items to switches with names that contain the search term or are an exact match.
- `network_ids` (List of String) networkIds query parameter. Optional parameter to filter items to switches in one of the provided networks.
- `per_page` (Number) perPage query parameter. The number of entries per page returned. Acceptable range is 3 50. Default is 10.
- `port_profile_ids` (List of String) portProfileIds query parameter. Optional parameter to filter items to switches that contain switchports belonging to one of the specified port profiles.
- `serial` (String) serial query parameter. Optional parameter to filter items to switches with serial number that contains the search term or are an exact match.
- `serials` (List of String) serials query parameter. Optional parameter to filter items to switches that have one of the provided serials.
- `starting_after` (String) startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `t0` (String) t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
- `t1` (String) t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
- `timespan` (Number) timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. If interval is provided, the timespan will be autocalculated.

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `items` (Attributes List) Switches (see [below for nested schema](#nestedatt--item--items))
- `meta` (Attributes) Metadata relevant to the paginated dataset (see [below for nested schema](#nestedatt--item--meta))

<a id="nestedatt--item--items"></a>
### Nested Schema for `item.items`

Read-Only:

- `mac` (String) The MAC address of the switch.
- `model` (String) The model of the switch.
- `name` (String) The name of the switch.
- `network` (Attributes) Identifying information of the switch's network. (see [below for nested schema](#nestedatt--item--items--network))
- `ports` (Attributes Set) The number of ports on the switch with usage data. (see [below for nested schema](#nestedatt--item--items--ports))
- `serial` (String) The serial number of the switch.

<a id="nestedatt--item--items--network"></a>
### Nested Schema for `item.items.network`

Read-Only:

- `id` (String) The ID of the network.
- `name` (String) The name of the network.


<a id="nestedatt--item--items--ports"></a>
### Nested Schema for `item.items.ports`

Read-Only:

- `intervals` (Attributes Set) An array of intervals for a port with bandwidth, traffic, and power usage data. (see [below for nested schema](#nestedatt--item--items--ports--intervals))
- `port_id` (String) The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module.

<a id="nestedatt--item--items--ports--intervals"></a>
### Nested Schema for `item.items.ports.intervals`

Read-Only:

- `bandwidth` (Attributes) A breakdown of the average speed of data that has passed through this port during the interval. (see [below for nested schema](#nestedatt--item--items--ports--intervals--bandwidth))
- `data` (Attributes) A breakdown of how many kilobytes have passed through this port during the interval timespan. (see [below for nested schema](#nestedatt--item--items--ports--intervals--data))
- `end_ts` (String) The end timestamp of the given interval.
- `energy` (Attributes) How much energy (in watt-hours) has been delivered by this port during the interval. (see [below for nested schema](#nestedatt--item--items--ports--intervals--energy))
- `start_ts` (String) The starting timestamp of the given interval.

<a id="nestedatt--item--items--ports--intervals--bandwidth"></a>
### Nested Schema for `item.items.ports.intervals.bandwidth`

Read-Only:

- `usage` (Attributes) Bandwidth usage data for the given interval. (see [below for nested schema](#nestedatt--item--items--ports--intervals--bandwidth--usage))

<a id="nestedatt--item--items--ports--intervals--bandwidth--usage"></a>
### Nested Schema for `item.items.ports.intervals.bandwidth.usage`

Read-Only:

- `downstream` (Number) The average speed of the data received (in kilobits-per-second).
- `total` (Number) The average speed of the data sent and received (in kilobits-per-second).
- `upstream` (Number) The average speed of the data sent (in kilobits-per-second).



<a id="nestedatt--item--items--ports--intervals--data"></a>
### Nested Schema for `item.items.ports.intervals.data`

Read-Only:

- `usage` (Attributes) Usage data for the given interval. (see [below for nested schema](#nestedatt--item--items--ports--intervals--data--usage))

<a id="nestedatt--item--items--ports--intervals--data--usage"></a>
### Nested Schema for `item.items.ports.intervals.data.usage`

Read-Only:

- `downstream` (Number) The amount of data received (in kilobytes).
- `total` (Number) The total amount of data sent and received (in kilobytes).
- `upstream` (Number) The amount of data sent (in kilobytes).



<a id="nestedatt--item--items--ports--intervals--energy"></a>
### Nested Schema for `item.items.ports.intervals.energy`

Read-Only:

- `usage` (Attributes) Energy data for the given interval. (see [below for nested schema](#nestedatt--item--items--ports--intervals--energy--usage))

<a id="nestedatt--item--items--ports--intervals--energy--usage"></a>
### Nested Schema for `item.items.ports.intervals.energy.usage`

Read-Only:

- `total` (Number) The total energy in watt-hours delivered by this port during the interval






<a id="nestedatt--item--meta"></a>
### Nested Schema for `item.meta`

Read-Only:

- `counts` (Attributes) Counts relating to the paginated dataset (see [below for nested schema](#nestedatt--item--meta--counts))

<a id="nestedatt--item--meta--counts"></a>
### Nested Schema for `item.meta.counts`

Read-Only:

- `items` (Attributes) Counts relating to the paginated items (see [below for nested schema](#nestedatt--item--meta--counts--items))

<a id="nestedatt--item--meta--counts--items"></a>
### Nested Schema for `item.meta.counts.items`

Read-Only:

- `remaining` (Number) The number of items in the dataset that are available on subsequent pages
- `total` (Number) The total number of items in the dataset
