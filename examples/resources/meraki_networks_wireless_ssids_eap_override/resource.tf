
resource "meraki_networks_wireless_ssids_eap_override" "example" {

  eapol_key = {

    retries       = 5
    timeout_in_ms = 5000
  }
  identity = {

    retries = 5
    timeout = 5
  }
  max_retries = 5
  network_id  = "string"
  number      = "string"
  timeout     = 5
}

output "meraki_networks_wireless_ssids_eap_override_example" {
  value = meraki_networks_wireless_ssids_eap_override.example
}