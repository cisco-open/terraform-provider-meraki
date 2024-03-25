
resource "meraki_devices_cellular_gateway_lan" "example" {

  fixed_ip_assignments = [{

    ip   = "192.168.0.10"
    mac  = "0b:00:00:00:00:ac"
    name = "server 1"
  }]
  reserved_ip_ranges = [{

    comment = "A reserved IP range"
    end     = "192.168.1.1"
    start   = "192.168.1.0"
  }]
  serial = "string"
}

output "meraki_devices_cellular_gateway_lan_example" {
  value = meraki_devices_cellular_gateway_lan.example
}