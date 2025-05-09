---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_wireless_electronic_shelf_label Resource - terraform-provider-meraki"
subcategory: "wireless"
description: |-
  
---

# meraki_networks_wireless_electronic_shelf_label (Resource)



## Example Usage

```terraform
resource "meraki_networks_wireless_electronic_shelf_label" "example" {

  enabled    = true
  hostname   = "example.com"
  mode       = "high frequency"
  network_id = "string"
}

output "meraki_networks_wireless_electronic_shelf_label_example" {
  value = meraki_networks_wireless_electronic_shelf_label.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Optional

- `enabled` (Boolean) Turn ESL features on and off for this network
- `hostname` (String) Desired ESL hostname of the network
- `mode` (String) Electronic shelf label mode of the network. Valid options are 'Bluetooth', 'high frequency'
                                  Allowed values: [Bluetooth,high frequency]

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_networks_wireless_electronic_shelf_label.example "network_id"
```
