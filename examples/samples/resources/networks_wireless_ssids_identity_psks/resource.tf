terraform {
  required_providers {
    meraki = {
      version = "0.2.13-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

resource "meraki_networks_wireless_ssids_identity_psks" "example" {

  expires_at      = "2018-02-11T00:00:00.090210Z"
  group_policy_id = "100"
  # identity_psk_id = "1284392014819"
  name       = "Sample Identity PSK"
  network_id = "L_828099381482771185"
  number     = "1"
  passphrase = "secret"
}

output "meraki_networks_wireless_ssids_identity_psks_example" {
  value = meraki_networks_wireless_ssids_identity_psks.example
}