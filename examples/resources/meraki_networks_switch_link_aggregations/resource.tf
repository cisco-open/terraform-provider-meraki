
resource "meraki_networks_switch_link_aggregations" "example" {

  network_id = "string"
  switch_ports = [{

    port_id = "1"
    serial  = "Q234-ABCD-0001"
  }]
}

output "meraki_networks_switch_link_aggregations_example" {
  value = meraki_networks_switch_link_aggregations.example
}