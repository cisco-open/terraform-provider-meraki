---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_sensor_alerts_profiles Resource - terraform-provider-meraki"
subcategory: "sensor"
description: |-
  
---

# meraki_networks_sensor_alerts_profiles (Resource)



## Example Usage

```terraform
resource "meraki_networks_sensor_alerts_profiles" "example" {

  conditions = [{

    direction = "above"
    duration  = 60
    metric    = "temperature"
    threshold = {

      apparent_power = {

        draw = 17.2
      }
      co2 = {

        concentration = 400
        quality       = "poor"
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
  includesensor_url = true
  message           = "Check with Miles on what to do."
  name              = "My Sensor Alert Profile"
  network_id        = "string"
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Optional

- `conditions` (Attributes Set) List of conditions that will cause the profile to send an alert. (see [below for nested schema](#nestedatt--conditions))
- `id` (String) id path parameter.
- `include_sensor_url` (Boolean) Include dashboard link to sensor in messages (default: true).
- `message` (String) A custom message that will appear in email and text message alerts.
- `name` (String) Name of the sensor alert profile.
- `recipients` (Attributes) List of recipients that will receive the alert. (see [below for nested schema](#nestedatt--recipients))
- `schedule` (Attributes) The sensor schedule to use with the alert profile. (see [below for nested schema](#nestedatt--schedule))
- `serials` (Set of String) List of device serials assigned to this sensor alert profile.

### Read-Only

- `conditions_response` (Attributes Set) List of conditions that will cause the profile to send an alert. (see [below for nested schema](#nestedatt--conditions_response))
- `profile_id` (String) ID of the sensor alert profile.

<a id="nestedatt--conditions"></a>
### Nested Schema for `conditions`

Optional:

- `direction` (String) If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature, humidity, realPower, apparentPower, powerFactor, voltage, current, and frequency thresholds.
                            Allowed values: [above,below]
- `duration` (Number) Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, 1 hour, 2 hours, 4 hours, and 8 hours. Default is 0.
                            Allowed values: [0,60,120,180,240,300,600,900,1800,3600,7200,14400,28800]
- `metric` (String) The type of sensor metric that will be monitored for changes.
                            Allowed values: [apparentPower,co2,current,door,frequency,humidity,indoorAirQuality,noise,pm25,powerFactor,realPower,temperature,tvoc,upstreamPower,voltage,water]
- `threshold` (Attributes) Threshold for sensor readings that will cause an alert to be sent. This object should contain a single property key matching the condition's 'metric' value. (see [below for nested schema](#nestedatt--conditions--threshold))

<a id="nestedatt--conditions--threshold"></a>
### Nested Schema for `conditions.threshold`

Optional:

- `apparent_power` (Attributes) Apparent power threshold. 'draw' must be provided. (see [below for nested schema](#nestedatt--conditions--threshold--apparent_power))
- `co2` (Attributes) CO2 concentration threshold. One of 'concentration' or 'quality' must be provided. (see [below for nested schema](#nestedatt--conditions--threshold--co2))
- `current` (Attributes) Electrical current threshold. 'level' must be provided. (see [below for nested schema](#nestedatt--conditions--threshold--current))
- `door` (Attributes) Door open threshold. 'open' must be provided and set to true. (see [below for nested schema](#nestedatt--conditions--threshold--door))
- `frequency` (Attributes) Electrical frequency threshold. 'level' must be provided. (see [below for nested schema](#nestedatt--conditions--threshold--frequency))
- `humidity` (Attributes) Humidity threshold. One of 'relativePercentage' or 'quality' must be provided. (see [below for nested schema](#nestedatt--conditions--threshold--humidity))
- `indoor_air_quality` (Attributes) Indoor air quality score threshold. One of 'score' or 'quality' must be provided. (see [below for nested schema](#nestedatt--conditions--threshold--indoor_air_quality))
- `noise` (Attributes) Noise threshold. 'ambient' must be provided. (see [below for nested schema](#nestedatt--conditions--threshold--noise))
- `pm25` (Attributes) PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided. (see [below for nested schema](#nestedatt--conditions--threshold--pm25))
- `power_factor` (Attributes) Power factor threshold. 'percentage' must be provided. (see [below for nested schema](#nestedatt--conditions--threshold--power_factor))
- `real_power` (Attributes) Real power threshold. 'draw' must be provided. (see [below for nested schema](#nestedatt--conditions--threshold--real_power))
- `temperature` (Attributes) Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided. (see [below for nested schema](#nestedatt--conditions--threshold--temperature))
- `tvoc` (Attributes) TVOC concentration threshold. One of 'concentration' or 'quality' must be provided. (see [below for nested schema](#nestedatt--conditions--threshold--tvoc))
- `upstream_power` (Attributes) Upstream power threshold. 'outageDetected' must be provided and set to true. (see [below for nested schema](#nestedatt--conditions--threshold--upstream_power))
- `voltage` (Attributes) Voltage threshold. 'level' must be provided. (see [below for nested schema](#nestedatt--conditions--threshold--voltage))
- `water` (Attributes) Water detection threshold. 'present' must be provided and set to true. (see [below for nested schema](#nestedatt--conditions--threshold--water))

<a id="nestedatt--conditions--threshold--apparent_power"></a>
### Nested Schema for `conditions.threshold.apparent_power`

Optional:

- `draw` (Number) Alerting threshold in volt-amps. Must be between 0 and 3750.


<a id="nestedatt--conditions--threshold--co2"></a>
### Nested Schema for `conditions.threshold.co2`

Optional:

- `concentration` (Number) Alerting threshold as CO2 parts per million.
- `quality` (String) Alerting threshold as a qualitative CO2 level.
                                        Allowed values: [fair,good,inadequate,poor]


<a id="nestedatt--conditions--threshold--current"></a>
### Nested Schema for `conditions.threshold.current`

Optional:

- `draw` (Number) Alerting threshold in amps. Must be between 0 and 15.


<a id="nestedatt--conditions--threshold--door"></a>
### Nested Schema for `conditions.threshold.door`

Optional:

- `open` (Boolean) Alerting threshold for a door open event. Must be set to true.


<a id="nestedatt--conditions--threshold--frequency"></a>
### Nested Schema for `conditions.threshold.frequency`

Optional:

- `level` (Number) Alerting threshold in hertz. Must be between 0 and 60.


<a id="nestedatt--conditions--threshold--humidity"></a>
### Nested Schema for `conditions.threshold.humidity`

Optional:

- `quality` (String) Alerting threshold as a qualitative humidity level.
                                        Allowed values: [fair,good,inadequate,poor]
- `relative_percentage` (Number) Alerting threshold in %RH.


<a id="nestedatt--conditions--threshold--indoor_air_quality"></a>
### Nested Schema for `conditions.threshold.indoor_air_quality`

Optional:

- `quality` (String) Alerting threshold as a qualitative indoor air quality level.
                                        Allowed values: [fair,good,inadequate,poor]
- `score` (Number) Alerting threshold as indoor air quality score.


<a id="nestedatt--conditions--threshold--noise"></a>
### Nested Schema for `conditions.threshold.noise`

Optional:

- `ambient` (Attributes) Ambient noise threshold. One of 'level' or 'quality' must be provided. (see [below for nested schema](#nestedatt--conditions--threshold--noise--ambient))

<a id="nestedatt--conditions--threshold--noise--ambient"></a>
### Nested Schema for `conditions.threshold.noise.ambient`

Optional:

- `level` (Number) Alerting threshold as adjusted decibels.
- `quality` (String) Alerting threshold as a qualitative ambient noise level.
                                              Allowed values: [fair,good,inadequate,poor]



<a id="nestedatt--conditions--threshold--pm25"></a>
### Nested Schema for `conditions.threshold.pm25`

Optional:

- `concentration` (Number) Alerting threshold as PM2.5 parts per million.
- `quality` (String) Alerting threshold as a qualitative PM2.5 level.
                                        Allowed values: [fair,good,inadequate,poor]


<a id="nestedatt--conditions--threshold--power_factor"></a>
### Nested Schema for `conditions.threshold.power_factor`

Optional:

- `percentage` (Number) Alerting threshold as the ratio of active power to apparent power. Must be between 0 and 100.


<a id="nestedatt--conditions--threshold--real_power"></a>
### Nested Schema for `conditions.threshold.real_power`

Optional:

- `draw` (Number) Alerting threshold in watts. Must be between 0 and 3750.


<a id="nestedatt--conditions--threshold--temperature"></a>
### Nested Schema for `conditions.threshold.temperature`

Optional:

- `celsius` (Number) Alerting threshold in degrees Celsius.
- `fahrenheit` (Number) Alerting threshold in degrees Fahrenheit.
- `quality` (String) Alerting threshold as a qualitative temperature level.
                                        Allowed values: [fair,good,inadequate,poor]


<a id="nestedatt--conditions--threshold--tvoc"></a>
### Nested Schema for `conditions.threshold.tvoc`

Optional:

- `concentration` (Number) Alerting threshold as TVOC micrograms per cubic meter.
- `quality` (String) Alerting threshold as a qualitative TVOC level.
                                        Allowed values: [fair,good,inadequate,poor]


<a id="nestedatt--conditions--threshold--upstream_power"></a>
### Nested Schema for `conditions.threshold.upstream_power`

Optional:

- `outage_detected` (Boolean) Alerting threshold for an upstream power event. Must be set to true.


<a id="nestedatt--conditions--threshold--voltage"></a>
### Nested Schema for `conditions.threshold.voltage`

Optional:

- `level` (Number) Alerting threshold in volts. Must be between 0 and 250.


<a id="nestedatt--conditions--threshold--water"></a>
### Nested Schema for `conditions.threshold.water`

Optional:

- `present` (Boolean) Alerting threshold for a water detection event. Must be set to true.




<a id="nestedatt--recipients"></a>
### Nested Schema for `recipients`

Optional:

- `emails` (Set of String) A list of emails that will receive information about the alert.
- `http_server_ids` (Set of String) A list of webhook endpoint IDs that will receive information about the alert.
- `sms_numbers` (Set of String) A list of SMS numbers that will receive information about the alert.


<a id="nestedatt--schedule"></a>
### Nested Schema for `schedule`

Optional:

- `id` (String) ID of the sensor schedule to use with the alert profile. If not defined, the alert profile will be active at all times.

Read-Only:

- `name` (String) Name of the sensor schedule to use with the alert profile.


<a id="nestedatt--conditions_response"></a>
### Nested Schema for `conditions_response`

Read-Only:

- `direction` (String) If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature, humidity, realPower, apparentPower, powerFactor, voltage, current, and frequency thresholds.
- `duration` (Number) Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, 1 hour, 2 hours, 4 hours, and 8 hours. Default is 0.
- `metric` (String) The type of sensor metric that will be monitored for changes. Available metrics are apparentPower, co2, current, door, frequency, humidity, indoorAirQuality, noise, pm25, powerFactor, realPower, temperature, tvoc, upstreamPower, voltage, and water.
- `threshold` (Attributes) Threshold for sensor readings that will cause an alert to be sent. This object should contain a single property key matching the condition's 'metric' value. (see [below for nested schema](#nestedatt--conditions_response--threshold))

<a id="nestedatt--conditions_response--threshold"></a>
### Nested Schema for `conditions_response.threshold`

Read-Only:

- `apparent_power` (Attributes) Apparent power threshold. 'draw' must be provided. (see [below for nested schema](#nestedatt--conditions_response--threshold--apparent_power))
- `co2` (Attributes) CO2 concentration threshold. One of 'concentration' or 'quality' must be provided. (see [below for nested schema](#nestedatt--conditions_response--threshold--co2))
- `current` (Attributes) Electrical current threshold. 'level' must be provided. (see [below for nested schema](#nestedatt--conditions_response--threshold--current))
- `door` (Attributes) Door open threshold. 'open' must be provided and set to true. (see [below for nested schema](#nestedatt--conditions_response--threshold--door))
- `frequency` (Attributes) Electrical frequency threshold. 'level' must be provided. (see [below for nested schema](#nestedatt--conditions_response--threshold--frequency))
- `humidity` (Attributes) Humidity threshold. One of 'relativePercentage' or 'quality' must be provided. (see [below for nested schema](#nestedatt--conditions_response--threshold--humidity))
- `indoor_air_quality` (Attributes) Indoor air quality score threshold. One of 'score' or 'quality' must be provided. (see [below for nested schema](#nestedatt--conditions_response--threshold--indoor_air_quality))
- `noise` (Attributes) Noise threshold. 'ambient' must be provided. (see [below for nested schema](#nestedatt--conditions_response--threshold--noise))
- `pm25` (Attributes) PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided. (see [below for nested schema](#nestedatt--conditions_response--threshold--pm25))
- `power_factor` (Attributes) Power factor threshold. 'percentage' must be provided. (see [below for nested schema](#nestedatt--conditions_response--threshold--power_factor))
- `real_power` (Attributes) Real power threshold. 'draw' must be provided. (see [below for nested schema](#nestedatt--conditions_response--threshold--real_power))
- `temperature` (Attributes) Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided. (see [below for nested schema](#nestedatt--conditions_response--threshold--temperature))
- `tvoc` (Attributes) TVOC concentration threshold. One of 'concentration' or 'quality' must be provided. (see [below for nested schema](#nestedatt--conditions_response--threshold--tvoc))
- `upstream_power` (Attributes) Upstream power threshold. 'outageDetected' must be provided and set to true. (see [below for nested schema](#nestedatt--conditions_response--threshold--upstream_power))
- `voltage` (Attributes) Voltage threshold. 'level' must be provided. (see [below for nested schema](#nestedatt--conditions_response--threshold--voltage))
- `water` (Attributes) Water detection threshold. 'present' must be provided and set to true. (see [below for nested schema](#nestedatt--conditions_response--threshold--water))

<a id="nestedatt--conditions_response--threshold--apparent_power"></a>
### Nested Schema for `conditions_response.threshold.apparent_power`

Read-Only:

- `draw` (Number) Alerting threshold in volt-amps. Must be between 0 and 3750.


<a id="nestedatt--conditions_response--threshold--co2"></a>
### Nested Schema for `conditions_response.threshold.co2`

Read-Only:

- `concentration` (Number) Alerting threshold as CO2 parts per million.
- `quality` (String) Alerting threshold as a qualitative CO2 level.
                                                    Allowed values: [fair,good,inadequate,poor]


<a id="nestedatt--conditions_response--threshold--current"></a>
### Nested Schema for `conditions_response.threshold.current`

Read-Only:

- `draw` (Number) Alerting threshold in amps. Must be between 0 and 15.


<a id="nestedatt--conditions_response--threshold--door"></a>
### Nested Schema for `conditions_response.threshold.door`

Read-Only:

- `open` (Boolean) Alerting threshold for a door open event. Must be set to true.


<a id="nestedatt--conditions_response--threshold--frequency"></a>
### Nested Schema for `conditions_response.threshold.frequency`

Read-Only:

- `level` (Number) Alerting threshold in hertz. Must be between 0 and 60.


<a id="nestedatt--conditions_response--threshold--humidity"></a>
### Nested Schema for `conditions_response.threshold.humidity`

Read-Only:

- `quality` (String) Alerting threshold as a qualitative humidity level.
- `relative_percentage` (Number) Alerting threshold in %RH.


<a id="nestedatt--conditions_response--threshold--indoor_air_quality"></a>
### Nested Schema for `conditions_response.threshold.indoor_air_quality`

Read-Only:

- `quality` (String) Alerting threshold as a qualitative indoor air quality level.
- `score` (Number) Alerting threshold as indoor air quality score.


<a id="nestedatt--conditions_response--threshold--noise"></a>
### Nested Schema for `conditions_response.threshold.noise`

Read-Only:

- `ambient` (Attributes) Ambient noise threshold. One of 'level' or 'quality' must be provided. (see [below for nested schema](#nestedatt--conditions_response--threshold--noise--ambient))

<a id="nestedatt--conditions_response--threshold--noise--ambient"></a>
### Nested Schema for `conditions_response.threshold.noise.ambient`

Read-Only:

- `level` (Number) Alerting threshold as adjusted decibels.
- `quality` (String) Alerting threshold as a qualitative ambient noise level.



<a id="nestedatt--conditions_response--threshold--pm25"></a>
### Nested Schema for `conditions_response.threshold.pm25`

Read-Only:

- `concentration` (Number) Alerting threshold as PM2.5 parts per million.
- `quality` (String) Alerting threshold as a qualitative PM2.5 level.


<a id="nestedatt--conditions_response--threshold--power_factor"></a>
### Nested Schema for `conditions_response.threshold.power_factor`

Read-Only:

- `percentage` (Number) Alerting threshold as the ratio of active power to apparent power. Must be between 0 and 100.


<a id="nestedatt--conditions_response--threshold--real_power"></a>
### Nested Schema for `conditions_response.threshold.real_power`

Read-Only:

- `draw` (Number) Alerting threshold in watts. Must be between 0 and 3750.


<a id="nestedatt--conditions_response--threshold--temperature"></a>
### Nested Schema for `conditions_response.threshold.temperature`

Read-Only:

- `celsius` (Number) Alerting threshold in degrees Celsius.
- `fahrenheit` (Number) Alerting threshold in degrees Fahrenheit.
- `quality` (String) Alerting threshold as a qualitative temperature level.


<a id="nestedatt--conditions_response--threshold--tvoc"></a>
### Nested Schema for `conditions_response.threshold.tvoc`

Read-Only:

- `concentration` (Number) Alerting threshold as TVOC micrograms per cubic meter.
- `quality` (String) Alerting threshold as a qualitative TVOC level.


<a id="nestedatt--conditions_response--threshold--upstream_power"></a>
### Nested Schema for `conditions_response.threshold.upstream_power`

Read-Only:

- `outage_detected` (Boolean) Alerting threshold for an upstream power event. Must be set to true.


<a id="nestedatt--conditions_response--threshold--voltage"></a>
### Nested Schema for `conditions_response.threshold.voltage`

Read-Only:

- `level` (Number) Alerting threshold in volts. Must be between 0 and 250.


<a id="nestedatt--conditions_response--threshold--water"></a>
### Nested Schema for `conditions_response.threshold.water`

Read-Only:

- `present` (Boolean) Alerting threshold for a water detection event. Must be set to true.

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_networks_sensor_alerts_profiles.example "id,network_id"
```
