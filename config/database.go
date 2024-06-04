package config

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

func ConnectDB() {
    clientOptions := options.Client().ApplyURI("mongodb+srv://mimika:1.Cambiame@databasecluster.ldcnr5i.mongodb.net/?retryWrites=true&w=majority&appName=DatabaseCluster")

    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal("Error conectando a MongoDB:", err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Ping(ctx, readpref.Primary())
    if err != nil {
        log.Fatal("Error verificando la conexión a MongoDB:", err)
    }

    log.Println("Conexión a MongoDB exitosa!")
    Client = client
}

func GetCollection(collectionName string) *mongo.Collection {
    return Client.Database("portfolio").Collection(collectionName)
}
