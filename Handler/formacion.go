package Handler

import (
    "context"
    "net/http"
    "portfolio-backend/config"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
)

func GetFormacion(c *gin.Context) {
    collection := config.GetCollection("Formacion")
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        return
    }
    defer cursor.Close(context.Background())

    var formacion []bson.M
    if err = cursor.All(context.Background(), &formacion); err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, formacion)
}
