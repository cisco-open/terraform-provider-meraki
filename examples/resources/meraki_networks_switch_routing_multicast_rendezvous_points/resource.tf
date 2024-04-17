
resource "meraki_networks_switch_routing_multicast_rendezvous_points" "example" {

  interface_ip    = "192.168.1.2"
  multicast_group = "Any"
  network_id      = "string"
}

output "meraki_networks_switch_routing_multicast_rendezvous_points_example" {
  value = meraki_networks_switch_routing_multicast_rendezvous_points.example
}