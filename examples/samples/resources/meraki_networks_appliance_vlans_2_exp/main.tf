terraform {
  required_providers {
    meraki = {
      version = "0.2.5-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}


resource "meraki_networks_appliance_vlans" "this_vlan_1" {
  network_id                = "L_828099381482775374"
  name                      = "MGMT"
  id                        = "1" 
  appliance_ip              = "10.76.237.17"
  subnet                    = "10.76.237.16/28"
  dhcp_handling             = "Run a DHCP server"
  dhcp_lease_time           = "1 day"
  dns_nameservers           = "upstream_dns"
  dhcp_boot_options_enabled = false
  ipv6 = {
    enabled = false
  }
  mandatory_dhcp = {
    enabled = false
  }
}