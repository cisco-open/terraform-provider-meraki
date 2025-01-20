
data "meraki_organizations_floor_plans_auto_locate_devices" "example" {

  ending_before   = "string"
  floor_plan_ids  = ["string"]
  network_ids     = ["string"]
  organization_id = "string"
  per_page        = 1
  starting_after  = "string"
}

output "meraki_organizations_floor_plans_auto_locate_devices_example" {
  value = data.meraki_organizations_floor_plans_auto_locate_devices.example.items
}
