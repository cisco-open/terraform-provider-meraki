
resource "meraki_organizations_policy_objects" "example" {

  category        = "network"
  name            = "Web Servers - Datacenter 10"
  organization_id = "string"
  type            = "cidr"
}

output "meraki_organizations_policy_objects_example" {
  value = meraki_organizations_policy_objects.example
}