package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", index).Methods(http.MethodGet)
	router.HandleFunc("/health_check", healthCheck).Methods(http.MethodGet)
	return router
}
