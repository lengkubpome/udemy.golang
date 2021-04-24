package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// Read data year
	readSingleDoc(collection, 1894)

	// Insert data
	// res, err := insertDoc(collection)
	// if err != nil {
	// 	log.Fatal("Main : insert failed =>", err)
	// }
	// fmt.Println("Insert success ", res.InsertedID)

	// Delete data
	// 1.ลบ 1 ข้อมูล
	// ctx, _ = context.WithTimeout(context.Background(), 3*time.Second)
	// oid, err := primitive.ObjectIDFromHex("6084563ed68e4ddaf9b90752")
	// if err != nil {
	// 	log.Fatal("Main : cannot create object id.")
	// }
	// dResult, err := collection.DeleteOne(ctx, bson.M{"_id": oid}) // ลบต้วเดียว
	// if err != nil {
	// 	log.Fatal("Main : cannot delete.")
	// }
	// fmt.Println("Delete result : ", dResult.DeletedCount)

	// 2.ลบทั้ง collection
	ctx, _ = context.WithTimeout(context.Background(), 3*time.Second)
	err = collection.Drop(ctx)
	if err != nil {
		log.Fatal("Main : cannot delete")
	}
	fmt.Println("Collection dropped.")

	// Update data
	// ข้อมูลเพิ่มเติม
	// https://docs.mongodb.com/manual/reference/operator/update/

	// ctx, _ = context.WithTimeout(context.Background(), 3*time.Second)
	// oid, err := primitive.ObjectIDFromHex("6084563ed68e4ddaf9b90752")
	// if err != nil {
	// 	log.Fatal("Main : cannot create object id.")
	// }

	// Update value
	// updateValue := bson.D{
	// 	{Key: "$set", Value: bson.D{
	// 		{Key: "imdbID", Value: 777731},
	// 	}}}
	// uResult, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, updateValue)
	// if err != nil {
	// 	log.Fatal("Main : cannot update document.")
	// }
	// fmt.Println("Update successfully : ", uResult.ModifiedCount)

	// Rename key
	// updateRename := bson.D{{Key: "$rename", Value: bson.D{{Key: "Title", Value: "title"}}}}
	// uResult2, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, updateRename)
	// if err != nil {
	// 	log.Fatal("Main : cannot update document.")
	// }
	// fmt.Println("Update successfully : ", uResult2.ModifiedCount)

}

type Movie struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	ImdbID      int                `json:"imdbID" bson:"imdbID"`
	Title       string             `json:"title" bson:"title"`
	Year        int                `json:"year" bson:"year"`
	Runtime     string             `json:"runtime" bson:"runtime"`
	Genre       string             `json:"genre" bson:"genre"`
	Released    string             `json:"released" bson:"released"`
	Director    string             `json:"director" bson:"director"`
	Cast        string             `json:"cast" bson:"cast"`
	ImdbRating  float64            `json:"imdbRating" bson:"imdbRating"`
	ImdbVotes   int                `json:"imdbVotes" bson:"imdbVotes"`
	Plot        string             `json:"plot" bson:"plot"`
	Fullplot    string             `json:"fullplot" bson:"fullplot"`
	Country     string             `json:"country" bson:"country"`
	Lastupdated string             `json:"lastupdated" bson:"lastupdated"`
	Type        string             `json:"type" bson:"type"`
}
