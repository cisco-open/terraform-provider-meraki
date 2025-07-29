terraform {
  required_providers {
    meraki = {
      version = "1.1.8-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}


resource "meraki_networks_switch_access_policies" "this_site_access_policy" {
  network_id = "L_828099381482775374"
  name       = "dot1xmab"
  radius_servers = [{
    host   = "192.168.1.2"
    port   = 1812
    secret = "<-blanked->"
  }]
  radius_accounting_enabled = true
  radius_accounting_servers = [{
    host   = "192.168.1.3"
    port   = 1813
    secret = "<-blanked->"
  }]
  radius_testing_enabled     = true
  radius_coa_support_enabled = false
  radius_group_attribute     = ""
  host_mode                  = "Single-Host"
  access_policy_type         = "Hybrid authentication"
  increase_access_speed      = false
  dot1x = {
    control_direction = "both"
  }
  radius = {
    critical_auth = {
      data_vlan_id        = 10
      suspend_port_bounce = true
      voice_vlan_id       = "11"
    }
    failed_auth_vlan_id        = 10
    re_authentication_interval = null
  }
  url_redirect_walled_garden_enabled = false
  guest_port_bouncing                = true
  guest_vlan_id                      = "10"
  voice_vlan_clients                 = true
}