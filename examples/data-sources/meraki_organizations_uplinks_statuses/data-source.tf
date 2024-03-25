
data "meraki_organizations_uplinks_statuses" "example" {

  ending_before   = "string"
  iccids          = ["string"]
  network_ids     = ["string"]
  organization_id = "string"
  per_page        = 1
  serials         = ["string"]
  starting_after  = "string"
}

output "meraki_organizations_uplinks_statuses_example" {
  value = data.meraki_organizations_uplinks_statuses.example.items
}
