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
    ec2instancetype = {
      source  = "local/custom/ec2-instance-type"
      version = "1.0.0"
    }
  }
}

provider "ec2instancetype" {
  region = "us-east-1"  # Change to your preferred region
}

resource "ec2instancetype_instance_type_changer" "example" {
  instance_id          = "i-1234567890abcdef0"  # Replace with your instance ID
  target_instance_type = "t3.medium"            # Replace with desired type
}

output "instance_info" {
  value = {
    instance_id          = ec2instancetype_instance_type_changer.example.instance_id
    current_type         = ec2instancetype_instance_type_changer.example.current_instance_type
    target_type          = ec2instancetype_instance_type_changer.example.target_instance_type
    state               = ec2instancetype_instance_type_changer.example.instance_state
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