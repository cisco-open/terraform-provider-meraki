
resource "meraki_networks_cellular_gateway_uplink" "example" {

  bandwidth_limits = {

    limit_down = 1000000
    limit_up   = 1000000
  }
  network_id = "string"
}

output "meraki_networks_cellular_gateway_uplink_example" {
  value = meraki_networks_cellular_gateway_uplink.example
}