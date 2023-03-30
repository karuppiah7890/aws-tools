package config

import (
	"fmt"
	"os"
)

// All configuration is through environment variables

const AWS_REGION_ENV_VAR = "AWS_REGION"
const AWS_ACCESS_KEY_ID_ENV_VAR = "AWS_ACCESS_KEY_ID"
const AWS_SECRET_ACCESS_KEY_ENV_VAR = "AWS_SECRET_ACCESS_KEY"

type Config struct {
	awsRegion          string
	awsAccessKeyId     string
	awsSecretAccessKey string
}

func NewConfigFromEnvVars() (*Config, error) {
	awsRegion, err := getAwsRegion()
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting AWS Region: %v", err)
	}

	awsAccessKeyId, err := getAwsAccessKeyId()
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting AWS Access Key ID: %v", err)
	}

	awsSecretAccessKey, err := getAwsSecretAccessKey()
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting AWS Secret Access Key: %v", err)
	}

	return &Config{
		awsRegion:          awsRegion,
		awsAccessKeyId:     awsAccessKeyId,
		awsSecretAccessKey: awsSecretAccessKey,
	}, nil
}

// Get AWS Region
func getAwsRegion() (string, error) {
	awsRegion, ok := os.LookupEnv(AWS_REGION_ENV_VAR)
	if !ok {
		return "", fmt.Errorf("%s environment variable value is a required value. Please define it", AWS_ACCESS_KEY_ID_ENV_VAR)
	}

	return awsRegion, nil
}

// Get AWS Access Key ID
func getAwsAccessKeyId() (string, error) {
	awsAccessKeyId, ok := os.LookupEnv(AWS_ACCESS_KEY_ID_ENV_VAR)
	if !ok {
		return "", fmt.Errorf("%s environment variable value is a required value. Please define it", AWS_ACCESS_KEY_ID_ENV_VAR)
	}

	return awsAccessKeyId, nil
}

// Get AWS Secret Access Key
func getAwsSecretAccessKey() (string, error) {
	awsSecretAccessKey, ok := os.LookupEnv(AWS_SECRET_ACCESS_KEY_ENV_VAR)
	if !ok {
		return "", fmt.Errorf("%s environment variable value is a required value. Please define it", AWS_SECRET_ACCESS_KEY_ENV_VAR)
	}

	return awsSecretAccessKey, nil
}

func (c *Config) GetAwsRegion() string {
	return c.awsRegion
}

func (c *Config) GetAwsAccessKeyId() string {
	return c.awsAccessKeyId
}

func (c *Config) GetAwsSecretAccessKey() string {
	return c.awsSecretAccessKey
}
