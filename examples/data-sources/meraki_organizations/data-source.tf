
data "meraki_organizations" "example" {

  ending_before  = "string"
  per_page       = 1
  starting_after = "string"
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
