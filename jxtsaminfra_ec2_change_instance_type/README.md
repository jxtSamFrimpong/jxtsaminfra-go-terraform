// ================================================================
// Build and Installation Instructions
// ================================================================

/*
To build and use this provider:

1. Initialize the Go module:
   go mod init terraform-provider-ec2-instance-type
   go mod tidy

2. Build the provider:
   go build -o terraform-provider-ec2-instance-type

3. Install locally for Terraform to find:
   # Create the plugins directory
   mkdir -p ~/.terraform.d/plugins/local/custom/ec2-instance-type/1.0.0/linux_amd64/
   
   # Copy the binary
   cp terraform-provider-ec2-instance-type ~/.terraform.d/plugins/local/custom/ec2-instance-type/1.0.0/linux_amd64/

4. Create a terraform configuration file (main.tf):
*/

// ================================================================
// Example Terraform Configuration (main.tf)
// ================================================================

/*
terraform {
  required_providers {
    jxtsaminfra = {
      source  = "yourusername/jxtsaminfra"
      version = "~> 1.0"
    }
  }
}

provider "jxtsaminfra" {
  region = "us-east-1"  # Change to your preferred region
  
  # Optional: Specify AWS credentials if not using environment/IAM roles
  # access_key    = "your-access-key"
  # secret_key    = "your-secret-key"
  # session_token = "your-session-token"  # Only needed for temporary credentials
}

resource "jxtsaminfra_ec2_change_instance_type" "example" {
  instance_id          = "i-1234567890abcdef0"  # Replace with your instance ID
  target_instance_type = "t3.medium"            # Replace with desired type
}

output "instance_info" {
  value = {
    instance_id          = jxtsaminfra_ec2_change_instance_type.example.instance_id
    current_type         = jxtsaminfra_ec2_change_instance_type.example.current_instance_type
    target_type          = jxtsaminfra_ec2_change_instance_type.example.target_instance_type
    state               = jxtsaminfra_ec2_change_instance_type.example.instance_state
  }
}
*/

// ================================================================
// Usage Commands
// ================================================================

/*
After creating main.tf:

1. Initialize Terraform:
   terraform init

2. Plan the changes:
   terraform plan

3. Apply the changes:
   terraform apply

4. To destroy (removes from state only, doesn't affect the instance):
   terraform destroy
*/

// ================================================================
// Additional Files Required for Terraform Registry
// ================================================================


# Terraform Provider JxtSamInfra

A Terraform provider for managing AWS EC2 instance type changes with proper state management.

## Features

- Change EC2 instance types safely
- Automatic instance stop/start handling
- Comprehensive error handling and validation
- Support for multiple AWS authentication methods

## Usage

```hcl
terraform {
  required_providers {
    jxtsaminfra = {
      source  = "jxtsaminfra/jxtsaminfra"
      version = "~> 1.0"
    }
  }
}

provider "jxtsaminfra" {
  region = "us-east-1"
}

resource "jxtsaminfra_ec2_change_instance_type" "example" {
  instance_id          = "i-1234567890abcdef0"
  target_instance_type = "t3.medium"
}
```

## Provider Configuration

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|----------|
| region | AWS region | string | us-east-1 | no |
| access_key | AWS access key ID | string | "" | no |
| secret_key | AWS secret access key | string | "" | no |
| session_token | AWS session token | string | "" | no |

## Resource: jxtsaminfra_ec2_change_instance_type

### Schema

| Name | Description | Type | Required |
|------|-------------|------|----------|
| instance_id | EC2 instance ID to modify | string | yes |
| target_instance_type | Target instance type | string | yes |
| current_instance_type | Current instance type (computed) | string | no |
| instance_state | Current instance state (computed) | string | no |

### Example

```hcl
resource "jxtsaminfra_ec2_change_instance_type" "web_server" {
  instance_id          = "i-1234567890abcdef0"
  target_instance_type = "t3.large"
}
```

## License

MIT License