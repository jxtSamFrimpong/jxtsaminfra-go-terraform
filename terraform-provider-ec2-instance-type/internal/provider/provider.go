// ================================================================
// internal/provider/provider.go - Main provider configuration
// ================================================================
package provider

import (
	"context"
	"terraform-provider-ec2-instance-type/internal/ec2manager"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AWS_REGION", "us-east-1"),
				Description: "AWS region",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ec2instancetype_instance_type_changer": resourceInstanceTypeChanger(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	region := d.Get("region").(string)
	
	manager, err := ec2manager.NewEC2Manager(region)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	
	return manager, nil
}
