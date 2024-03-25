
data "meraki_networks_appliance_prefixes_delegated_statics" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_prefixes_delegated_statics_example" {
  value = data.meraki_networks_appliance_prefixes_delegated_statics.example.items
}

data "meraki_networks_appliance_prefixes_delegated_statics" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_prefixes_delegated_statics_example" {
  value = data.meraki_networks_appliance_prefixes_delegated_statics.example.item
}
