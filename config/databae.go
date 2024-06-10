package config

import (
    "context"
    "log"
    "os"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

// Declarar la variable DB como una variable global
var DB *mongo.Client

func LoadEnvVariables() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}

func ConnectDB() {
    uri := os.Getenv("MONGO_URI")
    if uri == "" {
        log.Fatal("MONGO_URI is not set in .env file")
    }
    
    // Crear las opciones del cliente
    clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))

    // Conectar a MongoDB
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatalf("Error conectando a MongoDB: %v", err)
    }

    // Verificar la conexión
    err = client.Ping(context.TODO(), readpref.Primary())
    if err != nil {
        log.Fatalf("No se pudo conectar a la base de datos: %v", err)
    }

    log.Println("Conexión a la base de datos exitosa")
    DB = client
}

func GetCollection(collectionName string) *mongo.Collection {
    return DB.Database("portfolio").Collection(collectionName)
}

func GetPort() string {
    PORT := os.Getenv("PORT")
    if PORT == "" {
        log.Fatal("PORT is not set in .env file")
    }
    return PORT
}
