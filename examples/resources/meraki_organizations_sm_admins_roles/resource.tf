
resource "meraki_organizations_sm_admins_roles" "example" {

  name            = "sample name"
  organization_id = "string"
  scope           = "all_tags"
  tags            = ["tag"]
}

output "meraki_organizations_sm_admins_roles_example" {
  value = meraki_organizations_sm_admins_roles.example
}