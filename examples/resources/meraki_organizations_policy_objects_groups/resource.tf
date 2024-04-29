
resource "meraki_organizations_policy_objects_groups" "example" {

  category        = "NetworkObjectGroup"
  name            = "Web Servers - Datacenter 10"
  object_ids      = [100]
  organization_id = "string"
}

output "meraki_organizations_policy_objects_groups_example" {
  value = meraki_organizations_policy_objects_groups.example
}