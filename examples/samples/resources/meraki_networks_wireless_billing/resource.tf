terraform {
  required_providers {
    meraki = {
      version = "1.1.6-beta"
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
  default = "L_828099381482775375" # site 3
}

################################################################

resource "meraki_networks_wireless_billing" "example" {

  currency   = "USD"
  network_id = var.my_network_id
  plans = [{

    bandwidth_limits = {

      limit_down = 1000
      limit_up   = 1000
    }
    price      = 5
    time_limit = "1 hour"
    },
    {

      bandwidth_limits = {

        limit_down = 500
        limit_up   = 1000
      }
      price      = 10
      time_limit = "1 hour"
  }]
}

output "meraki_networks_wireless_billing_example" {
  value = meraki_networks_wireless_billing.example
}