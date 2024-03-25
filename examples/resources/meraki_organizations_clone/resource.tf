
resource "meraki_organizations_clone" "example" {

  organization_id = "string"
  parameters = {

    name = "My organization"
  }
}

output "meraki_organizations_clone_example" {
  value = meraki_organizations_clone.example
}