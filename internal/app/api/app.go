// Package api for handlers
package api

import (
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/mixanemca/example-gorilla-rest-api/internal/app/api/handler/v1"
)

func Routers() {
	router := mux.NewRouter()

	userHandler := &v1.UsersController{}
	router.HandleFunc("/", OpenAPI).Methods("GET")
	router.HandleFunc("/user", userHandler.CreateUserByID).Methods("POST")
	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", userHandler.DeleteUser).Methods("DELETE")

	http.Handle("/", router)
}

// ShowEndpoints is the temporary, while haven't swagger
func OpenAPI(w http.ResponseWriter, r *http.Request) {
	endpoints := []string{
		"POST /user - Создать пользователя",
		"GET /users - Получить список пользователей",
		"GET /user/{id} - Получить информацию о пользователе по ID",
		"PUT /user/{id} - Обновить информацию о пользователе по ID",
		"DELETE /user/{id} - Удалить пользователя по ID",
	}

	response := "Доступные эндпоинты:\n"
	for _, endpoint := range endpoints {
		response += endpoint + "\n"
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
