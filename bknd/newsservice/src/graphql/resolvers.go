package graphql

import (
	"dbclient"
	"fmt"
	"model"

	"github.com/graphql-go/graphql"
)

// var DBClient dbclient.IDbClient

// GraphQLResolvers provides resolvers methods for  graphql schema
type GraphQLResolvers interface {
	NewsResolverFunc(p graphql.ResolveParams) (interface{}, error)
	AllNewsResolverFunc(p graphql.ResolveParams) (interface{}, error)
	CreateNewsResolverFunc(p graphql.ResolveParams) (interface{}, error)
	AddCommentToNewsResolverFunc(p graphql.ResolveParams) (interface{}, error)
}

// LiveGraphQLResolvers implementations
type LiveGraphQLResolvers struct {
}

// NewsResolverFunc grapgql resolver for getNews query
func (gqlres *LiveGraphQLResolvers) NewsResolverFunc(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println("opening NewsResolverFunc() ")
	newsID, _ := p.Args["id"].(string)

	fmt.Println(" NewsResolverFunc() ID from argument ", newsID)
	// if newsID == nil {
	// 	panic("erroor parsing news id as an argument")
	// }

	news, err := fetchNews(newsID)
	if err != nil {
		return nil, err
	}
	return news, nil
}

// AllNewsResolverFunc graphql resolver for AllNews query
func (gqlres *LiveGraphQLResolvers) AllNewsResolverFunc(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println("opening AllNewsResolverFunc() ")

	news, err := dbclient.DBClient.QueryAllNews()

	fmt.Println("leaving  AllNewsResolverFunc() found ", news)
	return news, err
}

// CreateNewsResolverFunc graphql resolver for Create news mutation query
func (gqlres *LiveGraphQLResolvers) CreateNewsResolverFunc(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println("opening CreateNewsResolverFunc() ")

	// var news model.News

	news := model.News{
		Title:  p.Args["title"].(string),
		Teaser: p.Args["teaser"].(string),
		Body:   p.Args["body"].(string),
	}

	fmt.Println("CreateNewsResolverFunc: news from body ", news)

	result, err := dbclient.DBClient.CreateNews(news)

	fmt.Println("leaving  CreateNewsResolverFunc() found ", result)
	return result, err
}

// AddCommentToNewsResolverFunc add comment mutation resolver
func (gqlres *LiveGraphQLResolvers) AddCommentToNewsResolverFunc(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println("opening AddCommentToNewsResolverFunc() ")
	newsID, _ := p.Args["newsID"].(string)

	username, _ := p.Args["username"].(string)

	body, _ := p.Args["body"].(string)

	comment := model.Comment{
		Username: username,
		Body:     body,
	}

	result, err := dbclient.DBClient.CreateComment(newsID, comment)

	fmt.Println("leaving AddCommentToNewsResolverFunc() result ", result)
	return result, err
}

func fetchNews(newsID string) (model.News, error) {

	fmt.Println("fetchNews newsID ", newsID)

	news, err := dbclient.DBClient.QueryNews(newsID)

	fmt.Println("fetchNews result ", news)

	return news, err
}

// TestGraphQLResolvers test implementations
// type TestGraphQLResolvers struct {
// }

// func (gqlres *TestGraphQLResolvers) NewsResolverFunc(p graphql.ResolveParams) (interface{}, error) {
// 	logrus.Infof("ENTER - resolve function for Account with params %v", p.Args)
// 	id, _ := p.Args["id"].(string)
// 	for _, account := range accounts {
// 		if account.ID == id {
// 			return account, nil
// 		}
// 	}
// 	return nil, fmt.Errorf("No account found matching ID %v", id)
// }

// func (gqlres *TestGraphQLResolvers) AllNewsResolverFunc(p graphql.ResolveParams) (interface{}, error) {
// 	logrus.Infof("ENTER - resolve function for AllAccounts with params %v", p.Args)
// 	return accounts, nil
// }
