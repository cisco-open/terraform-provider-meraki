terraform {
  required_providers {
    meraki = {
      version = "1.0.6-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}


resource "meraki_organizations_snmp" "example" {

  organization_id = "828099381482762270"
  peer_ips        = ["123.123.123.1"]
  v2c_enabled     = false
  v3_auth_mode    = "SHA"
  v3_auth_pass    = "password"
  v3_enabled      = true
  v3_priv_mode    = "AES128"
  v3_priv_pass    = "password"
}

output "meraki_organizations_snmp_example" {
  value = meraki_organizations_snmp.example
}