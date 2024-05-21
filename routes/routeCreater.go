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
}