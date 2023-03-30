package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// connection to the MongoDB database using connection string
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Get a handle to the "movies" collection from my mongodb local compass
	collection := client.Database("H1_fs").Collection("h")

	// retrieving documents where the "name" field equals "ned stark"
	filter1 := bson.M{"name": bson.M{"$eq": "Ned Stark"}}//eq selector used to return the value with the given value 
	cursor1, err := collection.Find(context.Background(), filter1)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor1.Close(context.Background())

	// retrieving documents where the "name" field is greater than 5
	filter2 := bson.M{"name": bson.M{"$gt": 5}} //gt selector is used to retrieve greater than values
	cursor2, err := collection.Find(context.Background(), filter2)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor2.Close(context.Background())

	// retrieving documents where the "h.name" field contains "Jorah Mormont"
	filter3 := bson.M{"h.name": bson.M{"$regex": "Jorah Mormont"}}
	cursor3, err := collection.Find(context.Background(), filter3)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor3.Close(context.Background())

	// retrieving documents where the "h.email" field ends with "@example.com"
	filter4 := bson.M{"h.email": bson.M{"$regex": "@gamofthron.es$"}}
	cursor4, err := collection.Find(context.Background(), filter4)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor4.Close(context.Background())

	// retrieving documents where the "name" field is not equal to "Ned stark"
	filter5 := bson.M{"name": bson.M{"$ne": "Ned stark"}}// ne selector is used to retrieve data which are not equal to and it prints all the keyvalue pairs other than the nedstark
	cursor5, err := collection.Find(context.Background(), filter5)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor5.Close(context.Background())


	for cursor1.Next(context.Background()) {
		var doc bson.M
		if err := cursor1.Decode(&doc); err != nil {
			log.Fatal(err)
		}
		fmt.Println(doc)
	}

	for cursor2.Next(context.Background()) {
		var doc bson.M
		if err := cursor2.Decode(&doc); err != nil {
			log.Fatal(err)
		}
		fmt.Println(doc)
	}

	for cursor3.Next(context.Background()) {
		var doc bson.M
		if err := cursor3.Decode(&doc); err != nil {
			log.Fatal(err)
		}
		fmt.Println(doc)
	}

	for cursor4.Next(context.Background()) {
		var doc bson.M
		if err := cursor4.Decode(&doc); err != nil {
			log.Fatal(err)
		}
		fmt.Println(doc)
	}

	for cursor5.Next(context.Background()) {
		var doc bson.M
		if err := cursor5.Decode(&doc); err != nil {
			log.Fatal(err)
		}
		fmt.Println(doc)
	}
}
