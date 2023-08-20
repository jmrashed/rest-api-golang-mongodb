package router

import (
	"net/http"

	"github.com/jmrashed/rest-api-golang-mongodb/internal/handler"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/users", handler.ListUsers).Methods(http.MethodGet)
	r.HandleFunc("/api/users/{id}", handler.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/api/users", handler.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/api/users/{id}", handler.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/api/users/{id}", handler.DeleteUser).Methods(http.MethodDelete)

	return r
}
