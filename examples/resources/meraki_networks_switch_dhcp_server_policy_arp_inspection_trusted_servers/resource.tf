
resource "meraki_networks_switch_dhcp_server_policy_arp_inspection_trusted_servers" "example" {

  ipv4 = {

    address = "1.2.3.4"
  }
  mac        = "00:11:22:33:44:55"
  network_id = "string"
  vlan       = 100
}

output "meraki_networks_switch_dhcp_server_policy_arp_inspection_trusted_servers_example" {
  value = meraki_networks_switch_dhcp_server_policy_arp_inspection_trusted_servers.example
}