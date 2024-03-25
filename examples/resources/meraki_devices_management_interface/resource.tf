
resource "meraki_devices_management_interface" "example" {

  serial = "string"
  wan1 = {

    static_dns         = ["1.2.3.2", "1.2.3.3"]
    static_gateway_ip  = "1.2.3.1"
    static_ip          = "1.2.3.4"
    static_subnet_mask = "255.255.255.0"
    using_static_ip    = true
    vlan               = 7
    wan_enabled        = "not configured"
  }
  wan2 = {

    using_static_ip = false
    vlan            = 2
    wan_enabled     = "enabled"
  }
}

output "meraki_devices_management_interface_example" {
  value = meraki_devices_management_interface.example
}