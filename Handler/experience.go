package Handler

import (
    "context"
    "net/http"
    "portfolio-backend/config"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
)

func GetExperience(c *gin.Context) {
    collection := config.GetCollection("experience")
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        return
    }
    defer cursor.Close(context.Background())

    var experience []bson.M
    if err = cursor.All(context.Background(), &experience); err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, experience)
}
