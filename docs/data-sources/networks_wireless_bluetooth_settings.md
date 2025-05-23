---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_wireless_bluetooth_settings Data Source - terraform-provider-meraki"
subcategory: "wireless"
description: |-
  
---

# meraki_networks_wireless_bluetooth_settings (Data Source)



## Example Usage

```terraform
data "meraki_networks_wireless_bluetooth_settings" "example" {

  network_id = "string"
}

output "meraki_networks_wireless_bluetooth_settings_example" {
  value = data.meraki_networks_wireless_bluetooth_settings.example.item
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

- `advertising_enabled` (Boolean) Whether APs will advertise beacons.
- `esl_enabled` (Boolean) Whether ESL is enabled on this network.
- `major` (Number) The major number to be used in the beacon identifier. Only valid in 'Non-unique' mode.
- `major_minor_assignment_mode` (String) The way major and minor number should be assigned to nodes in the network. ('Unique', 'Non-unique')
- `minor` (Number) The minor number to be used in the beacon identifier. Only valid in 'Non-unique' mode.
- `scanning_enabled` (Boolean) Whether APs will scan for Bluetooth enabled clients.
- `uuid` (String) The UUID to be used in the beacon identifier.
