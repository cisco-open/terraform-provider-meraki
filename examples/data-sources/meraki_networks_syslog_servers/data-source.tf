
data "meraki_networks_syslog_servers" "example" {

  network_id = "string"
}

output "meraki_networks_syslog_servers_example" {
  value = data.meraki_networks_syslog_servers.example.item
}
