
data "meraki_organizations" "example" {

}

output "meraki_organizations_example" {
  value = data.meraki_organizations.example.items
}

data "meraki_organizations" "example" {

  organization_id = "string"
}

output "meraki_organizations_example" {
  value = data.meraki_organizations.example.item
}
