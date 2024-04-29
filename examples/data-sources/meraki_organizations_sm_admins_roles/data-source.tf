
data "meraki_organizations_sm_admins_roles" "example" {

  ending_before   = "string"
  organization_id = "string"
  per_page        = 1
  starting_after  = "string"
}

output "meraki_organizations_sm_admins_roles_example" {
  value = data.meraki_organizations_sm_admins_roles.example.item
}
