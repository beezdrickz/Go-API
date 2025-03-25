package main

import (
	"fmt"
	"go-api-frame/config"
	"go-api-frame/controller"
	"go-api-frame/initiate_dependency"
	"go-api-frame/logger"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// for initiate logger
	logger.InitializeLogger()

	// for connecting to DATABASE
	config.Connect()
	defer config.Close()

	router := mux.NewRouter()
	router.Use(controller.CorsMiddleware)
	apiRouter := router.PathPrefix("/test/api/v1").Subrouter()

	//for initiate all DI
	initiate_dependency.DependencyInjection(config.GetDB(), apiRouter)

	apiRouter.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"status": "UP"}`)
	})

	logger.LogServerStart("9190", "/test/api/v1")
	if err := http.ListenAndServe(":9190", router); err != nil {
		logger.LogError("Error starting server:", "", err.Error())
	}
}
