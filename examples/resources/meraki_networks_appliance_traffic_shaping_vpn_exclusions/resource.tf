
resource "meraki_networks_appliance_traffic_shaping_vpn_exclusions" "example" {

  network_id = "string"
  parameters = {

    custom = [{

      destination = "192.168.3.0/24"
      port        = "8000"
      protocol    = "tcp"
    }]
    major_applications = [{

      id   = "meraki:vpnExclusion/application/2"
      name = "Office 365 Sharepoint"
    }]
  }
}

output "meraki_networks_appliance_traffic_shaping_vpn_exclusions_example" {
  value = meraki_networks_appliance_traffic_shaping_vpn_exclusions.example
}