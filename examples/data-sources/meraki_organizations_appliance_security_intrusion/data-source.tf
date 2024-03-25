
data "meraki_organizations_appliance_security_intrusion" "example" {

  organization_id = "string"
}

output "meraki_organizations_appliance_security_intrusion_example" {
  value = data.meraki_organizations_appliance_security_intrusion.example.item
}
