
data "meraki_organizations_action_batches" "example" {

  organization_id = "string"
  status          = "string"
}

output "meraki_organizations_action_batches_example" {
  value = data.meraki_organizations_action_batches.example.items
}

data "meraki_organizations_action_batches" "example" {

  organization_id = "string"
  status          = "string"
}

output "meraki_organizations_action_batches_example" {
  value = data.meraki_organizations_action_batches.example.item
}
