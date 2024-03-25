
resource "meraki_networks_switch_alternate_management_interface" "example" {

  enabled    = true
  network_id = "string"
  protocols  = ["radius", "snmp", "syslog"]
  switches = [{

    alternate_management_ip = "1.2.3.4"
    gateway                 = "1.2.3.5"
    serial                  = "Q234-ABCD-5678"
    subnet_mask             = "255.255.255.0"
  }]
  vlan_id = 100
}

output "meraki_networks_switch_alternate_management_interface_example" {
  value = meraki_networks_switch_alternate_management_interface.example
}