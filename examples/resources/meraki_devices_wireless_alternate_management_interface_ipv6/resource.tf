
resource "meraki_devices_wireless_alternate_management_interface_ipv6" "example" {

  serial = "string"
  parameters = {

    addresses = [{

      address         = "2001:db8:3c4d:15::1"
      assignment_mode = "static"
      gateway         = "fe80:db8:c15:c0:d0c::10ca:1d02"
      nameservers = {

        addresses = ["2001:db8:3c4d:15::1", "2001:db8:3c4d:15::1"]
      }
      prefix   = "2001:db8:3c4d:15::/64"
      protocol = "ipv6"
    }]
  }
}

output "meraki_devices_wireless_alternate_management_interface_ipv6_example" {
  value = meraki_devices_wireless_alternate_management_interface_ipv6.example
}