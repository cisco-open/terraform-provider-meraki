# - name: change name of device
#   cisco.meraki.devices:
#     name: new name 4
#     serial: QBSB-D5ZD-9CXT
#     # organizationId: "{{org_id}}"
#     state: present 
#     meraki_suppress_logging: false

terraform {
  required_providers {
    meraki = {
      version = "1.0.3-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

resource "meraki_devices" "example" {

  name = "new name 6"
  # floor_plan_id = "string"
  # lat = 1.0
  # lng = 1.0
  # mac = "string"
  # move_map_marker = "false"
  # name = "string"
  notes  = "This is a test from terraform 2"
  serial = "QBSA-TFWJ-U4L9"
  # switch_profile_id = "string"
  # tags = ["string"]
}

output "meraki_devices_example" {
  value = meraki_devices.example
}