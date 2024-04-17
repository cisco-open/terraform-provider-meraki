
data "meraki_organizations_camera_detections_history_by_boundary_by_interval" "example" {

  boundary_ids    = ["string"]
  boundary_types  = ["string"]
  duration        = 1
  organization_id = "string"
  per_page        = 1
  ranges          = ["string"]
}

output "meraki_organizations_camera_detections_history_by_boundary_by_interval_example" {
  value = data.meraki_organizations_camera_detections_history_by_boundary_by_interval.example.items
}
