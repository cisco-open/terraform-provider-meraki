
resource "meraki_organizations_adaptive_policy_acls" "example" {

  description     = "Blocks sensitive web traffic"
  ip_version      = "ipv6"
  name            = "Block sensitive web traffic"
  organization_id = "string"
  rules = [{

    dst_port        = "22-30"
    log             = true
    policy          = "deny"
    protocol        = "tcp"
    src_port        = "1,33"
    tcp_established = true
  }]
}

output "meraki_organizations_adaptive_policy_acls_example" {
  value = meraki_organizations_adaptive_policy_acls.example
}