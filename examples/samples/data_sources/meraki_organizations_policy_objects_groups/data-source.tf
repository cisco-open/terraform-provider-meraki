terraform {
  required_providers {
    meraki = {
      version = "0.2.12-alpha"
      source  = "hashicorp.com/edu/meraki"
    }
  }
}

data "meraki_organizations_policy_objects_groups" "example" {

  # ending_before   = "string"
  # organization_id = "string"
  per_page        = 1
  # starting_after  = "string"
}

output "meraki_organizations_policy_objects_groups_example" {
  value = data.meraki_organizations_policy_objects_groups.example.item
}