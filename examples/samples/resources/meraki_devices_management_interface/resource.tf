terraform {
  required_providers {
    meraki = {
      version = "0.2.12-alpha"
      source  = "hashicorp.com/edu/meraki"
    }
  }
}



provider "meraki" {
  meraki_debug = "true"
}


resource "meraki_devices_management_interface" "sample_firewall_interface" {
  serial = "QBSA-TFWJ-U4L9"

  wan1 = {
    using_static_ip    = true
    static_ip          = "192.168.1.2"
    static_gateway_ip  = "192.168.1.1"
    static_subnet_mask = "255.255.255.0"
    static_dns         = ["8.8.8.8", "8.8.4.4"]
    vlan               = null
    wan_enabled        = "enabled"
  }

  wan2 = {
    using_static_ip = false
    wan_enabled     = "enabled"
  }
}