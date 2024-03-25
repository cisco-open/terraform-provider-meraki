
data "meraki_networks_appliance_warm_spare" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_warm_spare_example" {
  value = data.meraki_networks_appliance_warm_spare.example.item
}
