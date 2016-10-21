package rgw

import (
	"log"

	"github.com/minio/minio-go"
)

type Config struct {
	AccessKey  string
	SecretKey  string
	Insecure   bool
	S3Endpoint string
	Region     string
}

type RGWClient struct {
	region string
	s3conn *minio.Client
}

// NewClient() returns a new client.
func (c *Config) NewClient() (interface{}, error) {

	var secure bool

	if c.Insecure {
		secure = !c.Insecure
	} else {
		secure = true
	}

	var client RGWClient

	client.region = c.Region

	s3conn, err := minio.New(c.S3Endpoint, c.AccessKey, c.SecretKey, secure)
	if err != nil {
		log.Printf("[DEBUG] RGW S3 bucket: minio.New error")
		log.Printf(err.Error())
		return nil, nil
	}

	client.s3conn = s3conn

	return &client, nil
}
