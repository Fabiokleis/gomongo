package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Database struct {
	Ctx    context.Context
	Client *mongo.Client
	Name   string
}

func New(name string) Database {
	return Database{Name: name}
}

func (db *Database) Collection(name string) *mongo.Collection {
	return db.Client.Database(db.Name).Collection(name)
}

func (db *Database) Connect() {
	db.Ctx = context.Background()
	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		panic("set your 'MONGODB_URI' environment variable!")
	}
	client, err := mongo.Connect(options.Client().ApplyURI(uri + "/?compressors=snappy,zlib,zstd"))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	db.Client = client
	var ping bson.M
	if err := client.Database("admin").RunCommand(ctx, bson.D{{Key: "ping", Value: 1}}).Decode(&ping); err != nil {
		panic(err)
	}

	log.Printf("ping result: %v\n", ping)

	if err := client.Database("universidade").CreateCollection(ctx, "alunos"); err != nil {
		panic(err)
	}
}

func (db *Database) Disconnect() {
	if db.Client != nil {
		if err := db.Client.Disconnect(db.Ctx); err != nil {
			panic(err)
		}
		log.Println("mongodb disconnected")
		return
	}
	panic("trying to close nil mongo connection")
}
