terraform {
  required_providers {
    meraki = {
      version = "1.2.4-beta"
      source  = "hashicorp.com/edu/meraki"
    }
  }
}

resource "meraki_organizations_policy_objects" "example" {

  category        = "network"
  cidr            = "10.0.0.0/24"
  fqdn            = "example.com"
  group_ids       = ["8"]
  ip              = "1.2.3.4"
  mask            = "255.255.0.0"
  name            = "Web Servers - Datacenter 10"
  organization_id = "string"
  type            = "cidr"
}

output "meraki_organizations_policy_objects_example" {
  value = meraki_organizations_policy_objects.example
}