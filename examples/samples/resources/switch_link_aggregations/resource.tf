terraform {
  required_providers {
    meraki = {
      version = "1.0.7-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_switch_link_aggregations" "example" {

  network_id = "L_828099381482775374"
  switch_ports = [{
    port_id = "9"
    serial  = "QBSB-VLHZ-JQCN"
    },
    {
      port_id = "10"
      serial  = "QBSB-VLHZ-JQCN"
    },
    {
      port_id = "11"
      serial  = "QBSB-VLHZ-JQCN"
    }
  ]
}

output "meraki_networks_switch_link_aggregations_example" {
  value = meraki_networks_switch_link_aggregations.example
}