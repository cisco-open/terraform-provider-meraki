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
  meraki_debug = "true"
}
resource "meraki_networks_sm_devices_refresh_details" "example" {
  provider = meraki


  device_id  = "QBSC-ALSL-3GXN"
  network_id = "L_828099381482775374"

}