
resource "meraki_devices" "example" {

  address           = "1600 Pennsylvania Ave"
  floor_plan_id     = "g_2176982374"
  lat               = 37.4180951010362
  lng               = -122.098531723022
  move_map_marker   = true
  name              = "My AP"
  notes             = "My AP's note"
  serial            = "string"
  switch_profile_id = "1234"
  tags              = [" recently-added "]
}

output "meraki_devices_example" {
  value = meraki_devices.example
}