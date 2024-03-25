
data "meraki_networks_switch_stacks" "example" {

  network_id = "string"
}

output "meraki_networks_switch_stacks_example" {
  value = data.meraki_networks_switch_stacks.example.items
}

data "meraki_networks_switch_stacks" "example" {

  network_id = "string"
}

output "meraki_networks_switch_stacks_example" {
  value = data.meraki_networks_switch_stacks.example.item
}
