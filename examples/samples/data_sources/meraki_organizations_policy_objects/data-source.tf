terraform {
  required_providers {
    meraki = {
      version = "1.0.7-beta"
      source  = "hashicorp.com/edu/meraki"
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
  meraki_base_url = "http://localhost:3002"
}

data "meraki_organizations_policy_objects" "example" {

  # ending_before   = "string"
  organization_id = "799301"
  # per_page        = 1
  # starting_after  = "string"}
  policy_object_id = "803329583532213295"
}

output "meraki_organizations_policy_objects_example" {
  value = data.meraki_organizations_policy_objects.example.item
}
