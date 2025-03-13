package initiate_dependency

import (
	"go-api-frame/controller"

	"github.com/gorilla/mux"
)

type ControllerContainer struct {
	LicenseController *controller.LicenseController
}

func InitController(router *mux.Router, endpoint *EndpointContainer) *ControllerContainer {
	return &ControllerContainer{
		LicenseController: controller.NewLicenseController(router, endpoint.LicenseEndpoint),
	}
}
