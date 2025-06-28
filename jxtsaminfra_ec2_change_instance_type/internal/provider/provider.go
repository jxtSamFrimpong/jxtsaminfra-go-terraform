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

package provider

import (
	"context"
	"github.com/jxtSamFrimpong/jxtsaminfra-go-terraform/jxtsaminfra_ec2_change_instance_type/internal/ec2manager"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"region": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("AWS_REGION", "us-east-1"),
					Description: "AWS region",
				},
				"access_key": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("AWS_ACCESS_KEY_ID", ""),
					Description: "AWS access key ID",
					Sensitive:   true,
				},
				"secret_key": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("AWS_SECRET_ACCESS_KEY", ""),
					Description: "AWS secret access key",
					Sensitive:   true,
				},
				"session_token": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("AWS_SESSION_TOKEN", ""),
					Description: "AWS session token (for temporary credentials)",
					Sensitive:   true,
				},
			},
			ResourcesMap: map[string]*schema.Resource{
				"jxtsaminfra_ec2_change_instance_type": resourceEC2ChangeInstanceType(),
			},
			ConfigureContextFunc: providerConfigure,
		}

		return p
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	region := d.Get("region").(string)
	accessKey := d.Get("access_key").(string)
	secretKey := d.Get("secret_key").(string)
	sessionToken := d.Get("session_token").(string)
	
	config := &ec2manager.AWSConfig{
		Region:       region,
		AccessKey:    accessKey,
		SecretKey:    secretKey,
		SessionToken: sessionToken,
	}
	
	manager, err := ec2manager.NewEC2ManagerWithConfig(config)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	
	return manager, nil
}