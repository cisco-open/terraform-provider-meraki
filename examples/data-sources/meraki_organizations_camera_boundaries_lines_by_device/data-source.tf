
data "meraki_organizations_camera_boundaries_lines_by_device" "example" {

  organization_id = "string"
  serials         = ["string"]
}

output "meraki_organizations_camera_boundaries_lines_by_device_example" {
  value = data.meraki_organizations_camera_boundaries_lines_by_device.example.items
}
