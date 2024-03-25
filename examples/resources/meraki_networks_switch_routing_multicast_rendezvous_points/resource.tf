
resource "meraki_networks_switch_routing_multicast_rendezvous_points" "example" {

  interface_ip    = "192.168.1.2"
  multicast_group = "192.168.128.0/24"
  network_id      = "string"
}

output "meraki_networks_switch_routing_multicast_rendezvous_points_example" {
  value = meraki_networks_switch_routing_multicast_rendezvous_points.example
}