---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_floor_plans_auto_locate_jobs_batch Resource - terraform-provider-meraki"
subcategory: "networks"
description: |-
  
---

# meraki_networks_floor_plans_auto_locate_jobs_batch (Resource)





~>Warning: This resource does not represent a real-world entity in Meraki Dashboard, therefore changing or deleting this resource on its own has no immediate effect. Instead, it is a task part of a Meraki Dashboard workflow. It is executed in Meraki without any additional verification. It does not check if it was executed before or if a similar configuration or action 
already existed previously.


## Example Usage

```terraform
resource "meraki_networks_floor_plans_auto_locate_jobs_batch" "example" {

  network_id = "string"
  parameters = {

    jobs = [{

      floor_plan_id = "g_2176982374"
      refresh       = ["gnss", "ranging"]
      scheduled_at  = "2018-02-11T00:00:00Z"
    }]
  }
}

output "meraki_networks_floor_plans_auto_locate_jobs_batch_example" {
  value = meraki_networks_floor_plans_auto_locate_jobs_batch.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID
- `parameters` (Attributes) (see [below for nested schema](#nestedatt--parameters))

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `jobs` (Attributes Set) The list of auto locate jobs to be scheduled. Up to 100 jobs can be provided in a request. (see [below for nested schema](#nestedatt--parameters--jobs))

<a id="nestedatt--parameters--jobs"></a>
### Nested Schema for `parameters.jobs`

Optional:

- `floor_plan_id` (String) The ID of the floor plan to run auto locate for
- `refresh` (List of String) The types of location data that should be refreshed for this job. The list must either contain both 'gnss' and 'ranging' or be empty, as we currently only support refreshing both 'gnss' and 'ranging', or neither.
- `scheduled_at` (String) Timestamp in ISO8601 format which indicates when the auto locate job should be run. If omitted, the auto locate job will start immediately.



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `jobs` (Attributes Set) The newly created jobs (see [below for nested schema](#nestedatt--item--jobs))

<a id="nestedatt--item--jobs"></a>
### Nested Schema for `item.jobs`

Read-Only:

- `completed` (Attributes) Auto locate job progress information (see [below for nested schema](#nestedatt--item--jobs--completed))
- `errors` (Attributes Set) List of errors that occurred during a failed run of auto locate (see [below for nested schema](#nestedatt--item--jobs--errors))
- `floor_plan_id` (String) Floor plan ID
- `gnss` (Attributes) GNSS (e.g. GPS) status and progress information (see [below for nested schema](#nestedatt--item--jobs--gnss))
- `id` (String) Auto locate job ID
- `network_id` (String) Network ID
- `ranging` (Attributes) Ranging status and progress information (see [below for nested schema](#nestedatt--item--jobs--ranging))
- `scheduled_at` (String) Scheduled start time for auto locate job
- `status` (String) Auto locate job status. Possible values: 'scheduled', 'in progress', 'canceling', 'error', 'finished', 'published', 'canceled'

<a id="nestedatt--item--jobs--completed"></a>
### Nested Schema for `item.jobs.completed`

Read-Only:

- `percentage` (Number) Approximate auto locate job completion percentage


<a id="nestedatt--item--jobs--errors"></a>
### Nested Schema for `item.jobs.errors`

Read-Only:

- `source` (String) The step of the auto locate process when the error occurred. Possible values: 'gnss', 'ranging', 'positioning'
- `type` (String) The type of error that occurred. Possible values: 'failure', 'no neighbors', 'missing anchors', 'wrong anchors', 'missing ranging data', 'calculation failure', 'scheduling failure'


<a id="nestedatt--item--jobs--gnss"></a>
### Nested Schema for `item.jobs.gnss`

Read-Only:

- `completed` (Attributes) Progress information for the GNSS acquisition process (see [below for nested schema](#nestedatt--item--jobs--gnss--completed))
- `status` (String) GNSS status. Possible values: 'scheduled', 'in progress', 'error', 'finished', 'not applicable', 'canceled'

<a id="nestedatt--item--jobs--gnss--completed"></a>
### Nested Schema for `item.jobs.gnss.completed`

Read-Only:

- `percentage` (Number) Completion percentage of the GNSS acquisition process



<a id="nestedatt--item--jobs--ranging"></a>
### Nested Schema for `item.jobs.ranging`

Read-Only:

- `completed` (Attributes) Progress information for the ranging process (see [below for nested schema](#nestedatt--item--jobs--ranging--completed))
- `status` (String) Ranging status. Possible values: 'scheduled', 'in progress', 'error', 'finished', 'no neighbors'

<a id="nestedatt--item--jobs--ranging--completed"></a>
### Nested Schema for `item.jobs.ranging.completed`

Read-Only:

- `percentage` (Number) Completion percentage of the ranging process
