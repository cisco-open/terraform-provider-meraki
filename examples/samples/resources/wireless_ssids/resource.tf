terraform {
  required_providers {
    meraki = {
      version = "0.2.2-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_wireless_ssids" "test" {
  network_id                      = "L_828099381482771185"
  number                          = 1
  enabled                         = true
  name                            = "test 2"
  adult_content_filtering_enabled = false
  auth_mode                       = "psk"
  available_on_all_aps            = true
  band_selection                  = "Dual band operation with Band Steering"
  ip_assignment_mode              = "Bridge mode"
  psk                             = "<redacted>"
  encryption_mode                 = "wpa"
  wpa_encryption_mode             = "WPA2 only"
  mandatory_dhcp_enabled          = true
  default_vlan_id                 = "10"
  lan_isolation_enabled           = false
}