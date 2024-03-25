
resource "meraki_organizations_snmp" "example" {

  organization_id = "string"
  peer_ips        = ["123.123.123.1"]
  v2c_enabled     = false
  v3_auth_mode    = "SHA"
  v3_enabled      = true
  v3_priv_mode    = "AES128"
}

output "meraki_organizations_snmp_example" {
  value = meraki_organizations_snmp.example
}