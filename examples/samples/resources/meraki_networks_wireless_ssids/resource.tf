terraform {
  required_providers {
    meraki = {
      version = "0.1.0-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-open/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug=true
}

resource "meraki_networks_wireless_ssids" "example" {
    provider = meraki
    network_id = "L_3695766444210915322"
    number = 0
    name = "tf-test-wifi2"
    auth_mode = "psk"
    encryption_mode = "wpa"
    psk = "thisismylab"
    enabled = true
    # splash_guest_sponsor_domains=[]
}

output "meraki_wifi" {
  value = resource.meraki_networks_wireless_ssids.example
}
