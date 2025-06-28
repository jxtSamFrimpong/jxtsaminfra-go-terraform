terraform {
  required_providers {
    jxtsaminfra = {
      source  = "local/custom/jxtsaminfra"
      version = "1.0.3"
    }
  }
}

provider "jxtsaminfra" {
#   access_key = ""
#   secret_key = ""
#   session_token = ""
  region = "eu-west-1"
}