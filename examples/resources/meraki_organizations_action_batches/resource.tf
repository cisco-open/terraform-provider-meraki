
resource "meraki_organizations_action_batches" "example" {

  actions = [{

    operation = "create"
    resource  = "/devices/QXXX-XXXX-XXXX/switch/ports/3"
  }]
  confirmed       = true
  organization_id = "string"
  synchronous     = true
}

output "meraki_organizations_action_batches_example" {
  value = meraki_organizations_action_batches.example
}