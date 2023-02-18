package routes

import (
	"gos/app/http/controllers"

	"github.com/gorilla/mux"
)

func RegisterWebRoutes(r *mux.Router) {
	pc := new(controllers.PagesController)
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
}
