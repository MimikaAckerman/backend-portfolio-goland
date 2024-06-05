package routes

import (
    "portfolio-backend/api"
    "github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
    r := mux.NewRouter()

    // Definir las rutas
    r.HandleFunc("/formacion", api.GetFormacion).Methods("GET")
    r.HandleFunc("/experience", api.GetExperience).Methods("GET")

    return r
}
