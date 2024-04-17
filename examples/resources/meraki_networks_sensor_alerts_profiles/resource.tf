
resource "meraki_networks_sensor_alerts_profiles" "example" {

  conditions = [{

    direction = "above"
    duration  = 60
    metric    = "temperature"
    threshold = {

      apparent_power = {

        draw = 17.2
      }
      current = {

        draw = 0.14
      }
      door = {

        open = true
      }
      frequency = {

        level = 58.8
      }
      humidity = {

        quality             = "inadequate"
        relative_percentage = 65
      }
      indoor_air_quality = {

        quality = "fair"
        score   = 80
      }
      noise = {

        ambient = {

          level   = 120
          quality = "poor"
        }
      }
      pm25 = {

        concentration = 90
        quality       = "fair"
      }
      power_factor = {

        percentage = 81
      }
      real_power = {

        draw = 14.1
      }
      temperature = {

        celsius    = 20.5
        fahrenheit = 70.0
        quality    = "good"
      }
      tvoc = {

        concentration = 400
        quality       = "poor"
      }
      upstream_power = {

        outage_detected = true
      }
      voltage = {

        level = 119.5
      }
      water = {

        present = true
      }
    }
  }]
  name       = "My Sensor Alert Profile"
  network_id = "string"
  recipients = {

    emails          = ["miles@meraki.com"]
    http_server_ids = ["aHR0cHM6Ly93d3cuZXhhbXBsZS5jb20vd2ViaG9va3M="]
    sms_numbers     = ["+15555555555"]
  }
  schedule = {

    id = "5"
  }
  serials = ["Q234-ABCD-0001", "Q234-ABCD-0002", "Q234-ABCD-0003"]
}

output "meraki_networks_sensor_alerts_profiles_example" {
  value = meraki_networks_sensor_alerts_profiles.example
}