package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(env *Env) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	//dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass

	log.Printf(dbHost)
	mongodbURI := fmt.Sprintf("mongodb+srv://%s:%s@%s", dbUser, dbPass, dbHost)

	if dbUser == "" || dbPass == "" {
		mongodbURI = fmt.Sprintf("mongodb://%s", dbHost)
	}

	clientOptions := options.Client().ApplyURI(mongodbURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to create MongoDB client:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	log.Println("Connected to MongoDB successfully.")

	InitDatabaseCollections(client, env.DBName)

	return client
}

func CloseMongoDBConnection(client *mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal("Failed to close MongoDB connection:", err)
	}

	log.Println("Connection to MongoDB closed.")
}

// InitDatabaseCollections inicializa as coleções e os indices necessários do banco de dados.
func InitDatabaseCollections(client *mongo.Client, dbName string) {
	db := client.Database(dbName)

	initStatusCollection(db)
	initPriorityCollection(db)

	//inicialize os índices
	InitDatabaseIndexes(db)
}

func initStatusCollection(db *mongo.Database) {
	statusCollection := db.Collection("status")
	statuses := []interface{}{
		bson.M{"name": "Pendente"},
		bson.M{"name": "Em Andamento"},
		bson.M{"name": "Finalizado"},
	}

	for _, status := range statuses {
		_, err := statusCollection.UpdateOne(
			context.TODO(),
			bson.M{"name": status.(bson.M)["name"]},
			bson.D{{Key: "$setOnInsert", Value: status}},

			options.Update().SetUpsert(true),
		)
		if err != nil {
			log.Fatal("Failed to initialize status collection:", err)
		}
	}

	log.Println("Status collection initialized.")
}

func initPriorityCollection(db *mongo.Database) {
	priorityCollection := db.Collection("priorities")
	priorities := []interface{}{
		bson.M{"name": "Baixa"},
		bson.M{"name": "Media"},
		bson.M{"name": "Alta"},
	}

	for _, priority := range priorities {
		_, err := priorityCollection.UpdateOne(
			context.TODO(),
			bson.M{"name": priority.(bson.M)["name"]},
			bson.D{{Key: "$setOnInsert", Value: priority}},

			options.Update().SetUpsert(true),
		)
		if err != nil {
			log.Fatal("Failed to initialize priority collection:", err)
		}
	}

	log.Println("Priority collection initialized.")
}

func InitDatabaseIndexes(db *mongo.Database) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Índices para a coleção 'tasks'
	indexModels := []mongo.IndexModel{
		{
			Keys: bson.M{"title": 1},
		},
		{
			Keys: bson.M{"creationDate": 1},
		},
		{
			Keys: bson.M{"userID": 1},
		},
	}

	for _, model := range indexModels {
		_, err := db.Collection("tasks").Indexes().CreateOne(ctx, model)
		if err != nil {
			log.Fatal("Failed to create index:", err)
		}
	}

	log.Println("All necessary indexes for 'tasks' collection created successfully.")
}
