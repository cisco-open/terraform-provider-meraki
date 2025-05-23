---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_appliance_traffic_shaping_rules Data Source - terraform-provider-meraki"
subcategory: "appliance"
description: |-
  
---

# meraki_networks_appliance_traffic_shaping_rules (Data Source)



## Example Usage

```terraform
data "meraki_networks_appliance_traffic_shaping_rules" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_traffic_shaping_rules_example" {
  value = data.meraki_networks_appliance_traffic_shaping_rules.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `default_rules_enabled` (Boolean)
- `rules` (Attributes Set) (see [below for nested schema](#nestedatt--item--rules))

<a id="nestedatt--item--rules"></a>
### Nested Schema for `item.rules`

Read-Only:

- `definitions` (Attributes Set) (see [below for nested schema](#nestedatt--item--rules--definitions))
- `dscp_tag_value` (Number)
- `per_client_bandwidth_limits` (Attributes) (see [below for nested schema](#nestedatt--item--rules--per_client_bandwidth_limits))
- `priority` (String)

<a id="nestedatt--item--rules--definitions"></a>
### Nested Schema for `item.rules.definitions`

Read-Only:

- `type` (String)
- `value` (String)
- `value_list` (Set of String)
- `value_obj` (Attributes) (see [below for nested schema](#nestedatt--item--rules--definitions--value_obj))

<a id="nestedatt--item--rules--definitions--value_obj"></a>
### Nested Schema for `item.rules.definitions.value_obj`

Read-Only:

- `id` (String)
- `name` (String)



<a id="nestedatt--item--rules--per_client_bandwidth_limits"></a>
### Nested Schema for `item.rules.per_client_bandwidth_limits`

Read-Only:

- `bandwidth_limits` (Attributes) (see [below for nested schema](#nestedatt--item--rules--per_client_bandwidth_limits--bandwidth_limits))
- `settings` (String)

<a id="nestedatt--item--rules--per_client_bandwidth_limits--bandwidth_limits"></a>
### Nested Schema for `item.rules.per_client_bandwidth_limits.bandwidth_limits`

Read-Only:

- `limit_down` (Number)
- `limit_up` (Number)
