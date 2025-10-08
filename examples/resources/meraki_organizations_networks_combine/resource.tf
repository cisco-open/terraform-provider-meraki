
resource "meraki_organizations_networks_combine" "example" {

  organization_id = "string"
  parameters = {

    enrollment_string = "my-enrollment-string"
    name              = "Long Island Office"
    network_ids       = ["N_828099381482850157", "N_828099381482850161"]
  }
}

output "meraki_organizations_networks_combine_example" {
  value = meraki_organizations_networks_combine.example
}