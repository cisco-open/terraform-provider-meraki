---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_wireless_air_marshal_settings Resource - terraform-provider-meraki"
subcategory: "wireless"
description: |-
  
---

# meraki_networks_wireless_air_marshal_settings (Resource)





~>Warning: This resource does not represent a real-world entity in Meraki Dashboard, therefore changing or deleting this resource on its own has no immediate effect. Instead, it is a task part of a Meraki Dashboard workflow. It is executed in Meraki without any additional verification. It does not check if it was executed before or if a similar configuration or action 
already existed previously.


## Example Usage

```terraform
resource "meraki_networks_wireless_air_marshal_settings" "example" {

  network_id = "string"
  parameters = {

    default_policy = "allow"
  }
}

output "meraki_networks_wireless_air_marshal_settings_example" {
  value = meraki_networks_wireless_air_marshal_settings.example
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

- `default_policy` (String) Allows clients to access rogue networks. Blocked by default.
                                        Allowed values: [allow,block]


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `default_policy` (String) Indicates whether or not clients are allowed to       connect to rogue SSIDs. (blocked by default)
- `network_id` (String) The network ID
