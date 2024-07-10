
resource "meraki_networks_wireless_ssids_identity_psks" "example" {

  expires_at      = "2018-02-11T00:00:00.090209Z"
  group_policy_id = "101"
  name            = "Sample Identity PSK"
  network_id      = "string"
  number          = "string"
  passphrase      = "secret"
}

output "meraki_networks_wireless_ssids_identity_psks_example" {
  value = meraki_networks_wireless_ssids_identity_psks.example
}