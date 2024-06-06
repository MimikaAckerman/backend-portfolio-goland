package main

import (
    "log"
    "net/http"
    "os"
    "portfolio-backend/config"
    "portfolio-backend/routes"
    "github.com/rs/cors"
)

var (
    handler http.Handler
)

func init() {
    // Conectar a la base de datos
    config.ConnectDB()

    // Registrar rutas
    r := routes.RegisterRoutes()

    // Configurar CORS
    handler = cors.Default().Handler(r)
}

// Handler maneja todas las solicitudes HTTP
func Handler(w http.ResponseWriter, r *http.Request) {
    handler.ServeHTTP(w, r)
}

// Main se utiliza para ejecutar la aplicación localmente
func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000" // Puerto predeterminado si no se establece el puerto en el entorno
    }

    log.Println("Server running on port", port)
    log.Fatal(http.ListenAndServe(":"+port, handler))
}