
data "meraki_networks_floor_plans" "example" {

  network_id = "string"
}

output "meraki_networks_floor_plans_example" {
  value = data.meraki_networks_floor_plans.example.items
}

data "meraki_networks_floor_plans" "example" {

  network_id = "string"
}

output "meraki_networks_floor_plans_example" {
  value = data.meraki_networks_floor_plans.example.item
}
