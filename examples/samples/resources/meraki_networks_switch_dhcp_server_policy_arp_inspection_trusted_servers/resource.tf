
terraform {
  required_providers {
    meraki = {
      version = "1.1.3-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_switch_dhcp_server_policy_arp_inspection_trusted_servers" "example" {

  ipv4 = {

    address = "1.2.3.4"
  }
  mac        = "00:11:22:33:44:55"
  network_id = "L_828099381482771185"
  vlan       = 100
}

output "meraki_networks_switch_dhcp_server_policy_arp_inspection_trusted_servers_example" {
  value = meraki_networks_switch_dhcp_server_policy_arp_inspection_trusted_servers.example
}