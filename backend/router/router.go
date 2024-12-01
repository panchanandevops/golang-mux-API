package router

import (
	"database/sql"
	"net/http"

	"go-crud-app/handlers"
	"go-crud-app/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter(db *sql.DB) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/api/go/users", handlers.GetUsers(db)).Methods("GET")
	r.HandleFunc("/api/go/users", handlers.CreateUser(db)).Methods("POST")
	r.HandleFunc("/api/go/users/{id}", handlers.GetUser(db)).Methods("GET")
	r.HandleFunc("/api/go/users/{id}", handlers.UpdateUser(db)).Methods("PUT")
	r.HandleFunc("/api/go/users/{id}", handlers.DeleteUser(db)).Methods("DELETE")

	return middleware.EnableCORS(middleware.JsonContentTypeMiddleware(r))
}
