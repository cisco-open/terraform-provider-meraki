terraform {
  required_providers {
    meraki = {
      source  = "hashicorp.com/edu/meraki"
      version = "1.1.8-beta"
    }
  }

  required_version = ">= 1.2.0"
}

provider "meraki" {
  meraki_debug = "true"
}


resource "meraki_networks_appliance_traffic_shaping_rules" "example" {
  default_rules_enabled = true
  network_id            = "L_828099381482775486"

  rules = [{
    definitions = [
      {
        type  = "host"
        value = "google2.com"
      },
      {
        type = "applicationCategory"
        value_obj = {
          id   = "meraki:layer7/category/8"
          name = "Peer-to-peer (P2P)"
        }
      }
    ]
    per_client_bandwidth_limits = {
      bandwidth_limits = {
        limit_down = 1000000
        limit_up   = 1000000
      }
      settings = "custom"
    }
    priority = "high"
    }, {
    definitions = [
      {
        type  = "host"
        value = "spain2.com"
      },
      {
        type = "applicationCategory"
        value_obj = {
          id   = "meraki:layer7/category/8"
          name = "Peer-to-peer (P2P)"
        }
      }
    ]
    per_client_bandwidth_limits = {
      bandwidth_limits = {
        limit_down = 1000000
        limit_up   = 1000000
      }
      settings = "custom"
    }
    priority = "high"
  }]
}


output "meraki_networks_appliance_traffic_shaping_rules_example" {
  value = meraki_networks_appliance_traffic_shaping_rules.example
}