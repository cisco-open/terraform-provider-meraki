terraform {
  required_providers {
    meraki = {
      version = "1.0.6-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_appliance_firewall_l7_firewall_rules" "my_mx" {

  network_id = "L_828099381482771185"
  rules = [{

    policy = "deny"
    type   = "applicationCategory"
    # value                   = "10.11.12.00/24"
    value_obj = {
      name = "Sports 2"
      id   = "meraki:layer7/category/5"
    }
    }
  ]
}