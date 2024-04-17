
data "meraki_organizations_camera_permissions" "example" {

  organization_id     = "string"
  permission_scope_id = "string"
}

output "meraki_organizations_camera_permissions_example" {
  value = data.meraki_organizations_camera_permissions.example.item
}
