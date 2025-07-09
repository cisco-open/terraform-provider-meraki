terraform {
  required_providers {
    meraki = {
      version = "1.1.6-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug    = "true"
  meraki_base_url = "http://localhost:3002"
}

resource "meraki_organizations_saml_idps" "example" {

  organization_id           = "123456"
  slo_logout_url            = "https://login.remerge.io"
  x509cert_sha1_fingerprint = "80:63:7E:86:9A:90:99:30:DF:50:F2:CD:51:15:2D:67:81:BB:8E:6B"
}

output "meraki_organizations_saml_idps_example" {
  value = meraki_organizations_saml_idps.example
}