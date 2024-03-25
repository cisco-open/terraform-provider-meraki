
resource "meraki_networks_bind" "example" {

  network_id = "string"
  parameters = {

    auto_bind          = false
    config_template_id = "N_23952905"
  }
}

output "meraki_networks_bind_example" {
  value = meraki_networks_bind.example
}