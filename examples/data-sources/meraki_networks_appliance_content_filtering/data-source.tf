
data "meraki_networks_appliance_content_filtering" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_content_filtering_example" {
  value = data.meraki_networks_appliance_content_filtering.example.item
}
