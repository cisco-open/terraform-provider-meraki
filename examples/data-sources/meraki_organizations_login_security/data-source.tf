
data "meraki_organizations_login_security" "example" {

  organization_id = "string"
}

output "meraki_organizations_login_security_example" {
  value = data.meraki_organizations_login_security.example.item
}
