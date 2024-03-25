
data "meraki_devices_cellular_sims" "example" {

  serial = "string"
}

output "meraki_devices_cellular_sims_example" {
  value = data.meraki_devices_cellular_sims.example.item
}
