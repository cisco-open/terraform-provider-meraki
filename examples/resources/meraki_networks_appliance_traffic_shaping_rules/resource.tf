
resource "meraki_networks_appliance_traffic_shaping_rules" "example" {

  default_rules_enabled = true
  network_id            = "string"
  rules = [{

    definitions = [{

      type  = "host"
      value = "google.com"
    }]
    dscp_tag_value = 1
    per_client_bandwidth_limits = {

      bandwidth_limits = {

        limit_down = 1000000
        limit_up   = 1000000
      }
      settings = "custom"
    }
    priority = "normal"
  }]
}

output "meraki_networks_appliance_traffic_shaping_rules_example" {
  value = meraki_networks_appliance_traffic_shaping_rules.example
}