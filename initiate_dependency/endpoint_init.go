package initiate_dependency

import (
	"go-api-frame/endpoint"
)

type EndpointContainer struct {
	LicenseEndpoint *endpoint.LicenseEndpoint
}

func InitEndpoint(service *ServiceContainer) *EndpointContainer {
	return &EndpointContainer{
		LicenseEndpoint: endpoint.NewLicenseEndpoint(service.LicenseService),
	}
}
