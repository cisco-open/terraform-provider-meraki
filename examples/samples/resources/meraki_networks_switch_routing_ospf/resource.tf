
terraform {
  required_providers {
    meraki = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}



resource "meraki_networks_switch_routing_ospf" "example" {

  areas = [{

    area_id   = "6"
    area_name = "Backbone"
    area_type = "normal"
  }]
  dead_timer_in_seconds      = 40
  enabled                    = true
  hello_timer_in_seconds     = 10
  md5_authentication_enabled = true
  md5_authentication_key = {

    id         = 1
    passphrase = "abc1234"
  }
  network_id = "L_828099381482775374"
  v3 = {

    # areas                  = []
    dead_timer_in_seconds  = 40
    enabled                = true
    hello_timer_in_seconds = 10
  }
}

output "meraki_networks_switch_routing_ospf_example" {
  value = meraki_networks_switch_routing_ospf.example
}