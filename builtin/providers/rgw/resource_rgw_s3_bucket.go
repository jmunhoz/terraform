package rgw

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceRGWS3Bucket() *schema.Resource {
	return &schema.Resource{
		Create: resourceRGWS3BucketCreate,
		Read:   resourceRGWS3BucketRead,
		Update: resourceRGWS3BucketUpdate,
		Delete: resourceRGWS3BucketDelete,

		Schema: map[string]*schema.Schema{
			"bucket": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceRGWS3BucketCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] RGW S3 bucket: resourceRGWS3BucketCreate")
	return nil
}

func resourceRGWS3BucketRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] RGW S3 bucket: resourceRGWS3BucketRead")
	return nil
}

func resourceRGWS3BucketUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] RGW S3 bucket: resourceRGWS3BucketUpdate")
	return nil
}

func resourceRGWS3BucketDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] RGW S3 bucket: resourceRGWS3BucketDelete")
	return nil
}
