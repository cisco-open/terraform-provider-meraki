
resource "meraki_networks_devices_claim_vmx" "example" {

  network_id = "string"
  parameters = {

    size = "small"
  }
}

output "meraki_networks_devices_claim_vmx_example" {
  value = meraki_networks_devices_claim_vmx.example
}