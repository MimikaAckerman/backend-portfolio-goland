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

var DB *mongo.Client

func ConnectDB() {
    // Cargar el archivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error cargando el archivo .env: %v", err)
    }

    // Obtener la URI de MongoDB desde las variables de entorno
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        log.Fatal("MONGO_URI no está configurada en el archivo .env")
    }

    // Crear las opciones del cliente
    clientOptions := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))

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
