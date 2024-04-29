
resource "meraki_networks_wireless_ethernet_ports_profiles" "example" {

  name       = "string"
  network_id = "string"
  ports = [{

    enabled      = false
    name         = "string"
    psk_group_id = "string"
    ssid         = 1
  }]
  profile_id = "string"
  usb_ports = [{

    enabled = false
    name    = "string"
    ssid    = 1
  }]
}

output "meraki_networks_wireless_ethernet_ports_profiles_example" {
  value = meraki_networks_wireless_ethernet_ports_profiles.example
}