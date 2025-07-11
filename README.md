# terraform-provider-meraki

terraform-provider-meraki is a Terraform Provider for [Cisco Meraki]()

<img src="https://upload.wikimedia.org/wikipedia/commons/0/04/Terraform_Logo.svg" width="400px">

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.13.x
- [Go](https://golang.org/doc/install) 1.21 (to build the provider plugin)

## Introduction

The terraform-provider-meraki provides a Terraform provider for managing and automating your Cisco Meraki environment. It consists of a set of resources and data-sources for performing tasks related to Meraki.

This collection has been tested and supports Cisco Meraki 1.33.0.

## Using the provider

There are two ways to get and use the provider.
1. Downloading & installing it from registry.terraform.io
2. Building it from source

### From registry

To install this provider, copy and paste this code into your Terraform configuration. Then, run terraform init. 

```hcl
terraform {
  required_providers {
    meraki = {
      source = "cisco-open/meraki"
      version = "1.1.6-beta"
    }
  }
}

provider "meraki" {
  # Configuration options
  # More info at https://registry.terraform.io/providers/cisco-open/meraki/latest/docs#example-usage
}
```

### From build

Clone this repository to: `$GOPATH/src/github.com/cisco-open/terraform-provider-meraki`

```sh
mkdir -p $GOPATH/src/github.com/meraki/
cd $GOPATH/src/github.com/meraki/
git clone https://github.com/cisco-open/terraform-provider-meraki.git
```

Enter the provider directory and build the provider

```sh
cd $GOPATH/src/github.com/cisco-open/terraform-provider-meraki
make build
```

If the Makefile values (HOSTNAME, NAMESPACE, NAME, VERSION) were not changed, then the following code could used without changes.
Otherwise change the values accordingly.

To use this provider, copy and paste this code into your Terraform configuration. Then, run terraform init.

```hcl
terraform {
  required_providers {
    meraki = {
      source = "hashicorp.com/edu/meraki"
      version = "1.1.6-beta"
    }
  }
}

provider "meraki" {
  # Configuration options
  # More info at https://registry.terraform.io/providers/cisco-open/meraki/latest/docs#example-usage
}
```


## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed
on your machine (version 1.15+ is _required_). You'll also need to correctly setup a
[GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-meraki
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

_Note:_ Acceptance tests create real resources.

```sh
$ make testacc
```

## Documentation

In the docs directory, you can find the documentation.

## Compatibility matrix
The following table shows the supported versions.

| Dashboard Api version | Terraform "meraki" provider version | Go "dashboard-api-go" version|
|-----------------------|-------------------------------------|------------------------------|
| 1.33.0                | 0.1.0-alpha                         | 2.0.9                        |
| 1.44.1                | 0.2.0-alpha                         | 3.0.0                        |
| 1.53.0                | 1.1.6-beta                          | 4.0.0                        |

If your SDK, Terraform provider is older please consider updating it first.

## Fetch All Items of an Endpoint with Pagination

- **Support for fetching all items with `per_page=-1`**  
  A new feature has been introduced to the API endpoints, enabling clients to fetch all available items in a single request by setting the `per_page` parameter to `-1`. This enhancement allows you to retrieve the full dataset without needing to make multiple paginated requests.

### Behavior

- When `per_page` is set to `-1`, the server will return **all available items** for that endpoint, bypassing the pagination logic.
- If a positive integer is passed for `per_page`, the endpoint will continue using traditional pagination and return only the number of items specified by `per_page`.

## Provider Configuration for Retry Options

> **Note:** Configuration via environment variables (`MERAKI_RETRIES`, `MERAKI_RETRY_DELAY`, etc.) is now deprecated. All configuration must be set directly in the `provider` block of your `.tf` file.

You can customize retry options as follows:

```hcl
provider "meraki" {
  max_retries      = 3      # Maximum number of retries
  retry_delay      = 1000   # Base wait time between retries in ms
  retry_jitter     = 3000   # Maximum random jitter in ms
  use_retry_header = false  # Whether to respect the Retry-After header
  # ...other configuration parameters...
}
```

## Documentation

In the docs directory, you can find the documentation.

## Compatibility matrix
The following table shows the supported versions.

| Dashboard Api version | Terraform "meraki" provider version | Go "dashboard-api-go" version|
|-----------------------|-------------------------------------|------------------------------|
| 1.33.0                | 0.1.0-alpha                         | 2.0.9                        |
| 1.44.1                | 0.2.0-alpha                         | 3.0.0                        |
| 1.53.0                | 1.1.6-beta                          | 4.0.0                        |

If your SDK, Terraform provider is older please consider updating it first.

## Fetch All Items of an Endpoint with Pagination

- **Support for fetching all items with `per_page=-1`**  
  A new feature has been introduced to the API endpoints, enabling clients to fetch all available items in a single request by setting the `per_page` parameter to `-1`. This enhancement allows you to retrieve the full dataset without needing to make multiple paginated requests.

### Behavior

- When `per_page` is set to `-1`, the server will return **all available items** for that endpoint, bypassing the pagination logic.
- If a positive integer is passed for `per_page`, the endpoint will continue using traditional pagination and return only the number of items specified by `per_page`.

# Contributing

Ongoing development efforts and contributions to this provider are tracked as issues in this repository.

We welcome community contributions to this project. If you find problems, need an enhancement or need a new data-source or resource, please open an issue or create a PR against the [Terraform Provider for Cisco Meraki repository](https://github.com/cisco-open/terraform-provider-meraki/issues).

# Change log

All notable changes to this project will be documented in the [CHANGELOG](./CHANGELOG.md) file.

The development team may make additional changes as the library evolves with the Cisco Meraki.

## License

This library is distributed under the license found in the [LICENSE](./LICENSE) file.
