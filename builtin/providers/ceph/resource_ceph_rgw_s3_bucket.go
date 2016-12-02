package ceph

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCephRGWS3Bucket() *schema.Resource {
	return &schema.Resource{
		Create: resourceCephRGWS3BucketCreate,
		Read:   resourceCephRGWS3BucketRead,
		Update: resourceCephRGWS3BucketUpdate,
		Delete: resourceCephRGWS3BucketDelete,

		Schema: map[string]*schema.Schema{
			"bucket": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"acl": &schema.Schema{
				Type:     schema.TypeString,
				Default:  "private",
				Optional: true,
			},
		},
	}
}

func resourceCephRGWS3BucketCreate(d *schema.ResourceData, meta interface{}) error {
	s3conn := meta.(*CephClient).s3conn

	// Get the bucket and acl
	bucket := d.Get("bucket").(string)
	acl := d.Get("acl").(string)

	log.Printf("[DEBUG] S3 bucket create: %s", bucket)

	req := &s3.CreateBucketInput{
		Bucket: aws.String(bucket),
		ACL:    aws.String(acl),
	}

	if _, err := s3conn.CreateBucket(req); err != nil {
		return fmt.Errorf("Error creating S3 bucket: %s, %s", err)
	}

	return nil
}

func resourceCephRGWS3BucketRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Ceph RGW S3 bucket: resourceCephRGWS3BucketRead")
	return nil
}

func resourceCephRGWS3BucketUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Ceph RGW S3 bucket: resourceCephRGWS3BucketUpdate")
	return nil
}

func resourceCephRGWS3BucketDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] RGW S3 bucket: resourceCephRGWS3BucketDelete")
	return nil
}
