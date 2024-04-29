
resource "meraki_networks_sm_devices_install_apps" "example" {

  device_id  = "string"
  network_id = "string"
  parameters = {

    app_ids = ["1284392014819", "2983092129865"]
    force   = false
  }
}

output "meraki_networks_sm_devices_install_apps_example" {
  value = meraki_networks_sm_devices_install_apps.example
}