package main

import (
    "log"
    "net/http"
    "portfolio-backend/config"
    "portfolio-backend/routes"

    "github.com/gorilla/mux"
)

func main() {
    // Conectar a la base de datos
    config.ConnectDB()

    // Crear un nuevo enrutador
    r := mux.NewRouter()

    // Configurar las rutas
    routes.RegisterRoutes(r)

    // Iniciar el servidor
    log.Println("Servidor iniciado en el puerto 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
