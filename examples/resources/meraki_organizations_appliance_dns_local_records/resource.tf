
resource "meraki_organizations_appliance_dns_local_records" "example" {

  address         = "10.1.1.0"
  hostname        = "www.test.com"
  organization_id = "string"
  profile = {

    id = "1"
  }
}

output "meraki_organizations_appliance_dns_local_records_example" {
  value = meraki_organizations_appliance_dns_local_records.example
}