package main

import (
	"context"

	"github.com/is2ei/conoha-api-go-client/conoha"
)

// NewConohaClient makes new conoha client
func NewConohaClient(s *ConohaSettings) (*conoha.Conoha, error) {
	conohaClient := conoha.NewConoha(
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

	access, _, err := conohaClient.IdentityToken(context.Background())
	if err != nil {
		return nil, err
	}

	conohaClient.Token = access.Token.ID

	return conohaClient, nil
}
