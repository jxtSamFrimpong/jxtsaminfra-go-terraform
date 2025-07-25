// ================================================================
// internal/provider/resource_instance_type_changer.go - Terraform resource
// ================================================================
package provider

import (
	"context"
	// "fmt"
	"jxtsaminfra_ec2_change_instance_type/internal/ec2manager"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInstanceTypeChanger() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInstanceTypeChangerCreate,
		ReadContext:   resourceInstanceTypeChangerRead,
		UpdateContext: resourceInstanceTypeChangerUpdate,
		DeleteContext: resourceInstanceTypeChangerDelete,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "EC2 instance ID to modify",
			},
			"target_instance_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Target instance type (e.g., t3.medium, m5.large)",
			},
			"current_instance_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current instance type",
			},
			"instance_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current instance state",
			},
		},
	}
}

func resourceInstanceTypeChangerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceInstanceTypeChangerUpdate(ctx, d, meta)
}

func resourceInstanceTypeChangerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	manager := meta.(*ec2manager.EC2Manager)
	instanceID := d.Get("instance_id").(string)

	instanceInfo, err := manager.GetInstance(instanceID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("current_instance_type", instanceInfo.InstanceType) 
	d.Set("instance_state", instanceInfo.State) 
	d.SetId(instanceID)

	return nil
}

func resourceInstanceTypeChangerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	manager := meta.(*ec2manager.EC2Manager)
	instanceID := d.Get("instance_id").(string)
	targetInstanceType := d.Get("target_instance_type").(string)

	// Change the instance type
	instanceInfo, err := manager.ChangeInstanceType(instanceID, targetInstanceType)
	if err != nil {
		return diag.FromErr(err)
	}

	// Update state
	d.Set("current_instance_type", instanceInfo.InstanceType) 
	d.Set("instance_state", instanceInfo.State)
	d.SetId(instanceID)

	return nil
}

func resourceInstanceTypeChangerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// For this resource, delete just removes it from state
	// The EC2 instance remains unchanged
	d.SetId("")
	return nil
}