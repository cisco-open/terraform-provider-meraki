
resource "meraki_devices" "example" {

  lat    = 37.4180951010362
  lng    = -122.098531723022
  name   = "My AP"
  serial = "string"
  tags   = ["recently-added"]
}

output "meraki_devices_example" {
  value = meraki_devices.example
}