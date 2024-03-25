
data "meraki_organizations_sensor_readings_latest" "example" {

  ending_before   = "string"
  metrics         = ["string"]
  network_ids     = ["string"]
  organization_id = "string"
  per_page        = 1
  serials         = ["string"]
  starting_after  = "string"
}

output "meraki_organizations_sensor_readings_latest_example" {
  value = data.meraki_organizations_sensor_readings_latest.example.items
}
