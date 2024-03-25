
resource "meraki_networks_appliance_traffic_shaping_uplink_selection" "example" {

  active_active_auto_vpn_enabled = true
  default_uplink                 = "wan1"
  failover_and_failback = {

    immediate = {

      enabled = true
    }
  }
  load_balancing_enabled = true
  network_id             = "string"
  vpn_traffic_uplink_preferences = [{

    fail_over_criterion = "poorPerformance"
    performance_class = {

      builtin_performance_class_name = "VoIP"
      custom_performance_class_id    = "123456"
      type                           = "custom"
    }
    preferred_uplink = "bestForVoIP"
    traffic_filters = [{

      type = "applicationCategory"
      value = {

        destination = {

          cidr    = "any"
          fqdn    = "www.google.com"
          host    = 254
          network = "L_12345678"
          port    = "1-1024"
          vlan    = 10
        }
        id       = "meraki:layer7/category/1"
        protocol = "tcp"
        source = {

          cidr    = "192.168.1.0/24"
          host    = 200
          network = "L_23456789"
          port    = "any"
          vlan    = 20
        }
      }
    }]
  }]
  wan_traffic_uplink_preferences = [{

    preferred_uplink = "wan1"
    traffic_filters = [{

      type = "custom"
      value = {

        destination = {

          cidr = "any"
          port = "any"
        }
        protocol = "tcp"
        source = {

          cidr = "192.168.1.0/24"
          host = 254
          port = "1-1024"
          vlan = 10
        }
      }
    }]
  }]
}

output "meraki_networks_appliance_traffic_shaping_uplink_selection_example" {
  value = meraki_networks_appliance_traffic_shaping_uplink_selection.example
}