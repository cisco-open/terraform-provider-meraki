---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_config_templates_switch_profiles_ports Data Source - terraform-provider-meraki"
subcategory: "switch"
description: |-
  
---

# meraki_organizations_config_templates_switch_profiles_ports (Data Source)



## Example Usage

```terraform
data "meraki_organizations_config_templates_switch_profiles_ports" "example" {

  config_template_id = "string"
  organization_id    = "string"
  profile_id         = "string"
}

output "meraki_organizations_config_templates_switch_profiles_ports_example" {
  value = data.meraki_organizations_config_templates_switch_profiles_ports.example.items
}

data "meraki_organizations_config_templates_switch_profiles_ports" "example" {

  config_template_id = "string"
  organization_id    = "string"
  profile_id         = "string"
}

output "meraki_organizations_config_templates_switch_profiles_ports_example" {
  value = data.meraki_organizations_config_templates_switch_profiles_ports.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `config_template_id` (String) configTemplateId path parameter. Config template ID
- `organization_id` (String) organizationId path parameter. Organization ID
- `port_id` (String) portId path parameter. Port ID
- `profile_id` (String) profileId path parameter. Profile ID

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))
- `items` (Attributes List) Array of ResponseSwitchGetOrganizationConfigTemplateSwitchProfilePorts (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `access_policy_number` (Number) The number of a custom access policy to configure on the switch template port. Only applicable when 'accessPolicyType' is 'Custom access policy'.
- `access_policy_type` (String) The type of the access policy of the switch template port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
- `allowed_vlans` (String) The VLANs allowed on the switch template port. Only applicable to trunk ports.
- `dai_trusted` (Boolean) If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
- `dot3az` (Attributes) dot3az settings for the port (see [below for nested schema](#nestedatt--item--dot3az))
- `enabled` (Boolean) The status of the switch template port.
- `flexible_stacking_enabled` (Boolean) For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
- `isolation_enabled` (Boolean) The isolation status of the switch template port.
- `link_negotiation` (String) The link speed for the switch template port.
- `link_negotiation_capabilities` (List of String) Available link speeds for the switch template port.
- `mac_allow_list` (List of String) Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.
- `mirror` (Attributes) Port mirror (see [below for nested schema](#nestedatt--item--mirror))
- `module` (Attributes) Expansion module (see [below for nested schema](#nestedatt--item--module))
- `name` (String) The name of the switch template port.
- `poe_enabled` (Boolean) The PoE status of the switch template port.
- `port_id` (String) The identifier of the switch template port.
- `port_schedule_id` (String) The ID of the port schedule. A value of null will clear the port schedule.
- `profile` (Attributes) Profile attributes (see [below for nested schema](#nestedatt--item--profile))
- `rstp_enabled` (Boolean) The rapid spanning tree protocol status.
- `schedule` (Attributes) The port schedule data. (see [below for nested schema](#nestedatt--item--schedule))
- `stackwise_virtual` (Attributes) Stackwise Virtual settings for the port (see [below for nested schema](#nestedatt--item--stackwise_virtual))
- `sticky_mac_allow_list` (List of String) The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
- `sticky_mac_allow_list_limit` (Number) The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
- `storm_control_enabled` (Boolean) The storm control status of the switch template port.
- `stp_guard` (String) The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
- `tags` (List of String) The list of tags of the switch template port.
- `type` (String) The type of the switch template port ('trunk', 'access', 'stack' or 'routed').
- `udld` (String) The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
- `vlan` (Number) The VLAN of the switch template port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.
- `voice_vlan` (Number) The voice VLAN of the switch template port. Only applicable to access ports.

<a id="nestedatt--item--dot3az"></a>
### Nested Schema for `item.dot3az`

Read-Only:

- `enabled` (Boolean) The Energy Efficient Ethernet status of the switch template port.


<a id="nestedatt--item--mirror"></a>
### Nested Schema for `item.mirror`

Read-Only:

- `mode` (String) The port mirror mode. Can be one of ('Destination port', 'Source port' or 'Not mirroring traffic').


<a id="nestedatt--item--module"></a>
### Nested Schema for `item.module`

Read-Only:

- `model` (String) The model of the expansion module.


<a id="nestedatt--item--profile"></a>
### Nested Schema for `item.profile`

Read-Only:

- `enabled` (Boolean) When enabled, override this port's configuration with a port profile.
- `id` (String) When enabled, the ID of the port profile used to override the port's configuration.
- `iname` (String) When enabled, the IName of the profile.


<a id="nestedatt--item--schedule"></a>
### Nested Schema for `item.schedule`

Read-Only:

- `id` (String) The ID of the port schedule.
- `name` (String) The name of the port schedule.


<a id="nestedatt--item--stackwise_virtual"></a>
### Nested Schema for `item.stackwise_virtual`

Read-Only:

- `is_dual_active_detector` (Boolean) For SVL devices, whether or not the port is used for Dual Active Detection.
- `is_stack_wise_virtual_link` (Boolean) For SVL devices, whether or not the port is used for StackWise Virtual Link.



<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `access_policy_number` (Number) The number of a custom access policy to configure on the switch template port. Only applicable when 'accessPolicyType' is 'Custom access policy'.
- `access_policy_type` (String) The type of the access policy of the switch template port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
- `allowed_vlans` (String) The VLANs allowed on the switch template port. Only applicable to trunk ports.
- `dai_trusted` (Boolean) If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
- `dot3az` (Attributes) dot3az settings for the port (see [below for nested schema](#nestedatt--items--dot3az))
- `enabled` (Boolean) The status of the switch template port.
- `flexible_stacking_enabled` (Boolean) For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
- `isolation_enabled` (Boolean) The isolation status of the switch template port.
- `link_negotiation` (String) The link speed for the switch template port.
- `link_negotiation_capabilities` (List of String) Available link speeds for the switch template port.
- `mac_allow_list` (List of String) Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.
- `mirror` (Attributes) Port mirror (see [below for nested schema](#nestedatt--items--mirror))
- `module` (Attributes) Expansion module (see [below for nested schema](#nestedatt--items--module))
- `name` (String) The name of the switch template port.
- `poe_enabled` (Boolean) The PoE status of the switch template port.
- `port_id` (String) The identifier of the switch template port.
- `port_schedule_id` (String) The ID of the port schedule. A value of null will clear the port schedule.
- `profile` (Attributes) Profile attributes (see [below for nested schema](#nestedatt--items--profile))
- `rstp_enabled` (Boolean) The rapid spanning tree protocol status.
- `schedule` (Attributes) The port schedule data. (see [below for nested schema](#nestedatt--items--schedule))
- `stackwise_virtual` (Attributes) Stackwise Virtual settings for the port (see [below for nested schema](#nestedatt--items--stackwise_virtual))
- `sticky_mac_allow_list` (List of String) The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
- `sticky_mac_allow_list_limit` (Number) The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
- `storm_control_enabled` (Boolean) The storm control status of the switch template port.
- `stp_guard` (String) The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
- `tags` (List of String) The list of tags of the switch template port.
- `type` (String) The type of the switch template port ('trunk', 'access', 'stack' or 'routed').
- `udld` (String) The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
- `vlan` (Number) The VLAN of the switch template port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.
- `voice_vlan` (Number) The voice VLAN of the switch template port. Only applicable to access ports.

<a id="nestedatt--items--dot3az"></a>
### Nested Schema for `items.dot3az`

Read-Only:

- `enabled` (Boolean) The Energy Efficient Ethernet status of the switch template port.


<a id="nestedatt--items--mirror"></a>
### Nested Schema for `items.mirror`

Read-Only:

- `mode` (String) The port mirror mode. Can be one of ('Destination port', 'Source port' or 'Not mirroring traffic').


<a id="nestedatt--items--module"></a>
### Nested Schema for `items.module`

Read-Only:

- `model` (String) The model of the expansion module.


<a id="nestedatt--items--profile"></a>
### Nested Schema for `items.profile`

Read-Only:

- `enabled` (Boolean) When enabled, override this port's configuration with a port profile.
- `id` (String) When enabled, the ID of the port profile used to override the port's configuration.
- `iname` (String) When enabled, the IName of the profile.


<a id="nestedatt--items--schedule"></a>
### Nested Schema for `items.schedule`

Read-Only:

- `id` (String) The ID of the port schedule.
- `name` (String) The name of the port schedule.


<a id="nestedatt--items--stackwise_virtual"></a>
### Nested Schema for `items.stackwise_virtual`

Read-Only:

- `is_dual_active_detector` (Boolean) For SVL devices, whether or not the port is used for Dual Active Detection.
- `is_stack_wise_virtual_link` (Boolean) For SVL devices, whether or not the port is used for StackWise Virtual Link.
