---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_admins Resource - terraform-provider-meraki"
subcategory: "organizations"
description: |-
  
---

# meraki_organizations_admins (Resource)



## Example Usage

```terraform
resource "meraki_organizations_admins" "example" {

  authentication_method = "Email"
  email                 = "miles@meraki.com"
  name                  = "Miles Meraki"
  networks = [{

    access = "full"
    id     = "N_24329156"
  }]
  org_access      = "none"
  organization_id = "string"
  tags = [{

    access = "read-only"
    tag    = "west"
  }]
}

output "meraki_organizations_admins_example" {
  value = meraki_organizations_admins.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Optional

- `admin_id` (String) adminId path parameter. Admin ID
- `authentication_method` (String) Admin's authentication method
                                  Allowed values: [Cisco SecureX Sign-On,Email]
- `email` (String) Admin's email address
- `name` (String) Admin's username
- `networks` (Attributes Set) Admin network access information (see [below for nested schema](#nestedatt--networks))
- `org_access` (String) Admin's level of access to the organization
                                  Allowed values: [enterprise,full,none,read-only]
- `tags` (Attributes Set) Admin tag information (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `account_status` (String) Status of the admin's account
                                  Allowed values: [locked,ok,pending,unverified]
- `has_api_key` (Boolean) Indicates whether the admin has an API key
- `id` (String) Admin's ID
- `last_active` (String) Time when the admin was last active
- `two_factor_auth_enabled` (Boolean) Indicates whether two-factor authentication is enabled

<a id="nestedatt--networks"></a>
### Nested Schema for `networks`

Optional:

- `access` (String) Admin's level of access to the network
- `id` (String) Network ID


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `access` (String) Access level for the tag
                                        Allowed values: [full,guest-ambassador,monitor-only,read-only]
- `tag` (String) Tag value

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_organizations_admins.example "organization_id"
```
