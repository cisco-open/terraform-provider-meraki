
resource "meraki_organizations_adaptive_policy_policies" "example" {

  acls = [{

    id   = "444"
    name = "Block web"
  }]
  destination_group = {

    id   = "333"
    name = "IoT Servers"
    sgt  = 51
  }
  last_entry_rule = "allow"
  organization_id = "string"
  source_group = {

    id   = "222"
    name = "IoT Devices"
    sgt  = 50
  }
}

output "meraki_organizations_adaptive_policy_policies_example" {
  value = meraki_organizations_adaptive_policy_policies.example
}