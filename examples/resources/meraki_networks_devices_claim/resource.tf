
resource "meraki_networks_devices_claim" "example" {

  add_atomically = false
  network_id     = "string"
  parameters = {

    serials = ["Q234-ABCD-0001", "Q234-ABCD-0002", "Q234-ABCD-0003"]
  }
}

output "meraki_networks_devices_claim_example" {
  value = meraki_networks_devices_claim.example
}