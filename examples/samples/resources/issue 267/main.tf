terraform {
  required_providers {
    meraki = {
      version = "1.2.4-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}


variable "static_routes" {
  description = "Values for static routes"
  type = list(object({
    enabled = bool
    gateway_ip = string
    gateway_vlan_id = string
    name = string
    subnet = string
    static_route_id = string
  }))
  default = [
    {
      enabled = true
      gateway_ip = "192.168.200.2"
      gateway_vlan_id = null
      name = "route01"
      subnet = "192.168.144.0/24"
      static_route_id = null
    },
    {
      enabled = true
      gateway_ip = "192.168.200.2"
      gateway_vlan_id = null
      name = "route02"
      subnet = "192.168.145.0/24"
      static_route_id = null
    },
  ]
}
resource "meraki_networks_appliance_static_routes" "appliance_static_routes" {
  for_each = { for static_route in var.static_routes : static_route.name => static_route }
  enabled  = each.value.enabled
  gateway_ip = each.value.gateway_ip
  gateway_vlan_id = each.value.gateway_vlan_id
  name  = each.value.name
  network_id ="L_828099381482771185"
  subnet = each.value.subnet
  static_route_id = each.value.static_route_id
}