terraform {
  required_providers {
    meraki = {
      version = "1.1.6-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}


data "meraki_networks_wireless_rf_profiles" "example" {
  network_id    = "L_828099381482775342"
  rf_profile_id = "outdoor"
}

output "meraki_networks_wireless_rf_profiles_example" {
  value = data.meraki_networks_wireless_rf_profiles.example.item
}