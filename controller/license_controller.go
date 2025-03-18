package controller

import (
	"go-api-frame/endpoint"

	"github.com/gorilla/mux"
)

type LicenseController struct {
	licenseEndpoint *endpoint.LicenseEndpoint // Injected Endpoint
	router          *mux.Router
}

// Constructor for LicenseController
func NewLicenseController(router *mux.Router, endpoint *endpoint.LicenseEndpoint) *LicenseController {
	controller := &LicenseController{
		router:          router.PathPrefix("/license").Subrouter(),
		licenseEndpoint: endpoint,
	}
	controller.RegisterRoutes()
	return controller
}

// Setup Routes
func (lc *LicenseController) RegisterRoutes() {
	lc.router.HandleFunc("/", lc.licenseEndpoint.EndpointWithoutParam).Methods("POST", "GET", "OPTION")
	lc.router.HandleFunc("/file", lc.licenseEndpoint.EndpointWithFile).Methods("POST", "OPTION")
	lc.router.HandleFunc("/{ID}", lc.licenseEndpoint.EndpointWithParam).Methods("GET", "OPTION")
}
