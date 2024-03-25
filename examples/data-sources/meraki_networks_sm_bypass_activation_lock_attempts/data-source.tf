
data "meraki_networks_sm_bypass_activation_lock_attempts" "example" {

  attempt_id = "string"
  network_id = "string"
}

output "meraki_networks_sm_bypass_activation_lock_attempts_example" {
  value = data.meraki_networks_sm_bypass_activation_lock_attempts.example.item
}
