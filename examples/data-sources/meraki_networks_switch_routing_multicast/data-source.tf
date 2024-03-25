
data "meraki_networks_switch_routing_multicast" "example" {

  network_id = "string"
}

output "meraki_networks_switch_routing_multicast_example" {
  value = data.meraki_networks_switch_routing_multicast.example.item
}
