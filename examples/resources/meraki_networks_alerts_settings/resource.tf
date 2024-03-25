
resource "meraki_networks_alerts_settings" "example" {

  alerts = [{

    alert_destinations = {

      all_admins      = false
      emails          = ["miles@meraki.com"]
      http_server_ids = ["aHR0cHM6Ly93d3cuZXhhbXBsZS5jb20vd2ViaG9va3M="]
      snmp            = false
    }
    enabled = true
    filters = {

      timeout = 60
    }
    type = "gatewayDown"
  }]
  default_destinations = {

    all_admins      = true
    emails          = ["miles@meraki.com"]
    http_server_ids = ["aHR0cHM6Ly93d3cuZXhhbXBsZS5jb20vd2ViaG9va3M="]
    snmp            = true
  }
  network_id = "string"
}

output "meraki_networks_alerts_settings_example" {
  value = meraki_networks_alerts_settings.example
}