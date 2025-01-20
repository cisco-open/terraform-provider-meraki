terraform {
  required_providers {
    meraki = {
      version = "0.2.13-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

resource "meraki_networks_appliance_sdwan_internet_policies" "example" {

  network_id = "L_828099381482771185"
  parameters = {

    wan_traffic_uplink_preferences = [{

      fail_over_criterion = "poorPerformance"
      performance_class = {

        builtin_performance_class_name = "VoIP"
        custom_performance_class_id    = "123456"
        type                           = "custom"
      }
      preferred_uplink = "wan1"
      traffic_filters = [{

        type = "custom"
        value = {

          destination = {

            applications = [{

              id   = "meraki:layer7/application/3"
              name = "DNS"
              type = "major"
            }]
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
}

output "meraki_networks_appliance_sdwan_internet_policies_example" {
  value = meraki_networks_appliance_sdwan_internet_policies.example
}