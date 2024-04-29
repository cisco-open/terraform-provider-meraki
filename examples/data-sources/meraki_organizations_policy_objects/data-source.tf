
data "meraki_organizations_policy_objects" "example" {

  ending_before   = "string"
  organization_id = "string"
  per_page        = 1
  starting_after  = "string"
}

output "meraki_organizations_policy_objects_example" {
  value = data.meraki_organizations_policy_objects.example.item
}
