
resource "meraki_organizations_saml_idps" "example" {

  organization_id           = "string"
  slo_logout_url            = "https://somewhere.com"
  x509cert_sha1_fingerprint = "00:11:22:33:44:55:66:77:88:99:00:11:22:33:44:55:66:77:88:99"
}

output "meraki_organizations_saml_idps_example" {
  value = meraki_organizations_saml_idps.example
}