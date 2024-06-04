package routes

import(
	"portfolio-backend/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router){

// Definir las rutas
r.HandleFunc("/formacion", handlers.GetFormacion).Methods("GET")
r.HandleFunc("/experience", handlers.GetExperience).Methods("GET")

}
