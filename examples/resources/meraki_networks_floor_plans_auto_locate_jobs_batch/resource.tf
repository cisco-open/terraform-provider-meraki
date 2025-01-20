
resource "meraki_networks_floor_plans_auto_locate_jobs_batch" "example" {

  network_id = "string"
  parameters = {

    jobs = [{

      floor_plan_id = "g_2176982374"
      refresh       = ["gnss", "ranging"]
      scheduled_at  = "2018-02-11T00:00:00Z"
    }]
  }
}

output "meraki_networks_floor_plans_auto_locate_jobs_batch_example" {
  value = meraki_networks_floor_plans_auto_locate_jobs_batch.example
}