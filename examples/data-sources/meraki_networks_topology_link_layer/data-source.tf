
data "meraki_networks_topology_link_layer" "example" {

  network_id = "string"
}

output "meraki_networks_topology_link_layer_example" {
  value = data.meraki_networks_topology_link_layer.example.item
}
