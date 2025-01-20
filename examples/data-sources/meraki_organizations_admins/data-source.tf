
data "meraki_organizations_admins" "example" {

  network_ids     = ["string"]
  organization_id = "string"
}

output "meraki_organizations_admins_example" {
  value = data.meraki_organizations_admins.example.items
}
