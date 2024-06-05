package main

import (
    "log"
    "net/http"
    "os"
    "portfolio-backend/config"
    "portfolio-backend/routes"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000" // Puerto predeterminado si no se establece el puerto en el entorno
    }

    config.ConnectDB()

    r := routes.RegisterRoutes()

    log.Println("Server running on port", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}
