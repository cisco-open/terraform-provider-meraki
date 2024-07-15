terraform {
  required_providers {
    meraki = {
      version = "0.2.6-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_sensor_alerts_profiles" "this" {

  conditions = [
    {
      metric = "temperature"
      duration = 60
      threshold = {
        temperature = {
          quality = "good"
        }
      },
    },
    {
      metric = "humidity"
      duration = 60
      threshold = {
        humidity = {
          quality = "good"
        }
      },
    },
    {
      metric = "water"
      threshold = {
        water = {
          present = true
        }
      }
    }
  ]
  name       = "Sensor Alerts2"
  network_id = "L_828099381482775342"
  recipients = {
    http_server_ids = ["aHR0cHM6Ly8xLjIuMy40"]
  }
  # serials = ["Q2FV-VYGH-ZVB3", "QBSA-D8CD-5LR6"]
}


