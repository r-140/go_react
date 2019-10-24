package dbclient

import (
	"context"
	"fmt"
	"log"
	"model"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoClient  Real implementation
type MongoClient struct {
	client *mongo.Client
}

// OpenDbClient ...
func (mc *MongoClient) OpenDbClient() {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	mc.client = client
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

}

// Check Naive healthcheck, just makes sure the DB connection has been initialized.
func (mc *MongoClient) Check() bool {

	return mc.client != nil

}

//QueryAllNews ...
func (mc *MongoClient) QueryAllNews() ([]model.News, error) {

	collection := mc.client.Database("newsDb").Collection("news")

	results := make([]model.News, 0)

	findOptions := options.Find()

	cur, err := collection.Find(context.TODO(), findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem model.News
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	return results, err
}

// QueryNews ...
func (mc *MongoClient) QueryNews(newsID string) (model.News, error) {

	fmt.Println("newsId ", newsID)

	_id, err := primitive.ObjectIDFromHex(newsID)
	if err != nil {
		panic("wrong _id format")
	}

	filter := bson.D{{"_id", _id}}
	var result model.News

	collection := mc.client.Database("newsDb").Collection("news")

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	return result, err
}

// CreateNews ...
func (mc *MongoClient) CreateNews(news model.News) (string, error) {

	news.Id = primitive.NewObjectID()

	fmt.Println("news ", news)

	// TODO: add validation structure
	result, err := mc.addNews(&news)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created a single document: %+v\n", result)

	return result, err
}

// Seed Start seeding news
func (mc *MongoClient) Seed() {
	mc.seedNews()
}

func (mc *MongoClient) addNews(news *model.News) (string, error) {
	collection := mc.client.Database("newsDb").Collection("news")

	insertResult, err := collection.InsertOne(context.TODO(), news)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID.(primitive.ObjectID).Hex())

	return insertResult.InsertedID.(primitive.ObjectID).Hex(), err

}

// Seed (n) make-believe account objects into the NewsBucket bucket.
func (mc *MongoClient) seedNews() {

	total := 10
	for i := 0; i < total; i++ {

		key := primitive.NewObjectID()

		// Create an instance of our News struct
		news := model.News{
			Id:     key,
			Title:  "Nachrichte_" + strconv.Itoa(i),
			Teaser: "Teaser_" + strconv.Itoa(i),
			Body:   "This is body of nachrichte_" + strconv.Itoa(i),
		}

		result, err := mc.addNews(&news)
		if err != nil {
			panic(err)
		}

		fmt.Println("created sib=ngle doc =", result)
	}
	fmt.Printf("Seeded %v fake news...\n", total)
}
