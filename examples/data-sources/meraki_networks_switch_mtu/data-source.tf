
data "meraki_networks_switch_mtu" "example" {

  network_id = "string"
}

output "meraki_networks_switch_mtu_example" {
  value = data.meraki_networks_switch_mtu.example.item
}
