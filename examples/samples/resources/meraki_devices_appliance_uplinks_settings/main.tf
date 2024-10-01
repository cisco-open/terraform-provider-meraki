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

resource "meraki_devices_appliance_uplinks_settings" "this" {

  serial = "QBSA-TFWJ-U4L9"
  interfaces = {
    wan1 = {
      enabled = true
      svis = {
        ipv4 = {
          assignment_mode = "dynamic"
        }
      }
      vlan_tagging = {
        enabled = true
        vlan_id = "10"
      }
    }
  }
}