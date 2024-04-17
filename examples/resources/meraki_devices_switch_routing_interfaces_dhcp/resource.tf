
resource "meraki_devices_switch_routing_interfaces_dhcp" "example" {

  boot_file_name       = "home_boot_file"
  boot_next_server     = "1.2.3.4"
  boot_options_enabled = true
  dhcp_lease_time      = "1 day"
  dhcp_mode            = "dhcpServer"
  dhcp_options = [{

    code  = "5"
    type  = "text"
    value = "five"
  }]
  dhcp_relay_server_ips  = ["1.2.3.4"]
  dns_custom_nameservers = ["8.8.8.8, 8.8.4.4"]
  dns_nameservers_option = "custom"
  fixed_ip_assignments = [{

    ip   = "192.168.1.12"
    mac  = "22:33:44:55:66:77"
    name = "Cisco Meraki valued client"
  }]
  interface_id = "string"
  reserved_ip_ranges = [{

    comment = "A reserved IP range"
    end     = "192.168.1.10"
    start   = "192.168.1.1"
  }]
  serial = "string"
}

output "meraki_devices_switch_routing_interfaces_dhcp_example" {
  value = meraki_devices_switch_routing_interfaces_dhcp.example
}