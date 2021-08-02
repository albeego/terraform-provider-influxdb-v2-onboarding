---
layout: "influxdbv2-onboarding"
page_title: "InfluxDB V2 onboarding: influxdbv2-onboarding_setup"
sidebar_current: "docs-influxdbv2-onboarding-resource-setup"
description: |-
  The influxdbv2-oboarding_setup resource manages influxdb v2 onboarding setup.
---

## Example Usage

```hcl
resource "influxdbv2-onboarding_setup" "setup" {
  username = "<some_user>"
  password = "<some_password>"
  bucket = "<some_bucket>"
  org = "<some_organization"
  retention_period = <retention_period_in_hour>
}
```

## Argument Reference

The following arguments are supported:

* ``username`` (Required) The username the provider will create at first - Default: `administrator` - Can be set with `INFLUXDB_V2_USERNAME` environment variable.
* ``password`` (Required) The password created for the username - Default: `Administrator1.` - Can be set with `INFLUXDB_V2_PASSWORD` environment variable.
* ``bucket`` (Required) The bucket the provider creates at first.
* ``org`` (Required) The organization the provider creates with initial user.
* ``retention_period`` (Optional) Duration to keep data in hours.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* ``token`` - The token linked with the new user
* ``org_id`` - The id of the new organization
* ``user_id`` - The id of the new user
* ``bucket_id`` - The id of the new bucket
* ``auth_id`` - The id of the new authentication
* ``allowed`` - The status of the allowed variable - can be `true` or `false`
* ``server_url`` - The server URL the provider called