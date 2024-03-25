
resource "meraki_networks_firmware_upgrades_staged_groups" "example" {

  assigned_devices = {

    devices = [{

      name   = "Device Name"
      serial = "Q234-ABCD-5678"
    }]
    switch_stacks = [{

      id   = "1234"
      name = "Stack Name"
    }]
  }
  description = "The description of the group"
  is_default  = false
  name        = "My Staged Upgrade Group"
  network_id  = "string"
}

output "meraki_networks_firmware_upgrades_staged_groups_example" {
  value = meraki_networks_firmware_upgrades_staged_groups.example
}