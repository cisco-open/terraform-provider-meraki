---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_policy_objects Resource - terraform-provider-meraki"
subcategory: "organizations"
description: |-
  
---

# meraki_organizations_policy_objects (Resource)



## Example Usage

```terraform
resource "meraki_organizations_policy_objects" "example" {

  category        = "network"
  cidr            = "10.0.0.0/24"
  fqdn            = "example.com"
  group_ids       = ["8"]
  ip              = "1.2.3.4"
  mask            = "255.255.0.0"
  name            = "Web Servers - Datacenter 10"
  organization_id = "string"
  type            = "cidr"
}

output "meraki_organizations_policy_objects_example" {
  value = meraki_organizations_policy_objects.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Optional

- `category` (String) Category of a policy object (one of: adaptivePolicy, network)
- `cidr` (String) CIDR Value of a policy object
- `fqdn` (String) Fully qualified domain name of policy object (e.g. "example.com")
- `group_ids` (Set of String) The IDs of policy object groups the policy object belongs to.
- `ip` (String) IP Address of a policy object (e.g. "1.2.3.4")
- `mask` (String) Mask of a policy object (e.g. "255.255.0.0")
- `name` (String) Name of policy object (alphanumeric, space, dash, or underscore characters only).
- `policy_object_id` (String) policyObjectId path parameter. Policy object ID
- `type` (String) Type of a policy object (one of: adaptivePolicyIpv4Cidr, cidr, fqdn, ipAndMask)

### Read-Only

- `created_at` (String) Time Stamp of policy object creation.
- `id` (String) Policy object ID
- `network_ids` (Set of String) The IDs of the networks that use the policy object.
- `updated_at` (String) Time Stamp of policy object updation.

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_organizations_policy_objects.example "organization_id,policy_object_id"
```
