package initiate_dependency

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type AppContainer struct {
	DAO        *DAOContainer
	Service    *ServiceContainer
	Endpoint   *EndpointContainer
	Controller *ControllerContainer
}

func DependencyInjection(db *sql.DB, router *mux.Router) *AppContainer {
	dao := InitDAO(db)
	service := InitService(db, dao)
	endpoint := InitEndpoint(service)
	controller := InitController(router, endpoint)

	return &AppContainer{
		DAO:        dao,
		Service:    service,
		Endpoint:   endpoint,
		Controller: controller,
	}
}
