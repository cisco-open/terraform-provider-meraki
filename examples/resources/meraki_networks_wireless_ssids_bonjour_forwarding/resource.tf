
resource "meraki_networks_wireless_ssids_bonjour_forwarding" "example" {

  enabled = true
  exception = {

    enabled = true
  }
  network_id = "string"
  number     = "string"
  rules = [{

    description = "A simple bonjour rule"
    services    = ["All Services"]
    vlan_id     = "1"
  }]
}

output "meraki_networks_wireless_ssids_bonjour_forwarding_example" {
  value = meraki_networks_wireless_ssids_bonjour_forwarding.example
}