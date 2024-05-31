terraform {
  required_providers {
    meraki = {
      version = "0.2.3-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

resource "meraki_devices_camera_generate_snapshot" "example" {

  serial = "QBSD-WABS-BH7V"
  parameters = {

    fullframe = "false"
    timestamp = "2024-02-22T14:30:00Z"
  }
}