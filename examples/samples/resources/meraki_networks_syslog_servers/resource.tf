terraform {
  required_providers {
    meraki = {
      version = "0.2.9-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_syslog_servers" "example" {

  network_id = "L_828099381482775375"
  servers = [{

    host  = "1.2.3.42"
    port  = 443
    roles = ["Wireless event log", "URLs"]
  }]
}

output "meraki_networks_syslog_servers_example" {
  value = meraki_networks_syslog_servers.example
}