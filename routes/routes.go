package routes

import (
    "portfolio-backend/Handler"
    "github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
    r := mux.NewRouter()

    // Definir las rutas
    r.HandleFunc("/", Handler.Welcome).Methods("GET")
    r.HandleFunc("/formacion", Handler.GetFormacion).Methods("GET")
    r.HandleFunc("/experience", Handler.GetExperience).Methods("GET")

    return r
}
