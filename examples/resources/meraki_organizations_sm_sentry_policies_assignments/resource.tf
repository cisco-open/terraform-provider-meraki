
resource "meraki_organizations_sm_sentry_policies_assignments" "example" {

  organization_id = "string"
  parameters = {

    items = [{

      network_id = "N_24329156"
      policies = [{

        group_policy_id = "1284392014819"
        policy_id       = "1284392014819"
        scope           = "withAny"
        sm_network_id   = "N_24329156"
        tags            = ["tag1", "tag2"]
      }]
    }]
  }
}

output "meraki_organizations_sm_sentry_policies_assignments_example" {
  value = meraki_organizations_sm_sentry_policies_assignments.example
}