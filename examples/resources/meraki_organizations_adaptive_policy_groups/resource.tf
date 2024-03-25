
resource "meraki_organizations_adaptive_policy_groups" "example" {

  description     = "Group of XYZ Corp Employees"
  name            = "Employee Group"
  organization_id = "string"
  policy_objects = [{

    id   = "2345"
    name = "Example Policy Object"
  }]
  sgt = 1000
}

output "meraki_organizations_adaptive_policy_groups_example" {
  value = meraki_organizations_adaptive_policy_groups.example
}