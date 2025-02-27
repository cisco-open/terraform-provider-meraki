---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_wireless_ssids_schedules Data Source - terraform-provider-meraki"
subcategory: "wireless"
description: |-
  
---

# meraki_networks_wireless_ssids_schedules (Data Source)



## Example Usage

```terraform
data "meraki_networks_wireless_ssids_schedules" "example" {

  network_id = "string"
  number     = "string"
}

output "meraki_networks_wireless_ssids_schedules_example" {
  value = data.meraki_networks_wireless_ssids_schedules.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID
- `number` (String) number path parameter.

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `enabled` (Boolean) If true, the SSID outage schedule is enabled.
- `ranges` (Attributes Set) List of outage ranges. Has a start date and time, and end date and time. If this parameter is passed in along with rangesInSeconds parameter, this will take precedence. (see [below for nested schema](#nestedatt--item--ranges))
- `ranges_in_seconds` (Attributes Set) List of outage ranges in seconds since Sunday at Midnight. Has a start and end. If this parameter is passed in along with the ranges parameter, ranges will take precedence. (see [below for nested schema](#nestedatt--item--ranges_in_seconds))

<a id="nestedatt--item--ranges"></a>
### Nested Schema for `item.ranges`

Read-Only:

- `end_day` (String) Day of when the outage ends. Can be either full day name, or three letter abbreviation
- `end_time` (String) 24 hour time when the outage ends.
- `start_day` (String) Day of when the outage starts. Can be either full day name, or three letter abbreviation.
- `start_time` (String) 24 hour time when the outage starts.


<a id="nestedatt--item--ranges_in_seconds"></a>
### Nested Schema for `item.ranges_in_seconds`

Read-Only:

- `end` (Number) Seconds since Sunday at midnight when that outage range ends.
- `start` (Number) Seconds since Sunday at midnight when the outage range starts.
