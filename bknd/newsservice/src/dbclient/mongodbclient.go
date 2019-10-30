package dbclient

import (
	"context"

	"log"
	"model"
	"os"

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
	mongoUrl := os.Getenv("MONGODB_URL")

	clientOptions := options.Client().ApplyURI(mongoUrl)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	mc.client = client
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

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
		// elem.ID =
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	log.Printf("Found multiple documents (array of pointers): %+v\n", results)

	return results, err
}

// QueryNews ...
func (mc *MongoClient) QueryNews(newsID string) (model.News, error) {

	log.Println("newsId ", newsID)

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

	if result.Comments == nil {
		result.Comments = make([]model.Comment, 0)
	}

	log.Printf("Found a single document: %+v\n", result)

	return result, err
}

// CreateNews ...
func (mc *MongoClient) CreateNews(news model.News) (string, error) {

	news.ID = primitive.NewObjectID()

	log.Println("news ", news)

	// TODO: add validation structure
	result, err := mc.addNews(&news)
	if err != nil {
		panic(err)
	}

	log.Printf("Created a single document: %+v\n", result)

	return result, err
}

// CreateComment creates comments to database
func (mc *MongoClient) CreateComment(newsID string, comment model.Comment) (string, error) {
	log.Println("CreateComment(): comment ", comment)

	news, err := mc.QueryNews(newsID)

	log.Println("CreateComment(): found news ", news)

	if err != nil {
		panic(err)
	}

	comment.ID = primitive.NewObjectID()
	news.Comments = append(news.Comments, comment)

	log.Println("CreateComment(): fcomments ", news.Comments)

	collection := mc.client.Database("newsDb").Collection("news")

	_id, err := primitive.ObjectIDFromHex(newsID)
	if err != nil {
		panic("wrong _id format")
	}

	filter := bson.D{{"_id", _id}}

	update := bson.D{{"$set",
		bson.D{
			{"comments", news.Comments},
		},
	}}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return "comment has been created", err
}

func (mc *MongoClient) addNews(news *model.News) (string, error) {
	collection := mc.client.Database("newsDb").Collection("news")

	insertResult, err := collection.InsertOne(context.TODO(), news)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted a Single Document: ", insertResult.InsertedID.(primitive.ObjectID).Hex())

	return insertResult.InsertedID.(primitive.ObjectID).Hex(), err

}
