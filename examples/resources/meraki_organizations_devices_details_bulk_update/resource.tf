
resource "meraki_organizations_devices_details_bulk_update" "example" {

  organization_id = "string"
  parameters = {

    details = [{

      name  = "username"
      value = "ABC"
    }]
    serials = ["Q234-ABCD-0001", "Q234-ABCD-0002", "Q234-ABCD-0003"]
  }
}

output "meraki_organizations_devices_details_bulk_update_example" {
  value = meraki_organizations_devices_details_bulk_update.example
}