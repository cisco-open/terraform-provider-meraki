
data "meraki_organizations_sm_apns_cert" "example" {

  organization_id = "string"
}

output "meraki_organizations_sm_apns_cert_example" {
  value = data.meraki_organizations_sm_apns_cert.example.item
}
