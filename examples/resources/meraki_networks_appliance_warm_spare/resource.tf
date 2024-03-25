
resource "meraki_networks_appliance_warm_spare" "example" {

  enabled      = true
  network_id   = "string"
  spare_serial = "Q234-ABCD-5678"
  uplink_mode  = "virtual"
  virtual_ip1  = "1.2.3.4"
  virtual_ip2  = "1.2.3.4"
}

output "meraki_networks_appliance_warm_spare_example" {
  value = meraki_networks_appliance_warm_spare.example
}