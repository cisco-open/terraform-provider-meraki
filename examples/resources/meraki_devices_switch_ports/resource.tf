
resource "meraki_devices_switch_ports" "example" {

  access_policy_number     = 2
  access_policy_type       = "Sticky MAC allow list"
  adaptive_policy_group_id = "123"
  allowed_vlans            = "1,3,5-10"
  dai_trusted              = false
  dot3az = {

    enabled = false
  }
  enabled                   = true
  flexible_stacking_enabled = true
  isolation_enabled         = false
  link_negotiation          = "Auto negotiate"
  mac_allow_list            = ["34:56:fe:ce:8e:a0", "34:56:fe:ce:8e:a1"]
  name                      = "My switch port"
  peer_sgt_capable          = false
  poe_enabled               = true
  port_id                   = "string"
  port_schedule_id          = "1234"
  profile = {

    enabled = false
    id      = "1284392014819"
    iname   = "iname"
  }
  rstp_enabled                = true
  serial                      = "string"
  sticky_mac_allow_list       = ["34:56:fe:ce:8e:b0", "34:56:fe:ce:8e:b1"]
  sticky_mac_allow_list_limit = 5
  storm_control_enabled       = true
  stp_guard                   = "disabled"
  tags                        = ["tag1", "tag2"]
  type                        = "access"
  udld                        = "Alert only"
  vlan                        = 10
  voice_vlan                  = 20
}

output "meraki_devices_switch_ports_example" {
  value = meraki_devices_switch_ports.example
}