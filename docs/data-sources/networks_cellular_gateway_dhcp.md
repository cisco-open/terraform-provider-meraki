---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_cellular_gateway_dhcp Data Source - terraform-provider-meraki"
subcategory: "cellularGateway"
description: |-
  
---

# meraki_networks_cellular_gateway_dhcp (Data Source)



## Example Usage

```terraform
data "meraki_networks_cellular_gateway_dhcp" "example" {

  network_id = "string"
}

output "meraki_networks_cellular_gateway_dhcp_example" {
  value = data.meraki_networks_cellular_gateway_dhcp.example.item
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

- `dhcp_lease_time` (String) DHCP Lease time for all MG in the network.
- `dns_custom_nameservers` (List of String) List of fixed IPs representing the the DNS Name servers when the mode is 'custom'.
- `dns_nameservers` (String) DNS name servers mode for all MG in the network.
