
data "meraki_organizations_snmp" "example" {

  organization_id = "string"
}

output "meraki_organizations_snmp_example" {
  value = data.meraki_organizations_snmp.example.item
}
