
resource "meraki_organizations_policy_objects_groups" "example" {

  name            = "Web Servers - Datacenter 10"
  organization_id = "string"
}

output "meraki_organizations_policy_objects_groups_example" {
  value = meraki_organizations_policy_objects_groups.example
}