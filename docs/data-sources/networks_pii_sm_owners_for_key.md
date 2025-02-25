---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_pii_sm_owners_for_key Data Source - terraform-provider-meraki"
subcategory: "networks"
description: |-
  
---

# meraki_networks_pii_sm_owners_for_key (Data Source)



## Example Usage

```terraform
data "meraki_networks_pii_sm_owners_for_key" "example" {

  bluetooth_mac = "string"
  email         = "string"
  imei          = "string"
  mac           = "string"
  network_id    = "string"
  serial        = "string"
  username      = "string"
}

output "meraki_networks_pii_sm_owners_for_key_example" {
  value = data.meraki_networks_pii_sm_owners_for_key.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Optional

- `bluetooth_mac` (String) bluetoothMac query parameter. The MAC of a Bluetooth client
- `email` (String) email query parameter. The email of a network user account or a Systems Manager device
- `imei` (String) imei query parameter. The IMEI of a Systems Manager device
- `mac` (String) mac query parameter. The MAC of a network client device or a Systems Manager device
- `serial` (String) serial query parameter. The serial of a Systems Manager device
- `username` (String) username query parameter. The username of a Systems Manager user

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `n_1234` (List of String)
