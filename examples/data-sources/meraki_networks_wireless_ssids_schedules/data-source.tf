
data "meraki_networks_wireless_ssids_schedules" "example" {

  network_id = "string"
  number     = "string"
}

output "meraki_networks_wireless_ssids_schedules_example" {
  value = data.meraki_networks_wireless_ssids_schedules.example.item
}
