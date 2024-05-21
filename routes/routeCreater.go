package routes

import (
	"encoding/json"
	"example/models"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRouter() *models.Router {
	router := &models.Router{
		Router : mux.NewRouter(),
	}
	return router
}

func CreateRoutes(routerObject *models.Router) {
	router := routerObject.Router

	router.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := struct {
			Message string
			Status int
		}{
			Message: "Server Running",
			Status: 200,
		}

		json.NewEncoder(w).Encode(response)
	})

	//Authentication Routes
	router.HandleFunc("/users/signup",func(w http.ResponseWriter, r *http.Request) {})
	router.HandleFunc("/users/login",func(w http.ResponseWriter, r *http.Request) {})

	router.HandleFunc("/users",func(w http.ResponseWriter, r *http.Request) {})
	router.HandleFunc("/users/:user_id",func(w http.ResponseWriter, r *http.Request) {})
}