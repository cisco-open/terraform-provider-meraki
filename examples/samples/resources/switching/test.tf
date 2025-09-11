terraform {
  required_providers {
    meraki = {
      version = "1.2.3-beta"
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
  default = "L_828099381482775374" # site 1
}

################################################################

# resource "meraki_networks_switch_mtu" "example" {

#   default_mtu_size = 9216
#   network_id       = var.my_network_id
# }

# output "meraki_networks_switch_mtu_example" {
#   value = meraki_networks_switch_mtu.example
# }

################################################################

# resource "meraki_networks_switch_stp" "example" {

#   network_id   = var.my_network_id
#   rstp_enabled = true
#     stp_bridge_priority = [{

#     stp_priority = 4096
#     switches     = ["QBSB-VLHZ-JQCN"]
#     }]

# }

# output "meraki_networks_switch_stp_example" {
#   value = meraki_networks_switch_stp.example
# }
################################################################

# resource "meraki_devices_switch_ports" "example" {

#   allowed_vlans             = "1,3,5-10"
#   enabled                   = true
#   name                      = "My switch port"
#   poe_enabled               = true
#   port_id                   = "11"
#   rstp_enabled                = true
#   serial                      = "QBSB-VLHZ-JQCN"
#   type                        = "trunk"
# }

# output "meraki_devices_switch_ports_example" {
#   value = meraki_devices_switch_ports.example
# }
################################################################

resource "meraki_networks_switch_access_control_lists" "example" {

  network_id = var.my_network_id
  rules = [{

    comment    = "Deny SSH"
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

output "meraki_networks_switch_access_control_lists_example" {
  value = meraki_networks_switch_access_control_lists.example
}
################################################################

# resource "meraki_networks_switch_link_aggregations" "example" {

#   network_id = var.my_network_id
#   switch_ports = [{
#     port_id = "5"
#     serial  = "QBSB-VLHZ-JQCN"
#   },
#   {
#     port_id = "6"
#     serial  = "QBSB-VLHZ-JQCN"
#   }]
# }

# output "meraki_networks_switch_link_aggregations_example" {
#   value = meraki_networks_switch_link_aggregations.example
# }
################################################################

################################################################

################################################################
