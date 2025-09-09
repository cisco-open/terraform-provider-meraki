terraform {
  required_providers {
    meraki = {
      version = "1.2.2-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}

################################################################

################################################################

################################################################

################################################################

################################################################

################################################################

################################################################

# resource "meraki_organizations_admins" "example" {

#   authentication_method = "Email"
#   email                 = "miles@yopmail.com"
#   name                  = "Miles Meraki"
#   org_access      = "read-only"
#   organization_id = "828099381482762270"
#   tags = [{

#     access = "read-only"
#     tag    = "west"
#   }]
# }

# output "meraki_organizations_admins_example" {
#   value = meraki_organizations_admins.example
# }
################################################################

# resource "meraki_networks_webhooks_http_servers" "example" {

#   name       = "Example Webhook Server"
#   network_id = "L_828099381482775375"
#   payload_template = {

#     name                = "Meraki (included)"
#     payload_template_id = "wpt_00001"
#   }
#   shared_secret = "shhh"
#   url           = "https://example.com"
# }

# output "meraki_networks_webhooks_http_servers_example" {
#   value = meraki_networks_webhooks_http_servers.example
# }
################################################################
# resource "meraki_networks_syslog_servers" "example" {

#   network_id = "L_828099381482775375"
#   servers = [{

#     host  = "1.2.3.4"
#     port  = 443
#     roles = ["Wireless event log", "URLs"]
#   }]
# }

# output "meraki_networks_syslog_servers_example" {
#   value = meraki_networks_syslog_servers.example
# }
################################################################

# resource "meraki_networks_snmp" "example" {

#   access     = "users"
#   network_id = "L_828099381482775375"
#   users = [{

#     passphrase = "hunter2"
#     username   = "AzureDiamond"
#   }]
# }

# output "meraki_networks_snmp_example" {
#   value = meraki_networks_snmp.example
# }
################################################################

# resource "meraki_networks_alerts_settings" "example" {

#   alerts = [{

#     alert_destinations = {

#       all_admins      = false
#       emails          = ["miles@meraki.com"]
#     }
#     enabled = true
#     type = "gatewayDown"
#   }]
#   default_destinations = {

#     all_admins      = true
#     emails          = ["miles@meraki.com"]
#   }
#   network_id = "L_828099381482775375"
# }

# output "meraki_networks_alerts_settings_example" {
#   value = meraki_networks_alerts_settings.example
# }
################################################################

# resource "meraki_networks_netflow" "example" {

#   collector_ip      = "1.2.3.4"
#   collector_port    = 443
#   network_id        = "L_828099381482771185"
#   reporting_enabled = true
# }

# output "meraki_networks_netflow_example" {
#   value = meraki_networks_netflow.example
# }
################################################################

resource "meraki_networks_settings" "example" {

  local_status_page = {

    authentication = {
      username = "miles"
      enabled  = false
      password = "miles123*Asadsadsa"
    }
  }
  local_status_page_enabled  = true
  network_id                 = "L_828099381482775375"
  remote_status_page_enabled = true
  secure_port = {

    enabled = false
  }
}

output "meraki_networks_settings_example" {
  value     = meraki_networks_settings.example
  sensitive = true
}