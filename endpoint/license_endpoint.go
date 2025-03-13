package endpoint

import (
	"go-api-frame/services/license"
	"net/http"
)

type LicenseEndpoint struct {
	licenseService license.LicenseService
}

func NewLicenseEndpoint(licenseService license.LicenseService) *LicenseEndpoint {
	return &LicenseEndpoint{
		licenseService: licenseService,
	}
}

func (l *LicenseEndpoint) EndpointWithoutParam(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		writeHttpResponseResult(false, w, r, l.licenseService.InsertLicense)
	case "GET":
	}
}

func (l *LicenseEndpoint) EndpointWithParam(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		writeHttpResponseResult(false, w, r, l.licenseService.ViewDetail)
	}
}
