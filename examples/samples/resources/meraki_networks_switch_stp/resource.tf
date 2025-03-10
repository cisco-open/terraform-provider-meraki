terraform {
  required_providers {
    meraki = {
      version = "1.0.6-beta"
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

resource "meraki_networks_switch_stp" "example" {

  network_id   = var.my_network_id
  rstp_enabled = true
  stp_bridge_priority = [{

    stp_priority = 4096
    switches     = ["QBSB-VLHZ-JQCN"]
    },
  ]

}

# output "meraki_networks_switch_stp_example" {
#   value = meraki_networks_switch_stp.example
# }