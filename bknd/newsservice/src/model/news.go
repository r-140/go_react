package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// News model
type News struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
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

// User represents user collection in db
type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// ComparePwd compares pwds
func (user *User) ComparePwd(password string) bool {
	return user.Password == password
}

// func (a *News) String() string {
// 	return a.Id + " " + a.Title + " " + a.Teaser
// }
