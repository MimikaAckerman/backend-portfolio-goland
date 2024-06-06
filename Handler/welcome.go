package Handler

import (
    "github.com/gin-gonic/gin"
    "net/http"

)

func Welcome(c *gin.Context) {
    c.String(http.StatusOK, "¡Bienvenido a la API de mi portfolio!")
}
