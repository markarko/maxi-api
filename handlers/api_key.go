package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"os"

	"github.com/markarko/maxi-api/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ApiKeyHandler struct {
	l *log.Logger
}

var apiKeyCollection *mongo.Collection

func InitMongodb() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	apiKeyCollection = client.Database("maxi").Collection("api-keys")
}

func NewApiKeyHandler(l *log.Logger) *ApiKeyHandler {
	return &ApiKeyHandler{l}
}

func (apiKeyHandler *ApiKeyHandler) generateAPIKey() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	apiKey := base64.URLEncoding.EncodeToString(b)

	_, err = apiKeyHandler.AddApiKey(apiKey)
	if err != nil {
		panic(err)
	}
	return apiKey
}

func (apiKeyHandler *ApiKeyHandler) AddApiKey(apiKey string) (*mongo.InsertOneResult, error) {
	return apiKeyCollection.InsertOne(context.Background(), data.ApiKey{Key: apiKey})
}

func ValidApiKey(key string) bool {

	if key == "" {
		return false
	}

	filter := bson.M{"key": key}

	var apiKey data.ApiKey
	err := apiKeyCollection.FindOne(context.Background(), filter).Decode(&apiKey)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("No matching document found")
		} else {
			panic(err)
		}
		return false
	}

	log.Println("Found matching document", apiKey)
	return true
}

func (apiKeyHandler *ApiKeyHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		if r.Header.Get("Token") == "" || r.Header.Get("Token") != os.Getenv("MAXI_API_TOKEN") {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		rw.Write([]byte(apiKeyHandler.generateAPIKey()))
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}
