// Package api for handlers
package api

import (
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/mixanemca/example-gorilla-rest-api/internal/app/api/handler/v1"
)

func Routers() {
	router := mux.NewRouter()

	router.HandleFunc("/user", v1.Instance.CreateUser).Methods("POST")

	http.Handle("/", router)
}
