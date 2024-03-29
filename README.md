# terraform-provider-influxdb-v2-onboarding
A terraform provider for influxdb v2 onboarding step specifically

This provider is only used to setup the initial configuration on influxdbv2 startup.

The InfluxDB V2 provider allows Terraform to setup
[InfluxDB v2](https://www.influxdata.com/products/influxdb-overview/).

The provider configuration block accepts the following arguments:

* ``url`` (Optional) The root URL of a Influxdb V2 server. May alternativly be set via the INFLUXDB_V2_URL environment variable. Default to `http://localhost:9999`.
* ``username`` (Optional) The username that will be created as administrator. Defaults to `administrator`
* ``password`` (Optional) The password that will be set for the initial user. Defaults to `Administrator1.`

## Build

```bash
go build -o terraform-provider-influxdb-v2-onboarding
```

Don't forget to copy `terraform-provider-influxdbv2` to your terraform plugin directory (eg. `~/.terraform.d/plugins/linux_amd64` on linux).

## Test

To run test, at first run this command to check fmt requirements:
 
```bash
make fmt
```

Then run this command to test the provider instance test: 

```bash
make test
```

And finally to run acceptance test, run this command: 

```bash
make testacc
```

## How to use

At first, you need to start a new Influxdb V2 instance. To do so, you can follow the official documentation [here](https://v2.docs.influxdata.com/v2.0/get-started/#start-with-influxdb-oss)

### Initialize the provider
```hcl
provider "influxdb-v2-onboarding" {
  url = "http://influxdb.example.com:8086"
  username = "influxdbUsername"
  password = "influxdbPassword"
}
 ```

### Available functionalities

* **setup** to setup initial user, bucket and organization, documentation [here](website/docs/r/setup.html.md)

### Examples file
Find more examples in `examples/`. To run them:
```bash
terraform init
terraform apply
```

## Dev

This provider uses the official Go client developed by influxdata itself. For informations, follow this [link](https://github.com/influxdata/influxdb-client-go)
Don't forget to run `go mod tidy` from time to time to remove useless dependencies.