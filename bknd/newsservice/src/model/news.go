package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// News model
type News struct {
	Id     primitive.ObjectID `json:"_id" bson:"_id"`
	Title  string             `json:"title" bson:"title"`
	Teaser string             `json:"teaser" bson:"teaser"`
	Body   string             `json:"body" json:"body"`
}

// func (a *News) String() string {
// 	return a.Id + " " + a.Title + " " + a.Teaser
// }
