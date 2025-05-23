---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_devices_camera_quality_and_retention Resource - terraform-provider-meraki"
subcategory: "camera"
description: |-
  
---

# meraki_devices_camera_quality_and_retention (Resource)



## Example Usage

```terraform
resource "meraki_devices_camera_quality_and_retention" "example" {

  audio_recording_enabled           = false
  motion_based_retention_enabled    = false
  motion_detector_version           = 2
  profile_id                        = "1234"
  quality                           = "Standard"
  resolution                        = "1280x720"
  restricted_bandwidth_mode_enabled = false
  serial                            = "string"
}

output "meraki_devices_camera_quality_and_retention_example" {
  value = meraki_devices_camera_quality_and_retention.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `serial` (String) serial path parameter.

### Optional

- `audio_recording_enabled` (Boolean) Boolean indicating if audio recording is enabled(true) or disabled(false) on the camera
- `motion_based_retention_enabled` (Boolean) Boolean indicating if motion-based retention is enabled(true) or disabled(false) on the camera.
- `motion_detector_version` (Number) The version of the motion detector that will be used by the camera. Only applies to Gen 2 cameras. Defaults to v2.
                                  Allowed values: [1,2]
- `profile_id` (String) The ID of a quality and retention profile to assign to the camera. The profile's settings will override all of the per-camera quality and retention settings. If the value of this parameter is null, any existing profile will be unassigned from the camera.
- `quality` (String) Quality of the camera. Can be one of 'Standard', 'High', 'Enhanced' or 'Ultra'. Not all qualities are supported by every camera model.
                                  Allowed values: [Enhanced,High,Standard,Ultra]
- `resolution` (String) Resolution of the camera. Can be one of '1280x720', '1920x1080', '1080x1080', '2112x2112', '2880x2880', '2688x1512' or '3840x2160'.Not all resolutions are supported by every camera model.
                                  Allowed values: [1080x1080,1280x720,1920x1080,2112x2112,2688x1512,2880x2880,3840x2160]
- `restricted_bandwidth_mode_enabled` (Boolean) Boolean indicating if restricted bandwidth is enabled(true) or disabled(false) on the camera. This setting does not apply to MV2 cameras.

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_devices_camera_quality_and_retention.example "serial"
```
