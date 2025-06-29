# `jxtSamFrimpong/jxtsaminfra` Terraform Provider

[![CI / Release Tagging](https://github.com/jxtSamFrimpong/terraform-provider-jxtsaminfra/actions/workflows/test.yaml/badge.svg)](https://github.com/jxtSamFrimpong/terraform-provider-jxtsaminfra/actions/workflows/test.yaml)
[![Release Provider](https://github.com/jxtSamFrimpong/terraform-provider-jxtsaminfra/actions/workflows/release.yaml/badge.svg)](https://github.com/jxtSamFrimpong/terraform-provider-jxtsaminfra/actions/workflows/release.yaml)
[![Terraform Registry](https://img.shields.io/badge/Terraform_Registry-Provider-blue?logo=terraform)](https://registry.terraform.io/providers/jxtSamFrimpong/jxtsaminfra/latest)

This repository contains the source code for the `jxtSamFrimpong/jxtsaminfra` Terraform provider. This provider is designed to manage specific infrastructure operations within AWS, focusing on the ability to change the instance type of an existing EC2 instance.

## 🚀 Features

The current version of this provider offers the following core functionality:

* **`jxtsaminfra_instance_type_changer` resource**: Allows you to dynamically change the instance type of an existing AWS EC2 instance. This operation involves stopping the instance (if running), changing its type, and then restarting it.

## 📖 Requirements

To use this provider, you need:

* [Terraform CLI](https://developer.hashicorp.com/terraform/downloads) (v1.0.x or later recommended).

* AWS credentials configured for Terraform to access your AWS environment (e.g., via environment variables, AWS CLI config, or IAM roles).

To build the provider from source, you also need:

* [Go](https://go.dev/doc/install) (v1.21 or later).

## 📦 Installation (Using the Released Provider)

The easiest way to use the `jxtsaminfra` provider is to declare it in your Terraform configuration. The provider is published on the [Terraform Registry](https://registry.terraform.io/providers/jxtSamFrimpong/jxtsaminfra/latest).

1.  **Declare the provider** in your Terraform configuration (`.tf` file):

    ```terraform
    terraform {
      required_providers {
        jxtsaminfra = {
          source = "jxtSamFrimpong/jxtsaminfra"
          version = "1.0.22" # Use the latest published version
        }
      }
    }

    provider "jxtsaminfra" {
        #   access_key = ""
        #   secret_key = ""
        #   session_token = ""
            region = "eu-west-1"
        }
    ```

2.  **Initialize your Terraform working directory**:

    ```bash
    terraform init
    ```

    Terraform will automatically download the `jxtsaminfra` provider plugin.

## 🛠️ Building the Provider from Source (For Developers)

If you wish to contribute to the provider or use a custom build, follow these steps:

1.  **Clone the repository**:

    ```bash
    git clone [https://github.com/jxtSamFrimpong/terraform-provider-jxtsaminfra.git](https://github.com/jxtSamFrimpong/terraform-provider-jxtsaminfra.git)
    cd terraform-provider-jxtsaminfra
    ```

2.  **Navigate to the Go module directory**:

    ```bash
    cd jxtsaminfra-ec2-change-instance-type
    ```

3.  **Tidy and download Go modules**:

    ```bash
    go mod tidy
    go mod download
    ```

4.  **Build the provider binary**:

    ```bash
    go build -o terraform-provider-jxtsaminfra
    ```

    This will compile the provider binary in the current directory.

5.  **Install the custom provider for local use**:
    To allow Terraform CLI to find your custom-built provider, you need to place it in the correct plugin directory.

    * **Linux/macOS**:

        ```bash
        mkdir -p ~/.terraform.d/plugins/registry.terraform.io/jxtSamFrimpong/jxtsaminfra/<VERSION>/<OS_ARCH>/
        cp terraform-provider-jxtsaminfra ~/.terraform.d/plugins/registry.terraform.io/jxtSamFrimpong/jxtsaminfra/<VERSION>/<OS_ARCH>/
        ```

        Replace `<VERSION>` with a version string (e.g., `1.0.0-dev`) and `<OS_ARCH>` with your operating system and architecture (e.g., `darwin_amd64`, `linux_amd64`).

    * **Windows**:

        ```powershell
        New-Item -ItemType Directory -Path "$env:APPDATA\terraform.d\plugins\registry.terraform.io\jxtSamFrimpong\jxtsaminfra\<VERSION>\<OS_ARCH>" -Force
        Copy-Item terraform-provider-jxtsaminfra.exe "$env:APPDATA\terraform.d\plugins\registry.terraform.io\jxtsaminfra\jxtsaminfra\<VERSION>\<OS_ARCH>\"
        ```

        Replace `<VERSION>` and `<OS_ARCH>` appropriately (e.g., `windows_amd64`).

    After placing the binary, run `terraform init` in your configuration directory to pick up the new local provider.

## 🚀 Usage Examples

Here's a basic example of how to use the `jxtsaminfra_instance_type_changer` resource:

```terraform
# main.tf

# Configure the AWS provider for other AWS resources if needed
# provider "aws" {
#   region = "us-east-1"
# }

resource "aws_instance" "example" {
  ami           = "ami-0abcdef1234567890" # Replace with a valid AMI ID
  instance_type = "t2.micro" # Initial instance type
  tags = {
    Name = "MyTestInstance"
  }
}

resource "jxtsaminfra_instance_type_changer" "change_type_example" {
  instance_id          = aws_instance.example.id
  target_instance_type = "t2.small" # The desired instance type
}