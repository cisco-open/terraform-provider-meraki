terraform {
  required_providers {
    meraki = {
      source  = "hashicorp.com/edu/meraki"
      version = "1.0.7-beta"
    }
  }

  required_version = ">= 1.2.0"
}
resource "meraki_networks_appliance_static_routes" "example" {
  provider   = meraki
  gateway_ip = "1.2.3.5"
  name       = "My route"
  network_id = "L_828099381482771185"
  subnet     = "192.168.1.0/24"
}

output "meraki_networks_appliance_static_routes_example" {
  value = meraki_networks_appliance_static_routes.example
}