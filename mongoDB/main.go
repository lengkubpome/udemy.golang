package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://dbGoUser:Leng24082531@gocourse.syltw.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal("Main : cannot create new Client => ", err)
	}
	collection := client.Database("uncleBobDvd").Collection("movie_initial")
	fmt.Println(collection.Name())
	ctx, cancelFindOne := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFindOne()

	filter := bson.M{"year": "1894"}
	singleResult := collection.FindOne(ctx, filter)
	raw, err := singleResult.DecodeBytes()
	if err != nil {
		log.Fatal("Main : find document => ", err)
	}
	fmt.Println(raw)

}
