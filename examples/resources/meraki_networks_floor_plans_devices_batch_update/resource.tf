
resource "meraki_networks_floor_plans_devices_batch_update" "example" {

  network_id = "string"
  parameters = {

    assignments = [{

      floor_plan = {

        id = "g_2176982374"
      }
      serial = "Q234-ABCD-5678"
    }]
  }
}

output "meraki_networks_floor_plans_devices_batch_update_example" {
  value = meraki_networks_floor_plans_devices_batch_update.example
}