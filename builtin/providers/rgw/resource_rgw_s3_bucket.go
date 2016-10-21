package rgw

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	_ "github.com/minio/minio-go"
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

	bucketName := d.Get("bucket").(string)
	log.Printf("[DEBUG] bucket = " + bucketName)

	region := meta.(*RGWClient).region
	log.Printf("[DEBUG] region = " + region)

	s3conn := meta.(*RGWClient).s3conn
	if s3conn == nil {
		return fmt.Errorf("Error grabbing s3conn")
	}

	err := s3conn.MakeBucket(bucketName, "us-east-1")
	if err != nil {
		return fmt.Errorf("s3conn.MakeBucket failed " + err.Error())
	}

	log.Printf("[DEBUG] Bucket created successfully.")

	return nil
}

func resourceRGWS3BucketRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] RGW S3 bucket: resourceRGWS3BucketRead")
	return nil
}

func resourceRGWS3BucketUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] RGW S3 bucket: resourceRGWS3BucketUpdate")
	return nil
}

func resourceRGWS3BucketDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] RGW S3 bucket: resourceRGWS3BucketDelete")

	s3conn := meta.(*RGWClient).s3conn
	if s3conn == nil {
		return fmt.Errorf("Error grabbing s3conn")
	}

	bucketName := d.Get("bucket").(string)
	err := s3conn.RemoveBucket(bucketName)
	if err != nil {
		return fmt.Errorf("s3conn.RemoveBucket failed " + err.Error())
	}

	log.Printf("[DEBUG] Bucket removed successfully.")

	return nil
}
