terraform {
  required_providers {
    ec2instancetype = {
      source  = "local/custom/ec2-instance-type"
      version = "1.0.1"
    }
  }
}

provider "ec2instancetype" {
  access_key = ""
  secret_key = ""
  session_token = ""
  region = "eu-west-1"
}