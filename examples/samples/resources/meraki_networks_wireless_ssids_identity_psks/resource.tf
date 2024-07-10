terraform {
  required_providers {
    meraki = {
      version = "0.2.5-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}


resource "meraki_networks_wireless_ssids_identity_psks" "example" {
# lifecycle {
#   # ignore_changes = [ expires_at ]
# }
  expires_at      = "2018-02-11T00:00:00.090209Z"
  group_policy_id = "100"
  name            = "Sample Identity PSK"
  network_id      = "L_828099381482775375"
  number          = "12"
  passphrase      = "secret1221121"
}

output "meraki_networks_wireless_ssids_identity_psks_example" {
  value = meraki_networks_wireless_ssids_identity_psks.example
}