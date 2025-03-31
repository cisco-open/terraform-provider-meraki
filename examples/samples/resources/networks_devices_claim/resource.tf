# - name: Create
#       cisco.meraki.networks_devices_claim:
#         meraki_output_log: true
#         networkId: "{{network_id}}"
#         serials:
#         - QBSD-WABS-BH7V

terraform {
  required_providers {
    meraki = {
      version = "1.0.7-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

resource "meraki_networks_devices_claim" "example" {

  network_id = "L_828099381482771185"
  parameters = {
    serials = ["QBSD-WABS-BH7V"]
  }
}

output "meraki_networks_devices_claim_example" {
  value = meraki_networks_devices_claim.example
}