
resource "meraki_networks_floor_plans" "example" {

  image_contents = "Q2lzY28gTWVyYWtp"
  name           = "HQ Floor Plan"
  network_id     = "string"
}

output "meraki_networks_floor_plans_example" {
  value = meraki_networks_floor_plans.example
}