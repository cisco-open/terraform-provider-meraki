
data "meraki_networks_switch_routing_multicast_rendezvous_points" "example" {

  network_id = "string"
}

output "meraki_networks_switch_routing_multicast_rendezvous_points_example" {
  value = data.meraki_networks_switch_routing_multicast_rendezvous_points.example.items
}

data "meraki_networks_switch_routing_multicast_rendezvous_points" "example" {

  network_id = "string"
}

output "meraki_networks_switch_routing_multicast_rendezvous_points_example" {
  value = data.meraki_networks_switch_routing_multicast_rendezvous_points.example.item
}
