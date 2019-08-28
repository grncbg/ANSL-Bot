package main

import (
	"context"
)

// ComputeServer is info of compute server
type ComputeServer struct {
	ID            string
	ImageID       string
	FlavorID      string
	AdminPassword string
	IPv4Adder     string
}

// AddComputeServer adds a ComputeServer
func (c ConohaClient) AddComputeServer(ctx context.Context, spec ComputeServer) (*ComputeServer, error) {

	if spec.ImageID == "" {
		spec.ImageID = c.settings.Image
	}

	if spec.FlavorID == "" {
		spec.FlavorID = c.settings.Flavor
	}

	if spec.AdminPassword == "" {
		spec.AdminPassword = c.settings.AdminPassword
	}

	server, _, err := c.client.AddComputeServer(ctx, spec.ImageID, spec.FlavorID, spec.AdminPassword, "", "")
	if err != nil {
		return nil, err
	}
	spec.ID = server.ID

	return &spec, nil

}
