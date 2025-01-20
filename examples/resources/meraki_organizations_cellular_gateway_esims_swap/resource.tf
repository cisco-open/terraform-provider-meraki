
resource "meraki_organizations_cellular_gateway_esims_swap" "example" {

  organization_id = "string"
  parameters = {

    swaps = [{

      eid = "1234567890"
      target = {

        account_id         = "456"
        communication_plan = "A comm plan"
        rate_plan          = "A rate plan"
      }
    }]
  }
}

output "meraki_organizations_cellular_gateway_esims_swap_example" {
  value = meraki_organizations_cellular_gateway_esims_swap.example
}