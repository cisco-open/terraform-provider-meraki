
data "meraki_networks_appliance_content_filtering_categories" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_content_filtering_categories_example" {
  value = data.meraki_networks_appliance_content_filtering_categories.example.item
}
