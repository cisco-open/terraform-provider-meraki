
resource "meraki_devices_blink_leds" "example" {

  serial = "string"
  parameters = {

    duration = 20
    duty     = 50
    period   = 160
  }
}

output "meraki_devices_blink_leds_example" {
  value = meraki_devices_blink_leds.example
}