terraform {
  required_providers {
    meraki = {
      version = "0.2.12-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}


data "meraki_devices_management_interface" "example" {

  serial = "QBSB-BNH2-KDXJ"
}

output "meraki_devices_management_interface_example" {
  value = data.meraki_devices_management_interface.example.item
}
