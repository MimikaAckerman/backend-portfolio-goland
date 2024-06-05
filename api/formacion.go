package api

import (
    "context"
    "encoding/json"
    "net/http"
    "portfolio-backend/config"
    "go.mongodb.org/mongo-driver/bson"
)

func GetFormacion(w http.ResponseWriter, r *http.Request) {
    collection := config.GetCollection("Formacion")
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(context.Background())

    var formacion []bson.M
    if err = cursor.All(context.Background(), &formacion); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(formacion)
}
