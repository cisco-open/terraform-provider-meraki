
resource "meraki_networks_unbind" "example" {

  network_id = "string"
  parameters = {

    retain_configs = true
  }
}

output "meraki_networks_unbind_example" {
  value = meraki_networks_unbind.example
}