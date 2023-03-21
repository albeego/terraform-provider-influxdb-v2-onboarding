terraform {
  required_providers {
    influxdb-v2-onboarding = {
      source = "albeego/influxdb-v2-onboarding"
      version = "0.2.0"
    }
  }
}

provider "influxdb-v2-onboarding" {
  url = "http://localhost:8086"
}

resource "influxdb-v2-onboarding_setup" "setup" {
  username = "test"
  password = "test1234"
  bucket = "test-bucket"
  org = "test-org"
  retention_period = 4
}

output "token" {
  value = influxdb-v2-onboarding_setup.setup.token
}
output "user_id" {
  value = influxdb-v2-onboarding_setup.setup.user_id
}
output "org_id" {
  value = influxdb-v2-onboarding_setup.setup.org_id
}
output "bucket_id" {
  value = influxdb-v2-onboarding_setup.setup.bucket_id
}
output "auth_id" {
  value = influxdb-v2-onboarding_setup.setup.auth_id
}