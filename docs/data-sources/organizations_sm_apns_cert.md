---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_sm_apns_cert Data Source - terraform-provider-meraki"
subcategory: "sm"
description: |-
  
---

# meraki_organizations_sm_apns_cert (Data Source)



## Example Usage

```terraform
data "meraki_organizations_sm_apns_cert" "example" {

  organization_id = "string"
}

output "meraki_organizations_sm_apns_cert_example" {
  value = data.meraki_organizations_sm_apns_cert.example.item
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

- `certificate` (String) Organization APNS Certificate used by devices to communication with Apple
