
resource "meraki_networks_wireless_ssids_traffic_shaping_rules" "example" {

  default_rules_enabled = true
  network_id            = "string"
  number                = "string"
  rules = [{

    definitions = [{

      type  = "host"
      value = "google.com"
    }]
    dscp_tag_value = 1
    pcp_tag_value  = 1
    per_client_bandwidth_limits = {

      bandwidth_limits = {

        limit_down = 1000000
        limit_up   = 1000000
      }
      settings = "custom"
    }
  }]
  traffic_shaping_enabled = true
}

output "meraki_networks_wireless_ssids_traffic_shaping_rules_example" {
  value = meraki_networks_wireless_ssids_traffic_shaping_rules.example
}