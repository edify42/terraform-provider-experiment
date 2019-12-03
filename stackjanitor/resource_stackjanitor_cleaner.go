package stackjanitor

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func newCleaner() *schema.Resource {
	return &schema.Resource{

		Create: createCleaner,
		Read:   readCleaner,
		Update: updateCleaner,
		Delete: deleteCleaner,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"time_to_live": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"enabled_clients": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"realms": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
	}
}

func createCleaner(d *schema.ResourceData, m interface{}) error {
	return nil
}

func readCleaner(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateCleaner(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteCleaner(d *schema.ResourceData, m interface{}) error {
	return nil
}
