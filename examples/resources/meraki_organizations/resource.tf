
resource "meraki_organizations" "example" {

  management = {

    details = [{

      name  = "MSP ID"
      value = "123456"
    }]
  }
  name = "My organization"
}

output "meraki_organizations_example" {
  value = meraki_organizations.example
}