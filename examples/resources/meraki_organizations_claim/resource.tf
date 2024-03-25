
resource "meraki_organizations_claim" "example" {

  organization_id = "string"
  parameters = {

    licenses = [{

      key  = "Z2XXXXXXXXXX"
      mode = "addDevices"
    }]
    orders  = ["4CXXXXXXX"]
    serials = ["Q234-ABCD-5678"]
  }
}

output "meraki_organizations_claim_example" {
  value = meraki_organizations_claim.example
}