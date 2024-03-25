
resource "meraki_networks_devices_remove" "example" {

  network_id = "string"
  parameters = {

    serial = "Q234-ABCD-5678"
  }
}

output "meraki_networks_devices_remove_example" {
  value = meraki_networks_devices_remove.example
}