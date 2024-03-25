
resource "meraki_organizations_licensing_coterm_licenses_move" "example" {

  organization_id = "string"
  parameters = {

    destination = {

      mode            = "addDevices"
      organization_id = "123"
    }
    licenses = [{

      counts = [{

        count = 5
        model = "MR Enterprise"
      }]
      key = "Z2AA-BBBB-CCCC"
    }]
  }
}

output "meraki_organizations_licensing_coterm_licenses_move_example" {
  value = meraki_organizations_licensing_coterm_licenses_move.example
}