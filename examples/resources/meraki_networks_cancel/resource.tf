
resource "meraki_networks_cancel" "example" {

  job_id     = "string"
  network_id = "string"

}

output "meraki_networks_cancel_example" {
  value = meraki_networks_cancel.example
}