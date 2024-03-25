
resource "meraki_organizations_inventory_release" "example" {

  organization_id = "string"
  parameters = {

    serials = ["Q234-ABCD-5678"]
  }
}

output "meraki_organizations_inventory_release_example" {
  value = meraki_organizations_inventory_release.example
}