
resource "meraki_organizations_appliance_dns_split_profiles" "example" {

  hostnames = ["*.test1.com", "*.test2.com"]
  name      = "Default profile"
  nameservers = {

    addresses = ["12.1.10.1"]
  }
  organization_id = "string"
}

output "meraki_organizations_appliance_dns_split_profiles_example" {
  value = meraki_organizations_appliance_dns_split_profiles.example
}