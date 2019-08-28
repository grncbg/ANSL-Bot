package main

import (
	"context"

	"github.com/is2ei/conoha-api-go-client/conoha"
)

// ConohaClient is client of conoha API
type ConohaClient struct {
	settings ConohaSettings
	client   *conoha.Conoha
}

// NewConohaClient makes new conoha client
func NewConohaClient(s ConohaSettings) (*ConohaClient, error) {
	client := conoha.NewConoha(
		s.Endpoints.IdentityServiceURL,
		s.Endpoints.AccountServiceURL,
		s.Endpoints.ComputeServiceURL,
		s.Endpoints.VolumeServiceURL,
		s.Endpoints.ImageServiceURL,
		s.Endpoints.NetworkServiceURL,
		s.Endpoints.ObjectStorageServiceURL,
		s.Endpoints.DatabaseServiceURL,
		s.Endpoints.DNSServiceURL,
		s.Endpoints.MailServiceURL,
		s.Username,
		s.Password,
		s.TenantID,
		"",
	)

	access, _, err := client.IdentityToken(context.Background())
	if err != nil {
		return nil, err
	}

	client.Token = access.Token.ID

	conohaClient := ConohaClient{
		settings: s,
		client:   client,
	}
	conohaClient.settings.Token = access.Token.ID

	return &conohaClient, nil
}
