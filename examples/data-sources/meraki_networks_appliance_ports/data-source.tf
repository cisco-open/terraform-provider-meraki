
data "meraki_networks_appliance_ports" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_ports_example" {
  value = data.meraki_networks_appliance_ports.example.items
}

data "meraki_networks_appliance_ports" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_ports_example" {
  value = data.meraki_networks_appliance_ports.example.item
}
