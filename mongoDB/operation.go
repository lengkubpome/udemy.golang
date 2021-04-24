package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func readSingleDoc(collection *mongo.Collection, year int) {
	fmt.Println(collection.Name())
	ctx, cancelFindOne := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFindOne()

	filter := bson.M{"year": year}
	singleResult := collection.FindOne(ctx, filter) //.Decord(&result)
	raw, err := singleResult.DecodeBytes()
	if err != nil {
		log.Fatal("Main : find document => ", err)
	}
	fmt.Println(raw)

	result := Movie{}
	singleResult.Decode(&result)
	fmt.Println("--------------------------")
	resultByte, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatal("Main : unabled to marshl json => ", err)
	}
	fmt.Println(string(resultByte))
}

func insertDoc(collection *mongo.Collection) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// res, err := collection.InsertOne(ctx, Movie{ID: primitive.NewObjectID(), ImdbID: 13456789, Title: "Avengers EndGame"})
	res, err := collection.InsertOne(ctx, bson.M{"_id": primitive.NewObjectID(), "imdbID": 77777, "title": "Avengers EndGame"})
	if err != nil {
		return nil, err
	}
	return res, nil

}
