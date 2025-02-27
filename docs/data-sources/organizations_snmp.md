---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_snmp Data Source - terraform-provider-meraki"
subcategory: "organizations"
description: |-
  
---

# meraki_organizations_snmp (Data Source)



## Example Usage

```terraform
data "meraki_organizations_snmp" "example" {

  organization_id = "string"
}

output "meraki_organizations_snmp_example" {
  value = data.meraki_organizations_snmp.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `hostname` (String) The hostname of the SNMP server.
- `peer_ips` (List of String) The list of IPv4 addresses that are allowed to access the SNMP server.
- `port` (Number) The port of the SNMP server.
- `v2_community_string` (String) The community string for SNMP version 2c, if enabled.
- `v2c_enabled` (Boolean) Boolean indicating whether SNMP version 2c is enabled for the organization.
- `v3_auth_mode` (String) The SNMP version 3 authentication mode. Can be either 'MD5' or 'SHA'.
- `v3_enabled` (Boolean) Boolean indicating whether SNMP version 3 is enabled for the organization.
- `v3_priv_mode` (String) The SNMP version 3 privacy mode. Can be either 'DES' or 'AES128'.
- `v3_user` (String) The user for SNMP version 3, if enabled.
