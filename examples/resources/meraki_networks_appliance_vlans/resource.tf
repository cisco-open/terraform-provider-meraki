
resource "meraki_networks_appliance_vlans" "example" {

  appliance_ip              = "192.168.1.2"
  cidr                      = "192.168.1.0/24"
  dhcp_boot_options_enabled = true
  dhcp_handling             = "Run a DHCP server"
  dhcp_lease_time           = "30 minutes"
  dhcp_options = [{

    code  = "3"
    type  = "text"
    value = "five"
  }]
  group_policy_id = "101"
  id              = "1234"
  ipv6 = {

    enabled = true
    prefix_assignments = [{

      autonomous = false
      origin = {

        interfaces = ["wan0"]
        type       = "internet"
      }
      static_appliance_ip6 = "2001:db8:3c4d:15::1"
      static_prefix        = "2001:db8:3c4d:15::/64"
    }]
  }
  mandatory_dhcp = {

    enabled = true
  }
  mask               = 28
  name               = "My VLAN"
  network_id         = "string"
  subnet             = "192.168.1.0/24"
  template_vlan_type = "same"
}

output "meraki_networks_appliance_vlans_example" {
  value = meraki_networks_appliance_vlans.example
}