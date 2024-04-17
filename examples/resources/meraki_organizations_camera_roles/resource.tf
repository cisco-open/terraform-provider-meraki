
resource "meraki_organizations_camera_roles" "example" {

  applied_on_devices = [{

    id                  = ""
    permission_scope_id = "1"
    tag                 = "reception-desk"
  }]
  applied_on_networks = [{

    id                  = ""
    permission_scope_id = "2"
    tag                 = "building-a"
  }]
  applied_org_wide = [{

    permission_scope_id = "2"
  }]
  name            = "Security_Guard"
  organization_id = "string"
}

output "meraki_organizations_camera_roles_example" {
  value = meraki_organizations_camera_roles.example
}