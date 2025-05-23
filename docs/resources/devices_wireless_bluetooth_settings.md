---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_devices_wireless_bluetooth_settings Resource - terraform-provider-meraki"
subcategory: "wireless"
description: |-
  
---

# meraki_devices_wireless_bluetooth_settings (Resource)



## Example Usage

```terraform
resource "meraki_devices_wireless_bluetooth_settings" "example" {

  major  = 13
  minor  = 125
  serial = "string"
  uuid   = "00000000-0000-0000-000-000000000000"
}

output "meraki_devices_wireless_bluetooth_settings_example" {
  value = meraki_devices_wireless_bluetooth_settings.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `serial` (String) serial path parameter.

### Optional

- `major` (Number) Desired major value of the beacon. If the value is set to null it will reset to
          Dashboard's automatically generated value.
- `minor` (Number) Desired minor value of the beacon. If the value is set to null it will reset to
          Dashboard's automatically generated value.
- `uuid` (String) Desired UUID of the beacon. If the value is set to null it will reset to Dashboard's
          automatically generated value.

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_devices_wireless_bluetooth_settings.example "serial"
```
