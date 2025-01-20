
resource "meraki_devices_cellular_sims" "example" {

  serial = "string"
  sim_failover = {

    enabled = true
    timeout = 300
  }
  sim_ordering = ["sim1", "sim2", "sim3"]
  sims = [{

    apns = [{

      allowed_ip_types = ["ipv4", "ipv6"]
      authentication = {

        password = "secret"
        type     = "pap"
        username = "milesmeraki"
      }
      name = "internet"
    }]
    is_primary = false
    sim_order  = 3
    slot       = "sim1"
  }]
}

output "meraki_devices_cellular_sims_example" {
  value = meraki_devices_cellular_sims.example
}