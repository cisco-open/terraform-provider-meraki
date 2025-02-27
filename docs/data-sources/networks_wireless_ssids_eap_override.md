---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_wireless_ssids_eap_override Data Source - terraform-provider-meraki"
subcategory: "wireless"
description: |-
  
---

# meraki_networks_wireless_ssids_eap_override (Data Source)



## Example Usage

```terraform
data "meraki_networks_wireless_ssids_eap_override" "example" {

  network_id = "string"
  number     = "string"
}

output "meraki_networks_wireless_ssids_eap_override_example" {
  value = data.meraki_networks_wireless_ssids_eap_override.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID
- `number` (String) number path parameter.

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `eapol_key` (Attributes) EAPOL Key settings. (see [below for nested schema](#nestedatt--item--eapol_key))
- `identity` (Attributes) EAP settings for identity requests. (see [below for nested schema](#nestedatt--item--identity))
- `max_retries` (Number) Maximum number of general EAP retries.
- `timeout` (Number) General EAP timeout in seconds.

<a id="nestedatt--item--eapol_key"></a>
### Nested Schema for `item.eapol_key`

Read-Only:

- `retries` (Number) Maximum number of EAPOL key retries.
- `timeout_in_ms` (Number) EAPOL Key timeout in milliseconds.


<a id="nestedatt--item--identity"></a>
### Nested Schema for `item.identity`

Read-Only:

- `retries` (Number) Maximum number of EAP retries.
- `timeout` (Number) EAP timeout in seconds.
