terraform {
  required_providers {
    meraki = {
      version = "1.0.3-beta"
      source  = "hashicorp.com/edu/meraki"
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_organizations_policy_objects_groups" "example" {

  category = "NetworkObjectGroup"
  name     = "Issue 178"
  # object_ids      = [828099381482759082, 828099381482759083]
  object_ids      = [828099381482759082]
  organization_id = "828099381482762270"
}

output "meraki_organizations_policy_objects_groups_example" {
  value = meraki_organizations_policy_objects_groups.example
} 