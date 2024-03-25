
data "meraki_networks_switch_port_schedules" "example" {

  network_id = "string"
}

output "meraki_networks_switch_port_schedules_example" {
  value = data.meraki_networks_switch_port_schedules.example.items
}
