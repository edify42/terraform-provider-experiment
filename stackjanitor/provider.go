package stackjanitor

import (
	"os"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider is the standard entrypoint function
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"aws_region": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AWS_REGION", nil),
			},
			"debug": {
				Type:     schema.TypeBool,
				Optional: true,
				DefaultFunc: func() (interface{}, error) {
					v := os.Getenv("SJ_DEBUG")
					if v == "" {
						return false, nil
					}

					return v == "1" || v == "true" || v == "on", nil
				},
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"stackjanitor_cleaner": newCleaner(),
		},
	}
}
