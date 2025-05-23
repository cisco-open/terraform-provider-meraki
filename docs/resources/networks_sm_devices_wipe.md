---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_sm_devices_wipe Resource - terraform-provider-meraki"
subcategory: "sm"
description: |-
  
---

# meraki_networks_sm_devices_wipe (Resource)





~>Warning: This resource does not represent a real-world entity in Meraki Dashboard, therefore changing or deleting this resource on its own has no immediate effect. Instead, it is a task part of a Meraki Dashboard workflow. It is executed in Meraki without any additional verification. It does not check if it was executed before or if a similar configuration or action 
already existed previously.


## Example Usage

```terraform
resource "meraki_networks_sm_devices_wipe" "example" {

  network_id = "string"
  parameters = {

    id       = "1284392014819"
    pin      = 123456
    serial   = "XY0XX0Y0X0"
    wifi_mac = "00:11:22:33:44:55"
  }
}

output "meraki_networks_sm_devices_wipe_example" {
  value = meraki_networks_sm_devices_wipe.example
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

- `id` (String) The id of the device to be wiped.
- `pin` (Number) The pin number (a six digit value) for wiping a macOS device. Required only for macOS devices.
- `serial` (String) The serial of the device to be wiped.
- `wifi_mac` (String) The wifiMac of the device to be wiped.


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `id` (String) The Meraki Id of the devices.
