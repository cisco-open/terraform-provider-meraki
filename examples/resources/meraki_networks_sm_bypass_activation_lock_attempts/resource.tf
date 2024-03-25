
resource "meraki_networks_sm_bypass_activation_lock_attempts" "example" {

  ids        = ["1284392014819", "2983092129865"]
  network_id = "string"
}

output "meraki_networks_sm_bypass_activation_lock_attempts_example" {
  value = meraki_networks_sm_bypass_activation_lock_attempts.example
}