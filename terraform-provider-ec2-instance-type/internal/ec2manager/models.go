// ================================================================
// internal/ec2manager/models.go - Data structures
// ================================================================
package ec2manager

// InstanceInfo represents EC2 instance information
type InstanceInfo struct {
	InstanceID   string
	InstanceType string
	State        string
}

// InstanceState represents possible EC2 instance states
type InstanceState string

const (
	StateRunning    InstanceState = "running"
	StateStopped    InstanceState = "stopped"
	StateTerminated InstanceState = "terminated"
	StatePending    InstanceState = "pending"
)
