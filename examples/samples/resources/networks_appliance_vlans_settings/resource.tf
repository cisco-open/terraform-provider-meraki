terraform {
  required_providers {
    meraki = {
      version = "0.2.8-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}


resource "meraki_networks_appliance_vlans_settings" "example" {

  network_id    = "L_828099381482771185"
  vlans_enabled = true
}

output "meraki_networks_appliance_vlans_settings_example" {
  value = resource.meraki_networks_appliance_vlans_settings.example
}

resource "meraki_networks_appliance_vlans" "example" {

  network_id   = "L_828099381482771185"
  appliance_ip = "192.168.2.2"
  # cidr         = "192.168.2.1/24"
  name   = "VLAN Terraform"
  subnet = "192.168.2.0/24"
  id = 1001
  depends_on = [meraki_networks_appliance_vlans_settings.example]
}

output "meraki_networks_appliance_vlans_example" {
  value = resource.meraki_networks_appliance_vlans.example
}


resource "meraki_networks_appliance_ssids" "example" {

  auth_mode       = "8021x-radius"
  default_vlan_id = 1001
  enabled         = true
  #encryption_mode = "wep"
  name       = "My SSID 2"
  network_id = "L_828099381482771185"
  number     = 1
  radius_servers = [
    {
      host   = "1.2.3.4"
      port   = 1000
      secret = "secret"
    },
    {
      host   = "1.2.3.5"
      port   = 1002
      secret = "secret2"
    },
    {
      host   = "1.2.3.6"
      port   = 1002
      secret = "secret3"
    },
  ]
  visible             = true
  wpa_encryption_mode = "WPA2 only"
  depends_on          = [meraki_networks_appliance_vlans_settings.example, meraki_networks_appliance_vlans.example]
  lifecycle {
    ignore_changes = [radius_servers]
  }
}

output "meraki_networks_appliance_ssids_example" {
  value = meraki_networks_appliance_ssids.example
}

resource "meraki_networks_appliance_traffic_shaping" "example" {


  global_bandwidth_limits = {

    limit_down = 5121
    limit_up   = 2048
  }
  network_id = "L_828099381482771185"
  depends_on = [meraki_networks_appliance_ssids.example]

}

resource "meraki_networks_appliance_firewall_l3_firewall_rules" "example" {



  network_id = "L_828099381482771185"
  rules = [{

    comment        = "Block internetbadguys.com"
    dest_cidr      = "internetbadguys.com"
    dest_port      = "Any"
    policy         = "deny"
    protocol       = "any"
    src_cidr       = "Any"
    src_port       = "Any"
    syslog_enabled = false
    }, {

    comment        = "Block internetbadguys3.com"
    dest_cidr      = "internetbadguys3.com"
    dest_port      = "Any"
    policy         = "deny"
    protocol       = "any"
    src_cidr       = "Any"
    src_port       = "Any"
    syslog_enabled = false
    },
    {
      comment        = "Default rule"
      dest_cidr      = "Any"
      dest_port      = "Any"
      policy         = "allow"
      protocol       = "any"
      src_cidr       = "Any"
      src_port       = "Any"
      syslog_enabled = true
  }, ]
  depends_on = [meraki_networks_appliance_traffic_shaping.example]
  lifecycle {
    ignore_changes = [rules]
  }
}

output "meraki_networks_appliance_firewall_l3_firewall_rules_example" {
  value = meraki_networks_appliance_firewall_l3_firewall_rules.example
}

# resource "meraki_networks_appliance_firewall_l7_firewall_rules" "example" {
#     




#       network_id = "L_828099381482771185"
#       rules == [{

#         policy = "deny"
#         type = "applicationCategory"
#         value = "applicationCategory"
#       }]
#     depends_on = [ meraki_networks_appliance_firewall_l3_firewall_rules.example ]
# }

# output "meraki_networks_appliance_firewall_l7_firewall_rules_example" {
#     value = meraki_networks_appliance_firewall_l7_firewall_rules.example
# }

# resource "meraki_networks_appliance_ports" "example" {
#       
#     # allowed_vlans = "a"
#     # drop_untagged_traffic=false
#     enabled = true
#     network_id = "L_828099381482771185"
#     port_id = "4"
#     type = "access"
#     depends_on = [ meraki_networks_appliance_traffic_shaping.example ]
# }