
resource "meraki_organizations_assurance_alerts_dismiss" "example" {

  organization_id = "string"
  parameters = {

    alert_ids = ["1234", "4567"]
  }
}

output "meraki_organizations_assurance_alerts_dismiss_example" {
  value = meraki_organizations_assurance_alerts_dismiss.example
}