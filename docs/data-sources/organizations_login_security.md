---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_login_security Data Source - terraform-provider-meraki"
subcategory: "organizations"
description: |-
  
---

# meraki_organizations_login_security (Data Source)



## Example Usage

```terraform
data "meraki_organizations_login_security" "example" {

  organization_id = "string"
}

output "meraki_organizations_login_security_example" {
  value = data.meraki_organizations_login_security.example.item
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

- `account_lockout_attempts` (Number) Number of consecutive failed login attempts after which users' accounts will be locked.
- `api_authentication` (Attributes) Details for indicating whether organization will restrict access to API (but not Dashboard) to certain IP addresses. (see [below for nested schema](#nestedatt--item--api_authentication))
- `enforce_account_lockout` (Boolean) Boolean indicating whether users' Dashboard accounts will be locked out after a specified number of consecutive failed login attempts.
- `enforce_different_passwords` (Boolean) Boolean indicating whether users, when setting a new password, are forced to choose a new password that is different from any past passwords.
- `enforce_idle_timeout` (Boolean) Boolean indicating whether users will be logged out after being idle for the specified number of minutes.
- `enforce_login_ip_ranges` (Boolean) Boolean indicating whether organization will restrict access to Dashboard (including the API) from certain IP addresses.
- `enforce_password_expiration` (Boolean) Boolean indicating whether users are forced to change their password every X number of days.
- `enforce_strong_passwords` (Boolean) Deprecated. This will always be 'true'.
- `enforce_two_factor_auth` (Boolean) Boolean indicating whether users in this organization will be required to use an extra verification code when logging in to Dashboard. This code will be sent to their mobile phone via SMS, or can be generated by the authenticator application.
- `idle_timeout_minutes` (Number) Number of minutes users can remain idle before being logged out of their accounts.
- `login_ip_ranges` (List of String) List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.
- `minimum_password_length` (Number) The minimum number of characters required in admins' passwords.
- `num_different_passwords` (Number) Number of recent passwords that new password must be distinct from.
- `password_expiration_days` (Number) Number of days after which users will be forced to change their password.

<a id="nestedatt--item--api_authentication"></a>
### Nested Schema for `item.api_authentication`

Read-Only:

- `ip_restrictions_for_keys` (Attributes) Details for API-only IP restrictions. (see [below for nested schema](#nestedatt--item--api_authentication--ip_restrictions_for_keys))

<a id="nestedatt--item--api_authentication--ip_restrictions_for_keys"></a>
### Nested Schema for `item.api_authentication.ip_restrictions_for_keys`

Read-Only:

- `enabled` (Boolean) Boolean indicating whether the organization will restrict API key (not Dashboard GUI) usage to a specific list of IP addresses or CIDR ranges.
- `ranges` (List of String) List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.
