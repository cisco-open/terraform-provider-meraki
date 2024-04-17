
resource "meraki_networks_switch_dhcp_server_policy" "example" {

  alerts = {

    email = {

      enabled = true
    }
  }
  allowed_servers = ["00:50:56:00:00:01", "00:50:56:00:00:02"]
  arp_inspection = {

    enabled = true
  }
  blocked_servers = ["00:50:56:00:00:03", "00:50:56:00:00:04"]
  default_policy  = "block"
  network_id      = "string"
}

output "meraki_networks_switch_dhcp_server_policy_example" {
  value = meraki_networks_switch_dhcp_server_policy.example
}