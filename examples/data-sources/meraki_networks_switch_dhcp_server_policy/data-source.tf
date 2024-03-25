
data "meraki_networks_switch_dhcp_server_policy" "example" {

  network_id = "string"
}

output "meraki_networks_switch_dhcp_server_policy_example" {
  value = data.meraki_networks_switch_dhcp_server_policy.example.item
}
