---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_appliance_vlans_settings Data Source - terraform-provider-meraki"
subcategory: "appliance"
description: |-
  
---

# meraki_networks_appliance_vlans_settings (Data Source)



## Example Usage

```terraform
data "meraki_networks_appliance_vlans_settings" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_vlans_settings_example" {
  value = data.meraki_networks_appliance_vlans_settings.example.item
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

- `vlans_enabled` (Boolean) Boolean indicating whether VLANs are enabled (true) or disabled (false) for the network
