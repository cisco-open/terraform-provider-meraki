
resource "meraki_organizations_appliance_security_intrusion" "example" {

  allowed_rules = [{

    message = "SQL sa login failed"
    rule_id = "meraki:intrusion/snort/GID/01/SID/688"
  }]
  organization_id = "string"
}

output "meraki_organizations_appliance_security_intrusion_example" {
  value = meraki_organizations_appliance_security_intrusion.example
}