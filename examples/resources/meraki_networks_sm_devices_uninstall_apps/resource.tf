
resource "meraki_networks_sm_devices_uninstall_apps" "example" {

  device_id  = "string"
  network_id = "string"
  parameters = {

    app_ids = ["1284392014819", "2983092129865"]
  }
}

output "meraki_networks_sm_devices_uninstall_apps_example" {
  value = meraki_networks_sm_devices_uninstall_apps.example
}