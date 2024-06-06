package routes

import (
    "portfolio-backend/Handler"
    "github.com/gin-gonic/gin" // Agrega esta línea
)

func RegisterRoutes() *gin.Engine {
    r := gin.Default()

    // Definir las rutas
    r.GET("/", Handler.Welcome)
    r.GET("/formacion", Handler.GetFormacion)
    r.GET("/experience", Handler.GetExperience)

    return r
}
