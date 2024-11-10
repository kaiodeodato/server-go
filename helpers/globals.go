package helpers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"fmt"
	"log"
	"github.com/joho/godotenv"
	"os"
)

var Client *mongo.Client

// SetMongoClient sets the MongoDB client to be used in the package
func SetMongoClient(c *mongo.Client) {
	Client = c
}

// PingMongo checks if the MongoDB client can successfully connect to the database
func PingMongo() error {
	if Client == nil {
		return fmt.Errorf("MongoDB client is not initialized")
	}
	err := Client.Ping(context.TODO(), nil)
	if err != nil {
		return fmt.Errorf("Erro ao verificar conexão com o MongoDB: %v", err)
	}
	return nil
}

var Tarefas []TarefaModel

func init() {
    // Carrega o arquivo .env
    envFilePath := ".env" // Padrão local

	// Verifica se o código está sendo executado no Docker (caminho padrão para contêineres)
	if _, err := os.Stat("/app/.env"); err == nil {
		envFilePath = "/app/.env" // Caminho do .env no Docker
	}

	// Carrega o arquivo .env de acordo com o ambiente
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")

    // Monta a URI de conexão
    uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.hdkvcpc.mongodb.net/", dbUser, dbPass)

    // Conecta ao MongoDB
    clientOptions := options.Client().ApplyURI(uri)
    Client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal("Erro ao conectar ao MongoDB:", err)
    }

    // Verifica a conexão
    err = Client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal("Erro ao verificar conexão com o MongoDB:", err)
    }

    // Agora, definimos o client no pacote helpers
    SetMongoClient(Client) // Esse passo é crucial!

    log.Println("Conectado ao MongoDB com sucesso")
}
