---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_saml Data Source - terraform-provider-meraki"
subcategory: "organizations"
description: |-
  
---

# meraki_organizations_saml (Data Source)



## Example Usage

```terraform
data "meraki_organizations_saml" "example" {

  organization_id = "string"
}

output "meraki_organizations_saml_example" {
  value = data.meraki_organizations_saml.example.item
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

- `enabled` (Boolean) Toggle depicting if SAML SSO settings are enabled
