terraform {
  required_providers {
    meraki = {
      version = "0.2.6-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_wireless_ssids" "this_ssid" {
  network_id              = "L_828099381482771185"
  number                  = 0
  name                    = "SSID Test Terraform"
  enabled                 = true
  splash_page             = "None"
  auth_mode               = "8021x-radius"
  dot11w = {
    enabled  = false
    required = false
  }
  dot11r = {
    enabled  = true
    adaptive = false
  }
  wpa_encryption_mode = "WPA2 only"
  radius_servers = [{
    host           = "192.168.10.2"
    port           = 1812
    radsec_enabled = false
    secret         = "<-blanked->"
  }]
  radius_accounting_enabled = true
  radius_accounting_servers = [{
    host           = "192.168.10.3"
    port           = 1813
    radsec_enabled = false
    secret         = "<-blanked->"
  }]
  radius_testing_enabled             = false
  radius_server_timeout              = 1
  radius_server_attempts_limit       = 3
  radius_fallback_enabled            = false
  radius_accounting_interim_interval = 60
  radius_proxy_enabled               = false
  radius_coa_enabled                 = false
  radius_called_station_id           = "$NODE_MAC$:$VAP_NAME$"
  radius_authentication_nas_id       = "$NODE_MAC$:$VAP_NUM$"
  ip_assignment_mode                 = "Bridge mode"
  use_vlan_tagging                   = false
  radius_override                    = true
  min_bitrate                        = 11
  band_selection                     = "Dual band operation"
  per_client_bandwidth_limit_down    = 0
  per_client_bandwidth_limit_up      = 0
  per_ssid_bandwidth_limit_down      = 0
  per_ssid_bandwidth_limit_up        = 0
  mandatory_dhcp_enabled             = false
  lan_isolation_enabled              = true
  visible                            = true
  available_on_all_aps               = true
  availability_tags                  = []
  speed_burst = {
    enabled = false
  }
}