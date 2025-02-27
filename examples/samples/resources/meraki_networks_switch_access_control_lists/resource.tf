terraform {
  required_providers {
    meraki = {
      version = "1.0.4-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}
resource "meraki_networks_switch_access_control_lists" "example" {

  network_id = "L_828099381482775374"
  rules = [{

    comment    = "Allow SSH"
    dst_cidr   = "172.16.30.0/24"
    dst_port   = "22"
    ip_version = "ipv4"
    policy     = "deny"
    protocol   = "tcp"
    src_cidr   = "10.1.10.0/24"
    src_port   = "any"
    vlan       = "10"
  }]
}

# output "meraki_networks_switch_access_control_lists_example" {
#   value = meraki_networks_switch_access_control_lists.example
# }