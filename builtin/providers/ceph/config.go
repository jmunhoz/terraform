package ceph

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/go-cleanhttp"
)

type Config struct {
	AccessKey  string
	SecretKey  string
	Insecure   bool
	S3Endpoint string
	Region     string
}

type CephClient struct {
	region string
	s3conn *s3.S3
}

// NewClient() returns a new client.
func (c *Config) NewClient() (interface{}, error) {

	var client CephClient

	client.region = c.Region

	log.Println("[INFO] Building AWS auth structure")

	creds := credentials.NewStaticCredentials(c.AccessKey, c.SecretKey, "")
	_, err := creds.Get()

	if err != nil {
		return nil, errwrap.Wrapf("Error credentials: {{err}}", err)
	}

	S3ForcePathStyle := false

	awsConfig := &aws.Config{
		Credentials:      creds,
		Region:           aws.String(c.Region),
		HTTPClient:       cleanhttp.DefaultClient(),
		S3ForcePathStyle: &S3ForcePathStyle,
	}

	// Set up base session
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, errwrap.Wrapf("Error creating AWS session: {{err}}", err)
	}

	awsS3Sess := sess.Copy(&aws.Config{Endpoint: aws.String(c.S3Endpoint)})

	client.s3conn = s3.New(awsS3Sess)

	return &client, nil
}
