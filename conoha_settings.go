package main

import "github.com/kelseyhightower/envconfig"

// ConohaSettings has ConoHa API settings
type ConohaSettings struct {
	Endpoints struct {
		AccountServiceURL       string `split_words:"true"`
		ComputeServiceURL       string `split_words:"true"`
		VolumeServiceURL        string `split_words:"true"`
		DatabaseServiceURL      string `split_words:"true"`
		ImageServiceURL         string `split_words:"true"`
		DNSServiceURL           string `split_words:"true"`
		ObjectStorageServiceURL string `split_words:"true"`
		MailServiceURL          string `split_words:"true"`
		IdentityServiceURL      string `split_words:"true"`
		NetworkServiceURL       string `split_words:"true"`
	}
	Username string `required:"true"`
	Password string `required:"true"`
	TenantID string `required:"true" split_words:"true"`
	Token    string
	Image    string `required:"true"`
}

// NewConohaSettings is constructor of ConohaSettings
func NewConohaSettings() (*ConohaSettings, error) {
	var conohaSettings ConohaSettings
	err := envconfig.Process("conoha", &conohaSettings)
	return &conohaSettings, err
}
