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

    PORT := config.GetPort()

    log.Printf("Server running on port %s", PORT)
    log.Fatal(http.ListenAndServe(":"+PORT, r))
}
