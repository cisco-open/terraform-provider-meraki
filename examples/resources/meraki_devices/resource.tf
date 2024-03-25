
resource "meraki_devices" "example" {

  lat    = 37.4180951010362
  lng    = -122.098531723022
  mac    = "00:11:22:33:44:55"
  name   = "My AP"
  serial = "string"
  tags   = ["recently-added"]
}

output "meraki_devices_example" {
  value = meraki_devices.example
}