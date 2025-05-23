---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_api_requests Data Source - terraform-provider-meraki"
subcategory: "organizations"
description: |-
  
---

# meraki_organizations_api_requests (Data Source)



## Example Usage

```terraform
data "meraki_organizations_api_requests" "example" {

  admin_id        = "string"
  ending_before   = "string"
  method          = "string"
  operation_ids   = ["string"]
  organization_id = "string"
  path            = "string"
  per_page        = 1
  response_code   = 1
  source_ip       = "string"
  starting_after  = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
  user_agent      = "string"
  version         = 1
}

output "meraki_organizations_api_requests_example" {
  value = data.meraki_organizations_api_requests.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Optional

- `admin_id` (String) adminId query parameter. Filter the results by the ID of the admin who made the API requests
- `ending_before` (String) endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `method` (String) method query parameter. Filter the results by the method of the API requests (must be 'GET', 'PUT', 'POST' or 'DELETE')
- `operation_ids` (List of String) operationIds query parameter. Filter the results by one or more operation IDs for the API request
- `path` (String) path query parameter. Filter the results by the path of the API requests
- `per_page` (Number) perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 50.
- `response_code` (Number) responseCode query parameter. Filter the results by the response code of the API requests
- `source_ip` (String) sourceIp query parameter. Filter the results by the IP address of the originating API request
- `starting_after` (String) startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `t0` (String) t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
- `t1` (String) t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
- `timespan` (Number) timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days.
- `user_agent` (String) userAgent query parameter. Filter the results by the user agent string of the API request
- `version` (Number) version query parameter. Filter the results by the API version of the API request

### Read-Only

- `items` (Attributes List) Array of ResponseOrganizationsGetOrganizationApiRequests (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `admin_id` (String) Database ID for the admin user who made the API request.
- `client` (Attributes) Client information (see [below for nested schema](#nestedatt--items--client))
- `host` (String) The host which the API request was directed at.
- `method` (String) HTTP method used in the API request.
- `operation_id` (String) Operation ID for the endpoint.
- `path` (String) The API request path.
- `query_string` (String) The query string sent with the API request.
- `response_code` (Number) API request response code.
- `source_ip` (String) Public IP address from which the API request was made.
- `ts` (String) Timestamp, in iso8601 format, indicating when the API request was made.
- `user_agent` (String) The API request user agent.
- `version` (Number) API version of the endpoint.

<a id="nestedatt--items--client"></a>
### Nested Schema for `items.client`

Read-Only:

- `id` (String) ID for the client which made the request, if applicable.
- `type` (String) Type of client which made the request, if applicable. Available options are: oauth, api_key
