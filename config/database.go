package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

// ConnectDatabase establece la conexión a la base de datos MongoDB
func ConnectDatabase() {
	var err error

	// Configura las opciones de la base de datos
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Reemplaza con tu URI de MongoDB

	// Conectar a MongoDB
	DB, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Error al conectar a MongoDB: %v", err)
	}

	// Verifica la conexión
	err = DB.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("No se pudo verificar la conexión con MongoDB: %v", err)
	}

	fmt.Println("Conexión a MongoDB establecida correctamente")
}

// DisconnectDatabase cierra la conexión con la base de datos
func DisconnectDatabase() {
	if err := DB.Disconnect(context.Background()); err != nil {
		log.Fatalf("Error al desconectar de MongoDB: %v", err)
	}
	fmt.Println("Conexión con MongoDB cerrada")
}

// GetDatabase devuelve una instancia de la base de datos específica
func GetDatabase(dbName string) *mongo.Database {
	return DB.Database(dbName)
}
