terraform {
  required_providers {
    meraki = {
      version = "0.2.13-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}


resource "meraki_networks_wireless_ssids_firewall_l7_firewall_rules" "example" {

  network_id = "L_828099381482771185"
  number     = "1"
  rules = [{

    policy = "deny"
    type   = "host"
    value  = "google.com"
    },
    {

      policy = "deny"
      type   = "applicationCategory"
      value_obj = {
        id   = "meraki:layer7/category/8"
        name = "Peer-to-peer (P2P)"
      }
  }]
}

output "meraki_networks_wireless_ssids_firewall_l7_firewall_rules_example" {
  value = meraki_networks_wireless_ssids_firewall_l7_firewall_rules.example
}