---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_appliance_traffic_shaping_custom_performance_classes Resource - terraform-provider-meraki"
subcategory: "appliance"
description: |-
  
---

# meraki_networks_appliance_traffic_shaping_custom_performance_classes (Resource)





~>Warning: This resource does not represent a real-world entity in Meraki Dashboard, therefore changing or deleting this resource on its own has no immediate effect. Instead, it is a task part of a Meraki Dashboard workflow. It is executed in Meraki without any additional verification. It does not check if it was executed before or if a similar configuration or action 
already existed previously.


## Example Usage

```terraform
resource "meraki_networks_appliance_traffic_shaping_custom_performance_classes" "example" {

  network_id = "string"
  parameters = {

    max_jitter          = 100
    max_latency         = 100
    max_loss_percentage = 5
    name                = "myCustomPerformanceClass"
  }
}

output "meraki_networks_appliance_traffic_shaping_custom_performance_classes_example" {
  value = meraki_networks_appliance_traffic_shaping_custom_performance_classes.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID
- `parameters` (Attributes) (see [below for nested schema](#nestedatt--parameters))

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `max_jitter` (Number) Maximum jitter in milliseconds
- `max_latency` (Number) Maximum latency in milliseconds
- `max_loss_percentage` (Number) Maximum percentage of packet loss
- `name` (String) Name of the custom performance class


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `custom_performance_class_id` (String) ID of the custom performance class
- `max_jitter` (Number) Maximum jitter in milliseconds
- `max_latency` (Number) Maximum latency in milliseconds
- `max_loss_percentage` (Number) Maximum percentage of packet loss
- `name` (String) Name of the custom performance class
