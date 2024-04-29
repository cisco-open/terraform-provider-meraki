
data "meraki_organizations_camera_roles" "example" {

  organization_id = "string"
}

output "meraki_organizations_camera_roles_example" {
  value = data.meraki_organizations_camera_roles.example.items
}

data "meraki_organizations_camera_roles" "example" {

  organization_id = "string"
}

output "meraki_organizations_camera_roles_example" {
  value = data.meraki_organizations_camera_roles.example.item
}
