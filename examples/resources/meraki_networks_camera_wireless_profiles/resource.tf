
resource "meraki_networks_camera_wireless_profiles" "example" {

  identity = {

    password = "password123"
    username = "identityname"
  }
  name       = "wireless profile A"
  network_id = "string"
  ssid = {

    auth_mode       = "8021x-radius"
    encryption_mode = "wpa-eap"
    name            = "ssid test"
    psk             = "sampleKey"
  }
}

output "meraki_networks_camera_wireless_profiles_example" {
  value = meraki_networks_camera_wireless_profiles.example
}