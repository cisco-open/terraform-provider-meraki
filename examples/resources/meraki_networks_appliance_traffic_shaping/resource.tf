
resource "meraki_networks_appliance_traffic_shaping" "example" {

  global_bandwidth_limits = {

    limit_down = 5120
    limit_up   = 2048
  }
  network_id = "string"
}

output "meraki_networks_appliance_traffic_shaping_example" {
  value = meraki_networks_appliance_traffic_shaping.example
}