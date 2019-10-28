package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// News model
type News struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Title    string             `json:"title" bson:"title"`
	Teaser   string             `json:"teaser" bson:"teaser"`
	Body     string             `json:"body" json:"body"`
	Comments []Comment          `json:"comments" json:"comments"`
}

// Comment ...
type Comment struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Body     string             `json:"body" bson:"body"`
}
