terraform {
  required_providers {
    jxtsaminfra = {
      source = "jxtSamFrimpong/jxtsaminfra"
      version = "1.0.13"
    }
  }
}

provider "jxtsaminfra" {
#   access_key = ""
#   secret_key = ""
#   session_token = ""
  region = "eu-west-1"
}