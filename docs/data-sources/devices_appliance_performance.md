---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_devices_appliance_performance Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_devices_appliance_performance (Data Source)



## Example Usage

```terraform
data "meraki_devices_appliance_performance" "example" {

  serial = "string"
}

output "meraki_devices_appliance_performance_example" {
  value = data.meraki_devices_appliance_performance.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `serial` (String) serial path parameter.

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `perf_score` (Number)
