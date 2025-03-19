package endpoint

import (
	"go-api-frame/endpoint/utils"
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
		utils.WriteHttpResponseResult(true, w, r, l.licenseService.InsertLicense)
	case "GET":
		utils.WriteHttpResponseResult(true, w, r, l.licenseService.GetList)
	}
}

func (l *LicenseEndpoint) EndpointCount(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		utils.WriteHttpResponseResult(true, w, r, l.licenseService.GetCount)
	}
}

func (l *LicenseEndpoint) EndpointWithParam(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		utils.WriteHttpResponseResult(true, w, r, l.licenseService.ViewDetail)
	}
}

func (l *LicenseEndpoint) EndpointWithFile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		utils.WriteMultipartHttpResponseResult(false, w, r, l.licenseService.InsertLicenseWithFile)
	case "GET":
	}
}
