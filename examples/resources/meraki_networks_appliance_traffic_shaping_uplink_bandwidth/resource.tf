
resource "meraki_networks_appliance_traffic_shaping_uplink_bandwidth" "example" {

  bandwidth_limits = {

    cellular = {

      limit_down = 51200
      limit_up   = 51200
    }
    wan1 = {

      limit_down = 1000000
      limit_up   = 1000000
    }
    wan2 = {

      limit_down = 1000000
      limit_up   = 1000000
    }
  }
  network_id = "string"
}

output "meraki_networks_appliance_traffic_shaping_uplink_bandwidth_example" {
  value = meraki_networks_appliance_traffic_shaping_uplink_bandwidth.example
}