terraform {
  required_providers {
    meraki = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug    = "true"
  meraki_base_url = "http://localhost:3002"
}



# resource "meraki_organizations_camera_roles" "example" {
#   organization_id = "441848"

#   name = "test"

#   applied_org_wide = [
#     {
#       # permission_level    = "full_access"
#       # permission_scope    = "sensor"
#       permission_scope_id = "6"
#     },
#     {
#       # permission_level    = "view_and_export"
#       # permission_scope    = "camera_video"
#       permission_scope_id = "2"
#     }
#   ]
# }