package main

import (
    "log"
    "net/http"
    "portfolio-backend/config"
    "portfolio-backend/routes"
)

func main() {
    config.LoadEnvVariables()
    config.ConnectDB()

    r := routes.RegisterRoutes()

    // Obtén el puerto desde la variable de entorno PORT, si no está establecida usa el puerto 8080 por defecto.
    PORT := config.GetPort()

    log.Printf("Server running on port %s", PORT)
    log.Fatal(http.ListenAndServe(":"+PORT, r))
}
