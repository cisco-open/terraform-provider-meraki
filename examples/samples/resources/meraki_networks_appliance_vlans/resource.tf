terraform {
  required_providers {
    meraki = {
      version = "1.1.2-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}

variable "my_network_id" {
  type    = string
  default = "*******" # Branch-1234
}

resource "meraki_networks_appliance_vlans" "example" {
  network_id                = "L_828099381482775374"
  appliance_ip              = "192.168.2.1"
  id                        = "2"
  subnet                    = "192.168.2.0/24"
  name                      = "guest 2"
  dhcp_handling             = "Do not respond to DHCP requests"
  dns_nameservers           = "upstream_dns"
  dhcp_lease_time           = "1 day"
  dhcp_boot_options_enabled = false
  reserved_ip_ranges = [
    { comment : "range1", start : "192.168.2.2", end : "192.168.2.3" }, # Not sent to the api on create
    { comment : "range2", start : "192.168.2.4", end : "192.168.2.5" }, # Not sent to the api on create
    { comment : "range3", start : "192.168.2.6", end : "192.168.2.7" }, # Not sent to the api on create
  ]
}