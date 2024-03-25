# - hosts: localhost
#   gather_facts: false
#   tasks:
#     - name: Create
#       cisco.meraki.organizations_adaptive_policy_groups:
#         meraki_suppress_logging: true
#         state: present
#         # description: Group of XYZ Corp Employees
#         # isDefaultGroup: false
#         name: Employee Group
#         organizationId: "828099381482762270"
#         # policyObjects:
#         # - id: '2345'
#         #   name: Example Policy Object
#         # requiredIpMappings: []
#         sgt: 10005

terraform {
  required_providers {
    meraki = {
      version = "0.1.0-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

resource "meraki_organizations_adaptive_policy_groups" "example" {
  provider    = meraki
  description = "Group of XYZ Corp Employees"
  name            = "Employee Group"
  organization_id = "828099381482762270"
  #   policy_objects {

  #     id = "string"
  #     name = "string"
  #   }
  sgt = 1005

}

output "meraki_organizations_adaptive_policy_groups_example" {
  value = meraki_organizations_adaptive_policy_groups.example
}
