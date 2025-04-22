
resource "meraki_networks_settings" "example" {

  local_status_page = {

    authentication = {

      enabled  = false
      password = "miles123"
      username = "admin"
    }
  }
  local_status_page_enabled = true
  named_vlans = {

    enabled = true
  }
  network_id                 = "string"
  remote_status_page_enabled = true
  secure_port = {

    enabled = false
  }
}

output "meraki_networks_settings_example" {
  value = meraki_networks_settings.example
}