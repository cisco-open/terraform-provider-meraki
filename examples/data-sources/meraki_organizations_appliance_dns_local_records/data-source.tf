
data "meraki_organizations_appliance_dns_local_records" "example" {

  organization_id = "string"
  profile_ids     = ["string"]
}

output "meraki_organizations_appliance_dns_local_records_example" {
  value = data.meraki_organizations_appliance_dns_local_records.example.items
}
