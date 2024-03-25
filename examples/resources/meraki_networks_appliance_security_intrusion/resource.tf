
resource "meraki_networks_appliance_security_intrusion" "example" {

  ids_rulesets = "balanced"
  mode         = "prevention"
  network_id   = "string"
  protected_networks = {

    excluded_cidr = ["10.0.0.0/8", "127.0.0.0/8"]
    included_cidr = ["10.0.0.0/8", "127.0.0.0/8", "169.254.0.0/16", "172.16.0.0/12"]
    use_default   = false
  }
}

output "meraki_networks_appliance_security_intrusion_example" {
  value = meraki_networks_appliance_security_intrusion.example
}