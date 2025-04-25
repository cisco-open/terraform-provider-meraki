terraform {
  required_providers {
    meraki = {
      version = "1.1.2-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
  # meraki_user_agent = "new user agent"
}

resource "meraki_devices_live_tools_ping" "example" {

  serial = "Q2FV-VYGH-ZVB3"
  parameters = {
    count  = 2
    target = "8.8.8.8"
  }
}

output "meraki_devices_live_tools_ping_example" {
  value = meraki_devices_live_tools_ping.example
}