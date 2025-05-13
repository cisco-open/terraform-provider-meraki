terraform {
  required_providers {
    meraki = {
      version = "1.1.3-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = true
}

resource "meraki_networks_wireless_ssids" "example" {

  network_id      = "L_828099381482771185"
  number          = 0
  name            = "tf-test-wifi2"
  auth_mode       = "psk"
  encryption_mode = "wpa"
  psk             = "thisismylab"
  enabled         = true
  # splash_guest_sponsor_domains=[]
}

output "meraki_wifi" {
  value = resource.meraki_networks_wireless_ssids.example
}
