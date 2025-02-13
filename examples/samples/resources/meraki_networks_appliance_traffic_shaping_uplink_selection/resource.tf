terraform {
  required_providers {
    meraki = {
      source  = "hashicorp.com/edu/meraki"
      version = "1.0.2-beta"
    }
  }

  required_version = ">= 1.2.0"
}

provider "meraki" {
  meraki_debug = "true"
}


# data "meraki_networks" "my_networks" {
#   provider        = meraki
#   organization_id = "828099381482762270"

# }

resource "meraki_networks_appliance_traffic_shaping_uplink_selection" "this" {
  network_id                     = "L_828099381482775486"
  default_uplink                 = "wan1"
  active_active_auto_vpn_enabled = true
  load_balancing_enabled         = true
  failover_and_failback = {
    immediate = {
      enabled = false
    }
  }
}


output "meraki_networks_appliance_traffic_shaping_uplink_selection_example" {
  value = meraki_networks_appliance_traffic_shaping_uplink_selection.this
}