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
	Seed()
	Check() bool
}
