
resource "meraki_networks_syslog_servers" "example" {

  network_id = "string"
  servers = [{

    host  = "1.2.3.4"
    port  = 443
    roles = ["Wireless event log", "URLs"]
  }]
}

output "meraki_networks_syslog_servers_example" {
  value = meraki_networks_syslog_servers.example
}