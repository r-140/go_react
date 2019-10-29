package dbclient

import (
	"model"
)

// IDbClient ...
type IDbClient interface {
	OpenDbClient()
	QueryNews(newsID string) (model.News, error)
	QueryAllNews() ([]model.News, error)
	CreateNews(model.News) (string, error)
	Check() bool
	CreateComment(newsID string, comment model.Comment) (string, error)
}

var DBClient IDbClient
