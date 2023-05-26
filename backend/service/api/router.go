package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health_check", healthCheck).Methods(http.MethodGet)
	router.HandleFunc("/job", getJob).Methods(http.MethodGet)
	router.HandleFunc("/job", createJob).Methods(http.MethodPost)
	router.HandleFunc("/job", updateJob).Methods(http.MethodPut)
	router.HandleFunc("/job", deleteJob).Methods(http.MethodDelete)
	return router
}
