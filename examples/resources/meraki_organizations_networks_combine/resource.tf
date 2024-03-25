
resource "meraki_organizations_networks_combine" "example" {

  organization_id = "string"
  parameters = {

    enrollment_string = "my-enrollment-string"
    name              = "Long Island Office"
    network_ids       = ["N_1234", "N_5678"]
  }
}

output "meraki_organizations_networks_combine_example" {
  value = meraki_organizations_networks_combine.example
}