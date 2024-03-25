
data "meraki_networks_switch_qos_rules_order" "example" {

  network_id = "string"
}

output "meraki_networks_switch_qos_rules_order_example" {
  value = data.meraki_networks_switch_qos_rules_order.example.items
}

data "meraki_networks_switch_qos_rules_order" "example" {

  network_id = "string"
}

output "meraki_networks_switch_qos_rules_order_example" {
  value = data.meraki_networks_switch_qos_rules_order.example.item
}
