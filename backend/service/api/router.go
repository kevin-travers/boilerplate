package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter(srv *server) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health_check", srv.healthCheck).Methods(http.MethodGet)
	router.HandleFunc("/job", srv.getJob).Methods(http.MethodGet)
	router.HandleFunc("/job", srv.createJob).Methods(http.MethodPost)
	router.HandleFunc("/job", srv.updateJob).Methods(http.MethodPut)
	router.HandleFunc("/job", srv.deleteJob).Methods(http.MethodDelete)
	return router
}
