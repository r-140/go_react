package dbclient

import (
	"model"
)

// IDbClient ...
type IDbClient interface {
	OpenDbClient()
	QueryNews(newsID string) (model.News, error)
	QueryAllNews() ([]model.News, error)
	Seed()
	Check() bool
}
