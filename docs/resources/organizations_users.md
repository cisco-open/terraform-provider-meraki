---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_users Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_organizations_users (Resource)



~>Warning: This resource does not represent a real-world entity in Meraki Dashboard, therefore changing or deleting this resource on its own has no immediate effect. Instead, it is a task part of a Meraki Dashboard workflow. It is executed in Meraki without any additional verification. It does not check if it was executed before or if a similar configuration or action 
already existed previously.

## Example Usage

```terraform
resource "meraki_organizations_users" "example" {

  organization_id = "string"
  user_id         = "string"
}

output "meraki_organizations_users_example" {
  value = meraki_organizations_users.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID
- `user_id` (String) userId path parameter. User ID
