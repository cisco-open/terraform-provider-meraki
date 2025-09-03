terraform {
  required_providers {
    meraki = {
      version = "1.2.1-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}

variable "my_network_id" {
  type    = string
  default = "L_828099381482775375" # site 3
}
resource "meraki_networks_wireless_settings" "example" {

  ipv6_bridge_enabled        = true
  led_lights_on              = false
  location_analytics_enabled = false
  meshing_enabled            = true
  network_id                 = var.my_network_id
  upgradestrategy            = "minimizeUpgradeTime"
  # named_vlans = {
  #   pool_dhcp_monitoring = {
  #     enabled  = false
  #     duration = 5
  #   }
  # }
}

