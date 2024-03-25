
resource "meraki_organizations_switch_devices_clone" "example" {

  organization_id = "string"
  parameters = {

    source_serial  = "Q234-ABCD-5678"
    target_serials = ["Q234-ABCD-0001", "Q234-ABCD-0002", "Q234-ABCD-0003"]
  }
}

output "meraki_organizations_switch_devices_clone_example" {
  value = meraki_organizations_switch_devices_clone.example
}