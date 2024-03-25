
data "meraki_organizations_devices_uplinks_loss_and_latency" "example" {

  ip              = "string"
  organization_id = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
  uplink          = "string"
}

output "meraki_organizations_devices_uplinks_loss_and_latency_example" {
  value = data.meraki_organizations_devices_uplinks_loss_and_latency.example.items
}
