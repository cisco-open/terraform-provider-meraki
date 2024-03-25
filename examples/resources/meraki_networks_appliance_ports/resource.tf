
resource "meraki_networks_appliance_ports" "example" {

  access_policy         = "open"
  allowed_vlans         = "all"
  drop_untagged_traffic = false
  enabled               = true
  network_id            = "string"
  port_id               = "string"
  type                  = "access"
  vlan                  = 3
}

output "meraki_networks_appliance_ports_example" {
  value = meraki_networks_appliance_ports.example
}