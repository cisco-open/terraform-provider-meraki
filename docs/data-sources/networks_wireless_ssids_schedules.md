---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_wireless_ssids_schedules Data Source - terraform-provider-meraki"
subcategory: ""
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

- `enabled` (Boolean)
- `ranges` (Attributes Set) (see [below for nested schema](#nestedatt--item--ranges))

<a id="nestedatt--item--ranges"></a>
### Nested Schema for `item.ranges`

Read-Only:

- `end_day` (String)
- `end_time` (String)
- `start_day` (String)
- `start_time` (String)
