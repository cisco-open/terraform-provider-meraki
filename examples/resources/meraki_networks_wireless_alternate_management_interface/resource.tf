
resource "meraki_networks_wireless_alternate_management_interface" "example" {

  access_points = [{

    alternate_management_ip = "1.2.3.4"
    dns1                    = "8.8.8.8"
    dns2                    = "8.8.4.4"
    gateway                 = "1.2.3.5"
    serial                  = "Q234-ABCD-5678"
    subnet_mask             = "255.255.255.0"
  }]
  enabled    = true
  network_id = "string"
  protocols  = ["radius", "snmp", "syslog", "ldap"]
  vlan_id    = 100
}

output "meraki_networks_wireless_alternate_management_interface_example" {
  value = meraki_networks_wireless_alternate_management_interface.example
}