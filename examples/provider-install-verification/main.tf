terraform {
  required_providers {
    keyhubpreview = {
      source = "registry.terraform.io/hashicorp/keyhubpreview"
    }
  }
}

variable "keyhub_secret" {
  type        = string
  description = "Client secret on KeyHub"
}

variable "keyhub_secret_local" {
  type        = string
  description = "Client secret on KeyHub"
}

provider "keyhubpreview" {
  #  issuer       = "https://keyhub.topicusonderwijs.nl"
  #  clientid     = "3a5e82ad-3f0d-4a63-846b-4b3e431f1135"
  issuer       = "https://keyhub.localhost:8443"
  clientid     = "ebdf81ac-b02b-4335-9dc4-4a9bc4eb406d"
  clientsecret = var.keyhub_secret_local
}

#data "keyhubpreview_group" "test" {
#  #  uuid = "2fb85263-6406-44f9-9e8a-b1a6d1f43250"
#  uuid       = "c6c98d08-2cbf-45e9-937a-c5c0427348e2"
#  additional = ["accounts"]
#}

#output "data_group" {
#  value = data.keyhubpreview_group.test
#}

#data "keyhubpreview_client" "test" {
#  uuid = "ebdf81ac-b02b-4335-9dc4-4a9bc4eb406d"
#}

#output "data_client" {
#  value = data.keyhubpreview_client.test
#}

resource "keyhubpreview_group_vaultrecord" "terrarecord" {
  name       = "Terraform Record"
  group_uuid = resource.keyhubpreview_group.terra.uuid
  secret = {
    password = "test3"
  }
}

resource "keyhubpreview_group" "terra" {
  name = "Terraform"
  accounts = [{
    uuid   = "7ea6622b-f9d2-4e52-a799-217b26f88376"
    rights = "MANAGER"
  }]
}

resource "keyhubpreview_grouponsystem" "terragos" {
  provisioned_system_uuid = "47923975-b1af-47c8-bd7a-e52ebb4b9b84"
  owner_uuid              = resource.keyhubpreview_group.terra.uuid
  name_in_system          = "cn=terraform,ou=groups,ou=dev,dc=ad01,dc=keyhub,dc=s25,dc=topicus,dc=education"
  type                    = "GROUP"
  provgroups = [{
    activation_required = "false"
    group_uuid          = "c6c98d08-2cbf-45e9-937a-c5c0427348e2"
  }]
}

#output "resource_group" {
#  value = resource.keyhubpreview_group.terra
#}
