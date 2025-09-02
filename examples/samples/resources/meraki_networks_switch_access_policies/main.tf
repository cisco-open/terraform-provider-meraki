terraform {
  required_providers {
    meraki = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug    = "true"
}

variable "ise_servers" {
  description = "Lista de servidores ISE con su IP, puerto y secreto compartido"
  type = list(object({
    server_ip         = string
    ise_shared_secret = string
  }))
  default = [
    {
      server_ip         = "192.168.1.10"
      ise_shared_secret = "secreto123"
    },
    {
      server_ip         = "192.168.1.11"
      ise_shared_secret = "secreto456"
    }
  ]
}


resource "meraki_networks_switch_access_policies" "example" {

  network_id         = "L_828099381482775486"
  access_policy_type = "Hybrid authentication"
  dot1x = {
    control_direction = "both"
  }
  guest_port_bouncing   = true
  host_mode             = "Single-Host"
  increase_access_speed = false
  name                  = "dot1x/mab"
  radius = {
    critical_auth = {
      data_vlan_id        = 10
      suspend_port_bounce = true
    }
    failed_auth_vlan_id        = 10
    re_authentication_interval = 3600
  }
  guest_vlan_id             = 100
  radius_accounting_enabled = true
  radius_accounting_servers = [
    for server in var.ise_servers : {
      host   = server.server_ip
      port   = 1813
      secret = server.ise_shared_secret
  }]
  radius_coa_support_enabled = true
  radius_group_attribute     = ""
  radius_servers = [
    for server in var.ise_servers : {
      host   = server.server_ip
      port   = 1812
      secret = server.ise_shared_secret
  }]
  radius_testing_enabled             = true
  url_redirect_walled_garden_enabled = false
  voice_vlan_clients                 = false
}

# output "meraki_networks_switch_access_policies_example" {
#   value = meraki_networks_switch_access_policies.example
# }