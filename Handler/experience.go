package Handler

import (
    "context"
    "encoding/json"
    "net/http"
    "portfolio-backend/config"
    "go.mongodb.org/mongo-driver/bson"
)

func GetExperience(w http.ResponseWriter, r *http.Request) {
    collection := config.GetCollection("experience")
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(context.Background())

    var experience []bson.M
    if err = cursor.All(context.Background(), &experience); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(experience)
}
