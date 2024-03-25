
resource "meraki_networks_switch_access_policies" "example" {

  access_policy_type = "Hybrid authentication"
  dot1x = {

    control_direction = "inbound"
  }
  guest_port_bouncing   = false
  guest_vlan_id         = 100
  host_mode             = "Single-Host"
  increase_access_speed = false
  name                  = "Access policy #1"
  network_id            = "string"
  radius = {

    critical_auth = {

      data_vlan_id        = 100
      suspend_port_bounce = true
      voice_vlan_id       = 100
    }
    failed_auth_vlan_id        = 100
    re_authentication_interval = 120
  }
  radius_accounting_enabled = true
  radius_accounting_servers = [{

    host   = "1.2.3.4"
    port   = 22
    secret = "secret"
  }]
  radius_coa_support_enabled = false
  radius_group_attribute     = "11"
  radius_servers = [{

    host   = "1.2.3.4"
    port   = 22
    secret = "secret"
  }]
  radius_testing_enabled             = false
  url_redirect_walled_garden_enabled = true
  url_redirect_walled_garden_ranges  = ["192.168.1.0/24"]
  voice_vlan_clients                 = true
}

output "meraki_networks_switch_access_policies_example" {
  value = meraki_networks_switch_access_policies.example
}