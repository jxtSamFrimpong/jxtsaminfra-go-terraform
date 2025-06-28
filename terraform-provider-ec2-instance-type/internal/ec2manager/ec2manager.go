// ================================================================
// internal/ec2manager/ec2manager.go - Main EC2 management class
// ================================================================
package ec2manager

import (
	"fmt"
	// "time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// EC2Manager handles all EC2 operations following OOP principles
type EC2Manager struct {
	client *ec2.EC2
	region string
}

// NewEC2Manager creates a new EC2Manager instance using default AWS credentials
func NewEC2Manager(region string) (*EC2Manager, error) {
	config := &AWSConfig{
		Region: region,
	}
	return NewEC2ManagerWithConfig(config)
}

// NewEC2ManagerWithConfig creates a new EC2Manager instance with custom AWS configuration
func NewEC2ManagerWithConfig(config *AWSConfig) (*EC2Manager, error) {
	awsConfig := &aws.Config{
		Region: aws.String(config.Region),
	}
	
	// Only set credentials if access key is provided
	if config.AccessKey != "" {
		awsConfig.Credentials = credentials.NewStaticCredentials(
			config.AccessKey,
			config.SecretKey,
			config.SessionToken,
		)
	}
	// If no access key provided, AWS SDK will use default credential chain
	// (environment variables, IAM roles, shared credentials file, etc.)

	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	return &EC2Manager{
		client: ec2.New(sess),
		region: config.Region,
	}, nil
}

// GetInstance tries to get instance info from AWS
// Returns instance info object if exists, throws error if not
func (e *EC2Manager) GetInstance(instanceID string) (*InstanceInfo, error) {
	fmt.Printf("Getting instance info for: %s\n", instanceID)
	
	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{aws.String(instanceID)},
	}

	result, err := e.client.DescribeInstances(input)
	if err != nil {
		return nil, fmt.Errorf("failed to describe instance %s: %w", instanceID, err)
	}

	if len(result.Reservations) == 0 || len(result.Reservations[0].Instances) == 0 {
		return nil, fmt.Errorf("instance %s not found", instanceID)
	}

	instance := result.Reservations[0].Instances[0]
	
	instanceInfo := &InstanceInfo{
		InstanceID:   instanceID,
		InstanceType: aws.StringValue(instance.InstanceType),
		State:        aws.StringValue(instance.State.Name),
	}

	fmt.Printf("Instance found - Type: %s, State: %s\n", instanceInfo.InstanceType, instanceInfo.State)
	return instanceInfo, nil
}

// GetInstanceState tries to get the instance's state from AWS
func (e *EC2Manager) GetInstanceState(instanceID string) (InstanceState, error) {
	fmt.Printf("Getting instance state for: %s\n", instanceID)
	
	instanceInfo, err := e.GetInstance(instanceID)
	if err != nil {
		return "", err
	}

	state := InstanceState(instanceInfo.State)
	
	switch state {
	case StateRunning:
		return StateRunning, nil
	case StateStopped:
		return StateStopped, nil
	case StateTerminated:
		return "", fmt.Errorf("instance %s is terminated - cannot proceed", instanceID)
	case StatePending:
		return "", fmt.Errorf("instance %s is pending - cannot proceed", instanceID)
	default:
		return "", fmt.Errorf("instance %s is in unexpected state: %s", instanceID, state)
	}
}

// GetInstanceIntoStoppedState tries to stop the instance on AWS
func (e *EC2Manager) GetInstanceIntoStoppedState(instanceID string) (InstanceState, error) {
	fmt.Printf("Getting instance into stopped state: %s\n", instanceID)
	
	// Check current state
	currentState, err := e.GetInstanceState(instanceID)
	if err != nil {
		return "", err
	}

	// If already stopped, return immediately
	if currentState == StateStopped {
		fmt.Printf("Instance %s is already stopped\n", instanceID)
		return StateStopped, nil
	}

	// If running, stop it
	if currentState == StateRunning {
		fmt.Printf("Stopping instance %s...\n", instanceID)
		
		stopInput := &ec2.StopInstancesInput{
			InstanceIds: []*string{aws.String(instanceID)},
		}

		_, err := e.client.StopInstances(stopInput)
		if err != nil {
			return "", fmt.Errorf("failed to stop instance %s: %w", instanceID, err)
		}

		// Wait for instance to be stopped
		fmt.Printf("Waiting for instance %s to stop...\n", instanceID)
		err = e.client.WaitUntilInstanceStopped(&ec2.DescribeInstancesInput{
			InstanceIds: []*string{aws.String(instanceID)},
		})
		if err != nil {
			return "", fmt.Errorf("failed waiting for instance %s to stop: %w", instanceID, err)
		}

		// Verify state
		finalState, err := e.GetInstanceState(instanceID)
		if err != nil {
			return "", err
		}

		if finalState != StateStopped {
			return "", fmt.Errorf("instance %s failed to reach stopped state, current state: %s", instanceID, finalState)
		}

		fmt.Printf("Instance %s successfully stopped\n", instanceID)
		return StateStopped, nil
	}

	return "", fmt.Errorf("cannot stop instance %s from state: %s", instanceID, currentState)
}

// GetInstanceIntoRunningState tries to start an instance on AWS
func (e *EC2Manager) GetInstanceIntoRunningState(instanceID string) (InstanceState, error) {
	fmt.Printf("Getting instance into running state: %s\n", instanceID)
	
	// Check current state
	currentState, err := e.GetInstanceState(instanceID)
	if err != nil {
		return "", err
	}

	// If already running, return immediately
	if currentState == StateRunning {
		fmt.Printf("Instance %s is already running\n", instanceID)
		return StateRunning, nil
	}

	// If stopped, start it
	if currentState == StateStopped {
		fmt.Printf("Starting instance %s...\n", instanceID)
		
		startInput := &ec2.StartInstancesInput{
			InstanceIds: []*string{aws.String(instanceID)},
		}

		_, err := e.client.StartInstances(startInput)
		if err != nil {
			return "", fmt.Errorf("failed to start instance %s: %w", instanceID, err)
		}

		// Wait for instance to be running
		fmt.Printf("Waiting for instance %s to start...\n", instanceID)
		err = e.client.WaitUntilInstanceRunning(&ec2.DescribeInstancesInput{
			InstanceIds: []*string{aws.String(instanceID)},
		})
		if err != nil {
			return "", fmt.Errorf("failed waiting for instance %s to start: %w", instanceID, err)
		}

		// Verify state
		finalState, err := e.GetInstanceState(instanceID)
		if err != nil {
			return "", err
		}

		if finalState != StateRunning {
			return "", fmt.Errorf("instance %s failed to reach running state, current state: %s", instanceID, finalState)
		}

		fmt.Printf("Instance %s successfully started\n", instanceID)
		return StateRunning, nil
	}

	return "", fmt.Errorf("cannot start instance %s from state: %s", instanceID, currentState)
}

// ChangeInstanceType tries to change the instance type of an existing AWS instance
func (e *EC2Manager) ChangeInstanceType(instanceID, targetInstanceType string) (*InstanceInfo, error) {
	fmt.Printf("Starting instance type change for %s to %s\n", instanceID, targetInstanceType)
	
	// Get current instance info
	instanceInfo, err := e.GetInstance(instanceID)
	if err != nil {
		return nil, err
	}

	// Check if target type is different from current
	if instanceInfo.InstanceType == targetInstanceType {
		return nil, fmt.Errorf("instance %s is already of type %s", instanceID, targetInstanceType)
	}

	fmt.Printf("Changing instance type from %s to %s\n", instanceInfo.InstanceType, targetInstanceType)

	// Stop the instance
	_, err = e.GetInstanceIntoStoppedState(instanceID)
	if err != nil {
		return nil, fmt.Errorf("failed to stop instance for type change: %w", err)
	}
	fmt.Printf("✓ Instance %s successfully stopped for type change\n", instanceID)

	// Change instance type
	fmt.Printf("Modifying instance type to %s...\n", targetInstanceType)
	modifyInput := &ec2.ModifyInstanceAttributeInput{
		InstanceId: aws.String(instanceID),
		InstanceType: &ec2.AttributeValue{
			Value: aws.String(targetInstanceType),
		},
	}

	_, err = e.client.ModifyInstanceAttribute(modifyInput)
	if err != nil {
		fmt.Printf("✗ Failed to change instance type: %v\n", err)
		// Still try to start the instance even if type change failed
		e.GetInstanceIntoRunningState(instanceID)
		return nil, fmt.Errorf("failed to modify instance type: %w", err)
	}
	fmt.Printf("✓ Instance type successfully changed to %s\n", targetInstanceType)

	// Start the instance regardless of whether type change succeeded
	_, err = e.GetInstanceIntoRunningState(instanceID)
	if err != nil {
		fmt.Printf("✗ Failed to start instance after type change: %v\n", err)
	} else {
		fmt.Printf("✓ Instance %s successfully started\n", instanceID)
	}

	// Get final instance info
	finalInstanceInfo, err := e.GetInstance(instanceID)
	if err != nil {
		return nil, err
	}

	// Print summary
	fmt.Printf("\n=== SUMMARY ===\n")
	fmt.Printf("Instance ID: %s\n", instanceID)
	fmt.Printf("Previous Type: %s\n", instanceInfo.InstanceType)
	fmt.Printf("New Type: %s\n", finalInstanceInfo.InstanceType)
	fmt.Printf("Current State: %s\n", finalInstanceInfo.State)
	fmt.Printf("Operation: SUCCESS\n")
	fmt.Printf("===============\n\n")

	return finalInstanceInfo, nil
}

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
  
  # Optional: Specify AWS credentials if not using environment/IAM roles
  # access_key    = "your-access-key"
  # secret_key    = "your-secret-key"
  # session_token = "your-session-token"  # Only needed for temporary credentials
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
