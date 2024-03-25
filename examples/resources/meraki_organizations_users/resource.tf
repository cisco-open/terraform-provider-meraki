
resource "meraki_organizations_users" "example" {

  organization_id = "string"
  user_id         = "string"
}

output "meraki_organizations_users_example" {
  value = meraki_organizations_users.example
}