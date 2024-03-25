
data "meraki_networks_switch_routing_ospf" "example" {

  network_id = "string"
}

output "meraki_networks_switch_routing_ospf_example" {
  value = data.meraki_networks_switch_routing_ospf.example.item
}
