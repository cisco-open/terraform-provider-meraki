
resource "meraki_networks_switch_qos_rules_order" "example" {

  dscp           = 1
  dst_port       = 3000
  dst_port_range = "3000-3100"
  network_id     = "string"
  protocol       = "TCP"
  src_port       = 2000
  src_port_range = "70-80"
  vlan           = 100
}

output "meraki_networks_switch_qos_rules_order_example" {
  value = meraki_networks_switch_qos_rules_order.example
}