
resource "meraki_networks_appliance_ssids" "example" {

  auth_mode       = "8021x-radius"
  default_vlan_id = 1
  dhcp_enforced_deauthentication = {

    enabled = true
  }
  dot11w = {

    enabled  = true
    required = true
  }
  enabled         = true
  encryption_mode = "wpa"
  name            = "My SSID"
  network_id      = "string"
  number          = "string"
  psk             = "psk"
  radius_servers = [{

    host   = "0.0.0.0"
    port   = 1000
    secret = "secret"
  }]
  visible             = true
  wpa_encryption_mode = "WPA2 only"
}

output "meraki_networks_appliance_ssids_example" {
  value = meraki_networks_appliance_ssids.example
}