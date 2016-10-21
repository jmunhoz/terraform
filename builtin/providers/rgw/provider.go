package rgw

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secret_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"s3endpoint": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "betanzos",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"rgw_s3_bucket": resourceRGWS3Bucket(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AccessKey:  d.Get("access_key").(string),
		SecretKey:  d.Get("secret_key").(string),
		S3Endpoint: d.Get("s3endpoint").(string),
		Region:     d.Get("region").(string),
	}
	return config.NewClient()
}
