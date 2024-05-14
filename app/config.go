package app

import "fmt"

type CloudPlatformType string

const (
	AwsCloudPlatform CloudPlatformType = "AWS"
	GcpCloudPlatform CloudPlatformType = "GCP"
)

type Config struct {
	CloudPlatform             CloudPlatformType `yaml:"CloudPlatform"`
	AwsIgnoreAuthHeaderRegion bool              `yaml:"AwsIgnoreAuthHeaderRegion"`
}

func (c Config) Validate() error {
	if c.CloudPlatform != AwsCloudPlatform && c.CloudPlatform != GcpCloudPlatform {
		return fmt.Errorf("CloudPlatform must be one of: %s, %s", AwsCloudPlatform, GcpCloudPlatform)
	}

	return nil
}
